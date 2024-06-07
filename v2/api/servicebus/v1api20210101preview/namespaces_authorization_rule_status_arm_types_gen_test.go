// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210101preview

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

func Test_Namespaces_AuthorizationRule_Properties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_AuthorizationRule_Properties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_AuthorizationRule_Properties_STATUS_ARM, Namespaces_AuthorizationRule_Properties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_AuthorizationRule_Properties_STATUS_ARM runs a test to see if a specific instance of Namespaces_AuthorizationRule_Properties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_AuthorizationRule_Properties_STATUS_ARM(subject Namespaces_AuthorizationRule_Properties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_AuthorizationRule_Properties_STATUS_ARM
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

// Generator of Namespaces_AuthorizationRule_Properties_STATUS_ARM instances for property testing - lazily instantiated
// by Namespaces_AuthorizationRule_Properties_STATUS_ARMGenerator()
var namespaces_AuthorizationRule_Properties_STATUS_ARMGenerator gopter.Gen

// Namespaces_AuthorizationRule_Properties_STATUS_ARMGenerator returns a generator of Namespaces_AuthorizationRule_Properties_STATUS_ARM instances for property testing.
func Namespaces_AuthorizationRule_Properties_STATUS_ARMGenerator() gopter.Gen {
	if namespaces_AuthorizationRule_Properties_STATUS_ARMGenerator != nil {
		return namespaces_AuthorizationRule_Properties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Properties_STATUS_ARM(generators)
	namespaces_AuthorizationRule_Properties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_AuthorizationRule_Properties_STATUS_ARM{}), generators)

	return namespaces_AuthorizationRule_Properties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Properties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Properties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(Namespaces_AuthorizationRule_Properties_Rights_STATUS_Listen, Namespaces_AuthorizationRule_Properties_Rights_STATUS_Manage, Namespaces_AuthorizationRule_Properties_Rights_STATUS_Send))
}

func Test_Namespaces_AuthorizationRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_AuthorizationRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_AuthorizationRule_STATUS_ARM, Namespaces_AuthorizationRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_AuthorizationRule_STATUS_ARM runs a test to see if a specific instance of Namespaces_AuthorizationRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_AuthorizationRule_STATUS_ARM(subject Namespaces_AuthorizationRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_AuthorizationRule_STATUS_ARM
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

// Generator of Namespaces_AuthorizationRule_STATUS_ARM instances for property testing - lazily instantiated by
// Namespaces_AuthorizationRule_STATUS_ARMGenerator()
var namespaces_AuthorizationRule_STATUS_ARMGenerator gopter.Gen

// Namespaces_AuthorizationRule_STATUS_ARMGenerator returns a generator of Namespaces_AuthorizationRule_STATUS_ARM instances for property testing.
// We first initialize namespaces_AuthorizationRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_AuthorizationRule_STATUS_ARMGenerator() gopter.Gen {
	if namespaces_AuthorizationRule_STATUS_ARMGenerator != nil {
		return namespaces_AuthorizationRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_STATUS_ARM(generators)
	namespaces_AuthorizationRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_AuthorizationRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_AuthorizationRule_STATUS_ARM(generators)
	namespaces_AuthorizationRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_AuthorizationRule_STATUS_ARM{}), generators)

	return namespaces_AuthorizationRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespaces_AuthorizationRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_AuthorizationRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Namespaces_AuthorizationRule_Properties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}
