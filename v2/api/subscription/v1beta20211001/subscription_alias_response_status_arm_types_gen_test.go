// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

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

func Test_SubscriptionAliasResponse_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionAliasResponse_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionAliasResponse_STATUS_ARM, SubscriptionAliasResponse_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionAliasResponse_STATUS_ARM runs a test to see if a specific instance of SubscriptionAliasResponse_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionAliasResponse_STATUS_ARM(subject SubscriptionAliasResponse_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionAliasResponse_STATUS_ARM
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

// Generator of SubscriptionAliasResponse_STATUS_ARM instances for property testing - lazily instantiated by
// SubscriptionAliasResponse_STATUS_ARMGenerator()
var subscriptionAliasResponse_STATUS_ARMGenerator gopter.Gen

// SubscriptionAliasResponse_STATUS_ARMGenerator returns a generator of SubscriptionAliasResponse_STATUS_ARM instances for property testing.
// We first initialize subscriptionAliasResponse_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SubscriptionAliasResponse_STATUS_ARMGenerator() gopter.Gen {
	if subscriptionAliasResponse_STATUS_ARMGenerator != nil {
		return subscriptionAliasResponse_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponse_STATUS_ARM(generators)
	subscriptionAliasResponse_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponse_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponse_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForSubscriptionAliasResponse_STATUS_ARM(generators)
	subscriptionAliasResponse_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponse_STATUS_ARM{}), generators)

	return subscriptionAliasResponse_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionAliasResponse_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionAliasResponse_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSubscriptionAliasResponse_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSubscriptionAliasResponse_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SubscriptionAliasResponseProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_SubscriptionAliasResponseProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionAliasResponseProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionAliasResponseProperties_STATUS_ARM, SubscriptionAliasResponseProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionAliasResponseProperties_STATUS_ARM runs a test to see if a specific instance of SubscriptionAliasResponseProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionAliasResponseProperties_STATUS_ARM(subject SubscriptionAliasResponseProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionAliasResponseProperties_STATUS_ARM
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

// Generator of SubscriptionAliasResponseProperties_STATUS_ARM instances for property testing - lazily instantiated by
// SubscriptionAliasResponseProperties_STATUS_ARMGenerator()
var subscriptionAliasResponseProperties_STATUS_ARMGenerator gopter.Gen

// SubscriptionAliasResponseProperties_STATUS_ARMGenerator returns a generator of SubscriptionAliasResponseProperties_STATUS_ARM instances for property testing.
func SubscriptionAliasResponseProperties_STATUS_ARMGenerator() gopter.Gen {
	if subscriptionAliasResponseProperties_STATUS_ARMGenerator != nil {
		return subscriptionAliasResponseProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponseProperties_STATUS_ARM(generators)
	subscriptionAliasResponseProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponseProperties_STATUS_ARM{}), generators)

	return subscriptionAliasResponseProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionAliasResponseProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionAliasResponseProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AcceptOwnershipState"] = gen.PtrOf(gen.OneConstOf(AcceptOwnershipState_STATUS_Completed, AcceptOwnershipState_STATUS_Expired, AcceptOwnershipState_STATUS_Pending))
	gens["AcceptOwnershipUrl"] = gen.PtrOf(gen.AlphaString())
	gens["BillingScope"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["ManagementGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(SubscriptionAliasResponseProperties_ProvisioningState_STATUS_Accepted, SubscriptionAliasResponseProperties_ProvisioningState_STATUS_Failed, SubscriptionAliasResponseProperties_ProvisioningState_STATUS_Succeeded))
	gens["ResellerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionOwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Workload"] = gen.PtrOf(gen.OneConstOf(Workload_STATUS_DevTest, Workload_STATUS_Production))
}

func Test_SystemData_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS_ARM, SystemData_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS_ARM runs a test to see if a specific instance of SystemData_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS_ARM(subject SystemData_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS_ARM
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

// Generator of SystemData_STATUS_ARM instances for property testing - lazily instantiated by
// SystemData_STATUS_ARMGenerator()
var systemData_STATUS_ARMGenerator gopter.Gen

// SystemData_STATUS_ARMGenerator returns a generator of SystemData_STATUS_ARM instances for property testing.
func SystemData_STATUS_ARMGenerator() gopter.Gen {
	if systemData_STATUS_ARMGenerator != nil {
		return systemData_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(generators)
	systemData_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS_ARM{}), generators)

	return systemData_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_Application,
		SystemData_CreatedByType_STATUS_Key,
		SystemData_CreatedByType_STATUS_ManagedIdentity,
		SystemData_CreatedByType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_Application,
		SystemData_LastModifiedByType_STATUS_Key,
		SystemData_LastModifiedByType_STATUS_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_User))
}
