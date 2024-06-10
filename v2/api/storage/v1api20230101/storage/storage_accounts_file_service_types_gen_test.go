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

func Test_Multichannel_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Multichannel via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMultichannel, MultichannelGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMultichannel runs a test to see if a specific instance of Multichannel round trips to JSON and back losslessly
func RunJSONSerializationTestForMultichannel(subject Multichannel) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Multichannel
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

// Generator of Multichannel instances for property testing - lazily instantiated by MultichannelGenerator()
var multichannelGenerator gopter.Gen

// MultichannelGenerator returns a generator of Multichannel instances for property testing.
func MultichannelGenerator() gopter.Gen {
	if multichannelGenerator != nil {
		return multichannelGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMultichannel(generators)
	multichannelGenerator = gen.Struct(reflect.TypeOf(Multichannel{}), generators)

	return multichannelGenerator
}

// AddIndependentPropertyGeneratorsForMultichannel is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMultichannel(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_Multichannel_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Multichannel_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMultichannel_STATUS, Multichannel_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMultichannel_STATUS runs a test to see if a specific instance of Multichannel_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMultichannel_STATUS(subject Multichannel_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Multichannel_STATUS
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

// Generator of Multichannel_STATUS instances for property testing - lazily instantiated by
// Multichannel_STATUSGenerator()
var multichannel_STATUSGenerator gopter.Gen

// Multichannel_STATUSGenerator returns a generator of Multichannel_STATUS instances for property testing.
func Multichannel_STATUSGenerator() gopter.Gen {
	if multichannel_STATUSGenerator != nil {
		return multichannel_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMultichannel_STATUS(generators)
	multichannel_STATUSGenerator = gen.Struct(reflect.TypeOf(Multichannel_STATUS{}), generators)

	return multichannel_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMultichannel_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMultichannel_STATUS(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_ProtocolSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProtocolSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProtocolSettings, ProtocolSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProtocolSettings runs a test to see if a specific instance of ProtocolSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForProtocolSettings(subject ProtocolSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProtocolSettings
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

// Generator of ProtocolSettings instances for property testing - lazily instantiated by ProtocolSettingsGenerator()
var protocolSettingsGenerator gopter.Gen

// ProtocolSettingsGenerator returns a generator of ProtocolSettings instances for property testing.
func ProtocolSettingsGenerator() gopter.Gen {
	if protocolSettingsGenerator != nil {
		return protocolSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForProtocolSettings(generators)
	protocolSettingsGenerator = gen.Struct(reflect.TypeOf(ProtocolSettings{}), generators)

	return protocolSettingsGenerator
}

// AddRelatedPropertyGeneratorsForProtocolSettings is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProtocolSettings(gens map[string]gopter.Gen) {
	gens["Smb"] = gen.PtrOf(SmbSettingGenerator())
}

func Test_ProtocolSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProtocolSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProtocolSettings_STATUS, ProtocolSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProtocolSettings_STATUS runs a test to see if a specific instance of ProtocolSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProtocolSettings_STATUS(subject ProtocolSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProtocolSettings_STATUS
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

// Generator of ProtocolSettings_STATUS instances for property testing - lazily instantiated by
// ProtocolSettings_STATUSGenerator()
var protocolSettings_STATUSGenerator gopter.Gen

// ProtocolSettings_STATUSGenerator returns a generator of ProtocolSettings_STATUS instances for property testing.
func ProtocolSettings_STATUSGenerator() gopter.Gen {
	if protocolSettings_STATUSGenerator != nil {
		return protocolSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForProtocolSettings_STATUS(generators)
	protocolSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(ProtocolSettings_STATUS{}), generators)

	return protocolSettings_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForProtocolSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProtocolSettings_STATUS(gens map[string]gopter.Gen) {
	gens["Smb"] = gen.PtrOf(SmbSetting_STATUSGenerator())
}

func Test_SmbSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SmbSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSmbSetting, SmbSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSmbSetting runs a test to see if a specific instance of SmbSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForSmbSetting(subject SmbSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SmbSetting
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

// Generator of SmbSetting instances for property testing - lazily instantiated by SmbSettingGenerator()
var smbSettingGenerator gopter.Gen

// SmbSettingGenerator returns a generator of SmbSetting instances for property testing.
// We first initialize smbSettingGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SmbSettingGenerator() gopter.Gen {
	if smbSettingGenerator != nil {
		return smbSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSmbSetting(generators)
	smbSettingGenerator = gen.Struct(reflect.TypeOf(SmbSetting{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSmbSetting(generators)
	AddRelatedPropertyGeneratorsForSmbSetting(generators)
	smbSettingGenerator = gen.Struct(reflect.TypeOf(SmbSetting{}), generators)

	return smbSettingGenerator
}

// AddIndependentPropertyGeneratorsForSmbSetting is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSmbSetting(gens map[string]gopter.Gen) {
	gens["AuthenticationMethods"] = gen.PtrOf(gen.AlphaString())
	gens["ChannelEncryption"] = gen.PtrOf(gen.AlphaString())
	gens["KerberosTicketEncryption"] = gen.PtrOf(gen.AlphaString())
	gens["Versions"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSmbSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSmbSetting(gens map[string]gopter.Gen) {
	gens["Multichannel"] = gen.PtrOf(MultichannelGenerator())
}

func Test_SmbSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SmbSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSmbSetting_STATUS, SmbSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSmbSetting_STATUS runs a test to see if a specific instance of SmbSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSmbSetting_STATUS(subject SmbSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SmbSetting_STATUS
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

// Generator of SmbSetting_STATUS instances for property testing - lazily instantiated by SmbSetting_STATUSGenerator()
var smbSetting_STATUSGenerator gopter.Gen

// SmbSetting_STATUSGenerator returns a generator of SmbSetting_STATUS instances for property testing.
// We first initialize smbSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SmbSetting_STATUSGenerator() gopter.Gen {
	if smbSetting_STATUSGenerator != nil {
		return smbSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSmbSetting_STATUS(generators)
	smbSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(SmbSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSmbSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForSmbSetting_STATUS(generators)
	smbSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(SmbSetting_STATUS{}), generators)

	return smbSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSmbSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSmbSetting_STATUS(gens map[string]gopter.Gen) {
	gens["AuthenticationMethods"] = gen.PtrOf(gen.AlphaString())
	gens["ChannelEncryption"] = gen.PtrOf(gen.AlphaString())
	gens["KerberosTicketEncryption"] = gen.PtrOf(gen.AlphaString())
	gens["Versions"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSmbSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSmbSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Multichannel"] = gen.PtrOf(Multichannel_STATUSGenerator())
}

func Test_StorageAccountsFileService_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsFileService via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsFileService, StorageAccountsFileServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsFileService runs a test to see if a specific instance of StorageAccountsFileService round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsFileService(subject StorageAccountsFileService) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsFileService
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

// Generator of StorageAccountsFileService instances for property testing - lazily instantiated by
// StorageAccountsFileServiceGenerator()
var storageAccountsFileServiceGenerator gopter.Gen

// StorageAccountsFileServiceGenerator returns a generator of StorageAccountsFileService instances for property testing.
func StorageAccountsFileServiceGenerator() gopter.Gen {
	if storageAccountsFileServiceGenerator != nil {
		return storageAccountsFileServiceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsFileService(generators)
	storageAccountsFileServiceGenerator = gen.Struct(reflect.TypeOf(StorageAccountsFileService{}), generators)

	return storageAccountsFileServiceGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsFileService is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsFileService(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccounts_FileService_SpecGenerator()
	gens["Status"] = StorageAccounts_FileService_STATUSGenerator()
}

func Test_StorageAccounts_FileService_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_FileService_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_FileService_STATUS, StorageAccounts_FileService_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_FileService_STATUS runs a test to see if a specific instance of StorageAccounts_FileService_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_FileService_STATUS(subject StorageAccounts_FileService_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_FileService_STATUS
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

// Generator of StorageAccounts_FileService_STATUS instances for property testing - lazily instantiated by
// StorageAccounts_FileService_STATUSGenerator()
var storageAccounts_FileService_STATUSGenerator gopter.Gen

// StorageAccounts_FileService_STATUSGenerator returns a generator of StorageAccounts_FileService_STATUS instances for property testing.
// We first initialize storageAccounts_FileService_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_FileService_STATUSGenerator() gopter.Gen {
	if storageAccounts_FileService_STATUSGenerator != nil {
		return storageAccounts_FileService_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_FileService_STATUS(generators)
	storageAccounts_FileService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_FileService_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_FileService_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_FileService_STATUS(generators)
	storageAccounts_FileService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_FileService_STATUS{}), generators)

	return storageAccounts_FileService_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_FileService_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_FileService_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccounts_FileService_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_FileService_STATUS(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRules_STATUSGenerator())
	gens["ProtocolSettings"] = gen.PtrOf(ProtocolSettings_STATUSGenerator())
	gens["ShareDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
}

func Test_StorageAccounts_FileService_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_FileService_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_FileService_Spec, StorageAccounts_FileService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_FileService_Spec runs a test to see if a specific instance of StorageAccounts_FileService_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_FileService_Spec(subject StorageAccounts_FileService_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_FileService_Spec
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

// Generator of StorageAccounts_FileService_Spec instances for property testing - lazily instantiated by
// StorageAccounts_FileService_SpecGenerator()
var storageAccounts_FileService_SpecGenerator gopter.Gen

// StorageAccounts_FileService_SpecGenerator returns a generator of StorageAccounts_FileService_Spec instances for property testing.
// We first initialize storageAccounts_FileService_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_FileService_SpecGenerator() gopter.Gen {
	if storageAccounts_FileService_SpecGenerator != nil {
		return storageAccounts_FileService_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_FileService_Spec(generators)
	storageAccounts_FileService_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_FileService_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_FileService_Spec(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_FileService_Spec(generators)
	storageAccounts_FileService_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_FileService_Spec{}), generators)

	return storageAccounts_FileService_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_FileService_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_FileService_Spec(gens map[string]gopter.Gen) {
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccounts_FileService_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_FileService_Spec(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRulesGenerator())
	gens["ProtocolSettings"] = gen.PtrOf(ProtocolSettingsGenerator())
	gens["ShareDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyGenerator())
}
