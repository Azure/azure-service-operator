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

func Test_ChangeFeed_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeed via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeed, ChangeFeedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeed runs a test to see if a specific instance of ChangeFeed round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeed(subject ChangeFeed) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ChangeFeed
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

// Generator of ChangeFeed instances for property testing - lazily instantiated by ChangeFeedGenerator()
var changeFeedGenerator gopter.Gen

// ChangeFeedGenerator returns a generator of ChangeFeed instances for property testing.
func ChangeFeedGenerator() gopter.Gen {
	if changeFeedGenerator != nil {
		return changeFeedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeed(generators)
	changeFeedGenerator = gen.Struct(reflect.TypeOf(ChangeFeed{}), generators)

	return changeFeedGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeed is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeed(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_ChangeFeed_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeed_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeed_STATUS, ChangeFeed_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeed_STATUS runs a test to see if a specific instance of ChangeFeed_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeed_STATUS(subject ChangeFeed_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ChangeFeed_STATUS
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

// Generator of ChangeFeed_STATUS instances for property testing - lazily instantiated by ChangeFeed_STATUSGenerator()
var changeFeed_STATUSGenerator gopter.Gen

// ChangeFeed_STATUSGenerator returns a generator of ChangeFeed_STATUS instances for property testing.
func ChangeFeed_STATUSGenerator() gopter.Gen {
	if changeFeed_STATUSGenerator != nil {
		return changeFeed_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeed_STATUS(generators)
	changeFeed_STATUSGenerator = gen.Struct(reflect.TypeOf(ChangeFeed_STATUS{}), generators)

	return changeFeed_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeed_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeed_STATUS(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_CorsRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRule, CorsRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRule runs a test to see if a specific instance of CorsRule round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRule(subject CorsRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRule
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

// Generator of CorsRule instances for property testing - lazily instantiated by CorsRuleGenerator()
var corsRuleGenerator gopter.Gen

// CorsRuleGenerator returns a generator of CorsRule instances for property testing.
func CorsRuleGenerator() gopter.Gen {
	if corsRuleGenerator != nil {
		return corsRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRule(generators)
	corsRuleGenerator = gen.Struct(reflect.TypeOf(CorsRule{}), generators)

	return corsRuleGenerator
}

// AddIndependentPropertyGeneratorsForCorsRule is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRule(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.PtrOf(gen.Int())
}

func Test_CorsRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRule_STATUS, CorsRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRule_STATUS runs a test to see if a specific instance of CorsRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRule_STATUS(subject CorsRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRule_STATUS
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

// Generator of CorsRule_STATUS instances for property testing - lazily instantiated by CorsRule_STATUSGenerator()
var corsRule_STATUSGenerator gopter.Gen

// CorsRule_STATUSGenerator returns a generator of CorsRule_STATUS instances for property testing.
func CorsRule_STATUSGenerator() gopter.Gen {
	if corsRule_STATUSGenerator != nil {
		return corsRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRule_STATUS(generators)
	corsRule_STATUSGenerator = gen.Struct(reflect.TypeOf(CorsRule_STATUS{}), generators)

	return corsRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCorsRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRule_STATUS(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.PtrOf(gen.Int())
}

func Test_CorsRules_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRules via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRules, CorsRulesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRules runs a test to see if a specific instance of CorsRules round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRules(subject CorsRules) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRules
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

// Generator of CorsRules instances for property testing - lazily instantiated by CorsRulesGenerator()
var corsRulesGenerator gopter.Gen

// CorsRulesGenerator returns a generator of CorsRules instances for property testing.
func CorsRulesGenerator() gopter.Gen {
	if corsRulesGenerator != nil {
		return corsRulesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRules(generators)
	corsRulesGenerator = gen.Struct(reflect.TypeOf(CorsRules{}), generators)

	return corsRulesGenerator
}

// AddRelatedPropertyGeneratorsForCorsRules is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRules(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRuleGenerator())
}

func Test_CorsRules_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRules_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRules_STATUS, CorsRules_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRules_STATUS runs a test to see if a specific instance of CorsRules_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRules_STATUS(subject CorsRules_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRules_STATUS
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

// Generator of CorsRules_STATUS instances for property testing - lazily instantiated by CorsRules_STATUSGenerator()
var corsRules_STATUSGenerator gopter.Gen

// CorsRules_STATUSGenerator returns a generator of CorsRules_STATUS instances for property testing.
func CorsRules_STATUSGenerator() gopter.Gen {
	if corsRules_STATUSGenerator != nil {
		return corsRules_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRules_STATUS(generators)
	corsRules_STATUSGenerator = gen.Struct(reflect.TypeOf(CorsRules_STATUS{}), generators)

	return corsRules_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForCorsRules_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRules_STATUS(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRule_STATUSGenerator())
}

func Test_DeleteRetentionPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicy, DeleteRetentionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicy runs a test to see if a specific instance of DeleteRetentionPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicy(subject DeleteRetentionPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeleteRetentionPolicy
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

// Generator of DeleteRetentionPolicy instances for property testing - lazily instantiated by
// DeleteRetentionPolicyGenerator()
var deleteRetentionPolicyGenerator gopter.Gen

// DeleteRetentionPolicyGenerator returns a generator of DeleteRetentionPolicy instances for property testing.
func DeleteRetentionPolicyGenerator() gopter.Gen {
	if deleteRetentionPolicyGenerator != nil {
		return deleteRetentionPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicy(generators)
	deleteRetentionPolicyGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicy{}), generators)

	return deleteRetentionPolicyGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicy(gens map[string]gopter.Gen) {
	gens["AllowPermanentDelete"] = gen.PtrOf(gen.Bool())
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_DeleteRetentionPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicy_STATUS, DeleteRetentionPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicy_STATUS runs a test to see if a specific instance of DeleteRetentionPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicy_STATUS(subject DeleteRetentionPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeleteRetentionPolicy_STATUS
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

// Generator of DeleteRetentionPolicy_STATUS instances for property testing - lazily instantiated by
// DeleteRetentionPolicy_STATUSGenerator()
var deleteRetentionPolicy_STATUSGenerator gopter.Gen

// DeleteRetentionPolicy_STATUSGenerator returns a generator of DeleteRetentionPolicy_STATUS instances for property testing.
func DeleteRetentionPolicy_STATUSGenerator() gopter.Gen {
	if deleteRetentionPolicy_STATUSGenerator != nil {
		return deleteRetentionPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_STATUS(generators)
	deleteRetentionPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicy_STATUS{}), generators)

	return deleteRetentionPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["AllowPermanentDelete"] = gen.PtrOf(gen.Bool())
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_LastAccessTimeTrackingPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicy, LastAccessTimeTrackingPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicy runs a test to see if a specific instance of LastAccessTimeTrackingPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicy(subject LastAccessTimeTrackingPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LastAccessTimeTrackingPolicy
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

// Generator of LastAccessTimeTrackingPolicy instances for property testing - lazily instantiated by
// LastAccessTimeTrackingPolicyGenerator()
var lastAccessTimeTrackingPolicyGenerator gopter.Gen

// LastAccessTimeTrackingPolicyGenerator returns a generator of LastAccessTimeTrackingPolicy instances for property testing.
func LastAccessTimeTrackingPolicyGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicyGenerator != nil {
		return lastAccessTimeTrackingPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy(generators)
	lastAccessTimeTrackingPolicyGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicy{}), generators)

	return lastAccessTimeTrackingPolicyGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_LastAccessTimeTrackingPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicy_STATUS, LastAccessTimeTrackingPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicy_STATUS runs a test to see if a specific instance of LastAccessTimeTrackingPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicy_STATUS(subject LastAccessTimeTrackingPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LastAccessTimeTrackingPolicy_STATUS
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

// Generator of LastAccessTimeTrackingPolicy_STATUS instances for property testing - lazily instantiated by
// LastAccessTimeTrackingPolicy_STATUSGenerator()
var lastAccessTimeTrackingPolicy_STATUSGenerator gopter.Gen

// LastAccessTimeTrackingPolicy_STATUSGenerator returns a generator of LastAccessTimeTrackingPolicy_STATUS instances for property testing.
func LastAccessTimeTrackingPolicy_STATUSGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicy_STATUSGenerator != nil {
		return lastAccessTimeTrackingPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_STATUS(generators)
	lastAccessTimeTrackingPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicy_STATUS{}), generators)

	return lastAccessTimeTrackingPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_RestorePolicyProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyProperties, RestorePolicyPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyProperties runs a test to see if a specific instance of RestorePolicyProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyProperties(subject RestorePolicyProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RestorePolicyProperties
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

// Generator of RestorePolicyProperties instances for property testing - lazily instantiated by
// RestorePolicyPropertiesGenerator()
var restorePolicyPropertiesGenerator gopter.Gen

// RestorePolicyPropertiesGenerator returns a generator of RestorePolicyProperties instances for property testing.
func RestorePolicyPropertiesGenerator() gopter.Gen {
	if restorePolicyPropertiesGenerator != nil {
		return restorePolicyPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyProperties(generators)
	restorePolicyPropertiesGenerator = gen.Struct(reflect.TypeOf(RestorePolicyProperties{}), generators)

	return restorePolicyPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyProperties(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_RestorePolicyProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyProperties_STATUS, RestorePolicyProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyProperties_STATUS runs a test to see if a specific instance of RestorePolicyProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyProperties_STATUS(subject RestorePolicyProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RestorePolicyProperties_STATUS
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

// Generator of RestorePolicyProperties_STATUS instances for property testing - lazily instantiated by
// RestorePolicyProperties_STATUSGenerator()
var restorePolicyProperties_STATUSGenerator gopter.Gen

// RestorePolicyProperties_STATUSGenerator returns a generator of RestorePolicyProperties_STATUS instances for property testing.
func RestorePolicyProperties_STATUSGenerator() gopter.Gen {
	if restorePolicyProperties_STATUSGenerator != nil {
		return restorePolicyProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyProperties_STATUS(generators)
	restorePolicyProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(RestorePolicyProperties_STATUS{}), generators)

	return restorePolicyProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["LastEnabledTime"] = gen.PtrOf(gen.AlphaString())
	gens["MinRestoreTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_StorageAccountsBlobService_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobService via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobService, StorageAccountsBlobServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobService runs a test to see if a specific instance of StorageAccountsBlobService round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobService(subject StorageAccountsBlobService) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobService
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

// Generator of StorageAccountsBlobService instances for property testing - lazily instantiated by
// StorageAccountsBlobServiceGenerator()
var storageAccountsBlobServiceGenerator gopter.Gen

// StorageAccountsBlobServiceGenerator returns a generator of StorageAccountsBlobService instances for property testing.
func StorageAccountsBlobServiceGenerator() gopter.Gen {
	if storageAccountsBlobServiceGenerator != nil {
		return storageAccountsBlobServiceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobService(generators)
	storageAccountsBlobServiceGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService{}), generators)

	return storageAccountsBlobServiceGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobService is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobService(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccountsBlobService_SpecGenerator()
	gens["Status"] = StorageAccountsBlobService_STATUSGenerator()
}

func Test_StorageAccountsBlobService_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobService_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobService_STATUS, StorageAccountsBlobService_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobService_STATUS runs a test to see if a specific instance of StorageAccountsBlobService_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobService_STATUS(subject StorageAccountsBlobService_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobService_STATUS
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

// Generator of StorageAccountsBlobService_STATUS instances for property testing - lazily instantiated by
// StorageAccountsBlobService_STATUSGenerator()
var storageAccountsBlobService_STATUSGenerator gopter.Gen

// StorageAccountsBlobService_STATUSGenerator returns a generator of StorageAccountsBlobService_STATUS instances for property testing.
// We first initialize storageAccountsBlobService_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobService_STATUSGenerator() gopter.Gen {
	if storageAccountsBlobService_STATUSGenerator != nil {
		return storageAccountsBlobService_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobService_STATUS(generators)
	storageAccountsBlobService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobService_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobService_STATUS(generators)
	storageAccountsBlobService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService_STATUS{}), generators)

	return storageAccountsBlobService_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobService_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobService_STATUS(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobService_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobService_STATUS(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeed_STATUSGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_STATUSGenerator())
	gens["Cors"] = gen.PtrOf(CorsRules_STATUSGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_STATUSGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicy_STATUSGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyProperties_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
}

func Test_StorageAccountsBlobService_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobService_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobService_Spec, StorageAccountsBlobService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobService_Spec runs a test to see if a specific instance of StorageAccountsBlobService_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobService_Spec(subject StorageAccountsBlobService_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobService_Spec
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

// Generator of StorageAccountsBlobService_Spec instances for property testing - lazily instantiated by
// StorageAccountsBlobService_SpecGenerator()
var storageAccountsBlobService_SpecGenerator gopter.Gen

// StorageAccountsBlobService_SpecGenerator returns a generator of StorageAccountsBlobService_Spec instances for property testing.
// We first initialize storageAccountsBlobService_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobService_SpecGenerator() gopter.Gen {
	if storageAccountsBlobService_SpecGenerator != nil {
		return storageAccountsBlobService_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobService_Spec(generators)
	storageAccountsBlobService_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobService_Spec(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobService_Spec(generators)
	storageAccountsBlobService_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService_Spec{}), generators)

	return storageAccountsBlobService_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobService_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobService_Spec(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobService_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobService_Spec(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeedGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyGenerator())
	gens["Cors"] = gen.PtrOf(CorsRulesGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicyGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyPropertiesGenerator())
}
