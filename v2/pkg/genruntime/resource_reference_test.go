/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var validARMIDRef = genruntime.ResourceReference{ARMID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/microsoft.compute/VirtualMachine/myvm"}
var validKubRef = genruntime.ResourceReference{Group: "microsoft.resources.azure.com", Kind: "ResourceGroup", Name: "myrg"}
var invalidRefBothSpecified = genruntime.ResourceReference{Group: "microsoft.resources.azure.com", Kind: "ResourceGroup", Name: "myrg", ARMID: "oops"}
var invalidRefNeitherSpecified = genruntime.ResourceReference{}
var invalidRefIncompleteKubReference = genruntime.ResourceReference{Group: "microsoft.resources.azure.com", Name: "myrg"}

func Test_ResourceReference_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		ref          genruntime.ResourceReference
		errSubstring string
	}{
		{
			name:         "valid ARM reference is valid",
			ref:          validARMIDRef,
			errSubstring: "",
		},
		{
			name:         "valid Kubernetes reference is valid",
			ref:          validKubRef,
			errSubstring: "",
		},
		{
			name:         "both ARM and Kubernetes fields filled out, reference is invalid",
			ref:          invalidRefBothSpecified,
			errSubstring: "'ARMID' field is mutually exclusive with",
		},
		{
			name:         "nothing filled out, reference is invalid",
			ref:          invalidRefNeitherSpecified,
			errSubstring: "at least one of ['ARMID'] or ['Group', 'Kind', 'Namespace', 'Name'] must be set for ResourceReference",
		},
		{
			name:         "incomplete Kubernetes reference is invalid",
			ref:          invalidRefIncompleteKubReference,
			errSubstring: "'Group', 'Kind', 'Namespace', and 'Name' must all be specified",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			err := tt.ref.Validate()
			if tt.errSubstring != "" {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(ContainSubstring(tt.errSubstring))
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}
		})
	}
}

func Test_ResourceReference_IsARMOrKubernetes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		ref          genruntime.ResourceReference
		isARM        bool
		isKubernetes bool
	}{
		{
			name:         "valid ARM reference is ARM",
			ref:          validARMIDRef,
			isARM:        true,
			isKubernetes: false,
		},
		{
			name:         "valid Kubernetes reference is Kubernetes",
			ref:          validKubRef,
			isARM:        false,
			isKubernetes: true,
		},
		{
			name:         "both ARM and Kubernetes fields filled out, reference is neither",
			ref:          invalidRefBothSpecified,
			isARM:        false,
			isKubernetes: false,
		},
		{
			name:         "nothing filled out, reference is neither",
			ref:          invalidRefNeitherSpecified,
			isARM:        false,
			isKubernetes: false,
		},
		{
			name:         "incomplete Kubernetes reference is neither",
			ref:          invalidRefIncompleteKubReference,
			isARM:        false,
			isKubernetes: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(tt.ref.IsDirectARMReference()).To(Equal(tt.isARM))
			g.Expect(tt.ref.IsKubernetesReference()).To(Equal(tt.isKubernetes))
		})
	}
}

func Test_ResourceReference_AsNamespacedReference_ARMReferenceHasNoNamespace(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	namespace := "default"

	namespacedRef := validARMIDRef.AsNamespacedRef(namespace)
	g.Expect(namespacedRef.Namespace).To(BeEmpty())
}

func Test_ResourceReference_AsNamespacedReference_KubernetesReferenceHasNamespace(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	namespace := "default"

	namespacedRef := validKubRef.AsNamespacedRef(namespace)
	g.Expect(namespacedRef.Namespace).To(Equal(namespace))
}
