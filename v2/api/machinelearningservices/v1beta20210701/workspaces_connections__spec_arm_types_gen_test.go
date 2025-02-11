// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701

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

func Test_WorkspacesConnections_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspacesConnections_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspacesConnectionsSpecARM, WorkspacesConnectionsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspacesConnectionsSpecARM runs a test to see if a specific instance of WorkspacesConnections_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspacesConnectionsSpecARM(subject WorkspacesConnections_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspacesConnections_SpecARM
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

// Generator of WorkspacesConnections_SpecARM instances for property testing - lazily instantiated by
// WorkspacesConnectionsSpecARMGenerator()
var workspacesConnectionsSpecARMGenerator gopter.Gen

// WorkspacesConnectionsSpecARMGenerator returns a generator of WorkspacesConnections_SpecARM instances for property testing.
// We first initialize workspacesConnectionsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspacesConnectionsSpecARMGenerator() gopter.Gen {
	if workspacesConnectionsSpecARMGenerator != nil {
		return workspacesConnectionsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacesConnectionsSpecARM(generators)
	workspacesConnectionsSpecARMGenerator = gen.Struct(reflect.TypeOf(WorkspacesConnections_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacesConnectionsSpecARM(generators)
	AddRelatedPropertyGeneratorsForWorkspacesConnectionsSpecARM(generators)
	workspacesConnectionsSpecARMGenerator = gen.Struct(reflect.TypeOf(WorkspacesConnections_SpecARM{}), generators)

	return workspacesConnectionsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspacesConnectionsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspacesConnectionsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspacesConnectionsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspacesConnectionsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WorkspaceConnectionPropsARMGenerator())
}

func Test_WorkspaceConnectionPropsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceConnectionPropsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceConnectionPropsARM, WorkspaceConnectionPropsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceConnectionPropsARM runs a test to see if a specific instance of WorkspaceConnectionPropsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceConnectionPropsARM(subject WorkspaceConnectionPropsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceConnectionPropsARM
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

// Generator of WorkspaceConnectionPropsARM instances for property testing - lazily instantiated by
// WorkspaceConnectionPropsARMGenerator()
var workspaceConnectionPropsARMGenerator gopter.Gen

// WorkspaceConnectionPropsARMGenerator returns a generator of WorkspaceConnectionPropsARM instances for property testing.
func WorkspaceConnectionPropsARMGenerator() gopter.Gen {
	if workspaceConnectionPropsARMGenerator != nil {
		return workspaceConnectionPropsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceConnectionPropsARM(generators)
	workspaceConnectionPropsARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceConnectionPropsARM{}), generators)

	return workspaceConnectionPropsARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceConnectionPropsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceConnectionPropsARM(gens map[string]gopter.Gen) {
	gens["AuthType"] = gen.PtrOf(gen.AlphaString())
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
	gens["ValueFormat"] = gen.PtrOf(gen.OneConstOf(WorkspaceConnectionPropsValueFormatJSON))
}
