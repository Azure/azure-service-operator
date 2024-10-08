// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

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

func Test_AFDOriginProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AFDOriginProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAFDOriginProperties_ARM, AFDOriginProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAFDOriginProperties_ARM runs a test to see if a specific instance of AFDOriginProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAFDOriginProperties_ARM(subject AFDOriginProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AFDOriginProperties_ARM
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

// Generator of AFDOriginProperties_ARM instances for property testing - lazily instantiated by
// AFDOriginProperties_ARMGenerator()
var afdOriginProperties_ARMGenerator gopter.Gen

// AFDOriginProperties_ARMGenerator returns a generator of AFDOriginProperties_ARM instances for property testing.
// We first initialize afdOriginProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AFDOriginProperties_ARMGenerator() gopter.Gen {
	if afdOriginProperties_ARMGenerator != nil {
		return afdOriginProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDOriginProperties_ARM(generators)
	afdOriginProperties_ARMGenerator = gen.Struct(reflect.TypeOf(AFDOriginProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDOriginProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForAFDOriginProperties_ARM(generators)
	afdOriginProperties_ARMGenerator = gen.Struct(reflect.TypeOf(AFDOriginProperties_ARM{}), generators)

	return afdOriginProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAFDOriginProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAFDOriginProperties_ARM(gens map[string]gopter.Gen) {
	gens["EnabledState"] = gen.PtrOf(gen.OneConstOf(AFDOriginProperties_EnabledState_ARM_Disabled, AFDOriginProperties_EnabledState_ARM_Enabled))
	gens["EnforceCertificateNameCheck"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["HttpPort"] = gen.PtrOf(gen.Int())
	gens["HttpsPort"] = gen.PtrOf(gen.Int())
	gens["OriginHostHeader"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAFDOriginProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAFDOriginProperties_ARM(gens map[string]gopter.Gen) {
	gens["AzureOrigin"] = gen.PtrOf(ResourceReference_ARMGenerator())
	gens["SharedPrivateLinkResource"] = gen.PtrOf(SharedPrivateLinkResourceProperties_ARMGenerator())
}

func Test_Profiles_OriginGroups_Origin_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_OriginGroups_Origin_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_OriginGroups_Origin_Spec_ARM, Profiles_OriginGroups_Origin_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_OriginGroups_Origin_Spec_ARM runs a test to see if a specific instance of Profiles_OriginGroups_Origin_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_OriginGroups_Origin_Spec_ARM(subject Profiles_OriginGroups_Origin_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_OriginGroups_Origin_Spec_ARM
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

// Generator of Profiles_OriginGroups_Origin_Spec_ARM instances for property testing - lazily instantiated by
// Profiles_OriginGroups_Origin_Spec_ARMGenerator()
var profiles_OriginGroups_Origin_Spec_ARMGenerator gopter.Gen

// Profiles_OriginGroups_Origin_Spec_ARMGenerator returns a generator of Profiles_OriginGroups_Origin_Spec_ARM instances for property testing.
// We first initialize profiles_OriginGroups_Origin_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_OriginGroups_Origin_Spec_ARMGenerator() gopter.Gen {
	if profiles_OriginGroups_Origin_Spec_ARMGenerator != nil {
		return profiles_OriginGroups_Origin_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_OriginGroups_Origin_Spec_ARM(generators)
	profiles_OriginGroups_Origin_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_OriginGroups_Origin_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_OriginGroups_Origin_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForProfiles_OriginGroups_Origin_Spec_ARM(generators)
	profiles_OriginGroups_Origin_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_OriginGroups_Origin_Spec_ARM{}), generators)

	return profiles_OriginGroups_Origin_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_OriginGroups_Origin_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_OriginGroups_Origin_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForProfiles_OriginGroups_Origin_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_OriginGroups_Origin_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AFDOriginProperties_ARMGenerator())
}

func Test_SharedPrivateLinkResourceProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResourceProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResourceProperties_ARM, SharedPrivateLinkResourceProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResourceProperties_ARM runs a test to see if a specific instance of SharedPrivateLinkResourceProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResourceProperties_ARM(subject SharedPrivateLinkResourceProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResourceProperties_ARM
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

// Generator of SharedPrivateLinkResourceProperties_ARM instances for property testing - lazily instantiated by
// SharedPrivateLinkResourceProperties_ARMGenerator()
var sharedPrivateLinkResourceProperties_ARMGenerator gopter.Gen

// SharedPrivateLinkResourceProperties_ARMGenerator returns a generator of SharedPrivateLinkResourceProperties_ARM instances for property testing.
// We first initialize sharedPrivateLinkResourceProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SharedPrivateLinkResourceProperties_ARMGenerator() gopter.Gen {
	if sharedPrivateLinkResourceProperties_ARMGenerator != nil {
		return sharedPrivateLinkResourceProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties_ARM(generators)
	sharedPrivateLinkResourceProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties_ARM(generators)
	sharedPrivateLinkResourceProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceProperties_ARM{}), generators)

	return sharedPrivateLinkResourceProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties_ARM(gens map[string]gopter.Gen) {
	gens["GroupId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateLinkLocation"] = gen.PtrOf(gen.AlphaString())
	gens["RequestMessage"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		SharedPrivateLinkResourceProperties_Status_ARM_Approved,
		SharedPrivateLinkResourceProperties_Status_ARM_Disconnected,
		SharedPrivateLinkResourceProperties_Status_ARM_Pending,
		SharedPrivateLinkResourceProperties_Status_ARM_Rejected,
		SharedPrivateLinkResourceProperties_Status_ARM_Timeout))
}

// AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties_ARM(gens map[string]gopter.Gen) {
	gens["PrivateLink"] = gen.PtrOf(ResourceReference_ARMGenerator())
}
