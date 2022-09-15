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

func Test_RedisResource_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisResource_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisResource_STATUS_ARM, RedisResource_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisResource_STATUS_ARM runs a test to see if a specific instance of RedisResource_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisResource_STATUS_ARM(subject RedisResource_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisResource_STATUS_ARM
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

// Generator of RedisResource_STATUS_ARM instances for property testing - lazily instantiated by
// RedisResource_STATUS_ARMGenerator()
var redisResource_STATUS_ARMGenerator gopter.Gen

// RedisResource_STATUS_ARMGenerator returns a generator of RedisResource_STATUS_ARM instances for property testing.
// We first initialize redisResource_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisResource_STATUS_ARMGenerator() gopter.Gen {
	if redisResource_STATUS_ARMGenerator != nil {
		return redisResource_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResource_STATUS_ARM(generators)
	redisResource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RedisResource_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResource_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForRedisResource_STATUS_ARM(generators)
	redisResource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RedisResource_STATUS_ARM{}), generators)

	return redisResource_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisResource_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisResource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisResource_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisResource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisProperties_STATUS_ARMGenerator())
}

func Test_RedisProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisProperties_STATUS_ARM, RedisProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisProperties_STATUS_ARM runs a test to see if a specific instance of RedisProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisProperties_STATUS_ARM(subject RedisProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisProperties_STATUS_ARM
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

// Generator of RedisProperties_STATUS_ARM instances for property testing - lazily instantiated by
// RedisProperties_STATUS_ARMGenerator()
var redisProperties_STATUS_ARMGenerator gopter.Gen

// RedisProperties_STATUS_ARMGenerator returns a generator of RedisProperties_STATUS_ARM instances for property testing.
// We first initialize redisProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisProperties_STATUS_ARMGenerator() gopter.Gen {
	if redisProperties_STATUS_ARMGenerator != nil {
		return redisProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisProperties_STATUS_ARM(generators)
	redisProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RedisProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForRedisProperties_STATUS_ARM(generators)
	redisProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RedisProperties_STATUS_ARM{}), generators)

	return redisProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(RedisProperties_MinimumTlsVersion_STATUS_10, RedisProperties_MinimumTlsVersion_STATUS_11, RedisProperties_MinimumTlsVersion_STATUS_12))
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		RedisProperties_ProvisioningState_STATUS_Creating,
		RedisProperties_ProvisioningState_STATUS_Deleting,
		RedisProperties_ProvisioningState_STATUS_Disabled,
		RedisProperties_ProvisioningState_STATUS_Failed,
		RedisProperties_ProvisioningState_STATUS_Linking,
		RedisProperties_ProvisioningState_STATUS_Provisioning,
		RedisProperties_ProvisioningState_STATUS_RecoveringScaleFailure,
		RedisProperties_ProvisioningState_STATUS_Scaling,
		RedisProperties_ProvisioningState_STATUS_Succeeded,
		RedisProperties_ProvisioningState_STATUS_Unlinking,
		RedisProperties_ProvisioningState_STATUS_Unprovisioning,
		RedisProperties_ProvisioningState_STATUS_Updating))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RedisProperties_PublicNetworkAccess_STATUS_Disabled, RedisProperties_PublicNetworkAccess_STATUS_Enabled))
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

// AddRelatedPropertyGeneratorsForRedisProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Instances"] = gen.SliceOf(RedisInstanceDetails_STATUS_ARMGenerator())
	gens["LinkedServers"] = gen.SliceOf(RedisLinkedServer_STATUS_ARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUS_ARMGenerator())
}

func Test_PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM, PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM(subject PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM
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

// Generator of PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM instances for property testing - lazily
// instantiated by PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARMGenerator()
var privateEndpointConnection_STATUS_SubResourceEmbedded_ARMGenerator gopter.Gen

// PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARMGenerator returns a generator of PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM instances for property testing.
func PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if privateEndpointConnection_STATUS_SubResourceEmbedded_ARMGenerator != nil {
		return privateEndpointConnection_STATUS_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM(generators)
	privateEndpointConnection_STATUS_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM{}), generators)

	return privateEndpointConnection_STATUS_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisInstanceDetails_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisInstanceDetails_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisInstanceDetails_STATUS_ARM, RedisInstanceDetails_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisInstanceDetails_STATUS_ARM runs a test to see if a specific instance of RedisInstanceDetails_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisInstanceDetails_STATUS_ARM(subject RedisInstanceDetails_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisInstanceDetails_STATUS_ARM
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

// Generator of RedisInstanceDetails_STATUS_ARM instances for property testing - lazily instantiated by
// RedisInstanceDetails_STATUS_ARMGenerator()
var redisInstanceDetails_STATUS_ARMGenerator gopter.Gen

// RedisInstanceDetails_STATUS_ARMGenerator returns a generator of RedisInstanceDetails_STATUS_ARM instances for property testing.
func RedisInstanceDetails_STATUS_ARMGenerator() gopter.Gen {
	if redisInstanceDetails_STATUS_ARMGenerator != nil {
		return redisInstanceDetails_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisInstanceDetails_STATUS_ARM(generators)
	redisInstanceDetails_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RedisInstanceDetails_STATUS_ARM{}), generators)

	return redisInstanceDetails_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisInstanceDetails_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisInstanceDetails_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["IsMaster"] = gen.PtrOf(gen.Bool())
	gens["IsPrimary"] = gen.PtrOf(gen.Bool())
	gens["NonSslPort"] = gen.PtrOf(gen.Int())
	gens["ShardId"] = gen.PtrOf(gen.Int())
	gens["SslPort"] = gen.PtrOf(gen.Int())
	gens["Zone"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisLinkedServer_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer_STATUS_ARM, RedisLinkedServer_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer_STATUS_ARM runs a test to see if a specific instance of RedisLinkedServer_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer_STATUS_ARM(subject RedisLinkedServer_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer_STATUS_ARM
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

// Generator of RedisLinkedServer_STATUS_ARM instances for property testing - lazily instantiated by
// RedisLinkedServer_STATUS_ARMGenerator()
var redisLinkedServer_STATUS_ARMGenerator gopter.Gen

// RedisLinkedServer_STATUS_ARMGenerator returns a generator of RedisLinkedServer_STATUS_ARM instances for property testing.
func RedisLinkedServer_STATUS_ARMGenerator() gopter.Gen {
	if redisLinkedServer_STATUS_ARMGenerator != nil {
		return redisLinkedServer_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServer_STATUS_ARM(generators)
	redisLinkedServer_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_STATUS_ARM{}), generators)

	return redisLinkedServer_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServer_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServer_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS_ARM, Sku_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS_ARM runs a test to see if a specific instance of Sku_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS_ARM(subject Sku_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS_ARM
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

// Generator of Sku_STATUS_ARM instances for property testing - lazily instantiated by Sku_STATUS_ARMGenerator()
var sku_STATUS_ARMGenerator gopter.Gen

// Sku_STATUS_ARMGenerator returns a generator of Sku_STATUS_ARM instances for property testing.
func Sku_STATUS_ARMGenerator() gopter.Gen {
	if sku_STATUS_ARMGenerator != nil {
		return sku_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS_ARM(generators)
	sku_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS_ARM{}), generators)

	return sku_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.OneConstOf(Sku_Family_STATUS_C, Sku_Family_STATUS_P))
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_STATUS_Basic, Sku_Name_STATUS_Premium, Sku_Name_STATUS_Standard))
}
