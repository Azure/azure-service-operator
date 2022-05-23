// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

import (
	"encoding/json"
	v20180601s "github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1beta20180601storage"
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

func Test_Database_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Database to hub returns original",
		prop.ForAll(RunResourceConversionTestForDatabase, DatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDatabase tests if a specific instance of Database round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDatabase(subject Database) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20180601s.Database
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Database
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

func Test_Database_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Database to Database via AssignPropertiesToDatabase & AssignPropertiesFromDatabase returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabase, DatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabase tests if a specific instance of Database can be assigned to v1beta20180601storage and back losslessly
func RunPropertyAssignmentTestForDatabase(subject Database) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180601s.Database
	err := copied.AssignPropertiesToDatabase(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Database
	err = actual.AssignPropertiesFromDatabase(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Database_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabase, DatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabase runs a test to see if a specific instance of Database round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabase(subject Database) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database
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

// Generator of Database instances for property testing - lazily instantiated by DatabaseGenerator()
var databaseGenerator gopter.Gen

// DatabaseGenerator returns a generator of Database instances for property testing.
func DatabaseGenerator() gopter.Gen {
	if databaseGenerator != nil {
		return databaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDatabase(generators)
	databaseGenerator = gen.Struct(reflect.TypeOf(Database{}), generators)

	return databaseGenerator
}

// AddRelatedPropertyGeneratorsForDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = ServersDatabasesSpecGenerator()
	gens["Status"] = DatabaseStatusGenerator()
}

func Test_Database_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Database_Status to Database_Status via AssignPropertiesToDatabaseStatus & AssignPropertiesFromDatabaseStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseStatus, DatabaseStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseStatus tests if a specific instance of Database_Status can be assigned to v1beta20180601storage and back losslessly
func RunPropertyAssignmentTestForDatabaseStatus(subject Database_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180601s.Database_Status
	err := copied.AssignPropertiesToDatabaseStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Database_Status
	err = actual.AssignPropertiesFromDatabaseStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Database_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseStatus, DatabaseStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseStatus runs a test to see if a specific instance of Database_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseStatus(subject Database_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_Status
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

// Generator of Database_Status instances for property testing - lazily instantiated by DatabaseStatusGenerator()
var databaseStatusGenerator gopter.Gen

// DatabaseStatusGenerator returns a generator of Database_Status instances for property testing.
func DatabaseStatusGenerator() gopter.Gen {
	if databaseStatusGenerator != nil {
		return databaseStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseStatus(generators)
	databaseStatusGenerator = gen.Struct(reflect.TypeOf(Database_Status{}), generators)

	return databaseStatusGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseStatus(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersDatabases_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersDatabases_Spec to ServersDatabases_Spec via AssignPropertiesToServersDatabasesSpec & AssignPropertiesFromServersDatabasesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersDatabasesSpec, ServersDatabasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersDatabasesSpec tests if a specific instance of ServersDatabases_Spec can be assigned to v1beta20180601storage and back losslessly
func RunPropertyAssignmentTestForServersDatabasesSpec(subject ServersDatabases_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180601s.ServersDatabases_Spec
	err := copied.AssignPropertiesToServersDatabasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersDatabases_Spec
	err = actual.AssignPropertiesFromServersDatabasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersDatabases_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabases_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesSpec, ServersDatabasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesSpec runs a test to see if a specific instance of ServersDatabases_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesSpec(subject ServersDatabases_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabases_Spec
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

// Generator of ServersDatabases_Spec instances for property testing - lazily instantiated by
// ServersDatabasesSpecGenerator()
var serversDatabasesSpecGenerator gopter.Gen

// ServersDatabasesSpecGenerator returns a generator of ServersDatabases_Spec instances for property testing.
func ServersDatabasesSpecGenerator() gopter.Gen {
	if serversDatabasesSpecGenerator != nil {
		return serversDatabasesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesSpec(generators)
	serversDatabasesSpecGenerator = gen.Struct(reflect.TypeOf(ServersDatabases_Spec{}), generators)

	return serversDatabasesSpecGenerator
}

// AddIndependentPropertyGeneratorsForServersDatabasesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersDatabasesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
