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

<<<<<<<< HEAD:v2/api/machinelearningservices/v1beta20210701/workspaces_connection__spec_arm_types_gen_test.go
func Test_WorkspacesConnection_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
========
func Test_Workspaces_Connections_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
>>>>>>>> main:v2/api/machinelearningservices/v1beta20210701/workspaces_connections_spec_arm_types_gen_test.go
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<<< HEAD:v2/api/machinelearningservices/v1beta20210701/workspaces_connection__spec_arm_types_gen_test.go
		"Round trip of WorkspacesConnection_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspacesConnection_SpecARM, WorkspacesConnection_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspacesConnection_SpecARM runs a test to see if a specific instance of WorkspacesConnection_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspacesConnection_SpecARM(subject WorkspacesConnection_SpecARM) string {
========
		"Round trip of Workspaces_Connections_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaces_Connections_SpecARM, Workspaces_Connections_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaces_Connections_SpecARM runs a test to see if a specific instance of Workspaces_Connections_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaces_Connections_SpecARM(subject Workspaces_Connections_SpecARM) string {
>>>>>>>> main:v2/api/machinelearningservices/v1beta20210701/workspaces_connections_spec_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
<<<<<<<< HEAD:v2/api/machinelearningservices/v1beta20210701/workspaces_connection__spec_arm_types_gen_test.go
	var actual WorkspacesConnection_SpecARM
========
	var actual Workspaces_Connections_SpecARM
>>>>>>>> main:v2/api/machinelearningservices/v1beta20210701/workspaces_connections_spec_arm_types_gen_test.go
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

<<<<<<<< HEAD:v2/api/machinelearningservices/v1beta20210701/workspaces_connection__spec_arm_types_gen_test.go
// Generator of WorkspacesConnection_SpecARM instances for property testing - lazily instantiated by
// WorkspacesConnection_SpecARMGenerator()
var workspacesConnection_SpecARMGenerator gopter.Gen

// WorkspacesConnection_SpecARMGenerator returns a generator of WorkspacesConnection_SpecARM instances for property testing.
// We first initialize workspacesConnection_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspacesConnection_SpecARMGenerator() gopter.Gen {
	if workspacesConnection_SpecARMGenerator != nil {
		return workspacesConnection_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacesConnection_SpecARM(generators)
	workspacesConnection_SpecARMGenerator = gen.Struct(reflect.TypeOf(WorkspacesConnection_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacesConnection_SpecARM(generators)
	AddRelatedPropertyGeneratorsForWorkspacesConnection_SpecARM(generators)
	workspacesConnection_SpecARMGenerator = gen.Struct(reflect.TypeOf(WorkspacesConnection_SpecARM{}), generators)

	return workspacesConnection_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspacesConnection_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspacesConnection_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
========
// Generator of Workspaces_Connections_SpecARM instances for property testing - lazily instantiated by
// Workspaces_Connections_SpecARMGenerator()
var workspaces_Connections_SpecARMGenerator gopter.Gen

// Workspaces_Connections_SpecARMGenerator returns a generator of Workspaces_Connections_SpecARM instances for property testing.
// We first initialize workspaces_Connections_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Workspaces_Connections_SpecARMGenerator() gopter.Gen {
	if workspaces_Connections_SpecARMGenerator != nil {
		return workspaces_Connections_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaces_Connections_SpecARM(generators)
	workspaces_Connections_SpecARMGenerator = gen.Struct(reflect.TypeOf(Workspaces_Connections_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaces_Connections_SpecARM(generators)
	AddRelatedPropertyGeneratorsForWorkspaces_Connections_SpecARM(generators)
	workspaces_Connections_SpecARMGenerator = gen.Struct(reflect.TypeOf(Workspaces_Connections_SpecARM{}), generators)

	return workspaces_Connections_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaces_Connections_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaces_Connections_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
>>>>>>>> main:v2/api/machinelearningservices/v1beta20210701/workspaces_connections_spec_arm_types_gen_test.go
	gens["Name"] = gen.AlphaString()
}

<<<<<<<< HEAD:v2/api/machinelearningservices/v1beta20210701/workspaces_connection__spec_arm_types_gen_test.go
// AddRelatedPropertyGeneratorsForWorkspacesConnection_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspacesConnection_SpecARM(gens map[string]gopter.Gen) {
========
// AddRelatedPropertyGeneratorsForWorkspaces_Connections_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaces_Connections_SpecARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/machinelearningservices/v1beta20210701/workspaces_connections_spec_arm_types_gen_test.go
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
	gens["ValueFormat"] = gen.PtrOf(gen.OneConstOf(WorkspaceConnectionProps_ValueFormat_JSON))
}
