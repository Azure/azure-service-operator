// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_AFDOriginProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AFDOriginProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAFDOriginProperties, AFDOriginPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAFDOriginProperties runs a test to see if a specific instance of AFDOriginProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForAFDOriginProperties(subject AFDOriginProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AFDOriginProperties
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

// Generator of AFDOriginProperties instances for property testing - lazily instantiated by
// AFDOriginPropertiesGenerator()
var afdOriginPropertiesGenerator gopter.Gen

// AFDOriginPropertiesGenerator returns a generator of AFDOriginProperties instances for property testing.
// We first initialize afdOriginPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AFDOriginPropertiesGenerator() gopter.Gen {
	if afdOriginPropertiesGenerator != nil {
		return afdOriginPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDOriginProperties(generators)
	afdOriginPropertiesGenerator = gen.Struct(reflect.TypeOf(AFDOriginProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDOriginProperties(generators)
	AddRelatedPropertyGeneratorsForAFDOriginProperties(generators)
	afdOriginPropertiesGenerator = gen.Struct(reflect.TypeOf(AFDOriginProperties{}), generators)

	return afdOriginPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForAFDOriginProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAFDOriginProperties(gens map[string]gopter.Gen) {
	gens["EnabledState"] = gen.PtrOf(gen.OneConstOf(AFDOriginProperties_EnabledState_Disabled, AFDOriginProperties_EnabledState_Enabled))
	gens["EnforceCertificateNameCheck"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["HttpPort"] = gen.PtrOf(gen.Int())
	gens["HttpsPort"] = gen.PtrOf(gen.Int())
	gens["OriginHostHeader"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAFDOriginProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAFDOriginProperties(gens map[string]gopter.Gen) {
	gens["AzureOrigin"] = gen.PtrOf(ResourceReferenceGenerator())
	gens["SharedPrivateLinkResource"] = gen.PtrOf(SharedPrivateLinkResourcePropertiesGenerator())
}

func Test_AfdOrigin_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdOrigin_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdOrigin_Spec, AfdOrigin_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdOrigin_Spec runs a test to see if a specific instance of AfdOrigin_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdOrigin_Spec(subject AfdOrigin_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdOrigin_Spec
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

// Generator of AfdOrigin_Spec instances for property testing - lazily instantiated by AfdOrigin_SpecGenerator()
var afdOrigin_SpecGenerator gopter.Gen

// AfdOrigin_SpecGenerator returns a generator of AfdOrigin_Spec instances for property testing.
// We first initialize afdOrigin_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AfdOrigin_SpecGenerator() gopter.Gen {
	if afdOrigin_SpecGenerator != nil {
		return afdOrigin_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdOrigin_Spec(generators)
	afdOrigin_SpecGenerator = gen.Struct(reflect.TypeOf(AfdOrigin_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdOrigin_Spec(generators)
	AddRelatedPropertyGeneratorsForAfdOrigin_Spec(generators)
	afdOrigin_SpecGenerator = gen.Struct(reflect.TypeOf(AfdOrigin_Spec{}), generators)

	return afdOrigin_SpecGenerator
}

// AddIndependentPropertyGeneratorsForAfdOrigin_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAfdOrigin_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForAfdOrigin_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAfdOrigin_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AFDOriginPropertiesGenerator())
}

func Test_SharedPrivateLinkResourceProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResourceProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResourceProperties, SharedPrivateLinkResourcePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResourceProperties runs a test to see if a specific instance of SharedPrivateLinkResourceProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResourceProperties(subject SharedPrivateLinkResourceProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResourceProperties
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

// Generator of SharedPrivateLinkResourceProperties instances for property testing - lazily instantiated by
// SharedPrivateLinkResourcePropertiesGenerator()
var sharedPrivateLinkResourcePropertiesGenerator gopter.Gen

// SharedPrivateLinkResourcePropertiesGenerator returns a generator of SharedPrivateLinkResourceProperties instances for property testing.
// We first initialize sharedPrivateLinkResourcePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SharedPrivateLinkResourcePropertiesGenerator() gopter.Gen {
	if sharedPrivateLinkResourcePropertiesGenerator != nil {
		return sharedPrivateLinkResourcePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties(generators)
	sharedPrivateLinkResourcePropertiesGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties(generators)
	AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties(generators)
	sharedPrivateLinkResourcePropertiesGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceProperties{}), generators)

	return sharedPrivateLinkResourcePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperties(gens map[string]gopter.Gen) {
	gens["GroupId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateLinkLocation"] = gen.PtrOf(gen.AlphaString())
	gens["RequestMessage"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		SharedPrivateLinkResourceProperties_Status_Approved,
		SharedPrivateLinkResourceProperties_Status_Disconnected,
		SharedPrivateLinkResourceProperties_Status_Pending,
		SharedPrivateLinkResourceProperties_Status_Rejected,
		SharedPrivateLinkResourceProperties_Status_Timeout))
}

// AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceProperties(gens map[string]gopter.Gen) {
	gens["PrivateLink"] = gen.PtrOf(ResourceReferenceGenerator())
}
