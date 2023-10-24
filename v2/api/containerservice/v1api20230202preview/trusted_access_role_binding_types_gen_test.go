// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230202preview

import (
	"encoding/json"
	v20230202ps "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202previewstorage"
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

func Test_TrustedAccessRoleBinding_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from TrustedAccessRoleBinding to hub returns original",
		prop.ForAll(RunResourceConversionTestForTrustedAccessRoleBinding, TrustedAccessRoleBindingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForTrustedAccessRoleBinding tests if a specific instance of TrustedAccessRoleBinding round trips to the hub storage version and back losslessly
func RunResourceConversionTestForTrustedAccessRoleBinding(subject TrustedAccessRoleBinding) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20230202ps.TrustedAccessRoleBinding
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual TrustedAccessRoleBinding
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

func Test_TrustedAccessRoleBinding_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from TrustedAccessRoleBinding to TrustedAccessRoleBinding via AssignProperties_To_TrustedAccessRoleBinding & AssignProperties_From_TrustedAccessRoleBinding returns original",
		prop.ForAll(RunPropertyAssignmentTestForTrustedAccessRoleBinding, TrustedAccessRoleBindingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTrustedAccessRoleBinding tests if a specific instance of TrustedAccessRoleBinding can be assigned to v1api20230202previewstorage and back losslessly
func RunPropertyAssignmentTestForTrustedAccessRoleBinding(subject TrustedAccessRoleBinding) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230202ps.TrustedAccessRoleBinding
	err := copied.AssignProperties_To_TrustedAccessRoleBinding(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual TrustedAccessRoleBinding
	err = actual.AssignProperties_From_TrustedAccessRoleBinding(&other)
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

func Test_TrustedAccessRoleBinding_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustedAccessRoleBinding via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustedAccessRoleBinding, TrustedAccessRoleBindingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustedAccessRoleBinding runs a test to see if a specific instance of TrustedAccessRoleBinding round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustedAccessRoleBinding(subject TrustedAccessRoleBinding) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustedAccessRoleBinding
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

// Generator of TrustedAccessRoleBinding instances for property testing - lazily instantiated by
// TrustedAccessRoleBindingGenerator()
var trustedAccessRoleBindingGenerator gopter.Gen

// TrustedAccessRoleBindingGenerator returns a generator of TrustedAccessRoleBinding instances for property testing.
func TrustedAccessRoleBindingGenerator() gopter.Gen {
	if trustedAccessRoleBindingGenerator != nil {
		return trustedAccessRoleBindingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding(generators)
	trustedAccessRoleBindingGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBinding{}), generators)

	return trustedAccessRoleBindingGenerator
}

// AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding(gens map[string]gopter.Gen) {
	gens["Spec"] = ManagedClusters_TrustedAccessRoleBinding_SpecGenerator()
	gens["Status"] = ManagedClusters_TrustedAccessRoleBinding_STATUSGenerator()
}

func Test_ManagedClusters_TrustedAccessRoleBinding_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ManagedClusters_TrustedAccessRoleBinding_Spec to ManagedClusters_TrustedAccessRoleBinding_Spec via AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_Spec & AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForManagedClusters_TrustedAccessRoleBinding_Spec, ManagedClusters_TrustedAccessRoleBinding_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForManagedClusters_TrustedAccessRoleBinding_Spec tests if a specific instance of ManagedClusters_TrustedAccessRoleBinding_Spec can be assigned to v1api20230202previewstorage and back losslessly
func RunPropertyAssignmentTestForManagedClusters_TrustedAccessRoleBinding_Spec(subject ManagedClusters_TrustedAccessRoleBinding_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230202ps.ManagedClusters_TrustedAccessRoleBinding_Spec
	err := copied.AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ManagedClusters_TrustedAccessRoleBinding_Spec
	err = actual.AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_Spec(&other)
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

func Test_ManagedClusters_TrustedAccessRoleBinding_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusters_TrustedAccessRoleBinding_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusters_TrustedAccessRoleBinding_Spec, ManagedClusters_TrustedAccessRoleBinding_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusters_TrustedAccessRoleBinding_Spec runs a test to see if a specific instance of ManagedClusters_TrustedAccessRoleBinding_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusters_TrustedAccessRoleBinding_Spec(subject ManagedClusters_TrustedAccessRoleBinding_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusters_TrustedAccessRoleBinding_Spec
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

// Generator of ManagedClusters_TrustedAccessRoleBinding_Spec instances for property testing - lazily instantiated by
// ManagedClusters_TrustedAccessRoleBinding_SpecGenerator()
var managedClusters_TrustedAccessRoleBinding_SpecGenerator gopter.Gen

// ManagedClusters_TrustedAccessRoleBinding_SpecGenerator returns a generator of ManagedClusters_TrustedAccessRoleBinding_Spec instances for property testing.
func ManagedClusters_TrustedAccessRoleBinding_SpecGenerator() gopter.Gen {
	if managedClusters_TrustedAccessRoleBinding_SpecGenerator != nil {
		return managedClusters_TrustedAccessRoleBinding_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_Spec(generators)
	managedClusters_TrustedAccessRoleBinding_SpecGenerator = gen.Struct(reflect.TypeOf(ManagedClusters_TrustedAccessRoleBinding_Spec{}), generators)

	return managedClusters_TrustedAccessRoleBinding_SpecGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Roles"] = gen.SliceOf(gen.AlphaString())
}

func Test_ManagedClusters_TrustedAccessRoleBinding_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ManagedClusters_TrustedAccessRoleBinding_STATUS to ManagedClusters_TrustedAccessRoleBinding_STATUS via AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_STATUS & AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForManagedClusters_TrustedAccessRoleBinding_STATUS, ManagedClusters_TrustedAccessRoleBinding_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForManagedClusters_TrustedAccessRoleBinding_STATUS tests if a specific instance of ManagedClusters_TrustedAccessRoleBinding_STATUS can be assigned to v1api20230202previewstorage and back losslessly
func RunPropertyAssignmentTestForManagedClusters_TrustedAccessRoleBinding_STATUS(subject ManagedClusters_TrustedAccessRoleBinding_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230202ps.ManagedClusters_TrustedAccessRoleBinding_STATUS
	err := copied.AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ManagedClusters_TrustedAccessRoleBinding_STATUS
	err = actual.AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_STATUS(&other)
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

func Test_ManagedClusters_TrustedAccessRoleBinding_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusters_TrustedAccessRoleBinding_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusters_TrustedAccessRoleBinding_STATUS, ManagedClusters_TrustedAccessRoleBinding_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusters_TrustedAccessRoleBinding_STATUS runs a test to see if a specific instance of ManagedClusters_TrustedAccessRoleBinding_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusters_TrustedAccessRoleBinding_STATUS(subject ManagedClusters_TrustedAccessRoleBinding_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusters_TrustedAccessRoleBinding_STATUS
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

// Generator of ManagedClusters_TrustedAccessRoleBinding_STATUS instances for property testing - lazily instantiated by
// ManagedClusters_TrustedAccessRoleBinding_STATUSGenerator()
var managedClusters_TrustedAccessRoleBinding_STATUSGenerator gopter.Gen

// ManagedClusters_TrustedAccessRoleBinding_STATUSGenerator returns a generator of ManagedClusters_TrustedAccessRoleBinding_STATUS instances for property testing.
// We first initialize managedClusters_TrustedAccessRoleBinding_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClusters_TrustedAccessRoleBinding_STATUSGenerator() gopter.Gen {
	if managedClusters_TrustedAccessRoleBinding_STATUSGenerator != nil {
		return managedClusters_TrustedAccessRoleBinding_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS(generators)
	managedClusters_TrustedAccessRoleBinding_STATUSGenerator = gen.Struct(reflect.TypeOf(ManagedClusters_TrustedAccessRoleBinding_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS(generators)
	AddRelatedPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS(generators)
	managedClusters_TrustedAccessRoleBinding_STATUSGenerator = gen.Struct(reflect.TypeOf(ManagedClusters_TrustedAccessRoleBinding_STATUS{}), generators)

	return managedClusters_TrustedAccessRoleBinding_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Canceled,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Deleting,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Failed,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Succeeded,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Updating))
	gens["Roles"] = gen.SliceOf(gen.AlphaString())
	gens["SourceResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
