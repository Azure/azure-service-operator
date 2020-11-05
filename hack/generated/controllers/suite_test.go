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

	"github.com/Azure/k8s-infra/hack/generated/controllers"
	"github.com/Azure/k8s-infra/hack/generated/pkg/testcommon"
)

const (
	TestNamespace          = "k8s-infra-test-ns"
	DefaultResourceTimeout = 2 * time.Minute
)

var testContext testcommon.KubeGlobalContext

func setup(options Options) {
	log.Println("Running test setup")

	gomega.SetDefaultEventuallyTimeout(DefaultResourceTimeout)
	gomega.SetDefaultEventuallyPollingInterval(5 * time.Second)

	// set global context var
	testContext = testcommon.NewKubeContext(
		options.useEnvTest,
		options.recordReplay,
		TestNamespace,
		testcommon.DefaultTestRegion,
		controllers.ResourceStateAnnotation,
		controllers.ResourceErrorAnnotation)

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
		useEnvTest:   os.Getenv("ENVTEST") != "",
		recordReplay: os.Getenv("RECORD_REPLAY") != "",
	}
}
