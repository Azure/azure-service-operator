// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

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

func Test_NetworkSecurityGroup_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
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
		"Round trip from NetworkSecurityGroup to NetworkSecurityGroup via AssignPropertiesToNetworkSecurityGroup & AssignPropertiesFromNetworkSecurityGroup returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroup, NetworkSecurityGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroup tests if a specific instance of NetworkSecurityGroup can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroup(subject NetworkSecurityGroup) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.NetworkSecurityGroup
	err := copied.AssignPropertiesToNetworkSecurityGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroup
	err = actual.AssignPropertiesFromNetworkSecurityGroup(&other)
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
	gens["Spec"] = NetworkSecurityGroupsSpecGenerator()
	gens["Status"] = NetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()
}

func Test_NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded to NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded via AssignPropertiesToNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded & AssignPropertiesFromNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded, NetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded tests if a specific instance of NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(subject NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
	err := copied.AssignPropertiesToNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
	err = actual.AssignPropertiesFromNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(&other)
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

func Test_NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded, NetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded runs a test to see if a specific instance of NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(subject NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
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

// Generator of NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded instances for property testing -
// lazily instantiated by NetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()
var networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator gopter.Gen

// NetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator returns a generator of NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded instances for property testing.
// We first initialize networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator() gopter.Gen {
	if networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator != nil {
		return networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(generators)
	networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(generators)
	networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	return networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["DefaultSecurityRules"] = gen.SliceOf(SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator())
	gens["FlowLogs"] = gen.SliceOf(FlowLogStatusSubResourceEmbeddedGenerator())
	gens["NetworkInterfaces"] = gen.SliceOf(NetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator())
	gens["SecurityRules"] = gen.SliceOf(SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator())
	gens["Subnets"] = gen.SliceOf(SubnetStatusNetworkSecurityGroupSubResourceEmbeddedGenerator())
}

func Test_NetworkSecurityGroups_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroups_Spec to NetworkSecurityGroups_Spec via AssignPropertiesToNetworkSecurityGroupsSpec & AssignPropertiesFromNetworkSecurityGroupsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroupsSpec, NetworkSecurityGroupsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroupsSpec tests if a specific instance of NetworkSecurityGroups_Spec can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroupsSpec(subject NetworkSecurityGroups_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.NetworkSecurityGroups_Spec
	err := copied.AssignPropertiesToNetworkSecurityGroupsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroups_Spec
	err = actual.AssignPropertiesFromNetworkSecurityGroupsSpec(&other)
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

func Test_NetworkSecurityGroups_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroups_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSpec, NetworkSecurityGroupsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSpec runs a test to see if a specific instance of NetworkSecurityGroups_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSpec(subject NetworkSecurityGroups_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroups_Spec
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

// Generator of NetworkSecurityGroups_Spec instances for property testing - lazily instantiated by
// NetworkSecurityGroupsSpecGenerator()
var networkSecurityGroupsSpecGenerator gopter.Gen

// NetworkSecurityGroupsSpecGenerator returns a generator of NetworkSecurityGroups_Spec instances for property testing.
func NetworkSecurityGroupsSpecGenerator() gopter.Gen {
	if networkSecurityGroupsSpecGenerator != nil {
		return networkSecurityGroupsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSpec(generators)
	networkSecurityGroupsSpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_Spec{}), generators)

	return networkSecurityGroupsSpecGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_FlowLog_Status_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlowLog_Status_SubResourceEmbedded to FlowLog_Status_SubResourceEmbedded via AssignPropertiesToFlowLogStatusSubResourceEmbedded & AssignPropertiesFromFlowLogStatusSubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlowLogStatusSubResourceEmbedded, FlowLogStatusSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlowLogStatusSubResourceEmbedded tests if a specific instance of FlowLog_Status_SubResourceEmbedded can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForFlowLogStatusSubResourceEmbedded(subject FlowLog_Status_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.FlowLog_Status_SubResourceEmbedded
	err := copied.AssignPropertiesToFlowLogStatusSubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlowLog_Status_SubResourceEmbedded
	err = actual.AssignPropertiesFromFlowLogStatusSubResourceEmbedded(&other)
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

func Test_FlowLog_Status_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlowLog_Status_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlowLogStatusSubResourceEmbedded, FlowLogStatusSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlowLogStatusSubResourceEmbedded runs a test to see if a specific instance of FlowLog_Status_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForFlowLogStatusSubResourceEmbedded(subject FlowLog_Status_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlowLog_Status_SubResourceEmbedded
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

// Generator of FlowLog_Status_SubResourceEmbedded instances for property testing - lazily instantiated by
// FlowLogStatusSubResourceEmbeddedGenerator()
var flowLogStatusSubResourceEmbeddedGenerator gopter.Gen

// FlowLogStatusSubResourceEmbeddedGenerator returns a generator of FlowLog_Status_SubResourceEmbedded instances for property testing.
func FlowLogStatusSubResourceEmbeddedGenerator() gopter.Gen {
	if flowLogStatusSubResourceEmbeddedGenerator != nil {
		return flowLogStatusSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlowLogStatusSubResourceEmbedded(generators)
	flowLogStatusSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(FlowLog_Status_SubResourceEmbedded{}), generators)

	return flowLogStatusSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForFlowLogStatusSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlowLogStatusSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded to NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded via AssignPropertiesToNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded & AssignPropertiesFromNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded, NetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded tests if a specific instance of NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(subject NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded
	err := copied.AssignPropertiesToNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded
	err = actual.AssignPropertiesFromNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(&other)
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

func Test_NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded, NetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded runs a test to see if a specific instance of NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(subject NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded
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

// Generator of NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded instances for property testing - lazily
// instantiated by NetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()
var networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator gopter.Gen

// NetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator returns a generator of NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded instances for property testing.
// We first initialize networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator() gopter.Gen {
	if networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator != nil {
		return networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(generators)
	networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(generators)
	AddRelatedPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(generators)
	networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	return networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationStatusGenerator())
}

func Test_SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded to SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded via AssignPropertiesToSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded & AssignPropertiesFromSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded, SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded tests if a specific instance of SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(subject SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
	err := copied.AssignPropertiesToSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
	err = actual.AssignPropertiesFromSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(&other)
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

func Test_SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded, SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded runs a test to see if a specific instance of SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(subject SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
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

// Generator of SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded instances for property testing - lazily
// instantiated by SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()
var securityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator gopter.Gen

// SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator returns a generator of SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded instances for property testing.
func SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator() gopter.Gen {
	if securityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator != nil {
		return securityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(generators)
	securityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	return securityRuleStatusNetworkSecurityGroupSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded to Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded via AssignPropertiesToSubnetStatusNetworkSecurityGroupSubResourceEmbedded & AssignPropertiesFromSubnetStatusNetworkSecurityGroupSubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubnetStatusNetworkSecurityGroupSubResourceEmbedded, SubnetStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubnetStatusNetworkSecurityGroupSubResourceEmbedded tests if a specific instance of Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForSubnetStatusNetworkSecurityGroupSubResourceEmbedded(subject Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded
	err := copied.AssignPropertiesToSubnetStatusNetworkSecurityGroupSubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded
	err = actual.AssignPropertiesFromSubnetStatusNetworkSecurityGroupSubResourceEmbedded(&other)
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

func Test_Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnetStatusNetworkSecurityGroupSubResourceEmbedded, SubnetStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnetStatusNetworkSecurityGroupSubResourceEmbedded runs a test to see if a specific instance of Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnetStatusNetworkSecurityGroupSubResourceEmbedded(subject Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded
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

// Generator of Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded instances for property testing - lazily
// instantiated by SubnetStatusNetworkSecurityGroupSubResourceEmbeddedGenerator()
var subnetStatusNetworkSecurityGroupSubResourceEmbeddedGenerator gopter.Gen

// SubnetStatusNetworkSecurityGroupSubResourceEmbeddedGenerator returns a generator of Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded instances for property testing.
func SubnetStatusNetworkSecurityGroupSubResourceEmbeddedGenerator() gopter.Gen {
	if subnetStatusNetworkSecurityGroupSubResourceEmbeddedGenerator != nil {
		return subnetStatusNetworkSecurityGroupSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnetStatusNetworkSecurityGroupSubResourceEmbedded(generators)
	subnetStatusNetworkSecurityGroupSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	return subnetStatusNetworkSecurityGroupSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForSubnetStatusNetworkSecurityGroupSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnetStatusNetworkSecurityGroupSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
