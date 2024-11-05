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

func Test_DynamicMetricCriteria_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DynamicMetricCriteria via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDynamicMetricCriteria, DynamicMetricCriteriaGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDynamicMetricCriteria runs a test to see if a specific instance of DynamicMetricCriteria round trips to JSON and back losslessly
func RunJSONSerializationTestForDynamicMetricCriteria(subject DynamicMetricCriteria) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DynamicMetricCriteria
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

// Generator of DynamicMetricCriteria instances for property testing - lazily instantiated by
// DynamicMetricCriteriaGenerator()
var dynamicMetricCriteriaGenerator gopter.Gen

// DynamicMetricCriteriaGenerator returns a generator of DynamicMetricCriteria instances for property testing.
// We first initialize dynamicMetricCriteriaGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DynamicMetricCriteriaGenerator() gopter.Gen {
	if dynamicMetricCriteriaGenerator != nil {
		return dynamicMetricCriteriaGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDynamicMetricCriteria(generators)
	dynamicMetricCriteriaGenerator = gen.Struct(reflect.TypeOf(DynamicMetricCriteria{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDynamicMetricCriteria(generators)
	AddRelatedPropertyGeneratorsForDynamicMetricCriteria(generators)
	dynamicMetricCriteriaGenerator = gen.Struct(reflect.TypeOf(DynamicMetricCriteria{}), generators)

	return dynamicMetricCriteriaGenerator
}

// AddIndependentPropertyGeneratorsForDynamicMetricCriteria is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDynamicMetricCriteria(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForDynamicMetricCriteria is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDynamicMetricCriteria(gens map[string]gopter.Gen) {
	gens["Dimensions"] = gen.SliceOf(MetricDimensionGenerator())
	gens["FailingPeriods"] = gen.PtrOf(DynamicThresholdFailingPeriodsGenerator())
}

func Test_DynamicThresholdFailingPeriods_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DynamicThresholdFailingPeriods via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDynamicThresholdFailingPeriods, DynamicThresholdFailingPeriodsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDynamicThresholdFailingPeriods runs a test to see if a specific instance of DynamicThresholdFailingPeriods round trips to JSON and back losslessly
func RunJSONSerializationTestForDynamicThresholdFailingPeriods(subject DynamicThresholdFailingPeriods) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DynamicThresholdFailingPeriods
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

// Generator of DynamicThresholdFailingPeriods instances for property testing - lazily instantiated by
// DynamicThresholdFailingPeriodsGenerator()
var dynamicThresholdFailingPeriodsGenerator gopter.Gen

// DynamicThresholdFailingPeriodsGenerator returns a generator of DynamicThresholdFailingPeriods instances for property testing.
func DynamicThresholdFailingPeriodsGenerator() gopter.Gen {
	if dynamicThresholdFailingPeriodsGenerator != nil {
		return dynamicThresholdFailingPeriodsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDynamicThresholdFailingPeriods(generators)
	dynamicThresholdFailingPeriodsGenerator = gen.Struct(reflect.TypeOf(DynamicThresholdFailingPeriods{}), generators)

	return dynamicThresholdFailingPeriodsGenerator
}

// AddIndependentPropertyGeneratorsForDynamicThresholdFailingPeriods is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDynamicThresholdFailingPeriods(gens map[string]gopter.Gen) {
	gens["MinFailingPeriodsToAlert"] = gen.PtrOf(gen.Float64())
	gens["NumberOfEvaluationPeriods"] = gen.PtrOf(gen.Float64())
}

func Test_MetricAlertAction_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlertAction via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlertAction, MetricAlertActionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlertAction runs a test to see if a specific instance of MetricAlertAction round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlertAction(subject MetricAlertAction) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlertAction
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

// Generator of MetricAlertAction instances for property testing - lazily instantiated by MetricAlertActionGenerator()
var metricAlertActionGenerator gopter.Gen

// MetricAlertActionGenerator returns a generator of MetricAlertAction instances for property testing.
func MetricAlertActionGenerator() gopter.Gen {
	if metricAlertActionGenerator != nil {
		return metricAlertActionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertAction(generators)
	metricAlertActionGenerator = gen.Struct(reflect.TypeOf(MetricAlertAction{}), generators)

	return metricAlertActionGenerator
}

// AddIndependentPropertyGeneratorsForMetricAlertAction is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricAlertAction(gens map[string]gopter.Gen) {
	gens["ActionGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["WebHookProperties"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

func Test_MetricAlertCriteria_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlertCriteria via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlertCriteria, MetricAlertCriteriaGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlertCriteria runs a test to see if a specific instance of MetricAlertCriteria round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlertCriteria(subject MetricAlertCriteria) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlertCriteria
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

// Generator of MetricAlertCriteria instances for property testing - lazily instantiated by
// MetricAlertCriteriaGenerator()
var metricAlertCriteriaGenerator gopter.Gen

// MetricAlertCriteriaGenerator returns a generator of MetricAlertCriteria instances for property testing.
func MetricAlertCriteriaGenerator() gopter.Gen {
	if metricAlertCriteriaGenerator != nil {
		return metricAlertCriteriaGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMetricAlertCriteria(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(MetricAlertCriteria{}), map[string]gopter.Gen{propName: propGen}))
	}
	metricAlertCriteriaGenerator = gen.OneGenOf(gens...)

	return metricAlertCriteriaGenerator
}

// AddRelatedPropertyGeneratorsForMetricAlertCriteria is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricAlertCriteria(gens map[string]gopter.Gen) {
	gens["MicrosoftAzureMonitorMultipleResourceMultipleMetric"] = MetricAlertMultipleResourceMultipleMetricCriteriaGenerator().Map(func(it MetricAlertMultipleResourceMultipleMetricCriteria) *MetricAlertMultipleResourceMultipleMetricCriteria {
		return &it
	}) // generate one case for OneOf type
	gens["MicrosoftAzureMonitorSingleResourceMultipleMetric"] = MetricAlertSingleResourceMultipleMetricCriteriaGenerator().Map(func(it MetricAlertSingleResourceMultipleMetricCriteria) *MetricAlertSingleResourceMultipleMetricCriteria {
		return &it
	}) // generate one case for OneOf type
	gens["MicrosoftAzureMonitorWebtestLocationAvailability"] = WebtestLocationAvailabilityCriteriaGenerator().Map(func(it WebtestLocationAvailabilityCriteria) *WebtestLocationAvailabilityCriteria {
		return &it
	}) // generate one case for OneOf type
}

func Test_MetricAlertMultipleResourceMultipleMetricCriteria_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlertMultipleResourceMultipleMetricCriteria via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlertMultipleResourceMultipleMetricCriteria, MetricAlertMultipleResourceMultipleMetricCriteriaGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlertMultipleResourceMultipleMetricCriteria runs a test to see if a specific instance of MetricAlertMultipleResourceMultipleMetricCriteria round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlertMultipleResourceMultipleMetricCriteria(subject MetricAlertMultipleResourceMultipleMetricCriteria) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlertMultipleResourceMultipleMetricCriteria
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

// Generator of MetricAlertMultipleResourceMultipleMetricCriteria instances for property testing - lazily instantiated
// by MetricAlertMultipleResourceMultipleMetricCriteriaGenerator()
var metricAlertMultipleResourceMultipleMetricCriteriaGenerator gopter.Gen

// MetricAlertMultipleResourceMultipleMetricCriteriaGenerator returns a generator of MetricAlertMultipleResourceMultipleMetricCriteria instances for property testing.
// We first initialize metricAlertMultipleResourceMultipleMetricCriteriaGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricAlertMultipleResourceMultipleMetricCriteriaGenerator() gopter.Gen {
	if metricAlertMultipleResourceMultipleMetricCriteriaGenerator != nil {
		return metricAlertMultipleResourceMultipleMetricCriteriaGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria(generators)
	metricAlertMultipleResourceMultipleMetricCriteriaGenerator = gen.Struct(reflect.TypeOf(MetricAlertMultipleResourceMultipleMetricCriteria{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria(generators)
	AddRelatedPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria(generators)
	metricAlertMultipleResourceMultipleMetricCriteriaGenerator = gen.Struct(reflect.TypeOf(MetricAlertMultipleResourceMultipleMetricCriteria{}), generators)

	return metricAlertMultipleResourceMultipleMetricCriteriaGenerator
}

// AddIndependentPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria(gens map[string]gopter.Gen) {
	gens["OdataType"] = gen.OneConstOf(MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_MicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria)
}

// AddRelatedPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricAlertMultipleResourceMultipleMetricCriteria(gens map[string]gopter.Gen) {
	gens["AllOf"] = gen.SliceOf(MultiMetricCriteriaGenerator())
}

func Test_MetricAlertProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlertProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlertProperties, MetricAlertPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlertProperties runs a test to see if a specific instance of MetricAlertProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlertProperties(subject MetricAlertProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlertProperties
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

// Generator of MetricAlertProperties instances for property testing - lazily instantiated by
// MetricAlertPropertiesGenerator()
var metricAlertPropertiesGenerator gopter.Gen

// MetricAlertPropertiesGenerator returns a generator of MetricAlertProperties instances for property testing.
// We first initialize metricAlertPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricAlertPropertiesGenerator() gopter.Gen {
	if metricAlertPropertiesGenerator != nil {
		return metricAlertPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertProperties(generators)
	metricAlertPropertiesGenerator = gen.Struct(reflect.TypeOf(MetricAlertProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertProperties(generators)
	AddRelatedPropertyGeneratorsForMetricAlertProperties(generators)
	metricAlertPropertiesGenerator = gen.Struct(reflect.TypeOf(MetricAlertProperties{}), generators)

	return metricAlertPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForMetricAlertProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricAlertProperties(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForMetricAlertProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricAlertProperties(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.SliceOf(MetricAlertActionGenerator())
	gens["Criteria"] = gen.PtrOf(MetricAlertCriteriaGenerator())
}

func Test_MetricAlertSingleResourceMultipleMetricCriteria_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlertSingleResourceMultipleMetricCriteria via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlertSingleResourceMultipleMetricCriteria, MetricAlertSingleResourceMultipleMetricCriteriaGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlertSingleResourceMultipleMetricCriteria runs a test to see if a specific instance of MetricAlertSingleResourceMultipleMetricCriteria round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlertSingleResourceMultipleMetricCriteria(subject MetricAlertSingleResourceMultipleMetricCriteria) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlertSingleResourceMultipleMetricCriteria
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

// Generator of MetricAlertSingleResourceMultipleMetricCriteria instances for property testing - lazily instantiated by
// MetricAlertSingleResourceMultipleMetricCriteriaGenerator()
var metricAlertSingleResourceMultipleMetricCriteriaGenerator gopter.Gen

// MetricAlertSingleResourceMultipleMetricCriteriaGenerator returns a generator of MetricAlertSingleResourceMultipleMetricCriteria instances for property testing.
// We first initialize metricAlertSingleResourceMultipleMetricCriteriaGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricAlertSingleResourceMultipleMetricCriteriaGenerator() gopter.Gen {
	if metricAlertSingleResourceMultipleMetricCriteriaGenerator != nil {
		return metricAlertSingleResourceMultipleMetricCriteriaGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria(generators)
	metricAlertSingleResourceMultipleMetricCriteriaGenerator = gen.Struct(reflect.TypeOf(MetricAlertSingleResourceMultipleMetricCriteria{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria(generators)
	AddRelatedPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria(generators)
	metricAlertSingleResourceMultipleMetricCriteriaGenerator = gen.Struct(reflect.TypeOf(MetricAlertSingleResourceMultipleMetricCriteria{}), generators)

	return metricAlertSingleResourceMultipleMetricCriteriaGenerator
}

// AddIndependentPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria(gens map[string]gopter.Gen) {
	gens["OdataType"] = gen.OneConstOf(MetricAlertSingleResourceMultipleMetricCriteria_OdataType_MicrosoftAzureMonitorSingleResourceMultipleMetricCriteria)
}

// AddRelatedPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricAlertSingleResourceMultipleMetricCriteria(gens map[string]gopter.Gen) {
	gens["AllOf"] = gen.SliceOf(MetricCriteriaGenerator())
}

func Test_MetricAlert_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricAlert_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricAlert_Spec, MetricAlert_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricAlert_Spec runs a test to see if a specific instance of MetricAlert_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricAlert_Spec(subject MetricAlert_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricAlert_Spec
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

// Generator of MetricAlert_Spec instances for property testing - lazily instantiated by MetricAlert_SpecGenerator()
var metricAlert_SpecGenerator gopter.Gen

// MetricAlert_SpecGenerator returns a generator of MetricAlert_Spec instances for property testing.
// We first initialize metricAlert_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricAlert_SpecGenerator() gopter.Gen {
	if metricAlert_SpecGenerator != nil {
		return metricAlert_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlert_Spec(generators)
	metricAlert_SpecGenerator = gen.Struct(reflect.TypeOf(MetricAlert_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricAlert_Spec(generators)
	AddRelatedPropertyGeneratorsForMetricAlert_Spec(generators)
	metricAlert_SpecGenerator = gen.Struct(reflect.TypeOf(MetricAlert_Spec{}), generators)

	return metricAlert_SpecGenerator
}

// AddIndependentPropertyGeneratorsForMetricAlert_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricAlert_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMetricAlert_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricAlert_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MetricAlertPropertiesGenerator())
}

func Test_MetricCriteria_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricCriteria via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricCriteria, MetricCriteriaGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricCriteria runs a test to see if a specific instance of MetricCriteria round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricCriteria(subject MetricCriteria) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricCriteria
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

// Generator of MetricCriteria instances for property testing - lazily instantiated by MetricCriteriaGenerator()
var metricCriteriaGenerator gopter.Gen

// MetricCriteriaGenerator returns a generator of MetricCriteria instances for property testing.
// We first initialize metricCriteriaGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricCriteriaGenerator() gopter.Gen {
	if metricCriteriaGenerator != nil {
		return metricCriteriaGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricCriteria(generators)
	metricCriteriaGenerator = gen.Struct(reflect.TypeOf(MetricCriteria{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricCriteria(generators)
	AddRelatedPropertyGeneratorsForMetricCriteria(generators)
	metricCriteriaGenerator = gen.Struct(reflect.TypeOf(MetricCriteria{}), generators)

	return metricCriteriaGenerator
}

// AddIndependentPropertyGeneratorsForMetricCriteria is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricCriteria(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForMetricCriteria is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricCriteria(gens map[string]gopter.Gen) {
	gens["Dimensions"] = gen.SliceOf(MetricDimensionGenerator())
}

func Test_MetricDimension_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricDimension via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricDimension, MetricDimensionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricDimension runs a test to see if a specific instance of MetricDimension round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricDimension(subject MetricDimension) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricDimension
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

// Generator of MetricDimension instances for property testing - lazily instantiated by MetricDimensionGenerator()
var metricDimensionGenerator gopter.Gen

// MetricDimensionGenerator returns a generator of MetricDimension instances for property testing.
func MetricDimensionGenerator() gopter.Gen {
	if metricDimensionGenerator != nil {
		return metricDimensionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricDimension(generators)
	metricDimensionGenerator = gen.Struct(reflect.TypeOf(MetricDimension{}), generators)

	return metricDimensionGenerator
}

// AddIndependentPropertyGeneratorsForMetricDimension is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricDimension(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.AlphaString())
	gens["Values"] = gen.SliceOf(gen.AlphaString())
}

func Test_MultiMetricCriteria_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MultiMetricCriteria via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMultiMetricCriteria, MultiMetricCriteriaGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMultiMetricCriteria runs a test to see if a specific instance of MultiMetricCriteria round trips to JSON and back losslessly
func RunJSONSerializationTestForMultiMetricCriteria(subject MultiMetricCriteria) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MultiMetricCriteria
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

// Generator of MultiMetricCriteria instances for property testing - lazily instantiated by
// MultiMetricCriteriaGenerator()
var multiMetricCriteriaGenerator gopter.Gen

// MultiMetricCriteriaGenerator returns a generator of MultiMetricCriteria instances for property testing.
func MultiMetricCriteriaGenerator() gopter.Gen {
	if multiMetricCriteriaGenerator != nil {
		return multiMetricCriteriaGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMultiMetricCriteria(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(MultiMetricCriteria{}), map[string]gopter.Gen{propName: propGen}))
	}
	multiMetricCriteriaGenerator = gen.OneGenOf(gens...)

	return multiMetricCriteriaGenerator
}

// AddRelatedPropertyGeneratorsForMultiMetricCriteria is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMultiMetricCriteria(gens map[string]gopter.Gen) {
	gens["Dynamic"] = DynamicMetricCriteriaGenerator().Map(func(it DynamicMetricCriteria) *DynamicMetricCriteria {
		return &it
	}) // generate one case for OneOf type
	gens["Static"] = MetricCriteriaGenerator().Map(func(it MetricCriteria) *MetricCriteria {
		return &it
	}) // generate one case for OneOf type
}

func Test_WebtestLocationAvailabilityCriteria_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebtestLocationAvailabilityCriteria via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebtestLocationAvailabilityCriteria, WebtestLocationAvailabilityCriteriaGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebtestLocationAvailabilityCriteria runs a test to see if a specific instance of WebtestLocationAvailabilityCriteria round trips to JSON and back losslessly
func RunJSONSerializationTestForWebtestLocationAvailabilityCriteria(subject WebtestLocationAvailabilityCriteria) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebtestLocationAvailabilityCriteria
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

// Generator of WebtestLocationAvailabilityCriteria instances for property testing - lazily instantiated by
// WebtestLocationAvailabilityCriteriaGenerator()
var webtestLocationAvailabilityCriteriaGenerator gopter.Gen

// WebtestLocationAvailabilityCriteriaGenerator returns a generator of WebtestLocationAvailabilityCriteria instances for property testing.
func WebtestLocationAvailabilityCriteriaGenerator() gopter.Gen {
	if webtestLocationAvailabilityCriteriaGenerator != nil {
		return webtestLocationAvailabilityCriteriaGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebtestLocationAvailabilityCriteria(generators)
	webtestLocationAvailabilityCriteriaGenerator = gen.Struct(reflect.TypeOf(WebtestLocationAvailabilityCriteria{}), generators)

	return webtestLocationAvailabilityCriteriaGenerator
}

// AddIndependentPropertyGeneratorsForWebtestLocationAvailabilityCriteria is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebtestLocationAvailabilityCriteria(gens map[string]gopter.Gen) {
	gens["ComponentId"] = gen.PtrOf(gen.AlphaString())
	gens["FailedLocationCount"] = gen.PtrOf(gen.Float64())
	gens["OdataType"] = gen.OneConstOf(WebtestLocationAvailabilityCriteria_OdataType_MicrosoftAzureMonitorWebtestLocationAvailabilityCriteria)
	gens["WebTestId"] = gen.PtrOf(gen.AlphaString())
}