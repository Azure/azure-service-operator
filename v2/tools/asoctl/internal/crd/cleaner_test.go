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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	fake2 "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var fakeApiExtClient = fake.NewSimpleClientset().ApiextensionsV1()
var fakeClient = fake2.NewClientBuilder().WithScheme(controllers.CreateScheme()).Build()

var cleaner = NewCleaner(fakeApiExtClient.CustomResourceDefinitions(), fakeClient, false)

func Test_CleanDeprecatedCRDVersions_CleansAlphaVersion_IfExists(t *testing.T) {
	t.Parallel()
	asserter := NewGomegaWithT(t)

	crdName := "resourcegroups.resources.azure.com"
	alphaVersion := "v1alpha1api20200601"
	betaVersion := "v1beta20200601"

	definition := crdWithStoredVersions(crdName, alphaVersion, betaVersion)

	_, err := fakeApiExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	asserter.Expect(err).To(BeNil())

	err = cleaner.Run(context.TODO())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), crdName, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(crd.Status.StoredVersions).ToNot(HaveLen(2))
	asserter.Expect(crd.Status.StoredVersions).ToNot(ContainElement(alphaVersion))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
}

func Test_MigrateDeprecatedCRDResources_MigrateAlphaVersion_IfStorage(t *testing.T) {
	t.Parallel()
	asserter := NewGomegaWithT(t)

	crdName := "resourcegroups.resources.azure.com"
	alphaVersion := "v1alpha1api20200601"
	betaVersion := "v1beta20200601"

	// create CRD
	definition := crdWithStoredVersions(crdName, alphaVersion, betaVersion)
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

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), crdName, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	var updatedRG resources.ResourceGroup
	err = fakeClient.Get(context.TODO(), types.NamespacedName{Name: rg.Name, Namespace: rg.Namespace}, &updatedRG)
	asserter.Expect(err).To(BeNil())
	asserter.Expect(updatedRG.ResourceVersion).ToNot(BeEquivalentTo(rg.ResourceVersion))

	println(updatedRG.APIVersion)

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(crd.Status.StoredVersions).ToNot(HaveLen(2))
	asserter.Expect(crd.Status.StoredVersions).ToNot(ContainElement(alphaVersion))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
}

func Test_MigrateDeprecatedCRDResources_NoMigration_IfNotStorage(t *testing.T) {
	t.Parallel()
	asserter := NewGomegaWithT(t)

	crdName := "resourcegroups.resources.azure.com"
	alphaVersion := "v1alpha1api20200601"
	betaVersion := "v1beta20200601"

	// create CRD
	definition := crdWithStoredVersions(crdName, alphaVersion, betaVersion)
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
	asserter.Expect(err).To(BeNil())

	// create ResourceGroup
	rg := newResourceGroup("test-rg", ns.Name)
	err = fakeClient.Create(context.TODO(), rg)
	asserter.Expect(err).To(BeNil())

	err = cleaner.Run(context.TODO())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), crdName, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	var updatedRG resources.ResourceGroup
	err = fakeClient.Get(context.TODO(), types.NamespacedName{Name: rg.Name, Namespace: rg.Namespace}, &updatedRG)
	asserter.Expect(err).To(BeNil())
	asserter.Expect(updatedRG.ResourceVersion).To(BeEquivalentTo(rg.ResourceVersion))

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(crd.Status.StoredVersions).ToNot(HaveLen(2))
	asserter.Expect(crd.Status.StoredVersions).ToNot(ContainElement(alphaVersion))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
}

func Test_CleanDeprecatedCRDVersions_DoesNothing_IfAlphaVersionDoesNotExist(t *testing.T) {
	t.Parallel()

	asserter := NewGomegaWithT(t)
	crdName := "foo.bar.azure.com"
	betaVersion := "v1beta20230101storage"

	definition := crdWithStoredVersions(crdName, betaVersion)

	_, err2 := fakeApiExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	if err2 != nil {
		return
	}

	err := cleaner.Run(context.TODO())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), crdName, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(crd.Status.StoredVersions).To(HaveLen(1))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(betaVersion))
}

func Test_CleanDeprecatedCRDVersions_DoesNothing_BetaVersionDoesNotExist(t *testing.T) {
	t.Parallel()

	asserter := NewGomegaWithT(t)
	crdName := "foo.bar.azure.com"
	alphaVersion := "v1alpha1api20230101storage"

	definition := crdWithStoredVersions(crdName, alphaVersion)

	_, err2 := fakeApiExtClient.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	if err2 != nil {
		return
	}

	err := cleaner.Run(context.TODO())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), crdName, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(crd.Status.StoredVersions).To(HaveLen(1))
	asserter.Expect(crd.Status.StoredVersions).To(ContainElement(alphaVersion))
}

func crdWithStoredVersions(crdName string, versions ...string) *v1.CustomResourceDefinition {
	definition := &v1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: crdName,
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

func Test_MigrateAndCleanDeprecatedCRDResources_DryRun_NoAction(t *testing.T) {
	t.Parallel()
	asserter := NewGomegaWithT(t)

	cleanerDryRun := NewCleaner(fakeApiExtClient.CustomResourceDefinitions(), fakeClient, true)

	crdName := "resourcegroups.resources.azure.com"
	alphaVersion := "v1alpha1api20200601"
	betaVersion := "v1beta20200601"

	// create CRD
	definition := crdWithStoredVersions(crdName, alphaVersion, betaVersion)
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

	crd, err := fakeApiExtClient.CustomResourceDefinitions().Get(context.TODO(), crdName, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

	var updatedRG resources.ResourceGroup
	err = fakeClient.Get(context.TODO(), types.NamespacedName{Name: rg.Name, Namespace: rg.Namespace}, &updatedRG)
	asserter.Expect(err).To(BeNil())
	asserter.Expect(updatedRG.ResourceVersion).To(BeEquivalentTo(rg.ResourceVersion))

	println(updatedRG.APIVersion)

	asserter.Expect(crd.Status.StoredVersions).ToNot(BeNil())
	asserter.Expect(definition.Status.StoredVersions).To(BeEquivalentTo(crd.Status.StoredVersions))
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
