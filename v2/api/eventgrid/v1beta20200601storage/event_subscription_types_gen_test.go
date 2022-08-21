// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601storage

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

func Test_EventSubscription_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscription via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscription, EventSubscriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscription runs a test to see if a specific instance of EventSubscription round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscription(subject EventSubscription) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscription
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

// Generator of EventSubscription instances for property testing - lazily instantiated by EventSubscriptionGenerator()
var eventSubscriptionGenerator gopter.Gen

// EventSubscriptionGenerator returns a generator of EventSubscription instances for property testing.
func EventSubscriptionGenerator() gopter.Gen {
	if eventSubscriptionGenerator != nil {
		return eventSubscriptionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForEventSubscription(generators)
	eventSubscriptionGenerator = gen.Struct(reflect.TypeOf(EventSubscription{}), generators)

	return eventSubscriptionGenerator
}

// AddRelatedPropertyGeneratorsForEventSubscription is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventSubscription(gens map[string]gopter.Gen) {
	gens["Spec"] = EventSubscription_SpecGenerator()
	gens["Status"] = EventSubscription_STATUSGenerator()
}

func Test_EventSubscription_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscription_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscription_Spec, EventSubscription_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscription_Spec runs a test to see if a specific instance of EventSubscription_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscription_Spec(subject EventSubscription_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscription_Spec
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

// Generator of EventSubscription_Spec instances for property testing - lazily instantiated by
// EventSubscription_SpecGenerator()
var eventSubscription_SpecGenerator gopter.Gen

// EventSubscription_SpecGenerator returns a generator of EventSubscription_Spec instances for property testing.
// We first initialize eventSubscription_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventSubscription_SpecGenerator() gopter.Gen {
	if eventSubscription_SpecGenerator != nil {
		return eventSubscription_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscription_Spec(generators)
	eventSubscription_SpecGenerator = gen.Struct(reflect.TypeOf(EventSubscription_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscription_Spec(generators)
	AddRelatedPropertyGeneratorsForEventSubscription_Spec(generators)
	eventSubscription_SpecGenerator = gen.Struct(reflect.TypeOf(EventSubscription_Spec{}), generators)

	return eventSubscription_SpecGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscription_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscription_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EventDeliverySchema"] = gen.PtrOf(gen.AlphaString())
	gens["ExpirationTimeUtc"] = gen.PtrOf(gen.AlphaString())
	gens["Labels"] = gen.SliceOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForEventSubscription_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventSubscription_Spec(gens map[string]gopter.Gen) {
	gens["DeadLetterDestination"] = gen.PtrOf(DeadLetterDestinationGenerator())
	gens["Destination"] = gen.PtrOf(EventSubscriptionDestinationGenerator())
	gens["Filter"] = gen.PtrOf(EventSubscriptionFilterGenerator())
	gens["RetryPolicy"] = gen.PtrOf(RetryPolicyGenerator())
}

func Test_EventSubscription_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscription_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscription_STATUS, EventSubscription_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscription_STATUS runs a test to see if a specific instance of EventSubscription_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscription_STATUS(subject EventSubscription_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscription_STATUS
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

// Generator of EventSubscription_STATUS instances for property testing - lazily instantiated by
// EventSubscription_STATUSGenerator()
var eventSubscription_STATUSGenerator gopter.Gen

// EventSubscription_STATUSGenerator returns a generator of EventSubscription_STATUS instances for property testing.
// We first initialize eventSubscription_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventSubscription_STATUSGenerator() gopter.Gen {
	if eventSubscription_STATUSGenerator != nil {
		return eventSubscription_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscription_STATUS(generators)
	eventSubscription_STATUSGenerator = gen.Struct(reflect.TypeOf(EventSubscription_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscription_STATUS(generators)
	AddRelatedPropertyGeneratorsForEventSubscription_STATUS(generators)
	eventSubscription_STATUSGenerator = gen.Struct(reflect.TypeOf(EventSubscription_STATUS{}), generators)

	return eventSubscription_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscription_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscription_STATUS(gens map[string]gopter.Gen) {
	gens["EventDeliverySchema"] = gen.PtrOf(gen.AlphaString())
	gens["ExpirationTimeUtc"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Labels"] = gen.SliceOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Topic"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventSubscription_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventSubscription_STATUS(gens map[string]gopter.Gen) {
	gens["DeadLetterDestination"] = gen.PtrOf(DeadLetterDestination_STATUSGenerator())
	gens["Destination"] = gen.PtrOf(EventSubscriptionDestination_STATUSGenerator())
	gens["Filter"] = gen.PtrOf(EventSubscriptionFilter_STATUSGenerator())
	gens["RetryPolicy"] = gen.PtrOf(RetryPolicy_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_DeadLetterDestination_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeadLetterDestination via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeadLetterDestination, DeadLetterDestinationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeadLetterDestination runs a test to see if a specific instance of DeadLetterDestination round trips to JSON and back losslessly
func RunJSONSerializationTestForDeadLetterDestination(subject DeadLetterDestination) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeadLetterDestination
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

// Generator of DeadLetterDestination instances for property testing - lazily instantiated by
// DeadLetterDestinationGenerator()
var deadLetterDestinationGenerator gopter.Gen

// DeadLetterDestinationGenerator returns a generator of DeadLetterDestination instances for property testing.
func DeadLetterDestinationGenerator() gopter.Gen {
	if deadLetterDestinationGenerator != nil {
		return deadLetterDestinationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeadLetterDestination(generators)
	deadLetterDestinationGenerator = gen.Struct(reflect.TypeOf(DeadLetterDestination{}), generators)

	return deadLetterDestinationGenerator
}

// AddIndependentPropertyGeneratorsForDeadLetterDestination is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeadLetterDestination(gens map[string]gopter.Gen) {
	gens["EndpointType"] = gen.PtrOf(gen.AlphaString())
}

func Test_DeadLetterDestination_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeadLetterDestination_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeadLetterDestination_STATUS, DeadLetterDestination_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeadLetterDestination_STATUS runs a test to see if a specific instance of DeadLetterDestination_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDeadLetterDestination_STATUS(subject DeadLetterDestination_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeadLetterDestination_STATUS
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

// Generator of DeadLetterDestination_STATUS instances for property testing - lazily instantiated by
// DeadLetterDestination_STATUSGenerator()
var deadLetterDestination_STATUSGenerator gopter.Gen

// DeadLetterDestination_STATUSGenerator returns a generator of DeadLetterDestination_STATUS instances for property testing.
func DeadLetterDestination_STATUSGenerator() gopter.Gen {
	if deadLetterDestination_STATUSGenerator != nil {
		return deadLetterDestination_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeadLetterDestination_STATUS(generators)
	deadLetterDestination_STATUSGenerator = gen.Struct(reflect.TypeOf(DeadLetterDestination_STATUS{}), generators)

	return deadLetterDestination_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDeadLetterDestination_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeadLetterDestination_STATUS(gens map[string]gopter.Gen) {
	gens["EndpointType"] = gen.PtrOf(gen.AlphaString())
}

func Test_EventSubscriptionDestination_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscriptionDestination via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionDestination, EventSubscriptionDestinationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionDestination runs a test to see if a specific instance of EventSubscriptionDestination round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionDestination(subject EventSubscriptionDestination) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscriptionDestination
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

// Generator of EventSubscriptionDestination instances for property testing - lazily instantiated by
// EventSubscriptionDestinationGenerator()
var eventSubscriptionDestinationGenerator gopter.Gen

// EventSubscriptionDestinationGenerator returns a generator of EventSubscriptionDestination instances for property testing.
func EventSubscriptionDestinationGenerator() gopter.Gen {
	if eventSubscriptionDestinationGenerator != nil {
		return eventSubscriptionDestinationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionDestination(generators)
	eventSubscriptionDestinationGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionDestination{}), generators)

	return eventSubscriptionDestinationGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionDestination is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionDestination(gens map[string]gopter.Gen) {
	gens["EndpointType"] = gen.PtrOf(gen.AlphaString())
}

func Test_EventSubscriptionDestination_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscriptionDestination_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionDestination_STATUS, EventSubscriptionDestination_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionDestination_STATUS runs a test to see if a specific instance of EventSubscriptionDestination_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionDestination_STATUS(subject EventSubscriptionDestination_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscriptionDestination_STATUS
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

// Generator of EventSubscriptionDestination_STATUS instances for property testing - lazily instantiated by
// EventSubscriptionDestination_STATUSGenerator()
var eventSubscriptionDestination_STATUSGenerator gopter.Gen

// EventSubscriptionDestination_STATUSGenerator returns a generator of EventSubscriptionDestination_STATUS instances for property testing.
func EventSubscriptionDestination_STATUSGenerator() gopter.Gen {
	if eventSubscriptionDestination_STATUSGenerator != nil {
		return eventSubscriptionDestination_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionDestination_STATUS(generators)
	eventSubscriptionDestination_STATUSGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionDestination_STATUS{}), generators)

	return eventSubscriptionDestination_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionDestination_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionDestination_STATUS(gens map[string]gopter.Gen) {
	gens["EndpointType"] = gen.PtrOf(gen.AlphaString())
}

func Test_EventSubscriptionFilter_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscriptionFilter via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionFilter, EventSubscriptionFilterGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionFilter runs a test to see if a specific instance of EventSubscriptionFilter round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionFilter(subject EventSubscriptionFilter) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscriptionFilter
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

// Generator of EventSubscriptionFilter instances for property testing - lazily instantiated by
// EventSubscriptionFilterGenerator()
var eventSubscriptionFilterGenerator gopter.Gen

// EventSubscriptionFilterGenerator returns a generator of EventSubscriptionFilter instances for property testing.
// We first initialize eventSubscriptionFilterGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventSubscriptionFilterGenerator() gopter.Gen {
	if eventSubscriptionFilterGenerator != nil {
		return eventSubscriptionFilterGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionFilter(generators)
	eventSubscriptionFilterGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionFilter{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionFilter(generators)
	AddRelatedPropertyGeneratorsForEventSubscriptionFilter(generators)
	eventSubscriptionFilterGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionFilter{}), generators)

	return eventSubscriptionFilterGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionFilter is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionFilter(gens map[string]gopter.Gen) {
	gens["IncludedEventTypes"] = gen.SliceOf(gen.AlphaString())
	gens["IsSubjectCaseSensitive"] = gen.PtrOf(gen.Bool())
	gens["SubjectBeginsWith"] = gen.PtrOf(gen.AlphaString())
	gens["SubjectEndsWith"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventSubscriptionFilter is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventSubscriptionFilter(gens map[string]gopter.Gen) {
	gens["AdvancedFilters"] = gen.SliceOf(AdvancedFilterGenerator())
}

func Test_EventSubscriptionFilter_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscriptionFilter_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionFilter_STATUS, EventSubscriptionFilter_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionFilter_STATUS runs a test to see if a specific instance of EventSubscriptionFilter_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionFilter_STATUS(subject EventSubscriptionFilter_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscriptionFilter_STATUS
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

// Generator of EventSubscriptionFilter_STATUS instances for property testing - lazily instantiated by
// EventSubscriptionFilter_STATUSGenerator()
var eventSubscriptionFilter_STATUSGenerator gopter.Gen

// EventSubscriptionFilter_STATUSGenerator returns a generator of EventSubscriptionFilter_STATUS instances for property testing.
// We first initialize eventSubscriptionFilter_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventSubscriptionFilter_STATUSGenerator() gopter.Gen {
	if eventSubscriptionFilter_STATUSGenerator != nil {
		return eventSubscriptionFilter_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionFilter_STATUS(generators)
	eventSubscriptionFilter_STATUSGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionFilter_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionFilter_STATUS(generators)
	AddRelatedPropertyGeneratorsForEventSubscriptionFilter_STATUS(generators)
	eventSubscriptionFilter_STATUSGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionFilter_STATUS{}), generators)

	return eventSubscriptionFilter_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionFilter_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionFilter_STATUS(gens map[string]gopter.Gen) {
	gens["IncludedEventTypes"] = gen.SliceOf(gen.AlphaString())
	gens["IsSubjectCaseSensitive"] = gen.PtrOf(gen.Bool())
	gens["SubjectBeginsWith"] = gen.PtrOf(gen.AlphaString())
	gens["SubjectEndsWith"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventSubscriptionFilter_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventSubscriptionFilter_STATUS(gens map[string]gopter.Gen) {
	gens["AdvancedFilters"] = gen.SliceOf(AdvancedFilter_STATUSGenerator())
}

func Test_RetryPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RetryPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRetryPolicy, RetryPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRetryPolicy runs a test to see if a specific instance of RetryPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForRetryPolicy(subject RetryPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RetryPolicy
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

// Generator of RetryPolicy instances for property testing - lazily instantiated by RetryPolicyGenerator()
var retryPolicyGenerator gopter.Gen

// RetryPolicyGenerator returns a generator of RetryPolicy instances for property testing.
func RetryPolicyGenerator() gopter.Gen {
	if retryPolicyGenerator != nil {
		return retryPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRetryPolicy(generators)
	retryPolicyGenerator = gen.Struct(reflect.TypeOf(RetryPolicy{}), generators)

	return retryPolicyGenerator
}

// AddIndependentPropertyGeneratorsForRetryPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRetryPolicy(gens map[string]gopter.Gen) {
	gens["EventTimeToLiveInMinutes"] = gen.PtrOf(gen.Int())
	gens["MaxDeliveryAttempts"] = gen.PtrOf(gen.Int())
}

func Test_RetryPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RetryPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRetryPolicy_STATUS, RetryPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRetryPolicy_STATUS runs a test to see if a specific instance of RetryPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRetryPolicy_STATUS(subject RetryPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RetryPolicy_STATUS
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

// Generator of RetryPolicy_STATUS instances for property testing - lazily instantiated by RetryPolicy_STATUSGenerator()
var retryPolicy_STATUSGenerator gopter.Gen

// RetryPolicy_STATUSGenerator returns a generator of RetryPolicy_STATUS instances for property testing.
func RetryPolicy_STATUSGenerator() gopter.Gen {
	if retryPolicy_STATUSGenerator != nil {
		return retryPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRetryPolicy_STATUS(generators)
	retryPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(RetryPolicy_STATUS{}), generators)

	return retryPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRetryPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRetryPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["EventTimeToLiveInMinutes"] = gen.PtrOf(gen.Int())
	gens["MaxDeliveryAttempts"] = gen.PtrOf(gen.Int())
}

func Test_AdvancedFilter_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AdvancedFilter via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAdvancedFilter, AdvancedFilterGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAdvancedFilter runs a test to see if a specific instance of AdvancedFilter round trips to JSON and back losslessly
func RunJSONSerializationTestForAdvancedFilter(subject AdvancedFilter) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AdvancedFilter
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

// Generator of AdvancedFilter instances for property testing - lazily instantiated by AdvancedFilterGenerator()
var advancedFilterGenerator gopter.Gen

// AdvancedFilterGenerator returns a generator of AdvancedFilter instances for property testing.
func AdvancedFilterGenerator() gopter.Gen {
	if advancedFilterGenerator != nil {
		return advancedFilterGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAdvancedFilter(generators)
	advancedFilterGenerator = gen.Struct(reflect.TypeOf(AdvancedFilter{}), generators)

	return advancedFilterGenerator
}

// AddIndependentPropertyGeneratorsForAdvancedFilter is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAdvancedFilter(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(gen.AlphaString())
	gens["OperatorType"] = gen.PtrOf(gen.AlphaString())
}

func Test_AdvancedFilter_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AdvancedFilter_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAdvancedFilter_STATUS, AdvancedFilter_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAdvancedFilter_STATUS runs a test to see if a specific instance of AdvancedFilter_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAdvancedFilter_STATUS(subject AdvancedFilter_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AdvancedFilter_STATUS
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

// Generator of AdvancedFilter_STATUS instances for property testing - lazily instantiated by
// AdvancedFilter_STATUSGenerator()
var advancedFilter_STATUSGenerator gopter.Gen

// AdvancedFilter_STATUSGenerator returns a generator of AdvancedFilter_STATUS instances for property testing.
func AdvancedFilter_STATUSGenerator() gopter.Gen {
	if advancedFilter_STATUSGenerator != nil {
		return advancedFilter_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAdvancedFilter_STATUS(generators)
	advancedFilter_STATUSGenerator = gen.Struct(reflect.TypeOf(AdvancedFilter_STATUS{}), generators)

	return advancedFilter_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAdvancedFilter_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAdvancedFilter_STATUS(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(gen.AlphaString())
	gens["OperatorType"] = gen.PtrOf(gen.AlphaString())
}
