// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(TrustedAccessRoleBindingProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_TrustedAccessRoleBindingProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustedAccessRoleBindingProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustedAccessRoleBindingProperties_STATUS, TrustedAccessRoleBindingProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustedAccessRoleBindingProperties_STATUS runs a test to see if a specific instance of TrustedAccessRoleBindingProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustedAccessRoleBindingProperties_STATUS(subject TrustedAccessRoleBindingProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustedAccessRoleBindingProperties_STATUS
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

// Generator of TrustedAccessRoleBindingProperties_STATUS instances for property testing - lazily instantiated by
// TrustedAccessRoleBindingProperties_STATUSGenerator()
var trustedAccessRoleBindingProperties_STATUSGenerator gopter.Gen

// TrustedAccessRoleBindingProperties_STATUSGenerator returns a generator of TrustedAccessRoleBindingProperties_STATUS instances for property testing.
func TrustedAccessRoleBindingProperties_STATUSGenerator() gopter.Gen {
	if trustedAccessRoleBindingProperties_STATUSGenerator != nil {
		return trustedAccessRoleBindingProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrustedAccessRoleBindingProperties_STATUS(generators)
	trustedAccessRoleBindingProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBindingProperties_STATUS{}), generators)

	return trustedAccessRoleBindingProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTrustedAccessRoleBindingProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrustedAccessRoleBindingProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Canceled,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Deleting,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Failed,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Succeeded,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Updating))
	gens["Roles"] = gen.SliceOf(gen.AlphaString())
	gens["SourceResourceId"] = gen.PtrOf(gen.AlphaString())
}
