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
	"time"

	creatorv1 "Telstra.Dx.AzureOperator/api/v1"
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("ResourceGroup Controller", func() {

	const timeout = time.Second * 240

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create", func() {
		It("should create resource groups", func() {
			namespacedName := types.NamespacedName{Name: "testdx", Namespace: "default"}
			instance := &creatorv1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      namespacedName.Name,
					Namespace: namespacedName.Namespace,
				},
				Spec: creatorv1.ResourceGroupSpec{
					Location: "westus",
				},
			}

			// Create the Resource Group object and expect the Reconcile to be created
			err := k8sClient.Create(context.Background(), instance)

			time.Sleep(2 * time.Second)
			// The instance object may not be a valid object because it might be missing some required fields.
			// Please modify the instance object by adding required fields and then remove the following if statement.
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(2 * time.Second)

			instance2 := &creatorv1.Eventhub{}
			err = k8sClient.Get(context.Background(), namespacedName, instance2)
			Expect(err).NotTo(HaveOccurred())
			err = k8sClient.Delete(context.Background(), instance2)
			Expect(err).NotTo(HaveOccurred())

		})
	})
})
