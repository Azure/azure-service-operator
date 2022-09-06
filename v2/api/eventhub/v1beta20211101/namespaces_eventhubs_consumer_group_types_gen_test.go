// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101storage"
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

func Test_NamespacesEventhubsConsumerGroup_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsConsumerGroup to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesEventhubsConsumerGroup, NamespacesEventhubsConsumerGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesEventhubsConsumerGroup tests if a specific instance of NamespacesEventhubsConsumerGroup round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesEventhubsConsumerGroup(subject NamespacesEventhubsConsumerGroup) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.NamespacesEventhubsConsumerGroup
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesEventhubsConsumerGroup
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

func Test_NamespacesEventhubsConsumerGroup_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsConsumerGroup to NamespacesEventhubsConsumerGroup via AssignProperties_To_NamespacesEventhubsConsumerGroup & AssignProperties_From_NamespacesEventhubsConsumerGroup returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup, NamespacesEventhubsConsumerGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup tests if a specific instance of NamespacesEventhubsConsumerGroup can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup(subject NamespacesEventhubsConsumerGroup) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.NamespacesEventhubsConsumerGroup
	err := copied.AssignProperties_To_NamespacesEventhubsConsumerGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubsConsumerGroup
	err = actual.AssignProperties_From_NamespacesEventhubsConsumerGroup(&other)
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

func Test_NamespacesEventhubsConsumerGroup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsConsumerGroup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsConsumerGroup, NamespacesEventhubsConsumerGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsConsumerGroup runs a test to see if a specific instance of NamespacesEventhubsConsumerGroup round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsConsumerGroup(subject NamespacesEventhubsConsumerGroup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsConsumerGroup
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

// Generator of NamespacesEventhubsConsumerGroup instances for property testing - lazily instantiated by
// NamespacesEventhubsConsumerGroupGenerator()
var namespacesEventhubsConsumerGroupGenerator gopter.Gen

// NamespacesEventhubsConsumerGroupGenerator returns a generator of NamespacesEventhubsConsumerGroup instances for property testing.
func NamespacesEventhubsConsumerGroupGenerator() gopter.Gen {
	if namespacesEventhubsConsumerGroupGenerator != nil {
		return namespacesEventhubsConsumerGroupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup(generators)
	namespacesEventhubsConsumerGroupGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumerGroup{}), generators)

	return namespacesEventhubsConsumerGroupGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup(gens map[string]gopter.Gen) {
	gens["Spec"] = Namespaces_Eventhubs_Consumergroups_SpecGenerator()
	gens["Status"] = ConsumerGroup_STATUSGenerator()
}

func Test_ConsumerGroup_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ConsumerGroup_STATUS to ConsumerGroup_STATUS via AssignProperties_To_ConsumerGroup_STATUS & AssignProperties_From_ConsumerGroup_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForConsumerGroup_STATUS, ConsumerGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForConsumerGroup_STATUS tests if a specific instance of ConsumerGroup_STATUS can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForConsumerGroup_STATUS(subject ConsumerGroup_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ConsumerGroup_STATUS
	err := copied.AssignProperties_To_ConsumerGroup_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ConsumerGroup_STATUS
	err = actual.AssignProperties_From_ConsumerGroup_STATUS(&other)
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

func Test_ConsumerGroup_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConsumerGroup_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConsumerGroup_STATUS, ConsumerGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConsumerGroup_STATUS runs a test to see if a specific instance of ConsumerGroup_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForConsumerGroup_STATUS(subject ConsumerGroup_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConsumerGroup_STATUS
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

// Generator of ConsumerGroup_STATUS instances for property testing - lazily instantiated by
// ConsumerGroup_STATUSGenerator()
var consumerGroup_STATUSGenerator gopter.Gen

// ConsumerGroup_STATUSGenerator returns a generator of ConsumerGroup_STATUS instances for property testing.
// We first initialize consumerGroup_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConsumerGroup_STATUSGenerator() gopter.Gen {
	if consumerGroup_STATUSGenerator != nil {
		return consumerGroup_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConsumerGroup_STATUS(generators)
	consumerGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(ConsumerGroup_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConsumerGroup_STATUS(generators)
	AddRelatedPropertyGeneratorsForConsumerGroup_STATUS(generators)
	consumerGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(ConsumerGroup_STATUS{}), generators)

	return consumerGroup_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForConsumerGroup_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConsumerGroup_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["UserMetadata"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForConsumerGroup_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConsumerGroup_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Namespaces_Eventhubs_Consumergroups_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Namespaces_Eventhubs_Consumergroups_Spec to Namespaces_Eventhubs_Consumergroups_Spec via AssignProperties_To_Namespaces_Eventhubs_Consumergroups_Spec & AssignProperties_From_Namespaces_Eventhubs_Consumergroups_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespaces_Eventhubs_Consumergroups_Spec, Namespaces_Eventhubs_Consumergroups_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespaces_Eventhubs_Consumergroups_Spec tests if a specific instance of Namespaces_Eventhubs_Consumergroups_Spec can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespaces_Eventhubs_Consumergroups_Spec(subject Namespaces_Eventhubs_Consumergroups_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Namespaces_Eventhubs_Consumergroups_Spec
	err := copied.AssignProperties_To_Namespaces_Eventhubs_Consumergroups_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Namespaces_Eventhubs_Consumergroups_Spec
	err = actual.AssignProperties_From_Namespaces_Eventhubs_Consumergroups_Spec(&other)
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

func Test_Namespaces_Eventhubs_Consumergroups_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhubs_Consumergroups_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhubs_Consumergroups_Spec, Namespaces_Eventhubs_Consumergroups_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhubs_Consumergroups_Spec runs a test to see if a specific instance of Namespaces_Eventhubs_Consumergroups_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhubs_Consumergroups_Spec(subject Namespaces_Eventhubs_Consumergroups_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhubs_Consumergroups_Spec
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

// Generator of Namespaces_Eventhubs_Consumergroups_Spec instances for property testing - lazily instantiated by
// Namespaces_Eventhubs_Consumergroups_SpecGenerator()
var namespaces_Eventhubs_Consumergroups_SpecGenerator gopter.Gen

// Namespaces_Eventhubs_Consumergroups_SpecGenerator returns a generator of Namespaces_Eventhubs_Consumergroups_Spec instances for property testing.
func Namespaces_Eventhubs_Consumergroups_SpecGenerator() gopter.Gen {
	if namespaces_Eventhubs_Consumergroups_SpecGenerator != nil {
		return namespaces_Eventhubs_Consumergroups_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Consumergroups_Spec(generators)
	namespaces_Eventhubs_Consumergroups_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhubs_Consumergroups_Spec{}), generators)

	return namespaces_Eventhubs_Consumergroups_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Consumergroups_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Consumergroups_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["UserMetadata"] = gen.PtrOf(gen.AlphaString())
}
