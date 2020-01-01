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

var _ = Describe("AppInsights", func() {
	var (
		key              types.NamespacedName
		created, fetched *AppInsights
	)

	AfterEach(func() {
		By("deleting the AppInsights operator")
		Expect(k8sClient.Delete(context.TODO(), created)).To(Succeed())
		Expect(k8sClient.Get(context.TODO(), key, created)).ToNot(Succeed())
	})

	Context("Create Resource", func() {

		It("should create an AppInsights operator successfully", func() {

			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}
			created = &AppInsights{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: AppInsightsSpec{
					Location:      "westus",
					Kind:          "web",
					ResourceGroup: "app-insights-dev-rg",
				}}

			By("creating an AppInsights operator")
			Expect(k8sClient.Create(context.TODO(), created)).To(Succeed())

			fetched = &AppInsights{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			Expect(fetched).To(Equal(created))

		})

	})

})
