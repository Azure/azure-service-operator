// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220301

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

func Test_Image_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Image_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageSTATUSARM, ImageSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageSTATUSARM runs a test to see if a specific instance of Image_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageSTATUSARM(subject Image_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Image_STATUSARM
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

// Generator of Image_STATUSARM instances for property testing - lazily instantiated by ImageSTATUSARMGenerator()
var imageSTATUSARMGenerator gopter.Gen

// ImageSTATUSARMGenerator returns a generator of Image_STATUSARM instances for property testing.
// We first initialize imageSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageSTATUSARMGenerator() gopter.Gen {
	if imageSTATUSARMGenerator != nil {
		return imageSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageSTATUSARM(generators)
	imageSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Image_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForImageSTATUSARM(generators)
	imageSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Image_STATUSARM{}), generators)

	return imageSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForImageSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImageSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageSTATUSARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationSTATUSARMGenerator())
	gens["Properties"] = gen.PtrOf(ImagePropertiesSTATUSARMGenerator())
}

func Test_ExtendedLocation_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExtendedLocation_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtendedLocationSTATUSARM, ExtendedLocationSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtendedLocationSTATUSARM runs a test to see if a specific instance of ExtendedLocation_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExtendedLocationSTATUSARM(subject ExtendedLocation_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExtendedLocation_STATUSARM
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

// Generator of ExtendedLocation_STATUSARM instances for property testing - lazily instantiated by
// ExtendedLocationSTATUSARMGenerator()
var extendedLocationSTATUSARMGenerator gopter.Gen

// ExtendedLocationSTATUSARMGenerator returns a generator of ExtendedLocation_STATUSARM instances for property testing.
func ExtendedLocationSTATUSARMGenerator() gopter.Gen {
	if extendedLocationSTATUSARMGenerator != nil {
		return extendedLocationSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtendedLocationSTATUSARM(generators)
	extendedLocationSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ExtendedLocation_STATUSARM{}), generators)

	return extendedLocationSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForExtendedLocationSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtendedLocationSTATUSARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ExtendedLocationType_STATUS_EdgeZone))
}

func Test_ImageProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImagePropertiesSTATUSARM, ImagePropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImagePropertiesSTATUSARM runs a test to see if a specific instance of ImageProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImagePropertiesSTATUSARM(subject ImageProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageProperties_STATUSARM
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

// Generator of ImageProperties_STATUSARM instances for property testing - lazily instantiated by
// ImagePropertiesSTATUSARMGenerator()
var imagePropertiesSTATUSARMGenerator gopter.Gen

// ImagePropertiesSTATUSARMGenerator returns a generator of ImageProperties_STATUSARM instances for property testing.
// We first initialize imagePropertiesSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImagePropertiesSTATUSARMGenerator() gopter.Gen {
	if imagePropertiesSTATUSARMGenerator != nil {
		return imagePropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImagePropertiesSTATUSARM(generators)
	imagePropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImageProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImagePropertiesSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForImagePropertiesSTATUSARM(generators)
	imagePropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImageProperties_STATUSARM{}), generators)

	return imagePropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForImagePropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImagePropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(HyperVGenerationType_STATUS_V1, HyperVGenerationType_STATUS_V2))
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImagePropertiesSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImagePropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["SourceVirtualMachine"] = gen.PtrOf(SubResourceSTATUSARMGenerator())
	gens["StorageProfile"] = gen.PtrOf(ImageStorageProfileSTATUSARMGenerator())
}

func Test_ImageStorageProfile_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageStorageProfile_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageStorageProfileSTATUSARM, ImageStorageProfileSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageStorageProfileSTATUSARM runs a test to see if a specific instance of ImageStorageProfile_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageStorageProfileSTATUSARM(subject ImageStorageProfile_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageStorageProfile_STATUSARM
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

// Generator of ImageStorageProfile_STATUSARM instances for property testing - lazily instantiated by
// ImageStorageProfileSTATUSARMGenerator()
var imageStorageProfileSTATUSARMGenerator gopter.Gen

// ImageStorageProfileSTATUSARMGenerator returns a generator of ImageStorageProfile_STATUSARM instances for property testing.
// We first initialize imageStorageProfileSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageStorageProfileSTATUSARMGenerator() gopter.Gen {
	if imageStorageProfileSTATUSARMGenerator != nil {
		return imageStorageProfileSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfileSTATUSARM(generators)
	imageStorageProfileSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfileSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForImageStorageProfileSTATUSARM(generators)
	imageStorageProfileSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile_STATUSARM{}), generators)

	return imageStorageProfileSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForImageStorageProfileSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageStorageProfileSTATUSARM(gens map[string]gopter.Gen) {
	gens["ZoneResilient"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForImageStorageProfileSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageStorageProfileSTATUSARM(gens map[string]gopter.Gen) {
	gens["DataDisks"] = gen.SliceOf(ImageDataDiskSTATUSARMGenerator())
	gens["OsDisk"] = gen.PtrOf(ImageOSDiskSTATUSARMGenerator())
}

func Test_SubResource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResourceSTATUSARM, SubResourceSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResourceSTATUSARM runs a test to see if a specific instance of SubResource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResourceSTATUSARM(subject SubResource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_STATUSARM
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

// Generator of SubResource_STATUSARM instances for property testing - lazily instantiated by
// SubResourceSTATUSARMGenerator()
var subResourceSTATUSARMGenerator gopter.Gen

// SubResourceSTATUSARMGenerator returns a generator of SubResource_STATUSARM instances for property testing.
func SubResourceSTATUSARMGenerator() gopter.Gen {
	if subResourceSTATUSARMGenerator != nil {
		return subResourceSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResourceSTATUSARM(generators)
	subResourceSTATUSARMGenerator = gen.Struct(reflect.TypeOf(SubResource_STATUSARM{}), generators)

	return subResourceSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSubResourceSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResourceSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ImageDataDisk_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageDataDisk_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageDataDiskSTATUSARM, ImageDataDiskSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageDataDiskSTATUSARM runs a test to see if a specific instance of ImageDataDisk_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageDataDiskSTATUSARM(subject ImageDataDisk_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageDataDisk_STATUSARM
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

// Generator of ImageDataDisk_STATUSARM instances for property testing - lazily instantiated by
// ImageDataDiskSTATUSARMGenerator()
var imageDataDiskSTATUSARMGenerator gopter.Gen

// ImageDataDiskSTATUSARMGenerator returns a generator of ImageDataDisk_STATUSARM instances for property testing.
// We first initialize imageDataDiskSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageDataDiskSTATUSARMGenerator() gopter.Gen {
	if imageDataDiskSTATUSARMGenerator != nil {
		return imageDataDiskSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDiskSTATUSARM(generators)
	imageDataDiskSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDiskSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForImageDataDiskSTATUSARM(generators)
	imageDataDiskSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk_STATUSARM{}), generators)

	return imageDataDiskSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForImageDataDiskSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageDataDiskSTATUSARM(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.OneConstOf(ImageDataDiskSTATUSCaching_None, ImageDataDiskSTATUSCaching_ReadOnly, ImageDataDiskSTATUSCaching_ReadWrite))
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["Lun"] = gen.PtrOf(gen.Int())
	gens["StorageAccountType"] = gen.PtrOf(gen.OneConstOf(
		StorageAccountType_STATUS_PremiumLRS,
		StorageAccountType_STATUS_PremiumV2LRS,
		StorageAccountType_STATUS_PremiumZRS,
		StorageAccountType_STATUS_StandardLRS,
		StorageAccountType_STATUS_StandardSSDLRS,
		StorageAccountType_STATUS_StandardSSDZRS,
		StorageAccountType_STATUS_UltraSSDLRS))
}

// AddRelatedPropertyGeneratorsForImageDataDiskSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageDataDiskSTATUSARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResourceSTATUSARMGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResourceSTATUSARMGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResourceSTATUSARMGenerator())
}

func Test_ImageOSDisk_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageOSDisk_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageOSDiskSTATUSARM, ImageOSDiskSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageOSDiskSTATUSARM runs a test to see if a specific instance of ImageOSDisk_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageOSDiskSTATUSARM(subject ImageOSDisk_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageOSDisk_STATUSARM
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

// Generator of ImageOSDisk_STATUSARM instances for property testing - lazily instantiated by
// ImageOSDiskSTATUSARMGenerator()
var imageOSDiskSTATUSARMGenerator gopter.Gen

// ImageOSDiskSTATUSARMGenerator returns a generator of ImageOSDisk_STATUSARM instances for property testing.
// We first initialize imageOSDiskSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageOSDiskSTATUSARMGenerator() gopter.Gen {
	if imageOSDiskSTATUSARMGenerator != nil {
		return imageOSDiskSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDiskSTATUSARM(generators)
	imageOSDiskSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDiskSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForImageOSDiskSTATUSARM(generators)
	imageOSDiskSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk_STATUSARM{}), generators)

	return imageOSDiskSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForImageOSDiskSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageOSDiskSTATUSARM(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.OneConstOf(ImageOSDiskSTATUSCaching_None, ImageOSDiskSTATUSCaching_ReadOnly, ImageOSDiskSTATUSCaching_ReadWrite))
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsState"] = gen.PtrOf(gen.OneConstOf(ImageOSDiskSTATUSOsState_Generalized, ImageOSDiskSTATUSOsState_Specialized))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(ImageOSDiskSTATUSOsType_Linux, ImageOSDiskSTATUSOsType_Windows))
	gens["StorageAccountType"] = gen.PtrOf(gen.OneConstOf(
		StorageAccountType_STATUS_PremiumLRS,
		StorageAccountType_STATUS_PremiumV2LRS,
		StorageAccountType_STATUS_PremiumZRS,
		StorageAccountType_STATUS_StandardLRS,
		StorageAccountType_STATUS_StandardSSDLRS,
		StorageAccountType_STATUS_StandardSSDZRS,
		StorageAccountType_STATUS_UltraSSDLRS))
}

// AddRelatedPropertyGeneratorsForImageOSDiskSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageOSDiskSTATUSARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResourceSTATUSARMGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResourceSTATUSARMGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResourceSTATUSARMGenerator())
}
