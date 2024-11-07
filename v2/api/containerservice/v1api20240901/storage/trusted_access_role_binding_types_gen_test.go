// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_TrustedAccessRoleBinding_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustedAccessRoleBinding via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustedAccessRoleBinding, TrustedAccessRoleBindingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustedAccessRoleBinding runs a test to see if a specific instance of TrustedAccessRoleBinding round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustedAccessRoleBinding(subject TrustedAccessRoleBinding) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustedAccessRoleBinding
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

// Generator of TrustedAccessRoleBinding instances for property testing - lazily instantiated by
// TrustedAccessRoleBindingGenerator()
var trustedAccessRoleBindingGenerator gopter.Gen

// TrustedAccessRoleBindingGenerator returns a generator of TrustedAccessRoleBinding instances for property testing.
func TrustedAccessRoleBindingGenerator() gopter.Gen {
	if trustedAccessRoleBindingGenerator != nil {
		return trustedAccessRoleBindingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding(generators)
	trustedAccessRoleBindingGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBinding{}), generators)

	return trustedAccessRoleBindingGenerator
}

// AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding(gens map[string]gopter.Gen) {
	gens["Spec"] = TrustedAccessRoleBinding_SpecGenerator()
	gens["Status"] = TrustedAccessRoleBinding_STATUSGenerator()
}

func Test_TrustedAccessRoleBinding_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustedAccessRoleBinding_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustedAccessRoleBinding_STATUS, TrustedAccessRoleBinding_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustedAccessRoleBinding_STATUS runs a test to see if a specific instance of TrustedAccessRoleBinding_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustedAccessRoleBinding_STATUS(subject TrustedAccessRoleBinding_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustedAccessRoleBinding_STATUS
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

// Generator of TrustedAccessRoleBinding_STATUS instances for property testing - lazily instantiated by
// TrustedAccessRoleBinding_STATUSGenerator()
var trustedAccessRoleBinding_STATUSGenerator gopter.Gen

// TrustedAccessRoleBinding_STATUSGenerator returns a generator of TrustedAccessRoleBinding_STATUS instances for property testing.
// We first initialize trustedAccessRoleBinding_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TrustedAccessRoleBinding_STATUSGenerator() gopter.Gen {
	if trustedAccessRoleBinding_STATUSGenerator != nil {
		return trustedAccessRoleBinding_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_STATUS(generators)
	trustedAccessRoleBinding_STATUSGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBinding_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_STATUS(generators)
	AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding_STATUS(generators)
	trustedAccessRoleBinding_STATUSGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBinding_STATUS{}), generators)

	return trustedAccessRoleBinding_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Roles"] = gen.SliceOf(gen.AlphaString())
	gens["SourceResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_TrustedAccessRoleBinding_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustedAccessRoleBinding_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustedAccessRoleBinding_Spec, TrustedAccessRoleBinding_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustedAccessRoleBinding_Spec runs a test to see if a specific instance of TrustedAccessRoleBinding_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustedAccessRoleBinding_Spec(subject TrustedAccessRoleBinding_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustedAccessRoleBinding_Spec
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

// Generator of TrustedAccessRoleBinding_Spec instances for property testing - lazily instantiated by
// TrustedAccessRoleBinding_SpecGenerator()
var trustedAccessRoleBinding_SpecGenerator gopter.Gen

// TrustedAccessRoleBinding_SpecGenerator returns a generator of TrustedAccessRoleBinding_Spec instances for property testing.
func TrustedAccessRoleBinding_SpecGenerator() gopter.Gen {
	if trustedAccessRoleBinding_SpecGenerator != nil {
		return trustedAccessRoleBinding_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_Spec(generators)
	trustedAccessRoleBinding_SpecGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBinding_Spec{}), generators)

	return trustedAccessRoleBinding_SpecGenerator
}

// AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Roles"] = gen.SliceOf(gen.AlphaString())
}
