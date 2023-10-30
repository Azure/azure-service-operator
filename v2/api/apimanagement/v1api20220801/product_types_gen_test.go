// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import (
	"encoding/json"
	v20220801s "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801storage"
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

func Test_Product_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Product to hub returns original",
		prop.ForAll(RunResourceConversionTestForProduct, ProductGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForProduct tests if a specific instance of Product round trips to the hub storage version and back losslessly
func RunResourceConversionTestForProduct(subject Product) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20220801s.Product
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Product
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

func Test_Product_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Product to Product via AssignProperties_To_Product & AssignProperties_From_Product returns original",
		prop.ForAll(RunPropertyAssignmentTestForProduct, ProductGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProduct tests if a specific instance of Product can be assigned to v1api20220801storage and back losslessly
func RunPropertyAssignmentTestForProduct(subject Product) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220801s.Product
	err := copied.AssignProperties_To_Product(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Product
	err = actual.AssignProperties_From_Product(&other)
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

func Test_Product_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Product via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProduct, ProductGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProduct runs a test to see if a specific instance of Product round trips to JSON and back losslessly
func RunJSONSerializationTestForProduct(subject Product) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Product
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

// Generator of Product instances for property testing - lazily instantiated by ProductGenerator()
var productGenerator gopter.Gen

// ProductGenerator returns a generator of Product instances for property testing.
func ProductGenerator() gopter.Gen {
	if productGenerator != nil {
		return productGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForProduct(generators)
	productGenerator = gen.Struct(reflect.TypeOf(Product{}), generators)

	return productGenerator
}

// AddRelatedPropertyGeneratorsForProduct is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProduct(gens map[string]gopter.Gen) {
	gens["Spec"] = Service_Product_SpecGenerator()
	gens["Status"] = Service_Product_STATUSGenerator()
}

func Test_Service_Product_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Service_Product_Spec to Service_Product_Spec via AssignProperties_To_Service_Product_Spec & AssignProperties_From_Service_Product_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForService_Product_Spec, Service_Product_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForService_Product_Spec tests if a specific instance of Service_Product_Spec can be assigned to v1api20220801storage and back losslessly
func RunPropertyAssignmentTestForService_Product_Spec(subject Service_Product_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220801s.Service_Product_Spec
	err := copied.AssignProperties_To_Service_Product_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Service_Product_Spec
	err = actual.AssignProperties_From_Service_Product_Spec(&other)
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

func Test_Service_Product_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Product_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Product_Spec, Service_Product_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Product_Spec runs a test to see if a specific instance of Service_Product_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Product_Spec(subject Service_Product_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Product_Spec
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

// Generator of Service_Product_Spec instances for property testing - lazily instantiated by
// Service_Product_SpecGenerator()
var service_Product_SpecGenerator gopter.Gen

// Service_Product_SpecGenerator returns a generator of Service_Product_Spec instances for property testing.
func Service_Product_SpecGenerator() gopter.Gen {
	if service_Product_SpecGenerator != nil {
		return service_Product_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Product_Spec(generators)
	service_Product_SpecGenerator = gen.Struct(reflect.TypeOf(Service_Product_Spec{}), generators)

	return service_Product_SpecGenerator
}

// AddIndependentPropertyGeneratorsForService_Product_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Product_Spec(gens map[string]gopter.Gen) {
	gens["ApprovalRequired"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ProductContractProperties_State_NotPublished, ProductContractProperties_State_Published))
	gens["SubscriptionRequired"] = gen.PtrOf(gen.Bool())
	gens["SubscriptionsLimit"] = gen.PtrOf(gen.Int())
	gens["Terms"] = gen.PtrOf(gen.AlphaString())
}

func Test_Service_Product_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Service_Product_STATUS to Service_Product_STATUS via AssignProperties_To_Service_Product_STATUS & AssignProperties_From_Service_Product_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForService_Product_STATUS, Service_Product_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForService_Product_STATUS tests if a specific instance of Service_Product_STATUS can be assigned to v1api20220801storage and back losslessly
func RunPropertyAssignmentTestForService_Product_STATUS(subject Service_Product_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220801s.Service_Product_STATUS
	err := copied.AssignProperties_To_Service_Product_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Service_Product_STATUS
	err = actual.AssignProperties_From_Service_Product_STATUS(&other)
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

func Test_Service_Product_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Product_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Product_STATUS, Service_Product_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Product_STATUS runs a test to see if a specific instance of Service_Product_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Product_STATUS(subject Service_Product_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Product_STATUS
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

// Generator of Service_Product_STATUS instances for property testing - lazily instantiated by
// Service_Product_STATUSGenerator()
var service_Product_STATUSGenerator gopter.Gen

// Service_Product_STATUSGenerator returns a generator of Service_Product_STATUS instances for property testing.
func Service_Product_STATUSGenerator() gopter.Gen {
	if service_Product_STATUSGenerator != nil {
		return service_Product_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Product_STATUS(generators)
	service_Product_STATUSGenerator = gen.Struct(reflect.TypeOf(Service_Product_STATUS{}), generators)

	return service_Product_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForService_Product_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Product_STATUS(gens map[string]gopter.Gen) {
	gens["ApprovalRequired"] = gen.PtrOf(gen.Bool())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ProductContractProperties_State_STATUS_NotPublished, ProductContractProperties_State_STATUS_Published))
	gens["SubscriptionRequired"] = gen.PtrOf(gen.Bool())
	gens["SubscriptionsLimit"] = gen.PtrOf(gen.Int())
	gens["Terms"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
