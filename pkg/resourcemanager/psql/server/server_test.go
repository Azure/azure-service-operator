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

package server

import (
	"fmt"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("PSQL server", func() {

	var rgName string
	var location string
	var psqlServer string
	var PSQLManager PostgreSQLServerManager

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.ResourceGroupName
		location = tc.ResourceGroupLocation
		PSQLManager = tc.postgreSQLServerManager
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete PSQL server in azure", func() {

			defer GinkgoRecover()

			psqlServer = "t-dev-psql-srv-" + helpers.RandomString(10)

			pSQLSku := azurev1alpha1.PSQLSku{
				Name:     "B_Gen5_2",
				Tier:     "Basic",
				Capacity: 2,
				Size:     "51200",
				Family:   "Gen5",
			}

			pSQLServerInstance := &azurev1alpha1.PostgreSQLServer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      psqlServer,
					Namespace: "default",
				},
				Spec: azurev1alpha1.PostgreSQLServerSpec{
					Location:      location,
					ResourceGroup: rgName,
					Sku:           pSQLSku,
				},
			}

			Eventually(func() bool {
				time.Sleep(3 * time.Second)
				result, err := PSQLManager.Ensure(ctx, pSQLServerInstance)
				if err != nil {
					fmt.Println(err.Error())
					if !errhelp.IsAsynchronousOperationNotComplete(err) {
						fmt.Println("error occured")
						return false
					}
				}
				return result
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())

			time.Sleep(10 * time.Minute)

		})

	})
})
