// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

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

func Test_RedisResource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisResource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisResourceSTATUSARM, RedisResourceSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisResourceSTATUSARM runs a test to see if a specific instance of RedisResource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisResourceSTATUSARM(subject RedisResource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisResource_STATUSARM
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

// Generator of RedisResource_STATUSARM instances for property testing - lazily instantiated by
// RedisResourceSTATUSARMGenerator()
var redisResourceSTATUSARMGenerator gopter.Gen

// RedisResourceSTATUSARMGenerator returns a generator of RedisResource_STATUSARM instances for property testing.
// We first initialize redisResourceSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisResourceSTATUSARMGenerator() gopter.Gen {
	if redisResourceSTATUSARMGenerator != nil {
		return redisResourceSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResourceSTATUSARM(generators)
	redisResourceSTATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisResource_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResourceSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForRedisResourceSTATUSARM(generators)
	redisResourceSTATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisResource_STATUSARM{}), generators)

	return redisResourceSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisResourceSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisResourceSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisResourceSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisResourceSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisPropertiesSTATUSARMGenerator())
}

func Test_RedisProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPropertiesSTATUSARM, RedisPropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPropertiesSTATUSARM runs a test to see if a specific instance of RedisProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPropertiesSTATUSARM(subject RedisProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisProperties_STATUSARM
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

// Generator of RedisProperties_STATUSARM instances for property testing - lazily instantiated by
// RedisPropertiesSTATUSARMGenerator()
var redisPropertiesSTATUSARMGenerator gopter.Gen

// RedisPropertiesSTATUSARMGenerator returns a generator of RedisProperties_STATUSARM instances for property testing.
// We first initialize redisPropertiesSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisPropertiesSTATUSARMGenerator() gopter.Gen {
	if redisPropertiesSTATUSARMGenerator != nil {
		return redisPropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPropertiesSTATUSARM(generators)
	redisPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPropertiesSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForRedisPropertiesSTATUSARM(generators)
	redisPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisProperties_STATUSARM{}), generators)

	return redisPropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisPropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(RedisPropertiesSTATUSMinimumTlsVersion_10, RedisPropertiesSTATUSMinimumTlsVersion_11, RedisPropertiesSTATUSMinimumTlsVersion_12))
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		RedisPropertiesSTATUSProvisioningState_Creating,
		RedisPropertiesSTATUSProvisioningState_Deleting,
		RedisPropertiesSTATUSProvisioningState_Disabled,
		RedisPropertiesSTATUSProvisioningState_Failed,
		RedisPropertiesSTATUSProvisioningState_Linking,
		RedisPropertiesSTATUSProvisioningState_Provisioning,
		RedisPropertiesSTATUSProvisioningState_RecoveringScaleFailure,
		RedisPropertiesSTATUSProvisioningState_Scaling,
		RedisPropertiesSTATUSProvisioningState_Succeeded,
		RedisPropertiesSTATUSProvisioningState_Unlinking,
		RedisPropertiesSTATUSProvisioningState_Unprovisioning,
		RedisPropertiesSTATUSProvisioningState_Updating))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RedisPropertiesSTATUSPublicNetworkAccess_Disabled, RedisPropertiesSTATUSPublicNetworkAccess_Enabled))
	gens["RedisConfiguration"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["SslPort"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["SubnetId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisPropertiesSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Instances"] = gen.SliceOf(RedisInstanceDetailsSTATUSARMGenerator())
	gens["LinkedServers"] = gen.SliceOf(RedisLinkedServerSTATUSARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionSTATUSSubResourceEmbeddedARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuSTATUSARMGenerator())
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

func Test_RedisInstanceDetails_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisInstanceDetails_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisInstanceDetailsSTATUSARM, RedisInstanceDetailsSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisInstanceDetailsSTATUSARM runs a test to see if a specific instance of RedisInstanceDetails_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisInstanceDetailsSTATUSARM(subject RedisInstanceDetails_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisInstanceDetails_STATUSARM
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

// Generator of RedisInstanceDetails_STATUSARM instances for property testing - lazily instantiated by
// RedisInstanceDetailsSTATUSARMGenerator()
var redisInstanceDetailsSTATUSARMGenerator gopter.Gen

// RedisInstanceDetailsSTATUSARMGenerator returns a generator of RedisInstanceDetails_STATUSARM instances for property testing.
func RedisInstanceDetailsSTATUSARMGenerator() gopter.Gen {
	if redisInstanceDetailsSTATUSARMGenerator != nil {
		return redisInstanceDetailsSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisInstanceDetailsSTATUSARM(generators)
	redisInstanceDetailsSTATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisInstanceDetails_STATUSARM{}), generators)

	return redisInstanceDetailsSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisInstanceDetailsSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisInstanceDetailsSTATUSARM(gens map[string]gopter.Gen) {
	gens["IsMaster"] = gen.PtrOf(gen.Bool())
	gens["IsPrimary"] = gen.PtrOf(gen.Bool())
	gens["NonSslPort"] = gen.PtrOf(gen.Int())
	gens["ShardId"] = gen.PtrOf(gen.Int())
	gens["SslPort"] = gen.PtrOf(gen.Int())
	gens["Zone"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisLinkedServer_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerSTATUSARM, RedisLinkedServerSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerSTATUSARM runs a test to see if a specific instance of RedisLinkedServer_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerSTATUSARM(subject RedisLinkedServer_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer_STATUSARM
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

// Generator of RedisLinkedServer_STATUSARM instances for property testing - lazily instantiated by
// RedisLinkedServerSTATUSARMGenerator()
var redisLinkedServerSTATUSARMGenerator gopter.Gen

// RedisLinkedServerSTATUSARMGenerator returns a generator of RedisLinkedServer_STATUSARM instances for property testing.
func RedisLinkedServerSTATUSARMGenerator() gopter.Gen {
	if redisLinkedServerSTATUSARMGenerator != nil {
		return redisLinkedServerSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerSTATUSARM(generators)
	redisLinkedServerSTATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_STATUSARM{}), generators)

	return redisLinkedServerSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
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
	gens["Family"] = gen.PtrOf(gen.OneConstOf(SkuSTATUSFamily_C, SkuSTATUSFamily_P))
	gens["Name"] = gen.PtrOf(gen.OneConstOf(SkuSTATUSName_Basic, SkuSTATUSName_Premium, SkuSTATUSName_Standard))
}
