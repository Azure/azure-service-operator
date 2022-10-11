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

func Test_Workspaces_Connection_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspaces_Connection_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaces_Connection_Spec_ARM, Workspaces_Connection_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaces_Connection_Spec_ARM runs a test to see if a specific instance of Workspaces_Connection_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaces_Connection_Spec_ARM(subject Workspaces_Connection_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspaces_Connection_Spec_ARM
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

// Generator of Workspaces_Connection_Spec_ARM instances for property testing - lazily instantiated by
// Workspaces_Connection_Spec_ARMGenerator()
var workspaces_Connection_Spec_ARMGenerator gopter.Gen

// Workspaces_Connection_Spec_ARMGenerator returns a generator of Workspaces_Connection_Spec_ARM instances for property testing.
// We first initialize workspaces_Connection_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Workspaces_Connection_Spec_ARMGenerator() gopter.Gen {
	if workspaces_Connection_Spec_ARMGenerator != nil {
		return workspaces_Connection_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaces_Connection_Spec_ARM(generators)
	workspaces_Connection_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Workspaces_Connection_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaces_Connection_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForWorkspaces_Connection_Spec_ARM(generators)
	workspaces_Connection_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Workspaces_Connection_Spec_ARM{}), generators)

	return workspaces_Connection_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaces_Connection_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaces_Connection_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspaces_Connection_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaces_Connection_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WorkspaceConnectionProps_ARMGenerator())
}

func Test_WorkspaceConnectionProps_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceConnectionProps_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceConnectionProps_ARM, WorkspaceConnectionProps_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceConnectionProps_ARM runs a test to see if a specific instance of WorkspaceConnectionProps_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceConnectionProps_ARM(subject WorkspaceConnectionProps_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceConnectionProps_ARM
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

// Generator of WorkspaceConnectionProps_ARM instances for property testing - lazily instantiated by
// WorkspaceConnectionProps_ARMGenerator()
var workspaceConnectionProps_ARMGenerator gopter.Gen

// WorkspaceConnectionProps_ARMGenerator returns a generator of WorkspaceConnectionProps_ARM instances for property testing.
func WorkspaceConnectionProps_ARMGenerator() gopter.Gen {
	if workspaceConnectionProps_ARMGenerator != nil {
		return workspaceConnectionProps_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceConnectionProps_ARM(generators)
	workspaceConnectionProps_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceConnectionProps_ARM{}), generators)

	return workspaceConnectionProps_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceConnectionProps_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceConnectionProps_ARM(gens map[string]gopter.Gen) {
	gens["AuthType"] = gen.PtrOf(gen.AlphaString())
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
	gens["ValueFormat"] = gen.PtrOf(gen.OneConstOf(WorkspaceConnectionProps_ValueFormat_JSON))
}
