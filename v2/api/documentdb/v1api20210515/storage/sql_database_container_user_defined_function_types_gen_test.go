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

func Test_SqlDatabaseContainerUserDefinedFunction_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerUserDefinedFunction via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction, SqlDatabaseContainerUserDefinedFunctionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction runs a test to see if a specific instance of SqlDatabaseContainerUserDefinedFunction round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction(subject SqlDatabaseContainerUserDefinedFunction) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerUserDefinedFunction
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

// Generator of SqlDatabaseContainerUserDefinedFunction instances for property testing - lazily instantiated by
// SqlDatabaseContainerUserDefinedFunctionGenerator()
var sqlDatabaseContainerUserDefinedFunctionGenerator gopter.Gen

// SqlDatabaseContainerUserDefinedFunctionGenerator returns a generator of SqlDatabaseContainerUserDefinedFunction instances for property testing.
func SqlDatabaseContainerUserDefinedFunctionGenerator() gopter.Gen {
	if sqlDatabaseContainerUserDefinedFunctionGenerator != nil {
		return sqlDatabaseContainerUserDefinedFunctionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction(generators)
	sqlDatabaseContainerUserDefinedFunctionGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerUserDefinedFunction{}), generators)

	return sqlDatabaseContainerUserDefinedFunctionGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator()
	gens["Status"] = DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator()
}

func Test_DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec, DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec(subject DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec
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

// Generator of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec instances for property testing -
// lazily instantiated by DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator()
var databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator returns a generator of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator != nil {
		return databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec(generators)
	databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec(generators)
	databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec{}), generators)

	return databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlUserDefinedFunctionResourceGenerator())
}

func Test_DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS, DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS(subject DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS
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

// Generator of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS instances for property testing -
// lazily instantiated by DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator()
var databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator returns a generator of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator != nil {
		return databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS(generators)
	databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS(generators)
	databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS{}), generators)

	return databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator())
}

func Test_SqlUserDefinedFunctionGetProperties_Resource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionGetProperties_Resource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionGetProperties_Resource_STATUS, SqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionGetProperties_Resource_STATUS runs a test to see if a specific instance of SqlUserDefinedFunctionGetProperties_Resource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionGetProperties_Resource_STATUS(subject SqlUserDefinedFunctionGetProperties_Resource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionGetProperties_Resource_STATUS
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

// Generator of SqlUserDefinedFunctionGetProperties_Resource_STATUS instances for property testing - lazily instantiated
// by SqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator()
var sqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator gopter.Gen

// SqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator returns a generator of SqlUserDefinedFunctionGetProperties_Resource_STATUS instances for property testing.
func SqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator() gopter.Gen {
	if sqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator != nil {
		return sqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetProperties_Resource_STATUS(generators)
	sqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetProperties_Resource_STATUS{}), generators)

	return sqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

func Test_SqlUserDefinedFunctionResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionResource, SqlUserDefinedFunctionResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionResource runs a test to see if a specific instance of SqlUserDefinedFunctionResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionResource(subject SqlUserDefinedFunctionResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionResource
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

// Generator of SqlUserDefinedFunctionResource instances for property testing - lazily instantiated by
// SqlUserDefinedFunctionResourceGenerator()
var sqlUserDefinedFunctionResourceGenerator gopter.Gen

// SqlUserDefinedFunctionResourceGenerator returns a generator of SqlUserDefinedFunctionResource instances for property testing.
func SqlUserDefinedFunctionResourceGenerator() gopter.Gen {
	if sqlUserDefinedFunctionResourceGenerator != nil {
		return sqlUserDefinedFunctionResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource(generators)
	sqlUserDefinedFunctionResourceGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionResource{}), generators)

	return sqlUserDefinedFunctionResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
