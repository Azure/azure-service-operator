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

func Test_RoleAssignmentProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignmentProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignmentProperties_STATUS, RoleAssignmentProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignmentProperties_STATUS runs a test to see if a specific instance of RoleAssignmentProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignmentProperties_STATUS(subject RoleAssignmentProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignmentProperties_STATUS
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

// Generator of RoleAssignmentProperties_STATUS instances for property testing - lazily instantiated by
// RoleAssignmentProperties_STATUSGenerator()
var roleAssignmentProperties_STATUSGenerator gopter.Gen

// RoleAssignmentProperties_STATUSGenerator returns a generator of RoleAssignmentProperties_STATUS instances for property testing.
func RoleAssignmentProperties_STATUSGenerator() gopter.Gen {
	if roleAssignmentProperties_STATUSGenerator != nil {
		return roleAssignmentProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentProperties_STATUS(generators)
	roleAssignmentProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(RoleAssignmentProperties_STATUS{}), generators)

	return roleAssignmentProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignmentProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignmentProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedOn"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.OneConstOf(
		RoleAssignmentProperties_PrincipalType_STATUS_ForeignGroup,
		RoleAssignmentProperties_PrincipalType_STATUS_Group,
		RoleAssignmentProperties_PrincipalType_STATUS_ServicePrincipal,
		RoleAssignmentProperties_PrincipalType_STATUS_User))
	gens["RoleDefinitionId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedOn"] = gen.PtrOf(gen.AlphaString())
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
// We first initialize roleAssignment_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RoleAssignment_STATUSGenerator() gopter.Gen {
	if roleAssignment_STATUSGenerator != nil {
		return roleAssignment_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_STATUS(generators)
	roleAssignment_STATUSGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_STATUS(generators)
	AddRelatedPropertyGeneratorsForRoleAssignment_STATUS(generators)
	roleAssignment_STATUSGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_STATUS{}), generators)

	return roleAssignment_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignment_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignment_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoleAssignment_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleAssignment_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RoleAssignmentProperties_STATUSGenerator())
}
