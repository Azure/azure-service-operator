/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"
	"testing"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/fake"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fake2 "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

// TODO: Currently we need to create clientsets for each test as they run in parallel and we run into `resource already exists` error.
// TODO: We may require a testing suite re-use the clientsets efficiently.
func getClientSets() (apiextensions.ApiextensionsV1Interface, client.WithWatch, *Cleaner) {
	fakeApiExtClient := fake.NewSimpleClientset().ApiextensionsV1()
	fakeClient := fake2.NewClientBuilder().WithScheme(controllers.CreateScheme()).Build()
	cleaner := NewCleaner(fakeApiExtClient.CustomResourceDefinitions(), fakeClient, false)
	return fakeApiExtClient, fakeClient, cleaner
}

func Test_CleanDeprecatedCRDVersions_CleansAlphaVersion_IfExists(t *testing.T) {
	t.Parallel()

	fakeApiExtClient, _, cleaner := getClientSets()

	asserter := NewGomegaWithT(t)

	alphaVersion := "v1alpha1api20200601"
	betaVersion := "v1beta20200601"

	definition := newCRDWithStoredVersions("v1alpha1api20200601", "v1beta20200601")

	_, err := fakeApiExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	asserter.Expect(err).To(BeNil())

	err = cleaner.Run(context.TODO())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(crd.Status.StoredVersions).ToNot(BeEquivalentTo(definition.Status.StoredVersions))
	asserter.Expect(crd.Status.StoredVersions).ToNot(ContainElement(alphaVersion))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
}

func Test_MigrateDeprecatedCRDResources_MigrateAlphaVersion_IfStorage(t *testing.T) {
	t.Parallel()
	asserter := NewGomegaWithT(t)

	fakeApiExtClient, fakeClient, cleaner := getClientSets()

	alphaVersion := "v1alpha1api20200601"
	betaVersion := "v1beta20200601"

	// create CRD
	definition := newCRDWithStoredVersions(alphaVersion, betaVersion)
	definition.Spec.Versions = []v1.CustomResourceDefinitionVersion{
		{
			Name:    alphaVersion,
			Storage: true,
		},
		{
			Name: betaVersion,
		},
	}

	ns := newNamespace("test-ns")

	_, err := fakeApiExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	asserter.Expect(err).To(BeNil())

	// create Namespace
	err = fakeClient.Create(context.TODO(), ns)
	asserter.Expect(err).To(BeNil())

	// create ResourceGroup
	rg := newResourceGroup("test-rg", ns.Name)
	err = fakeClient.Create(context.TODO(), rg)
	asserter.Expect(err).To(BeNil())

	err = cleaner.Run(context.TODO())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	var updatedRG resources.ResourceGroup
	err = fakeClient.Get(context.TODO(), types.NamespacedName{Name: rg.Name, Namespace: rg.Namespace}, &updatedRG)
	asserter.Expect(err).To(BeNil())
	asserter.Expect(updatedRG.ResourceVersion).ToNot(BeEquivalentTo(rg.ResourceVersion))

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(crd.Status.StoredVersions).ToNot(BeEquivalentTo(definition.Status.StoredVersions))
	asserter.Expect(crd.Status.StoredVersions).ToNot(ContainElement(alphaVersion))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
}

func Test_MigrateDeprecatedCRDResources_NoMigration_IfNotStorage(t *testing.T) {
	t.Parallel()
	asserter := NewGomegaWithT(t)

	fakeApiExtClient, fakeClient, cleaner := getClientSets()

	alphaVersion := "v1alpha1api20200601"
	betaVersion := "v1beta20200601"

	// create CRD
	definition := newCRDWithStoredVersions(alphaVersion, betaVersion)
	definition.Spec.Versions = []v1.CustomResourceDefinitionVersion{
		{
			Name: alphaVersion,
		},
		{
			Name:    betaVersion,
			Storage: true,
		},
	}

	ns := newNamespace("test-ns")

	_, err := fakeApiExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	asserter.Expect(err).To(BeNil())

	// create Namespace
	err = fakeClient.Create(context.TODO(), ns)

	// create ResourceGroup
	rg := newResourceGroup("test-rg", ns.Name)
	err = fakeClient.Create(context.TODO(), rg)
	asserter.Expect(err).To(BeNil())

	err = cleaner.Run(context.TODO())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	var updatedRG resources.ResourceGroup
	err = fakeClient.Get(context.TODO(), types.NamespacedName{Name: rg.Name, Namespace: rg.Namespace}, &updatedRG)
	asserter.Expect(err).To(BeNil())
	asserter.Expect(updatedRG.ResourceVersion).To(BeEquivalentTo(rg.ResourceVersion))

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(crd.Status.StoredVersions).ToNot(BeEquivalentTo(definition.Status.StoredVersions))
	asserter.Expect(crd.Status.StoredVersions).ToNot(ContainElement(alphaVersion))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
}

func Test_CleanDeprecatedCRDVersions_DoesNothing_IfAlphaVersionDoesNotExist(t *testing.T) {
	t.Parallel()
	asserter := NewGomegaWithT(t)

	fakeApiExtClient, _, cleaner := getClientSets()

	betaVersion := "v1beta20230101storage"

	definition := newCRDWithStoredVersions(betaVersion)

	_, err := fakeApiExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	if err != nil {
		return
	}

	err = cleaner.Run(context.TODO())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(crd.Status.StoredVersions).To(BeEquivalentTo(definition.Status.StoredVersions))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
}

func Test_CleanDeprecatedCRDVersions_DoesNothing_BetaVersionDoesNotExist(t *testing.T) {
	t.Parallel()
	asserter := NewGomegaWithT(t)

	fakeApiExtClient, _, cleaner := getClientSets()

	alphaVersion := "v1alpha1api20230101storage"

	definition := newCRDWithStoredVersions(alphaVersion)

	_, err := fakeApiExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	if err != nil {
		return
	}

	err = cleaner.Run(context.TODO())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(crd.Status.StoredVersions).To(HaveLen(1))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(alphaVersion))
}

func Test_MigrateAndCleanDeprecatedCRDResources_DryRun_NoAction(t *testing.T) {
	t.Parallel()
	asserter := NewGomegaWithT(t)

	fakeApiExtClient := fake.NewSimpleClientset().ApiextensionsV1()
	fakeClient := fake2.NewClientBuilder().WithScheme(controllers.CreateScheme()).Build()
	cleanerDryRun := NewCleaner(fakeApiExtClient.CustomResourceDefinitions(), fakeClient, true)

	alphaVersion := "v1alpha1api20200601"
	betaVersion := "v1beta20200601"

	// create CRD
	definition := newCRDWithStoredVersions(alphaVersion, betaVersion)
	definition.Spec.Versions = []v1.CustomResourceDefinitionVersion{
		{
			Name:    alphaVersion,
			Storage: true,
		},
		{
			Name: betaVersion,
		},
	}

	ns := newNamespace("test-rg")

	_, err := fakeApiExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	asserter.Expect(err).To(BeNil())

	// create Namespace
	err = fakeClient.Create(context.TODO(), ns)
	asserter.Expect(err).To(BeNil())

	// create ResourceGroup
	rg := newResourceGroup("test-rg", ns.Name)
	err = fakeClient.Create(context.TODO(), rg)
	asserter.Expect(err).To(BeNil())

	err = cleanerDryRun.Run(context.TODO())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), definition.Name, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	var updatedRG resources.ResourceGroup
	err = fakeClient.Get(context.TODO(), types.NamespacedName{Name: rg.Name, Namespace: rg.Namespace}, &updatedRG)
	asserter.Expect(err).To(BeNil())
	asserter.Expect(updatedRG.ResourceVersion).To(BeEquivalentTo(rg.ResourceVersion))

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(definition.Status.StoredVersions).To(BeEquivalentTo(crd.Status.StoredVersions))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(alphaVersion))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
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
		Spec: resources.ResourceGroupSpec{
			Location: to.StringPtr("westus2"),
		},
	}
}

func newCRDWithStoredVersions(versions ...string) *v1.CustomResourceDefinition {
	definition := &v1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "resourcegroups.resources.azure.com",
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: "resources.azure.com",
			Names: v1.CustomResourceDefinitionNames{
				ListKind: "ResourceGroup",
			},
		},

		Status: v1.CustomResourceDefinitionStatus{
			StoredVersions: versions,
		},
	}
	return definition
}
