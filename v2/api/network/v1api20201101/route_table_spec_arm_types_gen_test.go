// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101

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

func Test_RouteTablePropertiesFormat_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTablePropertiesFormat_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTablePropertiesFormat_ARM, RouteTablePropertiesFormat_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTablePropertiesFormat_ARM runs a test to see if a specific instance of RouteTablePropertiesFormat_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTablePropertiesFormat_ARM(subject RouteTablePropertiesFormat_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTablePropertiesFormat_ARM
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

// Generator of RouteTablePropertiesFormat_ARM instances for property testing - lazily instantiated by
// RouteTablePropertiesFormat_ARMGenerator()
var routeTablePropertiesFormat_ARMGenerator gopter.Gen

// RouteTablePropertiesFormat_ARMGenerator returns a generator of RouteTablePropertiesFormat_ARM instances for property testing.
// We first initialize routeTablePropertiesFormat_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RouteTablePropertiesFormat_ARMGenerator() gopter.Gen {
	if routeTablePropertiesFormat_ARMGenerator != nil {
		return routeTablePropertiesFormat_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTablePropertiesFormat_ARM(generators)
	routeTablePropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(RouteTablePropertiesFormat_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTablePropertiesFormat_ARM(generators)
	AddRelatedPropertyGeneratorsForRouteTablePropertiesFormat_ARM(generators)
	routeTablePropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(RouteTablePropertiesFormat_ARM{}), generators)

	return routeTablePropertiesFormat_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRouteTablePropertiesFormat_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTablePropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["DisableBgpRoutePropagation"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForRouteTablePropertiesFormat_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTablePropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["Routes"] = gen.SliceOf(Route_ARMGenerator())
}

func Test_RouteTable_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable_Spec_ARM, RouteTable_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable_Spec_ARM runs a test to see if a specific instance of RouteTable_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable_Spec_ARM(subject RouteTable_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable_Spec_ARM
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

// Generator of RouteTable_Spec_ARM instances for property testing - lazily instantiated by
// RouteTable_Spec_ARMGenerator()
var routeTable_Spec_ARMGenerator gopter.Gen

// RouteTable_Spec_ARMGenerator returns a generator of RouteTable_Spec_ARM instances for property testing.
// We first initialize routeTable_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RouteTable_Spec_ARMGenerator() gopter.Gen {
	if routeTable_Spec_ARMGenerator != nil {
		return routeTable_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_Spec_ARM(generators)
	routeTable_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(RouteTable_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForRouteTable_Spec_ARM(generators)
	routeTable_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(RouteTable_Spec_ARM{}), generators)

	return routeTable_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRouteTable_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTable_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRouteTable_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTable_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RouteTablePropertiesFormat_ARMGenerator())
}

func Test_Route_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Route_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoute_ARM, Route_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoute_ARM runs a test to see if a specific instance of Route_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoute_ARM(subject Route_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Route_ARM
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

// Generator of Route_ARM instances for property testing - lazily instantiated by Route_ARMGenerator()
var route_ARMGenerator gopter.Gen

// Route_ARMGenerator returns a generator of Route_ARM instances for property testing.
// We first initialize route_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Route_ARMGenerator() gopter.Gen {
	if route_ARMGenerator != nil {
		return route_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoute_ARM(generators)
	route_ARMGenerator = gen.Struct(reflect.TypeOf(Route_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoute_ARM(generators)
	AddRelatedPropertyGeneratorsForRoute_ARM(generators)
	route_ARMGenerator = gen.Struct(reflect.TypeOf(Route_ARM{}), generators)

	return route_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRoute_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoute_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoute_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoute_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RoutePropertiesFormat_ARMGenerator())
}
