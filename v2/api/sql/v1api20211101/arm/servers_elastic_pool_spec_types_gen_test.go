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

func Test_ElasticPoolPerDatabaseSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ElasticPoolPerDatabaseSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForElasticPoolPerDatabaseSettings, ElasticPoolPerDatabaseSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForElasticPoolPerDatabaseSettings runs a test to see if a specific instance of ElasticPoolPerDatabaseSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForElasticPoolPerDatabaseSettings(subject ElasticPoolPerDatabaseSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ElasticPoolPerDatabaseSettings
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

// Generator of ElasticPoolPerDatabaseSettings instances for property testing - lazily instantiated by
// ElasticPoolPerDatabaseSettingsGenerator()
var elasticPoolPerDatabaseSettingsGenerator gopter.Gen

// ElasticPoolPerDatabaseSettingsGenerator returns a generator of ElasticPoolPerDatabaseSettings instances for property testing.
func ElasticPoolPerDatabaseSettingsGenerator() gopter.Gen {
	if elasticPoolPerDatabaseSettingsGenerator != nil {
		return elasticPoolPerDatabaseSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings(generators)
	elasticPoolPerDatabaseSettingsGenerator = gen.Struct(reflect.TypeOf(ElasticPoolPerDatabaseSettings{}), generators)

	return elasticPoolPerDatabaseSettingsGenerator
}

// AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings(gens map[string]gopter.Gen) {
	gens["MaxCapacity"] = gen.PtrOf(gen.Float64())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
}

func Test_ElasticPoolProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ElasticPoolProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForElasticPoolProperties, ElasticPoolPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForElasticPoolProperties runs a test to see if a specific instance of ElasticPoolProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForElasticPoolProperties(subject ElasticPoolProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ElasticPoolProperties
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

// Generator of ElasticPoolProperties instances for property testing - lazily instantiated by
// ElasticPoolPropertiesGenerator()
var elasticPoolPropertiesGenerator gopter.Gen

// ElasticPoolPropertiesGenerator returns a generator of ElasticPoolProperties instances for property testing.
// We first initialize elasticPoolPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ElasticPoolPropertiesGenerator() gopter.Gen {
	if elasticPoolPropertiesGenerator != nil {
		return elasticPoolPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolProperties(generators)
	elasticPoolPropertiesGenerator = gen.Struct(reflect.TypeOf(ElasticPoolProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolProperties(generators)
	AddRelatedPropertyGeneratorsForElasticPoolProperties(generators)
	elasticPoolPropertiesGenerator = gen.Struct(reflect.TypeOf(ElasticPoolProperties{}), generators)

	return elasticPoolPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForElasticPoolProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForElasticPoolProperties(gens map[string]gopter.Gen) {
	gens["HighAvailabilityReplicaCount"] = gen.PtrOf(gen.Int())
	gens["LicenseType"] = gen.PtrOf(gen.OneConstOf(ElasticPoolProperties_LicenseType_BasePrice, ElasticPoolProperties_LicenseType_LicenseIncluded))
	gens["MaintenanceConfigurationId"] = gen.PtrOf(gen.AlphaString())
	gens["MaxSizeBytes"] = gen.PtrOf(gen.Int())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForElasticPoolProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForElasticPoolProperties(gens map[string]gopter.Gen) {
	gens["PerDatabaseSettings"] = gen.PtrOf(ElasticPoolPerDatabaseSettingsGenerator())
}

func Test_ServersElasticPool_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersElasticPool_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersElasticPool_Spec, ServersElasticPool_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersElasticPool_Spec runs a test to see if a specific instance of ServersElasticPool_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersElasticPool_Spec(subject ServersElasticPool_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersElasticPool_Spec
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

// Generator of ServersElasticPool_Spec instances for property testing - lazily instantiated by
// ServersElasticPool_SpecGenerator()
var serversElasticPool_SpecGenerator gopter.Gen

// ServersElasticPool_SpecGenerator returns a generator of ServersElasticPool_Spec instances for property testing.
// We first initialize serversElasticPool_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersElasticPool_SpecGenerator() gopter.Gen {
	if serversElasticPool_SpecGenerator != nil {
		return serversElasticPool_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersElasticPool_Spec(generators)
	serversElasticPool_SpecGenerator = gen.Struct(reflect.TypeOf(ServersElasticPool_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersElasticPool_Spec(generators)
	AddRelatedPropertyGeneratorsForServersElasticPool_Spec(generators)
	serversElasticPool_SpecGenerator = gen.Struct(reflect.TypeOf(ServersElasticPool_Spec{}), generators)

	return serversElasticPool_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersElasticPool_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersElasticPool_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServersElasticPool_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersElasticPool_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ElasticPoolPropertiesGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}
