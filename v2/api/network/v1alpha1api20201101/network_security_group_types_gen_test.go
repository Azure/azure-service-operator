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

func Test_NetworkSecurityGroup_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroup to hub returns original",
		prop.ForAll(RunResourceConversionTestForNetworkSecurityGroup, NetworkSecurityGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNetworkSecurityGroup tests if a specific instance of NetworkSecurityGroup round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNetworkSecurityGroup(subject NetworkSecurityGroup) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20201101s.NetworkSecurityGroup
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NetworkSecurityGroup
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

func Test_NetworkSecurityGroup_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroup to NetworkSecurityGroup via AssignProperties_To_NetworkSecurityGroup & AssignProperties_From_NetworkSecurityGroup returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroup, NetworkSecurityGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroup tests if a specific instance of NetworkSecurityGroup can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroup(subject NetworkSecurityGroup) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.NetworkSecurityGroup
	err := copied.AssignProperties_To_NetworkSecurityGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroup
	err = actual.AssignProperties_From_NetworkSecurityGroup(&other)
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

func Test_NetworkSecurityGroup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroup, NetworkSecurityGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroup runs a test to see if a specific instance of NetworkSecurityGroup round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroup(subject NetworkSecurityGroup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup
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

// Generator of NetworkSecurityGroup instances for property testing - lazily instantiated by
// NetworkSecurityGroupGenerator()
var networkSecurityGroupGenerator gopter.Gen

// NetworkSecurityGroupGenerator returns a generator of NetworkSecurityGroup instances for property testing.
func NetworkSecurityGroupGenerator() gopter.Gen {
	if networkSecurityGroupGenerator != nil {
		return networkSecurityGroupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroup(generators)
	networkSecurityGroupGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup{}), generators)

	return networkSecurityGroupGenerator
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroup is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroup(gens map[string]gopter.Gen) {
	gens["Spec"] = NetworkSecurityGroup_SpecGenerator()
	gens["Status"] = NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()
}

func Test_NetworkSecurityGroup_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroup_Spec to NetworkSecurityGroup_Spec via AssignProperties_To_NetworkSecurityGroup_Spec & AssignProperties_From_NetworkSecurityGroup_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroup_Spec, NetworkSecurityGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroup_Spec tests if a specific instance of NetworkSecurityGroup_Spec can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroup_Spec(subject NetworkSecurityGroup_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.NetworkSecurityGroup_Spec
	err := copied.AssignProperties_To_NetworkSecurityGroup_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroup_Spec
	err = actual.AssignProperties_From_NetworkSecurityGroup_Spec(&other)
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

func Test_NetworkSecurityGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroup_Spec, NetworkSecurityGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroup_Spec runs a test to see if a specific instance of NetworkSecurityGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroup_Spec(subject NetworkSecurityGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_Spec
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

// Generator of NetworkSecurityGroup_Spec instances for property testing - lazily instantiated by
// NetworkSecurityGroup_SpecGenerator()
var networkSecurityGroup_SpecGenerator gopter.Gen

// NetworkSecurityGroup_SpecGenerator returns a generator of NetworkSecurityGroup_Spec instances for property testing.
func NetworkSecurityGroup_SpecGenerator() gopter.Gen {
	if networkSecurityGroup_SpecGenerator != nil {
		return networkSecurityGroup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_Spec(generators)
	networkSecurityGroup_SpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_Spec{}), generators)

	return networkSecurityGroup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroup_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded to NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded via AssignProperties_To_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded & AssignProperties_From_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded, NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded tests if a specific instance of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err := copied.AssignProperties_To_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err = actual.AssignProperties_From_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
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

func Test_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded, NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded runs a test to see if a specific instance of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded
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

// Generator of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing -
// lazily instantiated by NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()
var networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator gopter.Gen

// NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator returns a generator of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing.
// We first initialize networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator() gopter.Gen {
	if networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator != nil {
		return networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	return networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["DefaultSecurityRules"] = gen.SliceOf(SecurityRule_STATUSGenerator())
	gens["FlowLogs"] = gen.SliceOf(FlowLog_STATUSGenerator())
	gens["NetworkInterfaces"] = gen.SliceOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator())
	gens["SecurityRules"] = gen.SliceOf(SecurityRule_STATUSGenerator())
	gens["Subnets"] = gen.SliceOf(Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator())
}

func Test_FlowLog_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlowLog_STATUS to FlowLog_STATUS via AssignProperties_To_FlowLog_STATUS & AssignProperties_From_FlowLog_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlowLog_STATUS, FlowLog_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlowLog_STATUS tests if a specific instance of FlowLog_STATUS can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForFlowLog_STATUS(subject FlowLog_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.FlowLog_STATUS
	err := copied.AssignProperties_To_FlowLog_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlowLog_STATUS
	err = actual.AssignProperties_From_FlowLog_STATUS(&other)
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

func Test_FlowLog_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlowLog_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlowLog_STATUS, FlowLog_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlowLog_STATUS runs a test to see if a specific instance of FlowLog_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlowLog_STATUS(subject FlowLog_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlowLog_STATUS
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

// Generator of FlowLog_STATUS instances for property testing - lazily instantiated by FlowLog_STATUSGenerator()
var flowLog_STATUSGenerator gopter.Gen

// FlowLog_STATUSGenerator returns a generator of FlowLog_STATUS instances for property testing.
func FlowLog_STATUSGenerator() gopter.Gen {
	if flowLog_STATUSGenerator != nil {
		return flowLog_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlowLog_STATUS(generators)
	flowLog_STATUSGenerator = gen.Struct(reflect.TypeOf(FlowLog_STATUS{}), generators)

	return flowLog_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlowLog_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlowLog_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded to NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded via AssignProperties_To_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded & AssignProperties_From_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded, NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded tests if a specific instance of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err := copied.AssignProperties_To_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err = actual.AssignProperties_From_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
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

func Test_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded, NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded runs a test to see if a specific instance of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded
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

// Generator of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing - lazily
// instantiated by NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()
var networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator gopter.Gen

// NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator returns a generator of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing.
func NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator() gopter.Gen {
	if networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator != nil {
		return networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	return networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_SecurityRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SecurityRule_STATUS to SecurityRule_STATUS via AssignProperties_To_SecurityRule_STATUS & AssignProperties_From_SecurityRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSecurityRule_STATUS, SecurityRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSecurityRule_STATUS tests if a specific instance of SecurityRule_STATUS can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForSecurityRule_STATUS(subject SecurityRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.SecurityRule_STATUS
	err := copied.AssignProperties_To_SecurityRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SecurityRule_STATUS
	err = actual.AssignProperties_From_SecurityRule_STATUS(&other)
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

func Test_SecurityRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRule_STATUS, SecurityRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRule_STATUS runs a test to see if a specific instance of SecurityRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRule_STATUS(subject SecurityRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_STATUS
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

// Generator of SecurityRule_STATUS instances for property testing - lazily instantiated by
// SecurityRule_STATUSGenerator()
var securityRule_STATUSGenerator gopter.Gen

// SecurityRule_STATUSGenerator returns a generator of SecurityRule_STATUS instances for property testing.
func SecurityRule_STATUSGenerator() gopter.Gen {
	if securityRule_STATUSGenerator != nil {
		return securityRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRule_STATUS(generators)
	securityRule_STATUSGenerator = gen.Struct(reflect.TypeOf(SecurityRule_STATUS{}), generators)

	return securityRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded to Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded via AssignProperties_To_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded & AssignProperties_From_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded, Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded tests if a specific instance of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err := copied.AssignProperties_To_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err = actual.AssignProperties_From_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
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

func Test_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded, Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded runs a test to see if a specific instance of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded
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

// Generator of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing - lazily
// instantiated by Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()
var subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator gopter.Gen

// Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator returns a generator of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing.
func Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator() gopter.Gen {
	if subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator != nil {
		return subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	return subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
