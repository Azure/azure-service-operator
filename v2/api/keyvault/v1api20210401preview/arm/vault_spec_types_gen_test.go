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

func Test_AccessPolicyEntry_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AccessPolicyEntry via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAccessPolicyEntry, AccessPolicyEntryGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAccessPolicyEntry runs a test to see if a specific instance of AccessPolicyEntry round trips to JSON and back losslessly
func RunJSONSerializationTestForAccessPolicyEntry(subject AccessPolicyEntry) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AccessPolicyEntry
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

// Generator of AccessPolicyEntry instances for property testing - lazily instantiated by AccessPolicyEntryGenerator()
var accessPolicyEntryGenerator gopter.Gen

// AccessPolicyEntryGenerator returns a generator of AccessPolicyEntry instances for property testing.
// We first initialize accessPolicyEntryGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AccessPolicyEntryGenerator() gopter.Gen {
	if accessPolicyEntryGenerator != nil {
		return accessPolicyEntryGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAccessPolicyEntry(generators)
	accessPolicyEntryGenerator = gen.Struct(reflect.TypeOf(AccessPolicyEntry{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAccessPolicyEntry(generators)
	AddRelatedPropertyGeneratorsForAccessPolicyEntry(generators)
	accessPolicyEntryGenerator = gen.Struct(reflect.TypeOf(AccessPolicyEntry{}), generators)

	return accessPolicyEntryGenerator
}

// AddIndependentPropertyGeneratorsForAccessPolicyEntry is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAccessPolicyEntry(gens map[string]gopter.Gen) {
	gens["ApplicationId"] = gen.PtrOf(gen.AlphaString())
	gens["ObjectId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAccessPolicyEntry is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAccessPolicyEntry(gens map[string]gopter.Gen) {
	gens["Permissions"] = gen.PtrOf(PermissionsGenerator())
}

func Test_IPRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IPRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIPRule, IPRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIPRule runs a test to see if a specific instance of IPRule round trips to JSON and back losslessly
func RunJSONSerializationTestForIPRule(subject IPRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IPRule
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

// Generator of IPRule instances for property testing - lazily instantiated by IPRuleGenerator()
var ipRuleGenerator gopter.Gen

// IPRuleGenerator returns a generator of IPRule instances for property testing.
func IPRuleGenerator() gopter.Gen {
	if ipRuleGenerator != nil {
		return ipRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPRule(generators)
	ipRuleGenerator = gen.Struct(reflect.TypeOf(IPRule{}), generators)

	return ipRuleGenerator
}

// AddIndependentPropertyGeneratorsForIPRule is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIPRule(gens map[string]gopter.Gen) {
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkRuleSet_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkRuleSet via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkRuleSet, NetworkRuleSetGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkRuleSet runs a test to see if a specific instance of NetworkRuleSet round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkRuleSet(subject NetworkRuleSet) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkRuleSet
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

// Generator of NetworkRuleSet instances for property testing - lazily instantiated by NetworkRuleSetGenerator()
var networkRuleSetGenerator gopter.Gen

// NetworkRuleSetGenerator returns a generator of NetworkRuleSet instances for property testing.
// We first initialize networkRuleSetGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkRuleSetGenerator() gopter.Gen {
	if networkRuleSetGenerator != nil {
		return networkRuleSetGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkRuleSet(generators)
	networkRuleSetGenerator = gen.Struct(reflect.TypeOf(NetworkRuleSet{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkRuleSet(generators)
	AddRelatedPropertyGeneratorsForNetworkRuleSet(generators)
	networkRuleSetGenerator = gen.Struct(reflect.TypeOf(NetworkRuleSet{}), generators)

	return networkRuleSetGenerator
}

// AddIndependentPropertyGeneratorsForNetworkRuleSet is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkRuleSet(gens map[string]gopter.Gen) {
	gens["Bypass"] = gen.PtrOf(gen.OneConstOf(NetworkRuleSet_Bypass_AzureServices, NetworkRuleSet_Bypass_None))
	gens["DefaultAction"] = gen.PtrOf(gen.OneConstOf(NetworkRuleSet_DefaultAction_Allow, NetworkRuleSet_DefaultAction_Deny))
}

// AddRelatedPropertyGeneratorsForNetworkRuleSet is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkRuleSet(gens map[string]gopter.Gen) {
	gens["IpRules"] = gen.SliceOf(IPRuleGenerator())
	gens["VirtualNetworkRules"] = gen.SliceOf(VirtualNetworkRuleGenerator())
}

func Test_Permissions_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Permissions via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPermissions, PermissionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPermissions runs a test to see if a specific instance of Permissions round trips to JSON and back losslessly
func RunJSONSerializationTestForPermissions(subject Permissions) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Permissions
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

// Generator of Permissions instances for property testing - lazily instantiated by PermissionsGenerator()
var permissionsGenerator gopter.Gen

// PermissionsGenerator returns a generator of Permissions instances for property testing.
func PermissionsGenerator() gopter.Gen {
	if permissionsGenerator != nil {
		return permissionsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPermissions(generators)
	permissionsGenerator = gen.Struct(reflect.TypeOf(Permissions{}), generators)

	return permissionsGenerator
}

// AddIndependentPropertyGeneratorsForPermissions is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPermissions(gens map[string]gopter.Gen) {
	gens["Certificates"] = gen.SliceOf(gen.OneConstOf(
		Permissions_Certificates_Backup,
		Permissions_Certificates_Create,
		Permissions_Certificates_Delete,
		Permissions_Certificates_Deleteissuers,
		Permissions_Certificates_Get,
		Permissions_Certificates_Getissuers,
		Permissions_Certificates_Import,
		Permissions_Certificates_List,
		Permissions_Certificates_Listissuers,
		Permissions_Certificates_Managecontacts,
		Permissions_Certificates_Manageissuers,
		Permissions_Certificates_Purge,
		Permissions_Certificates_Recover,
		Permissions_Certificates_Restore,
		Permissions_Certificates_Setissuers,
		Permissions_Certificates_Update))
	gens["Keys"] = gen.SliceOf(gen.OneConstOf(
		Permissions_Keys_Backup,
		Permissions_Keys_Create,
		Permissions_Keys_Decrypt,
		Permissions_Keys_Delete,
		Permissions_Keys_Encrypt,
		Permissions_Keys_Get,
		Permissions_Keys_Import,
		Permissions_Keys_List,
		Permissions_Keys_Purge,
		Permissions_Keys_Recover,
		Permissions_Keys_Release,
		Permissions_Keys_Restore,
		Permissions_Keys_Sign,
		Permissions_Keys_UnwrapKey,
		Permissions_Keys_Update,
		Permissions_Keys_Verify,
		Permissions_Keys_WrapKey))
	gens["Secrets"] = gen.SliceOf(gen.OneConstOf(
		Permissions_Secrets_Backup,
		Permissions_Secrets_Delete,
		Permissions_Secrets_Get,
		Permissions_Secrets_List,
		Permissions_Secrets_Purge,
		Permissions_Secrets_Recover,
		Permissions_Secrets_Restore,
		Permissions_Secrets_Set))
	gens["Storage"] = gen.SliceOf(gen.OneConstOf(
		Permissions_Storage_Backup,
		Permissions_Storage_Delete,
		Permissions_Storage_Deletesas,
		Permissions_Storage_Get,
		Permissions_Storage_Getsas,
		Permissions_Storage_List,
		Permissions_Storage_Listsas,
		Permissions_Storage_Purge,
		Permissions_Storage_Recover,
		Permissions_Storage_Regeneratekey,
		Permissions_Storage_Restore,
		Permissions_Storage_Set,
		Permissions_Storage_Setsas,
		Permissions_Storage_Update))
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
	gens["Family"] = gen.PtrOf(gen.OneConstOf(Sku_Family_A))
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_Premium, Sku_Name_Standard))
}

func Test_VaultProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VaultProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVaultProperties, VaultPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVaultProperties runs a test to see if a specific instance of VaultProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForVaultProperties(subject VaultProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VaultProperties
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

// Generator of VaultProperties instances for property testing - lazily instantiated by VaultPropertiesGenerator()
var vaultPropertiesGenerator gopter.Gen

// VaultPropertiesGenerator returns a generator of VaultProperties instances for property testing.
// We first initialize vaultPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VaultPropertiesGenerator() gopter.Gen {
	if vaultPropertiesGenerator != nil {
		return vaultPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVaultProperties(generators)
	vaultPropertiesGenerator = gen.Struct(reflect.TypeOf(VaultProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVaultProperties(generators)
	AddRelatedPropertyGeneratorsForVaultProperties(generators)
	vaultPropertiesGenerator = gen.Struct(reflect.TypeOf(VaultProperties{}), generators)

	return vaultPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForVaultProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVaultProperties(gens map[string]gopter.Gen) {
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(
		VaultProperties_CreateMode_CreateOrRecover,
		VaultProperties_CreateMode_Default,
		VaultProperties_CreateMode_PurgeThenCreate,
		VaultProperties_CreateMode_Recover))
	gens["EnablePurgeProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableRbacAuthorization"] = gen.PtrOf(gen.Bool())
	gens["EnableSoftDelete"] = gen.PtrOf(gen.Bool())
	gens["EnabledForDeployment"] = gen.PtrOf(gen.Bool())
	gens["EnabledForDiskEncryption"] = gen.PtrOf(gen.Bool())
	gens["EnabledForTemplateDeployment"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(VaultProperties_ProvisioningState_RegisteringDns, VaultProperties_ProvisioningState_Succeeded))
	gens["SoftDeleteRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["VaultUri"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVaultProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVaultProperties(gens map[string]gopter.Gen) {
	gens["AccessPolicies"] = gen.SliceOf(AccessPolicyEntryGenerator())
	gens["NetworkAcls"] = gen.PtrOf(NetworkRuleSetGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_Vault_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Vault_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVault_Spec, Vault_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVault_Spec runs a test to see if a specific instance of Vault_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVault_Spec(subject Vault_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Vault_Spec
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

// Generator of Vault_Spec instances for property testing - lazily instantiated by Vault_SpecGenerator()
var vault_SpecGenerator gopter.Gen

// Vault_SpecGenerator returns a generator of Vault_Spec instances for property testing.
// We first initialize vault_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Vault_SpecGenerator() gopter.Gen {
	if vault_SpecGenerator != nil {
		return vault_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVault_Spec(generators)
	vault_SpecGenerator = gen.Struct(reflect.TypeOf(Vault_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVault_Spec(generators)
	AddRelatedPropertyGeneratorsForVault_Spec(generators)
	vault_SpecGenerator = gen.Struct(reflect.TypeOf(Vault_Spec{}), generators)

	return vault_SpecGenerator
}

// AddIndependentPropertyGeneratorsForVault_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVault_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVault_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVault_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VaultPropertiesGenerator())
}

func Test_VirtualNetworkRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkRule, VirtualNetworkRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkRule runs a test to see if a specific instance of VirtualNetworkRule round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkRule(subject VirtualNetworkRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkRule
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

// Generator of VirtualNetworkRule instances for property testing - lazily instantiated by VirtualNetworkRuleGenerator()
var virtualNetworkRuleGenerator gopter.Gen

// VirtualNetworkRuleGenerator returns a generator of VirtualNetworkRule instances for property testing.
func VirtualNetworkRuleGenerator() gopter.Gen {
	if virtualNetworkRuleGenerator != nil {
		return virtualNetworkRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkRule(generators)
	virtualNetworkRuleGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkRule{}), generators)

	return virtualNetworkRuleGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkRule is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkRule(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IgnoreMissingVnetServiceEndpoint"] = gen.PtrOf(gen.Bool())
}
