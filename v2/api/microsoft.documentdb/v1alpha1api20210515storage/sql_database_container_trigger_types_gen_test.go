// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

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
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
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
//SqlDatabaseContainerTriggerGenerator()
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
	gens["Spec"] = DatabaseAccountsSqlDatabasesContainersTriggersSpecGenerator()
	gens["Status"] = SqlTriggerGetResultsStatusGenerator()
}

func Test_DatabaseAccountsSqlDatabasesContainersTriggers_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersTriggers_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTriggersSpec, DatabaseAccountsSqlDatabasesContainersTriggersSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTriggersSpec runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersTriggers_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTriggersSpec(subject DatabaseAccountsSqlDatabasesContainersTriggers_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersTriggers_Spec
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

// Generator of DatabaseAccountsSqlDatabasesContainersTriggers_Spec instances for property testing - lazily instantiated
//by DatabaseAccountsSqlDatabasesContainersTriggersSpecGenerator()
var databaseAccountsSqlDatabasesContainersTriggersSpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersTriggersSpecGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersTriggers_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersTriggersSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersTriggersSpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersTriggersSpecGenerator != nil {
		return databaseAccountsSqlDatabasesContainersTriggersSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpec(generators)
	databaseAccountsSqlDatabasesContainersTriggersSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersTriggers_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpec(generators)
	databaseAccountsSqlDatabasesContainersTriggersSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersTriggers_Spec{}), generators)

	return databaseAccountsSqlDatabasesContainersTriggersSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlTriggerResourceGenerator())
}

func Test_SqlTriggerGetResults_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetResults_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetResultsStatus, SqlTriggerGetResultsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetResultsStatus runs a test to see if a specific instance of SqlTriggerGetResults_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetResultsStatus(subject SqlTriggerGetResults_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetResults_Status
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

// Generator of SqlTriggerGetResults_Status instances for property testing - lazily instantiated by
//SqlTriggerGetResultsStatusGenerator()
var sqlTriggerGetResultsStatusGenerator gopter.Gen

// SqlTriggerGetResultsStatusGenerator returns a generator of SqlTriggerGetResults_Status instances for property testing.
// We first initialize sqlTriggerGetResultsStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlTriggerGetResultsStatusGenerator() gopter.Gen {
	if sqlTriggerGetResultsStatusGenerator != nil {
		return sqlTriggerGetResultsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetResultsStatus(generators)
	sqlTriggerGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetResults_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetResultsStatus(generators)
	AddRelatedPropertyGeneratorsForSqlTriggerGetResultsStatus(generators)
	sqlTriggerGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetResults_Status{}), generators)

	return sqlTriggerGetResultsStatusGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerGetResultsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlTriggerGetResultsStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTriggerGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlTriggerGetPropertiesStatusResourceGenerator())
}

func Test_SqlTriggerGetProperties_Status_Resource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetProperties_Status_Resource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetPropertiesStatusResource, SqlTriggerGetPropertiesStatusResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetPropertiesStatusResource runs a test to see if a specific instance of SqlTriggerGetProperties_Status_Resource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetPropertiesStatusResource(subject SqlTriggerGetProperties_Status_Resource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetProperties_Status_Resource
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

// Generator of SqlTriggerGetProperties_Status_Resource instances for property testing - lazily instantiated by
//SqlTriggerGetPropertiesStatusResourceGenerator()
var sqlTriggerGetPropertiesStatusResourceGenerator gopter.Gen

// SqlTriggerGetPropertiesStatusResourceGenerator returns a generator of SqlTriggerGetProperties_Status_Resource instances for property testing.
func SqlTriggerGetPropertiesStatusResourceGenerator() gopter.Gen {
	if sqlTriggerGetPropertiesStatusResourceGenerator != nil {
		return sqlTriggerGetPropertiesStatusResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetPropertiesStatusResource(generators)
	sqlTriggerGetPropertiesStatusResourceGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetProperties_Status_Resource{}), generators)

	return sqlTriggerGetPropertiesStatusResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerGetPropertiesStatusResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerGetPropertiesStatusResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerOperation"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerType"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

func Test_SqlTriggerResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
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
