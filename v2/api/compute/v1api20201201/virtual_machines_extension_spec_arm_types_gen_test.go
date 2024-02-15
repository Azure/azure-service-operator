// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201201

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

func Test_VirtualMachines_Extension_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachines_Extension_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachines_Extension_Spec_ARM, VirtualMachines_Extension_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachines_Extension_Spec_ARM runs a test to see if a specific instance of VirtualMachines_Extension_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachines_Extension_Spec_ARM(subject VirtualMachines_Extension_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachines_Extension_Spec_ARM
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

// Generator of VirtualMachines_Extension_Spec_ARM instances for property testing - lazily instantiated by
// VirtualMachines_Extension_Spec_ARMGenerator()
var virtualMachines_Extension_Spec_ARMGenerator gopter.Gen

// VirtualMachines_Extension_Spec_ARMGenerator returns a generator of VirtualMachines_Extension_Spec_ARM instances for property testing.
// We first initialize virtualMachines_Extension_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachines_Extension_Spec_ARMGenerator() gopter.Gen {
	if virtualMachines_Extension_Spec_ARMGenerator != nil {
		return virtualMachines_Extension_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachines_Extension_Spec_ARM(generators)
	virtualMachines_Extension_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachines_Extension_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachines_Extension_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualMachines_Extension_Spec_ARM(generators)
	virtualMachines_Extension_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachines_Extension_Spec_ARM{}), generators)

	return virtualMachines_Extension_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachines_Extension_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachines_Extension_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachines_Extension_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachines_Extension_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualMachineExtensionProperties_ARMGenerator())
}

func Test_VirtualMachineExtensionProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineExtensionProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineExtensionProperties_ARM, VirtualMachineExtensionProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineExtensionProperties_ARM runs a test to see if a specific instance of VirtualMachineExtensionProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineExtensionProperties_ARM(subject VirtualMachineExtensionProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineExtensionProperties_ARM
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

// Generator of VirtualMachineExtensionProperties_ARM instances for property testing - lazily instantiated by
// VirtualMachineExtensionProperties_ARMGenerator()
var virtualMachineExtensionProperties_ARMGenerator gopter.Gen

// VirtualMachineExtensionProperties_ARMGenerator returns a generator of VirtualMachineExtensionProperties_ARM instances for property testing.
// We first initialize virtualMachineExtensionProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineExtensionProperties_ARMGenerator() gopter.Gen {
	if virtualMachineExtensionProperties_ARMGenerator != nil {
		return virtualMachineExtensionProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionProperties_ARM(generators)
	virtualMachineExtensionProperties_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineExtensionProperties_ARM(generators)
	virtualMachineExtensionProperties_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionProperties_ARM{}), generators)

	return virtualMachineExtensionProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineExtensionProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineExtensionProperties_ARM(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["EnableAutomaticUpgrade"] = gen.PtrOf(gen.Bool())
	gens["ForceUpdateTag"] = gen.PtrOf(gen.AlphaString())
	gens["ProtectedSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineExtensionProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineExtensionProperties_ARM(gens map[string]gopter.Gen) {
	gens["InstanceView"] = gen.PtrOf(VirtualMachineExtensionInstanceView_ARMGenerator())
}

func Test_VirtualMachineExtensionInstanceView_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineExtensionInstanceView_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineExtensionInstanceView_ARM, VirtualMachineExtensionInstanceView_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineExtensionInstanceView_ARM runs a test to see if a specific instance of VirtualMachineExtensionInstanceView_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineExtensionInstanceView_ARM(subject VirtualMachineExtensionInstanceView_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineExtensionInstanceView_ARM
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

// Generator of VirtualMachineExtensionInstanceView_ARM instances for property testing - lazily instantiated by
// VirtualMachineExtensionInstanceView_ARMGenerator()
var virtualMachineExtensionInstanceView_ARMGenerator gopter.Gen

// VirtualMachineExtensionInstanceView_ARMGenerator returns a generator of VirtualMachineExtensionInstanceView_ARM instances for property testing.
// We first initialize virtualMachineExtensionInstanceView_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineExtensionInstanceView_ARMGenerator() gopter.Gen {
	if virtualMachineExtensionInstanceView_ARMGenerator != nil {
		return virtualMachineExtensionInstanceView_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_ARM(generators)
	virtualMachineExtensionInstanceView_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_ARM(generators)
	virtualMachineExtensionInstanceView_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineExtensionInstanceView_ARM{}), generators)

	return virtualMachineExtensionInstanceView_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineExtensionInstanceView_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineExtensionInstanceView_ARM(gens map[string]gopter.Gen) {
	gens["Statuses"] = gen.SliceOf(InstanceViewStatus_ARMGenerator())
	gens["Substatuses"] = gen.SliceOf(InstanceViewStatus_ARMGenerator())
}

func Test_InstanceViewStatus_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InstanceViewStatus_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInstanceViewStatus_ARM, InstanceViewStatus_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInstanceViewStatus_ARM runs a test to see if a specific instance of InstanceViewStatus_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInstanceViewStatus_ARM(subject InstanceViewStatus_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InstanceViewStatus_ARM
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

// Generator of InstanceViewStatus_ARM instances for property testing - lazily instantiated by
// InstanceViewStatus_ARMGenerator()
var instanceViewStatus_ARMGenerator gopter.Gen

// InstanceViewStatus_ARMGenerator returns a generator of InstanceViewStatus_ARM instances for property testing.
func InstanceViewStatus_ARMGenerator() gopter.Gen {
	if instanceViewStatus_ARMGenerator != nil {
		return instanceViewStatus_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInstanceViewStatus_ARM(generators)
	instanceViewStatus_ARMGenerator = gen.Struct(reflect.TypeOf(InstanceViewStatus_ARM{}), generators)

	return instanceViewStatus_ARMGenerator
}

// AddIndependentPropertyGeneratorsForInstanceViewStatus_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInstanceViewStatus_ARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayStatus"] = gen.PtrOf(gen.AlphaString())
	gens["Level"] = gen.PtrOf(gen.OneConstOf(InstanceViewStatus_Level_Error, InstanceViewStatus_Level_Info, InstanceViewStatus_Level_Warning))
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Time"] = gen.PtrOf(gen.AlphaString())
}
