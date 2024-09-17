/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

// Note: This is following the pattern set up in v2/internal/controllers/suite_test.go and should be kept mostly
// in sync with that

const (
	DefaultResourceTimeout = 10 * time.Minute
)

var globalTestContext testcommon.KubeGlobalContext

func setup() error {
	options := getOptions()
	log.Println("Running test setup")

	// Note: These are set just so we have somewhat reasonable defaults. Almost all
	// usage of Eventually is done through the testContext wrapper which understands
	// replay vs record modes and passes a different timeout and polling interval for each,
	// meaning that there are very few instances where these timeouts are actually used.
	gomega.SetDefaultEventuallyTimeout(DefaultResourceTimeout)
	gomega.SetDefaultEventuallyPollingInterval(5 * time.Second)

	format.TruncateThreshold = 4000 // Force a longer truncate threshold

	// If you need to debug envtest setup/teardown,
	// set a global logger for controller-runtime:
	//
	// import (ctrl "sigs.k8s.io/controller-runtime")
	// cfg := textlogger.NewConfig(textlogger.Verbosity(Debug)) // Use verbose logging in tests
	// log := textlogger.NewLogger(cfg)
	// ctrl.SetLogger(log)

	nameConfig := testcommon.NewResourceNameConfig(
		testcommon.ResourcePrefix,
		"-",
		6,
		testcommon.ResourceNamerModeRandomBasedOnTestName)

	// set global context var
	newGlobalTestContext, err := testcommon.NewKubeContext(
		options.useEnvTest,
		options.recordReplay,
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
