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

func Test_ApplicationSecurityGroup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroup, ApplicationSecurityGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroup runs a test to see if a specific instance of ApplicationSecurityGroup round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroup(subject ApplicationSecurityGroup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup
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

// Generator of ApplicationSecurityGroup instances for property testing - lazily instantiated by
// ApplicationSecurityGroupGenerator()
var applicationSecurityGroupGenerator gopter.Gen

// ApplicationSecurityGroupGenerator returns a generator of ApplicationSecurityGroup instances for property testing.
func ApplicationSecurityGroupGenerator() gopter.Gen {
	if applicationSecurityGroupGenerator != nil {
		return applicationSecurityGroupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForApplicationSecurityGroup(generators)
	applicationSecurityGroupGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup{}), generators)

	return applicationSecurityGroupGenerator
}

// AddRelatedPropertyGeneratorsForApplicationSecurityGroup is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApplicationSecurityGroup(gens map[string]gopter.Gen) {
	gens["Spec"] = ApplicationSecurityGroup_SpecGenerator()
	gens["Status"] = ApplicationSecurityGroup_STATUSGenerator()
}

func Test_ApplicationSecurityGroup_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroup_STATUS, ApplicationSecurityGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroup_STATUS runs a test to see if a specific instance of ApplicationSecurityGroup_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroup_STATUS(subject ApplicationSecurityGroup_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_STATUS
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

// Generator of ApplicationSecurityGroup_STATUS instances for property testing - lazily instantiated by
// ApplicationSecurityGroup_STATUSGenerator()
var applicationSecurityGroup_STATUSGenerator gopter.Gen

// ApplicationSecurityGroup_STATUSGenerator returns a generator of ApplicationSecurityGroup_STATUS instances for property testing.
func ApplicationSecurityGroup_STATUSGenerator() gopter.Gen {
	if applicationSecurityGroup_STATUSGenerator != nil {
		return applicationSecurityGroup_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS(generators)
	applicationSecurityGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_STATUS{}), generators)

	return applicationSecurityGroup_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ApplicationSecurityGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroup_Spec, ApplicationSecurityGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroup_Spec runs a test to see if a specific instance of ApplicationSecurityGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroup_Spec(subject ApplicationSecurityGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_Spec
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

// Generator of ApplicationSecurityGroup_Spec instances for property testing - lazily instantiated by
// ApplicationSecurityGroup_SpecGenerator()
var applicationSecurityGroup_SpecGenerator gopter.Gen

// ApplicationSecurityGroup_SpecGenerator returns a generator of ApplicationSecurityGroup_Spec instances for property testing.
func ApplicationSecurityGroup_SpecGenerator() gopter.Gen {
	if applicationSecurityGroup_SpecGenerator != nil {
		return applicationSecurityGroup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroup_Spec(generators)
	applicationSecurityGroup_SpecGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_Spec{}), generators)

	return applicationSecurityGroup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroup_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}
