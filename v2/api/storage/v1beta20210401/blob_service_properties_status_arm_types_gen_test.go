// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

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

func Test_BlobServiceProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobServiceProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobServiceProperties_STATUS_ARM, BlobServiceProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobServiceProperties_STATUS_ARM runs a test to see if a specific instance of BlobServiceProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobServiceProperties_STATUS_ARM(subject BlobServiceProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobServiceProperties_STATUS_ARM
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

// Generator of BlobServiceProperties_STATUS_ARM instances for property testing - lazily instantiated by
// BlobServiceProperties_STATUS_ARMGenerator()
var blobServiceProperties_STATUS_ARMGenerator gopter.Gen

// BlobServiceProperties_STATUS_ARMGenerator returns a generator of BlobServiceProperties_STATUS_ARM instances for property testing.
// We first initialize blobServiceProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobServiceProperties_STATUS_ARMGenerator() gopter.Gen {
	if blobServiceProperties_STATUS_ARMGenerator != nil {
		return blobServiceProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUS_ARM(generators)
	blobServiceProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForBlobServiceProperties_STATUS_ARM(generators)
	blobServiceProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_STATUS_ARM{}), generators)

	return blobServiceProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBlobServiceProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobServiceProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(BlobServiceProperties_Properties_STATUS_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUS_ARMGenerator())
}

func Test_BlobServiceProperties_Properties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobServiceProperties_Properties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobServiceProperties_Properties_STATUS_ARM, BlobServiceProperties_Properties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobServiceProperties_Properties_STATUS_ARM runs a test to see if a specific instance of BlobServiceProperties_Properties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobServiceProperties_Properties_STATUS_ARM(subject BlobServiceProperties_Properties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobServiceProperties_Properties_STATUS_ARM
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

// Generator of BlobServiceProperties_Properties_STATUS_ARM instances for property testing - lazily instantiated by
// BlobServiceProperties_Properties_STATUS_ARMGenerator()
var blobServiceProperties_Properties_STATUS_ARMGenerator gopter.Gen

// BlobServiceProperties_Properties_STATUS_ARMGenerator returns a generator of BlobServiceProperties_Properties_STATUS_ARM instances for property testing.
// We first initialize blobServiceProperties_Properties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobServiceProperties_Properties_STATUS_ARMGenerator() gopter.Gen {
	if blobServiceProperties_Properties_STATUS_ARMGenerator != nil {
		return blobServiceProperties_Properties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_Properties_STATUS_ARM(generators)
	blobServiceProperties_Properties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_Properties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_Properties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForBlobServiceProperties_Properties_STATUS_ARM(generators)
	blobServiceProperties_Properties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_Properties_STATUS_ARM{}), generators)

	return blobServiceProperties_Properties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBlobServiceProperties_Properties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobServiceProperties_Properties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForBlobServiceProperties_Properties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobServiceProperties_Properties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeed_STATUS_ARMGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_STATUS_ARMGenerator())
	gens["Cors"] = gen.PtrOf(CorsRules_STATUS_ARMGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_STATUS_ARMGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicy_STATUS_ARMGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyProperties_STATUS_ARMGenerator())
}

func Test_Sku_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS_ARM, Sku_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS_ARM runs a test to see if a specific instance of Sku_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS_ARM(subject Sku_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS_ARM
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

// Generator of Sku_STATUS_ARM instances for property testing - lazily instantiated by Sku_STATUS_ARMGenerator()
var sku_STATUS_ARMGenerator gopter.Gen

// Sku_STATUS_ARMGenerator returns a generator of Sku_STATUS_ARM instances for property testing.
func Sku_STATUS_ARMGenerator() gopter.Gen {
	if sku_STATUS_ARMGenerator != nil {
		return sku_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS_ARM(generators)
	sku_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS_ARM{}), generators)

	return sku_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		SkuName_STATUS_Premium_LRS,
		SkuName_STATUS_Premium_ZRS,
		SkuName_STATUS_Standard_GRS,
		SkuName_STATUS_Standard_GZRS,
		SkuName_STATUS_Standard_LRS,
		SkuName_STATUS_Standard_RAGRS,
		SkuName_STATUS_Standard_RAGZRS,
		SkuName_STATUS_Standard_ZRS))
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Tier_STATUS_Premium, Tier_STATUS_Standard))
}

func Test_ChangeFeed_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeed_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeed_STATUS_ARM, ChangeFeed_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeed_STATUS_ARM runs a test to see if a specific instance of ChangeFeed_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeed_STATUS_ARM(subject ChangeFeed_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ChangeFeed_STATUS_ARM
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

// Generator of ChangeFeed_STATUS_ARM instances for property testing - lazily instantiated by
// ChangeFeed_STATUS_ARMGenerator()
var changeFeed_STATUS_ARMGenerator gopter.Gen

// ChangeFeed_STATUS_ARMGenerator returns a generator of ChangeFeed_STATUS_ARM instances for property testing.
func ChangeFeed_STATUS_ARMGenerator() gopter.Gen {
	if changeFeed_STATUS_ARMGenerator != nil {
		return changeFeed_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeed_STATUS_ARM(generators)
	changeFeed_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ChangeFeed_STATUS_ARM{}), generators)

	return changeFeed_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeed_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeed_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_CorsRules_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRules_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRules_STATUS_ARM, CorsRules_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRules_STATUS_ARM runs a test to see if a specific instance of CorsRules_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRules_STATUS_ARM(subject CorsRules_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRules_STATUS_ARM
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

// Generator of CorsRules_STATUS_ARM instances for property testing - lazily instantiated by
// CorsRules_STATUS_ARMGenerator()
var corsRules_STATUS_ARMGenerator gopter.Gen

// CorsRules_STATUS_ARMGenerator returns a generator of CorsRules_STATUS_ARM instances for property testing.
func CorsRules_STATUS_ARMGenerator() gopter.Gen {
	if corsRules_STATUS_ARMGenerator != nil {
		return corsRules_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRules_STATUS_ARM(generators)
	corsRules_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(CorsRules_STATUS_ARM{}), generators)

	return corsRules_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForCorsRules_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRules_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRule_STATUS_ARMGenerator())
}

func Test_DeleteRetentionPolicy_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicy_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicy_STATUS_ARM, DeleteRetentionPolicy_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicy_STATUS_ARM runs a test to see if a specific instance of DeleteRetentionPolicy_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicy_STATUS_ARM(subject DeleteRetentionPolicy_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeleteRetentionPolicy_STATUS_ARM
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

// Generator of DeleteRetentionPolicy_STATUS_ARM instances for property testing - lazily instantiated by
// DeleteRetentionPolicy_STATUS_ARMGenerator()
var deleteRetentionPolicy_STATUS_ARMGenerator gopter.Gen

// DeleteRetentionPolicy_STATUS_ARMGenerator returns a generator of DeleteRetentionPolicy_STATUS_ARM instances for property testing.
func DeleteRetentionPolicy_STATUS_ARMGenerator() gopter.Gen {
	if deleteRetentionPolicy_STATUS_ARMGenerator != nil {
		return deleteRetentionPolicy_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_STATUS_ARM(generators)
	deleteRetentionPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicy_STATUS_ARM{}), generators)

	return deleteRetentionPolicy_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_LastAccessTimeTrackingPolicy_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicy_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicy_STATUS_ARM, LastAccessTimeTrackingPolicy_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicy_STATUS_ARM runs a test to see if a specific instance of LastAccessTimeTrackingPolicy_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicy_STATUS_ARM(subject LastAccessTimeTrackingPolicy_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LastAccessTimeTrackingPolicy_STATUS_ARM
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

// Generator of LastAccessTimeTrackingPolicy_STATUS_ARM instances for property testing - lazily instantiated by
// LastAccessTimeTrackingPolicy_STATUS_ARMGenerator()
var lastAccessTimeTrackingPolicy_STATUS_ARMGenerator gopter.Gen

// LastAccessTimeTrackingPolicy_STATUS_ARMGenerator returns a generator of LastAccessTimeTrackingPolicy_STATUS_ARM instances for property testing.
func LastAccessTimeTrackingPolicy_STATUS_ARMGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicy_STATUS_ARMGenerator != nil {
		return lastAccessTimeTrackingPolicy_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_STATUS_ARM(generators)
	lastAccessTimeTrackingPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicy_STATUS_ARM{}), generators)

	return lastAccessTimeTrackingPolicy_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(LastAccessTimeTrackingPolicy_Name_STATUS_AccessTimeTracking))
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_RestorePolicyProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyProperties_STATUS_ARM, RestorePolicyProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyProperties_STATUS_ARM runs a test to see if a specific instance of RestorePolicyProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyProperties_STATUS_ARM(subject RestorePolicyProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RestorePolicyProperties_STATUS_ARM
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

// Generator of RestorePolicyProperties_STATUS_ARM instances for property testing - lazily instantiated by
// RestorePolicyProperties_STATUS_ARMGenerator()
var restorePolicyProperties_STATUS_ARMGenerator gopter.Gen

// RestorePolicyProperties_STATUS_ARMGenerator returns a generator of RestorePolicyProperties_STATUS_ARM instances for property testing.
func RestorePolicyProperties_STATUS_ARMGenerator() gopter.Gen {
	if restorePolicyProperties_STATUS_ARMGenerator != nil {
		return restorePolicyProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyProperties_STATUS_ARM(generators)
	restorePolicyProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RestorePolicyProperties_STATUS_ARM{}), generators)

	return restorePolicyProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["LastEnabledTime"] = gen.PtrOf(gen.AlphaString())
	gens["MinRestoreTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_CorsRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRule_STATUS_ARM, CorsRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRule_STATUS_ARM runs a test to see if a specific instance of CorsRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRule_STATUS_ARM(subject CorsRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRule_STATUS_ARM
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

// Generator of CorsRule_STATUS_ARM instances for property testing - lazily instantiated by
// CorsRule_STATUS_ARMGenerator()
var corsRule_STATUS_ARMGenerator gopter.Gen

// CorsRule_STATUS_ARMGenerator returns a generator of CorsRule_STATUS_ARM instances for property testing.
func CorsRule_STATUS_ARMGenerator() gopter.Gen {
	if corsRule_STATUS_ARMGenerator != nil {
		return corsRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRule_STATUS_ARM(generators)
	corsRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(CorsRule_STATUS_ARM{}), generators)

	return corsRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCorsRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.OneConstOf(
		CorsRule_AllowedMethods_STATUS_DELETE,
		CorsRule_AllowedMethods_STATUS_GET,
		CorsRule_AllowedMethods_STATUS_HEAD,
		CorsRule_AllowedMethods_STATUS_MERGE,
		CorsRule_AllowedMethods_STATUS_OPTIONS,
		CorsRule_AllowedMethods_STATUS_POST,
		CorsRule_AllowedMethods_STATUS_PUT))
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.PtrOf(gen.Int())
}
