// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_FleetMemberProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FleetMemberProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleetMemberProperties, FleetMemberPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleetMemberProperties runs a test to see if a specific instance of FleetMemberProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForFleetMemberProperties(subject FleetMemberProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FleetMemberProperties
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

// Generator of FleetMemberProperties instances for property testing - lazily instantiated by
// FleetMemberPropertiesGenerator()
var fleetMemberPropertiesGenerator gopter.Gen

// FleetMemberPropertiesGenerator returns a generator of FleetMemberProperties instances for property testing.
func FleetMemberPropertiesGenerator() gopter.Gen {
	if fleetMemberPropertiesGenerator != nil {
		return fleetMemberPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleetMemberProperties(generators)
	fleetMemberPropertiesGenerator = gen.Struct(reflect.TypeOf(FleetMemberProperties{}), generators)

	return fleetMemberPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForFleetMemberProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleetMemberProperties(gens map[string]gopter.Gen) {
	gens["ClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Group"] = gen.PtrOf(gen.AlphaString())
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
// We first initialize fleetsMember_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FleetsMember_SpecGenerator() gopter.Gen {
	if fleetsMember_SpecGenerator != nil {
		return fleetsMember_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleetsMember_Spec(generators)
	fleetsMember_SpecGenerator = gen.Struct(reflect.TypeOf(FleetsMember_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleetsMember_Spec(generators)
	AddRelatedPropertyGeneratorsForFleetsMember_Spec(generators)
	fleetsMember_SpecGenerator = gen.Struct(reflect.TypeOf(FleetsMember_Spec{}), generators)

	return fleetsMember_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFleetsMember_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleetsMember_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForFleetsMember_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFleetsMember_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FleetMemberPropertiesGenerator())
}
