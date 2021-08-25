/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resources "github.com/Azure/azure-service-operator/hack/generated/apis/microsoft.resources/v1alpha1api20200601"
)

func GetKnownStorageTypes() []client.Object {
	knownTypes := getKnownStorageTypes()

	knownTypes = append(knownTypes, new(resources.ResourceGroup))

	return knownTypes
}

func GetKnownTypes() []client.Object {
	knownTypes := getKnownTypes()

	knownTypes = append(knownTypes, new(resources.ResourceGroup))

	return knownTypes
}

func CreateScheme() *runtime.Scheme {
	scheme := createScheme()
	_ = resources.AddToScheme(scheme)

	return scheme
}
