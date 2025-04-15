/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	"github.com/rotisserie/eris"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

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
	nameConfig *ResourceNameConfig,
) (KubeGlobalContext, error) {
	var err error
	var cbtc BaseTestContextFactory
	cleanup := func() {}

	if useEnvTest {
		cbtc, cleanup = createEnvtestContext()
	} else {
		cbtc, err = createRealKubeContext()
		if err != nil {
			return KubeGlobalContext{}, eris.Wrap(err, "unable to create real Kube context")
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

func AsKubernetesOwner(obj client.Object) *genruntime.KubernetesOwnerReference {
	return &genruntime.KubernetesOwnerReference{
		Name: obj.GetName(),
	}
}

func AsARMIDOwner(id string) *genruntime.KnownResourceReference {
	return &genruntime.KnownResourceReference{
		ARMID: id,
	}
}
