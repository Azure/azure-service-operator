// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230630

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

func Test_AdministratorProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AdministratorProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAdministratorProperties_STATUS_ARM, AdministratorProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAdministratorProperties_STATUS_ARM runs a test to see if a specific instance of AdministratorProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAdministratorProperties_STATUS_ARM(subject AdministratorProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AdministratorProperties_STATUS_ARM
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

// Generator of AdministratorProperties_STATUS_ARM instances for property testing - lazily instantiated by
// AdministratorProperties_STATUS_ARMGenerator()
var administratorProperties_STATUS_ARMGenerator gopter.Gen

// AdministratorProperties_STATUS_ARMGenerator returns a generator of AdministratorProperties_STATUS_ARM instances for property testing.
func AdministratorProperties_STATUS_ARMGenerator() gopter.Gen {
	if administratorProperties_STATUS_ARMGenerator != nil {
		return administratorProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAdministratorProperties_STATUS_ARM(generators)
	administratorProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AdministratorProperties_STATUS_ARM{}), generators)

	return administratorProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAdministratorProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAdministratorProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AdministratorType"] = gen.PtrOf(gen.OneConstOf(AdministratorProperties_AdministratorType_STATUS_ARM_ActiveDirectory))
	gens["IdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Login"] = gen.PtrOf(gen.AlphaString())
	gens["Sid"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_FlexibleServersAdministrator_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersAdministrator_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersAdministrator_STATUS_ARM, FlexibleServersAdministrator_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersAdministrator_STATUS_ARM runs a test to see if a specific instance of FlexibleServersAdministrator_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersAdministrator_STATUS_ARM(subject FlexibleServersAdministrator_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersAdministrator_STATUS_ARM
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

// Generator of FlexibleServersAdministrator_STATUS_ARM instances for property testing - lazily instantiated by
// FlexibleServersAdministrator_STATUS_ARMGenerator()
var flexibleServersAdministrator_STATUS_ARMGenerator gopter.Gen

// FlexibleServersAdministrator_STATUS_ARMGenerator returns a generator of FlexibleServersAdministrator_STATUS_ARM instances for property testing.
// We first initialize flexibleServersAdministrator_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersAdministrator_STATUS_ARMGenerator() gopter.Gen {
	if flexibleServersAdministrator_STATUS_ARMGenerator != nil {
		return flexibleServersAdministrator_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_STATUS_ARM(generators)
	flexibleServersAdministrator_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdministrator_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_STATUS_ARM(generators)
	flexibleServersAdministrator_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdministrator_STATUS_ARM{}), generators)

	return flexibleServersAdministrator_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AdministratorProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}
