/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

func createRealKubeContext(perTestContext PerTestContext) (*KubeBaseTestContext, error) {
	config, err := ctrl.GetConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to retrieve kubeconfig")
	}

	return &KubeBaseTestContext{
		PerTestContext: perTestContext,
		KubeConfig:     config,
	}, nil
}
