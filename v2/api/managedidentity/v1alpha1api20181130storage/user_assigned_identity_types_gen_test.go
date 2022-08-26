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
		"Round trip from UserAssignedIdentity to UserAssignedIdentity via AssignProperties_To_UserAssignedIdentity & AssignProperties_From_UserAssignedIdentity returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentity, UserAssignedIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentity tests if a specific instance of UserAssignedIdentity can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentity(subject UserAssignedIdentity) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20181130s.UserAssignedIdentity
	err := copied.AssignProperties_To_UserAssignedIdentity(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentity
	err = actual.AssignProperties_From_UserAssignedIdentity(&other)
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
<<<<<<< HEAD
	gens["Spec"] = UserAssignedIdentity_SpecGenerator()
	gens["Status"] = UserAssignedIdentity_STATUSGenerator()
=======
	gens["Spec"] = UserAssignedIdentities_SpecGenerator()
	gens["Status"] = Identity_STATUSGenerator()
>>>>>>> main
}

func Test_UserAssignedIdentity_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip from UserAssignedIdentity_Spec to UserAssignedIdentity_Spec via AssignPropertiesToUserAssignedIdentity_Spec & AssignPropertiesFromUserAssignedIdentity_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentity_Spec, UserAssignedIdentity_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentity_Spec tests if a specific instance of UserAssignedIdentity_Spec can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentity_Spec(subject UserAssignedIdentity_Spec) string {
=======
		"Round trip from Identity_STATUS to Identity_STATUS via AssignProperties_To_Identity_STATUS & AssignProperties_From_Identity_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForIdentity_STATUS, Identity_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForIdentity_STATUS tests if a specific instance of Identity_STATUS can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForIdentity_STATUS(subject Identity_STATUS) string {
>>>>>>> main
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
<<<<<<< HEAD
	var other v20181130s.UserAssignedIdentity_Spec
	err := copied.AssignPropertiesToUserAssignedIdentity_Spec(&other)
=======
	var other v20181130s.Identity_STATUS
	err := copied.AssignProperties_To_Identity_STATUS(&other)
>>>>>>> main
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
<<<<<<< HEAD
	var actual UserAssignedIdentity_Spec
	err = actual.AssignPropertiesFromUserAssignedIdentity_Spec(&other)
=======
	var actual Identity_STATUS
	err = actual.AssignProperties_From_Identity_STATUS(&other)
>>>>>>> main
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

func Test_UserAssignedIdentity_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip of UserAssignedIdentity_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_Spec, UserAssignedIdentity_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_Spec runs a test to see if a specific instance of UserAssignedIdentity_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_Spec(subject UserAssignedIdentity_Spec) string {
=======
		"Round trip of Identity_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity_STATUS, Identity_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity_STATUS runs a test to see if a specific instance of Identity_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity_STATUS(subject Identity_STATUS) string {
>>>>>>> main
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_Spec
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

<<<<<<< HEAD
// Generator of UserAssignedIdentity_Spec instances for property testing - lazily instantiated by
// UserAssignedIdentity_SpecGenerator()
var userAssignedIdentity_SpecGenerator gopter.Gen

// UserAssignedIdentity_SpecGenerator returns a generator of UserAssignedIdentity_Spec instances for property testing.
func UserAssignedIdentity_SpecGenerator() gopter.Gen {
	if userAssignedIdentity_SpecGenerator != nil {
		return userAssignedIdentity_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec(generators)
	userAssignedIdentity_SpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_Spec{}), generators)

	return userAssignedIdentity_SpecGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_UserAssignedIdentity_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentity_STATUS to UserAssignedIdentity_STATUS via AssignPropertiesToUserAssignedIdentity_STATUS & AssignPropertiesFromUserAssignedIdentity_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentity_STATUS, UserAssignedIdentity_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentity_STATUS tests if a specific instance of UserAssignedIdentity_STATUS can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentity_STATUS(subject UserAssignedIdentity_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20181130s.UserAssignedIdentity_STATUS
	err := copied.AssignPropertiesToUserAssignedIdentity_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentity_STATUS
	err = actual.AssignPropertiesFromUserAssignedIdentity_STATUS(&other)
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

func Test_UserAssignedIdentity_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_STATUS, UserAssignedIdentity_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_STATUS runs a test to see if a specific instance of UserAssignedIdentity_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_STATUS(subject UserAssignedIdentity_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_STATUS
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

// Generator of UserAssignedIdentity_STATUS instances for property testing - lazily instantiated by
// UserAssignedIdentity_STATUSGenerator()
var userAssignedIdentity_STATUSGenerator gopter.Gen

// UserAssignedIdentity_STATUSGenerator returns a generator of UserAssignedIdentity_STATUS instances for property testing.
func UserAssignedIdentity_STATUSGenerator() gopter.Gen {
	if userAssignedIdentity_STATUSGenerator != nil {
		return userAssignedIdentity_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS(generators)
	userAssignedIdentity_STATUSGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUS{}), generators)

	return userAssignedIdentity_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS(gens map[string]gopter.Gen) {
=======
// Generator of Identity_STATUS instances for property testing - lazily instantiated by Identity_STATUSGenerator()
var identity_STATUSGenerator gopter.Gen

// Identity_STATUSGenerator returns a generator of Identity_STATUS instances for property testing.
func Identity_STATUSGenerator() gopter.Gen {
	if identity_STATUSGenerator != nil {
		return identity_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUS(generators)
	identity_STATUSGenerator = gen.Struct(reflect.TypeOf(Identity_STATUS{}), generators)

	return identity_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForIdentity_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity_STATUS(gens map[string]gopter.Gen) {
>>>>>>> main
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
<<<<<<< HEAD
=======

func Test_UserAssignedIdentities_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentities_Spec to UserAssignedIdentities_Spec via AssignProperties_To_UserAssignedIdentities_Spec & AssignProperties_From_UserAssignedIdentities_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentities_Spec, UserAssignedIdentities_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentities_Spec tests if a specific instance of UserAssignedIdentities_Spec can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentities_Spec(subject UserAssignedIdentities_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20181130s.UserAssignedIdentities_Spec
	err := copied.AssignProperties_To_UserAssignedIdentities_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentities_Spec
	err = actual.AssignProperties_From_UserAssignedIdentities_Spec(&other)
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
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentities_Spec, UserAssignedIdentities_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentities_Spec runs a test to see if a specific instance of UserAssignedIdentities_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentities_Spec(subject UserAssignedIdentities_Spec) string {
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
// UserAssignedIdentities_SpecGenerator()
var userAssignedIdentities_SpecGenerator gopter.Gen

// UserAssignedIdentities_SpecGenerator returns a generator of UserAssignedIdentities_Spec instances for property testing.
func UserAssignedIdentities_SpecGenerator() gopter.Gen {
	if userAssignedIdentities_SpecGenerator != nil {
		return userAssignedIdentities_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentities_Spec(generators)
	userAssignedIdentities_SpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentities_Spec{}), generators)

	return userAssignedIdentities_SpecGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentities_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentities_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
>>>>>>> main
