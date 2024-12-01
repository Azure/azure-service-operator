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

func Test_SqlDatabaseCreateUpdateProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseCreateUpdateProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseCreateUpdateProperties, SqlDatabaseCreateUpdatePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseCreateUpdateProperties runs a test to see if a specific instance of SqlDatabaseCreateUpdateProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseCreateUpdateProperties(subject SqlDatabaseCreateUpdateProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseCreateUpdateProperties
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

// Generator of SqlDatabaseCreateUpdateProperties instances for property testing - lazily instantiated by
// SqlDatabaseCreateUpdatePropertiesGenerator()
var sqlDatabaseCreateUpdatePropertiesGenerator gopter.Gen

// SqlDatabaseCreateUpdatePropertiesGenerator returns a generator of SqlDatabaseCreateUpdateProperties instances for property testing.
func SqlDatabaseCreateUpdatePropertiesGenerator() gopter.Gen {
	if sqlDatabaseCreateUpdatePropertiesGenerator != nil {
		return sqlDatabaseCreateUpdatePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseCreateUpdateProperties(generators)
	sqlDatabaseCreateUpdatePropertiesGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseCreateUpdateProperties{}), generators)

	return sqlDatabaseCreateUpdatePropertiesGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseCreateUpdateProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseCreateUpdateProperties(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlDatabaseResourceGenerator())
}

func Test_SqlDatabaseResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseResource, SqlDatabaseResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseResource runs a test to see if a specific instance of SqlDatabaseResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseResource(subject SqlDatabaseResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseResource
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

// Generator of SqlDatabaseResource instances for property testing - lazily instantiated by
// SqlDatabaseResourceGenerator()
var sqlDatabaseResourceGenerator gopter.Gen

// SqlDatabaseResourceGenerator returns a generator of SqlDatabaseResource instances for property testing.
func SqlDatabaseResourceGenerator() gopter.Gen {
	if sqlDatabaseResourceGenerator != nil {
		return sqlDatabaseResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseResource(generators)
	sqlDatabaseResourceGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseResource{}), generators)

	return sqlDatabaseResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseResource(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_SqlDatabase_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabase_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabase_Spec, SqlDatabase_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabase_Spec runs a test to see if a specific instance of SqlDatabase_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabase_Spec(subject SqlDatabase_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabase_Spec
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

// Generator of SqlDatabase_Spec instances for property testing - lazily instantiated by SqlDatabase_SpecGenerator()
var sqlDatabase_SpecGenerator gopter.Gen

// SqlDatabase_SpecGenerator returns a generator of SqlDatabase_Spec instances for property testing.
// We first initialize sqlDatabase_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabase_SpecGenerator() gopter.Gen {
	if sqlDatabase_SpecGenerator != nil {
		return sqlDatabase_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabase_Spec(generators)
	sqlDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabase_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabase_Spec(generators)
	AddRelatedPropertyGeneratorsForSqlDatabase_Spec(generators)
	sqlDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabase_Spec{}), generators)

	return sqlDatabase_SpecGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabase_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabase_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabase_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlDatabaseCreateUpdatePropertiesGenerator())
}
