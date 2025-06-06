/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/gomega"
)

func Test_getEntraID_givenAnnotation_returnsExpectedResults(t *testing.T) {
	t.Parallel()

	objWithEntraID := &v1.ObjectMeta{
		Annotations: map[string]string{
			genruntime.ResourceIDAnnotation: "entra-id",
		},
	}

	objWithoutEntraID := &v1.ObjectMeta{
		Annotations: map[string]string{
			genruntime.ChildResourceIDOverrideAnnotation: "child-entra-id",
		},
	}

	cases := map[string]struct {
		obj      v1.Object
		expected string
		found    bool
	}{
		"entra-id": {
			obj:      objWithEntraID,
			expected: "entra-id",
			found:    true,
		},
		"non-existent": {
			obj:      objWithoutEntraID,
			expected: "",
			found:    false,
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result, found := getEntraID(c.obj)
			if found {
				g.Expect(result).To(Equal(c.expected), "Expected annotation value to match")
			} else {
				g.Expect(result).To(BeEmpty(), "Expected annotation value to be empty")
			}
		})
	}
}

func Test_setEntraID_givenAnnotation_returnsExpectedResults(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup   func(obj *v1.ObjectMeta)
		entraID string
	}{
		"set initial value": {
			entraID: "entra-id",
		},
		"update existing value": {
			setup: func(obj *v1.ObjectMeta) {
				obj.Annotations = map[string]string{
					genruntime.ResourceIDAnnotation: "old-entra-id",
				}
			},
			entraID: "new-entra-id",
		},
		"remove existing value": {
			setup: func(obj *v1.ObjectMeta) {
				obj.Annotations = map[string]string{
					genruntime.ResourceIDAnnotation: "old-entra-id",
				}
			},
			entraID: "",
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			obj := &v1.ObjectMeta{
				Annotations: map[string]string{},
			}

			if c.setup != nil {
				c.setup(obj)
			}

			setEntraID(obj, c.entraID)

			if c.entraID != "" {
				g.Expect(obj.GetAnnotations()).To(HaveKeyWithValue(genruntime.ResourceIDAnnotation, c.entraID))
			} else {
				g.Expect(obj.GetAnnotations()).ToNot(HaveKey(genruntime.ResourceIDAnnotation))
			}
		})
	}
}
