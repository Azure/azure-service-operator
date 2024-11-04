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

func Test_FederatedIdentityCredentialProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FederatedIdentityCredentialProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFederatedIdentityCredentialProperties_STATUS, FederatedIdentityCredentialProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFederatedIdentityCredentialProperties_STATUS runs a test to see if a specific instance of FederatedIdentityCredentialProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFederatedIdentityCredentialProperties_STATUS(subject FederatedIdentityCredentialProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FederatedIdentityCredentialProperties_STATUS
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

// Generator of FederatedIdentityCredentialProperties_STATUS instances for property testing - lazily instantiated by
// FederatedIdentityCredentialProperties_STATUSGenerator()
var federatedIdentityCredentialProperties_STATUSGenerator gopter.Gen

// FederatedIdentityCredentialProperties_STATUSGenerator returns a generator of FederatedIdentityCredentialProperties_STATUS instances for property testing.
func FederatedIdentityCredentialProperties_STATUSGenerator() gopter.Gen {
	if federatedIdentityCredentialProperties_STATUSGenerator != nil {
		return federatedIdentityCredentialProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_STATUS(generators)
	federatedIdentityCredentialProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredentialProperties_STATUS{}), generators)

	return federatedIdentityCredentialProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Audiences"] = gen.SliceOf(gen.AlphaString())
	gens["Issuer"] = gen.PtrOf(gen.AlphaString())
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
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
// We first initialize federatedIdentityCredential_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FederatedIdentityCredential_STATUSGenerator() gopter.Gen {
	if federatedIdentityCredential_STATUSGenerator != nil {
		return federatedIdentityCredential_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUS(generators)
	federatedIdentityCredential_STATUSGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredential_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUS(generators)
	AddRelatedPropertyGeneratorsForFederatedIdentityCredential_STATUS(generators)
	federatedIdentityCredential_STATUSGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredential_STATUS{}), generators)

	return federatedIdentityCredential_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFederatedIdentityCredential_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFederatedIdentityCredential_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FederatedIdentityCredentialProperties_STATUSGenerator())
}