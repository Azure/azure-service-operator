// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package apimgmt

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resources"
	"github.com/Azure/azure-service-operator/pkg/util"
)

// TestCreateAPIMgmtSvc tests creating and delete API Mgmt svcs
func TestCreateAPIMgmtSvc(t *testing.T) {

	// skip this test for now due to length of time constraints, comment out to execute
	t.SkipNow()

	var groupName = config.GenerateGroupName("APIMSTest")
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
		ServiceName:       generateName("apimsvc"),
		Email:             "test@microsoft.com",
		Name:              "test",
	}

	// wait for the service to be created, then only proceed once activated
	_, err = sdk.CreateAPIMgmtSvc()
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("cannot create api management service: %v", err))
		t.FailNow()
	}
	for true {
		time.Sleep(time.Second)
		activated, err := sdk.IsAPIMgmtSvcActivated()
		if err != nil {
			util.PrintAndLog(fmt.Sprintf("error checking for activation: %v", err))
			t.FailNow()
			break
		}
		if activated == true {
			util.PrintAndLog("api management service created")
			break
		}
	}

	// delete the service
	_, err = sdk.DeleteAPIMgmtSvc()
	if err != nil {
		util.PrintAndLog(fmt.Sprintf("cannot delete api management service: %v", err))
		t.FailNow()
	} else {
		util.PrintAndLog("api management service deleted")
	}
}
