// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

import (
	"encoding/json"
	v1api20200601s "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601storage"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601storage"
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

func Test_ResourceGroup_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ResourceGroup to hub returns original",
		prop.ForAll(RunResourceConversionTestForResourceGroup, ResourceGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForResourceGroup tests if a specific instance of ResourceGroup round trips to the hub storage version and back losslessly
func RunResourceConversionTestForResourceGroup(subject ResourceGroup) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1api20200601s.ResourceGroup
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ResourceGroup
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

func Test_ResourceGroup_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ResourceGroup to ResourceGroup via AssignProperties_To_ResourceGroup & AssignProperties_From_ResourceGroup returns original",
		prop.ForAll(RunPropertyAssignmentTestForResourceGroup, ResourceGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForResourceGroup tests if a specific instance of ResourceGroup can be assigned to v1beta20200601storage and back losslessly
func RunPropertyAssignmentTestForResourceGroup(subject ResourceGroup) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.ResourceGroup
	err := copied.AssignProperties_To_ResourceGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ResourceGroup
	err = actual.AssignProperties_From_ResourceGroup(&other)
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

func Test_ResourceGroup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceGroup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceGroup, ResourceGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceGroup runs a test to see if a specific instance of ResourceGroup round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceGroup(subject ResourceGroup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceGroup
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

// Generator of ResourceGroup instances for property testing - lazily instantiated by ResourceGroupGenerator()
var resourceGroupGenerator gopter.Gen

// ResourceGroupGenerator returns a generator of ResourceGroup instances for property testing.
func ResourceGroupGenerator() gopter.Gen {
	if resourceGroupGenerator != nil {
		return resourceGroupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForResourceGroup(generators)
	resourceGroupGenerator = gen.Struct(reflect.TypeOf(ResourceGroup{}), generators)

	return resourceGroupGenerator
}

// AddRelatedPropertyGeneratorsForResourceGroup is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForResourceGroup(gens map[string]gopter.Gen) {
	gens["Spec"] = ResourceGroup_SpecGenerator()
	gens["Status"] = ResourceGroup_STATUSGenerator()
}

func Test_ResourceGroup_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ResourceGroup_Spec to ResourceGroup_Spec via AssignProperties_To_ResourceGroup_Spec & AssignProperties_From_ResourceGroup_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForResourceGroup_Spec, ResourceGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForResourceGroup_Spec tests if a specific instance of ResourceGroup_Spec can be assigned to v1beta20200601storage and back losslessly
func RunPropertyAssignmentTestForResourceGroup_Spec(subject ResourceGroup_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.ResourceGroup_Spec
	err := copied.AssignProperties_To_ResourceGroup_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ResourceGroup_Spec
	err = actual.AssignProperties_From_ResourceGroup_Spec(&other)
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

func Test_ResourceGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceGroup_Spec, ResourceGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceGroup_Spec runs a test to see if a specific instance of ResourceGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceGroup_Spec(subject ResourceGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceGroup_Spec
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

// Generator of ResourceGroup_Spec instances for property testing - lazily instantiated by ResourceGroup_SpecGenerator()
var resourceGroup_SpecGenerator gopter.Gen

// ResourceGroup_SpecGenerator returns a generator of ResourceGroup_Spec instances for property testing.
func ResourceGroup_SpecGenerator() gopter.Gen {
	if resourceGroup_SpecGenerator != nil {
		return resourceGroup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceGroup_Spec(generators)
	resourceGroup_SpecGenerator = gen.Struct(reflect.TypeOf(ResourceGroup_Spec{}), generators)

	return resourceGroup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForResourceGroup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceGroup_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["ManagedBy"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_ResourceGroup_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ResourceGroup_STATUS to ResourceGroup_STATUS via AssignProperties_To_ResourceGroup_STATUS & AssignProperties_From_ResourceGroup_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForResourceGroup_STATUS, ResourceGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForResourceGroup_STATUS tests if a specific instance of ResourceGroup_STATUS can be assigned to v1beta20200601storage and back losslessly
func RunPropertyAssignmentTestForResourceGroup_STATUS(subject ResourceGroup_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.ResourceGroup_STATUS
	err := copied.AssignProperties_To_ResourceGroup_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ResourceGroup_STATUS
	err = actual.AssignProperties_From_ResourceGroup_STATUS(&other)
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

func Test_ResourceGroup_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceGroup_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceGroup_STATUS, ResourceGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceGroup_STATUS runs a test to see if a specific instance of ResourceGroup_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceGroup_STATUS(subject ResourceGroup_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceGroup_STATUS
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

// Generator of ResourceGroup_STATUS instances for property testing - lazily instantiated by
// ResourceGroup_STATUSGenerator()
var resourceGroup_STATUSGenerator gopter.Gen

// ResourceGroup_STATUSGenerator returns a generator of ResourceGroup_STATUS instances for property testing.
// We first initialize resourceGroup_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ResourceGroup_STATUSGenerator() gopter.Gen {
	if resourceGroup_STATUSGenerator != nil {
		return resourceGroup_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceGroup_STATUS(generators)
	resourceGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(ResourceGroup_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceGroup_STATUS(generators)
	AddRelatedPropertyGeneratorsForResourceGroup_STATUS(generators)
	resourceGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(ResourceGroup_STATUS{}), generators)

	return resourceGroup_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForResourceGroup_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceGroup_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["ManagedBy"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForResourceGroup_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForResourceGroup_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ResourceGroupProperties_STATUSGenerator())
}

func Test_ResourceGroupProperties_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ResourceGroupProperties_STATUS to ResourceGroupProperties_STATUS via AssignProperties_To_ResourceGroupProperties_STATUS & AssignProperties_From_ResourceGroupProperties_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForResourceGroupProperties_STATUS, ResourceGroupProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForResourceGroupProperties_STATUS tests if a specific instance of ResourceGroupProperties_STATUS can be assigned to v1beta20200601storage and back losslessly
func RunPropertyAssignmentTestForResourceGroupProperties_STATUS(subject ResourceGroupProperties_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.ResourceGroupProperties_STATUS
	err := copied.AssignProperties_To_ResourceGroupProperties_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ResourceGroupProperties_STATUS
	err = actual.AssignProperties_From_ResourceGroupProperties_STATUS(&other)
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

func Test_ResourceGroupProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceGroupProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceGroupProperties_STATUS, ResourceGroupProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceGroupProperties_STATUS runs a test to see if a specific instance of ResourceGroupProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceGroupProperties_STATUS(subject ResourceGroupProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceGroupProperties_STATUS
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

// Generator of ResourceGroupProperties_STATUS instances for property testing - lazily instantiated by
// ResourceGroupProperties_STATUSGenerator()
var resourceGroupProperties_STATUSGenerator gopter.Gen

// ResourceGroupProperties_STATUSGenerator returns a generator of ResourceGroupProperties_STATUS instances for property testing.
func ResourceGroupProperties_STATUSGenerator() gopter.Gen {
	if resourceGroupProperties_STATUSGenerator != nil {
		return resourceGroupProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceGroupProperties_STATUS(generators)
	resourceGroupProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ResourceGroupProperties_STATUS{}), generators)

	return resourceGroupProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForResourceGroupProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceGroupProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
}
