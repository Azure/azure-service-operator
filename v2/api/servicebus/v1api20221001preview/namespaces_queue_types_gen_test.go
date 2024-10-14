// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20221001preview

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/storage"
	v20221001ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview/storage"
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

func Test_MessageCountDetails_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from MessageCountDetails_STATUS to MessageCountDetails_STATUS via AssignProperties_To_MessageCountDetails_STATUS & AssignProperties_From_MessageCountDetails_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForMessageCountDetails_STATUS, MessageCountDetails_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForMessageCountDetails_STATUS tests if a specific instance of MessageCountDetails_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForMessageCountDetails_STATUS(subject MessageCountDetails_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20221001ps.MessageCountDetails_STATUS
	err := copied.AssignProperties_To_MessageCountDetails_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual MessageCountDetails_STATUS
	err = actual.AssignProperties_From_MessageCountDetails_STATUS(&other)
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

func Test_MessageCountDetails_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MessageCountDetails_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMessageCountDetails_STATUS, MessageCountDetails_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMessageCountDetails_STATUS runs a test to see if a specific instance of MessageCountDetails_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMessageCountDetails_STATUS(subject MessageCountDetails_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MessageCountDetails_STATUS
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

// Generator of MessageCountDetails_STATUS instances for property testing - lazily instantiated by
// MessageCountDetails_STATUSGenerator()
var messageCountDetails_STATUSGenerator gopter.Gen

// MessageCountDetails_STATUSGenerator returns a generator of MessageCountDetails_STATUS instances for property testing.
func MessageCountDetails_STATUSGenerator() gopter.Gen {
	if messageCountDetails_STATUSGenerator != nil {
		return messageCountDetails_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMessageCountDetails_STATUS(generators)
	messageCountDetails_STATUSGenerator = gen.Struct(reflect.TypeOf(MessageCountDetails_STATUS{}), generators)

	return messageCountDetails_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMessageCountDetails_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMessageCountDetails_STATUS(gens map[string]gopter.Gen) {
	gens["ActiveMessageCount"] = gen.PtrOf(gen.Int())
	gens["DeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["ScheduledMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferDeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferMessageCount"] = gen.PtrOf(gen.Int())
}

func Test_NamespacesQueue_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesQueue to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesQueue, NamespacesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesQueue tests if a specific instance of NamespacesQueue round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesQueue(subject NamespacesQueue) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.NamespacesQueue
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesQueue
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

func Test_NamespacesQueue_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesQueue to NamespacesQueue via AssignProperties_To_NamespacesQueue & AssignProperties_From_NamespacesQueue returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesQueue, NamespacesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesQueue tests if a specific instance of NamespacesQueue can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesQueue(subject NamespacesQueue) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20221001ps.NamespacesQueue
	err := copied.AssignProperties_To_NamespacesQueue(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesQueue
	err = actual.AssignProperties_From_NamespacesQueue(&other)
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

func Test_NamespacesQueue_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesQueue via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesQueue, NamespacesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesQueue runs a test to see if a specific instance of NamespacesQueue round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesQueue(subject NamespacesQueue) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesQueue
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

// Generator of NamespacesQueue instances for property testing - lazily instantiated by NamespacesQueueGenerator()
var namespacesQueueGenerator gopter.Gen

// NamespacesQueueGenerator returns a generator of NamespacesQueue instances for property testing.
func NamespacesQueueGenerator() gopter.Gen {
	if namespacesQueueGenerator != nil {
		return namespacesQueueGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesQueue(generators)
	namespacesQueueGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue{}), generators)

	return namespacesQueueGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesQueue is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesQueue(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesQueue_SpecGenerator()
	gens["Status"] = NamespacesQueue_STATUSGenerator()
}

func Test_NamespacesQueue_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesQueue_STATUS to NamespacesQueue_STATUS via AssignProperties_To_NamespacesQueue_STATUS & AssignProperties_From_NamespacesQueue_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesQueue_STATUS, NamespacesQueue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesQueue_STATUS tests if a specific instance of NamespacesQueue_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesQueue_STATUS(subject NamespacesQueue_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20221001ps.NamespacesQueue_STATUS
	err := copied.AssignProperties_To_NamespacesQueue_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesQueue_STATUS
	err = actual.AssignProperties_From_NamespacesQueue_STATUS(&other)
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

func Test_NamespacesQueue_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesQueue_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesQueue_STATUS, NamespacesQueue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesQueue_STATUS runs a test to see if a specific instance of NamespacesQueue_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesQueue_STATUS(subject NamespacesQueue_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesQueue_STATUS
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

// Generator of NamespacesQueue_STATUS instances for property testing - lazily instantiated by
// NamespacesQueue_STATUSGenerator()
var namespacesQueue_STATUSGenerator gopter.Gen

// NamespacesQueue_STATUSGenerator returns a generator of NamespacesQueue_STATUS instances for property testing.
// We first initialize namespacesQueue_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesQueue_STATUSGenerator() gopter.Gen {
	if namespacesQueue_STATUSGenerator != nil {
		return namespacesQueue_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueue_STATUS(generators)
	namespacesQueue_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueue_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespacesQueue_STATUS(generators)
	namespacesQueue_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue_STATUS{}), generators)

	return namespacesQueue_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesQueue_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesQueue_STATUS(gens map[string]gopter.Gen) {
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxMessageSizeInKilobytes"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["MessageCount"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
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

// AddRelatedPropertyGeneratorsForNamespacesQueue_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesQueue_STATUS(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetails_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_NamespacesQueue_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesQueue_Spec to NamespacesQueue_Spec via AssignProperties_To_NamespacesQueue_Spec & AssignProperties_From_NamespacesQueue_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesQueue_Spec, NamespacesQueue_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesQueue_Spec tests if a specific instance of NamespacesQueue_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesQueue_Spec(subject NamespacesQueue_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20221001ps.NamespacesQueue_Spec
	err := copied.AssignProperties_To_NamespacesQueue_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesQueue_Spec
	err = actual.AssignProperties_From_NamespacesQueue_Spec(&other)
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

func Test_NamespacesQueue_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesQueue_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesQueue_Spec, NamespacesQueue_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesQueue_Spec runs a test to see if a specific instance of NamespacesQueue_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesQueue_Spec(subject NamespacesQueue_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesQueue_Spec
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

// Generator of NamespacesQueue_Spec instances for property testing - lazily instantiated by
// NamespacesQueue_SpecGenerator()
var namespacesQueue_SpecGenerator gopter.Gen

// NamespacesQueue_SpecGenerator returns a generator of NamespacesQueue_Spec instances for property testing.
func NamespacesQueue_SpecGenerator() gopter.Gen {
	if namespacesQueue_SpecGenerator != nil {
		return namespacesQueue_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueue_Spec(generators)
	namespacesQueue_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue_Spec{}), generators)

	return namespacesQueue_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesQueue_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesQueue_Spec(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxMessageSizeInKilobytes"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
}
