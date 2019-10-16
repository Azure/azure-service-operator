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
	"log"
	"net/http"
	"strings"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("KeyVault Controller", func() {

	keyVaultName := "t-kv-dev-" + helpers.RandomString(10)
	const poll = time.Second * 10

	Context("Key Vault Controller", func() {

		It("Should Create and Delete Key Vault instances", func() {

			keyVaultLocation := tc.resourceGroupLocation

			// Declare KeyVault object
			keyVaultInstance := &azurev1alpha1.KeyVault{
				ObjectMeta: metav1.ObjectMeta{
					Name:      keyVaultName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.KeyVaultSpec{
					Location:          keyVaultLocation,
					ResourceGroupName: tc.resourceGroupName,
				},
			}

			// Create the Keyvault object and expect the Reconcile to be created
			log.Print("Create")
			err := tc.k8sClient.Create(context.Background(), keyVaultInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			// Prep query for get
			keyVaultNamespacedName := types.NamespacedName{Name: keyVaultName, Namespace: "default"}

			// Wait until key vault is provisioned
			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), keyVaultNamespacedName, keyVaultInstance)
				//log.Print(keyVaultInstance.Status)
				return keyVaultInstance.Status.ID != nil
			}, tc.timeout,
			).Should(BeTrue())

			// verify key vault exists in Azure
			Eventually(func() bool {
				result, _ := tc.keyVaultManager.GetVault(context.Background(), tc.resourceGroupName, keyVaultInstance.Name)
				return result.Response.StatusCode == http.StatusOK
			}, tc.timeout,
			).Should(BeTrue())

			// delete key vault
			err = tc.k8sClient.Delete(context.Background(), keyVaultInstance)
			Expect(err).NotTo(HaveOccurred())

			// verify key vault is gone from kubernetes
			Eventually(func() bool {
				err := tc.k8sClient.Get(context.Background(), keyVaultNamespacedName, keyVaultInstance)
				if err == nil {
					err = fmt.Errorf("")
				}
				return strings.Contains(err.Error(), "not found")
			}, tc.timeout,
			).Should(BeTrue())

			// confirm key vault is gone from Azure
			Eventually(func() bool {
				result, _ := tc.keyVaultManager.GetVault(context.Background(), tc.resourceGroupName, keyVaultInstance.Name)
				return result.Response.StatusCode == http.StatusNotFound
			}, tc.timeout, poll,
			).Should(BeTrue())
		})
	})
})
