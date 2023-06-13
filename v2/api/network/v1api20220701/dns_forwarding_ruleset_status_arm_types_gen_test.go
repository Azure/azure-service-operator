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

func Test_DnsForwardingRuleset_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRuleset_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRuleset_STATUS_ARM, DnsForwardingRuleset_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRuleset_STATUS_ARM runs a test to see if a specific instance of DnsForwardingRuleset_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRuleset_STATUS_ARM(subject DnsForwardingRuleset_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRuleset_STATUS_ARM
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

// Generator of DnsForwardingRuleset_STATUS_ARM instances for property testing - lazily instantiated by
// DnsForwardingRuleset_STATUS_ARMGenerator()
var dnsForwardingRuleset_STATUS_ARMGenerator gopter.Gen

// DnsForwardingRuleset_STATUS_ARMGenerator returns a generator of DnsForwardingRuleset_STATUS_ARM instances for property testing.
// We first initialize dnsForwardingRuleset_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRuleset_STATUS_ARMGenerator() gopter.Gen {
	if dnsForwardingRuleset_STATUS_ARMGenerator != nil {
		return dnsForwardingRuleset_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS_ARM(generators)
	dnsForwardingRuleset_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleset_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRuleset_STATUS_ARM(generators)
	dnsForwardingRuleset_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleset_STATUS_ARM{}), generators)

	return dnsForwardingRuleset_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsForwardingRuleset_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRuleset_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DnsForwardingRulesetProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_DnsForwardingRulesetProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRulesetProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRulesetProperties_STATUS_ARM, DnsForwardingRulesetProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRulesetProperties_STATUS_ARM runs a test to see if a specific instance of DnsForwardingRulesetProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRulesetProperties_STATUS_ARM(subject DnsForwardingRulesetProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRulesetProperties_STATUS_ARM
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

// Generator of DnsForwardingRulesetProperties_STATUS_ARM instances for property testing - lazily instantiated by
// DnsForwardingRulesetProperties_STATUS_ARMGenerator()
var dnsForwardingRulesetProperties_STATUS_ARMGenerator gopter.Gen

// DnsForwardingRulesetProperties_STATUS_ARMGenerator returns a generator of DnsForwardingRulesetProperties_STATUS_ARM instances for property testing.
// We first initialize dnsForwardingRulesetProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRulesetProperties_STATUS_ARMGenerator() gopter.Gen {
	if dnsForwardingRulesetProperties_STATUS_ARMGenerator != nil {
		return dnsForwardingRulesetProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS_ARM(generators)
	dnsForwardingRulesetProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRulesetProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS_ARM(generators)
	dnsForwardingRulesetProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRulesetProperties_STATUS_ARM{}), generators)

	return dnsForwardingRulesetProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DnsresolverProvisioningState_STATUS_Canceled,
		DnsresolverProvisioningState_STATUS_Creating,
		DnsresolverProvisioningState_STATUS_Deleting,
		DnsresolverProvisioningState_STATUS_Failed,
		DnsresolverProvisioningState_STATUS_Succeeded,
		DnsresolverProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DnsResolverOutboundEndpoints"] = gen.SliceOf(DnsresolverSubResource_STATUS_ARMGenerator())
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

func Test_DnsresolverSubResource_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsresolverSubResource_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsresolverSubResource_STATUS_ARM, DnsresolverSubResource_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsresolverSubResource_STATUS_ARM runs a test to see if a specific instance of DnsresolverSubResource_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsresolverSubResource_STATUS_ARM(subject DnsresolverSubResource_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsresolverSubResource_STATUS_ARM
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

// Generator of DnsresolverSubResource_STATUS_ARM instances for property testing - lazily instantiated by
// DnsresolverSubResource_STATUS_ARMGenerator()
var dnsresolverSubResource_STATUS_ARMGenerator gopter.Gen

// DnsresolverSubResource_STATUS_ARMGenerator returns a generator of DnsresolverSubResource_STATUS_ARM instances for property testing.
func DnsresolverSubResource_STATUS_ARMGenerator() gopter.Gen {
	if dnsresolverSubResource_STATUS_ARMGenerator != nil {
		return dnsresolverSubResource_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsresolverSubResource_STATUS_ARM(generators)
	dnsresolverSubResource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsresolverSubResource_STATUS_ARM{}), generators)

	return dnsresolverSubResource_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsresolverSubResource_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsresolverSubResource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
