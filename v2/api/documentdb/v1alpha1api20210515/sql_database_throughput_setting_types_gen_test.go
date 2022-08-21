// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	alpha20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
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

func Test_SqlDatabaseThroughputSetting_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseThroughputSetting to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseThroughputSetting tests if a specific instance of SqlDatabaseThroughputSetting round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210515s.SqlDatabaseThroughputSetting
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseThroughputSetting
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

func Test_SqlDatabaseThroughputSetting_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseThroughputSetting to SqlDatabaseThroughputSetting via AssignPropertiesToSqlDatabaseThroughputSetting & AssignPropertiesFromSqlDatabaseThroughputSetting returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseThroughputSetting tests if a specific instance of SqlDatabaseThroughputSetting can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210515s.SqlDatabaseThroughputSetting
	err := copied.AssignPropertiesToSqlDatabaseThroughputSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseThroughputSetting
	err = actual.AssignPropertiesFromSqlDatabaseThroughputSetting(&other)
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

func Test_SqlDatabaseThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseThroughputSetting runs a test to see if a specific instance of SqlDatabaseThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseThroughputSetting
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

// Generator of SqlDatabaseThroughputSetting instances for property testing - lazily instantiated by
// SqlDatabaseThroughputSettingGenerator()
var sqlDatabaseThroughputSettingGenerator gopter.Gen

// SqlDatabaseThroughputSettingGenerator returns a generator of SqlDatabaseThroughputSetting instances for property testing.
func SqlDatabaseThroughputSettingGenerator() gopter.Gen {
	if sqlDatabaseThroughputSettingGenerator != nil {
		return sqlDatabaseThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting(generators)
	sqlDatabaseThroughputSettingGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseThroughputSetting{}), generators)

	return sqlDatabaseThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesThroughputSetting_SpecGenerator()
	gens["Status"] = DatabaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator()
}

func Test_DatabaseAccountsSqlDatabasesThroughputSetting_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabasesThroughputSetting_Spec to DatabaseAccountsSqlDatabasesThroughputSetting_Spec via AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSetting_Spec & AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSetting_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesThroughputSetting_Spec, DatabaseAccountsSqlDatabasesThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesThroughputSetting_Spec tests if a specific instance of DatabaseAccountsSqlDatabasesThroughputSetting_Spec can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesThroughputSetting_Spec(subject DatabaseAccountsSqlDatabasesThroughputSetting_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210515s.DatabaseAccountsSqlDatabasesThroughputSetting_Spec
	err := copied.AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSetting_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabasesThroughputSetting_Spec
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSetting_Spec(&other)
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

func Test_DatabaseAccountsSqlDatabasesThroughputSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesThroughputSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSetting_Spec, DatabaseAccountsSqlDatabasesThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSetting_Spec runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesThroughputSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSetting_Spec(subject DatabaseAccountsSqlDatabasesThroughputSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesThroughputSetting_Spec
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

// Generator of DatabaseAccountsSqlDatabasesThroughputSetting_Spec instances for property testing - lazily instantiated
// by DatabaseAccountsSqlDatabasesThroughputSetting_SpecGenerator()
var databaseAccountsSqlDatabasesThroughputSetting_SpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesThroughputSetting_SpecGenerator returns a generator of DatabaseAccountsSqlDatabasesThroughputSetting_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesThroughputSetting_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesThroughputSetting_SpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesThroughputSetting_SpecGenerator != nil {
		return databaseAccountsSqlDatabasesThroughputSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_Spec(generators)
	databaseAccountsSqlDatabasesThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesThroughputSetting_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_Spec(generators)
	databaseAccountsSqlDatabasesThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesThroughputSetting_Spec{}), generators)

	return databaseAccountsSqlDatabasesThroughputSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
}

func Test_DatabaseAccountsSqlDatabasesThroughputSetting_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabasesThroughputSetting_STATUS to DatabaseAccountsSqlDatabasesThroughputSetting_STATUS via AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSetting_STATUS & AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSetting_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS, DatabaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS tests if a specific instance of DatabaseAccountsSqlDatabasesThroughputSetting_STATUS can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS(subject DatabaseAccountsSqlDatabasesThroughputSetting_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210515s.DatabaseAccountsSqlDatabasesThroughputSetting_STATUS
	err := copied.AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabasesThroughputSetting_STATUS
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSetting_STATUS(&other)
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

func Test_DatabaseAccountsSqlDatabasesThroughputSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesThroughputSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS, DatabaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesThroughputSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS(subject DatabaseAccountsSqlDatabasesThroughputSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesThroughputSetting_STATUS
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

// Generator of DatabaseAccountsSqlDatabasesThroughputSetting_STATUS instances for property testing - lazily
// instantiated by DatabaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator()
var databaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator returns a generator of DatabaseAccountsSqlDatabasesThroughputSetting_STATUS instances for property testing.
// We first initialize databaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator != nil {
		return databaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS(generators)
	databaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesThroughputSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS(generators)
	databaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesThroughputSetting_STATUS{}), generators)

	return databaseAccountsSqlDatabasesThroughputSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsGetProperties_Resource_STATUSGenerator())
}
