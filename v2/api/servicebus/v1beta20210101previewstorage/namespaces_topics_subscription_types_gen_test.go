// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101previewstorage

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
	gens["Spec"] = Namespaces_Topics_Subscription_SpecGenerator()
	gens["Status"] = SBSubscription_STATUSGenerator()
}

func Test_Namespaces_Topics_Subscription_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Topics_Subscription_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Topics_Subscription_Spec, Namespaces_Topics_Subscription_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Topics_Subscription_Spec runs a test to see if a specific instance of Namespaces_Topics_Subscription_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Topics_Subscription_Spec(subject Namespaces_Topics_Subscription_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Topics_Subscription_Spec
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

// Generator of Namespaces_Topics_Subscription_Spec instances for property testing - lazily instantiated by
// Namespaces_Topics_Subscription_SpecGenerator()
var namespaces_Topics_Subscription_SpecGenerator gopter.Gen

// Namespaces_Topics_Subscription_SpecGenerator returns a generator of Namespaces_Topics_Subscription_Spec instances for property testing.
func Namespaces_Topics_Subscription_SpecGenerator() gopter.Gen {
	if namespaces_Topics_Subscription_SpecGenerator != nil {
		return namespaces_Topics_Subscription_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topics_Subscription_Spec(generators)
	namespaces_Topics_Subscription_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topics_Subscription_Spec{}), generators)

	return namespaces_Topics_Subscription_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Topics_Subscription_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Topics_Subscription_Spec(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["DeadLetteringOnFilterEvaluationExceptions"] = gen.PtrOf(gen.Bool())
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_SBSubscription_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBSubscription_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBSubscription_STATUS, SBSubscription_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBSubscription_STATUS runs a test to see if a specific instance of SBSubscription_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSBSubscription_STATUS(subject SBSubscription_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBSubscription_STATUS
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

// Generator of SBSubscription_STATUS instances for property testing - lazily instantiated by
// SBSubscription_STATUSGenerator()
var sbSubscription_STATUSGenerator gopter.Gen

// SBSubscription_STATUSGenerator returns a generator of SBSubscription_STATUS instances for property testing.
// We first initialize sbSubscription_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBSubscription_STATUSGenerator() gopter.Gen {
	if sbSubscription_STATUSGenerator != nil {
		return sbSubscription_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBSubscription_STATUS(generators)
	sbSubscription_STATUSGenerator = gen.Struct(reflect.TypeOf(SBSubscription_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBSubscription_STATUS(generators)
	AddRelatedPropertyGeneratorsForSBSubscription_STATUS(generators)
	sbSubscription_STATUSGenerator = gen.Struct(reflect.TypeOf(SBSubscription_STATUS{}), generators)

	return sbSubscription_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSBSubscription_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBSubscription_STATUS(gens map[string]gopter.Gen) {
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
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MessageCount"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBSubscription_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBSubscription_STATUS(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetails_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
