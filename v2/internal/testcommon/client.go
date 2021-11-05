/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func CreateClient(scheme *runtime.Scheme) client.Client {
	testClient := fake.NewClientBuilder().WithScheme(scheme).Build()
	return testClient
}
