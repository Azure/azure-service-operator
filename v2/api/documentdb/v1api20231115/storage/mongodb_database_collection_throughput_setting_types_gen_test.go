// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_AutoUpgradePolicyResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoUpgradePolicyResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoUpgradePolicyResource_STATUS, AutoUpgradePolicyResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoUpgradePolicyResource_STATUS runs a test to see if a specific instance of AutoUpgradePolicyResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoUpgradePolicyResource_STATUS(subject AutoUpgradePolicyResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoUpgradePolicyResource_STATUS
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

// Generator of AutoUpgradePolicyResource_STATUS instances for property testing - lazily instantiated by
// AutoUpgradePolicyResource_STATUSGenerator()
var autoUpgradePolicyResource_STATUSGenerator gopter.Gen

// AutoUpgradePolicyResource_STATUSGenerator returns a generator of AutoUpgradePolicyResource_STATUS instances for property testing.
func AutoUpgradePolicyResource_STATUSGenerator() gopter.Gen {
	if autoUpgradePolicyResource_STATUSGenerator != nil {
		return autoUpgradePolicyResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource_STATUS(generators)
	autoUpgradePolicyResource_STATUSGenerator = gen.Struct(reflect.TypeOf(AutoUpgradePolicyResource_STATUS{}), generators)

	return autoUpgradePolicyResource_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource_STATUS(gens map[string]gopter.Gen) {
	gens["ThroughputPolicy"] = gen.PtrOf(ThroughputPolicyResource_STATUSGenerator())
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

func Test_AutoscaleSettingsResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettingsResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsResource_STATUS, AutoscaleSettingsResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsResource_STATUS runs a test to see if a specific instance of AutoscaleSettingsResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsResource_STATUS(subject AutoscaleSettingsResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettingsResource_STATUS
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

// Generator of AutoscaleSettingsResource_STATUS instances for property testing - lazily instantiated by
// AutoscaleSettingsResource_STATUSGenerator()
var autoscaleSettingsResource_STATUSGenerator gopter.Gen

// AutoscaleSettingsResource_STATUSGenerator returns a generator of AutoscaleSettingsResource_STATUS instances for property testing.
// We first initialize autoscaleSettingsResource_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AutoscaleSettingsResource_STATUSGenerator() gopter.Gen {
	if autoscaleSettingsResource_STATUSGenerator != nil {
		return autoscaleSettingsResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_STATUS(generators)
	autoscaleSettingsResource_STATUSGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_STATUS(generators)
	AddRelatedPropertyGeneratorsForAutoscaleSettingsResource_STATUS(generators)
	autoscaleSettingsResource_STATUSGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource_STATUS{}), generators)

	return autoscaleSettingsResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_STATUS(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
	gens["TargetMaxThroughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAutoscaleSettingsResource_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoscaleSettingsResource_STATUS(gens map[string]gopter.Gen) {
	gens["AutoUpgradePolicy"] = gen.PtrOf(AutoUpgradePolicyResource_STATUSGenerator())
}

func Test_MongodbDatabaseCollectionThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabaseCollectionThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting, MongodbDatabaseCollectionThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting runs a test to see if a specific instance of MongodbDatabaseCollectionThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting(subject MongodbDatabaseCollectionThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabaseCollectionThroughputSetting
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

// Generator of MongodbDatabaseCollectionThroughputSetting instances for property testing - lazily instantiated by
// MongodbDatabaseCollectionThroughputSettingGenerator()
var mongodbDatabaseCollectionThroughputSettingGenerator gopter.Gen

// MongodbDatabaseCollectionThroughputSettingGenerator returns a generator of MongodbDatabaseCollectionThroughputSetting instances for property testing.
func MongodbDatabaseCollectionThroughputSettingGenerator() gopter.Gen {
	if mongodbDatabaseCollectionThroughputSettingGenerator != nil {
		return mongodbDatabaseCollectionThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting(generators)
	mongodbDatabaseCollectionThroughputSettingGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollectionThroughputSetting{}), generators)

	return mongodbDatabaseCollectionThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = MongodbDatabaseCollectionThroughputSetting_SpecGenerator()
	gens["Status"] = MongodbDatabaseCollectionThroughputSetting_STATUSGenerator()
}

func Test_MongodbDatabaseCollectionThroughputSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabaseCollectionThroughputSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting_STATUS, MongodbDatabaseCollectionThroughputSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting_STATUS runs a test to see if a specific instance of MongodbDatabaseCollectionThroughputSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting_STATUS(subject MongodbDatabaseCollectionThroughputSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabaseCollectionThroughputSetting_STATUS
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

// Generator of MongodbDatabaseCollectionThroughputSetting_STATUS instances for property testing - lazily instantiated
// by MongodbDatabaseCollectionThroughputSetting_STATUSGenerator()
var mongodbDatabaseCollectionThroughputSetting_STATUSGenerator gopter.Gen

// MongodbDatabaseCollectionThroughputSetting_STATUSGenerator returns a generator of MongodbDatabaseCollectionThroughputSetting_STATUS instances for property testing.
// We first initialize mongodbDatabaseCollectionThroughputSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongodbDatabaseCollectionThroughputSetting_STATUSGenerator() gopter.Gen {
	if mongodbDatabaseCollectionThroughputSetting_STATUSGenerator != nil {
		return mongodbDatabaseCollectionThroughputSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_STATUS(generators)
	mongodbDatabaseCollectionThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollectionThroughputSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_STATUS(generators)
	mongodbDatabaseCollectionThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollectionThroughputSetting_STATUS{}), generators)

	return mongodbDatabaseCollectionThroughputSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsGetProperties_Resource_STATUSGenerator())
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
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
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

func Test_ThroughputPolicyResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputPolicyResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputPolicyResource_STATUS, ThroughputPolicyResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputPolicyResource_STATUS runs a test to see if a specific instance of ThroughputPolicyResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputPolicyResource_STATUS(subject ThroughputPolicyResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputPolicyResource_STATUS
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

// Generator of ThroughputPolicyResource_STATUS instances for property testing - lazily instantiated by
// ThroughputPolicyResource_STATUSGenerator()
var throughputPolicyResource_STATUSGenerator gopter.Gen

// ThroughputPolicyResource_STATUSGenerator returns a generator of ThroughputPolicyResource_STATUS instances for property testing.
func ThroughputPolicyResource_STATUSGenerator() gopter.Gen {
	if throughputPolicyResource_STATUSGenerator != nil {
		return throughputPolicyResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputPolicyResource_STATUS(generators)
	throughputPolicyResource_STATUSGenerator = gen.Struct(reflect.TypeOf(ThroughputPolicyResource_STATUS{}), generators)

	return throughputPolicyResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForThroughputPolicyResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputPolicyResource_STATUS(gens map[string]gopter.Gen) {
	gens["IncrementPercent"] = gen.PtrOf(gen.Int())
	gens["IsEnabled"] = gen.PtrOf(gen.Bool())
}

func Test_ThroughputSettingsGetProperties_Resource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsGetProperties_Resource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsGetProperties_Resource_STATUS, ThroughputSettingsGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsGetProperties_Resource_STATUS runs a test to see if a specific instance of ThroughputSettingsGetProperties_Resource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsGetProperties_Resource_STATUS(subject ThroughputSettingsGetProperties_Resource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsGetProperties_Resource_STATUS
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

// Generator of ThroughputSettingsGetProperties_Resource_STATUS instances for property testing - lazily instantiated by
// ThroughputSettingsGetProperties_Resource_STATUSGenerator()
var throughputSettingsGetProperties_Resource_STATUSGenerator gopter.Gen

// ThroughputSettingsGetProperties_Resource_STATUSGenerator returns a generator of ThroughputSettingsGetProperties_Resource_STATUS instances for property testing.
// We first initialize throughputSettingsGetProperties_Resource_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ThroughputSettingsGetProperties_Resource_STATUSGenerator() gopter.Gen {
	if throughputSettingsGetProperties_Resource_STATUSGenerator != nil {
		return throughputSettingsGetProperties_Resource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUS(generators)
	throughputSettingsGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsGetProperties_Resource_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUS(generators)
	AddRelatedPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUS(generators)
	throughputSettingsGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsGetProperties_Resource_STATUS{}), generators)

	return throughputSettingsGetProperties_Resource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["InstantMaximumThroughput"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumThroughput"] = gen.PtrOf(gen.AlphaString())
	gens["OfferReplacePending"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["SoftAllowedMaximumThroughput"] = gen.PtrOf(gen.AlphaString())
	gens["Throughput"] = gen.PtrOf(gen.Int())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

// AddRelatedPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsResource_STATUSGenerator())
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
