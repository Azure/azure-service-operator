// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

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

func Test_NamespacesAuthorizationRules_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRules_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRulesSpecARM, NamespacesAuthorizationRulesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRulesSpecARM runs a test to see if a specific instance of NamespacesAuthorizationRules_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRulesSpecARM(subject NamespacesAuthorizationRules_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRules_SpecARM
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

// Generator of NamespacesAuthorizationRules_SpecARM instances for property testing - lazily instantiated by
// NamespacesAuthorizationRulesSpecARMGenerator()
var namespacesAuthorizationRulesSpecARMGenerator gopter.Gen

// NamespacesAuthorizationRulesSpecARMGenerator returns a generator of NamespacesAuthorizationRules_SpecARM instances for property testing.
// We first initialize namespacesAuthorizationRulesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesAuthorizationRulesSpecARMGenerator() gopter.Gen {
	if namespacesAuthorizationRulesSpecARMGenerator != nil {
		return namespacesAuthorizationRulesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSpecARM(generators)
	namespacesAuthorizationRulesSpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRules_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSpecARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRulesSpecARM(generators)
	namespacesAuthorizationRulesSpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRules_SpecARM{}), generators)

	return namespacesAuthorizationRulesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRulesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRulesSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AuthorizationRulePropertiesARMGenerator())
}

func Test_AuthorizationRulePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationRulePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationRulePropertiesARM, AuthorizationRulePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationRulePropertiesARM runs a test to see if a specific instance of AuthorizationRulePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationRulePropertiesARM(subject AuthorizationRulePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationRulePropertiesARM
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

// Generator of AuthorizationRulePropertiesARM instances for property testing - lazily instantiated by
// AuthorizationRulePropertiesARMGenerator()
var authorizationRulePropertiesARMGenerator gopter.Gen

// AuthorizationRulePropertiesARMGenerator returns a generator of AuthorizationRulePropertiesARM instances for property testing.
func AuthorizationRulePropertiesARMGenerator() gopter.Gen {
	if authorizationRulePropertiesARMGenerator != nil {
		return authorizationRulePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRulePropertiesARM(generators)
	authorizationRulePropertiesARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationRulePropertiesARM{}), generators)

	return authorizationRulePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationRulePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationRulePropertiesARM(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(AuthorizationRulePropertiesRights_Listen, AuthorizationRulePropertiesRights_Manage, AuthorizationRulePropertiesRights_Send))
}
