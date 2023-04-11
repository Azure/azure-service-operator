// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

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

func Test_Servers_AdvancedThreatProtectionSetting_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_AdvancedThreatProtectionSetting_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_AdvancedThreatProtectionSetting_STATUS_ARM, Servers_AdvancedThreatProtectionSetting_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_AdvancedThreatProtectionSetting_STATUS_ARM runs a test to see if a specific instance of Servers_AdvancedThreatProtectionSetting_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_AdvancedThreatProtectionSetting_STATUS_ARM(subject Servers_AdvancedThreatProtectionSetting_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_AdvancedThreatProtectionSetting_STATUS_ARM
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

// Generator of Servers_AdvancedThreatProtectionSetting_STATUS_ARM instances for property testing - lazily instantiated
// by Servers_AdvancedThreatProtectionSetting_STATUS_ARMGenerator()
var servers_AdvancedThreatProtectionSetting_STATUS_ARMGenerator gopter.Gen

// Servers_AdvancedThreatProtectionSetting_STATUS_ARMGenerator returns a generator of Servers_AdvancedThreatProtectionSetting_STATUS_ARM instances for property testing.
// We first initialize servers_AdvancedThreatProtectionSetting_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_AdvancedThreatProtectionSetting_STATUS_ARMGenerator() gopter.Gen {
	if servers_AdvancedThreatProtectionSetting_STATUS_ARMGenerator != nil {
		return servers_AdvancedThreatProtectionSetting_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_AdvancedThreatProtectionSetting_STATUS_ARM(generators)
	servers_AdvancedThreatProtectionSetting_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_AdvancedThreatProtectionSetting_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_AdvancedThreatProtectionSetting_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForServers_AdvancedThreatProtectionSetting_STATUS_ARM(generators)
	servers_AdvancedThreatProtectionSetting_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_AdvancedThreatProtectionSetting_STATUS_ARM{}), generators)

	return servers_AdvancedThreatProtectionSetting_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServers_AdvancedThreatProtectionSetting_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_AdvancedThreatProtectionSetting_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServers_AdvancedThreatProtectionSetting_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_AdvancedThreatProtectionSetting_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AdvancedThreatProtectionProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_AdvancedThreatProtectionProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AdvancedThreatProtectionProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAdvancedThreatProtectionProperties_STATUS_ARM, AdvancedThreatProtectionProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAdvancedThreatProtectionProperties_STATUS_ARM runs a test to see if a specific instance of AdvancedThreatProtectionProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAdvancedThreatProtectionProperties_STATUS_ARM(subject AdvancedThreatProtectionProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AdvancedThreatProtectionProperties_STATUS_ARM
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

// Generator of AdvancedThreatProtectionProperties_STATUS_ARM instances for property testing - lazily instantiated by
// AdvancedThreatProtectionProperties_STATUS_ARMGenerator()
var advancedThreatProtectionProperties_STATUS_ARMGenerator gopter.Gen

// AdvancedThreatProtectionProperties_STATUS_ARMGenerator returns a generator of AdvancedThreatProtectionProperties_STATUS_ARM instances for property testing.
func AdvancedThreatProtectionProperties_STATUS_ARMGenerator() gopter.Gen {
	if advancedThreatProtectionProperties_STATUS_ARMGenerator != nil {
		return advancedThreatProtectionProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAdvancedThreatProtectionProperties_STATUS_ARM(generators)
	advancedThreatProtectionProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AdvancedThreatProtectionProperties_STATUS_ARM{}), generators)

	return advancedThreatProtectionProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAdvancedThreatProtectionProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAdvancedThreatProtectionProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreationTime"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(AdvancedThreatProtectionProperties_State_STATUS_Disabled, AdvancedThreatProtectionProperties_State_STATUS_Enabled, AdvancedThreatProtectionProperties_State_STATUS_New))
}

func Test_SystemData_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS_ARM, SystemData_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS_ARM runs a test to see if a specific instance of SystemData_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS_ARM(subject SystemData_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS_ARM
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

// Generator of SystemData_STATUS_ARM instances for property testing - lazily instantiated by
// SystemData_STATUS_ARMGenerator()
var systemData_STATUS_ARMGenerator gopter.Gen

// SystemData_STATUS_ARMGenerator returns a generator of SystemData_STATUS_ARM instances for property testing.
func SystemData_STATUS_ARMGenerator() gopter.Gen {
	if systemData_STATUS_ARMGenerator != nil {
		return systemData_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(generators)
	systemData_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS_ARM{}), generators)

	return systemData_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_Application,
		SystemData_CreatedByType_STATUS_Key,
		SystemData_CreatedByType_STATUS_ManagedIdentity,
		SystemData_CreatedByType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_Application,
		SystemData_LastModifiedByType_STATUS_Key,
		SystemData_LastModifiedByType_STATUS_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_User))
}
