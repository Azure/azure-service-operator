// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

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

func Test_VirtualNetworkPeering_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeering_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeering_STATUS_ARM, VirtualNetworkPeering_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeering_STATUS_ARM runs a test to see if a specific instance of VirtualNetworkPeering_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeering_STATUS_ARM(subject VirtualNetworkPeering_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeering_STATUS_ARM
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

// Generator of VirtualNetworkPeering_STATUS_ARM instances for property testing - lazily instantiated by
// VirtualNetworkPeering_STATUS_ARMGenerator()
var virtualNetworkPeering_STATUS_ARMGenerator gopter.Gen

// VirtualNetworkPeering_STATUS_ARMGenerator returns a generator of VirtualNetworkPeering_STATUS_ARM instances for property testing.
// We first initialize virtualNetworkPeering_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeering_STATUS_ARMGenerator() gopter.Gen {
	if virtualNetworkPeering_STATUS_ARMGenerator != nil {
		return virtualNetworkPeering_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUS_ARM(generators)
	virtualNetworkPeering_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeering_STATUS_ARM(generators)
	virtualNetworkPeering_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_STATUS_ARM{}), generators)

	return virtualNetworkPeering_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeering_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeering_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator())
}

func Test_VirtualNetworkPeeringPropertiesFormat_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeeringPropertiesFormat_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormat_STATUS_ARM, VirtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormat_STATUS_ARM runs a test to see if a specific instance of VirtualNetworkPeeringPropertiesFormat_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormat_STATUS_ARM(subject VirtualNetworkPeeringPropertiesFormat_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeeringPropertiesFormat_STATUS_ARM
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

// Generator of VirtualNetworkPeeringPropertiesFormat_STATUS_ARM instances for property testing - lazily instantiated by
// VirtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator()
var virtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator gopter.Gen

// VirtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator returns a generator of VirtualNetworkPeeringPropertiesFormat_STATUS_ARM instances for property testing.
// We first initialize virtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator() gopter.Gen {
	if virtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator != nil {
		return virtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUS_ARM(generators)
	virtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormat_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUS_ARM(generators)
	virtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormat_STATUS_ARM{}), generators)

	return virtualNetworkPeeringPropertiesFormat_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Connected, VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Disconnected, VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Initiated))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpace_STATUS_ARMGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_STATUS_ARMGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResource_STATUS_ARMGenerator())
}

func Test_VirtualNetworkBgpCommunities_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunities_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUS_ARM, VirtualNetworkBgpCommunities_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUS_ARM runs a test to see if a specific instance of VirtualNetworkBgpCommunities_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUS_ARM(subject VirtualNetworkBgpCommunities_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunities_STATUS_ARM
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

// Generator of VirtualNetworkBgpCommunities_STATUS_ARM instances for property testing - lazily instantiated by
// VirtualNetworkBgpCommunities_STATUS_ARMGenerator()
var virtualNetworkBgpCommunities_STATUS_ARMGenerator gopter.Gen

// VirtualNetworkBgpCommunities_STATUS_ARMGenerator returns a generator of VirtualNetworkBgpCommunities_STATUS_ARM instances for property testing.
func VirtualNetworkBgpCommunities_STATUS_ARMGenerator() gopter.Gen {
	if virtualNetworkBgpCommunities_STATUS_ARMGenerator != nil {
		return virtualNetworkBgpCommunities_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUS_ARM(generators)
	virtualNetworkBgpCommunities_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunities_STATUS_ARM{}), generators)

	return virtualNetworkBgpCommunities_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["RegionalCommunity"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkCommunity"] = gen.PtrOf(gen.AlphaString())
}
