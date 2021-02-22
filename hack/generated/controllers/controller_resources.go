/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"k8s.io/apimachinery/pkg/runtime"

	resources "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.resources/v20200601"
)

func GetKnownTypes() []runtime.Object {
	knownTypes := getKnownTypes()

	knownTypes = append(knownTypes, new(resources.ResourceGroup))

	return knownTypes
}

func CreateScheme() *runtime.Scheme {
	scheme := createScheme()
	_ = resources.AddToScheme(scheme)

	return scheme
}
