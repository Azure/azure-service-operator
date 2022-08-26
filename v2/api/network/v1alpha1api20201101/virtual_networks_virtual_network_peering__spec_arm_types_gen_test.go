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

<<<<<<<< HEAD:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peering__spec_arm_types_gen_test.go
func Test_VirtualNetworksVirtualNetworkPeering_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
========
func Test_VirtualNetworks_VirtualNetworkPeerings_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
>>>>>>>> main:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peerings_spec_arm_types_gen_test.go
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<<< HEAD:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peering__spec_arm_types_gen_test.go
		"Round trip of VirtualNetworksVirtualNetworkPeering_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_SpecARM, VirtualNetworksVirtualNetworkPeering_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_SpecARM runs a test to see if a specific instance of VirtualNetworksVirtualNetworkPeering_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_SpecARM(subject VirtualNetworksVirtualNetworkPeering_SpecARM) string {
========
		"Round trip of VirtualNetworks_VirtualNetworkPeerings_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeerings_SpecARM, VirtualNetworks_VirtualNetworkPeerings_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeerings_SpecARM runs a test to see if a specific instance of VirtualNetworks_VirtualNetworkPeerings_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeerings_SpecARM(subject VirtualNetworks_VirtualNetworkPeerings_SpecARM) string {
>>>>>>>> main:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peerings_spec_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
<<<<<<<< HEAD:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peering__spec_arm_types_gen_test.go
	var actual VirtualNetworksVirtualNetworkPeering_SpecARM
========
	var actual VirtualNetworks_VirtualNetworkPeerings_SpecARM
>>>>>>>> main:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peerings_spec_arm_types_gen_test.go
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

<<<<<<<< HEAD:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peering__spec_arm_types_gen_test.go
// Generator of VirtualNetworksVirtualNetworkPeering_SpecARM instances for property testing - lazily instantiated by
// VirtualNetworksVirtualNetworkPeering_SpecARMGenerator()
var virtualNetworksVirtualNetworkPeering_SpecARMGenerator gopter.Gen

// VirtualNetworksVirtualNetworkPeering_SpecARMGenerator returns a generator of VirtualNetworksVirtualNetworkPeering_SpecARM instances for property testing.
// We first initialize virtualNetworksVirtualNetworkPeering_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksVirtualNetworkPeering_SpecARMGenerator() gopter.Gen {
	if virtualNetworksVirtualNetworkPeering_SpecARMGenerator != nil {
		return virtualNetworksVirtualNetworkPeering_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_SpecARM(generators)
	virtualNetworksVirtualNetworkPeering_SpecARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_SpecARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_SpecARM(generators)
	virtualNetworksVirtualNetworkPeering_SpecARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering_SpecARM{}), generators)

	return virtualNetworksVirtualNetworkPeering_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Id"] = gen.PtrOf(gen.AlphaString())
========
// Generator of VirtualNetworks_VirtualNetworkPeerings_SpecARM instances for property testing - lazily instantiated by
// VirtualNetworks_VirtualNetworkPeerings_SpecARMGenerator()
var virtualNetworks_VirtualNetworkPeerings_SpecARMGenerator gopter.Gen

// VirtualNetworks_VirtualNetworkPeerings_SpecARMGenerator returns a generator of VirtualNetworks_VirtualNetworkPeerings_SpecARM instances for property testing.
// We first initialize virtualNetworks_VirtualNetworkPeerings_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworks_VirtualNetworkPeerings_SpecARMGenerator() gopter.Gen {
	if virtualNetworks_VirtualNetworkPeerings_SpecARMGenerator != nil {
		return virtualNetworks_VirtualNetworkPeerings_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_SpecARM(generators)
	virtualNetworks_VirtualNetworkPeerings_SpecARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_VirtualNetworkPeerings_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_SpecARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_SpecARM(generators)
	virtualNetworks_VirtualNetworkPeerings_SpecARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_VirtualNetworkPeerings_SpecARM{}), generators)

	return virtualNetworks_VirtualNetworkPeerings_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
>>>>>>>> main:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peerings_spec_arm_types_gen_test.go
	gens["Name"] = gen.AlphaString()
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

<<<<<<<< HEAD:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peering__spec_arm_types_gen_test.go
// AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_SpecARM(gens map[string]gopter.Gen) {
========
// AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_SpecARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peerings_spec_arm_types_gen_test.go
	gens["Properties"] = gen.PtrOf(VirtualNetworkPeeringPropertiesFormatARMGenerator())
}

func Test_VirtualNetworkPeeringPropertiesFormatARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeeringPropertiesFormatARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormatARM, VirtualNetworkPeeringPropertiesFormatARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormatARM runs a test to see if a specific instance of VirtualNetworkPeeringPropertiesFormatARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormatARM(subject VirtualNetworkPeeringPropertiesFormatARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeeringPropertiesFormatARM
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

// Generator of VirtualNetworkPeeringPropertiesFormatARM instances for property testing - lazily instantiated by
// VirtualNetworkPeeringPropertiesFormatARMGenerator()
var virtualNetworkPeeringPropertiesFormatARMGenerator gopter.Gen

// VirtualNetworkPeeringPropertiesFormatARMGenerator returns a generator of VirtualNetworkPeeringPropertiesFormatARM instances for property testing.
// We first initialize virtualNetworkPeeringPropertiesFormatARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeeringPropertiesFormatARMGenerator() gopter.Gen {
	if virtualNetworkPeeringPropertiesFormatARMGenerator != nil {
		return virtualNetworkPeeringPropertiesFormatARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM(generators)
	virtualNetworkPeeringPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormatARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM(generators)
	virtualNetworkPeeringPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormatARM{}), generators)

	return virtualNetworkPeeringPropertiesFormatARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
<<<<<<<< HEAD:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peering__spec_arm_types_gen_test.go
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
========
>>>>>>>> main:v2/api/network/v1alpha1api20201101/virtual_networks_virtual_network_peerings_spec_arm_types_gen_test.go
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_PeeringState_Connected, VirtualNetworkPeeringPropertiesFormat_PeeringState_Disconnected, VirtualNetworkPeeringPropertiesFormat_PeeringState_Initiated))
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpaceARMGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesARMGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResourceARMGenerator())
}
