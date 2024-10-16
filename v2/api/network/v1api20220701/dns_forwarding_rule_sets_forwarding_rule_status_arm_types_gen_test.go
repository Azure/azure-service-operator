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

func Test_DnsForwardingRuleSetsForwardingRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRuleSetsForwardingRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRuleSetsForwardingRule_STATUS_ARM, DnsForwardingRuleSetsForwardingRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRuleSetsForwardingRule_STATUS_ARM runs a test to see if a specific instance of DnsForwardingRuleSetsForwardingRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRuleSetsForwardingRule_STATUS_ARM(subject DnsForwardingRuleSetsForwardingRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRuleSetsForwardingRule_STATUS_ARM
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

// Generator of DnsForwardingRuleSetsForwardingRule_STATUS_ARM instances for property testing - lazily instantiated by
// DnsForwardingRuleSetsForwardingRule_STATUS_ARMGenerator()
var dnsForwardingRuleSetsForwardingRule_STATUS_ARMGenerator gopter.Gen

// DnsForwardingRuleSetsForwardingRule_STATUS_ARMGenerator returns a generator of DnsForwardingRuleSetsForwardingRule_STATUS_ARM instances for property testing.
// We first initialize dnsForwardingRuleSetsForwardingRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRuleSetsForwardingRule_STATUS_ARMGenerator() gopter.Gen {
	if dnsForwardingRuleSetsForwardingRule_STATUS_ARMGenerator != nil {
		return dnsForwardingRuleSetsForwardingRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsForwardingRule_STATUS_ARM(generators)
	dnsForwardingRuleSetsForwardingRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleSetsForwardingRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsForwardingRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsForwardingRule_STATUS_ARM(generators)
	dnsForwardingRuleSetsForwardingRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleSetsForwardingRule_STATUS_ARM{}), generators)

	return dnsForwardingRuleSetsForwardingRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsForwardingRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsForwardingRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsForwardingRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsForwardingRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ForwardingRuleProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_ForwardingRuleProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ForwardingRuleProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForForwardingRuleProperties_STATUS_ARM, ForwardingRuleProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForForwardingRuleProperties_STATUS_ARM runs a test to see if a specific instance of ForwardingRuleProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForForwardingRuleProperties_STATUS_ARM(subject ForwardingRuleProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ForwardingRuleProperties_STATUS_ARM
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

// Generator of ForwardingRuleProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ForwardingRuleProperties_STATUS_ARMGenerator()
var forwardingRuleProperties_STATUS_ARMGenerator gopter.Gen

// ForwardingRuleProperties_STATUS_ARMGenerator returns a generator of ForwardingRuleProperties_STATUS_ARM instances for property testing.
// We first initialize forwardingRuleProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ForwardingRuleProperties_STATUS_ARMGenerator() gopter.Gen {
	if forwardingRuleProperties_STATUS_ARMGenerator != nil {
		return forwardingRuleProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForForwardingRuleProperties_STATUS_ARM(generators)
	forwardingRuleProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ForwardingRuleProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForForwardingRuleProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForForwardingRuleProperties_STATUS_ARM(generators)
	forwardingRuleProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ForwardingRuleProperties_STATUS_ARM{}), generators)

	return forwardingRuleProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForForwardingRuleProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForForwardingRuleProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DomainName"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardingRuleState"] = gen.PtrOf(gen.OneConstOf(ForwardingRuleProperties_ForwardingRuleState_STATUS_ARM_Disabled, ForwardingRuleProperties_ForwardingRuleState_STATUS_ARM_Enabled))
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DnsresolverProvisioningState_STATUS_ARM_Canceled,
		DnsresolverProvisioningState_STATUS_ARM_Creating,
		DnsresolverProvisioningState_STATUS_ARM_Deleting,
		DnsresolverProvisioningState_STATUS_ARM_Failed,
		DnsresolverProvisioningState_STATUS_ARM_Succeeded,
		DnsresolverProvisioningState_STATUS_ARM_Updating))
}

// AddRelatedPropertyGeneratorsForForwardingRuleProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForForwardingRuleProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["TargetDnsServers"] = gen.SliceOf(TargetDnsServer_STATUS_ARMGenerator())
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
		SystemData_CreatedByType_STATUS_ARM_Application,
		SystemData_CreatedByType_STATUS_ARM_Key,
		SystemData_CreatedByType_STATUS_ARM_ManagedIdentity,
		SystemData_CreatedByType_STATUS_ARM_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_ARM_Application,
		SystemData_LastModifiedByType_STATUS_ARM_Key,
		SystemData_LastModifiedByType_STATUS_ARM_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_ARM_User))
}

func Test_TargetDnsServer_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TargetDnsServer_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTargetDnsServer_STATUS_ARM, TargetDnsServer_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTargetDnsServer_STATUS_ARM runs a test to see if a specific instance of TargetDnsServer_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTargetDnsServer_STATUS_ARM(subject TargetDnsServer_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TargetDnsServer_STATUS_ARM
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

// Generator of TargetDnsServer_STATUS_ARM instances for property testing - lazily instantiated by
// TargetDnsServer_STATUS_ARMGenerator()
var targetDnsServer_STATUS_ARMGenerator gopter.Gen

// TargetDnsServer_STATUS_ARMGenerator returns a generator of TargetDnsServer_STATUS_ARM instances for property testing.
func TargetDnsServer_STATUS_ARMGenerator() gopter.Gen {
	if targetDnsServer_STATUS_ARMGenerator != nil {
		return targetDnsServer_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTargetDnsServer_STATUS_ARM(generators)
	targetDnsServer_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(TargetDnsServer_STATUS_ARM{}), generators)

	return targetDnsServer_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTargetDnsServer_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTargetDnsServer_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["IpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Port"] = gen.PtrOf(gen.Int())
}
