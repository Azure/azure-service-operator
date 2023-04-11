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

func Test_Servers_FailoverGroup_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_FailoverGroup_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_FailoverGroup_Spec_ARM, Servers_FailoverGroup_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_FailoverGroup_Spec_ARM runs a test to see if a specific instance of Servers_FailoverGroup_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_FailoverGroup_Spec_ARM(subject Servers_FailoverGroup_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_FailoverGroup_Spec_ARM
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

// Generator of Servers_FailoverGroup_Spec_ARM instances for property testing - lazily instantiated by
// Servers_FailoverGroup_Spec_ARMGenerator()
var servers_FailoverGroup_Spec_ARMGenerator gopter.Gen

// Servers_FailoverGroup_Spec_ARMGenerator returns a generator of Servers_FailoverGroup_Spec_ARM instances for property testing.
// We first initialize servers_FailoverGroup_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_FailoverGroup_Spec_ARMGenerator() gopter.Gen {
	if servers_FailoverGroup_Spec_ARMGenerator != nil {
		return servers_FailoverGroup_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_FailoverGroup_Spec_ARM(generators)
	servers_FailoverGroup_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_FailoverGroup_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_FailoverGroup_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForServers_FailoverGroup_Spec_ARM(generators)
	servers_FailoverGroup_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_FailoverGroup_Spec_ARM{}), generators)

	return servers_FailoverGroup_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServers_FailoverGroup_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_FailoverGroup_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServers_FailoverGroup_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_FailoverGroup_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FailoverGroupProperties_ARMGenerator())
}

func Test_FailoverGroupProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FailoverGroupProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFailoverGroupProperties_ARM, FailoverGroupProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFailoverGroupProperties_ARM runs a test to see if a specific instance of FailoverGroupProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFailoverGroupProperties_ARM(subject FailoverGroupProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FailoverGroupProperties_ARM
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

// Generator of FailoverGroupProperties_ARM instances for property testing - lazily instantiated by
// FailoverGroupProperties_ARMGenerator()
var failoverGroupProperties_ARMGenerator gopter.Gen

// FailoverGroupProperties_ARMGenerator returns a generator of FailoverGroupProperties_ARM instances for property testing.
// We first initialize failoverGroupProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FailoverGroupProperties_ARMGenerator() gopter.Gen {
	if failoverGroupProperties_ARMGenerator != nil {
		return failoverGroupProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFailoverGroupProperties_ARM(generators)
	failoverGroupProperties_ARMGenerator = gen.Struct(reflect.TypeOf(FailoverGroupProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFailoverGroupProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForFailoverGroupProperties_ARM(generators)
	failoverGroupProperties_ARMGenerator = gen.Struct(reflect.TypeOf(FailoverGroupProperties_ARM{}), generators)

	return failoverGroupProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFailoverGroupProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFailoverGroupProperties_ARM(gens map[string]gopter.Gen) {
	gens["Databases"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFailoverGroupProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFailoverGroupProperties_ARM(gens map[string]gopter.Gen) {
	gens["PartnerServers"] = gen.SliceOf(PartnerInfo_ARMGenerator())
	gens["ReadOnlyEndpoint"] = gen.PtrOf(FailoverGroupReadOnlyEndpoint_ARMGenerator())
	gens["ReadWriteEndpoint"] = gen.PtrOf(FailoverGroupReadWriteEndpoint_ARMGenerator())
}

func Test_FailoverGroupReadOnlyEndpoint_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FailoverGroupReadOnlyEndpoint_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFailoverGroupReadOnlyEndpoint_ARM, FailoverGroupReadOnlyEndpoint_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFailoverGroupReadOnlyEndpoint_ARM runs a test to see if a specific instance of FailoverGroupReadOnlyEndpoint_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFailoverGroupReadOnlyEndpoint_ARM(subject FailoverGroupReadOnlyEndpoint_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FailoverGroupReadOnlyEndpoint_ARM
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

// Generator of FailoverGroupReadOnlyEndpoint_ARM instances for property testing - lazily instantiated by
// FailoverGroupReadOnlyEndpoint_ARMGenerator()
var failoverGroupReadOnlyEndpoint_ARMGenerator gopter.Gen

// FailoverGroupReadOnlyEndpoint_ARMGenerator returns a generator of FailoverGroupReadOnlyEndpoint_ARM instances for property testing.
func FailoverGroupReadOnlyEndpoint_ARMGenerator() gopter.Gen {
	if failoverGroupReadOnlyEndpoint_ARMGenerator != nil {
		return failoverGroupReadOnlyEndpoint_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFailoverGroupReadOnlyEndpoint_ARM(generators)
	failoverGroupReadOnlyEndpoint_ARMGenerator = gen.Struct(reflect.TypeOf(FailoverGroupReadOnlyEndpoint_ARM{}), generators)

	return failoverGroupReadOnlyEndpoint_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFailoverGroupReadOnlyEndpoint_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFailoverGroupReadOnlyEndpoint_ARM(gens map[string]gopter.Gen) {
	gens["FailoverPolicy"] = gen.PtrOf(gen.OneConstOf(FailoverGroupReadOnlyEndpoint_FailoverPolicy_Disabled, FailoverGroupReadOnlyEndpoint_FailoverPolicy_Enabled))
}

func Test_FailoverGroupReadWriteEndpoint_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FailoverGroupReadWriteEndpoint_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFailoverGroupReadWriteEndpoint_ARM, FailoverGroupReadWriteEndpoint_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFailoverGroupReadWriteEndpoint_ARM runs a test to see if a specific instance of FailoverGroupReadWriteEndpoint_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFailoverGroupReadWriteEndpoint_ARM(subject FailoverGroupReadWriteEndpoint_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FailoverGroupReadWriteEndpoint_ARM
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

// Generator of FailoverGroupReadWriteEndpoint_ARM instances for property testing - lazily instantiated by
// FailoverGroupReadWriteEndpoint_ARMGenerator()
var failoverGroupReadWriteEndpoint_ARMGenerator gopter.Gen

// FailoverGroupReadWriteEndpoint_ARMGenerator returns a generator of FailoverGroupReadWriteEndpoint_ARM instances for property testing.
func FailoverGroupReadWriteEndpoint_ARMGenerator() gopter.Gen {
	if failoverGroupReadWriteEndpoint_ARMGenerator != nil {
		return failoverGroupReadWriteEndpoint_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFailoverGroupReadWriteEndpoint_ARM(generators)
	failoverGroupReadWriteEndpoint_ARMGenerator = gen.Struct(reflect.TypeOf(FailoverGroupReadWriteEndpoint_ARM{}), generators)

	return failoverGroupReadWriteEndpoint_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFailoverGroupReadWriteEndpoint_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFailoverGroupReadWriteEndpoint_ARM(gens map[string]gopter.Gen) {
	gens["FailoverPolicy"] = gen.PtrOf(gen.OneConstOf(FailoverGroupReadWriteEndpoint_FailoverPolicy_Automatic, FailoverGroupReadWriteEndpoint_FailoverPolicy_Manual))
	gens["FailoverWithDataLossGracePeriodMinutes"] = gen.PtrOf(gen.Int())
}

func Test_PartnerInfo_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PartnerInfo_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPartnerInfo_ARM, PartnerInfo_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPartnerInfo_ARM runs a test to see if a specific instance of PartnerInfo_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPartnerInfo_ARM(subject PartnerInfo_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PartnerInfo_ARM
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

// Generator of PartnerInfo_ARM instances for property testing - lazily instantiated by PartnerInfo_ARMGenerator()
var partnerInfo_ARMGenerator gopter.Gen

// PartnerInfo_ARMGenerator returns a generator of PartnerInfo_ARM instances for property testing.
func PartnerInfo_ARMGenerator() gopter.Gen {
	if partnerInfo_ARMGenerator != nil {
		return partnerInfo_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPartnerInfo_ARM(generators)
	partnerInfo_ARMGenerator = gen.Struct(reflect.TypeOf(PartnerInfo_ARM{}), generators)

	return partnerInfo_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPartnerInfo_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPartnerInfo_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
