/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/config"
)

func createRealKubeContext() (BaseTestContextFactory, error) {
	kubeConfig, err := ctrl.GetConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to retrieve kubeconfig")
	}

	return func(perTestContext PerTestContext, _ config.Values) (*KubeBaseTestContext, error) {
		return &KubeBaseTestContext{
			PerTestContext: perTestContext,
			KubeConfig:     kubeConfig,
		}, nil
	}, nil
}
