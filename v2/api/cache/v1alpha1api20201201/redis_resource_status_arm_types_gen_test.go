// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

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
		prop.ForAll(RunJSONSerializationTestForRedisResource_STATUSARM, RedisResource_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisResource_STATUSARM runs a test to see if a specific instance of RedisResource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisResource_STATUSARM(subject RedisResource_STATUSARM) string {
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
// RedisResource_STATUSARMGenerator()
var redisResource_STATUSARMGenerator gopter.Gen

// RedisResource_STATUSARMGenerator returns a generator of RedisResource_STATUSARM instances for property testing.
// We first initialize redisResource_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisResource_STATUSARMGenerator() gopter.Gen {
	if redisResource_STATUSARMGenerator != nil {
		return redisResource_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResource_STATUSARM(generators)
	redisResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisResource_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResource_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForRedisResource_STATUSARM(generators)
	redisResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisResource_STATUSARM{}), generators)

	return redisResource_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisResource_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisResource_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisProperties_STATUSARMGenerator())
}

func Test_RedisProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisProperties_STATUSARM, RedisProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisProperties_STATUSARM runs a test to see if a specific instance of RedisProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisProperties_STATUSARM(subject RedisProperties_STATUSARM) string {
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
// RedisProperties_STATUSARMGenerator()
var redisProperties_STATUSARMGenerator gopter.Gen

// RedisProperties_STATUSARMGenerator returns a generator of RedisProperties_STATUSARM instances for property testing.
// We first initialize redisProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisProperties_STATUSARMGenerator() gopter.Gen {
	if redisProperties_STATUSARMGenerator != nil {
		return redisProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisProperties_STATUSARM(generators)
	redisProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForRedisProperties_STATUSARM(generators)
	redisProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisProperties_STATUSARM{}), generators)

	return redisProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(RedisProperties_STATUS_MinimumTlsVersion_10, RedisProperties_STATUS_MinimumTlsVersion_11, RedisProperties_STATUS_MinimumTlsVersion_12))
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		RedisProperties_STATUS_ProvisioningState_Creating,
		RedisProperties_STATUS_ProvisioningState_Deleting,
		RedisProperties_STATUS_ProvisioningState_Disabled,
		RedisProperties_STATUS_ProvisioningState_Failed,
		RedisProperties_STATUS_ProvisioningState_Linking,
		RedisProperties_STATUS_ProvisioningState_Provisioning,
		RedisProperties_STATUS_ProvisioningState_RecoveringScaleFailure,
		RedisProperties_STATUS_ProvisioningState_Scaling,
		RedisProperties_STATUS_ProvisioningState_Succeeded,
		RedisProperties_STATUS_ProvisioningState_Unlinking,
		RedisProperties_STATUS_ProvisioningState_Unprovisioning,
		RedisProperties_STATUS_ProvisioningState_Updating))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RedisProperties_STATUS_PublicNetworkAccess_Disabled, RedisProperties_STATUS_PublicNetworkAccess_Enabled))
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

// AddRelatedPropertyGeneratorsForRedisProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Instances"] = gen.SliceOf(RedisInstanceDetails_STATUSARMGenerator())
	gens["LinkedServers"] = gen.SliceOf(RedisLinkedServer_STATUSARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSARMGenerator())
}

func Test_PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM, PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(subject PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM) string {
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
// instantiated by PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator()
var privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator gopter.Gen

// PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator returns a generator of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM instances for property testing.
func PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator != nil {
		return privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(generators)
	privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM{}), generators)

	return privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
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
		prop.ForAll(RunJSONSerializationTestForRedisInstanceDetails_STATUSARM, RedisInstanceDetails_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisInstanceDetails_STATUSARM runs a test to see if a specific instance of RedisInstanceDetails_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisInstanceDetails_STATUSARM(subject RedisInstanceDetails_STATUSARM) string {
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
// RedisInstanceDetails_STATUSARMGenerator()
var redisInstanceDetails_STATUSARMGenerator gopter.Gen

// RedisInstanceDetails_STATUSARMGenerator returns a generator of RedisInstanceDetails_STATUSARM instances for property testing.
func RedisInstanceDetails_STATUSARMGenerator() gopter.Gen {
	if redisInstanceDetails_STATUSARMGenerator != nil {
		return redisInstanceDetails_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisInstanceDetails_STATUSARM(generators)
	redisInstanceDetails_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisInstanceDetails_STATUSARM{}), generators)

	return redisInstanceDetails_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisInstanceDetails_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisInstanceDetails_STATUSARM(gens map[string]gopter.Gen) {
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
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer_STATUSARM, RedisLinkedServer_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer_STATUSARM runs a test to see if a specific instance of RedisLinkedServer_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer_STATUSARM(subject RedisLinkedServer_STATUSARM) string {
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
// RedisLinkedServer_STATUSARMGenerator()
var redisLinkedServer_STATUSARMGenerator gopter.Gen

// RedisLinkedServer_STATUSARMGenerator returns a generator of RedisLinkedServer_STATUSARM instances for property testing.
func RedisLinkedServer_STATUSARMGenerator() gopter.Gen {
	if redisLinkedServer_STATUSARMGenerator != nil {
		return redisLinkedServer_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServer_STATUSARM(generators)
	redisLinkedServer_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_STATUSARM{}), generators)

	return redisLinkedServer_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServer_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServer_STATUSARM(gens map[string]gopter.Gen) {
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
		prop.ForAll(RunJSONSerializationTestForSku_STATUSARM, Sku_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUSARM runs a test to see if a specific instance of Sku_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUSARM(subject Sku_STATUSARM) string {
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

// Generator of Sku_STATUSARM instances for property testing - lazily instantiated by Sku_STATUSARMGenerator()
var sku_STATUSARMGenerator gopter.Gen

// Sku_STATUSARMGenerator returns a generator of Sku_STATUSARM instances for property testing.
func Sku_STATUSARMGenerator() gopter.Gen {
	if sku_STATUSARMGenerator != nil {
		return sku_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUSARM(generators)
	sku_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUSARM{}), generators)

	return sku_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUSARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.OneConstOf(Sku_STATUS_Family_C, Sku_STATUS_Family_P))
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_STATUS_Name_Basic, Sku_STATUS_Name_Premium, Sku_STATUS_Name_Standard))
}
