// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	"encoding/json"
	v20231001s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
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

func Test_ClusterUpgradeSettings_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ClusterUpgradeSettings to ClusterUpgradeSettings via AssignProperties_To_ClusterUpgradeSettings & AssignProperties_From_ClusterUpgradeSettings returns original",
		prop.ForAll(RunPropertyAssignmentTestForClusterUpgradeSettings, ClusterUpgradeSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForClusterUpgradeSettings tests if a specific instance of ClusterUpgradeSettings can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForClusterUpgradeSettings(subject ClusterUpgradeSettings) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20231001s.ClusterUpgradeSettings
	err := copied.AssignProperties_To_ClusterUpgradeSettings(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ClusterUpgradeSettings
	err = actual.AssignProperties_From_ClusterUpgradeSettings(&other)
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

func Test_ClusterUpgradeSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ClusterUpgradeSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForClusterUpgradeSettings, ClusterUpgradeSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForClusterUpgradeSettings runs a test to see if a specific instance of ClusterUpgradeSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForClusterUpgradeSettings(subject ClusterUpgradeSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ClusterUpgradeSettings
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

// Generator of ClusterUpgradeSettings instances for property testing - lazily instantiated by
// ClusterUpgradeSettingsGenerator()
var clusterUpgradeSettingsGenerator gopter.Gen

// ClusterUpgradeSettingsGenerator returns a generator of ClusterUpgradeSettings instances for property testing.
func ClusterUpgradeSettingsGenerator() gopter.Gen {
	if clusterUpgradeSettingsGenerator != nil {
		return clusterUpgradeSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForClusterUpgradeSettings(generators)
	clusterUpgradeSettingsGenerator = gen.Struct(reflect.TypeOf(ClusterUpgradeSettings{}), generators)

	return clusterUpgradeSettingsGenerator
}

// AddRelatedPropertyGeneratorsForClusterUpgradeSettings is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForClusterUpgradeSettings(gens map[string]gopter.Gen) {
	gens["OverrideSettings"] = gen.PtrOf(UpgradeOverrideSettingsGenerator())
}
