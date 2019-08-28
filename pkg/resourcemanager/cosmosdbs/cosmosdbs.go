package cosmosdbs

import (
	"context"
	//"encoding/json"
	//"errors"
	"fmt"
	"log"

	//uuid "github.com/satori/go.uuid"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	//"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

func getCosmosDBClient() documentdb.DatabaseAccountsClient {
	cosmosDBClient := documentdb.NewDatabaseAccountsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	cosmosDBClient.Authorizer = a
	cosmosDBClient.AddToUserAgent(config.UserAgent())
	return cosmosDBClient
}

// CreateCosmosDB creates a new CosmosDB
func CreateCosmosDB(ctx context.Context, groupName string,
	cosmosDBName string,
	location string,
	kind azurev1.CosmosDBKind,
	dbType azurev1.CosmosDBDatabaseAccountOfferType,
	tags map[string]*string) (documentdb.DatabaseAccount, error) {
	cosmosDBClient := getCosmosDBClient()

	log.Println("CosmosDB:CosmosDBName" + cosmosDBName)

	/* Uncomment and update if we should be checking for name exists first
	result, err = cosmosDBClient.CheckNameExists(ctx, cosmosDBName)
	if err != nil {
		return documentdb.DatabaseAccount.{}, err
	}
	result.
	if *result.NameAvailable == false {
		log.Fatalf("storage account not available: %v\n", result.Reason)
		return storage.Account{}, errors.New("storage account not available")
	}*/

	dbKind := documentdb.DatabaseAccountKind(kind)

	sDBType := string(dbType)

	/*
	*   Current state of Locations and CosmosDB properties:
	*   Creating a Database account with CosmosDB requires
	*   that DatabaseAccountCreateUpdateProperties be sent over
	*   and currently we are not reading most of these values in
	*   as part of the Spec for CosmosDB.  We are currently
	*   specifying a single Location as part of a location array
	*   which matches the location set for the overall CosmosDB
	*   instance.  This matches the general behavior of creating
	*   a CosmosDB instance in the portal where the only
	*   geo-relicated region is the sole region the CosmosDB
	*   is created in.
	 */
	locationObj := documentdb.Location{
		ID:               to.StringPtr(fmt.Sprintf("%s-%s", cosmosDBName, location)),
		FailoverPriority: to.Int32Ptr(0),
		LocationName:     to.StringPtr(location),
	}

	locationsArray := []documentdb.Location{
		locationObj,
	}

	createUpdateParams := documentdb.DatabaseAccountCreateUpdateParameters{
		Location: to.StringPtr(location),
		Tags:     tags,
		Name:     &cosmosDBName,
		Kind:     dbKind,
		Type:     to.StringPtr("Microsoft.DocumentDb/databaseAccounts"),
		ID:       &cosmosDBName,
		DatabaseAccountCreateUpdateProperties: &documentdb.DatabaseAccountCreateUpdateProperties{
			DatabaseAccountOfferType:      &sDBType,
			EnableMultipleWriteLocations:  to.BoolPtr(false),
			IsVirtualNetworkFilterEnabled: to.BoolPtr(false),
			Locations:                     &locationsArray,
		},
	}

	log.Println(fmt.Sprintf("creating cosmosDB '%s' in resource group '%s' and location: %v", cosmosDBName, groupName, location))

	future, err := cosmosDBClient.CreateOrUpdate(
		ctx, groupName, cosmosDBName, createUpdateParams)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR creating cosmosDB '%s' in resource group '%s' and location: %v", cosmosDBName, groupName, location))
		log.Println(fmt.Printf("failed to initialize cosmosdb: %v\n", err))
	}
	return future.Result(cosmosDBClient)
}

// DeleteCosmosDB removes the resource group named by env var
func DeleteCosmosDB(ctx context.Context, groupName string, cosmosDBName string) (result documentdb.DatabaseAccountsDeleteFuture, err error) {
	cosmosDBClient := getCosmosDBClient()
	return cosmosDBClient.Delete(ctx, groupName, cosmosDBName)
}

/*  Pre-Refactor
// New generates a new object
func New(cosmosdb *azureV1alpha1.CosmosDB) *Template {
	return &Template{
		CosmosDB: cosmosdb,
	}
}

// Template defines the dynamodb cfts
type Template struct {
	CosmosDB *azureV1alpha1.CosmosDB
}

func (t *Template) CreateDeployment(ctx context.Context, resourceGroupName string) (string, error) {
	deploymentName := uuid.NewV4().String()
	asset, err := template.Asset("cosmosdb.json")
	templateContents := make(map[string]interface{})
	json.Unmarshal(asset, &templateContents)
	params := map[string]interface{}{
		"location": map[string]interface{}{
			"value": t.CosmosDB.Spec.Location,
		},
		"kind": map[string]interface{}{
			"value": t.CosmosDB.Spec.Kind,
		},
		"properties": map[string]interface{}{
			"value": t.CosmosDB.Spec.Properties,
		},
	}

	err = deployment.CreateDeployment(ctx, resourceGroupName, deploymentName, &templateContents, &params)
	return deploymentName, err
}
*/
