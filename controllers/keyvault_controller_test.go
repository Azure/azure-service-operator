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
	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	keyvaultresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"log"
	"strings"
	"time"
)

var _ = Describe("KeyVault Controller", func() {

	Context("Create and Delete", func() {

		It("should create and delete keyvault instances", func() {

			//t.Parallel()
			//RegisterTestingT(t)
			const timeout = time.Second * 240

			keyVaultName := "t-kv-dev-sample"
			keyVaultLocation := "westus"
			resourceGroupName := "t-rg-dev-controller"

			// declare resource group object
			resourceGroupInstance := &azurev1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resourceGroupName,
					Namespace: "default",
				},
				Spec: azurev1.ResourceGroupSpec{
					Location: "westus",
				},
			}

			// Create the resource group object and expect the Reconcile to be created
			err := k8sClient.Create(context.Background(), resourceGroupInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			resourceGroupNamespacedName := types.NamespacedName{Name: resourceGroupName, Namespace: "default"}
			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
				return resourceGroupInstance.Status.Provisioned == true
			}, timeout,
			).Should(BeTrue())

			// declare KeyVault object
			keyVaultInstance := &azurev1.KeyVault{
				ObjectMeta: metav1.ObjectMeta{
					Name:      keyVaultName,
					Namespace: "default",
				},
				Spec: azurev1.KeyVaultSpec{
					Location:          keyVaultLocation,
					ResourceGroupName: resourceGroupName,
				},
			}

			// Create the Keyvault object and expect the Reconcile to be created
			log.Print("create")
			err = k8sClient.Create(context.Background(), keyVaultInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			// prep query for get (?)
			keyVaultNamespacedName := types.NamespacedName{Name: keyVaultName, Namespace: "default"}

			// wait until resource is provisioned
			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), keyVaultNamespacedName, keyVaultInstance)
				//log.Print(keyVaultInstance.Status)
				return keyVaultInstance.Status.Provisioned == true
			}, timeout,
			).Should(BeTrue())

			// verify keyvault exists in Azure
			Eventually(func() bool {
				result, _ := keyvaultresourcemanager.GetVault(context.Background(), resourceGroupName, keyVaultInstance.Name)
				return result.Response.StatusCode == 200
			}, timeout,
			).Should(BeTrue())

			// delete keyvault

			k8sClient.Delete(context.Background(), keyVaultInstance)

			// verify its gone from kubernetes
			Eventually(func() bool {
				err := k8sClient.Get(context.Background(), keyVaultNamespacedName, keyVaultInstance)
				if err == nil {
					err = fmt.Errorf("")
				}
				return strings.Contains(err.Error(), "not found")
			}, timeout,
			).Should(BeTrue())

			// confirm resource is gone from Azure
			Eventually(func() bool {
				result, _ := keyvaultresourcemanager.GetVault(context.Background(), resourceGroupName, keyVaultInstance.Name)
				return result.Response.StatusCode == 200
			}, timeout,
			).Should(BeTrue())

		})
	})
})
