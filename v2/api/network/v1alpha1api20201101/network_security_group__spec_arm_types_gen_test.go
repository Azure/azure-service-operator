// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

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

func Test_NetworkSecurityGroup_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroup_SpecARM, NetworkSecurityGroup_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroup_SpecARM runs a test to see if a specific instance of NetworkSecurityGroup_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroup_SpecARM(subject NetworkSecurityGroup_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_SpecARM
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

// Generator of NetworkSecurityGroup_SpecARM instances for property testing - lazily instantiated by
//NetworkSecurityGroup_SpecARMGenerator()
var networkSecurityGroup_specARMGenerator gopter.Gen

// NetworkSecurityGroup_SpecARMGenerator returns a generator of NetworkSecurityGroup_SpecARM instances for property testing.
// We first initialize networkSecurityGroup_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroup_SpecARMGenerator() gopter.Gen {
	if networkSecurityGroup_specARMGenerator != nil {
		return networkSecurityGroup_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_SpecARM(generators)
	networkSecurityGroup_specARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_SpecARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroup_SpecARM(generators)
	networkSecurityGroup_specARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_SpecARM{}), generators)

	return networkSecurityGroup_specARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroup_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroup_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroup_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroup_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NetworkSecurityGroupPropertiesFormatARMGenerator())
}

func Test_NetworkSecurityGroupPropertiesFormatARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupPropertiesFormatARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatARM, NetworkSecurityGroupPropertiesFormatARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatARM runs a test to see if a specific instance of NetworkSecurityGroupPropertiesFormatARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatARM(subject NetworkSecurityGroupPropertiesFormatARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupPropertiesFormatARM
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

// Generator of NetworkSecurityGroupPropertiesFormatARM instances for property testing - lazily instantiated by
//NetworkSecurityGroupPropertiesFormatARMGenerator()
var networkSecurityGroupPropertiesFormatARMGenerator gopter.Gen

// NetworkSecurityGroupPropertiesFormatARMGenerator returns a generator of NetworkSecurityGroupPropertiesFormatARM instances for property testing.
func NetworkSecurityGroupPropertiesFormatARMGenerator() gopter.Gen {
	if networkSecurityGroupPropertiesFormatARMGenerator != nil {
		return networkSecurityGroupPropertiesFormatARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	networkSecurityGroupPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormatARM{}), generators)

	return networkSecurityGroupPropertiesFormatARMGenerator
}
