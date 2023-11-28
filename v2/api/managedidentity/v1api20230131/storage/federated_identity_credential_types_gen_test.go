// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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
	gens["Spec"] = UserAssignedIdentities_FederatedIdentityCredential_SpecGenerator()
	gens["Status"] = UserAssignedIdentities_FederatedIdentityCredential_STATUSGenerator()
}

func Test_UserAssignedIdentities_FederatedIdentityCredential_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentities_FederatedIdentityCredential_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_Spec, UserAssignedIdentities_FederatedIdentityCredential_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_Spec runs a test to see if a specific instance of UserAssignedIdentities_FederatedIdentityCredential_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_Spec(subject UserAssignedIdentities_FederatedIdentityCredential_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentities_FederatedIdentityCredential_Spec
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

// Generator of UserAssignedIdentities_FederatedIdentityCredential_Spec instances for property testing - lazily
// instantiated by UserAssignedIdentities_FederatedIdentityCredential_SpecGenerator()
var userAssignedIdentities_FederatedIdentityCredential_SpecGenerator gopter.Gen

// UserAssignedIdentities_FederatedIdentityCredential_SpecGenerator returns a generator of UserAssignedIdentities_FederatedIdentityCredential_Spec instances for property testing.
func UserAssignedIdentities_FederatedIdentityCredential_SpecGenerator() gopter.Gen {
	if userAssignedIdentities_FederatedIdentityCredential_SpecGenerator != nil {
		return userAssignedIdentities_FederatedIdentityCredential_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_Spec(generators)
	userAssignedIdentities_FederatedIdentityCredential_SpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentities_FederatedIdentityCredential_Spec{}), generators)

	return userAssignedIdentities_FederatedIdentityCredential_SpecGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_Spec(gens map[string]gopter.Gen) {
	gens["Audiences"] = gen.SliceOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["Issuer"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentities_FederatedIdentityCredential_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentities_FederatedIdentityCredential_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_STATUS, UserAssignedIdentities_FederatedIdentityCredential_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_STATUS runs a test to see if a specific instance of UserAssignedIdentities_FederatedIdentityCredential_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_STATUS(subject UserAssignedIdentities_FederatedIdentityCredential_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentities_FederatedIdentityCredential_STATUS
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

// Generator of UserAssignedIdentities_FederatedIdentityCredential_STATUS instances for property testing - lazily
// instantiated by UserAssignedIdentities_FederatedIdentityCredential_STATUSGenerator()
var userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator gopter.Gen

// UserAssignedIdentities_FederatedIdentityCredential_STATUSGenerator returns a generator of UserAssignedIdentities_FederatedIdentityCredential_STATUS instances for property testing.
// We first initialize userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func UserAssignedIdentities_FederatedIdentityCredential_STATUSGenerator() gopter.Gen {
	if userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator != nil {
		return userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_STATUS(generators)
	userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentities_FederatedIdentityCredential_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_STATUS(generators)
	AddRelatedPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_STATUS(generators)
	userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentities_FederatedIdentityCredential_STATUS{}), generators)

	return userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_STATUS(gens map[string]gopter.Gen) {
	gens["Audiences"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Issuer"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
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
	gens["CreatedByType"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.AlphaString())
}
