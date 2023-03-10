// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package crdmanagement_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/yaml"

	"github.com/Azure/azure-service-operator/v2/internal/crdmanagement"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
)

func newFakeKubeClient(s *runtime.Scheme) kubeclient.Client {
	fakeClient := fake.NewClientBuilder().WithScheme(s).Build()
	return kubeclient.NewClient(fakeClient)
}

func newTestScheme() *runtime.Scheme {
	s := runtime.NewScheme()
	_ = apiextensions.AddToScheme(s)

	return s
}

func makeBasicCRD(name string) apiextensions.CustomResourceDefinition {
	return apiextensions.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apiextensions.k8s.io/v1",
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s.testrp.azure.com", name),
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group: "testrp.azure.com",
			Scope: apiextensions.NamespaceScoped,
			Versions: []apiextensions.CustomResourceDefinitionVersion{
				{
					Name:    "v1alpha1api20181130",
					Served:  true,
					Storage: true,
					Schema: &apiextensions.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
							Properties: map[string]apiextensions.JSONSchemaProps{
								"apiVersion": {
									Type:        "string",
									Description: "APIVersion description",
								},
								"kind": {
									Type:        "string",
									Description: "Kind description",
								},
							},
						},
					},
				},
			},
		},
	}
}

func Test_LoadCRDs(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	crd := makeBasicCRD("test")
	dir := t.TempDir()
	logger := testcommon.NewTestLogger(t)

	bytes, err := yaml.Marshal(crd)
	g.Expect(err).ToNot(HaveOccurred())

	crdPath := filepath.Join(dir, "crd.yaml")
	g.Expect(os.WriteFile(crdPath, bytes, 0600)).To(Succeed())

	crdManager := crdmanagement.NewManager(logger, nil)

	loadedCRDs, err := crdManager.LoadOperatorCRDs(dir)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(loadedCRDs).To(HaveLen(1))
	g.Expect(loadedCRDs[0]).To(Equal(crd))
}

func Test_CompareExistingCRDsWithGoal_EqualCRDsCompareAsEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	existingCRD := makeBasicCRD("test")
	goalCRD := makeBasicCRD("test")
	existing := []apiextensions.CustomResourceDefinition{existingCRD}
	goal := []apiextensions.CustomResourceDefinition{goalCRD}

	logger := testcommon.NewTestLogger(t)
	crdManager := crdmanagement.NewManager(logger, nil)

	needUpdate := crdManager.FindGoalCRDsNeedingUpdate(existing, goal, crdmanagement.SpecEqual)

	g.Expect(needUpdate).To(BeEmpty())
}

func Test_CompareExistingCRDsWithGoal_CRDsWithDifferentConversionsCompareAsEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	existingCRD := makeBasicCRD("test")
	existingCRD.Spec.Conversion = &apiextensions.CustomResourceConversion{
		Strategy: apiextensions.WebhookConverter,
		Webhook: &apiextensions.WebhookConversion{
			ClientConfig: &apiextensions.WebhookClientConfig{
				Service: &apiextensions.ServiceReference{
					Name:      "azureserviceoperator-webhook-service",
					Namespace: "azureserviceoperator-system",
					Path:      to.StringPtr("/convert"),
					Port:      to.Int32Ptr(443),
				},
				CABundle: []byte{17, 14, 12, 21, 33, 61, 25, 99, 111},
			},
		},
	}

	goalCRD := makeBasicCRD("test")
	goalCRD.Spec.Conversion = &apiextensions.CustomResourceConversion{
		Strategy: apiextensions.WebhookConverter,
		Webhook: &apiextensions.WebhookConversion{
			ClientConfig: &apiextensions.WebhookClientConfig{
				Service: &apiextensions.ServiceReference{
					Name:      "azureserviceoperator-webhook-service",
					Namespace: "azureserviceoperator-system",
					Path:      to.StringPtr("/convert"),
					Port:      to.Int32Ptr(443),
				},
			},
		},
	}
	existing := []apiextensions.CustomResourceDefinition{existingCRD}
	goal := []apiextensions.CustomResourceDefinition{goalCRD}

	logger := testcommon.NewTestLogger(t)
	crdManager := crdmanagement.NewManager(logger, nil)

	needUpdate := crdManager.FindGoalCRDsNeedingUpdate(existing, goal, crdmanagement.SpecEqual)

	g.Expect(needUpdate).To(BeEmpty())
	// Ensure that we still have CABundle set here
	g.Expect(existing[0].Spec.Conversion.Webhook.ClientConfig.CABundle).ToNot(BeEmpty())
}

func Test_ListCRDs_ListsOnlyCRDsMatchingLabel(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	kubeClient := newFakeKubeClient(newTestScheme())
	crd1 := makeBasicCRD("test1")
	crd2 := makeBasicCRD("test2")
	crd3 := makeBasicCRD("test3")

	crd3.Labels = map[string]string{
		crdmanagement.ServiceOperatorVersionLabel: "123",
	}

	g.Expect(kubeClient.Create(ctx, &crd1)).To(Succeed())
	g.Expect(kubeClient.Create(ctx, &crd2)).To(Succeed())
	g.Expect(kubeClient.Create(ctx, &crd3)).To(Succeed())

	logger := testcommon.NewTestLogger(t)
	crdManager := crdmanagement.NewManager(logger, kubeClient)

	crds, err := crdManager.ListOperatorCRDs(ctx)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(crds).To(HaveLen(1))
}

// This test requires that the task target `bundle-crds` has been run
func Test_BundledCRDs_HaveExactlyTwoInstancesOfNamespace(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	path := "../../out/crds"
	logger := testcommon.NewTestLogger(t)
	crdManager := crdmanagement.NewManager(logger, nil)

	defaultNamespace := "azureserviceoperator-system"

	loadedCRDs, err := crdManager.LoadOperatorCRDs(path)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(loadedCRDs).ToNot(BeEmpty())
	// The raw JSON should contain exactly 2 locations where the namespace is referenced. If this changes, we need
	// to update the code at crdManager.fixCRDNamespace

	crd := loadedCRDs[0]
	bytes, err := json.Marshal(crd)
	g.Expect(err).ToNot(HaveOccurred())

	count := strings.Count(string(bytes), defaultNamespace)
	g.Expect(count).To(Equal(2))
}
