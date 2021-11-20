/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/v2/api/batch/v1alpha1api20210101"
	"github.com/Azure/azure-service-operator/v2/api/resources/v1alpha1api20200601"
)

func CreateScheme() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()
	err := v1alpha1api20210101.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}

	err = v1alpha1api20200601.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}

	return scheme, nil
}
