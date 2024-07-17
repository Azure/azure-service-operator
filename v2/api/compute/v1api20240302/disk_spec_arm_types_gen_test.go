// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240302

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

func Test_CreationData_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CreationData_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCreationData_ARM, CreationData_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCreationData_ARM runs a test to see if a specific instance of CreationData_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCreationData_ARM(subject CreationData_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CreationData_ARM
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

// Generator of CreationData_ARM instances for property testing - lazily instantiated by CreationData_ARMGenerator()
var creationData_ARMGenerator gopter.Gen

// CreationData_ARMGenerator returns a generator of CreationData_ARM instances for property testing.
// We first initialize creationData_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CreationData_ARMGenerator() gopter.Gen {
	if creationData_ARMGenerator != nil {
		return creationData_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreationData_ARM(generators)
	creationData_ARMGenerator = gen.Struct(reflect.TypeOf(CreationData_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreationData_ARM(generators)
	AddRelatedPropertyGeneratorsForCreationData_ARM(generators)
	creationData_ARMGenerator = gen.Struct(reflect.TypeOf(CreationData_ARM{}), generators)

	return creationData_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCreationData_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCreationData_ARM(gens map[string]gopter.Gen) {
	gens["CreateOption"] = gen.PtrOf(gen.OneConstOf(
		CreationData_CreateOption_Attach,
		CreationData_CreateOption_Copy,
		CreationData_CreateOption_CopyFromSanSnapshot,
		CreationData_CreateOption_CopyStart,
		CreationData_CreateOption_Empty,
		CreationData_CreateOption_FromImage,
		CreationData_CreateOption_Import,
		CreationData_CreateOption_ImportSecure,
		CreationData_CreateOption_Restore,
		CreationData_CreateOption_Upload,
		CreationData_CreateOption_UploadPreparedSecure))
	gens["ElasticSanResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["LogicalSectorSize"] = gen.PtrOf(gen.Int())
	gens["PerformancePlus"] = gen.PtrOf(gen.Bool())
	gens["ProvisionedBandwidthCopySpeed"] = gen.PtrOf(gen.OneConstOf(CreationData_ProvisionedBandwidthCopySpeed_Enhanced, CreationData_ProvisionedBandwidthCopySpeed_None))
	gens["SecurityDataUri"] = gen.PtrOf(gen.AlphaString())
	gens["SourceResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["SourceUri"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountId"] = gen.PtrOf(gen.AlphaString())
	gens["UploadSizeBytes"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForCreationData_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCreationData_ARM(gens map[string]gopter.Gen) {
	gens["GalleryImageReference"] = gen.PtrOf(ImageDiskReference_ARMGenerator())
	gens["ImageReference"] = gen.PtrOf(ImageDiskReference_ARMGenerator())
}

func Test_DiskProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskProperties_ARM, DiskProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskProperties_ARM runs a test to see if a specific instance of DiskProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskProperties_ARM(subject DiskProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskProperties_ARM
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

// Generator of DiskProperties_ARM instances for property testing - lazily instantiated by DiskProperties_ARMGenerator()
var diskProperties_ARMGenerator gopter.Gen

// DiskProperties_ARMGenerator returns a generator of DiskProperties_ARM instances for property testing.
// We first initialize diskProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiskProperties_ARMGenerator() gopter.Gen {
	if diskProperties_ARMGenerator != nil {
		return diskProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskProperties_ARM(generators)
	diskProperties_ARMGenerator = gen.Struct(reflect.TypeOf(DiskProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForDiskProperties_ARM(generators)
	diskProperties_ARMGenerator = gen.Struct(reflect.TypeOf(DiskProperties_ARM{}), generators)

	return diskProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDiskProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskProperties_ARM(gens map[string]gopter.Gen) {
	gens["BurstingEnabled"] = gen.PtrOf(gen.Bool())
	gens["CompletionPercent"] = gen.PtrOf(gen.Float64())
	gens["DataAccessAuthMode"] = gen.PtrOf(gen.OneConstOf(DataAccessAuthMode_AzureActiveDirectory, DataAccessAuthMode_None))
	gens["DiskAccessId"] = gen.PtrOf(gen.AlphaString())
	gens["DiskIOPSReadOnly"] = gen.PtrOf(gen.Int())
	gens["DiskIOPSReadWrite"] = gen.PtrOf(gen.Int())
	gens["DiskMBpsReadOnly"] = gen.PtrOf(gen.Int())
	gens["DiskMBpsReadWrite"] = gen.PtrOf(gen.Int())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(DiskProperties_HyperVGeneration_V1, DiskProperties_HyperVGeneration_V2))
	gens["MaxShares"] = gen.PtrOf(gen.Int())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.OneConstOf(NetworkAccessPolicy_AllowAll, NetworkAccessPolicy_AllowPrivate, NetworkAccessPolicy_DenyAll))
	gens["OptimizedForFrequentAttach"] = gen.PtrOf(gen.Bool())
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(DiskProperties_OsType_Linux, DiskProperties_OsType_Windows))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccess_Disabled, PublicNetworkAccess_Enabled))
	gens["SupportsHibernation"] = gen.PtrOf(gen.Bool())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiskProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskProperties_ARM(gens map[string]gopter.Gen) {
	gens["CreationData"] = gen.PtrOf(CreationData_ARMGenerator())
	gens["Encryption"] = gen.PtrOf(Encryption_ARMGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollection_ARMGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlan_ARMGenerator())
	gens["SecurityProfile"] = gen.PtrOf(DiskSecurityProfile_ARMGenerator())
	gens["SupportedCapabilities"] = gen.PtrOf(SupportedCapabilities_ARMGenerator())
}

func Test_DiskSecurityProfile_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskSecurityProfile_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskSecurityProfile_ARM, DiskSecurityProfile_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskSecurityProfile_ARM runs a test to see if a specific instance of DiskSecurityProfile_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskSecurityProfile_ARM(subject DiskSecurityProfile_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskSecurityProfile_ARM
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

// Generator of DiskSecurityProfile_ARM instances for property testing - lazily instantiated by
// DiskSecurityProfile_ARMGenerator()
var diskSecurityProfile_ARMGenerator gopter.Gen

// DiskSecurityProfile_ARMGenerator returns a generator of DiskSecurityProfile_ARM instances for property testing.
func DiskSecurityProfile_ARMGenerator() gopter.Gen {
	if diskSecurityProfile_ARMGenerator != nil {
		return diskSecurityProfile_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskSecurityProfile_ARM(generators)
	diskSecurityProfile_ARMGenerator = gen.Struct(reflect.TypeOf(DiskSecurityProfile_ARM{}), generators)

	return diskSecurityProfile_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDiskSecurityProfile_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskSecurityProfile_ARM(gens map[string]gopter.Gen) {
	gens["SecureVMDiskEncryptionSetId"] = gen.PtrOf(gen.AlphaString())
	gens["SecurityType"] = gen.PtrOf(gen.OneConstOf(
		DiskSecurityType_ConfidentialVM_DiskEncryptedWithCustomerKey,
		DiskSecurityType_ConfidentialVM_DiskEncryptedWithPlatformKey,
		DiskSecurityType_ConfidentialVM_NonPersistedTPM,
		DiskSecurityType_ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey,
		DiskSecurityType_TrustedLaunch))
}

func Test_DiskSku_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskSku_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskSku_ARM, DiskSku_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskSku_ARM runs a test to see if a specific instance of DiskSku_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskSku_ARM(subject DiskSku_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskSku_ARM
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

// Generator of DiskSku_ARM instances for property testing - lazily instantiated by DiskSku_ARMGenerator()
var diskSku_ARMGenerator gopter.Gen

// DiskSku_ARMGenerator returns a generator of DiskSku_ARM instances for property testing.
func DiskSku_ARMGenerator() gopter.Gen {
	if diskSku_ARMGenerator != nil {
		return diskSku_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskSku_ARM(generators)
	diskSku_ARMGenerator = gen.Struct(reflect.TypeOf(DiskSku_ARM{}), generators)

	return diskSku_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDiskSku_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskSku_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		DiskSku_Name_PremiumV2_LRS,
		DiskSku_Name_Premium_LRS,
		DiskSku_Name_Premium_ZRS,
		DiskSku_Name_StandardSSD_LRS,
		DiskSku_Name_StandardSSD_ZRS,
		DiskSku_Name_Standard_LRS,
		DiskSku_Name_UltraSSD_LRS))
}

func Test_Disk_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Disk_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDisk_Spec_ARM, Disk_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDisk_Spec_ARM runs a test to see if a specific instance of Disk_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDisk_Spec_ARM(subject Disk_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Disk_Spec_ARM
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

// Generator of Disk_Spec_ARM instances for property testing - lazily instantiated by Disk_Spec_ARMGenerator()
var disk_Spec_ARMGenerator gopter.Gen

// Disk_Spec_ARMGenerator returns a generator of Disk_Spec_ARM instances for property testing.
// We first initialize disk_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Disk_Spec_ARMGenerator() gopter.Gen {
	if disk_Spec_ARMGenerator != nil {
		return disk_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDisk_Spec_ARM(generators)
	disk_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Disk_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDisk_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDisk_Spec_ARM(generators)
	disk_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Disk_Spec_ARM{}), generators)

	return disk_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDisk_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDisk_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDisk_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDisk_Spec_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_ARMGenerator())
	gens["Properties"] = gen.PtrOf(DiskProperties_ARMGenerator())
	gens["Sku"] = gen.PtrOf(DiskSku_ARMGenerator())
}

func Test_EncryptionSettingsCollection_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionSettingsCollection_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionSettingsCollection_ARM, EncryptionSettingsCollection_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionSettingsCollection_ARM runs a test to see if a specific instance of EncryptionSettingsCollection_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionSettingsCollection_ARM(subject EncryptionSettingsCollection_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionSettingsCollection_ARM
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

// Generator of EncryptionSettingsCollection_ARM instances for property testing - lazily instantiated by
// EncryptionSettingsCollection_ARMGenerator()
var encryptionSettingsCollection_ARMGenerator gopter.Gen

// EncryptionSettingsCollection_ARMGenerator returns a generator of EncryptionSettingsCollection_ARM instances for property testing.
// We first initialize encryptionSettingsCollection_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionSettingsCollection_ARMGenerator() gopter.Gen {
	if encryptionSettingsCollection_ARMGenerator != nil {
		return encryptionSettingsCollection_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSettingsCollection_ARM(generators)
	encryptionSettingsCollection_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSettingsCollection_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSettingsCollection_ARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionSettingsCollection_ARM(generators)
	encryptionSettingsCollection_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSettingsCollection_ARM{}), generators)

	return encryptionSettingsCollection_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionSettingsCollection_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionSettingsCollection_ARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["EncryptionSettingsVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEncryptionSettingsCollection_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionSettingsCollection_ARM(gens map[string]gopter.Gen) {
	gens["EncryptionSettings"] = gen.SliceOf(EncryptionSettingsElement_ARMGenerator())
}

func Test_EncryptionSettingsElement_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionSettingsElement_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionSettingsElement_ARM, EncryptionSettingsElement_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionSettingsElement_ARM runs a test to see if a specific instance of EncryptionSettingsElement_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionSettingsElement_ARM(subject EncryptionSettingsElement_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionSettingsElement_ARM
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

// Generator of EncryptionSettingsElement_ARM instances for property testing - lazily instantiated by
// EncryptionSettingsElement_ARMGenerator()
var encryptionSettingsElement_ARMGenerator gopter.Gen

// EncryptionSettingsElement_ARMGenerator returns a generator of EncryptionSettingsElement_ARM instances for property testing.
func EncryptionSettingsElement_ARMGenerator() gopter.Gen {
	if encryptionSettingsElement_ARMGenerator != nil {
		return encryptionSettingsElement_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForEncryptionSettingsElement_ARM(generators)
	encryptionSettingsElement_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSettingsElement_ARM{}), generators)

	return encryptionSettingsElement_ARMGenerator
}

// AddRelatedPropertyGeneratorsForEncryptionSettingsElement_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionSettingsElement_ARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionKey"] = gen.PtrOf(KeyVaultAndSecretReference_ARMGenerator())
	gens["KeyEncryptionKey"] = gen.PtrOf(KeyVaultAndKeyReference_ARMGenerator())
}

func Test_Encryption_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Encryption_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryption_ARM, Encryption_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryption_ARM runs a test to see if a specific instance of Encryption_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryption_ARM(subject Encryption_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Encryption_ARM
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

// Generator of Encryption_ARM instances for property testing - lazily instantiated by Encryption_ARMGenerator()
var encryption_ARMGenerator gopter.Gen

// Encryption_ARMGenerator returns a generator of Encryption_ARM instances for property testing.
func Encryption_ARMGenerator() gopter.Gen {
	if encryption_ARMGenerator != nil {
		return encryption_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryption_ARM(generators)
	encryption_ARMGenerator = gen.Struct(reflect.TypeOf(Encryption_ARM{}), generators)

	return encryption_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryption_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryption_ARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSetId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(EncryptionType_EncryptionAtRestWithCustomerKey, EncryptionType_EncryptionAtRestWithPlatformAndCustomerKeys, EncryptionType_EncryptionAtRestWithPlatformKey))
}

func Test_ImageDiskReference_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageDiskReference_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageDiskReference_ARM, ImageDiskReference_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageDiskReference_ARM runs a test to see if a specific instance of ImageDiskReference_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageDiskReference_ARM(subject ImageDiskReference_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageDiskReference_ARM
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

// Generator of ImageDiskReference_ARM instances for property testing - lazily instantiated by
// ImageDiskReference_ARMGenerator()
var imageDiskReference_ARMGenerator gopter.Gen

// ImageDiskReference_ARMGenerator returns a generator of ImageDiskReference_ARM instances for property testing.
func ImageDiskReference_ARMGenerator() gopter.Gen {
	if imageDiskReference_ARMGenerator != nil {
		return imageDiskReference_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDiskReference_ARM(generators)
	imageDiskReference_ARMGenerator = gen.Struct(reflect.TypeOf(ImageDiskReference_ARM{}), generators)

	return imageDiskReference_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImageDiskReference_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageDiskReference_ARM(gens map[string]gopter.Gen) {
	gens["CommunityGalleryImageId"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Lun"] = gen.PtrOf(gen.Int())
	gens["SharedGalleryImageId"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultAndKeyReference_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultAndKeyReference_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultAndKeyReference_ARM, KeyVaultAndKeyReference_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultAndKeyReference_ARM runs a test to see if a specific instance of KeyVaultAndKeyReference_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultAndKeyReference_ARM(subject KeyVaultAndKeyReference_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultAndKeyReference_ARM
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

// Generator of KeyVaultAndKeyReference_ARM instances for property testing - lazily instantiated by
// KeyVaultAndKeyReference_ARMGenerator()
var keyVaultAndKeyReference_ARMGenerator gopter.Gen

// KeyVaultAndKeyReference_ARMGenerator returns a generator of KeyVaultAndKeyReference_ARM instances for property testing.
// We first initialize keyVaultAndKeyReference_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultAndKeyReference_ARMGenerator() gopter.Gen {
	if keyVaultAndKeyReference_ARMGenerator != nil {
		return keyVaultAndKeyReference_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultAndKeyReference_ARM(generators)
	keyVaultAndKeyReference_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultAndKeyReference_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultAndKeyReference_ARM(generators)
	AddRelatedPropertyGeneratorsForKeyVaultAndKeyReference_ARM(generators)
	keyVaultAndKeyReference_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultAndKeyReference_ARM{}), generators)

	return keyVaultAndKeyReference_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultAndKeyReference_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultAndKeyReference_ARM(gens map[string]gopter.Gen) {
	gens["KeyUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultAndKeyReference_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultAndKeyReference_ARM(gens map[string]gopter.Gen) {
	gens["SourceVault"] = gen.PtrOf(SourceVault_ARMGenerator())
}

func Test_KeyVaultAndSecretReference_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultAndSecretReference_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultAndSecretReference_ARM, KeyVaultAndSecretReference_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultAndSecretReference_ARM runs a test to see if a specific instance of KeyVaultAndSecretReference_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultAndSecretReference_ARM(subject KeyVaultAndSecretReference_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultAndSecretReference_ARM
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

// Generator of KeyVaultAndSecretReference_ARM instances for property testing - lazily instantiated by
// KeyVaultAndSecretReference_ARMGenerator()
var keyVaultAndSecretReference_ARMGenerator gopter.Gen

// KeyVaultAndSecretReference_ARMGenerator returns a generator of KeyVaultAndSecretReference_ARM instances for property testing.
// We first initialize keyVaultAndSecretReference_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultAndSecretReference_ARMGenerator() gopter.Gen {
	if keyVaultAndSecretReference_ARMGenerator != nil {
		return keyVaultAndSecretReference_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultAndSecretReference_ARM(generators)
	keyVaultAndSecretReference_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultAndSecretReference_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultAndSecretReference_ARM(generators)
	AddRelatedPropertyGeneratorsForKeyVaultAndSecretReference_ARM(generators)
	keyVaultAndSecretReference_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultAndSecretReference_ARM{}), generators)

	return keyVaultAndSecretReference_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultAndSecretReference_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultAndSecretReference_ARM(gens map[string]gopter.Gen) {
	gens["SecretUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultAndSecretReference_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultAndSecretReference_ARM(gens map[string]gopter.Gen) {
	gens["SourceVault"] = gen.PtrOf(SourceVault_ARMGenerator())
}

func Test_PurchasePlan_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PurchasePlan_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPurchasePlan_ARM, PurchasePlan_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPurchasePlan_ARM runs a test to see if a specific instance of PurchasePlan_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPurchasePlan_ARM(subject PurchasePlan_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PurchasePlan_ARM
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

// Generator of PurchasePlan_ARM instances for property testing - lazily instantiated by PurchasePlan_ARMGenerator()
var purchasePlan_ARMGenerator gopter.Gen

// PurchasePlan_ARMGenerator returns a generator of PurchasePlan_ARM instances for property testing.
func PurchasePlan_ARMGenerator() gopter.Gen {
	if purchasePlan_ARMGenerator != nil {
		return purchasePlan_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPurchasePlan_ARM(generators)
	purchasePlan_ARMGenerator = gen.Struct(reflect.TypeOf(PurchasePlan_ARM{}), generators)

	return purchasePlan_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPurchasePlan_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPurchasePlan_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Product"] = gen.PtrOf(gen.AlphaString())
	gens["PromotionCode"] = gen.PtrOf(gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
}

func Test_SupportedCapabilities_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SupportedCapabilities_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSupportedCapabilities_ARM, SupportedCapabilities_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSupportedCapabilities_ARM runs a test to see if a specific instance of SupportedCapabilities_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSupportedCapabilities_ARM(subject SupportedCapabilities_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SupportedCapabilities_ARM
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

// Generator of SupportedCapabilities_ARM instances for property testing - lazily instantiated by
// SupportedCapabilities_ARMGenerator()
var supportedCapabilities_ARMGenerator gopter.Gen

// SupportedCapabilities_ARMGenerator returns a generator of SupportedCapabilities_ARM instances for property testing.
func SupportedCapabilities_ARMGenerator() gopter.Gen {
	if supportedCapabilities_ARMGenerator != nil {
		return supportedCapabilities_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSupportedCapabilities_ARM(generators)
	supportedCapabilities_ARMGenerator = gen.Struct(reflect.TypeOf(SupportedCapabilities_ARM{}), generators)

	return supportedCapabilities_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSupportedCapabilities_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSupportedCapabilities_ARM(gens map[string]gopter.Gen) {
	gens["AcceleratedNetwork"] = gen.PtrOf(gen.Bool())
	gens["Architecture"] = gen.PtrOf(gen.OneConstOf(SupportedCapabilities_Architecture_Arm64, SupportedCapabilities_Architecture_X64))
	gens["DiskControllerTypes"] = gen.PtrOf(gen.AlphaString())
}
