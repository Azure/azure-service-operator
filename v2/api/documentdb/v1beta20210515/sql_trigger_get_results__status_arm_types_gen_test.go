// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

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

func Test_SqlTriggerGetResults_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetResults_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetResultsStatusARM, SqlTriggerGetResultsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetResultsStatusARM runs a test to see if a specific instance of SqlTriggerGetResults_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetResultsStatusARM(subject SqlTriggerGetResults_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetResults_StatusARM
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

// Generator of SqlTriggerGetResults_StatusARM instances for property testing - lazily instantiated by
// SqlTriggerGetResultsStatusARMGenerator()
var sqlTriggerGetResultsStatusARMGenerator gopter.Gen

// SqlTriggerGetResultsStatusARMGenerator returns a generator of SqlTriggerGetResults_StatusARM instances for property testing.
// We first initialize sqlTriggerGetResultsStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlTriggerGetResultsStatusARMGenerator() gopter.Gen {
	if sqlTriggerGetResultsStatusARMGenerator != nil {
		return sqlTriggerGetResultsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetResultsStatusARM(generators)
	sqlTriggerGetResultsStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetResults_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetResultsStatusARM(generators)
	AddRelatedPropertyGeneratorsForSqlTriggerGetResultsStatusARM(generators)
	sqlTriggerGetResultsStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetResults_StatusARM{}), generators)

	return sqlTriggerGetResultsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerGetResultsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerGetResultsStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlTriggerGetResultsStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTriggerGetResultsStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlTriggerGetPropertiesStatusARMGenerator())
}

func Test_SqlTriggerGetProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetPropertiesStatusARM, SqlTriggerGetPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetPropertiesStatusARM runs a test to see if a specific instance of SqlTriggerGetProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetPropertiesStatusARM(subject SqlTriggerGetProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetProperties_StatusARM
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

// Generator of SqlTriggerGetProperties_StatusARM instances for property testing - lazily instantiated by
// SqlTriggerGetPropertiesStatusARMGenerator()
var sqlTriggerGetPropertiesStatusARMGenerator gopter.Gen

// SqlTriggerGetPropertiesStatusARMGenerator returns a generator of SqlTriggerGetProperties_StatusARM instances for property testing.
func SqlTriggerGetPropertiesStatusARMGenerator() gopter.Gen {
	if sqlTriggerGetPropertiesStatusARMGenerator != nil {
		return sqlTriggerGetPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlTriggerGetPropertiesStatusARM(generators)
	sqlTriggerGetPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetProperties_StatusARM{}), generators)

	return sqlTriggerGetPropertiesStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlTriggerGetPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTriggerGetPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlTriggerGetPropertiesStatusResourceARMGenerator())
}

func Test_SqlTriggerGetProperties_Status_ResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetProperties_Status_ResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetPropertiesStatusResourceARM, SqlTriggerGetPropertiesStatusResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetPropertiesStatusResourceARM runs a test to see if a specific instance of SqlTriggerGetProperties_Status_ResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetPropertiesStatusResourceARM(subject SqlTriggerGetProperties_Status_ResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetProperties_Status_ResourceARM
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

// Generator of SqlTriggerGetProperties_Status_ResourceARM instances for property testing - lazily instantiated by
// SqlTriggerGetPropertiesStatusResourceARMGenerator()
var sqlTriggerGetPropertiesStatusResourceARMGenerator gopter.Gen

// SqlTriggerGetPropertiesStatusResourceARMGenerator returns a generator of SqlTriggerGetProperties_Status_ResourceARM instances for property testing.
func SqlTriggerGetPropertiesStatusResourceARMGenerator() gopter.Gen {
	if sqlTriggerGetPropertiesStatusResourceARMGenerator != nil {
		return sqlTriggerGetPropertiesStatusResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetPropertiesStatusResourceARM(generators)
	sqlTriggerGetPropertiesStatusResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetProperties_Status_ResourceARM{}), generators)

	return sqlTriggerGetPropertiesStatusResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerGetPropertiesStatusResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerGetPropertiesStatusResourceARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerOperation"] = gen.PtrOf(gen.OneConstOf(
		SqlTriggerGetPropertiesStatusResourceTriggerOperation_All,
		SqlTriggerGetPropertiesStatusResourceTriggerOperation_Create,
		SqlTriggerGetPropertiesStatusResourceTriggerOperation_Delete,
		SqlTriggerGetPropertiesStatusResourceTriggerOperation_Replace,
		SqlTriggerGetPropertiesStatusResourceTriggerOperation_Update))
	gens["TriggerType"] = gen.PtrOf(gen.OneConstOf(SqlTriggerGetPropertiesStatusResourceTriggerType_Post, SqlTriggerGetPropertiesStatusResourceTriggerType_Pre))
	gens["Ts"] = gen.PtrOf(gen.Float64())
}
