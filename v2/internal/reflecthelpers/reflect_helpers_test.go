/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reflecthelpers_test

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//nolint:staticcheck // ignoring deprecation (SA1019) to unblock CI builds
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"

	. "github.com/onsi/gomega"
)

type ResourceWithReferences struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceWithReferencesSpec `json:"spec,omitempty"`
}

type ResourceWithReferencesSpec struct {
	Owner genruntime.KnownResourceReference `json:"owner"`

	AzureName string `json:"azureName"`

	Ref *ResourceReference `json:"ref,omitempty"`

	Location string `json:"location,omitempty"`
}

type ResourceReference struct {
	Reference genruntime.ResourceReference `armReference:"Id" json:"reference"`
}

func Test_FindReferences(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := genruntime.ResourceReference{ARMID: "test"}

	res := ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
			Ref: &ResourceReference{
				Reference: ref,
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	refs, err := reflecthelpers.FindResourceReferences(res)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(refs).To(HaveLen(1))
	g.Expect(refs).To(HaveKey(ref))
}
