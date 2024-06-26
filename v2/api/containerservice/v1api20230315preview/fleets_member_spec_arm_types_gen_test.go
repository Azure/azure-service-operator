// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230315preview

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

func Test_FleetMemberProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FleetMemberProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleetMemberProperties_ARM, FleetMemberProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleetMemberProperties_ARM runs a test to see if a specific instance of FleetMemberProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFleetMemberProperties_ARM(subject FleetMemberProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FleetMemberProperties_ARM
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

// Generator of FleetMemberProperties_ARM instances for property testing - lazily instantiated by
// FleetMemberProperties_ARMGenerator()
var fleetMemberProperties_ARMGenerator gopter.Gen

// FleetMemberProperties_ARMGenerator returns a generator of FleetMemberProperties_ARM instances for property testing.
func FleetMemberProperties_ARMGenerator() gopter.Gen {
	if fleetMemberProperties_ARMGenerator != nil {
		return fleetMemberProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleetMemberProperties_ARM(generators)
	fleetMemberProperties_ARMGenerator = gen.Struct(reflect.TypeOf(FleetMemberProperties_ARM{}), generators)

	return fleetMemberProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFleetMemberProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleetMemberProperties_ARM(gens map[string]gopter.Gen) {
	gens["ClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Group"] = gen.PtrOf(gen.AlphaString())
}

func Test_Fleets_Member_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Fleets_Member_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleets_Member_Spec_ARM, Fleets_Member_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleets_Member_Spec_ARM runs a test to see if a specific instance of Fleets_Member_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFleets_Member_Spec_ARM(subject Fleets_Member_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Fleets_Member_Spec_ARM
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

// Generator of Fleets_Member_Spec_ARM instances for property testing - lazily instantiated by
// Fleets_Member_Spec_ARMGenerator()
var fleets_Member_Spec_ARMGenerator gopter.Gen

// Fleets_Member_Spec_ARMGenerator returns a generator of Fleets_Member_Spec_ARM instances for property testing.
// We first initialize fleets_Member_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Fleets_Member_Spec_ARMGenerator() gopter.Gen {
	if fleets_Member_Spec_ARMGenerator != nil {
		return fleets_Member_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleets_Member_Spec_ARM(generators)
	fleets_Member_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Fleets_Member_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleets_Member_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForFleets_Member_Spec_ARM(generators)
	fleets_Member_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Fleets_Member_Spec_ARM{}), generators)

	return fleets_Member_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFleets_Member_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleets_Member_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForFleets_Member_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFleets_Member_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FleetMemberProperties_ARMGenerator())
}
