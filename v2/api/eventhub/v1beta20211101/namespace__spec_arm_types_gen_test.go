// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

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

func Test_Namespace_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespace__spec_arm_types_gen_test.go
		"Round trip of Namespace_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespace_SpecARM, Namespace_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespace_SpecARM runs a test to see if a specific instance of Namespace_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespace_SpecARM(subject Namespace_SpecARM) string {
========
		"Round trip of Namespaces_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_SpecARM, Namespaces_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_SpecARM runs a test to see if a specific instance of Namespaces_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_SpecARM(subject Namespaces_SpecARM) string {
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_spec_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespace_SpecARM
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

<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespace__spec_arm_types_gen_test.go
// Generator of Namespace_SpecARM instances for property testing - lazily instantiated by Namespace_SpecARMGenerator()
var namespace_SpecARMGenerator gopter.Gen

// Namespace_SpecARMGenerator returns a generator of Namespace_SpecARM instances for property testing.
// We first initialize namespace_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespace_SpecARMGenerator() gopter.Gen {
	if namespace_SpecARMGenerator != nil {
		return namespace_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_SpecARM(generators)
	namespace_SpecARMGenerator = gen.Struct(reflect.TypeOf(Namespace_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_SpecARM(generators)
	AddRelatedPropertyGeneratorsForNamespace_SpecARM(generators)
	namespace_SpecARMGenerator = gen.Struct(reflect.TypeOf(Namespace_SpecARM{}), generators)

	return namespace_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespace_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespace_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
========
// Generator of Namespaces_SpecARM instances for property testing - lazily instantiated by Namespaces_SpecARMGenerator()
var namespaces_SpecARMGenerator gopter.Gen

// Namespaces_SpecARMGenerator returns a generator of Namespaces_SpecARM instances for property testing.
// We first initialize namespaces_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_SpecARMGenerator() gopter.Gen {
	if namespaces_SpecARMGenerator != nil {
		return namespaces_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_SpecARM(generators)
	namespaces_SpecARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_SpecARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_SpecARM(generators)
	namespaces_SpecARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_SpecARM{}), generators)

	return namespaces_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_SpecARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_spec_arm_types_gen_test.go
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespace__spec_arm_types_gen_test.go
// AddRelatedPropertyGeneratorsForNamespace_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespace_SpecARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityARMGenerator())
	gens["Properties"] = gen.PtrOf(Namespace_Spec_PropertiesARMGenerator())
========
// AddRelatedPropertyGeneratorsForNamespaces_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_SpecARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityARMGenerator())
	gens["Properties"] = gen.PtrOf(Namespaces_Spec_PropertiesARMGenerator())
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_spec_arm_types_gen_test.go
	gens["Sku"] = gen.PtrOf(SkuARMGenerator())
}

func Test_IdentityARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IdentityARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityARM, IdentityARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityARM runs a test to see if a specific instance of IdentityARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityARM(subject IdentityARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IdentityARM
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

// Generator of IdentityARM instances for property testing - lazily instantiated by IdentityARMGenerator()
var identityARMGenerator gopter.Gen

// IdentityARMGenerator returns a generator of IdentityARM instances for property testing.
func IdentityARMGenerator() gopter.Gen {
	if identityARMGenerator != nil {
		return identityARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityARM(generators)
	identityARMGenerator = gen.Struct(reflect.TypeOf(IdentityARM{}), generators)

	return identityARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentityARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		Identity_Type_None,
		Identity_Type_SystemAssigned,
		Identity_Type_SystemAssignedUserAssigned,
		Identity_Type_UserAssigned))
}

func Test_Namespace_Spec_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespace__spec_arm_types_gen_test.go
		"Round trip of Namespace_Spec_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespace_Spec_PropertiesARM, Namespace_Spec_PropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespace_Spec_PropertiesARM runs a test to see if a specific instance of Namespace_Spec_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespace_Spec_PropertiesARM(subject Namespace_Spec_PropertiesARM) string {
========
		"Round trip of Namespaces_Spec_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Spec_PropertiesARM, Namespaces_Spec_PropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Spec_PropertiesARM runs a test to see if a specific instance of Namespaces_Spec_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Spec_PropertiesARM(subject Namespaces_Spec_PropertiesARM) string {
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_spec_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespace_Spec_PropertiesARM
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

<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespace__spec_arm_types_gen_test.go
// Generator of Namespace_Spec_PropertiesARM instances for property testing - lazily instantiated by
// Namespace_Spec_PropertiesARMGenerator()
var namespace_Spec_PropertiesARMGenerator gopter.Gen

// Namespace_Spec_PropertiesARMGenerator returns a generator of Namespace_Spec_PropertiesARM instances for property testing.
// We first initialize namespace_Spec_PropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespace_Spec_PropertiesARMGenerator() gopter.Gen {
	if namespace_Spec_PropertiesARMGenerator != nil {
		return namespace_Spec_PropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_Spec_PropertiesARM(generators)
	namespace_Spec_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(Namespace_Spec_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_Spec_PropertiesARM(generators)
	AddRelatedPropertyGeneratorsForNamespace_Spec_PropertiesARM(generators)
	namespace_Spec_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(Namespace_Spec_PropertiesARM{}), generators)

	return namespace_Spec_PropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespace_Spec_PropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespace_Spec_PropertiesARM(gens map[string]gopter.Gen) {
========
// Generator of Namespaces_Spec_PropertiesARM instances for property testing - lazily instantiated by
// Namespaces_Spec_PropertiesARMGenerator()
var namespaces_Spec_PropertiesARMGenerator gopter.Gen

// Namespaces_Spec_PropertiesARMGenerator returns a generator of Namespaces_Spec_PropertiesARM instances for property testing.
// We first initialize namespaces_Spec_PropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Spec_PropertiesARMGenerator() gopter.Gen {
	if namespaces_Spec_PropertiesARMGenerator != nil {
		return namespaces_Spec_PropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Spec_PropertiesARM(generators)
	namespaces_Spec_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Spec_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Spec_PropertiesARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Spec_PropertiesARM(generators)
	namespaces_Spec_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Spec_PropertiesARM{}), generators)

	return namespaces_Spec_PropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Spec_PropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Spec_PropertiesARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_spec_arm_types_gen_test.go
	gens["AlternateName"] = gen.PtrOf(gen.AlphaString())
	gens["ClusterArmId"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["IsAutoInflateEnabled"] = gen.PtrOf(gen.Bool())
	gens["KafkaEnabled"] = gen.PtrOf(gen.Bool())
	gens["MaximumThroughputUnits"] = gen.PtrOf(gen.Int())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespace__spec_arm_types_gen_test.go
// AddRelatedPropertyGeneratorsForNamespace_Spec_PropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespace_Spec_PropertiesARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionARMGenerator())
========
// AddRelatedPropertyGeneratorsForNamespaces_Spec_PropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Spec_PropertiesARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(Namespaces_Spec_Properties_PrivateEndpointConnectionsARMGenerator())
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_spec_arm_types_gen_test.go
}

func Test_SkuARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SkuARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuARM, SkuARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuARM runs a test to see if a specific instance of SkuARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuARM(subject SkuARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SkuARM
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

// Generator of SkuARM instances for property testing - lazily instantiated by SkuARMGenerator()
var skuARMGenerator gopter.Gen

// SkuARMGenerator returns a generator of SkuARM instances for property testing.
func SkuARMGenerator() gopter.Gen {
	if skuARMGenerator != nil {
		return skuARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuARM(generators)
	skuARMGenerator = gen.Struct(reflect.TypeOf(SkuARM{}), generators)

	return skuARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_Basic, Sku_Name_Premium, Sku_Name_Standard))
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Sku_Tier_Basic, Sku_Tier_Premium, Sku_Tier_Standard))
}

func Test_EncryptionARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionARM, EncryptionARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionARM runs a test to see if a specific instance of EncryptionARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionARM(subject EncryptionARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionARM
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

// Generator of EncryptionARM instances for property testing - lazily instantiated by EncryptionARMGenerator()
var encryptionARMGenerator gopter.Gen

// EncryptionARMGenerator returns a generator of EncryptionARM instances for property testing.
// We first initialize encryptionARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionARMGenerator() gopter.Gen {
	if encryptionARMGenerator != nil {
		return encryptionARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionARM(generators)
	encryptionARMGenerator = gen.Struct(reflect.TypeOf(EncryptionARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionARM(generators)
	encryptionARMGenerator = gen.Struct(reflect.TypeOf(EncryptionARM{}), generators)

	return encryptionARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionARM(gens map[string]gopter.Gen) {
	gens["KeySource"] = gen.PtrOf(gen.OneConstOf(Encryption_KeySource_MicrosoftKeyVault))
	gens["RequireInfrastructureEncryption"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForEncryptionARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.SliceOf(KeyVaultPropertiesARMGenerator())
}

func Test_PrivateEndpointConnectionARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespace__spec_arm_types_gen_test.go
		"Round trip of PrivateEndpointConnectionARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionARM, PrivateEndpointConnectionARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionARM runs a test to see if a specific instance of PrivateEndpointConnectionARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionARM(subject PrivateEndpointConnectionARM) string {
========
		"Round trip of Namespaces_Spec_Properties_PrivateEndpointConnectionsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Spec_Properties_PrivateEndpointConnectionsARM, Namespaces_Spec_Properties_PrivateEndpointConnectionsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Spec_Properties_PrivateEndpointConnectionsARM runs a test to see if a specific instance of Namespaces_Spec_Properties_PrivateEndpointConnectionsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Spec_Properties_PrivateEndpointConnectionsARM(subject Namespaces_Spec_Properties_PrivateEndpointConnectionsARM) string {
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_spec_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnectionARM
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

<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespace__spec_arm_types_gen_test.go
// Generator of PrivateEndpointConnectionARM instances for property testing - lazily instantiated by
// PrivateEndpointConnectionARMGenerator()
var privateEndpointConnectionARMGenerator gopter.Gen

// PrivateEndpointConnectionARMGenerator returns a generator of PrivateEndpointConnectionARM instances for property testing.
func PrivateEndpointConnectionARMGenerator() gopter.Gen {
	if privateEndpointConnectionARMGenerator != nil {
		return privateEndpointConnectionARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateEndpointConnectionARM(generators)
	privateEndpointConnectionARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnectionARM{}), generators)

	return privateEndpointConnectionARMGenerator
}

// AddRelatedPropertyGeneratorsForPrivateEndpointConnectionARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointConnectionARM(gens map[string]gopter.Gen) {
========
// Generator of Namespaces_Spec_Properties_PrivateEndpointConnectionsARM instances for property testing - lazily
// instantiated by Namespaces_Spec_Properties_PrivateEndpointConnectionsARMGenerator()
var namespaces_Spec_Properties_PrivateEndpointConnectionsARMGenerator gopter.Gen

// Namespaces_Spec_Properties_PrivateEndpointConnectionsARMGenerator returns a generator of Namespaces_Spec_Properties_PrivateEndpointConnectionsARM instances for property testing.
func Namespaces_Spec_Properties_PrivateEndpointConnectionsARMGenerator() gopter.Gen {
	if namespaces_Spec_Properties_PrivateEndpointConnectionsARMGenerator != nil {
		return namespaces_Spec_Properties_PrivateEndpointConnectionsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespaces_Spec_Properties_PrivateEndpointConnectionsARM(generators)
	namespaces_Spec_Properties_PrivateEndpointConnectionsARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Spec_Properties_PrivateEndpointConnectionsARM{}), generators)

	return namespaces_Spec_Properties_PrivateEndpointConnectionsARMGenerator
}

// AddRelatedPropertyGeneratorsForNamespaces_Spec_Properties_PrivateEndpointConnectionsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Spec_Properties_PrivateEndpointConnectionsARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_spec_arm_types_gen_test.go
	gens["Properties"] = gen.PtrOf(PrivateEndpointConnectionPropertiesARMGenerator())
}

func Test_KeyVaultPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultPropertiesARM, KeyVaultPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultPropertiesARM runs a test to see if a specific instance of KeyVaultPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultPropertiesARM(subject KeyVaultPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultPropertiesARM
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

// Generator of KeyVaultPropertiesARM instances for property testing - lazily instantiated by
// KeyVaultPropertiesARMGenerator()
var keyVaultPropertiesARMGenerator gopter.Gen

// KeyVaultPropertiesARMGenerator returns a generator of KeyVaultPropertiesARM instances for property testing.
// We first initialize keyVaultPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultPropertiesARMGenerator() gopter.Gen {
	if keyVaultPropertiesARMGenerator != nil {
		return keyVaultPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM(generators)
	keyVaultPropertiesARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultPropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForKeyVaultPropertiesARM(generators)
	keyVaultPropertiesARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultPropertiesARM{}), generators)

	return keyVaultPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM(gens map[string]gopter.Gen) {
	gens["KeyName"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultUri"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultPropertiesARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(UserAssignedIdentityPropertiesARMGenerator())
}

func Test_PrivateEndpointConnectionPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnectionPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionPropertiesARM, PrivateEndpointConnectionPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionPropertiesARM runs a test to see if a specific instance of PrivateEndpointConnectionPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionPropertiesARM(subject PrivateEndpointConnectionPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnectionPropertiesARM
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

// Generator of PrivateEndpointConnectionPropertiesARM instances for property testing - lazily instantiated by
// PrivateEndpointConnectionPropertiesARMGenerator()
var privateEndpointConnectionPropertiesARMGenerator gopter.Gen

// PrivateEndpointConnectionPropertiesARMGenerator returns a generator of PrivateEndpointConnectionPropertiesARM instances for property testing.
func PrivateEndpointConnectionPropertiesARMGenerator() gopter.Gen {
	if privateEndpointConnectionPropertiesARMGenerator != nil {
		return privateEndpointConnectionPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateEndpointConnectionPropertiesARM(generators)
	privateEndpointConnectionPropertiesARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnectionPropertiesARM{}), generators)

	return privateEndpointConnectionPropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForPrivateEndpointConnectionPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointConnectionPropertiesARM(gens map[string]gopter.Gen) {
	gens["PrivateEndpoint"] = gen.PtrOf(PrivateEndpointARMGenerator())
}

func Test_PrivateEndpointARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointARM, PrivateEndpointARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointARM runs a test to see if a specific instance of PrivateEndpointARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointARM(subject PrivateEndpointARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointARM
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

// Generator of PrivateEndpointARM instances for property testing - lazily instantiated by PrivateEndpointARMGenerator()
var privateEndpointARMGenerator gopter.Gen

// PrivateEndpointARMGenerator returns a generator of PrivateEndpointARM instances for property testing.
func PrivateEndpointARMGenerator() gopter.Gen {
	if privateEndpointARMGenerator != nil {
		return privateEndpointARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointARM(generators)
	privateEndpointARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointARM{}), generators)

	return privateEndpointARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentityPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityPropertiesARM, UserAssignedIdentityPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityPropertiesARM runs a test to see if a specific instance of UserAssignedIdentityPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityPropertiesARM(subject UserAssignedIdentityPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityPropertiesARM
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

// Generator of UserAssignedIdentityPropertiesARM instances for property testing - lazily instantiated by
// UserAssignedIdentityPropertiesARMGenerator()
var userAssignedIdentityPropertiesARMGenerator gopter.Gen

// UserAssignedIdentityPropertiesARMGenerator returns a generator of UserAssignedIdentityPropertiesARM instances for property testing.
func UserAssignedIdentityPropertiesARMGenerator() gopter.Gen {
	if userAssignedIdentityPropertiesARMGenerator != nil {
		return userAssignedIdentityPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesARM(generators)
	userAssignedIdentityPropertiesARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityPropertiesARM{}), generators)

	return userAssignedIdentityPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}
