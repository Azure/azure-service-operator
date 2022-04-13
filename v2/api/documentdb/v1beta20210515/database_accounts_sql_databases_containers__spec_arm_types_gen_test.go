// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

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

func Test_DatabaseAccountsSqlDatabasesContainers_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainers_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersSpecARM, DatabaseAccountsSqlDatabasesContainersSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersSpecARM runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainers_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersSpecARM(subject DatabaseAccountsSqlDatabasesContainers_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainers_SpecARM
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

// Generator of DatabaseAccountsSqlDatabasesContainers_SpecARM instances for property testing - lazily instantiated by
//DatabaseAccountsSqlDatabasesContainersSpecARMGenerator()
var databaseAccountsSqlDatabasesContainersSpecARMGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersSpecARMGenerator returns a generator of DatabaseAccountsSqlDatabasesContainers_SpecARM instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersSpecARMGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersSpecARMGenerator != nil {
		return databaseAccountsSqlDatabasesContainersSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSpecARM(generators)
	databaseAccountsSqlDatabasesContainersSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainers_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSpecARM(generators)
	databaseAccountsSqlDatabasesContainersSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainers_SpecARM{}), generators)

	return databaseAccountsSqlDatabasesContainersSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlContainerCreateUpdatePropertiesARMGenerator())
}

func Test_SqlContainerCreateUpdatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerCreateUpdatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerCreateUpdatePropertiesARM, SqlContainerCreateUpdatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerCreateUpdatePropertiesARM runs a test to see if a specific instance of SqlContainerCreateUpdatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerCreateUpdatePropertiesARM(subject SqlContainerCreateUpdatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerCreateUpdatePropertiesARM
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

// Generator of SqlContainerCreateUpdatePropertiesARM instances for property testing - lazily instantiated by
//SqlContainerCreateUpdatePropertiesARMGenerator()
var sqlContainerCreateUpdatePropertiesARMGenerator gopter.Gen

// SqlContainerCreateUpdatePropertiesARMGenerator returns a generator of SqlContainerCreateUpdatePropertiesARM instances for property testing.
func SqlContainerCreateUpdatePropertiesARMGenerator() gopter.Gen {
	if sqlContainerCreateUpdatePropertiesARMGenerator != nil {
		return sqlContainerCreateUpdatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlContainerCreateUpdatePropertiesARM(generators)
	sqlContainerCreateUpdatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerCreateUpdatePropertiesARM{}), generators)

	return sqlContainerCreateUpdatePropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlContainerCreateUpdatePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerCreateUpdatePropertiesARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsARMGenerator())
	gens["Resource"] = gen.PtrOf(SqlContainerResourceARMGenerator())
}

func Test_SqlContainerResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerResourceARM, SqlContainerResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerResourceARM runs a test to see if a specific instance of SqlContainerResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerResourceARM(subject SqlContainerResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerResourceARM
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

// Generator of SqlContainerResourceARM instances for property testing - lazily instantiated by
//SqlContainerResourceARMGenerator()
var sqlContainerResourceARMGenerator gopter.Gen

// SqlContainerResourceARMGenerator returns a generator of SqlContainerResourceARM instances for property testing.
// We first initialize sqlContainerResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlContainerResourceARMGenerator() gopter.Gen {
	if sqlContainerResourceARMGenerator != nil {
		return sqlContainerResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerResourceARM(generators)
	sqlContainerResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerResourceARM(generators)
	AddRelatedPropertyGeneratorsForSqlContainerResourceARM(generators)
	sqlContainerResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerResourceARM{}), generators)

	return sqlContainerResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlContainerResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlContainerResourceARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["DefaultTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlContainerResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerResourceARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPolicy"] = gen.PtrOf(ConflictResolutionPolicyARMGenerator())
	gens["IndexingPolicy"] = gen.PtrOf(IndexingPolicyARMGenerator())
	gens["PartitionKey"] = gen.PtrOf(ContainerPartitionKeyARMGenerator())
	gens["UniqueKeyPolicy"] = gen.PtrOf(UniqueKeyPolicyARMGenerator())
}

func Test_ConflictResolutionPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConflictResolutionPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConflictResolutionPolicyARM, ConflictResolutionPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConflictResolutionPolicyARM runs a test to see if a specific instance of ConflictResolutionPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConflictResolutionPolicyARM(subject ConflictResolutionPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConflictResolutionPolicyARM
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

// Generator of ConflictResolutionPolicyARM instances for property testing - lazily instantiated by
//ConflictResolutionPolicyARMGenerator()
var conflictResolutionPolicyARMGenerator gopter.Gen

// ConflictResolutionPolicyARMGenerator returns a generator of ConflictResolutionPolicyARM instances for property testing.
func ConflictResolutionPolicyARMGenerator() gopter.Gen {
	if conflictResolutionPolicyARMGenerator != nil {
		return conflictResolutionPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConflictResolutionPolicyARM(generators)
	conflictResolutionPolicyARMGenerator = gen.Struct(reflect.TypeOf(ConflictResolutionPolicyARM{}), generators)

	return conflictResolutionPolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForConflictResolutionPolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConflictResolutionPolicyARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPath"] = gen.PtrOf(gen.AlphaString())
	gens["ConflictResolutionProcedure"] = gen.PtrOf(gen.AlphaString())
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(ConflictResolutionPolicyModeCustom, ConflictResolutionPolicyModeLastWriterWins))
}

func Test_ContainerPartitionKeyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerPartitionKeyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerPartitionKeyARM, ContainerPartitionKeyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerPartitionKeyARM runs a test to see if a specific instance of ContainerPartitionKeyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerPartitionKeyARM(subject ContainerPartitionKeyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerPartitionKeyARM
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

// Generator of ContainerPartitionKeyARM instances for property testing - lazily instantiated by
//ContainerPartitionKeyARMGenerator()
var containerPartitionKeyARMGenerator gopter.Gen

// ContainerPartitionKeyARMGenerator returns a generator of ContainerPartitionKeyARM instances for property testing.
func ContainerPartitionKeyARMGenerator() gopter.Gen {
	if containerPartitionKeyARMGenerator != nil {
		return containerPartitionKeyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPartitionKeyARM(generators)
	containerPartitionKeyARMGenerator = gen.Struct(reflect.TypeOf(ContainerPartitionKeyARM{}), generators)

	return containerPartitionKeyARMGenerator
}

// AddIndependentPropertyGeneratorsForContainerPartitionKeyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerPartitionKeyARM(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(ContainerPartitionKeyKindHash, ContainerPartitionKeyKindMultiHash, ContainerPartitionKeyKindRange))
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.Int())
}

func Test_IndexingPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IndexingPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexingPolicyARM, IndexingPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexingPolicyARM runs a test to see if a specific instance of IndexingPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexingPolicyARM(subject IndexingPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IndexingPolicyARM
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

// Generator of IndexingPolicyARM instances for property testing - lazily instantiated by IndexingPolicyARMGenerator()
var indexingPolicyARMGenerator gopter.Gen

// IndexingPolicyARMGenerator returns a generator of IndexingPolicyARM instances for property testing.
// We first initialize indexingPolicyARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IndexingPolicyARMGenerator() gopter.Gen {
	if indexingPolicyARMGenerator != nil {
		return indexingPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicyARM(generators)
	indexingPolicyARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicyARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicyARM(generators)
	AddRelatedPropertyGeneratorsForIndexingPolicyARM(generators)
	indexingPolicyARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicyARM{}), generators)

	return indexingPolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexingPolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexingPolicyARM(gens map[string]gopter.Gen) {
	gens["Automatic"] = gen.PtrOf(gen.Bool())
	gens["IndexingMode"] = gen.PtrOf(gen.OneConstOf(IndexingPolicyIndexingModeConsistent, IndexingPolicyIndexingModeLazy, IndexingPolicyIndexingModeNone))
}

// AddRelatedPropertyGeneratorsForIndexingPolicyARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIndexingPolicyARM(gens map[string]gopter.Gen) {
	gens["CompositeIndexes"] = gen.SliceOf(gen.SliceOf(CompositePathARMGenerator()))
	gens["ExcludedPaths"] = gen.SliceOf(ExcludedPathARMGenerator())
	gens["IncludedPaths"] = gen.SliceOf(IncludedPathARMGenerator())
	gens["SpatialIndexes"] = gen.SliceOf(SpatialSpecARMGenerator())
}

func Test_UniqueKeyPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKeyPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKeyPolicyARM, UniqueKeyPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKeyPolicyARM runs a test to see if a specific instance of UniqueKeyPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKeyPolicyARM(subject UniqueKeyPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKeyPolicyARM
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

// Generator of UniqueKeyPolicyARM instances for property testing - lazily instantiated by UniqueKeyPolicyARMGenerator()
var uniqueKeyPolicyARMGenerator gopter.Gen

// UniqueKeyPolicyARMGenerator returns a generator of UniqueKeyPolicyARM instances for property testing.
func UniqueKeyPolicyARMGenerator() gopter.Gen {
	if uniqueKeyPolicyARMGenerator != nil {
		return uniqueKeyPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUniqueKeyPolicyARM(generators)
	uniqueKeyPolicyARMGenerator = gen.Struct(reflect.TypeOf(UniqueKeyPolicyARM{}), generators)

	return uniqueKeyPolicyARMGenerator
}

// AddRelatedPropertyGeneratorsForUniqueKeyPolicyARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUniqueKeyPolicyARM(gens map[string]gopter.Gen) {
	gens["UniqueKeys"] = gen.SliceOf(UniqueKeyARMGenerator())
}

func Test_CompositePathARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CompositePathARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCompositePathARM, CompositePathARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCompositePathARM runs a test to see if a specific instance of CompositePathARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCompositePathARM(subject CompositePathARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CompositePathARM
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

// Generator of CompositePathARM instances for property testing - lazily instantiated by CompositePathARMGenerator()
var compositePathARMGenerator gopter.Gen

// CompositePathARMGenerator returns a generator of CompositePathARM instances for property testing.
func CompositePathARMGenerator() gopter.Gen {
	if compositePathARMGenerator != nil {
		return compositePathARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCompositePathARM(generators)
	compositePathARMGenerator = gen.Struct(reflect.TypeOf(CompositePathARM{}), generators)

	return compositePathARMGenerator
}

// AddIndependentPropertyGeneratorsForCompositePathARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCompositePathARM(gens map[string]gopter.Gen) {
	gens["Order"] = gen.PtrOf(gen.OneConstOf(CompositePathOrderAscending, CompositePathOrderDescending))
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_ExcludedPathARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExcludedPathARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExcludedPathARM, ExcludedPathARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExcludedPathARM runs a test to see if a specific instance of ExcludedPathARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExcludedPathARM(subject ExcludedPathARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExcludedPathARM
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

// Generator of ExcludedPathARM instances for property testing - lazily instantiated by ExcludedPathARMGenerator()
var excludedPathARMGenerator gopter.Gen

// ExcludedPathARMGenerator returns a generator of ExcludedPathARM instances for property testing.
func ExcludedPathARMGenerator() gopter.Gen {
	if excludedPathARMGenerator != nil {
		return excludedPathARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExcludedPathARM(generators)
	excludedPathARMGenerator = gen.Struct(reflect.TypeOf(ExcludedPathARM{}), generators)

	return excludedPathARMGenerator
}

// AddIndependentPropertyGeneratorsForExcludedPathARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExcludedPathARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_IncludedPathARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IncludedPathARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIncludedPathARM, IncludedPathARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIncludedPathARM runs a test to see if a specific instance of IncludedPathARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIncludedPathARM(subject IncludedPathARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IncludedPathARM
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

// Generator of IncludedPathARM instances for property testing - lazily instantiated by IncludedPathARMGenerator()
var includedPathARMGenerator gopter.Gen

// IncludedPathARMGenerator returns a generator of IncludedPathARM instances for property testing.
// We first initialize includedPathARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IncludedPathARMGenerator() gopter.Gen {
	if includedPathARMGenerator != nil {
		return includedPathARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPathARM(generators)
	includedPathARMGenerator = gen.Struct(reflect.TypeOf(IncludedPathARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPathARM(generators)
	AddRelatedPropertyGeneratorsForIncludedPathARM(generators)
	includedPathARMGenerator = gen.Struct(reflect.TypeOf(IncludedPathARM{}), generators)

	return includedPathARMGenerator
}

// AddIndependentPropertyGeneratorsForIncludedPathARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIncludedPathARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForIncludedPathARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIncludedPathARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(IndexesARMGenerator())
}

func Test_SpatialSpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SpatialSpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSpatialSpecARM, SpatialSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSpatialSpecARM runs a test to see if a specific instance of SpatialSpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSpatialSpecARM(subject SpatialSpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SpatialSpecARM
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

// Generator of SpatialSpecARM instances for property testing - lazily instantiated by SpatialSpecARMGenerator()
var spatialSpecARMGenerator gopter.Gen

// SpatialSpecARMGenerator returns a generator of SpatialSpecARM instances for property testing.
func SpatialSpecARMGenerator() gopter.Gen {
	if spatialSpecARMGenerator != nil {
		return spatialSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSpatialSpecARM(generators)
	spatialSpecARMGenerator = gen.Struct(reflect.TypeOf(SpatialSpecARM{}), generators)

	return spatialSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForSpatialSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSpatialSpecARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
	gens["Types"] = gen.SliceOf(gen.OneConstOf(
		SpatialSpecTypesLineString,
		SpatialSpecTypesMultiPolygon,
		SpatialSpecTypesPoint,
		SpatialSpecTypesPolygon))
}

func Test_UniqueKeyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKeyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKeyARM, UniqueKeyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKeyARM runs a test to see if a specific instance of UniqueKeyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKeyARM(subject UniqueKeyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKeyARM
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

// Generator of UniqueKeyARM instances for property testing - lazily instantiated by UniqueKeyARMGenerator()
var uniqueKeyARMGenerator gopter.Gen

// UniqueKeyARMGenerator returns a generator of UniqueKeyARM instances for property testing.
func UniqueKeyARMGenerator() gopter.Gen {
	if uniqueKeyARMGenerator != nil {
		return uniqueKeyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUniqueKeyARM(generators)
	uniqueKeyARMGenerator = gen.Struct(reflect.TypeOf(UniqueKeyARM{}), generators)

	return uniqueKeyARMGenerator
}

// AddIndependentPropertyGeneratorsForUniqueKeyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUniqueKeyARM(gens map[string]gopter.Gen) {
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
}

func Test_IndexesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IndexesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexesARM, IndexesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexesARM runs a test to see if a specific instance of IndexesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexesARM(subject IndexesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IndexesARM
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

// Generator of IndexesARM instances for property testing - lazily instantiated by IndexesARMGenerator()
var indexesARMGenerator gopter.Gen

// IndexesARMGenerator returns a generator of IndexesARM instances for property testing.
func IndexesARMGenerator() gopter.Gen {
	if indexesARMGenerator != nil {
		return indexesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexesARM(generators)
	indexesARMGenerator = gen.Struct(reflect.TypeOf(IndexesARM{}), generators)

	return indexesARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexesARM(gens map[string]gopter.Gen) {
	gens["DataType"] = gen.PtrOf(gen.OneConstOf(
		IndexesDataTypeLineString,
		IndexesDataTypeMultiPolygon,
		IndexesDataTypeNumber,
		IndexesDataTypePoint,
		IndexesDataTypePolygon,
		IndexesDataTypeString))
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(IndexesKindHash, IndexesKindRange, IndexesKindSpatial))
	gens["Precision"] = gen.PtrOf(gen.Int())
}
