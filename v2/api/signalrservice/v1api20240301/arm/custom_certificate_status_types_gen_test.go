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

func Test_CustomCertificateProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CustomCertificateProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCustomCertificateProperties_STATUS, CustomCertificateProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCustomCertificateProperties_STATUS runs a test to see if a specific instance of CustomCertificateProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCustomCertificateProperties_STATUS(subject CustomCertificateProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CustomCertificateProperties_STATUS
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

// Generator of CustomCertificateProperties_STATUS instances for property testing - lazily instantiated by
// CustomCertificateProperties_STATUSGenerator()
var customCertificateProperties_STATUSGenerator gopter.Gen

// CustomCertificateProperties_STATUSGenerator returns a generator of CustomCertificateProperties_STATUS instances for property testing.
func CustomCertificateProperties_STATUSGenerator() gopter.Gen {
	if customCertificateProperties_STATUSGenerator != nil {
		return customCertificateProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomCertificateProperties_STATUS(generators)
	customCertificateProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(CustomCertificateProperties_STATUS{}), generators)

	return customCertificateProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCustomCertificateProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCustomCertificateProperties_STATUS(gens map[string]gopter.Gen) {
	gens["KeyVaultBaseUri"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultSecretName"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultSecretVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Canceled,
		ProvisioningState_STATUS_Creating,
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Moving,
		ProvisioningState_STATUS_Running,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Unknown,
		ProvisioningState_STATUS_Updating))
}

func Test_CustomCertificate_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CustomCertificate_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCustomCertificate_STATUS, CustomCertificate_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCustomCertificate_STATUS runs a test to see if a specific instance of CustomCertificate_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCustomCertificate_STATUS(subject CustomCertificate_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CustomCertificate_STATUS
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

// Generator of CustomCertificate_STATUS instances for property testing - lazily instantiated by
// CustomCertificate_STATUSGenerator()
var customCertificate_STATUSGenerator gopter.Gen

// CustomCertificate_STATUSGenerator returns a generator of CustomCertificate_STATUS instances for property testing.
// We first initialize customCertificate_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CustomCertificate_STATUSGenerator() gopter.Gen {
	if customCertificate_STATUSGenerator != nil {
		return customCertificate_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomCertificate_STATUS(generators)
	customCertificate_STATUSGenerator = gen.Struct(reflect.TypeOf(CustomCertificate_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomCertificate_STATUS(generators)
	AddRelatedPropertyGeneratorsForCustomCertificate_STATUS(generators)
	customCertificate_STATUSGenerator = gen.Struct(reflect.TypeOf(CustomCertificate_STATUS{}), generators)

	return customCertificate_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCustomCertificate_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCustomCertificate_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForCustomCertificate_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCustomCertificate_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(CustomCertificateProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_SystemData_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS runs a test to see if a specific instance of SystemData_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS
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

// Generator of SystemData_STATUS instances for property testing - lazily instantiated by SystemData_STATUSGenerator()
var systemData_STATUSGenerator gopter.Gen

// SystemData_STATUSGenerator returns a generator of SystemData_STATUS instances for property testing.
func SystemData_STATUSGenerator() gopter.Gen {
	if systemData_STATUSGenerator != nil {
		return systemData_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS(generators)
	systemData_STATUSGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS{}), generators)

	return systemData_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_Application,
		SystemData_CreatedByType_STATUS_Key,
		SystemData_CreatedByType_STATUS_ManagedIdentity,
		SystemData_CreatedByType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_Application,
		SystemData_LastModifiedByType_STATUS_Key,
		SystemData_LastModifiedByType_STATUS_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_User))
}
