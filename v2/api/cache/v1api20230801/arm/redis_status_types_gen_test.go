// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_ManagedServiceIdentity_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedServiceIdentity_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedServiceIdentity_STATUS, ManagedServiceIdentity_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedServiceIdentity_STATUS runs a test to see if a specific instance of ManagedServiceIdentity_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedServiceIdentity_STATUS(subject ManagedServiceIdentity_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedServiceIdentity_STATUS
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

// Generator of ManagedServiceIdentity_STATUS instances for property testing - lazily instantiated by
// ManagedServiceIdentity_STATUSGenerator()
var managedServiceIdentity_STATUSGenerator gopter.Gen

// ManagedServiceIdentity_STATUSGenerator returns a generator of ManagedServiceIdentity_STATUS instances for property testing.
// We first initialize managedServiceIdentity_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedServiceIdentity_STATUSGenerator() gopter.Gen {
	if managedServiceIdentity_STATUSGenerator != nil {
		return managedServiceIdentity_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedServiceIdentity_STATUS(generators)
	managedServiceIdentity_STATUSGenerator = gen.Struct(reflect.TypeOf(ManagedServiceIdentity_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedServiceIdentity_STATUS(generators)
	AddRelatedPropertyGeneratorsForManagedServiceIdentity_STATUS(generators)
	managedServiceIdentity_STATUSGenerator = gen.Struct(reflect.TypeOf(ManagedServiceIdentity_STATUS{}), generators)

	return managedServiceIdentity_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForManagedServiceIdentity_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedServiceIdentity_STATUS(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		ManagedServiceIdentityType_STATUS_None,
		ManagedServiceIdentityType_STATUS_SystemAssigned,
		ManagedServiceIdentityType_STATUS_SystemAssignedUserAssigned,
		ManagedServiceIdentityType_STATUS_UserAssigned))
}

// AddRelatedPropertyGeneratorsForManagedServiceIdentity_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedServiceIdentity_STATUS(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		UserAssignedIdentity_STATUSGenerator())
}

func Test_PrivateEndpointConnection_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS, PrivateEndpointConnection_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS runs a test to see if a specific instance of PrivateEndpointConnection_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS(subject PrivateEndpointConnection_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS
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

// Generator of PrivateEndpointConnection_STATUS instances for property testing - lazily instantiated by
// PrivateEndpointConnection_STATUSGenerator()
var privateEndpointConnection_STATUSGenerator gopter.Gen

// PrivateEndpointConnection_STATUSGenerator returns a generator of PrivateEndpointConnection_STATUS instances for property testing.
func PrivateEndpointConnection_STATUSGenerator() gopter.Gen {
	if privateEndpointConnection_STATUSGenerator != nil {
		return privateEndpointConnection_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS(generators)
	privateEndpointConnection_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS{}), generators)

	return privateEndpointConnection_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisInstanceDetails_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisInstanceDetails_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisInstanceDetails_STATUS, RedisInstanceDetails_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisInstanceDetails_STATUS runs a test to see if a specific instance of RedisInstanceDetails_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisInstanceDetails_STATUS(subject RedisInstanceDetails_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisInstanceDetails_STATUS
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

// Generator of RedisInstanceDetails_STATUS instances for property testing - lazily instantiated by
// RedisInstanceDetails_STATUSGenerator()
var redisInstanceDetails_STATUSGenerator gopter.Gen

// RedisInstanceDetails_STATUSGenerator returns a generator of RedisInstanceDetails_STATUS instances for property testing.
func RedisInstanceDetails_STATUSGenerator() gopter.Gen {
	if redisInstanceDetails_STATUSGenerator != nil {
		return redisInstanceDetails_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisInstanceDetails_STATUS(generators)
	redisInstanceDetails_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisInstanceDetails_STATUS{}), generators)

	return redisInstanceDetails_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisInstanceDetails_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisInstanceDetails_STATUS(gens map[string]gopter.Gen) {
	gens["IsMaster"] = gen.PtrOf(gen.Bool())
	gens["IsPrimary"] = gen.PtrOf(gen.Bool())
	gens["NonSslPort"] = gen.PtrOf(gen.Int())
	gens["ShardId"] = gen.PtrOf(gen.Int())
	gens["SslPort"] = gen.PtrOf(gen.Int())
	gens["Zone"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisLinkedServer_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer_STATUS, RedisLinkedServer_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer_STATUS runs a test to see if a specific instance of RedisLinkedServer_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer_STATUS(subject RedisLinkedServer_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer_STATUS
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

// Generator of RedisLinkedServer_STATUS instances for property testing - lazily instantiated by
// RedisLinkedServer_STATUSGenerator()
var redisLinkedServer_STATUSGenerator gopter.Gen

// RedisLinkedServer_STATUSGenerator returns a generator of RedisLinkedServer_STATUS instances for property testing.
func RedisLinkedServer_STATUSGenerator() gopter.Gen {
	if redisLinkedServer_STATUSGenerator != nil {
		return redisLinkedServer_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServer_STATUS(generators)
	redisLinkedServer_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_STATUS{}), generators)

	return redisLinkedServer_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServer_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServer_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisProperties_RedisConfiguration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisProperties_RedisConfiguration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisProperties_RedisConfiguration_STATUS, RedisProperties_RedisConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisProperties_RedisConfiguration_STATUS runs a test to see if a specific instance of RedisProperties_RedisConfiguration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisProperties_RedisConfiguration_STATUS(subject RedisProperties_RedisConfiguration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisProperties_RedisConfiguration_STATUS
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

// Generator of RedisProperties_RedisConfiguration_STATUS instances for property testing - lazily instantiated by
// RedisProperties_RedisConfiguration_STATUSGenerator()
var redisProperties_RedisConfiguration_STATUSGenerator gopter.Gen

// RedisProperties_RedisConfiguration_STATUSGenerator returns a generator of RedisProperties_RedisConfiguration_STATUS instances for property testing.
func RedisProperties_RedisConfiguration_STATUSGenerator() gopter.Gen {
	if redisProperties_RedisConfiguration_STATUSGenerator != nil {
		return redisProperties_RedisConfiguration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisProperties_RedisConfiguration_STATUS(generators)
	redisProperties_RedisConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisProperties_RedisConfiguration_STATUS{}), generators)

	return redisProperties_RedisConfiguration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisProperties_RedisConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisProperties_RedisConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["AadEnabled"] = gen.PtrOf(gen.AlphaString())
	gens["AofBackupEnabled"] = gen.PtrOf(gen.AlphaString())
	gens["AofStorageConnectionString0"] = gen.PtrOf(gen.AlphaString())
	gens["AofStorageConnectionString1"] = gen.PtrOf(gen.AlphaString())
	gens["Authnotrequired"] = gen.PtrOf(gen.AlphaString())
	gens["Maxclients"] = gen.PtrOf(gen.AlphaString())
	gens["MaxfragmentationmemoryReserved"] = gen.PtrOf(gen.AlphaString())
	gens["MaxmemoryDelta"] = gen.PtrOf(gen.AlphaString())
	gens["MaxmemoryPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["MaxmemoryReserved"] = gen.PtrOf(gen.AlphaString())
	gens["NotifyKeyspaceEvents"] = gen.PtrOf(gen.AlphaString())
	gens["PreferredDataArchiveAuthMethod"] = gen.PtrOf(gen.AlphaString())
	gens["PreferredDataPersistenceAuthMethod"] = gen.PtrOf(gen.AlphaString())
	gens["RdbBackupEnabled"] = gen.PtrOf(gen.AlphaString())
	gens["RdbBackupFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["RdbBackupMaxSnapshotCount"] = gen.PtrOf(gen.AlphaString())
	gens["RdbStorageConnectionString"] = gen.PtrOf(gen.AlphaString())
	gens["StorageSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["ZonalConfiguration"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisProperties_STATUS, RedisProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisProperties_STATUS runs a test to see if a specific instance of RedisProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisProperties_STATUS(subject RedisProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisProperties_STATUS
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

// Generator of RedisProperties_STATUS instances for property testing - lazily instantiated by
// RedisProperties_STATUSGenerator()
var redisProperties_STATUSGenerator gopter.Gen

// RedisProperties_STATUSGenerator returns a generator of RedisProperties_STATUS instances for property testing.
// We first initialize redisProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisProperties_STATUSGenerator() gopter.Gen {
	if redisProperties_STATUSGenerator != nil {
		return redisProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisProperties_STATUS(generators)
	redisProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForRedisProperties_STATUS(generators)
	redisProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisProperties_STATUS{}), generators)

	return redisProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisProperties_STATUS(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(RedisProperties_MinimumTlsVersion_STATUS_10, RedisProperties_MinimumTlsVersion_STATUS_11, RedisProperties_MinimumTlsVersion_STATUS_12))
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		RedisProperties_ProvisioningState_STATUS_ConfiguringAAD,
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
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["SslPort"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["SubnetId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["UpdateChannel"] = gen.PtrOf(gen.OneConstOf(RedisProperties_UpdateChannel_STATUS_Preview, RedisProperties_UpdateChannel_STATUS_Stable))
}

// AddRelatedPropertyGeneratorsForRedisProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Instances"] = gen.SliceOf(RedisInstanceDetails_STATUSGenerator())
	gens["LinkedServers"] = gen.SliceOf(RedisLinkedServer_STATUSGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUSGenerator())
	gens["RedisConfiguration"] = gen.PtrOf(RedisProperties_RedisConfiguration_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
}

func Test_Redis_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_STATUS, Redis_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_STATUS runs a test to see if a specific instance of Redis_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_STATUS(subject Redis_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_STATUS
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

// Generator of Redis_STATUS instances for property testing - lazily instantiated by Redis_STATUSGenerator()
var redis_STATUSGenerator gopter.Gen

// Redis_STATUSGenerator returns a generator of Redis_STATUS instances for property testing.
// We first initialize redis_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Redis_STATUSGenerator() gopter.Gen {
	if redis_STATUSGenerator != nil {
		return redis_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_STATUS(generators)
	redis_STATUSGenerator = gen.Struct(reflect.TypeOf(Redis_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_STATUS(generators)
	AddRelatedPropertyGeneratorsForRedis_STATUS(generators)
	redis_STATUSGenerator = gen.Struct(reflect.TypeOf(Redis_STATUS{}), generators)

	return redis_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedis_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedis_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis_STATUS(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(ManagedServiceIdentity_STATUSGenerator())
	gens["Properties"] = gen.PtrOf(RedisProperties_STATUSGenerator())
}

func Test_Sku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS, Sku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS runs a test to see if a specific instance of Sku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS(subject Sku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS
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

// Generator of Sku_STATUS instances for property testing - lazily instantiated by Sku_STATUSGenerator()
var sku_STATUSGenerator gopter.Gen

// Sku_STATUSGenerator returns a generator of Sku_STATUS instances for property testing.
func Sku_STATUSGenerator() gopter.Gen {
	if sku_STATUSGenerator != nil {
		return sku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS(generators)
	sku_STATUSGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS{}), generators)

	return sku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.OneConstOf(Sku_Family_STATUS_C, Sku_Family_STATUS_P))
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_STATUS_Basic, Sku_Name_STATUS_Premium, Sku_Name_STATUS_Standard))
}

func Test_UserAssignedIdentity_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_STATUS, UserAssignedIdentity_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_STATUS runs a test to see if a specific instance of UserAssignedIdentity_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_STATUS(subject UserAssignedIdentity_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_STATUS
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

// Generator of UserAssignedIdentity_STATUS instances for property testing - lazily instantiated by
// UserAssignedIdentity_STATUSGenerator()
var userAssignedIdentity_STATUSGenerator gopter.Gen

// UserAssignedIdentity_STATUSGenerator returns a generator of UserAssignedIdentity_STATUS instances for property testing.
func UserAssignedIdentity_STATUSGenerator() gopter.Gen {
	if userAssignedIdentity_STATUSGenerator != nil {
		return userAssignedIdentity_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS(generators)
	userAssignedIdentity_STATUSGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUS{}), generators)

	return userAssignedIdentity_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}
