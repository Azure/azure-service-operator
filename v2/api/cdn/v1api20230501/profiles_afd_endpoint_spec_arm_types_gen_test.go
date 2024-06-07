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

func Test_AFDEndpointProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AFDEndpointProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAFDEndpointProperties_ARM, AFDEndpointProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAFDEndpointProperties_ARM runs a test to see if a specific instance of AFDEndpointProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAFDEndpointProperties_ARM(subject AFDEndpointProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AFDEndpointProperties_ARM
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

// Generator of AFDEndpointProperties_ARM instances for property testing - lazily instantiated by
// AFDEndpointProperties_ARMGenerator()
var afdEndpointProperties_ARMGenerator gopter.Gen

// AFDEndpointProperties_ARMGenerator returns a generator of AFDEndpointProperties_ARM instances for property testing.
func AFDEndpointProperties_ARMGenerator() gopter.Gen {
	if afdEndpointProperties_ARMGenerator != nil {
		return afdEndpointProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDEndpointProperties_ARM(generators)
	afdEndpointProperties_ARMGenerator = gen.Struct(reflect.TypeOf(AFDEndpointProperties_ARM{}), generators)

	return afdEndpointProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAFDEndpointProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAFDEndpointProperties_ARM(gens map[string]gopter.Gen) {
	gens["AutoGeneratedDomainNameLabelScope"] = gen.PtrOf(gen.OneConstOf(
		AutoGeneratedDomainNameLabelScope_NoReuse,
		AutoGeneratedDomainNameLabelScope_ResourceGroupReuse,
		AutoGeneratedDomainNameLabelScope_SubscriptionReuse,
		AutoGeneratedDomainNameLabelScope_TenantReuse))
	gens["EnabledState"] = gen.PtrOf(gen.OneConstOf(AFDEndpointProperties_EnabledState_Disabled, AFDEndpointProperties_EnabledState_Enabled))
}

func Test_Profiles_AfdEndpoint_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_AfdEndpoint_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_AfdEndpoint_Spec_ARM, Profiles_AfdEndpoint_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_AfdEndpoint_Spec_ARM runs a test to see if a specific instance of Profiles_AfdEndpoint_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_AfdEndpoint_Spec_ARM(subject Profiles_AfdEndpoint_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_AfdEndpoint_Spec_ARM
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

// Generator of Profiles_AfdEndpoint_Spec_ARM instances for property testing - lazily instantiated by
// Profiles_AfdEndpoint_Spec_ARMGenerator()
var profiles_AfdEndpoint_Spec_ARMGenerator gopter.Gen

// Profiles_AfdEndpoint_Spec_ARMGenerator returns a generator of Profiles_AfdEndpoint_Spec_ARM instances for property testing.
// We first initialize profiles_AfdEndpoint_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_AfdEndpoint_Spec_ARMGenerator() gopter.Gen {
	if profiles_AfdEndpoint_Spec_ARMGenerator != nil {
		return profiles_AfdEndpoint_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_AfdEndpoint_Spec_ARM(generators)
	profiles_AfdEndpoint_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_AfdEndpoint_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_AfdEndpoint_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForProfiles_AfdEndpoint_Spec_ARM(generators)
	profiles_AfdEndpoint_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_AfdEndpoint_Spec_ARM{}), generators)

	return profiles_AfdEndpoint_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_AfdEndpoint_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_AfdEndpoint_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfiles_AfdEndpoint_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_AfdEndpoint_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AFDEndpointProperties_ARMGenerator())
}
