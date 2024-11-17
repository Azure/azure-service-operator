/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package multitenant_test

import (
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

const (
	DefaultResourceTimeout = 10 * time.Minute
)

var globalTestContext testcommon.KubeGlobalContext

func setup() error {
	// Note: These are set just so we have somewhat reasonable defaults. Almost all
	// usage of Eventually is done through the testContext wrapper which manages its
	// own timeouts.
	gomega.SetDefaultEventuallyTimeout(DefaultResourceTimeout)
	gomega.SetDefaultEventuallyPollingInterval(5 * time.Second)

	format.TruncateThreshold = 4000 // Force a longer truncate threshold

	// setup global logger for controller-runtime:
	cfg := textlogger.NewConfig(textlogger.Verbosity(Debug)) // Use verbose logging in tests
	log := textlogger.NewLogger(cfg)
	ctrl.SetLogger(log)

	log.Info("Running test setup")

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

	log.Info("Done with test setup")
	globalTestContext = newGlobalTestContext
	return nil
}

func teardown() error {
	return globalTestContext.Cleanup()
}

func TestMain(m *testing.M) {
	os.Exit(testcommon.SetupTeardownTestMain(m, setup, teardown))
}
