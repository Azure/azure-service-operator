/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package api

import (
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/v2/internal/controllers"
)

// CreateScheme is a helper method to export 'internal/controllers/controller_resources.go/CreateScheme()' method outside internal package for external use.
func CreateScheme() *runtime.Scheme {
	return controllers.CreateScheme()
}
