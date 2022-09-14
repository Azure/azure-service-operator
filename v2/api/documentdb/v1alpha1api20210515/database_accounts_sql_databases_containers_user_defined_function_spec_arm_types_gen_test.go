// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

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

func Test_DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM, DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM(subject DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM
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

// Generator of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM instances for property testing -
// lazily instantiated by DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARMGenerator()
var databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARMGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARMGenerator returns a generator of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARMGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARMGenerator != nil {
		return databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM(generators)
	databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM(generators)
	databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM{}), generators)

	return databaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlUserDefinedFunctionCreateUpdateProperties_ARMGenerator())
}

func Test_SqlUserDefinedFunctionCreateUpdateProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionCreateUpdateProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdateProperties_ARM, SqlUserDefinedFunctionCreateUpdateProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdateProperties_ARM runs a test to see if a specific instance of SqlUserDefinedFunctionCreateUpdateProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdateProperties_ARM(subject SqlUserDefinedFunctionCreateUpdateProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionCreateUpdateProperties_ARM
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

// Generator of SqlUserDefinedFunctionCreateUpdateProperties_ARM instances for property testing - lazily instantiated by
// SqlUserDefinedFunctionCreateUpdateProperties_ARMGenerator()
var sqlUserDefinedFunctionCreateUpdateProperties_ARMGenerator gopter.Gen

// SqlUserDefinedFunctionCreateUpdateProperties_ARMGenerator returns a generator of SqlUserDefinedFunctionCreateUpdateProperties_ARM instances for property testing.
func SqlUserDefinedFunctionCreateUpdateProperties_ARMGenerator() gopter.Gen {
	if sqlUserDefinedFunctionCreateUpdateProperties_ARMGenerator != nil {
		return sqlUserDefinedFunctionCreateUpdateProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdateProperties_ARM(generators)
	sqlUserDefinedFunctionCreateUpdateProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionCreateUpdateProperties_ARM{}), generators)

	return sqlUserDefinedFunctionCreateUpdateProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdateProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdateProperties_ARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptions_ARMGenerator())
	gens["Resource"] = gen.PtrOf(SqlUserDefinedFunctionResource_ARMGenerator())
}

func Test_SqlUserDefinedFunctionResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionResource_ARM, SqlUserDefinedFunctionResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionResource_ARM runs a test to see if a specific instance of SqlUserDefinedFunctionResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionResource_ARM(subject SqlUserDefinedFunctionResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionResource_ARM
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

// Generator of SqlUserDefinedFunctionResource_ARM instances for property testing - lazily instantiated by
// SqlUserDefinedFunctionResource_ARMGenerator()
var sqlUserDefinedFunctionResource_ARMGenerator gopter.Gen

// SqlUserDefinedFunctionResource_ARMGenerator returns a generator of SqlUserDefinedFunctionResource_ARM instances for property testing.
func SqlUserDefinedFunctionResource_ARMGenerator() gopter.Gen {
	if sqlUserDefinedFunctionResource_ARMGenerator != nil {
		return sqlUserDefinedFunctionResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource_ARM(generators)
	sqlUserDefinedFunctionResource_ARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionResource_ARM{}), generators)

	return sqlUserDefinedFunctionResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource_ARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
