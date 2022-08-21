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

func Test_SqlDatabaseContainerTrigger_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerTrigger via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerTrigger, SqlDatabaseContainerTriggerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerTrigger runs a test to see if a specific instance of SqlDatabaseContainerTrigger round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerTrigger(subject SqlDatabaseContainerTrigger) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerTrigger
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

// Generator of SqlDatabaseContainerTrigger instances for property testing - lazily instantiated by
// SqlDatabaseContainerTriggerGenerator()
var sqlDatabaseContainerTriggerGenerator gopter.Gen

// SqlDatabaseContainerTriggerGenerator returns a generator of SqlDatabaseContainerTrigger instances for property testing.
func SqlDatabaseContainerTriggerGenerator() gopter.Gen {
	if sqlDatabaseContainerTriggerGenerator != nil {
		return sqlDatabaseContainerTriggerGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerTrigger(generators)
	sqlDatabaseContainerTriggerGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerTrigger{}), generators)

	return sqlDatabaseContainerTriggerGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerTrigger is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerTrigger(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesContainersTrigger_SpecGenerator()
	gens["Status"] = DatabaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator()
}

func Test_DatabaseAccountsSqlDatabasesContainersTrigger_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersTrigger_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTrigger_Spec, DatabaseAccountsSqlDatabasesContainersTrigger_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTrigger_Spec runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersTrigger_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTrigger_Spec(subject DatabaseAccountsSqlDatabasesContainersTrigger_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersTrigger_Spec
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

// Generator of DatabaseAccountsSqlDatabasesContainersTrigger_Spec instances for property testing - lazily instantiated
// by DatabaseAccountsSqlDatabasesContainersTrigger_SpecGenerator()
var databaseAccountsSqlDatabasesContainersTrigger_SpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersTrigger_SpecGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersTrigger_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersTrigger_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersTrigger_SpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersTrigger_SpecGenerator != nil {
		return databaseAccountsSqlDatabasesContainersTrigger_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_Spec(generators)
	databaseAccountsSqlDatabasesContainersTrigger_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersTrigger_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_Spec(generators)
	databaseAccountsSqlDatabasesContainersTrigger_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersTrigger_Spec{}), generators)

	return databaseAccountsSqlDatabasesContainersTrigger_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_Spec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlTriggerResourceGenerator())
}

func Test_DatabaseAccountsSqlDatabasesContainersTrigger_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersTrigger_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTrigger_STATUS, DatabaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTrigger_STATUS runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersTrigger_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTrigger_STATUS(subject DatabaseAccountsSqlDatabasesContainersTrigger_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersTrigger_STATUS
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

// Generator of DatabaseAccountsSqlDatabasesContainersTrigger_STATUS instances for property testing - lazily
// instantiated by DatabaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator()
var databaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersTrigger_STATUS instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator != nil {
		return databaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_STATUS(generators)
	databaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersTrigger_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_STATUS(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_STATUS(generators)
	databaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersTrigger_STATUS{}), generators)

	return databaseAccountsSqlDatabasesContainersTrigger_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTrigger_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlTriggerGetProperties_Resource_STATUSGenerator())
}

func Test_SqlTriggerGetProperties_Resource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetProperties_Resource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetProperties_Resource_STATUS, SqlTriggerGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetProperties_Resource_STATUS runs a test to see if a specific instance of SqlTriggerGetProperties_Resource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetProperties_Resource_STATUS(subject SqlTriggerGetProperties_Resource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetProperties_Resource_STATUS
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

// Generator of SqlTriggerGetProperties_Resource_STATUS instances for property testing - lazily instantiated by
// SqlTriggerGetProperties_Resource_STATUSGenerator()
var sqlTriggerGetProperties_Resource_STATUSGenerator gopter.Gen

// SqlTriggerGetProperties_Resource_STATUSGenerator returns a generator of SqlTriggerGetProperties_Resource_STATUS instances for property testing.
func SqlTriggerGetProperties_Resource_STATUSGenerator() gopter.Gen {
	if sqlTriggerGetProperties_Resource_STATUSGenerator != nil {
		return sqlTriggerGetProperties_Resource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetProperties_Resource_STATUS(generators)
	sqlTriggerGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetProperties_Resource_STATUS{}), generators)

	return sqlTriggerGetProperties_Resource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerOperation"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerType"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

func Test_SqlTriggerResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerResource, SqlTriggerResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerResource runs a test to see if a specific instance of SqlTriggerResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerResource(subject SqlTriggerResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerResource
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

// Generator of SqlTriggerResource instances for property testing - lazily instantiated by SqlTriggerResourceGenerator()
var sqlTriggerResourceGenerator gopter.Gen

// SqlTriggerResourceGenerator returns a generator of SqlTriggerResource instances for property testing.
func SqlTriggerResourceGenerator() gopter.Gen {
	if sqlTriggerResourceGenerator != nil {
		return sqlTriggerResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerResource(generators)
	sqlTriggerResourceGenerator = gen.Struct(reflect.TypeOf(SqlTriggerResource{}), generators)

	return sqlTriggerResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerOperation"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerType"] = gen.PtrOf(gen.AlphaString())
}
