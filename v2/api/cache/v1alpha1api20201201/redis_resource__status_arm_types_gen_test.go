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

func Test_RedisResource_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisResource_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisResourceStatusARM, RedisResourceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisResourceStatusARM runs a test to see if a specific instance of RedisResource_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisResourceStatusARM(subject RedisResource_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisResource_StatusARM
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

// Generator of RedisResource_StatusARM instances for property testing - lazily instantiated by
//RedisResourceStatusARMGenerator()
var redisResourceStatusARMGenerator gopter.Gen

// RedisResourceStatusARMGenerator returns a generator of RedisResource_StatusARM instances for property testing.
// We first initialize redisResourceStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisResourceStatusARMGenerator() gopter.Gen {
	if redisResourceStatusARMGenerator != nil {
		return redisResourceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResourceStatusARM(generators)
	redisResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisResource_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResourceStatusARM(generators)
	AddRelatedPropertyGeneratorsForRedisResourceStatusARM(generators)
	redisResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisResource_StatusARM{}), generators)

	return redisResourceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisResourceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisResourceStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisResourceStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisResourceStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisPropertiesStatusARMGenerator())
}

func Test_RedisProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPropertiesStatusARM, RedisPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPropertiesStatusARM runs a test to see if a specific instance of RedisProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPropertiesStatusARM(subject RedisProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisProperties_StatusARM
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

// Generator of RedisProperties_StatusARM instances for property testing - lazily instantiated by
//RedisPropertiesStatusARMGenerator()
var redisPropertiesStatusARMGenerator gopter.Gen

// RedisPropertiesStatusARMGenerator returns a generator of RedisProperties_StatusARM instances for property testing.
// We first initialize redisPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisPropertiesStatusARMGenerator() gopter.Gen {
	if redisPropertiesStatusARMGenerator != nil {
		return redisPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPropertiesStatusARM(generators)
	redisPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForRedisPropertiesStatusARM(generators)
	redisPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisProperties_StatusARM{}), generators)

	return redisPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(RedisPropertiesStatusMinimumTlsVersion10, RedisPropertiesStatusMinimumTlsVersion11, RedisPropertiesStatusMinimumTlsVersion12))
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		RedisPropertiesStatusProvisioningStateCreating,
		RedisPropertiesStatusProvisioningStateDeleting,
		RedisPropertiesStatusProvisioningStateDisabled,
		RedisPropertiesStatusProvisioningStateFailed,
		RedisPropertiesStatusProvisioningStateLinking,
		RedisPropertiesStatusProvisioningStateProvisioning,
		RedisPropertiesStatusProvisioningStateRecoveringScaleFailure,
		RedisPropertiesStatusProvisioningStateScaling,
		RedisPropertiesStatusProvisioningStateSucceeded,
		RedisPropertiesStatusProvisioningStateUnlinking,
		RedisPropertiesStatusProvisioningStateUnprovisioning,
		RedisPropertiesStatusProvisioningStateUpdating))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RedisPropertiesStatusPublicNetworkAccessDisabled, RedisPropertiesStatusPublicNetworkAccessEnabled))
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

// AddRelatedPropertyGeneratorsForRedisPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["AccessKeys"] = gen.PtrOf(RedisAccessKeysStatusARMGenerator())
	gens["Instances"] = gen.SliceOf(RedisInstanceDetailsStatusARMGenerator())
	gens["LinkedServers"] = gen.SliceOf(RedisLinkedServerStatusARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator())
	gens["Sku"] = SkuStatusARMGenerator()
}

func Test_PrivateEndpointConnection_Status_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
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
//instantiated by PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator()
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

func Test_RedisAccessKeys_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisAccessKeys_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisAccessKeysStatusARM, RedisAccessKeysStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisAccessKeysStatusARM runs a test to see if a specific instance of RedisAccessKeys_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisAccessKeysStatusARM(subject RedisAccessKeys_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisAccessKeys_StatusARM
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

// Generator of RedisAccessKeys_StatusARM instances for property testing - lazily instantiated by
//RedisAccessKeysStatusARMGenerator()
var redisAccessKeysStatusARMGenerator gopter.Gen

// RedisAccessKeysStatusARMGenerator returns a generator of RedisAccessKeys_StatusARM instances for property testing.
func RedisAccessKeysStatusARMGenerator() gopter.Gen {
	if redisAccessKeysStatusARMGenerator != nil {
		return redisAccessKeysStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisAccessKeysStatusARM(generators)
	redisAccessKeysStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisAccessKeys_StatusARM{}), generators)

	return redisAccessKeysStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisAccessKeysStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisAccessKeysStatusARM(gens map[string]gopter.Gen) {
	gens["PrimaryKey"] = gen.PtrOf(gen.AlphaString())
	gens["SecondaryKey"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisInstanceDetails_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisInstanceDetails_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisInstanceDetailsStatusARM, RedisInstanceDetailsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisInstanceDetailsStatusARM runs a test to see if a specific instance of RedisInstanceDetails_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisInstanceDetailsStatusARM(subject RedisInstanceDetails_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisInstanceDetails_StatusARM
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

// Generator of RedisInstanceDetails_StatusARM instances for property testing - lazily instantiated by
//RedisInstanceDetailsStatusARMGenerator()
var redisInstanceDetailsStatusARMGenerator gopter.Gen

// RedisInstanceDetailsStatusARMGenerator returns a generator of RedisInstanceDetails_StatusARM instances for property testing.
func RedisInstanceDetailsStatusARMGenerator() gopter.Gen {
	if redisInstanceDetailsStatusARMGenerator != nil {
		return redisInstanceDetailsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisInstanceDetailsStatusARM(generators)
	redisInstanceDetailsStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisInstanceDetails_StatusARM{}), generators)

	return redisInstanceDetailsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisInstanceDetailsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisInstanceDetailsStatusARM(gens map[string]gopter.Gen) {
	gens["IsMaster"] = gen.PtrOf(gen.Bool())
	gens["IsPrimary"] = gen.PtrOf(gen.Bool())
	gens["NonSslPort"] = gen.PtrOf(gen.Int())
	gens["ShardId"] = gen.PtrOf(gen.Int())
	gens["SslPort"] = gen.PtrOf(gen.Int())
	gens["Zone"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisLinkedServer_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerStatusARM, RedisLinkedServerStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerStatusARM runs a test to see if a specific instance of RedisLinkedServer_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerStatusARM(subject RedisLinkedServer_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer_StatusARM
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

// Generator of RedisLinkedServer_StatusARM instances for property testing - lazily instantiated by
//RedisLinkedServerStatusARMGenerator()
var redisLinkedServerStatusARMGenerator gopter.Gen

// RedisLinkedServerStatusARMGenerator returns a generator of RedisLinkedServer_StatusARM instances for property testing.
func RedisLinkedServerStatusARMGenerator() gopter.Gen {
	if redisLinkedServerStatusARMGenerator != nil {
		return redisLinkedServerStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerStatusARM(generators)
	redisLinkedServerStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_StatusARM{}), generators)

	return redisLinkedServerStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
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
	gens["Capacity"] = gen.Int()
	gens["Family"] = gen.OneConstOf(SkuStatusFamilyC, SkuStatusFamilyP)
	gens["Name"] = gen.OneConstOf(SkuStatusNameBasic, SkuStatusNamePremium, SkuStatusNameStandard)
}
