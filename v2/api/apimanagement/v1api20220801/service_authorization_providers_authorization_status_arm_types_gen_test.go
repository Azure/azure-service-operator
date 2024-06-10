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

func Test_AuthorizationContractProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationContractProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationContractProperties_STATUS_ARM, AuthorizationContractProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationContractProperties_STATUS_ARM runs a test to see if a specific instance of AuthorizationContractProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationContractProperties_STATUS_ARM(subject AuthorizationContractProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationContractProperties_STATUS_ARM
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

// Generator of AuthorizationContractProperties_STATUS_ARM instances for property testing - lazily instantiated by
// AuthorizationContractProperties_STATUS_ARMGenerator()
var authorizationContractProperties_STATUS_ARMGenerator gopter.Gen

// AuthorizationContractProperties_STATUS_ARMGenerator returns a generator of AuthorizationContractProperties_STATUS_ARM instances for property testing.
// We first initialize authorizationContractProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationContractProperties_STATUS_ARMGenerator() gopter.Gen {
	if authorizationContractProperties_STATUS_ARMGenerator != nil {
		return authorizationContractProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationContractProperties_STATUS_ARM(generators)
	authorizationContractProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationContractProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationContractProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForAuthorizationContractProperties_STATUS_ARM(generators)
	authorizationContractProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationContractProperties_STATUS_ARM{}), generators)

	return authorizationContractProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationContractProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationContractProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AuthorizationType"] = gen.PtrOf(gen.OneConstOf(AuthorizationContractProperties_AuthorizationType_STATUS_OAuth2))
	gens["Oauth2GrantType"] = gen.PtrOf(gen.OneConstOf(AuthorizationContractProperties_Oauth2GrantType_STATUS_AuthorizationCode, AuthorizationContractProperties_Oauth2GrantType_STATUS_ClientCredentials))
	gens["Parameters"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationContractProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationContractProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Error"] = gen.PtrOf(AuthorizationError_STATUS_ARMGenerator())
}

func Test_AuthorizationError_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationError_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationError_STATUS_ARM, AuthorizationError_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationError_STATUS_ARM runs a test to see if a specific instance of AuthorizationError_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationError_STATUS_ARM(subject AuthorizationError_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationError_STATUS_ARM
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

// Generator of AuthorizationError_STATUS_ARM instances for property testing - lazily instantiated by
// AuthorizationError_STATUS_ARMGenerator()
var authorizationError_STATUS_ARMGenerator gopter.Gen

// AuthorizationError_STATUS_ARMGenerator returns a generator of AuthorizationError_STATUS_ARM instances for property testing.
func AuthorizationError_STATUS_ARMGenerator() gopter.Gen {
	if authorizationError_STATUS_ARMGenerator != nil {
		return authorizationError_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationError_STATUS_ARM(generators)
	authorizationError_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationError_STATUS_ARM{}), generators)

	return authorizationError_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationError_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationError_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
}

func Test_Service_AuthorizationProviders_Authorization_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_AuthorizationProviders_Authorization_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_AuthorizationProviders_Authorization_STATUS_ARM, Service_AuthorizationProviders_Authorization_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_AuthorizationProviders_Authorization_STATUS_ARM runs a test to see if a specific instance of Service_AuthorizationProviders_Authorization_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForService_AuthorizationProviders_Authorization_STATUS_ARM(subject Service_AuthorizationProviders_Authorization_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_AuthorizationProviders_Authorization_STATUS_ARM
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

// Generator of Service_AuthorizationProviders_Authorization_STATUS_ARM instances for property testing - lazily
// instantiated by Service_AuthorizationProviders_Authorization_STATUS_ARMGenerator()
var service_AuthorizationProviders_Authorization_STATUS_ARMGenerator gopter.Gen

// Service_AuthorizationProviders_Authorization_STATUS_ARMGenerator returns a generator of Service_AuthorizationProviders_Authorization_STATUS_ARM instances for property testing.
// We first initialize service_AuthorizationProviders_Authorization_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_AuthorizationProviders_Authorization_STATUS_ARMGenerator() gopter.Gen {
	if service_AuthorizationProviders_Authorization_STATUS_ARMGenerator != nil {
		return service_AuthorizationProviders_Authorization_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorization_STATUS_ARM(generators)
	service_AuthorizationProviders_Authorization_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProviders_Authorization_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorization_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForService_AuthorizationProviders_Authorization_STATUS_ARM(generators)
	service_AuthorizationProviders_Authorization_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProviders_Authorization_STATUS_ARM{}), generators)

	return service_AuthorizationProviders_Authorization_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorization_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorization_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForService_AuthorizationProviders_Authorization_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_AuthorizationProviders_Authorization_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AuthorizationContractProperties_STATUS_ARMGenerator())
}
