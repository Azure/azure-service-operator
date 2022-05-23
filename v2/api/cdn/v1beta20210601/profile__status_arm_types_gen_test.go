// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

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

func Test_Profile_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profile_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfileStatusARM, ProfileStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfileStatusARM runs a test to see if a specific instance of Profile_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfileStatusARM(subject Profile_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profile_StatusARM
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

// Generator of Profile_StatusARM instances for property testing - lazily instantiated by ProfileStatusARMGenerator()
var profileStatusARMGenerator gopter.Gen

// ProfileStatusARMGenerator returns a generator of Profile_StatusARM instances for property testing.
// We first initialize profileStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ProfileStatusARMGenerator() gopter.Gen {
	if profileStatusARMGenerator != nil {
		return profileStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfileStatusARM(generators)
	profileStatusARMGenerator = gen.Struct(reflect.TypeOf(Profile_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfileStatusARM(generators)
	AddRelatedPropertyGeneratorsForProfileStatusARM(generators)
	profileStatusARMGenerator = gen.Struct(reflect.TypeOf(Profile_StatusARM{}), generators)

	return profileStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForProfileStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfileStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfileStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfileStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ProfilePropertiesStatusARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuStatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_ProfileProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProfileProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfilePropertiesStatusARM, ProfilePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfilePropertiesStatusARM runs a test to see if a specific instance of ProfileProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfilePropertiesStatusARM(subject ProfileProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProfileProperties_StatusARM
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

// Generator of ProfileProperties_StatusARM instances for property testing - lazily instantiated by
// ProfilePropertiesStatusARMGenerator()
var profilePropertiesStatusARMGenerator gopter.Gen

// ProfilePropertiesStatusARMGenerator returns a generator of ProfileProperties_StatusARM instances for property testing.
func ProfilePropertiesStatusARMGenerator() gopter.Gen {
	if profilePropertiesStatusARMGenerator != nil {
		return profilePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfilePropertiesStatusARM(generators)
	profilePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ProfileProperties_StatusARM{}), generators)

	return profilePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForProfilePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfilePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["FrontDoorId"] = gen.PtrOf(gen.AlphaString())
	gens["OriginResponseTimeoutSeconds"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProfilePropertiesStatusProvisioningStateCreating,
		ProfilePropertiesStatusProvisioningStateDeleting,
		ProfilePropertiesStatusProvisioningStateFailed,
		ProfilePropertiesStatusProvisioningStateSucceeded,
		ProfilePropertiesStatusProvisioningStateUpdating))
	gens["ResourceState"] = gen.PtrOf(gen.OneConstOf(
		ProfilePropertiesStatusResourceStateActive,
		ProfilePropertiesStatusResourceStateCreating,
		ProfilePropertiesStatusResourceStateDeleting,
		ProfilePropertiesStatusResourceStateDisabled))
}

func Test_Sku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuStatusARM, SkuStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuStatusARM runs a test to see if a specific instance of Sku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuStatusARM(subject Sku_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_StatusARM
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

// Generator of Sku_StatusARM instances for property testing - lazily instantiated by SkuStatusARMGenerator()
var skuStatusARMGenerator gopter.Gen

// SkuStatusARMGenerator returns a generator of Sku_StatusARM instances for property testing.
func SkuStatusARMGenerator() gopter.Gen {
	if skuStatusARMGenerator != nil {
		return skuStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuStatusARM(generators)
	skuStatusARMGenerator = gen.Struct(reflect.TypeOf(Sku_StatusARM{}), generators)

	return skuStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuStatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		SkuStatusNameCustomVerizon,
		SkuStatusNamePremiumAzureFrontDoor,
		SkuStatusNamePremiumVerizon,
		SkuStatusNameStandard955BandWidthChinaCdn,
		SkuStatusNameStandardAkamai,
		SkuStatusNameStandardAvgBandWidthChinaCdn,
		SkuStatusNameStandardAzureFrontDoor,
		SkuStatusNameStandardChinaCdn,
		SkuStatusNameStandardMicrosoft,
		SkuStatusNameStandardPlus955BandWidthChinaCdn,
		SkuStatusNameStandardPlusAvgBandWidthChinaCdn,
		SkuStatusNameStandardPlusChinaCdn,
		SkuStatusNameStandardVerizon))
}
