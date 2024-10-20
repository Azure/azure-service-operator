// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220131preview

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

func Test_FederatedIdentityCredentialProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FederatedIdentityCredentialProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFederatedIdentityCredentialProperties_ARM, FederatedIdentityCredentialProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFederatedIdentityCredentialProperties_ARM runs a test to see if a specific instance of FederatedIdentityCredentialProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFederatedIdentityCredentialProperties_ARM(subject FederatedIdentityCredentialProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FederatedIdentityCredentialProperties_ARM
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

// Generator of FederatedIdentityCredentialProperties_ARM instances for property testing - lazily instantiated by
// FederatedIdentityCredentialProperties_ARMGenerator()
var federatedIdentityCredentialProperties_ARMGenerator gopter.Gen

// FederatedIdentityCredentialProperties_ARMGenerator returns a generator of FederatedIdentityCredentialProperties_ARM instances for property testing.
func FederatedIdentityCredentialProperties_ARMGenerator() gopter.Gen {
	if federatedIdentityCredentialProperties_ARMGenerator != nil {
		return federatedIdentityCredentialProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_ARM(generators)
	federatedIdentityCredentialProperties_ARMGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredentialProperties_ARM{}), generators)

	return federatedIdentityCredentialProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_ARM(gens map[string]gopter.Gen) {
	gens["Audiences"] = gen.SliceOf(gen.AlphaString())
	gens["Issuer"] = gen.PtrOf(gen.AlphaString())
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
}

func Test_FederatedIdentityCredential_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FederatedIdentityCredential_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFederatedIdentityCredential_Spec_ARM, FederatedIdentityCredential_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFederatedIdentityCredential_Spec_ARM runs a test to see if a specific instance of FederatedIdentityCredential_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFederatedIdentityCredential_Spec_ARM(subject FederatedIdentityCredential_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FederatedIdentityCredential_Spec_ARM
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

// Generator of FederatedIdentityCredential_Spec_ARM instances for property testing - lazily instantiated by
// FederatedIdentityCredential_Spec_ARMGenerator()
var federatedIdentityCredential_Spec_ARMGenerator gopter.Gen

// FederatedIdentityCredential_Spec_ARMGenerator returns a generator of FederatedIdentityCredential_Spec_ARM instances for property testing.
// We first initialize federatedIdentityCredential_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FederatedIdentityCredential_Spec_ARMGenerator() gopter.Gen {
	if federatedIdentityCredential_Spec_ARMGenerator != nil {
		return federatedIdentityCredential_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredential_Spec_ARM(generators)
	federatedIdentityCredential_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredential_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredential_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForFederatedIdentityCredential_Spec_ARM(generators)
	federatedIdentityCredential_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredential_Spec_ARM{}), generators)

	return federatedIdentityCredential_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFederatedIdentityCredential_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFederatedIdentityCredential_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForFederatedIdentityCredential_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFederatedIdentityCredential_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FederatedIdentityCredentialProperties_ARMGenerator())
}
