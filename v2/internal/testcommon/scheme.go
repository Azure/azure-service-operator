/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"k8s.io/apimachinery/pkg/runtime"

	batch "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
)

func CreateScheme() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()
	err := batch.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}

	err = resources.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}

	return scheme, nil
}
