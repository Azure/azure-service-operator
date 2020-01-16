/*
Copyright 2019 Microsoft.
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

package vnet

import (
	"fmt"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VNet", func() {

	var location string
	var rgName string
	var addressSpace string
	var subnetName string
	var subnetPrefix string
	var vnetManager VNetManager

	BeforeEach(func() {
		location = tc.ResourceGroupLocation
		rgName = tc.ResourceGroupName
		addressSpace = tc.AddressSpace
		subnetName = tc.SubnetName
		subnetPrefix = tc.SubnetAddressPrefix
		vnetManager = tc.VirtualNetworkManager
	})

	Context("Create and Delete", func() {
		It("should create and delete a VNet instance in azure", func() {

			defer GinkgoRecover()

			vnetName := "t-vnet-" + helpers.RandomString(10)

			// Create vnet instance
			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				exists, _ := vnetManager.VNetExists(ctx, rgName, vnetName)
				if exists {
					return true
				}
				_, err := vnetManager.CreateVNet(ctx, location, rgName, vnetName, addressSpace, []azurev1alpha1.VNetSubnets{
					azurev1alpha1.VNetSubnets{
						SubnetName:          subnetName,
						SubnetAddressPrefix: subnetPrefix,
					},
				})
				if err != nil {
					fmt.Println(err.Error())
					if !errhelp.IsAsynchronousOperationNotComplete(err) {
						fmt.Println("error occured")
						return false
					}
				}
				return true
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())

			time.Sleep(3 * time.Minute)

			// Delete vnet instance
			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				exists, _ := vnetManager.VNetExists(ctx, rgName, vnetName)
				if !exists {
					return true
				}
				_, err := vnetManager.DeleteVNet(ctx, rgName, vnetName)
				if err != nil {
					fmt.Println(err.Error())
					if !errhelp.IsAsynchronousOperationNotComplete(err) {
						fmt.Println("error occured")
						return false
					}
				}
				return err == nil
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())
		})
	})
})
