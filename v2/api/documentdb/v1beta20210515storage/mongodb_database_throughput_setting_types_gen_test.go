// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

import (
	"encoding/json"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515storage"
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

func Test_MongodbDatabaseThroughputSetting_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from MongodbDatabaseThroughputSetting to hub returns original",
		prop.ForAll(RunResourceConversionTestForMongodbDatabaseThroughputSetting, MongodbDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForMongodbDatabaseThroughputSetting tests if a specific instance of MongodbDatabaseThroughputSetting round trips to the hub storage version and back losslessly
func RunResourceConversionTestForMongodbDatabaseThroughputSetting(subject MongodbDatabaseThroughputSetting) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210515s.MongodbDatabaseThroughputSetting
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual MongodbDatabaseThroughputSetting
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

func Test_MongodbDatabaseThroughputSetting_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from MongodbDatabaseThroughputSetting to MongodbDatabaseThroughputSetting via AssignProperties_To_MongodbDatabaseThroughputSetting & AssignProperties_From_MongodbDatabaseThroughputSetting returns original",
		prop.ForAll(RunPropertyAssignmentTestForMongodbDatabaseThroughputSetting, MongodbDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForMongodbDatabaseThroughputSetting tests if a specific instance of MongodbDatabaseThroughputSetting can be assigned to v1api20210515storage and back losslessly
func RunPropertyAssignmentTestForMongodbDatabaseThroughputSetting(subject MongodbDatabaseThroughputSetting) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.MongodbDatabaseThroughputSetting
	err := copied.AssignProperties_To_MongodbDatabaseThroughputSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual MongodbDatabaseThroughputSetting
	err = actual.AssignProperties_From_MongodbDatabaseThroughputSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_MongodbDatabaseThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabaseThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabaseThroughputSetting, MongodbDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabaseThroughputSetting runs a test to see if a specific instance of MongodbDatabaseThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabaseThroughputSetting(subject MongodbDatabaseThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabaseThroughputSetting
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

// Generator of MongodbDatabaseThroughputSetting instances for property testing - lazily instantiated by
// MongodbDatabaseThroughputSettingGenerator()
var mongodbDatabaseThroughputSettingGenerator gopter.Gen

// MongodbDatabaseThroughputSettingGenerator returns a generator of MongodbDatabaseThroughputSetting instances for property testing.
func MongodbDatabaseThroughputSettingGenerator() gopter.Gen {
	if mongodbDatabaseThroughputSettingGenerator != nil {
		return mongodbDatabaseThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongodbDatabaseThroughputSetting(generators)
	mongodbDatabaseThroughputSettingGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseThroughputSetting{}), generators)

	return mongodbDatabaseThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator()
	gens["Status"] = DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator()
}

func Test_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec to DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec via AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec & AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec, DatabaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec tests if a specific instance of DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec can be assigned to v1api20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(subject DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec
	err := copied.AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec
	err = actual.AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec, DatabaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec runs a test to see if a specific instance of DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(subject DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec
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

// Generator of DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec instances for property testing - lazily
// instantiated by DatabaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator()
var databaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator gopter.Gen

// DatabaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator returns a generator of DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec instances for property testing.
// We first initialize databaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator() gopter.Gen {
	if databaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator != nil {
		return databaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(generators)
	databaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(generators)
	databaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec{}), generators)

	return databaseAccounts_MongodbDatabases_ThroughputSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
}

func Test_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS to DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS via AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS & AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS, DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS tests if a specific instance of DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS can be assigned to v1api20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(subject DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS
	err := copied.AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS
	err = actual.AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS, DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS runs a test to see if a specific instance of DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(subject DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS
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

// Generator of DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS instances for property testing - lazily
// instantiated by DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator()
var databaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator gopter.Gen

// DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator returns a generator of DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS instances for property testing.
// We first initialize databaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator() gopter.Gen {
	if databaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator != nil {
		return databaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(generators)
	databaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(generators)
	databaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS{}), generators)

	return databaseAccounts_MongodbDatabases_ThroughputSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsGetProperties_Resource_STATUSGenerator())
}
