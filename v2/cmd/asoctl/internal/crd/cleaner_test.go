/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/fake"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fake2 "sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/Azure/azure-service-operator/v2/api"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

type clientSet struct {
	fakeAPIExtClient apiextensions.ApiextensionsV1Interface
	fakeClient       client.WithWatch
	cleaner          *Cleaner
}

// TODO: Currently we need to create clientsets for each test as they run in parallel and we run into `resource already exists` error.
// TODO: We may require a testing suite re-use the clientsets efficiently.
func makeClientSets() *clientSet {
	fakeAPIExtClient := fake.NewSimpleClientset().ApiextensionsV1()
	fakeClient := fake2.NewClientBuilder().WithScheme(api.CreateScheme()).Build()
	cleaner := NewCleaner(
		fakeAPIExtClient.CustomResourceDefinitions(),
		fakeClient,
		false, // dry-run
		logr.Discard())
	return &clientSet{
		fakeAPIExtClient: fakeAPIExtClient,
		fakeClient:       fakeClient,
		cleaner:          cleaner,
	}
}

func Test_CleanDeprecatedCRDVersions_CleansBetaVersion_IfExists(t *testing.T) {
	t.Parallel()

	c := makeClientSets()

	g := NewGomegaWithT(t)

	betaVersion := "v1beta20200601"
	gaVersion := "v1api20200601"

	definition := newCRDWithStoredVersions(betaVersion, gaVersion)

	_, err := c.fakeAPIExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	g.Expect(err).To(BeNil())

	err = c.cleaner.Run(context.TODO())
	g.Expect(err).To(BeNil())

	crd, err := c.fakeAPIExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	g.Expect(err).To(BeNil())

	g.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	g.Expect(crd.Status.StoredVersions).ToNot(BeEquivalentTo(definition.Status.StoredVersions))
	g.Expect(crd.Status.StoredVersions).ToNot(ContainElement(betaVersion))
	g.Expect(crd.Status.StoredVersions).To(ContainElement(gaVersion))
}

func Test_CleanDeprecatedCRDVersions_CleansHandcraftedBetaVersion_IfExists(t *testing.T) {
	t.Parallel()

	c := makeClientSets()

	g := NewGomegaWithT(t)

	betaVersion := "v1beta1"
	gaVersion := "v1"

	definition := newCRDWithStoredVersions(betaVersion, gaVersion)

	_, err := c.fakeAPIExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	g.Expect(err).To(BeNil())

	err = c.cleaner.Run(context.TODO())
	g.Expect(err).To(BeNil())

	crd, err := c.fakeAPIExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	g.Expect(err).To(BeNil())

	g.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	g.Expect(crd.Status.StoredVersions).ToNot(BeEquivalentTo(definition.Status.StoredVersions))
	g.Expect(crd.Status.StoredVersions).ToNot(ContainElement(betaVersion))
	g.Expect(crd.Status.StoredVersions).To(ContainElement(gaVersion))
}

func Test_CleanDeprecatedCRDVersions_CleansTrustedAccessRoleBindings(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := makeClientSets()

	oldVersion := "v1api20230202previewstorage"
	newVersion := "v1api20231001storage"

	// create CRD
	definition := newCRDWithStoredVersionsAndName(
		"containerservice.azure.com",
		"trustedaccessrolebindings",
		"TrustedAccessRoleBindingList",
		oldVersion,
		newVersion)

	_, err := c.fakeAPIExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	g.Expect(err).To(BeNil())

	err = c.cleaner.Run(context.TODO())
	g.Expect(err).To(BeNil())

	crd, err := c.fakeAPIExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	g.Expect(err).To(BeNil())

	g.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	g.Expect(crd.Status.StoredVersions).ToNot(BeEquivalentTo(definition.Status.StoredVersions))
	g.Expect(crd.Status.StoredVersions).ToNot(ContainElement(oldVersion))
	g.Expect(crd.Status.StoredVersions).To(ContainElement(newVersion))
}

func Test_MigrateDeprecatedCRDResources_DoesNotMigrateBetaVersion_IfStorage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := makeClientSets()

	// This test does not include GA non-storage version, as that would never be possible. Always the latest version would be set to storage
	betaVersion := "v1beta20200601"

	// create CRD
	definition := newCRDWithStoredVersions(betaVersion)
	definition.Spec.Versions = []v1.CustomResourceDefinitionVersion{
		{
			Name:    betaVersion,
			Storage: true,
		},
	}

	ns := newNamespace("test-ns")

	_, err := c.fakeAPIExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	g.Expect(err).To(BeNil())

	// create Namespace
	err = c.fakeClient.Create(context.TODO(), ns)
	g.Expect(err).To(BeNil())

	// create ResourceGroup
	rg := newResourceGroup("test-rg", ns.Name)
	err = c.fakeClient.Create(context.TODO(), rg)
	g.Expect(err).To(BeNil())

	err = c.cleaner.Run(context.TODO())
	g.Expect(err).ToNot(BeNil())

	crd, err := c.fakeAPIExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	g.Expect(err).To(BeNil())

	var updatedRG resources.ResourceGroup
	err = c.fakeClient.Get(context.TODO(), types.NamespacedName{Name: rg.Name, Namespace: rg.Namespace}, &updatedRG)
	g.Expect(err).To(BeNil())
	g.Expect(updatedRG.ResourceVersion).To(BeEquivalentTo(rg.ResourceVersion))

	g.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	g.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
}

func Test_MigrateDeprecatedCRDResources_MigratesBeta_IfNotStorage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := makeClientSets()

	betaVersion := "v1beta20200601"
	gaVersion := "v1api20200601"

	// create CRD
	definition := newCRDWithStoredVersions(betaVersion, gaVersion)
	definition.Spec.Versions = []v1.CustomResourceDefinitionVersion{
		{
			Name: betaVersion,
		},
		{
			Name:    gaVersion,
			Storage: true,
		},
	}

	ns := newNamespace("test-ns")

	_, err := c.fakeAPIExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	g.Expect(err).To(BeNil())

	// create Namespace
	err = c.fakeClient.Create(context.TODO(), ns)
	g.Expect(err).To(BeNil())

	// create ResourceGroup
	rg := newResourceGroup("test-rg", ns.Name)
	err = c.fakeClient.Create(context.TODO(), rg)
	g.Expect(err).To(BeNil())

	err = c.cleaner.Run(context.TODO())
	g.Expect(err).To(BeNil())

	crd, err := c.fakeAPIExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	g.Expect(err).To(BeNil())

	var updatedRG resources.ResourceGroup
	err = c.fakeClient.Get(context.TODO(), types.NamespacedName{Name: rg.Name, Namespace: rg.Namespace}, &updatedRG)
	g.Expect(err).To(BeNil())
	g.Expect(updatedRG.ResourceVersion).ToNot(BeEquivalentTo(rg.ResourceVersion))

	g.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	g.Expect(crd.Status.StoredVersions).ToNot(BeEquivalentTo(definition.Status.StoredVersions))
	g.Expect(crd.Status.StoredVersions).ToNot(ContainElement(betaVersion))
	g.Expect(crd.Status.StoredVersions).To(ContainElement(gaVersion))
}

func Test_CleanDeprecatedCRDVersions_DoesNothing_IfBetaVersionDoesNotExist(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := makeClientSets()

	betaVersion := "v1api20230101storage"

	definition := newCRDWithStoredVersions(betaVersion)

	_, err := c.fakeAPIExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	if err != nil {
		return
	}

	err = c.cleaner.Run(context.TODO())
	g.Expect(err).To(BeNil())

	crd, err := c.fakeAPIExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	g.Expect(err).To(BeNil())

	g.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	g.Expect(crd.Status.StoredVersions).To(BeEquivalentTo(definition.Status.StoredVersions))
	g.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
}

func Test_CleanDeprecatedCRDVersions_ReturnsError_IfGAVersionDoesNotExist(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := makeClientSets()

	betaVersion := "v1beta20230101storage"

	definition := newCRDWithStoredVersions(betaVersion)

	_, err := c.fakeAPIExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	if err != nil {
		return
	}

	err = c.cleaner.Run(context.TODO())
	g.Expect(err).ToNot(BeNil())

	crd, err := c.fakeAPIExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	g.Expect(err).To(BeNil())

	g.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	g.Expect(crd.Status.StoredVersions).To(HaveLen(1))
	g.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
}

func Test_MigrateAndCleanDeprecatedCRDResources_DryRun_NoAction(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fakeAPIExtClient := fake.NewSimpleClientset().ApiextensionsV1()
	fakeClient := fake2.NewClientBuilder().WithScheme(api.CreateScheme()).Build()
	cleanerDryRun := NewCleaner(
		fakeAPIExtClient.CustomResourceDefinitions(),
		fakeClient,
		true, // dry-run
		logr.Discard())

	betaVersion := "v1beta20200601"
	gaVersion := "v1api20200601"

	// create CRD
	definition := newCRDWithStoredVersions(betaVersion, gaVersion)
	definition.Spec.Versions = []v1.CustomResourceDefinitionVersion{
		{
			Name:    betaVersion,
			Storage: true,
		},
		{
			Name: gaVersion,
		},
	}

	ns := newNamespace("test-rg")

	_, err := fakeAPIExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	g.Expect(err).To(BeNil())

	// create Namespace
	err = fakeClient.Create(context.TODO(), ns)
	g.Expect(err).To(BeNil())

	// create ResourceGroup
	rg := newResourceGroup("test-rg", ns.Name)
	err = fakeClient.Create(context.TODO(), rg)
	g.Expect(err).To(BeNil())

	err = cleanerDryRun.Run(context.TODO())
	g.Expect(err).To(BeNil())

	crd, err := fakeAPIExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	g.Expect(err).To(BeNil())

	var updatedRG resources.ResourceGroup
	err = fakeClient.Get(context.TODO(), types.NamespacedName{Name: rg.Name, Namespace: rg.Namespace}, &updatedRG)
	g.Expect(err).To(BeNil())
	g.Expect(updatedRG.ResourceVersion).To(BeEquivalentTo(rg.ResourceVersion))

	g.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	g.Expect(definition.Status.StoredVersions).To(BeEquivalentTo(crd.Status.StoredVersions))
	g.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
	g.Expect(crd.Status.StoredVersions).To(ContainElement(gaVersion))
}

func newNamespace(name string) *corev1.Namespace {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	return ns
}

func newResourceGroup(name, namespace string) *resources.ResourceGroup {
	return &resources.ResourceGroup{
		ObjectMeta: ctrl.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: resources.ResourceGroup_Spec{
			Location: to.Ptr("westus2"),
		},
	}
}

func newCRDWithStoredVersions(versions ...string) *v1.CustomResourceDefinition {
	return newCRDWithStoredVersionsAndName("resources.azure.com", "resourcegroups", "ResourceGroup", versions...)
}

func newCRDWithStoredVersionsAndName(group string, name string, listKind string, versions ...string) *v1.CustomResourceDefinition {
	definition := &v1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s.%s", name, group),
			Labels: map[string]string{
				"app.kubernetes.io/name": "azure-service-operator",
			},
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: group,
			Names: v1.CustomResourceDefinitionNames{
				ListKind: listKind,
			},
		},
		Status: v1.CustomResourceDefinitionStatus{
			StoredVersions: versions,
		},
	}

	return definition
}
