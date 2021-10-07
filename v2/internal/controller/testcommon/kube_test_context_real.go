/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/internal/controller/config"
)

func createRealKubeContext(perTestContext PerTestContext, _ config.Values) (*KubeBaseTestContext, error) {
	kubeConfig, err := ctrl.GetConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to retrieve kubeconfig")
	}

	return &KubeBaseTestContext{
		PerTestContext: perTestContext,
		KubeConfig:     kubeConfig,
	}, nil
}
