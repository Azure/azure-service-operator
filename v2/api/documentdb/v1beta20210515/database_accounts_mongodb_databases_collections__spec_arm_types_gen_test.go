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

func Test_DatabaseAccountsMongodbDatabasesCollections_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesCollections_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsSpecARM, DatabaseAccountsMongodbDatabasesCollectionsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsSpecARM runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesCollections_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsSpecARM(subject DatabaseAccountsMongodbDatabasesCollections_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesCollections_SpecARM
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

// Generator of DatabaseAccountsMongodbDatabasesCollections_SpecARM instances for property testing - lazily instantiated
// by DatabaseAccountsMongodbDatabasesCollectionsSpecARMGenerator()
var databaseAccountsMongodbDatabasesCollectionsSpecARMGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesCollectionsSpecARMGenerator returns a generator of DatabaseAccountsMongodbDatabasesCollections_SpecARM instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesCollectionsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesCollectionsSpecARMGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesCollectionsSpecARMGenerator != nil {
		return databaseAccountsMongodbDatabasesCollectionsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpecARM(generators)
	databaseAccountsMongodbDatabasesCollectionsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollections_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpecARM(generators)
	databaseAccountsMongodbDatabasesCollectionsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollections_SpecARM{}), generators)

	return databaseAccountsMongodbDatabasesCollectionsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBCollectionCreateUpdatePropertiesARMGenerator())
}

func Test_MongoDBCollectionCreateUpdatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionCreateUpdatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionCreateUpdatePropertiesARM, MongoDBCollectionCreateUpdatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionCreateUpdatePropertiesARM runs a test to see if a specific instance of MongoDBCollectionCreateUpdatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionCreateUpdatePropertiesARM(subject MongoDBCollectionCreateUpdatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionCreateUpdatePropertiesARM
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

// Generator of MongoDBCollectionCreateUpdatePropertiesARM instances for property testing - lazily instantiated by
// MongoDBCollectionCreateUpdatePropertiesARMGenerator()
var mongoDBCollectionCreateUpdatePropertiesARMGenerator gopter.Gen

// MongoDBCollectionCreateUpdatePropertiesARMGenerator returns a generator of MongoDBCollectionCreateUpdatePropertiesARM instances for property testing.
func MongoDBCollectionCreateUpdatePropertiesARMGenerator() gopter.Gen {
	if mongoDBCollectionCreateUpdatePropertiesARMGenerator != nil {
		return mongoDBCollectionCreateUpdatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBCollectionCreateUpdatePropertiesARM(generators)
	mongoDBCollectionCreateUpdatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionCreateUpdatePropertiesARM{}), generators)

	return mongoDBCollectionCreateUpdatePropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionCreateUpdatePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionCreateUpdatePropertiesARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsARMGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBCollectionResourceARMGenerator())
}

func Test_CreateUpdateOptionsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CreateUpdateOptionsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCreateUpdateOptionsARM, CreateUpdateOptionsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCreateUpdateOptionsARM runs a test to see if a specific instance of CreateUpdateOptionsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCreateUpdateOptionsARM(subject CreateUpdateOptionsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CreateUpdateOptionsARM
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

// Generator of CreateUpdateOptionsARM instances for property testing - lazily instantiated by
// CreateUpdateOptionsARMGenerator()
var createUpdateOptionsARMGenerator gopter.Gen

// CreateUpdateOptionsARMGenerator returns a generator of CreateUpdateOptionsARM instances for property testing.
// We first initialize createUpdateOptionsARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CreateUpdateOptionsARMGenerator() gopter.Gen {
	if createUpdateOptionsARMGenerator != nil {
		return createUpdateOptionsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptionsARM(generators)
	createUpdateOptionsARMGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptionsARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptionsARM(generators)
	AddRelatedPropertyGeneratorsForCreateUpdateOptionsARM(generators)
	createUpdateOptionsARMGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptionsARM{}), generators)

	return createUpdateOptionsARMGenerator
}

// AddIndependentPropertyGeneratorsForCreateUpdateOptionsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCreateUpdateOptionsARM(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForCreateUpdateOptionsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCreateUpdateOptionsARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsARMGenerator())
}

func Test_MongoDBCollectionResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionResourceARM, MongoDBCollectionResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionResourceARM runs a test to see if a specific instance of MongoDBCollectionResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionResourceARM(subject MongoDBCollectionResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionResourceARM
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

// Generator of MongoDBCollectionResourceARM instances for property testing - lazily instantiated by
// MongoDBCollectionResourceARMGenerator()
var mongoDBCollectionResourceARMGenerator gopter.Gen

// MongoDBCollectionResourceARMGenerator returns a generator of MongoDBCollectionResourceARM instances for property testing.
// We first initialize mongoDBCollectionResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionResourceARMGenerator() gopter.Gen {
	if mongoDBCollectionResourceARMGenerator != nil {
		return mongoDBCollectionResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResourceARM(generators)
	mongoDBCollectionResourceARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResourceARM(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionResourceARM(generators)
	mongoDBCollectionResourceARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResourceARM{}), generators)

	return mongoDBCollectionResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionResourceARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["ShardKey"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionResourceARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndexARMGenerator())
}

func Test_AutoscaleSettingsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettingsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsARM, AutoscaleSettingsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsARM runs a test to see if a specific instance of AutoscaleSettingsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsARM(subject AutoscaleSettingsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettingsARM
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

// Generator of AutoscaleSettingsARM instances for property testing - lazily instantiated by
// AutoscaleSettingsARMGenerator()
var autoscaleSettingsARMGenerator gopter.Gen

// AutoscaleSettingsARMGenerator returns a generator of AutoscaleSettingsARM instances for property testing.
func AutoscaleSettingsARMGenerator() gopter.Gen {
	if autoscaleSettingsARMGenerator != nil {
		return autoscaleSettingsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsARM(generators)
	autoscaleSettingsARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsARM{}), generators)

	return autoscaleSettingsARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

func Test_MongoIndexARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexARM, MongoIndexARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexARM runs a test to see if a specific instance of MongoIndexARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexARM(subject MongoIndexARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexARM
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

// Generator of MongoIndexARM instances for property testing - lazily instantiated by MongoIndexARMGenerator()
var mongoIndexARMGenerator gopter.Gen

// MongoIndexARMGenerator returns a generator of MongoIndexARM instances for property testing.
func MongoIndexARMGenerator() gopter.Gen {
	if mongoIndexARMGenerator != nil {
		return mongoIndexARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndexARM(generators)
	mongoIndexARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexARM{}), generators)

	return mongoIndexARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndexARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndexARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeysARMGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptionsARMGenerator())
}

func Test_MongoIndexKeysARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeysARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeysARM, MongoIndexKeysARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeysARM runs a test to see if a specific instance of MongoIndexKeysARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeysARM(subject MongoIndexKeysARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeysARM
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

// Generator of MongoIndexKeysARM instances for property testing - lazily instantiated by MongoIndexKeysARMGenerator()
var mongoIndexKeysARMGenerator gopter.Gen

// MongoIndexKeysARMGenerator returns a generator of MongoIndexKeysARM instances for property testing.
func MongoIndexKeysARMGenerator() gopter.Gen {
	if mongoIndexKeysARMGenerator != nil {
		return mongoIndexKeysARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeysARM(generators)
	mongoIndexKeysARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeysARM{}), generators)

	return mongoIndexKeysARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeysARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeysARM(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexOptionsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptionsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptionsARM, MongoIndexOptionsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptionsARM runs a test to see if a specific instance of MongoIndexOptionsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptionsARM(subject MongoIndexOptionsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptionsARM
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

// Generator of MongoIndexOptionsARM instances for property testing - lazily instantiated by
// MongoIndexOptionsARMGenerator()
var mongoIndexOptionsARMGenerator gopter.Gen

// MongoIndexOptionsARMGenerator returns a generator of MongoIndexOptionsARM instances for property testing.
func MongoIndexOptionsARMGenerator() gopter.Gen {
	if mongoIndexOptionsARMGenerator != nil {
		return mongoIndexOptionsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptionsARM(generators)
	mongoIndexOptionsARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptionsARM{}), generators)

	return mongoIndexOptionsARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptionsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptionsARM(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}
