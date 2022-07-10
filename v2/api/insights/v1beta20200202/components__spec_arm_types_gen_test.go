// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200202

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

func Test_Components_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Components_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComponentsSpecARM, ComponentsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComponentsSpecARM runs a test to see if a specific instance of Components_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForComponentsSpecARM(subject Components_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Components_SpecARM
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

// Generator of Components_SpecARM instances for property testing - lazily instantiated by ComponentsSpecARMGenerator()
var componentsSpecARMGenerator gopter.Gen

// ComponentsSpecARMGenerator returns a generator of Components_SpecARM instances for property testing.
// We first initialize componentsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ComponentsSpecARMGenerator() gopter.Gen {
	if componentsSpecARMGenerator != nil {
		return componentsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComponentsSpecARM(generators)
	componentsSpecARMGenerator = gen.Struct(reflect.TypeOf(Components_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComponentsSpecARM(generators)
	AddRelatedPropertyGeneratorsForComponentsSpecARM(generators)
	componentsSpecARMGenerator = gen.Struct(reflect.TypeOf(Components_SpecARM{}), generators)

	return componentsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForComponentsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForComponentsSpecARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForComponentsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForComponentsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ApplicationInsightsComponentPropertiesARMGenerator())
}

func Test_ApplicationInsightsComponentPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationInsightsComponentPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationInsightsComponentPropertiesARM, ApplicationInsightsComponentPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationInsightsComponentPropertiesARM runs a test to see if a specific instance of ApplicationInsightsComponentPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationInsightsComponentPropertiesARM(subject ApplicationInsightsComponentPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationInsightsComponentPropertiesARM
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

// Generator of ApplicationInsightsComponentPropertiesARM instances for property testing - lazily instantiated by
// ApplicationInsightsComponentPropertiesARMGenerator()
var applicationInsightsComponentPropertiesARMGenerator gopter.Gen

// ApplicationInsightsComponentPropertiesARMGenerator returns a generator of ApplicationInsightsComponentPropertiesARM instances for property testing.
func ApplicationInsightsComponentPropertiesARMGenerator() gopter.Gen {
	if applicationInsightsComponentPropertiesARMGenerator != nil {
		return applicationInsightsComponentPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationInsightsComponentPropertiesARM(generators)
	applicationInsightsComponentPropertiesARMGenerator = gen.Struct(reflect.TypeOf(ApplicationInsightsComponentPropertiesARM{}), generators)

	return applicationInsightsComponentPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForApplicationInsightsComponentPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationInsightsComponentPropertiesARM(gens map[string]gopter.Gen) {
	gens["ApplicationType"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesApplicationTypeOther, ApplicationInsightsComponentPropertiesApplicationTypeWeb))
	gens["DisableIpMasking"] = gen.PtrOf(gen.Bool())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["FlowType"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesFlowTypeBluefield))
	gens["ForceCustomerStorageForProfiler"] = gen.PtrOf(gen.Bool())
	gens["HockeyAppId"] = gen.PtrOf(gen.AlphaString())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
	gens["IngestionMode"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesIngestionModeApplicationInsights, ApplicationInsightsComponentPropertiesIngestionModeApplicationInsightsWithDiagnosticSettings, ApplicationInsightsComponentPropertiesIngestionModeLogAnalytics))
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestionDisabled, ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestionEnabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesPublicNetworkAccessForQueryDisabled, ApplicationInsightsComponentPropertiesPublicNetworkAccessForQueryEnabled))
	gens["RequestSource"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesRequestSourceRest))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
	gens["SamplingPercentage"] = gen.PtrOf(gen.Float64())
	gens["WorkspaceResourceId"] = gen.PtrOf(gen.AlphaString())
}
