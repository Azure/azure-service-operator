// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_AutoUpgradePolicyResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoUpgradePolicyResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoUpgradePolicyResource, AutoUpgradePolicyResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoUpgradePolicyResource runs a test to see if a specific instance of AutoUpgradePolicyResource round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoUpgradePolicyResource(subject AutoUpgradePolicyResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoUpgradePolicyResource
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

// Generator of AutoUpgradePolicyResource instances for property testing - lazily instantiated by
// AutoUpgradePolicyResourceGenerator()
var autoUpgradePolicyResourceGenerator gopter.Gen

// AutoUpgradePolicyResourceGenerator returns a generator of AutoUpgradePolicyResource instances for property testing.
func AutoUpgradePolicyResourceGenerator() gopter.Gen {
	if autoUpgradePolicyResourceGenerator != nil {
		return autoUpgradePolicyResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource(generators)
	autoUpgradePolicyResourceGenerator = gen.Struct(reflect.TypeOf(AutoUpgradePolicyResource{}), generators)

	return autoUpgradePolicyResourceGenerator
}

// AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource(gens map[string]gopter.Gen) {
	gens["ThroughputPolicy"] = gen.PtrOf(ThroughputPolicyResourceGenerator())
}

func Test_AutoscaleSettingsResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettingsResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsResource, AutoscaleSettingsResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsResource runs a test to see if a specific instance of AutoscaleSettingsResource round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsResource(subject AutoscaleSettingsResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettingsResource
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

// Generator of AutoscaleSettingsResource instances for property testing - lazily instantiated by
// AutoscaleSettingsResourceGenerator()
var autoscaleSettingsResourceGenerator gopter.Gen

// AutoscaleSettingsResourceGenerator returns a generator of AutoscaleSettingsResource instances for property testing.
// We first initialize autoscaleSettingsResourceGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AutoscaleSettingsResourceGenerator() gopter.Gen {
	if autoscaleSettingsResourceGenerator != nil {
		return autoscaleSettingsResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResource(generators)
	autoscaleSettingsResourceGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResource(generators)
	AddRelatedPropertyGeneratorsForAutoscaleSettingsResource(generators)
	autoscaleSettingsResourceGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource{}), generators)

	return autoscaleSettingsResourceGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsResource(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAutoscaleSettingsResource is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoscaleSettingsResource(gens map[string]gopter.Gen) {
	gens["AutoUpgradePolicy"] = gen.PtrOf(AutoUpgradePolicyResourceGenerator())
}

func Test_MongodbDatabaseCollectionThroughputSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabaseCollectionThroughputSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting_Spec, MongodbDatabaseCollectionThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting_Spec runs a test to see if a specific instance of MongodbDatabaseCollectionThroughputSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting_Spec(subject MongodbDatabaseCollectionThroughputSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabaseCollectionThroughputSetting_Spec
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

// Generator of MongodbDatabaseCollectionThroughputSetting_Spec instances for property testing - lazily instantiated by
// MongodbDatabaseCollectionThroughputSetting_SpecGenerator()
var mongodbDatabaseCollectionThroughputSetting_SpecGenerator gopter.Gen

// MongodbDatabaseCollectionThroughputSetting_SpecGenerator returns a generator of MongodbDatabaseCollectionThroughputSetting_Spec instances for property testing.
// We first initialize mongodbDatabaseCollectionThroughputSetting_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongodbDatabaseCollectionThroughputSetting_SpecGenerator() gopter.Gen {
	if mongodbDatabaseCollectionThroughputSetting_SpecGenerator != nil {
		return mongodbDatabaseCollectionThroughputSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec(generators)
	mongodbDatabaseCollectionThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollectionThroughputSetting_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec(generators)
	AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec(generators)
	mongodbDatabaseCollectionThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollectionThroughputSetting_Spec{}), generators)

	return mongodbDatabaseCollectionThroughputSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ThroughputSettingsUpdatePropertiesGenerator())
}

func Test_ThroughputPolicyResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputPolicyResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputPolicyResource, ThroughputPolicyResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputPolicyResource runs a test to see if a specific instance of ThroughputPolicyResource round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputPolicyResource(subject ThroughputPolicyResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputPolicyResource
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

// Generator of ThroughputPolicyResource instances for property testing - lazily instantiated by
// ThroughputPolicyResourceGenerator()
var throughputPolicyResourceGenerator gopter.Gen

// ThroughputPolicyResourceGenerator returns a generator of ThroughputPolicyResource instances for property testing.
func ThroughputPolicyResourceGenerator() gopter.Gen {
	if throughputPolicyResourceGenerator != nil {
		return throughputPolicyResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputPolicyResource(generators)
	throughputPolicyResourceGenerator = gen.Struct(reflect.TypeOf(ThroughputPolicyResource{}), generators)

	return throughputPolicyResourceGenerator
}

// AddIndependentPropertyGeneratorsForThroughputPolicyResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputPolicyResource(gens map[string]gopter.Gen) {
	gens["IncrementPercent"] = gen.PtrOf(gen.Int())
	gens["IsEnabled"] = gen.PtrOf(gen.Bool())
}

func Test_ThroughputSettingsResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsResource, ThroughputSettingsResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsResource runs a test to see if a specific instance of ThroughputSettingsResource round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsResource(subject ThroughputSettingsResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsResource
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

// Generator of ThroughputSettingsResource instances for property testing - lazily instantiated by
// ThroughputSettingsResourceGenerator()
var throughputSettingsResourceGenerator gopter.Gen

// ThroughputSettingsResourceGenerator returns a generator of ThroughputSettingsResource instances for property testing.
// We first initialize throughputSettingsResourceGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ThroughputSettingsResourceGenerator() gopter.Gen {
	if throughputSettingsResourceGenerator != nil {
		return throughputSettingsResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsResource(generators)
	throughputSettingsResourceGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsResource{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsResource(generators)
	AddRelatedPropertyGeneratorsForThroughputSettingsResource(generators)
	throughputSettingsResourceGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsResource{}), generators)

	return throughputSettingsResourceGenerator
}

// AddIndependentPropertyGeneratorsForThroughputSettingsResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputSettingsResource(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForThroughputSettingsResource is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsResource(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsResourceGenerator())
}

func Test_ThroughputSettingsUpdateProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsUpdateProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsUpdateProperties, ThroughputSettingsUpdatePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsUpdateProperties runs a test to see if a specific instance of ThroughputSettingsUpdateProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsUpdateProperties(subject ThroughputSettingsUpdateProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsUpdateProperties
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

// Generator of ThroughputSettingsUpdateProperties instances for property testing - lazily instantiated by
// ThroughputSettingsUpdatePropertiesGenerator()
var throughputSettingsUpdatePropertiesGenerator gopter.Gen

// ThroughputSettingsUpdatePropertiesGenerator returns a generator of ThroughputSettingsUpdateProperties instances for property testing.
func ThroughputSettingsUpdatePropertiesGenerator() gopter.Gen {
	if throughputSettingsUpdatePropertiesGenerator != nil {
		return throughputSettingsUpdatePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForThroughputSettingsUpdateProperties(generators)
	throughputSettingsUpdatePropertiesGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsUpdateProperties{}), generators)

	return throughputSettingsUpdatePropertiesGenerator
}

// AddRelatedPropertyGeneratorsForThroughputSettingsUpdateProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsUpdateProperties(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
}
