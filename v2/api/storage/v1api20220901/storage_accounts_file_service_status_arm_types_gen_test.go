// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

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

func Test_Multichannel_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Multichannel_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMultichannel_STATUS_ARM, Multichannel_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMultichannel_STATUS_ARM runs a test to see if a specific instance of Multichannel_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMultichannel_STATUS_ARM(subject Multichannel_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Multichannel_STATUS_ARM
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

// Generator of Multichannel_STATUS_ARM instances for property testing - lazily instantiated by
// Multichannel_STATUS_ARMGenerator()
var multichannel_STATUS_ARMGenerator gopter.Gen

// Multichannel_STATUS_ARMGenerator returns a generator of Multichannel_STATUS_ARM instances for property testing.
func Multichannel_STATUS_ARMGenerator() gopter.Gen {
	if multichannel_STATUS_ARMGenerator != nil {
		return multichannel_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMultichannel_STATUS_ARM(generators)
	multichannel_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Multichannel_STATUS_ARM{}), generators)

	return multichannel_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMultichannel_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMultichannel_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_ProtocolSettings_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProtocolSettings_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProtocolSettings_STATUS_ARM, ProtocolSettings_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProtocolSettings_STATUS_ARM runs a test to see if a specific instance of ProtocolSettings_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProtocolSettings_STATUS_ARM(subject ProtocolSettings_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProtocolSettings_STATUS_ARM
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

// Generator of ProtocolSettings_STATUS_ARM instances for property testing - lazily instantiated by
// ProtocolSettings_STATUS_ARMGenerator()
var protocolSettings_STATUS_ARMGenerator gopter.Gen

// ProtocolSettings_STATUS_ARMGenerator returns a generator of ProtocolSettings_STATUS_ARM instances for property testing.
func ProtocolSettings_STATUS_ARMGenerator() gopter.Gen {
	if protocolSettings_STATUS_ARMGenerator != nil {
		return protocolSettings_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForProtocolSettings_STATUS_ARM(generators)
	protocolSettings_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ProtocolSettings_STATUS_ARM{}), generators)

	return protocolSettings_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForProtocolSettings_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProtocolSettings_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Smb"] = gen.PtrOf(SmbSetting_STATUS_ARMGenerator())
}

func Test_SmbSetting_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SmbSetting_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSmbSetting_STATUS_ARM, SmbSetting_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSmbSetting_STATUS_ARM runs a test to see if a specific instance of SmbSetting_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSmbSetting_STATUS_ARM(subject SmbSetting_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SmbSetting_STATUS_ARM
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

// Generator of SmbSetting_STATUS_ARM instances for property testing - lazily instantiated by
// SmbSetting_STATUS_ARMGenerator()
var smbSetting_STATUS_ARMGenerator gopter.Gen

// SmbSetting_STATUS_ARMGenerator returns a generator of SmbSetting_STATUS_ARM instances for property testing.
// We first initialize smbSetting_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SmbSetting_STATUS_ARMGenerator() gopter.Gen {
	if smbSetting_STATUS_ARMGenerator != nil {
		return smbSetting_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSmbSetting_STATUS_ARM(generators)
	smbSetting_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SmbSetting_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSmbSetting_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForSmbSetting_STATUS_ARM(generators)
	smbSetting_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SmbSetting_STATUS_ARM{}), generators)

	return smbSetting_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSmbSetting_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSmbSetting_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AuthenticationMethods"] = gen.PtrOf(gen.AlphaString())
	gens["ChannelEncryption"] = gen.PtrOf(gen.AlphaString())
	gens["KerberosTicketEncryption"] = gen.PtrOf(gen.AlphaString())
	gens["Versions"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSmbSetting_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSmbSetting_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Multichannel"] = gen.PtrOf(Multichannel_STATUS_ARMGenerator())
}

func Test_StorageAccounts_FileService_Properties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_FileService_Properties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_FileService_Properties_STATUS_ARM, StorageAccounts_FileService_Properties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_FileService_Properties_STATUS_ARM runs a test to see if a specific instance of StorageAccounts_FileService_Properties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_FileService_Properties_STATUS_ARM(subject StorageAccounts_FileService_Properties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_FileService_Properties_STATUS_ARM
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

// Generator of StorageAccounts_FileService_Properties_STATUS_ARM instances for property testing - lazily instantiated
// by StorageAccounts_FileService_Properties_STATUS_ARMGenerator()
var storageAccounts_FileService_Properties_STATUS_ARMGenerator gopter.Gen

// StorageAccounts_FileService_Properties_STATUS_ARMGenerator returns a generator of StorageAccounts_FileService_Properties_STATUS_ARM instances for property testing.
func StorageAccounts_FileService_Properties_STATUS_ARMGenerator() gopter.Gen {
	if storageAccounts_FileService_Properties_STATUS_ARMGenerator != nil {
		return storageAccounts_FileService_Properties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccounts_FileService_Properties_STATUS_ARM(generators)
	storageAccounts_FileService_Properties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_FileService_Properties_STATUS_ARM{}), generators)

	return storageAccounts_FileService_Properties_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccounts_FileService_Properties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_FileService_Properties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRules_STATUS_ARMGenerator())
	gens["ProtocolSettings"] = gen.PtrOf(ProtocolSettings_STATUS_ARMGenerator())
	gens["ShareDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_STATUS_ARMGenerator())
}

func Test_StorageAccounts_FileService_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_FileService_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_FileService_STATUS_ARM, StorageAccounts_FileService_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_FileService_STATUS_ARM runs a test to see if a specific instance of StorageAccounts_FileService_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_FileService_STATUS_ARM(subject StorageAccounts_FileService_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_FileService_STATUS_ARM
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

// Generator of StorageAccounts_FileService_STATUS_ARM instances for property testing - lazily instantiated by
// StorageAccounts_FileService_STATUS_ARMGenerator()
var storageAccounts_FileService_STATUS_ARMGenerator gopter.Gen

// StorageAccounts_FileService_STATUS_ARMGenerator returns a generator of StorageAccounts_FileService_STATUS_ARM instances for property testing.
// We first initialize storageAccounts_FileService_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_FileService_STATUS_ARMGenerator() gopter.Gen {
	if storageAccounts_FileService_STATUS_ARMGenerator != nil {
		return storageAccounts_FileService_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_FileService_STATUS_ARM(generators)
	storageAccounts_FileService_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_FileService_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_FileService_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_FileService_STATUS_ARM(generators)
	storageAccounts_FileService_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_FileService_STATUS_ARM{}), generators)

	return storageAccounts_FileService_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_FileService_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_FileService_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccounts_FileService_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_FileService_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(StorageAccounts_FileService_Properties_STATUS_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUS_ARMGenerator())
}
