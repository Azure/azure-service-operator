// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

import (
	"encoding/json"
	v20210101ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1beta20210101previewstorage"
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
	var hub v20210101ps.NamespacesQueue
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
		"Round trip from NamespacesQueue to NamespacesQueue via AssignPropertiesToNamespacesQueue & AssignPropertiesFromNamespacesQueue returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesQueue, NamespacesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesQueue tests if a specific instance of NamespacesQueue can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespacesQueue(subject NamespacesQueue) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.NamespacesQueue
	err := copied.AssignPropertiesToNamespacesQueue(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesQueue
	err = actual.AssignPropertiesFromNamespacesQueue(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
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
	gens["Spec"] = NamespacesQueuesSpecGenerator()
	gens["Status"] = SBQueueSTATUSGenerator()
}

func Test_NamespacesQueues_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesQueues_Spec to NamespacesQueues_Spec via AssignPropertiesToNamespacesQueuesSpec & AssignPropertiesFromNamespacesQueuesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesQueuesSpec, NamespacesQueuesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesQueuesSpec tests if a specific instance of NamespacesQueues_Spec can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespacesQueuesSpec(subject NamespacesQueues_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.NamespacesQueues_Spec
	err := copied.AssignPropertiesToNamespacesQueuesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesQueues_Spec
	err = actual.AssignPropertiesFromNamespacesQueuesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesQueues_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesQueues_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesQueuesSpec, NamespacesQueuesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesQueuesSpec runs a test to see if a specific instance of NamespacesQueues_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesQueuesSpec(subject NamespacesQueues_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesQueues_Spec
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

// Generator of NamespacesQueues_Spec instances for property testing - lazily instantiated by
// NamespacesQueuesSpecGenerator()
var namespacesQueuesSpecGenerator gopter.Gen

// NamespacesQueuesSpecGenerator returns a generator of NamespacesQueues_Spec instances for property testing.
func NamespacesQueuesSpecGenerator() gopter.Gen {
	if namespacesQueuesSpecGenerator != nil {
		return namespacesQueuesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueuesSpec(generators)
	namespacesQueuesSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesQueues_Spec{}), generators)

	return namespacesQueuesSpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesQueuesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesQueuesSpec(gens map[string]gopter.Gen) {
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
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_SBQueue_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SBQueue_STATUS to SBQueue_STATUS via AssignPropertiesToSBQueueSTATUS & AssignPropertiesFromSBQueueSTATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSBQueueSTATUS, SBQueueSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSBQueueSTATUS tests if a specific instance of SBQueue_STATUS can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForSBQueueSTATUS(subject SBQueue_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.SBQueue_STATUS
	err := copied.AssignPropertiesToSBQueueSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SBQueue_STATUS
	err = actual.AssignPropertiesFromSBQueueSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SBQueue_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBQueue_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBQueueSTATUS, SBQueueSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBQueueSTATUS runs a test to see if a specific instance of SBQueue_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSBQueueSTATUS(subject SBQueue_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBQueue_STATUS
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

// Generator of SBQueue_STATUS instances for property testing - lazily instantiated by SBQueueSTATUSGenerator()
var sbQueueSTATUSGenerator gopter.Gen

// SBQueueSTATUSGenerator returns a generator of SBQueue_STATUS instances for property testing.
// We first initialize sbQueueSTATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBQueueSTATUSGenerator() gopter.Gen {
	if sbQueueSTATUSGenerator != nil {
		return sbQueueSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueueSTATUS(generators)
	sbQueueSTATUSGenerator = gen.Struct(reflect.TypeOf(SBQueue_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueueSTATUS(generators)
	AddRelatedPropertyGeneratorsForSBQueueSTATUS(generators)
	sbQueueSTATUSGenerator = gen.Struct(reflect.TypeOf(SBQueue_STATUS{}), generators)

	return sbQueueSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForSBQueueSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBQueueSTATUS(gens map[string]gopter.Gen) {
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
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
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

// AddRelatedPropertyGeneratorsForSBQueueSTATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBQueueSTATUS(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetailsSTATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataSTATUSGenerator())
}

func Test_MessageCountDetails_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from MessageCountDetails_STATUS to MessageCountDetails_STATUS via AssignPropertiesToMessageCountDetailsSTATUS & AssignPropertiesFromMessageCountDetailsSTATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForMessageCountDetailsSTATUS, MessageCountDetailsSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForMessageCountDetailsSTATUS tests if a specific instance of MessageCountDetails_STATUS can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForMessageCountDetailsSTATUS(subject MessageCountDetails_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.MessageCountDetails_STATUS
	err := copied.AssignPropertiesToMessageCountDetailsSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual MessageCountDetails_STATUS
	err = actual.AssignPropertiesFromMessageCountDetailsSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
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
		prop.ForAll(RunJSONSerializationTestForMessageCountDetailsSTATUS, MessageCountDetailsSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMessageCountDetailsSTATUS runs a test to see if a specific instance of MessageCountDetails_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMessageCountDetailsSTATUS(subject MessageCountDetails_STATUS) string {
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
// MessageCountDetailsSTATUSGenerator()
var messageCountDetailsSTATUSGenerator gopter.Gen

// MessageCountDetailsSTATUSGenerator returns a generator of MessageCountDetails_STATUS instances for property testing.
func MessageCountDetailsSTATUSGenerator() gopter.Gen {
	if messageCountDetailsSTATUSGenerator != nil {
		return messageCountDetailsSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMessageCountDetailsSTATUS(generators)
	messageCountDetailsSTATUSGenerator = gen.Struct(reflect.TypeOf(MessageCountDetails_STATUS{}), generators)

	return messageCountDetailsSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForMessageCountDetailsSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMessageCountDetailsSTATUS(gens map[string]gopter.Gen) {
	gens["ActiveMessageCount"] = gen.PtrOf(gen.Int())
	gens["DeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["ScheduledMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferDeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferMessageCount"] = gen.PtrOf(gen.Int())
}
