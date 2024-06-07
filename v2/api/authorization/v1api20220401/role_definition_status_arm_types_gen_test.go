// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220401

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

func Test_Permission_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Permission_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPermission_STATUS_ARM, Permission_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPermission_STATUS_ARM runs a test to see if a specific instance of Permission_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPermission_STATUS_ARM(subject Permission_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Permission_STATUS_ARM
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

// Generator of Permission_STATUS_ARM instances for property testing - lazily instantiated by
// Permission_STATUS_ARMGenerator()
var permission_STATUS_ARMGenerator gopter.Gen

// Permission_STATUS_ARMGenerator returns a generator of Permission_STATUS_ARM instances for property testing.
func Permission_STATUS_ARMGenerator() gopter.Gen {
	if permission_STATUS_ARMGenerator != nil {
		return permission_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPermission_STATUS_ARM(generators)
	permission_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Permission_STATUS_ARM{}), generators)

	return permission_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPermission_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPermission_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.SliceOf(gen.AlphaString())
	gens["DataActions"] = gen.SliceOf(gen.AlphaString())
	gens["NotActions"] = gen.SliceOf(gen.AlphaString())
	gens["NotDataActions"] = gen.SliceOf(gen.AlphaString())
}

func Test_RoleDefinitionProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleDefinitionProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleDefinitionProperties_STATUS_ARM, RoleDefinitionProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleDefinitionProperties_STATUS_ARM runs a test to see if a specific instance of RoleDefinitionProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleDefinitionProperties_STATUS_ARM(subject RoleDefinitionProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleDefinitionProperties_STATUS_ARM
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

// Generator of RoleDefinitionProperties_STATUS_ARM instances for property testing - lazily instantiated by
// RoleDefinitionProperties_STATUS_ARMGenerator()
var roleDefinitionProperties_STATUS_ARMGenerator gopter.Gen

// RoleDefinitionProperties_STATUS_ARMGenerator returns a generator of RoleDefinitionProperties_STATUS_ARM instances for property testing.
// We first initialize roleDefinitionProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RoleDefinitionProperties_STATUS_ARMGenerator() gopter.Gen {
	if roleDefinitionProperties_STATUS_ARMGenerator != nil {
		return roleDefinitionProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleDefinitionProperties_STATUS_ARM(generators)
	roleDefinitionProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RoleDefinitionProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleDefinitionProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForRoleDefinitionProperties_STATUS_ARM(generators)
	roleDefinitionProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RoleDefinitionProperties_STATUS_ARM{}), generators)

	return roleDefinitionProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRoleDefinitionProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleDefinitionProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AssignableScopes"] = gen.SliceOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedOn"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["RoleName"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedOn"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoleDefinitionProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleDefinitionProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Permissions"] = gen.SliceOf(Permission_STATUS_ARMGenerator())
}

func Test_RoleDefinition_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleDefinition_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleDefinition_STATUS_ARM, RoleDefinition_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleDefinition_STATUS_ARM runs a test to see if a specific instance of RoleDefinition_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleDefinition_STATUS_ARM(subject RoleDefinition_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleDefinition_STATUS_ARM
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

// Generator of RoleDefinition_STATUS_ARM instances for property testing - lazily instantiated by
// RoleDefinition_STATUS_ARMGenerator()
var roleDefinition_STATUS_ARMGenerator gopter.Gen

// RoleDefinition_STATUS_ARMGenerator returns a generator of RoleDefinition_STATUS_ARM instances for property testing.
// We first initialize roleDefinition_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RoleDefinition_STATUS_ARMGenerator() gopter.Gen {
	if roleDefinition_STATUS_ARMGenerator != nil {
		return roleDefinition_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleDefinition_STATUS_ARM(generators)
	roleDefinition_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RoleDefinition_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleDefinition_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForRoleDefinition_STATUS_ARM(generators)
	roleDefinition_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RoleDefinition_STATUS_ARM{}), generators)

	return roleDefinition_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRoleDefinition_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleDefinition_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoleDefinition_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleDefinition_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RoleDefinitionProperties_STATUS_ARMGenerator())
}
