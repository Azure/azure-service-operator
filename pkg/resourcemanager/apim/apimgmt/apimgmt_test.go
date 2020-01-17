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

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App Insights", func() {

	var rgName string
	var psqlServer string
	var APIManager APIManager

	BeforeEach(func() {
		rgName = tc.ResourceGroupName
		APIManager = tc.APIManager
	})

	Context("Create and Delete", func() {
		It("should create and delete App insights instance in azure", func() {

			defer GinkgoRecover()

			apiMgmtInstance := "t-apimgmt-" + helpers.RandomString(10)

			// Create app insights instance
			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				_, err := APIManager.GetAPI(ctx, rgName, apiMgmtInstance)
				if err == nil {
					return true
				}
				_, err = APIManager.CreateAPI(ctx, rgName, "my-api-svc", "my-api", v1alpha1.APIProperties{
					APIRevision:            "v1",
					APIRevisionDescription: "revision description",
					IsCurrent:              true,
					IsOnline:               true,
					DisplayName:            "my-api",
					Description:            "API description",
					APIVersionDescription:  "version description",
					Path:                   "/api/test",
					Protocols:              []string{"http", "udp"},
					SubscriptionRequired:   false,
					ServiceURL:             "https://my-api/api",
				}, "")
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

			// Delete app insights instance
			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				_, err := APIManager.GetAPI(ctx, rgName, apiMgmtInstance)
				if err != nil {
					return true
				}
				_, err = APIManager.DeleteAPI(ctx, psqlServer, apiMgmtInstance)
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
