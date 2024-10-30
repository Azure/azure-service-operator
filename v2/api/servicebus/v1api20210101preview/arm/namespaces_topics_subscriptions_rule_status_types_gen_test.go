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

func Test_Action_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Action_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAction_STATUS, Action_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAction_STATUS runs a test to see if a specific instance of Action_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAction_STATUS(subject Action_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Action_STATUS
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

// Generator of Action_STATUS instances for property testing - lazily instantiated by Action_STATUSGenerator()
var action_STATUSGenerator gopter.Gen

// Action_STATUSGenerator returns a generator of Action_STATUS instances for property testing.
func Action_STATUSGenerator() gopter.Gen {
	if action_STATUSGenerator != nil {
		return action_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAction_STATUS(generators)
	action_STATUSGenerator = gen.Struct(reflect.TypeOf(Action_STATUS{}), generators)

	return action_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAction_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAction_STATUS(gens map[string]gopter.Gen) {
	gens["CompatibilityLevel"] = gen.PtrOf(gen.Int())
	gens["RequiresPreprocessing"] = gen.PtrOf(gen.Bool())
	gens["SqlExpression"] = gen.PtrOf(gen.AlphaString())
}

func Test_CorrelationFilter_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorrelationFilter_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorrelationFilter_STATUS, CorrelationFilter_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorrelationFilter_STATUS runs a test to see if a specific instance of CorrelationFilter_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCorrelationFilter_STATUS(subject CorrelationFilter_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorrelationFilter_STATUS
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

// Generator of CorrelationFilter_STATUS instances for property testing - lazily instantiated by
// CorrelationFilter_STATUSGenerator()
var correlationFilter_STATUSGenerator gopter.Gen

// CorrelationFilter_STATUSGenerator returns a generator of CorrelationFilter_STATUS instances for property testing.
func CorrelationFilter_STATUSGenerator() gopter.Gen {
	if correlationFilter_STATUSGenerator != nil {
		return correlationFilter_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorrelationFilter_STATUS(generators)
	correlationFilter_STATUSGenerator = gen.Struct(reflect.TypeOf(CorrelationFilter_STATUS{}), generators)

	return correlationFilter_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCorrelationFilter_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorrelationFilter_STATUS(gens map[string]gopter.Gen) {
	gens["ContentType"] = gen.PtrOf(gen.AlphaString())
	gens["CorrelationId"] = gen.PtrOf(gen.AlphaString())
	gens["Label"] = gen.PtrOf(gen.AlphaString())
	gens["MessageId"] = gen.PtrOf(gen.AlphaString())
	gens["Properties"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["ReplyTo"] = gen.PtrOf(gen.AlphaString())
	gens["ReplyToSessionId"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresPreprocessing"] = gen.PtrOf(gen.Bool())
	gens["SessionId"] = gen.PtrOf(gen.AlphaString())
	gens["To"] = gen.PtrOf(gen.AlphaString())
}

func Test_NamespacesTopicsSubscriptionsRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopicsSubscriptionsRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopicsSubscriptionsRule_STATUS, NamespacesTopicsSubscriptionsRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopicsSubscriptionsRule_STATUS runs a test to see if a specific instance of NamespacesTopicsSubscriptionsRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopicsSubscriptionsRule_STATUS(subject NamespacesTopicsSubscriptionsRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopicsSubscriptionsRule_STATUS
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

// Generator of NamespacesTopicsSubscriptionsRule_STATUS instances for property testing - lazily instantiated by
// NamespacesTopicsSubscriptionsRule_STATUSGenerator()
var namespacesTopicsSubscriptionsRule_STATUSGenerator gopter.Gen

// NamespacesTopicsSubscriptionsRule_STATUSGenerator returns a generator of NamespacesTopicsSubscriptionsRule_STATUS instances for property testing.
// We first initialize namespacesTopicsSubscriptionsRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesTopicsSubscriptionsRule_STATUSGenerator() gopter.Gen {
	if namespacesTopicsSubscriptionsRule_STATUSGenerator != nil {
		return namespacesTopicsSubscriptionsRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopicsSubscriptionsRule_STATUS(generators)
	namespacesTopicsSubscriptionsRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesTopicsSubscriptionsRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopicsSubscriptionsRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespacesTopicsSubscriptionsRule_STATUS(generators)
	namespacesTopicsSubscriptionsRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesTopicsSubscriptionsRule_STATUS{}), generators)

	return namespacesTopicsSubscriptionsRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesTopicsSubscriptionsRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesTopicsSubscriptionsRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesTopicsSubscriptionsRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopicsSubscriptionsRule_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Ruleproperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Ruleproperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Ruleproperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRuleproperties_STATUS, Ruleproperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRuleproperties_STATUS runs a test to see if a specific instance of Ruleproperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRuleproperties_STATUS(subject Ruleproperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Ruleproperties_STATUS
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

// Generator of Ruleproperties_STATUS instances for property testing - lazily instantiated by
// Ruleproperties_STATUSGenerator()
var ruleproperties_STATUSGenerator gopter.Gen

// Ruleproperties_STATUSGenerator returns a generator of Ruleproperties_STATUS instances for property testing.
// We first initialize ruleproperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Ruleproperties_STATUSGenerator() gopter.Gen {
	if ruleproperties_STATUSGenerator != nil {
		return ruleproperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleproperties_STATUS(generators)
	ruleproperties_STATUSGenerator = gen.Struct(reflect.TypeOf(Ruleproperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleproperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForRuleproperties_STATUS(generators)
	ruleproperties_STATUSGenerator = gen.Struct(reflect.TypeOf(Ruleproperties_STATUS{}), generators)

	return ruleproperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRuleproperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRuleproperties_STATUS(gens map[string]gopter.Gen) {
	gens["FilterType"] = gen.PtrOf(gen.OneConstOf(FilterType_STATUS_CorrelationFilter, FilterType_STATUS_SqlFilter))
}

// AddRelatedPropertyGeneratorsForRuleproperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRuleproperties_STATUS(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(Action_STATUSGenerator())
	gens["CorrelationFilter"] = gen.PtrOf(CorrelationFilter_STATUSGenerator())
	gens["SqlFilter"] = gen.PtrOf(SqlFilter_STATUSGenerator())
}

func Test_SqlFilter_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlFilter_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlFilter_STATUS, SqlFilter_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlFilter_STATUS runs a test to see if a specific instance of SqlFilter_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlFilter_STATUS(subject SqlFilter_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlFilter_STATUS
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

// Generator of SqlFilter_STATUS instances for property testing - lazily instantiated by SqlFilter_STATUSGenerator()
var sqlFilter_STATUSGenerator gopter.Gen

// SqlFilter_STATUSGenerator returns a generator of SqlFilter_STATUS instances for property testing.
func SqlFilter_STATUSGenerator() gopter.Gen {
	if sqlFilter_STATUSGenerator != nil {
		return sqlFilter_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlFilter_STATUS(generators)
	sqlFilter_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlFilter_STATUS{}), generators)

	return sqlFilter_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlFilter_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlFilter_STATUS(gens map[string]gopter.Gen) {
	gens["CompatibilityLevel"] = gen.PtrOf(gen.Int())
	gens["RequiresPreprocessing"] = gen.PtrOf(gen.Bool())
	gens["SqlExpression"] = gen.PtrOf(gen.AlphaString())
}
