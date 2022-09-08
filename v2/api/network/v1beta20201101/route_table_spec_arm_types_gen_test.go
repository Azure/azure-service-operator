// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

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

func Test_RouteTable_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable_SpecARM, RouteTable_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable_SpecARM runs a test to see if a specific instance of RouteTable_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable_SpecARM(subject RouteTable_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable_SpecARM
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

// Generator of RouteTable_SpecARM instances for property testing - lazily instantiated by RouteTable_SpecARMGenerator()
var routeTable_SpecARMGenerator gopter.Gen

// RouteTable_SpecARMGenerator returns a generator of RouteTable_SpecARM instances for property testing.
// We first initialize routeTable_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RouteTable_SpecARMGenerator() gopter.Gen {
	if routeTable_SpecARMGenerator != nil {
		return routeTable_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_SpecARM(generators)
	routeTable_SpecARMGenerator = gen.Struct(reflect.TypeOf(RouteTable_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_SpecARM(generators)
	AddRelatedPropertyGeneratorsForRouteTable_SpecARM(generators)
	routeTable_SpecARMGenerator = gen.Struct(reflect.TypeOf(RouteTable_SpecARM{}), generators)

	return routeTable_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForRouteTable_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTable_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRouteTable_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTable_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RouteTable_Spec_PropertiesARMGenerator())
}

func Test_RouteTable_Spec_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable_Spec_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable_Spec_PropertiesARM, RouteTable_Spec_PropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable_Spec_PropertiesARM runs a test to see if a specific instance of RouteTable_Spec_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable_Spec_PropertiesARM(subject RouteTable_Spec_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable_Spec_PropertiesARM
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

// Generator of RouteTable_Spec_PropertiesARM instances for property testing - lazily instantiated by
// RouteTable_Spec_PropertiesARMGenerator()
var routeTable_Spec_PropertiesARMGenerator gopter.Gen

// RouteTable_Spec_PropertiesARMGenerator returns a generator of RouteTable_Spec_PropertiesARM instances for property testing.
// We first initialize routeTable_Spec_PropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RouteTable_Spec_PropertiesARMGenerator() gopter.Gen {
	if routeTable_Spec_PropertiesARMGenerator != nil {
		return routeTable_Spec_PropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_Spec_PropertiesARM(generators)
	routeTable_Spec_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(RouteTable_Spec_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_Spec_PropertiesARM(generators)
	AddRelatedPropertyGeneratorsForRouteTable_Spec_PropertiesARM(generators)
	routeTable_Spec_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(RouteTable_Spec_PropertiesARM{}), generators)

	return routeTable_Spec_PropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForRouteTable_Spec_PropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTable_Spec_PropertiesARM(gens map[string]gopter.Gen) {
	gens["DisableBgpRoutePropagation"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForRouteTable_Spec_PropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTable_Spec_PropertiesARM(gens map[string]gopter.Gen) {
	gens["Routes"] = gen.SliceOf(RouteTable_Spec_Properties_RoutesARMGenerator())
}

func Test_RouteTable_Spec_Properties_RoutesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable_Spec_Properties_RoutesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable_Spec_Properties_RoutesARM, RouteTable_Spec_Properties_RoutesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable_Spec_Properties_RoutesARM runs a test to see if a specific instance of RouteTable_Spec_Properties_RoutesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable_Spec_Properties_RoutesARM(subject RouteTable_Spec_Properties_RoutesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable_Spec_Properties_RoutesARM
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

// Generator of RouteTable_Spec_Properties_RoutesARM instances for property testing - lazily instantiated by
// RouteTable_Spec_Properties_RoutesARMGenerator()
var routeTable_Spec_Properties_RoutesARMGenerator gopter.Gen

// RouteTable_Spec_Properties_RoutesARMGenerator returns a generator of RouteTable_Spec_Properties_RoutesARM instances for property testing.
// We first initialize routeTable_Spec_Properties_RoutesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RouteTable_Spec_Properties_RoutesARMGenerator() gopter.Gen {
	if routeTable_Spec_Properties_RoutesARMGenerator != nil {
		return routeTable_Spec_Properties_RoutesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_Spec_Properties_RoutesARM(generators)
	routeTable_Spec_Properties_RoutesARMGenerator = gen.Struct(reflect.TypeOf(RouteTable_Spec_Properties_RoutesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_Spec_Properties_RoutesARM(generators)
	AddRelatedPropertyGeneratorsForRouteTable_Spec_Properties_RoutesARM(generators)
	routeTable_Spec_Properties_RoutesARMGenerator = gen.Struct(reflect.TypeOf(RouteTable_Spec_Properties_RoutesARM{}), generators)

	return routeTable_Spec_Properties_RoutesARMGenerator
}

// AddIndependentPropertyGeneratorsForRouteTable_Spec_Properties_RoutesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTable_Spec_Properties_RoutesARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRouteTable_Spec_Properties_RoutesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTable_Spec_Properties_RoutesARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RoutePropertiesFormatARMGenerator())
}
