// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

import (
	"encoding/json"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/cdn/v1beta20210601storage"
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
	var hub v20210601s.Profile
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
		"Round trip from Profile to Profile via AssignPropertiesToProfile & AssignPropertiesFromProfile returns original",
		prop.ForAll(RunPropertyAssignmentTestForProfile, ProfileGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProfile tests if a specific instance of Profile can be assigned to v1beta20210601storage and back losslessly
func RunPropertyAssignmentTestForProfile(subject Profile) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.Profile
	err := copied.AssignPropertiesToProfile(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Profile
	err = actual.AssignPropertiesFromProfile(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
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
	parameters.MaxSize = 10
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
	gens["Spec"] = ProfilesSpecGenerator()
	gens["Status"] = ProfileStatusGenerator()
}

func Test_Profile_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Profile_Status to Profile_Status via AssignPropertiesToProfileStatus & AssignPropertiesFromProfileStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForProfileStatus, ProfileStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProfileStatus tests if a specific instance of Profile_Status can be assigned to v1beta20210601storage and back losslessly
func RunPropertyAssignmentTestForProfileStatus(subject Profile_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.Profile_Status
	err := copied.AssignPropertiesToProfileStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Profile_Status
	err = actual.AssignPropertiesFromProfileStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Profile_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profile_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfileStatus, ProfileStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfileStatus runs a test to see if a specific instance of Profile_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForProfileStatus(subject Profile_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profile_Status
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

// Generator of Profile_Status instances for property testing - lazily instantiated by ProfileStatusGenerator()
var profileStatusGenerator gopter.Gen

// ProfileStatusGenerator returns a generator of Profile_Status instances for property testing.
// We first initialize profileStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ProfileStatusGenerator() gopter.Gen {
	if profileStatusGenerator != nil {
		return profileStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfileStatus(generators)
	profileStatusGenerator = gen.Struct(reflect.TypeOf(Profile_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfileStatus(generators)
	AddRelatedPropertyGeneratorsForProfileStatus(generators)
	profileStatusGenerator = gen.Struct(reflect.TypeOf(Profile_Status{}), generators)

	return profileStatusGenerator
}

// AddIndependentPropertyGeneratorsForProfileStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfileStatus(gens map[string]gopter.Gen) {
	gens["FrontDoorId"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
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
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfileStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfileStatus(gens map[string]gopter.Gen) {
	gens["Sku"] = gen.PtrOf(SkuStatusGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_Profiles_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Profiles_Spec to Profiles_Spec via AssignPropertiesToProfilesSpec & AssignPropertiesFromProfilesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForProfilesSpec, ProfilesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProfilesSpec tests if a specific instance of Profiles_Spec can be assigned to v1beta20210601storage and back losslessly
func RunPropertyAssignmentTestForProfilesSpec(subject Profiles_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.Profiles_Spec
	err := copied.AssignPropertiesToProfilesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Profiles_Spec
	err = actual.AssignPropertiesFromProfilesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Profiles_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfilesSpec, ProfilesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfilesSpec runs a test to see if a specific instance of Profiles_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForProfilesSpec(subject Profiles_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_Spec
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

// Generator of Profiles_Spec instances for property testing - lazily instantiated by ProfilesSpecGenerator()
var profilesSpecGenerator gopter.Gen

// ProfilesSpecGenerator returns a generator of Profiles_Spec instances for property testing.
// We first initialize profilesSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ProfilesSpecGenerator() gopter.Gen {
	if profilesSpecGenerator != nil {
		return profilesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfilesSpec(generators)
	profilesSpecGenerator = gen.Struct(reflect.TypeOf(Profiles_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfilesSpec(generators)
	AddRelatedPropertyGeneratorsForProfilesSpec(generators)
	profilesSpecGenerator = gen.Struct(reflect.TypeOf(Profiles_Spec{}), generators)

	return profilesSpecGenerator
}

// AddIndependentPropertyGeneratorsForProfilesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfilesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginResponseTimeoutSeconds"] = gen.PtrOf(gen.Int())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfilesSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfilesSpec(gens map[string]gopter.Gen) {
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_Sku_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Sku to Sku via AssignPropertiesToSku & AssignPropertiesFromSku returns original",
		prop.ForAll(RunPropertyAssignmentTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSku tests if a specific instance of Sku can be assigned to v1beta20210601storage and back losslessly
func RunPropertyAssignmentTestForSku(subject Sku) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.Sku
	err := copied.AssignPropertiesToSku(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Sku
	err = actual.AssignPropertiesFromSku(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
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
	parameters.MaxSize = 10
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
		SkuNameCustomVerizon,
		SkuNamePremiumAzureFrontDoor,
		SkuNamePremiumVerizon,
		SkuNameStandard955BandWidthChinaCdn,
		SkuNameStandardAkamai,
		SkuNameStandardAvgBandWidthChinaCdn,
		SkuNameStandardAzureFrontDoor,
		SkuNameStandardChinaCdn,
		SkuNameStandardMicrosoft,
		SkuNameStandardPlus955BandWidthChinaCdn,
		SkuNameStandardPlusAvgBandWidthChinaCdn,
		SkuNameStandardPlusChinaCdn,
		SkuNameStandardVerizon))
}

func Test_Sku_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Sku_Status to Sku_Status via AssignPropertiesToSkuStatus & AssignPropertiesFromSkuStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSkuStatus, SkuStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSkuStatus tests if a specific instance of Sku_Status can be assigned to v1beta20210601storage and back losslessly
func RunPropertyAssignmentTestForSkuStatus(subject Sku_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.Sku_Status
	err := copied.AssignPropertiesToSkuStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Sku_Status
	err = actual.AssignPropertiesFromSkuStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Sku_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuStatus, SkuStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuStatus runs a test to see if a specific instance of Sku_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuStatus(subject Sku_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_Status
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

// Generator of Sku_Status instances for property testing - lazily instantiated by SkuStatusGenerator()
var skuStatusGenerator gopter.Gen

// SkuStatusGenerator returns a generator of Sku_Status instances for property testing.
func SkuStatusGenerator() gopter.Gen {
	if skuStatusGenerator != nil {
		return skuStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuStatus(generators)
	skuStatusGenerator = gen.Struct(reflect.TypeOf(Sku_Status{}), generators)

	return skuStatusGenerator
}

// AddIndependentPropertyGeneratorsForSkuStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuStatus(gens map[string]gopter.Gen) {
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

func Test_SystemData_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SystemData_Status to SystemData_Status via AssignPropertiesToSystemDataStatus & AssignPropertiesFromSystemDataStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSystemDataStatus, SystemDataStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSystemDataStatus tests if a specific instance of SystemData_Status can be assigned to v1beta20210601storage and back losslessly
func RunPropertyAssignmentTestForSystemDataStatus(subject SystemData_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.SystemData_Status
	err := copied.AssignPropertiesToSystemDataStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SystemData_Status
	err = actual.AssignPropertiesFromSystemDataStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SystemData_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataStatus, SystemDataStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataStatus runs a test to see if a specific instance of SystemData_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataStatus(subject SystemData_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_Status
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

// Generator of SystemData_Status instances for property testing - lazily instantiated by SystemDataStatusGenerator()
var systemDataStatusGenerator gopter.Gen

// SystemDataStatusGenerator returns a generator of SystemData_Status instances for property testing.
func SystemDataStatusGenerator() gopter.Gen {
	if systemDataStatusGenerator != nil {
		return systemDataStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataStatus(generators)
	systemDataStatusGenerator = gen.Struct(reflect.TypeOf(SystemData_Status{}), generators)

	return systemDataStatusGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataStatus(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		IdentityType_StatusApplication,
		IdentityType_StatusKey,
		IdentityType_StatusManagedIdentity,
		IdentityType_StatusUser))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		IdentityType_StatusApplication,
		IdentityType_StatusKey,
		IdentityType_StatusManagedIdentity,
		IdentityType_StatusUser))
}
