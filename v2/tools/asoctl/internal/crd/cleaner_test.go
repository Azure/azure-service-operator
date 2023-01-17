/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/fake"
)

var fakeClientSet = fake.NewSimpleClientset().ApiextensionsV1()

func Test_CleanDeprecatedCRDVersions_CleansAlphaVersion_IfExists(t *testing.T) {
	t.Parallel()

	asserter := NewGomegaWithT(t)
	crdName := "foo.bar.azure.com"
	alphaVersion := "v1alpha1api20230101storage"
	betaVersion := "v1beta20230101storage"

	definition := crdWithStoredVersions(crdName, alphaVersion, betaVersion)

	_, err2 := fakeClientSet.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	if err2 != nil {
		return
	}

	err := CleanDeprecatedCRDVersions(context.TODO(), fakeClientSet.CustomResourceDefinitions())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeClientSet.CustomResourceDefinitions().Get(context.TODO(), crdName, metav1.GetOptions{})
	asserter.Expect(err).To(BeNil())

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

	_, err2 := fakeClientSet.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	if err2 != nil {
		return
	}

	err := CleanDeprecatedCRDVersions(context.TODO(), fakeClientSet.CustomResourceDefinitions())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeClientSet.CustomResourceDefinitions().Get(context.TODO(), crdName, metav1.GetOptions{})
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

	_, err2 := fakeClientSet.CustomResourceDefinitions().Create(context.TODO(), definition, metav1.CreateOptions{})
	if err2 != nil {
		return
	}

	err := CleanDeprecatedCRDVersions(context.TODO(), fakeClientSet.CustomResourceDefinitions())
	asserter.Expect(err).To(BeNil())

	crd, err := fakeClientSet.CustomResourceDefinitions().Get(context.TODO(), crdName, metav1.GetOptions{})
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
		Status: v1.CustomResourceDefinitionStatus{
			StoredVersions: versions,
		},
	}
	return definition
}
