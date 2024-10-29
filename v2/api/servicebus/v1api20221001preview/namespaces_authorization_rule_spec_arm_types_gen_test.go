// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20221001preview

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

func Test_NamespacesAuthorizationRule_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRule_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRule_Spec_ARM, NamespacesAuthorizationRule_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRule_Spec_ARM runs a test to see if a specific instance of NamespacesAuthorizationRule_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRule_Spec_ARM(subject NamespacesAuthorizationRule_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRule_Spec_ARM
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

// Generator of NamespacesAuthorizationRule_Spec_ARM instances for property testing - lazily instantiated by
// NamespacesAuthorizationRule_Spec_ARMGenerator()
var namespacesAuthorizationRule_Spec_ARMGenerator gopter.Gen

// NamespacesAuthorizationRule_Spec_ARMGenerator returns a generator of NamespacesAuthorizationRule_Spec_ARM instances for property testing.
// We first initialize namespacesAuthorizationRule_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesAuthorizationRule_Spec_ARMGenerator() gopter.Gen {
	if namespacesAuthorizationRule_Spec_ARMGenerator != nil {
		return namespacesAuthorizationRule_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec_ARM(generators)
	namespacesAuthorizationRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_Spec_ARM(generators)
	namespacesAuthorizationRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_Spec_ARM{}), generators)

	return namespacesAuthorizationRule_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Namespaces_AuthorizationRule_Properties_Spec_ARMGenerator())
}

func Test_Namespaces_AuthorizationRule_Properties_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_AuthorizationRule_Properties_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_AuthorizationRule_Properties_Spec_ARM, Namespaces_AuthorizationRule_Properties_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_AuthorizationRule_Properties_Spec_ARM runs a test to see if a specific instance of Namespaces_AuthorizationRule_Properties_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_AuthorizationRule_Properties_Spec_ARM(subject Namespaces_AuthorizationRule_Properties_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_AuthorizationRule_Properties_Spec_ARM
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

// Generator of Namespaces_AuthorizationRule_Properties_Spec_ARM instances for property testing - lazily instantiated by
// Namespaces_AuthorizationRule_Properties_Spec_ARMGenerator()
var namespaces_AuthorizationRule_Properties_Spec_ARMGenerator gopter.Gen

// Namespaces_AuthorizationRule_Properties_Spec_ARMGenerator returns a generator of Namespaces_AuthorizationRule_Properties_Spec_ARM instances for property testing.
func Namespaces_AuthorizationRule_Properties_Spec_ARMGenerator() gopter.Gen {
	if namespaces_AuthorizationRule_Properties_Spec_ARMGenerator != nil {
		return namespaces_AuthorizationRule_Properties_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Properties_Spec_ARM(generators)
	namespaces_AuthorizationRule_Properties_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_AuthorizationRule_Properties_Spec_ARM{}), generators)

	return namespaces_AuthorizationRule_Properties_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Properties_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Properties_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM_Listen, Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM_Manage, Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM_Send))
}
