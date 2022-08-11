// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210301

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

func Test_RedisEnterpriseDatabases_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterpriseDatabases_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterpriseDatabasesSpecARM, RedisEnterpriseDatabasesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterpriseDatabasesSpecARM runs a test to see if a specific instance of RedisEnterpriseDatabases_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterpriseDatabasesSpecARM(subject RedisEnterpriseDatabases_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterpriseDatabases_SpecARM
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

// Generator of RedisEnterpriseDatabases_SpecARM instances for property testing - lazily instantiated by
// RedisEnterpriseDatabasesSpecARMGenerator()
var redisEnterpriseDatabasesSpecARMGenerator gopter.Gen

// RedisEnterpriseDatabasesSpecARMGenerator returns a generator of RedisEnterpriseDatabases_SpecARM instances for property testing.
// We first initialize redisEnterpriseDatabasesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterpriseDatabasesSpecARMGenerator() gopter.Gen {
	if redisEnterpriseDatabasesSpecARMGenerator != nil {
		return redisEnterpriseDatabasesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabasesSpecARM(generators)
	redisEnterpriseDatabasesSpecARMGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabases_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabasesSpecARM(generators)
	AddRelatedPropertyGeneratorsForRedisEnterpriseDatabasesSpecARM(generators)
	redisEnterpriseDatabasesSpecARMGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabases_SpecARM{}), generators)

	return redisEnterpriseDatabasesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterpriseDatabasesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterpriseDatabasesSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisEnterpriseDatabasesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterpriseDatabasesSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabasePropertiesARMGenerator())
}

func Test_DatabasePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabasePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabasePropertiesARM, DatabasePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabasePropertiesARM runs a test to see if a specific instance of DatabasePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabasePropertiesARM(subject DatabasePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabasePropertiesARM
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

// Generator of DatabasePropertiesARM instances for property testing - lazily instantiated by
// DatabasePropertiesARMGenerator()
var databasePropertiesARMGenerator gopter.Gen

// DatabasePropertiesARMGenerator returns a generator of DatabasePropertiesARM instances for property testing.
// We first initialize databasePropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabasePropertiesARMGenerator() gopter.Gen {
	if databasePropertiesARMGenerator != nil {
		return databasePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabasePropertiesARM(generators)
	databasePropertiesARMGenerator = gen.Struct(reflect.TypeOf(DatabasePropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabasePropertiesARM(generators)
	AddRelatedPropertyGeneratorsForDatabasePropertiesARM(generators)
	databasePropertiesARMGenerator = gen.Struct(reflect.TypeOf(DatabasePropertiesARM{}), generators)

	return databasePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabasePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabasePropertiesARM(gens map[string]gopter.Gen) {
	gens["ClientProtocol"] = gen.PtrOf(gen.OneConstOf(DatabasePropertiesClientProtocol_Encrypted, DatabasePropertiesClientProtocol_Plaintext))
	gens["ClusteringPolicy"] = gen.PtrOf(gen.OneConstOf(DatabasePropertiesClusteringPolicy_EnterpriseCluster, DatabasePropertiesClusteringPolicy_OSSCluster))
	gens["EvictionPolicy"] = gen.PtrOf(gen.OneConstOf(
		DatabasePropertiesEvictionPolicy_AllKeysLFU,
		DatabasePropertiesEvictionPolicy_AllKeysLRU,
		DatabasePropertiesEvictionPolicy_AllKeysRandom,
		DatabasePropertiesEvictionPolicy_NoEviction,
		DatabasePropertiesEvictionPolicy_VolatileLFU,
		DatabasePropertiesEvictionPolicy_VolatileLRU,
		DatabasePropertiesEvictionPolicy_VolatileRandom,
		DatabasePropertiesEvictionPolicy_VolatileTTL))
	gens["Port"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDatabasePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabasePropertiesARM(gens map[string]gopter.Gen) {
	gens["Modules"] = gen.SliceOf(ModuleARMGenerator())
	gens["Persistence"] = gen.PtrOf(PersistenceARMGenerator())
}

func Test_ModuleARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ModuleARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForModuleARM, ModuleARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForModuleARM runs a test to see if a specific instance of ModuleARM round trips to JSON and back losslessly
func RunJSONSerializationTestForModuleARM(subject ModuleARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ModuleARM
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

// Generator of ModuleARM instances for property testing - lazily instantiated by ModuleARMGenerator()
var moduleARMGenerator gopter.Gen

// ModuleARMGenerator returns a generator of ModuleARM instances for property testing.
func ModuleARMGenerator() gopter.Gen {
	if moduleARMGenerator != nil {
		return moduleARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForModuleARM(generators)
	moduleARMGenerator = gen.Struct(reflect.TypeOf(ModuleARM{}), generators)

	return moduleARMGenerator
}

// AddIndependentPropertyGeneratorsForModuleARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForModuleARM(gens map[string]gopter.Gen) {
	gens["Args"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_PersistenceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PersistenceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPersistenceARM, PersistenceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPersistenceARM runs a test to see if a specific instance of PersistenceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPersistenceARM(subject PersistenceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PersistenceARM
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

// Generator of PersistenceARM instances for property testing - lazily instantiated by PersistenceARMGenerator()
var persistenceARMGenerator gopter.Gen

// PersistenceARMGenerator returns a generator of PersistenceARM instances for property testing.
func PersistenceARMGenerator() gopter.Gen {
	if persistenceARMGenerator != nil {
		return persistenceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPersistenceARM(generators)
	persistenceARMGenerator = gen.Struct(reflect.TypeOf(PersistenceARM{}), generators)

	return persistenceARMGenerator
}

// AddIndependentPropertyGeneratorsForPersistenceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPersistenceARM(gens map[string]gopter.Gen) {
	gens["AofEnabled"] = gen.PtrOf(gen.Bool())
	gens["AofFrequency"] = gen.PtrOf(gen.OneConstOf(PersistenceAofFrequency_1S, PersistenceAofFrequency_Always))
	gens["RdbEnabled"] = gen.PtrOf(gen.Bool())
	gens["RdbFrequency"] = gen.PtrOf(gen.OneConstOf(PersistenceRdbFrequency_12H, PersistenceRdbFrequency_1H, PersistenceRdbFrequency_6H))
}
