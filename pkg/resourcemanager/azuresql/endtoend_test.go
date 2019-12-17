// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package azuresql

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resources"
	azuresqldb "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	azuresqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	azuresqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	azuresqlfailovergroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
	azuresqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfirewallrule"

	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/util"
	"github.com/Azure/go-autorest/autorest/to"
	ctrl "sigs.k8s.io/controller-runtime"
)

type TestContext struct {
	SqlServerManager        azuresqlserver.SqlServerManager
	sqlDbManager            azuresqldb.SqlDbManager
	sqlFirewallRuleManager  azuresqlfirewaullrule.SqlFirewallRuleManager
	sqlFailoverGroupManager azuresqlfailovergroup.SqlFailoverGroupManager
	sqlUserManager          azuresqluser.SqlUserManager
}

var tc TestContext

// TestCreateOrUpdateSQLServer tests creating and delete a SQL server
func TestCreateOrUpdateSQLServer(t *testing.T) {

	sqlServerManager := NewAzureSqlServerManager(ctrl.Log.WithName("sqlservermanager").WithName("AzureSqlServer"))
	sqlDbManager := NewAzureSqlDbManager(ctrl.Log.WithName("sqldbmanager").WithName("AzureSqlDb"))
	sqlFirewallRuleManager := NewAzureSqlFirewallRuleManager(ctrl.Log.WithName("sqlfirewallrulemanager").WithName("AzureSqlFirewallRule"))
	sqlFailoverGroupManager := NewAzureSqlFailoverGroupManager(ctrl.Log.WithName("sqlfailovergroupmanager").WithName("AzureSqlFailoverGroup"))
	sqlUserManager := NewAzureSqlUserManager(ctrl.Log.WithName("sqlusermanager").WithName("AzureSqlUser"))

	tc = TestContext{
		SqlServerManager:        sqlServerManager,
		sqlDbManager:            sqlDbManager,
		sqlFirewallRuleManager:  sqlFirewallRuleManager,
		sqlFailoverGroupManager: sqlFailoverGroupManager,
		sqlUserManager:          sqlUserManager,
	}

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

	location := config.DefaultLocation()
	serverName := generateName("sqlsrvtest")

	// create the server
	sqlServerProperties := SQLServerProperties{
		AdministratorLogin:         to.StringPtr("Moss"),
		AdministratorLoginPassword: to.StringPtr("TheITCrowd_{01}!"),
	}

	ignorableErrors := []string{errhelp.AsyncOpIncompleteError}

	// wait for server to be created, then only proceed once activated
	server, err := tc.SqlServerManager.CreateOrUpdateSQLServer(ctx, groupName, location, serverName, sqlServerProperties)
	azerr := errhelp.NewAzureErrorAzureError(err)
	if err != nil && !helpers.ContainsString(ignorableErrors, azerr.Type) {
		util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v", err))
		t.FailNow()
	}
	for {
		time.Sleep(time.Second)

		server, err = tc.SqlServerManager.GetServer(ctx, groupName, serverName)
		if err == nil {
			if *server.State == "Ready" {
				util.PrintAndLog("sql server ready")
				break
			} else {
				util.PrintAndLog("waiting for sql server to be ready...")
				continue
			}
		} else {
			if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) || errhelp.IsResourceNotFound(err) {
				util.PrintAndLog("waiting for sql server to be ready...")
				continue
			} else {
				util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v", err))
				util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v", serverName))

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
	future, err := tc.sqlDbManager.CreateOrUpdateDB(ctx, groupName, location, serverName, sqlDBProperties)
	for {
		time.Sleep(time.Second)
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
	_, err = tc.sqlFirewallRuleManager.CreateOrUpdateSQLFirewallRule(ctx, groupName, serverName, "test-rule1", "1.1.1.1", "2.2.2.2")
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
	server, err = tc.SqlServerManager.CreateOrUpdateSQLServer(ctx, groupName, secLocation, secSrvName, sqlServerProperties)
	azerr = errhelp.NewAzureErrorAzureError(err)
	if err != nil && !helpers.ContainsString(ignorableErrors, azerr.Type) {
		util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v", err))
		t.FailNow()
	}
	for {
		time.Sleep(time.Second)

		server, err := tc.SqlServerManager.GetServer(ctx, groupName, secSrvName)
		if err == nil {
			if *server.State == "Ready" {
				util.PrintAndLog("sql server ready")
				break
			} else {
				util.PrintAndLog("waiting for sql server to be ready...")
				continue
			}
		} else {
			if errhelp.IsAsynchronousOperationNotComplete(err) ||
				errhelp.IsGroupNotFound(err) ||
				errhelp.IsResourceNotFound(err) {
				util.PrintAndLog("waiting for sql server to be ready...")
				continue
			} else {
				util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v err: %v", secSrvName, err))
				t.FailNow()
				break
			}
		}
	}

	// Initialize struct for failover group
	sqlFailoverGroupProperties := SQLFailoverGroupProperties{
		FailoverPolicy:               Automatic,
		FailoverGracePeriod:          30,
		SecondaryServerName:          secSrvName,
		SecondaryServerResourceGroup: groupName,
		DatabaseList:                 []string{"sqldatabase-sample"},
	}

	failoverGroupName := generateName("failovergroup")
	_, err = tc.sqlFailoverGroupManager.CreateOrUpdateFailoverGroup(ctx, groupName, serverName, failoverGroupName, sqlFailoverGroupProperties)
	for {
		time.Sleep(time.Second)
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
	err = tc.sqlFirewallRuleManager.DeleteSQLFirewallRule(ctx, groupName, serverName, "test-rule1")
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("cannot delete firewall rule: %v", err))
		t.FailNow()
	}
	util.PrintAndLog("firewall rule deleted")

	// delete the failover group
	util.PrintAndLog("deleting failover group...")
	response, err := tc.sqlFailoverGroupManager.DeleteFailoverGroup(ctx, groupName, serverName, failoverGroupName)
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
	util.PrintAndLog("deleting db...")
	response, err = tc.sqlDbManager.DeleteDB(ctx, groupName, secSrvName, "sqldatabase-sample")
	if err == nil {
		if response.StatusCode == 200 {
			util.PrintAndLog("db deleted")
		}
	} else {
		util.PrintAndLog(fmt.Sprintf("cannot delete db: %v", err))
		t.FailNow()
	}

	// delete the server
	util.PrintAndLog("deleting server...")
	time.Sleep(time.Second)
	response, err = tc.SqlServerManager.DeleteSQLServer(ctx, groupName, serverName)
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
	util.PrintAndLog("deleting second server...")
	time.Sleep(time.Second)
	response, err = tc.SqlServerManager.DeleteSQLServer(ctx, groupName, secSrvName)
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

	// delete the resource group
	util.PrintAndLog("deleting resource group...")
	time.Sleep(time.Second)
	_, err = resources.DeleteGroup(ctx, config.GroupName())
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("Cannot delete resourcegroup: %v", err))
		t.FailNow()
	}

}
