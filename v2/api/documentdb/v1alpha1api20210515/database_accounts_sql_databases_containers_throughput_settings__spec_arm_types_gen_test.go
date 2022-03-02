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

func Test_DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARM, DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARM runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARM(subject DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM
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

// Generator of DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM instances for property testing - lazily
//instantiated by DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARMGenerator()
var databaseAccountsSqlDatabasesContainersThroughputSettingsSpecARMGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARMGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersThroughputSettingsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARMGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersThroughputSettingsSpecARMGenerator != nil {
		return databaseAccountsSqlDatabasesContainersThroughputSettingsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARM(generators)
	databaseAccountsSqlDatabasesContainersThroughputSettingsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARM(generators)
	databaseAccountsSqlDatabasesContainersThroughputSettingsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM{}), generators)

	return databaseAccountsSqlDatabasesContainersThroughputSettingsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ThroughputSettingsUpdatePropertiesARMGenerator())
}
