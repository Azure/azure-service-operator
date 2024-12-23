// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_Permission_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Permission via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPermission, PermissionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPermission runs a test to see if a specific instance of Permission round trips to JSON and back losslessly
func RunJSONSerializationTestForPermission(subject Permission) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Permission
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

// Generator of Permission instances for property testing - lazily instantiated by PermissionGenerator()
var permissionGenerator gopter.Gen

// PermissionGenerator returns a generator of Permission instances for property testing.
func PermissionGenerator() gopter.Gen {
	if permissionGenerator != nil {
		return permissionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPermission(generators)
	permissionGenerator = gen.Struct(reflect.TypeOf(Permission{}), generators)

	return permissionGenerator
}

// AddIndependentPropertyGeneratorsForPermission is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPermission(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.SliceOf(gen.AlphaString())
	gens["DataActions"] = gen.SliceOf(gen.AlphaString())
	gens["NotActions"] = gen.SliceOf(gen.AlphaString())
	gens["NotDataActions"] = gen.SliceOf(gen.AlphaString())
}

func Test_RoleDefinitionProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleDefinitionProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleDefinitionProperties, RoleDefinitionPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleDefinitionProperties runs a test to see if a specific instance of RoleDefinitionProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleDefinitionProperties(subject RoleDefinitionProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleDefinitionProperties
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

// Generator of RoleDefinitionProperties instances for property testing - lazily instantiated by
// RoleDefinitionPropertiesGenerator()
var roleDefinitionPropertiesGenerator gopter.Gen

// RoleDefinitionPropertiesGenerator returns a generator of RoleDefinitionProperties instances for property testing.
// We first initialize roleDefinitionPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RoleDefinitionPropertiesGenerator() gopter.Gen {
	if roleDefinitionPropertiesGenerator != nil {
		return roleDefinitionPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleDefinitionProperties(generators)
	roleDefinitionPropertiesGenerator = gen.Struct(reflect.TypeOf(RoleDefinitionProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleDefinitionProperties(generators)
	AddRelatedPropertyGeneratorsForRoleDefinitionProperties(generators)
	roleDefinitionPropertiesGenerator = gen.Struct(reflect.TypeOf(RoleDefinitionProperties{}), generators)

	return roleDefinitionPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForRoleDefinitionProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleDefinitionProperties(gens map[string]gopter.Gen) {
	gens["AssignableScopes"] = gen.SliceOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["RoleName"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoleDefinitionProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleDefinitionProperties(gens map[string]gopter.Gen) {
	gens["Permissions"] = gen.SliceOf(PermissionGenerator())
}

func Test_RoleDefinition_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleDefinition_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleDefinition_Spec, RoleDefinition_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleDefinition_Spec runs a test to see if a specific instance of RoleDefinition_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleDefinition_Spec(subject RoleDefinition_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleDefinition_Spec
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

// Generator of RoleDefinition_Spec instances for property testing - lazily instantiated by
// RoleDefinition_SpecGenerator()
var roleDefinition_SpecGenerator gopter.Gen

// RoleDefinition_SpecGenerator returns a generator of RoleDefinition_Spec instances for property testing.
// We first initialize roleDefinition_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RoleDefinition_SpecGenerator() gopter.Gen {
	if roleDefinition_SpecGenerator != nil {
		return roleDefinition_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleDefinition_Spec(generators)
	roleDefinition_SpecGenerator = gen.Struct(reflect.TypeOf(RoleDefinition_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleDefinition_Spec(generators)
	AddRelatedPropertyGeneratorsForRoleDefinition_Spec(generators)
	roleDefinition_SpecGenerator = gen.Struct(reflect.TypeOf(RoleDefinition_Spec{}), generators)

	return roleDefinition_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRoleDefinition_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleDefinition_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRoleDefinition_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleDefinition_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RoleDefinitionPropertiesGenerator())
}
