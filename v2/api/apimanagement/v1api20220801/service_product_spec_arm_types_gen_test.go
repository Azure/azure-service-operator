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

func Test_ProductContractProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductContractProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductContractProperties_ARM, ProductContractProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductContractProperties_ARM runs a test to see if a specific instance of ProductContractProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProductContractProperties_ARM(subject ProductContractProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductContractProperties_ARM
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

// Generator of ProductContractProperties_ARM instances for property testing - lazily instantiated by
// ProductContractProperties_ARMGenerator()
var productContractProperties_ARMGenerator gopter.Gen

// ProductContractProperties_ARMGenerator returns a generator of ProductContractProperties_ARM instances for property testing.
func ProductContractProperties_ARMGenerator() gopter.Gen {
	if productContractProperties_ARMGenerator != nil {
		return productContractProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductContractProperties_ARM(generators)
	productContractProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ProductContractProperties_ARM{}), generators)

	return productContractProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProductContractProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProductContractProperties_ARM(gens map[string]gopter.Gen) {
	gens["ApprovalRequired"] = gen.PtrOf(gen.Bool())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ProductContractProperties_State_ARM_NotPublished, ProductContractProperties_State_ARM_Published))
	gens["SubscriptionRequired"] = gen.PtrOf(gen.Bool())
	gens["SubscriptionsLimit"] = gen.PtrOf(gen.Int())
	gens["Terms"] = gen.PtrOf(gen.AlphaString())
}

func Test_Service_Product_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Product_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Product_Spec_ARM, Service_Product_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Product_Spec_ARM runs a test to see if a specific instance of Service_Product_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Product_Spec_ARM(subject Service_Product_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Product_Spec_ARM
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

// Generator of Service_Product_Spec_ARM instances for property testing - lazily instantiated by
// Service_Product_Spec_ARMGenerator()
var service_Product_Spec_ARMGenerator gopter.Gen

// Service_Product_Spec_ARMGenerator returns a generator of Service_Product_Spec_ARM instances for property testing.
// We first initialize service_Product_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_Product_Spec_ARMGenerator() gopter.Gen {
	if service_Product_Spec_ARMGenerator != nil {
		return service_Product_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Product_Spec_ARM(generators)
	service_Product_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Service_Product_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Product_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForService_Product_Spec_ARM(generators)
	service_Product_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Service_Product_Spec_ARM{}), generators)

	return service_Product_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForService_Product_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Product_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForService_Product_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_Product_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ProductContractProperties_ARMGenerator())
}
