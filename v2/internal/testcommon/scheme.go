/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	batch_v20210101 "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101"
	batch_v20210101s "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101/storage"
	batch_v20240701 "github.com/Azure/azure-service-operator/v2/api/batch/v20240701"
	batch_v20240701s "github.com/Azure/azure-service-operator/v2/api/batch/v20240701/storage"
	dbforpostgresqlpreview "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20230601preview"
	dbforpostgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v20250801"
	dbforpostgresqlstorage "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v20250801/storage"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	resourcesstorage "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601/storage"
)

func CreateScheme() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()

	addToScheme := []func(scheme *runtime.Scheme) error{
		batch_v20210101.AddToScheme,
		batch_v20210101s.AddToScheme,
		batch_v20240701.AddToScheme,
		batch_v20240701s.AddToScheme,
		resources.AddToScheme,
		resourcesstorage.AddToScheme,
		dbforpostgresqlpreview.AddToScheme,
		// These dbforpostgresql resources are an example that illustrates
		// https://github.com/Azure/azure-service-operator/issues/4316
		dbforpostgresql.AddToScheme,
		dbforpostgresqlstorage.AddToScheme,
		v1.AddToScheme,
	}

	for _, f := range addToScheme {
		err := f(scheme)
		if err != nil {
			return nil, err
		}
	}

	return scheme, nil
}
