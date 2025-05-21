/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package labels_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr/testr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	containerservice "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901"
	"github.com/Azure/azure-service-operator/v2/pkg/common/labels"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func TestSetOwnerNameLabel(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	err := containerservice.AddToScheme(scheme)
	g.Expect(err).To(Succeed())

	cases := []struct {
		name     string
		input    *containerservice.ManagedCluster
		expected string
	}{
		{
			name: "short label is saved",
			input: &containerservice.ManagedCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-mc",
				},
				Spec: containerservice.ManagedCluster_Spec{
					Owner: &genruntime.KnownResourceReference{
						Name: "test-resourcegroup",
					},
				},
			},
			expected: "test-resourcegroup",
		},
		{
			name: "long label is truncated",
			input: &containerservice.ManagedCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-mc",
				},
				Spec: containerservice.ManagedCluster_Spec{
					Owner: &genruntime.KnownResourceReference{
						Name: "this-is-a-long-test-rg-1234567890123456789012345678901234567890123456789012345678901234567890",
					},
				},
			},
			expected: "this-is-a-long-test-rg-1234567890123456789012345678901234567890",
		},
		{
			name: "ARM id owner name is not saved",
			input: &containerservice.ManagedCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-mc",
				},
				Spec: containerservice.ManagedCluster_Spec{
					Owner: &genruntime.KnownResourceReference{
						ARMID: "myarmid",
					},
				},
			},
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g = NewGomegaWithT(t)
			log := testr.New(t)

			labels.SetOwnerNameLabel(log, c.input)
			if c.expected == "" {
				g.Expect(c.input.GetLabels()).NotTo(HaveKey(labels.OwnerNameLabel))
			} else {
				g.Expect(c.input.GetLabels()).To(HaveKeyWithValue(labels.OwnerNameLabel, c.expected))
			}
		})
	}
}

type testArbitraryOwnerGVKResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	ownerGK           metav1.GroupKind
	ownerName         string
}

var _ genruntime.ARMMetaObject = &testArbitraryOwnerGVKResource{}

func (t *testArbitraryOwnerGVKResource) DeepCopyObject() runtime.Object {
	// TODO implement me
	panic("implement me")
}

func (t *testArbitraryOwnerGVKResource) GetConditions() conditions.Conditions {
	// TODO implement me
	panic("implement me")
}

func (t *testArbitraryOwnerGVKResource) SetConditions(conditions conditions.Conditions) {
	// TODO implement me
	panic("implement me")
}

func (t *testArbitraryOwnerGVKResource) Owner() *genruntime.ResourceReference {
	return &genruntime.ResourceReference{
		Name:  t.ownerName,
		Group: t.ownerGK.Group,
		Kind:  t.ownerGK.Kind,
	}
}

func (t *testArbitraryOwnerGVKResource) GetSupportedOperations() []genruntime.ResourceOperation {
	// TODO implement me
	panic("implement me")
}

func (t *testArbitraryOwnerGVKResource) AzureName() string {
	// TODO implement me
	panic("implement me")
}

func (t *testArbitraryOwnerGVKResource) GetType() string {
	// TODO implement me
	panic("implement me")
}

func (t *testArbitraryOwnerGVKResource) GetResourceScope() genruntime.ResourceScope {
	// TODO implement me
	panic("implement me")
}

func (t *testArbitraryOwnerGVKResource) GetAPIVersion() string {
	// TODO implement me
	panic("implement me")
}

func (t *testArbitraryOwnerGVKResource) GetSpec() genruntime.ConvertibleSpec {
	// TODO implement me
	panic("implement me")
}

func (t *testArbitraryOwnerGVKResource) GetStatus() genruntime.ConvertibleStatus {
	// TODO implement me
	panic("implement me")
}

func (t *testArbitraryOwnerGVKResource) NewEmptyStatus() genruntime.ConvertibleStatus {
	// TODO implement me
	panic("implement me")
}

func (t *testArbitraryOwnerGVKResource) SetStatus(status genruntime.ConvertibleStatus) error {
	// TODO implement me
	panic("implement me")
}

func TestSetOwnerGroupKindLabel(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	err := containerservice.AddToScheme(scheme)
	g.Expect(err).To(Succeed())

	cases := []struct {
		name     string
		input    genruntime.ARMMetaObject
		expected string
	}{
		{
			name: "short label is saved",
			input: &containerservice.ManagedCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-mc",
				},
				Spec: containerservice.ManagedCluster_Spec{
					Owner: &genruntime.KnownResourceReference{
						Name: "test-resourcegroup",
					},
				},
			},
			expected: "ResourceGroup.resources.azure.com",
		},
		{
			name: "long label is truncated",
			input: &testArbitraryOwnerGVKResource{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-mc",
				},
				ownerGK: metav1.GroupKind{
					Group: "thisgroupisextremelylonglikesolongitcausesissues.azure.com",
					Kind:  "MyResource",
				},
				ownerName: "myresource",
			},
			expected: "MyResource.thisgroupisextremelylonglikesolongitcausesissues.azu",
		},
		{
			name: "ARM id owner GK is not saved",
			input: &containerservice.ManagedCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-mc",
				},
				Spec: containerservice.ManagedCluster_Spec{
					Owner: &genruntime.KnownResourceReference{
						ARMID: "myarmid",
					},
				},
			},
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g = NewGomegaWithT(t)
			log := testr.New(t)

			labels.SetOwnerGroupKindLabel(log, c.input)
			if c.expected == "" {
				g.Expect(c.input.GetLabels()).NotTo(HaveKey(labels.OwnerGroupKindLabel))
			} else {
				g.Expect(c.input.GetLabels()).To(HaveKeyWithValue(labels.OwnerGroupKindLabel, c.expected))
			}
		})
	}
}

func TestSetOwnerUIDLabel(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	err := containerservice.AddToScheme(scheme)
	g.Expect(err).To(Succeed())

	cases := []struct {
		name     string
		input    *containerservice.ManagedCluster
		expected string
	}{
		{
			name: "UID is saved when owner reference is found",
			input: &containerservice.ManagedCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-mc",
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion: "resources.azure.com/v1",
							Kind:       "ResourceGroup",
							Name:       "test-rg",
							UID:        "12345678-1234-1234-1234-123456789012",
						},
					},
				},
				Spec: containerservice.ManagedCluster_Spec{
					Owner: &genruntime.KnownResourceReference{
						Name: "test-rg",
					},
				},
			},
			expected: "12345678-1234-1234-1234-123456789012",
		},
		{
			name: "UID is not saved when owner reference is not found",
			input: &containerservice.ManagedCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-account",
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion: "something.else.com/v1",
							Kind:       "SomethingElse",
							Name:       "test-rg",
							UID:        "12345678-1234-1234-1234-123456789012",
						},
					},
				},
				Spec: containerservice.ManagedCluster_Spec{
					Owner: &genruntime.KnownResourceReference{
						Name: "test-rg",
					},
				},
			},
			expected: "",
		},
		{
			name: "UID is not saved when multiple owner references but one matches",
			input: &containerservice.ManagedCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-account",
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion: "something.else.com/v1",
							Kind:       "SomethingElse",
							Name:       "test-rg",
							UID:        "12345678-1234-1234-1234-123456789010",
						},
						{
							APIVersion: "resources.azure.com/v1",
							Kind:       "ResourceGroup",
							Name:       "test-rg",
							UID:        "12345678-1234-1234-1234-123456789011",
						},
					},
				},
				Spec: containerservice.ManagedCluster_Spec{
					Owner: &genruntime.KnownResourceReference{
						Name: "test-rg",
					},
				},
			},
			expected: "12345678-1234-1234-1234-123456789011",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g = NewGomegaWithT(t)

			labels.SetOwnerUIDLabel(c.input)
			if c.expected == "" {
				g.Expect(c.input.GetLabels()).NotTo(HaveKey(labels.OwnerUIDLabel))
			} else {
				g.Expect(c.input.GetLabels()).To(HaveKeyWithValue(labels.OwnerUIDLabel, c.expected))
			}
		})
	}
}
