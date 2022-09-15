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

func Test_MongoDBDatabaseGetResults_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseGetResults_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseGetResults_STATUS_ARM, MongoDBDatabaseGetResults_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseGetResults_STATUS_ARM runs a test to see if a specific instance of MongoDBDatabaseGetResults_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseGetResults_STATUS_ARM(subject MongoDBDatabaseGetResults_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseGetResults_STATUS_ARM
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

// Generator of MongoDBDatabaseGetResults_STATUS_ARM instances for property testing - lazily instantiated by
// MongoDBDatabaseGetResults_STATUS_ARMGenerator()
var mongoDBDatabaseGetResults_STATUS_ARMGenerator gopter.Gen

// MongoDBDatabaseGetResults_STATUS_ARMGenerator returns a generator of MongoDBDatabaseGetResults_STATUS_ARM instances for property testing.
// We first initialize mongoDBDatabaseGetResults_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBDatabaseGetResults_STATUS_ARMGenerator() gopter.Gen {
	if mongoDBDatabaseGetResults_STATUS_ARMGenerator != nil {
		return mongoDBDatabaseGetResults_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseGetResults_STATUS_ARM(generators)
	mongoDBDatabaseGetResults_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetResults_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseGetResults_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseGetResults_STATUS_ARM(generators)
	mongoDBDatabaseGetResults_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetResults_STATUS_ARM{}), generators)

	return mongoDBDatabaseGetResults_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseGetResults_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseGetResults_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseGetResults_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseGetResults_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBDatabaseGetProperties_STATUS_ARMGenerator())
}

func Test_MongoDBDatabaseGetProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseGetProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseGetProperties_STATUS_ARM, MongoDBDatabaseGetProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseGetProperties_STATUS_ARM runs a test to see if a specific instance of MongoDBDatabaseGetProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseGetProperties_STATUS_ARM(subject MongoDBDatabaseGetProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseGetProperties_STATUS_ARM
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

// Generator of MongoDBDatabaseGetProperties_STATUS_ARM instances for property testing - lazily instantiated by
// MongoDBDatabaseGetProperties_STATUS_ARMGenerator()
var mongoDBDatabaseGetProperties_STATUS_ARMGenerator gopter.Gen

// MongoDBDatabaseGetProperties_STATUS_ARMGenerator returns a generator of MongoDBDatabaseGetProperties_STATUS_ARM instances for property testing.
func MongoDBDatabaseGetProperties_STATUS_ARMGenerator() gopter.Gen {
	if mongoDBDatabaseGetProperties_STATUS_ARMGenerator != nil {
		return mongoDBDatabaseGetProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_STATUS_ARM(generators)
	mongoDBDatabaseGetProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetProperties_STATUS_ARM{}), generators)

	return mongoDBDatabaseGetProperties_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResource_STATUS_ARMGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseGetProperties_Resource_STATUS_ARMGenerator())
}

func Test_MongoDBDatabaseGetProperties_Resource_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseGetProperties_Resource_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUS_ARM, MongoDBDatabaseGetProperties_Resource_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUS_ARM runs a test to see if a specific instance of MongoDBDatabaseGetProperties_Resource_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUS_ARM(subject MongoDBDatabaseGetProperties_Resource_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseGetProperties_Resource_STATUS_ARM
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

// Generator of MongoDBDatabaseGetProperties_Resource_STATUS_ARM instances for property testing - lazily instantiated by
// MongoDBDatabaseGetProperties_Resource_STATUS_ARMGenerator()
var mongoDBDatabaseGetProperties_Resource_STATUS_ARMGenerator gopter.Gen

// MongoDBDatabaseGetProperties_Resource_STATUS_ARMGenerator returns a generator of MongoDBDatabaseGetProperties_Resource_STATUS_ARM instances for property testing.
func MongoDBDatabaseGetProperties_Resource_STATUS_ARMGenerator() gopter.Gen {
	if mongoDBDatabaseGetProperties_Resource_STATUS_ARMGenerator != nil {
		return mongoDBDatabaseGetProperties_Resource_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS_ARM(generators)
	mongoDBDatabaseGetProperties_Resource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetProperties_Resource_STATUS_ARM{}), generators)

	return mongoDBDatabaseGetProperties_Resource_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}
