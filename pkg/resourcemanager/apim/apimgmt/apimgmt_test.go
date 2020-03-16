/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
*/
// +build all apimgmt
package apimgmt

import (
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	APIServiceName string = "AzureOperatorsTestAPIM"
	APIId          string = "apiId1"
	APIETag        string = "eTagASOTest"
)

var _ = Describe("API Management", func() {

	var rgName string
	var APIManager APIManager
	var contract apimanagement.APIContract
	var err error

	BeforeEach(func() {
		rgName = tc.ResourceGroupName
		APIManager = tc.APIManager
	})

	Context("Create and Delete", func() {
		It("should create and delete API instance in azure", func() {

			defer GinkgoRecover()

			// Create API instance
			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				contract, err = APIManager.GetAPI(ctx, rgName, APIServiceName, "test-revision")
				if err == nil {
					return true
				}
				_, err = APIManager.CreateAPI(
					ctx,
					rgName,
					APIServiceName,
					APIId,
					v1alpha1.APIProperties{
						Format: "Openapi",
						APIVersionSet: v1alpha1.APIVersionSet{
							Name: "apiversionsetdetails1",
						},
						APIVersionSetID:        "",
						APIRevision:            "test-revision",
						APIRevisionDescription: "revision description",
						IsCurrent:              true,
						IsOnline:               true,
						DisplayName:            "aso-apimgmt-test",
						Description:            "API description",
						APIVersionDescription:  "version description",
						Path:                   "/api/test",
						Protocols:              []string{"http", "udp"},
						SubscriptionRequired:   false,
					},
					APIETag)
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

			time.Sleep(5 * time.Minute)

			// Delete API instance
			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				_, err := APIManager.GetAPI(ctx, rgName, APIServiceName, "test-revision")
				if err != nil {
					return true
				}
				_, err = APIManager.DeleteAPI(ctx, tc.ResourceGroupName, APIServiceName, *contract.ID, *contract.APIRevision, true)
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
