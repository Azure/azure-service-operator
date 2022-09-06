// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220131previewstorage

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

func Test_FederatedIdentityCredential_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FederatedIdentityCredential via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFederatedIdentityCredential, FederatedIdentityCredentialGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFederatedIdentityCredential runs a test to see if a specific instance of FederatedIdentityCredential round trips to JSON and back losslessly
func RunJSONSerializationTestForFederatedIdentityCredential(subject FederatedIdentityCredential) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FederatedIdentityCredential
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

// Generator of FederatedIdentityCredential instances for property testing - lazily instantiated by
// FederatedIdentityCredentialGenerator()
var federatedIdentityCredentialGenerator gopter.Gen

// FederatedIdentityCredentialGenerator returns a generator of FederatedIdentityCredential instances for property testing.
func FederatedIdentityCredentialGenerator() gopter.Gen {
	if federatedIdentityCredentialGenerator != nil {
		return federatedIdentityCredentialGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFederatedIdentityCredential(generators)
	federatedIdentityCredentialGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredential{}), generators)

	return federatedIdentityCredentialGenerator
}

// AddRelatedPropertyGeneratorsForFederatedIdentityCredential is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFederatedIdentityCredential(gens map[string]gopter.Gen) {
	gens["Spec"] = UserAssignedIdentities_FederatedIdentityCredentials_SpecGenerator()
	gens["Status"] = FederatedIdentityCredential_STATUSGenerator()
}

func Test_FederatedIdentityCredential_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FederatedIdentityCredential_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFederatedIdentityCredential_STATUS, FederatedIdentityCredential_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFederatedIdentityCredential_STATUS runs a test to see if a specific instance of FederatedIdentityCredential_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFederatedIdentityCredential_STATUS(subject FederatedIdentityCredential_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FederatedIdentityCredential_STATUS
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

// Generator of FederatedIdentityCredential_STATUS instances for property testing - lazily instantiated by
// FederatedIdentityCredential_STATUSGenerator()
var federatedIdentityCredential_STATUSGenerator gopter.Gen

// FederatedIdentityCredential_STATUSGenerator returns a generator of FederatedIdentityCredential_STATUS instances for property testing.
func FederatedIdentityCredential_STATUSGenerator() gopter.Gen {
	if federatedIdentityCredential_STATUSGenerator != nil {
		return federatedIdentityCredential_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUS(generators)
	federatedIdentityCredential_STATUSGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredential_STATUS{}), generators)

	return federatedIdentityCredential_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUS(gens map[string]gopter.Gen) {
	gens["Audiences"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Issuer"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentities_FederatedIdentityCredentials_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentities_FederatedIdentityCredentials_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredentials_Spec, UserAssignedIdentities_FederatedIdentityCredentials_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredentials_Spec runs a test to see if a specific instance of UserAssignedIdentities_FederatedIdentityCredentials_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredentials_Spec(subject UserAssignedIdentities_FederatedIdentityCredentials_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentities_FederatedIdentityCredentials_Spec
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

// Generator of UserAssignedIdentities_FederatedIdentityCredentials_Spec instances for property testing - lazily
// instantiated by UserAssignedIdentities_FederatedIdentityCredentials_SpecGenerator()
var userAssignedIdentities_FederatedIdentityCredentials_SpecGenerator gopter.Gen

// UserAssignedIdentities_FederatedIdentityCredentials_SpecGenerator returns a generator of UserAssignedIdentities_FederatedIdentityCredentials_Spec instances for property testing.
func UserAssignedIdentities_FederatedIdentityCredentials_SpecGenerator() gopter.Gen {
	if userAssignedIdentities_FederatedIdentityCredentials_SpecGenerator != nil {
		return userAssignedIdentities_FederatedIdentityCredentials_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredentials_Spec(generators)
	userAssignedIdentities_FederatedIdentityCredentials_SpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentities_FederatedIdentityCredentials_Spec{}), generators)

	return userAssignedIdentities_FederatedIdentityCredentials_SpecGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredentials_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredentials_Spec(gens map[string]gopter.Gen) {
	gens["Audiences"] = gen.SliceOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["Issuer"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
}
