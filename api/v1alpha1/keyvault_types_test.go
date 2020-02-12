/*

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

package v1alpha1

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// These tests are written in BDD-style using Ginkgo framework. Refer to
// http://onsi.github.io/ginkgo to learn more.

var _ = Describe("KeyVault", func() {
	var (
		key              types.NamespacedName
		created, fetched *KeyVault
	)

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
	Context("KeyVault", func() {

		It("should create an object successfully", func() {

			ctx := context.Background()

			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}
			created = &KeyVault{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: KeyVaultSpec{
					Location:      "westus",
					ResourceGroup: "foo-rg",
				}}

			By("creating an API obj")
			Expect(k8sClient.Create(ctx, created)).To(Succeed())

			fetched = &KeyVault{}
			Expect(k8sClient.Get(ctx, key, fetched)).To(Succeed())
			Expect(fetched).To(Equal(created))

			By("deleting the created object")
			Expect(k8sClient.Delete(ctx, created)).To(Succeed())
			Expect(k8sClient.Get(ctx, key, created)).ToNot(Succeed())
		})

	})

})
