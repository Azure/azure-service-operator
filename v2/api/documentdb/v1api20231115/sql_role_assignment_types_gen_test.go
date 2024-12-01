// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231115

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20231115/storage"
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

func Test_SqlRoleAssignment_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlRoleAssignment to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlRoleAssignment, SqlRoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlRoleAssignment tests if a specific instance of SqlRoleAssignment round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlRoleAssignment(subject SqlRoleAssignment) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.SqlRoleAssignment
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlRoleAssignment
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

func Test_SqlRoleAssignment_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlRoleAssignment to SqlRoleAssignment via AssignProperties_To_SqlRoleAssignment & AssignProperties_From_SqlRoleAssignment returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlRoleAssignment, SqlRoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlRoleAssignment tests if a specific instance of SqlRoleAssignment can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlRoleAssignment(subject SqlRoleAssignment) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlRoleAssignment
	err := copied.AssignProperties_To_SqlRoleAssignment(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlRoleAssignment
	err = actual.AssignProperties_From_SqlRoleAssignment(&other)
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

func Test_SqlRoleAssignment_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlRoleAssignment via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlRoleAssignment, SqlRoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlRoleAssignment runs a test to see if a specific instance of SqlRoleAssignment round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlRoleAssignment(subject SqlRoleAssignment) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlRoleAssignment
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

// Generator of SqlRoleAssignment instances for property testing - lazily instantiated by SqlRoleAssignmentGenerator()
var sqlRoleAssignmentGenerator gopter.Gen

// SqlRoleAssignmentGenerator returns a generator of SqlRoleAssignment instances for property testing.
func SqlRoleAssignmentGenerator() gopter.Gen {
	if sqlRoleAssignmentGenerator != nil {
		return sqlRoleAssignmentGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlRoleAssignment(generators)
	sqlRoleAssignmentGenerator = gen.Struct(reflect.TypeOf(SqlRoleAssignment{}), generators)

	return sqlRoleAssignmentGenerator
}

// AddRelatedPropertyGeneratorsForSqlRoleAssignment is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlRoleAssignment(gens map[string]gopter.Gen) {
	gens["Spec"] = SqlRoleAssignment_SpecGenerator()
	gens["Status"] = SqlRoleAssignment_STATUSGenerator()
}

func Test_SqlRoleAssignmentOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlRoleAssignmentOperatorSpec to SqlRoleAssignmentOperatorSpec via AssignProperties_To_SqlRoleAssignmentOperatorSpec & AssignProperties_From_SqlRoleAssignmentOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlRoleAssignmentOperatorSpec, SqlRoleAssignmentOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlRoleAssignmentOperatorSpec tests if a specific instance of SqlRoleAssignmentOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlRoleAssignmentOperatorSpec(subject SqlRoleAssignmentOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlRoleAssignmentOperatorSpec
	err := copied.AssignProperties_To_SqlRoleAssignmentOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlRoleAssignmentOperatorSpec
	err = actual.AssignProperties_From_SqlRoleAssignmentOperatorSpec(&other)
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

func Test_SqlRoleAssignmentOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlRoleAssignmentOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlRoleAssignmentOperatorSpec, SqlRoleAssignmentOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlRoleAssignmentOperatorSpec runs a test to see if a specific instance of SqlRoleAssignmentOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlRoleAssignmentOperatorSpec(subject SqlRoleAssignmentOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlRoleAssignmentOperatorSpec
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

// Generator of SqlRoleAssignmentOperatorSpec instances for property testing - lazily instantiated by
// SqlRoleAssignmentOperatorSpecGenerator()
var sqlRoleAssignmentOperatorSpecGenerator gopter.Gen

// SqlRoleAssignmentOperatorSpecGenerator returns a generator of SqlRoleAssignmentOperatorSpec instances for property testing.
func SqlRoleAssignmentOperatorSpecGenerator() gopter.Gen {
	if sqlRoleAssignmentOperatorSpecGenerator != nil {
		return sqlRoleAssignmentOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	sqlRoleAssignmentOperatorSpecGenerator = gen.Struct(reflect.TypeOf(SqlRoleAssignmentOperatorSpec{}), generators)

	return sqlRoleAssignmentOperatorSpecGenerator
}

func Test_SqlRoleAssignment_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlRoleAssignment_STATUS to SqlRoleAssignment_STATUS via AssignProperties_To_SqlRoleAssignment_STATUS & AssignProperties_From_SqlRoleAssignment_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlRoleAssignment_STATUS, SqlRoleAssignment_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlRoleAssignment_STATUS tests if a specific instance of SqlRoleAssignment_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlRoleAssignment_STATUS(subject SqlRoleAssignment_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlRoleAssignment_STATUS
	err := copied.AssignProperties_To_SqlRoleAssignment_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlRoleAssignment_STATUS
	err = actual.AssignProperties_From_SqlRoleAssignment_STATUS(&other)
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

func Test_SqlRoleAssignment_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlRoleAssignment_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlRoleAssignment_STATUS, SqlRoleAssignment_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlRoleAssignment_STATUS runs a test to see if a specific instance of SqlRoleAssignment_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlRoleAssignment_STATUS(subject SqlRoleAssignment_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlRoleAssignment_STATUS
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

// Generator of SqlRoleAssignment_STATUS instances for property testing - lazily instantiated by
// SqlRoleAssignment_STATUSGenerator()
var sqlRoleAssignment_STATUSGenerator gopter.Gen

// SqlRoleAssignment_STATUSGenerator returns a generator of SqlRoleAssignment_STATUS instances for property testing.
func SqlRoleAssignment_STATUSGenerator() gopter.Gen {
	if sqlRoleAssignment_STATUSGenerator != nil {
		return sqlRoleAssignment_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlRoleAssignment_STATUS(generators)
	sqlRoleAssignment_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlRoleAssignment_STATUS{}), generators)

	return sqlRoleAssignment_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlRoleAssignment_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlRoleAssignment_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["RoleDefinitionId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_SqlRoleAssignment_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlRoleAssignment_Spec to SqlRoleAssignment_Spec via AssignProperties_To_SqlRoleAssignment_Spec & AssignProperties_From_SqlRoleAssignment_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlRoleAssignment_Spec, SqlRoleAssignment_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlRoleAssignment_Spec tests if a specific instance of SqlRoleAssignment_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlRoleAssignment_Spec(subject SqlRoleAssignment_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlRoleAssignment_Spec
	err := copied.AssignProperties_To_SqlRoleAssignment_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlRoleAssignment_Spec
	err = actual.AssignProperties_From_SqlRoleAssignment_Spec(&other)
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

func Test_SqlRoleAssignment_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlRoleAssignment_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlRoleAssignment_Spec, SqlRoleAssignment_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlRoleAssignment_Spec runs a test to see if a specific instance of SqlRoleAssignment_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlRoleAssignment_Spec(subject SqlRoleAssignment_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlRoleAssignment_Spec
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

// Generator of SqlRoleAssignment_Spec instances for property testing - lazily instantiated by
// SqlRoleAssignment_SpecGenerator()
var sqlRoleAssignment_SpecGenerator gopter.Gen

// SqlRoleAssignment_SpecGenerator returns a generator of SqlRoleAssignment_Spec instances for property testing.
// We first initialize sqlRoleAssignment_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlRoleAssignment_SpecGenerator() gopter.Gen {
	if sqlRoleAssignment_SpecGenerator != nil {
		return sqlRoleAssignment_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlRoleAssignment_Spec(generators)
	sqlRoleAssignment_SpecGenerator = gen.Struct(reflect.TypeOf(SqlRoleAssignment_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlRoleAssignment_Spec(generators)
	AddRelatedPropertyGeneratorsForSqlRoleAssignment_Spec(generators)
	sqlRoleAssignment_SpecGenerator = gen.Struct(reflect.TypeOf(SqlRoleAssignment_Spec{}), generators)

	return sqlRoleAssignment_SpecGenerator
}

// AddIndependentPropertyGeneratorsForSqlRoleAssignment_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlRoleAssignment_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["RoleDefinitionId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlRoleAssignment_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlRoleAssignment_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(SqlRoleAssignmentOperatorSpecGenerator())
}
