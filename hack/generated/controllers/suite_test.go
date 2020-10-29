/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	resources "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.resources/v20200601"
	"github.com/Azure/k8s-infra/hack/generated/controllers"
	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/testcommon"
)

const (
	TestRegion             = "westus"
	TestNamespace          = "k8s-infra-test-ns"
	DefaultResourceTimeout = 2 * time.Minute
)

var testContext ControllerTestContext
var envtestContext *EnvtestContext

type EnvtestContext struct {
	testenv     envtest.Environment
	manager     ctrl.Manager
	stopManager chan struct{}
}

type ControllerTestContext struct {
	*testcommon.TestContext
	SharedResourceGroup *resources.ResourceGroup
}

func (tc *ControllerTestContext) SharedResourceGroupOwner() genruntime.KnownResourceReference {
	return genruntime.KnownResourceReference{Name: tc.SharedResourceGroup.Name}
}

func setupEnvTest() (*rest.Config, error) {
	envtestContext = &EnvtestContext{
		testenv: envtest.Environment{
			CRDDirectoryPaths: []string{
				"../config/crd/bases/valid",
			},
		},
		stopManager: make(chan struct{}),
	}

	log.Print("Starting envtest")
	config, err := envtestContext.testenv.Start()
	if err != nil {
		return nil, errors.Wrapf(err, "starting envtest environment")
	}

	log.Print("Creating & starting controller-runtime manager")
	mgr, err := ctrl.NewManager(config, ctrl.Options{Scheme: testcommon.CreateScheme()})
	if err != nil {
		return nil, errors.Wrapf(err, "creating controller-runtime manager")
	}

	envtestContext.manager = mgr

	go func() {
		// this blocks until the input chan is closed
		err := envtestContext.manager.Start(envtestContext.stopManager)
		if err != nil {
			log.Fatal(errors.Wrapf(err, "running controller-runtime manager"))
		}
	}()

	log.Print("Creating ARM client")
	armClient, err := armclient.NewAzureTemplateClient()
	if err != nil {
		return nil, errors.Wrapf(err, "creating ARM client")
	}

	log.Print("Registering custom controllers")
	errs := controllers.RegisterAll(
		envtestContext.manager,
		armClient,
		controllers.KnownTypes,
		klogr.New(),
		controller.Options{})

	if errs != nil {
		return nil, errors.Wrapf(kerrors.NewAggregate(errs), "registering reconcilers")
	}

	return config, nil
}

func teardownEnvTest() error {
	if envtestContext != nil {
		log.Print("Stopping controller-runtime manager")
		close(envtestContext.stopManager)

		log.Print("Stopping envtest")
		err := envtestContext.testenv.Stop()
		if err != nil {
			return errors.Wrapf(err, "stopping envtest environment")
		}
	}

	return nil
}

func setup(options Options) error {
	ctx := context.Background()
	log.Println("Running test setup")

	gomega.SetDefaultEventuallyTimeout(DefaultResourceTimeout)
	gomega.SetDefaultEventuallyPollingInterval(5 * time.Second)

	var err error
	var config *rest.Config
	if options.useEnvTest {
		config, err = setupEnvTest()
		if err != nil {
			return errors.Wrapf(err, "setting up envtest")
		}
	} else {
		config, err = ctrl.GetConfig()
		if err != nil {
			return errors.Wrapf(err, "unable to retrieve kubeconfig")
		}
	}

	newCtx, err := testcommon.NewTestContext(
		config,
		TestRegion,
		TestNamespace,
		controllers.ResourceStateAnnotation,
		controllers.ResourceErrorAnnotation)
	if err != nil {
		return err
	}

	err = newCtx.CreateTestNamespace()
	if err != nil {
		return err
	}

	// Create a shared resource group, for tests to use
	sharedResourceGroup := newCtx.NewTestResourceGroup()
	err = newCtx.KubeClient.Create(ctx, sharedResourceGroup)
	if err != nil {
		return errors.Wrapf(err, "creating shared resource group")
	}

	// TODO: Should use AzureName rather than Name once it's always set
	log.Printf("Created shared resource group '%s'\n", sharedResourceGroup.Name)

	// It should be created in Kubernetes
	err = testcommon.WaitFor(ctx, DefaultResourceTimeout, func(waitCtx context.Context) (bool, error) {
		return newCtx.Ensure.Provisioned(waitCtx, sharedResourceGroup)
	})

	if err != nil {
		return errors.Wrapf(err, "waiting for shared resource group")
	}

	log.Print("Done with test setup")

	testContext = ControllerTestContext{newCtx, sharedResourceGroup}

	return nil
}

func teardown() error {
	log.Println("Started common controller test teardown")

	ctx := context.Background()

	// List all of the resource groups
	rgList := &resources.ResourceGroupList{}
	err := testContext.KubeClient.List(ctx, rgList, &client.ListOptions{Namespace: testContext.Namespace})
	if err != nil {
		return errors.Wrap(err, "listing resource groups")
	}

	// Delete any leaked resource groups
	var errs []error

	var resourceGroups []runtime.Object

	for _, rg := range rgList.Items {
		rg := rg // Copy so that we can safely take addr
		resourceGroups = append(resourceGroups, &rg)
		err := testContext.KubeClient.Delete(ctx, &rg)
		if err != nil {
			errs = append(errs, err)
		}
	}
	err = kerrors.NewAggregate(errs)
	if err != nil {
		return err
	}

	// Don't block forever waiting for delete to complete
	err = testcommon.WaitFor(ctx, DefaultResourceTimeout, func(waitCtx context.Context) (bool, error) {
		return testContext.Ensure.AllDeleted(waitCtx, resourceGroups)
	})

	if err != nil {
		return errors.Wrapf(err, "waiting for all resource groups to delete")
	}

	err = teardownEnvTest()
	if err != nil {
		return errors.Wrapf(err, "tearing down envtest")
	}

	log.Println("Finished common controller test teardown")
	return nil
}

func TestMain(m *testing.M) {
	options := getOptions()
	os.Exit(testcommon.SetupTeardownTestMain(
		m,
		true,
		func() error {
			return setup(options)
		},
		teardown))
}

type Options struct {
	useEnvTest bool
}

func getOptions() Options {
	options := Options{}
	flag.BoolVar(&options.useEnvTest, "envtest", false, "Use the envtest package to run tests? If not, a cluster must be configured already in .kubeconfig.")
	flag.Parse()
	return options
}
