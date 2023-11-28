// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	"encoding/json"
	v20230701s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230701/storage"
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

func Test_ClusterUpgradeSettings_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ClusterUpgradeSettings_STATUS to ClusterUpgradeSettings_STATUS via AssignProperties_To_ClusterUpgradeSettings_STATUS & AssignProperties_From_ClusterUpgradeSettings_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForClusterUpgradeSettings_STATUS, ClusterUpgradeSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForClusterUpgradeSettings_STATUS tests if a specific instance of ClusterUpgradeSettings_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForClusterUpgradeSettings_STATUS(subject ClusterUpgradeSettings_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230701s.ClusterUpgradeSettings_STATUS
	err := copied.AssignProperties_To_ClusterUpgradeSettings_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ClusterUpgradeSettings_STATUS
	err = actual.AssignProperties_From_ClusterUpgradeSettings_STATUS(&other)
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

func Test_ClusterUpgradeSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ClusterUpgradeSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForClusterUpgradeSettings_STATUS, ClusterUpgradeSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForClusterUpgradeSettings_STATUS runs a test to see if a specific instance of ClusterUpgradeSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForClusterUpgradeSettings_STATUS(subject ClusterUpgradeSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ClusterUpgradeSettings_STATUS
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

// Generator of ClusterUpgradeSettings_STATUS instances for property testing - lazily instantiated by
// ClusterUpgradeSettings_STATUSGenerator()
var clusterUpgradeSettings_STATUSGenerator gopter.Gen

// ClusterUpgradeSettings_STATUSGenerator returns a generator of ClusterUpgradeSettings_STATUS instances for property testing.
func ClusterUpgradeSettings_STATUSGenerator() gopter.Gen {
	if clusterUpgradeSettings_STATUSGenerator != nil {
		return clusterUpgradeSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForClusterUpgradeSettings_STATUS(generators)
	clusterUpgradeSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(ClusterUpgradeSettings_STATUS{}), generators)

	return clusterUpgradeSettings_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForClusterUpgradeSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForClusterUpgradeSettings_STATUS(gens map[string]gopter.Gen) {
	gens["OverrideSettings"] = gen.PtrOf(UpgradeOverrideSettings_STATUSGenerator())
}
