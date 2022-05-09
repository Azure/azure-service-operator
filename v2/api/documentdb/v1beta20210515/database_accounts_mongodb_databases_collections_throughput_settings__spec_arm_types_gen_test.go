// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

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

func Test_DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM, DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM(subject DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM
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

// Generator of DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM instances for property testing -
// lazily instantiated by DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARMGenerator()
var databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARMGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARMGenerator returns a generator of DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARMGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARMGenerator != nil {
		return databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM(generators)
	databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM(generators)
	databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM{}), generators)

	return databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ThroughputSettingsUpdatePropertiesARMGenerator())
}

func Test_ThroughputSettingsUpdatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsUpdatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsUpdatePropertiesARM, ThroughputSettingsUpdatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsUpdatePropertiesARM runs a test to see if a specific instance of ThroughputSettingsUpdatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsUpdatePropertiesARM(subject ThroughputSettingsUpdatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsUpdatePropertiesARM
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

// Generator of ThroughputSettingsUpdatePropertiesARM instances for property testing - lazily instantiated by
// ThroughputSettingsUpdatePropertiesARMGenerator()
var throughputSettingsUpdatePropertiesARMGenerator gopter.Gen

// ThroughputSettingsUpdatePropertiesARMGenerator returns a generator of ThroughputSettingsUpdatePropertiesARM instances for property testing.
func ThroughputSettingsUpdatePropertiesARMGenerator() gopter.Gen {
	if throughputSettingsUpdatePropertiesARMGenerator != nil {
		return throughputSettingsUpdatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForThroughputSettingsUpdatePropertiesARM(generators)
	throughputSettingsUpdatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsUpdatePropertiesARM{}), generators)

	return throughputSettingsUpdatePropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForThroughputSettingsUpdatePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsUpdatePropertiesARM(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceARMGenerator())
}

func Test_ThroughputSettingsResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsResourceARM, ThroughputSettingsResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsResourceARM runs a test to see if a specific instance of ThroughputSettingsResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsResourceARM(subject ThroughputSettingsResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsResourceARM
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

// Generator of ThroughputSettingsResourceARM instances for property testing - lazily instantiated by
// ThroughputSettingsResourceARMGenerator()
var throughputSettingsResourceARMGenerator gopter.Gen

// ThroughputSettingsResourceARMGenerator returns a generator of ThroughputSettingsResourceARM instances for property testing.
// We first initialize throughputSettingsResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ThroughputSettingsResourceARMGenerator() gopter.Gen {
	if throughputSettingsResourceARMGenerator != nil {
		return throughputSettingsResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsResourceARM(generators)
	throughputSettingsResourceARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsResourceARM(generators)
	AddRelatedPropertyGeneratorsForThroughputSettingsResourceARM(generators)
	throughputSettingsResourceARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsResourceARM{}), generators)

	return throughputSettingsResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputSettingsResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputSettingsResourceARM(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForThroughputSettingsResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsResourceARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsResourceARMGenerator())
}

func Test_AutoscaleSettingsResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettingsResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsResourceARM, AutoscaleSettingsResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsResourceARM runs a test to see if a specific instance of AutoscaleSettingsResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsResourceARM(subject AutoscaleSettingsResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettingsResourceARM
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

// Generator of AutoscaleSettingsResourceARM instances for property testing - lazily instantiated by
// AutoscaleSettingsResourceARMGenerator()
var autoscaleSettingsResourceARMGenerator gopter.Gen

// AutoscaleSettingsResourceARMGenerator returns a generator of AutoscaleSettingsResourceARM instances for property testing.
// We first initialize autoscaleSettingsResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AutoscaleSettingsResourceARMGenerator() gopter.Gen {
	if autoscaleSettingsResourceARMGenerator != nil {
		return autoscaleSettingsResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceARM(generators)
	autoscaleSettingsResourceARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceARM(generators)
	AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceARM(generators)
	autoscaleSettingsResourceARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResourceARM{}), generators)

	return autoscaleSettingsResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceARM(gens map[string]gopter.Gen) {
	gens["AutoUpgradePolicy"] = gen.PtrOf(AutoUpgradePolicyResourceARMGenerator())
}

func Test_AutoUpgradePolicyResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoUpgradePolicyResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoUpgradePolicyResourceARM, AutoUpgradePolicyResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoUpgradePolicyResourceARM runs a test to see if a specific instance of AutoUpgradePolicyResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoUpgradePolicyResourceARM(subject AutoUpgradePolicyResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoUpgradePolicyResourceARM
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

// Generator of AutoUpgradePolicyResourceARM instances for property testing - lazily instantiated by
// AutoUpgradePolicyResourceARMGenerator()
var autoUpgradePolicyResourceARMGenerator gopter.Gen

// AutoUpgradePolicyResourceARMGenerator returns a generator of AutoUpgradePolicyResourceARM instances for property testing.
func AutoUpgradePolicyResourceARMGenerator() gopter.Gen {
	if autoUpgradePolicyResourceARMGenerator != nil {
		return autoUpgradePolicyResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceARM(generators)
	autoUpgradePolicyResourceARMGenerator = gen.Struct(reflect.TypeOf(AutoUpgradePolicyResourceARM{}), generators)

	return autoUpgradePolicyResourceARMGenerator
}

// AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceARM(gens map[string]gopter.Gen) {
	gens["ThroughputPolicy"] = gen.PtrOf(ThroughputPolicyResourceARMGenerator())
}

func Test_ThroughputPolicyResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputPolicyResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputPolicyResourceARM, ThroughputPolicyResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputPolicyResourceARM runs a test to see if a specific instance of ThroughputPolicyResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputPolicyResourceARM(subject ThroughputPolicyResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputPolicyResourceARM
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

// Generator of ThroughputPolicyResourceARM instances for property testing - lazily instantiated by
// ThroughputPolicyResourceARMGenerator()
var throughputPolicyResourceARMGenerator gopter.Gen

// ThroughputPolicyResourceARMGenerator returns a generator of ThroughputPolicyResourceARM instances for property testing.
func ThroughputPolicyResourceARMGenerator() gopter.Gen {
	if throughputPolicyResourceARMGenerator != nil {
		return throughputPolicyResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputPolicyResourceARM(generators)
	throughputPolicyResourceARMGenerator = gen.Struct(reflect.TypeOf(ThroughputPolicyResourceARM{}), generators)

	return throughputPolicyResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputPolicyResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputPolicyResourceARM(gens map[string]gopter.Gen) {
	gens["IncrementPercent"] = gen.PtrOf(gen.Int())
	gens["IsEnabled"] = gen.PtrOf(gen.Bool())
}
