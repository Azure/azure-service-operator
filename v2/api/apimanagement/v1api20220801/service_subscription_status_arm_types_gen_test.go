// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

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

func Test_Service_Subscription_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Subscription_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Subscription_STATUS_ARM, Service_Subscription_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Subscription_STATUS_ARM runs a test to see if a specific instance of Service_Subscription_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Subscription_STATUS_ARM(subject Service_Subscription_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Subscription_STATUS_ARM
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

// Generator of Service_Subscription_STATUS_ARM instances for property testing - lazily instantiated by
// Service_Subscription_STATUS_ARMGenerator()
var service_Subscription_STATUS_ARMGenerator gopter.Gen

// Service_Subscription_STATUS_ARMGenerator returns a generator of Service_Subscription_STATUS_ARM instances for property testing.
// We first initialize service_Subscription_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_Subscription_STATUS_ARMGenerator() gopter.Gen {
	if service_Subscription_STATUS_ARMGenerator != nil {
		return service_Subscription_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Subscription_STATUS_ARM(generators)
	service_Subscription_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_Subscription_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Subscription_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForService_Subscription_STATUS_ARM(generators)
	service_Subscription_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_Subscription_STATUS_ARM{}), generators)

	return service_Subscription_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForService_Subscription_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Subscription_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForService_Subscription_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_Subscription_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SubscriptionContractProperties_STATUS_ARMGenerator())
}

func Test_SubscriptionContractProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionContractProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionContractProperties_STATUS_ARM, SubscriptionContractProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionContractProperties_STATUS_ARM runs a test to see if a specific instance of SubscriptionContractProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionContractProperties_STATUS_ARM(subject SubscriptionContractProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionContractProperties_STATUS_ARM
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

// Generator of SubscriptionContractProperties_STATUS_ARM instances for property testing - lazily instantiated by
// SubscriptionContractProperties_STATUS_ARMGenerator()
var subscriptionContractProperties_STATUS_ARMGenerator gopter.Gen

// SubscriptionContractProperties_STATUS_ARMGenerator returns a generator of SubscriptionContractProperties_STATUS_ARM instances for property testing.
func SubscriptionContractProperties_STATUS_ARMGenerator() gopter.Gen {
	if subscriptionContractProperties_STATUS_ARMGenerator != nil {
		return subscriptionContractProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionContractProperties_STATUS_ARM(generators)
	subscriptionContractProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SubscriptionContractProperties_STATUS_ARM{}), generators)

	return subscriptionContractProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionContractProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionContractProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AllowTracing"] = gen.PtrOf(gen.Bool())
	gens["CreatedDate"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["EndDate"] = gen.PtrOf(gen.AlphaString())
	gens["ExpirationDate"] = gen.PtrOf(gen.AlphaString())
	gens["NotificationDate"] = gen.PtrOf(gen.AlphaString())
	gens["OwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["StartDate"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		SubscriptionContractProperties_State_STATUS_ARM_Active,
		SubscriptionContractProperties_State_STATUS_ARM_Cancelled,
		SubscriptionContractProperties_State_STATUS_ARM_Expired,
		SubscriptionContractProperties_State_STATUS_ARM_Rejected,
		SubscriptionContractProperties_State_STATUS_ARM_Submitted,
		SubscriptionContractProperties_State_STATUS_ARM_Suspended))
	gens["StateComment"] = gen.PtrOf(gen.AlphaString())
}
