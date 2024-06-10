// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220615

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

func Test_Actions_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Actions_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForActions_ARM, Actions_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForActions_ARM runs a test to see if a specific instance of Actions_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForActions_ARM(subject Actions_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Actions_ARM
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

// Generator of Actions_ARM instances for property testing - lazily instantiated by Actions_ARMGenerator()
var actions_ARMGenerator gopter.Gen

// Actions_ARMGenerator returns a generator of Actions_ARM instances for property testing.
func Actions_ARMGenerator() gopter.Gen {
	if actions_ARMGenerator != nil {
		return actions_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForActions_ARM(generators)
	actions_ARMGenerator = gen.Struct(reflect.TypeOf(Actions_ARM{}), generators)

	return actions_ARMGenerator
}

// AddIndependentPropertyGeneratorsForActions_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForActions_ARM(gens map[string]gopter.Gen) {
	gens["ActionGroups"] = gen.SliceOf(gen.AlphaString())
	gens["CustomProperties"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

func Test_Condition_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Condition_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCondition_ARM, Condition_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCondition_ARM runs a test to see if a specific instance of Condition_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCondition_ARM(subject Condition_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Condition_ARM
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

// Generator of Condition_ARM instances for property testing - lazily instantiated by Condition_ARMGenerator()
var condition_ARMGenerator gopter.Gen

// Condition_ARMGenerator returns a generator of Condition_ARM instances for property testing.
// We first initialize condition_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Condition_ARMGenerator() gopter.Gen {
	if condition_ARMGenerator != nil {
		return condition_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition_ARM(generators)
	condition_ARMGenerator = gen.Struct(reflect.TypeOf(Condition_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition_ARM(generators)
	AddRelatedPropertyGeneratorsForCondition_ARM(generators)
	condition_ARMGenerator = gen.Struct(reflect.TypeOf(Condition_ARM{}), generators)

	return condition_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCondition_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCondition_ARM(gens map[string]gopter.Gen) {
	gens["MetricMeasureColumn"] = gen.PtrOf(gen.AlphaString())
	gens["MetricName"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.OneConstOf(
		Condition_Operator_Equals,
		Condition_Operator_GreaterThan,
		Condition_Operator_GreaterThanOrEqual,
		Condition_Operator_LessThan,
		Condition_Operator_LessThanOrEqual))
	gens["Query"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceIdColumn"] = gen.PtrOf(gen.AlphaString())
	gens["Threshold"] = gen.PtrOf(gen.Float64())
	gens["TimeAggregation"] = gen.PtrOf(gen.OneConstOf(
		Condition_TimeAggregation_Average,
		Condition_TimeAggregation_Count,
		Condition_TimeAggregation_Maximum,
		Condition_TimeAggregation_Minimum,
		Condition_TimeAggregation_Total))
}

// AddRelatedPropertyGeneratorsForCondition_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCondition_ARM(gens map[string]gopter.Gen) {
	gens["Dimensions"] = gen.SliceOf(Dimension_ARMGenerator())
	gens["FailingPeriods"] = gen.PtrOf(Condition_FailingPeriods_ARMGenerator())
}

func Test_Condition_FailingPeriods_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Condition_FailingPeriods_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCondition_FailingPeriods_ARM, Condition_FailingPeriods_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCondition_FailingPeriods_ARM runs a test to see if a specific instance of Condition_FailingPeriods_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCondition_FailingPeriods_ARM(subject Condition_FailingPeriods_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Condition_FailingPeriods_ARM
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

// Generator of Condition_FailingPeriods_ARM instances for property testing - lazily instantiated by
// Condition_FailingPeriods_ARMGenerator()
var condition_FailingPeriods_ARMGenerator gopter.Gen

// Condition_FailingPeriods_ARMGenerator returns a generator of Condition_FailingPeriods_ARM instances for property testing.
func Condition_FailingPeriods_ARMGenerator() gopter.Gen {
	if condition_FailingPeriods_ARMGenerator != nil {
		return condition_FailingPeriods_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition_FailingPeriods_ARM(generators)
	condition_FailingPeriods_ARMGenerator = gen.Struct(reflect.TypeOf(Condition_FailingPeriods_ARM{}), generators)

	return condition_FailingPeriods_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCondition_FailingPeriods_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCondition_FailingPeriods_ARM(gens map[string]gopter.Gen) {
	gens["MinFailingPeriodsToAlert"] = gen.PtrOf(gen.Int())
	gens["NumberOfEvaluationPeriods"] = gen.PtrOf(gen.Int())
}

func Test_Dimension_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Dimension_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDimension_ARM, Dimension_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDimension_ARM runs a test to see if a specific instance of Dimension_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDimension_ARM(subject Dimension_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Dimension_ARM
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

// Generator of Dimension_ARM instances for property testing - lazily instantiated by Dimension_ARMGenerator()
var dimension_ARMGenerator gopter.Gen

// Dimension_ARMGenerator returns a generator of Dimension_ARM instances for property testing.
func Dimension_ARMGenerator() gopter.Gen {
	if dimension_ARMGenerator != nil {
		return dimension_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDimension_ARM(generators)
	dimension_ARMGenerator = gen.Struct(reflect.TypeOf(Dimension_ARM{}), generators)

	return dimension_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDimension_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDimension_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.OneConstOf(Dimension_Operator_Exclude, Dimension_Operator_Include))
	gens["Values"] = gen.SliceOf(gen.AlphaString())
}

func Test_ScheduledQueryRuleCriteria_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRuleCriteria_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRuleCriteria_ARM, ScheduledQueryRuleCriteria_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRuleCriteria_ARM runs a test to see if a specific instance of ScheduledQueryRuleCriteria_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRuleCriteria_ARM(subject ScheduledQueryRuleCriteria_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRuleCriteria_ARM
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

// Generator of ScheduledQueryRuleCriteria_ARM instances for property testing - lazily instantiated by
// ScheduledQueryRuleCriteria_ARMGenerator()
var scheduledQueryRuleCriteria_ARMGenerator gopter.Gen

// ScheduledQueryRuleCriteria_ARMGenerator returns a generator of ScheduledQueryRuleCriteria_ARM instances for property testing.
func ScheduledQueryRuleCriteria_ARMGenerator() gopter.Gen {
	if scheduledQueryRuleCriteria_ARMGenerator != nil {
		return scheduledQueryRuleCriteria_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria_ARM(generators)
	scheduledQueryRuleCriteria_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRuleCriteria_ARM{}), generators)

	return scheduledQueryRuleCriteria_ARMGenerator
}

// AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria_ARM(gens map[string]gopter.Gen) {
	gens["AllOf"] = gen.SliceOf(Condition_ARMGenerator())
}

func Test_ScheduledQueryRuleProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRuleProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRuleProperties_ARM, ScheduledQueryRuleProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRuleProperties_ARM runs a test to see if a specific instance of ScheduledQueryRuleProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRuleProperties_ARM(subject ScheduledQueryRuleProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRuleProperties_ARM
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

// Generator of ScheduledQueryRuleProperties_ARM instances for property testing - lazily instantiated by
// ScheduledQueryRuleProperties_ARMGenerator()
var scheduledQueryRuleProperties_ARMGenerator gopter.Gen

// ScheduledQueryRuleProperties_ARMGenerator returns a generator of ScheduledQueryRuleProperties_ARM instances for property testing.
// We first initialize scheduledQueryRuleProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ScheduledQueryRuleProperties_ARMGenerator() gopter.Gen {
	if scheduledQueryRuleProperties_ARMGenerator != nil {
		return scheduledQueryRuleProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties_ARM(generators)
	scheduledQueryRuleProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRuleProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForScheduledQueryRuleProperties_ARM(generators)
	scheduledQueryRuleProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRuleProperties_ARM{}), generators)

	return scheduledQueryRuleProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties_ARM(gens map[string]gopter.Gen) {
	gens["AutoMitigate"] = gen.PtrOf(gen.Bool())
	gens["CheckWorkspaceAlertsStorageConfigured"] = gen.PtrOf(gen.Bool())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["EvaluationFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["MuteActionsDuration"] = gen.PtrOf(gen.AlphaString())
	gens["OverrideQueryTimeRange"] = gen.PtrOf(gen.AlphaString())
	gens["Scopes"] = gen.SliceOf(gen.AlphaString())
	gens["Severity"] = gen.PtrOf(gen.OneConstOf(
		ScheduledQueryRuleProperties_Severity_0,
		ScheduledQueryRuleProperties_Severity_1,
		ScheduledQueryRuleProperties_Severity_2,
		ScheduledQueryRuleProperties_Severity_3,
		ScheduledQueryRuleProperties_Severity_4))
	gens["SkipQueryValidation"] = gen.PtrOf(gen.Bool())
	gens["TargetResourceTypes"] = gen.SliceOf(gen.AlphaString())
	gens["WindowSize"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForScheduledQueryRuleProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRuleProperties_ARM(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.PtrOf(Actions_ARMGenerator())
	gens["Criteria"] = gen.PtrOf(ScheduledQueryRuleCriteria_ARMGenerator())
}

func Test_ScheduledQueryRule_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRule_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRule_Spec_ARM, ScheduledQueryRule_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRule_Spec_ARM runs a test to see if a specific instance of ScheduledQueryRule_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRule_Spec_ARM(subject ScheduledQueryRule_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRule_Spec_ARM
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

// Generator of ScheduledQueryRule_Spec_ARM instances for property testing - lazily instantiated by
// ScheduledQueryRule_Spec_ARMGenerator()
var scheduledQueryRule_Spec_ARMGenerator gopter.Gen

// ScheduledQueryRule_Spec_ARMGenerator returns a generator of ScheduledQueryRule_Spec_ARM instances for property testing.
// We first initialize scheduledQueryRule_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ScheduledQueryRule_Spec_ARMGenerator() gopter.Gen {
	if scheduledQueryRule_Spec_ARMGenerator != nil {
		return scheduledQueryRule_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec_ARM(generators)
	scheduledQueryRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRule_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForScheduledQueryRule_Spec_ARM(generators)
	scheduledQueryRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRule_Spec_ARM{}), generators)

	return scheduledQueryRule_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(ScheduledQueryRule_Kind_Spec_LogAlert, ScheduledQueryRule_Kind_Spec_LogToMetric))
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForScheduledQueryRule_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ScheduledQueryRuleProperties_ARMGenerator())
}
