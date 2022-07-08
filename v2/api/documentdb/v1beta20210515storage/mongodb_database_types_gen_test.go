// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

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

func Test_MongodbDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabase, MongodbDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabase runs a test to see if a specific instance of MongodbDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabase(subject MongodbDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabase
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

// Generator of MongodbDatabase instances for property testing - lazily instantiated by MongodbDatabaseGenerator()
var mongodbDatabaseGenerator gopter.Gen

// MongodbDatabaseGenerator returns a generator of MongodbDatabase instances for property testing.
func MongodbDatabaseGenerator() gopter.Gen {
	if mongodbDatabaseGenerator != nil {
		return mongodbDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongodbDatabase(generators)
	mongodbDatabaseGenerator = gen.Struct(reflect.TypeOf(MongodbDatabase{}), generators)

	return mongodbDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForMongodbDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsMongodbDatabasesSpecGenerator()
	gens["Status"] = MongoDBDatabaseGetResultsStatusGenerator()
}

func Test_DatabaseAccountsMongodbDatabases_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabases_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesSpec, DatabaseAccountsMongodbDatabasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesSpec runs a test to see if a specific instance of DatabaseAccountsMongodbDatabases_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesSpec(subject DatabaseAccountsMongodbDatabases_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabases_Spec
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

// Generator of DatabaseAccountsMongodbDatabases_Spec instances for property testing - lazily instantiated by
// DatabaseAccountsMongodbDatabasesSpecGenerator()
var databaseAccountsMongodbDatabasesSpecGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesSpecGenerator returns a generator of DatabaseAccountsMongodbDatabases_Spec instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesSpecGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesSpecGenerator != nil {
		return databaseAccountsMongodbDatabasesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpec(generators)
	databaseAccountsMongodbDatabasesSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabases_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpec(generators)
	databaseAccountsMongodbDatabasesSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabases_Spec{}), generators)

	return databaseAccountsMongodbDatabasesSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesSpec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseResourceGenerator())
}

func Test_MongoDBDatabaseGetResults_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseGetResults_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseGetResultsStatus, MongoDBDatabaseGetResultsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseGetResultsStatus runs a test to see if a specific instance of MongoDBDatabaseGetResults_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseGetResultsStatus(subject MongoDBDatabaseGetResults_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseGetResults_Status
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

// Generator of MongoDBDatabaseGetResults_Status instances for property testing - lazily instantiated by
// MongoDBDatabaseGetResultsStatusGenerator()
var mongoDBDatabaseGetResultsStatusGenerator gopter.Gen

// MongoDBDatabaseGetResultsStatusGenerator returns a generator of MongoDBDatabaseGetResults_Status instances for property testing.
// We first initialize mongoDBDatabaseGetResultsStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBDatabaseGetResultsStatusGenerator() gopter.Gen {
	if mongoDBDatabaseGetResultsStatusGenerator != nil {
		return mongoDBDatabaseGetResultsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseGetResultsStatus(generators)
	mongoDBDatabaseGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetResults_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseGetResultsStatus(generators)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseGetResultsStatus(generators)
	mongoDBDatabaseGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetResults_Status{}), generators)

	return mongoDBDatabaseGetResultsStatusGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseGetResultsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseGetResultsStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResourceStatusGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseGetPropertiesStatusResourceGenerator())
}

func Test_CreateUpdateOptions_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CreateUpdateOptions via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCreateUpdateOptions, CreateUpdateOptionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCreateUpdateOptions runs a test to see if a specific instance of CreateUpdateOptions round trips to JSON and back losslessly
func RunJSONSerializationTestForCreateUpdateOptions(subject CreateUpdateOptions) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CreateUpdateOptions
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

// Generator of CreateUpdateOptions instances for property testing - lazily instantiated by
// CreateUpdateOptionsGenerator()
var createUpdateOptionsGenerator gopter.Gen

// CreateUpdateOptionsGenerator returns a generator of CreateUpdateOptions instances for property testing.
// We first initialize createUpdateOptionsGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CreateUpdateOptionsGenerator() gopter.Gen {
	if createUpdateOptionsGenerator != nil {
		return createUpdateOptionsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions(generators)
	createUpdateOptionsGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions(generators)
	AddRelatedPropertyGeneratorsForCreateUpdateOptions(generators)
	createUpdateOptionsGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions{}), generators)

	return createUpdateOptionsGenerator
}

// AddIndependentPropertyGeneratorsForCreateUpdateOptions is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCreateUpdateOptions(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForCreateUpdateOptions is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCreateUpdateOptions(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsGenerator())
}

func Test_MongoDBDatabaseGetProperties_Status_Resource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseGetProperties_Status_Resource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseGetPropertiesStatusResource, MongoDBDatabaseGetPropertiesStatusResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseGetPropertiesStatusResource runs a test to see if a specific instance of MongoDBDatabaseGetProperties_Status_Resource round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseGetPropertiesStatusResource(subject MongoDBDatabaseGetProperties_Status_Resource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseGetProperties_Status_Resource
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

// Generator of MongoDBDatabaseGetProperties_Status_Resource instances for property testing - lazily instantiated by
// MongoDBDatabaseGetPropertiesStatusResourceGenerator()
var mongoDBDatabaseGetPropertiesStatusResourceGenerator gopter.Gen

// MongoDBDatabaseGetPropertiesStatusResourceGenerator returns a generator of MongoDBDatabaseGetProperties_Status_Resource instances for property testing.
func MongoDBDatabaseGetPropertiesStatusResourceGenerator() gopter.Gen {
	if mongoDBDatabaseGetPropertiesStatusResourceGenerator != nil {
		return mongoDBDatabaseGetPropertiesStatusResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseGetPropertiesStatusResource(generators)
	mongoDBDatabaseGetPropertiesStatusResourceGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetProperties_Status_Resource{}), generators)

	return mongoDBDatabaseGetPropertiesStatusResourceGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseGetPropertiesStatusResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseGetPropertiesStatusResource(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

func Test_MongoDBDatabaseResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseResource, MongoDBDatabaseResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseResource runs a test to see if a specific instance of MongoDBDatabaseResource round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseResource(subject MongoDBDatabaseResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseResource
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

// Generator of MongoDBDatabaseResource instances for property testing - lazily instantiated by
// MongoDBDatabaseResourceGenerator()
var mongoDBDatabaseResourceGenerator gopter.Gen

// MongoDBDatabaseResourceGenerator returns a generator of MongoDBDatabaseResource instances for property testing.
func MongoDBDatabaseResourceGenerator() gopter.Gen {
	if mongoDBDatabaseResourceGenerator != nil {
		return mongoDBDatabaseResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseResource(generators)
	mongoDBDatabaseResourceGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseResource{}), generators)

	return mongoDBDatabaseResourceGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseResource(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_OptionsResource_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OptionsResource_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOptionsResourceStatus, OptionsResourceStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOptionsResourceStatus runs a test to see if a specific instance of OptionsResource_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForOptionsResourceStatus(subject OptionsResource_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OptionsResource_Status
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

// Generator of OptionsResource_Status instances for property testing - lazily instantiated by
// OptionsResourceStatusGenerator()
var optionsResourceStatusGenerator gopter.Gen

// OptionsResourceStatusGenerator returns a generator of OptionsResource_Status instances for property testing.
// We first initialize optionsResourceStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func OptionsResourceStatusGenerator() gopter.Gen {
	if optionsResourceStatusGenerator != nil {
		return optionsResourceStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResourceStatus(generators)
	optionsResourceStatusGenerator = gen.Struct(reflect.TypeOf(OptionsResource_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResourceStatus(generators)
	AddRelatedPropertyGeneratorsForOptionsResourceStatus(generators)
	optionsResourceStatusGenerator = gen.Struct(reflect.TypeOf(OptionsResource_Status{}), generators)

	return optionsResourceStatusGenerator
}

// AddIndependentPropertyGeneratorsForOptionsResourceStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOptionsResourceStatus(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForOptionsResourceStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForOptionsResourceStatus(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsStatusGenerator())
}

func Test_AutoscaleSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettings, AutoscaleSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettings runs a test to see if a specific instance of AutoscaleSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettings(subject AutoscaleSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings
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

// Generator of AutoscaleSettings instances for property testing - lazily instantiated by AutoscaleSettingsGenerator()
var autoscaleSettingsGenerator gopter.Gen

// AutoscaleSettingsGenerator returns a generator of AutoscaleSettings instances for property testing.
func AutoscaleSettingsGenerator() gopter.Gen {
	if autoscaleSettingsGenerator != nil {
		return autoscaleSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettings(generators)
	autoscaleSettingsGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings{}), generators)

	return autoscaleSettingsGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettings(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

func Test_AutoscaleSettings_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsStatus, AutoscaleSettingsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsStatus runs a test to see if a specific instance of AutoscaleSettings_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsStatus(subject AutoscaleSettings_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings_Status
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

// Generator of AutoscaleSettings_Status instances for property testing - lazily instantiated by
// AutoscaleSettingsStatusGenerator()
var autoscaleSettingsStatusGenerator gopter.Gen

// AutoscaleSettingsStatusGenerator returns a generator of AutoscaleSettings_Status instances for property testing.
func AutoscaleSettingsStatusGenerator() gopter.Gen {
	if autoscaleSettingsStatusGenerator != nil {
		return autoscaleSettingsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsStatus(generators)
	autoscaleSettingsStatusGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings_Status{}), generators)

	return autoscaleSettingsStatusGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsStatus(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}
