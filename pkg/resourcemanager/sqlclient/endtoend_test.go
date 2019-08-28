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
		ServerName:        generateName("sqlsrvtest"),
		Location:          "eastus2",
	}

	// create the sql server properties struct
	sqlServerProperties := SQLServerProperties{
		AdministratorLogin:         to.StringPtr("Moss"),
		AdministratorLoginPassword: to.StringPtr("TheITCrowd_{01}!"),
	}

	// this firewall rule should fail (the server doesn't exist yet)
	result, err := sdk.CreateOrUpdateSQLFirewallRule("fail rule", "0.0.0.0", "1.2.3.4")
	if result {
		util.PrintAndLog("firewall rule add succeeded, but shouldn't have")
		t.FailNow()
	} else {
		util.PrintAndLog("firewall rule add failed (good)")
	}

	// wait for server to be created, then only proceed once activated
	for {
		time.Sleep(time.Second)
		server, err := sdk.CreateOrUpdateSQLServer(sqlServerProperties)
		if err == nil {
			if *server.State == "Ready" {
				util.PrintAndLog("sql server ready")
				break
			}
		} else {
			if sdk.IsAsyncNotCompleted(err) {
				util.PrintAndLog("waiting for sql server to be ready...")
				continue
			} else {
				util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v", err))
				t.FailNow()
				break
			}
		}
	}

	// this firewall rule should succeed
	result, err = sdk.CreateOrUpdateSQLFirewallRule("succeed rule", "0.0.0.0", "4.3.2.1")
	if result {
		util.PrintAndLog("firewall rule add succeeded")
	} else {
		util.PrintAndLog("firewall rule add failed, but should have succeeded")
		t.FailNow()
	}

	// this firewall rule deletion should succeed
	err = sdk.DeleteSQLFirewallRule("succeed rule")
	if err == nil {
		util.PrintAndLog("firewall rule deletion succeeded")
	} else {
		util.PrintAndLog(fmt.Sprintf("firewall rule deletion failed, but should have succeeded: %v", err))
		t.FailNow()
	}

	// create a DB
	sqlDBProperties := SQLDatabaseProperties{
		DatabaseName: "testDB",
		Edition:      Basic,
	}

	// wait for db to be created, then only proceed once activated
	for {
		time.Sleep(time.Second)
		db, err := sdk.CreateOrUpdateDB(sqlDBProperties)
		if err == nil {
			if *db.Status == "Online" {
				util.PrintAndLog("db ready")
				break
			}
		} else {
			if sdk.IsAsyncNotCompleted(err) {
				util.PrintAndLog("waiting for db to be ready...")
				continue
			} else {
				util.PrintAndLog(fmt.Sprintf("cannot create db: %v", err))
				t.FailNow()
				break
			}
		}
	}

	// delete the DB
	time.Sleep(time.Second)
	response, err := sdk.DeleteDB("testDB")
	if err == nil {
		if response.StatusCode == 200 {
			util.PrintAndLog("db deleted")
		}
	} else {
		util.PrintAndLog(fmt.Sprintf("cannot delete db: %v", err))
		t.FailNow()
	}

	// delete the server
	for {
		time.Sleep(time.Second)
		response, err := sdk.DeleteSQLServer()
		if err == nil {
			if response.StatusCode == 200 {
				util.PrintAndLog("sql server deleted")
				break
			}
		} else {
			if sdk.IsAsyncNotCompleted(err) {
				util.PrintAndLog("waiting for sql server to be deleted...")
				continue
			} else {
				util.PrintAndLog(fmt.Sprintf("cannot delete sql server: %v", err))
				t.FailNow()
				break
			}
		}
	}
}
