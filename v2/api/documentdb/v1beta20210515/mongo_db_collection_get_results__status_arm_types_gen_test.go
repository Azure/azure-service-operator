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

func Test_MongoDBCollectionGetResults_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionGetResults_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionGetResultsStatusARM, MongoDBCollectionGetResultsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionGetResultsStatusARM runs a test to see if a specific instance of MongoDBCollectionGetResults_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionGetResultsStatusARM(subject MongoDBCollectionGetResults_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionGetResults_StatusARM
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

// Generator of MongoDBCollectionGetResults_StatusARM instances for property testing - lazily instantiated by
// MongoDBCollectionGetResultsStatusARMGenerator()
var mongoDBCollectionGetResultsStatusARMGenerator gopter.Gen

// MongoDBCollectionGetResultsStatusARMGenerator returns a generator of MongoDBCollectionGetResults_StatusARM instances for property testing.
// We first initialize mongoDBCollectionGetResultsStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionGetResultsStatusARMGenerator() gopter.Gen {
	if mongoDBCollectionGetResultsStatusARMGenerator != nil {
		return mongoDBCollectionGetResultsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsStatusARM(generators)
	mongoDBCollectionGetResultsStatusARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetResults_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsStatusARM(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionGetResultsStatusARM(generators)
	mongoDBCollectionGetResultsStatusARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetResults_StatusARM{}), generators)

	return mongoDBCollectionGetResultsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionGetResultsStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionGetResultsStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBCollectionGetPropertiesStatusARMGenerator())
}

func Test_MongoDBCollectionGetProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionGetProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionGetPropertiesStatusARM, MongoDBCollectionGetPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionGetPropertiesStatusARM runs a test to see if a specific instance of MongoDBCollectionGetProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionGetPropertiesStatusARM(subject MongoDBCollectionGetProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionGetProperties_StatusARM
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

// Generator of MongoDBCollectionGetProperties_StatusARM instances for property testing - lazily instantiated by
// MongoDBCollectionGetPropertiesStatusARMGenerator()
var mongoDBCollectionGetPropertiesStatusARMGenerator gopter.Gen

// MongoDBCollectionGetPropertiesStatusARMGenerator returns a generator of MongoDBCollectionGetProperties_StatusARM instances for property testing.
func MongoDBCollectionGetPropertiesStatusARMGenerator() gopter.Gen {
	if mongoDBCollectionGetPropertiesStatusARMGenerator != nil {
		return mongoDBCollectionGetPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusARM(generators)
	mongoDBCollectionGetPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetProperties_StatusARM{}), generators)

	return mongoDBCollectionGetPropertiesStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResourceStatusARMGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBCollectionGetPropertiesStatusResourceARMGenerator())
}

func Test_MongoDBCollectionGetProperties_Status_ResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionGetProperties_Status_ResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionGetPropertiesStatusResourceARM, MongoDBCollectionGetPropertiesStatusResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionGetPropertiesStatusResourceARM runs a test to see if a specific instance of MongoDBCollectionGetProperties_Status_ResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionGetPropertiesStatusResourceARM(subject MongoDBCollectionGetProperties_Status_ResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionGetProperties_Status_ResourceARM
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

// Generator of MongoDBCollectionGetProperties_Status_ResourceARM instances for property testing - lazily instantiated
// by MongoDBCollectionGetPropertiesStatusResourceARMGenerator()
var mongoDBCollectionGetPropertiesStatusResourceARMGenerator gopter.Gen

// MongoDBCollectionGetPropertiesStatusResourceARMGenerator returns a generator of MongoDBCollectionGetProperties_Status_ResourceARM instances for property testing.
// We first initialize mongoDBCollectionGetPropertiesStatusResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionGetPropertiesStatusResourceARMGenerator() gopter.Gen {
	if mongoDBCollectionGetPropertiesStatusResourceARMGenerator != nil {
		return mongoDBCollectionGetPropertiesStatusResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResourceARM(generators)
	mongoDBCollectionGetPropertiesStatusResourceARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetProperties_Status_ResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResourceARM(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResourceARM(generators)
	mongoDBCollectionGetPropertiesStatusResourceARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetProperties_Status_ResourceARM{}), generators)

	return mongoDBCollectionGetPropertiesStatusResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResourceARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["ShardKey"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResourceARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndexStatusARMGenerator())
}

func Test_OptionsResource_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OptionsResource_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOptionsResourceStatusARM, OptionsResourceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOptionsResourceStatusARM runs a test to see if a specific instance of OptionsResource_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForOptionsResourceStatusARM(subject OptionsResource_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OptionsResource_StatusARM
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

// Generator of OptionsResource_StatusARM instances for property testing - lazily instantiated by
// OptionsResourceStatusARMGenerator()
var optionsResourceStatusARMGenerator gopter.Gen

// OptionsResourceStatusARMGenerator returns a generator of OptionsResource_StatusARM instances for property testing.
// We first initialize optionsResourceStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func OptionsResourceStatusARMGenerator() gopter.Gen {
	if optionsResourceStatusARMGenerator != nil {
		return optionsResourceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResourceStatusARM(generators)
	optionsResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(OptionsResource_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResourceStatusARM(generators)
	AddRelatedPropertyGeneratorsForOptionsResourceStatusARM(generators)
	optionsResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(OptionsResource_StatusARM{}), generators)

	return optionsResourceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForOptionsResourceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOptionsResourceStatusARM(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForOptionsResourceStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForOptionsResourceStatusARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsStatusARMGenerator())
}

func Test_AutoscaleSettings_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsStatusARM, AutoscaleSettingsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsStatusARM runs a test to see if a specific instance of AutoscaleSettings_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsStatusARM(subject AutoscaleSettings_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings_StatusARM
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

// Generator of AutoscaleSettings_StatusARM instances for property testing - lazily instantiated by
// AutoscaleSettingsStatusARMGenerator()
var autoscaleSettingsStatusARMGenerator gopter.Gen

// AutoscaleSettingsStatusARMGenerator returns a generator of AutoscaleSettings_StatusARM instances for property testing.
func AutoscaleSettingsStatusARMGenerator() gopter.Gen {
	if autoscaleSettingsStatusARMGenerator != nil {
		return autoscaleSettingsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsStatusARM(generators)
	autoscaleSettingsStatusARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings_StatusARM{}), generators)

	return autoscaleSettingsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsStatusARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

func Test_MongoIndex_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndex_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexStatusARM, MongoIndexStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexStatusARM runs a test to see if a specific instance of MongoIndex_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexStatusARM(subject MongoIndex_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndex_StatusARM
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

// Generator of MongoIndex_StatusARM instances for property testing - lazily instantiated by
// MongoIndexStatusARMGenerator()
var mongoIndexStatusARMGenerator gopter.Gen

// MongoIndexStatusARMGenerator returns a generator of MongoIndex_StatusARM instances for property testing.
func MongoIndexStatusARMGenerator() gopter.Gen {
	if mongoIndexStatusARMGenerator != nil {
		return mongoIndexStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndexStatusARM(generators)
	mongoIndexStatusARMGenerator = gen.Struct(reflect.TypeOf(MongoIndex_StatusARM{}), generators)

	return mongoIndexStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndexStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndexStatusARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeysStatusARMGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptionsStatusARMGenerator())
}

func Test_MongoIndexKeys_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeys_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeysStatusARM, MongoIndexKeysStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeysStatusARM runs a test to see if a specific instance of MongoIndexKeys_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeysStatusARM(subject MongoIndexKeys_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeys_StatusARM
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

// Generator of MongoIndexKeys_StatusARM instances for property testing - lazily instantiated by
// MongoIndexKeysStatusARMGenerator()
var mongoIndexKeysStatusARMGenerator gopter.Gen

// MongoIndexKeysStatusARMGenerator returns a generator of MongoIndexKeys_StatusARM instances for property testing.
func MongoIndexKeysStatusARMGenerator() gopter.Gen {
	if mongoIndexKeysStatusARMGenerator != nil {
		return mongoIndexKeysStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeysStatusARM(generators)
	mongoIndexKeysStatusARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeys_StatusARM{}), generators)

	return mongoIndexKeysStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeysStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeysStatusARM(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexOptions_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptions_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptionsStatusARM, MongoIndexOptionsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptionsStatusARM runs a test to see if a specific instance of MongoIndexOptions_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptionsStatusARM(subject MongoIndexOptions_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptions_StatusARM
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

// Generator of MongoIndexOptions_StatusARM instances for property testing - lazily instantiated by
// MongoIndexOptionsStatusARMGenerator()
var mongoIndexOptionsStatusARMGenerator gopter.Gen

// MongoIndexOptionsStatusARMGenerator returns a generator of MongoIndexOptions_StatusARM instances for property testing.
func MongoIndexOptionsStatusARMGenerator() gopter.Gen {
	if mongoIndexOptionsStatusARMGenerator != nil {
		return mongoIndexOptionsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptionsStatusARM(generators)
	mongoIndexOptionsStatusARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptions_StatusARM{}), generators)

	return mongoIndexOptionsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptionsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptionsStatusARM(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}
