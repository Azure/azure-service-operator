// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180601

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

func Test_Factory_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Factory_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactory_Spec_ARM, Factory_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactory_Spec_ARM runs a test to see if a specific instance of Factory_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactory_Spec_ARM(subject Factory_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Factory_Spec_ARM
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

// Generator of Factory_Spec_ARM instances for property testing - lazily instantiated by Factory_Spec_ARMGenerator()
var factory_Spec_ARMGenerator gopter.Gen

// Factory_Spec_ARMGenerator returns a generator of Factory_Spec_ARM instances for property testing.
// We first initialize factory_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Factory_Spec_ARMGenerator() gopter.Gen {
	if factory_Spec_ARMGenerator != nil {
		return factory_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactory_Spec_ARM(generators)
	factory_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Factory_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactory_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForFactory_Spec_ARM(generators)
	factory_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Factory_Spec_ARM{}), generators)

	return factory_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFactory_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFactory_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFactory_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFactory_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(FactoryIdentity_ARMGenerator())
	gens["Properties"] = gen.PtrOf(FactoryProperties_ARMGenerator())
}

func Test_FactoryIdentity_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FactoryIdentity_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactoryIdentity_ARM, FactoryIdentity_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactoryIdentity_ARM runs a test to see if a specific instance of FactoryIdentity_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactoryIdentity_ARM(subject FactoryIdentity_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FactoryIdentity_ARM
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

// Generator of FactoryIdentity_ARM instances for property testing - lazily instantiated by
// FactoryIdentity_ARMGenerator()
var factoryIdentity_ARMGenerator gopter.Gen

// FactoryIdentity_ARMGenerator returns a generator of FactoryIdentity_ARM instances for property testing.
// We first initialize factoryIdentity_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FactoryIdentity_ARMGenerator() gopter.Gen {
	if factoryIdentity_ARMGenerator != nil {
		return factoryIdentity_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryIdentity_ARM(generators)
	factoryIdentity_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryIdentity_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryIdentity_ARM(generators)
	AddRelatedPropertyGeneratorsForFactoryIdentity_ARM(generators)
	factoryIdentity_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryIdentity_ARM{}), generators)

	return factoryIdentity_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFactoryIdentity_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFactoryIdentity_ARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(FactoryIdentity_Type_SystemAssigned, FactoryIdentity_Type_SystemAssignedUserAssigned, FactoryIdentity_Type_UserAssigned))
}

// AddRelatedPropertyGeneratorsForFactoryIdentity_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFactoryIdentity_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		UserAssignedIdentityDetails_ARMGenerator())
}

func Test_FactoryProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FactoryProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactoryProperties_ARM, FactoryProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactoryProperties_ARM runs a test to see if a specific instance of FactoryProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactoryProperties_ARM(subject FactoryProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FactoryProperties_ARM
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

// Generator of FactoryProperties_ARM instances for property testing - lazily instantiated by
// FactoryProperties_ARMGenerator()
var factoryProperties_ARMGenerator gopter.Gen

// FactoryProperties_ARMGenerator returns a generator of FactoryProperties_ARM instances for property testing.
// We first initialize factoryProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FactoryProperties_ARMGenerator() gopter.Gen {
	if factoryProperties_ARMGenerator != nil {
		return factoryProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryProperties_ARM(generators)
	factoryProperties_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForFactoryProperties_ARM(generators)
	factoryProperties_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryProperties_ARM{}), generators)

	return factoryProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFactoryProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFactoryProperties_ARM(gens map[string]gopter.Gen) {
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(FactoryProperties_PublicNetworkAccess_Disabled, FactoryProperties_PublicNetworkAccess_Enabled))
}

// AddRelatedPropertyGeneratorsForFactoryProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFactoryProperties_ARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionConfiguration_ARMGenerator())
	gens["GlobalParameters"] = gen.MapOf(
		gen.AlphaString(),
		GlobalParameterSpecification_ARMGenerator())
	gens["PurviewConfiguration"] = gen.PtrOf(PurviewConfiguration_ARMGenerator())
	gens["RepoConfiguration"] = gen.PtrOf(FactoryRepoConfiguration_ARMGenerator())
}

func Test_EncryptionConfiguration_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionConfiguration_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionConfiguration_ARM, EncryptionConfiguration_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionConfiguration_ARM runs a test to see if a specific instance of EncryptionConfiguration_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionConfiguration_ARM(subject EncryptionConfiguration_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionConfiguration_ARM
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

// Generator of EncryptionConfiguration_ARM instances for property testing - lazily instantiated by
// EncryptionConfiguration_ARMGenerator()
var encryptionConfiguration_ARMGenerator gopter.Gen

// EncryptionConfiguration_ARMGenerator returns a generator of EncryptionConfiguration_ARM instances for property testing.
// We first initialize encryptionConfiguration_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionConfiguration_ARMGenerator() gopter.Gen {
	if encryptionConfiguration_ARMGenerator != nil {
		return encryptionConfiguration_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionConfiguration_ARM(generators)
	encryptionConfiguration_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionConfiguration_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionConfiguration_ARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionConfiguration_ARM(generators)
	encryptionConfiguration_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionConfiguration_ARM{}), generators)

	return encryptionConfiguration_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionConfiguration_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionConfiguration_ARM(gens map[string]gopter.Gen) {
	gens["KeyName"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVersion"] = gen.PtrOf(gen.AlphaString())
	gens["VaultBaseUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEncryptionConfiguration_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionConfiguration_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(CMKIdentityDefinition_ARMGenerator())
}

func Test_FactoryRepoConfiguration_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FactoryRepoConfiguration_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactoryRepoConfiguration_ARM, FactoryRepoConfiguration_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactoryRepoConfiguration_ARM runs a test to see if a specific instance of FactoryRepoConfiguration_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactoryRepoConfiguration_ARM(subject FactoryRepoConfiguration_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FactoryRepoConfiguration_ARM
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

// Generator of FactoryRepoConfiguration_ARM instances for property testing - lazily instantiated by
// FactoryRepoConfiguration_ARMGenerator()
var factoryRepoConfiguration_ARMGenerator gopter.Gen

// FactoryRepoConfiguration_ARMGenerator returns a generator of FactoryRepoConfiguration_ARM instances for property testing.
func FactoryRepoConfiguration_ARMGenerator() gopter.Gen {
	if factoryRepoConfiguration_ARMGenerator != nil {
		return factoryRepoConfiguration_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFactoryRepoConfiguration_ARM(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(FactoryRepoConfiguration_ARM{}), map[string]gopter.Gen{propName: propGen}))
	}
	factoryRepoConfiguration_ARMGenerator = gen.OneGenOf(gens...)

	return factoryRepoConfiguration_ARMGenerator
}

// AddRelatedPropertyGeneratorsForFactoryRepoConfiguration_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFactoryRepoConfiguration_ARM(gens map[string]gopter.Gen) {
	gens["FactoryGitHub"] = FactoryGitHubConfiguration_ARMGenerator().Map(func(it FactoryGitHubConfiguration_ARM) *FactoryGitHubConfiguration_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["FactoryVSTS"] = FactoryVSTSConfiguration_ARMGenerator().Map(func(it FactoryVSTSConfiguration_ARM) *FactoryVSTSConfiguration_ARM {
		return &it
	}) // generate one case for OneOf type
}

func Test_GlobalParameterSpecification_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of GlobalParameterSpecification_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForGlobalParameterSpecification_ARM, GlobalParameterSpecification_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForGlobalParameterSpecification_ARM runs a test to see if a specific instance of GlobalParameterSpecification_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForGlobalParameterSpecification_ARM(subject GlobalParameterSpecification_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual GlobalParameterSpecification_ARM
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

// Generator of GlobalParameterSpecification_ARM instances for property testing - lazily instantiated by
// GlobalParameterSpecification_ARMGenerator()
var globalParameterSpecification_ARMGenerator gopter.Gen

// GlobalParameterSpecification_ARMGenerator returns a generator of GlobalParameterSpecification_ARM instances for property testing.
func GlobalParameterSpecification_ARMGenerator() gopter.Gen {
	if globalParameterSpecification_ARMGenerator != nil {
		return globalParameterSpecification_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForGlobalParameterSpecification_ARM(generators)
	globalParameterSpecification_ARMGenerator = gen.Struct(reflect.TypeOf(GlobalParameterSpecification_ARM{}), generators)

	return globalParameterSpecification_ARMGenerator
}

// AddIndependentPropertyGeneratorsForGlobalParameterSpecification_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForGlobalParameterSpecification_ARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		GlobalParameterSpecification_Type_Array,
		GlobalParameterSpecification_Type_Bool,
		GlobalParameterSpecification_Type_Float,
		GlobalParameterSpecification_Type_Int,
		GlobalParameterSpecification_Type_Object,
		GlobalParameterSpecification_Type_String))
}

func Test_PurviewConfiguration_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PurviewConfiguration_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPurviewConfiguration_ARM, PurviewConfiguration_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPurviewConfiguration_ARM runs a test to see if a specific instance of PurviewConfiguration_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPurviewConfiguration_ARM(subject PurviewConfiguration_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PurviewConfiguration_ARM
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

// Generator of PurviewConfiguration_ARM instances for property testing - lazily instantiated by
// PurviewConfiguration_ARMGenerator()
var purviewConfiguration_ARMGenerator gopter.Gen

// PurviewConfiguration_ARMGenerator returns a generator of PurviewConfiguration_ARM instances for property testing.
func PurviewConfiguration_ARMGenerator() gopter.Gen {
	if purviewConfiguration_ARMGenerator != nil {
		return purviewConfiguration_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPurviewConfiguration_ARM(generators)
	purviewConfiguration_ARMGenerator = gen.Struct(reflect.TypeOf(PurviewConfiguration_ARM{}), generators)

	return purviewConfiguration_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPurviewConfiguration_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPurviewConfiguration_ARM(gens map[string]gopter.Gen) {
	gens["PurviewResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentityDetails_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityDetails_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityDetails_ARM, UserAssignedIdentityDetails_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityDetails_ARM runs a test to see if a specific instance of UserAssignedIdentityDetails_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityDetails_ARM(subject UserAssignedIdentityDetails_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityDetails_ARM
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

// Generator of UserAssignedIdentityDetails_ARM instances for property testing - lazily instantiated by
// UserAssignedIdentityDetails_ARMGenerator()
var userAssignedIdentityDetails_ARMGenerator gopter.Gen

// UserAssignedIdentityDetails_ARMGenerator returns a generator of UserAssignedIdentityDetails_ARM instances for property testing.
func UserAssignedIdentityDetails_ARMGenerator() gopter.Gen {
	if userAssignedIdentityDetails_ARMGenerator != nil {
		return userAssignedIdentityDetails_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	userAssignedIdentityDetails_ARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityDetails_ARM{}), generators)

	return userAssignedIdentityDetails_ARMGenerator
}

func Test_CMKIdentityDefinition_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CMKIdentityDefinition_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCMKIdentityDefinition_ARM, CMKIdentityDefinition_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCMKIdentityDefinition_ARM runs a test to see if a specific instance of CMKIdentityDefinition_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCMKIdentityDefinition_ARM(subject CMKIdentityDefinition_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CMKIdentityDefinition_ARM
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

// Generator of CMKIdentityDefinition_ARM instances for property testing - lazily instantiated by
// CMKIdentityDefinition_ARMGenerator()
var cmkIdentityDefinition_ARMGenerator gopter.Gen

// CMKIdentityDefinition_ARMGenerator returns a generator of CMKIdentityDefinition_ARM instances for property testing.
func CMKIdentityDefinition_ARMGenerator() gopter.Gen {
	if cmkIdentityDefinition_ARMGenerator != nil {
		return cmkIdentityDefinition_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCMKIdentityDefinition_ARM(generators)
	cmkIdentityDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(CMKIdentityDefinition_ARM{}), generators)

	return cmkIdentityDefinition_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCMKIdentityDefinition_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCMKIdentityDefinition_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}

func Test_FactoryGitHubConfiguration_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FactoryGitHubConfiguration_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactoryGitHubConfiguration_ARM, FactoryGitHubConfiguration_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactoryGitHubConfiguration_ARM runs a test to see if a specific instance of FactoryGitHubConfiguration_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactoryGitHubConfiguration_ARM(subject FactoryGitHubConfiguration_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FactoryGitHubConfiguration_ARM
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

// Generator of FactoryGitHubConfiguration_ARM instances for property testing - lazily instantiated by
// FactoryGitHubConfiguration_ARMGenerator()
var factoryGitHubConfiguration_ARMGenerator gopter.Gen

// FactoryGitHubConfiguration_ARMGenerator returns a generator of FactoryGitHubConfiguration_ARM instances for property testing.
// We first initialize factoryGitHubConfiguration_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FactoryGitHubConfiguration_ARMGenerator() gopter.Gen {
	if factoryGitHubConfiguration_ARMGenerator != nil {
		return factoryGitHubConfiguration_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryGitHubConfiguration_ARM(generators)
	factoryGitHubConfiguration_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryGitHubConfiguration_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryGitHubConfiguration_ARM(generators)
	AddRelatedPropertyGeneratorsForFactoryGitHubConfiguration_ARM(generators)
	factoryGitHubConfiguration_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryGitHubConfiguration_ARM{}), generators)

	return factoryGitHubConfiguration_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFactoryGitHubConfiguration_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFactoryGitHubConfiguration_ARM(gens map[string]gopter.Gen) {
	gens["AccountName"] = gen.PtrOf(gen.AlphaString())
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["CollaborationBranch"] = gen.PtrOf(gen.AlphaString())
	gens["DisablePublish"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["LastCommitId"] = gen.PtrOf(gen.AlphaString())
	gens["RepositoryName"] = gen.PtrOf(gen.AlphaString())
	gens["RootFolder"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(FactoryGitHubConfiguration_Type_FactoryGitHubConfiguration)
}

// AddRelatedPropertyGeneratorsForFactoryGitHubConfiguration_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFactoryGitHubConfiguration_ARM(gens map[string]gopter.Gen) {
	gens["ClientSecret"] = gen.PtrOf(GitHubClientSecret_ARMGenerator())
}

func Test_FactoryVSTSConfiguration_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FactoryVSTSConfiguration_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactoryVSTSConfiguration_ARM, FactoryVSTSConfiguration_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactoryVSTSConfiguration_ARM runs a test to see if a specific instance of FactoryVSTSConfiguration_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactoryVSTSConfiguration_ARM(subject FactoryVSTSConfiguration_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FactoryVSTSConfiguration_ARM
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

// Generator of FactoryVSTSConfiguration_ARM instances for property testing - lazily instantiated by
// FactoryVSTSConfiguration_ARMGenerator()
var factoryVSTSConfiguration_ARMGenerator gopter.Gen

// FactoryVSTSConfiguration_ARMGenerator returns a generator of FactoryVSTSConfiguration_ARM instances for property testing.
func FactoryVSTSConfiguration_ARMGenerator() gopter.Gen {
	if factoryVSTSConfiguration_ARMGenerator != nil {
		return factoryVSTSConfiguration_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryVSTSConfiguration_ARM(generators)
	factoryVSTSConfiguration_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryVSTSConfiguration_ARM{}), generators)

	return factoryVSTSConfiguration_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFactoryVSTSConfiguration_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFactoryVSTSConfiguration_ARM(gens map[string]gopter.Gen) {
	gens["AccountName"] = gen.PtrOf(gen.AlphaString())
	gens["CollaborationBranch"] = gen.PtrOf(gen.AlphaString())
	gens["DisablePublish"] = gen.PtrOf(gen.Bool())
	gens["LastCommitId"] = gen.PtrOf(gen.AlphaString())
	gens["ProjectName"] = gen.PtrOf(gen.AlphaString())
	gens["RepositoryName"] = gen.PtrOf(gen.AlphaString())
	gens["RootFolder"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(FactoryVSTSConfiguration_Type_FactoryVSTSConfiguration)
}

func Test_GitHubClientSecret_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of GitHubClientSecret_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForGitHubClientSecret_ARM, GitHubClientSecret_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForGitHubClientSecret_ARM runs a test to see if a specific instance of GitHubClientSecret_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForGitHubClientSecret_ARM(subject GitHubClientSecret_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual GitHubClientSecret_ARM
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

// Generator of GitHubClientSecret_ARM instances for property testing - lazily instantiated by
// GitHubClientSecret_ARMGenerator()
var gitHubClientSecret_ARMGenerator gopter.Gen

// GitHubClientSecret_ARMGenerator returns a generator of GitHubClientSecret_ARM instances for property testing.
func GitHubClientSecret_ARMGenerator() gopter.Gen {
	if gitHubClientSecret_ARMGenerator != nil {
		return gitHubClientSecret_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForGitHubClientSecret_ARM(generators)
	gitHubClientSecret_ARMGenerator = gen.Struct(reflect.TypeOf(GitHubClientSecret_ARM{}), generators)

	return gitHubClientSecret_ARMGenerator
}

// AddIndependentPropertyGeneratorsForGitHubClientSecret_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForGitHubClientSecret_ARM(gens map[string]gopter.Gen) {
	gens["ByoaSecretAkvUrl"] = gen.PtrOf(gen.AlphaString())
	gens["ByoaSecretName"] = gen.PtrOf(gen.AlphaString())
}
