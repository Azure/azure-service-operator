// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlshared

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/Azure/azure-service-operator/api/v1beta1"
)

// SQLServerProperties contains values needed for adding / updating SQL servers,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql#Server
// also wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql#ServerProperties
type SQLServerProperties struct {

	// AdministratorLogin - Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string

	// AdministratorLoginPassword - The administrator login password (required for server creation).
	AdministratorLoginPassword *string
}

// SQLDatabaseProperties contains values needed for adding / updating SQL servers,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql#Database
// also wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql#DatabaseProperties
type SQLDatabaseProperties struct {

	// DatabaseName is the name of the database
	DatabaseName string

	// MaxSize - The max size of the database
	MaxSize *resource.Quantity

	// ElasticPoolID - The resource identifier of the elastic pool containing this database.
	ElasticPoolID string

	// Sku - The database SKU.
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
	Sku *Sku
}

// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql#Sku
type Sku struct {
	// Name - The name of the SKU, typically, a letter + Number code, e.g. P3.
	Name string
	// Tier - The tier or edition of the particular SKU, e.g. Basic, Premium.
	Tier string
	// Size - Size of the particular SKU
	Size string
	// Family - If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family string
	// Capacity - Capacity of the particular SKU.
	Capacity *int32
}

type SQLDatabaseBackupLongTermRetentionPolicy struct {
	WeeklyRetention  string
	MonthlyRetention string
	YearlyRetention  string
	WeekOfYear       int32
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

	var maxSizeBytes *int64
	if properties.MaxSize != nil {
		maxSizeBytes = to.Int64Ptr(properties.MaxSize.Value())
	}

	var elasticPoolID *string
	if properties.ElasticPoolID != "" {
		elasticPoolID = &properties.ElasticPoolID
	}

	result = sql.DatabaseProperties{
		MaxSizeBytes:  maxSizeBytes,
		ElasticPoolID: elasticPoolID,
	}

	return result
}

// SQLDatabasePropertiesToDatabase translates SQLDatabaseProperties to DatabaseProperties
func SQLDatabasePropertiesToSku(properties SQLDatabaseProperties) (result *sql.Sku) {

	if properties.Sku == nil {
		return nil
	}

	result = &sql.Sku{
		Name:     &properties.Sku.Name,
		Tier:     &properties.Sku.Tier,
		Size:     &properties.Sku.Size,
		Family:   &properties.Sku.Family,
		Capacity: properties.Sku.Capacity,
	}

	return result
}

func MergeDBEditionAndSku(in v1beta1.DBEdition, sku *v1beta1.SqlDatabaseSku) (*Sku, error) {

	// There are 3 cases:
	// - Customer just sets Spec.Sku - Spec.Sku is sent as is to ARM.
	// - Customer just sets Spec.Edition - Edition is converted to Spec.Sku and that is sent to ARM
	// - Customer sets both Spec.Sku and Spec.Edition - Spec.Edition is ignored and Spec.Sku is used

	if sku != nil {
		return &Sku{
			Name:     sku.Name,
			Tier:     sku.Tier,
			Size:     sku.Size,
			Family:   sku.Family,
			Capacity: sku.Capacity,
		}, nil
	}

	// These defaults were determined based on what SQL defaults to in the 2015-03-01-preview API.
	// The way this works is that for editions that were DTU based, Edition maps to name: Edition, tier: Edition,
	// so for example edition: "Basic" became name: "Basic", tier: "Basic". For VCore offerings, the name can't actually
	// be derived from just the old edition, so we had to hardcode some actual values (the defaults which SQL happens
	// to use - i.e. HS_Gen5_2 is the default when you deploy a 2015-03-01-preview SQL database with Edition: "Hyperscale").
	// There is some risk that if SQL changes their defaults we'll be out of sync, but given we're planning to remove
	// edition in a future API version and customers have full access to SKU to specify exactly what they want, that
	// shouldn't really be a problem.
	switch in {
	case v1beta1.SqlDatabaseEditionBasic:
		return &Sku{Name: "Basic", Tier: "Basic"}, nil
	case v1beta1.SqlDatabaseEditionDataWarehouse:
		return &Sku{Name: "DataWarehouse", Tier: "DataWarehouse"}, nil
	case v1beta1.SqlDatabaseEditionFree:
		return &Sku{Name: "Free", Tier: "Free"}, nil
	case v1beta1.SqlDatabaseEditionGeneralPurpose:
		return &Sku{Name: "GP_Gen5_2", Tier: "GeneralPurpose"}, nil
	case v1beta1.SqlDatabaseEditionHyperscale:
		return &Sku{Name: "HS_Gen5_2", Tier: "Hyperscale"}, nil
	case v1beta1.SqlDatabaseEditionPremium:
		return &Sku{Name: "Premium", Tier: "Premium"}, nil
	case v1beta1.SqlDatabaseEditionStandard:
		return &Sku{Name: "Standard", Tier: "Standard"}, nil
	case v1beta1.SqlDatabaseEditionStretch:
		return &Sku{Name: "Stretch", Tier: "Stretch"}, nil
	default:
		// There are some numbers that aren't supported -- that's because SQL actually doesn't support them anymore.
		// Since they don't actually work, we're failing fast here rather than trying to come up with a conversion for them.
		return nil, fmt.Errorf("unsupported sql edition: %d", in)
	}
}

// TranslateFailoverPolicy translates the enum
func TranslateFailoverPolicy(in v1beta1.ReadWriteEndpointFailoverPolicy) (sql.ReadWriteEndpointFailoverPolicy, error) {
	var result sql.ReadWriteEndpointFailoverPolicy

	switch in {
	case v1beta1.FailoverPolicyAutomatic:
		result = sql.Automatic
	case v1beta1.FailoverPolicyManual:
		result = sql.Manual
	default:
		return result, errors.Errorf("couldn't translate failover policy %s", in)
	}

	return result, nil
}
