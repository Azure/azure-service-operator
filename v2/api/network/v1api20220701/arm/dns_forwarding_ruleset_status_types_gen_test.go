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

func Test_DnsForwardingRulesetProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRulesetProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRulesetProperties_STATUS, DnsForwardingRulesetProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRulesetProperties_STATUS runs a test to see if a specific instance of DnsForwardingRulesetProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRulesetProperties_STATUS(subject DnsForwardingRulesetProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRulesetProperties_STATUS
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

// Generator of DnsForwardingRulesetProperties_STATUS instances for property testing - lazily instantiated by
// DnsForwardingRulesetProperties_STATUSGenerator()
var dnsForwardingRulesetProperties_STATUSGenerator gopter.Gen

// DnsForwardingRulesetProperties_STATUSGenerator returns a generator of DnsForwardingRulesetProperties_STATUS instances for property testing.
// We first initialize dnsForwardingRulesetProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRulesetProperties_STATUSGenerator() gopter.Gen {
	if dnsForwardingRulesetProperties_STATUSGenerator != nil {
		return dnsForwardingRulesetProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS(generators)
	dnsForwardingRulesetProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRulesetProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS(generators)
	dnsForwardingRulesetProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRulesetProperties_STATUS{}), generators)

	return dnsForwardingRulesetProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DnsresolverProvisioningState_STATUS_Canceled,
		DnsresolverProvisioningState_STATUS_Creating,
		DnsresolverProvisioningState_STATUS_Deleting,
		DnsresolverProvisioningState_STATUS_Failed,
		DnsresolverProvisioningState_STATUS_Succeeded,
		DnsresolverProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRulesetProperties_STATUS(gens map[string]gopter.Gen) {
	gens["DnsResolverOutboundEndpoints"] = gen.SliceOf(SubResource_STATUSGenerator())
}

func Test_DnsForwardingRuleset_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRuleset_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRuleset_STATUS, DnsForwardingRuleset_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRuleset_STATUS runs a test to see if a specific instance of DnsForwardingRuleset_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRuleset_STATUS(subject DnsForwardingRuleset_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRuleset_STATUS
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

// Generator of DnsForwardingRuleset_STATUS instances for property testing - lazily instantiated by
// DnsForwardingRuleset_STATUSGenerator()
var dnsForwardingRuleset_STATUSGenerator gopter.Gen

// DnsForwardingRuleset_STATUSGenerator returns a generator of DnsForwardingRuleset_STATUS instances for property testing.
// We first initialize dnsForwardingRuleset_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRuleset_STATUSGenerator() gopter.Gen {
	if dnsForwardingRuleset_STATUSGenerator != nil {
		return dnsForwardingRuleset_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS(generators)
	dnsForwardingRuleset_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleset_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRuleset_STATUS(generators)
	dnsForwardingRuleset_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleset_STATUS{}), generators)

	return dnsForwardingRuleset_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsForwardingRuleset_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRuleset_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DnsForwardingRulesetProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
