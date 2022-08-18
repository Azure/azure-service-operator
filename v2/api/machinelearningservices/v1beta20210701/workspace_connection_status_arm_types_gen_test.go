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

func Test_WorkspaceConnection_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceConnection_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceConnectionSTATUSARM, WorkspaceConnectionSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceConnectionSTATUSARM runs a test to see if a specific instance of WorkspaceConnection_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceConnectionSTATUSARM(subject WorkspaceConnection_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceConnection_STATUSARM
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

// Generator of WorkspaceConnection_STATUSARM instances for property testing - lazily instantiated by
// WorkspaceConnectionSTATUSARMGenerator()
var workspaceConnectionSTATUSARMGenerator gopter.Gen

// WorkspaceConnectionSTATUSARMGenerator returns a generator of WorkspaceConnection_STATUSARM instances for property testing.
// We first initialize workspaceConnectionSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspaceConnectionSTATUSARMGenerator() gopter.Gen {
	if workspaceConnectionSTATUSARMGenerator != nil {
		return workspaceConnectionSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceConnectionSTATUSARM(generators)
	workspaceConnectionSTATUSARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceConnection_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceConnectionSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForWorkspaceConnectionSTATUSARM(generators)
	workspaceConnectionSTATUSARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceConnection_STATUSARM{}), generators)

	return workspaceConnectionSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceConnectionSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceConnectionSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspaceConnectionSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaceConnectionSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WorkspaceConnectionPropsSTATUSARMGenerator())
}

func Test_WorkspaceConnectionProps_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceConnectionProps_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceConnectionPropsSTATUSARM, WorkspaceConnectionPropsSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceConnectionPropsSTATUSARM runs a test to see if a specific instance of WorkspaceConnectionProps_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceConnectionPropsSTATUSARM(subject WorkspaceConnectionProps_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceConnectionProps_STATUSARM
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

// Generator of WorkspaceConnectionProps_STATUSARM instances for property testing - lazily instantiated by
// WorkspaceConnectionPropsSTATUSARMGenerator()
var workspaceConnectionPropsSTATUSARMGenerator gopter.Gen

// WorkspaceConnectionPropsSTATUSARMGenerator returns a generator of WorkspaceConnectionProps_STATUSARM instances for property testing.
func WorkspaceConnectionPropsSTATUSARMGenerator() gopter.Gen {
	if workspaceConnectionPropsSTATUSARMGenerator != nil {
		return workspaceConnectionPropsSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceConnectionPropsSTATUSARM(generators)
	workspaceConnectionPropsSTATUSARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceConnectionProps_STATUSARM{}), generators)

	return workspaceConnectionPropsSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceConnectionPropsSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceConnectionPropsSTATUSARM(gens map[string]gopter.Gen) {
	gens["AuthType"] = gen.PtrOf(gen.AlphaString())
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
	gens["ValueFormat"] = gen.PtrOf(gen.OneConstOf(WorkspaceConnectionPropsSTATUSValueFormat_JSON))
}
