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

package appinsights

import (
	"fmt"
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App Insights", func() {

	var rgName string
	var location string
	var psqlServer string
	var AppInsightsManager ApplicationInsightsManager

	BeforeEach(func() {
		rgName = tc.ResourceGroupName
		location = tc.ResourceGroupLocation
		AppInsightsManager = tc.AppInsightsManager
	})

	Context("Create and Delete", func() {
		It("should create and delete App insights instance in azure", func() {

			defer GinkgoRecover()

			appinsightsInstance := "t-appinsight-" + helpers.RandomString(10)

			// Create app insights instance
			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				_, err := AppInsightsManager.GetAppInsights(ctx, rgName, appinsightsInstance)
				if err == nil {
					return true
				}
				_, err = AppInsightsManager.CreateAppInsights(ctx, rgName, "web", "other", location, appinsightsInstance)
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
				_, err := AppInsightsManager.GetAppInsights(ctx, rgName, appinsightsInstance)
				if err != nil {
					return true
				}
				_, err = AppInsightsManager.DeleteAppInsights(ctx, psqlServer, appinsightsInstance)
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
