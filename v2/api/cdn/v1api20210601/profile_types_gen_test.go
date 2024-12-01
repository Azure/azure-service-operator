// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210601

import (
	"encoding/json"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20210601/storage"
	v20230501s "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/storage"
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

func Test_Profile_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Profile to hub returns original",
		prop.ForAll(RunResourceConversionTestForProfile, ProfileGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForProfile tests if a specific instance of Profile round trips to the hub storage version and back losslessly
func RunResourceConversionTestForProfile(subject Profile) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20230501s.Profile
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Profile
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

func Test_Profile_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Profile to Profile via AssignProperties_To_Profile & AssignProperties_From_Profile returns original",
		prop.ForAll(RunPropertyAssignmentTestForProfile, ProfileGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProfile tests if a specific instance of Profile can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProfile(subject Profile) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.Profile
	err := copied.AssignProperties_To_Profile(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Profile
	err = actual.AssignProperties_From_Profile(&other)
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

func Test_Profile_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profile via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfile, ProfileGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfile runs a test to see if a specific instance of Profile round trips to JSON and back losslessly
func RunJSONSerializationTestForProfile(subject Profile) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profile
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

// Generator of Profile instances for property testing - lazily instantiated by ProfileGenerator()
var profileGenerator gopter.Gen

// ProfileGenerator returns a generator of Profile instances for property testing.
func ProfileGenerator() gopter.Gen {
	if profileGenerator != nil {
		return profileGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForProfile(generators)
	profileGenerator = gen.Struct(reflect.TypeOf(Profile{}), generators)

	return profileGenerator
}

// AddRelatedPropertyGeneratorsForProfile is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfile(gens map[string]gopter.Gen) {
	gens["Spec"] = Profile_SpecGenerator()
	gens["Status"] = Profile_STATUSGenerator()
}

func Test_ProfileOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ProfileOperatorSpec to ProfileOperatorSpec via AssignProperties_To_ProfileOperatorSpec & AssignProperties_From_ProfileOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForProfileOperatorSpec, ProfileOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProfileOperatorSpec tests if a specific instance of ProfileOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProfileOperatorSpec(subject ProfileOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.ProfileOperatorSpec
	err := copied.AssignProperties_To_ProfileOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ProfileOperatorSpec
	err = actual.AssignProperties_From_ProfileOperatorSpec(&other)
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

func Test_ProfileOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProfileOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfileOperatorSpec, ProfileOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfileOperatorSpec runs a test to see if a specific instance of ProfileOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForProfileOperatorSpec(subject ProfileOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProfileOperatorSpec
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

// Generator of ProfileOperatorSpec instances for property testing - lazily instantiated by
// ProfileOperatorSpecGenerator()
var profileOperatorSpecGenerator gopter.Gen

// ProfileOperatorSpecGenerator returns a generator of ProfileOperatorSpec instances for property testing.
func ProfileOperatorSpecGenerator() gopter.Gen {
	if profileOperatorSpecGenerator != nil {
		return profileOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	profileOperatorSpecGenerator = gen.Struct(reflect.TypeOf(ProfileOperatorSpec{}), generators)

	return profileOperatorSpecGenerator
}

func Test_Profile_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Profile_STATUS to Profile_STATUS via AssignProperties_To_Profile_STATUS & AssignProperties_From_Profile_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForProfile_STATUS, Profile_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProfile_STATUS tests if a specific instance of Profile_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProfile_STATUS(subject Profile_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.Profile_STATUS
	err := copied.AssignProperties_To_Profile_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Profile_STATUS
	err = actual.AssignProperties_From_Profile_STATUS(&other)
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

func Test_Profile_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profile_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfile_STATUS, Profile_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfile_STATUS runs a test to see if a specific instance of Profile_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProfile_STATUS(subject Profile_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profile_STATUS
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

// Generator of Profile_STATUS instances for property testing - lazily instantiated by Profile_STATUSGenerator()
var profile_STATUSGenerator gopter.Gen

// Profile_STATUSGenerator returns a generator of Profile_STATUS instances for property testing.
// We first initialize profile_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profile_STATUSGenerator() gopter.Gen {
	if profile_STATUSGenerator != nil {
		return profile_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfile_STATUS(generators)
	profile_STATUSGenerator = gen.Struct(reflect.TypeOf(Profile_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfile_STATUS(generators)
	AddRelatedPropertyGeneratorsForProfile_STATUS(generators)
	profile_STATUSGenerator = gen.Struct(reflect.TypeOf(Profile_STATUS{}), generators)

	return profile_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForProfile_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfile_STATUS(gens map[string]gopter.Gen) {
	gens["FrontDoorId"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["OriginResponseTimeoutSeconds"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProfileProperties_ProvisioningState_STATUS_Creating,
		ProfileProperties_ProvisioningState_STATUS_Deleting,
		ProfileProperties_ProvisioningState_STATUS_Failed,
		ProfileProperties_ProvisioningState_STATUS_Succeeded,
		ProfileProperties_ProvisioningState_STATUS_Updating))
	gens["ResourceState"] = gen.PtrOf(gen.OneConstOf(
		ProfileProperties_ResourceState_STATUS_Active,
		ProfileProperties_ResourceState_STATUS_Creating,
		ProfileProperties_ResourceState_STATUS_Deleting,
		ProfileProperties_ResourceState_STATUS_Disabled))
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfile_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfile_STATUS(gens map[string]gopter.Gen) {
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Profile_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Profile_Spec to Profile_Spec via AssignProperties_To_Profile_Spec & AssignProperties_From_Profile_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForProfile_Spec, Profile_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProfile_Spec tests if a specific instance of Profile_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProfile_Spec(subject Profile_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.Profile_Spec
	err := copied.AssignProperties_To_Profile_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Profile_Spec
	err = actual.AssignProperties_From_Profile_Spec(&other)
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
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginResponseTimeoutSeconds"] = gen.PtrOf(gen.Int())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfile_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfile_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(ProfileOperatorSpecGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_Sku_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Sku to Sku via AssignProperties_To_Sku & AssignProperties_From_Sku returns original",
		prop.ForAll(RunPropertyAssignmentTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSku tests if a specific instance of Sku can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSku(subject Sku) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.Sku
	err := copied.AssignProperties_To_Sku(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Sku
	err = actual.AssignProperties_From_Sku(&other)
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

func Test_Sku_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Sku_STATUS to Sku_STATUS via AssignProperties_To_Sku_STATUS & AssignProperties_From_Sku_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSku_STATUS, Sku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSku_STATUS tests if a specific instance of Sku_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSku_STATUS(subject Sku_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.Sku_STATUS
	err := copied.AssignProperties_To_Sku_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Sku_STATUS
	err = actual.AssignProperties_From_Sku_STATUS(&other)
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

func Test_Sku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS, Sku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS runs a test to see if a specific instance of Sku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS(subject Sku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS
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

// Generator of Sku_STATUS instances for property testing - lazily instantiated by Sku_STATUSGenerator()
var sku_STATUSGenerator gopter.Gen

// Sku_STATUSGenerator returns a generator of Sku_STATUS instances for property testing.
func Sku_STATUSGenerator() gopter.Gen {
	if sku_STATUSGenerator != nil {
		return sku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS(generators)
	sku_STATUSGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS{}), generators)

	return sku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		Sku_Name_STATUS_Custom_Verizon,
		Sku_Name_STATUS_Premium_AzureFrontDoor,
		Sku_Name_STATUS_Premium_Verizon,
		Sku_Name_STATUS_StandardPlus_955BandWidth_ChinaCdn,
		Sku_Name_STATUS_StandardPlus_AvgBandWidth_ChinaCdn,
		Sku_Name_STATUS_StandardPlus_ChinaCdn,
		Sku_Name_STATUS_Standard_955BandWidth_ChinaCdn,
		Sku_Name_STATUS_Standard_Akamai,
		Sku_Name_STATUS_Standard_AvgBandWidth_ChinaCdn,
		Sku_Name_STATUS_Standard_AzureFrontDoor,
		Sku_Name_STATUS_Standard_ChinaCdn,
		Sku_Name_STATUS_Standard_Microsoft,
		Sku_Name_STATUS_Standard_Verizon))
}

func Test_SystemData_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SystemData_STATUS to SystemData_STATUS via AssignProperties_To_SystemData_STATUS & AssignProperties_From_SystemData_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSystemData_STATUS tests if a specific instance of SystemData_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.SystemData_STATUS
	err := copied.AssignProperties_To_SystemData_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SystemData_STATUS
	err = actual.AssignProperties_From_SystemData_STATUS(&other)
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

func Test_SystemData_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS runs a test to see if a specific instance of SystemData_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS
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

// Generator of SystemData_STATUS instances for property testing - lazily instantiated by SystemData_STATUSGenerator()
var systemData_STATUSGenerator gopter.Gen

// SystemData_STATUSGenerator returns a generator of SystemData_STATUS instances for property testing.
func SystemData_STATUSGenerator() gopter.Gen {
	if systemData_STATUSGenerator != nil {
		return systemData_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS(generators)
	systemData_STATUSGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS{}), generators)

	return systemData_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		IdentityType_STATUS_Application,
		IdentityType_STATUS_Key,
		IdentityType_STATUS_ManagedIdentity,
		IdentityType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		IdentityType_STATUS_Application,
		IdentityType_STATUS_Key,
		IdentityType_STATUS_ManagedIdentity,
		IdentityType_STATUS_User))
}
