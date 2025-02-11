/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package main

import (
	"flag"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/benbjohnson/clock"
	"github.com/go-logr/logr"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/healthz"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	armreconciler "github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/version"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

func main() {
	exeName := os.Args[0] + " " + version.BuildVersion
	flagSet := flag.NewFlagSet(exeName, flag.ExitOnError)
	klog.InitFlags(flagSet)
	setupLog := ctrl.Log.WithName("setup")

	var metricsAddr string
	var healthAddr string
	var enableLeaderElection bool

	// default here for 'metricsAddr' is set to "0", which sets metrics to be disabled if 'metrics-addr' flag is omitted.
	flagSet.StringVar(&metricsAddr, "metrics-addr", "0", "The address the metric endpoint binds to.")
	flagSet.StringVar(&healthAddr, "health-addr", "", "The address the healthz endpoint binds to.")
	flagSet.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controllers manager. Enabling this will ensure there is only one active controllers manager.")
	flagSet.Parse(os.Args[1:]) //nolint:errcheck

	armMetrics := asometrics.NewARMClientMetrics()
	asometrics.RegisterMetrics(armMetrics)

	scheme := controllers.CreateScheme()

	ctrl.SetLogger(klogr.New())

	cfg, err := config.ReadAndValidate()
	if err != nil {
		setupLog.Error(err, "unable to get env configuration values")
		os.Exit(1)
	}

	var cacheFunc cache.NewCacheFunc
	if cfg.TargetNamespaces != nil {
		cacheFunc = cache.MultiNamespacedCacheBuilder(cfg.TargetNamespaces)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		NewCache:               cacheFunc,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "controllers-leader-election-azinfra-generated",
		Port:                   9443,
		HealthProbeBindAddress: healthAddr,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		setupLog.Error(err, "unable to get default azure credential")
		os.Exit(1)
	}

	armClient, err := genericarmclient.NewGenericClient(cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint, creds, cfg.SubscriptionID, armMetrics)
	if err != nil {
		setupLog.Error(err, "failed to get new genericArmClient")
		os.Exit(1)
	}

	var clientFactory armreconciler.ARMClientFactory = func(_ genruntime.ARMMetaObject) *genericarmclient.GenericClient {
		// always use the configured ARM client
		return armClient
	}

	log := ctrl.Log.WithName("controllers")
	log.V(Status).Info("Configuration details", "config", cfg.String())

	if cfg.OperatorMode.IncludesWatchers() {
		kubeClient := kubeclient.NewClient(mgr.GetClient())
		positiveConditions := conditions.NewPositiveConditionBuilder(clock.New())

		options := makeControllerOptions(log, cfg)
		var objs []*registration.StorageType
		objs, err = controllers.GetKnownStorageTypes(
			mgr,
			clientFactory,
			kubeClient,
			positiveConditions,
			options)
		if err != nil {
			setupLog.Error(err, "failed getting storage types and reconcilers")
			os.Exit(1)
		}

		err = controllers.RegisterAll(
			mgr,
			mgr.GetFieldIndexer(),
			kubeClient,
			positiveConditions,
			objs,
			options)
		if err != nil {
			setupLog.Error(err, "failed to register gvks")
			os.Exit(1)
		}
	}

	if cfg.OperatorMode.IncludesWebhooks() {
		if errs := controllers.RegisterWebhooks(mgr, controllers.GetKnownTypes()); errs != nil {
			setupLog.Error(err, "failed to register webhook for gvks")
			os.Exit(1)
		}
	}

	// Healthz liveness probe endpoint
	err = mgr.AddHealthzCheck("healthz", healthz.Ping)
	if err != nil {
		setupLog.Error(err, "Failed setting up health check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func makeControllerOptions(log logr.Logger, cfg config.Values) controllers.Options {
	return controllers.Options{
		Config: cfg,
		Options: controller.Options{
			MaxConcurrentReconciles: 1,
			Log:                     log,
			RateLimiter:             controllers.NewRateLimiter(1*time.Second, 1*time.Minute),
		},
	}
}
