// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resources"
	"github.com/Azure/azure-service-operator/pkg/util"
	"github.com/Azure/go-autorest/autorest/to"
)

// TestCreateOrUpdateSQLServer tests creating and delete a SQL server
func TestCreateOrUpdateSQLServer(t *testing.T) {

	var groupName = config.GenerateGroupName("SQLCreateDeleteTest")
	config.SetGroupName(groupName)

	ctx := context.Background()
	defer resources.Cleanup(ctx)

	// create the resource group
	_, err := resources.CreateGroup(ctx, config.GroupName())
	if err != nil {
		util.PrintAndLog(err.Error())
		t.FailNow()
	}

	// create the Go SDK client with relevant info
	sdk := GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        generateName("apimsvc"),
		Location:          "eastus2",
	}

	// create the service
	sqlServerProperties := SQLServerProperties{
		AdministratorLogin:         to.StringPtr("Moss"),
		AdministratorLoginPassword: to.StringPtr("TheITCrowd_{01}!"),
		AllowAzureServicesAccess:   true,
	}
	_, err = CreateOrUpdateSQLServer(sdk, sqlServerProperties)
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v", err))
		t.FailNow()
	}
	util.PrintAndLog("sql server created")

	// create a DB
	sqlDBProperties := SQLDatabaseProperties{
		DatabaseName: "testDB",
		Edition:      Free,
	}
	_, err = CreateOrUpdateDB(sdk, sqlDBProperties)
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("cannot create db: %v", err))
		t.FailNow()
	}
	util.PrintAndLog("db created")

	// wait a minute
	time.Sleep(time.Minute)

	// delete the DB
	_, err = DeleteDB(sdk, "testDB")
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("cannot delete the db: %v", err))
		t.FailNow()
	} else {
		util.PrintAndLog("db deleted")
	}

	// delete the service
	_, err = DeleteSQLServer(sdk)
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("cannot delete the sql server: %v", err))
		t.FailNow()
	} else {
		util.PrintAndLog("sql server deleted")
	}
}
