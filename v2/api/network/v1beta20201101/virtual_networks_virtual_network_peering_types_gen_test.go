// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import (
	"encoding/json"
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

// RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering tests if a specific instance of VirtualNetworksVirtualNetworkPeering can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering(subject VirtualNetworksVirtualNetworkPeering) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.VirtualNetworksVirtualNetworkPeering
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
<<<<<<< HEAD
	gens["Spec"] = VirtualNetworksVirtualNetworkPeering_SpecGenerator()
	gens["Status"] = VirtualNetworksVirtualNetworkPeering_STATUSGenerator()
=======
	gens["Spec"] = VirtualNetworks_VirtualNetworkPeerings_SpecGenerator()
	gens["Status"] = VirtualNetworkPeering_STATUSGenerator()
>>>>>>> main
}

func Test_VirtualNetworksVirtualNetworkPeering_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip from VirtualNetworksVirtualNetworkPeering_Spec to VirtualNetworksVirtualNetworkPeering_Spec via AssignPropertiesToVirtualNetworksVirtualNetworkPeering_Spec & AssignPropertiesFromVirtualNetworksVirtualNetworkPeering_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_Spec, VirtualNetworksVirtualNetworkPeering_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_Spec tests if a specific instance of VirtualNetworksVirtualNetworkPeering_Spec can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_Spec(subject VirtualNetworksVirtualNetworkPeering_Spec) string {
=======
		"Round trip from VirtualNetworkPeering_STATUS to VirtualNetworkPeering_STATUS via AssignProperties_To_VirtualNetworkPeering_STATUS & AssignProperties_From_VirtualNetworkPeering_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworkPeering_STATUS, VirtualNetworkPeering_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworkPeering_STATUS tests if a specific instance of VirtualNetworkPeering_STATUS can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworkPeering_STATUS(subject VirtualNetworkPeering_STATUS) string {
>>>>>>> main
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
<<<<<<< HEAD
	var other v20201101s.VirtualNetworksVirtualNetworkPeering_Spec
	err := copied.AssignPropertiesToVirtualNetworksVirtualNetworkPeering_Spec(&other)
=======
	var other v20201101s.VirtualNetworkPeering_STATUS
	err := copied.AssignProperties_To_VirtualNetworkPeering_STATUS(&other)
>>>>>>> main
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
<<<<<<< HEAD
	var actual VirtualNetworksVirtualNetworkPeering_Spec
	err = actual.AssignPropertiesFromVirtualNetworksVirtualNetworkPeering_Spec(&other)
=======
	var actual VirtualNetworkPeering_STATUS
	err = actual.AssignProperties_From_VirtualNetworkPeering_STATUS(&other)
>>>>>>> main
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

func Test_VirtualNetworksVirtualNetworkPeering_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip of VirtualNetworksVirtualNetworkPeering_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_Spec, VirtualNetworksVirtualNetworkPeering_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_Spec runs a test to see if a specific instance of VirtualNetworksVirtualNetworkPeering_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_Spec(subject VirtualNetworksVirtualNetworkPeering_Spec) string {
=======
		"Round trip of VirtualNetworkPeering_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeering_STATUS, VirtualNetworkPeering_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeering_STATUS runs a test to see if a specific instance of VirtualNetworkPeering_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeering_STATUS(subject VirtualNetworkPeering_STATUS) string {
>>>>>>> main
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksVirtualNetworkPeering_Spec
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

<<<<<<< HEAD
// Generator of VirtualNetworksVirtualNetworkPeering_Spec instances for property testing - lazily instantiated by
// VirtualNetworksVirtualNetworkPeering_SpecGenerator()
var virtualNetworksVirtualNetworkPeering_SpecGenerator gopter.Gen

// VirtualNetworksVirtualNetworkPeering_SpecGenerator returns a generator of VirtualNetworksVirtualNetworkPeering_Spec instances for property testing.
// We first initialize virtualNetworksVirtualNetworkPeering_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksVirtualNetworkPeering_SpecGenerator() gopter.Gen {
	if virtualNetworksVirtualNetworkPeering_SpecGenerator != nil {
		return virtualNetworksVirtualNetworkPeering_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec(generators)
	virtualNetworksVirtualNetworkPeering_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec(generators)
	virtualNetworksVirtualNetworkPeering_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering_Spec{}), generators)

	return virtualNetworksVirtualNetworkPeering_SpecGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_PeeringState_Connected, VirtualNetworkPeeringPropertiesFormat_PeeringState_Disconnected, VirtualNetworkPeeringPropertiesFormat_PeeringState_Initiated))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpaceGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResourceGenerator())
}

func Test_VirtualNetworksVirtualNetworkPeering_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworksVirtualNetworkPeering_STATUS to VirtualNetworksVirtualNetworkPeering_STATUS via AssignPropertiesToVirtualNetworksVirtualNetworkPeering_STATUS & AssignPropertiesFromVirtualNetworksVirtualNetworkPeering_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_STATUS, VirtualNetworksVirtualNetworkPeering_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_STATUS tests if a specific instance of VirtualNetworksVirtualNetworkPeering_STATUS can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_STATUS(subject VirtualNetworksVirtualNetworkPeering_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.VirtualNetworksVirtualNetworkPeering_STATUS
	err := copied.AssignPropertiesToVirtualNetworksVirtualNetworkPeering_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworksVirtualNetworkPeering_STATUS
	err = actual.AssignPropertiesFromVirtualNetworksVirtualNetworkPeering_STATUS(&other)
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

func Test_VirtualNetworksVirtualNetworkPeering_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworksVirtualNetworkPeering_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_STATUS, VirtualNetworksVirtualNetworkPeering_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_STATUS runs a test to see if a specific instance of VirtualNetworksVirtualNetworkPeering_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_STATUS(subject VirtualNetworksVirtualNetworkPeering_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksVirtualNetworkPeering_STATUS
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

// Generator of VirtualNetworksVirtualNetworkPeering_STATUS instances for property testing - lazily instantiated by
// VirtualNetworksVirtualNetworkPeering_STATUSGenerator()
var virtualNetworksVirtualNetworkPeering_STATUSGenerator gopter.Gen

// VirtualNetworksVirtualNetworkPeering_STATUSGenerator returns a generator of VirtualNetworksVirtualNetworkPeering_STATUS instances for property testing.
// We first initialize virtualNetworksVirtualNetworkPeering_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksVirtualNetworkPeering_STATUSGenerator() gopter.Gen {
	if virtualNetworksVirtualNetworkPeering_STATUSGenerator != nil {
		return virtualNetworksVirtualNetworkPeering_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS(generators)
	virtualNetworksVirtualNetworkPeering_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS(generators)
	virtualNetworksVirtualNetworkPeering_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering_STATUS{}), generators)

	return virtualNetworksVirtualNetworkPeering_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS(gens map[string]gopter.Gen) {
=======
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
>>>>>>> main
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
<<<<<<< HEAD
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_PeeringState_Connected_STATUS, VirtualNetworkPeeringPropertiesFormat_PeeringState_Disconnected_STATUS, VirtualNetworkPeeringPropertiesFormat_PeeringState_Initiated_STATUS))
=======
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState_Connected, VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState_Disconnected, VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState_Initiated))
>>>>>>> main
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_Deleting_STATUS,
		ProvisioningState_Failed_STATUS,
		ProvisioningState_Succeeded_STATUS,
		ProvisioningState_Updating_STATUS))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

<<<<<<< HEAD
// AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpace_STATUSGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_STATUSGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResource_STATUSGenerator())
=======
// AddRelatedPropertyGeneratorsForVirtualNetworkPeering_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeering_STATUS(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpace_STATUSGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_STATUSGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_VirtualNetworks_VirtualNetworkPeerings_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworks_VirtualNetworkPeerings_Spec to VirtualNetworks_VirtualNetworkPeerings_Spec via AssignProperties_To_VirtualNetworks_VirtualNetworkPeerings_Spec & AssignProperties_From_VirtualNetworks_VirtualNetworkPeerings_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworks_VirtualNetworkPeerings_Spec, VirtualNetworks_VirtualNetworkPeerings_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworks_VirtualNetworkPeerings_Spec tests if a specific instance of VirtualNetworks_VirtualNetworkPeerings_Spec can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworks_VirtualNetworkPeerings_Spec(subject VirtualNetworks_VirtualNetworkPeerings_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.VirtualNetworks_VirtualNetworkPeerings_Spec
	err := copied.AssignProperties_To_VirtualNetworks_VirtualNetworkPeerings_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworks_VirtualNetworkPeerings_Spec
	err = actual.AssignProperties_From_VirtualNetworks_VirtualNetworkPeerings_Spec(&other)
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

func Test_VirtualNetworks_VirtualNetworkPeerings_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworks_VirtualNetworkPeerings_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeerings_Spec, VirtualNetworks_VirtualNetworkPeerings_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeerings_Spec runs a test to see if a specific instance of VirtualNetworks_VirtualNetworkPeerings_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworks_VirtualNetworkPeerings_Spec(subject VirtualNetworks_VirtualNetworkPeerings_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworks_VirtualNetworkPeerings_Spec
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

// Generator of VirtualNetworks_VirtualNetworkPeerings_Spec instances for property testing - lazily instantiated by
// VirtualNetworks_VirtualNetworkPeerings_SpecGenerator()
var virtualNetworks_VirtualNetworkPeerings_SpecGenerator gopter.Gen

// VirtualNetworks_VirtualNetworkPeerings_SpecGenerator returns a generator of VirtualNetworks_VirtualNetworkPeerings_Spec instances for property testing.
// We first initialize virtualNetworks_VirtualNetworkPeerings_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworks_VirtualNetworkPeerings_SpecGenerator() gopter.Gen {
	if virtualNetworks_VirtualNetworkPeerings_SpecGenerator != nil {
		return virtualNetworks_VirtualNetworkPeerings_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_Spec(generators)
	virtualNetworks_VirtualNetworkPeerings_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_VirtualNetworkPeerings_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_Spec(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_Spec(generators)
	virtualNetworks_VirtualNetworkPeerings_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_VirtualNetworkPeerings_Spec{}), generators)

	return virtualNetworks_VirtualNetworkPeerings_SpecGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_Spec(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_PeeringState_Connected, VirtualNetworkPeeringPropertiesFormat_PeeringState_Disconnected, VirtualNetworkPeeringPropertiesFormat_PeeringState_Initiated))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworks_VirtualNetworkPeerings_Spec(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpaceGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResourceGenerator())
>>>>>>> main
}
