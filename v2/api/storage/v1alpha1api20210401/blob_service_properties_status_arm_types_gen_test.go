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

func Test_BlobServiceProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobServiceProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobServiceProperties_STATUSARM, BlobServiceProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobServiceProperties_STATUSARM runs a test to see if a specific instance of BlobServiceProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobServiceProperties_STATUSARM(subject BlobServiceProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobServiceProperties_STATUSARM
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

// Generator of BlobServiceProperties_STATUSARM instances for property testing - lazily instantiated by
// BlobServiceProperties_STATUSARMGenerator()
var blobServiceProperties_STATUSARMGenerator gopter.Gen

// BlobServiceProperties_STATUSARMGenerator returns a generator of BlobServiceProperties_STATUSARM instances for property testing.
// We first initialize blobServiceProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobServiceProperties_STATUSARMGenerator() gopter.Gen {
	if blobServiceProperties_STATUSARMGenerator != nil {
		return blobServiceProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUSARM(generators)
	blobServiceProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForBlobServiceProperties_STATUSARM(generators)
	blobServiceProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_STATUSARM{}), generators)

	return blobServiceProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBlobServiceProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobServiceProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(BlobServiceProperties_STATUS_PropertiesARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSARMGenerator())
}

func Test_BlobServiceProperties_STATUS_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobServiceProperties_STATUS_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobServiceProperties_STATUS_PropertiesARM, BlobServiceProperties_STATUS_PropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobServiceProperties_STATUS_PropertiesARM runs a test to see if a specific instance of BlobServiceProperties_STATUS_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobServiceProperties_STATUS_PropertiesARM(subject BlobServiceProperties_STATUS_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobServiceProperties_STATUS_PropertiesARM
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

// Generator of BlobServiceProperties_STATUS_PropertiesARM instances for property testing - lazily instantiated by
// BlobServiceProperties_STATUS_PropertiesARMGenerator()
var blobServiceProperties_STATUS_PropertiesARMGenerator gopter.Gen

// BlobServiceProperties_STATUS_PropertiesARMGenerator returns a generator of BlobServiceProperties_STATUS_PropertiesARM instances for property testing.
// We first initialize blobServiceProperties_STATUS_PropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobServiceProperties_STATUS_PropertiesARMGenerator() gopter.Gen {
	if blobServiceProperties_STATUS_PropertiesARMGenerator != nil {
		return blobServiceProperties_STATUS_PropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUS_PropertiesARM(generators)
	blobServiceProperties_STATUS_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_STATUS_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUS_PropertiesARM(generators)
	AddRelatedPropertyGeneratorsForBlobServiceProperties_STATUS_PropertiesARM(generators)
	blobServiceProperties_STATUS_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_STATUS_PropertiesARM{}), generators)

	return blobServiceProperties_STATUS_PropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUS_PropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobServiceProperties_STATUS_PropertiesARM(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForBlobServiceProperties_STATUS_PropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobServiceProperties_STATUS_PropertiesARM(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeed_STATUSARMGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_STATUSARMGenerator())
	gens["Cors"] = gen.PtrOf(CorsRules_STATUSARMGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_STATUSARMGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicy_STATUSARMGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyProperties_STATUSARMGenerator())
}

func Test_Sku_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUSARM, Sku_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUSARM runs a test to see if a specific instance of Sku_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUSARM(subject Sku_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUSARM
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

// Generator of Sku_STATUSARM instances for property testing - lazily instantiated by Sku_STATUSARMGenerator()
var sku_STATUSARMGenerator gopter.Gen

// Sku_STATUSARMGenerator returns a generator of Sku_STATUSARM instances for property testing.
func Sku_STATUSARMGenerator() gopter.Gen {
	if sku_STATUSARMGenerator != nil {
		return sku_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUSARM(generators)
	sku_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUSARM{}), generators)

	return sku_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUSARM(gens map[string]gopter.Gen) {
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

func Test_ChangeFeed_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeed_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeed_STATUSARM, ChangeFeed_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeed_STATUSARM runs a test to see if a specific instance of ChangeFeed_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeed_STATUSARM(subject ChangeFeed_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ChangeFeed_STATUSARM
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

// Generator of ChangeFeed_STATUSARM instances for property testing - lazily instantiated by
// ChangeFeed_STATUSARMGenerator()
var changeFeed_STATUSARMGenerator gopter.Gen

// ChangeFeed_STATUSARMGenerator returns a generator of ChangeFeed_STATUSARM instances for property testing.
func ChangeFeed_STATUSARMGenerator() gopter.Gen {
	if changeFeed_STATUSARMGenerator != nil {
		return changeFeed_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeed_STATUSARM(generators)
	changeFeed_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ChangeFeed_STATUSARM{}), generators)

	return changeFeed_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeed_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeed_STATUSARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_CorsRules_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRules_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRules_STATUSARM, CorsRules_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRules_STATUSARM runs a test to see if a specific instance of CorsRules_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRules_STATUSARM(subject CorsRules_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRules_STATUSARM
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

// Generator of CorsRules_STATUSARM instances for property testing - lazily instantiated by
// CorsRules_STATUSARMGenerator()
var corsRules_STATUSARMGenerator gopter.Gen

// CorsRules_STATUSARMGenerator returns a generator of CorsRules_STATUSARM instances for property testing.
func CorsRules_STATUSARMGenerator() gopter.Gen {
	if corsRules_STATUSARMGenerator != nil {
		return corsRules_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRules_STATUSARM(generators)
	corsRules_STATUSARMGenerator = gen.Struct(reflect.TypeOf(CorsRules_STATUSARM{}), generators)

	return corsRules_STATUSARMGenerator
}

// AddRelatedPropertyGeneratorsForCorsRules_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRules_STATUSARM(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRule_STATUSARMGenerator())
}

func Test_DeleteRetentionPolicy_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicy_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicy_STATUSARM, DeleteRetentionPolicy_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicy_STATUSARM runs a test to see if a specific instance of DeleteRetentionPolicy_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicy_STATUSARM(subject DeleteRetentionPolicy_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeleteRetentionPolicy_STATUSARM
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

// Generator of DeleteRetentionPolicy_STATUSARM instances for property testing - lazily instantiated by
// DeleteRetentionPolicy_STATUSARMGenerator()
var deleteRetentionPolicy_STATUSARMGenerator gopter.Gen

// DeleteRetentionPolicy_STATUSARMGenerator returns a generator of DeleteRetentionPolicy_STATUSARM instances for property testing.
func DeleteRetentionPolicy_STATUSARMGenerator() gopter.Gen {
	if deleteRetentionPolicy_STATUSARMGenerator != nil {
		return deleteRetentionPolicy_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_STATUSARM(generators)
	deleteRetentionPolicy_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicy_STATUSARM{}), generators)

	return deleteRetentionPolicy_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_STATUSARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_LastAccessTimeTrackingPolicy_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicy_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicy_STATUSARM, LastAccessTimeTrackingPolicy_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicy_STATUSARM runs a test to see if a specific instance of LastAccessTimeTrackingPolicy_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicy_STATUSARM(subject LastAccessTimeTrackingPolicy_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LastAccessTimeTrackingPolicy_STATUSARM
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

// Generator of LastAccessTimeTrackingPolicy_STATUSARM instances for property testing - lazily instantiated by
// LastAccessTimeTrackingPolicy_STATUSARMGenerator()
var lastAccessTimeTrackingPolicy_STATUSARMGenerator gopter.Gen

// LastAccessTimeTrackingPolicy_STATUSARMGenerator returns a generator of LastAccessTimeTrackingPolicy_STATUSARM instances for property testing.
func LastAccessTimeTrackingPolicy_STATUSARMGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicy_STATUSARMGenerator != nil {
		return lastAccessTimeTrackingPolicy_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_STATUSARM(generators)
	lastAccessTimeTrackingPolicy_STATUSARMGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicy_STATUSARM{}), generators)

	return lastAccessTimeTrackingPolicy_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_STATUSARM(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(LastAccessTimeTrackingPolicy_STATUS_Name_AccessTimeTracking))
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_RestorePolicyProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyProperties_STATUSARM, RestorePolicyProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyProperties_STATUSARM runs a test to see if a specific instance of RestorePolicyProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyProperties_STATUSARM(subject RestorePolicyProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RestorePolicyProperties_STATUSARM
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

// Generator of RestorePolicyProperties_STATUSARM instances for property testing - lazily instantiated by
// RestorePolicyProperties_STATUSARMGenerator()
var restorePolicyProperties_STATUSARMGenerator gopter.Gen

// RestorePolicyProperties_STATUSARMGenerator returns a generator of RestorePolicyProperties_STATUSARM instances for property testing.
func RestorePolicyProperties_STATUSARMGenerator() gopter.Gen {
	if restorePolicyProperties_STATUSARMGenerator != nil {
		return restorePolicyProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyProperties_STATUSARM(generators)
	restorePolicyProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RestorePolicyProperties_STATUSARM{}), generators)

	return restorePolicyProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["LastEnabledTime"] = gen.PtrOf(gen.AlphaString())
	gens["MinRestoreTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_CorsRule_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRule_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRule_STATUSARM, CorsRule_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRule_STATUSARM runs a test to see if a specific instance of CorsRule_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRule_STATUSARM(subject CorsRule_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRule_STATUSARM
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

// Generator of CorsRule_STATUSARM instances for property testing - lazily instantiated by CorsRule_STATUSARMGenerator()
var corsRule_STATUSARMGenerator gopter.Gen

// CorsRule_STATUSARMGenerator returns a generator of CorsRule_STATUSARM instances for property testing.
func CorsRule_STATUSARMGenerator() gopter.Gen {
	if corsRule_STATUSARMGenerator != nil {
		return corsRule_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRule_STATUSARM(generators)
	corsRule_STATUSARMGenerator = gen.Struct(reflect.TypeOf(CorsRule_STATUSARM{}), generators)

	return corsRule_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForCorsRule_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRule_STATUSARM(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.OneConstOf(
		CorsRule_STATUS_AllowedMethods_DELETE,
		CorsRule_STATUS_AllowedMethods_GET,
		CorsRule_STATUS_AllowedMethods_HEAD,
		CorsRule_STATUS_AllowedMethods_MERGE,
		CorsRule_STATUS_AllowedMethods_OPTIONS,
		CorsRule_STATUS_AllowedMethods_POST,
		CorsRule_STATUS_AllowedMethods_PUT))
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.PtrOf(gen.Int())
}
