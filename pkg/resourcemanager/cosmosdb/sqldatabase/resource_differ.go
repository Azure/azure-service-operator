// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package sqldatabase

import "github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2021-03-15/documentdb"

func doResourcesMatch(expected *documentdb.SQLDatabaseResource, actual *documentdb.SQLDatabaseGetPropertiesResource) bool {
	var expectedIDCoalesced *string
	if expected != nil {
		expectedIDCoalesced = expected.ID
	}

	var actualIDCoalesced *string
	if actual != nil {
		actualIDCoalesced = actual.ID
	}

	if (expectedIDCoalesced == nil) != (actualIDCoalesced == nil) {
		return false
	}
	if expectedIDCoalesced != nil && actualIDCoalesced != nil && *expectedIDCoalesced != *actualIDCoalesced {
		return false
	}

	return true
}

// doesThroughputMatch checks if the throughput specified on an Azure SQL Cosmos DB matches the throughput settings on the
// child resource. We check the child resource because Cosmos DB has a bug where the throughput stuff doesn't get returned
// on a GET of the SQL DB itself.
func doesThroughputMatch(expected *documentdb.CreateUpdateOptions, actual *documentdb.ThroughputSettingsGetProperties) bool {
	var expectedThroughputCoalesced *int32
	var expectedAutoscaleCoalesced *documentdb.AutoscaleSettings
	if expected != nil {
		expectedThroughputCoalesced = expected.Throughput
		expectedAutoscaleCoalesced = expected.AutoscaleSettings
	}

	var actualThroughputCoalesced *int32
	var actualAutoscaleCoalesced *documentdb.AutoscaleSettingsResource
	if actual != nil && actual.Resource != nil {
		actualThroughputCoalesced = actual.Resource.Throughput
		actualAutoscaleCoalesced = actual.Resource.AutoscaleSettings
	}

	// Throughput
	if (expectedThroughputCoalesced == nil) != (actualThroughputCoalesced == nil) {
		return false
	}
	if expectedThroughputCoalesced != nil && actualThroughputCoalesced != nil &&
		*expectedThroughputCoalesced != *actualThroughputCoalesced {
		return false
	}

	if !doAutoscaleSettingsMatch(expectedAutoscaleCoalesced, actualAutoscaleCoalesced) {
		return false
	}

	return true
}

func doAutoscaleSettingsMatch(expected *documentdb.AutoscaleSettings, actual *documentdb.AutoscaleSettingsResource) bool {
	var expectedMaxThroughputCoalesced *int32
	if expected != nil {
		expectedMaxThroughputCoalesced = expected.MaxThroughput
	}

	var actualMaxThroughputCoalesced *int32
	if actual != nil {
		actualMaxThroughputCoalesced = actual.MaxThroughput
	}

	if (expectedMaxThroughputCoalesced == nil) != (actualMaxThroughputCoalesced == nil) {
		return false
	}

	if expectedMaxThroughputCoalesced != nil && actualMaxThroughputCoalesced != nil {
		if *expectedMaxThroughputCoalesced != *actualMaxThroughputCoalesced {
			return false
		}
	}

	return true
}

func doSQLDatabasePropertiesMatch(expected *documentdb.SQLDatabaseCreateUpdateProperties, actual *documentdb.SQLDatabaseGetProperties) bool {
	var expectedResourceCoalesced *documentdb.SQLDatabaseResource
	if expected != nil {
		expectedResourceCoalesced = expected.Resource
	}

	var actualResourceCoalesced *documentdb.SQLDatabaseGetPropertiesResource
	if actual != nil {
		actualResourceCoalesced = actual.Resource
	}

	if !doResourcesMatch(expectedResourceCoalesced, actualResourceCoalesced) {
		return false
	}

	return true
}

func doTagsMatch(expected map[string]*string, actual map[string]*string) bool {
	if len(expected) != len(actual) {
		return false
	}
	for expectedKey, expectedValue := range expected {
		actualValue, ok := actual[expectedKey]
		if !ok {
			return false
		}

		if (expectedValue == nil) != (actualValue == nil) {
			return false
		}

		if expectedValue != nil && actualValue != nil && *expectedValue != *actualValue {
			return false
		}
	}

	return true
}

func DoesResourceMatchAzure(
	expected documentdb.SQLDatabaseCreateUpdateParameters,
	actual documentdb.SQLDatabaseGetResults,
	actualThroughput documentdb.ThroughputSettingsGetResults) bool {

	if !doTagsMatch(expected.Tags, actual.Tags) {
		return false
	}

	if !doSQLDatabasePropertiesMatch(expected.SQLDatabaseCreateUpdateProperties, actual.SQLDatabaseGetProperties) {
		return false
	}

	if expected.SQLDatabaseCreateUpdateProperties != nil {
		if !doesThroughputMatch(expected.SQLDatabaseCreateUpdateProperties.Options, actualThroughput.ThroughputSettingsGetProperties) {
			return false
		}
	}

	return true
}
