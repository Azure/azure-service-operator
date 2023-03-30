// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101

import (
	"encoding/json"
	v1api20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101storage"
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

func Test_NetworkSecurityGroupsSecurityRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroupsSecurityRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForNetworkSecurityGroupsSecurityRule, NetworkSecurityGroupsSecurityRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNetworkSecurityGroupsSecurityRule tests if a specific instance of NetworkSecurityGroupsSecurityRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNetworkSecurityGroupsSecurityRule(subject NetworkSecurityGroupsSecurityRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1api20201101s.NetworkSecurityGroupsSecurityRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NetworkSecurityGroupsSecurityRule
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

func Test_NetworkSecurityGroupsSecurityRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroupsSecurityRule to NetworkSecurityGroupsSecurityRule via AssignProperties_To_NetworkSecurityGroupsSecurityRule & AssignProperties_From_NetworkSecurityGroupsSecurityRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule, NetworkSecurityGroupsSecurityRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule tests if a specific instance of NetworkSecurityGroupsSecurityRule can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule(subject NetworkSecurityGroupsSecurityRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20201101s.NetworkSecurityGroupsSecurityRule
	err := copied.AssignProperties_To_NetworkSecurityGroupsSecurityRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroupsSecurityRule
	err = actual.AssignProperties_From_NetworkSecurityGroupsSecurityRule(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroupsSecurityRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule, NetworkSecurityGroupsSecurityRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRule round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule(subject NetworkSecurityGroupsSecurityRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRule
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

// Generator of NetworkSecurityGroupsSecurityRule instances for property testing - lazily instantiated by
// NetworkSecurityGroupsSecurityRuleGenerator()
var networkSecurityGroupsSecurityRuleGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRuleGenerator returns a generator of NetworkSecurityGroupsSecurityRule instances for property testing.
func NetworkSecurityGroupsSecurityRuleGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRuleGenerator != nil {
		return networkSecurityGroupsSecurityRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule(generators)
	networkSecurityGroupsSecurityRuleGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule{}), generators)

	return networkSecurityGroupsSecurityRuleGenerator
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule(gens map[string]gopter.Gen) {
	gens["Spec"] = NetworkSecurityGroups_SecurityRule_SpecGenerator()
	gens["Status"] = NetworkSecurityGroups_SecurityRule_STATUSGenerator()
}

func Test_NetworkSecurityGroups_SecurityRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroups_SecurityRule_Spec to NetworkSecurityGroups_SecurityRule_Spec via AssignProperties_To_NetworkSecurityGroups_SecurityRule_Spec & AssignProperties_From_NetworkSecurityGroups_SecurityRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroups_SecurityRule_Spec, NetworkSecurityGroups_SecurityRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroups_SecurityRule_Spec tests if a specific instance of NetworkSecurityGroups_SecurityRule_Spec can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroups_SecurityRule_Spec(subject NetworkSecurityGroups_SecurityRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20201101s.NetworkSecurityGroups_SecurityRule_Spec
	err := copied.AssignProperties_To_NetworkSecurityGroups_SecurityRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroups_SecurityRule_Spec
	err = actual.AssignProperties_From_NetworkSecurityGroups_SecurityRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroups_SecurityRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroups_SecurityRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_Spec, NetworkSecurityGroups_SecurityRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_Spec runs a test to see if a specific instance of NetworkSecurityGroups_SecurityRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_Spec(subject NetworkSecurityGroups_SecurityRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroups_SecurityRule_Spec
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

// Generator of NetworkSecurityGroups_SecurityRule_Spec instances for property testing - lazily instantiated by
// NetworkSecurityGroups_SecurityRule_SpecGenerator()
var networkSecurityGroups_SecurityRule_SpecGenerator gopter.Gen

// NetworkSecurityGroups_SecurityRule_SpecGenerator returns a generator of NetworkSecurityGroups_SecurityRule_Spec instances for property testing.
// We first initialize networkSecurityGroups_SecurityRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroups_SecurityRule_SpecGenerator() gopter.Gen {
	if networkSecurityGroups_SecurityRule_SpecGenerator != nil {
		return networkSecurityGroups_SecurityRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec(generators)
	networkSecurityGroups_SecurityRule_SpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SecurityRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec(generators)
	networkSecurityGroups_SecurityRule_SpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SecurityRule_Spec{}), generators)

	return networkSecurityGroups_SecurityRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.OneConstOf(SecurityRuleAccess_Allow, SecurityRuleAccess_Deny))
	gens["AzureName"] = gen.AlphaString()
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.OneConstOf(SecurityRuleDirection_Inbound, SecurityRuleDirection_Outbound))
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(
		SecurityRulePropertiesFormat_Protocol_Ah,
		SecurityRulePropertiesFormat_Protocol_Esp,
		SecurityRulePropertiesFormat_Protocol_Icmp,
		SecurityRulePropertiesFormat_Protocol_Star,
		SecurityRulePropertiesFormat_Protocol_Tcp,
		SecurityRulePropertiesFormat_Protocol_Udp))
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator())
}

func Test_NetworkSecurityGroups_SecurityRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroups_SecurityRule_STATUS to NetworkSecurityGroups_SecurityRule_STATUS via AssignProperties_To_NetworkSecurityGroups_SecurityRule_STATUS & AssignProperties_From_NetworkSecurityGroups_SecurityRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroups_SecurityRule_STATUS, NetworkSecurityGroups_SecurityRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroups_SecurityRule_STATUS tests if a specific instance of NetworkSecurityGroups_SecurityRule_STATUS can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroups_SecurityRule_STATUS(subject NetworkSecurityGroups_SecurityRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20201101s.NetworkSecurityGroups_SecurityRule_STATUS
	err := copied.AssignProperties_To_NetworkSecurityGroups_SecurityRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroups_SecurityRule_STATUS
	err = actual.AssignProperties_From_NetworkSecurityGroups_SecurityRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroups_SecurityRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroups_SecurityRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_STATUS, NetworkSecurityGroups_SecurityRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_STATUS runs a test to see if a specific instance of NetworkSecurityGroups_SecurityRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_STATUS(subject NetworkSecurityGroups_SecurityRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroups_SecurityRule_STATUS
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

// Generator of NetworkSecurityGroups_SecurityRule_STATUS instances for property testing - lazily instantiated by
// NetworkSecurityGroups_SecurityRule_STATUSGenerator()
var networkSecurityGroups_SecurityRule_STATUSGenerator gopter.Gen

// NetworkSecurityGroups_SecurityRule_STATUSGenerator returns a generator of NetworkSecurityGroups_SecurityRule_STATUS instances for property testing.
// We first initialize networkSecurityGroups_SecurityRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroups_SecurityRule_STATUSGenerator() gopter.Gen {
	if networkSecurityGroups_SecurityRule_STATUSGenerator != nil {
		return networkSecurityGroups_SecurityRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS(generators)
	networkSecurityGroups_SecurityRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SecurityRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS(generators)
	networkSecurityGroups_SecurityRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SecurityRule_STATUS{}), generators)

	return networkSecurityGroups_SecurityRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.OneConstOf(SecurityRuleAccess_STATUS_Allow, SecurityRuleAccess_STATUS_Deny))
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.OneConstOf(SecurityRuleDirection_STATUS_Inbound, SecurityRuleDirection_STATUS_Outbound))
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(
		SecurityRulePropertiesFormat_Protocol_STATUS_Ah,
		SecurityRulePropertiesFormat_Protocol_STATUS_Esp,
		SecurityRulePropertiesFormat_Protocol_STATUS_Icmp,
		SecurityRulePropertiesFormat_Protocol_STATUS_Star,
		SecurityRulePropertiesFormat_Protocol_STATUS_Tcp,
		SecurityRulePropertiesFormat_Protocol_STATUS_Udp))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator())
}

func Test_ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded to ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded via AssignProperties_To_ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded & AssignProperties_From_ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded, ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded tests if a specific instance of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(subject ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20201101s.ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
	err := copied.AssignProperties_To_ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
	err = actual.AssignProperties_From_ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded, ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded runs a test to see if a specific instance of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(subject ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
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

// Generator of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded instances for
// property testing - lazily instantiated by
// ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator()
var applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator gopter.Gen

// ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator returns a generator of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded instances for property testing.
func ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator() gopter.Gen {
	if applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator != nil {
		return applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(generators)
	applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded{}), generators)

	return applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded to ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded via AssignProperties_To_ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded & AssignProperties_From_ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded, ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded tests if a specific instance of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(subject ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20201101s.ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
	err := copied.AssignProperties_To_ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
	err = actual.AssignProperties_From_ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded, ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded runs a test to see if a specific instance of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(subject ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
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

// Generator of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded instances for
// property testing - lazily instantiated by
// ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator()
var applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator gopter.Gen

// ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator returns a generator of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded instances for property testing.
func ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator() gopter.Gen {
	if applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator != nil {
		return applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded{}), generators)

	return applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator
}
