// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20181130storage

import (
	"encoding/json"
	v20181130s "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1beta20181130storage"
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

func Test_UserAssignedIdentity_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentity to hub returns original",
		prop.ForAll(RunResourceConversionTestForUserAssignedIdentity, UserAssignedIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForUserAssignedIdentity tests if a specific instance of UserAssignedIdentity round trips to the hub storage version and back losslessly
func RunResourceConversionTestForUserAssignedIdentity(subject UserAssignedIdentity) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20181130s.UserAssignedIdentity
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual UserAssignedIdentity
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

func Test_UserAssignedIdentity_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentity to UserAssignedIdentity via AssignPropertiesToUserAssignedIdentity & AssignPropertiesFromUserAssignedIdentity returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentity, UserAssignedIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentity tests if a specific instance of UserAssignedIdentity can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentity(subject UserAssignedIdentity) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20181130s.UserAssignedIdentity
	err := copied.AssignPropertiesToUserAssignedIdentity(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentity
	err = actual.AssignPropertiesFromUserAssignedIdentity(&other)
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

func Test_UserAssignedIdentity_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity, UserAssignedIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity runs a test to see if a specific instance of UserAssignedIdentity round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity(subject UserAssignedIdentity) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity
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

// Generator of UserAssignedIdentity instances for property testing - lazily instantiated by
// UserAssignedIdentityGenerator()
var userAssignedIdentityGenerator gopter.Gen

// UserAssignedIdentityGenerator returns a generator of UserAssignedIdentity instances for property testing.
func UserAssignedIdentityGenerator() gopter.Gen {
	if userAssignedIdentityGenerator != nil {
		return userAssignedIdentityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUserAssignedIdentity(generators)
	userAssignedIdentityGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity{}), generators)

	return userAssignedIdentityGenerator
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentity is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentity(gens map[string]gopter.Gen) {
	gens["Spec"] = UserAssignedIdentitiesSpecGenerator()
	gens["Status"] = IdentitySTATUSGenerator()
}

func Test_Identity_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Identity_STATUS to Identity_STATUS via AssignPropertiesToIdentitySTATUS & AssignPropertiesFromIdentitySTATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForIdentitySTATUS, IdentitySTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForIdentitySTATUS tests if a specific instance of Identity_STATUS can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForIdentitySTATUS(subject Identity_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20181130s.Identity_STATUS
	err := copied.AssignPropertiesToIdentitySTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Identity_STATUS
	err = actual.AssignPropertiesFromIdentitySTATUS(&other)
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

func Test_Identity_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentitySTATUS, IdentitySTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentitySTATUS runs a test to see if a specific instance of Identity_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentitySTATUS(subject Identity_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_STATUS
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

// Generator of Identity_STATUS instances for property testing - lazily instantiated by IdentitySTATUSGenerator()
var identitySTATUSGenerator gopter.Gen

// IdentitySTATUSGenerator returns a generator of Identity_STATUS instances for property testing.
func IdentitySTATUSGenerator() gopter.Gen {
	if identitySTATUSGenerator != nil {
		return identitySTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentitySTATUS(generators)
	identitySTATUSGenerator = gen.Struct(reflect.TypeOf(Identity_STATUS{}), generators)

	return identitySTATUSGenerator
}

// AddIndependentPropertyGeneratorsForIdentitySTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentitySTATUS(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentities_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentities_Spec to UserAssignedIdentities_Spec via AssignPropertiesToUserAssignedIdentitiesSpec & AssignPropertiesFromUserAssignedIdentitiesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentitiesSpec, UserAssignedIdentitiesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentitiesSpec tests if a specific instance of UserAssignedIdentities_Spec can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentitiesSpec(subject UserAssignedIdentities_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20181130s.UserAssignedIdentities_Spec
	err := copied.AssignPropertiesToUserAssignedIdentitiesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentities_Spec
	err = actual.AssignPropertiesFromUserAssignedIdentitiesSpec(&other)
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

func Test_UserAssignedIdentities_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentities_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentitiesSpec, UserAssignedIdentitiesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentitiesSpec runs a test to see if a specific instance of UserAssignedIdentities_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentitiesSpec(subject UserAssignedIdentities_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentities_Spec
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

// Generator of UserAssignedIdentities_Spec instances for property testing - lazily instantiated by
// UserAssignedIdentitiesSpecGenerator()
var userAssignedIdentitiesSpecGenerator gopter.Gen

// UserAssignedIdentitiesSpecGenerator returns a generator of UserAssignedIdentities_Spec instances for property testing.
func UserAssignedIdentitiesSpecGenerator() gopter.Gen {
	if userAssignedIdentitiesSpecGenerator != nil {
		return userAssignedIdentitiesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSpec(generators)
	userAssignedIdentitiesSpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentities_Spec{}), generators)

	return userAssignedIdentitiesSpecGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
