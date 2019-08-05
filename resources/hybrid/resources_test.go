// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package resources

import (
	"context"
	"flag"
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/config"
	"github.com/Azure/azure-service-operator/pkg/util"
)

func setupEnvironment() error {
	err1 := config.ParseEnvironment()
	err2 := config.AddFlags()
	err3 := addLocalConfig()

	for _, err := range []error{err1, err2, err3} {
		if err != nil {
			return err
		}
	}

	flag.Parse()
	return nil
}

func addLocalConfig() error {
	return nil
}

func TestGroupsHybrid(t *testing.T) {
	err := setupEnvironment()
	if err != nil {
		t.Fatalf("could not set up environment: %v\n", err)
	}

	groupName := config.GenerateGroupName("resource-groups-hybrid")
	config.SetGroupName(groupName) // TODO: don't rely on globals

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	defer Cleanup(ctx)

	_, err = CreateGroup(ctx)
	if err != nil {
		util.PrintAndLog(err.Error())
	}
	util.PrintAndLog("resource group created")

	// Output:
	// resource group created
}
