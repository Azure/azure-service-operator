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

func Test_DatabaseProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseProperties, DatabasePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseProperties runs a test to see if a specific instance of DatabaseProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseProperties(subject DatabaseProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties
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

// Generator of DatabaseProperties instances for property testing - lazily instantiated by DatabasePropertiesGenerator()
var databasePropertiesGenerator gopter.Gen

// DatabasePropertiesGenerator returns a generator of DatabaseProperties instances for property testing.
// We first initialize databasePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabasePropertiesGenerator() gopter.Gen {
	if databasePropertiesGenerator != nil {
		return databasePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties(generators)
	databasePropertiesGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties(generators)
	AddRelatedPropertyGeneratorsForDatabaseProperties(generators)
	databasePropertiesGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties{}), generators)

	return databasePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseProperties(gens map[string]gopter.Gen) {
	gens["ClientProtocol"] = gen.PtrOf(gen.OneConstOf(DatabaseProperties_ClientProtocol_Encrypted, DatabaseProperties_ClientProtocol_Plaintext))
	gens["ClusteringPolicy"] = gen.PtrOf(gen.OneConstOf(DatabaseProperties_ClusteringPolicy_EnterpriseCluster, DatabaseProperties_ClusteringPolicy_OSSCluster))
	gens["EvictionPolicy"] = gen.PtrOf(gen.OneConstOf(
		DatabaseProperties_EvictionPolicy_AllKeysLFU,
		DatabaseProperties_EvictionPolicy_AllKeysLRU,
		DatabaseProperties_EvictionPolicy_AllKeysRandom,
		DatabaseProperties_EvictionPolicy_NoEviction,
		DatabaseProperties_EvictionPolicy_VolatileLFU,
		DatabaseProperties_EvictionPolicy_VolatileLRU,
		DatabaseProperties_EvictionPolicy_VolatileRandom,
		DatabaseProperties_EvictionPolicy_VolatileTTL))
	gens["Port"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDatabaseProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseProperties(gens map[string]gopter.Gen) {
	gens["GeoReplication"] = gen.PtrOf(DatabaseProperties_GeoReplicationGenerator())
	gens["Modules"] = gen.SliceOf(ModuleGenerator())
	gens["Persistence"] = gen.PtrOf(PersistenceGenerator())
}

func Test_DatabaseProperties_GeoReplication_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_GeoReplication via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseProperties_GeoReplication, DatabaseProperties_GeoReplicationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseProperties_GeoReplication runs a test to see if a specific instance of DatabaseProperties_GeoReplication round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseProperties_GeoReplication(subject DatabaseProperties_GeoReplication) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_GeoReplication
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

// Generator of DatabaseProperties_GeoReplication instances for property testing - lazily instantiated by
// DatabaseProperties_GeoReplicationGenerator()
var databaseProperties_GeoReplicationGenerator gopter.Gen

// DatabaseProperties_GeoReplicationGenerator returns a generator of DatabaseProperties_GeoReplication instances for property testing.
// We first initialize databaseProperties_GeoReplicationGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseProperties_GeoReplicationGenerator() gopter.Gen {
	if databaseProperties_GeoReplicationGenerator != nil {
		return databaseProperties_GeoReplicationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_GeoReplication(generators)
	databaseProperties_GeoReplicationGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_GeoReplication{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_GeoReplication(generators)
	AddRelatedPropertyGeneratorsForDatabaseProperties_GeoReplication(generators)
	databaseProperties_GeoReplicationGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_GeoReplication{}), generators)

	return databaseProperties_GeoReplicationGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseProperties_GeoReplication is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseProperties_GeoReplication(gens map[string]gopter.Gen) {
	gens["GroupNickname"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseProperties_GeoReplication is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseProperties_GeoReplication(gens map[string]gopter.Gen) {
	gens["LinkedDatabases"] = gen.SliceOf(LinkedDatabaseGenerator())
}

func Test_LinkedDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinkedDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinkedDatabase, LinkedDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinkedDatabase runs a test to see if a specific instance of LinkedDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForLinkedDatabase(subject LinkedDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinkedDatabase
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

// Generator of LinkedDatabase instances for property testing - lazily instantiated by LinkedDatabaseGenerator()
var linkedDatabaseGenerator gopter.Gen

// LinkedDatabaseGenerator returns a generator of LinkedDatabase instances for property testing.
func LinkedDatabaseGenerator() gopter.Gen {
	if linkedDatabaseGenerator != nil {
		return linkedDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinkedDatabase(generators)
	linkedDatabaseGenerator = gen.Struct(reflect.TypeOf(LinkedDatabase{}), generators)

	return linkedDatabaseGenerator
}

// AddIndependentPropertyGeneratorsForLinkedDatabase is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinkedDatabase(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Module_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Module via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForModule, ModuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForModule runs a test to see if a specific instance of Module round trips to JSON and back losslessly
func RunJSONSerializationTestForModule(subject Module) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Module
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

// Generator of Module instances for property testing - lazily instantiated by ModuleGenerator()
var moduleGenerator gopter.Gen

// ModuleGenerator returns a generator of Module instances for property testing.
func ModuleGenerator() gopter.Gen {
	if moduleGenerator != nil {
		return moduleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForModule(generators)
	moduleGenerator = gen.Struct(reflect.TypeOf(Module{}), generators)

	return moduleGenerator
}

// AddIndependentPropertyGeneratorsForModule is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForModule(gens map[string]gopter.Gen) {
	gens["Args"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_Persistence_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Persistence via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPersistence, PersistenceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPersistence runs a test to see if a specific instance of Persistence round trips to JSON and back losslessly
func RunJSONSerializationTestForPersistence(subject Persistence) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Persistence
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

// Generator of Persistence instances for property testing - lazily instantiated by PersistenceGenerator()
var persistenceGenerator gopter.Gen

// PersistenceGenerator returns a generator of Persistence instances for property testing.
func PersistenceGenerator() gopter.Gen {
	if persistenceGenerator != nil {
		return persistenceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPersistence(generators)
	persistenceGenerator = gen.Struct(reflect.TypeOf(Persistence{}), generators)

	return persistenceGenerator
}

// AddIndependentPropertyGeneratorsForPersistence is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPersistence(gens map[string]gopter.Gen) {
	gens["AofEnabled"] = gen.PtrOf(gen.Bool())
	gens["AofFrequency"] = gen.PtrOf(gen.OneConstOf(Persistence_AofFrequency_1S, Persistence_AofFrequency_Always))
	gens["RdbEnabled"] = gen.PtrOf(gen.Bool())
	gens["RdbFrequency"] = gen.PtrOf(gen.OneConstOf(Persistence_RdbFrequency_12H, Persistence_RdbFrequency_1H, Persistence_RdbFrequency_6H))
}

func Test_RedisEnterpriseDatabase_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterpriseDatabase_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterpriseDatabase_Spec, RedisEnterpriseDatabase_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterpriseDatabase_Spec runs a test to see if a specific instance of RedisEnterpriseDatabase_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterpriseDatabase_Spec(subject RedisEnterpriseDatabase_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterpriseDatabase_Spec
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

// Generator of RedisEnterpriseDatabase_Spec instances for property testing - lazily instantiated by
// RedisEnterpriseDatabase_SpecGenerator()
var redisEnterpriseDatabase_SpecGenerator gopter.Gen

// RedisEnterpriseDatabase_SpecGenerator returns a generator of RedisEnterpriseDatabase_Spec instances for property testing.
// We first initialize redisEnterpriseDatabase_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterpriseDatabase_SpecGenerator() gopter.Gen {
	if redisEnterpriseDatabase_SpecGenerator != nil {
		return redisEnterpriseDatabase_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_Spec(generators)
	redisEnterpriseDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabase_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_Spec(generators)
	AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_Spec(generators)
	redisEnterpriseDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabase_Spec{}), generators)

	return redisEnterpriseDatabase_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabasePropertiesGenerator())
}
