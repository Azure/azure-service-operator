// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301

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

func Test_Database_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseStatusARM, DatabaseStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseStatusARM runs a test to see if a specific instance of Database_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseStatusARM(subject Database_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_StatusARM
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

// Generator of Database_StatusARM instances for property testing - lazily instantiated by DatabaseStatusARMGenerator()
var databaseStatusARMGenerator gopter.Gen

// DatabaseStatusARMGenerator returns a generator of Database_StatusARM instances for property testing.
// We first initialize databaseStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseStatusARMGenerator() gopter.Gen {
	if databaseStatusARMGenerator != nil {
		return databaseStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseStatusARM(generators)
	databaseStatusARMGenerator = gen.Struct(reflect.TypeOf(Database_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseStatusARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseStatusARM(generators)
	databaseStatusARMGenerator = gen.Struct(reflect.TypeOf(Database_StatusARM{}), generators)

	return databaseStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabasePropertiesStatusARMGenerator())
}

func Test_DatabaseProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabasePropertiesStatusARM, DatabasePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabasePropertiesStatusARM runs a test to see if a specific instance of DatabaseProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabasePropertiesStatusARM(subject DatabaseProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_StatusARM
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

// Generator of DatabaseProperties_StatusARM instances for property testing - lazily instantiated by
//DatabasePropertiesStatusARMGenerator()
var databasePropertiesStatusARMGenerator gopter.Gen

// DatabasePropertiesStatusARMGenerator returns a generator of DatabaseProperties_StatusARM instances for property testing.
// We first initialize databasePropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabasePropertiesStatusARMGenerator() gopter.Gen {
	if databasePropertiesStatusARMGenerator != nil {
		return databasePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabasePropertiesStatusARM(generators)
	databasePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabasePropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForDatabasePropertiesStatusARM(generators)
	databasePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_StatusARM{}), generators)

	return databasePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabasePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabasePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["ClientProtocol"] = gen.PtrOf(gen.OneConstOf(DatabasePropertiesStatusClientProtocolEncrypted, DatabasePropertiesStatusClientProtocolPlaintext))
	gens["ClusteringPolicy"] = gen.PtrOf(gen.OneConstOf(DatabasePropertiesStatusClusteringPolicyEnterpriseCluster, DatabasePropertiesStatusClusteringPolicyOSSCluster))
	gens["EvictionPolicy"] = gen.PtrOf(gen.OneConstOf(
		DatabasePropertiesStatusEvictionPolicyAllKeysLFU,
		DatabasePropertiesStatusEvictionPolicyAllKeysLRU,
		DatabasePropertiesStatusEvictionPolicyAllKeysRandom,
		DatabasePropertiesStatusEvictionPolicyNoEviction,
		DatabasePropertiesStatusEvictionPolicyVolatileLFU,
		DatabasePropertiesStatusEvictionPolicyVolatileLRU,
		DatabasePropertiesStatusEvictionPolicyVolatileRandom,
		DatabasePropertiesStatusEvictionPolicyVolatileTTL))
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_StatusCanceled,
		ProvisioningState_StatusCreating,
		ProvisioningState_StatusDeleting,
		ProvisioningState_StatusFailed,
		ProvisioningState_StatusSucceeded,
		ProvisioningState_StatusUpdating))
	gens["ResourceState"] = gen.PtrOf(gen.OneConstOf(
		ResourceState_StatusCreateFailed,
		ResourceState_StatusCreating,
		ResourceState_StatusDeleteFailed,
		ResourceState_StatusDeleting,
		ResourceState_StatusDisableFailed,
		ResourceState_StatusDisabled,
		ResourceState_StatusDisabling,
		ResourceState_StatusEnableFailed,
		ResourceState_StatusEnabling,
		ResourceState_StatusRunning,
		ResourceState_StatusUpdateFailed,
		ResourceState_StatusUpdating))
}

// AddRelatedPropertyGeneratorsForDatabasePropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabasePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Modules"] = gen.SliceOf(ModuleStatusARMGenerator())
	gens["Persistence"] = gen.PtrOf(PersistenceStatusARMGenerator())
}

func Test_Module_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Module_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForModuleStatusARM, ModuleStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForModuleStatusARM runs a test to see if a specific instance of Module_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForModuleStatusARM(subject Module_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Module_StatusARM
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

// Generator of Module_StatusARM instances for property testing - lazily instantiated by ModuleStatusARMGenerator()
var moduleStatusARMGenerator gopter.Gen

// ModuleStatusARMGenerator returns a generator of Module_StatusARM instances for property testing.
func ModuleStatusARMGenerator() gopter.Gen {
	if moduleStatusARMGenerator != nil {
		return moduleStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForModuleStatusARM(generators)
	moduleStatusARMGenerator = gen.Struct(reflect.TypeOf(Module_StatusARM{}), generators)

	return moduleStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForModuleStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForModuleStatusARM(gens map[string]gopter.Gen) {
	gens["Args"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

func Test_Persistence_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Persistence_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPersistenceStatusARM, PersistenceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPersistenceStatusARM runs a test to see if a specific instance of Persistence_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPersistenceStatusARM(subject Persistence_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Persistence_StatusARM
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

// Generator of Persistence_StatusARM instances for property testing - lazily instantiated by
//PersistenceStatusARMGenerator()
var persistenceStatusARMGenerator gopter.Gen

// PersistenceStatusARMGenerator returns a generator of Persistence_StatusARM instances for property testing.
func PersistenceStatusARMGenerator() gopter.Gen {
	if persistenceStatusARMGenerator != nil {
		return persistenceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPersistenceStatusARM(generators)
	persistenceStatusARMGenerator = gen.Struct(reflect.TypeOf(Persistence_StatusARM{}), generators)

	return persistenceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForPersistenceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPersistenceStatusARM(gens map[string]gopter.Gen) {
	gens["AofEnabled"] = gen.PtrOf(gen.Bool())
	gens["AofFrequency"] = gen.PtrOf(gen.OneConstOf(PersistenceStatusAofFrequency1S, PersistenceStatusAofFrequencyAlways))
	gens["RdbEnabled"] = gen.PtrOf(gen.Bool())
	gens["RdbFrequency"] = gen.PtrOf(gen.OneConstOf(PersistenceStatusRdbFrequency12H, PersistenceStatusRdbFrequency1H, PersistenceStatusRdbFrequency6H))
}
