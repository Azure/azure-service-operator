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

func Test_MongoDBDatabaseCreateUpdateProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseCreateUpdateProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseCreateUpdateProperties, MongoDBDatabaseCreateUpdatePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseCreateUpdateProperties runs a test to see if a specific instance of MongoDBDatabaseCreateUpdateProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseCreateUpdateProperties(subject MongoDBDatabaseCreateUpdateProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseCreateUpdateProperties
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

// Generator of MongoDBDatabaseCreateUpdateProperties instances for property testing - lazily instantiated by
// MongoDBDatabaseCreateUpdatePropertiesGenerator()
var mongoDBDatabaseCreateUpdatePropertiesGenerator gopter.Gen

// MongoDBDatabaseCreateUpdatePropertiesGenerator returns a generator of MongoDBDatabaseCreateUpdateProperties instances for property testing.
func MongoDBDatabaseCreateUpdatePropertiesGenerator() gopter.Gen {
	if mongoDBDatabaseCreateUpdatePropertiesGenerator != nil {
		return mongoDBDatabaseCreateUpdatePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdateProperties(generators)
	mongoDBDatabaseCreateUpdatePropertiesGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseCreateUpdateProperties{}), generators)

	return mongoDBDatabaseCreateUpdatePropertiesGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdateProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdateProperties(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseResourceGenerator())
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

func Test_MongodbDatabase_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabase_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabase_Spec, MongodbDatabase_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabase_Spec runs a test to see if a specific instance of MongodbDatabase_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabase_Spec(subject MongodbDatabase_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabase_Spec
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

// Generator of MongodbDatabase_Spec instances for property testing - lazily instantiated by
// MongodbDatabase_SpecGenerator()
var mongodbDatabase_SpecGenerator gopter.Gen

// MongodbDatabase_SpecGenerator returns a generator of MongodbDatabase_Spec instances for property testing.
// We first initialize mongodbDatabase_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongodbDatabase_SpecGenerator() gopter.Gen {
	if mongodbDatabase_SpecGenerator != nil {
		return mongodbDatabase_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabase_Spec(generators)
	mongodbDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(MongodbDatabase_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabase_Spec(generators)
	AddRelatedPropertyGeneratorsForMongodbDatabase_Spec(generators)
	mongodbDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(MongodbDatabase_Spec{}), generators)

	return mongodbDatabase_SpecGenerator
}

// AddIndependentPropertyGeneratorsForMongodbDatabase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongodbDatabase_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongodbDatabase_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabase_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBDatabaseCreateUpdatePropertiesGenerator())
}
