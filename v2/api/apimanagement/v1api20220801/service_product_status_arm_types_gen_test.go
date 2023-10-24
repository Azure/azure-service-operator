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

func Test_Service_Product_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Product_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Product_STATUS_ARM, Service_Product_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Product_STATUS_ARM runs a test to see if a specific instance of Service_Product_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Product_STATUS_ARM(subject Service_Product_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Product_STATUS_ARM
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

// Generator of Service_Product_STATUS_ARM instances for property testing - lazily instantiated by
// Service_Product_STATUS_ARMGenerator()
var service_Product_STATUS_ARMGenerator gopter.Gen

// Service_Product_STATUS_ARMGenerator returns a generator of Service_Product_STATUS_ARM instances for property testing.
// We first initialize service_Product_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_Product_STATUS_ARMGenerator() gopter.Gen {
	if service_Product_STATUS_ARMGenerator != nil {
		return service_Product_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Product_STATUS_ARM(generators)
	service_Product_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_Product_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Product_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForService_Product_STATUS_ARM(generators)
	service_Product_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_Product_STATUS_ARM{}), generators)

	return service_Product_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForService_Product_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Product_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForService_Product_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_Product_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ProductContractProperties_STATUS_ARMGenerator())
}

func Test_ProductContractProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductContractProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductContractProperties_STATUS_ARM, ProductContractProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductContractProperties_STATUS_ARM runs a test to see if a specific instance of ProductContractProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProductContractProperties_STATUS_ARM(subject ProductContractProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductContractProperties_STATUS_ARM
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

// Generator of ProductContractProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ProductContractProperties_STATUS_ARMGenerator()
var productContractProperties_STATUS_ARMGenerator gopter.Gen

// ProductContractProperties_STATUS_ARMGenerator returns a generator of ProductContractProperties_STATUS_ARM instances for property testing.
func ProductContractProperties_STATUS_ARMGenerator() gopter.Gen {
	if productContractProperties_STATUS_ARMGenerator != nil {
		return productContractProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductContractProperties_STATUS_ARM(generators)
	productContractProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ProductContractProperties_STATUS_ARM{}), generators)

	return productContractProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProductContractProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProductContractProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ApprovalRequired"] = gen.PtrOf(gen.Bool())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ProductContractProperties_State_STATUS_NotPublished, ProductContractProperties_State_STATUS_Published))
	gens["SubscriptionRequired"] = gen.PtrOf(gen.Bool())
	gens["SubscriptionsLimit"] = gen.PtrOf(gen.Int())
	gens["Terms"] = gen.PtrOf(gen.AlphaString())
}
