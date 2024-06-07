// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

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

func Test_PrivateDnsZoneConfig_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZoneConfig_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZoneConfig_STATUS_ARM, PrivateDnsZoneConfig_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZoneConfig_STATUS_ARM runs a test to see if a specific instance of PrivateDnsZoneConfig_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZoneConfig_STATUS_ARM(subject PrivateDnsZoneConfig_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZoneConfig_STATUS_ARM
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

// Generator of PrivateDnsZoneConfig_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateDnsZoneConfig_STATUS_ARMGenerator()
var privateDnsZoneConfig_STATUS_ARMGenerator gopter.Gen

// PrivateDnsZoneConfig_STATUS_ARMGenerator returns a generator of PrivateDnsZoneConfig_STATUS_ARM instances for property testing.
// We first initialize privateDnsZoneConfig_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZoneConfig_STATUS_ARMGenerator() gopter.Gen {
	if privateDnsZoneConfig_STATUS_ARMGenerator != nil {
		return privateDnsZoneConfig_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_STATUS_ARM(generators)
	privateDnsZoneConfig_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneConfig_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig_STATUS_ARM(generators)
	privateDnsZoneConfig_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneConfig_STATUS_ARM{}), generators)

	return privateDnsZoneConfig_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateDnsZonePropertiesFormat_STATUS_ARMGenerator())
}

func Test_PrivateDnsZoneGroupPropertiesFormat_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZoneGroupPropertiesFormat_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZoneGroupPropertiesFormat_STATUS_ARM, PrivateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZoneGroupPropertiesFormat_STATUS_ARM runs a test to see if a specific instance of PrivateDnsZoneGroupPropertiesFormat_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZoneGroupPropertiesFormat_STATUS_ARM(subject PrivateDnsZoneGroupPropertiesFormat_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZoneGroupPropertiesFormat_STATUS_ARM
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

// Generator of PrivateDnsZoneGroupPropertiesFormat_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator()
var privateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator gopter.Gen

// PrivateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator returns a generator of PrivateDnsZoneGroupPropertiesFormat_STATUS_ARM instances for property testing.
// We first initialize privateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator() gopter.Gen {
	if privateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator != nil {
		return privateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat_STATUS_ARM(generators)
	privateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneGroupPropertiesFormat_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat_STATUS_ARM(generators)
	privateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneGroupPropertiesFormat_STATUS_ARM{}), generators)

	return privateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		PrivateEndpointProvisioningState_STATUS_Deleting,
		PrivateEndpointProvisioningState_STATUS_Failed,
		PrivateEndpointProvisioningState_STATUS_Succeeded,
		PrivateEndpointProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrivateDnsZoneConfigs"] = gen.SliceOf(PrivateDnsZoneConfig_STATUS_ARMGenerator())
}

func Test_PrivateDnsZonePropertiesFormat_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonePropertiesFormat_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonePropertiesFormat_STATUS_ARM, PrivateDnsZonePropertiesFormat_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonePropertiesFormat_STATUS_ARM runs a test to see if a specific instance of PrivateDnsZonePropertiesFormat_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonePropertiesFormat_STATUS_ARM(subject PrivateDnsZonePropertiesFormat_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonePropertiesFormat_STATUS_ARM
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

// Generator of PrivateDnsZonePropertiesFormat_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateDnsZonePropertiesFormat_STATUS_ARMGenerator()
var privateDnsZonePropertiesFormat_STATUS_ARMGenerator gopter.Gen

// PrivateDnsZonePropertiesFormat_STATUS_ARMGenerator returns a generator of PrivateDnsZonePropertiesFormat_STATUS_ARM instances for property testing.
// We first initialize privateDnsZonePropertiesFormat_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonePropertiesFormat_STATUS_ARMGenerator() gopter.Gen {
	if privateDnsZonePropertiesFormat_STATUS_ARMGenerator != nil {
		return privateDnsZonePropertiesFormat_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonePropertiesFormat_STATUS_ARM(generators)
	privateDnsZonePropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonePropertiesFormat_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonePropertiesFormat_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonePropertiesFormat_STATUS_ARM(generators)
	privateDnsZonePropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonePropertiesFormat_STATUS_ARM{}), generators)

	return privateDnsZonePropertiesFormat_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonePropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonePropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrivateDnsZoneId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonePropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonePropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["RecordSets"] = gen.SliceOf(RecordSet_STATUS_ARMGenerator())
}

func Test_PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM, PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM runs a test to see if a specific instance of PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM(subject PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM
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

// Generator of PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARMGenerator()
var privateEndpoints_PrivateDnsZoneGroup_STATUS_ARMGenerator gopter.Gen

// PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARMGenerator returns a generator of PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM instances for property testing.
// We first initialize privateEndpoints_PrivateDnsZoneGroup_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARMGenerator() gopter.Gen {
	if privateEndpoints_PrivateDnsZoneGroup_STATUS_ARMGenerator != nil {
		return privateEndpoints_PrivateDnsZoneGroup_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM(generators)
	privateEndpoints_PrivateDnsZoneGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM(generators)
	privateEndpoints_PrivateDnsZoneGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM{}), generators)

	return privateEndpoints_PrivateDnsZoneGroup_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateDnsZoneGroupPropertiesFormat_STATUS_ARMGenerator())
}

func Test_RecordSet_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RecordSet_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRecordSet_STATUS_ARM, RecordSet_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRecordSet_STATUS_ARM runs a test to see if a specific instance of RecordSet_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRecordSet_STATUS_ARM(subject RecordSet_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RecordSet_STATUS_ARM
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

// Generator of RecordSet_STATUS_ARM instances for property testing - lazily instantiated by
// RecordSet_STATUS_ARMGenerator()
var recordSet_STATUS_ARMGenerator gopter.Gen

// RecordSet_STATUS_ARMGenerator returns a generator of RecordSet_STATUS_ARM instances for property testing.
func RecordSet_STATUS_ARMGenerator() gopter.Gen {
	if recordSet_STATUS_ARMGenerator != nil {
		return recordSet_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRecordSet_STATUS_ARM(generators)
	recordSet_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RecordSet_STATUS_ARM{}), generators)

	return recordSet_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRecordSet_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRecordSet_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["IpAddresses"] = gen.SliceOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		PrivateEndpointProvisioningState_STATUS_Deleting,
		PrivateEndpointProvisioningState_STATUS_Failed,
		PrivateEndpointProvisioningState_STATUS_Succeeded,
		PrivateEndpointProvisioningState_STATUS_Updating))
	gens["RecordSetName"] = gen.PtrOf(gen.AlphaString())
	gens["RecordType"] = gen.PtrOf(gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
}
