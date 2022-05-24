// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

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

func Test_VirtualNetwork_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkStatusARM, VirtualNetworkStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkStatusARM runs a test to see if a specific instance of VirtualNetwork_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkStatusARM(subject VirtualNetwork_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_StatusARM
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

// Generator of VirtualNetwork_StatusARM instances for property testing - lazily instantiated by
// VirtualNetworkStatusARMGenerator()
var virtualNetworkStatusARMGenerator gopter.Gen

// VirtualNetworkStatusARMGenerator returns a generator of VirtualNetwork_StatusARM instances for property testing.
// We first initialize virtualNetworkStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkStatusARMGenerator() gopter.Gen {
	if virtualNetworkStatusARMGenerator != nil {
		return virtualNetworkStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkStatusARM(generators)
	virtualNetworkStatusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkStatusARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkStatusARM(generators)
	virtualNetworkStatusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_StatusARM{}), generators)

	return virtualNetworkStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkStatusARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkStatusARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationStatusARMGenerator())
	gens["Properties"] = gen.PtrOf(VirtualNetworkPropertiesFormatStatusARMGenerator())
}

func Test_VirtualNetworkPropertiesFormat_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPropertiesFormat_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPropertiesFormatStatusARM, VirtualNetworkPropertiesFormatStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPropertiesFormatStatusARM runs a test to see if a specific instance of VirtualNetworkPropertiesFormat_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPropertiesFormatStatusARM(subject VirtualNetworkPropertiesFormat_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPropertiesFormat_StatusARM
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

// Generator of VirtualNetworkPropertiesFormat_StatusARM instances for property testing - lazily instantiated by
// VirtualNetworkPropertiesFormatStatusARMGenerator()
var virtualNetworkPropertiesFormatStatusARMGenerator gopter.Gen

// VirtualNetworkPropertiesFormatStatusARMGenerator returns a generator of VirtualNetworkPropertiesFormat_StatusARM instances for property testing.
// We first initialize virtualNetworkPropertiesFormatStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPropertiesFormatStatusARMGenerator() gopter.Gen {
	if virtualNetworkPropertiesFormatStatusARMGenerator != nil {
		return virtualNetworkPropertiesFormatStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormatStatusARM(generators)
	virtualNetworkPropertiesFormatStatusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPropertiesFormat_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormatStatusARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormatStatusARM(generators)
	virtualNetworkPropertiesFormatStatusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPropertiesFormat_StatusARM{}), generators)

	return virtualNetworkPropertiesFormatStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormatStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormatStatusARM(gens map[string]gopter.Gen) {
	gens["EnableDdosProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableVmProtection"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormatStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormatStatusARM(gens map[string]gopter.Gen) {
	gens["AddressSpace"] = gen.PtrOf(AddressSpaceStatusARMGenerator())
	gens["BgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesStatusARMGenerator())
	gens["DdosProtectionPlan"] = gen.PtrOf(SubResourceStatusARMGenerator())
	gens["DhcpOptions"] = gen.PtrOf(DhcpOptionsStatusARMGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResourceStatusARMGenerator())
}

func Test_DhcpOptions_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DhcpOptions_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDhcpOptionsStatusARM, DhcpOptionsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDhcpOptionsStatusARM runs a test to see if a specific instance of DhcpOptions_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDhcpOptionsStatusARM(subject DhcpOptions_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DhcpOptions_StatusARM
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

// Generator of DhcpOptions_StatusARM instances for property testing - lazily instantiated by
// DhcpOptionsStatusARMGenerator()
var dhcpOptionsStatusARMGenerator gopter.Gen

// DhcpOptionsStatusARMGenerator returns a generator of DhcpOptions_StatusARM instances for property testing.
func DhcpOptionsStatusARMGenerator() gopter.Gen {
	if dhcpOptionsStatusARMGenerator != nil {
		return dhcpOptionsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDhcpOptionsStatusARM(generators)
	dhcpOptionsStatusARMGenerator = gen.Struct(reflect.TypeOf(DhcpOptions_StatusARM{}), generators)

	return dhcpOptionsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDhcpOptionsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDhcpOptionsStatusARM(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
}
