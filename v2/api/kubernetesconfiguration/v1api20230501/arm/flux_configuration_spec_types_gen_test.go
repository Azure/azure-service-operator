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

func Test_AzureBlobDefinition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AzureBlobDefinition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAzureBlobDefinition, AzureBlobDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAzureBlobDefinition runs a test to see if a specific instance of AzureBlobDefinition round trips to JSON and back losslessly
func RunJSONSerializationTestForAzureBlobDefinition(subject AzureBlobDefinition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AzureBlobDefinition
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

// Generator of AzureBlobDefinition instances for property testing - lazily instantiated by
// AzureBlobDefinitionGenerator()
var azureBlobDefinitionGenerator gopter.Gen

// AzureBlobDefinitionGenerator returns a generator of AzureBlobDefinition instances for property testing.
// We first initialize azureBlobDefinitionGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AzureBlobDefinitionGenerator() gopter.Gen {
	if azureBlobDefinitionGenerator != nil {
		return azureBlobDefinitionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureBlobDefinition(generators)
	azureBlobDefinitionGenerator = gen.Struct(reflect.TypeOf(AzureBlobDefinition{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureBlobDefinition(generators)
	AddRelatedPropertyGeneratorsForAzureBlobDefinition(generators)
	azureBlobDefinitionGenerator = gen.Struct(reflect.TypeOf(AzureBlobDefinition{}), generators)

	return azureBlobDefinitionGenerator
}

// AddIndependentPropertyGeneratorsForAzureBlobDefinition is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAzureBlobDefinition(gens map[string]gopter.Gen) {
	gens["AccountKey"] = gen.PtrOf(gen.AlphaString())
	gens["ContainerName"] = gen.PtrOf(gen.AlphaString())
	gens["LocalAuthRef"] = gen.PtrOf(gen.AlphaString())
	gens["SasToken"] = gen.PtrOf(gen.AlphaString())
	gens["SyncIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["TimeoutInSeconds"] = gen.PtrOf(gen.Int())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAzureBlobDefinition is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAzureBlobDefinition(gens map[string]gopter.Gen) {
	gens["ManagedIdentity"] = gen.PtrOf(ManagedIdentityDefinitionGenerator())
	gens["ServicePrincipal"] = gen.PtrOf(ServicePrincipalDefinitionGenerator())
}

func Test_BucketDefinition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BucketDefinition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBucketDefinition, BucketDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBucketDefinition runs a test to see if a specific instance of BucketDefinition round trips to JSON and back losslessly
func RunJSONSerializationTestForBucketDefinition(subject BucketDefinition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BucketDefinition
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

// Generator of BucketDefinition instances for property testing - lazily instantiated by BucketDefinitionGenerator()
var bucketDefinitionGenerator gopter.Gen

// BucketDefinitionGenerator returns a generator of BucketDefinition instances for property testing.
func BucketDefinitionGenerator() gopter.Gen {
	if bucketDefinitionGenerator != nil {
		return bucketDefinitionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBucketDefinition(generators)
	bucketDefinitionGenerator = gen.Struct(reflect.TypeOf(BucketDefinition{}), generators)

	return bucketDefinitionGenerator
}

// AddIndependentPropertyGeneratorsForBucketDefinition is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBucketDefinition(gens map[string]gopter.Gen) {
	gens["AccessKey"] = gen.PtrOf(gen.AlphaString())
	gens["BucketName"] = gen.PtrOf(gen.AlphaString())
	gens["Insecure"] = gen.PtrOf(gen.Bool())
	gens["LocalAuthRef"] = gen.PtrOf(gen.AlphaString())
	gens["SyncIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["TimeoutInSeconds"] = gen.PtrOf(gen.Int())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

func Test_FluxConfiguration_Properties_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FluxConfiguration_Properties_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFluxConfiguration_Properties_Spec, FluxConfiguration_Properties_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFluxConfiguration_Properties_Spec runs a test to see if a specific instance of FluxConfiguration_Properties_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFluxConfiguration_Properties_Spec(subject FluxConfiguration_Properties_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FluxConfiguration_Properties_Spec
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

// Generator of FluxConfiguration_Properties_Spec instances for property testing - lazily instantiated by
// FluxConfiguration_Properties_SpecGenerator()
var fluxConfiguration_Properties_SpecGenerator gopter.Gen

// FluxConfiguration_Properties_SpecGenerator returns a generator of FluxConfiguration_Properties_Spec instances for property testing.
// We first initialize fluxConfiguration_Properties_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FluxConfiguration_Properties_SpecGenerator() gopter.Gen {
	if fluxConfiguration_Properties_SpecGenerator != nil {
		return fluxConfiguration_Properties_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFluxConfiguration_Properties_Spec(generators)
	fluxConfiguration_Properties_SpecGenerator = gen.Struct(reflect.TypeOf(FluxConfiguration_Properties_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFluxConfiguration_Properties_Spec(generators)
	AddRelatedPropertyGeneratorsForFluxConfiguration_Properties_Spec(generators)
	fluxConfiguration_Properties_SpecGenerator = gen.Struct(reflect.TypeOf(FluxConfiguration_Properties_Spec{}), generators)

	return fluxConfiguration_Properties_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFluxConfiguration_Properties_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFluxConfiguration_Properties_Spec(gens map[string]gopter.Gen) {
	gens["ConfigurationProtectedSettings"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Namespace"] = gen.PtrOf(gen.AlphaString())
	gens["ReconciliationWaitDuration"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.OneConstOf(ScopeDefinition_Cluster, ScopeDefinition_Namespace))
	gens["SourceKind"] = gen.PtrOf(gen.OneConstOf(SourceKindDefinition_AzureBlob, SourceKindDefinition_Bucket, SourceKindDefinition_GitRepository))
	gens["Suspend"] = gen.PtrOf(gen.Bool())
	gens["WaitForReconciliation"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForFluxConfiguration_Properties_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFluxConfiguration_Properties_Spec(gens map[string]gopter.Gen) {
	gens["AzureBlob"] = gen.PtrOf(AzureBlobDefinitionGenerator())
	gens["Bucket"] = gen.PtrOf(BucketDefinitionGenerator())
	gens["GitRepository"] = gen.PtrOf(GitRepositoryDefinitionGenerator())
	gens["Kustomizations"] = gen.MapOf(
		gen.AlphaString(),
		KustomizationDefinitionGenerator())
}

func Test_FluxConfiguration_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FluxConfiguration_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFluxConfiguration_Spec, FluxConfiguration_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFluxConfiguration_Spec runs a test to see if a specific instance of FluxConfiguration_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFluxConfiguration_Spec(subject FluxConfiguration_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FluxConfiguration_Spec
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

// Generator of FluxConfiguration_Spec instances for property testing - lazily instantiated by
// FluxConfiguration_SpecGenerator()
var fluxConfiguration_SpecGenerator gopter.Gen

// FluxConfiguration_SpecGenerator returns a generator of FluxConfiguration_Spec instances for property testing.
// We first initialize fluxConfiguration_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FluxConfiguration_SpecGenerator() gopter.Gen {
	if fluxConfiguration_SpecGenerator != nil {
		return fluxConfiguration_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFluxConfiguration_Spec(generators)
	fluxConfiguration_SpecGenerator = gen.Struct(reflect.TypeOf(FluxConfiguration_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFluxConfiguration_Spec(generators)
	AddRelatedPropertyGeneratorsForFluxConfiguration_Spec(generators)
	fluxConfiguration_SpecGenerator = gen.Struct(reflect.TypeOf(FluxConfiguration_Spec{}), generators)

	return fluxConfiguration_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFluxConfiguration_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFluxConfiguration_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForFluxConfiguration_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFluxConfiguration_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FluxConfiguration_Properties_SpecGenerator())
}

func Test_GitRepositoryDefinition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of GitRepositoryDefinition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForGitRepositoryDefinition, GitRepositoryDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForGitRepositoryDefinition runs a test to see if a specific instance of GitRepositoryDefinition round trips to JSON and back losslessly
func RunJSONSerializationTestForGitRepositoryDefinition(subject GitRepositoryDefinition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual GitRepositoryDefinition
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

// Generator of GitRepositoryDefinition instances for property testing - lazily instantiated by
// GitRepositoryDefinitionGenerator()
var gitRepositoryDefinitionGenerator gopter.Gen

// GitRepositoryDefinitionGenerator returns a generator of GitRepositoryDefinition instances for property testing.
// We first initialize gitRepositoryDefinitionGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func GitRepositoryDefinitionGenerator() gopter.Gen {
	if gitRepositoryDefinitionGenerator != nil {
		return gitRepositoryDefinitionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForGitRepositoryDefinition(generators)
	gitRepositoryDefinitionGenerator = gen.Struct(reflect.TypeOf(GitRepositoryDefinition{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForGitRepositoryDefinition(generators)
	AddRelatedPropertyGeneratorsForGitRepositoryDefinition(generators)
	gitRepositoryDefinitionGenerator = gen.Struct(reflect.TypeOf(GitRepositoryDefinition{}), generators)

	return gitRepositoryDefinitionGenerator
}

// AddIndependentPropertyGeneratorsForGitRepositoryDefinition is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForGitRepositoryDefinition(gens map[string]gopter.Gen) {
	gens["HttpsCACert"] = gen.PtrOf(gen.AlphaString())
	gens["HttpsUser"] = gen.PtrOf(gen.AlphaString())
	gens["LocalAuthRef"] = gen.PtrOf(gen.AlphaString())
	gens["SshKnownHosts"] = gen.PtrOf(gen.AlphaString())
	gens["SyncIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["TimeoutInSeconds"] = gen.PtrOf(gen.Int())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForGitRepositoryDefinition is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForGitRepositoryDefinition(gens map[string]gopter.Gen) {
	gens["RepositoryRef"] = gen.PtrOf(RepositoryRefDefinitionGenerator())
}

func Test_KustomizationDefinition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KustomizationDefinition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKustomizationDefinition, KustomizationDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKustomizationDefinition runs a test to see if a specific instance of KustomizationDefinition round trips to JSON and back losslessly
func RunJSONSerializationTestForKustomizationDefinition(subject KustomizationDefinition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KustomizationDefinition
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

// Generator of KustomizationDefinition instances for property testing - lazily instantiated by
// KustomizationDefinitionGenerator()
var kustomizationDefinitionGenerator gopter.Gen

// KustomizationDefinitionGenerator returns a generator of KustomizationDefinition instances for property testing.
// We first initialize kustomizationDefinitionGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KustomizationDefinitionGenerator() gopter.Gen {
	if kustomizationDefinitionGenerator != nil {
		return kustomizationDefinitionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKustomizationDefinition(generators)
	kustomizationDefinitionGenerator = gen.Struct(reflect.TypeOf(KustomizationDefinition{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKustomizationDefinition(generators)
	AddRelatedPropertyGeneratorsForKustomizationDefinition(generators)
	kustomizationDefinitionGenerator = gen.Struct(reflect.TypeOf(KustomizationDefinition{}), generators)

	return kustomizationDefinitionGenerator
}

// AddIndependentPropertyGeneratorsForKustomizationDefinition is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKustomizationDefinition(gens map[string]gopter.Gen) {
	gens["DependsOn"] = gen.SliceOf(gen.AlphaString())
	gens["Force"] = gen.PtrOf(gen.Bool())
	gens["Path"] = gen.PtrOf(gen.AlphaString())
	gens["Prune"] = gen.PtrOf(gen.Bool())
	gens["RetryIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SyncIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["TimeoutInSeconds"] = gen.PtrOf(gen.Int())
	gens["Wait"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForKustomizationDefinition is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKustomizationDefinition(gens map[string]gopter.Gen) {
	gens["PostBuild"] = gen.PtrOf(PostBuildDefinitionGenerator())
}

func Test_ManagedIdentityDefinition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedIdentityDefinition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedIdentityDefinition, ManagedIdentityDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedIdentityDefinition runs a test to see if a specific instance of ManagedIdentityDefinition round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedIdentityDefinition(subject ManagedIdentityDefinition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedIdentityDefinition
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

// Generator of ManagedIdentityDefinition instances for property testing - lazily instantiated by
// ManagedIdentityDefinitionGenerator()
var managedIdentityDefinitionGenerator gopter.Gen

// ManagedIdentityDefinitionGenerator returns a generator of ManagedIdentityDefinition instances for property testing.
func ManagedIdentityDefinitionGenerator() gopter.Gen {
	if managedIdentityDefinitionGenerator != nil {
		return managedIdentityDefinitionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedIdentityDefinition(generators)
	managedIdentityDefinitionGenerator = gen.Struct(reflect.TypeOf(ManagedIdentityDefinition{}), generators)

	return managedIdentityDefinitionGenerator
}

// AddIndependentPropertyGeneratorsForManagedIdentityDefinition is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedIdentityDefinition(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
}

func Test_PostBuildDefinition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PostBuildDefinition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPostBuildDefinition, PostBuildDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPostBuildDefinition runs a test to see if a specific instance of PostBuildDefinition round trips to JSON and back losslessly
func RunJSONSerializationTestForPostBuildDefinition(subject PostBuildDefinition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PostBuildDefinition
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

// Generator of PostBuildDefinition instances for property testing - lazily instantiated by
// PostBuildDefinitionGenerator()
var postBuildDefinitionGenerator gopter.Gen

// PostBuildDefinitionGenerator returns a generator of PostBuildDefinition instances for property testing.
// We first initialize postBuildDefinitionGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PostBuildDefinitionGenerator() gopter.Gen {
	if postBuildDefinitionGenerator != nil {
		return postBuildDefinitionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPostBuildDefinition(generators)
	postBuildDefinitionGenerator = gen.Struct(reflect.TypeOf(PostBuildDefinition{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPostBuildDefinition(generators)
	AddRelatedPropertyGeneratorsForPostBuildDefinition(generators)
	postBuildDefinitionGenerator = gen.Struct(reflect.TypeOf(PostBuildDefinition{}), generators)

	return postBuildDefinitionGenerator
}

// AddIndependentPropertyGeneratorsForPostBuildDefinition is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPostBuildDefinition(gens map[string]gopter.Gen) {
	gens["Substitute"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPostBuildDefinition is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPostBuildDefinition(gens map[string]gopter.Gen) {
	gens["SubstituteFrom"] = gen.SliceOf(SubstituteFromDefinitionGenerator())
}

func Test_RepositoryRefDefinition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RepositoryRefDefinition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRepositoryRefDefinition, RepositoryRefDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRepositoryRefDefinition runs a test to see if a specific instance of RepositoryRefDefinition round trips to JSON and back losslessly
func RunJSONSerializationTestForRepositoryRefDefinition(subject RepositoryRefDefinition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RepositoryRefDefinition
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

// Generator of RepositoryRefDefinition instances for property testing - lazily instantiated by
// RepositoryRefDefinitionGenerator()
var repositoryRefDefinitionGenerator gopter.Gen

// RepositoryRefDefinitionGenerator returns a generator of RepositoryRefDefinition instances for property testing.
func RepositoryRefDefinitionGenerator() gopter.Gen {
	if repositoryRefDefinitionGenerator != nil {
		return repositoryRefDefinitionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRepositoryRefDefinition(generators)
	repositoryRefDefinitionGenerator = gen.Struct(reflect.TypeOf(RepositoryRefDefinition{}), generators)

	return repositoryRefDefinitionGenerator
}

// AddIndependentPropertyGeneratorsForRepositoryRefDefinition is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRepositoryRefDefinition(gens map[string]gopter.Gen) {
	gens["Branch"] = gen.PtrOf(gen.AlphaString())
	gens["Commit"] = gen.PtrOf(gen.AlphaString())
	gens["Semver"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServicePrincipalDefinition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServicePrincipalDefinition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServicePrincipalDefinition, ServicePrincipalDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServicePrincipalDefinition runs a test to see if a specific instance of ServicePrincipalDefinition round trips to JSON and back losslessly
func RunJSONSerializationTestForServicePrincipalDefinition(subject ServicePrincipalDefinition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServicePrincipalDefinition
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

// Generator of ServicePrincipalDefinition instances for property testing - lazily instantiated by
// ServicePrincipalDefinitionGenerator()
var servicePrincipalDefinitionGenerator gopter.Gen

// ServicePrincipalDefinitionGenerator returns a generator of ServicePrincipalDefinition instances for property testing.
func ServicePrincipalDefinitionGenerator() gopter.Gen {
	if servicePrincipalDefinitionGenerator != nil {
		return servicePrincipalDefinitionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServicePrincipalDefinition(generators)
	servicePrincipalDefinitionGenerator = gen.Struct(reflect.TypeOf(ServicePrincipalDefinition{}), generators)

	return servicePrincipalDefinitionGenerator
}

// AddIndependentPropertyGeneratorsForServicePrincipalDefinition is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServicePrincipalDefinition(gens map[string]gopter.Gen) {
	gens["ClientCertificate"] = gen.PtrOf(gen.AlphaString())
	gens["ClientCertificatePassword"] = gen.PtrOf(gen.AlphaString())
	gens["ClientCertificateSendChain"] = gen.PtrOf(gen.Bool())
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["ClientSecret"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_SubstituteFromDefinition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubstituteFromDefinition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubstituteFromDefinition, SubstituteFromDefinitionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubstituteFromDefinition runs a test to see if a specific instance of SubstituteFromDefinition round trips to JSON and back losslessly
func RunJSONSerializationTestForSubstituteFromDefinition(subject SubstituteFromDefinition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubstituteFromDefinition
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

// Generator of SubstituteFromDefinition instances for property testing - lazily instantiated by
// SubstituteFromDefinitionGenerator()
var substituteFromDefinitionGenerator gopter.Gen

// SubstituteFromDefinitionGenerator returns a generator of SubstituteFromDefinition instances for property testing.
func SubstituteFromDefinitionGenerator() gopter.Gen {
	if substituteFromDefinitionGenerator != nil {
		return substituteFromDefinitionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubstituteFromDefinition(generators)
	substituteFromDefinitionGenerator = gen.Struct(reflect.TypeOf(SubstituteFromDefinition{}), generators)

	return substituteFromDefinitionGenerator
}

// AddIndependentPropertyGeneratorsForSubstituteFromDefinition is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubstituteFromDefinition(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Optional"] = gen.PtrOf(gen.Bool())
}
