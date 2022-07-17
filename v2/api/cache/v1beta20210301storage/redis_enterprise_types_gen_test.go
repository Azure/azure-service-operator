// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210301storage

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

func Test_RedisEnterprise_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterprise via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterprise, RedisEnterpriseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterprise runs a test to see if a specific instance of RedisEnterprise round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterprise(subject RedisEnterprise) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterprise
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

// Generator of RedisEnterprise instances for property testing - lazily instantiated by RedisEnterpriseGenerator()
var redisEnterpriseGenerator gopter.Gen

// RedisEnterpriseGenerator returns a generator of RedisEnterprise instances for property testing.
func RedisEnterpriseGenerator() gopter.Gen {
	if redisEnterpriseGenerator != nil {
		return redisEnterpriseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisEnterprise(generators)
	redisEnterpriseGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise{}), generators)

	return redisEnterpriseGenerator
}

// AddRelatedPropertyGeneratorsForRedisEnterprise is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterprise(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisEnterpriseSpecGenerator()
	gens["Status"] = ClusterStatusGenerator()
}

func Test_Cluster_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Cluster_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForClusterStatus, ClusterStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForClusterStatus runs a test to see if a specific instance of Cluster_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForClusterStatus(subject Cluster_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Cluster_Status
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

// Generator of Cluster_Status instances for property testing - lazily instantiated by ClusterStatusGenerator()
var clusterStatusGenerator gopter.Gen

// ClusterStatusGenerator returns a generator of Cluster_Status instances for property testing.
// We first initialize clusterStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ClusterStatusGenerator() gopter.Gen {
	if clusterStatusGenerator != nil {
		return clusterStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterStatus(generators)
	clusterStatusGenerator = gen.Struct(reflect.TypeOf(Cluster_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterStatus(generators)
	AddRelatedPropertyGeneratorsForClusterStatus(generators)
	clusterStatusGenerator = gen.Struct(reflect.TypeOf(Cluster_Status{}), generators)

	return clusterStatusGenerator
}

// AddIndependentPropertyGeneratorsForClusterStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForClusterStatus(gens map[string]gopter.Gen) {
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceState"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForClusterStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForClusterStatus(gens map[string]gopter.Gen) {
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionStatusSubResourceEmbeddedGenerator())
	gens["Sku"] = gen.PtrOf(SkuStatusGenerator())
}

func Test_RedisEnterprise_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterprise_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterpriseSpec, RedisEnterpriseSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterpriseSpec runs a test to see if a specific instance of RedisEnterprise_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterpriseSpec(subject RedisEnterprise_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterprise_Spec
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

// Generator of RedisEnterprise_Spec instances for property testing - lazily instantiated by
// RedisEnterpriseSpecGenerator()
var redisEnterpriseSpecGenerator gopter.Gen

// RedisEnterpriseSpecGenerator returns a generator of RedisEnterprise_Spec instances for property testing.
// We first initialize redisEnterpriseSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterpriseSpecGenerator() gopter.Gen {
	if redisEnterpriseSpecGenerator != nil {
		return redisEnterpriseSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseSpec(generators)
	redisEnterpriseSpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseSpec(generators)
	AddRelatedPropertyGeneratorsForRedisEnterpriseSpec(generators)
	redisEnterpriseSpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_Spec{}), generators)

	return redisEnterpriseSpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterpriseSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterpriseSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisEnterpriseSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterpriseSpec(gens map[string]gopter.Gen) {
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_PrivateEndpointConnection_Status_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_Status_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbedded, PrivateEndpointConnectionStatusSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbedded runs a test to see if a specific instance of PrivateEndpointConnection_Status_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbedded(subject PrivateEndpointConnection_Status_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_Status_SubResourceEmbedded
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

// Generator of PrivateEndpointConnection_Status_SubResourceEmbedded instances for property testing - lazily
// instantiated by PrivateEndpointConnectionStatusSubResourceEmbeddedGenerator()
var privateEndpointConnectionStatusSubResourceEmbeddedGenerator gopter.Gen

// PrivateEndpointConnectionStatusSubResourceEmbeddedGenerator returns a generator of PrivateEndpointConnection_Status_SubResourceEmbedded instances for property testing.
func PrivateEndpointConnectionStatusSubResourceEmbeddedGenerator() gopter.Gen {
	if privateEndpointConnectionStatusSubResourceEmbeddedGenerator != nil {
		return privateEndpointConnectionStatusSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbedded(generators)
	privateEndpointConnectionStatusSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_Status_SubResourceEmbedded{}), generators)

	return privateEndpointConnectionStatusSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuStatus, SkuStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuStatus runs a test to see if a specific instance of Sku_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuStatus(subject Sku_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_Status
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

// Generator of Sku_Status instances for property testing - lazily instantiated by SkuStatusGenerator()
var skuStatusGenerator gopter.Gen

// SkuStatusGenerator returns a generator of Sku_Status instances for property testing.
func SkuStatusGenerator() gopter.Gen {
	if skuStatusGenerator != nil {
		return skuStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuStatus(generators)
	skuStatusGenerator = gen.Struct(reflect.TypeOf(Sku_Status{}), generators)

	return skuStatusGenerator
}

// AddIndependentPropertyGeneratorsForSkuStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuStatus(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}
