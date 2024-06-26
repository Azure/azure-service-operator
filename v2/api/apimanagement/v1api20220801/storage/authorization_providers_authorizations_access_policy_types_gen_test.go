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

func Test_AuthorizationProvidersAuthorizationsAccessPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProvidersAuthorizationsAccessPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy, AuthorizationProvidersAuthorizationsAccessPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy runs a test to see if a specific instance of AuthorizationProvidersAuthorizationsAccessPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy(subject AuthorizationProvidersAuthorizationsAccessPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProvidersAuthorizationsAccessPolicy
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

// Generator of AuthorizationProvidersAuthorizationsAccessPolicy instances for property testing - lazily instantiated by
// AuthorizationProvidersAuthorizationsAccessPolicyGenerator()
var authorizationProvidersAuthorizationsAccessPolicyGenerator gopter.Gen

// AuthorizationProvidersAuthorizationsAccessPolicyGenerator returns a generator of AuthorizationProvidersAuthorizationsAccessPolicy instances for property testing.
func AuthorizationProvidersAuthorizationsAccessPolicyGenerator() gopter.Gen {
	if authorizationProvidersAuthorizationsAccessPolicyGenerator != nil {
		return authorizationProvidersAuthorizationsAccessPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy(generators)
	authorizationProvidersAuthorizationsAccessPolicyGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorizationsAccessPolicy{}), generators)

	return authorizationProvidersAuthorizationsAccessPolicyGenerator
}

// AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = Service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator()
	gens["Status"] = Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator()
}

func Test_Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS, Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS runs a test to see if a specific instance of Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS(subject Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS
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

// Generator of Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS instances for property testing -
// lazily instantiated by Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator()
var service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator gopter.Gen

// Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator returns a generator of Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS instances for property testing.
func Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator() gopter.Gen {
	if service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator != nil {
		return service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS(generators)
	service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS{}), generators)

	return service_AuthorizationProviders_Authorizations_AccessPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ObjectId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec, Service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec runs a test to see if a specific instance of Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec(subject Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec
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

// Generator of Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec instances for property testing - lazily
// instantiated by Service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator()
var service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator gopter.Gen

// Service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator returns a generator of Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec instances for property testing.
func Service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator() gopter.Gen {
	if service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator != nil {
		return service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec(generators)
	service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec{}), generators)

	return service_AuthorizationProviders_Authorizations_AccessPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorizations_AccessPolicy_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["ObjectId"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}
