// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

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

func Test_SqlDatabaseContainerThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting, SqlDatabaseContainerThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting runs a test to see if a specific instance of SqlDatabaseContainerThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting(subject SqlDatabaseContainerThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerThroughputSetting
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

// Generator of SqlDatabaseContainerThroughputSetting instances for property testing - lazily instantiated by
// SqlDatabaseContainerThroughputSettingGenerator()
var sqlDatabaseContainerThroughputSettingGenerator gopter.Gen

// SqlDatabaseContainerThroughputSettingGenerator returns a generator of SqlDatabaseContainerThroughputSetting instances for property testing.
func SqlDatabaseContainerThroughputSettingGenerator() gopter.Gen {
	if sqlDatabaseContainerThroughputSettingGenerator != nil {
		return sqlDatabaseContainerThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting(generators)
	sqlDatabaseContainerThroughputSettingGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerThroughputSetting{}), generators)

	return sqlDatabaseContainerThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator()
	gens["Status"] = ThroughputSettingsGetResultsStatusGenerator()
}

func Test_DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec, DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(subject DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
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

// Generator of DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec instances for property testing - lazily
// instantiated by DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator()
var databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator != nil {
		return databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(generators)
	databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(generators)
	databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}), generators)

	return databaseAccountsSqlDatabasesContainersThroughputSettingsSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
}
