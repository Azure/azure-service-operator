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

func Test_ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded, ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded runs a test to see if a specific instance of ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded(subject ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded
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

// Generator of ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded instances for property
// testing - lazily instantiated by ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbeddedGenerator()
var activatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbeddedGenerator gopter.Gen

// ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbeddedGenerator returns a generator of ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded instances for property testing.
func ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbeddedGenerator() gopter.Gen {
	if activatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbeddedGenerator != nil {
		return activatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded(generators)
	activatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded{}), generators)

	return activatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_SecurityPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicy, SecurityPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicy runs a test to see if a specific instance of SecurityPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicy(subject SecurityPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicy
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

// Generator of SecurityPolicy instances for property testing - lazily instantiated by SecurityPolicyGenerator()
var securityPolicyGenerator gopter.Gen

// SecurityPolicyGenerator returns a generator of SecurityPolicy instances for property testing.
func SecurityPolicyGenerator() gopter.Gen {
	if securityPolicyGenerator != nil {
		return securityPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecurityPolicy(generators)
	securityPolicyGenerator = gen.Struct(reflect.TypeOf(SecurityPolicy{}), generators)

	return securityPolicyGenerator
}

// AddRelatedPropertyGeneratorsForSecurityPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = SecurityPolicy_SpecGenerator()
	gens["Status"] = SecurityPolicy_STATUSGenerator()
}

func Test_SecurityPolicyPropertiesParameters_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicyPropertiesParameters via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicyPropertiesParameters, SecurityPolicyPropertiesParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicyPropertiesParameters runs a test to see if a specific instance of SecurityPolicyPropertiesParameters round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicyPropertiesParameters(subject SecurityPolicyPropertiesParameters) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicyPropertiesParameters
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

// Generator of SecurityPolicyPropertiesParameters instances for property testing - lazily instantiated by
// SecurityPolicyPropertiesParametersGenerator()
var securityPolicyPropertiesParametersGenerator gopter.Gen

// SecurityPolicyPropertiesParametersGenerator returns a generator of SecurityPolicyPropertiesParameters instances for property testing.
func SecurityPolicyPropertiesParametersGenerator() gopter.Gen {
	if securityPolicyPropertiesParametersGenerator != nil {
		return securityPolicyPropertiesParametersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecurityPolicyPropertiesParameters(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(SecurityPolicyPropertiesParameters{}), map[string]gopter.Gen{propName: propGen}))
	}
	securityPolicyPropertiesParametersGenerator = gen.OneGenOf(gens...)

	return securityPolicyPropertiesParametersGenerator
}

// AddRelatedPropertyGeneratorsForSecurityPolicyPropertiesParameters is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicyPropertiesParameters(gens map[string]gopter.Gen) {
	gens["WebApplicationFirewall"] = SecurityPolicyWebApplicationFirewallParametersGenerator().Map(func(it SecurityPolicyWebApplicationFirewallParameters) *SecurityPolicyWebApplicationFirewallParameters {
		return &it
	}) // generate one case for OneOf type
}

func Test_SecurityPolicyPropertiesParameters_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicyPropertiesParameters_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicyPropertiesParameters_STATUS, SecurityPolicyPropertiesParameters_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicyPropertiesParameters_STATUS runs a test to see if a specific instance of SecurityPolicyPropertiesParameters_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicyPropertiesParameters_STATUS(subject SecurityPolicyPropertiesParameters_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicyPropertiesParameters_STATUS
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

// Generator of SecurityPolicyPropertiesParameters_STATUS instances for property testing - lazily instantiated by
// SecurityPolicyPropertiesParameters_STATUSGenerator()
var securityPolicyPropertiesParameters_STATUSGenerator gopter.Gen

// SecurityPolicyPropertiesParameters_STATUSGenerator returns a generator of SecurityPolicyPropertiesParameters_STATUS instances for property testing.
func SecurityPolicyPropertiesParameters_STATUSGenerator() gopter.Gen {
	if securityPolicyPropertiesParameters_STATUSGenerator != nil {
		return securityPolicyPropertiesParameters_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecurityPolicyPropertiesParameters_STATUS(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(SecurityPolicyPropertiesParameters_STATUS{}), map[string]gopter.Gen{propName: propGen}))
	}
	securityPolicyPropertiesParameters_STATUSGenerator = gen.OneGenOf(gens...)

	return securityPolicyPropertiesParameters_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForSecurityPolicyPropertiesParameters_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicyPropertiesParameters_STATUS(gens map[string]gopter.Gen) {
	gens["WebApplicationFirewall"] = SecurityPolicyWebApplicationFirewallParameters_STATUSGenerator().Map(func(it SecurityPolicyWebApplicationFirewallParameters_STATUS) *SecurityPolicyWebApplicationFirewallParameters_STATUS {
		return &it
	}) // generate one case for OneOf type
}

func Test_SecurityPolicyWebApplicationFirewallAssociation_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicyWebApplicationFirewallAssociation via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallAssociation, SecurityPolicyWebApplicationFirewallAssociationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallAssociation runs a test to see if a specific instance of SecurityPolicyWebApplicationFirewallAssociation round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallAssociation(subject SecurityPolicyWebApplicationFirewallAssociation) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicyWebApplicationFirewallAssociation
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

// Generator of SecurityPolicyWebApplicationFirewallAssociation instances for property testing - lazily instantiated by
// SecurityPolicyWebApplicationFirewallAssociationGenerator()
var securityPolicyWebApplicationFirewallAssociationGenerator gopter.Gen

// SecurityPolicyWebApplicationFirewallAssociationGenerator returns a generator of SecurityPolicyWebApplicationFirewallAssociation instances for property testing.
// We first initialize securityPolicyWebApplicationFirewallAssociationGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityPolicyWebApplicationFirewallAssociationGenerator() gopter.Gen {
	if securityPolicyWebApplicationFirewallAssociationGenerator != nil {
		return securityPolicyWebApplicationFirewallAssociationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation(generators)
	securityPolicyWebApplicationFirewallAssociationGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallAssociation{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation(generators)
	AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation(generators)
	securityPolicyWebApplicationFirewallAssociationGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallAssociation{}), generators)

	return securityPolicyWebApplicationFirewallAssociationGenerator
}

// AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation(gens map[string]gopter.Gen) {
	gens["PatternsToMatch"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation(gens map[string]gopter.Gen) {
	gens["Domains"] = gen.SliceOf(ActivatedResourceReferenceGenerator())
}

func Test_SecurityPolicyWebApplicationFirewallAssociation_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicyWebApplicationFirewallAssociation_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallAssociation_STATUS, SecurityPolicyWebApplicationFirewallAssociation_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallAssociation_STATUS runs a test to see if a specific instance of SecurityPolicyWebApplicationFirewallAssociation_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallAssociation_STATUS(subject SecurityPolicyWebApplicationFirewallAssociation_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicyWebApplicationFirewallAssociation_STATUS
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

// Generator of SecurityPolicyWebApplicationFirewallAssociation_STATUS instances for property testing - lazily
// instantiated by SecurityPolicyWebApplicationFirewallAssociation_STATUSGenerator()
var securityPolicyWebApplicationFirewallAssociation_STATUSGenerator gopter.Gen

// SecurityPolicyWebApplicationFirewallAssociation_STATUSGenerator returns a generator of SecurityPolicyWebApplicationFirewallAssociation_STATUS instances for property testing.
// We first initialize securityPolicyWebApplicationFirewallAssociation_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityPolicyWebApplicationFirewallAssociation_STATUSGenerator() gopter.Gen {
	if securityPolicyWebApplicationFirewallAssociation_STATUSGenerator != nil {
		return securityPolicyWebApplicationFirewallAssociation_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_STATUS(generators)
	securityPolicyWebApplicationFirewallAssociation_STATUSGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallAssociation_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_STATUS(generators)
	AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_STATUS(generators)
	securityPolicyWebApplicationFirewallAssociation_STATUSGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallAssociation_STATUS{}), generators)

	return securityPolicyWebApplicationFirewallAssociation_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_STATUS(gens map[string]gopter.Gen) {
	gens["PatternsToMatch"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallAssociation_STATUS(gens map[string]gopter.Gen) {
	gens["Domains"] = gen.SliceOf(ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbeddedGenerator())
}

func Test_SecurityPolicyWebApplicationFirewallParameters_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicyWebApplicationFirewallParameters via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallParameters, SecurityPolicyWebApplicationFirewallParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallParameters runs a test to see if a specific instance of SecurityPolicyWebApplicationFirewallParameters round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallParameters(subject SecurityPolicyWebApplicationFirewallParameters) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicyWebApplicationFirewallParameters
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

// Generator of SecurityPolicyWebApplicationFirewallParameters instances for property testing - lazily instantiated by
// SecurityPolicyWebApplicationFirewallParametersGenerator()
var securityPolicyWebApplicationFirewallParametersGenerator gopter.Gen

// SecurityPolicyWebApplicationFirewallParametersGenerator returns a generator of SecurityPolicyWebApplicationFirewallParameters instances for property testing.
// We first initialize securityPolicyWebApplicationFirewallParametersGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityPolicyWebApplicationFirewallParametersGenerator() gopter.Gen {
	if securityPolicyWebApplicationFirewallParametersGenerator != nil {
		return securityPolicyWebApplicationFirewallParametersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters(generators)
	securityPolicyWebApplicationFirewallParametersGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallParameters{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters(generators)
	AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters(generators)
	securityPolicyWebApplicationFirewallParametersGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallParameters{}), generators)

	return securityPolicyWebApplicationFirewallParametersGenerator
}

// AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters(gens map[string]gopter.Gen) {
	gens["Associations"] = gen.SliceOf(SecurityPolicyWebApplicationFirewallAssociationGenerator())
	gens["WafPolicy"] = gen.PtrOf(ResourceReferenceGenerator())
}

func Test_SecurityPolicyWebApplicationFirewallParameters_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicyWebApplicationFirewallParameters_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallParameters_STATUS, SecurityPolicyWebApplicationFirewallParameters_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallParameters_STATUS runs a test to see if a specific instance of SecurityPolicyWebApplicationFirewallParameters_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicyWebApplicationFirewallParameters_STATUS(subject SecurityPolicyWebApplicationFirewallParameters_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicyWebApplicationFirewallParameters_STATUS
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

// Generator of SecurityPolicyWebApplicationFirewallParameters_STATUS instances for property testing - lazily
// instantiated by SecurityPolicyWebApplicationFirewallParameters_STATUSGenerator()
var securityPolicyWebApplicationFirewallParameters_STATUSGenerator gopter.Gen

// SecurityPolicyWebApplicationFirewallParameters_STATUSGenerator returns a generator of SecurityPolicyWebApplicationFirewallParameters_STATUS instances for property testing.
// We first initialize securityPolicyWebApplicationFirewallParameters_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityPolicyWebApplicationFirewallParameters_STATUSGenerator() gopter.Gen {
	if securityPolicyWebApplicationFirewallParameters_STATUSGenerator != nil {
		return securityPolicyWebApplicationFirewallParameters_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_STATUS(generators)
	securityPolicyWebApplicationFirewallParameters_STATUSGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallParameters_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_STATUS(generators)
	AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_STATUS(generators)
	securityPolicyWebApplicationFirewallParameters_STATUSGenerator = gen.Struct(reflect.TypeOf(SecurityPolicyWebApplicationFirewallParameters_STATUS{}), generators)

	return securityPolicyWebApplicationFirewallParameters_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_STATUS(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicyWebApplicationFirewallParameters_STATUS(gens map[string]gopter.Gen) {
	gens["Associations"] = gen.SliceOf(SecurityPolicyWebApplicationFirewallAssociation_STATUSGenerator())
	gens["WafPolicy"] = gen.PtrOf(ResourceReference_STATUSGenerator())
}

func Test_SecurityPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicy_STATUS, SecurityPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicy_STATUS runs a test to see if a specific instance of SecurityPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicy_STATUS(subject SecurityPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicy_STATUS
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

// Generator of SecurityPolicy_STATUS instances for property testing - lazily instantiated by
// SecurityPolicy_STATUSGenerator()
var securityPolicy_STATUSGenerator gopter.Gen

// SecurityPolicy_STATUSGenerator returns a generator of SecurityPolicy_STATUS instances for property testing.
// We first initialize securityPolicy_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityPolicy_STATUSGenerator() gopter.Gen {
	if securityPolicy_STATUSGenerator != nil {
		return securityPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicy_STATUS(generators)
	securityPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(SecurityPolicy_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicy_STATUS(generators)
	AddRelatedPropertyGeneratorsForSecurityPolicy_STATUS(generators)
	securityPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(SecurityPolicy_STATUS{}), generators)

	return securityPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSecurityPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["DeploymentStatus"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProfileName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityPolicy_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Parameters"] = gen.PtrOf(SecurityPolicyPropertiesParameters_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_SecurityPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityPolicy_Spec, SecurityPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityPolicy_Spec runs a test to see if a specific instance of SecurityPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityPolicy_Spec(subject SecurityPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityPolicy_Spec
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

// Generator of SecurityPolicy_Spec instances for property testing - lazily instantiated by
// SecurityPolicy_SpecGenerator()
var securityPolicy_SpecGenerator gopter.Gen

// SecurityPolicy_SpecGenerator returns a generator of SecurityPolicy_Spec instances for property testing.
// We first initialize securityPolicy_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityPolicy_SpecGenerator() gopter.Gen {
	if securityPolicy_SpecGenerator != nil {
		return securityPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicy_Spec(generators)
	securityPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(SecurityPolicy_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityPolicy_Spec(generators)
	AddRelatedPropertyGeneratorsForSecurityPolicy_Spec(generators)
	securityPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(SecurityPolicy_Spec{}), generators)

	return securityPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForSecurityPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityPolicy_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForSecurityPolicy_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityPolicy_Spec(gens map[string]gopter.Gen) {
	gens["Parameters"] = gen.PtrOf(SecurityPolicyPropertiesParametersGenerator())
}
