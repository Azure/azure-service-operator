// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101previewstorage

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

func Test_NamespacesTopic_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopic to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesTopic, NamespacesTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesTopic tests if a specific instance of NamespacesTopic round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesTopic(subject NamespacesTopic) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210101ps.NamespacesTopic
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesTopic
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

func Test_NamespacesTopic_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopic to NamespacesTopic via AssignProperties_To_NamespacesTopic & AssignProperties_From_NamespacesTopic returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopic, NamespacesTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopic tests if a specific instance of NamespacesTopic can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopic(subject NamespacesTopic) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.NamespacesTopic
	err := copied.AssignProperties_To_NamespacesTopic(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopic
	err = actual.AssignProperties_From_NamespacesTopic(&other)
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

func Test_NamespacesTopic_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopic via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopic, NamespacesTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopic runs a test to see if a specific instance of NamespacesTopic round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopic(subject NamespacesTopic) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopic
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

// Generator of NamespacesTopic instances for property testing - lazily instantiated by NamespacesTopicGenerator()
var namespacesTopicGenerator gopter.Gen

// NamespacesTopicGenerator returns a generator of NamespacesTopic instances for property testing.
func NamespacesTopicGenerator() gopter.Gen {
	if namespacesTopicGenerator != nil {
		return namespacesTopicGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesTopic(generators)
	namespacesTopicGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic{}), generators)

	return namespacesTopicGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesTopic is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopic(gens map[string]gopter.Gen) {
<<<<<<< HEAD
	gens["Spec"] = NamespacesTopic_SpecGenerator()
	gens["Status"] = NamespacesTopic_STATUSGenerator()
}

func Test_NamespacesTopic_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
=======
	gens["Spec"] = Namespaces_Topics_SpecGenerator()
	gens["Status"] = SBTopic_STATUSGenerator()
}

func Test_Namespaces_Topics_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
>>>>>>> main
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip from NamespacesTopic_Spec to NamespacesTopic_Spec via AssignPropertiesToNamespacesTopic_Spec & AssignPropertiesFromNamespacesTopic_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopic_Spec, NamespacesTopic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopic_Spec tests if a specific instance of NamespacesTopic_Spec can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopic_Spec(subject NamespacesTopic_Spec) string {
=======
		"Round trip from Namespaces_Topics_Spec to Namespaces_Topics_Spec via AssignProperties_To_Namespaces_Topics_Spec & AssignProperties_From_Namespaces_Topics_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespaces_Topics_Spec, Namespaces_Topics_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespaces_Topics_Spec tests if a specific instance of Namespaces_Topics_Spec can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespaces_Topics_Spec(subject Namespaces_Topics_Spec) string {
>>>>>>> main
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
<<<<<<< HEAD
	var other v20210101ps.NamespacesTopic_Spec
	err := copied.AssignPropertiesToNamespacesTopic_Spec(&other)
=======
	var other v20210101ps.Namespaces_Topics_Spec
	err := copied.AssignProperties_To_Namespaces_Topics_Spec(&other)
>>>>>>> main
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
<<<<<<< HEAD
	var actual NamespacesTopic_Spec
	err = actual.AssignPropertiesFromNamespacesTopic_Spec(&other)
=======
	var actual Namespaces_Topics_Spec
	err = actual.AssignProperties_From_Namespaces_Topics_Spec(&other)
>>>>>>> main
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

<<<<<<< HEAD
func Test_NamespacesTopic_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
=======
func Test_Namespaces_Topics_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
>>>>>>> main
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip of NamespacesTopic_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopic_Spec, NamespacesTopic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopic_Spec runs a test to see if a specific instance of NamespacesTopic_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopic_Spec(subject NamespacesTopic_Spec) string {
=======
		"Round trip of Namespaces_Topics_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Topics_Spec, Namespaces_Topics_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Topics_Spec runs a test to see if a specific instance of Namespaces_Topics_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Topics_Spec(subject Namespaces_Topics_Spec) string {
>>>>>>> main
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
<<<<<<< HEAD
	var actual NamespacesTopic_Spec
=======
	var actual Namespaces_Topics_Spec
>>>>>>> main
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

<<<<<<< HEAD
// Generator of NamespacesTopic_Spec instances for property testing - lazily instantiated by
// NamespacesTopic_SpecGenerator()
var namespacesTopic_SpecGenerator gopter.Gen

// NamespacesTopic_SpecGenerator returns a generator of NamespacesTopic_Spec instances for property testing.
func NamespacesTopic_SpecGenerator() gopter.Gen {
	if namespacesTopic_SpecGenerator != nil {
		return namespacesTopic_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopic_Spec(generators)
	namespacesTopic_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic_Spec{}), generators)

	return namespacesTopic_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesTopic_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesTopic_Spec(gens map[string]gopter.Gen) {
=======
// Generator of Namespaces_Topics_Spec instances for property testing - lazily instantiated by
// Namespaces_Topics_SpecGenerator()
var namespaces_Topics_SpecGenerator gopter.Gen

// Namespaces_Topics_SpecGenerator returns a generator of Namespaces_Topics_Spec instances for property testing.
func Namespaces_Topics_SpecGenerator() gopter.Gen {
	if namespaces_Topics_SpecGenerator != nil {
		return namespaces_Topics_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topics_Spec(generators)
	namespaces_Topics_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topics_Spec{}), generators)

	return namespaces_Topics_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Topics_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Topics_Spec(gens map[string]gopter.Gen) {
>>>>>>> main
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
}

func Test_NamespacesTopic_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip from NamespacesTopic_STATUS to NamespacesTopic_STATUS via AssignPropertiesToNamespacesTopic_STATUS & AssignPropertiesFromNamespacesTopic_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopic_STATUS, NamespacesTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopic_STATUS tests if a specific instance of NamespacesTopic_STATUS can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopic_STATUS(subject NamespacesTopic_STATUS) string {
=======
		"Round trip from SBTopic_STATUS to SBTopic_STATUS via AssignProperties_To_SBTopic_STATUS & AssignProperties_From_SBTopic_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSBTopic_STATUS, SBTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSBTopic_STATUS tests if a specific instance of SBTopic_STATUS can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForSBTopic_STATUS(subject SBTopic_STATUS) string {
>>>>>>> main
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
<<<<<<< HEAD
	var other v20210101ps.NamespacesTopic_STATUS
	err := copied.AssignPropertiesToNamespacesTopic_STATUS(&other)
=======
	var other v20210101ps.SBTopic_STATUS
	err := copied.AssignProperties_To_SBTopic_STATUS(&other)
>>>>>>> main
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
<<<<<<< HEAD
	var actual NamespacesTopic_STATUS
	err = actual.AssignPropertiesFromNamespacesTopic_STATUS(&other)
=======
	var actual SBTopic_STATUS
	err = actual.AssignProperties_From_SBTopic_STATUS(&other)
>>>>>>> main
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

func Test_NamespacesTopic_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip of NamespacesTopic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopic_STATUS, NamespacesTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopic_STATUS runs a test to see if a specific instance of NamespacesTopic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopic_STATUS(subject NamespacesTopic_STATUS) string {
=======
		"Round trip of SBTopic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBTopic_STATUS, SBTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBTopic_STATUS runs a test to see if a specific instance of SBTopic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSBTopic_STATUS(subject SBTopic_STATUS) string {
>>>>>>> main
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopic_STATUS
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

<<<<<<< HEAD
// Generator of NamespacesTopic_STATUS instances for property testing - lazily instantiated by
// NamespacesTopic_STATUSGenerator()
var namespacesTopic_STATUSGenerator gopter.Gen

// NamespacesTopic_STATUSGenerator returns a generator of NamespacesTopic_STATUS instances for property testing.
// We first initialize namespacesTopic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesTopic_STATUSGenerator() gopter.Gen {
	if namespacesTopic_STATUSGenerator != nil {
		return namespacesTopic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS(generators)
	namespacesTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespacesTopic_STATUS(generators)
	namespacesTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic_STATUS{}), generators)

	return namespacesTopic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS(gens map[string]gopter.Gen) {
=======
// Generator of SBTopic_STATUS instances for property testing - lazily instantiated by SBTopic_STATUSGenerator()
var sbTopic_STATUSGenerator gopter.Gen

// SBTopic_STATUSGenerator returns a generator of SBTopic_STATUS instances for property testing.
// We first initialize sbTopic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBTopic_STATUSGenerator() gopter.Gen {
	if sbTopic_STATUSGenerator != nil {
		return sbTopic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopic_STATUS(generators)
	sbTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(SBTopic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopic_STATUS(generators)
	AddRelatedPropertyGeneratorsForSBTopic_STATUS(generators)
	sbTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(SBTopic_STATUS{}), generators)

	return sbTopic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSBTopic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBTopic_STATUS(gens map[string]gopter.Gen) {
>>>>>>> main
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionCount"] = gen.PtrOf(gen.Int())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

<<<<<<< HEAD
// AddRelatedPropertyGeneratorsForNamespacesTopic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopic_STATUS(gens map[string]gopter.Gen) {
=======
// AddRelatedPropertyGeneratorsForSBTopic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBTopic_STATUS(gens map[string]gopter.Gen) {
>>>>>>> main
	gens["CountDetails"] = gen.PtrOf(MessageCountDetails_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
