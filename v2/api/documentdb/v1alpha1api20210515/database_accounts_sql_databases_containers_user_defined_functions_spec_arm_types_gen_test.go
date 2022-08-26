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

<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_containers_user_defined_function__spec_arm_types_gen_test.go
func Test_DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
========
func Test_DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_containers_user_defined_functions_spec_arm_types_gen_test.go
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_containers_user_defined_function__spec_arm_types_gen_test.go
		"Round trip of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM, DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM(subject DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM) string {
========
		"Round trip of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM, DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM(subject DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM) string {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_containers_user_defined_functions_spec_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_containers_user_defined_function__spec_arm_types_gen_test.go
	var actual DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM
========
	var actual DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_containers_user_defined_functions_spec_arm_types_gen_test.go
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

<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_containers_user_defined_function__spec_arm_types_gen_test.go
// Generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM instances for property testing -
// lazily instantiated by DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARMGenerator()
var databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARMGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARMGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARMGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARMGenerator != nil {
		return databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM{}), generators)

	return databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
========
// Generator of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM instances for property testing -
// lazily instantiated by DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARMGenerator()
var databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARMGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARMGenerator returns a generator of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARMGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARMGenerator != nil {
		return databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM(generators)
	databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM(generators)
	databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM{}), generators)

	return databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_containers_user_defined_functions_spec_arm_types_gen_test.go
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_containers_user_defined_function__spec_arm_types_gen_test.go
// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecARM(gens map[string]gopter.Gen) {
========
// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_containers_user_defined_functions_spec_arm_types_gen_test.go
	gens["Properties"] = gen.PtrOf(SqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator())
}

func Test_SqlUserDefinedFunctionCreateUpdatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionCreateUpdatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdatePropertiesARM, SqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdatePropertiesARM runs a test to see if a specific instance of SqlUserDefinedFunctionCreateUpdatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdatePropertiesARM(subject SqlUserDefinedFunctionCreateUpdatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionCreateUpdatePropertiesARM
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

// Generator of SqlUserDefinedFunctionCreateUpdatePropertiesARM instances for property testing - lazily instantiated by
// SqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator()
var sqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator gopter.Gen

// SqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator returns a generator of SqlUserDefinedFunctionCreateUpdatePropertiesARM instances for property testing.
func SqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator() gopter.Gen {
	if sqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator != nil {
		return sqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdatePropertiesARM(generators)
	sqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionCreateUpdatePropertiesARM{}), generators)

	return sqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdatePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdatePropertiesARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsARMGenerator())
	gens["Resource"] = gen.PtrOf(SqlUserDefinedFunctionResourceARMGenerator())
}

func Test_SqlUserDefinedFunctionResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionResourceARM, SqlUserDefinedFunctionResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionResourceARM runs a test to see if a specific instance of SqlUserDefinedFunctionResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionResourceARM(subject SqlUserDefinedFunctionResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionResourceARM
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

// Generator of SqlUserDefinedFunctionResourceARM instances for property testing - lazily instantiated by
// SqlUserDefinedFunctionResourceARMGenerator()
var sqlUserDefinedFunctionResourceARMGenerator gopter.Gen

// SqlUserDefinedFunctionResourceARMGenerator returns a generator of SqlUserDefinedFunctionResourceARM instances for property testing.
func SqlUserDefinedFunctionResourceARMGenerator() gopter.Gen {
	if sqlUserDefinedFunctionResourceARMGenerator != nil {
		return sqlUserDefinedFunctionResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResourceARM(generators)
	sqlUserDefinedFunctionResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionResourceARM{}), generators)

	return sqlUserDefinedFunctionResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResourceARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
