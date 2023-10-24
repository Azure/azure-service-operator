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

	. "github.com/onsi/gomega"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/yaml"

	"github.com/Azure/azure-service-operator/v2/internal/crdmanagement"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

/*
 * LoadCRDs tests
 */

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

	loadedCRDs, err := crdManager.LoadOperatorCRDs(dir, "azureserviceoperator-system")
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(loadedCRDs).To(HaveLen(1))
	g.Expect(loadedCRDs[0]).To(Equal(crd))
}

func Test_LoadCRDs_FixesNamespace(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var port int32 = 443

	crd := makeBasicCRD("test")
	crd.Spec.Conversion = &apiextensions.CustomResourceConversion{
		Strategy: apiextensions.WebhookConverter,
		Webhook: &apiextensions.WebhookConversion{
			ClientConfig: &apiextensions.WebhookClientConfig{
				Service: &apiextensions.ServiceReference{
					Name:      "azureserviceoperator-webhook-service",
					Namespace: "azureserviceoperator-system",
					Path:      to.Ptr("/convert"),
					Port:      to.Ptr(port),
				},
				CABundle: makeFakeCABundle(),
			},
		},
	}
	dir := t.TempDir()
	logger := testcommon.NewTestLogger(t)

	bytes, err := yaml.Marshal(crd)
	g.Expect(err).ToNot(HaveOccurred())

	crdPath := filepath.Join(dir, "crd.yaml")
	g.Expect(os.WriteFile(crdPath, bytes, 0600)).To(Succeed())

	crdManager := crdmanagement.NewManager(logger, nil)

	loadedCRDs, err := crdManager.LoadOperatorCRDs(dir, "other-namespace")
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(loadedCRDs).To(HaveLen(1))
	g.Expect(loadedCRDs[0].Annotations).To(HaveLen(1))
	g.Expect(loadedCRDs[0].Annotations["cert-manager.io/inject-ca-from"]).To(Equal("other-namespace/azureserviceoperator-serving-cert"))
	g.Expect(loadedCRDs[0].Spec.Conversion.Webhook.ClientConfig.Service.Namespace).To(Equal("other-namespace"))
}

/*
 * FindMatchingCRDs tests
 */

func Test_FindMatchingCRDs_EqualCRDsCompareAsEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	existingCRD := makeBasicCRD("test")
	goalCRD := makeBasicCRD("test")
	existing := []apiextensions.CustomResourceDefinition{existingCRD}
	goal := []apiextensions.CustomResourceDefinition{goalCRD}

	logger := testcommon.NewTestLogger(t)
	crdManager := crdmanagement.NewManager(logger, nil)

	matching := crdManager.FindMatchingCRDs(existing, goal, crdmanagement.SpecEqual)

	g.Expect(matching).To(HaveLen(1))
}

func Test_FindMatchingCRDs_MissingCRD(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	goalCRD := makeBasicCRD("test")
	var existing []apiextensions.CustomResourceDefinition
	goal := []apiextensions.CustomResourceDefinition{goalCRD}

	logger := testcommon.NewTestLogger(t)
	crdManager := crdmanagement.NewManager(logger, nil)

	matching := crdManager.FindMatchingCRDs(existing, goal, crdmanagement.SpecEqual)

	g.Expect(matching).To(BeEmpty())
}

func Test_FindMatchingCRDs_CRDsWithDifferentConversionsCompareAsEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var port int32 = 443

	existingCRD := makeBasicCRD("test")
	existingCRD.Spec.Conversion = &apiextensions.CustomResourceConversion{
		Strategy: apiextensions.WebhookConverter,
		Webhook: &apiextensions.WebhookConversion{
			ClientConfig: &apiextensions.WebhookClientConfig{
				Service: &apiextensions.ServiceReference{
					Name:      "azureserviceoperator-webhook-service",
					Namespace: "azureserviceoperator-system",
					Path:      to.Ptr("/convert"),
					Port:      to.Ptr(port),
				},
				CABundle: makeFakeCABundle(),
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
					Path:      to.Ptr("/convert"),
					Port:      to.Ptr(port),
				},
			},
		},
	}
	existing := []apiextensions.CustomResourceDefinition{existingCRD}
	goal := []apiextensions.CustomResourceDefinition{goalCRD}

	logger := testcommon.NewTestLogger(t)
	crdManager := crdmanagement.NewManager(logger, nil)

	matching := crdManager.FindMatchingCRDs(existing, goal, crdmanagement.SpecEqual)

	g.Expect(matching).To(HaveLen(1))
	// Ensure that we still have CABundle set here
	g.Expect(existing[0].Spec.Conversion.Webhook.ClientConfig.CABundle).ToNot(BeEmpty())
}

/*
 * FindNonMatchingCRDs tests
 */

func Test_FindNonMatchingCRDs_EqualCRDsCompareAsEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	existingCRD := makeBasicCRD("test")
	goalCRD := makeBasicCRD("test")
	existing := []apiextensions.CustomResourceDefinition{existingCRD}
	goal := []apiextensions.CustomResourceDefinition{goalCRD}

	logger := testcommon.NewTestLogger(t)
	crdManager := crdmanagement.NewManager(logger, nil)

	nonMatching := crdManager.FindNonMatchingCRDs(existing, goal, crdmanagement.SpecEqual)

	g.Expect(nonMatching).To(BeEmpty())
}

func Test_FindNonMatchingCRDs_MissingCRD(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	goalCRD := makeBasicCRD("test")
	var existing []apiextensions.CustomResourceDefinition
	goal := []apiextensions.CustomResourceDefinition{goalCRD}

	logger := testcommon.NewTestLogger(t)
	crdManager := crdmanagement.NewManager(logger, nil)

	nonMatching := crdManager.FindNonMatchingCRDs(existing, goal, crdmanagement.SpecEqual)

	g.Expect(nonMatching).To(HaveLen(1))
}

func Test_FindNonMatchingCRDs_CRDsWithDifferentConversionsCompareAsEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var port int32 = 443

	existingCRD := makeBasicCRD("test")
	existingCRD.Spec.Conversion = &apiextensions.CustomResourceConversion{
		Strategy: apiextensions.WebhookConverter,
		Webhook: &apiextensions.WebhookConversion{
			ClientConfig: &apiextensions.WebhookClientConfig{
				Service: &apiextensions.ServiceReference{
					Name:      "azureserviceoperator-webhook-service",
					Namespace: "azureserviceoperator-system",
					Path:      to.Ptr("/convert"),
					Port:      to.Ptr(port),
				},
				CABundle: makeFakeCABundle(),
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
					Path:      to.Ptr("/convert"),
					Port:      to.Ptr(port),
				},
			},
		},
	}
	existing := []apiextensions.CustomResourceDefinition{existingCRD}
	goal := []apiextensions.CustomResourceDefinition{goalCRD}

	logger := testcommon.NewTestLogger(t)
	crdManager := crdmanagement.NewManager(logger, nil)

	nonMatching := crdManager.FindNonMatchingCRDs(existing, goal, crdmanagement.SpecEqual)

	g.Expect(nonMatching).To(BeEmpty())
	// Ensure that we still have CABundle set here
	g.Expect(existing[0].Spec.Conversion.Webhook.ClientConfig.CABundle).ToNot(BeEmpty())
}

/*
 * DetermineCRDsToInstallOrUpgrade tests
 */

func Test_DetermineCRDsToInstallOrUpgrade(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		goal     []apiextensions.CustomResourceDefinition
		existing []apiextensions.CustomResourceDefinition
		patterns string
		validate func(g *WithT, instructions []*crdmanagement.CRDInstallationInstruction)
	}{
		{
			name:     "Skips CRDs if no pattern or installed",
			goal:     []apiextensions.CustomResourceDefinition{makeBasicCRD("test")},
			existing: nil,
			patterns: "",
			validate: func(g *WithT, instructions []*crdmanagement.CRDInstallationInstruction) {
				g.Expect(instructions).To(HaveLen(1))
				g.Expect(instructions[0].FilterResult).To(Equal(crdmanagement.Excluded))
				g.Expect(instructions[0].FilterReason).To(ContainSubstring("was not matched by CRD pattern and did not already exist in cluster"))
				apply, _ := instructions[0].ShouldApply()
				g.Expect(apply).To(BeFalse())
			},
		},
		{
			name:     "Matches CRDs if pattern matches",
			goal:     []apiextensions.CustomResourceDefinition{makeBasicCRD("test")},
			existing: nil,
			patterns: "testrp.azure.com/*",
			validate: func(g *WithT, instructions []*crdmanagement.CRDInstallationInstruction) {
				g.Expect(instructions).To(HaveLen(1))
				g.Expect(instructions[0].FilterResult).To(Equal(crdmanagement.MatchedPattern))
				g.Expect(instructions[0].FilterReason).To(Equal("CRD named \"testrp.azure.com/test\" matched pattern \"testrp.azure.com/*\""))
				g.Expect(instructions[0].DiffResult).To(Equal(crdmanagement.SpecDifferent))
				apply, _ := instructions[0].ShouldApply()
				g.Expect(apply).To(BeTrue())
			},
		},
		{
			name:     "Matches if CRD already installed",
			goal:     []apiextensions.CustomResourceDefinition{makeBasicCRD("test")},
			existing: []apiextensions.CustomResourceDefinition{makeBasicCRDWithVersion("test", "1.1")},
			patterns: "",
			validate: func(g *WithT, instructions []*crdmanagement.CRDInstallationInstruction) {
				g.Expect(instructions).To(HaveLen(1))
				g.Expect(instructions[0].FilterResult).To(Equal(crdmanagement.MatchedExistingCRD))
				g.Expect(instructions[0].FilterReason).To(Equal("A CRD named \"testrp.azure.com/test\" was already installed, considering that existing CRD for update"))
				g.Expect(instructions[0].DiffResult).To(Equal(crdmanagement.VersionDifferent))
				apply, _ := instructions[0].ShouldApply()
				g.Expect(apply).To(BeTrue())
			},
		},
		{
			name:     "Matches if CRD already installed and pattern matches, pattern matches wins",
			goal:     []apiextensions.CustomResourceDefinition{makeBasicCRD("test")},
			existing: []apiextensions.CustomResourceDefinition{makeBasicCRDWithVersion("test", "1.1")},
			patterns: "testrp.azure.com/*",
			validate: func(g *WithT, instructions []*crdmanagement.CRDInstallationInstruction) {
				g.Expect(instructions).To(HaveLen(1))
				g.Expect(instructions[0].FilterResult).To(Equal(crdmanagement.MatchedPattern))
				g.Expect(instructions[0].FilterReason).To(Equal("CRD named \"testrp.azure.com/test\" matched pattern \"testrp.azure.com/*\""))
				g.Expect(instructions[0].DiffResult).To(Equal(crdmanagement.VersionDifferent))
				apply, _ := instructions[0].ShouldApply()
				g.Expect(apply).To(BeTrue())
			},
		},
		{
			name:     "Pattern matches but CRD already installed, nothing to do",
			goal:     []apiextensions.CustomResourceDefinition{makeBasicCRD("test")},
			existing: []apiextensions.CustomResourceDefinition{makeBasicCRD("test")},
			patterns: "testrp.azure.com/*",
			validate: func(g *WithT, instructions []*crdmanagement.CRDInstallationInstruction) {
				g.Expect(instructions).To(HaveLen(1))
				g.Expect(instructions[0].FilterResult).To(Equal(crdmanagement.MatchedPattern))
				g.Expect(instructions[0].FilterReason).To(Equal("CRD named \"testrp.azure.com/test\" matched pattern \"testrp.azure.com/*\""))
				g.Expect(instructions[0].DiffResult).To(Equal(crdmanagement.NoDifference))
				apply, _ := instructions[0].ShouldApply()
				g.Expect(apply).To(BeFalse())
			},
		},
		{
			name:     "Pattern matches subset of CRDs from group, only that subset is installed",
			goal:     []apiextensions.CustomResourceDefinition{makeBasicCRD("test1"), makeBasicCRD("test2")},
			existing: nil,
			patterns: "testrp.azure.com/test1",
			validate: func(g *WithT, instructions []*crdmanagement.CRDInstallationInstruction) {
				g.Expect(instructions).To(HaveLen(2))

				// use loop here rather than index as installation instruction order may not match CRD order
				for _, instruction := range instructions {
					if instruction.CRD.Name == "test1.testrp.azure.com" {
						g.Expect(instruction.FilterResult).To(Equal(crdmanagement.MatchedPattern))
						g.Expect(instruction.FilterReason).To(Equal("CRD named \"testrp.azure.com/test1\" matched pattern \"testrp.azure.com/test1\""))
						g.Expect(instruction.DiffResult).To(Equal(crdmanagement.SpecDifferent))
						apply, _ := instruction.ShouldApply()
						g.Expect(apply).To(BeTrue())
					} else {
						g.Expect(instruction.FilterResult).To(Equal(crdmanagement.Excluded))
						g.Expect(instruction.FilterReason).To(ContainSubstring("was not matched by CRD pattern and did not already exist in cluster"))
						apply, _ := instruction.ShouldApply()
						g.Expect(apply).To(BeFalse())
					}
				}
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)
			logger := testcommon.NewTestLogger(t)
			crdManager := crdmanagement.NewManager(logger, nil)

			instructions, err := crdManager.DetermineCRDsToInstallOrUpgrade(c.goal, c.existing, c.patterns)
			g.Expect(err).ToNot(HaveOccurred())

			c.validate(g, instructions)
		})
	}
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
		crdmanagement.ServiceOperatorAppLabel:     crdmanagement.ServiceOperatorAppValue,
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

	loadedCRDs, err := crdManager.LoadOperatorCRDs(path, defaultNamespace)
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

/*
 * Helper functions
 */

func makeFakeCABundle() []byte {
	return []byte{17, 14, 12, 21, 33, 61, 25, 99, 111}
}

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
			Annotations: map[string]string{
				"cert-manager.io/inject-ca-from": "azureserviceoperator-system/azureserviceoperator-serving-cert",
			},
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group: "testrp.azure.com",
			Names: apiextensions.CustomResourceDefinitionNames{
				Kind: name,
			},
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

func makeBasicCRDWithVersion(name string, version string) apiextensions.CustomResourceDefinition {
	crd := makeBasicCRD(name)
	crd.Labels = map[string]string{
		crdmanagement.ServiceOperatorVersionLabelOld: version,
	}

	return crd
}
