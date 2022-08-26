// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

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

<<<<<<<< HEAD:v2/api/dbforpostgresql/v1beta20210601/flexible_servers_database__spec_arm_types_gen_test.go
func Test_FlexibleServersDatabase_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
========
func Test_FlexibleServers_Databases_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
>>>>>>>> main:v2/api/dbforpostgresql/v1beta20210601/flexible_servers_databases_spec_arm_types_gen_test.go
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<<< HEAD:v2/api/dbforpostgresql/v1beta20210601/flexible_servers_database__spec_arm_types_gen_test.go
		"Round trip of FlexibleServersDatabase_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersDatabase_SpecARM, FlexibleServersDatabase_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersDatabase_SpecARM runs a test to see if a specific instance of FlexibleServersDatabase_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersDatabase_SpecARM(subject FlexibleServersDatabase_SpecARM) string {
========
		"Round trip of FlexibleServers_Databases_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_Databases_SpecARM, FlexibleServers_Databases_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_Databases_SpecARM runs a test to see if a specific instance of FlexibleServers_Databases_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_Databases_SpecARM(subject FlexibleServers_Databases_SpecARM) string {
>>>>>>>> main:v2/api/dbforpostgresql/v1beta20210601/flexible_servers_databases_spec_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
<<<<<<<< HEAD:v2/api/dbforpostgresql/v1beta20210601/flexible_servers_database__spec_arm_types_gen_test.go
	var actual FlexibleServersDatabase_SpecARM
========
	var actual FlexibleServers_Databases_SpecARM
>>>>>>>> main:v2/api/dbforpostgresql/v1beta20210601/flexible_servers_databases_spec_arm_types_gen_test.go
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

<<<<<<<< HEAD:v2/api/dbforpostgresql/v1beta20210601/flexible_servers_database__spec_arm_types_gen_test.go
// Generator of FlexibleServersDatabase_SpecARM instances for property testing - lazily instantiated by
// FlexibleServersDatabase_SpecARMGenerator()
var flexibleServersDatabase_SpecARMGenerator gopter.Gen

// FlexibleServersDatabase_SpecARMGenerator returns a generator of FlexibleServersDatabase_SpecARM instances for property testing.
// We first initialize flexibleServersDatabase_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersDatabase_SpecARMGenerator() gopter.Gen {
	if flexibleServersDatabase_SpecARMGenerator != nil {
		return flexibleServersDatabase_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersDatabase_SpecARM(generators)
	flexibleServersDatabase_SpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabase_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersDatabase_SpecARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersDatabase_SpecARM(generators)
	flexibleServersDatabase_SpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabase_SpecARM{}), generators)

	return flexibleServersDatabase_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersDatabase_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersDatabase_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
========
// Generator of FlexibleServers_Databases_SpecARM instances for property testing - lazily instantiated by
// FlexibleServers_Databases_SpecARMGenerator()
var flexibleServers_Databases_SpecARMGenerator gopter.Gen

// FlexibleServers_Databases_SpecARMGenerator returns a generator of FlexibleServers_Databases_SpecARM instances for property testing.
// We first initialize flexibleServers_Databases_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServers_Databases_SpecARMGenerator() gopter.Gen {
	if flexibleServers_Databases_SpecARMGenerator != nil {
		return flexibleServers_Databases_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Databases_SpecARM(generators)
	flexibleServers_Databases_SpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Databases_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Databases_SpecARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServers_Databases_SpecARM(generators)
	flexibleServers_Databases_SpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Databases_SpecARM{}), generators)

	return flexibleServers_Databases_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_Databases_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_Databases_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
>>>>>>>> main:v2/api/dbforpostgresql/v1beta20210601/flexible_servers_databases_spec_arm_types_gen_test.go
	gens["Name"] = gen.AlphaString()
}

<<<<<<<< HEAD:v2/api/dbforpostgresql/v1beta20210601/flexible_servers_database__spec_arm_types_gen_test.go
// AddRelatedPropertyGeneratorsForFlexibleServersDatabase_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersDatabase_SpecARM(gens map[string]gopter.Gen) {
========
// AddRelatedPropertyGeneratorsForFlexibleServers_Databases_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServers_Databases_SpecARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/dbforpostgresql/v1beta20210601/flexible_servers_databases_spec_arm_types_gen_test.go
	gens["Properties"] = gen.PtrOf(DatabasePropertiesARMGenerator())
}

func Test_DatabasePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabasePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabasePropertiesARM, DatabasePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabasePropertiesARM runs a test to see if a specific instance of DatabasePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabasePropertiesARM(subject DatabasePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabasePropertiesARM
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

// Generator of DatabasePropertiesARM instances for property testing - lazily instantiated by
// DatabasePropertiesARMGenerator()
var databasePropertiesARMGenerator gopter.Gen

// DatabasePropertiesARMGenerator returns a generator of DatabasePropertiesARM instances for property testing.
func DatabasePropertiesARMGenerator() gopter.Gen {
	if databasePropertiesARMGenerator != nil {
		return databasePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabasePropertiesARM(generators)
	databasePropertiesARMGenerator = gen.Struct(reflect.TypeOf(DatabasePropertiesARM{}), generators)

	return databasePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabasePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabasePropertiesARM(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
}
