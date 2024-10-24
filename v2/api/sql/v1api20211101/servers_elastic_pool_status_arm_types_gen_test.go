// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

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

func Test_ElasticPoolPerDatabaseSettings_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ElasticPoolPerDatabaseSettings_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForElasticPoolPerDatabaseSettings_STATUS_ARM, ElasticPoolPerDatabaseSettings_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForElasticPoolPerDatabaseSettings_STATUS_ARM runs a test to see if a specific instance of ElasticPoolPerDatabaseSettings_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForElasticPoolPerDatabaseSettings_STATUS_ARM(subject ElasticPoolPerDatabaseSettings_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ElasticPoolPerDatabaseSettings_STATUS_ARM
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

// Generator of ElasticPoolPerDatabaseSettings_STATUS_ARM instances for property testing - lazily instantiated by
// ElasticPoolPerDatabaseSettings_STATUS_ARMGenerator()
var elasticPoolPerDatabaseSettings_STATUS_ARMGenerator gopter.Gen

// ElasticPoolPerDatabaseSettings_STATUS_ARMGenerator returns a generator of ElasticPoolPerDatabaseSettings_STATUS_ARM instances for property testing.
func ElasticPoolPerDatabaseSettings_STATUS_ARMGenerator() gopter.Gen {
	if elasticPoolPerDatabaseSettings_STATUS_ARMGenerator != nil {
		return elasticPoolPerDatabaseSettings_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_STATUS_ARM(generators)
	elasticPoolPerDatabaseSettings_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ElasticPoolPerDatabaseSettings_STATUS_ARM{}), generators)

	return elasticPoolPerDatabaseSettings_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["MaxCapacity"] = gen.PtrOf(gen.Float64())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
}

func Test_ElasticPoolProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ElasticPoolProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForElasticPoolProperties_STATUS_ARM, ElasticPoolProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForElasticPoolProperties_STATUS_ARM runs a test to see if a specific instance of ElasticPoolProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForElasticPoolProperties_STATUS_ARM(subject ElasticPoolProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ElasticPoolProperties_STATUS_ARM
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

// Generator of ElasticPoolProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ElasticPoolProperties_STATUS_ARMGenerator()
var elasticPoolProperties_STATUS_ARMGenerator gopter.Gen

// ElasticPoolProperties_STATUS_ARMGenerator returns a generator of ElasticPoolProperties_STATUS_ARM instances for property testing.
// We first initialize elasticPoolProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ElasticPoolProperties_STATUS_ARMGenerator() gopter.Gen {
	if elasticPoolProperties_STATUS_ARMGenerator != nil {
		return elasticPoolProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolProperties_STATUS_ARM(generators)
	elasticPoolProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ElasticPoolProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForElasticPoolProperties_STATUS_ARM(generators)
	elasticPoolProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ElasticPoolProperties_STATUS_ARM{}), generators)

	return elasticPoolProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForElasticPoolProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForElasticPoolProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreationDate"] = gen.PtrOf(gen.AlphaString())
	gens["HighAvailabilityReplicaCount"] = gen.PtrOf(gen.Int())
	gens["LicenseType"] = gen.PtrOf(gen.OneConstOf(ElasticPoolProperties_LicenseType_STATUS_ARM_BasePrice, ElasticPoolProperties_LicenseType_STATUS_ARM_LicenseIncluded))
	gens["MaintenanceConfigurationId"] = gen.PtrOf(gen.AlphaString())
	gens["MaxSizeBytes"] = gen.PtrOf(gen.Int())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ElasticPoolProperties_State_STATUS_ARM_Creating, ElasticPoolProperties_State_STATUS_ARM_Disabled, ElasticPoolProperties_State_STATUS_ARM_Ready))
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForElasticPoolProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForElasticPoolProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PerDatabaseSettings"] = gen.PtrOf(ElasticPoolPerDatabaseSettings_STATUS_ARMGenerator())
}

func Test_ServersElasticPool_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersElasticPool_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersElasticPool_STATUS_ARM, ServersElasticPool_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersElasticPool_STATUS_ARM runs a test to see if a specific instance of ServersElasticPool_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServersElasticPool_STATUS_ARM(subject ServersElasticPool_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersElasticPool_STATUS_ARM
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

// Generator of ServersElasticPool_STATUS_ARM instances for property testing - lazily instantiated by
// ServersElasticPool_STATUS_ARMGenerator()
var serversElasticPool_STATUS_ARMGenerator gopter.Gen

// ServersElasticPool_STATUS_ARMGenerator returns a generator of ServersElasticPool_STATUS_ARM instances for property testing.
// We first initialize serversElasticPool_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersElasticPool_STATUS_ARMGenerator() gopter.Gen {
	if serversElasticPool_STATUS_ARMGenerator != nil {
		return serversElasticPool_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersElasticPool_STATUS_ARM(generators)
	serversElasticPool_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ServersElasticPool_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersElasticPool_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForServersElasticPool_STATUS_ARM(generators)
	serversElasticPool_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ServersElasticPool_STATUS_ARM{}), generators)

	return serversElasticPool_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServersElasticPool_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersElasticPool_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServersElasticPool_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersElasticPool_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ElasticPoolProperties_STATUS_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUS_ARMGenerator())
}
