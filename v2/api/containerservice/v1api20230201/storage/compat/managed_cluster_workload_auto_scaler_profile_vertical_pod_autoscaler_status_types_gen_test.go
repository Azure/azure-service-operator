// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
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

func Test_ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS to ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS via AssignProperties_To_ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS & AssignProperties_From_ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS, ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS tests if a specific instance of ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS(subject ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS
	err := copied.AssignProperties_To_ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS
	err = actual.AssignProperties_From_ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS(&other)
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

func Test_ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS, ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS runs a test to see if a specific instance of ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS(subject ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS
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

// Generator of ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS instances for property testing -
// lazily instantiated by ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUSGenerator()
var managedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUSGenerator gopter.Gen

// ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUSGenerator returns a generator of ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS instances for property testing.
func ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUSGenerator() gopter.Gen {
	if managedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUSGenerator != nil {
		return managedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS(generators)
	managedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUSGenerator = gen.Struct(reflect.TypeOf(ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS{}), generators)

	return managedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS(gens map[string]gopter.Gen) {
	gens["ControlledValues"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["UpdateMode"] = gen.PtrOf(gen.AlphaString())
}
