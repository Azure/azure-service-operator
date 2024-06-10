// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

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

func Test_Namespaces_Topics_Subscription_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Topics_Subscription_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Topics_Subscription_Spec_ARM, Namespaces_Topics_Subscription_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Topics_Subscription_Spec_ARM runs a test to see if a specific instance of Namespaces_Topics_Subscription_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Topics_Subscription_Spec_ARM(subject Namespaces_Topics_Subscription_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Topics_Subscription_Spec_ARM
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

// Generator of Namespaces_Topics_Subscription_Spec_ARM instances for property testing - lazily instantiated by
// Namespaces_Topics_Subscription_Spec_ARMGenerator()
var namespaces_Topics_Subscription_Spec_ARMGenerator gopter.Gen

// Namespaces_Topics_Subscription_Spec_ARMGenerator returns a generator of Namespaces_Topics_Subscription_Spec_ARM instances for property testing.
// We first initialize namespaces_Topics_Subscription_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Topics_Subscription_Spec_ARMGenerator() gopter.Gen {
	if namespaces_Topics_Subscription_Spec_ARMGenerator != nil {
		return namespaces_Topics_Subscription_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topics_Subscription_Spec_ARM(generators)
	namespaces_Topics_Subscription_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topics_Subscription_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topics_Subscription_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Topics_Subscription_Spec_ARM(generators)
	namespaces_Topics_Subscription_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topics_Subscription_Spec_ARM{}), generators)

	return namespaces_Topics_Subscription_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Topics_Subscription_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Topics_Subscription_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespaces_Topics_Subscription_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Topics_Subscription_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SBSubscriptionProperties_ARMGenerator())
}

func Test_SBClientAffineProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBClientAffineProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBClientAffineProperties_ARM, SBClientAffineProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBClientAffineProperties_ARM runs a test to see if a specific instance of SBClientAffineProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBClientAffineProperties_ARM(subject SBClientAffineProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBClientAffineProperties_ARM
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

// Generator of SBClientAffineProperties_ARM instances for property testing - lazily instantiated by
// SBClientAffineProperties_ARMGenerator()
var sbClientAffineProperties_ARMGenerator gopter.Gen

// SBClientAffineProperties_ARMGenerator returns a generator of SBClientAffineProperties_ARM instances for property testing.
func SBClientAffineProperties_ARMGenerator() gopter.Gen {
	if sbClientAffineProperties_ARMGenerator != nil {
		return sbClientAffineProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBClientAffineProperties_ARM(generators)
	sbClientAffineProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SBClientAffineProperties_ARM{}), generators)

	return sbClientAffineProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSBClientAffineProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBClientAffineProperties_ARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["IsDurable"] = gen.PtrOf(gen.Bool())
	gens["IsShared"] = gen.PtrOf(gen.Bool())
}

func Test_SBSubscriptionProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBSubscriptionProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBSubscriptionProperties_ARM, SBSubscriptionProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBSubscriptionProperties_ARM runs a test to see if a specific instance of SBSubscriptionProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBSubscriptionProperties_ARM(subject SBSubscriptionProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBSubscriptionProperties_ARM
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

// Generator of SBSubscriptionProperties_ARM instances for property testing - lazily instantiated by
// SBSubscriptionProperties_ARMGenerator()
var sbSubscriptionProperties_ARMGenerator gopter.Gen

// SBSubscriptionProperties_ARMGenerator returns a generator of SBSubscriptionProperties_ARM instances for property testing.
// We first initialize sbSubscriptionProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBSubscriptionProperties_ARMGenerator() gopter.Gen {
	if sbSubscriptionProperties_ARMGenerator != nil {
		return sbSubscriptionProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBSubscriptionProperties_ARM(generators)
	sbSubscriptionProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SBSubscriptionProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBSubscriptionProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForSBSubscriptionProperties_ARM(generators)
	sbSubscriptionProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SBSubscriptionProperties_ARM{}), generators)

	return sbSubscriptionProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSBSubscriptionProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBSubscriptionProperties_ARM(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
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

// AddRelatedPropertyGeneratorsForSBSubscriptionProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBSubscriptionProperties_ARM(gens map[string]gopter.Gen) {
	gens["ClientAffineProperties"] = gen.PtrOf(SBClientAffineProperties_ARMGenerator())
}
