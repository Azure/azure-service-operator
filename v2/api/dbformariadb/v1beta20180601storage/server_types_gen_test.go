// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601storage

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

func Test_Server_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Server via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServer, ServerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServer runs a test to see if a specific instance of Server round trips to JSON and back losslessly
func RunJSONSerializationTestForServer(subject Server) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Server
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

// Generator of Server instances for property testing - lazily instantiated by ServerGenerator()
var serverGenerator gopter.Gen

// ServerGenerator returns a generator of Server instances for property testing.
func ServerGenerator() gopter.Gen {
	if serverGenerator != nil {
		return serverGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServer(generators)
	serverGenerator = gen.Struct(reflect.TypeOf(Server{}), generators)

	return serverGenerator
}

// AddRelatedPropertyGeneratorsForServer is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServer(gens map[string]gopter.Gen) {
	gens["Spec"] = Server_SpecGenerator()
	gens["Status"] = Server_STATUSGenerator()
}

func Test_Server_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Server_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServer_STATUS, Server_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServer_STATUS runs a test to see if a specific instance of Server_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServer_STATUS(subject Server_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Server_STATUS
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

// Generator of Server_STATUS instances for property testing - lazily instantiated by Server_STATUSGenerator()
var server_STATUSGenerator gopter.Gen

// Server_STATUSGenerator returns a generator of Server_STATUS instances for property testing.
// We first initialize server_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Server_STATUSGenerator() gopter.Gen {
	if server_STATUSGenerator != nil {
		return server_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServer_STATUS(generators)
	server_STATUSGenerator = gen.Struct(reflect.TypeOf(Server_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServer_STATUS(generators)
	AddRelatedPropertyGeneratorsForServer_STATUS(generators)
	server_STATUSGenerator = gen.Struct(reflect.TypeOf(Server_STATUS{}), generators)

	return server_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServer_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServer_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServer_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServer_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerPropertiesForCreate_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
}

func Test_Server_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Server_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServer_Spec, Server_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServer_Spec runs a test to see if a specific instance of Server_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServer_Spec(subject Server_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Server_Spec
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

// Generator of Server_Spec instances for property testing - lazily instantiated by Server_SpecGenerator()
var server_SpecGenerator gopter.Gen

// Server_SpecGenerator returns a generator of Server_Spec instances for property testing.
// We first initialize server_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Server_SpecGenerator() gopter.Gen {
	if server_SpecGenerator != nil {
		return server_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServer_Spec(generators)
	server_SpecGenerator = gen.Struct(reflect.TypeOf(Server_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServer_Spec(generators)
	AddRelatedPropertyGeneratorsForServer_Spec(generators)
	server_SpecGenerator = gen.Struct(reflect.TypeOf(Server_Spec{}), generators)

	return server_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServer_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServer_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServer_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServer_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(ServerOperatorSpecGenerator())
	gens["Properties"] = gen.PtrOf(ServerPropertiesForCreateGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_ServerOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerOperatorSpec, ServerOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerOperatorSpec runs a test to see if a specific instance of ServerOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForServerOperatorSpec(subject ServerOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerOperatorSpec
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

// Generator of ServerOperatorSpec instances for property testing - lazily instantiated by ServerOperatorSpecGenerator()
var serverOperatorSpecGenerator gopter.Gen

// ServerOperatorSpecGenerator returns a generator of ServerOperatorSpec instances for property testing.
func ServerOperatorSpecGenerator() gopter.Gen {
	if serverOperatorSpecGenerator != nil {
		return serverOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServerOperatorSpec(generators)
	serverOperatorSpecGenerator = gen.Struct(reflect.TypeOf(ServerOperatorSpec{}), generators)

	return serverOperatorSpecGenerator
}

// AddRelatedPropertyGeneratorsForServerOperatorSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerOperatorSpec(gens map[string]gopter.Gen) {
	gens["Secrets"] = gen.PtrOf(ServerOperatorSecretsGenerator())
}

func Test_ServerPropertiesForCreate_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerPropertiesForCreate via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPropertiesForCreate, ServerPropertiesForCreateGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPropertiesForCreate runs a test to see if a specific instance of ServerPropertiesForCreate round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPropertiesForCreate(subject ServerPropertiesForCreate) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerPropertiesForCreate
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

// Generator of ServerPropertiesForCreate instances for property testing - lazily instantiated by
// ServerPropertiesForCreateGenerator()
var serverPropertiesForCreateGenerator gopter.Gen

// ServerPropertiesForCreateGenerator returns a generator of ServerPropertiesForCreate instances for property testing.
// We first initialize serverPropertiesForCreateGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerPropertiesForCreateGenerator() gopter.Gen {
	if serverPropertiesForCreateGenerator != nil {
		return serverPropertiesForCreateGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesForCreate(generators)
	serverPropertiesForCreateGenerator = gen.Struct(reflect.TypeOf(ServerPropertiesForCreate{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesForCreate(generators)
	AddRelatedPropertyGeneratorsForServerPropertiesForCreate(generators)
	serverPropertiesForCreateGenerator = gen.Struct(reflect.TypeOf(ServerPropertiesForCreate{}), generators)

	return serverPropertiesForCreateGenerator
}

// AddIndependentPropertyGeneratorsForServerPropertiesForCreate is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPropertiesForCreate(gens map[string]gopter.Gen) {
	gens["CreateMode"] = gen.PtrOf(gen.AlphaString())
	gens["MinimalTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["SslEnforcement"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServerPropertiesForCreate is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerPropertiesForCreate(gens map[string]gopter.Gen) {
	gens["StorageProfile"] = gen.PtrOf(StorageProfileGenerator())
}

func Test_ServerPropertiesForCreate_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerPropertiesForCreate_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPropertiesForCreate_STATUS, ServerPropertiesForCreate_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPropertiesForCreate_STATUS runs a test to see if a specific instance of ServerPropertiesForCreate_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPropertiesForCreate_STATUS(subject ServerPropertiesForCreate_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerPropertiesForCreate_STATUS
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

// Generator of ServerPropertiesForCreate_STATUS instances for property testing - lazily instantiated by
// ServerPropertiesForCreate_STATUSGenerator()
var serverPropertiesForCreate_STATUSGenerator gopter.Gen

// ServerPropertiesForCreate_STATUSGenerator returns a generator of ServerPropertiesForCreate_STATUS instances for property testing.
// We first initialize serverPropertiesForCreate_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerPropertiesForCreate_STATUSGenerator() gopter.Gen {
	if serverPropertiesForCreate_STATUSGenerator != nil {
		return serverPropertiesForCreate_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesForCreate_STATUS(generators)
	serverPropertiesForCreate_STATUSGenerator = gen.Struct(reflect.TypeOf(ServerPropertiesForCreate_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesForCreate_STATUS(generators)
	AddRelatedPropertyGeneratorsForServerPropertiesForCreate_STATUS(generators)
	serverPropertiesForCreate_STATUSGenerator = gen.Struct(reflect.TypeOf(ServerPropertiesForCreate_STATUS{}), generators)

	return serverPropertiesForCreate_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServerPropertiesForCreate_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPropertiesForCreate_STATUS(gens map[string]gopter.Gen) {
	gens["CreateMode"] = gen.PtrOf(gen.AlphaString())
	gens["MinimalTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["SslEnforcement"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServerPropertiesForCreate_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerPropertiesForCreate_STATUS(gens map[string]gopter.Gen) {
	gens["StorageProfile"] = gen.PtrOf(StorageProfile_STATUSGenerator())
}

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Size"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS, Sku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS runs a test to see if a specific instance of Sku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS(subject Sku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS
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

// Generator of Sku_STATUS instances for property testing - lazily instantiated by Sku_STATUSGenerator()
var sku_STATUSGenerator gopter.Gen

// Sku_STATUSGenerator returns a generator of Sku_STATUS instances for property testing.
func Sku_STATUSGenerator() gopter.Gen {
	if sku_STATUSGenerator != nil {
		return sku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS(generators)
	sku_STATUSGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS{}), generators)

	return sku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Size"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServerOperatorSecrets_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerOperatorSecrets via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerOperatorSecrets, ServerOperatorSecretsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerOperatorSecrets runs a test to see if a specific instance of ServerOperatorSecrets round trips to JSON and back losslessly
func RunJSONSerializationTestForServerOperatorSecrets(subject ServerOperatorSecrets) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerOperatorSecrets
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

// Generator of ServerOperatorSecrets instances for property testing - lazily instantiated by
// ServerOperatorSecretsGenerator()
var serverOperatorSecretsGenerator gopter.Gen

// ServerOperatorSecretsGenerator returns a generator of ServerOperatorSecrets instances for property testing.
func ServerOperatorSecretsGenerator() gopter.Gen {
	if serverOperatorSecretsGenerator != nil {
		return serverOperatorSecretsGenerator
	}

	generators := make(map[string]gopter.Gen)
	serverOperatorSecretsGenerator = gen.Struct(reflect.TypeOf(ServerOperatorSecrets{}), generators)

	return serverOperatorSecretsGenerator
}

func Test_StorageProfile_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageProfile via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageProfile, StorageProfileGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageProfile runs a test to see if a specific instance of StorageProfile round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageProfile(subject StorageProfile) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageProfile
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

// Generator of StorageProfile instances for property testing - lazily instantiated by StorageProfileGenerator()
var storageProfileGenerator gopter.Gen

// StorageProfileGenerator returns a generator of StorageProfile instances for property testing.
func StorageProfileGenerator() gopter.Gen {
	if storageProfileGenerator != nil {
		return storageProfileGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageProfile(generators)
	storageProfileGenerator = gen.Struct(reflect.TypeOf(StorageProfile{}), generators)

	return storageProfileGenerator
}

// AddIndependentPropertyGeneratorsForStorageProfile is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageProfile(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAutogrow"] = gen.PtrOf(gen.AlphaString())
	gens["StorageMB"] = gen.PtrOf(gen.Int())
}

func Test_StorageProfile_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageProfile_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageProfile_STATUS, StorageProfile_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageProfile_STATUS runs a test to see if a specific instance of StorageProfile_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageProfile_STATUS(subject StorageProfile_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageProfile_STATUS
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

// Generator of StorageProfile_STATUS instances for property testing - lazily instantiated by
// StorageProfile_STATUSGenerator()
var storageProfile_STATUSGenerator gopter.Gen

// StorageProfile_STATUSGenerator returns a generator of StorageProfile_STATUS instances for property testing.
func StorageProfile_STATUSGenerator() gopter.Gen {
	if storageProfile_STATUSGenerator != nil {
		return storageProfile_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageProfile_STATUS(generators)
	storageProfile_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageProfile_STATUS{}), generators)

	return storageProfile_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageProfile_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageProfile_STATUS(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAutogrow"] = gen.PtrOf(gen.AlphaString())
	gens["StorageMB"] = gen.PtrOf(gen.Int())
}
