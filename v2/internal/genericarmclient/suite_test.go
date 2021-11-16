/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient_test

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

const DefaultEventuallyTimeout = 3 * time.Minute

var testContext testcommon.TestContext

func setup() error {
	recordReplay := os.Getenv("RECORD_REPLAY") != "0"

	log.Println("Running test setup")

	gomega.SetDefaultEventuallyTimeout(DefaultEventuallyTimeout)
	gomega.SetDefaultEventuallyPollingInterval(5 * time.Second)

	// setup global logger for controller-runtime:
	ctrl.SetLogger(klogr.New())

	nameConfig := testcommon.NewResourceNameConfig(
		testcommon.ResourcePrefix,
		"-",
		6,
		testcommon.ResourceNamerModeRandomBasedOnTestName)

	// set global test context
	testContext = testcommon.NewTestContext(testcommon.DefaultTestRegion, recordReplay, nameConfig)

	log.Println("Done with test setup")

	return nil
}

func teardown() error {
	log.Println("Started common controller test teardown")
	log.Println("Finished common controller test teardown")
	return nil
}

func TestMain(m *testing.M) {
	os.Exit(testcommon.SetupTeardownTestMain(m, setup, teardown))
}
