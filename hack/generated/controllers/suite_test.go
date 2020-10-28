/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resources "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.resources/v20200601"
	"github.com/Azure/k8s-infra/hack/generated/controllers"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/testcommon"
)

const (
	TestRegion             = "westus"
	TestNamespace          = "k8s-infra-test-ns"
	DefaultResourceTimeout = 2 * time.Minute
)

var testContext ControllerTestContext

type ControllerTestContext struct {
	*testcommon.TestContext
	SharedResourceGroup *resources.ResourceGroup
}

func (tc *ControllerTestContext) SharedResourceGroupOwner() genruntime.KnownResourceReference {
	return genruntime.KnownResourceReference{Name: tc.SharedResourceGroup.Name}
}

func setup() error {
	ctx := context.Background()
	log.Println("Running test setup")

	gomega.SetDefaultEventuallyTimeout(DefaultResourceTimeout)
	gomega.SetDefaultEventuallyPollingInterval(5 * time.Second)

	newCtx, err := testcommon.NewTestContext(
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
	log.Printf("Created shared resource group %s\n", sharedResourceGroup.Name)

	// It should be created in Kubernetes
	err = testcommon.WaitFor(ctx, DefaultResourceTimeout, func(waitCtx context.Context) (bool, error) {
		return newCtx.Ensure.Provisioned(waitCtx, sharedResourceGroup)
	})
	if err != nil {
		return errors.Wrapf(err, "waiting for shared resource group")
	}

	log.Println("Done with test setup")

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

	log.Println("Finished common controller test teardown")
	return nil
}

func TestMain(m *testing.M) {
	os.Exit(testcommon.SetupTeardownTestMain(m, true, setup, teardown))
}

// TODO: Do we need this?
//func PanicRecover(t *testing.T) {
//	if err := recover(); err != nil {
//		t.Logf("caught panic in test: %v", err)
//		t.Logf("stacktrace from panic: \n%s", string(debug.Stack()))
//		t.Fail()
//	}
//}
