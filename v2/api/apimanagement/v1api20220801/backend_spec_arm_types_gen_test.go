// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

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

func Test_BackendAuthorizationHeaderCredentials_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendAuthorizationHeaderCredentials_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendAuthorizationHeaderCredentials_ARM, BackendAuthorizationHeaderCredentials_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendAuthorizationHeaderCredentials_ARM runs a test to see if a specific instance of BackendAuthorizationHeaderCredentials_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendAuthorizationHeaderCredentials_ARM(subject BackendAuthorizationHeaderCredentials_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendAuthorizationHeaderCredentials_ARM
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

// Generator of BackendAuthorizationHeaderCredentials_ARM instances for property testing - lazily instantiated by
// BackendAuthorizationHeaderCredentials_ARMGenerator()
var backendAuthorizationHeaderCredentials_ARMGenerator gopter.Gen

// BackendAuthorizationHeaderCredentials_ARMGenerator returns a generator of BackendAuthorizationHeaderCredentials_ARM instances for property testing.
func BackendAuthorizationHeaderCredentials_ARMGenerator() gopter.Gen {
	if backendAuthorizationHeaderCredentials_ARMGenerator != nil {
		return backendAuthorizationHeaderCredentials_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendAuthorizationHeaderCredentials_ARM(generators)
	backendAuthorizationHeaderCredentials_ARMGenerator = gen.Struct(reflect.TypeOf(BackendAuthorizationHeaderCredentials_ARM{}), generators)

	return backendAuthorizationHeaderCredentials_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBackendAuthorizationHeaderCredentials_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendAuthorizationHeaderCredentials_ARM(gens map[string]gopter.Gen) {
	gens["Parameter"] = gen.PtrOf(gen.AlphaString())
	gens["Scheme"] = gen.PtrOf(gen.AlphaString())
}

func Test_BackendContractProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendContractProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendContractProperties_ARM, BackendContractProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendContractProperties_ARM runs a test to see if a specific instance of BackendContractProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendContractProperties_ARM(subject BackendContractProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendContractProperties_ARM
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

// Generator of BackendContractProperties_ARM instances for property testing - lazily instantiated by
// BackendContractProperties_ARMGenerator()
var backendContractProperties_ARMGenerator gopter.Gen

// BackendContractProperties_ARMGenerator returns a generator of BackendContractProperties_ARM instances for property testing.
// We first initialize backendContractProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BackendContractProperties_ARMGenerator() gopter.Gen {
	if backendContractProperties_ARMGenerator != nil {
		return backendContractProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendContractProperties_ARM(generators)
	backendContractProperties_ARMGenerator = gen.Struct(reflect.TypeOf(BackendContractProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendContractProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForBackendContractProperties_ARM(generators)
	backendContractProperties_ARMGenerator = gen.Struct(reflect.TypeOf(BackendContractProperties_ARM{}), generators)

	return backendContractProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBackendContractProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendContractProperties_ARM(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(BackendContractProperties_Protocol_ARM_Http, BackendContractProperties_Protocol_ARM_Soap))
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Title"] = gen.PtrOf(gen.AlphaString())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBackendContractProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackendContractProperties_ARM(gens map[string]gopter.Gen) {
	gens["Credentials"] = gen.PtrOf(BackendCredentialsContract_ARMGenerator())
	gens["Properties"] = gen.PtrOf(BackendProperties_ARMGenerator())
	gens["Proxy"] = gen.PtrOf(BackendProxyContract_ARMGenerator())
	gens["Tls"] = gen.PtrOf(BackendTlsProperties_ARMGenerator())
}

func Test_BackendCredentialsContract_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendCredentialsContract_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendCredentialsContract_ARM, BackendCredentialsContract_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendCredentialsContract_ARM runs a test to see if a specific instance of BackendCredentialsContract_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendCredentialsContract_ARM(subject BackendCredentialsContract_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendCredentialsContract_ARM
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

// Generator of BackendCredentialsContract_ARM instances for property testing - lazily instantiated by
// BackendCredentialsContract_ARMGenerator()
var backendCredentialsContract_ARMGenerator gopter.Gen

// BackendCredentialsContract_ARMGenerator returns a generator of BackendCredentialsContract_ARM instances for property testing.
// We first initialize backendCredentialsContract_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BackendCredentialsContract_ARMGenerator() gopter.Gen {
	if backendCredentialsContract_ARMGenerator != nil {
		return backendCredentialsContract_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendCredentialsContract_ARM(generators)
	backendCredentialsContract_ARMGenerator = gen.Struct(reflect.TypeOf(BackendCredentialsContract_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendCredentialsContract_ARM(generators)
	AddRelatedPropertyGeneratorsForBackendCredentialsContract_ARM(generators)
	backendCredentialsContract_ARMGenerator = gen.Struct(reflect.TypeOf(BackendCredentialsContract_ARM{}), generators)

	return backendCredentialsContract_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBackendCredentialsContract_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendCredentialsContract_ARM(gens map[string]gopter.Gen) {
	gens["Certificate"] = gen.SliceOf(gen.AlphaString())
	gens["CertificateIds"] = gen.SliceOf(gen.AlphaString())
	gens["Header"] = gen.MapOf(
		gen.AlphaString(),
		gen.SliceOf(gen.AlphaString()))
	gens["Query"] = gen.MapOf(
		gen.AlphaString(),
		gen.SliceOf(gen.AlphaString()))
}

// AddRelatedPropertyGeneratorsForBackendCredentialsContract_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackendCredentialsContract_ARM(gens map[string]gopter.Gen) {
	gens["Authorization"] = gen.PtrOf(BackendAuthorizationHeaderCredentials_ARMGenerator())
}

func Test_BackendProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendProperties_ARM, BackendProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendProperties_ARM runs a test to see if a specific instance of BackendProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendProperties_ARM(subject BackendProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendProperties_ARM
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

// Generator of BackendProperties_ARM instances for property testing - lazily instantiated by
// BackendProperties_ARMGenerator()
var backendProperties_ARMGenerator gopter.Gen

// BackendProperties_ARMGenerator returns a generator of BackendProperties_ARM instances for property testing.
func BackendProperties_ARMGenerator() gopter.Gen {
	if backendProperties_ARMGenerator != nil {
		return backendProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForBackendProperties_ARM(generators)
	backendProperties_ARMGenerator = gen.Struct(reflect.TypeOf(BackendProperties_ARM{}), generators)

	return backendProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForBackendProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackendProperties_ARM(gens map[string]gopter.Gen) {
	gens["ServiceFabricCluster"] = gen.PtrOf(BackendServiceFabricClusterProperties_ARMGenerator())
}

func Test_BackendProxyContract_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendProxyContract_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendProxyContract_ARM, BackendProxyContract_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendProxyContract_ARM runs a test to see if a specific instance of BackendProxyContract_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendProxyContract_ARM(subject BackendProxyContract_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendProxyContract_ARM
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

// Generator of BackendProxyContract_ARM instances for property testing - lazily instantiated by
// BackendProxyContract_ARMGenerator()
var backendProxyContract_ARMGenerator gopter.Gen

// BackendProxyContract_ARMGenerator returns a generator of BackendProxyContract_ARM instances for property testing.
func BackendProxyContract_ARMGenerator() gopter.Gen {
	if backendProxyContract_ARMGenerator != nil {
		return backendProxyContract_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendProxyContract_ARM(generators)
	backendProxyContract_ARMGenerator = gen.Struct(reflect.TypeOf(BackendProxyContract_ARM{}), generators)

	return backendProxyContract_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBackendProxyContract_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendProxyContract_ARM(gens map[string]gopter.Gen) {
	gens["Password"] = gen.PtrOf(gen.AlphaString())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
	gens["Username"] = gen.PtrOf(gen.AlphaString())
}

func Test_BackendServiceFabricClusterProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendServiceFabricClusterProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendServiceFabricClusterProperties_ARM, BackendServiceFabricClusterProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendServiceFabricClusterProperties_ARM runs a test to see if a specific instance of BackendServiceFabricClusterProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendServiceFabricClusterProperties_ARM(subject BackendServiceFabricClusterProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendServiceFabricClusterProperties_ARM
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

// Generator of BackendServiceFabricClusterProperties_ARM instances for property testing - lazily instantiated by
// BackendServiceFabricClusterProperties_ARMGenerator()
var backendServiceFabricClusterProperties_ARMGenerator gopter.Gen

// BackendServiceFabricClusterProperties_ARMGenerator returns a generator of BackendServiceFabricClusterProperties_ARM instances for property testing.
// We first initialize backendServiceFabricClusterProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BackendServiceFabricClusterProperties_ARMGenerator() gopter.Gen {
	if backendServiceFabricClusterProperties_ARMGenerator != nil {
		return backendServiceFabricClusterProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendServiceFabricClusterProperties_ARM(generators)
	backendServiceFabricClusterProperties_ARMGenerator = gen.Struct(reflect.TypeOf(BackendServiceFabricClusterProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendServiceFabricClusterProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForBackendServiceFabricClusterProperties_ARM(generators)
	backendServiceFabricClusterProperties_ARMGenerator = gen.Struct(reflect.TypeOf(BackendServiceFabricClusterProperties_ARM{}), generators)

	return backendServiceFabricClusterProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBackendServiceFabricClusterProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendServiceFabricClusterProperties_ARM(gens map[string]gopter.Gen) {
	gens["ClientCertificateId"] = gen.PtrOf(gen.AlphaString())
	gens["ClientCertificatethumbprint"] = gen.PtrOf(gen.AlphaString())
	gens["ManagementEndpoints"] = gen.SliceOf(gen.AlphaString())
	gens["MaxPartitionResolutionRetries"] = gen.PtrOf(gen.Int())
	gens["ServerCertificateThumbprints"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBackendServiceFabricClusterProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackendServiceFabricClusterProperties_ARM(gens map[string]gopter.Gen) {
	gens["ServerX509Names"] = gen.SliceOf(X509CertificateName_ARMGenerator())
}

func Test_BackendTlsProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendTlsProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendTlsProperties_ARM, BackendTlsProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendTlsProperties_ARM runs a test to see if a specific instance of BackendTlsProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendTlsProperties_ARM(subject BackendTlsProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendTlsProperties_ARM
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

// Generator of BackendTlsProperties_ARM instances for property testing - lazily instantiated by
// BackendTlsProperties_ARMGenerator()
var backendTlsProperties_ARMGenerator gopter.Gen

// BackendTlsProperties_ARMGenerator returns a generator of BackendTlsProperties_ARM instances for property testing.
func BackendTlsProperties_ARMGenerator() gopter.Gen {
	if backendTlsProperties_ARMGenerator != nil {
		return backendTlsProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendTlsProperties_ARM(generators)
	backendTlsProperties_ARMGenerator = gen.Struct(reflect.TypeOf(BackendTlsProperties_ARM{}), generators)

	return backendTlsProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBackendTlsProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendTlsProperties_ARM(gens map[string]gopter.Gen) {
	gens["ValidateCertificateChain"] = gen.PtrOf(gen.Bool())
	gens["ValidateCertificateName"] = gen.PtrOf(gen.Bool())
}

func Test_Backend_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backend_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackend_Spec_ARM, Backend_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackend_Spec_ARM runs a test to see if a specific instance of Backend_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackend_Spec_ARM(subject Backend_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backend_Spec_ARM
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

// Generator of Backend_Spec_ARM instances for property testing - lazily instantiated by Backend_Spec_ARMGenerator()
var backend_Spec_ARMGenerator gopter.Gen

// Backend_Spec_ARMGenerator returns a generator of Backend_Spec_ARM instances for property testing.
// We first initialize backend_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Backend_Spec_ARMGenerator() gopter.Gen {
	if backend_Spec_ARMGenerator != nil {
		return backend_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackend_Spec_ARM(generators)
	backend_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Backend_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackend_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForBackend_Spec_ARM(generators)
	backend_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Backend_Spec_ARM{}), generators)

	return backend_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBackend_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackend_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForBackend_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackend_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(BackendContractProperties_ARMGenerator())
}

func Test_X509CertificateName_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of X509CertificateName_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForX509CertificateName_ARM, X509CertificateName_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForX509CertificateName_ARM runs a test to see if a specific instance of X509CertificateName_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForX509CertificateName_ARM(subject X509CertificateName_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual X509CertificateName_ARM
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

// Generator of X509CertificateName_ARM instances for property testing - lazily instantiated by
// X509CertificateName_ARMGenerator()
var x509CertificateName_ARMGenerator gopter.Gen

// X509CertificateName_ARMGenerator returns a generator of X509CertificateName_ARM instances for property testing.
func X509CertificateName_ARMGenerator() gopter.Gen {
	if x509CertificateName_ARMGenerator != nil {
		return x509CertificateName_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForX509CertificateName_ARM(generators)
	x509CertificateName_ARMGenerator = gen.Struct(reflect.TypeOf(X509CertificateName_ARM{}), generators)

	return x509CertificateName_ARMGenerator
}

// AddIndependentPropertyGeneratorsForX509CertificateName_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForX509CertificateName_ARM(gens map[string]gopter.Gen) {
	gens["IssuerCertificateThumbprint"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}
