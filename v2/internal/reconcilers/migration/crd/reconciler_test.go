/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package crd_test

import (
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/crdmanagement"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/migration/crd"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
)

// This test requires that the task target `bundle-crds` has been run
func TestReconcileCRDs(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		instances        []*unstructured.Unstructured
		expectDeprecated bool
	}{
		"NoInstancesExist_Migrated": {
			instances:        nil,
			expectDeprecated: true,
		},
		"InstancesExistWithoutLatestVersion_NotMigrated": {
			instances: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "containerservice.azure.com/v1api20240901storage",
						"kind":       "ManagedCluster",
						"metadata": map[string]interface{}{
							"name":      "test-cluster",
							"namespace": "default",
						},
					},
				},
			},
			expectDeprecated: false,
		},
		"InstancesExistWithLatestVersion_Migrated": {
			instances: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "containerservice.azure.com/v1api20240901storage",
						"kind":       "ManagedCluster",
						"metadata": map[string]interface{}{
							"name":      "test-cluster",
							"namespace": "default",
							"labels": map[string]interface{}{
								"serviceoperator.azure.com/last-reconciled-version": "v1.0.0",
							},
						},
					},
				},
			},
			expectDeprecated: true,
		},
		"InstancesExistWithOlderVersion_NotMigrated": {
			instances: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "containerservice.azure.com/v1api20240901storage",
						"kind":       "ManagedCluster",
						"metadata": map[string]interface{}{
							"name":      "test-cluster",
							"namespace": "default",
							"labels": map[string]interface{}{
								"serviceoperator.azure.com/last-reconciled-version": "v0.0.9",
							},
						},
					},
				},
			},
			expectDeprecated: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			testData := testSetup(t)

			// Add instances if provided
			for _, instance := range tt.instances {
				g.Expect(testData.kubeClient.Create(t.Context(), instance)).To(Succeed())
			}

			deprecatedCRDVersions := map[string][]string{
				"managedclusters.containerservice.azure.com": {"v1api20231102previewstorage"},
			}

			options := crd.Options{
				VersionProvider: func() string {
					return "v1.0.0" // use a single fixed version for testing
				},
			}

			reconciler := crd.NewReconciler(testData.kubeClient, nil, deprecatedCRDVersions, options)
			result, err := reconciler.Reconcile(t.Context())
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(result).To(Equal(ctrl.Result{RequeueAfter: 1 * time.Hour}))

			// The CRD should have the deprecated versions removed from its status.storedVersions
			crd := &apiextensions.CustomResourceDefinition{}
			g.Expect(testData.kubeClient.Get(
				t.Context(),
				types.NamespacedName{
					Name: "managedclusters.containerservice.azure.com",
				},
				crd)).To(Succeed())

			if tt.expectDeprecated {
				g.Expect(crd.Status.StoredVersions).ToNot(ContainElement("v1api20231102previewstorage"))
			} else {
				g.Expect(crd.Status.StoredVersions).To(ContainElement("v1api20231102previewstorage"))
			}
		})
	}
}

func NewFakeKubeClient(s *runtime.Scheme) kubeclient.Client {
	fakeClient := fake.NewClientBuilder().WithScheme(s).Build()
	return kubeclient.NewClient(fakeClient)
}

type testData struct {
	cfg        config.Values
	kubeClient kubeclient.Client
	logger     logr.Logger
	crdManager *crdmanagement.Manager
}

func testSetup(t *testing.T) *testData {
	asoScheme := api.CreateScheme()
	_ = apiextensions.AddToScheme(asoScheme)
	g := NewGomegaWithT(t)

	kubeClient := NewFakeKubeClient(asoScheme)

	logger := testcommon.NewTestLogger(t)
	cfg := config.Values{}

	crdManager := crdmanagement.NewManager(logger, kubeClient, nil)
	crdPath := "../../../../out/crds"
	namespace := "azureserviceoperator-system"

	testData := &testData{
		cfg: cfg,
		// s:          s,
		kubeClient: kubeClient,
		logger:     logger,
		crdManager: crdManager,
	}

	goalCRDs, err := crdManager.LoadOperatorCRDs(crdPath, namespace)
	g.Expect(err).ToNot(HaveOccurred())

	for _, crd := range goalCRDs {
		// We mark all storage versions as "storedVersions" for testing purposes
		crd.Status.StoredVersions = []string{}
		for _, version := range crd.Spec.Versions {
			if strings.HasSuffix(version.Name, "storage") {
				crd.Status.StoredVersions = append(crd.Status.StoredVersions, version.Name)
			}
		}

		g.Expect(kubeClient.Create(t.Context(), &crd)).To(Succeed())
	}

	return testData
}
