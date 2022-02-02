// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

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
				_, err := vnetManager.GetVNet(ctx, rgName, vnetName)
				if err == nil {
					return true
				}
				_, err = vnetManager.CreateVNet(ctx, location, rgName, vnetName, addressSpace, []azurev1alpha1.VNetSubnets{
					{
						SubnetName:          subnetName,
						SubnetAddressPrefix: subnetPrefix,
					},
				})
				if err != nil {
					fmt.Println(err.Error())
					ignore := []string{
						errhelp.AsyncOpIncompleteError,
					}
					azerr := errhelp.NewAzureError(err)
					if !helpers.ContainsString(ignore, azerr.Type) {
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
				_, err := vnetManager.GetVNet(ctx, rgName, vnetName)
				if err != nil {
					return true
				}
				_, err = vnetManager.DeleteVNet(ctx, rgName, vnetName)
				if err != nil {
					fmt.Println(err.Error())
					ignore := []string{
						errhelp.AsyncOpIncompleteError,
					}
					azerr := errhelp.NewAzureError(err)
					if !helpers.ContainsString(ignore, azerr.Type) {
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
