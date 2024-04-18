// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

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

func Test_NatGateway_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGateway_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGateway_Spec_ARM, NatGateway_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGateway_Spec_ARM runs a test to see if a specific instance of NatGateway_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGateway_Spec_ARM(subject NatGateway_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGateway_Spec_ARM
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

// Generator of NatGateway_Spec_ARM instances for property testing - lazily instantiated by
// NatGateway_Spec_ARMGenerator()
var natGateway_Spec_ARMGenerator gopter.Gen

// NatGateway_Spec_ARMGenerator returns a generator of NatGateway_Spec_ARM instances for property testing.
// We first initialize natGateway_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NatGateway_Spec_ARMGenerator() gopter.Gen {
	if natGateway_Spec_ARMGenerator != nil {
		return natGateway_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGateway_Spec_ARM(generators)
	natGateway_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NatGateway_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGateway_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNatGateway_Spec_ARM(generators)
	natGateway_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NatGateway_Spec_ARM{}), generators)

	return natGateway_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNatGateway_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGateway_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNatGateway_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNatGateway_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NatGatewayPropertiesFormat_ARMGenerator())
	gens["Sku"] = gen.PtrOf(NatGatewaySku_ARMGenerator())
}

func Test_NatGatewayPropertiesFormat_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGatewayPropertiesFormat_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGatewayPropertiesFormat_ARM, NatGatewayPropertiesFormat_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGatewayPropertiesFormat_ARM runs a test to see if a specific instance of NatGatewayPropertiesFormat_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGatewayPropertiesFormat_ARM(subject NatGatewayPropertiesFormat_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGatewayPropertiesFormat_ARM
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

// Generator of NatGatewayPropertiesFormat_ARM instances for property testing - lazily instantiated by
// NatGatewayPropertiesFormat_ARMGenerator()
var natGatewayPropertiesFormat_ARMGenerator gopter.Gen

// NatGatewayPropertiesFormat_ARMGenerator returns a generator of NatGatewayPropertiesFormat_ARM instances for property testing.
// We first initialize natGatewayPropertiesFormat_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NatGatewayPropertiesFormat_ARMGenerator() gopter.Gen {
	if natGatewayPropertiesFormat_ARMGenerator != nil {
		return natGatewayPropertiesFormat_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewayPropertiesFormat_ARM(generators)
	natGatewayPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(NatGatewayPropertiesFormat_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewayPropertiesFormat_ARM(generators)
	AddRelatedPropertyGeneratorsForNatGatewayPropertiesFormat_ARM(generators)
	natGatewayPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(NatGatewayPropertiesFormat_ARM{}), generators)

	return natGatewayPropertiesFormat_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNatGatewayPropertiesFormat_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGatewayPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForNatGatewayPropertiesFormat_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNatGatewayPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["PublicIpAddresses"] = gen.SliceOf(ApplicationGatewaySubResource_ARMGenerator())
	gens["PublicIpPrefixes"] = gen.SliceOf(ApplicationGatewaySubResource_ARMGenerator())
}

func Test_NatGatewaySku_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGatewaySku_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGatewaySku_ARM, NatGatewaySku_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGatewaySku_ARM runs a test to see if a specific instance of NatGatewaySku_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGatewaySku_ARM(subject NatGatewaySku_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGatewaySku_ARM
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

// Generator of NatGatewaySku_ARM instances for property testing - lazily instantiated by NatGatewaySku_ARMGenerator()
var natGatewaySku_ARMGenerator gopter.Gen

// NatGatewaySku_ARMGenerator returns a generator of NatGatewaySku_ARM instances for property testing.
func NatGatewaySku_ARMGenerator() gopter.Gen {
	if natGatewaySku_ARMGenerator != nil {
		return natGatewaySku_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewaySku_ARM(generators)
	natGatewaySku_ARMGenerator = gen.Struct(reflect.TypeOf(NatGatewaySku_ARM{}), generators)

	return natGatewaySku_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNatGatewaySku_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGatewaySku_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(NatGatewaySku_Name_Standard))
}
