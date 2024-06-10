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

func Test_DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_Collection_Spec_ARM, DatabaseAccounts_MongodbDatabases_Collection_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_Collection_Spec_ARM runs a test to see if a specific instance of DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_Collection_Spec_ARM(subject DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM
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

// Generator of DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM instances for property testing - lazily
// instantiated by DatabaseAccounts_MongodbDatabases_Collection_Spec_ARMGenerator()
var databaseAccounts_MongodbDatabases_Collection_Spec_ARMGenerator gopter.Gen

// DatabaseAccounts_MongodbDatabases_Collection_Spec_ARMGenerator returns a generator of DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM instances for property testing.
// We first initialize databaseAccounts_MongodbDatabases_Collection_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_MongodbDatabases_Collection_Spec_ARMGenerator() gopter.Gen {
	if databaseAccounts_MongodbDatabases_Collection_Spec_ARMGenerator != nil {
		return databaseAccounts_MongodbDatabases_Collection_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_Collection_Spec_ARM(generators)
	databaseAccounts_MongodbDatabases_Collection_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_Collection_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_Collection_Spec_ARM(generators)
	databaseAccounts_MongodbDatabases_Collection_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM{}), generators)

	return databaseAccounts_MongodbDatabases_Collection_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_Collection_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_Collection_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_Collection_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_Collection_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBCollectionCreateUpdateProperties_ARMGenerator())
}

func Test_MongoDBCollectionCreateUpdateProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionCreateUpdateProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionCreateUpdateProperties_ARM, MongoDBCollectionCreateUpdateProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionCreateUpdateProperties_ARM runs a test to see if a specific instance of MongoDBCollectionCreateUpdateProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionCreateUpdateProperties_ARM(subject MongoDBCollectionCreateUpdateProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionCreateUpdateProperties_ARM
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

// Generator of MongoDBCollectionCreateUpdateProperties_ARM instances for property testing - lazily instantiated by
// MongoDBCollectionCreateUpdateProperties_ARMGenerator()
var mongoDBCollectionCreateUpdateProperties_ARMGenerator gopter.Gen

// MongoDBCollectionCreateUpdateProperties_ARMGenerator returns a generator of MongoDBCollectionCreateUpdateProperties_ARM instances for property testing.
func MongoDBCollectionCreateUpdateProperties_ARMGenerator() gopter.Gen {
	if mongoDBCollectionCreateUpdateProperties_ARMGenerator != nil {
		return mongoDBCollectionCreateUpdateProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBCollectionCreateUpdateProperties_ARM(generators)
	mongoDBCollectionCreateUpdateProperties_ARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionCreateUpdateProperties_ARM{}), generators)

	return mongoDBCollectionCreateUpdateProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionCreateUpdateProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionCreateUpdateProperties_ARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptions_ARMGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBCollectionResource_ARMGenerator())
}

func Test_MongoDBCollectionResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionResource_ARM, MongoDBCollectionResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionResource_ARM runs a test to see if a specific instance of MongoDBCollectionResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionResource_ARM(subject MongoDBCollectionResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionResource_ARM
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

// Generator of MongoDBCollectionResource_ARM instances for property testing - lazily instantiated by
// MongoDBCollectionResource_ARMGenerator()
var mongoDBCollectionResource_ARMGenerator gopter.Gen

// MongoDBCollectionResource_ARMGenerator returns a generator of MongoDBCollectionResource_ARM instances for property testing.
// We first initialize mongoDBCollectionResource_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionResource_ARMGenerator() gopter.Gen {
	if mongoDBCollectionResource_ARMGenerator != nil {
		return mongoDBCollectionResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResource_ARM(generators)
	mongoDBCollectionResource_ARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResource_ARM(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionResource_ARM(generators)
	mongoDBCollectionResource_ARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource_ARM{}), generators)

	return mongoDBCollectionResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionResource_ARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["ShardKey"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionResource_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionResource_ARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndex_ARMGenerator())
}

func Test_MongoIndexKeys_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeys_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeys_ARM, MongoIndexKeys_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeys_ARM runs a test to see if a specific instance of MongoIndexKeys_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeys_ARM(subject MongoIndexKeys_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeys_ARM
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

// Generator of MongoIndexKeys_ARM instances for property testing - lazily instantiated by MongoIndexKeys_ARMGenerator()
var mongoIndexKeys_ARMGenerator gopter.Gen

// MongoIndexKeys_ARMGenerator returns a generator of MongoIndexKeys_ARM instances for property testing.
func MongoIndexKeys_ARMGenerator() gopter.Gen {
	if mongoIndexKeys_ARMGenerator != nil {
		return mongoIndexKeys_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeys_ARM(generators)
	mongoIndexKeys_ARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeys_ARM{}), generators)

	return mongoIndexKeys_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeys_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeys_ARM(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexOptions_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptions_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptions_ARM, MongoIndexOptions_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptions_ARM runs a test to see if a specific instance of MongoIndexOptions_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptions_ARM(subject MongoIndexOptions_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptions_ARM
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

// Generator of MongoIndexOptions_ARM instances for property testing - lazily instantiated by
// MongoIndexOptions_ARMGenerator()
var mongoIndexOptions_ARMGenerator gopter.Gen

// MongoIndexOptions_ARMGenerator returns a generator of MongoIndexOptions_ARM instances for property testing.
func MongoIndexOptions_ARMGenerator() gopter.Gen {
	if mongoIndexOptions_ARMGenerator != nil {
		return mongoIndexOptions_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptions_ARM(generators)
	mongoIndexOptions_ARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptions_ARM{}), generators)

	return mongoIndexOptions_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptions_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptions_ARM(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}

func Test_MongoIndex_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndex_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndex_ARM, MongoIndex_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndex_ARM runs a test to see if a specific instance of MongoIndex_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndex_ARM(subject MongoIndex_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndex_ARM
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

// Generator of MongoIndex_ARM instances for property testing - lazily instantiated by MongoIndex_ARMGenerator()
var mongoIndex_ARMGenerator gopter.Gen

// MongoIndex_ARMGenerator returns a generator of MongoIndex_ARM instances for property testing.
func MongoIndex_ARMGenerator() gopter.Gen {
	if mongoIndex_ARMGenerator != nil {
		return mongoIndex_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndex_ARM(generators)
	mongoIndex_ARMGenerator = gen.Struct(reflect.TypeOf(MongoIndex_ARM{}), generators)

	return mongoIndex_ARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndex_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndex_ARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeys_ARMGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptions_ARMGenerator())
}
