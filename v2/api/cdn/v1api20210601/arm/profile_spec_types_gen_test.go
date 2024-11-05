// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_ProfileProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProfileProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfileProperties, ProfilePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfileProperties runs a test to see if a specific instance of ProfileProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForProfileProperties(subject ProfileProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProfileProperties
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

// Generator of ProfileProperties instances for property testing - lazily instantiated by ProfilePropertiesGenerator()
var profilePropertiesGenerator gopter.Gen

// ProfilePropertiesGenerator returns a generator of ProfileProperties instances for property testing.
func ProfilePropertiesGenerator() gopter.Gen {
	if profilePropertiesGenerator != nil {
		return profilePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfileProperties(generators)
	profilePropertiesGenerator = gen.Struct(reflect.TypeOf(ProfileProperties{}), generators)

	return profilePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForProfileProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfileProperties(gens map[string]gopter.Gen) {
	gens["OriginResponseTimeoutSeconds"] = gen.PtrOf(gen.Int())
}

func Test_Profile_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profile_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfile_Spec, Profile_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfile_Spec runs a test to see if a specific instance of Profile_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForProfile_Spec(subject Profile_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profile_Spec
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

// Generator of Profile_Spec instances for property testing - lazily instantiated by Profile_SpecGenerator()
var profile_SpecGenerator gopter.Gen

// Profile_SpecGenerator returns a generator of Profile_Spec instances for property testing.
// We first initialize profile_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profile_SpecGenerator() gopter.Gen {
	if profile_SpecGenerator != nil {
		return profile_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfile_Spec(generators)
	profile_SpecGenerator = gen.Struct(reflect.TypeOf(Profile_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfile_Spec(generators)
	AddRelatedPropertyGeneratorsForProfile_Spec(generators)
	profile_SpecGenerator = gen.Struct(reflect.TypeOf(Profile_Spec{}), generators)

	return profile_SpecGenerator
}

// AddIndependentPropertyGeneratorsForProfile_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfile_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfile_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfile_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ProfilePropertiesGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		Sku_Name_Custom_Verizon,
		Sku_Name_Premium_AzureFrontDoor,
		Sku_Name_Premium_Verizon,
		Sku_Name_StandardPlus_955BandWidth_ChinaCdn,
		Sku_Name_StandardPlus_AvgBandWidth_ChinaCdn,
		Sku_Name_StandardPlus_ChinaCdn,
		Sku_Name_Standard_955BandWidth_ChinaCdn,
		Sku_Name_Standard_Akamai,
		Sku_Name_Standard_AvgBandWidth_ChinaCdn,
		Sku_Name_Standard_AzureFrontDoor,
		Sku_Name_Standard_ChinaCdn,
		Sku_Name_Standard_Microsoft,
		Sku_Name_Standard_Verizon))
}
