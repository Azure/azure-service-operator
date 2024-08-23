// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

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

func Test_Profiles_SecurityPolicy_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_SecurityPolicy_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_SecurityPolicy_Spec_ARM, Profiles_SecurityPolicy_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_SecurityPolicy_Spec_ARM runs a test to see if a specific instance of Profiles_SecurityPolicy_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_SecurityPolicy_Spec_ARM(subject Profiles_SecurityPolicy_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_SecurityPolicy_Spec_ARM
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

// Generator of Profiles_SecurityPolicy_Spec_ARM instances for property testing - lazily instantiated by
// Profiles_SecurityPolicy_Spec_ARMGenerator()
var profiles_SecurityPolicy_Spec_ARMGenerator gopter.Gen

// Profiles_SecurityPolicy_Spec_ARMGenerator returns a generator of Profiles_SecurityPolicy_Spec_ARM instances for property testing.
// We first initialize profiles_SecurityPolicy_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_SecurityPolicy_Spec_ARMGenerator() gopter.Gen {
	if profiles_SecurityPolicy_Spec_ARMGenerator != nil {
		return profiles_SecurityPolicy_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_SecurityPolicy_Spec_ARM(generators)
	profiles_SecurityPolicy_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_SecurityPolicy_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_SecurityPolicy_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForProfiles_SecurityPolicy_Spec_ARM(generators)
	profiles_SecurityPolicy_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_SecurityPolicy_Spec_ARM{}), generators)

	return profiles_SecurityPolicy_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_SecurityPolicy_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_SecurityPolicy_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForProfiles_SecurityPolicy_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_SecurityPolicy_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SecurityPolicyProperties_ARMGenerator())
}

func Test_SecurityPolicyPropertiesParameters_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicyPropertiesParameters_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicyPropertiesParameters_ARM, SecurityPolicyPropertiesParameters_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicyPropertiesParameters_ARM runs a test to see if a specific instance of SecurityPolicyPropertiesParameters_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicyPropertiesParameters_ARM(subject SecurityPolicyPropertiesParameters_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicyPropertiesParameters_ARM
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

// Generator of SecurityPolicyPropertiesParameters_ARM instances for property testing - lazily instantiated by
// SecurityPolicyPropertiesParameters_ARMGenerator()
var securityPolicyPropertiesParameters_ARMGenerator gopter.Gen

// SecurityPolicyPropertiesParameters_ARMGenerator returns a generator of SecurityPolicyPropertiesParameters_ARM instances for property testing.
func SecurityPolicyPropertiesParameters_ARMGenerator() gopter.Gen {
	if securityPolicyPropertiesParameters_ARMGenerator != nil {
		return securityPolicyPropertiesParameters_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecurityPolicyPropertiesParameters_ARM(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(SecurityPolicyPropertiesParameters_ARM{}), map[string]gopter.Gen{propName: propGen}))
	}
	securityPolicyPropertiesParameters_ARMGenerator = gen.OneGenOf(gens...)

	return securityPolicyPropertiesParameters_ARMGenerator
}

// AddRelatedPropertyGeneratorsForSecurityPolicyPropertiesParameters_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicyPropertiesParameters_ARM(gens map[string]gopter.Gen) {
	gens["WebApplicationFirewall"] = SecurityPolicyWebApplicationFirewallParameters_ARMGenerator().Map(func(it SecurityPolicyWebApplicationFirewallParameters_ARM) *SecurityPolicyWebApplicationFirewallParameters_ARM {
		return &it
	}) // generate one case for OneOf type
}

func Test_SecurityPolicyProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicyProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicyProperties_ARM, SecurityPolicyProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicyProperties_ARM runs a test to see if a specific instance of SecurityPolicyProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicyProperties_ARM(subject SecurityPolicyProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicyProperties_ARM
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

// Generator of SecurityPolicyProperties_ARM instances for property testing - lazily instantiated by
// SecurityPolicyProperties_ARMGenerator()
var securityPolicyProperties_ARMGenerator gopter.Gen

// SecurityPolicyProperties_ARMGenerator returns a generator of SecurityPolicyProperties_ARM instances for property testing.
func SecurityPolicyProperties_ARMGenerator() gopter.Gen {
	if securityPolicyProperties_ARMGenerator != nil {
		return securityPolicyProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecurityPolicyProperties_ARM(generators)
	securityPolicyProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyProperties_ARM{}), generators)

	return securityPolicyProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForSecurityPolicyProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicyProperties_ARM(gens map[string]gopter.Gen) {
	gens["Parameters"] = gen.PtrOf(SecurityPolicyPropertiesParameters_ARMGenerator())
}

func Test_SecurityPolicyWebApplicationFirewallAssociation_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicyWebApplicationFirewallAssociation_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallAssociation_ARM, SecurityPolicyWebApplicationFirewallAssociation_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallAssociation_ARM runs a test to see if a specific instance of SecurityPolicyWebApplicationFirewallAssociation_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallAssociation_ARM(subject SecurityPolicyWebApplicationFirewallAssociation_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicyWebApplicationFirewallAssociation_ARM
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

// Generator of SecurityPolicyWebApplicationFirewallAssociation_ARM instances for property testing - lazily instantiated
// by SecurityPolicyWebApplicationFirewallAssociation_ARMGenerator()
var securityPolicyWebApplicationFirewallAssociation_ARMGenerator gopter.Gen

// SecurityPolicyWebApplicationFirewallAssociation_ARMGenerator returns a generator of SecurityPolicyWebApplicationFirewallAssociation_ARM instances for property testing.
// We first initialize securityPolicyWebApplicationFirewallAssociation_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityPolicyWebApplicationFirewallAssociation_ARMGenerator() gopter.Gen {
	if securityPolicyWebApplicationFirewallAssociation_ARMGenerator != nil {
		return securityPolicyWebApplicationFirewallAssociation_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_ARM(generators)
	securityPolicyWebApplicationFirewallAssociation_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallAssociation_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_ARM(generators)
	AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_ARM(generators)
	securityPolicyWebApplicationFirewallAssociation_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallAssociation_ARM{}), generators)

	return securityPolicyWebApplicationFirewallAssociation_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_ARM(gens map[string]gopter.Gen) {
	gens["PatternsToMatch"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_ARM(gens map[string]gopter.Gen) {
	gens["Domains"] = gen.SliceOf(ActivatedResourceReference_ARMGenerator())
}

func Test_SecurityPolicyWebApplicationFirewallParameters_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicyWebApplicationFirewallParameters_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallParameters_ARM, SecurityPolicyWebApplicationFirewallParameters_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallParameters_ARM runs a test to see if a specific instance of SecurityPolicyWebApplicationFirewallParameters_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallParameters_ARM(subject SecurityPolicyWebApplicationFirewallParameters_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicyWebApplicationFirewallParameters_ARM
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

// Generator of SecurityPolicyWebApplicationFirewallParameters_ARM instances for property testing - lazily instantiated
// by SecurityPolicyWebApplicationFirewallParameters_ARMGenerator()
var securityPolicyWebApplicationFirewallParameters_ARMGenerator gopter.Gen

// SecurityPolicyWebApplicationFirewallParameters_ARMGenerator returns a generator of SecurityPolicyWebApplicationFirewallParameters_ARM instances for property testing.
// We first initialize securityPolicyWebApplicationFirewallParameters_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityPolicyWebApplicationFirewallParameters_ARMGenerator() gopter.Gen {
	if securityPolicyWebApplicationFirewallParameters_ARMGenerator != nil {
		return securityPolicyWebApplicationFirewallParameters_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_ARM(generators)
	securityPolicyWebApplicationFirewallParameters_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallParameters_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_ARM(generators)
	AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_ARM(generators)
	securityPolicyWebApplicationFirewallParameters_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallParameters_ARM{}), generators)

	return securityPolicyWebApplicationFirewallParameters_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_ARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.OneConstOf(SecurityPolicyWebApplicationFirewallParameters_Type_ARM_WebApplicationFirewall)
}

// AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_ARM(gens map[string]gopter.Gen) {
	gens["Associations"] = gen.SliceOf(SecurityPolicyWebApplicationFirewallAssociation_ARMGenerator())
	gens["WafPolicy"] = gen.PtrOf(ResourceReference_ARMGenerator())
}
