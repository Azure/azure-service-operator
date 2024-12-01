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

func Test_Actions_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Actions via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForActions, ActionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForActions runs a test to see if a specific instance of Actions round trips to JSON and back losslessly
func RunJSONSerializationTestForActions(subject Actions) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Actions
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

// Generator of Actions instances for property testing - lazily instantiated by ActionsGenerator()
var actionsGenerator gopter.Gen

// ActionsGenerator returns a generator of Actions instances for property testing.
func ActionsGenerator() gopter.Gen {
	if actionsGenerator != nil {
		return actionsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForActions(generators)
	actionsGenerator = gen.Struct(reflect.TypeOf(Actions{}), generators)

	return actionsGenerator
}

// AddIndependentPropertyGeneratorsForActions is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForActions(gens map[string]gopter.Gen) {
	gens["ActionGroups"] = gen.SliceOf(gen.AlphaString())
	gens["CustomProperties"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

func Test_Condition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Condition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCondition, ConditionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCondition runs a test to see if a specific instance of Condition round trips to JSON and back losslessly
func RunJSONSerializationTestForCondition(subject Condition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Condition
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

// Generator of Condition instances for property testing - lazily instantiated by ConditionGenerator()
var conditionGenerator gopter.Gen

// ConditionGenerator returns a generator of Condition instances for property testing.
// We first initialize conditionGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConditionGenerator() gopter.Gen {
	if conditionGenerator != nil {
		return conditionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition(generators)
	conditionGenerator = gen.Struct(reflect.TypeOf(Condition{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition(generators)
	AddRelatedPropertyGeneratorsForCondition(generators)
	conditionGenerator = gen.Struct(reflect.TypeOf(Condition{}), generators)

	return conditionGenerator
}

// AddIndependentPropertyGeneratorsForCondition is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCondition(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForCondition is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCondition(gens map[string]gopter.Gen) {
	gens["Dimensions"] = gen.SliceOf(DimensionGenerator())
	gens["FailingPeriods"] = gen.PtrOf(Condition_FailingPeriodsGenerator())
}

func Test_Condition_FailingPeriods_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Condition_FailingPeriods via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCondition_FailingPeriods, Condition_FailingPeriodsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCondition_FailingPeriods runs a test to see if a specific instance of Condition_FailingPeriods round trips to JSON and back losslessly
func RunJSONSerializationTestForCondition_FailingPeriods(subject Condition_FailingPeriods) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Condition_FailingPeriods
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

// Generator of Condition_FailingPeriods instances for property testing - lazily instantiated by
// Condition_FailingPeriodsGenerator()
var condition_FailingPeriodsGenerator gopter.Gen

// Condition_FailingPeriodsGenerator returns a generator of Condition_FailingPeriods instances for property testing.
func Condition_FailingPeriodsGenerator() gopter.Gen {
	if condition_FailingPeriodsGenerator != nil {
		return condition_FailingPeriodsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition_FailingPeriods(generators)
	condition_FailingPeriodsGenerator = gen.Struct(reflect.TypeOf(Condition_FailingPeriods{}), generators)

	return condition_FailingPeriodsGenerator
}

// AddIndependentPropertyGeneratorsForCondition_FailingPeriods is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCondition_FailingPeriods(gens map[string]gopter.Gen) {
	gens["MinFailingPeriodsToAlert"] = gen.PtrOf(gen.Int())
	gens["NumberOfEvaluationPeriods"] = gen.PtrOf(gen.Int())
}

func Test_Dimension_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Dimension via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDimension, DimensionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDimension runs a test to see if a specific instance of Dimension round trips to JSON and back losslessly
func RunJSONSerializationTestForDimension(subject Dimension) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Dimension
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

// Generator of Dimension instances for property testing - lazily instantiated by DimensionGenerator()
var dimensionGenerator gopter.Gen

// DimensionGenerator returns a generator of Dimension instances for property testing.
func DimensionGenerator() gopter.Gen {
	if dimensionGenerator != nil {
		return dimensionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDimension(generators)
	dimensionGenerator = gen.Struct(reflect.TypeOf(Dimension{}), generators)

	return dimensionGenerator
}

// AddIndependentPropertyGeneratorsForDimension is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDimension(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.OneConstOf(Dimension_Operator_Exclude, Dimension_Operator_Include))
	gens["Values"] = gen.SliceOf(gen.AlphaString())
}

func Test_ScheduledQueryRuleCriteria_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRuleCriteria via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRuleCriteria, ScheduledQueryRuleCriteriaGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRuleCriteria runs a test to see if a specific instance of ScheduledQueryRuleCriteria round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRuleCriteria(subject ScheduledQueryRuleCriteria) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRuleCriteria
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

// Generator of ScheduledQueryRuleCriteria instances for property testing - lazily instantiated by
// ScheduledQueryRuleCriteriaGenerator()
var scheduledQueryRuleCriteriaGenerator gopter.Gen

// ScheduledQueryRuleCriteriaGenerator returns a generator of ScheduledQueryRuleCriteria instances for property testing.
func ScheduledQueryRuleCriteriaGenerator() gopter.Gen {
	if scheduledQueryRuleCriteriaGenerator != nil {
		return scheduledQueryRuleCriteriaGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria(generators)
	scheduledQueryRuleCriteriaGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRuleCriteria{}), generators)

	return scheduledQueryRuleCriteriaGenerator
}

// AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria(gens map[string]gopter.Gen) {
	gens["AllOf"] = gen.SliceOf(ConditionGenerator())
}

func Test_ScheduledQueryRuleProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRuleProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRuleProperties, ScheduledQueryRulePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRuleProperties runs a test to see if a specific instance of ScheduledQueryRuleProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRuleProperties(subject ScheduledQueryRuleProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRuleProperties
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

// Generator of ScheduledQueryRuleProperties instances for property testing - lazily instantiated by
// ScheduledQueryRulePropertiesGenerator()
var scheduledQueryRulePropertiesGenerator gopter.Gen

// ScheduledQueryRulePropertiesGenerator returns a generator of ScheduledQueryRuleProperties instances for property testing.
// We first initialize scheduledQueryRulePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ScheduledQueryRulePropertiesGenerator() gopter.Gen {
	if scheduledQueryRulePropertiesGenerator != nil {
		return scheduledQueryRulePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties(generators)
	scheduledQueryRulePropertiesGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRuleProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties(generators)
	AddRelatedPropertyGeneratorsForScheduledQueryRuleProperties(generators)
	scheduledQueryRulePropertiesGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRuleProperties{}), generators)

	return scheduledQueryRulePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForScheduledQueryRuleProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRuleProperties(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.PtrOf(ActionsGenerator())
	gens["Criteria"] = gen.PtrOf(ScheduledQueryRuleCriteriaGenerator())
}

func Test_ScheduledQueryRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRule_Spec, ScheduledQueryRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRule_Spec runs a test to see if a specific instance of ScheduledQueryRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRule_Spec(subject ScheduledQueryRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRule_Spec
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

// Generator of ScheduledQueryRule_Spec instances for property testing - lazily instantiated by
// ScheduledQueryRule_SpecGenerator()
var scheduledQueryRule_SpecGenerator gopter.Gen

// ScheduledQueryRule_SpecGenerator returns a generator of ScheduledQueryRule_Spec instances for property testing.
// We first initialize scheduledQueryRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ScheduledQueryRule_SpecGenerator() gopter.Gen {
	if scheduledQueryRule_SpecGenerator != nil {
		return scheduledQueryRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec(generators)
	scheduledQueryRule_SpecGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec(generators)
	AddRelatedPropertyGeneratorsForScheduledQueryRule_Spec(generators)
	scheduledQueryRule_SpecGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRule_Spec{}), generators)

	return scheduledQueryRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(ScheduledQueryRule_Kind_Spec_LogAlert, ScheduledQueryRule_Kind_Spec_LogToMetric))
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForScheduledQueryRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRule_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ScheduledQueryRulePropertiesGenerator())
}
