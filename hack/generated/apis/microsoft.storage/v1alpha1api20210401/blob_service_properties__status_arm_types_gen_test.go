// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

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

func Test_BlobServiceProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobServiceProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobServicePropertiesStatusARM, BlobServicePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobServicePropertiesStatusARM runs a test to see if a specific instance of BlobServiceProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobServicePropertiesStatusARM(subject BlobServiceProperties_StatusARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual BlobServiceProperties_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of BlobServiceProperties_StatusARM instances for property testing -
//lazily instantiated by BlobServicePropertiesStatusARMGenerator()
var blobServicePropertiesStatusARMGenerator gopter.Gen

// BlobServicePropertiesStatusARMGenerator returns a generator of BlobServiceProperties_StatusARM instances for property testing.
// We first initialize blobServicePropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobServicePropertiesStatusARMGenerator() gopter.Gen {
	if blobServicePropertiesStatusARMGenerator != nil {
		return blobServicePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServicePropertiesStatusARM(generators)
	blobServicePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServicePropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForBlobServicePropertiesStatusARM(generators)
	blobServicePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_StatusARM{}), generators)

	return blobServicePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForBlobServicePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobServicePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBlobServicePropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobServicePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(BlobServicePropertiesStatusPropertiesARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuStatusARMGenerator())
}

func Test_BlobServiceProperties_Status_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobServiceProperties_Status_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobServicePropertiesStatusPropertiesARM, BlobServicePropertiesStatusPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobServicePropertiesStatusPropertiesARM runs a test to see if a specific instance of BlobServiceProperties_Status_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobServicePropertiesStatusPropertiesARM(subject BlobServiceProperties_Status_PropertiesARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual BlobServiceProperties_Status_PropertiesARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of BlobServiceProperties_Status_PropertiesARM instances for property
//testing - lazily instantiated by
//BlobServicePropertiesStatusPropertiesARMGenerator()
var blobServicePropertiesStatusPropertiesARMGenerator gopter.Gen

// BlobServicePropertiesStatusPropertiesARMGenerator returns a generator of BlobServiceProperties_Status_PropertiesARM instances for property testing.
// We first initialize blobServicePropertiesStatusPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobServicePropertiesStatusPropertiesARMGenerator() gopter.Gen {
	if blobServicePropertiesStatusPropertiesARMGenerator != nil {
		return blobServicePropertiesStatusPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServicePropertiesStatusPropertiesARM(generators)
	blobServicePropertiesStatusPropertiesARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_Status_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServicePropertiesStatusPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForBlobServicePropertiesStatusPropertiesARM(generators)
	blobServicePropertiesStatusPropertiesARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_Status_PropertiesARM{}), generators)

	return blobServicePropertiesStatusPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForBlobServicePropertiesStatusPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobServicePropertiesStatusPropertiesARM(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForBlobServicePropertiesStatusPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobServicePropertiesStatusPropertiesARM(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeedStatusARMGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyStatusARMGenerator())
	gens["Cors"] = gen.PtrOf(CorsRulesStatusARMGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyStatusARMGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicyStatusARMGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyPropertiesStatusARMGenerator())
}

func Test_Sku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuStatusARM, SkuStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuStatusARM runs a test to see if a specific instance of Sku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuStatusARM(subject Sku_StatusARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual Sku_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of Sku_StatusARM instances for property testing - lazily instantiated
//by SkuStatusARMGenerator()
var skuStatusARMGenerator gopter.Gen

// SkuStatusARMGenerator returns a generator of Sku_StatusARM instances for property testing.
func SkuStatusARMGenerator() gopter.Gen {
	if skuStatusARMGenerator != nil {
		return skuStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuStatusARM(generators)
	skuStatusARMGenerator = gen.Struct(reflect.TypeOf(Sku_StatusARM{}), generators)

	return skuStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuStatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.OneConstOf(SkuName_StatusPremiumLRS, SkuName_StatusPremiumZRS, SkuName_StatusStandardGRS, SkuName_StatusStandardGZRS, SkuName_StatusStandardLRS, SkuName_StatusStandardRAGRS, SkuName_StatusStandardRAGZRS, SkuName_StatusStandardZRS)
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Tier_StatusPremium, Tier_StatusStandard))
}

func Test_ChangeFeed_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeed_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeedStatusARM, ChangeFeedStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeedStatusARM runs a test to see if a specific instance of ChangeFeed_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeedStatusARM(subject ChangeFeed_StatusARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual ChangeFeed_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of ChangeFeed_StatusARM instances for property testing - lazily
//instantiated by ChangeFeedStatusARMGenerator()
var changeFeedStatusARMGenerator gopter.Gen

// ChangeFeedStatusARMGenerator returns a generator of ChangeFeed_StatusARM instances for property testing.
func ChangeFeedStatusARMGenerator() gopter.Gen {
	if changeFeedStatusARMGenerator != nil {
		return changeFeedStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeedStatusARM(generators)
	changeFeedStatusARMGenerator = gen.Struct(reflect.TypeOf(ChangeFeed_StatusARM{}), generators)

	return changeFeedStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeedStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeedStatusARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_CorsRules_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRules_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRulesStatusARM, CorsRulesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRulesStatusARM runs a test to see if a specific instance of CorsRules_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRulesStatusARM(subject CorsRules_StatusARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual CorsRules_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of CorsRules_StatusARM instances for property testing - lazily
//instantiated by CorsRulesStatusARMGenerator()
var corsRulesStatusARMGenerator gopter.Gen

// CorsRulesStatusARMGenerator returns a generator of CorsRules_StatusARM instances for property testing.
func CorsRulesStatusARMGenerator() gopter.Gen {
	if corsRulesStatusARMGenerator != nil {
		return corsRulesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRulesStatusARM(generators)
	corsRulesStatusARMGenerator = gen.Struct(reflect.TypeOf(CorsRules_StatusARM{}), generators)

	return corsRulesStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForCorsRulesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRulesStatusARM(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRuleStatusARMGenerator())
}

func Test_DeleteRetentionPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicyStatusARM, DeleteRetentionPolicyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicyStatusARM runs a test to see if a specific instance of DeleteRetentionPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicyStatusARM(subject DeleteRetentionPolicy_StatusARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual DeleteRetentionPolicy_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of DeleteRetentionPolicy_StatusARM instances for property testing -
//lazily instantiated by DeleteRetentionPolicyStatusARMGenerator()
var deleteRetentionPolicyStatusARMGenerator gopter.Gen

// DeleteRetentionPolicyStatusARMGenerator returns a generator of DeleteRetentionPolicy_StatusARM instances for property testing.
func DeleteRetentionPolicyStatusARMGenerator() gopter.Gen {
	if deleteRetentionPolicyStatusARMGenerator != nil {
		return deleteRetentionPolicyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicyStatusARM(generators)
	deleteRetentionPolicyStatusARMGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicy_StatusARM{}), generators)

	return deleteRetentionPolicyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicyStatusARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_LastAccessTimeTrackingPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicyStatusARM, LastAccessTimeTrackingPolicyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicyStatusARM runs a test to see if a specific instance of LastAccessTimeTrackingPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicyStatusARM(subject LastAccessTimeTrackingPolicy_StatusARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual LastAccessTimeTrackingPolicy_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of LastAccessTimeTrackingPolicy_StatusARM instances for property
//testing - lazily instantiated by LastAccessTimeTrackingPolicyStatusARMGenerator()
var lastAccessTimeTrackingPolicyStatusARMGenerator gopter.Gen

// LastAccessTimeTrackingPolicyStatusARMGenerator returns a generator of LastAccessTimeTrackingPolicy_StatusARM instances for property testing.
func LastAccessTimeTrackingPolicyStatusARMGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicyStatusARMGenerator != nil {
		return lastAccessTimeTrackingPolicyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicyStatusARM(generators)
	lastAccessTimeTrackingPolicyStatusARMGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicy_StatusARM{}), generators)

	return lastAccessTimeTrackingPolicyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicyStatusARM(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.Bool()
	gens["Name"] = gen.PtrOf(gen.OneConstOf(LastAccessTimeTrackingPolicyStatusNameAccessTimeTracking))
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_RestorePolicyProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyPropertiesStatusARM, RestorePolicyPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyPropertiesStatusARM runs a test to see if a specific instance of RestorePolicyProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyPropertiesStatusARM(subject RestorePolicyProperties_StatusARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual RestorePolicyProperties_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of RestorePolicyProperties_StatusARM instances for property testing -
//lazily instantiated by RestorePolicyPropertiesStatusARMGenerator()
var restorePolicyPropertiesStatusARMGenerator gopter.Gen

// RestorePolicyPropertiesStatusARMGenerator returns a generator of RestorePolicyProperties_StatusARM instances for property testing.
func RestorePolicyPropertiesStatusARMGenerator() gopter.Gen {
	if restorePolicyPropertiesStatusARMGenerator != nil {
		return restorePolicyPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyPropertiesStatusARM(generators)
	restorePolicyPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(RestorePolicyProperties_StatusARM{}), generators)

	return restorePolicyPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.Bool()
	gens["LastEnabledTime"] = gen.PtrOf(gen.AlphaString())
	gens["MinRestoreTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_CorsRule_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRule_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRuleStatusARM, CorsRuleStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRuleStatusARM runs a test to see if a specific instance of CorsRule_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRuleStatusARM(subject CorsRule_StatusARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual CorsRule_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of CorsRule_StatusARM instances for property testing - lazily
//instantiated by CorsRuleStatusARMGenerator()
var corsRuleStatusARMGenerator gopter.Gen

// CorsRuleStatusARMGenerator returns a generator of CorsRule_StatusARM instances for property testing.
func CorsRuleStatusARMGenerator() gopter.Gen {
	if corsRuleStatusARMGenerator != nil {
		return corsRuleStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRuleStatusARM(generators)
	corsRuleStatusARMGenerator = gen.Struct(reflect.TypeOf(CorsRule_StatusARM{}), generators)

	return corsRuleStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForCorsRuleStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRuleStatusARM(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.OneConstOf(CorsRuleStatusAllowedMethodsDELETE, CorsRuleStatusAllowedMethodsGET, CorsRuleStatusAllowedMethodsHEAD, CorsRuleStatusAllowedMethodsMERGE, CorsRuleStatusAllowedMethodsOPTIONS, CorsRuleStatusAllowedMethodsPOST, CorsRuleStatusAllowedMethodsPUT))
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.Int()
}
