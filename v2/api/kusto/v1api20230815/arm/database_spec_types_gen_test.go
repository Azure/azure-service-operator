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

func Test_Database_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabase_Spec, Database_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabase_Spec runs a test to see if a specific instance of Database_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabase_Spec(subject Database_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_Spec
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

// Generator of Database_Spec instances for property testing - lazily instantiated by Database_SpecGenerator()
var database_SpecGenerator gopter.Gen

// Database_SpecGenerator returns a generator of Database_Spec instances for property testing.
// We first initialize database_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Database_SpecGenerator() gopter.Gen {
	if database_SpecGenerator != nil {
		return database_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabase_Spec(generators)
	database_SpecGenerator = gen.OneGenOf(gens...)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabase_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabase_Spec(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(Database_Spec{}), map[string]gopter.Gen{propName: propGen}))
	}
	database_SpecGenerator = gen.OneGenOf(gens...)

	return database_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabase_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForDatabase_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabase_Spec(gens map[string]gopter.Gen) {
	gens["ReadOnlyFollowing"] = ReadOnlyFollowingDatabaseGenerator().Map(func(it ReadOnlyFollowingDatabase) *ReadOnlyFollowingDatabase {
		return &it
	}) // generate one case for OneOf type
	gens["ReadWrite"] = ReadWriteDatabaseGenerator().Map(func(it ReadWriteDatabase) *ReadWriteDatabase {
		return &it
	}) // generate one case for OneOf type
}

func Test_ReadOnlyFollowingDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadOnlyFollowingDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadOnlyFollowingDatabase, ReadOnlyFollowingDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadOnlyFollowingDatabase runs a test to see if a specific instance of ReadOnlyFollowingDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForReadOnlyFollowingDatabase(subject ReadOnlyFollowingDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadOnlyFollowingDatabase
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

// Generator of ReadOnlyFollowingDatabase instances for property testing - lazily instantiated by
// ReadOnlyFollowingDatabaseGenerator()
var readOnlyFollowingDatabaseGenerator gopter.Gen

// ReadOnlyFollowingDatabaseGenerator returns a generator of ReadOnlyFollowingDatabase instances for property testing.
// We first initialize readOnlyFollowingDatabaseGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadOnlyFollowingDatabaseGenerator() gopter.Gen {
	if readOnlyFollowingDatabaseGenerator != nil {
		return readOnlyFollowingDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase(generators)
	readOnlyFollowingDatabaseGenerator = gen.Struct(reflect.TypeOf(ReadOnlyFollowingDatabase{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase(generators)
	AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabase(generators)
	readOnlyFollowingDatabaseGenerator = gen.Struct(reflect.TypeOf(ReadOnlyFollowingDatabase{}), generators)

	return readOnlyFollowingDatabaseGenerator
}

// AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.OneConstOf(ReadOnlyFollowingDatabase_Kind_ReadOnlyFollowing)
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabase(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ReadOnlyFollowingDatabasePropertiesGenerator())
}

func Test_ReadOnlyFollowingDatabaseProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadOnlyFollowingDatabaseProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadOnlyFollowingDatabaseProperties, ReadOnlyFollowingDatabasePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadOnlyFollowingDatabaseProperties runs a test to see if a specific instance of ReadOnlyFollowingDatabaseProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForReadOnlyFollowingDatabaseProperties(subject ReadOnlyFollowingDatabaseProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadOnlyFollowingDatabaseProperties
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

// Generator of ReadOnlyFollowingDatabaseProperties instances for property testing - lazily instantiated by
// ReadOnlyFollowingDatabasePropertiesGenerator()
var readOnlyFollowingDatabasePropertiesGenerator gopter.Gen

// ReadOnlyFollowingDatabasePropertiesGenerator returns a generator of ReadOnlyFollowingDatabaseProperties instances for property testing.
func ReadOnlyFollowingDatabasePropertiesGenerator() gopter.Gen {
	if readOnlyFollowingDatabasePropertiesGenerator != nil {
		return readOnlyFollowingDatabasePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabaseProperties(generators)
	readOnlyFollowingDatabasePropertiesGenerator = gen.Struct(reflect.TypeOf(ReadOnlyFollowingDatabaseProperties{}), generators)

	return readOnlyFollowingDatabasePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabaseProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabaseProperties(gens map[string]gopter.Gen) {
	gens["DatabaseShareOrigin"] = gen.PtrOf(gen.OneConstOf(DatabaseShareOrigin_DataShare, DatabaseShareOrigin_Direct, DatabaseShareOrigin_Other))
	gens["HotCachePeriod"] = gen.PtrOf(gen.AlphaString())
}

func Test_ReadWriteDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadWriteDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadWriteDatabase, ReadWriteDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadWriteDatabase runs a test to see if a specific instance of ReadWriteDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForReadWriteDatabase(subject ReadWriteDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadWriteDatabase
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

// Generator of ReadWriteDatabase instances for property testing - lazily instantiated by ReadWriteDatabaseGenerator()
var readWriteDatabaseGenerator gopter.Gen

// ReadWriteDatabaseGenerator returns a generator of ReadWriteDatabase instances for property testing.
// We first initialize readWriteDatabaseGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadWriteDatabaseGenerator() gopter.Gen {
	if readWriteDatabaseGenerator != nil {
		return readWriteDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabase(generators)
	readWriteDatabaseGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabase{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabase(generators)
	AddRelatedPropertyGeneratorsForReadWriteDatabase(generators)
	readWriteDatabaseGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabase{}), generators)

	return readWriteDatabaseGenerator
}

// AddIndependentPropertyGeneratorsForReadWriteDatabase is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadWriteDatabase(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.OneConstOf(ReadWriteDatabase_Kind_ReadWrite)
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForReadWriteDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadWriteDatabase(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ReadWriteDatabasePropertiesGenerator())
}

func Test_ReadWriteDatabaseProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadWriteDatabaseProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadWriteDatabaseProperties, ReadWriteDatabasePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadWriteDatabaseProperties runs a test to see if a specific instance of ReadWriteDatabaseProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForReadWriteDatabaseProperties(subject ReadWriteDatabaseProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadWriteDatabaseProperties
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

// Generator of ReadWriteDatabaseProperties instances for property testing - lazily instantiated by
// ReadWriteDatabasePropertiesGenerator()
var readWriteDatabasePropertiesGenerator gopter.Gen

// ReadWriteDatabasePropertiesGenerator returns a generator of ReadWriteDatabaseProperties instances for property testing.
// We first initialize readWriteDatabasePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadWriteDatabasePropertiesGenerator() gopter.Gen {
	if readWriteDatabasePropertiesGenerator != nil {
		return readWriteDatabasePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties(generators)
	readWriteDatabasePropertiesGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabaseProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties(generators)
	AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties(generators)
	readWriteDatabasePropertiesGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabaseProperties{}), generators)

	return readWriteDatabasePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties(gens map[string]gopter.Gen) {
	gens["HotCachePeriod"] = gen.PtrOf(gen.AlphaString())
	gens["SoftDeletePeriod"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultPropertiesGenerator())
}
