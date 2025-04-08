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

func Test_DatabaseStatistics_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseStatistics_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseStatistics_STATUS, DatabaseStatistics_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseStatistics_STATUS runs a test to see if a specific instance of DatabaseStatistics_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseStatistics_STATUS(subject DatabaseStatistics_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseStatistics_STATUS
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

// Generator of DatabaseStatistics_STATUS instances for property testing - lazily instantiated by
// DatabaseStatistics_STATUSGenerator()
var databaseStatistics_STATUSGenerator gopter.Gen

// DatabaseStatistics_STATUSGenerator returns a generator of DatabaseStatistics_STATUS instances for property testing.
func DatabaseStatistics_STATUSGenerator() gopter.Gen {
	if databaseStatistics_STATUSGenerator != nil {
		return databaseStatistics_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseStatistics_STATUS(generators)
	databaseStatistics_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseStatistics_STATUS{}), generators)

	return databaseStatistics_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseStatistics_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseStatistics_STATUS(gens map[string]gopter.Gen) {
	gens["Size"] = gen.PtrOf(gen.Float64())
}

func Test_Database_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabase_STATUS, Database_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabase_STATUS runs a test to see if a specific instance of Database_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabase_STATUS(subject Database_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_STATUS
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

// Generator of Database_STATUS instances for property testing - lazily instantiated by Database_STATUSGenerator()
var database_STATUSGenerator gopter.Gen

// Database_STATUSGenerator returns a generator of Database_STATUS instances for property testing.
// We first initialize database_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Database_STATUSGenerator() gopter.Gen {
	if database_STATUSGenerator != nil {
		return database_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabase_STATUS(generators)
	database_STATUSGenerator = gen.OneGenOf(gens...)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabase_STATUS(generators)
	AddRelatedPropertyGeneratorsForDatabase_STATUS(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(Database_STATUS{}), map[string]gopter.Gen{propName: propGen}))
	}
	database_STATUSGenerator = gen.OneGenOf(gens...)

	return database_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["ReadOnlyFollowing"] = ReadOnlyFollowingDatabase_STATUSGenerator().Map(func(it ReadOnlyFollowingDatabase_STATUS) *ReadOnlyFollowingDatabase_STATUS {
		return &it
	}) // generate one case for OneOf type
	gens["ReadWrite"] = ReadWriteDatabase_STATUSGenerator().Map(func(it ReadWriteDatabase_STATUS) *ReadWriteDatabase_STATUS {
		return &it
	}) // generate one case for OneOf type
}

func Test_ReadOnlyFollowingDatabaseProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadOnlyFollowingDatabaseProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadOnlyFollowingDatabaseProperties_STATUS, ReadOnlyFollowingDatabaseProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadOnlyFollowingDatabaseProperties_STATUS runs a test to see if a specific instance of ReadOnlyFollowingDatabaseProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForReadOnlyFollowingDatabaseProperties_STATUS(subject ReadOnlyFollowingDatabaseProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadOnlyFollowingDatabaseProperties_STATUS
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

// Generator of ReadOnlyFollowingDatabaseProperties_STATUS instances for property testing - lazily instantiated by
// ReadOnlyFollowingDatabaseProperties_STATUSGenerator()
var readOnlyFollowingDatabaseProperties_STATUSGenerator gopter.Gen

// ReadOnlyFollowingDatabaseProperties_STATUSGenerator returns a generator of ReadOnlyFollowingDatabaseProperties_STATUS instances for property testing.
// We first initialize readOnlyFollowingDatabaseProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadOnlyFollowingDatabaseProperties_STATUSGenerator() gopter.Gen {
	if readOnlyFollowingDatabaseProperties_STATUSGenerator != nil {
		return readOnlyFollowingDatabaseProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS(generators)
	readOnlyFollowingDatabaseProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadOnlyFollowingDatabaseProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS(generators)
	readOnlyFollowingDatabaseProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadOnlyFollowingDatabaseProperties_STATUS{}), generators)

	return readOnlyFollowingDatabaseProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS(gens map[string]gopter.Gen) {
	gens["AttachedDatabaseConfigurationName"] = gen.PtrOf(gen.AlphaString())
	gens["DatabaseShareOrigin"] = gen.PtrOf(gen.OneConstOf(DatabaseShareOrigin_STATUS_DataShare, DatabaseShareOrigin_STATUS_Direct, DatabaseShareOrigin_STATUS_Other))
	gens["HotCachePeriod"] = gen.PtrOf(gen.AlphaString())
	gens["LeaderClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalDatabaseName"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalsModificationKind"] = gen.PtrOf(gen.OneConstOf(ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_None, ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_Replace, ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_Union))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Canceled,
		ProvisioningState_STATUS_Creating,
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Moving,
		ProvisioningState_STATUS_Running,
		ProvisioningState_STATUS_Succeeded))
	gens["SoftDeletePeriod"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Statistics"] = gen.PtrOf(DatabaseStatistics_STATUSGenerator())
	gens["SuspensionDetails"] = gen.PtrOf(SuspensionDetails_STATUSGenerator())
	gens["TableLevelSharingProperties"] = gen.PtrOf(TableLevelSharingProperties_STATUSGenerator())
}

func Test_ReadOnlyFollowingDatabase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadOnlyFollowingDatabase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadOnlyFollowingDatabase_STATUS, ReadOnlyFollowingDatabase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadOnlyFollowingDatabase_STATUS runs a test to see if a specific instance of ReadOnlyFollowingDatabase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForReadOnlyFollowingDatabase_STATUS(subject ReadOnlyFollowingDatabase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadOnlyFollowingDatabase_STATUS
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

// Generator of ReadOnlyFollowingDatabase_STATUS instances for property testing - lazily instantiated by
// ReadOnlyFollowingDatabase_STATUSGenerator()
var readOnlyFollowingDatabase_STATUSGenerator gopter.Gen

// ReadOnlyFollowingDatabase_STATUSGenerator returns a generator of ReadOnlyFollowingDatabase_STATUS instances for property testing.
// We first initialize readOnlyFollowingDatabase_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadOnlyFollowingDatabase_STATUSGenerator() gopter.Gen {
	if readOnlyFollowingDatabase_STATUSGenerator != nil {
		return readOnlyFollowingDatabase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS(generators)
	readOnlyFollowingDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadOnlyFollowingDatabase_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS(generators)
	AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS(generators)
	readOnlyFollowingDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadOnlyFollowingDatabase_STATUS{}), generators)

	return readOnlyFollowingDatabase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.OneConstOf(ReadOnlyFollowingDatabase_Kind_STATUS_ReadOnlyFollowing)
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ReadOnlyFollowingDatabaseProperties_STATUSGenerator())
}

func Test_ReadWriteDatabaseProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadWriteDatabaseProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadWriteDatabaseProperties_STATUS, ReadWriteDatabaseProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadWriteDatabaseProperties_STATUS runs a test to see if a specific instance of ReadWriteDatabaseProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForReadWriteDatabaseProperties_STATUS(subject ReadWriteDatabaseProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadWriteDatabaseProperties_STATUS
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

// Generator of ReadWriteDatabaseProperties_STATUS instances for property testing - lazily instantiated by
// ReadWriteDatabaseProperties_STATUSGenerator()
var readWriteDatabaseProperties_STATUSGenerator gopter.Gen

// ReadWriteDatabaseProperties_STATUSGenerator returns a generator of ReadWriteDatabaseProperties_STATUS instances for property testing.
// We first initialize readWriteDatabaseProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadWriteDatabaseProperties_STATUSGenerator() gopter.Gen {
	if readWriteDatabaseProperties_STATUSGenerator != nil {
		return readWriteDatabaseProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS(generators)
	readWriteDatabaseProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabaseProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties_STATUS(generators)
	readWriteDatabaseProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabaseProperties_STATUS{}), generators)

	return readWriteDatabaseProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS(gens map[string]gopter.Gen) {
	gens["HotCachePeriod"] = gen.PtrOf(gen.AlphaString())
	gens["IsFollowed"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Canceled,
		ProvisioningState_STATUS_Creating,
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Moving,
		ProvisioningState_STATUS_Running,
		ProvisioningState_STATUS_Succeeded))
	gens["SoftDeletePeriod"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties_STATUS(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultProperties_STATUSGenerator())
	gens["Statistics"] = gen.PtrOf(DatabaseStatistics_STATUSGenerator())
	gens["SuspensionDetails"] = gen.PtrOf(SuspensionDetails_STATUSGenerator())
}

func Test_ReadWriteDatabase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadWriteDatabase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadWriteDatabase_STATUS, ReadWriteDatabase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadWriteDatabase_STATUS runs a test to see if a specific instance of ReadWriteDatabase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForReadWriteDatabase_STATUS(subject ReadWriteDatabase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadWriteDatabase_STATUS
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

// Generator of ReadWriteDatabase_STATUS instances for property testing - lazily instantiated by
// ReadWriteDatabase_STATUSGenerator()
var readWriteDatabase_STATUSGenerator gopter.Gen

// ReadWriteDatabase_STATUSGenerator returns a generator of ReadWriteDatabase_STATUS instances for property testing.
// We first initialize readWriteDatabase_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadWriteDatabase_STATUSGenerator() gopter.Gen {
	if readWriteDatabase_STATUSGenerator != nil {
		return readWriteDatabase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS(generators)
	readWriteDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabase_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS(generators)
	AddRelatedPropertyGeneratorsForReadWriteDatabase_STATUS(generators)
	readWriteDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabase_STATUS{}), generators)

	return readWriteDatabase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.OneConstOf(ReadWriteDatabase_Kind_STATUS_ReadWrite)
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReadWriteDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadWriteDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ReadWriteDatabaseProperties_STATUSGenerator())
}

func Test_SuspensionDetails_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SuspensionDetails_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSuspensionDetails_STATUS, SuspensionDetails_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSuspensionDetails_STATUS runs a test to see if a specific instance of SuspensionDetails_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSuspensionDetails_STATUS(subject SuspensionDetails_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SuspensionDetails_STATUS
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

// Generator of SuspensionDetails_STATUS instances for property testing - lazily instantiated by
// SuspensionDetails_STATUSGenerator()
var suspensionDetails_STATUSGenerator gopter.Gen

// SuspensionDetails_STATUSGenerator returns a generator of SuspensionDetails_STATUS instances for property testing.
func SuspensionDetails_STATUSGenerator() gopter.Gen {
	if suspensionDetails_STATUSGenerator != nil {
		return suspensionDetails_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSuspensionDetails_STATUS(generators)
	suspensionDetails_STATUSGenerator = gen.Struct(reflect.TypeOf(SuspensionDetails_STATUS{}), generators)

	return suspensionDetails_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSuspensionDetails_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSuspensionDetails_STATUS(gens map[string]gopter.Gen) {
	gens["SuspensionStartDate"] = gen.PtrOf(gen.AlphaString())
}

func Test_TableLevelSharingProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableLevelSharingProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableLevelSharingProperties_STATUS, TableLevelSharingProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableLevelSharingProperties_STATUS runs a test to see if a specific instance of TableLevelSharingProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTableLevelSharingProperties_STATUS(subject TableLevelSharingProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableLevelSharingProperties_STATUS
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

// Generator of TableLevelSharingProperties_STATUS instances for property testing - lazily instantiated by
// TableLevelSharingProperties_STATUSGenerator()
var tableLevelSharingProperties_STATUSGenerator gopter.Gen

// TableLevelSharingProperties_STATUSGenerator returns a generator of TableLevelSharingProperties_STATUS instances for property testing.
func TableLevelSharingProperties_STATUSGenerator() gopter.Gen {
	if tableLevelSharingProperties_STATUSGenerator != nil {
		return tableLevelSharingProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableLevelSharingProperties_STATUS(generators)
	tableLevelSharingProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(TableLevelSharingProperties_STATUS{}), generators)

	return tableLevelSharingProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTableLevelSharingProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableLevelSharingProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ExternalTablesToExclude"] = gen.SliceOf(gen.AlphaString())
	gens["ExternalTablesToInclude"] = gen.SliceOf(gen.AlphaString())
	gens["FunctionsToExclude"] = gen.SliceOf(gen.AlphaString())
	gens["FunctionsToInclude"] = gen.SliceOf(gen.AlphaString())
	gens["MaterializedViewsToExclude"] = gen.SliceOf(gen.AlphaString())
	gens["MaterializedViewsToInclude"] = gen.SliceOf(gen.AlphaString())
	gens["TablesToExclude"] = gen.SliceOf(gen.AlphaString())
	gens["TablesToInclude"] = gen.SliceOf(gen.AlphaString())
}
