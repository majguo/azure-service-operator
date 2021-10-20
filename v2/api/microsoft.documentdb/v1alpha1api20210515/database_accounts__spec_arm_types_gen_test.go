// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_ManagedServiceIdentityARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedServiceIdentityARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedServiceIdentityARM, ManagedServiceIdentityARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedServiceIdentityARM runs a test to see if a specific instance of ManagedServiceIdentityARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedServiceIdentityARM(subject ManagedServiceIdentityARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedServiceIdentityARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ManagedServiceIdentityARM instances for property testing - lazily instantiated by
//ManagedServiceIdentityARMGenerator()
var managedServiceIdentityARMGenerator gopter.Gen

// ManagedServiceIdentityARMGenerator returns a generator of ManagedServiceIdentityARM instances for property testing.
func ManagedServiceIdentityARMGenerator() gopter.Gen {
	if managedServiceIdentityARMGenerator != nil {
		return managedServiceIdentityARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedServiceIdentityARM(generators)
	managedServiceIdentityARMGenerator = gen.Struct(reflect.TypeOf(ManagedServiceIdentityARM{}), generators)

	return managedServiceIdentityARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedServiceIdentityARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedServiceIdentityARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ManagedServiceIdentityTypeNone, ManagedServiceIdentityTypeSystemAssigned, ManagedServiceIdentityTypeSystemAssignedUserAssigned, ManagedServiceIdentityTypeUserAssigned))
}

func Test_AnalyticalStorageConfigurationARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AnalyticalStorageConfigurationARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAnalyticalStorageConfigurationARM, AnalyticalStorageConfigurationARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAnalyticalStorageConfigurationARM runs a test to see if a specific instance of AnalyticalStorageConfigurationARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAnalyticalStorageConfigurationARM(subject AnalyticalStorageConfigurationARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AnalyticalStorageConfigurationARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AnalyticalStorageConfigurationARM instances for property testing - lazily instantiated by
//AnalyticalStorageConfigurationARMGenerator()
var analyticalStorageConfigurationARMGenerator gopter.Gen

// AnalyticalStorageConfigurationARMGenerator returns a generator of AnalyticalStorageConfigurationARM instances for property testing.
func AnalyticalStorageConfigurationARMGenerator() gopter.Gen {
	if analyticalStorageConfigurationARMGenerator != nil {
		return analyticalStorageConfigurationARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAnalyticalStorageConfigurationARM(generators)
	analyticalStorageConfigurationARMGenerator = gen.Struct(reflect.TypeOf(AnalyticalStorageConfigurationARM{}), generators)

	return analyticalStorageConfigurationARMGenerator
}

// AddIndependentPropertyGeneratorsForAnalyticalStorageConfigurationARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAnalyticalStorageConfigurationARM(gens map[string]gopter.Gen) {
	gens["SchemaType"] = gen.PtrOf(gen.OneConstOf(AnalyticalStorageConfigurationSchemaTypeFullFidelity, AnalyticalStorageConfigurationSchemaTypeWellDefined))
}

func Test_ApiPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiPropertiesARM, ApiPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiPropertiesARM runs a test to see if a specific instance of ApiPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApiPropertiesARM(subject ApiPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiPropertiesARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ApiPropertiesARM instances for property testing - lazily instantiated by ApiPropertiesARMGenerator()
var apiPropertiesARMGenerator gopter.Gen

// ApiPropertiesARMGenerator returns a generator of ApiPropertiesARM instances for property testing.
func ApiPropertiesARMGenerator() gopter.Gen {
	if apiPropertiesARMGenerator != nil {
		return apiPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiPropertiesARM(generators)
	apiPropertiesARMGenerator = gen.Struct(reflect.TypeOf(ApiPropertiesARM{}), generators)

	return apiPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForApiPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiPropertiesARM(gens map[string]gopter.Gen) {
	gens["ServerVersion"] = gen.PtrOf(gen.OneConstOf(ApiPropertiesServerVersion32, ApiPropertiesServerVersion36, ApiPropertiesServerVersion40))
}

func Test_BackupPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackupPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupPolicyARM, BackupPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupPolicyARM runs a test to see if a specific instance of BackupPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupPolicyARM(subject BackupPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackupPolicyARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of BackupPolicyARM instances for property testing - lazily instantiated by BackupPolicyARMGenerator()
var backupPolicyARMGenerator gopter.Gen

// BackupPolicyARMGenerator returns a generator of BackupPolicyARM instances for property testing.
func BackupPolicyARMGenerator() gopter.Gen {
	if backupPolicyARMGenerator != nil {
		return backupPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForBackupPolicyARM(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(BackupPolicyARM{}), map[string]gopter.Gen{propName: propGen}))
	}
	backupPolicyARMGenerator = gen.OneGenOf(gens...)

	return backupPolicyARMGenerator
}

// AddRelatedPropertyGeneratorsForBackupPolicyARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackupPolicyARM(gens map[string]gopter.Gen) {
	gens["Continuous"] = ContinuousModeBackupPolicyARMGenerator().Map(func(it ContinuousModeBackupPolicyARM) *ContinuousModeBackupPolicyARM {
		return &it
	}) // generate one case for OneOf type
	gens["Periodic"] = PeriodicModeBackupPolicyARMGenerator().Map(func(it PeriodicModeBackupPolicyARM) *PeriodicModeBackupPolicyARM {
		return &it
	}) // generate one case for OneOf type
}

func Test_CapabilityARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CapabilityARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCapabilityARM, CapabilityARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCapabilityARM runs a test to see if a specific instance of CapabilityARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCapabilityARM(subject CapabilityARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CapabilityARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of CapabilityARM instances for property testing - lazily instantiated by CapabilityARMGenerator()
var capabilityARMGenerator gopter.Gen

// CapabilityARMGenerator returns a generator of CapabilityARM instances for property testing.
func CapabilityARMGenerator() gopter.Gen {
	if capabilityARMGenerator != nil {
		return capabilityARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCapabilityARM(generators)
	capabilityARMGenerator = gen.Struct(reflect.TypeOf(CapabilityARM{}), generators)

	return capabilityARMGenerator
}

// AddIndependentPropertyGeneratorsForCapabilityARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCapabilityARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_ConsistencyPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConsistencyPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConsistencyPolicyARM, ConsistencyPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConsistencyPolicyARM runs a test to see if a specific instance of ConsistencyPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConsistencyPolicyARM(subject ConsistencyPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConsistencyPolicyARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ConsistencyPolicyARM instances for property testing - lazily instantiated by
//ConsistencyPolicyARMGenerator()
var consistencyPolicyARMGenerator gopter.Gen

// ConsistencyPolicyARMGenerator returns a generator of ConsistencyPolicyARM instances for property testing.
func ConsistencyPolicyARMGenerator() gopter.Gen {
	if consistencyPolicyARMGenerator != nil {
		return consistencyPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConsistencyPolicyARM(generators)
	consistencyPolicyARMGenerator = gen.Struct(reflect.TypeOf(ConsistencyPolicyARM{}), generators)

	return consistencyPolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForConsistencyPolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConsistencyPolicyARM(gens map[string]gopter.Gen) {
	gens["DefaultConsistencyLevel"] = gen.OneConstOf(ConsistencyPolicyDefaultConsistencyLevelBoundedStaleness, ConsistencyPolicyDefaultConsistencyLevelConsistentPrefix, ConsistencyPolicyDefaultConsistencyLevelEventual, ConsistencyPolicyDefaultConsistencyLevelSession, ConsistencyPolicyDefaultConsistencyLevelStrong)
	gens["MaxIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["MaxStalenessPrefix"] = gen.PtrOf(gen.Int())
}

func Test_CorsPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsPolicyARM, CorsPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsPolicyARM runs a test to see if a specific instance of CorsPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsPolicyARM(subject CorsPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsPolicyARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of CorsPolicyARM instances for property testing - lazily instantiated by CorsPolicyARMGenerator()
var corsPolicyARMGenerator gopter.Gen

// CorsPolicyARMGenerator returns a generator of CorsPolicyARM instances for property testing.
func CorsPolicyARMGenerator() gopter.Gen {
	if corsPolicyARMGenerator != nil {
		return corsPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsPolicyARM(generators)
	corsPolicyARMGenerator = gen.Struct(reflect.TypeOf(CorsPolicyARM{}), generators)

	return corsPolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForCorsPolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsPolicyARM(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.PtrOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.PtrOf(gen.AlphaString())
	gens["AllowedOrigins"] = gen.AlphaString()
	gens["ExposedHeaders"] = gen.PtrOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.PtrOf(gen.Int())
}

func Test_IpAddressOrRangeARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpAddressOrRangeARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpAddressOrRangeARM, IpAddressOrRangeARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpAddressOrRangeARM runs a test to see if a specific instance of IpAddressOrRangeARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIpAddressOrRangeARM(subject IpAddressOrRangeARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpAddressOrRangeARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of IpAddressOrRangeARM instances for property testing - lazily instantiated by
//IpAddressOrRangeARMGenerator()
var ipAddressOrRangeARMGenerator gopter.Gen

// IpAddressOrRangeARMGenerator returns a generator of IpAddressOrRangeARM instances for property testing.
func IpAddressOrRangeARMGenerator() gopter.Gen {
	if ipAddressOrRangeARMGenerator != nil {
		return ipAddressOrRangeARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpAddressOrRangeARM(generators)
	ipAddressOrRangeARMGenerator = gen.Struct(reflect.TypeOf(IpAddressOrRangeARM{}), generators)

	return ipAddressOrRangeARMGenerator
}

// AddIndependentPropertyGeneratorsForIpAddressOrRangeARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpAddressOrRangeARM(gens map[string]gopter.Gen) {
	gens["IpAddressOrRange"] = gen.PtrOf(gen.AlphaString())
}

func Test_LocationARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LocationARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLocationARM, LocationARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLocationARM runs a test to see if a specific instance of LocationARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLocationARM(subject LocationARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LocationARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of LocationARM instances for property testing - lazily instantiated by LocationARMGenerator()
var locationARMGenerator gopter.Gen

// LocationARMGenerator returns a generator of LocationARM instances for property testing.
func LocationARMGenerator() gopter.Gen {
	if locationARMGenerator != nil {
		return locationARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLocationARM(generators)
	locationARMGenerator = gen.Struct(reflect.TypeOf(LocationARM{}), generators)

	return locationARMGenerator
}

// AddIndependentPropertyGeneratorsForLocationARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLocationARM(gens map[string]gopter.Gen) {
	gens["FailoverPriority"] = gen.PtrOf(gen.Int())
	gens["IsZoneRedundant"] = gen.PtrOf(gen.Bool())
	gens["LocationName"] = gen.PtrOf(gen.AlphaString())
}

func Test_VirtualNetworkRuleARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkRuleARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkRuleARM, VirtualNetworkRuleARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkRuleARM runs a test to see if a specific instance of VirtualNetworkRuleARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkRuleARM(subject VirtualNetworkRuleARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkRuleARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetworkRuleARM instances for property testing - lazily instantiated by
//VirtualNetworkRuleARMGenerator()
var virtualNetworkRuleARMGenerator gopter.Gen

// VirtualNetworkRuleARMGenerator returns a generator of VirtualNetworkRuleARM instances for property testing.
func VirtualNetworkRuleARMGenerator() gopter.Gen {
	if virtualNetworkRuleARMGenerator != nil {
		return virtualNetworkRuleARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkRuleARM(generators)
	virtualNetworkRuleARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkRuleARM{}), generators)

	return virtualNetworkRuleARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkRuleARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkRuleARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IgnoreMissingVNetServiceEndpoint"] = gen.PtrOf(gen.Bool())
}

func Test_ContinuousModeBackupPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContinuousModeBackupPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContinuousModeBackupPolicyARM, ContinuousModeBackupPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContinuousModeBackupPolicyARM runs a test to see if a specific instance of ContinuousModeBackupPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContinuousModeBackupPolicyARM(subject ContinuousModeBackupPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContinuousModeBackupPolicyARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ContinuousModeBackupPolicyARM instances for property testing - lazily instantiated by
//ContinuousModeBackupPolicyARMGenerator()
var continuousModeBackupPolicyARMGenerator gopter.Gen

// ContinuousModeBackupPolicyARMGenerator returns a generator of ContinuousModeBackupPolicyARM instances for property testing.
func ContinuousModeBackupPolicyARMGenerator() gopter.Gen {
	if continuousModeBackupPolicyARMGenerator != nil {
		return continuousModeBackupPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContinuousModeBackupPolicyARM(generators)
	continuousModeBackupPolicyARMGenerator = gen.Struct(reflect.TypeOf(ContinuousModeBackupPolicyARM{}), generators)

	return continuousModeBackupPolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForContinuousModeBackupPolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContinuousModeBackupPolicyARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.OneConstOf(ContinuousModeBackupPolicyTypeContinuous)
}

func Test_PeriodicModeBackupPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PeriodicModeBackupPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPeriodicModeBackupPolicyARM, PeriodicModeBackupPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPeriodicModeBackupPolicyARM runs a test to see if a specific instance of PeriodicModeBackupPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPeriodicModeBackupPolicyARM(subject PeriodicModeBackupPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PeriodicModeBackupPolicyARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of PeriodicModeBackupPolicyARM instances for property testing - lazily instantiated by
//PeriodicModeBackupPolicyARMGenerator()
var periodicModeBackupPolicyARMGenerator gopter.Gen

// PeriodicModeBackupPolicyARMGenerator returns a generator of PeriodicModeBackupPolicyARM instances for property testing.
// We first initialize periodicModeBackupPolicyARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PeriodicModeBackupPolicyARMGenerator() gopter.Gen {
	if periodicModeBackupPolicyARMGenerator != nil {
		return periodicModeBackupPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPeriodicModeBackupPolicyARM(generators)
	periodicModeBackupPolicyARMGenerator = gen.Struct(reflect.TypeOf(PeriodicModeBackupPolicyARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPeriodicModeBackupPolicyARM(generators)
	AddRelatedPropertyGeneratorsForPeriodicModeBackupPolicyARM(generators)
	periodicModeBackupPolicyARMGenerator = gen.Struct(reflect.TypeOf(PeriodicModeBackupPolicyARM{}), generators)

	return periodicModeBackupPolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForPeriodicModeBackupPolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPeriodicModeBackupPolicyARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.OneConstOf(PeriodicModeBackupPolicyTypePeriodic)
}

// AddRelatedPropertyGeneratorsForPeriodicModeBackupPolicyARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPeriodicModeBackupPolicyARM(gens map[string]gopter.Gen) {
	gens["PeriodicModeProperties"] = gen.PtrOf(PeriodicModePropertiesARMGenerator())
}

func Test_PeriodicModePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PeriodicModePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPeriodicModePropertiesARM, PeriodicModePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPeriodicModePropertiesARM runs a test to see if a specific instance of PeriodicModePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPeriodicModePropertiesARM(subject PeriodicModePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PeriodicModePropertiesARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of PeriodicModePropertiesARM instances for property testing - lazily instantiated by
//PeriodicModePropertiesARMGenerator()
var periodicModePropertiesARMGenerator gopter.Gen

// PeriodicModePropertiesARMGenerator returns a generator of PeriodicModePropertiesARM instances for property testing.
func PeriodicModePropertiesARMGenerator() gopter.Gen {
	if periodicModePropertiesARMGenerator != nil {
		return periodicModePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPeriodicModePropertiesARM(generators)
	periodicModePropertiesARMGenerator = gen.Struct(reflect.TypeOf(PeriodicModePropertiesARM{}), generators)

	return periodicModePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForPeriodicModePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPeriodicModePropertiesARM(gens map[string]gopter.Gen) {
	gens["BackupIntervalInMinutes"] = gen.PtrOf(gen.Int())
	gens["BackupRetentionIntervalInHours"] = gen.PtrOf(gen.Int())
}
