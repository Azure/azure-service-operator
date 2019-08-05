// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package resources

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/config"
)

const (
	CI_KEY_NAME = "TRAVIS"
)

func TestMain(m *testing.M) {
	err := setupEnvironment()
	if err != nil {
		log.Fatalf("could not set up environment: %v\n", err)
	}

	os.Exit(m.Run())
}

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

func TestGroups(t *testing.T) {
	groupName := config.GenerateGroupName("Groups")
	config.SetGroupName(groupName) // TODO: don't rely on globals

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	defer Cleanup(ctx)

	var err error
	_, err = CreateGroup(ctx, config.GroupName())
	if err != nil {
		t.Fatalf("failed to create group: %v\n", err.Error())
	}
	t.Logf("created group: %s\n", config.GroupName())
}

func TestGroupsWithAuthFile(t *testing.T) {
	if _, is_ci := os.LookupEnv(CI_KEY_NAME); is_ci == true {
		t.Skipf("skipping auth file test in CI")
	}
	groupName := config.GenerateGroupName("resource-groups-authfile")
	config.SetGroupName(groupName) // TODO: don't rely on globals
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	defer Cleanup(ctx)

	var err error
	_, err = CreateGroupWithAuthFile(ctx, config.GroupName())
	if err != nil {
		t.Fatalf("failed to create group: %v\n", err.Error())
	}
	t.Logf("created group with auth file: %s\n", config.GroupName())
}
