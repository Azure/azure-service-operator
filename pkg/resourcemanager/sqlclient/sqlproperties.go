// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
)

// DBEdition - wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#DatabaseEdition
type DBEdition byte

const (
	// Basic ...
	Basic DBEdition = 0
	// Business ...
	Business DBEdition = 1
	// BusinessCritical ...
	BusinessCritical DBEdition = 2
	// DataWarehouse ...
	DataWarehouse DBEdition = 3
	// Free ...
	Free DBEdition = 4
	// GeneralPurpose ...
	GeneralPurpose DBEdition = 5
	// Hyperscale ...
	Hyperscale DBEdition = 6
	// Premium ...
	Premium DBEdition = 7
	// PremiumRS ...
	PremiumRS DBEdition = 8
	// Standard ...
	Standard DBEdition = 9
	// Stretch ...
	Stretch DBEdition = 10
	// System ...
	System DBEdition = 11
	// System2 ...
	System2 DBEdition = 12
	// Web ...
	Web DBEdition = 13
)

// SQLServerProperties contains values needed for adding / updating SQL servers,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#Server
// also wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#ServerProperties
type SQLServerProperties struct {

	// AdministratorLogin - Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string

	// AdministratorLoginPassword - The administrator login password (required for server creation).
	AdministratorLoginPassword *string
}

// SQLDatabaseProperties contains values needed for adding / updating SQL servers,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#Database
// also wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#DatabaseProperties
type SQLDatabaseProperties struct {

	// DatabaseName is the name of the database
	DatabaseName string

	// Edition - The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	//
	// ```azurecli
	// az sql db list-editions -l <location> -o table
	// ````
	//
	// ```powershell
	// Get-AzSqlServerServiceObjective -Location <location>
	// ````
	// . Possible values include: 'Web', 'Business', 'Basic', 'Standard', 'Premium', 'PremiumRS', 'Free', 'Stretch', 'DataWarehouse', 'System', 'System2', 'GeneralPurpose', 'BusinessCritical', 'Hyperscale'
	Edition DBEdition
}

// SQLServerPropertiesToServer translates SQLServerProperties to ServerProperties
func SQLServerPropertiesToServer(properties SQLServerProperties) (result sql.ServerProperties) {

	result = sql.ServerProperties{
		AdministratorLogin:         properties.AdministratorLogin,
		AdministratorLoginPassword: properties.AdministratorLoginPassword,
	}

	return result
}

// SQLDatabasePropertiesToDatabase translates SQLDatabaseProperties to DatabaseProperties
func SQLDatabasePropertiesToDatabase(properties SQLDatabaseProperties) (result sql.DatabaseProperties) {

	result = sql.DatabaseProperties{
		Edition: translateDBEdition(properties.Edition),
	}

	return result
}

// translateDBEdition translates enums
func translateDBEdition(in DBEdition) (result sql.DatabaseEdition) {
	switch in {
	case 0:
		result = sql.Basic
	case 1:
		result = sql.Business
	case 2:
		result = sql.BusinessCritical
	case 3:
		result = sql.DataWarehouse
	case 4:
		result = sql.Free
	case 5:
		result = sql.GeneralPurpose
	case 6:
		result = sql.Hyperscale
	case 7:
		result = sql.Premium
	case 8:
		result = sql.PremiumRS
	case 9:
		result = sql.Standard
	case 10:
		result = sql.Stretch
	case 11:
		result = sql.System
	case 12:
		result = sql.System2
	case 13:
		result = sql.Web
	default:
		result = sql.Free
	}

	return result
}

// AvailabilityResponse is the response for checking name validation
type AvailabilityResponse struct {
	Available bool
	Message   string
	Name      string
}

// ToAvailabilityResponse converts CheckNameAvailabilityResponse to AvailabilityResponse
func ToAvailabilityResponse(response sql.CheckNameAvailabilityResponse) (result AvailabilityResponse) {
	result.Available = *response.Available
	if response.Message != nil {
		result.Message = *response.Message
	}
	if response.Name != nil {
		result.Name = *response.Name
	}

	return result
}
