// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240101

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20240101/storage"
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

func Test_NamespacesTopicsSubscription_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopicsSubscription to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesTopicsSubscription, NamespacesTopicsSubscriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesTopicsSubscription tests if a specific instance of NamespacesTopicsSubscription round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesTopicsSubscription(subject NamespacesTopicsSubscription) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.NamespacesTopicsSubscription
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesTopicsSubscription
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

func Test_NamespacesTopicsSubscription_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopicsSubscription to NamespacesTopicsSubscription via AssignProperties_To_NamespacesTopicsSubscription & AssignProperties_From_NamespacesTopicsSubscription returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopicsSubscription, NamespacesTopicsSubscriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopicsSubscription tests if a specific instance of NamespacesTopicsSubscription can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopicsSubscription(subject NamespacesTopicsSubscription) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.NamespacesTopicsSubscription
	err := copied.AssignProperties_To_NamespacesTopicsSubscription(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopicsSubscription
	err = actual.AssignProperties_From_NamespacesTopicsSubscription(&other)
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

func Test_NamespacesTopicsSubscription_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopicsSubscription via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopicsSubscription, NamespacesTopicsSubscriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopicsSubscription runs a test to see if a specific instance of NamespacesTopicsSubscription round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopicsSubscription(subject NamespacesTopicsSubscription) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopicsSubscription
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

// Generator of NamespacesTopicsSubscription instances for property testing - lazily instantiated by
// NamespacesTopicsSubscriptionGenerator()
var namespacesTopicsSubscriptionGenerator gopter.Gen

// NamespacesTopicsSubscriptionGenerator returns a generator of NamespacesTopicsSubscription instances for property testing.
func NamespacesTopicsSubscriptionGenerator() gopter.Gen {
	if namespacesTopicsSubscriptionGenerator != nil {
		return namespacesTopicsSubscriptionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesTopicsSubscription(generators)
	namespacesTopicsSubscriptionGenerator = gen.Struct(reflect.TypeOf(NamespacesTopicsSubscription{}), generators)

	return namespacesTopicsSubscriptionGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesTopicsSubscription is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopicsSubscription(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesTopicsSubscription_SpecGenerator()
	gens["Status"] = NamespacesTopicsSubscription_STATUSGenerator()
}

func Test_NamespacesTopicsSubscriptionOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopicsSubscriptionOperatorSpec to NamespacesTopicsSubscriptionOperatorSpec via AssignProperties_To_NamespacesTopicsSubscriptionOperatorSpec & AssignProperties_From_NamespacesTopicsSubscriptionOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopicsSubscriptionOperatorSpec, NamespacesTopicsSubscriptionOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopicsSubscriptionOperatorSpec tests if a specific instance of NamespacesTopicsSubscriptionOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopicsSubscriptionOperatorSpec(subject NamespacesTopicsSubscriptionOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.NamespacesTopicsSubscriptionOperatorSpec
	err := copied.AssignProperties_To_NamespacesTopicsSubscriptionOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopicsSubscriptionOperatorSpec
	err = actual.AssignProperties_From_NamespacesTopicsSubscriptionOperatorSpec(&other)
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

func Test_NamespacesTopicsSubscriptionOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopicsSubscriptionOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopicsSubscriptionOperatorSpec, NamespacesTopicsSubscriptionOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopicsSubscriptionOperatorSpec runs a test to see if a specific instance of NamespacesTopicsSubscriptionOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopicsSubscriptionOperatorSpec(subject NamespacesTopicsSubscriptionOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopicsSubscriptionOperatorSpec
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

// Generator of NamespacesTopicsSubscriptionOperatorSpec instances for property testing - lazily instantiated by
// NamespacesTopicsSubscriptionOperatorSpecGenerator()
var namespacesTopicsSubscriptionOperatorSpecGenerator gopter.Gen

// NamespacesTopicsSubscriptionOperatorSpecGenerator returns a generator of NamespacesTopicsSubscriptionOperatorSpec instances for property testing.
func NamespacesTopicsSubscriptionOperatorSpecGenerator() gopter.Gen {
	if namespacesTopicsSubscriptionOperatorSpecGenerator != nil {
		return namespacesTopicsSubscriptionOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	namespacesTopicsSubscriptionOperatorSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesTopicsSubscriptionOperatorSpec{}), generators)

	return namespacesTopicsSubscriptionOperatorSpecGenerator
}

func Test_NamespacesTopicsSubscription_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopicsSubscription_STATUS to NamespacesTopicsSubscription_STATUS via AssignProperties_To_NamespacesTopicsSubscription_STATUS & AssignProperties_From_NamespacesTopicsSubscription_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopicsSubscription_STATUS, NamespacesTopicsSubscription_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopicsSubscription_STATUS tests if a specific instance of NamespacesTopicsSubscription_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopicsSubscription_STATUS(subject NamespacesTopicsSubscription_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.NamespacesTopicsSubscription_STATUS
	err := copied.AssignProperties_To_NamespacesTopicsSubscription_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopicsSubscription_STATUS
	err = actual.AssignProperties_From_NamespacesTopicsSubscription_STATUS(&other)
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

func Test_NamespacesTopicsSubscription_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopicsSubscription_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopicsSubscription_STATUS, NamespacesTopicsSubscription_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopicsSubscription_STATUS runs a test to see if a specific instance of NamespacesTopicsSubscription_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopicsSubscription_STATUS(subject NamespacesTopicsSubscription_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopicsSubscription_STATUS
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

// Generator of NamespacesTopicsSubscription_STATUS instances for property testing - lazily instantiated by
// NamespacesTopicsSubscription_STATUSGenerator()
var namespacesTopicsSubscription_STATUSGenerator gopter.Gen

// NamespacesTopicsSubscription_STATUSGenerator returns a generator of NamespacesTopicsSubscription_STATUS instances for property testing.
// We first initialize namespacesTopicsSubscription_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesTopicsSubscription_STATUSGenerator() gopter.Gen {
	if namespacesTopicsSubscription_STATUSGenerator != nil {
		return namespacesTopicsSubscription_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopicsSubscription_STATUS(generators)
	namespacesTopicsSubscription_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesTopicsSubscription_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopicsSubscription_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespacesTopicsSubscription_STATUS(generators)
	namespacesTopicsSubscription_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesTopicsSubscription_STATUS{}), generators)

	return namespacesTopicsSubscription_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesTopicsSubscription_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesTopicsSubscription_STATUS(gens map[string]gopter.Gen) {
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DeadLetteringOnFilterEvaluationExceptions"] = gen.PtrOf(gen.Bool())
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsClientAffine"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MessageCount"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EntityStatus_STATUS_Active,
		EntityStatus_STATUS_Creating,
		EntityStatus_STATUS_Deleting,
		EntityStatus_STATUS_Disabled,
		EntityStatus_STATUS_ReceiveDisabled,
		EntityStatus_STATUS_Renaming,
		EntityStatus_STATUS_Restoring,
		EntityStatus_STATUS_SendDisabled,
		EntityStatus_STATUS_Unknown))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesTopicsSubscription_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopicsSubscription_STATUS(gens map[string]gopter.Gen) {
	gens["ClientAffineProperties"] = gen.PtrOf(SBClientAffineProperties_STATUSGenerator())
	gens["CountDetails"] = gen.PtrOf(MessageCountDetails_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_NamespacesTopicsSubscription_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopicsSubscription_Spec to NamespacesTopicsSubscription_Spec via AssignProperties_To_NamespacesTopicsSubscription_Spec & AssignProperties_From_NamespacesTopicsSubscription_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopicsSubscription_Spec, NamespacesTopicsSubscription_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopicsSubscription_Spec tests if a specific instance of NamespacesTopicsSubscription_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopicsSubscription_Spec(subject NamespacesTopicsSubscription_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.NamespacesTopicsSubscription_Spec
	err := copied.AssignProperties_To_NamespacesTopicsSubscription_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopicsSubscription_Spec
	err = actual.AssignProperties_From_NamespacesTopicsSubscription_Spec(&other)
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

func Test_NamespacesTopicsSubscription_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopicsSubscription_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopicsSubscription_Spec, NamespacesTopicsSubscription_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopicsSubscription_Spec runs a test to see if a specific instance of NamespacesTopicsSubscription_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopicsSubscription_Spec(subject NamespacesTopicsSubscription_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopicsSubscription_Spec
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

// Generator of NamespacesTopicsSubscription_Spec instances for property testing - lazily instantiated by
// NamespacesTopicsSubscription_SpecGenerator()
var namespacesTopicsSubscription_SpecGenerator gopter.Gen

// NamespacesTopicsSubscription_SpecGenerator returns a generator of NamespacesTopicsSubscription_Spec instances for property testing.
// We first initialize namespacesTopicsSubscription_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesTopicsSubscription_SpecGenerator() gopter.Gen {
	if namespacesTopicsSubscription_SpecGenerator != nil {
		return namespacesTopicsSubscription_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopicsSubscription_Spec(generators)
	namespacesTopicsSubscription_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesTopicsSubscription_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopicsSubscription_Spec(generators)
	AddRelatedPropertyGeneratorsForNamespacesTopicsSubscription_Spec(generators)
	namespacesTopicsSubscription_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesTopicsSubscription_Spec{}), generators)

	return namespacesTopicsSubscription_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesTopicsSubscription_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesTopicsSubscription_Spec(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["DeadLetteringOnFilterEvaluationExceptions"] = gen.PtrOf(gen.Bool())
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["IsClientAffine"] = gen.PtrOf(gen.Bool())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForNamespacesTopicsSubscription_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopicsSubscription_Spec(gens map[string]gopter.Gen) {
	gens["ClientAffineProperties"] = gen.PtrOf(SBClientAffinePropertiesGenerator())
	gens["OperatorSpec"] = gen.PtrOf(NamespacesTopicsSubscriptionOperatorSpecGenerator())
}

func Test_SBClientAffineProperties_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SBClientAffineProperties to SBClientAffineProperties via AssignProperties_To_SBClientAffineProperties & AssignProperties_From_SBClientAffineProperties returns original",
		prop.ForAll(RunPropertyAssignmentTestForSBClientAffineProperties, SBClientAffinePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSBClientAffineProperties tests if a specific instance of SBClientAffineProperties can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSBClientAffineProperties(subject SBClientAffineProperties) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SBClientAffineProperties
	err := copied.AssignProperties_To_SBClientAffineProperties(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SBClientAffineProperties
	err = actual.AssignProperties_From_SBClientAffineProperties(&other)
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

func Test_SBClientAffineProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBClientAffineProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBClientAffineProperties, SBClientAffinePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBClientAffineProperties runs a test to see if a specific instance of SBClientAffineProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForSBClientAffineProperties(subject SBClientAffineProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBClientAffineProperties
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

// Generator of SBClientAffineProperties instances for property testing - lazily instantiated by
// SBClientAffinePropertiesGenerator()
var sbClientAffinePropertiesGenerator gopter.Gen

// SBClientAffinePropertiesGenerator returns a generator of SBClientAffineProperties instances for property testing.
func SBClientAffinePropertiesGenerator() gopter.Gen {
	if sbClientAffinePropertiesGenerator != nil {
		return sbClientAffinePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBClientAffineProperties(generators)
	sbClientAffinePropertiesGenerator = gen.Struct(reflect.TypeOf(SBClientAffineProperties{}), generators)

	return sbClientAffinePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForSBClientAffineProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBClientAffineProperties(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["IsDurable"] = gen.PtrOf(gen.Bool())
	gens["IsShared"] = gen.PtrOf(gen.Bool())
}

func Test_SBClientAffineProperties_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SBClientAffineProperties_STATUS to SBClientAffineProperties_STATUS via AssignProperties_To_SBClientAffineProperties_STATUS & AssignProperties_From_SBClientAffineProperties_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSBClientAffineProperties_STATUS, SBClientAffineProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSBClientAffineProperties_STATUS tests if a specific instance of SBClientAffineProperties_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSBClientAffineProperties_STATUS(subject SBClientAffineProperties_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SBClientAffineProperties_STATUS
	err := copied.AssignProperties_To_SBClientAffineProperties_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SBClientAffineProperties_STATUS
	err = actual.AssignProperties_From_SBClientAffineProperties_STATUS(&other)
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

func Test_SBClientAffineProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBClientAffineProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBClientAffineProperties_STATUS, SBClientAffineProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBClientAffineProperties_STATUS runs a test to see if a specific instance of SBClientAffineProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSBClientAffineProperties_STATUS(subject SBClientAffineProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBClientAffineProperties_STATUS
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

// Generator of SBClientAffineProperties_STATUS instances for property testing - lazily instantiated by
// SBClientAffineProperties_STATUSGenerator()
var sbClientAffineProperties_STATUSGenerator gopter.Gen

// SBClientAffineProperties_STATUSGenerator returns a generator of SBClientAffineProperties_STATUS instances for property testing.
func SBClientAffineProperties_STATUSGenerator() gopter.Gen {
	if sbClientAffineProperties_STATUSGenerator != nil {
		return sbClientAffineProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBClientAffineProperties_STATUS(generators)
	sbClientAffineProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(SBClientAffineProperties_STATUS{}), generators)

	return sbClientAffineProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSBClientAffineProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBClientAffineProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["IsDurable"] = gen.PtrOf(gen.Bool())
	gens["IsShared"] = gen.PtrOf(gen.Bool())
}
