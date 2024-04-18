// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

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

func Test_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM, DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM(subject DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM
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

// Generator of DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM instances for property testing - lazily
// instantiated by DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARMGenerator()
var databaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARMGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARMGenerator returns a generator of DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARMGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARMGenerator != nil {
		return databaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM(generators)
	databaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM(generators)
	databaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM{}), generators)

	return databaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlTriggerCreateUpdateProperties_ARMGenerator())
}

func Test_SqlTriggerCreateUpdateProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerCreateUpdateProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerCreateUpdateProperties_ARM, SqlTriggerCreateUpdateProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerCreateUpdateProperties_ARM runs a test to see if a specific instance of SqlTriggerCreateUpdateProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerCreateUpdateProperties_ARM(subject SqlTriggerCreateUpdateProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerCreateUpdateProperties_ARM
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

// Generator of SqlTriggerCreateUpdateProperties_ARM instances for property testing - lazily instantiated by
// SqlTriggerCreateUpdateProperties_ARMGenerator()
var sqlTriggerCreateUpdateProperties_ARMGenerator gopter.Gen

// SqlTriggerCreateUpdateProperties_ARMGenerator returns a generator of SqlTriggerCreateUpdateProperties_ARM instances for property testing.
func SqlTriggerCreateUpdateProperties_ARMGenerator() gopter.Gen {
	if sqlTriggerCreateUpdateProperties_ARMGenerator != nil {
		return sqlTriggerCreateUpdateProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlTriggerCreateUpdateProperties_ARM(generators)
	sqlTriggerCreateUpdateProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerCreateUpdateProperties_ARM{}), generators)

	return sqlTriggerCreateUpdateProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlTriggerCreateUpdateProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTriggerCreateUpdateProperties_ARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptions_ARMGenerator())
	gens["Resource"] = gen.PtrOf(SqlTriggerResource_ARMGenerator())
}

func Test_SqlTriggerResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerResource_ARM, SqlTriggerResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerResource_ARM runs a test to see if a specific instance of SqlTriggerResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerResource_ARM(subject SqlTriggerResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerResource_ARM
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

// Generator of SqlTriggerResource_ARM instances for property testing - lazily instantiated by
// SqlTriggerResource_ARMGenerator()
var sqlTriggerResource_ARMGenerator gopter.Gen

// SqlTriggerResource_ARMGenerator returns a generator of SqlTriggerResource_ARM instances for property testing.
func SqlTriggerResource_ARMGenerator() gopter.Gen {
	if sqlTriggerResource_ARMGenerator != nil {
		return sqlTriggerResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerResource_ARM(generators)
	sqlTriggerResource_ARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerResource_ARM{}), generators)

	return sqlTriggerResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerResource_ARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerOperation"] = gen.PtrOf(gen.OneConstOf(
		SqlTriggerResource_TriggerOperation_All,
		SqlTriggerResource_TriggerOperation_Create,
		SqlTriggerResource_TriggerOperation_Delete,
		SqlTriggerResource_TriggerOperation_Replace,
		SqlTriggerResource_TriggerOperation_Update))
	gens["TriggerType"] = gen.PtrOf(gen.OneConstOf(SqlTriggerResource_TriggerType_Post, SqlTriggerResource_TriggerType_Pre))
}
