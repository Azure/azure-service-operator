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

func Test_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM, DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM(subject DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM
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

// Generator of DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM instances for property testing - lazily
// instantiated by DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARMGenerator()
var databaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARMGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARMGenerator returns a generator of DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARMGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARMGenerator != nil {
		return databaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM(generators)
	databaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM(generators)
	databaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM{}), generators)

	return databaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ThroughputSettingsGetProperties_STATUS_ARMGenerator())
}
