// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import (
	"encoding/json"
	alpha20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101storage"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
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

func Test_VirtualNetworksVirtualNetworkPeering_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworksVirtualNetworkPeering to hub returns original",
		prop.ForAll(RunResourceConversionTestForVirtualNetworksVirtualNetworkPeering, VirtualNetworksVirtualNetworkPeeringGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForVirtualNetworksVirtualNetworkPeering tests if a specific instance of VirtualNetworksVirtualNetworkPeering round trips to the hub storage version and back losslessly
func RunResourceConversionTestForVirtualNetworksVirtualNetworkPeering(subject VirtualNetworksVirtualNetworkPeering) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20201101s.VirtualNetworksVirtualNetworkPeering
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual VirtualNetworksVirtualNetworkPeering
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetworksVirtualNetworkPeering_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworksVirtualNetworkPeering to VirtualNetworksVirtualNetworkPeering via AssignProperties_To_VirtualNetworksVirtualNetworkPeering & AssignProperties_From_VirtualNetworksVirtualNetworkPeering returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering, VirtualNetworksVirtualNetworkPeeringGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering tests if a specific instance of VirtualNetworksVirtualNetworkPeering can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering(subject VirtualNetworksVirtualNetworkPeering) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.VirtualNetworksVirtualNetworkPeering
	err := copied.AssignProperties_To_VirtualNetworksVirtualNetworkPeering(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworksVirtualNetworkPeering
	err = actual.AssignProperties_From_VirtualNetworksVirtualNetworkPeering(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

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
	gens["Status"] = VirtualNetworkPeering_STATUSGenerator()
}

func Test_VirtualNetworkPeering_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworkPeering_STATUS to VirtualNetworkPeering_STATUS via AssignProperties_To_VirtualNetworkPeering_STATUS & AssignProperties_From_VirtualNetworkPeering_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworkPeering_STATUS, VirtualNetworkPeering_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworkPeering_STATUS tests if a specific instance of VirtualNetworkPeering_STATUS can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworkPeering_STATUS(subject VirtualNetworkPeering_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.VirtualNetworkPeering_STATUS
	err := copied.AssignProperties_To_VirtualNetworkPeering_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworkPeering_STATUS
	err = actual.AssignProperties_From_VirtualNetworkPeering_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetworkPeering_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeering_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeering_STATUS, VirtualNetworkPeering_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeering_STATUS runs a test to see if a specific instance of VirtualNetworkPeering_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeering_STATUS(subject VirtualNetworkPeering_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeering_STATUS
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

// Generator of VirtualNetworkPeering_STATUS instances for property testing - lazily instantiated by
// VirtualNetworkPeering_STATUSGenerator()
var virtualNetworkPeering_STATUSGenerator gopter.Gen

// VirtualNetworkPeering_STATUSGenerator returns a generator of VirtualNetworkPeering_STATUS instances for property testing.
// We first initialize virtualNetworkPeering_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeering_STATUSGenerator() gopter.Gen {
	if virtualNetworkPeering_STATUSGenerator != nil {
		return virtualNetworkPeering_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUS(generators)
	virtualNetworkPeering_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeering_STATUS(generators)
	virtualNetworkPeering_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_STATUS{}), generators)

	return virtualNetworkPeering_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUS(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Connected, VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Disconnected, VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Initiated))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeering_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeering_STATUS(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpace_STATUSGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_STATUSGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_VirtualNetworks_VirtualNetworkPeering_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworks_VirtualNetworkPeering_Spec to VirtualNetworks_VirtualNetworkPeering_Spec via AssignProperties_To_VirtualNetworks_VirtualNetworkPeering_Spec & AssignProperties_From_VirtualNetworks_VirtualNetworkPeering_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworks_VirtualNetworkPeering_Spec, VirtualNetworks_VirtualNetworkPeering_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworks_VirtualNetworkPeering_Spec tests if a specific instance of VirtualNetworks_VirtualNetworkPeering_Spec can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworks_VirtualNetworkPeering_Spec(subject VirtualNetworks_VirtualNetworkPeering_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.VirtualNetworks_VirtualNetworkPeering_Spec
	err := copied.AssignProperties_To_VirtualNetworks_VirtualNetworkPeering_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworks_VirtualNetworkPeering_Spec
	err = actual.AssignProperties_From_VirtualNetworks_VirtualNetworkPeering_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
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
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_PeeringState_Connected, VirtualNetworkPeeringPropertiesFormat_PeeringState_Disconnected, VirtualNetworkPeeringPropertiesFormat_PeeringState_Initiated))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeering_Spec(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpaceGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResourceGenerator())
}
