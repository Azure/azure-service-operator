package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("ResourceGroup Controller", func() {

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

	Context("Create and Delete", func() {
		It("should create and delete resource group instances", func() {
			resourceGroupName := "t-rg-dev-" + helpers.RandomString(10)

			defer GinkgoRecover()

			var err error

			// Create the ResourceGroup object and expect the Reconcile to be created
			resourceGroupInstance := &azurev1alpha1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resourceGroupName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.ResourceGroupSpec{
					Location: tc.resourceGroupLocation,
				},
			}

			// create rg
			err = tc.k8sClient.Create(context.Background(), resourceGroupInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			resourceGroupNamespacedName := types.NamespacedName{Name: resourceGroupName, Namespace: "default"}

			// verify sure rg has a finalizer
			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
				return resourceGroupInstance.HasFinalizer(resourceGroupFinalizerName)
			}, tc.timeout,
			).Should(BeTrue())

			// verify rg gets submitted
			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
				return resourceGroupInstance.IsSubmitted()
			}, tc.timeout,
			).Should(BeTrue())

			// verify rg exists in azure
			Eventually(func() bool {
				_, err := tc.resourceGroupManager.CheckExistence(context.Background(), resourceGroupName)
				return err == nil
			}, tc.timeout,
			).Should(BeTrue())

			// delete rg
			err = tc.k8sClient.Delete(context.Background(), resourceGroupInstance)
			Expect(err).NotTo(HaveOccurred())

			// verify rg is being deleted
			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
				return resourceGroupInstance.IsBeingDeleted()
			}, tc.timeout,
			).Should(BeTrue())

			// verify rg is gone from kubernetes
			Eventually(func() bool {
				err := tc.k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
				if err == nil {
					err = fmt.Errorf("")
				}
				return strings.Contains(err.Error(), "not found")
			}, tc.timeout,
			).Should(BeTrue())

			// verify rg is gone from Azure
			Eventually(func() bool {
				result, _ := tc.resourceGroupManager.CheckExistence(context.Background(), resourceGroupName)
				return result.Response.StatusCode == http.StatusNotFound
			}, tc.timeout,
			).Should(BeTrue())
		})
	})
})
