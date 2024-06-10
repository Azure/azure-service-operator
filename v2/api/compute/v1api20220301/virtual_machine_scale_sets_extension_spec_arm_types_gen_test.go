// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220301

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

func Test_KeyVaultSecretReference_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultSecretReference_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultSecretReference_ARM, KeyVaultSecretReference_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultSecretReference_ARM runs a test to see if a specific instance of KeyVaultSecretReference_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultSecretReference_ARM(subject KeyVaultSecretReference_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultSecretReference_ARM
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

// Generator of KeyVaultSecretReference_ARM instances for property testing - lazily instantiated by
// KeyVaultSecretReference_ARMGenerator()
var keyVaultSecretReference_ARMGenerator gopter.Gen

// KeyVaultSecretReference_ARMGenerator returns a generator of KeyVaultSecretReference_ARM instances for property testing.
// We first initialize keyVaultSecretReference_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultSecretReference_ARMGenerator() gopter.Gen {
	if keyVaultSecretReference_ARMGenerator != nil {
		return keyVaultSecretReference_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultSecretReference_ARM(generators)
	keyVaultSecretReference_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultSecretReference_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultSecretReference_ARM(generators)
	AddRelatedPropertyGeneratorsForKeyVaultSecretReference_ARM(generators)
	keyVaultSecretReference_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultSecretReference_ARM{}), generators)

	return keyVaultSecretReference_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultSecretReference_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultSecretReference_ARM(gens map[string]gopter.Gen) {
	gens["SecretUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultSecretReference_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultSecretReference_ARM(gens map[string]gopter.Gen) {
	gens["SourceVault"] = gen.PtrOf(SubResource_ARMGenerator())
}

func Test_VirtualMachineScaleSetExtensionProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineScaleSetExtensionProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineScaleSetExtensionProperties_ARM, VirtualMachineScaleSetExtensionProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineScaleSetExtensionProperties_ARM runs a test to see if a specific instance of VirtualMachineScaleSetExtensionProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineScaleSetExtensionProperties_ARM(subject VirtualMachineScaleSetExtensionProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineScaleSetExtensionProperties_ARM
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

// Generator of VirtualMachineScaleSetExtensionProperties_ARM instances for property testing - lazily instantiated by
// VirtualMachineScaleSetExtensionProperties_ARMGenerator()
var virtualMachineScaleSetExtensionProperties_ARMGenerator gopter.Gen

// VirtualMachineScaleSetExtensionProperties_ARMGenerator returns a generator of VirtualMachineScaleSetExtensionProperties_ARM instances for property testing.
// We first initialize virtualMachineScaleSetExtensionProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineScaleSetExtensionProperties_ARMGenerator() gopter.Gen {
	if virtualMachineScaleSetExtensionProperties_ARMGenerator != nil {
		return virtualMachineScaleSetExtensionProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_ARM(generators)
	virtualMachineScaleSetExtensionProperties_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetExtensionProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_ARM(generators)
	virtualMachineScaleSetExtensionProperties_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetExtensionProperties_ARM{}), generators)

	return virtualMachineScaleSetExtensionProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_ARM(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["EnableAutomaticUpgrade"] = gen.PtrOf(gen.Bool())
	gens["ForceUpdateTag"] = gen.PtrOf(gen.AlphaString())
	gens["ProtectedSettings"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["ProvisionAfterExtensions"] = gen.SliceOf(gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["SuppressFailures"] = gen.PtrOf(gen.Bool())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_ARM(gens map[string]gopter.Gen) {
	gens["ProtectedSettingsFromKeyVault"] = gen.PtrOf(KeyVaultSecretReference_ARMGenerator())
}

func Test_VirtualMachineScaleSets_Extension_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineScaleSets_Extension_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineScaleSets_Extension_Spec_ARM, VirtualMachineScaleSets_Extension_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineScaleSets_Extension_Spec_ARM runs a test to see if a specific instance of VirtualMachineScaleSets_Extension_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineScaleSets_Extension_Spec_ARM(subject VirtualMachineScaleSets_Extension_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineScaleSets_Extension_Spec_ARM
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

// Generator of VirtualMachineScaleSets_Extension_Spec_ARM instances for property testing - lazily instantiated by
// VirtualMachineScaleSets_Extension_Spec_ARMGenerator()
var virtualMachineScaleSets_Extension_Spec_ARMGenerator gopter.Gen

// VirtualMachineScaleSets_Extension_Spec_ARMGenerator returns a generator of VirtualMachineScaleSets_Extension_Spec_ARM instances for property testing.
// We first initialize virtualMachineScaleSets_Extension_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineScaleSets_Extension_Spec_ARMGenerator() gopter.Gen {
	if virtualMachineScaleSets_Extension_Spec_ARMGenerator != nil {
		return virtualMachineScaleSets_Extension_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSets_Extension_Spec_ARM(generators)
	virtualMachineScaleSets_Extension_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSets_Extension_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSets_Extension_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineScaleSets_Extension_Spec_ARM(generators)
	virtualMachineScaleSets_Extension_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSets_Extension_Spec_ARM{}), generators)

	return virtualMachineScaleSets_Extension_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineScaleSets_Extension_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineScaleSets_Extension_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForVirtualMachineScaleSets_Extension_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineScaleSets_Extension_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualMachineScaleSetExtensionProperties_ARMGenerator())
}
