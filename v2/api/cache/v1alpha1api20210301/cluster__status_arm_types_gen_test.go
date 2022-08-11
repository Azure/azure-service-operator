// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301

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

func Test_Cluster_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Cluster_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForClusterStatusARM, ClusterStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForClusterStatusARM runs a test to see if a specific instance of Cluster_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForClusterStatusARM(subject Cluster_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Cluster_StatusARM
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

// Generator of Cluster_StatusARM instances for property testing - lazily instantiated by ClusterStatusARMGenerator()
var clusterStatusARMGenerator gopter.Gen

// ClusterStatusARMGenerator returns a generator of Cluster_StatusARM instances for property testing.
// We first initialize clusterStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ClusterStatusARMGenerator() gopter.Gen {
	if clusterStatusARMGenerator != nil {
		return clusterStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterStatusARM(generators)
	clusterStatusARMGenerator = gen.Struct(reflect.TypeOf(Cluster_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterStatusARM(generators)
	AddRelatedPropertyGeneratorsForClusterStatusARM(generators)
	clusterStatusARMGenerator = gen.Struct(reflect.TypeOf(Cluster_StatusARM{}), generators)

	return clusterStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForClusterStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForClusterStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForClusterStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForClusterStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ClusterPropertiesStatusARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuStatusARMGenerator())
}

func Test_ClusterProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ClusterProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForClusterPropertiesStatusARM, ClusterPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForClusterPropertiesStatusARM runs a test to see if a specific instance of ClusterProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForClusterPropertiesStatusARM(subject ClusterProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ClusterProperties_StatusARM
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

// Generator of ClusterProperties_StatusARM instances for property testing - lazily instantiated by
// ClusterPropertiesStatusARMGenerator()
var clusterPropertiesStatusARMGenerator gopter.Gen

// ClusterPropertiesStatusARMGenerator returns a generator of ClusterProperties_StatusARM instances for property testing.
// We first initialize clusterPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ClusterPropertiesStatusARMGenerator() gopter.Gen {
	if clusterPropertiesStatusARMGenerator != nil {
		return clusterPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterPropertiesStatusARM(generators)
	clusterPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ClusterProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForClusterPropertiesStatusARM(generators)
	clusterPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ClusterProperties_StatusARM{}), generators)

	return clusterPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForClusterPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForClusterPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(ClusterPropertiesStatusMinimumTlsVersion_10, ClusterPropertiesStatusMinimumTlsVersion_11, ClusterPropertiesStatusMinimumTlsVersion_12))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_Status_Canceled,
		ProvisioningState_Status_Creating,
		ProvisioningState_Status_Deleting,
		ProvisioningState_Status_Failed,
		ProvisioningState_Status_Succeeded,
		ProvisioningState_Status_Updating))
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceState"] = gen.PtrOf(gen.OneConstOf(
		ResourceState_Status_CreateFailed,
		ResourceState_Status_Creating,
		ResourceState_Status_DeleteFailed,
		ResourceState_Status_Deleting,
		ResourceState_Status_DisableFailed,
		ResourceState_Status_Disabled,
		ResourceState_Status_Disabling,
		ResourceState_Status_EnableFailed,
		ResourceState_Status_Enabling,
		ResourceState_Status_Running,
		ResourceState_Status_UpdateFailed,
		ResourceState_Status_Updating))
}

// AddRelatedPropertyGeneratorsForClusterPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForClusterPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator())
}

func Test_Sku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuStatusARM, SkuStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuStatusARM runs a test to see if a specific instance of Sku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuStatusARM(subject Sku_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_StatusARM
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

// Generator of Sku_StatusARM instances for property testing - lazily instantiated by SkuStatusARMGenerator()
var skuStatusARMGenerator gopter.Gen

// SkuStatusARMGenerator returns a generator of Sku_StatusARM instances for property testing.
func SkuStatusARMGenerator() gopter.Gen {
	if skuStatusARMGenerator != nil {
		return skuStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuStatusARM(generators)
	skuStatusARMGenerator = gen.Struct(reflect.TypeOf(Sku_StatusARM{}), generators)

	return skuStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuStatusARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		SkuStatusName_EnterpriseE10,
		SkuStatusName_EnterpriseE100,
		SkuStatusName_EnterpriseE20,
		SkuStatusName_EnterpriseE50,
		SkuStatusName_EnterpriseFlashF1500,
		SkuStatusName_EnterpriseFlashF300,
		SkuStatusName_EnterpriseFlashF700))
}

func Test_PrivateEndpointConnection_Status_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_Status_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbeddedARM, PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbeddedARM runs a test to see if a specific instance of PrivateEndpointConnection_Status_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbeddedARM(subject PrivateEndpointConnection_Status_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_Status_SubResourceEmbeddedARM
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

// Generator of PrivateEndpointConnection_Status_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator()
var privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator gopter.Gen

// PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator returns a generator of PrivateEndpointConnection_Status_SubResourceEmbeddedARM instances for property testing.
func PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator() gopter.Gen {
	if privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator != nil {
		return privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbeddedARM(generators)
	privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_Status_SubResourceEmbeddedARM{}), generators)

	return privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
