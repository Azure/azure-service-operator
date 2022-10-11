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

func Test_AuthorizationRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationRule_STATUS_ARM, AuthorizationRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationRule_STATUS_ARM runs a test to see if a specific instance of AuthorizationRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationRule_STATUS_ARM(subject AuthorizationRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationRule_STATUS_ARM
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

// Generator of AuthorizationRule_STATUS_ARM instances for property testing - lazily instantiated by
// AuthorizationRule_STATUS_ARMGenerator()
var authorizationRule_STATUS_ARMGenerator gopter.Gen

// AuthorizationRule_STATUS_ARMGenerator returns a generator of AuthorizationRule_STATUS_ARM instances for property testing.
// We first initialize authorizationRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationRule_STATUS_ARMGenerator() gopter.Gen {
	if authorizationRule_STATUS_ARMGenerator != nil {
		return authorizationRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRule_STATUS_ARM(generators)
	authorizationRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForAuthorizationRule_STATUS_ARM(generators)
	authorizationRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationRule_STATUS_ARM{}), generators)

	return authorizationRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AuthorizationRule_Properties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_AuthorizationRule_Properties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationRule_Properties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationRule_Properties_STATUS_ARM, AuthorizationRule_Properties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationRule_Properties_STATUS_ARM runs a test to see if a specific instance of AuthorizationRule_Properties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationRule_Properties_STATUS_ARM(subject AuthorizationRule_Properties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationRule_Properties_STATUS_ARM
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

// Generator of AuthorizationRule_Properties_STATUS_ARM instances for property testing - lazily instantiated by
// AuthorizationRule_Properties_STATUS_ARMGenerator()
var authorizationRule_Properties_STATUS_ARMGenerator gopter.Gen

// AuthorizationRule_Properties_STATUS_ARMGenerator returns a generator of AuthorizationRule_Properties_STATUS_ARM instances for property testing.
func AuthorizationRule_Properties_STATUS_ARMGenerator() gopter.Gen {
	if authorizationRule_Properties_STATUS_ARMGenerator != nil {
		return authorizationRule_Properties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRule_Properties_STATUS_ARM(generators)
	authorizationRule_Properties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationRule_Properties_STATUS_ARM{}), generators)

	return authorizationRule_Properties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationRule_Properties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationRule_Properties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(AuthorizationRule_Properties_Rights_STATUS_Listen, AuthorizationRule_Properties_Rights_STATUS_Manage, AuthorizationRule_Properties_Rights_STATUS_Send))
}

func Test_SystemData_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS_ARM, SystemData_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS_ARM runs a test to see if a specific instance of SystemData_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS_ARM(subject SystemData_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS_ARM
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

// Generator of SystemData_STATUS_ARM instances for property testing - lazily instantiated by
// SystemData_STATUS_ARMGenerator()
var systemData_STATUS_ARMGenerator gopter.Gen

// SystemData_STATUS_ARMGenerator returns a generator of SystemData_STATUS_ARM instances for property testing.
func SystemData_STATUS_ARMGenerator() gopter.Gen {
	if systemData_STATUS_ARMGenerator != nil {
		return systemData_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(generators)
	systemData_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS_ARM{}), generators)

	return systemData_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_Application,
		SystemData_CreatedByType_STATUS_Key,
		SystemData_CreatedByType_STATUS_ManagedIdentity,
		SystemData_CreatedByType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_Application,
		SystemData_LastModifiedByType_STATUS_Key,
		SystemData_LastModifiedByType_STATUS_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_User))
}
