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

func Test_RouteTablesRoute_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTablesRoute via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTablesRoute, RouteTablesRouteGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTablesRoute runs a test to see if a specific instance of RouteTablesRoute round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTablesRoute(subject RouteTablesRoute) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTablesRoute
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

// Generator of RouteTablesRoute instances for property testing - lazily instantiated by RouteTablesRouteGenerator()
var routeTablesRouteGenerator gopter.Gen

// RouteTablesRouteGenerator returns a generator of RouteTablesRoute instances for property testing.
func RouteTablesRouteGenerator() gopter.Gen {
	if routeTablesRouteGenerator != nil {
		return routeTablesRouteGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRouteTablesRoute(generators)
	routeTablesRouteGenerator = gen.Struct(reflect.TypeOf(RouteTablesRoute{}), generators)

	return routeTablesRouteGenerator
}

// AddRelatedPropertyGeneratorsForRouteTablesRoute is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTablesRoute(gens map[string]gopter.Gen) {
	gens["Spec"] = RouteTablesRoute_SpecGenerator()
	gens["Status"] = RouteTablesRoute_STATUSGenerator()
}

func Test_RouteTablesRouteOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTablesRouteOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTablesRouteOperatorSpec, RouteTablesRouteOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTablesRouteOperatorSpec runs a test to see if a specific instance of RouteTablesRouteOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTablesRouteOperatorSpec(subject RouteTablesRouteOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTablesRouteOperatorSpec
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

// Generator of RouteTablesRouteOperatorSpec instances for property testing - lazily instantiated by
// RouteTablesRouteOperatorSpecGenerator()
var routeTablesRouteOperatorSpecGenerator gopter.Gen

// RouteTablesRouteOperatorSpecGenerator returns a generator of RouteTablesRouteOperatorSpec instances for property testing.
func RouteTablesRouteOperatorSpecGenerator() gopter.Gen {
	if routeTablesRouteOperatorSpecGenerator != nil {
		return routeTablesRouteOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	routeTablesRouteOperatorSpecGenerator = gen.Struct(reflect.TypeOf(RouteTablesRouteOperatorSpec{}), generators)

	return routeTablesRouteOperatorSpecGenerator
}

func Test_RouteTablesRoute_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTablesRoute_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTablesRoute_STATUS, RouteTablesRoute_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTablesRoute_STATUS runs a test to see if a specific instance of RouteTablesRoute_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTablesRoute_STATUS(subject RouteTablesRoute_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTablesRoute_STATUS
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

// Generator of RouteTablesRoute_STATUS instances for property testing - lazily instantiated by
// RouteTablesRoute_STATUSGenerator()
var routeTablesRoute_STATUSGenerator gopter.Gen

// RouteTablesRoute_STATUSGenerator returns a generator of RouteTablesRoute_STATUS instances for property testing.
func RouteTablesRoute_STATUSGenerator() gopter.Gen {
	if routeTablesRoute_STATUSGenerator != nil {
		return routeTablesRoute_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTablesRoute_STATUS(generators)
	routeTablesRoute_STATUSGenerator = gen.Struct(reflect.TypeOf(RouteTablesRoute_STATUS{}), generators)

	return routeTablesRoute_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRouteTablesRoute_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTablesRoute_STATUS(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["HasBgpOverride"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["NextHopIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["NextHopType"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_RouteTablesRoute_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTablesRoute_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTablesRoute_Spec, RouteTablesRoute_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTablesRoute_Spec runs a test to see if a specific instance of RouteTablesRoute_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTablesRoute_Spec(subject RouteTablesRoute_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTablesRoute_Spec
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

// Generator of RouteTablesRoute_Spec instances for property testing - lazily instantiated by
// RouteTablesRoute_SpecGenerator()
var routeTablesRoute_SpecGenerator gopter.Gen

// RouteTablesRoute_SpecGenerator returns a generator of RouteTablesRoute_Spec instances for property testing.
// We first initialize routeTablesRoute_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RouteTablesRoute_SpecGenerator() gopter.Gen {
	if routeTablesRoute_SpecGenerator != nil {
		return routeTablesRoute_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTablesRoute_Spec(generators)
	routeTablesRoute_SpecGenerator = gen.Struct(reflect.TypeOf(RouteTablesRoute_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTablesRoute_Spec(generators)
	AddRelatedPropertyGeneratorsForRouteTablesRoute_Spec(generators)
	routeTablesRoute_SpecGenerator = gen.Struct(reflect.TypeOf(RouteTablesRoute_Spec{}), generators)

	return routeTablesRoute_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRouteTablesRoute_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTablesRoute_Spec(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["NextHopIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["NextHopType"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRouteTablesRoute_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTablesRoute_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(RouteTablesRouteOperatorSpecGenerator())
}
