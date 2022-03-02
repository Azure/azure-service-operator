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

func Test_VirtualNetworkPeering_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeering_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeeringStatusARM, VirtualNetworkPeeringStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeeringStatusARM runs a test to see if a specific instance of VirtualNetworkPeering_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeeringStatusARM(subject VirtualNetworkPeering_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeering_StatusARM
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

// Generator of VirtualNetworkPeering_StatusARM instances for property testing - lazily instantiated by
//VirtualNetworkPeeringStatusARMGenerator()
var virtualNetworkPeeringStatusARMGenerator gopter.Gen

// VirtualNetworkPeeringStatusARMGenerator returns a generator of VirtualNetworkPeering_StatusARM instances for property testing.
// We first initialize virtualNetworkPeeringStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeeringStatusARMGenerator() gopter.Gen {
	if virtualNetworkPeeringStatusARMGenerator != nil {
		return virtualNetworkPeeringStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringStatusARM(generators)
	virtualNetworkPeeringStatusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringStatusARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeeringStatusARM(generators)
	virtualNetworkPeeringStatusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_StatusARM{}), generators)

	return virtualNetworkPeeringStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeeringStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeeringStatusARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeeringStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeeringStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkPeeringPropertiesFormatStatusARMGenerator())
}

func Test_VirtualNetworkPeeringPropertiesFormat_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeeringPropertiesFormat_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormatStatusARM, VirtualNetworkPeeringPropertiesFormatStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormatStatusARM runs a test to see if a specific instance of VirtualNetworkPeeringPropertiesFormat_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormatStatusARM(subject VirtualNetworkPeeringPropertiesFormat_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeeringPropertiesFormat_StatusARM
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

// Generator of VirtualNetworkPeeringPropertiesFormat_StatusARM instances for property testing - lazily instantiated by
//VirtualNetworkPeeringPropertiesFormatStatusARMGenerator()
var virtualNetworkPeeringPropertiesFormatStatusARMGenerator gopter.Gen

// VirtualNetworkPeeringPropertiesFormatStatusARMGenerator returns a generator of VirtualNetworkPeeringPropertiesFormat_StatusARM instances for property testing.
// We first initialize virtualNetworkPeeringPropertiesFormatStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeeringPropertiesFormatStatusARMGenerator() gopter.Gen {
	if virtualNetworkPeeringPropertiesFormatStatusARMGenerator != nil {
		return virtualNetworkPeeringPropertiesFormatStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatStatusARM(generators)
	virtualNetworkPeeringPropertiesFormatStatusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormat_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatStatusARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatStatusARM(generators)
	virtualNetworkPeeringPropertiesFormatStatusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormat_StatusARM{}), generators)

	return virtualNetworkPeeringPropertiesFormatStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatStatusARM(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormatStatusPeeringStateConnected, VirtualNetworkPeeringPropertiesFormatStatusPeeringStateDisconnected, VirtualNetworkPeeringPropertiesFormatStatusPeeringStateInitiated))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_StatusDeleting,
		ProvisioningState_StatusFailed,
		ProvisioningState_StatusSucceeded,
		ProvisioningState_StatusUpdating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatStatusARM(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpaceStatusARMGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesStatusARMGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResourceStatusARMGenerator())
}

func Test_VirtualNetworkBgpCommunities_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunities_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunitiesStatusARM, VirtualNetworkBgpCommunitiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunitiesStatusARM runs a test to see if a specific instance of VirtualNetworkBgpCommunities_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunitiesStatusARM(subject VirtualNetworkBgpCommunities_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunities_StatusARM
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

// Generator of VirtualNetworkBgpCommunities_StatusARM instances for property testing - lazily instantiated by
//VirtualNetworkBgpCommunitiesStatusARMGenerator()
var virtualNetworkBgpCommunitiesStatusARMGenerator gopter.Gen

// VirtualNetworkBgpCommunitiesStatusARMGenerator returns a generator of VirtualNetworkBgpCommunities_StatusARM instances for property testing.
func VirtualNetworkBgpCommunitiesStatusARMGenerator() gopter.Gen {
	if virtualNetworkBgpCommunitiesStatusARMGenerator != nil {
		return virtualNetworkBgpCommunitiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesStatusARM(generators)
	virtualNetworkBgpCommunitiesStatusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunities_StatusARM{}), generators)

	return virtualNetworkBgpCommunitiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesStatusARM(gens map[string]gopter.Gen) {
	gens["RegionalCommunity"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkCommunity"] = gen.PtrOf(gen.AlphaString())
}
