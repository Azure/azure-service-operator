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

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
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
	sdk := GoSDKClient{}

	location := "eastus2"
	serverName := generateName("sqlsrvtest")

	// create the server
	sqlServerProperties := SQLServerProperties{
		AdministratorLogin:         to.StringPtr("Moss"),
		AdministratorLoginPassword: to.StringPtr("TheITCrowd_{01}!"),
	}

	// wait for server to be created, then only proceed once activated
	for {
		time.Sleep(time.Second)
		server, err := sdk.CreateOrUpdateSQLServer(ctx, groupName, location, serverName, sqlServerProperties)
		if err == nil {
			if *server.State == "Ready" {
				util.PrintAndLog("sql server ready")
				break
			} else {
				util.PrintAndLog("waiting for sql server to be ready...")
				continue
			}
		} else {
			if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
				util.PrintAndLog("waiting for sql server to be ready...")
				continue
			} else {
				util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v", err))
				t.FailNow()
				break
			}
		}
	}

	// create a DB
	sqlDBProperties := SQLDatabaseProperties{
		DatabaseName: "sqldatabase-sample",
		Edition:      Basic,
	}

	// wait for db to be created, then only proceed once activated
	for {
		time.Sleep(time.Second)
		future, err := sdk.CreateOrUpdateDB(ctx, groupName, location, serverName, sqlDBProperties)
		if err == nil {
			db, err := future.Result(getGoDbClient())
			if err == nil {
				if *db.Status == "Online" {
					util.PrintAndLog("db ready")
					break
				}
			} else {
				util.PrintAndLog("waiting for db to be ready...")
				continue
			}
		} else {
			if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
				util.PrintAndLog("waiting for db to be ready...")
				continue
			} else {
				util.PrintAndLog(fmt.Sprintf("cannot create db: %v", err))
				t.FailNow()
				break
			}
		}
	}

	// create a firewall rule
	util.PrintAndLog("creating firewall rule...")
	_, err = sdk.CreateOrUpdateSQLFirewallRule(ctx, groupName, location, serverName, "test-rule1", "1.1.1.1", "2.2.2.2")
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("cannot create firewall rule: %v", err))
		t.FailNow()
	}
	util.PrintAndLog("firewall rule created")
	time.Sleep(time.Second)

	// create a failover group

	// create secondary SQL server
	// create the Go SDK client with relevant info
	secSrvName := generateName("sqlsrvsecondary")
	secLocation := "westus"

	// create the server
	sqlServerProperties = SQLServerProperties{
		AdministratorLogin:         to.StringPtr("Moss"),
		AdministratorLoginPassword: to.StringPtr("TheITCrowd_{01}!"),
	}

	// wait for server to be created, then only proceed once activated
	for {
		time.Sleep(time.Second)
		server, err := sdk.CreateOrUpdateSQLServer(ctx, groupName, secLocation, secSrvName, sqlServerProperties)
		if err == nil {
			if *server.State == "Ready" {
				util.PrintAndLog("sql server ready")
				break
			} else {
				util.PrintAndLog("waiting for sql server to be ready...")
				continue
			}
		} else {
			if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
				util.PrintAndLog("waiting for sql server to be ready...")
				continue
			} else {
				util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v", err))
				t.FailNow()
				break
			}
		}
	}

	// Initialize struct for failover group
	sqlFailoverGroupProperties := SQLFailoverGroupProperties{
		FailoverPolicy:               sql.Automatic,
		FailoverGracePeriod:          30,
		SecondaryServerName:          secSrvName,
		SecondaryServerResourceGroup: groupName,
		DatabaseList:                 []string{"sqldatabase-sample"},
	}

	failoverGroupName := generateName("failovergroup")
	for {
		time.Sleep(time.Second)
		_, err := sdk.CreateOrUpdateFailoverGroup(ctx, groupName, secLocation, secSrvName, failoverGroupName, sqlFailoverGroupProperties)
		if err == nil {
			util.PrintAndLog(fmt.Sprintf("failover group created successfully %s", failoverGroupName))
			break
		} else {
			if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
				util.PrintAndLog("waiting for failover group to be ready...")
				continue
			} else {
				util.PrintAndLog(fmt.Sprintf("cannot create failovergroup: %v", err))
				t.FailNow()
				break
			}
		}
	}

	// delete firewall rule
	util.PrintAndLog("deleting firewall rule...")
	err = sdk.DeleteSQLFirewallRule(ctx, groupName, location, serverName, "test-rule1")
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("cannot delete firewall rule: %v", err))
		t.FailNow()
	}
	util.PrintAndLog("firewall rule deleted")

	// delete the failover group
	util.PrintAndLog("deleting failover group...")
	response, err := sdk.DeleteFailoverGroup(ctx, groupName, secLocation, secSrvName, failoverGroupName)
	if err == nil {
		if response.StatusCode == 200 {
			util.PrintAndLog("failover group deleted")
		}
	} else {
		util.PrintAndLog(fmt.Sprintf("cannot delete failover group: %v", err))
		t.FailNow()
	}

	// delete the DB
	time.Sleep(time.Second)
	response, err = sdk.DeleteDB(ctx, groupName, secLocation, secSrvName, "sqldatabase-sample")
	if err == nil {
		if response.StatusCode == 200 {
			util.PrintAndLog("db deleted")
		}
	} else {
		util.PrintAndLog(fmt.Sprintf("cannot delete db: %v", err))
		t.FailNow()
	}

	// delete the server
	time.Sleep(time.Second)
	response, err = sdk.DeleteSQLServer(ctx, groupName, location, serverName)
	if err == nil {
		if response.StatusCode == 200 {
			util.PrintAndLog("sql server deleted")
		} else {
			util.PrintAndLog(fmt.Sprintf("cannot delete sql server, code: %v", response.StatusCode))
			t.FailNow()
		}
	} else {
		if !errhelp.IsAsynchronousOperationNotComplete(err) && !errhelp.IsGroupNotFound(err) {
			util.PrintAndLog(fmt.Sprintf("cannot delete sql server: %v", err))
			t.FailNow()
		}
	}

	// delete the secondary server
	time.Sleep(time.Second)
	response, err = sdk.DeleteSQLServer(ctx, groupName, secLocation, secSrvName)
	if err == nil {
		if response.StatusCode == 200 {
			util.PrintAndLog("sql server deleted")
		} else {
			util.PrintAndLog(fmt.Sprintf("cannot delete sql server, code: %v", response.StatusCode))
			t.FailNow()
		}
	} else {
		if !errhelp.IsAsynchronousOperationNotComplete(err) && !errhelp.IsGroupNotFound(err) {
			util.PrintAndLog(fmt.Sprintf("cannot delete sql server: %v", err))
			t.FailNow()
		}
	}

}
