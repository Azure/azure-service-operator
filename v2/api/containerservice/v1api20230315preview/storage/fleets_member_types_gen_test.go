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

func Test_FleetsMember_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FleetsMember via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleetsMember, FleetsMemberGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleetsMember runs a test to see if a specific instance of FleetsMember round trips to JSON and back losslessly
func RunJSONSerializationTestForFleetsMember(subject FleetsMember) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FleetsMember
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

// Generator of FleetsMember instances for property testing - lazily instantiated by FleetsMemberGenerator()
var fleetsMemberGenerator gopter.Gen

// FleetsMemberGenerator returns a generator of FleetsMember instances for property testing.
func FleetsMemberGenerator() gopter.Gen {
	if fleetsMemberGenerator != nil {
		return fleetsMemberGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFleetsMember(generators)
	fleetsMemberGenerator = gen.Struct(reflect.TypeOf(FleetsMember{}), generators)

	return fleetsMemberGenerator
}

// AddRelatedPropertyGeneratorsForFleetsMember is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFleetsMember(gens map[string]gopter.Gen) {
	gens["Spec"] = FleetsMember_SpecGenerator()
	gens["Status"] = FleetsMember_STATUSGenerator()
}

func Test_FleetsMember_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FleetsMember_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleetsMember_STATUS, FleetsMember_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleetsMember_STATUS runs a test to see if a specific instance of FleetsMember_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFleetsMember_STATUS(subject FleetsMember_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FleetsMember_STATUS
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

// Generator of FleetsMember_STATUS instances for property testing - lazily instantiated by
// FleetsMember_STATUSGenerator()
var fleetsMember_STATUSGenerator gopter.Gen

// FleetsMember_STATUSGenerator returns a generator of FleetsMember_STATUS instances for property testing.
// We first initialize fleetsMember_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FleetsMember_STATUSGenerator() gopter.Gen {
	if fleetsMember_STATUSGenerator != nil {
		return fleetsMember_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleetsMember_STATUS(generators)
	fleetsMember_STATUSGenerator = gen.Struct(reflect.TypeOf(FleetsMember_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleetsMember_STATUS(generators)
	AddRelatedPropertyGeneratorsForFleetsMember_STATUS(generators)
	fleetsMember_STATUSGenerator = gen.Struct(reflect.TypeOf(FleetsMember_STATUS{}), generators)

	return fleetsMember_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFleetsMember_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleetsMember_STATUS(gens map[string]gopter.Gen) {
	gens["ClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ETag"] = gen.PtrOf(gen.AlphaString())
	gens["Group"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFleetsMember_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFleetsMember_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_FleetsMember_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FleetsMember_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleetsMember_Spec, FleetsMember_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleetsMember_Spec runs a test to see if a specific instance of FleetsMember_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFleetsMember_Spec(subject FleetsMember_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FleetsMember_Spec
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

// Generator of FleetsMember_Spec instances for property testing - lazily instantiated by FleetsMember_SpecGenerator()
var fleetsMember_SpecGenerator gopter.Gen

// FleetsMember_SpecGenerator returns a generator of FleetsMember_Spec instances for property testing.
func FleetsMember_SpecGenerator() gopter.Gen {
	if fleetsMember_SpecGenerator != nil {
		return fleetsMember_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleetsMember_Spec(generators)
	fleetsMember_SpecGenerator = gen.Struct(reflect.TypeOf(FleetsMember_Spec{}), generators)

	return fleetsMember_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFleetsMember_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleetsMember_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Group"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}
