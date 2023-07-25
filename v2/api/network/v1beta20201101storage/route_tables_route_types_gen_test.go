// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101storage

import (
	"encoding/json"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101storage"
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

func Test_RouteTablesRoute_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTablesRoute to hub returns original",
		prop.ForAll(RunResourceConversionTestForRouteTablesRoute, RouteTablesRouteGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRouteTablesRoute tests if a specific instance of RouteTablesRoute round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRouteTablesRoute(subject RouteTablesRoute) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20201101s.RouteTablesRoute
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RouteTablesRoute
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RouteTablesRoute_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTablesRoute to RouteTablesRoute via AssignProperties_To_RouteTablesRoute & AssignProperties_From_RouteTablesRoute returns original",
		prop.ForAll(RunPropertyAssignmentTestForRouteTablesRoute, RouteTablesRouteGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRouteTablesRoute tests if a specific instance of RouteTablesRoute can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForRouteTablesRoute(subject RouteTablesRoute) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.RouteTablesRoute
	err := copied.AssignProperties_To_RouteTablesRoute(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RouteTablesRoute
	err = actual.AssignProperties_From_RouteTablesRoute(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

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
	gens["Spec"] = RouteTables_Route_SpecGenerator()
	gens["Status"] = RouteTables_Route_STATUSGenerator()
}

func Test_RouteTables_Route_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTables_Route_Spec to RouteTables_Route_Spec via AssignProperties_To_RouteTables_Route_Spec & AssignProperties_From_RouteTables_Route_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRouteTables_Route_Spec, RouteTables_Route_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRouteTables_Route_Spec tests if a specific instance of RouteTables_Route_Spec can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForRouteTables_Route_Spec(subject RouteTables_Route_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.RouteTables_Route_Spec
	err := copied.AssignProperties_To_RouteTables_Route_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RouteTables_Route_Spec
	err = actual.AssignProperties_From_RouteTables_Route_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RouteTables_Route_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTables_Route_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTables_Route_Spec, RouteTables_Route_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTables_Route_Spec runs a test to see if a specific instance of RouteTables_Route_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTables_Route_Spec(subject RouteTables_Route_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTables_Route_Spec
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

// Generator of RouteTables_Route_Spec instances for property testing - lazily instantiated by
// RouteTables_Route_SpecGenerator()
var routeTables_Route_SpecGenerator gopter.Gen

// RouteTables_Route_SpecGenerator returns a generator of RouteTables_Route_Spec instances for property testing.
func RouteTables_Route_SpecGenerator() gopter.Gen {
	if routeTables_Route_SpecGenerator != nil {
		return routeTables_Route_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTables_Route_Spec(generators)
	routeTables_Route_SpecGenerator = gen.Struct(reflect.TypeOf(RouteTables_Route_Spec{}), generators)

	return routeTables_Route_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRouteTables_Route_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTables_Route_Spec(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["HasBgpOverride"] = gen.PtrOf(gen.Bool())
	gens["NextHopIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["NextHopType"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

func Test_RouteTables_Route_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTables_Route_STATUS to RouteTables_Route_STATUS via AssignProperties_To_RouteTables_Route_STATUS & AssignProperties_From_RouteTables_Route_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRouteTables_Route_STATUS, RouteTables_Route_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRouteTables_Route_STATUS tests if a specific instance of RouteTables_Route_STATUS can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForRouteTables_Route_STATUS(subject RouteTables_Route_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.RouteTables_Route_STATUS
	err := copied.AssignProperties_To_RouteTables_Route_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RouteTables_Route_STATUS
	err = actual.AssignProperties_From_RouteTables_Route_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RouteTables_Route_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTables_Route_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTables_Route_STATUS, RouteTables_Route_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTables_Route_STATUS runs a test to see if a specific instance of RouteTables_Route_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTables_Route_STATUS(subject RouteTables_Route_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTables_Route_STATUS
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

// Generator of RouteTables_Route_STATUS instances for property testing - lazily instantiated by
// RouteTables_Route_STATUSGenerator()
var routeTables_Route_STATUSGenerator gopter.Gen

// RouteTables_Route_STATUSGenerator returns a generator of RouteTables_Route_STATUS instances for property testing.
func RouteTables_Route_STATUSGenerator() gopter.Gen {
	if routeTables_Route_STATUSGenerator != nil {
		return routeTables_Route_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTables_Route_STATUS(generators)
	routeTables_Route_STATUSGenerator = gen.Struct(reflect.TypeOf(RouteTables_Route_STATUS{}), generators)

	return routeTables_Route_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRouteTables_Route_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTables_Route_STATUS(gens map[string]gopter.Gen) {
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
