/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package multitenant_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

const (
	DefaultResourceTimeout = 10 * time.Minute
)

var globalTestContext testcommon.KubeGlobalContext

func setup() error {
	log.Println("Running test setup")

	// Note: These are set just so we have somewhat reasonable defaults. Almost all
	// usage of Eventually is done through the testContext wrapper which manages its
	// own timeouts.
	gomega.SetDefaultEventuallyTimeout(DefaultResourceTimeout)
	gomega.SetDefaultEventuallyPollingInterval(5 * time.Second)

	// setup global logger for controller-runtime:
	ctrl.SetLogger(klogr.New())

	nameConfig := testcommon.NewResourceNameConfig(
		testcommon.LiveResourcePrefix,
		"-",
		6,
		testcommon.ResourceNamerModeRandom)

	// set global context var
	newGlobalTestContext, err := testcommon.NewKubeContext(
		false,
		false,
		testcommon.DefaultTestRegion,
		nameConfig)
	if err != nil {
		return err
	}

	log.Print("Done with test setup")
	globalTestContext = newGlobalTestContext
	return nil
}

func teardown() error {
	return globalTestContext.Cleanup()
}

func TestMain(m *testing.M) {
	os.Exit(testcommon.SetupTeardownTestMain(m, setup, teardown))
}
