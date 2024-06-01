// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/storage"
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

func Test_DnsResolver_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsResolver to hub returns original",
		prop.ForAll(RunResourceConversionTestForDnsResolver, DnsResolverGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDnsResolver tests if a specific instance of DnsResolver round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDnsResolver(subject DnsResolver) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.DnsResolver
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual DnsResolver
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

func Test_DnsResolver_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsResolver to DnsResolver via AssignProperties_To_DnsResolver & AssignProperties_From_DnsResolver returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsResolver, DnsResolverGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsResolver tests if a specific instance of DnsResolver can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsResolver(subject DnsResolver) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsResolver
	err := copied.AssignProperties_To_DnsResolver(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsResolver
	err = actual.AssignProperties_From_DnsResolver(&other)
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

func Test_DnsResolver_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolver via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolver, DnsResolverGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolver runs a test to see if a specific instance of DnsResolver round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolver(subject DnsResolver) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolver
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

// Generator of DnsResolver instances for property testing - lazily instantiated by DnsResolverGenerator()
var dnsResolverGenerator gopter.Gen

// DnsResolverGenerator returns a generator of DnsResolver instances for property testing.
func DnsResolverGenerator() gopter.Gen {
	if dnsResolverGenerator != nil {
		return dnsResolverGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsResolver(generators)
	dnsResolverGenerator = gen.Struct(reflect.TypeOf(DnsResolver{}), generators)

	return dnsResolverGenerator
}

// AddRelatedPropertyGeneratorsForDnsResolver is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolver(gens map[string]gopter.Gen) {
	gens["Spec"] = DnsResolver_SpecGenerator()
	gens["Status"] = DnsResolver_STATUSGenerator()
}

func Test_DnsResolver_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsResolver_Spec to DnsResolver_Spec via AssignProperties_To_DnsResolver_Spec & AssignProperties_From_DnsResolver_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsResolver_Spec, DnsResolver_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsResolver_Spec tests if a specific instance of DnsResolver_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsResolver_Spec(subject DnsResolver_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsResolver_Spec
	err := copied.AssignProperties_To_DnsResolver_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsResolver_Spec
	err = actual.AssignProperties_From_DnsResolver_Spec(&other)
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

func Test_DnsResolver_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolver_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolver_Spec, DnsResolver_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolver_Spec runs a test to see if a specific instance of DnsResolver_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolver_Spec(subject DnsResolver_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolver_Spec
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

// Generator of DnsResolver_Spec instances for property testing - lazily instantiated by DnsResolver_SpecGenerator()
var dnsResolver_SpecGenerator gopter.Gen

// DnsResolver_SpecGenerator returns a generator of DnsResolver_Spec instances for property testing.
// We first initialize dnsResolver_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsResolver_SpecGenerator() gopter.Gen {
	if dnsResolver_SpecGenerator != nil {
		return dnsResolver_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolver_Spec(generators)
	dnsResolver_SpecGenerator = gen.Struct(reflect.TypeOf(DnsResolver_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolver_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsResolver_Spec(generators)
	dnsResolver_SpecGenerator = gen.Struct(reflect.TypeOf(DnsResolver_Spec{}), generators)

	return dnsResolver_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsResolver_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsResolver_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsResolver_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolver_Spec(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(DnsresolverSubResourceGenerator())
}

func Test_DnsResolver_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsResolver_STATUS to DnsResolver_STATUS via AssignProperties_To_DnsResolver_STATUS & AssignProperties_From_DnsResolver_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsResolver_STATUS, DnsResolver_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsResolver_STATUS tests if a specific instance of DnsResolver_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsResolver_STATUS(subject DnsResolver_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsResolver_STATUS
	err := copied.AssignProperties_To_DnsResolver_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsResolver_STATUS
	err = actual.AssignProperties_From_DnsResolver_STATUS(&other)
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

func Test_DnsResolver_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolver_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolver_STATUS, DnsResolver_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolver_STATUS runs a test to see if a specific instance of DnsResolver_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolver_STATUS(subject DnsResolver_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolver_STATUS
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

// Generator of DnsResolver_STATUS instances for property testing - lazily instantiated by DnsResolver_STATUSGenerator()
var dnsResolver_STATUSGenerator gopter.Gen

// DnsResolver_STATUSGenerator returns a generator of DnsResolver_STATUS instances for property testing.
// We first initialize dnsResolver_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsResolver_STATUSGenerator() gopter.Gen {
	if dnsResolver_STATUSGenerator != nil {
		return dnsResolver_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolver_STATUS(generators)
	dnsResolver_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsResolver_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolver_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsResolver_STATUS(generators)
	dnsResolver_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsResolver_STATUS{}), generators)

	return dnsResolver_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsResolver_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsResolver_STATUS(gens map[string]gopter.Gen) {
	gens["DnsResolverState"] = gen.PtrOf(gen.OneConstOf(DnsResolverProperties_DnsResolverState_STATUS_Connected, DnsResolverProperties_DnsResolverState_STATUS_Disconnected))
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DnsresolverProvisioningState_STATUS_Canceled,
		DnsresolverProvisioningState_STATUS_Creating,
		DnsresolverProvisioningState_STATUS_Deleting,
		DnsresolverProvisioningState_STATUS_Failed,
		DnsresolverProvisioningState_STATUS_Succeeded,
		DnsresolverProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsResolver_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolver_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
	gens["VirtualNetwork"] = gen.PtrOf(DnsresolverSubResource_STATUSGenerator())
}
