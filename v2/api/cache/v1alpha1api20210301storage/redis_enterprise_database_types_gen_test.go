// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301storage

import (
	"encoding/json"
	v20210301s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20210301storage"
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

func Test_RedisEnterpriseDatabase_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisEnterpriseDatabase to hub returns original",
		prop.ForAll(RunResourceConversionTestForRedisEnterpriseDatabase, RedisEnterpriseDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRedisEnterpriseDatabase tests if a specific instance of RedisEnterpriseDatabase round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRedisEnterpriseDatabase(subject RedisEnterpriseDatabase) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210301s.RedisEnterpriseDatabase
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RedisEnterpriseDatabase
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisEnterpriseDatabase_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisEnterpriseDatabase to RedisEnterpriseDatabase via AssignPropertiesToRedisEnterpriseDatabase & AssignPropertiesFromRedisEnterpriseDatabase returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisEnterpriseDatabase, RedisEnterpriseDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisEnterpriseDatabase tests if a specific instance of RedisEnterpriseDatabase can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForRedisEnterpriseDatabase(subject RedisEnterpriseDatabase) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.RedisEnterpriseDatabase
	err := copied.AssignPropertiesToRedisEnterpriseDatabase(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisEnterpriseDatabase
	err = actual.AssignPropertiesFromRedisEnterpriseDatabase(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisEnterpriseDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterpriseDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterpriseDatabase, RedisEnterpriseDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterpriseDatabase runs a test to see if a specific instance of RedisEnterpriseDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterpriseDatabase(subject RedisEnterpriseDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterpriseDatabase
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

// Generator of RedisEnterpriseDatabase instances for property testing - lazily instantiated by
// RedisEnterpriseDatabaseGenerator()
var redisEnterpriseDatabaseGenerator gopter.Gen

// RedisEnterpriseDatabaseGenerator returns a generator of RedisEnterpriseDatabase instances for property testing.
func RedisEnterpriseDatabaseGenerator() gopter.Gen {
	if redisEnterpriseDatabaseGenerator != nil {
		return redisEnterpriseDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase(generators)
	redisEnterpriseDatabaseGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabase{}), generators)

	return redisEnterpriseDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisEnterpriseDatabasesSpecGenerator()
	gens["Status"] = DatabaseStatusGenerator()
}

func Test_Database_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Database_Status to Database_Status via AssignPropertiesToDatabaseStatus & AssignPropertiesFromDatabaseStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseStatus, DatabaseStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseStatus tests if a specific instance of Database_Status can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForDatabaseStatus(subject Database_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.Database_Status
	err := copied.AssignPropertiesToDatabaseStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Database_Status
	err = actual.AssignPropertiesFromDatabaseStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Database_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseStatus, DatabaseStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseStatus runs a test to see if a specific instance of Database_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseStatus(subject Database_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_Status
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

// Generator of Database_Status instances for property testing - lazily instantiated by DatabaseStatusGenerator()
var databaseStatusGenerator gopter.Gen

// DatabaseStatusGenerator returns a generator of Database_Status instances for property testing.
// We first initialize databaseStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseStatusGenerator() gopter.Gen {
	if databaseStatusGenerator != nil {
		return databaseStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseStatus(generators)
	databaseStatusGenerator = gen.Struct(reflect.TypeOf(Database_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseStatus(generators)
	AddRelatedPropertyGeneratorsForDatabaseStatus(generators)
	databaseStatusGenerator = gen.Struct(reflect.TypeOf(Database_Status{}), generators)

	return databaseStatusGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseStatus(gens map[string]gopter.Gen) {
	gens["ClientProtocol"] = gen.PtrOf(gen.AlphaString())
	gens["ClusteringPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["EvictionPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceState"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseStatus(gens map[string]gopter.Gen) {
	gens["Modules"] = gen.SliceOf(ModuleStatusGenerator())
	gens["Persistence"] = gen.PtrOf(PersistenceStatusGenerator())
}

func Test_RedisEnterpriseDatabases_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisEnterpriseDatabases_Spec to RedisEnterpriseDatabases_Spec via AssignPropertiesToRedisEnterpriseDatabasesSpec & AssignPropertiesFromRedisEnterpriseDatabasesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisEnterpriseDatabasesSpec, RedisEnterpriseDatabasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisEnterpriseDatabasesSpec tests if a specific instance of RedisEnterpriseDatabases_Spec can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForRedisEnterpriseDatabasesSpec(subject RedisEnterpriseDatabases_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.RedisEnterpriseDatabases_Spec
	err := copied.AssignPropertiesToRedisEnterpriseDatabasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisEnterpriseDatabases_Spec
	err = actual.AssignPropertiesFromRedisEnterpriseDatabasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisEnterpriseDatabases_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterpriseDatabases_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterpriseDatabasesSpec, RedisEnterpriseDatabasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterpriseDatabasesSpec runs a test to see if a specific instance of RedisEnterpriseDatabases_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterpriseDatabasesSpec(subject RedisEnterpriseDatabases_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterpriseDatabases_Spec
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

// Generator of RedisEnterpriseDatabases_Spec instances for property testing - lazily instantiated by
// RedisEnterpriseDatabasesSpecGenerator()
var redisEnterpriseDatabasesSpecGenerator gopter.Gen

// RedisEnterpriseDatabasesSpecGenerator returns a generator of RedisEnterpriseDatabases_Spec instances for property testing.
// We first initialize redisEnterpriseDatabasesSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterpriseDatabasesSpecGenerator() gopter.Gen {
	if redisEnterpriseDatabasesSpecGenerator != nil {
		return redisEnterpriseDatabasesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabasesSpec(generators)
	redisEnterpriseDatabasesSpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabases_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabasesSpec(generators)
	AddRelatedPropertyGeneratorsForRedisEnterpriseDatabasesSpec(generators)
	redisEnterpriseDatabasesSpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabases_Spec{}), generators)

	return redisEnterpriseDatabasesSpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterpriseDatabasesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterpriseDatabasesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["ClientProtocol"] = gen.PtrOf(gen.AlphaString())
	gens["ClusteringPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["EvictionPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisEnterpriseDatabasesSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterpriseDatabasesSpec(gens map[string]gopter.Gen) {
	gens["Modules"] = gen.SliceOf(ModuleGenerator())
	gens["Persistence"] = gen.PtrOf(PersistenceGenerator())
}

func Test_Module_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Module to Module via AssignPropertiesToModule & AssignPropertiesFromModule returns original",
		prop.ForAll(RunPropertyAssignmentTestForModule, ModuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForModule tests if a specific instance of Module can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForModule(subject Module) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.Module
	err := copied.AssignPropertiesToModule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Module
	err = actual.AssignPropertiesFromModule(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
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

func Test_Module_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Module_Status to Module_Status via AssignPropertiesToModuleStatus & AssignPropertiesFromModuleStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForModuleStatus, ModuleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForModuleStatus tests if a specific instance of Module_Status can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForModuleStatus(subject Module_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.Module_Status
	err := copied.AssignPropertiesToModuleStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Module_Status
	err = actual.AssignPropertiesFromModuleStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Module_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Module_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForModuleStatus, ModuleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForModuleStatus runs a test to see if a specific instance of Module_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForModuleStatus(subject Module_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Module_Status
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

// Generator of Module_Status instances for property testing - lazily instantiated by ModuleStatusGenerator()
var moduleStatusGenerator gopter.Gen

// ModuleStatusGenerator returns a generator of Module_Status instances for property testing.
func ModuleStatusGenerator() gopter.Gen {
	if moduleStatusGenerator != nil {
		return moduleStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForModuleStatus(generators)
	moduleStatusGenerator = gen.Struct(reflect.TypeOf(Module_Status{}), generators)

	return moduleStatusGenerator
}

// AddIndependentPropertyGeneratorsForModuleStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForModuleStatus(gens map[string]gopter.Gen) {
	gens["Args"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

func Test_Persistence_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Persistence to Persistence via AssignPropertiesToPersistence & AssignPropertiesFromPersistence returns original",
		prop.ForAll(RunPropertyAssignmentTestForPersistence, PersistenceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPersistence tests if a specific instance of Persistence can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForPersistence(subject Persistence) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.Persistence
	err := copied.AssignPropertiesToPersistence(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Persistence
	err = actual.AssignPropertiesFromPersistence(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
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
	gens["AofFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["RdbEnabled"] = gen.PtrOf(gen.Bool())
	gens["RdbFrequency"] = gen.PtrOf(gen.AlphaString())
}

func Test_Persistence_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Persistence_Status to Persistence_Status via AssignPropertiesToPersistenceStatus & AssignPropertiesFromPersistenceStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForPersistenceStatus, PersistenceStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPersistenceStatus tests if a specific instance of Persistence_Status can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForPersistenceStatus(subject Persistence_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.Persistence_Status
	err := copied.AssignPropertiesToPersistenceStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Persistence_Status
	err = actual.AssignPropertiesFromPersistenceStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Persistence_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Persistence_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPersistenceStatus, PersistenceStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPersistenceStatus runs a test to see if a specific instance of Persistence_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForPersistenceStatus(subject Persistence_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Persistence_Status
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

// Generator of Persistence_Status instances for property testing - lazily instantiated by PersistenceStatusGenerator()
var persistenceStatusGenerator gopter.Gen

// PersistenceStatusGenerator returns a generator of Persistence_Status instances for property testing.
func PersistenceStatusGenerator() gopter.Gen {
	if persistenceStatusGenerator != nil {
		return persistenceStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPersistenceStatus(generators)
	persistenceStatusGenerator = gen.Struct(reflect.TypeOf(Persistence_Status{}), generators)

	return persistenceStatusGenerator
}

// AddIndependentPropertyGeneratorsForPersistenceStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPersistenceStatus(gens map[string]gopter.Gen) {
	gens["AofEnabled"] = gen.PtrOf(gen.Bool())
	gens["AofFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["RdbEnabled"] = gen.PtrOf(gen.Bool())
	gens["RdbFrequency"] = gen.PtrOf(gen.AlphaString())
}
