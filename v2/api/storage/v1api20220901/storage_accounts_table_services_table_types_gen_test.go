// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

import (
	"encoding/json"
	v20220901s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/storage"
	v20230101s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/storage"
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

func Test_StorageAccountsTableServicesTable_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsTableServicesTable to hub returns original",
		prop.ForAll(RunResourceConversionTestForStorageAccountsTableServicesTable, StorageAccountsTableServicesTableGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForStorageAccountsTableServicesTable tests if a specific instance of StorageAccountsTableServicesTable round trips to the hub storage version and back losslessly
func RunResourceConversionTestForStorageAccountsTableServicesTable(subject StorageAccountsTableServicesTable) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20230101s.StorageAccountsTableServicesTable
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual StorageAccountsTableServicesTable
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

func Test_StorageAccountsTableServicesTable_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsTableServicesTable to StorageAccountsTableServicesTable via AssignProperties_To_StorageAccountsTableServicesTable & AssignProperties_From_StorageAccountsTableServicesTable returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsTableServicesTable, StorageAccountsTableServicesTableGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsTableServicesTable tests if a specific instance of StorageAccountsTableServicesTable can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsTableServicesTable(subject StorageAccountsTableServicesTable) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.StorageAccountsTableServicesTable
	err := copied.AssignProperties_To_StorageAccountsTableServicesTable(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsTableServicesTable
	err = actual.AssignProperties_From_StorageAccountsTableServicesTable(&other)
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

func Test_StorageAccountsTableServicesTable_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsTableServicesTable via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsTableServicesTable, StorageAccountsTableServicesTableGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsTableServicesTable runs a test to see if a specific instance of StorageAccountsTableServicesTable round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsTableServicesTable(subject StorageAccountsTableServicesTable) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsTableServicesTable
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

// Generator of StorageAccountsTableServicesTable instances for property testing - lazily instantiated by
// StorageAccountsTableServicesTableGenerator()
var storageAccountsTableServicesTableGenerator gopter.Gen

// StorageAccountsTableServicesTableGenerator returns a generator of StorageAccountsTableServicesTable instances for property testing.
func StorageAccountsTableServicesTableGenerator() gopter.Gen {
	if storageAccountsTableServicesTableGenerator != nil {
		return storageAccountsTableServicesTableGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable(generators)
	storageAccountsTableServicesTableGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTable{}), generators)

	return storageAccountsTableServicesTableGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccountsTableServicesTable_SpecGenerator()
	gens["Status"] = StorageAccountsTableServicesTable_STATUSGenerator()
}

func Test_StorageAccountsTableServicesTableOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsTableServicesTableOperatorSpec to StorageAccountsTableServicesTableOperatorSpec via AssignProperties_To_StorageAccountsTableServicesTableOperatorSpec & AssignProperties_From_StorageAccountsTableServicesTableOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsTableServicesTableOperatorSpec, StorageAccountsTableServicesTableOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsTableServicesTableOperatorSpec tests if a specific instance of StorageAccountsTableServicesTableOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsTableServicesTableOperatorSpec(subject StorageAccountsTableServicesTableOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.StorageAccountsTableServicesTableOperatorSpec
	err := copied.AssignProperties_To_StorageAccountsTableServicesTableOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsTableServicesTableOperatorSpec
	err = actual.AssignProperties_From_StorageAccountsTableServicesTableOperatorSpec(&other)
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

func Test_StorageAccountsTableServicesTableOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsTableServicesTableOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsTableServicesTableOperatorSpec, StorageAccountsTableServicesTableOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsTableServicesTableOperatorSpec runs a test to see if a specific instance of StorageAccountsTableServicesTableOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsTableServicesTableOperatorSpec(subject StorageAccountsTableServicesTableOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsTableServicesTableOperatorSpec
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

// Generator of StorageAccountsTableServicesTableOperatorSpec instances for property testing - lazily instantiated by
// StorageAccountsTableServicesTableOperatorSpecGenerator()
var storageAccountsTableServicesTableOperatorSpecGenerator gopter.Gen

// StorageAccountsTableServicesTableOperatorSpecGenerator returns a generator of StorageAccountsTableServicesTableOperatorSpec instances for property testing.
func StorageAccountsTableServicesTableOperatorSpecGenerator() gopter.Gen {
	if storageAccountsTableServicesTableOperatorSpecGenerator != nil {
		return storageAccountsTableServicesTableOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	storageAccountsTableServicesTableOperatorSpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTableOperatorSpec{}), generators)

	return storageAccountsTableServicesTableOperatorSpecGenerator
}

func Test_StorageAccountsTableServicesTable_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsTableServicesTable_STATUS to StorageAccountsTableServicesTable_STATUS via AssignProperties_To_StorageAccountsTableServicesTable_STATUS & AssignProperties_From_StorageAccountsTableServicesTable_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsTableServicesTable_STATUS, StorageAccountsTableServicesTable_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsTableServicesTable_STATUS tests if a specific instance of StorageAccountsTableServicesTable_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsTableServicesTable_STATUS(subject StorageAccountsTableServicesTable_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.StorageAccountsTableServicesTable_STATUS
	err := copied.AssignProperties_To_StorageAccountsTableServicesTable_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsTableServicesTable_STATUS
	err = actual.AssignProperties_From_StorageAccountsTableServicesTable_STATUS(&other)
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

func Test_StorageAccountsTableServicesTable_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsTableServicesTable_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsTableServicesTable_STATUS, StorageAccountsTableServicesTable_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsTableServicesTable_STATUS runs a test to see if a specific instance of StorageAccountsTableServicesTable_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsTableServicesTable_STATUS(subject StorageAccountsTableServicesTable_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsTableServicesTable_STATUS
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

// Generator of StorageAccountsTableServicesTable_STATUS instances for property testing - lazily instantiated by
// StorageAccountsTableServicesTable_STATUSGenerator()
var storageAccountsTableServicesTable_STATUSGenerator gopter.Gen

// StorageAccountsTableServicesTable_STATUSGenerator returns a generator of StorageAccountsTableServicesTable_STATUS instances for property testing.
// We first initialize storageAccountsTableServicesTable_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsTableServicesTable_STATUSGenerator() gopter.Gen {
	if storageAccountsTableServicesTable_STATUSGenerator != nil {
		return storageAccountsTableServicesTable_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS(generators)
	storageAccountsTableServicesTable_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTable_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS(generators)
	storageAccountsTableServicesTable_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTable_STATUS{}), generators)

	return storageAccountsTableServicesTable_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["TableName"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS(gens map[string]gopter.Gen) {
	gens["SignedIdentifiers"] = gen.SliceOf(TableSignedIdentifier_STATUSGenerator())
}

func Test_StorageAccountsTableServicesTable_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsTableServicesTable_Spec to StorageAccountsTableServicesTable_Spec via AssignProperties_To_StorageAccountsTableServicesTable_Spec & AssignProperties_From_StorageAccountsTableServicesTable_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsTableServicesTable_Spec, StorageAccountsTableServicesTable_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsTableServicesTable_Spec tests if a specific instance of StorageAccountsTableServicesTable_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsTableServicesTable_Spec(subject StorageAccountsTableServicesTable_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.StorageAccountsTableServicesTable_Spec
	err := copied.AssignProperties_To_StorageAccountsTableServicesTable_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsTableServicesTable_Spec
	err = actual.AssignProperties_From_StorageAccountsTableServicesTable_Spec(&other)
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

func Test_StorageAccountsTableServicesTable_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsTableServicesTable_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsTableServicesTable_Spec, StorageAccountsTableServicesTable_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsTableServicesTable_Spec runs a test to see if a specific instance of StorageAccountsTableServicesTable_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsTableServicesTable_Spec(subject StorageAccountsTableServicesTable_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsTableServicesTable_Spec
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

// Generator of StorageAccountsTableServicesTable_Spec instances for property testing - lazily instantiated by
// StorageAccountsTableServicesTable_SpecGenerator()
var storageAccountsTableServicesTable_SpecGenerator gopter.Gen

// StorageAccountsTableServicesTable_SpecGenerator returns a generator of StorageAccountsTableServicesTable_Spec instances for property testing.
// We first initialize storageAccountsTableServicesTable_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsTableServicesTable_SpecGenerator() gopter.Gen {
	if storageAccountsTableServicesTable_SpecGenerator != nil {
		return storageAccountsTableServicesTable_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_Spec(generators)
	storageAccountsTableServicesTable_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTable_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_Spec(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_Spec(generators)
	storageAccountsTableServicesTable_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTable_Spec{}), generators)

	return storageAccountsTableServicesTable_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(StorageAccountsTableServicesTableOperatorSpecGenerator())
	gens["SignedIdentifiers"] = gen.SliceOf(TableSignedIdentifierGenerator())
}

func Test_TableAccessPolicy_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from TableAccessPolicy to TableAccessPolicy via AssignProperties_To_TableAccessPolicy & AssignProperties_From_TableAccessPolicy returns original",
		prop.ForAll(RunPropertyAssignmentTestForTableAccessPolicy, TableAccessPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTableAccessPolicy tests if a specific instance of TableAccessPolicy can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForTableAccessPolicy(subject TableAccessPolicy) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.TableAccessPolicy
	err := copied.AssignProperties_To_TableAccessPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual TableAccessPolicy
	err = actual.AssignProperties_From_TableAccessPolicy(&other)
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

func Test_TableAccessPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableAccessPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableAccessPolicy, TableAccessPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableAccessPolicy runs a test to see if a specific instance of TableAccessPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForTableAccessPolicy(subject TableAccessPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableAccessPolicy
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

// Generator of TableAccessPolicy instances for property testing - lazily instantiated by TableAccessPolicyGenerator()
var tableAccessPolicyGenerator gopter.Gen

// TableAccessPolicyGenerator returns a generator of TableAccessPolicy instances for property testing.
func TableAccessPolicyGenerator() gopter.Gen {
	if tableAccessPolicyGenerator != nil {
		return tableAccessPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableAccessPolicy(generators)
	tableAccessPolicyGenerator = gen.Struct(reflect.TypeOf(TableAccessPolicy{}), generators)

	return tableAccessPolicyGenerator
}

// AddIndependentPropertyGeneratorsForTableAccessPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableAccessPolicy(gens map[string]gopter.Gen) {
	gens["ExpiryTime"] = gen.PtrOf(gen.AlphaString())
	gens["Permission"] = gen.PtrOf(gen.AlphaString())
	gens["StartTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_TableAccessPolicy_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from TableAccessPolicy_STATUS to TableAccessPolicy_STATUS via AssignProperties_To_TableAccessPolicy_STATUS & AssignProperties_From_TableAccessPolicy_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForTableAccessPolicy_STATUS, TableAccessPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTableAccessPolicy_STATUS tests if a specific instance of TableAccessPolicy_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForTableAccessPolicy_STATUS(subject TableAccessPolicy_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.TableAccessPolicy_STATUS
	err := copied.AssignProperties_To_TableAccessPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual TableAccessPolicy_STATUS
	err = actual.AssignProperties_From_TableAccessPolicy_STATUS(&other)
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

func Test_TableAccessPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableAccessPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableAccessPolicy_STATUS, TableAccessPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableAccessPolicy_STATUS runs a test to see if a specific instance of TableAccessPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTableAccessPolicy_STATUS(subject TableAccessPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableAccessPolicy_STATUS
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

// Generator of TableAccessPolicy_STATUS instances for property testing - lazily instantiated by
// TableAccessPolicy_STATUSGenerator()
var tableAccessPolicy_STATUSGenerator gopter.Gen

// TableAccessPolicy_STATUSGenerator returns a generator of TableAccessPolicy_STATUS instances for property testing.
func TableAccessPolicy_STATUSGenerator() gopter.Gen {
	if tableAccessPolicy_STATUSGenerator != nil {
		return tableAccessPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableAccessPolicy_STATUS(generators)
	tableAccessPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(TableAccessPolicy_STATUS{}), generators)

	return tableAccessPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTableAccessPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableAccessPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["ExpiryTime"] = gen.PtrOf(gen.AlphaString())
	gens["Permission"] = gen.PtrOf(gen.AlphaString())
	gens["StartTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_TableSignedIdentifier_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from TableSignedIdentifier to TableSignedIdentifier via AssignProperties_To_TableSignedIdentifier & AssignProperties_From_TableSignedIdentifier returns original",
		prop.ForAll(RunPropertyAssignmentTestForTableSignedIdentifier, TableSignedIdentifierGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTableSignedIdentifier tests if a specific instance of TableSignedIdentifier can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForTableSignedIdentifier(subject TableSignedIdentifier) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.TableSignedIdentifier
	err := copied.AssignProperties_To_TableSignedIdentifier(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual TableSignedIdentifier
	err = actual.AssignProperties_From_TableSignedIdentifier(&other)
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

func Test_TableSignedIdentifier_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableSignedIdentifier via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableSignedIdentifier, TableSignedIdentifierGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableSignedIdentifier runs a test to see if a specific instance of TableSignedIdentifier round trips to JSON and back losslessly
func RunJSONSerializationTestForTableSignedIdentifier(subject TableSignedIdentifier) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableSignedIdentifier
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

// Generator of TableSignedIdentifier instances for property testing - lazily instantiated by
// TableSignedIdentifierGenerator()
var tableSignedIdentifierGenerator gopter.Gen

// TableSignedIdentifierGenerator returns a generator of TableSignedIdentifier instances for property testing.
func TableSignedIdentifierGenerator() gopter.Gen {
	if tableSignedIdentifierGenerator != nil {
		return tableSignedIdentifierGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTableSignedIdentifier(generators)
	tableSignedIdentifierGenerator = gen.Struct(reflect.TypeOf(TableSignedIdentifier{}), generators)

	return tableSignedIdentifierGenerator
}

// AddRelatedPropertyGeneratorsForTableSignedIdentifier is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTableSignedIdentifier(gens map[string]gopter.Gen) {
	gens["AccessPolicy"] = gen.PtrOf(TableAccessPolicyGenerator())
}

func Test_TableSignedIdentifier_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from TableSignedIdentifier_STATUS to TableSignedIdentifier_STATUS via AssignProperties_To_TableSignedIdentifier_STATUS & AssignProperties_From_TableSignedIdentifier_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForTableSignedIdentifier_STATUS, TableSignedIdentifier_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTableSignedIdentifier_STATUS tests if a specific instance of TableSignedIdentifier_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForTableSignedIdentifier_STATUS(subject TableSignedIdentifier_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.TableSignedIdentifier_STATUS
	err := copied.AssignProperties_To_TableSignedIdentifier_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual TableSignedIdentifier_STATUS
	err = actual.AssignProperties_From_TableSignedIdentifier_STATUS(&other)
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

func Test_TableSignedIdentifier_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableSignedIdentifier_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableSignedIdentifier_STATUS, TableSignedIdentifier_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableSignedIdentifier_STATUS runs a test to see if a specific instance of TableSignedIdentifier_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTableSignedIdentifier_STATUS(subject TableSignedIdentifier_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableSignedIdentifier_STATUS
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

// Generator of TableSignedIdentifier_STATUS instances for property testing - lazily instantiated by
// TableSignedIdentifier_STATUSGenerator()
var tableSignedIdentifier_STATUSGenerator gopter.Gen

// TableSignedIdentifier_STATUSGenerator returns a generator of TableSignedIdentifier_STATUS instances for property testing.
// We first initialize tableSignedIdentifier_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TableSignedIdentifier_STATUSGenerator() gopter.Gen {
	if tableSignedIdentifier_STATUSGenerator != nil {
		return tableSignedIdentifier_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS(generators)
	tableSignedIdentifier_STATUSGenerator = gen.Struct(reflect.TypeOf(TableSignedIdentifier_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS(generators)
	AddRelatedPropertyGeneratorsForTableSignedIdentifier_STATUS(generators)
	tableSignedIdentifier_STATUSGenerator = gen.Struct(reflect.TypeOf(TableSignedIdentifier_STATUS{}), generators)

	return tableSignedIdentifier_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTableSignedIdentifier_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTableSignedIdentifier_STATUS(gens map[string]gopter.Gen) {
	gens["AccessPolicy"] = gen.PtrOf(TableAccessPolicy_STATUSGenerator())
}
