// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

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
	gens["Spec"] = DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator()
	gens["Status"] = SqlStoredProcedureGetResultsSTATUSGenerator()
}

func Test_DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec, DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(subject DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
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

// Generator of DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec instances for property testing - lazily
// instantiated by DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator()
var databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator != nil {
		return databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(generators)
	databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(generators)
	databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}), generators)

	return databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlStoredProcedureResourceGenerator())
}

func Test_SqlStoredProcedureGetResults_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureGetResults_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureGetResultsSTATUS, SqlStoredProcedureGetResultsSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureGetResultsSTATUS runs a test to see if a specific instance of SqlStoredProcedureGetResults_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureGetResultsSTATUS(subject SqlStoredProcedureGetResults_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureGetResults_STATUS
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

// Generator of SqlStoredProcedureGetResults_STATUS instances for property testing - lazily instantiated by
// SqlStoredProcedureGetResultsSTATUSGenerator()
var sqlStoredProcedureGetResultsSTATUSGenerator gopter.Gen

// SqlStoredProcedureGetResultsSTATUSGenerator returns a generator of SqlStoredProcedureGetResults_STATUS instances for property testing.
// We first initialize sqlStoredProcedureGetResultsSTATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlStoredProcedureGetResultsSTATUSGenerator() gopter.Gen {
	if sqlStoredProcedureGetResultsSTATUSGenerator != nil {
		return sqlStoredProcedureGetResultsSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResultsSTATUS(generators)
	sqlStoredProcedureGetResultsSTATUSGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureGetResults_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResultsSTATUS(generators)
	AddRelatedPropertyGeneratorsForSqlStoredProcedureGetResultsSTATUS(generators)
	sqlStoredProcedureGetResultsSTATUSGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureGetResults_STATUS{}), generators)

	return sqlStoredProcedureGetResultsSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResultsSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResultsSTATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlStoredProcedureGetResultsSTATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlStoredProcedureGetResultsSTATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlStoredProcedureGetPropertiesSTATUSResourceGenerator())
}

func Test_SqlStoredProcedureGetProperties_STATUS_Resource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureGetProperties_STATUS_Resource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureGetPropertiesSTATUSResource, SqlStoredProcedureGetPropertiesSTATUSResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureGetPropertiesSTATUSResource runs a test to see if a specific instance of SqlStoredProcedureGetProperties_STATUS_Resource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureGetPropertiesSTATUSResource(subject SqlStoredProcedureGetProperties_STATUS_Resource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureGetProperties_STATUS_Resource
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

// Generator of SqlStoredProcedureGetProperties_STATUS_Resource instances for property testing - lazily instantiated by
// SqlStoredProcedureGetPropertiesSTATUSResourceGenerator()
var sqlStoredProcedureGetPropertiesSTATUSResourceGenerator gopter.Gen

// SqlStoredProcedureGetPropertiesSTATUSResourceGenerator returns a generator of SqlStoredProcedureGetProperties_STATUS_Resource instances for property testing.
func SqlStoredProcedureGetPropertiesSTATUSResourceGenerator() gopter.Gen {
	if sqlStoredProcedureGetPropertiesSTATUSResourceGenerator != nil {
		return sqlStoredProcedureGetPropertiesSTATUSResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureGetPropertiesSTATUSResource(generators)
	sqlStoredProcedureGetPropertiesSTATUSResourceGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureGetProperties_STATUS_Resource{}), generators)

	return sqlStoredProcedureGetPropertiesSTATUSResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureGetPropertiesSTATUSResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureGetPropertiesSTATUSResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
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
