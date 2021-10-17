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

func Test_SqlUserDefinedFunctionGetResults_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionGetResults_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionGetResultsStatusARM, SqlUserDefinedFunctionGetResultsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionGetResultsStatusARM runs a test to see if a specific instance of SqlUserDefinedFunctionGetResults_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionGetResultsStatusARM(subject SqlUserDefinedFunctionGetResults_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionGetResults_StatusARM
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

// Generator of SqlUserDefinedFunctionGetResults_StatusARM instances for property testing - lazily instantiated by
//SqlUserDefinedFunctionGetResultsStatusARMGenerator()
var sqlUserDefinedFunctionGetResultsStatusARMGenerator gopter.Gen

// SqlUserDefinedFunctionGetResultsStatusARMGenerator returns a generator of SqlUserDefinedFunctionGetResults_StatusARM instances for property testing.
// We first initialize sqlUserDefinedFunctionGetResultsStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlUserDefinedFunctionGetResultsStatusARMGenerator() gopter.Gen {
	if sqlUserDefinedFunctionGetResultsStatusARMGenerator != nil {
		return sqlUserDefinedFunctionGetResultsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatusARM(generators)
	sqlUserDefinedFunctionGetResultsStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetResults_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatusARM(generators)
	AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatusARM(generators)
	sqlUserDefinedFunctionGetResultsStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetResults_StatusARM{}), generators)

	return sqlUserDefinedFunctionGetResultsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlUserDefinedFunctionGetPropertiesStatusARMGenerator())
}

func Test_SqlUserDefinedFunctionGetProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionGetProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionGetPropertiesStatusARM, SqlUserDefinedFunctionGetPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionGetPropertiesStatusARM runs a test to see if a specific instance of SqlUserDefinedFunctionGetProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionGetPropertiesStatusARM(subject SqlUserDefinedFunctionGetProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionGetProperties_StatusARM
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

// Generator of SqlUserDefinedFunctionGetProperties_StatusARM instances for property testing - lazily instantiated by
//SqlUserDefinedFunctionGetPropertiesStatusARMGenerator()
var sqlUserDefinedFunctionGetPropertiesStatusARMGenerator gopter.Gen

// SqlUserDefinedFunctionGetPropertiesStatusARMGenerator returns a generator of SqlUserDefinedFunctionGetProperties_StatusARM instances for property testing.
func SqlUserDefinedFunctionGetPropertiesStatusARMGenerator() gopter.Gen {
	if sqlUserDefinedFunctionGetPropertiesStatusARMGenerator != nil {
		return sqlUserDefinedFunctionGetPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetPropertiesStatusARM(generators)
	sqlUserDefinedFunctionGetPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetProperties_StatusARM{}), generators)

	return sqlUserDefinedFunctionGetPropertiesStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlUserDefinedFunctionGetPropertiesStatusResourceARMGenerator())
}

func Test_SqlUserDefinedFunctionGetProperties_Status_ResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionGetProperties_Status_ResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionGetPropertiesStatusResourceARM, SqlUserDefinedFunctionGetPropertiesStatusResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionGetPropertiesStatusResourceARM runs a test to see if a specific instance of SqlUserDefinedFunctionGetProperties_Status_ResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionGetPropertiesStatusResourceARM(subject SqlUserDefinedFunctionGetProperties_Status_ResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionGetProperties_Status_ResourceARM
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

// Generator of SqlUserDefinedFunctionGetProperties_Status_ResourceARM instances for property testing - lazily
//instantiated by SqlUserDefinedFunctionGetPropertiesStatusResourceARMGenerator()
var sqlUserDefinedFunctionGetPropertiesStatusResourceARMGenerator gopter.Gen

// SqlUserDefinedFunctionGetPropertiesStatusResourceARMGenerator returns a generator of SqlUserDefinedFunctionGetProperties_Status_ResourceARM instances for property testing.
func SqlUserDefinedFunctionGetPropertiesStatusResourceARMGenerator() gopter.Gen {
	if sqlUserDefinedFunctionGetPropertiesStatusResourceARMGenerator != nil {
		return sqlUserDefinedFunctionGetPropertiesStatusResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetPropertiesStatusResourceARM(generators)
	sqlUserDefinedFunctionGetPropertiesStatusResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetProperties_Status_ResourceARM{}), generators)

	return sqlUserDefinedFunctionGetPropertiesStatusResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetPropertiesStatusResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetPropertiesStatusResourceARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.AlphaString()
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}
