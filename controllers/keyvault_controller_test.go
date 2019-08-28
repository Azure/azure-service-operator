/*
Copyright 2019 microsoft.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package controllers

import (
	"context"
	"fmt"
	"testing"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"k8s.io/apimachinery/pkg/types"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"


)

func TestKeyVault(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	const timeout = time.Second * 240

	resourcemanagerconfig.LoadSettings()

	keyVaultName := "t-kv-dev-sample"
	keyVaultLocation := "westus"
	resourceGroupName := "t-rg-dev-controller"


	// Create the Keyvault object and expect the Reconcile to be created
	keyVaultInstance := &azurev1.KeyVault {
		ObjectMeta: metav1.ObjectMeta{
			Name: keyVaultName,

		},
		Spec: azurev1.KeyVaultSpec {
			Location: keyVaultLocation,
			ResourceGroupName: resourceGroupName,
		},
	}

	err := k8sClient.Create(context.Background(), keyVaultInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	// prep query for get (?)
	keyVaultNamespacedName := types.NamespacedName{Name: keyVaultInstance.Name, Namespace: "default"}

	// wait until resource is provisioned
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), keyVaultNamespacedName, keyVaultInstance)
		return keyVaultInstance.Status.Provisioned == true
	}, tcfg.Timeout,
	).Should(BeTrue())

	// verify keyvault exists in Azure
	Eventually(func() bool {
		result, _ := resourcegroupsresourcemanager.CheckExistence(context.Background(), keyVaultInstance.Name)
		return result.Response.StatusCode == 404
	}, tcfg.Timeout, tcfg.Poll,
	).Should(BeTrue())

	// delete keyvault
	k8sClient.Delete(context.Background(), keyVaultInstance)

	// verify its gone from kubernetes
	Eventually(func() bool {
		err := k8sClient.Get(context.Background(), keyVaultNamespacedName, keyVaultGroupInstance)
		if err == nil {
			err = fmt.Errorf("")
		}
		return strings.Contains(err.Error(), "not found")
	}, tcfg.Timeout,
	).Should(BeTrue())	

	// confirm resource is gone from Azure
	Eventually(func() bool {
		result, _ := resourcegroupsresourcemanager.CheckExistence(context.Background(), keyVaultInstance.Name)
		return result.Response.StatusCode == 404
	}, tcfg.Timeout, tcfg.Poll,
	).Should(BeTrue())

}