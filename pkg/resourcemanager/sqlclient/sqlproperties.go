// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
)

// DBAddition - wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#DatabaseEdition
type DBAddition byte

const (
	// Basic ...
	Basic DBAddition = 0
	// Business ...
	Business DBAddition = 1
	// BusinessCritical ...
	BusinessCritical DBAddition = 2
	// DataWarehouse ...
	DataWarehouse DBAddition = 3
	// Free ...
	Free DBAddition = 4
	// GeneralPurpose ...
	GeneralPurpose DBAddition = 5
	// Hyperscale ...
	Hyperscale DBAddition = 6
	// Premium ...
	Premium DBAddition = 7
	// PremiumRS ...
	PremiumRS DBAddition = 8
	// Standard ...
	Standard DBAddition = 9
	// Stretch ...
	Stretch DBAddition = 10
	// System ...
	System DBAddition = 11
	// System2 ...
	System2 DBAddition = 12
	// Web ...
	Web DBAddition = 13
)

// SQLServerProperties contains values needed for adding / updating SQL servers,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#Server
// also wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#ServerProperties
type SQLServerProperties struct {

	// AdministratorLogin - Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string

	// AdministratorLoginPassword - The administrator login password (required for server creation).
	AdministratorLoginPassword *string

	// AllowAzureServicesAccess - allow Azure services and resources to access this server
	AllowAzureServicesAccess bool
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
	Edition DBAddition
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
		Edition: translateDBAddition(properties.Edition),
	}

	return result
}

// translateDBAddition translates enums
func translateDBAddition(in DBAddition) (result sql.DatabaseEdition) {
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
