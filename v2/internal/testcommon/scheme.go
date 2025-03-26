/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	batch "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101"
	batchstorage "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101/storage"
	dbforpostgresqlpreview "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20230601preview"
	dbforpostgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20240801"
	dbforpostgresqlstorage "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20240801/storage"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	resourcesstorage "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601/storage"
)

func CreateScheme() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()

	addToScheme := []func(scheme *runtime.Scheme) error{
		batch.AddToScheme,
		batchstorage.AddToScheme,
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
