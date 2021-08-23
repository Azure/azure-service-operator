/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

const (
	TestNamespace          = "aso-test-ns"
	DefaultResourceTimeout = 10 * time.Minute
)

var globalTestContext testcommon.KubeGlobalContext

func setup(options Options) {
	log.Println("Running test setup")

	// Note: These are set just so we have somewhat reasonable defaults. Almost all
	// usage of Eventually is done through the testContext wrapper which understands
	// replay vs record modes and passes a different timeout and polling interval for each,
	// meaning that there are very few instances where these timeouts are actually used.
	gomega.SetDefaultEventuallyTimeout(DefaultResourceTimeout)
	gomega.SetDefaultEventuallyPollingInterval(5 * time.Second)

	// set global context var
	globalTestContext = testcommon.NewKubeContext(
		options.useEnvTest,
		options.recordReplay,
		TestNamespace,
		testcommon.DefaultTestRegion)

	log.Print("Done with test setup")
}

func TestMain(m *testing.M) {
	options := getOptions()
	os.Exit(testcommon.SetupTeardownTestMain(
		m,
		true,
		func() error {
			setup(options)
			return nil
		},
		func() error {
			return nil
		}))
}

type Options struct {
	useEnvTest   bool
	recordReplay bool
}

func getOptions() Options {
	return Options{
		useEnvTest:   os.Getenv("ENVTEST") != "0",
		recordReplay: os.Getenv("RECORD_REPLAY") != "0",
	}
}
