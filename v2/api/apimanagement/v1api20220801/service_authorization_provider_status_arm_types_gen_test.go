// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

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

func Test_AuthorizationProviderContractProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProviderContractProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProviderContractProperties_STATUS_ARM, AuthorizationProviderContractProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProviderContractProperties_STATUS_ARM runs a test to see if a specific instance of AuthorizationProviderContractProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProviderContractProperties_STATUS_ARM(subject AuthorizationProviderContractProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProviderContractProperties_STATUS_ARM
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

// Generator of AuthorizationProviderContractProperties_STATUS_ARM instances for property testing - lazily instantiated
// by AuthorizationProviderContractProperties_STATUS_ARMGenerator()
var authorizationProviderContractProperties_STATUS_ARMGenerator gopter.Gen

// AuthorizationProviderContractProperties_STATUS_ARMGenerator returns a generator of AuthorizationProviderContractProperties_STATUS_ARM instances for property testing.
// We first initialize authorizationProviderContractProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationProviderContractProperties_STATUS_ARMGenerator() gopter.Gen {
	if authorizationProviderContractProperties_STATUS_ARMGenerator != nil {
		return authorizationProviderContractProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProviderContractProperties_STATUS_ARM(generators)
	authorizationProviderContractProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationProviderContractProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProviderContractProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForAuthorizationProviderContractProperties_STATUS_ARM(generators)
	authorizationProviderContractProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationProviderContractProperties_STATUS_ARM{}), generators)

	return authorizationProviderContractProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationProviderContractProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationProviderContractProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["IdentityProvider"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationProviderContractProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProviderContractProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Oauth2"] = gen.PtrOf(AuthorizationProviderOAuth2Settings_STATUS_ARMGenerator())
}

func Test_AuthorizationProviderOAuth2GrantTypes_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProviderOAuth2GrantTypes_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProviderOAuth2GrantTypes_STATUS_ARM, AuthorizationProviderOAuth2GrantTypes_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProviderOAuth2GrantTypes_STATUS_ARM runs a test to see if a specific instance of AuthorizationProviderOAuth2GrantTypes_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProviderOAuth2GrantTypes_STATUS_ARM(subject AuthorizationProviderOAuth2GrantTypes_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProviderOAuth2GrantTypes_STATUS_ARM
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

// Generator of AuthorizationProviderOAuth2GrantTypes_STATUS_ARM instances for property testing - lazily instantiated by
// AuthorizationProviderOAuth2GrantTypes_STATUS_ARMGenerator()
var authorizationProviderOAuth2GrantTypes_STATUS_ARMGenerator gopter.Gen

// AuthorizationProviderOAuth2GrantTypes_STATUS_ARMGenerator returns a generator of AuthorizationProviderOAuth2GrantTypes_STATUS_ARM instances for property testing.
func AuthorizationProviderOAuth2GrantTypes_STATUS_ARMGenerator() gopter.Gen {
	if authorizationProviderOAuth2GrantTypes_STATUS_ARMGenerator != nil {
		return authorizationProviderOAuth2GrantTypes_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2GrantTypes_STATUS_ARM(generators)
	authorizationProviderOAuth2GrantTypes_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationProviderOAuth2GrantTypes_STATUS_ARM{}), generators)

	return authorizationProviderOAuth2GrantTypes_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2GrantTypes_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2GrantTypes_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AuthorizationCode"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["ClientCredentials"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

func Test_AuthorizationProviderOAuth2Settings_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProviderOAuth2Settings_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProviderOAuth2Settings_STATUS_ARM, AuthorizationProviderOAuth2Settings_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProviderOAuth2Settings_STATUS_ARM runs a test to see if a specific instance of AuthorizationProviderOAuth2Settings_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProviderOAuth2Settings_STATUS_ARM(subject AuthorizationProviderOAuth2Settings_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProviderOAuth2Settings_STATUS_ARM
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

// Generator of AuthorizationProviderOAuth2Settings_STATUS_ARM instances for property testing - lazily instantiated by
// AuthorizationProviderOAuth2Settings_STATUS_ARMGenerator()
var authorizationProviderOAuth2Settings_STATUS_ARMGenerator gopter.Gen

// AuthorizationProviderOAuth2Settings_STATUS_ARMGenerator returns a generator of AuthorizationProviderOAuth2Settings_STATUS_ARM instances for property testing.
// We first initialize authorizationProviderOAuth2Settings_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationProviderOAuth2Settings_STATUS_ARMGenerator() gopter.Gen {
	if authorizationProviderOAuth2Settings_STATUS_ARMGenerator != nil {
		return authorizationProviderOAuth2Settings_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS_ARM(generators)
	authorizationProviderOAuth2Settings_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationProviderOAuth2Settings_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS_ARM(generators)
	authorizationProviderOAuth2Settings_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationProviderOAuth2Settings_STATUS_ARM{}), generators)

	return authorizationProviderOAuth2Settings_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["RedirectUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["GrantTypes"] = gen.PtrOf(AuthorizationProviderOAuth2GrantTypes_STATUS_ARMGenerator())
}

func Test_Service_AuthorizationProvider_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_AuthorizationProvider_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_AuthorizationProvider_STATUS_ARM, Service_AuthorizationProvider_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_AuthorizationProvider_STATUS_ARM runs a test to see if a specific instance of Service_AuthorizationProvider_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForService_AuthorizationProvider_STATUS_ARM(subject Service_AuthorizationProvider_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_AuthorizationProvider_STATUS_ARM
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

// Generator of Service_AuthorizationProvider_STATUS_ARM instances for property testing - lazily instantiated by
// Service_AuthorizationProvider_STATUS_ARMGenerator()
var service_AuthorizationProvider_STATUS_ARMGenerator gopter.Gen

// Service_AuthorizationProvider_STATUS_ARMGenerator returns a generator of Service_AuthorizationProvider_STATUS_ARM instances for property testing.
// We first initialize service_AuthorizationProvider_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_AuthorizationProvider_STATUS_ARMGenerator() gopter.Gen {
	if service_AuthorizationProvider_STATUS_ARMGenerator != nil {
		return service_AuthorizationProvider_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProvider_STATUS_ARM(generators)
	service_AuthorizationProvider_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProvider_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProvider_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForService_AuthorizationProvider_STATUS_ARM(generators)
	service_AuthorizationProvider_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProvider_STATUS_ARM{}), generators)

	return service_AuthorizationProvider_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForService_AuthorizationProvider_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_AuthorizationProvider_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForService_AuthorizationProvider_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_AuthorizationProvider_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AuthorizationProviderContractProperties_STATUS_ARMGenerator())
}
