// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

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
