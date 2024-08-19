// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220401

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401/storage"
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

func Test_Permission_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Permission to Permission via AssignProperties_To_Permission & AssignProperties_From_Permission returns original",
		prop.ForAll(RunPropertyAssignmentTestForPermission, PermissionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPermission tests if a specific instance of Permission can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPermission(subject Permission) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Permission
	err := copied.AssignProperties_To_Permission(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Permission
	err = actual.AssignProperties_From_Permission(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

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

func Test_Permission_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Permission_STATUS to Permission_STATUS via AssignProperties_To_Permission_STATUS & AssignProperties_From_Permission_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPermission_STATUS, Permission_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPermission_STATUS tests if a specific instance of Permission_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPermission_STATUS(subject Permission_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Permission_STATUS
	err := copied.AssignProperties_To_Permission_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Permission_STATUS
	err = actual.AssignProperties_From_Permission_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Permission_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Permission_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPermission_STATUS, Permission_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPermission_STATUS runs a test to see if a specific instance of Permission_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPermission_STATUS(subject Permission_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Permission_STATUS
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

// Generator of Permission_STATUS instances for property testing - lazily instantiated by Permission_STATUSGenerator()
var permission_STATUSGenerator gopter.Gen

// Permission_STATUSGenerator returns a generator of Permission_STATUS instances for property testing.
func Permission_STATUSGenerator() gopter.Gen {
	if permission_STATUSGenerator != nil {
		return permission_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPermission_STATUS(generators)
	permission_STATUSGenerator = gen.Struct(reflect.TypeOf(Permission_STATUS{}), generators)

	return permission_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPermission_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPermission_STATUS(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.SliceOf(gen.AlphaString())
	gens["DataActions"] = gen.SliceOf(gen.AlphaString())
	gens["NotActions"] = gen.SliceOf(gen.AlphaString())
	gens["NotDataActions"] = gen.SliceOf(gen.AlphaString())
}

func Test_RoleDefinition_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleDefinition to hub returns original",
		prop.ForAll(RunResourceConversionTestForRoleDefinition, RoleDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRoleDefinition tests if a specific instance of RoleDefinition round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRoleDefinition(subject RoleDefinition) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.RoleDefinition
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RoleDefinition
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RoleDefinition_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleDefinition to RoleDefinition via AssignProperties_To_RoleDefinition & AssignProperties_From_RoleDefinition returns original",
		prop.ForAll(RunPropertyAssignmentTestForRoleDefinition, RoleDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRoleDefinition tests if a specific instance of RoleDefinition can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRoleDefinition(subject RoleDefinition) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.RoleDefinition
	err := copied.AssignProperties_To_RoleDefinition(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RoleDefinition
	err = actual.AssignProperties_From_RoleDefinition(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RoleDefinition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleDefinition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleDefinition, RoleDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleDefinition runs a test to see if a specific instance of RoleDefinition round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleDefinition(subject RoleDefinition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleDefinition
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

// Generator of RoleDefinition instances for property testing - lazily instantiated by RoleDefinitionGenerator()
var roleDefinitionGenerator gopter.Gen

// RoleDefinitionGenerator returns a generator of RoleDefinition instances for property testing.
func RoleDefinitionGenerator() gopter.Gen {
	if roleDefinitionGenerator != nil {
		return roleDefinitionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRoleDefinition(generators)
	roleDefinitionGenerator = gen.Struct(reflect.TypeOf(RoleDefinition{}), generators)

	return roleDefinitionGenerator
}

// AddRelatedPropertyGeneratorsForRoleDefinition is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleDefinition(gens map[string]gopter.Gen) {
	gens["Spec"] = RoleDefinition_SpecGenerator()
	gens["Status"] = RoleDefinition_STATUSGenerator()
}

func Test_RoleDefinitionOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleDefinitionOperatorSpec to RoleDefinitionOperatorSpec via AssignProperties_To_RoleDefinitionOperatorSpec & AssignProperties_From_RoleDefinitionOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRoleDefinitionOperatorSpec, RoleDefinitionOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRoleDefinitionOperatorSpec tests if a specific instance of RoleDefinitionOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRoleDefinitionOperatorSpec(subject RoleDefinitionOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.RoleDefinitionOperatorSpec
	err := copied.AssignProperties_To_RoleDefinitionOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RoleDefinitionOperatorSpec
	err = actual.AssignProperties_From_RoleDefinitionOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RoleDefinitionOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleDefinitionOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleDefinitionOperatorSpec, RoleDefinitionOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleDefinitionOperatorSpec runs a test to see if a specific instance of RoleDefinitionOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleDefinitionOperatorSpec(subject RoleDefinitionOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleDefinitionOperatorSpec
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

// Generator of RoleDefinitionOperatorSpec instances for property testing - lazily instantiated by
// RoleDefinitionOperatorSpecGenerator()
var roleDefinitionOperatorSpecGenerator gopter.Gen

// RoleDefinitionOperatorSpecGenerator returns a generator of RoleDefinitionOperatorSpec instances for property testing.
func RoleDefinitionOperatorSpecGenerator() gopter.Gen {
	if roleDefinitionOperatorSpecGenerator != nil {
		return roleDefinitionOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleDefinitionOperatorSpec(generators)
	roleDefinitionOperatorSpecGenerator = gen.Struct(reflect.TypeOf(RoleDefinitionOperatorSpec{}), generators)

	return roleDefinitionOperatorSpecGenerator
}

// AddIndependentPropertyGeneratorsForRoleDefinitionOperatorSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleDefinitionOperatorSpec(gens map[string]gopter.Gen) {
	gens["UUIDGeneration"] = gen.PtrOf(gen.AlphaString())
}

func Test_RoleDefinition_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleDefinition_STATUS to RoleDefinition_STATUS via AssignProperties_To_RoleDefinition_STATUS & AssignProperties_From_RoleDefinition_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRoleDefinition_STATUS, RoleDefinition_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRoleDefinition_STATUS tests if a specific instance of RoleDefinition_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRoleDefinition_STATUS(subject RoleDefinition_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.RoleDefinition_STATUS
	err := copied.AssignProperties_To_RoleDefinition_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RoleDefinition_STATUS
	err = actual.AssignProperties_From_RoleDefinition_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RoleDefinition_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleDefinition_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleDefinition_STATUS, RoleDefinition_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleDefinition_STATUS runs a test to see if a specific instance of RoleDefinition_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleDefinition_STATUS(subject RoleDefinition_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleDefinition_STATUS
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

// Generator of RoleDefinition_STATUS instances for property testing - lazily instantiated by
// RoleDefinition_STATUSGenerator()
var roleDefinition_STATUSGenerator gopter.Gen

// RoleDefinition_STATUSGenerator returns a generator of RoleDefinition_STATUS instances for property testing.
// We first initialize roleDefinition_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RoleDefinition_STATUSGenerator() gopter.Gen {
	if roleDefinition_STATUSGenerator != nil {
		return roleDefinition_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleDefinition_STATUS(generators)
	roleDefinition_STATUSGenerator = gen.Struct(reflect.TypeOf(RoleDefinition_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleDefinition_STATUS(generators)
	AddRelatedPropertyGeneratorsForRoleDefinition_STATUS(generators)
	roleDefinition_STATUSGenerator = gen.Struct(reflect.TypeOf(RoleDefinition_STATUS{}), generators)

	return roleDefinition_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRoleDefinition_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleDefinition_STATUS(gens map[string]gopter.Gen) {
	gens["AssignableScopes"] = gen.SliceOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedOn"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesType"] = gen.PtrOf(gen.AlphaString())
	gens["RoleName"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedOn"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoleDefinition_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleDefinition_STATUS(gens map[string]gopter.Gen) {
	gens["Permissions"] = gen.SliceOf(Permission_STATUSGenerator())
}

func Test_RoleDefinition_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleDefinition_Spec to RoleDefinition_Spec via AssignProperties_To_RoleDefinition_Spec & AssignProperties_From_RoleDefinition_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRoleDefinition_Spec, RoleDefinition_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRoleDefinition_Spec tests if a specific instance of RoleDefinition_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRoleDefinition_Spec(subject RoleDefinition_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.RoleDefinition_Spec
	err := copied.AssignProperties_To_RoleDefinition_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RoleDefinition_Spec
	err = actual.AssignProperties_From_RoleDefinition_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
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
	gens["AzureName"] = gen.AlphaString()
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["RoleName"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoleDefinition_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleDefinition_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(RoleDefinitionOperatorSpecGenerator())
	gens["Permissions"] = gen.SliceOf(PermissionGenerator())
}
