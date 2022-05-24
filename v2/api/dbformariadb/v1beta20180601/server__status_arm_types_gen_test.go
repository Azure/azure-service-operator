// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

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

func Test_Server_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Server_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerStatusARM, ServerStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerStatusARM runs a test to see if a specific instance of Server_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerStatusARM(subject Server_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Server_StatusARM
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

// Generator of Server_StatusARM instances for property testing - lazily instantiated by ServerStatusARMGenerator()
var serverStatusARMGenerator gopter.Gen

// ServerStatusARMGenerator returns a generator of Server_StatusARM instances for property testing.
// We first initialize serverStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerStatusARMGenerator() gopter.Gen {
	if serverStatusARMGenerator != nil {
		return serverStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerStatusARM(generators)
	serverStatusARMGenerator = gen.Struct(reflect.TypeOf(Server_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerStatusARM(generators)
	AddRelatedPropertyGeneratorsForServerStatusARM(generators)
	serverStatusARMGenerator = gen.Struct(reflect.TypeOf(Server_StatusARM{}), generators)

	return serverStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForServerStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServerStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerPropertiesStatusARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuStatusARMGenerator())
}

func Test_ServerProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPropertiesStatusARM, ServerPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPropertiesStatusARM runs a test to see if a specific instance of ServerProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPropertiesStatusARM(subject ServerProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerProperties_StatusARM
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

// Generator of ServerProperties_StatusARM instances for property testing - lazily instantiated by
// ServerPropertiesStatusARMGenerator()
var serverPropertiesStatusARMGenerator gopter.Gen

// ServerPropertiesStatusARMGenerator returns a generator of ServerProperties_StatusARM instances for property testing.
// We first initialize serverPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerPropertiesStatusARMGenerator() gopter.Gen {
	if serverPropertiesStatusARMGenerator != nil {
		return serverPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesStatusARM(generators)
	serverPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForServerPropertiesStatusARM(generators)
	serverPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_StatusARM{}), generators)

	return serverPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForServerPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["EarliestRestoreDate"] = gen.PtrOf(gen.AlphaString())
	gens["FullyQualifiedDomainName"] = gen.PtrOf(gen.AlphaString())
	gens["MasterServerId"] = gen.PtrOf(gen.AlphaString())
	gens["MinimalTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicaCapacity"] = gen.PtrOf(gen.Int())
	gens["ReplicationRole"] = gen.PtrOf(gen.AlphaString())
	gens["SslEnforcement"] = gen.PtrOf(gen.AlphaString())
	gens["UserVisibleState"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServerPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["PrivateEndpointConnections"] = gen.SliceOf(ServerPrivateEndpointConnectionStatusARMGenerator())
	gens["StorageProfile"] = gen.PtrOf(StorageProfileStatusARMGenerator())
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
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Size"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServerPrivateEndpointConnection_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerPrivateEndpointConnection_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPrivateEndpointConnectionStatusARM, ServerPrivateEndpointConnectionStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPrivateEndpointConnectionStatusARM runs a test to see if a specific instance of ServerPrivateEndpointConnection_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPrivateEndpointConnectionStatusARM(subject ServerPrivateEndpointConnection_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerPrivateEndpointConnection_StatusARM
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

// Generator of ServerPrivateEndpointConnection_StatusARM instances for property testing - lazily instantiated by
// ServerPrivateEndpointConnectionStatusARMGenerator()
var serverPrivateEndpointConnectionStatusARMGenerator gopter.Gen

// ServerPrivateEndpointConnectionStatusARMGenerator returns a generator of ServerPrivateEndpointConnection_StatusARM instances for property testing.
// We first initialize serverPrivateEndpointConnectionStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerPrivateEndpointConnectionStatusARMGenerator() gopter.Gen {
	if serverPrivateEndpointConnectionStatusARMGenerator != nil {
		return serverPrivateEndpointConnectionStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionStatusARM(generators)
	serverPrivateEndpointConnectionStatusARMGenerator = gen.Struct(reflect.TypeOf(ServerPrivateEndpointConnection_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionStatusARM(generators)
	AddRelatedPropertyGeneratorsForServerPrivateEndpointConnectionStatusARM(generators)
	serverPrivateEndpointConnectionStatusARMGenerator = gen.Struct(reflect.TypeOf(ServerPrivateEndpointConnection_StatusARM{}), generators)

	return serverPrivateEndpointConnectionStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServerPrivateEndpointConnectionStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerPrivateEndpointConnectionStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerPrivateEndpointConnectionPropertiesStatusARMGenerator())
}

func Test_StorageProfile_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageProfile_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageProfileStatusARM, StorageProfileStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageProfileStatusARM runs a test to see if a specific instance of StorageProfile_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageProfileStatusARM(subject StorageProfile_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageProfile_StatusARM
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

// Generator of StorageProfile_StatusARM instances for property testing - lazily instantiated by
// StorageProfileStatusARMGenerator()
var storageProfileStatusARMGenerator gopter.Gen

// StorageProfileStatusARMGenerator returns a generator of StorageProfile_StatusARM instances for property testing.
func StorageProfileStatusARMGenerator() gopter.Gen {
	if storageProfileStatusARMGenerator != nil {
		return storageProfileStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageProfileStatusARM(generators)
	storageProfileStatusARMGenerator = gen.Struct(reflect.TypeOf(StorageProfile_StatusARM{}), generators)

	return storageProfileStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageProfileStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageProfileStatusARM(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAutogrow"] = gen.PtrOf(gen.AlphaString())
	gens["StorageMB"] = gen.PtrOf(gen.Int())
}

func Test_ServerPrivateEndpointConnectionProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerPrivateEndpointConnectionProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPrivateEndpointConnectionPropertiesStatusARM, ServerPrivateEndpointConnectionPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPrivateEndpointConnectionPropertiesStatusARM runs a test to see if a specific instance of ServerPrivateEndpointConnectionProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPrivateEndpointConnectionPropertiesStatusARM(subject ServerPrivateEndpointConnectionProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerPrivateEndpointConnectionProperties_StatusARM
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

// Generator of ServerPrivateEndpointConnectionProperties_StatusARM instances for property testing - lazily instantiated
// by ServerPrivateEndpointConnectionPropertiesStatusARMGenerator()
var serverPrivateEndpointConnectionPropertiesStatusARMGenerator gopter.Gen

// ServerPrivateEndpointConnectionPropertiesStatusARMGenerator returns a generator of ServerPrivateEndpointConnectionProperties_StatusARM instances for property testing.
// We first initialize serverPrivateEndpointConnectionPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerPrivateEndpointConnectionPropertiesStatusARMGenerator() gopter.Gen {
	if serverPrivateEndpointConnectionPropertiesStatusARMGenerator != nil {
		return serverPrivateEndpointConnectionPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionPropertiesStatusARM(generators)
	serverPrivateEndpointConnectionPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ServerPrivateEndpointConnectionProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForServerPrivateEndpointConnectionPropertiesStatusARM(generators)
	serverPrivateEndpointConnectionPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ServerPrivateEndpointConnectionProperties_StatusARM{}), generators)

	return serverPrivateEndpointConnectionPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPrivateEndpointConnectionPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServerPrivateEndpointConnectionPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerPrivateEndpointConnectionPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["PrivateEndpoint"] = gen.PtrOf(PrivateEndpointPropertyStatusARMGenerator())
	gens["PrivateLinkServiceConnectionState"] = gen.PtrOf(ServerPrivateLinkServiceConnectionStatePropertyStatusARMGenerator())
}

func Test_PrivateEndpointProperty_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointProperty_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointPropertyStatusARM, PrivateEndpointPropertyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointPropertyStatusARM runs a test to see if a specific instance of PrivateEndpointProperty_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointPropertyStatusARM(subject PrivateEndpointProperty_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointProperty_StatusARM
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

// Generator of PrivateEndpointProperty_StatusARM instances for property testing - lazily instantiated by
// PrivateEndpointPropertyStatusARMGenerator()
var privateEndpointPropertyStatusARMGenerator gopter.Gen

// PrivateEndpointPropertyStatusARMGenerator returns a generator of PrivateEndpointProperty_StatusARM instances for property testing.
func PrivateEndpointPropertyStatusARMGenerator() gopter.Gen {
	if privateEndpointPropertyStatusARMGenerator != nil {
		return privateEndpointPropertyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointPropertyStatusARM(generators)
	privateEndpointPropertyStatusARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointProperty_StatusARM{}), generators)

	return privateEndpointPropertyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointPropertyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointPropertyStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServerPrivateLinkServiceConnectionStateProperty_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerPrivateLinkServiceConnectionStateProperty_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPrivateLinkServiceConnectionStatePropertyStatusARM, ServerPrivateLinkServiceConnectionStatePropertyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPrivateLinkServiceConnectionStatePropertyStatusARM runs a test to see if a specific instance of ServerPrivateLinkServiceConnectionStateProperty_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPrivateLinkServiceConnectionStatePropertyStatusARM(subject ServerPrivateLinkServiceConnectionStateProperty_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerPrivateLinkServiceConnectionStateProperty_StatusARM
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

// Generator of ServerPrivateLinkServiceConnectionStateProperty_StatusARM instances for property testing - lazily
// instantiated by ServerPrivateLinkServiceConnectionStatePropertyStatusARMGenerator()
var serverPrivateLinkServiceConnectionStatePropertyStatusARMGenerator gopter.Gen

// ServerPrivateLinkServiceConnectionStatePropertyStatusARMGenerator returns a generator of ServerPrivateLinkServiceConnectionStateProperty_StatusARM instances for property testing.
func ServerPrivateLinkServiceConnectionStatePropertyStatusARMGenerator() gopter.Gen {
	if serverPrivateLinkServiceConnectionStatePropertyStatusARMGenerator != nil {
		return serverPrivateLinkServiceConnectionStatePropertyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPrivateLinkServiceConnectionStatePropertyStatusARM(generators)
	serverPrivateLinkServiceConnectionStatePropertyStatusARMGenerator = gen.Struct(reflect.TypeOf(ServerPrivateLinkServiceConnectionStateProperty_StatusARM{}), generators)

	return serverPrivateLinkServiceConnectionStatePropertyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForServerPrivateLinkServiceConnectionStatePropertyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPrivateLinkServiceConnectionStatePropertyStatusARM(gens map[string]gopter.Gen) {
	gens["ActionsRequired"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
}
