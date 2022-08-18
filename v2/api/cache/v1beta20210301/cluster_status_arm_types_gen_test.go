// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210301

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

func Test_Cluster_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Cluster_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForClusterSTATUSARM, ClusterSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForClusterSTATUSARM runs a test to see if a specific instance of Cluster_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForClusterSTATUSARM(subject Cluster_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Cluster_STATUSARM
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

// Generator of Cluster_STATUSARM instances for property testing - lazily instantiated by ClusterSTATUSARMGenerator()
var clusterSTATUSARMGenerator gopter.Gen

// ClusterSTATUSARMGenerator returns a generator of Cluster_STATUSARM instances for property testing.
// We first initialize clusterSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ClusterSTATUSARMGenerator() gopter.Gen {
	if clusterSTATUSARMGenerator != nil {
		return clusterSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterSTATUSARM(generators)
	clusterSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Cluster_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForClusterSTATUSARM(generators)
	clusterSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Cluster_STATUSARM{}), generators)

	return clusterSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForClusterSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForClusterSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForClusterSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForClusterSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ClusterPropertiesSTATUSARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuSTATUSARMGenerator())
}

func Test_ClusterProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ClusterProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForClusterPropertiesSTATUSARM, ClusterPropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForClusterPropertiesSTATUSARM runs a test to see if a specific instance of ClusterProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForClusterPropertiesSTATUSARM(subject ClusterProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ClusterProperties_STATUSARM
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

// Generator of ClusterProperties_STATUSARM instances for property testing - lazily instantiated by
// ClusterPropertiesSTATUSARMGenerator()
var clusterPropertiesSTATUSARMGenerator gopter.Gen

// ClusterPropertiesSTATUSARMGenerator returns a generator of ClusterProperties_STATUSARM instances for property testing.
// We first initialize clusterPropertiesSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ClusterPropertiesSTATUSARMGenerator() gopter.Gen {
	if clusterPropertiesSTATUSARMGenerator != nil {
		return clusterPropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterPropertiesSTATUSARM(generators)
	clusterPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ClusterProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterPropertiesSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForClusterPropertiesSTATUSARM(generators)
	clusterPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ClusterProperties_STATUSARM{}), generators)

	return clusterPropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForClusterPropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForClusterPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(ClusterPropertiesSTATUSMinimumTlsVersion_10, ClusterPropertiesSTATUSMinimumTlsVersion_11, ClusterPropertiesSTATUSMinimumTlsVersion_12))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Canceled,
		ProvisioningState_STATUS_Creating,
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceState"] = gen.PtrOf(gen.OneConstOf(
		ResourceState_STATUS_CreateFailed,
		ResourceState_STATUS_Creating,
		ResourceState_STATUS_DeleteFailed,
		ResourceState_STATUS_Deleting,
		ResourceState_STATUS_DisableFailed,
		ResourceState_STATUS_Disabled,
		ResourceState_STATUS_Disabling,
		ResourceState_STATUS_EnableFailed,
		ResourceState_STATUS_Enabling,
		ResourceState_STATUS_Running,
		ResourceState_STATUS_UpdateFailed,
		ResourceState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForClusterPropertiesSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForClusterPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionSTATUSSubResourceEmbeddedARMGenerator())
}

func Test_Sku_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuSTATUSARM, SkuSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuSTATUSARM runs a test to see if a specific instance of Sku_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuSTATUSARM(subject Sku_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUSARM
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

// Generator of Sku_STATUSARM instances for property testing - lazily instantiated by SkuSTATUSARMGenerator()
var skuSTATUSARMGenerator gopter.Gen

// SkuSTATUSARMGenerator returns a generator of Sku_STATUSARM instances for property testing.
func SkuSTATUSARMGenerator() gopter.Gen {
	if skuSTATUSARMGenerator != nil {
		return skuSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuSTATUSARM(generators)
	skuSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUSARM{}), generators)

	return skuSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuSTATUSARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		SkuSTATUSName_EnterpriseE10,
		SkuSTATUSName_EnterpriseE100,
		SkuSTATUSName_EnterpriseE20,
		SkuSTATUSName_EnterpriseE50,
		SkuSTATUSName_EnterpriseFlashF1500,
		SkuSTATUSName_EnterpriseFlashF300,
		SkuSTATUSName_EnterpriseFlashF700))
}

func Test_PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionSTATUSSubResourceEmbeddedARM, PrivateEndpointConnectionSTATUSSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionSTATUSSubResourceEmbeddedARM runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionSTATUSSubResourceEmbeddedARM(subject PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM
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

// Generator of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by PrivateEndpointConnectionSTATUSSubResourceEmbeddedARMGenerator()
var privateEndpointConnectionSTATUSSubResourceEmbeddedARMGenerator gopter.Gen

// PrivateEndpointConnectionSTATUSSubResourceEmbeddedARMGenerator returns a generator of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM instances for property testing.
func PrivateEndpointConnectionSTATUSSubResourceEmbeddedARMGenerator() gopter.Gen {
	if privateEndpointConnectionSTATUSSubResourceEmbeddedARMGenerator != nil {
		return privateEndpointConnectionSTATUSSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionSTATUSSubResourceEmbeddedARM(generators)
	privateEndpointConnectionSTATUSSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM{}), generators)

	return privateEndpointConnectionSTATUSSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionSTATUSSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionSTATUSSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
