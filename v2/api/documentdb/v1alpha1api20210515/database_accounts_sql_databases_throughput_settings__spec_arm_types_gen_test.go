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

func Test_DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSettingsSpecARM, DatabaseAccountsSqlDatabasesThroughputSettingsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSettingsSpecARM runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSettingsSpecARM(subject DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM
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

// Generator of DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM instances for property testing - lazily
// instantiated by DatabaseAccountsSqlDatabasesThroughputSettingsSpecARMGenerator()
var databaseAccountsSqlDatabasesThroughputSettingsSpecARMGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesThroughputSettingsSpecARMGenerator returns a generator of DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM instances for property testing.
// We first initialize databaseAccountsSqlDatabasesThroughputSettingsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesThroughputSettingsSpecARMGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesThroughputSettingsSpecARMGenerator != nil {
		return databaseAccountsSqlDatabasesThroughputSettingsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpecARM(generators)
	databaseAccountsSqlDatabasesThroughputSettingsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpecARM(generators)
	databaseAccountsSqlDatabasesThroughputSettingsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM{}), generators)

	return databaseAccountsSqlDatabasesThroughputSettingsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ThroughputSettingsUpdatePropertiesARMGenerator())
}
