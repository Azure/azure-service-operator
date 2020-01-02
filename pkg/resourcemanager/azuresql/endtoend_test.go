// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package azuresqlshared

import (
	"fmt"

	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	azuresqldb "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	azuresqlfailovergroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
	azuresqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfirewallrule"
	azuresqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	azuresqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/util"
	"github.com/Azure/go-autorest/autorest/to"
	ctrl "sigs.k8s.io/controller-runtime"
)

type TestContext struct {
	sqlServerManager        azuresqlserver.SqlServerManager
	sqlDbManager            azuresqldb.SqlDbManager
	sqlFirewallRuleManager  azuresqlfirewallrule.SqlFirewallRuleManager
	sqlFailoverGroupManager azuresqlfailovergroup.SqlFailoverGroupManager
	sqlUserManager          azuresqluser.SqlUserManager

	timeout time.Duration //timeout in mins
}

// TestCreateOrUpdateSQLServer tests creating and delete a SQL server
func TestCreateOrUpdateSQLServer(t *testing.T) {

	location := config.DefaultLocation()
	sqlServerManager := azuresqlserver.NewAzureSqlServerManager(ctrl.Log.WithName("sqlservermanager").WithName("AzureSqlServer"))
	sqlDbManager := azuresqldb.NewAzureSqlDbManager(ctrl.Log.WithName("sqldbmanager").WithName("AzureSqlDb"))
	sqlFirewallRuleManager := azuresqlfirewallrule.NewAzureSqlFirewallRuleManager(ctrl.Log.WithName("sqlfirewallrulemanager").WithName("AzureSqlFirewallRule"))
	sqlFailoverGroupManager := azuresqlfailovergroup.NewAzureSqlFailoverGroupManager(ctrl.Log.WithName("sqlfailovergroupmanager").WithName("AzureSqlFailoverGroup"))
	sqlUserManager := azuresqluser.NewAzureSqlUserManager(ctrl.Log.WithName("sqlusermanager").WithName("AzureSqlUser"))

	tc = TestContext{
		sqlServerManager:        sqlServerManager,
		sqlDbManager:            sqlDbManager,
		sqlFirewallRuleManager:  sqlFirewallRuleManager,
		sqlFailoverGroupManager: sqlFailoverGroupManager,
		sqlUserManager:          sqlUserManager,
		timeout:                 20 * time.Minute, // timeout in mins
	}

	ignorableErrors := []string{errhelp.AsyncOpIncompleteError}

	serverName := generateName("sqlsrvtest")

	// create the server
	sqlServerProperties := azuresqlshared.SQLServerProperties{
		AdministratorLogin:         to.StringPtr("Moss"),
		AdministratorLoginPassword: to.StringPtr("TheITCrowd_{01}!"),
	}

	// wait for server to be created, then only proceed once activated
	server, err := tc.sqlServerManager.CreateOrUpdateSQLServer(ctx, groupName, location, serverName, sqlServerProperties)
	azerr := errhelp.NewAzureErrorAzureError(err)
	if err != nil && !helpers.ContainsString(ignorableErrors, azerr.Type) {
		util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v", err))
		t.FailNow()
	}
	then := time.Now()
	for {
		if time.Since(then) > tc.timeout {
			util.PrintAndLog("test timed out")
			t.FailNow()
		}

		time.Sleep(time.Second)

		server, err = tc.sqlServerManager.GetServer(ctx, groupName, serverName)
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
			}
		}
	}

	// create a DB
	sqlDBProperties := azuresqlshared.SQLDatabaseProperties{
		DatabaseName: "sqldatabase-sample",
		Edition:      azuresqlshared.Basic,
	}

	// wait for db to be created, then only proceed once activated
	future, err := tc.sqlDbManager.CreateOrUpdateDB(ctx, groupName, location, serverName, sqlDBProperties)
	then = time.Now()
	for {
		if time.Since(then) > tc.timeout {
			util.PrintAndLog("test timed out")
			t.FailNow()
		}
		time.Sleep(time.Second)
		if err == nil {
			db, err := future.Result(azuresqlshared.GetGoDbClient())
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
	secLocation := "westus2"

	// create the server
	sqlServerProperties = azuresqlshared.SQLServerProperties{
		AdministratorLogin:         to.StringPtr("Moss"),
		AdministratorLoginPassword: to.StringPtr("TheITCrowd_{01}!"),
	}

	// wait for server to be created, then only proceed once activated
	server, err = tc.sqlServerManager.CreateOrUpdateSQLServer(ctx, groupName, secLocation, secSrvName, sqlServerProperties)
	azerr = errhelp.NewAzureErrorAzureError(err)

	if err != nil && !helpers.ContainsString(ignorableErrors, azerr.Type) {
		util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v", err))
		t.FailNow()
	}
	then = time.Now()
	for {
		if time.Since(then) > tc.timeout {
			util.PrintAndLog("test timed out")
			t.FailNow()
		}
		time.Sleep(time.Second)

		server, err := tc.sqlServerManager.GetServer(ctx, groupName, secSrvName)
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
			}
		}
	}

	// Initialize struct for failover group
	sqlFailoverGroupProperties := azuresqlshared.SQLFailoverGroupProperties{
		FailoverPolicy:               azuresqlshared.Automatic,
		FailoverGracePeriod:          30,
		SecondaryServerName:          secSrvName,
		SecondaryServerResourceGroup: groupName,
		DatabaseList:                 []string{"sqldatabase-sample"},
	}

	failoverGroupName := generateName("failovergroup")
	fogfuture, err := tc.sqlFailoverGroupManager.CreateOrUpdateFailoverGroup(ctx, groupName, serverName, failoverGroupName, sqlFailoverGroupProperties)
	then = time.Now()
	for {
		if time.Since(then) > tc.timeout {
			util.PrintAndLog("test timed out")
			t.FailNow()
		}
		time.Sleep(time.Second)
		if err == nil {
			_, err := fogfuture.Result(azuresqlshared.GetGoFailoverGroupsClient())
			if err == nil {
				util.PrintAndLog("Failover group ready")
				break

			} else {
				util.PrintAndLog("waiting for failover group to be ready...")
				continue
			}
		} else {
			if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
				util.PrintAndLog("waiting for failover group to be ready...")
				continue
			} else {
				util.PrintAndLog(fmt.Sprintf("cannot create failover group: %v", err))
				t.FailNow()
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
		if !errhelp.IsAsynchronousOperationNotComplete(err) && !errhelp.IsGroupNotFound(err) {
			util.PrintAndLog(fmt.Sprintf("cannot delete failover group: %v", err))
			t.FailNow()
		}
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
		if !errhelp.IsAsynchronousOperationNotComplete(err) && !errhelp.IsGroupNotFound(err) {
			util.PrintAndLog(fmt.Sprintf("cannot delete db: %v", err))
			t.FailNow()
		}
	}

	// delete the server
	util.PrintAndLog("deleting server...")
	time.Sleep(time.Second)
	response, err = tc.sqlServerManager.DeleteSQLServer(ctx, groupName, serverName)
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
	response, err = tc.sqlServerManager.DeleteSQLServer(ctx, groupName, secSrvName)
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
