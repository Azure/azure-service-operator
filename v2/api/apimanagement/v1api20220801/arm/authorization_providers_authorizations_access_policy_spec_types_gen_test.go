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

func Test_AuthorizationAccessPolicyContractProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationAccessPolicyContractProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationAccessPolicyContractProperties, AuthorizationAccessPolicyContractPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationAccessPolicyContractProperties runs a test to see if a specific instance of AuthorizationAccessPolicyContractProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationAccessPolicyContractProperties(subject AuthorizationAccessPolicyContractProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationAccessPolicyContractProperties
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

// Generator of AuthorizationAccessPolicyContractProperties instances for property testing - lazily instantiated by
// AuthorizationAccessPolicyContractPropertiesGenerator()
var authorizationAccessPolicyContractPropertiesGenerator gopter.Gen

// AuthorizationAccessPolicyContractPropertiesGenerator returns a generator of AuthorizationAccessPolicyContractProperties instances for property testing.
func AuthorizationAccessPolicyContractPropertiesGenerator() gopter.Gen {
	if authorizationAccessPolicyContractPropertiesGenerator != nil {
		return authorizationAccessPolicyContractPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationAccessPolicyContractProperties(generators)
	authorizationAccessPolicyContractPropertiesGenerator = gen.Struct(reflect.TypeOf(AuthorizationAccessPolicyContractProperties{}), generators)

	return authorizationAccessPolicyContractPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationAccessPolicyContractProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationAccessPolicyContractProperties(gens map[string]gopter.Gen) {
	gens["ObjectId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_AuthorizationProvidersAuthorizationsAccessPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProvidersAuthorizationsAccessPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy_Spec, AuthorizationProvidersAuthorizationsAccessPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy_Spec runs a test to see if a specific instance of AuthorizationProvidersAuthorizationsAccessPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy_Spec(subject AuthorizationProvidersAuthorizationsAccessPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProvidersAuthorizationsAccessPolicy_Spec
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

// Generator of AuthorizationProvidersAuthorizationsAccessPolicy_Spec instances for property testing - lazily
// instantiated by AuthorizationProvidersAuthorizationsAccessPolicy_SpecGenerator()
var authorizationProvidersAuthorizationsAccessPolicy_SpecGenerator gopter.Gen

// AuthorizationProvidersAuthorizationsAccessPolicy_SpecGenerator returns a generator of AuthorizationProvidersAuthorizationsAccessPolicy_Spec instances for property testing.
// We first initialize authorizationProvidersAuthorizationsAccessPolicy_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationProvidersAuthorizationsAccessPolicy_SpecGenerator() gopter.Gen {
	if authorizationProvidersAuthorizationsAccessPolicy_SpecGenerator != nil {
		return authorizationProvidersAuthorizationsAccessPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_Spec(generators)
	authorizationProvidersAuthorizationsAccessPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorizationsAccessPolicy_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_Spec(generators)
	AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_Spec(generators)
	authorizationProvidersAuthorizationsAccessPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorizationsAccessPolicy_Spec{}), generators)

	return authorizationProvidersAuthorizationsAccessPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AuthorizationAccessPolicyContractPropertiesGenerator())
}
