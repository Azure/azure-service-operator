// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220120preview

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

func Test_FlexibleServers_Database_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_Database_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_Database_STATUS_ARM, FlexibleServers_Database_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_Database_STATUS_ARM runs a test to see if a specific instance of FlexibleServers_Database_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_Database_STATUS_ARM(subject FlexibleServers_Database_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_Database_STATUS_ARM
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

// Generator of FlexibleServers_Database_STATUS_ARM instances for property testing - lazily instantiated by
// FlexibleServers_Database_STATUS_ARMGenerator()
var flexibleServers_Database_STATUS_ARMGenerator gopter.Gen

// FlexibleServers_Database_STATUS_ARMGenerator returns a generator of FlexibleServers_Database_STATUS_ARM instances for property testing.
// We first initialize flexibleServers_Database_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServers_Database_STATUS_ARMGenerator() gopter.Gen {
	if flexibleServers_Database_STATUS_ARMGenerator != nil {
		return flexibleServers_Database_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Database_STATUS_ARM(generators)
	flexibleServers_Database_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Database_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Database_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServers_Database_STATUS_ARM(generators)
	flexibleServers_Database_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Database_STATUS_ARM{}), generators)

	return flexibleServers_Database_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_Database_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_Database_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServers_Database_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServers_Database_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabaseProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_DatabaseProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseProperties_STATUS_ARM, DatabaseProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseProperties_STATUS_ARM runs a test to see if a specific instance of DatabaseProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseProperties_STATUS_ARM(subject DatabaseProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_STATUS_ARM
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

// Generator of DatabaseProperties_STATUS_ARM instances for property testing - lazily instantiated by
// DatabaseProperties_STATUS_ARMGenerator()
var databaseProperties_STATUS_ARMGenerator gopter.Gen

// DatabaseProperties_STATUS_ARMGenerator returns a generator of DatabaseProperties_STATUS_ARM instances for property testing.
func DatabaseProperties_STATUS_ARMGenerator() gopter.Gen {
	if databaseProperties_STATUS_ARMGenerator != nil {
		return databaseProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_STATUS_ARM(generators)
	databaseProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_STATUS_ARM{}), generators)

	return databaseProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
}
