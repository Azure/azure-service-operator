// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package crdmanagement

import (
	"testing"

	. "github.com/onsi/gomega"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	asolabels "github.com/Azure/azure-service-operator/v2/pkg/common/labels"
)

func TestMergeCRDMetadata(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	desired := &apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"managed-label": "desired",
			},
			Annotations: map[string]string{
				"managed-annotation": "desired",
			},
		},
	}
	existing := metav1.ObjectMeta{
		Labels: map[string]string{
			"external-label": "preserved",
			"managed-label":  "old",
		},
		Annotations: map[string]string{
			"external-annotation": "preserved",
			"managed-annotation":  "old",
		},
	}

	mergeCRDMetadata(existing, desired)

	g.Expect(desired.Labels).To(HaveLen(2))
	g.Expect(desired.Labels).To(HaveKeyWithValue("external-label", "preserved"))
	g.Expect(desired.Labels).To(HaveKeyWithValue("managed-label", "desired"))
	g.Expect(desired.Annotations).To(HaveLen(2))
	g.Expect(desired.Annotations).To(HaveKeyWithValue("external-annotation", "preserved"))
	g.Expect(desired.Annotations).To(HaveKeyWithValue("managed-annotation", "desired"))
}

func TestDesiredMetadataEqual(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		existing metav1.ObjectMeta
		desired  metav1.ObjectMeta
		want     bool
	}{
		{
			name: "desired metadata is present alongside unknown metadata",
			existing: metav1.ObjectMeta{
				Labels:      map[string]string{"configured": "value", "external": "preserved"},
				Annotations: map[string]string{"configured": "value", "external": "preserved"},
			},
			desired: metav1.ObjectMeta{
				Labels:      map[string]string{"configured": "value"},
				Annotations: map[string]string{"configured": "value"},
			},
			want: true,
		},
		{
			name:     "configured label is missing",
			existing: metav1.ObjectMeta{},
			desired:  metav1.ObjectMeta{Labels: map[string]string{"configured": "value"}},
			want:     false,
		},
		{
			name:     "managed annotation has a different value",
			existing: metav1.ObjectMeta{Annotations: map[string]string{"configured": "old"}},
			desired:  metav1.ObjectMeta{Annotations: map[string]string{"configured": "value"}},
			want:     false,
		},
		{
			name:     "empty valued label is missing",
			existing: metav1.ObjectMeta{Labels: map[string]string{"external": "preserved"}},
			desired:  metav1.ObjectMeta{Labels: map[string]string{"configured": ""}},
			want:     false,
		},
		{
			name:     "empty valued annotation is missing",
			existing: metav1.ObjectMeta{},
			desired:  metav1.ObjectMeta{Annotations: map[string]string{"configured": ""}},
			want:     false,
		},
		{
			name:     "empty valued label is present",
			existing: metav1.ObjectMeta{Labels: map[string]string{"configured": ""}},
			desired:  metav1.ObjectMeta{Labels: map[string]string{"configured": ""}},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			existing := apiextensions.CustomResourceDefinition{ObjectMeta: tt.existing}
			desired := apiextensions.CustomResourceDefinition{ObjectMeta: tt.desired}
			g.Expect(DesiredMetadataEqual(existing, desired)).To(Equal(tt.want))
		})
	}
}

func TestApplyCRDLabels(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	crds := []apiextensions.CustomResourceDefinition{
		{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"bundled": "preserved", "configured": "old"}}},
		{},
	}

	applyCRDLabels(crds, map[string]string{"configured": "new"})

	g.Expect(crds[0].Labels).To(HaveLen(2))
	g.Expect(crds[0].Labels).To(HaveKeyWithValue("bundled", "preserved"))
	g.Expect(crds[0].Labels).To(HaveKeyWithValue("configured", "new"))
	g.Expect(crds[1].Labels).To(HaveLen(1))
	g.Expect(crds[1].Labels).To(HaveKeyWithValue("configured", "new"))
}

func TestApplyCRDLabelsIgnoresReservedLabels(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	bundledLabels := map[string]string{
		asolabels.ServiceOperatorAppLabel:     asolabels.ServiceOperatorAppValue,
		asolabels.ServiceOperatorVersionLabel: "v2.0.0",
	}
	crds := []apiextensions.CustomResourceDefinition{
		{ObjectMeta: metav1.ObjectMeta{Labels: bundledLabels}},
	}

	configuredLabels := map[string]string{
		asolabels.ServiceOperatorAppLabel:        "hijacked",
		asolabels.ServiceOperatorVersionLabel:    "pinned",
		asolabels.ServiceOperatorVersionLabelOld: "pinned",
		"example.com/owner":                      "aso",
	}
	configuredLabels[asolabels.ServiceOperatorLabelPrefix+"anything"] = "hijacked"

	applyCRDLabels(crds, configuredLabels)

	g.Expect(crds[0].Labels).To(HaveLen(3))
	g.Expect(crds[0].Labels).To(HaveKeyWithValue(asolabels.ServiceOperatorAppLabel, asolabels.ServiceOperatorAppValue))
	g.Expect(crds[0].Labels).To(HaveKeyWithValue(asolabels.ServiceOperatorVersionLabel, "v2.0.0"))
	g.Expect(crds[0].Labels).To(HaveKeyWithValue("example.com/owner", "aso"))
}

// TestApplyCRDLabelsRemovalIsNotPropagated documents that removing a label from --crd-labels does not
// remove it from CRDs already in the cluster; mergeCRDMetadata intentionally preserves existing metadata.
func TestApplyCRDLabelsRemovalIsNotPropagated(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	existing := metav1.ObjectMeta{Labels: map[string]string{"bundled": "preserved", "removed": "old"}}
	desired := apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"bundled": "preserved"}},
	}

	mergeCRDMetadata(existing, &desired)

	g.Expect(desired.Labels).To(HaveLen(2))
	g.Expect(desired.Labels).To(HaveKeyWithValue("bundled", "preserved"))
	g.Expect(desired.Labels).To(HaveKeyWithValue("removed", "old"))
}

// TestDesiredMetadataEqualDetectsCertManagerNamespaceDrift pins the behaviour that the annotation half of
// DesiredMetadataEqual is load bearing: ASO ships cert-manager.io/inject-ca-from on every conversion webhook
// CRD and fixCRDNamespace rewrites its namespace at load time, so a redeploy into a different namespace must
// be detected as a metadata difference even when the ASO version is unchanged.
func TestDesiredMetadataEqualDetectsCertManagerNamespaceDrift(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	const annotation = "cert-manager.io/inject-ca-from"
	existing := apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{annotation: "old-namespace/azureserviceoperator-serving-cert"},
		},
	}
	desired := apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{annotation: "new-namespace/azureserviceoperator-serving-cert"},
		},
	}

	g.Expect(DesiredMetadataEqual(existing, desired)).To(BeFalse())
	g.Expect(DesiredMetadataEqual(desired, desired)).To(BeTrue())
}
