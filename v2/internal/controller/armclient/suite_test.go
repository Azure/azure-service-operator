/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package armclient_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/controller/testcommon"
)

const DefaultEventuallyTimeout = 3 * time.Minute

var testContext testcommon.TestContext

func setup() error {
	recordReplay := os.Getenv("RECORD_REPLAY") != "0"

	log.Println("Running test setup")

	gomega.SetDefaultEventuallyTimeout(DefaultEventuallyTimeout)
	gomega.SetDefaultEventuallyPollingInterval(5 * time.Second)

	// set global test context
	testContext = testcommon.NewTestContext(testcommon.DefaultTestRegion, recordReplay)

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
