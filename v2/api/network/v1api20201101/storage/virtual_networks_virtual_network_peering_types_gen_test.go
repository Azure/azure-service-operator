// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_VirtualNetworksVirtualNetworkPeering_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworksVirtualNetworkPeering via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering, VirtualNetworksVirtualNetworkPeeringGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering runs a test to see if a specific instance of VirtualNetworksVirtualNetworkPeering round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering(subject VirtualNetworksVirtualNetworkPeering) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksVirtualNetworkPeering
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

// Generator of VirtualNetworksVirtualNetworkPeering instances for property testing - lazily instantiated by
// VirtualNetworksVirtualNetworkPeeringGenerator()
var virtualNetworksVirtualNetworkPeeringGenerator gopter.Gen

// VirtualNetworksVirtualNetworkPeeringGenerator returns a generator of VirtualNetworksVirtualNetworkPeering instances for property testing.
func VirtualNetworksVirtualNetworkPeeringGenerator() gopter.Gen {
	if virtualNetworksVirtualNetworkPeeringGenerator != nil {
		return virtualNetworksVirtualNetworkPeeringGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering(generators)
	virtualNetworksVirtualNetworkPeeringGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering{}), generators)

	return virtualNetworksVirtualNetworkPeeringGenerator
}

// AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering(gens map[string]gopter.Gen) {
	gens["Spec"] = VirtualNetworks_VirtualNetworkPeering_SpecGenerator()
	gens["Status"] = VirtualNetworks_VirtualNetworkPeering_STATUSGenerator()
}

func Test_VirtualNetworks_VirtualNetworkPeering_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworks_VirtualNetworkPeering_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeering_STATUS, VirtualNetworks_VirtualNetworkPeering_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeering_STATUS runs a test to see if a specific instance of VirtualNetworks_VirtualNetworkPeering_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeering_STATUS(subject VirtualNetworks_VirtualNetworkPeering_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworks_VirtualNetworkPeering_STATUS
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

// Generator of VirtualNetworks_VirtualNetworkPeering_STATUS instances for property testing - lazily instantiated by
// VirtualNetworks_VirtualNetworkPeering_STATUSGenerator()
var virtualNetworks_VirtualNetworkPeering_STATUSGenerator gopter.Gen

// VirtualNetworks_VirtualNetworkPeering_STATUSGenerator returns a generator of VirtualNetworks_VirtualNetworkPeering_STATUS instances for property testing.
// We first initialize virtualNetworks_VirtualNetworkPeering_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworks_VirtualNetworkPeering_STATUSGenerator() gopter.Gen {
	if virtualNetworks_VirtualNetworkPeering_STATUSGenerator != nil {
		return virtualNetworks_VirtualNetworkPeering_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_STATUS(generators)
	virtualNetworks_VirtualNetworkPeering_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_VirtualNetworkPeering_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_STATUS(generators)
	virtualNetworks_VirtualNetworkPeering_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_VirtualNetworkPeering_STATUS{}), generators)

	return virtualNetworks_VirtualNetworkPeering_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_STATUS(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PeeringState"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_STATUS(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpace_STATUSGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_STATUSGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_VirtualNetworks_VirtualNetworkPeering_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworks_VirtualNetworkPeering_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeering_Spec, VirtualNetworks_VirtualNetworkPeering_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeering_Spec runs a test to see if a specific instance of VirtualNetworks_VirtualNetworkPeering_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeering_Spec(subject VirtualNetworks_VirtualNetworkPeering_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworks_VirtualNetworkPeering_Spec
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

// Generator of VirtualNetworks_VirtualNetworkPeering_Spec instances for property testing - lazily instantiated by
// VirtualNetworks_VirtualNetworkPeering_SpecGenerator()
var virtualNetworks_VirtualNetworkPeering_SpecGenerator gopter.Gen

// VirtualNetworks_VirtualNetworkPeering_SpecGenerator returns a generator of VirtualNetworks_VirtualNetworkPeering_Spec instances for property testing.
// We first initialize virtualNetworks_VirtualNetworkPeering_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworks_VirtualNetworkPeering_SpecGenerator() gopter.Gen {
	if virtualNetworks_VirtualNetworkPeering_SpecGenerator != nil {
		return virtualNetworks_VirtualNetworkPeering_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_Spec(generators)
	virtualNetworks_VirtualNetworkPeering_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_VirtualNetworkPeering_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_Spec(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_Spec(generators)
	virtualNetworks_VirtualNetworkPeering_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_VirtualNetworkPeering_Spec{}), generators)

	return virtualNetworks_VirtualNetworkPeering_SpecGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_Spec(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PeeringState"] = gen.PtrOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_Spec(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpaceGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResourceGenerator())
}
