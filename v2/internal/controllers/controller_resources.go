/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1alpha1api20200601"
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

//GetResourceExtensions returns a map between resource and resource extension
func GetResourceExtensions() map[string]genruntime.ResourceExtension {
	extensionMapping := make(map[string]genruntime.ResourceExtension)

	for _, extension := range getResourceExtensions() {
		for _, resource := range extension.GetExtendedResources() {
			extensionMapping[resource.Owner().String()] = extension
		}
	}
	return extensionMapping
}
