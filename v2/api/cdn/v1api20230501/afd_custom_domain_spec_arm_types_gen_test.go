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

func Test_AFDDomainHttpsParameters_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AFDDomainHttpsParameters_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAFDDomainHttpsParameters_ARM, AFDDomainHttpsParameters_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAFDDomainHttpsParameters_ARM runs a test to see if a specific instance of AFDDomainHttpsParameters_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAFDDomainHttpsParameters_ARM(subject AFDDomainHttpsParameters_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AFDDomainHttpsParameters_ARM
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

// Generator of AFDDomainHttpsParameters_ARM instances for property testing - lazily instantiated by
// AFDDomainHttpsParameters_ARMGenerator()
var afdDomainHttpsParameters_ARMGenerator gopter.Gen

// AFDDomainHttpsParameters_ARMGenerator returns a generator of AFDDomainHttpsParameters_ARM instances for property testing.
// We first initialize afdDomainHttpsParameters_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AFDDomainHttpsParameters_ARMGenerator() gopter.Gen {
	if afdDomainHttpsParameters_ARMGenerator != nil {
		return afdDomainHttpsParameters_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDDomainHttpsParameters_ARM(generators)
	afdDomainHttpsParameters_ARMGenerator = gen.Struct(reflect.TypeOf(AFDDomainHttpsParameters_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDDomainHttpsParameters_ARM(generators)
	AddRelatedPropertyGeneratorsForAFDDomainHttpsParameters_ARM(generators)
	afdDomainHttpsParameters_ARMGenerator = gen.Struct(reflect.TypeOf(AFDDomainHttpsParameters_ARM{}), generators)

	return afdDomainHttpsParameters_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAFDDomainHttpsParameters_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAFDDomainHttpsParameters_ARM(gens map[string]gopter.Gen) {
	gens["CertificateType"] = gen.PtrOf(gen.OneConstOf(AFDDomainHttpsParameters_CertificateType_ARM_AzureFirstPartyManagedCertificate, AFDDomainHttpsParameters_CertificateType_ARM_CustomerCertificate, AFDDomainHttpsParameters_CertificateType_ARM_ManagedCertificate))
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(AFDDomainHttpsParameters_MinimumTlsVersion_ARM_TLS10, AFDDomainHttpsParameters_MinimumTlsVersion_ARM_TLS12))
}

// AddRelatedPropertyGeneratorsForAFDDomainHttpsParameters_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAFDDomainHttpsParameters_ARM(gens map[string]gopter.Gen) {
	gens["Secret"] = gen.PtrOf(ResourceReference_ARMGenerator())
}

func Test_AFDDomainProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AFDDomainProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAFDDomainProperties_ARM, AFDDomainProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAFDDomainProperties_ARM runs a test to see if a specific instance of AFDDomainProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAFDDomainProperties_ARM(subject AFDDomainProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AFDDomainProperties_ARM
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

// Generator of AFDDomainProperties_ARM instances for property testing - lazily instantiated by
// AFDDomainProperties_ARMGenerator()
var afdDomainProperties_ARMGenerator gopter.Gen

// AFDDomainProperties_ARMGenerator returns a generator of AFDDomainProperties_ARM instances for property testing.
// We first initialize afdDomainProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AFDDomainProperties_ARMGenerator() gopter.Gen {
	if afdDomainProperties_ARMGenerator != nil {
		return afdDomainProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDDomainProperties_ARM(generators)
	afdDomainProperties_ARMGenerator = gen.Struct(reflect.TypeOf(AFDDomainProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDDomainProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForAFDDomainProperties_ARM(generators)
	afdDomainProperties_ARMGenerator = gen.Struct(reflect.TypeOf(AFDDomainProperties_ARM{}), generators)

	return afdDomainProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAFDDomainProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAFDDomainProperties_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedProperties"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAFDDomainProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAFDDomainProperties_ARM(gens map[string]gopter.Gen) {
	gens["AzureDnsZone"] = gen.PtrOf(ResourceReference_ARMGenerator())
	gens["PreValidatedCustomDomainResourceId"] = gen.PtrOf(ResourceReference_ARMGenerator())
	gens["TlsSettings"] = gen.PtrOf(AFDDomainHttpsParameters_ARMGenerator())
}

func Test_AfdCustomDomain_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdCustomDomain_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdCustomDomain_Spec_ARM, AfdCustomDomain_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdCustomDomain_Spec_ARM runs a test to see if a specific instance of AfdCustomDomain_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdCustomDomain_Spec_ARM(subject AfdCustomDomain_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdCustomDomain_Spec_ARM
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

// Generator of AfdCustomDomain_Spec_ARM instances for property testing - lazily instantiated by
// AfdCustomDomain_Spec_ARMGenerator()
var afdCustomDomain_Spec_ARMGenerator gopter.Gen

// AfdCustomDomain_Spec_ARMGenerator returns a generator of AfdCustomDomain_Spec_ARM instances for property testing.
// We first initialize afdCustomDomain_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AfdCustomDomain_Spec_ARMGenerator() gopter.Gen {
	if afdCustomDomain_Spec_ARMGenerator != nil {
		return afdCustomDomain_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdCustomDomain_Spec_ARM(generators)
	afdCustomDomain_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(AfdCustomDomain_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdCustomDomain_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForAfdCustomDomain_Spec_ARM(generators)
	afdCustomDomain_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(AfdCustomDomain_Spec_ARM{}), generators)

	return afdCustomDomain_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAfdCustomDomain_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAfdCustomDomain_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForAfdCustomDomain_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAfdCustomDomain_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AFDDomainProperties_ARMGenerator())
}

func Test_ResourceReference_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceReference_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceReference_ARM, ResourceReference_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceReference_ARM runs a test to see if a specific instance of ResourceReference_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceReference_ARM(subject ResourceReference_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceReference_ARM
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

// Generator of ResourceReference_ARM instances for property testing - lazily instantiated by
// ResourceReference_ARMGenerator()
var resourceReference_ARMGenerator gopter.Gen

// ResourceReference_ARMGenerator returns a generator of ResourceReference_ARM instances for property testing.
func ResourceReference_ARMGenerator() gopter.Gen {
	if resourceReference_ARMGenerator != nil {
		return resourceReference_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceReference_ARM(generators)
	resourceReference_ARMGenerator = gen.Struct(reflect.TypeOf(ResourceReference_ARM{}), generators)

	return resourceReference_ARMGenerator
}

// AddIndependentPropertyGeneratorsForResourceReference_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceReference_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}