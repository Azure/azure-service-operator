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
		"Round trip from RedisEnterpriseDatabase to RedisEnterpriseDatabase via AssignProperties_To_RedisEnterpriseDatabase & AssignProperties_From_RedisEnterpriseDatabase returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisEnterpriseDatabase, RedisEnterpriseDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisEnterpriseDatabase tests if a specific instance of RedisEnterpriseDatabase can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForRedisEnterpriseDatabase(subject RedisEnterpriseDatabase) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.RedisEnterpriseDatabase
	err := copied.AssignProperties_To_RedisEnterpriseDatabase(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisEnterpriseDatabase
	err = actual.AssignProperties_From_RedisEnterpriseDatabase(&other)
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
	gens["Spec"] = RedisEnterprise_Database_SpecGenerator()
	gens["Status"] = RedisEnterprise_Database_STATUSGenerator()
}

func Test_RedisEnterprise_Database_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisEnterprise_Database_Spec to RedisEnterprise_Database_Spec via AssignProperties_To_RedisEnterprise_Database_Spec & AssignProperties_From_RedisEnterprise_Database_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisEnterprise_Database_Spec, RedisEnterprise_Database_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisEnterprise_Database_Spec tests if a specific instance of RedisEnterprise_Database_Spec can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForRedisEnterprise_Database_Spec(subject RedisEnterprise_Database_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.RedisEnterprise_Database_Spec
	err := copied.AssignProperties_To_RedisEnterprise_Database_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisEnterprise_Database_Spec
	err = actual.AssignProperties_From_RedisEnterprise_Database_Spec(&other)
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

func Test_RedisEnterprise_Database_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterprise_Database_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterprise_Database_Spec, RedisEnterprise_Database_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterprise_Database_Spec runs a test to see if a specific instance of RedisEnterprise_Database_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterprise_Database_Spec(subject RedisEnterprise_Database_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterprise_Database_Spec
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

// Generator of RedisEnterprise_Database_Spec instances for property testing - lazily instantiated by
// RedisEnterprise_Database_SpecGenerator()
var redisEnterprise_Database_SpecGenerator gopter.Gen

// RedisEnterprise_Database_SpecGenerator returns a generator of RedisEnterprise_Database_Spec instances for property testing.
// We first initialize redisEnterprise_Database_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterprise_Database_SpecGenerator() gopter.Gen {
	if redisEnterprise_Database_SpecGenerator != nil {
		return redisEnterprise_Database_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_Database_Spec(generators)
	redisEnterprise_Database_SpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_Database_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_Database_Spec(generators)
	AddRelatedPropertyGeneratorsForRedisEnterprise_Database_Spec(generators)
	redisEnterprise_Database_SpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_Database_Spec{}), generators)

	return redisEnterprise_Database_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterprise_Database_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterprise_Database_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["ClientProtocol"] = gen.PtrOf(gen.AlphaString())
	gens["ClusteringPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["EvictionPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Port"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForRedisEnterprise_Database_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterprise_Database_Spec(gens map[string]gopter.Gen) {
	gens["Modules"] = gen.SliceOf(ModuleGenerator())
	gens["Persistence"] = gen.PtrOf(PersistenceGenerator())
}

func Test_RedisEnterprise_Database_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisEnterprise_Database_STATUS to RedisEnterprise_Database_STATUS via AssignProperties_To_RedisEnterprise_Database_STATUS & AssignProperties_From_RedisEnterprise_Database_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisEnterprise_Database_STATUS, RedisEnterprise_Database_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisEnterprise_Database_STATUS tests if a specific instance of RedisEnterprise_Database_STATUS can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForRedisEnterprise_Database_STATUS(subject RedisEnterprise_Database_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.RedisEnterprise_Database_STATUS
	err := copied.AssignProperties_To_RedisEnterprise_Database_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisEnterprise_Database_STATUS
	err = actual.AssignProperties_From_RedisEnterprise_Database_STATUS(&other)
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

func Test_RedisEnterprise_Database_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterprise_Database_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterprise_Database_STATUS, RedisEnterprise_Database_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterprise_Database_STATUS runs a test to see if a specific instance of RedisEnterprise_Database_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterprise_Database_STATUS(subject RedisEnterprise_Database_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterprise_Database_STATUS
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

// Generator of RedisEnterprise_Database_STATUS instances for property testing - lazily instantiated by
// RedisEnterprise_Database_STATUSGenerator()
var redisEnterprise_Database_STATUSGenerator gopter.Gen

// RedisEnterprise_Database_STATUSGenerator returns a generator of RedisEnterprise_Database_STATUS instances for property testing.
// We first initialize redisEnterprise_Database_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterprise_Database_STATUSGenerator() gopter.Gen {
	if redisEnterprise_Database_STATUSGenerator != nil {
		return redisEnterprise_Database_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_Database_STATUS(generators)
	redisEnterprise_Database_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_Database_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_Database_STATUS(generators)
	AddRelatedPropertyGeneratorsForRedisEnterprise_Database_STATUS(generators)
	redisEnterprise_Database_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_Database_STATUS{}), generators)

	return redisEnterprise_Database_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterprise_Database_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterprise_Database_STATUS(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForRedisEnterprise_Database_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterprise_Database_STATUS(gens map[string]gopter.Gen) {
	gens["Modules"] = gen.SliceOf(Module_STATUSGenerator())
	gens["Persistence"] = gen.PtrOf(Persistence_STATUSGenerator())
}

func Test_Module_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Module to Module via AssignProperties_To_Module & AssignProperties_From_Module returns original",
		prop.ForAll(RunPropertyAssignmentTestForModule, ModuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForModule tests if a specific instance of Module can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForModule(subject Module) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.Module
	err := copied.AssignProperties_To_Module(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Module
	err = actual.AssignProperties_From_Module(&other)
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

func Test_Module_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Module_STATUS to Module_STATUS via AssignProperties_To_Module_STATUS & AssignProperties_From_Module_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForModule_STATUS, Module_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForModule_STATUS tests if a specific instance of Module_STATUS can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForModule_STATUS(subject Module_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.Module_STATUS
	err := copied.AssignProperties_To_Module_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Module_STATUS
	err = actual.AssignProperties_From_Module_STATUS(&other)
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

func Test_Module_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Module_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForModule_STATUS, Module_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForModule_STATUS runs a test to see if a specific instance of Module_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForModule_STATUS(subject Module_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Module_STATUS
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

// Generator of Module_STATUS instances for property testing - lazily instantiated by Module_STATUSGenerator()
var module_STATUSGenerator gopter.Gen

// Module_STATUSGenerator returns a generator of Module_STATUS instances for property testing.
func Module_STATUSGenerator() gopter.Gen {
	if module_STATUSGenerator != nil {
		return module_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForModule_STATUS(generators)
	module_STATUSGenerator = gen.Struct(reflect.TypeOf(Module_STATUS{}), generators)

	return module_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForModule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForModule_STATUS(gens map[string]gopter.Gen) {
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
		"Round trip from Persistence to Persistence via AssignProperties_To_Persistence & AssignProperties_From_Persistence returns original",
		prop.ForAll(RunPropertyAssignmentTestForPersistence, PersistenceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPersistence tests if a specific instance of Persistence can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForPersistence(subject Persistence) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.Persistence
	err := copied.AssignProperties_To_Persistence(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Persistence
	err = actual.AssignProperties_From_Persistence(&other)
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

func Test_Persistence_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Persistence_STATUS to Persistence_STATUS via AssignProperties_To_Persistence_STATUS & AssignProperties_From_Persistence_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPersistence_STATUS, Persistence_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPersistence_STATUS tests if a specific instance of Persistence_STATUS can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForPersistence_STATUS(subject Persistence_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.Persistence_STATUS
	err := copied.AssignProperties_To_Persistence_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Persistence_STATUS
	err = actual.AssignProperties_From_Persistence_STATUS(&other)
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

func Test_Persistence_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Persistence_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPersistence_STATUS, Persistence_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPersistence_STATUS runs a test to see if a specific instance of Persistence_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPersistence_STATUS(subject Persistence_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Persistence_STATUS
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

// Generator of Persistence_STATUS instances for property testing - lazily instantiated by Persistence_STATUSGenerator()
var persistence_STATUSGenerator gopter.Gen

// Persistence_STATUSGenerator returns a generator of Persistence_STATUS instances for property testing.
func Persistence_STATUSGenerator() gopter.Gen {
	if persistence_STATUSGenerator != nil {
		return persistence_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPersistence_STATUS(generators)
	persistence_STATUSGenerator = gen.Struct(reflect.TypeOf(Persistence_STATUS{}), generators)

	return persistence_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPersistence_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPersistence_STATUS(gens map[string]gopter.Gen) {
	gens["AofEnabled"] = gen.PtrOf(gen.Bool())
	gens["AofFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["RdbEnabled"] = gen.PtrOf(gen.Bool())
	gens["RdbFrequency"] = gen.PtrOf(gen.AlphaString())
}
