// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101storage

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

func Test_ServersElasticPool_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersElasticPool via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersElasticPool, ServersElasticPoolGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersElasticPool runs a test to see if a specific instance of ServersElasticPool round trips to JSON and back losslessly
func RunJSONSerializationTestForServersElasticPool(subject ServersElasticPool) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersElasticPool
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

// Generator of ServersElasticPool instances for property testing - lazily instantiated by ServersElasticPoolGenerator()
var serversElasticPoolGenerator gopter.Gen

// ServersElasticPoolGenerator returns a generator of ServersElasticPool instances for property testing.
func ServersElasticPoolGenerator() gopter.Gen {
	if serversElasticPoolGenerator != nil {
		return serversElasticPoolGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersElasticPool(generators)
	serversElasticPoolGenerator = gen.Struct(reflect.TypeOf(ServersElasticPool{}), generators)

	return serversElasticPoolGenerator
}

// AddRelatedPropertyGeneratorsForServersElasticPool is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersElasticPool(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_ElasticPool_SpecGenerator()
	gens["Status"] = Servers_ElasticPool_STATUSGenerator()
}

func Test_Servers_ElasticPool_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_ElasticPool_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_ElasticPool_Spec, Servers_ElasticPool_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_ElasticPool_Spec runs a test to see if a specific instance of Servers_ElasticPool_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_ElasticPool_Spec(subject Servers_ElasticPool_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_ElasticPool_Spec
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

// Generator of Servers_ElasticPool_Spec instances for property testing - lazily instantiated by
// Servers_ElasticPool_SpecGenerator()
var servers_ElasticPool_SpecGenerator gopter.Gen

// Servers_ElasticPool_SpecGenerator returns a generator of Servers_ElasticPool_Spec instances for property testing.
// We first initialize servers_ElasticPool_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_ElasticPool_SpecGenerator() gopter.Gen {
	if servers_ElasticPool_SpecGenerator != nil {
		return servers_ElasticPool_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec(generators)
	servers_ElasticPool_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_ElasticPool_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec(generators)
	AddRelatedPropertyGeneratorsForServers_ElasticPool_Spec(generators)
	servers_ElasticPool_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_ElasticPool_Spec{}), generators)

	return servers_ElasticPool_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["HighAvailabilityReplicaCount"] = gen.PtrOf(gen.Int())
	gens["LicenseType"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MaintenanceConfigurationId"] = gen.PtrOf(gen.AlphaString())
	gens["MaxSizeBytes"] = gen.PtrOf(gen.Int())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForServers_ElasticPool_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_ElasticPool_Spec(gens map[string]gopter.Gen) {
	gens["PerDatabaseSettings"] = gen.PtrOf(ElasticPoolPerDatabaseSettingsGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_Servers_ElasticPool_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_ElasticPool_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_ElasticPool_STATUS, Servers_ElasticPool_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_ElasticPool_STATUS runs a test to see if a specific instance of Servers_ElasticPool_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_ElasticPool_STATUS(subject Servers_ElasticPool_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_ElasticPool_STATUS
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

// Generator of Servers_ElasticPool_STATUS instances for property testing - lazily instantiated by
// Servers_ElasticPool_STATUSGenerator()
var servers_ElasticPool_STATUSGenerator gopter.Gen

// Servers_ElasticPool_STATUSGenerator returns a generator of Servers_ElasticPool_STATUS instances for property testing.
// We first initialize servers_ElasticPool_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_ElasticPool_STATUSGenerator() gopter.Gen {
	if servers_ElasticPool_STATUSGenerator != nil {
		return servers_ElasticPool_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ElasticPool_STATUS(generators)
	servers_ElasticPool_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_ElasticPool_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ElasticPool_STATUS(generators)
	AddRelatedPropertyGeneratorsForServers_ElasticPool_STATUS(generators)
	servers_ElasticPool_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_ElasticPool_STATUS{}), generators)

	return servers_ElasticPool_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_ElasticPool_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_ElasticPool_STATUS(gens map[string]gopter.Gen) {
	gens["CreationDate"] = gen.PtrOf(gen.AlphaString())
	gens["HighAvailabilityReplicaCount"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["LicenseType"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MaintenanceConfigurationId"] = gen.PtrOf(gen.AlphaString())
	gens["MaxSizeBytes"] = gen.PtrOf(gen.Int())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForServers_ElasticPool_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_ElasticPool_STATUS(gens map[string]gopter.Gen) {
	gens["PerDatabaseSettings"] = gen.PtrOf(ElasticPoolPerDatabaseSettings_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
}

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

func Test_ElasticPoolPerDatabaseSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ElasticPoolPerDatabaseSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForElasticPoolPerDatabaseSettings_STATUS, ElasticPoolPerDatabaseSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForElasticPoolPerDatabaseSettings_STATUS runs a test to see if a specific instance of ElasticPoolPerDatabaseSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForElasticPoolPerDatabaseSettings_STATUS(subject ElasticPoolPerDatabaseSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ElasticPoolPerDatabaseSettings_STATUS
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

// Generator of ElasticPoolPerDatabaseSettings_STATUS instances for property testing - lazily instantiated by
// ElasticPoolPerDatabaseSettings_STATUSGenerator()
var elasticPoolPerDatabaseSettings_STATUSGenerator gopter.Gen

// ElasticPoolPerDatabaseSettings_STATUSGenerator returns a generator of ElasticPoolPerDatabaseSettings_STATUS instances for property testing.
func ElasticPoolPerDatabaseSettings_STATUSGenerator() gopter.Gen {
	if elasticPoolPerDatabaseSettings_STATUSGenerator != nil {
		return elasticPoolPerDatabaseSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_STATUS(generators)
	elasticPoolPerDatabaseSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(ElasticPoolPerDatabaseSettings_STATUS{}), generators)

	return elasticPoolPerDatabaseSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_STATUS(gens map[string]gopter.Gen) {
	gens["MaxCapacity"] = gen.PtrOf(gen.Float64())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
}
