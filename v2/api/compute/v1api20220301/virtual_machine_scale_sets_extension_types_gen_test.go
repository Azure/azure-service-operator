// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220301

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/compute/v1api20220301/storage"
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

func Test_KeyVaultSecretReference_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from KeyVaultSecretReference to KeyVaultSecretReference via AssignProperties_To_KeyVaultSecretReference & AssignProperties_From_KeyVaultSecretReference returns original",
		prop.ForAll(RunPropertyAssignmentTestForKeyVaultSecretReference, KeyVaultSecretReferenceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForKeyVaultSecretReference tests if a specific instance of KeyVaultSecretReference can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForKeyVaultSecretReference(subject KeyVaultSecretReference) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.KeyVaultSecretReference
	err := copied.AssignProperties_To_KeyVaultSecretReference(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual KeyVaultSecretReference
	err = actual.AssignProperties_From_KeyVaultSecretReference(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_KeyVaultSecretReference_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultSecretReference via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultSecretReference, KeyVaultSecretReferenceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultSecretReference runs a test to see if a specific instance of KeyVaultSecretReference round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultSecretReference(subject KeyVaultSecretReference) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultSecretReference
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

// Generator of KeyVaultSecretReference instances for property testing - lazily instantiated by
// KeyVaultSecretReferenceGenerator()
var keyVaultSecretReferenceGenerator gopter.Gen

// KeyVaultSecretReferenceGenerator returns a generator of KeyVaultSecretReference instances for property testing.
// We first initialize keyVaultSecretReferenceGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultSecretReferenceGenerator() gopter.Gen {
	if keyVaultSecretReferenceGenerator != nil {
		return keyVaultSecretReferenceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultSecretReference(generators)
	keyVaultSecretReferenceGenerator = gen.Struct(reflect.TypeOf(KeyVaultSecretReference{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultSecretReference(generators)
	AddRelatedPropertyGeneratorsForKeyVaultSecretReference(generators)
	keyVaultSecretReferenceGenerator = gen.Struct(reflect.TypeOf(KeyVaultSecretReference{}), generators)

	return keyVaultSecretReferenceGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultSecretReference is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultSecretReference(gens map[string]gopter.Gen) {
	gens["SecretUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultSecretReference is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultSecretReference(gens map[string]gopter.Gen) {
	gens["SourceVault"] = gen.PtrOf(SubResourceGenerator())
}

func Test_KeyVaultSecretReference_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from KeyVaultSecretReference_STATUS to KeyVaultSecretReference_STATUS via AssignProperties_To_KeyVaultSecretReference_STATUS & AssignProperties_From_KeyVaultSecretReference_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForKeyVaultSecretReference_STATUS, KeyVaultSecretReference_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForKeyVaultSecretReference_STATUS tests if a specific instance of KeyVaultSecretReference_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForKeyVaultSecretReference_STATUS(subject KeyVaultSecretReference_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.KeyVaultSecretReference_STATUS
	err := copied.AssignProperties_To_KeyVaultSecretReference_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual KeyVaultSecretReference_STATUS
	err = actual.AssignProperties_From_KeyVaultSecretReference_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_KeyVaultSecretReference_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultSecretReference_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultSecretReference_STATUS, KeyVaultSecretReference_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultSecretReference_STATUS runs a test to see if a specific instance of KeyVaultSecretReference_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultSecretReference_STATUS(subject KeyVaultSecretReference_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultSecretReference_STATUS
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

// Generator of KeyVaultSecretReference_STATUS instances for property testing - lazily instantiated by
// KeyVaultSecretReference_STATUSGenerator()
var keyVaultSecretReference_STATUSGenerator gopter.Gen

// KeyVaultSecretReference_STATUSGenerator returns a generator of KeyVaultSecretReference_STATUS instances for property testing.
// We first initialize keyVaultSecretReference_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultSecretReference_STATUSGenerator() gopter.Gen {
	if keyVaultSecretReference_STATUSGenerator != nil {
		return keyVaultSecretReference_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultSecretReference_STATUS(generators)
	keyVaultSecretReference_STATUSGenerator = gen.Struct(reflect.TypeOf(KeyVaultSecretReference_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultSecretReference_STATUS(generators)
	AddRelatedPropertyGeneratorsForKeyVaultSecretReference_STATUS(generators)
	keyVaultSecretReference_STATUSGenerator = gen.Struct(reflect.TypeOf(KeyVaultSecretReference_STATUS{}), generators)

	return keyVaultSecretReference_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultSecretReference_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultSecretReference_STATUS(gens map[string]gopter.Gen) {
	gens["SecretUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultSecretReference_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultSecretReference_STATUS(gens map[string]gopter.Gen) {
	gens["SourceVault"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_VirtualMachineScaleSetsExtension_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualMachineScaleSetsExtension to hub returns original",
		prop.ForAll(RunResourceConversionTestForVirtualMachineScaleSetsExtension, VirtualMachineScaleSetsExtensionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForVirtualMachineScaleSetsExtension tests if a specific instance of VirtualMachineScaleSetsExtension round trips to the hub storage version and back losslessly
func RunResourceConversionTestForVirtualMachineScaleSetsExtension(subject VirtualMachineScaleSetsExtension) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.VirtualMachineScaleSetsExtension
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual VirtualMachineScaleSetsExtension
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualMachineScaleSetsExtension_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualMachineScaleSetsExtension to VirtualMachineScaleSetsExtension via AssignProperties_To_VirtualMachineScaleSetsExtension & AssignProperties_From_VirtualMachineScaleSetsExtension returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualMachineScaleSetsExtension, VirtualMachineScaleSetsExtensionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualMachineScaleSetsExtension tests if a specific instance of VirtualMachineScaleSetsExtension can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualMachineScaleSetsExtension(subject VirtualMachineScaleSetsExtension) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.VirtualMachineScaleSetsExtension
	err := copied.AssignProperties_To_VirtualMachineScaleSetsExtension(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualMachineScaleSetsExtension
	err = actual.AssignProperties_From_VirtualMachineScaleSetsExtension(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualMachineScaleSetsExtension_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineScaleSetsExtension via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineScaleSetsExtension, VirtualMachineScaleSetsExtensionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineScaleSetsExtension runs a test to see if a specific instance of VirtualMachineScaleSetsExtension round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineScaleSetsExtension(subject VirtualMachineScaleSetsExtension) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineScaleSetsExtension
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

// Generator of VirtualMachineScaleSetsExtension instances for property testing - lazily instantiated by
// VirtualMachineScaleSetsExtensionGenerator()
var virtualMachineScaleSetsExtensionGenerator gopter.Gen

// VirtualMachineScaleSetsExtensionGenerator returns a generator of VirtualMachineScaleSetsExtension instances for property testing.
func VirtualMachineScaleSetsExtensionGenerator() gopter.Gen {
	if virtualMachineScaleSetsExtensionGenerator != nil {
		return virtualMachineScaleSetsExtensionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension(generators)
	virtualMachineScaleSetsExtensionGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetsExtension{}), generators)

	return virtualMachineScaleSetsExtensionGenerator
}

// AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension(gens map[string]gopter.Gen) {
	gens["Spec"] = VirtualMachineScaleSetsExtension_SpecGenerator()
	gens["Status"] = VirtualMachineScaleSetsExtension_STATUSGenerator()
}

func Test_VirtualMachineScaleSetsExtensionOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualMachineScaleSetsExtensionOperatorSpec to VirtualMachineScaleSetsExtensionOperatorSpec via AssignProperties_To_VirtualMachineScaleSetsExtensionOperatorSpec & AssignProperties_From_VirtualMachineScaleSetsExtensionOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualMachineScaleSetsExtensionOperatorSpec, VirtualMachineScaleSetsExtensionOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualMachineScaleSetsExtensionOperatorSpec tests if a specific instance of VirtualMachineScaleSetsExtensionOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualMachineScaleSetsExtensionOperatorSpec(subject VirtualMachineScaleSetsExtensionOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.VirtualMachineScaleSetsExtensionOperatorSpec
	err := copied.AssignProperties_To_VirtualMachineScaleSetsExtensionOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualMachineScaleSetsExtensionOperatorSpec
	err = actual.AssignProperties_From_VirtualMachineScaleSetsExtensionOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualMachineScaleSetsExtensionOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineScaleSetsExtensionOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineScaleSetsExtensionOperatorSpec, VirtualMachineScaleSetsExtensionOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineScaleSetsExtensionOperatorSpec runs a test to see if a specific instance of VirtualMachineScaleSetsExtensionOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineScaleSetsExtensionOperatorSpec(subject VirtualMachineScaleSetsExtensionOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineScaleSetsExtensionOperatorSpec
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

// Generator of VirtualMachineScaleSetsExtensionOperatorSpec instances for property testing - lazily instantiated by
// VirtualMachineScaleSetsExtensionOperatorSpecGenerator()
var virtualMachineScaleSetsExtensionOperatorSpecGenerator gopter.Gen

// VirtualMachineScaleSetsExtensionOperatorSpecGenerator returns a generator of VirtualMachineScaleSetsExtensionOperatorSpec instances for property testing.
func VirtualMachineScaleSetsExtensionOperatorSpecGenerator() gopter.Gen {
	if virtualMachineScaleSetsExtensionOperatorSpecGenerator != nil {
		return virtualMachineScaleSetsExtensionOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	virtualMachineScaleSetsExtensionOperatorSpecGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetsExtensionOperatorSpec{}), generators)

	return virtualMachineScaleSetsExtensionOperatorSpecGenerator
}

func Test_VirtualMachineScaleSetsExtension_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualMachineScaleSetsExtension_STATUS to VirtualMachineScaleSetsExtension_STATUS via AssignProperties_To_VirtualMachineScaleSetsExtension_STATUS & AssignProperties_From_VirtualMachineScaleSetsExtension_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualMachineScaleSetsExtension_STATUS, VirtualMachineScaleSetsExtension_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualMachineScaleSetsExtension_STATUS tests if a specific instance of VirtualMachineScaleSetsExtension_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualMachineScaleSetsExtension_STATUS(subject VirtualMachineScaleSetsExtension_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.VirtualMachineScaleSetsExtension_STATUS
	err := copied.AssignProperties_To_VirtualMachineScaleSetsExtension_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualMachineScaleSetsExtension_STATUS
	err = actual.AssignProperties_From_VirtualMachineScaleSetsExtension_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualMachineScaleSetsExtension_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineScaleSetsExtension_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineScaleSetsExtension_STATUS, VirtualMachineScaleSetsExtension_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineScaleSetsExtension_STATUS runs a test to see if a specific instance of VirtualMachineScaleSetsExtension_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineScaleSetsExtension_STATUS(subject VirtualMachineScaleSetsExtension_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineScaleSetsExtension_STATUS
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

// Generator of VirtualMachineScaleSetsExtension_STATUS instances for property testing - lazily instantiated by
// VirtualMachineScaleSetsExtension_STATUSGenerator()
var virtualMachineScaleSetsExtension_STATUSGenerator gopter.Gen

// VirtualMachineScaleSetsExtension_STATUSGenerator returns a generator of VirtualMachineScaleSetsExtension_STATUS instances for property testing.
// We first initialize virtualMachineScaleSetsExtension_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineScaleSetsExtension_STATUSGenerator() gopter.Gen {
	if virtualMachineScaleSetsExtension_STATUSGenerator != nil {
		return virtualMachineScaleSetsExtension_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS(generators)
	virtualMachineScaleSetsExtension_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetsExtension_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS(generators)
	virtualMachineScaleSetsExtension_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetsExtension_STATUS{}), generators)

	return virtualMachineScaleSetsExtension_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["EnableAutomaticUpgrade"] = gen.PtrOf(gen.Bool())
	gens["ForceUpdateTag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesType"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisionAfterExtensions"] = gen.SliceOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["SuppressFailures"] = gen.PtrOf(gen.Bool())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS(gens map[string]gopter.Gen) {
	gens["ProtectedSettingsFromKeyVault"] = gen.PtrOf(KeyVaultSecretReference_STATUSGenerator())
}

func Test_VirtualMachineScaleSetsExtension_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualMachineScaleSetsExtension_Spec to VirtualMachineScaleSetsExtension_Spec via AssignProperties_To_VirtualMachineScaleSetsExtension_Spec & AssignProperties_From_VirtualMachineScaleSetsExtension_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualMachineScaleSetsExtension_Spec, VirtualMachineScaleSetsExtension_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualMachineScaleSetsExtension_Spec tests if a specific instance of VirtualMachineScaleSetsExtension_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualMachineScaleSetsExtension_Spec(subject VirtualMachineScaleSetsExtension_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.VirtualMachineScaleSetsExtension_Spec
	err := copied.AssignProperties_To_VirtualMachineScaleSetsExtension_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualMachineScaleSetsExtension_Spec
	err = actual.AssignProperties_From_VirtualMachineScaleSetsExtension_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualMachineScaleSetsExtension_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineScaleSetsExtension_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineScaleSetsExtension_Spec, VirtualMachineScaleSetsExtension_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineScaleSetsExtension_Spec runs a test to see if a specific instance of VirtualMachineScaleSetsExtension_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineScaleSetsExtension_Spec(subject VirtualMachineScaleSetsExtension_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineScaleSetsExtension_Spec
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

// Generator of VirtualMachineScaleSetsExtension_Spec instances for property testing - lazily instantiated by
// VirtualMachineScaleSetsExtension_SpecGenerator()
var virtualMachineScaleSetsExtension_SpecGenerator gopter.Gen

// VirtualMachineScaleSetsExtension_SpecGenerator returns a generator of VirtualMachineScaleSetsExtension_Spec instances for property testing.
// We first initialize virtualMachineScaleSetsExtension_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineScaleSetsExtension_SpecGenerator() gopter.Gen {
	if virtualMachineScaleSetsExtension_SpecGenerator != nil {
		return virtualMachineScaleSetsExtension_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_Spec(generators)
	virtualMachineScaleSetsExtension_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetsExtension_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_Spec(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension_Spec(generators)
	virtualMachineScaleSetsExtension_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetsExtension_Spec{}), generators)

	return virtualMachineScaleSetsExtension_SpecGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_Spec(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["EnableAutomaticUpgrade"] = gen.PtrOf(gen.Bool())
	gens["ForceUpdateTag"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisionAfterExtensions"] = gen.SliceOf(gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["SuppressFailures"] = gen.PtrOf(gen.Bool())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(VirtualMachineScaleSetsExtensionOperatorSpecGenerator())
	gens["ProtectedSettingsFromKeyVault"] = gen.PtrOf(KeyVaultSecretReferenceGenerator())
}
