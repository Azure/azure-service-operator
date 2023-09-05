// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101previewstorage

import (
	"encoding/json"
	v20210101ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20210101previewstorage"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101storage"
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

// RunPropertyAssignmentTestForNamespacesQueue tests if a specific instance of NamespacesQueue can be assigned to v1api20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespacesQueue(subject NamespacesQueue) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.NamespacesQueue
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
	gens["Spec"] = Namespaces_Queue_SpecGenerator()
	gens["Status"] = Namespaces_Queue_STATUSGenerator()
}

func Test_Namespaces_Queue_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Namespaces_Queue_Spec to Namespaces_Queue_Spec via AssignProperties_To_Namespaces_Queue_Spec & AssignProperties_From_Namespaces_Queue_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespaces_Queue_Spec, Namespaces_Queue_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespaces_Queue_Spec tests if a specific instance of Namespaces_Queue_Spec can be assigned to v1api20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespaces_Queue_Spec(subject Namespaces_Queue_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.Namespaces_Queue_Spec
	err := copied.AssignProperties_To_Namespaces_Queue_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Namespaces_Queue_Spec
	err = actual.AssignProperties_From_Namespaces_Queue_Spec(&other)
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

func Test_Namespaces_Queue_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Queue_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Queue_Spec, Namespaces_Queue_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Queue_Spec runs a test to see if a specific instance of Namespaces_Queue_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Queue_Spec(subject Namespaces_Queue_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Queue_Spec
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

// Generator of Namespaces_Queue_Spec instances for property testing - lazily instantiated by
// Namespaces_Queue_SpecGenerator()
var namespaces_Queue_SpecGenerator gopter.Gen

// Namespaces_Queue_SpecGenerator returns a generator of Namespaces_Queue_Spec instances for property testing.
func Namespaces_Queue_SpecGenerator() gopter.Gen {
	if namespaces_Queue_SpecGenerator != nil {
		return namespaces_Queue_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Queue_Spec(generators)
	namespaces_Queue_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_Queue_Spec{}), generators)

	return namespaces_Queue_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Queue_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Queue_Spec(gens map[string]gopter.Gen) {
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
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
}

func Test_Namespaces_Queue_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Namespaces_Queue_STATUS to Namespaces_Queue_STATUS via AssignProperties_To_Namespaces_Queue_STATUS & AssignProperties_From_Namespaces_Queue_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespaces_Queue_STATUS, Namespaces_Queue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespaces_Queue_STATUS tests if a specific instance of Namespaces_Queue_STATUS can be assigned to v1api20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespaces_Queue_STATUS(subject Namespaces_Queue_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.Namespaces_Queue_STATUS
	err := copied.AssignProperties_To_Namespaces_Queue_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Namespaces_Queue_STATUS
	err = actual.AssignProperties_From_Namespaces_Queue_STATUS(&other)
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

func Test_Namespaces_Queue_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Queue_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Queue_STATUS, Namespaces_Queue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Queue_STATUS runs a test to see if a specific instance of Namespaces_Queue_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Queue_STATUS(subject Namespaces_Queue_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Queue_STATUS
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

// Generator of Namespaces_Queue_STATUS instances for property testing - lazily instantiated by
// Namespaces_Queue_STATUSGenerator()
var namespaces_Queue_STATUSGenerator gopter.Gen

// Namespaces_Queue_STATUSGenerator returns a generator of Namespaces_Queue_STATUS instances for property testing.
// We first initialize namespaces_Queue_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Queue_STATUSGenerator() gopter.Gen {
	if namespaces_Queue_STATUSGenerator != nil {
		return namespaces_Queue_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Queue_STATUS(generators)
	namespaces_Queue_STATUSGenerator = gen.Struct(reflect.TypeOf(Namespaces_Queue_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Queue_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Queue_STATUS(generators)
	namespaces_Queue_STATUSGenerator = gen.Struct(reflect.TypeOf(Namespaces_Queue_STATUS{}), generators)

	return namespaces_Queue_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Queue_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Queue_STATUS(gens map[string]gopter.Gen) {
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
	gens["Status"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespaces_Queue_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Queue_STATUS(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetails_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

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

// RunPropertyAssignmentTestForMessageCountDetails_STATUS tests if a specific instance of MessageCountDetails_STATUS can be assigned to v1api20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForMessageCountDetails_STATUS(subject MessageCountDetails_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.MessageCountDetails_STATUS
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
