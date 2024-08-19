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

func Test_RoleAssignment_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignment, RoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignment runs a test to see if a specific instance of RoleAssignment round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignment(subject RoleAssignment) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment
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

// Generator of RoleAssignment instances for property testing - lazily instantiated by RoleAssignmentGenerator()
var roleAssignmentGenerator gopter.Gen

// RoleAssignmentGenerator returns a generator of RoleAssignment instances for property testing.
func RoleAssignmentGenerator() gopter.Gen {
	if roleAssignmentGenerator != nil {
		return roleAssignmentGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRoleAssignment(generators)
	roleAssignmentGenerator = gen.Struct(reflect.TypeOf(RoleAssignment{}), generators)

	return roleAssignmentGenerator
}

// AddRelatedPropertyGeneratorsForRoleAssignment is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleAssignment(gens map[string]gopter.Gen) {
	gens["Spec"] = RoleAssignment_SpecGenerator()
	gens["Status"] = RoleAssignment_STATUSGenerator()
}

func Test_RoleAssignmentOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignmentOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignmentOperatorSpec, RoleAssignmentOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignmentOperatorSpec runs a test to see if a specific instance of RoleAssignmentOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignmentOperatorSpec(subject RoleAssignmentOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignmentOperatorSpec
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

// Generator of RoleAssignmentOperatorSpec instances for property testing - lazily instantiated by
// RoleAssignmentOperatorSpecGenerator()
var roleAssignmentOperatorSpecGenerator gopter.Gen

// RoleAssignmentOperatorSpecGenerator returns a generator of RoleAssignmentOperatorSpec instances for property testing.
func RoleAssignmentOperatorSpecGenerator() gopter.Gen {
	if roleAssignmentOperatorSpecGenerator != nil {
		return roleAssignmentOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentOperatorSpec(generators)
	roleAssignmentOperatorSpecGenerator = gen.Struct(reflect.TypeOf(RoleAssignmentOperatorSpec{}), generators)

	return roleAssignmentOperatorSpecGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignmentOperatorSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignmentOperatorSpec(gens map[string]gopter.Gen) {
	gens["UUIDGeneration"] = gen.PtrOf(gen.AlphaString())
}

func Test_RoleAssignment_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignment_STATUS, RoleAssignment_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignment_STATUS runs a test to see if a specific instance of RoleAssignment_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignment_STATUS(subject RoleAssignment_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment_STATUS
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

// Generator of RoleAssignment_STATUS instances for property testing - lazily instantiated by
// RoleAssignment_STATUSGenerator()
var roleAssignment_STATUSGenerator gopter.Gen

// RoleAssignment_STATUSGenerator returns a generator of RoleAssignment_STATUS instances for property testing.
func RoleAssignment_STATUSGenerator() gopter.Gen {
	if roleAssignment_STATUSGenerator != nil {
		return roleAssignment_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_STATUS(generators)
	roleAssignment_STATUSGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_STATUS{}), generators)

	return roleAssignment_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignment_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignment_STATUS(gens map[string]gopter.Gen) {
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedOn"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.AlphaString())
	gens["RoleDefinitionId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedOn"] = gen.PtrOf(gen.AlphaString())
}

func Test_RoleAssignment_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignment_Spec, RoleAssignment_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignment_Spec runs a test to see if a specific instance of RoleAssignment_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignment_Spec(subject RoleAssignment_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment_Spec
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

// Generator of RoleAssignment_Spec instances for property testing - lazily instantiated by
// RoleAssignment_SpecGenerator()
var roleAssignment_SpecGenerator gopter.Gen

// RoleAssignment_SpecGenerator returns a generator of RoleAssignment_Spec instances for property testing.
// We first initialize roleAssignment_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RoleAssignment_SpecGenerator() gopter.Gen {
	if roleAssignment_SpecGenerator != nil {
		return roleAssignment_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_Spec(generators)
	roleAssignment_SpecGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_Spec(generators)
	AddRelatedPropertyGeneratorsForRoleAssignment_Spec(generators)
	roleAssignment_SpecGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_Spec{}), generators)

	return roleAssignment_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignment_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignment_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoleAssignment_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleAssignment_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(RoleAssignmentOperatorSpecGenerator())
}
