// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

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

func Test_AzureFirstPartyManagedCertificateParameters_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AzureFirstPartyManagedCertificateParameters_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAzureFirstPartyManagedCertificateParameters_ARM, AzureFirstPartyManagedCertificateParameters_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAzureFirstPartyManagedCertificateParameters_ARM runs a test to see if a specific instance of AzureFirstPartyManagedCertificateParameters_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAzureFirstPartyManagedCertificateParameters_ARM(subject AzureFirstPartyManagedCertificateParameters_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AzureFirstPartyManagedCertificateParameters_ARM
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

// Generator of AzureFirstPartyManagedCertificateParameters_ARM instances for property testing - lazily instantiated by
// AzureFirstPartyManagedCertificateParameters_ARMGenerator()
var azureFirstPartyManagedCertificateParameters_ARMGenerator gopter.Gen

// AzureFirstPartyManagedCertificateParameters_ARMGenerator returns a generator of AzureFirstPartyManagedCertificateParameters_ARM instances for property testing.
func AzureFirstPartyManagedCertificateParameters_ARMGenerator() gopter.Gen {
	if azureFirstPartyManagedCertificateParameters_ARMGenerator != nil {
		return azureFirstPartyManagedCertificateParameters_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters_ARM(generators)
	azureFirstPartyManagedCertificateParameters_ARMGenerator = gen.Struct(reflect.TypeOf(AzureFirstPartyManagedCertificateParameters_ARM{}), generators)

	return azureFirstPartyManagedCertificateParameters_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters_ARM(gens map[string]gopter.Gen) {
	gens["SubjectAlternativeNames"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(AzureFirstPartyManagedCertificateParameters_Type_ARM_AzureFirstPartyManagedCertificate)
}

func Test_CustomerCertificateParameters_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CustomerCertificateParameters_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCustomerCertificateParameters_ARM, CustomerCertificateParameters_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCustomerCertificateParameters_ARM runs a test to see if a specific instance of CustomerCertificateParameters_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCustomerCertificateParameters_ARM(subject CustomerCertificateParameters_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CustomerCertificateParameters_ARM
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

// Generator of CustomerCertificateParameters_ARM instances for property testing - lazily instantiated by
// CustomerCertificateParameters_ARMGenerator()
var customerCertificateParameters_ARMGenerator gopter.Gen

// CustomerCertificateParameters_ARMGenerator returns a generator of CustomerCertificateParameters_ARM instances for property testing.
// We first initialize customerCertificateParameters_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CustomerCertificateParameters_ARMGenerator() gopter.Gen {
	if customerCertificateParameters_ARMGenerator != nil {
		return customerCertificateParameters_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomerCertificateParameters_ARM(generators)
	customerCertificateParameters_ARMGenerator = gen.Struct(reflect.TypeOf(CustomerCertificateParameters_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomerCertificateParameters_ARM(generators)
	AddRelatedPropertyGeneratorsForCustomerCertificateParameters_ARM(generators)
	customerCertificateParameters_ARMGenerator = gen.Struct(reflect.TypeOf(CustomerCertificateParameters_ARM{}), generators)

	return customerCertificateParameters_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCustomerCertificateParameters_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCustomerCertificateParameters_ARM(gens map[string]gopter.Gen) {
	gens["SecretVersion"] = gen.PtrOf(gen.AlphaString())
	gens["SubjectAlternativeNames"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(CustomerCertificateParameters_Type_ARM_CustomerCertificate)
	gens["UseLatestVersion"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCustomerCertificateParameters_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCustomerCertificateParameters_ARM(gens map[string]gopter.Gen) {
	gens["SecretSource"] = gen.PtrOf(ResourceReference_ARMGenerator())
}

func Test_ManagedCertificateParameters_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedCertificateParameters_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedCertificateParameters_ARM, ManagedCertificateParameters_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedCertificateParameters_ARM runs a test to see if a specific instance of ManagedCertificateParameters_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedCertificateParameters_ARM(subject ManagedCertificateParameters_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedCertificateParameters_ARM
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

// Generator of ManagedCertificateParameters_ARM instances for property testing - lazily instantiated by
// ManagedCertificateParameters_ARMGenerator()
var managedCertificateParameters_ARMGenerator gopter.Gen

// ManagedCertificateParameters_ARMGenerator returns a generator of ManagedCertificateParameters_ARM instances for property testing.
func ManagedCertificateParameters_ARMGenerator() gopter.Gen {
	if managedCertificateParameters_ARMGenerator != nil {
		return managedCertificateParameters_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedCertificateParameters_ARM(generators)
	managedCertificateParameters_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedCertificateParameters_ARM{}), generators)

	return managedCertificateParameters_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedCertificateParameters_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedCertificateParameters_ARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.OneConstOf(ManagedCertificateParameters_Type_ARM_ManagedCertificate)
}

func Test_Profiles_Secret_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_Secret_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_Secret_Spec_ARM, Profiles_Secret_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_Secret_Spec_ARM runs a test to see if a specific instance of Profiles_Secret_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_Secret_Spec_ARM(subject Profiles_Secret_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_Secret_Spec_ARM
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

// Generator of Profiles_Secret_Spec_ARM instances for property testing - lazily instantiated by
// Profiles_Secret_Spec_ARMGenerator()
var profiles_Secret_Spec_ARMGenerator gopter.Gen

// Profiles_Secret_Spec_ARMGenerator returns a generator of Profiles_Secret_Spec_ARM instances for property testing.
// We first initialize profiles_Secret_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_Secret_Spec_ARMGenerator() gopter.Gen {
	if profiles_Secret_Spec_ARMGenerator != nil {
		return profiles_Secret_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_Secret_Spec_ARM(generators)
	profiles_Secret_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_Secret_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_Secret_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForProfiles_Secret_Spec_ARM(generators)
	profiles_Secret_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_Secret_Spec_ARM{}), generators)

	return profiles_Secret_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_Secret_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_Secret_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForProfiles_Secret_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_Secret_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SecretProperties_ARMGenerator())
}

func Test_SecretParameters_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecretParameters_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecretParameters_ARM, SecretParameters_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecretParameters_ARM runs a test to see if a specific instance of SecretParameters_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecretParameters_ARM(subject SecretParameters_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecretParameters_ARM
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

// Generator of SecretParameters_ARM instances for property testing - lazily instantiated by
// SecretParameters_ARMGenerator()
var secretParameters_ARMGenerator gopter.Gen

// SecretParameters_ARMGenerator returns a generator of SecretParameters_ARM instances for property testing.
func SecretParameters_ARMGenerator() gopter.Gen {
	if secretParameters_ARMGenerator != nil {
		return secretParameters_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecretParameters_ARM(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(SecretParameters_ARM{}), map[string]gopter.Gen{propName: propGen}))
	}
	secretParameters_ARMGenerator = gen.OneGenOf(gens...)

	return secretParameters_ARMGenerator
}

// AddRelatedPropertyGeneratorsForSecretParameters_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecretParameters_ARM(gens map[string]gopter.Gen) {
	gens["AzureFirstPartyManagedCertificate"] = AzureFirstPartyManagedCertificateParameters_ARMGenerator().Map(func(it AzureFirstPartyManagedCertificateParameters_ARM) *AzureFirstPartyManagedCertificateParameters_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["CustomerCertificate"] = CustomerCertificateParameters_ARMGenerator().Map(func(it CustomerCertificateParameters_ARM) *CustomerCertificateParameters_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["ManagedCertificate"] = ManagedCertificateParameters_ARMGenerator().Map(func(it ManagedCertificateParameters_ARM) *ManagedCertificateParameters_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["UrlSigningKey"] = UrlSigningKeyParameters_ARMGenerator().Map(func(it UrlSigningKeyParameters_ARM) *UrlSigningKeyParameters_ARM {
		return &it
	}) // generate one case for OneOf type
}

func Test_SecretProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecretProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecretProperties_ARM, SecretProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecretProperties_ARM runs a test to see if a specific instance of SecretProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecretProperties_ARM(subject SecretProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecretProperties_ARM
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

// Generator of SecretProperties_ARM instances for property testing - lazily instantiated by
// SecretProperties_ARMGenerator()
var secretProperties_ARMGenerator gopter.Gen

// SecretProperties_ARMGenerator returns a generator of SecretProperties_ARM instances for property testing.
func SecretProperties_ARMGenerator() gopter.Gen {
	if secretProperties_ARMGenerator != nil {
		return secretProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecretProperties_ARM(generators)
	secretProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SecretProperties_ARM{}), generators)

	return secretProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForSecretProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecretProperties_ARM(gens map[string]gopter.Gen) {
	gens["Parameters"] = gen.PtrOf(SecretParameters_ARMGenerator())
}

func Test_UrlSigningKeyParameters_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UrlSigningKeyParameters_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUrlSigningKeyParameters_ARM, UrlSigningKeyParameters_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUrlSigningKeyParameters_ARM runs a test to see if a specific instance of UrlSigningKeyParameters_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUrlSigningKeyParameters_ARM(subject UrlSigningKeyParameters_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UrlSigningKeyParameters_ARM
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

// Generator of UrlSigningKeyParameters_ARM instances for property testing - lazily instantiated by
// UrlSigningKeyParameters_ARMGenerator()
var urlSigningKeyParameters_ARMGenerator gopter.Gen

// UrlSigningKeyParameters_ARMGenerator returns a generator of UrlSigningKeyParameters_ARM instances for property testing.
// We first initialize urlSigningKeyParameters_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func UrlSigningKeyParameters_ARMGenerator() gopter.Gen {
	if urlSigningKeyParameters_ARMGenerator != nil {
		return urlSigningKeyParameters_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUrlSigningKeyParameters_ARM(generators)
	urlSigningKeyParameters_ARMGenerator = gen.Struct(reflect.TypeOf(UrlSigningKeyParameters_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUrlSigningKeyParameters_ARM(generators)
	AddRelatedPropertyGeneratorsForUrlSigningKeyParameters_ARM(generators)
	urlSigningKeyParameters_ARMGenerator = gen.Struct(reflect.TypeOf(UrlSigningKeyParameters_ARM{}), generators)

	return urlSigningKeyParameters_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUrlSigningKeyParameters_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUrlSigningKeyParameters_ARM(gens map[string]gopter.Gen) {
	gens["KeyId"] = gen.PtrOf(gen.AlphaString())
	gens["SecretVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(UrlSigningKeyParameters_Type_ARM_UrlSigningKey)
}

// AddRelatedPropertyGeneratorsForUrlSigningKeyParameters_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUrlSigningKeyParameters_ARM(gens map[string]gopter.Gen) {
	gens["SecretSource"] = gen.PtrOf(ResourceReference_ARMGenerator())
}
