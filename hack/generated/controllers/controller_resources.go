/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

// TODO: Code generate this file
import (
	"k8s.io/apimachinery/pkg/runtime"

	batch "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.batch/v20170901"
	resources "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.resources/v20200601"
	storage "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.storage/v20190401"
)

// KnownTypes defines an array of runtime.Objects to be reconciled, where each
// object in the array will generate a controller. If the concrete type
// implements the owner interface, the generated controllers will inject the
// owned types supplied by the Owns() method of the CRD type. Each controllers
// may directly reconcile a single object, but may indirectly watch
// and reconcile many Owned objects. The singular type is necessary to generically
// produce a reconcile function aware of concrete types, as a closure.

var KnownTypes = []runtime.Object{
	new(batch.BatchAccount),
	// new(batch.BatchAccountsPool),
	new(resources.ResourceGroup),
	new(storage.StorageAccount),
	new(storage.StorageAccountsBlobService),
	new(storage.StorageAccountsBlobServicesContainer),
}
