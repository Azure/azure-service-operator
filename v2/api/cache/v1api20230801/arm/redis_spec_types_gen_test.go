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

func Test_ManagedServiceIdentity_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedServiceIdentity via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedServiceIdentity, ManagedServiceIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedServiceIdentity runs a test to see if a specific instance of ManagedServiceIdentity round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedServiceIdentity(subject ManagedServiceIdentity) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedServiceIdentity
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

// Generator of ManagedServiceIdentity instances for property testing - lazily instantiated by
// ManagedServiceIdentityGenerator()
var managedServiceIdentityGenerator gopter.Gen

// ManagedServiceIdentityGenerator returns a generator of ManagedServiceIdentity instances for property testing.
// We first initialize managedServiceIdentityGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedServiceIdentityGenerator() gopter.Gen {
	if managedServiceIdentityGenerator != nil {
		return managedServiceIdentityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedServiceIdentity(generators)
	managedServiceIdentityGenerator = gen.Struct(reflect.TypeOf(ManagedServiceIdentity{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedServiceIdentity(generators)
	AddRelatedPropertyGeneratorsForManagedServiceIdentity(generators)
	managedServiceIdentityGenerator = gen.Struct(reflect.TypeOf(ManagedServiceIdentity{}), generators)

	return managedServiceIdentityGenerator
}

// AddIndependentPropertyGeneratorsForManagedServiceIdentity is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedServiceIdentity(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		ManagedServiceIdentityType_None,
		ManagedServiceIdentityType_SystemAssigned,
		ManagedServiceIdentityType_SystemAssignedUserAssigned,
		ManagedServiceIdentityType_UserAssigned))
}

// AddRelatedPropertyGeneratorsForManagedServiceIdentity is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedServiceIdentity(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		UserAssignedIdentityDetailsGenerator())
}

func Test_RedisCreateProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisCreateProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisCreateProperties, RedisCreatePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisCreateProperties runs a test to see if a specific instance of RedisCreateProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisCreateProperties(subject RedisCreateProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisCreateProperties
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

// Generator of RedisCreateProperties instances for property testing - lazily instantiated by
// RedisCreatePropertiesGenerator()
var redisCreatePropertiesGenerator gopter.Gen

// RedisCreatePropertiesGenerator returns a generator of RedisCreateProperties instances for property testing.
// We first initialize redisCreatePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisCreatePropertiesGenerator() gopter.Gen {
	if redisCreatePropertiesGenerator != nil {
		return redisCreatePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateProperties(generators)
	redisCreatePropertiesGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateProperties(generators)
	AddRelatedPropertyGeneratorsForRedisCreateProperties(generators)
	redisCreatePropertiesGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties{}), generators)

	return redisCreatePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForRedisCreateProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisCreateProperties(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(RedisCreateProperties_MinimumTlsVersion_10, RedisCreateProperties_MinimumTlsVersion_11, RedisCreateProperties_MinimumTlsVersion_12))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RedisCreateProperties_PublicNetworkAccess_Disabled, RedisCreateProperties_PublicNetworkAccess_Enabled))
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["SubnetId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["UpdateChannel"] = gen.PtrOf(gen.OneConstOf(RedisCreateProperties_UpdateChannel_Preview, RedisCreateProperties_UpdateChannel_Stable))
}

// AddRelatedPropertyGeneratorsForRedisCreateProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisCreateProperties(gens map[string]gopter.Gen) {
	gens["RedisConfiguration"] = gen.PtrOf(RedisCreateProperties_RedisConfigurationGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_RedisCreateProperties_RedisConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisCreateProperties_RedisConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisCreateProperties_RedisConfiguration, RedisCreateProperties_RedisConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisCreateProperties_RedisConfiguration runs a test to see if a specific instance of RedisCreateProperties_RedisConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisCreateProperties_RedisConfiguration(subject RedisCreateProperties_RedisConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisCreateProperties_RedisConfiguration
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

// Generator of RedisCreateProperties_RedisConfiguration instances for property testing - lazily instantiated by
// RedisCreateProperties_RedisConfigurationGenerator()
var redisCreateProperties_RedisConfigurationGenerator gopter.Gen

// RedisCreateProperties_RedisConfigurationGenerator returns a generator of RedisCreateProperties_RedisConfiguration instances for property testing.
func RedisCreateProperties_RedisConfigurationGenerator() gopter.Gen {
	if redisCreateProperties_RedisConfigurationGenerator != nil {
		return redisCreateProperties_RedisConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateProperties_RedisConfiguration(generators)
	redisCreateProperties_RedisConfigurationGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties_RedisConfiguration{}), generators)

	return redisCreateProperties_RedisConfigurationGenerator
}

// AddIndependentPropertyGeneratorsForRedisCreateProperties_RedisConfiguration is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisCreateProperties_RedisConfiguration(gens map[string]gopter.Gen) {
	gens["AadEnabled"] = gen.PtrOf(gen.AlphaString())
	gens["AofBackupEnabled"] = gen.PtrOf(gen.AlphaString())
	gens["AofStorageConnectionString0"] = gen.PtrOf(gen.AlphaString())
	gens["AofStorageConnectionString1"] = gen.PtrOf(gen.AlphaString())
	gens["Authnotrequired"] = gen.PtrOf(gen.AlphaString())
	gens["MaxfragmentationmemoryReserved"] = gen.PtrOf(gen.AlphaString())
	gens["MaxmemoryDelta"] = gen.PtrOf(gen.AlphaString())
	gens["MaxmemoryPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["MaxmemoryReserved"] = gen.PtrOf(gen.AlphaString())
	gens["NotifyKeyspaceEvents"] = gen.PtrOf(gen.AlphaString())
	gens["PreferredDataPersistenceAuthMethod"] = gen.PtrOf(gen.AlphaString())
	gens["RdbBackupEnabled"] = gen.PtrOf(gen.AlphaString())
	gens["RdbBackupFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["RdbBackupMaxSnapshotCount"] = gen.PtrOf(gen.AlphaString())
	gens["RdbStorageConnectionString"] = gen.PtrOf(gen.AlphaString())
	gens["StorageSubscriptionId"] = gen.PtrOf(gen.AlphaString())
}

func Test_Redis_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_Spec, Redis_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_Spec runs a test to see if a specific instance of Redis_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_Spec(subject Redis_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_Spec
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

// Generator of Redis_Spec instances for property testing - lazily instantiated by Redis_SpecGenerator()
var redis_SpecGenerator gopter.Gen

// Redis_SpecGenerator returns a generator of Redis_Spec instances for property testing.
// We first initialize redis_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Redis_SpecGenerator() gopter.Gen {
	if redis_SpecGenerator != nil {
		return redis_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_Spec(generators)
	redis_SpecGenerator = gen.Struct(reflect.TypeOf(Redis_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_Spec(generators)
	AddRelatedPropertyGeneratorsForRedis_Spec(generators)
	redis_SpecGenerator = gen.Struct(reflect.TypeOf(Redis_Spec{}), generators)

	return redis_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedis_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedis_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis_Spec(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(ManagedServiceIdentityGenerator())
	gens["Properties"] = gen.PtrOf(RedisCreatePropertiesGenerator())
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
	gens["Family"] = gen.PtrOf(gen.OneConstOf(Sku_Family_C, Sku_Family_P))
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_Basic, Sku_Name_Premium, Sku_Name_Standard))
}

func Test_UserAssignedIdentityDetails_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityDetails via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityDetails, UserAssignedIdentityDetailsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityDetails runs a test to see if a specific instance of UserAssignedIdentityDetails round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityDetails(subject UserAssignedIdentityDetails) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityDetails
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

// Generator of UserAssignedIdentityDetails instances for property testing - lazily instantiated by
// UserAssignedIdentityDetailsGenerator()
var userAssignedIdentityDetailsGenerator gopter.Gen

// UserAssignedIdentityDetailsGenerator returns a generator of UserAssignedIdentityDetails instances for property testing.
func UserAssignedIdentityDetailsGenerator() gopter.Gen {
	if userAssignedIdentityDetailsGenerator != nil {
		return userAssignedIdentityDetailsGenerator
	}

	generators := make(map[string]gopter.Gen)
	userAssignedIdentityDetailsGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityDetails{}), generators)

	return userAssignedIdentityDetailsGenerator
}
