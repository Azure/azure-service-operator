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

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"k8s.io/klog/v2/textlogger"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

const DefaultEventuallyTimeout = 3 * time.Minute

var testContext testcommon.TestContext

func setup() error {
	recordReplay := os.Getenv("RECORD_REPLAY") != "0"

	gomega.SetDefaultEventuallyTimeout(DefaultEventuallyTimeout)
	gomega.SetDefaultEventuallyPollingInterval(5 * time.Second)
	format.TruncateThreshold = 4000 // Force a longer truncate threshold

	// setup global logger for controller-runtime:
	cfg := textlogger.NewConfig(textlogger.Verbosity(Debug)) // Use verbose logging in tests
	log := textlogger.NewLogger(cfg)
	ctrl.SetLogger(log)

	log.Info("Running test setup")

	nameConfig := testcommon.NewResourceNameConfig(
		testcommon.ResourcePrefix,
		"-",
		6,
		testcommon.ResourceNamerModeRandomBasedOnTestName)

	// set global test context
	testContext = testcommon.NewTestContext(testcommon.DefaultTestRegion, recordReplay, nameConfig)

	log.Info("Done with test setup")

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
