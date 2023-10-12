// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180301

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

func Test_MetricAlert_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlert_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlert_Spec_ARM, MetricAlert_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlert_Spec_ARM runs a test to see if a specific instance of MetricAlert_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlert_Spec_ARM(subject MetricAlert_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlert_Spec_ARM
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

// Generator of MetricAlert_Spec_ARM instances for property testing - lazily instantiated by
// MetricAlert_Spec_ARMGenerator()
var metricAlert_Spec_ARMGenerator gopter.Gen

// MetricAlert_Spec_ARMGenerator returns a generator of MetricAlert_Spec_ARM instances for property testing.
// We first initialize metricAlert_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricAlert_Spec_ARMGenerator() gopter.Gen {
	if metricAlert_Spec_ARMGenerator != nil {
		return metricAlert_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlert_Spec_ARM(generators)
	metricAlert_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(MetricAlert_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlert_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForMetricAlert_Spec_ARM(generators)
	metricAlert_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(MetricAlert_Spec_ARM{}), generators)

	return metricAlert_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMetricAlert_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricAlert_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMetricAlert_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricAlert_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MetricAlertProperties_ARMGenerator())
}

func Test_MetricAlertProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlertProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlertProperties_ARM, MetricAlertProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlertProperties_ARM runs a test to see if a specific instance of MetricAlertProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlertProperties_ARM(subject MetricAlertProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlertProperties_ARM
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

// Generator of MetricAlertProperties_ARM instances for property testing - lazily instantiated by
// MetricAlertProperties_ARMGenerator()
var metricAlertProperties_ARMGenerator gopter.Gen

// MetricAlertProperties_ARMGenerator returns a generator of MetricAlertProperties_ARM instances for property testing.
// We first initialize metricAlertProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricAlertProperties_ARMGenerator() gopter.Gen {
	if metricAlertProperties_ARMGenerator != nil {
		return metricAlertProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertProperties_ARM(generators)
	metricAlertProperties_ARMGenerator = gen.Struct(reflect.TypeOf(MetricAlertProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForMetricAlertProperties_ARM(generators)
	metricAlertProperties_ARMGenerator = gen.Struct(reflect.TypeOf(MetricAlertProperties_ARM{}), generators)

	return metricAlertProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMetricAlertProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricAlertProperties_ARM(gens map[string]gopter.Gen) {
	gens["AutoMitigate"] = gen.PtrOf(gen.Bool())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["EvaluationFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["Scopes"] = gen.SliceOf(gen.AlphaString())
	gens["Severity"] = gen.PtrOf(gen.Int())
	gens["TargetResourceRegion"] = gen.PtrOf(gen.AlphaString())
	gens["TargetResourceType"] = gen.PtrOf(gen.AlphaString())
	gens["WindowSize"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMetricAlertProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricAlertProperties_ARM(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.SliceOf(MetricAlertAction_ARMGenerator())
	gens["Criteria"] = gen.PtrOf(MetricAlertCriteria_ARMGenerator())
}

func Test_MetricAlertAction_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlertAction_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlertAction_ARM, MetricAlertAction_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlertAction_ARM runs a test to see if a specific instance of MetricAlertAction_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlertAction_ARM(subject MetricAlertAction_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlertAction_ARM
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

// Generator of MetricAlertAction_ARM instances for property testing - lazily instantiated by
// MetricAlertAction_ARMGenerator()
var metricAlertAction_ARMGenerator gopter.Gen

// MetricAlertAction_ARMGenerator returns a generator of MetricAlertAction_ARM instances for property testing.
func MetricAlertAction_ARMGenerator() gopter.Gen {
	if metricAlertAction_ARMGenerator != nil {
		return metricAlertAction_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertAction_ARM(generators)
	metricAlertAction_ARMGenerator = gen.Struct(reflect.TypeOf(MetricAlertAction_ARM{}), generators)

	return metricAlertAction_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMetricAlertAction_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricAlertAction_ARM(gens map[string]gopter.Gen) {
	gens["ActionGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["WebHookProperties"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_MetricAlertCriteria_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlertCriteria_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlertCriteria_ARM, MetricAlertCriteria_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlertCriteria_ARM runs a test to see if a specific instance of MetricAlertCriteria_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlertCriteria_ARM(subject MetricAlertCriteria_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlertCriteria_ARM
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

// Generator of MetricAlertCriteria_ARM instances for property testing - lazily instantiated by
// MetricAlertCriteria_ARMGenerator()
var metricAlertCriteria_ARMGenerator gopter.Gen

// MetricAlertCriteria_ARMGenerator returns a generator of MetricAlertCriteria_ARM instances for property testing.
func MetricAlertCriteria_ARMGenerator() gopter.Gen {
	if metricAlertCriteria_ARMGenerator != nil {
		return metricAlertCriteria_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMetricAlertCriteria_ARM(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(MetricAlertCriteria_ARM{}), map[string]gopter.Gen{propName: propGen}))
	}
	metricAlertCriteria_ARMGenerator = gen.OneGenOf(gens...)

	return metricAlertCriteria_ARMGenerator
}

// AddRelatedPropertyGeneratorsForMetricAlertCriteria_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricAlertCriteria_ARM(gens map[string]gopter.Gen) {
	gens["MicrosoftAzureMonitorMultipleResourceMultipleMetric"] = MetricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator().Map(func(it MetricAlertMultipleResourceMultipleMetricCriteria_ARM) *MetricAlertMultipleResourceMultipleMetricCriteria_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["MicrosoftAzureMonitorSingleResourceMultipleMetric"] = MetricAlertSingleResourceMultipleMetricCriteria_ARMGenerator().Map(func(it MetricAlertSingleResourceMultipleMetricCriteria_ARM) *MetricAlertSingleResourceMultipleMetricCriteria_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["MicrosoftAzureMonitorWebtestLocationAvailability"] = WebtestLocationAvailabilityCriteria_ARMGenerator().Map(func(it WebtestLocationAvailabilityCriteria_ARM) *WebtestLocationAvailabilityCriteria_ARM {
		return &it
	}) // generate one case for OneOf type
}

func Test_MetricAlertMultipleResourceMultipleMetricCriteria_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlertMultipleResourceMultipleMetricCriteria_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlertMultipleResourceMultipleMetricCriteria_ARM, MetricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlertMultipleResourceMultipleMetricCriteria_ARM runs a test to see if a specific instance of MetricAlertMultipleResourceMultipleMetricCriteria_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlertMultipleResourceMultipleMetricCriteria_ARM(subject MetricAlertMultipleResourceMultipleMetricCriteria_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlertMultipleResourceMultipleMetricCriteria_ARM
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

// Generator of MetricAlertMultipleResourceMultipleMetricCriteria_ARM instances for property testing - lazily
// instantiated by MetricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator()
var metricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator gopter.Gen

// MetricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator returns a generator of MetricAlertMultipleResourceMultipleMetricCriteria_ARM instances for property testing.
// We first initialize metricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator() gopter.Gen {
	if metricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator != nil {
		return metricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria_ARM(generators)
	metricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator = gen.Struct(reflect.TypeOf(MetricAlertMultipleResourceMultipleMetricCriteria_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria_ARM(generators)
	AddRelatedPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria_ARM(generators)
	metricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator = gen.Struct(reflect.TypeOf(MetricAlertMultipleResourceMultipleMetricCriteria_ARM{}), generators)

	return metricAlertMultipleResourceMultipleMetricCriteria_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria_ARM(gens map[string]gopter.Gen) {
	gens["OdataType"] = gen.OneConstOf(MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_MicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria)
}

// AddRelatedPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria_ARM(gens map[string]gopter.Gen) {
	gens["AllOf"] = gen.SliceOf(MultiMetricCriteria_ARMGenerator())
}

func Test_MetricAlertSingleResourceMultipleMetricCriteria_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlertSingleResourceMultipleMetricCriteria_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlertSingleResourceMultipleMetricCriteria_ARM, MetricAlertSingleResourceMultipleMetricCriteria_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlertSingleResourceMultipleMetricCriteria_ARM runs a test to see if a specific instance of MetricAlertSingleResourceMultipleMetricCriteria_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlertSingleResourceMultipleMetricCriteria_ARM(subject MetricAlertSingleResourceMultipleMetricCriteria_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlertSingleResourceMultipleMetricCriteria_ARM
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

// Generator of MetricAlertSingleResourceMultipleMetricCriteria_ARM instances for property testing - lazily instantiated
// by MetricAlertSingleResourceMultipleMetricCriteria_ARMGenerator()
var metricAlertSingleResourceMultipleMetricCriteria_ARMGenerator gopter.Gen

// MetricAlertSingleResourceMultipleMetricCriteria_ARMGenerator returns a generator of MetricAlertSingleResourceMultipleMetricCriteria_ARM instances for property testing.
// We first initialize metricAlertSingleResourceMultipleMetricCriteria_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricAlertSingleResourceMultipleMetricCriteria_ARMGenerator() gopter.Gen {
	if metricAlertSingleResourceMultipleMetricCriteria_ARMGenerator != nil {
		return metricAlertSingleResourceMultipleMetricCriteria_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria_ARM(generators)
	metricAlertSingleResourceMultipleMetricCriteria_ARMGenerator = gen.Struct(reflect.TypeOf(MetricAlertSingleResourceMultipleMetricCriteria_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria_ARM(generators)
	AddRelatedPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria_ARM(generators)
	metricAlertSingleResourceMultipleMetricCriteria_ARMGenerator = gen.Struct(reflect.TypeOf(MetricAlertSingleResourceMultipleMetricCriteria_ARM{}), generators)

	return metricAlertSingleResourceMultipleMetricCriteria_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria_ARM(gens map[string]gopter.Gen) {
	gens["OdataType"] = gen.OneConstOf(MetricAlertSingleResourceMultipleMetricCriteria_OdataType_MicrosoftAzureMonitorSingleResourceMultipleMetricCriteria)
}

// AddRelatedPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria_ARM(gens map[string]gopter.Gen) {
	gens["AllOf"] = gen.SliceOf(MetricCriteria_ARMGenerator())
}

func Test_WebtestLocationAvailabilityCriteria_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebtestLocationAvailabilityCriteria_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebtestLocationAvailabilityCriteria_ARM, WebtestLocationAvailabilityCriteria_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebtestLocationAvailabilityCriteria_ARM runs a test to see if a specific instance of WebtestLocationAvailabilityCriteria_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebtestLocationAvailabilityCriteria_ARM(subject WebtestLocationAvailabilityCriteria_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebtestLocationAvailabilityCriteria_ARM
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

// Generator of WebtestLocationAvailabilityCriteria_ARM instances for property testing - lazily instantiated by
// WebtestLocationAvailabilityCriteria_ARMGenerator()
var webtestLocationAvailabilityCriteria_ARMGenerator gopter.Gen

// WebtestLocationAvailabilityCriteria_ARMGenerator returns a generator of WebtestLocationAvailabilityCriteria_ARM instances for property testing.
func WebtestLocationAvailabilityCriteria_ARMGenerator() gopter.Gen {
	if webtestLocationAvailabilityCriteria_ARMGenerator != nil {
		return webtestLocationAvailabilityCriteria_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebtestLocationAvailabilityCriteria_ARM(generators)
	webtestLocationAvailabilityCriteria_ARMGenerator = gen.Struct(reflect.TypeOf(WebtestLocationAvailabilityCriteria_ARM{}), generators)

	return webtestLocationAvailabilityCriteria_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebtestLocationAvailabilityCriteria_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebtestLocationAvailabilityCriteria_ARM(gens map[string]gopter.Gen) {
	gens["ComponentId"] = gen.PtrOf(gen.AlphaString())
	gens["FailedLocationCount"] = gen.PtrOf(gen.Float64())
	gens["OdataType"] = gen.OneConstOf(WebtestLocationAvailabilityCriteria_OdataType_MicrosoftAzureMonitorWebtestLocationAvailabilityCriteria)
	gens["WebTestId"] = gen.PtrOf(gen.AlphaString())
}

func Test_MetricCriteria_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricCriteria_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricCriteria_ARM, MetricCriteria_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricCriteria_ARM runs a test to see if a specific instance of MetricCriteria_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricCriteria_ARM(subject MetricCriteria_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricCriteria_ARM
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

// Generator of MetricCriteria_ARM instances for property testing - lazily instantiated by MetricCriteria_ARMGenerator()
var metricCriteria_ARMGenerator gopter.Gen

// MetricCriteria_ARMGenerator returns a generator of MetricCriteria_ARM instances for property testing.
// We first initialize metricCriteria_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricCriteria_ARMGenerator() gopter.Gen {
	if metricCriteria_ARMGenerator != nil {
		return metricCriteria_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricCriteria_ARM(generators)
	metricCriteria_ARMGenerator = gen.Struct(reflect.TypeOf(MetricCriteria_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricCriteria_ARM(generators)
	AddRelatedPropertyGeneratorsForMetricCriteria_ARM(generators)
	metricCriteria_ARMGenerator = gen.Struct(reflect.TypeOf(MetricCriteria_ARM{}), generators)

	return metricCriteria_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMetricCriteria_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricCriteria_ARM(gens map[string]gopter.Gen) {
	gens["CriterionType"] = gen.OneConstOf(MetricCriteria_CriterionType_StaticThresholdCriterion)
	gens["MetricName"] = gen.PtrOf(gen.AlphaString())
	gens["MetricNamespace"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.OneConstOf(
		MetricCriteria_Operator_Equals,
		MetricCriteria_Operator_GreaterThan,
		MetricCriteria_Operator_GreaterThanOrEqual,
		MetricCriteria_Operator_LessThan,
		MetricCriteria_Operator_LessThanOrEqual))
	gens["SkipMetricValidation"] = gen.PtrOf(gen.Bool())
	gens["Threshold"] = gen.PtrOf(gen.Float64())
	gens["TimeAggregation"] = gen.PtrOf(gen.OneConstOf(
		MetricCriteria_TimeAggregation_Average,
		MetricCriteria_TimeAggregation_Count,
		MetricCriteria_TimeAggregation_Maximum,
		MetricCriteria_TimeAggregation_Minimum,
		MetricCriteria_TimeAggregation_Total))
}

// AddRelatedPropertyGeneratorsForMetricCriteria_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricCriteria_ARM(gens map[string]gopter.Gen) {
	gens["Dimensions"] = gen.SliceOf(MetricDimension_ARMGenerator())
}

func Test_MultiMetricCriteria_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MultiMetricCriteria_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMultiMetricCriteria_ARM, MultiMetricCriteria_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMultiMetricCriteria_ARM runs a test to see if a specific instance of MultiMetricCriteria_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMultiMetricCriteria_ARM(subject MultiMetricCriteria_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MultiMetricCriteria_ARM
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

// Generator of MultiMetricCriteria_ARM instances for property testing - lazily instantiated by
// MultiMetricCriteria_ARMGenerator()
var multiMetricCriteria_ARMGenerator gopter.Gen

// MultiMetricCriteria_ARMGenerator returns a generator of MultiMetricCriteria_ARM instances for property testing.
func MultiMetricCriteria_ARMGenerator() gopter.Gen {
	if multiMetricCriteria_ARMGenerator != nil {
		return multiMetricCriteria_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMultiMetricCriteria_ARM(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(MultiMetricCriteria_ARM{}), map[string]gopter.Gen{propName: propGen}))
	}
	multiMetricCriteria_ARMGenerator = gen.OneGenOf(gens...)

	return multiMetricCriteria_ARMGenerator
}

// AddRelatedPropertyGeneratorsForMultiMetricCriteria_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMultiMetricCriteria_ARM(gens map[string]gopter.Gen) {
	gens["Dynamic"] = DynamicMetricCriteria_ARMGenerator().Map(func(it DynamicMetricCriteria_ARM) *DynamicMetricCriteria_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["Static"] = MetricCriteria_ARMGenerator().Map(func(it MetricCriteria_ARM) *MetricCriteria_ARM {
		return &it
	}) // generate one case for OneOf type
}

func Test_DynamicMetricCriteria_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DynamicMetricCriteria_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDynamicMetricCriteria_ARM, DynamicMetricCriteria_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDynamicMetricCriteria_ARM runs a test to see if a specific instance of DynamicMetricCriteria_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDynamicMetricCriteria_ARM(subject DynamicMetricCriteria_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DynamicMetricCriteria_ARM
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

// Generator of DynamicMetricCriteria_ARM instances for property testing - lazily instantiated by
// DynamicMetricCriteria_ARMGenerator()
var dynamicMetricCriteria_ARMGenerator gopter.Gen

// DynamicMetricCriteria_ARMGenerator returns a generator of DynamicMetricCriteria_ARM instances for property testing.
// We first initialize dynamicMetricCriteria_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DynamicMetricCriteria_ARMGenerator() gopter.Gen {
	if dynamicMetricCriteria_ARMGenerator != nil {
		return dynamicMetricCriteria_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDynamicMetricCriteria_ARM(generators)
	dynamicMetricCriteria_ARMGenerator = gen.Struct(reflect.TypeOf(DynamicMetricCriteria_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDynamicMetricCriteria_ARM(generators)
	AddRelatedPropertyGeneratorsForDynamicMetricCriteria_ARM(generators)
	dynamicMetricCriteria_ARMGenerator = gen.Struct(reflect.TypeOf(DynamicMetricCriteria_ARM{}), generators)

	return dynamicMetricCriteria_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDynamicMetricCriteria_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDynamicMetricCriteria_ARM(gens map[string]gopter.Gen) {
	gens["AlertSensitivity"] = gen.PtrOf(gen.OneConstOf(DynamicMetricCriteria_AlertSensitivity_High, DynamicMetricCriteria_AlertSensitivity_Low, DynamicMetricCriteria_AlertSensitivity_Medium))
	gens["CriterionType"] = gen.OneConstOf(DynamicMetricCriteria_CriterionType_DynamicThresholdCriterion)
	gens["IgnoreDataBefore"] = gen.PtrOf(gen.AlphaString())
	gens["MetricName"] = gen.PtrOf(gen.AlphaString())
	gens["MetricNamespace"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.OneConstOf(DynamicMetricCriteria_Operator_GreaterOrLessThan, DynamicMetricCriteria_Operator_GreaterThan, DynamicMetricCriteria_Operator_LessThan))
	gens["SkipMetricValidation"] = gen.PtrOf(gen.Bool())
	gens["TimeAggregation"] = gen.PtrOf(gen.OneConstOf(
		DynamicMetricCriteria_TimeAggregation_Average,
		DynamicMetricCriteria_TimeAggregation_Count,
		DynamicMetricCriteria_TimeAggregation_Maximum,
		DynamicMetricCriteria_TimeAggregation_Minimum,
		DynamicMetricCriteria_TimeAggregation_Total))
}

// AddRelatedPropertyGeneratorsForDynamicMetricCriteria_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDynamicMetricCriteria_ARM(gens map[string]gopter.Gen) {
	gens["Dimensions"] = gen.SliceOf(MetricDimension_ARMGenerator())
	gens["FailingPeriods"] = gen.PtrOf(DynamicThresholdFailingPeriods_ARMGenerator())
}

func Test_MetricDimension_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricDimension_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricDimension_ARM, MetricDimension_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricDimension_ARM runs a test to see if a specific instance of MetricDimension_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricDimension_ARM(subject MetricDimension_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricDimension_ARM
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

// Generator of MetricDimension_ARM instances for property testing - lazily instantiated by
// MetricDimension_ARMGenerator()
var metricDimension_ARMGenerator gopter.Gen

// MetricDimension_ARMGenerator returns a generator of MetricDimension_ARM instances for property testing.
func MetricDimension_ARMGenerator() gopter.Gen {
	if metricDimension_ARMGenerator != nil {
		return metricDimension_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricDimension_ARM(generators)
	metricDimension_ARMGenerator = gen.Struct(reflect.TypeOf(MetricDimension_ARM{}), generators)

	return metricDimension_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMetricDimension_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricDimension_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.AlphaString())
	gens["Values"] = gen.SliceOf(gen.AlphaString())
}

func Test_DynamicThresholdFailingPeriods_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DynamicThresholdFailingPeriods_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDynamicThresholdFailingPeriods_ARM, DynamicThresholdFailingPeriods_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDynamicThresholdFailingPeriods_ARM runs a test to see if a specific instance of DynamicThresholdFailingPeriods_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDynamicThresholdFailingPeriods_ARM(subject DynamicThresholdFailingPeriods_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DynamicThresholdFailingPeriods_ARM
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

// Generator of DynamicThresholdFailingPeriods_ARM instances for property testing - lazily instantiated by
// DynamicThresholdFailingPeriods_ARMGenerator()
var dynamicThresholdFailingPeriods_ARMGenerator gopter.Gen

// DynamicThresholdFailingPeriods_ARMGenerator returns a generator of DynamicThresholdFailingPeriods_ARM instances for property testing.
func DynamicThresholdFailingPeriods_ARMGenerator() gopter.Gen {
	if dynamicThresholdFailingPeriods_ARMGenerator != nil {
		return dynamicThresholdFailingPeriods_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDynamicThresholdFailingPeriods_ARM(generators)
	dynamicThresholdFailingPeriods_ARMGenerator = gen.Struct(reflect.TypeOf(DynamicThresholdFailingPeriods_ARM{}), generators)

	return dynamicThresholdFailingPeriods_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDynamicThresholdFailingPeriods_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDynamicThresholdFailingPeriods_ARM(gens map[string]gopter.Gen) {
	gens["MinFailingPeriodsToAlert"] = gen.PtrOf(gen.Float64())
	gens["NumberOfEvaluationPeriods"] = gen.PtrOf(gen.Float64())
}
