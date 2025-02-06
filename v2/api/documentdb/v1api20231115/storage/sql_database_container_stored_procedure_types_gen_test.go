// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/storage"
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

func Test_SqlDatabaseContainerStoredProcedure_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerStoredProcedure to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseContainerStoredProcedure, SqlDatabaseContainerStoredProcedureGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseContainerStoredProcedure tests if a specific instance of SqlDatabaseContainerStoredProcedure round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseContainerStoredProcedure(subject SqlDatabaseContainerStoredProcedure) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.SqlDatabaseContainerStoredProcedure
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseContainerStoredProcedure
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

func Test_SqlDatabaseContainerStoredProcedure_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerStoredProcedure to SqlDatabaseContainerStoredProcedure via AssignProperties_To_SqlDatabaseContainerStoredProcedure & AssignProperties_From_SqlDatabaseContainerStoredProcedure returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure, SqlDatabaseContainerStoredProcedureGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure tests if a specific instance of SqlDatabaseContainerStoredProcedure can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure(subject SqlDatabaseContainerStoredProcedure) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlDatabaseContainerStoredProcedure
	err := copied.AssignProperties_To_SqlDatabaseContainerStoredProcedure(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerStoredProcedure
	err = actual.AssignProperties_From_SqlDatabaseContainerStoredProcedure(&other)
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

func Test_SqlDatabaseContainerStoredProcedure_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerStoredProcedure via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure, SqlDatabaseContainerStoredProcedureGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure runs a test to see if a specific instance of SqlDatabaseContainerStoredProcedure round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure(subject SqlDatabaseContainerStoredProcedure) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerStoredProcedure
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

// Generator of SqlDatabaseContainerStoredProcedure instances for property testing - lazily instantiated by
// SqlDatabaseContainerStoredProcedureGenerator()
var sqlDatabaseContainerStoredProcedureGenerator gopter.Gen

// SqlDatabaseContainerStoredProcedureGenerator returns a generator of SqlDatabaseContainerStoredProcedure instances for property testing.
func SqlDatabaseContainerStoredProcedureGenerator() gopter.Gen {
	if sqlDatabaseContainerStoredProcedureGenerator != nil {
		return sqlDatabaseContainerStoredProcedureGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure(generators)
	sqlDatabaseContainerStoredProcedureGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerStoredProcedure{}), generators)

	return sqlDatabaseContainerStoredProcedureGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure(gens map[string]gopter.Gen) {
	gens["Spec"] = SqlDatabaseContainerStoredProcedure_SpecGenerator()
	gens["Status"] = SqlDatabaseContainerStoredProcedure_STATUSGenerator()
}

func Test_SqlDatabaseContainerStoredProcedureOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerStoredProcedureOperatorSpec to SqlDatabaseContainerStoredProcedureOperatorSpec via AssignProperties_To_SqlDatabaseContainerStoredProcedureOperatorSpec & AssignProperties_From_SqlDatabaseContainerStoredProcedureOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedureOperatorSpec, SqlDatabaseContainerStoredProcedureOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedureOperatorSpec tests if a specific instance of SqlDatabaseContainerStoredProcedureOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedureOperatorSpec(subject SqlDatabaseContainerStoredProcedureOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlDatabaseContainerStoredProcedureOperatorSpec
	err := copied.AssignProperties_To_SqlDatabaseContainerStoredProcedureOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerStoredProcedureOperatorSpec
	err = actual.AssignProperties_From_SqlDatabaseContainerStoredProcedureOperatorSpec(&other)
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

func Test_SqlDatabaseContainerStoredProcedureOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerStoredProcedureOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerStoredProcedureOperatorSpec, SqlDatabaseContainerStoredProcedureOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerStoredProcedureOperatorSpec runs a test to see if a specific instance of SqlDatabaseContainerStoredProcedureOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerStoredProcedureOperatorSpec(subject SqlDatabaseContainerStoredProcedureOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerStoredProcedureOperatorSpec
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

// Generator of SqlDatabaseContainerStoredProcedureOperatorSpec instances for property testing - lazily instantiated by
// SqlDatabaseContainerStoredProcedureOperatorSpecGenerator()
var sqlDatabaseContainerStoredProcedureOperatorSpecGenerator gopter.Gen

// SqlDatabaseContainerStoredProcedureOperatorSpecGenerator returns a generator of SqlDatabaseContainerStoredProcedureOperatorSpec instances for property testing.
func SqlDatabaseContainerStoredProcedureOperatorSpecGenerator() gopter.Gen {
	if sqlDatabaseContainerStoredProcedureOperatorSpecGenerator != nil {
		return sqlDatabaseContainerStoredProcedureOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	sqlDatabaseContainerStoredProcedureOperatorSpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerStoredProcedureOperatorSpec{}), generators)

	return sqlDatabaseContainerStoredProcedureOperatorSpecGenerator
}

func Test_SqlDatabaseContainerStoredProcedure_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerStoredProcedure_STATUS to SqlDatabaseContainerStoredProcedure_STATUS via AssignProperties_To_SqlDatabaseContainerStoredProcedure_STATUS & AssignProperties_From_SqlDatabaseContainerStoredProcedure_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure_STATUS, SqlDatabaseContainerStoredProcedure_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure_STATUS tests if a specific instance of SqlDatabaseContainerStoredProcedure_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure_STATUS(subject SqlDatabaseContainerStoredProcedure_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlDatabaseContainerStoredProcedure_STATUS
	err := copied.AssignProperties_To_SqlDatabaseContainerStoredProcedure_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerStoredProcedure_STATUS
	err = actual.AssignProperties_From_SqlDatabaseContainerStoredProcedure_STATUS(&other)
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

func Test_SqlDatabaseContainerStoredProcedure_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerStoredProcedure_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure_STATUS, SqlDatabaseContainerStoredProcedure_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure_STATUS runs a test to see if a specific instance of SqlDatabaseContainerStoredProcedure_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure_STATUS(subject SqlDatabaseContainerStoredProcedure_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerStoredProcedure_STATUS
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

// Generator of SqlDatabaseContainerStoredProcedure_STATUS instances for property testing - lazily instantiated by
// SqlDatabaseContainerStoredProcedure_STATUSGenerator()
var sqlDatabaseContainerStoredProcedure_STATUSGenerator gopter.Gen

// SqlDatabaseContainerStoredProcedure_STATUSGenerator returns a generator of SqlDatabaseContainerStoredProcedure_STATUS instances for property testing.
// We first initialize sqlDatabaseContainerStoredProcedure_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseContainerStoredProcedure_STATUSGenerator() gopter.Gen {
	if sqlDatabaseContainerStoredProcedure_STATUSGenerator != nil {
		return sqlDatabaseContainerStoredProcedure_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_STATUS(generators)
	sqlDatabaseContainerStoredProcedure_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerStoredProcedure_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_STATUS(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_STATUS(generators)
	sqlDatabaseContainerStoredProcedure_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerStoredProcedure_STATUS{}), generators)

	return sqlDatabaseContainerStoredProcedure_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlStoredProcedureGetProperties_Resource_STATUSGenerator())
}

func Test_SqlDatabaseContainerStoredProcedure_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerStoredProcedure_Spec to SqlDatabaseContainerStoredProcedure_Spec via AssignProperties_To_SqlDatabaseContainerStoredProcedure_Spec & AssignProperties_From_SqlDatabaseContainerStoredProcedure_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure_Spec, SqlDatabaseContainerStoredProcedure_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure_Spec tests if a specific instance of SqlDatabaseContainerStoredProcedure_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure_Spec(subject SqlDatabaseContainerStoredProcedure_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlDatabaseContainerStoredProcedure_Spec
	err := copied.AssignProperties_To_SqlDatabaseContainerStoredProcedure_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerStoredProcedure_Spec
	err = actual.AssignProperties_From_SqlDatabaseContainerStoredProcedure_Spec(&other)
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

func Test_SqlDatabaseContainerStoredProcedure_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerStoredProcedure_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure_Spec, SqlDatabaseContainerStoredProcedure_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure_Spec runs a test to see if a specific instance of SqlDatabaseContainerStoredProcedure_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure_Spec(subject SqlDatabaseContainerStoredProcedure_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerStoredProcedure_Spec
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

// Generator of SqlDatabaseContainerStoredProcedure_Spec instances for property testing - lazily instantiated by
// SqlDatabaseContainerStoredProcedure_SpecGenerator()
var sqlDatabaseContainerStoredProcedure_SpecGenerator gopter.Gen

// SqlDatabaseContainerStoredProcedure_SpecGenerator returns a generator of SqlDatabaseContainerStoredProcedure_Spec instances for property testing.
// We first initialize sqlDatabaseContainerStoredProcedure_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseContainerStoredProcedure_SpecGenerator() gopter.Gen {
	if sqlDatabaseContainerStoredProcedure_SpecGenerator != nil {
		return sqlDatabaseContainerStoredProcedure_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_Spec(generators)
	sqlDatabaseContainerStoredProcedure_SpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerStoredProcedure_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_Spec(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_Spec(generators)
	sqlDatabaseContainerStoredProcedure_SpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerStoredProcedure_Spec{}), generators)

	return sqlDatabaseContainerStoredProcedure_SpecGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(SqlDatabaseContainerStoredProcedureOperatorSpecGenerator())
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlStoredProcedureResourceGenerator())
}

func Test_SqlStoredProcedureGetProperties_Resource_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlStoredProcedureGetProperties_Resource_STATUS to SqlStoredProcedureGetProperties_Resource_STATUS via AssignProperties_To_SqlStoredProcedureGetProperties_Resource_STATUS & AssignProperties_From_SqlStoredProcedureGetProperties_Resource_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlStoredProcedureGetProperties_Resource_STATUS, SqlStoredProcedureGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlStoredProcedureGetProperties_Resource_STATUS tests if a specific instance of SqlStoredProcedureGetProperties_Resource_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlStoredProcedureGetProperties_Resource_STATUS(subject SqlStoredProcedureGetProperties_Resource_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlStoredProcedureGetProperties_Resource_STATUS
	err := copied.AssignProperties_To_SqlStoredProcedureGetProperties_Resource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlStoredProcedureGetProperties_Resource_STATUS
	err = actual.AssignProperties_From_SqlStoredProcedureGetProperties_Resource_STATUS(&other)
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

func Test_SqlStoredProcedureGetProperties_Resource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureGetProperties_Resource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureGetProperties_Resource_STATUS, SqlStoredProcedureGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureGetProperties_Resource_STATUS runs a test to see if a specific instance of SqlStoredProcedureGetProperties_Resource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureGetProperties_Resource_STATUS(subject SqlStoredProcedureGetProperties_Resource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureGetProperties_Resource_STATUS
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

// Generator of SqlStoredProcedureGetProperties_Resource_STATUS instances for property testing - lazily instantiated by
// SqlStoredProcedureGetProperties_Resource_STATUSGenerator()
var sqlStoredProcedureGetProperties_Resource_STATUSGenerator gopter.Gen

// SqlStoredProcedureGetProperties_Resource_STATUSGenerator returns a generator of SqlStoredProcedureGetProperties_Resource_STATUS instances for property testing.
func SqlStoredProcedureGetProperties_Resource_STATUSGenerator() gopter.Gen {
	if sqlStoredProcedureGetProperties_Resource_STATUSGenerator != nil {
		return sqlStoredProcedureGetProperties_Resource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureGetProperties_Resource_STATUS(generators)
	sqlStoredProcedureGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureGetProperties_Resource_STATUS{}), generators)

	return sqlStoredProcedureGetProperties_Resource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

func Test_SqlStoredProcedureResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlStoredProcedureResource to SqlStoredProcedureResource via AssignProperties_To_SqlStoredProcedureResource & AssignProperties_From_SqlStoredProcedureResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlStoredProcedureResource, SqlStoredProcedureResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlStoredProcedureResource tests if a specific instance of SqlStoredProcedureResource can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlStoredProcedureResource(subject SqlStoredProcedureResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlStoredProcedureResource
	err := copied.AssignProperties_To_SqlStoredProcedureResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlStoredProcedureResource
	err = actual.AssignProperties_From_SqlStoredProcedureResource(&other)
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

func Test_SqlStoredProcedureResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureResource, SqlStoredProcedureResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureResource runs a test to see if a specific instance of SqlStoredProcedureResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureResource(subject SqlStoredProcedureResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureResource
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

// Generator of SqlStoredProcedureResource instances for property testing - lazily instantiated by
// SqlStoredProcedureResourceGenerator()
var sqlStoredProcedureResourceGenerator gopter.Gen

// SqlStoredProcedureResourceGenerator returns a generator of SqlStoredProcedureResource instances for property testing.
func SqlStoredProcedureResourceGenerator() gopter.Gen {
	if sqlStoredProcedureResourceGenerator != nil {
		return sqlStoredProcedureResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureResource(generators)
	sqlStoredProcedureResourceGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureResource{}), generators)

	return sqlStoredProcedureResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
