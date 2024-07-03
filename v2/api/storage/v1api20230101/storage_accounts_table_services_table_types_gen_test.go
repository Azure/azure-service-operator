// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230101

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/storage"
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
	var hub storage.StorageAccountsTableServicesTable
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
	var other storage.StorageAccountsTableServicesTable
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
	gens["Spec"] = StorageAccounts_TableServices_Table_SpecGenerator()
	gens["Status"] = StorageAccounts_TableServices_Table_STATUSGenerator()
}

func Test_StorageAccounts_TableServices_Table_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccounts_TableServices_Table_STATUS to StorageAccounts_TableServices_Table_STATUS via AssignProperties_To_StorageAccounts_TableServices_Table_STATUS & AssignProperties_From_StorageAccounts_TableServices_Table_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccounts_TableServices_Table_STATUS, StorageAccounts_TableServices_Table_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccounts_TableServices_Table_STATUS tests if a specific instance of StorageAccounts_TableServices_Table_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccounts_TableServices_Table_STATUS(subject StorageAccounts_TableServices_Table_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.StorageAccounts_TableServices_Table_STATUS
	err := copied.AssignProperties_To_StorageAccounts_TableServices_Table_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccounts_TableServices_Table_STATUS
	err = actual.AssignProperties_From_StorageAccounts_TableServices_Table_STATUS(&other)
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

func Test_StorageAccounts_TableServices_Table_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_TableServices_Table_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_TableServices_Table_STATUS, StorageAccounts_TableServices_Table_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_TableServices_Table_STATUS runs a test to see if a specific instance of StorageAccounts_TableServices_Table_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_TableServices_Table_STATUS(subject StorageAccounts_TableServices_Table_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_TableServices_Table_STATUS
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

// Generator of StorageAccounts_TableServices_Table_STATUS instances for property testing - lazily instantiated by
// StorageAccounts_TableServices_Table_STATUSGenerator()
var storageAccounts_TableServices_Table_STATUSGenerator gopter.Gen

// StorageAccounts_TableServices_Table_STATUSGenerator returns a generator of StorageAccounts_TableServices_Table_STATUS instances for property testing.
// We first initialize storageAccounts_TableServices_Table_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_TableServices_Table_STATUSGenerator() gopter.Gen {
	if storageAccounts_TableServices_Table_STATUSGenerator != nil {
		return storageAccounts_TableServices_Table_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_STATUS(generators)
	storageAccounts_TableServices_Table_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_TableServices_Table_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_TableServices_Table_STATUS(generators)
	storageAccounts_TableServices_Table_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_TableServices_Table_STATUS{}), generators)

	return storageAccounts_TableServices_Table_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["TableName"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccounts_TableServices_Table_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_TableServices_Table_STATUS(gens map[string]gopter.Gen) {
	gens["SignedIdentifiers"] = gen.SliceOf(TableSignedIdentifier_STATUSGenerator())
}

func Test_StorageAccounts_TableServices_Table_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccounts_TableServices_Table_Spec to StorageAccounts_TableServices_Table_Spec via AssignProperties_To_StorageAccounts_TableServices_Table_Spec & AssignProperties_From_StorageAccounts_TableServices_Table_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccounts_TableServices_Table_Spec, StorageAccounts_TableServices_Table_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccounts_TableServices_Table_Spec tests if a specific instance of StorageAccounts_TableServices_Table_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccounts_TableServices_Table_Spec(subject StorageAccounts_TableServices_Table_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.StorageAccounts_TableServices_Table_Spec
	err := copied.AssignProperties_To_StorageAccounts_TableServices_Table_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccounts_TableServices_Table_Spec
	err = actual.AssignProperties_From_StorageAccounts_TableServices_Table_Spec(&other)
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

func Test_StorageAccounts_TableServices_Table_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_TableServices_Table_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_TableServices_Table_Spec, StorageAccounts_TableServices_Table_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_TableServices_Table_Spec runs a test to see if a specific instance of StorageAccounts_TableServices_Table_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_TableServices_Table_Spec(subject StorageAccounts_TableServices_Table_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_TableServices_Table_Spec
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

// Generator of StorageAccounts_TableServices_Table_Spec instances for property testing - lazily instantiated by
// StorageAccounts_TableServices_Table_SpecGenerator()
var storageAccounts_TableServices_Table_SpecGenerator gopter.Gen

// StorageAccounts_TableServices_Table_SpecGenerator returns a generator of StorageAccounts_TableServices_Table_Spec instances for property testing.
// We first initialize storageAccounts_TableServices_Table_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_TableServices_Table_SpecGenerator() gopter.Gen {
	if storageAccounts_TableServices_Table_SpecGenerator != nil {
		return storageAccounts_TableServices_Table_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec(generators)
	storageAccounts_TableServices_Table_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_TableServices_Table_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec(generators)
	storageAccounts_TableServices_Table_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_TableServices_Table_Spec{}), generators)

	return storageAccounts_TableServices_Table_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec(gens map[string]gopter.Gen) {
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
	var other storage.TableAccessPolicy
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
	var other storage.TableAccessPolicy_STATUS
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
	var other storage.TableSignedIdentifier
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
	var other storage.TableSignedIdentifier_STATUS
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
