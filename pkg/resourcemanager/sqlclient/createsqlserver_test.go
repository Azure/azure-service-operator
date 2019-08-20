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
)

// TestCreateOrUpdateSQLServer tests creating and delete a SQL server
func TestCreateOrUpdateSQLServer(t *testing.T) {

	// skip this test for now due to length of time constraints, comment out to execute
	// t.SkipNow()

	var groupName = config.GenerateGroupName("SQLCreateDeleteTest")
	config.SetGroupName(groupName)

	ctx := context.Background()
	defer resources.Cleanup(ctx)

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

	// create the SQLServerProperties struct
	sqlServerProperties := SQLServerProperties{}

	// create the service
	_, err = CreateOrUpdateSQLServer(sdk, sqlServerProperties)
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("cannot create sql server: %v", err))
		t.FailNow()
	}
	util.PrintAndLog("sql server created")

	// sleep for a minute
	time.Sleep(time.Second * 60)

	// delete the service
	_, err = DeleteSQLServer(sdk)
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("cannot delete the sql server: %v", err))
		t.FailNow()
	} else {
		util.PrintAndLog("sql server deleted")
	}
}
