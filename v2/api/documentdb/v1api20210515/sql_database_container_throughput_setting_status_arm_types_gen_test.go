// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

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

func Test_SqlDatabaseContainerThroughputSetting_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerThroughputSetting_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting_STATUS_ARM, SqlDatabaseContainerThroughputSetting_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting_STATUS_ARM runs a test to see if a specific instance of SqlDatabaseContainerThroughputSetting_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting_STATUS_ARM(subject SqlDatabaseContainerThroughputSetting_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerThroughputSetting_STATUS_ARM
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

// Generator of SqlDatabaseContainerThroughputSetting_STATUS_ARM instances for property testing - lazily instantiated by
// SqlDatabaseContainerThroughputSetting_STATUS_ARMGenerator()
var sqlDatabaseContainerThroughputSetting_STATUS_ARMGenerator gopter.Gen

// SqlDatabaseContainerThroughputSetting_STATUS_ARMGenerator returns a generator of SqlDatabaseContainerThroughputSetting_STATUS_ARM instances for property testing.
// We first initialize sqlDatabaseContainerThroughputSetting_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseContainerThroughputSetting_STATUS_ARMGenerator() gopter.Gen {
	if sqlDatabaseContainerThroughputSetting_STATUS_ARMGenerator != nil {
		return sqlDatabaseContainerThroughputSetting_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS_ARM(generators)
	sqlDatabaseContainerThroughputSetting_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerThroughputSetting_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS_ARM(generators)
	sqlDatabaseContainerThroughputSetting_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerThroughputSetting_STATUS_ARM{}), generators)

	return sqlDatabaseContainerThroughputSetting_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ThroughputSettingsGetProperties_STATUS_ARMGenerator())
}
