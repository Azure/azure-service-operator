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

<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_database__spec_arm_types_gen_test.go
func Test_DatabaseAccountsSqlDatabase_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
========
func Test_DatabaseAccounts_SqlDatabases_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_spec_arm_types_gen_test.go
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_database__spec_arm_types_gen_test.go
		"Round trip of DatabaseAccountsSqlDatabase_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabase_SpecARM, DatabaseAccountsSqlDatabase_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabase_SpecARM runs a test to see if a specific instance of DatabaseAccountsSqlDatabase_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabase_SpecARM(subject DatabaseAccountsSqlDatabase_SpecARM) string {
========
		"Round trip of DatabaseAccounts_SqlDatabases_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_SpecARM, DatabaseAccounts_SqlDatabases_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_SpecARM runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_SpecARM(subject DatabaseAccounts_SqlDatabases_SpecARM) string {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_spec_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_database__spec_arm_types_gen_test.go
	var actual DatabaseAccountsSqlDatabase_SpecARM
========
	var actual DatabaseAccounts_SqlDatabases_SpecARM
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_spec_arm_types_gen_test.go
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

<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_database__spec_arm_types_gen_test.go
// Generator of DatabaseAccountsSqlDatabase_SpecARM instances for property testing - lazily instantiated by
// DatabaseAccountsSqlDatabase_SpecARMGenerator()
var databaseAccountsSqlDatabase_SpecARMGenerator gopter.Gen

// DatabaseAccountsSqlDatabase_SpecARMGenerator returns a generator of DatabaseAccountsSqlDatabase_SpecARM instances for property testing.
// We first initialize databaseAccountsSqlDatabase_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabase_SpecARMGenerator() gopter.Gen {
	if databaseAccountsSqlDatabase_SpecARMGenerator != nil {
		return databaseAccountsSqlDatabase_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabase_SpecARM(generators)
	databaseAccountsSqlDatabase_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabase_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabase_SpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabase_SpecARM(generators)
	databaseAccountsSqlDatabase_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabase_SpecARM{}), generators)

	return databaseAccountsSqlDatabase_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabase_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabase_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
========
// Generator of DatabaseAccounts_SqlDatabases_SpecARM instances for property testing - lazily instantiated by
// DatabaseAccounts_SqlDatabases_SpecARMGenerator()
var databaseAccounts_SqlDatabases_SpecARMGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_SpecARMGenerator returns a generator of DatabaseAccounts_SqlDatabases_SpecARM instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_SpecARMGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_SpecARMGenerator != nil {
		return databaseAccounts_SqlDatabases_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_SpecARM(generators)
	databaseAccounts_SqlDatabases_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_SpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_SpecARM(generators)
	databaseAccounts_SqlDatabases_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_SpecARM{}), generators)

	return databaseAccounts_SqlDatabases_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_SpecARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_spec_arm_types_gen_test.go
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_database__spec_arm_types_gen_test.go
// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabase_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabase_SpecARM(gens map[string]gopter.Gen) {
========
// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_SpecARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_spec_arm_types_gen_test.go
	gens["Properties"] = gen.PtrOf(SqlDatabaseCreateUpdatePropertiesARMGenerator())
}

func Test_SqlDatabaseCreateUpdatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseCreateUpdatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseCreateUpdatePropertiesARM, SqlDatabaseCreateUpdatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseCreateUpdatePropertiesARM runs a test to see if a specific instance of SqlDatabaseCreateUpdatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseCreateUpdatePropertiesARM(subject SqlDatabaseCreateUpdatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseCreateUpdatePropertiesARM
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

// Generator of SqlDatabaseCreateUpdatePropertiesARM instances for property testing - lazily instantiated by
// SqlDatabaseCreateUpdatePropertiesARMGenerator()
var sqlDatabaseCreateUpdatePropertiesARMGenerator gopter.Gen

// SqlDatabaseCreateUpdatePropertiesARMGenerator returns a generator of SqlDatabaseCreateUpdatePropertiesARM instances for property testing.
func SqlDatabaseCreateUpdatePropertiesARMGenerator() gopter.Gen {
	if sqlDatabaseCreateUpdatePropertiesARMGenerator != nil {
		return sqlDatabaseCreateUpdatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseCreateUpdatePropertiesARM(generators)
	sqlDatabaseCreateUpdatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseCreateUpdatePropertiesARM{}), generators)

	return sqlDatabaseCreateUpdatePropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseCreateUpdatePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseCreateUpdatePropertiesARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsARMGenerator())
	gens["Resource"] = gen.PtrOf(SqlDatabaseResourceARMGenerator())
}

func Test_SqlDatabaseResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseResourceARM, SqlDatabaseResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseResourceARM runs a test to see if a specific instance of SqlDatabaseResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseResourceARM(subject SqlDatabaseResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseResourceARM
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

// Generator of SqlDatabaseResourceARM instances for property testing - lazily instantiated by
// SqlDatabaseResourceARMGenerator()
var sqlDatabaseResourceARMGenerator gopter.Gen

// SqlDatabaseResourceARMGenerator returns a generator of SqlDatabaseResourceARM instances for property testing.
func SqlDatabaseResourceARMGenerator() gopter.Gen {
	if sqlDatabaseResourceARMGenerator != nil {
		return sqlDatabaseResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseResourceARM(generators)
	sqlDatabaseResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseResourceARM{}), generators)

	return sqlDatabaseResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseResourceARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
