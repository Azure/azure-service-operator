// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package randextensions_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/azure-service-operator/v2/internal/util/randextensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_MakeUniqueOwnerScopedString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		ref          *genruntime.ResourceReference
		objGK        schema.GroupKind
		objNamespace string
		objName      string
		expected     string
	}{
		{
			name:         "nil owner returns empty parent string",
			ref:          nil,
			objGK:        schema.GroupKind{Group: "resources.azure.com", Kind: "ResourceGroup"},
			objNamespace: "default",
			objName:      "myrg",
			// Note that group and kind are backwards for the object here because I typoed the ordering originally... This is OK as we just want a unique seed for a GUID.
			expected: "ResourceGroup/resources.azure.com:default/myrg",
		},
		{
			name:         "GVK-based owner, full owner string included",
			ref:          &genruntime.ResourceReference{Group: "resources.azure.com", Kind: "ResourceGroup", Name: "myrg"},
			objGK:        schema.GroupKind{Group: "authorization.azure.com", Kind: "RoleAssignment"},
			objNamespace: "default",
			objName:      "myroleassignment",
			// Note that group and kind are backwards for the object here because I typoed the ordering originally... This is OK as we just want a unique seed for a GUID.
			expected: "resources.azure.com/ResourceGroup:default/myrg:RoleAssignment/authorization.azure.com:default/myroleassignment",
		},
		{
			name:         "ARM-ID-based owner, ARM-ID included",
			ref:          &genruntime.ResourceReference{ARMID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg"},
			objGK:        schema.GroupKind{Group: "authorization.azure.com", Kind: "RoleAssignment"},
			objNamespace: "default",
			objName:      "myroleassignment",
			// Note that group and kind are backwards for the object here because I typoed the ordering originally... This is OK as we just want a unique seed for a GUID.
			expected: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg:RoleAssignment/authorization.azure.com:default/myroleassignment",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := randextensions.MakeUniqueOwnerScopedString(tt.ref, tt.objGK, tt.objNamespace, tt.objName)
			g.Expect(result).To(Equal(tt.expected))
		})
	}
}

func Test_MakeUniqueOwnerScopedStringLegacy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		ref          *genruntime.ResourceReference
		objGK        schema.GroupKind
		objNamespace string
		objName      string
		expected     string
	}{
		{
			name:         "nil owner returns empty parent string",
			ref:          nil,
			objGK:        schema.GroupKind{Group: "resources.azure.com", Kind: "ResourceGroup"},
			objNamespace: "default",
			objName:      "myrg",
			expected:     "ResourceGroup/resources.azure.com:default/myrg",
		},
		{
			name:         "GVK-based owner, full owner string included",
			ref:          &genruntime.ResourceReference{Group: "resources.azure.com", Kind: "ResourceGroup", Name: "myrg"},
			objGK:        schema.GroupKind{Group: "authorization.azure.com", Kind: "RoleAssignment"},
			objNamespace: "default",
			objName:      "myroleassignment",
			expected:     "resources.azure.com/ResourceGroup:default/myrg:RoleAssignment/authorization.azure.com:default/myroleassignment",
		},
		{
			name:         "ARM-ID-based owner, ARM-ID included",
			ref:          &genruntime.ResourceReference{ARMID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg"},
			objGK:        schema.GroupKind{Group: "authorization.azure.com", Kind: "RoleAssignment"},
			objNamespace: "default",
			objName:      "myroleassignment",
			// Note that group and kind are backwards for the object here because I typoed the ordering originally... This is OK as we just want a unique seed for a GUID.
			expected: "/:default/:RoleAssignment/authorization.azure.com:default/myroleassignment",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := randextensions.MakeUniqueOwnerScopedStringLegacy(tt.ref, tt.objGK, tt.objNamespace, tt.objName)
			g.Expect(result).To(Equal(tt.expected))
		})
	}
}
