// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

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

func Test_DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTriggersSpecARM, DatabaseAccountsSqlDatabasesContainersTriggersSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTriggersSpecARM runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersTriggersSpecARM(subject DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM
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

// Generator of DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM instances for property testing - lazily
//instantiated by DatabaseAccountsSqlDatabasesContainersTriggersSpecARMGenerator()
var databaseAccountsSqlDatabasesContainersTriggersSpecARMGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersTriggersSpecARMGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersTriggersSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersTriggersSpecARMGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersTriggersSpecARMGenerator != nil {
		return databaseAccountsSqlDatabasesContainersTriggersSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpecARM(generators)
	databaseAccountsSqlDatabasesContainersTriggersSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpecARM(generators)
	databaseAccountsSqlDatabasesContainersTriggersSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM{}), generators)

	return databaseAccountsSqlDatabasesContainersTriggersSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpecARM(gens map[string]gopter.Gen) {
	gens["APIVersion"] = gen.OneConstOf(DatabaseAccountsSqlDatabasesContainersTriggersSpecAPIVersion20210515)
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.OneConstOf(DatabaseAccountsSqlDatabasesContainersTriggersSpecTypeMicrosoftDocumentDBDatabaseAccountsSqlDatabasesContainersTriggers)
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersTriggersSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = SqlTriggerCreateUpdatePropertiesARMGenerator()
}

func Test_SqlTriggerCreateUpdatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerCreateUpdatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerCreateUpdatePropertiesARM, SqlTriggerCreateUpdatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerCreateUpdatePropertiesARM runs a test to see if a specific instance of SqlTriggerCreateUpdatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerCreateUpdatePropertiesARM(subject SqlTriggerCreateUpdatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerCreateUpdatePropertiesARM
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

// Generator of SqlTriggerCreateUpdatePropertiesARM instances for property testing - lazily instantiated by
//SqlTriggerCreateUpdatePropertiesARMGenerator()
var sqlTriggerCreateUpdatePropertiesARMGenerator gopter.Gen

// SqlTriggerCreateUpdatePropertiesARMGenerator returns a generator of SqlTriggerCreateUpdatePropertiesARM instances for property testing.
func SqlTriggerCreateUpdatePropertiesARMGenerator() gopter.Gen {
	if sqlTriggerCreateUpdatePropertiesARMGenerator != nil {
		return sqlTriggerCreateUpdatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlTriggerCreateUpdatePropertiesARM(generators)
	sqlTriggerCreateUpdatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerCreateUpdatePropertiesARM{}), generators)

	return sqlTriggerCreateUpdatePropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlTriggerCreateUpdatePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTriggerCreateUpdatePropertiesARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsARMGenerator())
	gens["Resource"] = SqlTriggerResourceARMGenerator()
}

func Test_SqlTriggerResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerResourceARM, SqlTriggerResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerResourceARM runs a test to see if a specific instance of SqlTriggerResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerResourceARM(subject SqlTriggerResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerResourceARM
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

// Generator of SqlTriggerResourceARM instances for property testing - lazily instantiated by
//SqlTriggerResourceARMGenerator()
var sqlTriggerResourceARMGenerator gopter.Gen

// SqlTriggerResourceARMGenerator returns a generator of SqlTriggerResourceARM instances for property testing.
func SqlTriggerResourceARMGenerator() gopter.Gen {
	if sqlTriggerResourceARMGenerator != nil {
		return sqlTriggerResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerResourceARM(generators)
	sqlTriggerResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerResourceARM{}), generators)

	return sqlTriggerResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerResourceARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.AlphaString()
	gens["TriggerOperation"] = gen.PtrOf(gen.OneConstOf(SqlTriggerResourceTriggerOperationAll, SqlTriggerResourceTriggerOperationCreate, SqlTriggerResourceTriggerOperationDelete, SqlTriggerResourceTriggerOperationReplace, SqlTriggerResourceTriggerOperationUpdate))
	gens["TriggerType"] = gen.PtrOf(gen.OneConstOf(SqlTriggerResourceTriggerTypePost, SqlTriggerResourceTriggerTypePre))
}
