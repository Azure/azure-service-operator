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

func Test_Policy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Policy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicy, PolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicy runs a test to see if a specific instance of Policy round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicy(subject Policy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Policy
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

// Generator of Policy instances for property testing - lazily instantiated by PolicyGenerator()
var policyGenerator gopter.Gen

// PolicyGenerator returns a generator of Policy instances for property testing.
func PolicyGenerator() gopter.Gen {
	if policyGenerator != nil {
		return policyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPolicy(generators)
	policyGenerator = gen.Struct(reflect.TypeOf(Policy{}), generators)

	return policyGenerator
}

// AddRelatedPropertyGeneratorsForPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = Service_Policy_SpecGenerator()
	gens["Status"] = Service_Policy_STATUSGenerator()
}

func Test_Service_Policy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Policy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Policy_Spec, Service_Policy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Policy_Spec runs a test to see if a specific instance of Service_Policy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Policy_Spec(subject Service_Policy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Policy_Spec
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

// Generator of Service_Policy_Spec instances for property testing - lazily instantiated by
// Service_Policy_SpecGenerator()
var service_Policy_SpecGenerator gopter.Gen

// Service_Policy_SpecGenerator returns a generator of Service_Policy_Spec instances for property testing.
func Service_Policy_SpecGenerator() gopter.Gen {
	if service_Policy_SpecGenerator != nil {
		return service_Policy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Policy_Spec(generators)
	service_Policy_SpecGenerator = gen.Struct(reflect.TypeOf(Service_Policy_Spec{}), generators)

	return service_Policy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForService_Policy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Policy_Spec(gens map[string]gopter.Gen) {
	gens["Format"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_Service_Policy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Policy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Policy_STATUS, Service_Policy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Policy_STATUS runs a test to see if a specific instance of Service_Policy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Policy_STATUS(subject Service_Policy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Policy_STATUS
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

// Generator of Service_Policy_STATUS instances for property testing - lazily instantiated by
// Service_Policy_STATUSGenerator()
var service_Policy_STATUSGenerator gopter.Gen

// Service_Policy_STATUSGenerator returns a generator of Service_Policy_STATUS instances for property testing.
func Service_Policy_STATUSGenerator() gopter.Gen {
	if service_Policy_STATUSGenerator != nil {
		return service_Policy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Policy_STATUS(generators)
	service_Policy_STATUSGenerator = gen.Struct(reflect.TypeOf(Service_Policy_STATUS{}), generators)

	return service_Policy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForService_Policy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Policy_STATUS(gens map[string]gopter.Gen) {
	gens["Format"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}
