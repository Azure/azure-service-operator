// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

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

func Test_Database_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseSTATUSARM, DatabaseSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseSTATUSARM runs a test to see if a specific instance of Database_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseSTATUSARM(subject Database_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_STATUSARM
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

// Generator of Database_STATUSARM instances for property testing - lazily instantiated by DatabaseSTATUSARMGenerator()
var databaseSTATUSARMGenerator gopter.Gen

// DatabaseSTATUSARMGenerator returns a generator of Database_STATUSARM instances for property testing.
// We first initialize databaseSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseSTATUSARMGenerator() gopter.Gen {
	if databaseSTATUSARMGenerator != nil {
		return databaseSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseSTATUSARM(generators)
	databaseSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Database_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseSTATUSARM(generators)
	databaseSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Database_STATUSARM{}), generators)

	return databaseSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabasePropertiesSTATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataSTATUSARMGenerator())
}

func Test_DatabaseProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabasePropertiesSTATUSARM, DatabasePropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabasePropertiesSTATUSARM runs a test to see if a specific instance of DatabaseProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabasePropertiesSTATUSARM(subject DatabaseProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_STATUSARM
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

// Generator of DatabaseProperties_STATUSARM instances for property testing - lazily instantiated by
// DatabasePropertiesSTATUSARMGenerator()
var databasePropertiesSTATUSARMGenerator gopter.Gen

// DatabasePropertiesSTATUSARMGenerator returns a generator of DatabaseProperties_STATUSARM instances for property testing.
func DatabasePropertiesSTATUSARMGenerator() gopter.Gen {
	if databasePropertiesSTATUSARMGenerator != nil {
		return databasePropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabasePropertiesSTATUSARM(generators)
	databasePropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_STATUSARM{}), generators)

	return databasePropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabasePropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabasePropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemData_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataSTATUSARM, SystemDataSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataSTATUSARM runs a test to see if a specific instance of SystemData_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataSTATUSARM(subject SystemData_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUSARM
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

// Generator of SystemData_STATUSARM instances for property testing - lazily instantiated by
// SystemDataSTATUSARMGenerator()
var systemDataSTATUSARMGenerator gopter.Gen

// SystemDataSTATUSARMGenerator returns a generator of SystemData_STATUSARM instances for property testing.
func SystemDataSTATUSARMGenerator() gopter.Gen {
	if systemDataSTATUSARMGenerator != nil {
		return systemDataSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataSTATUSARM(generators)
	systemDataSTATUSARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUSARM{}), generators)

	return systemDataSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataSTATUSARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataSTATUSCreatedByType_Application,
		SystemDataSTATUSCreatedByType_Key,
		SystemDataSTATUSCreatedByType_ManagedIdentity,
		SystemDataSTATUSCreatedByType_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataSTATUSLastModifiedByType_Application,
		SystemDataSTATUSLastModifiedByType_Key,
		SystemDataSTATUSLastModifiedByType_ManagedIdentity,
		SystemDataSTATUSLastModifiedByType_User))
}
