/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type KubeGlobalContext struct {
	TestContext

	createBaseTestContext BaseTestContextFactory
	cleanup               context.CancelFunc
}

type BaseTestContextFactory func(PerTestContext, config.Values) (*KubeBaseTestContext, error)

func NewKubeContext(
	useEnvTest bool,
	recordReplay bool,
	region string,
	nameConfig *ResourceNameConfig) (KubeGlobalContext, error) {

	var err error
	var cbtc BaseTestContextFactory
	var cleanup func() = func() {}

	if useEnvTest {
		cbtc, cleanup = createEnvtestContext()
	} else {
		cbtc, err = createRealKubeContext()
		if err != nil {
			return KubeGlobalContext{}, errors.Wrap(err, "unable to create real Kube context")
		}
	}

	return KubeGlobalContext{
		TestContext:           NewTestContext(region, recordReplay, nameConfig),
		createBaseTestContext: cbtc,
		cleanup:               cleanup,
	}, nil
}

func (ctx KubeGlobalContext) Cleanup() error {
	if ctx.cleanup != nil {
		ctx.cleanup()
	}
	return nil
}

type KubeBaseTestContext struct {
	PerTestContext

	KubeConfig *rest.Config
}

func AsOwner(obj client.Object) *genruntime.KnownResourceReference {
	return &genruntime.KnownResourceReference{
		Name: obj.GetName(),
	}
}
