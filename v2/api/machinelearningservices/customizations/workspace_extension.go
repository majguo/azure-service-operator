/*
* Copyright (c) Microsoft Corporation.
* Licensed under the MIT license.
 */

package customizations

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning"
	storage "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1beta20210701storage"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ extensions.SecretsRetriever = &WorkspaceExtension{}

func (ext *WorkspaceExtension) RetrieveSecrets(
	ctx context.Context,
	obj genruntime.ARMMetaObject,
	armClient *genericarmclient.GenericClient,
	log logr.Logger) ([]*v1.Secret, error) {

	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*storage.Workspace)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *storage.Workspace", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = typedObj

	hasSecrets := secretsSpecified(typedObj)
	if !hasSecrets {
		log.V(Debug).Info("No secrets retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	id, err := genruntime.GetAndParseResourceID(obj)
	if err != nil {
		return nil, err
	}

	var keys armmachinelearning.ListWorkspaceKeysResult
	// Only bother calling ListKeys if there are secrets to retrieve
	if hasSecrets {
		subscription := armClient.SubscriptionID()
		// Using armClient.ClientOptions() here ensures we share the same HTTP connection, so this is not opening a new
		// connection each time through
		var workspacesClient *armmachinelearning.WorkspacesClient
		workspacesClient, err = armmachinelearning.NewWorkspacesClient(subscription, armClient.Creds(), armClient.ClientOptions())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create new workspaceClient")
		}

		var resp armmachinelearning.WorkspacesClientListKeysResponse
		resp, err = workspacesClient.ListKeys(ctx, id.ResourceGroupName, obj.AzureName(), nil)
		if err != nil {
			return nil, errors.Wrapf(err, "failed listing keys")
		}
		keys = resp.ListWorkspaceKeysResult
	}

	secretSlice, err := secretsToWrite(typedObj, keys)
	if err != nil {
		return nil, err
	}

	return secretSlice, nil
}

func secretsSpecified(obj *storage.Workspace) bool {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false
	}

	operatorSecrets := obj.Spec.OperatorSpec.Secrets

	if operatorSecrets.AppInsightsInstrumentationKey != nil ||
		operatorSecrets.PrimaryNotebookAccessKey != nil ||
		operatorSecrets.SecondaryNotebookAccessKey != nil ||
		operatorSecrets.UserStorageKey != nil ||
		operatorSecrets.ContainerRegistryPassword != nil ||
		operatorSecrets.ContainerRegistryPassword2 != nil ||
		operatorSecrets.ContainerRegistryUserName != nil {
		return true
	}

	return false
}

func secretsToWrite(obj *storage.Workspace, keysResp armmachinelearning.ListWorkspaceKeysResult) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, errors.Errorf("unexpected nil operatorspec")
	}

	collector := secrets.NewSecretCollector(obj.Namespace)

	creds, crUsername := getContainerRegCreds(keysResp)

	collector.AddSecretValue(operatorSpecSecrets.ContainerRegistryPassword, creds["password"])
	collector.AddSecretValue(operatorSpecSecrets.ContainerRegistryPassword2, creds["password2"])
	collector.AddSecretValue(operatorSpecSecrets.ContainerRegistryUserName, crUsername)
	collector.AddSecretValue(operatorSpecSecrets.UserStorageKey, to.String(keysResp.UserStorageKey))
	collector.AddSecretValue(operatorSpecSecrets.AppInsightsInstrumentationKey, to.String(keysResp.AppInsightsInstrumentationKey))

	if keysResp.NotebookAccessKeys != nil {
		collector.AddSecretValue(operatorSpecSecrets.PrimaryNotebookAccessKey, to.String(keysResp.NotebookAccessKeys.PrimaryAccessKey))
		collector.AddSecretValue(operatorSpecSecrets.SecondaryNotebookAccessKey, to.String(keysResp.NotebookAccessKeys.SecondaryAccessKey))
	}
	if keysResp.ContainerRegistryCredentials != nil {
	}

	return collector.Secrets(), nil
}

func getContainerRegCreds(keysResp armmachinelearning.ListWorkspaceKeysResult) (map[string]string, string) {
	creds := make(map[string]string)

	if keysResp.ContainerRegistryCredentials == nil {
		return creds, ""
	}

	for _, password := range keysResp.ContainerRegistryCredentials.Passwords {
		if password.Name != nil && password.Value != nil {
			creds[to.String(password.Name)] = to.String(password.Value)
		}
	}
	return creds, to.String(keysResp.ContainerRegistryCredentials.Username)
}
