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

	"github.com/Azure/k8s-infra/hack/generated/pkg/testcommon"
)

const DefaultEventuallyTimeout = 3 * time.Minute

var testContext testcommon.TestContext

func setup(recordReplay bool) error {
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
	recordReplay := os.Getenv("RECORD_REPLAY") != ""
	os.Exit(testcommon.SetupTeardownTestMain(
		m,
		true,
		func() error {
			return setup(recordReplay)
		}, teardown))
}
