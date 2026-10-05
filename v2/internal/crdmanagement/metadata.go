/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package crdmanagement

import (
	"github.com/samber/lo"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DesiredMetadataEqual returns true when all labels and annotations specified by the desired CRD are present on the existing CRD.
// Additional metadata on the existing CRD is intentionally ignored and preserved during updates.
func DesiredMetadataEqual(existing apiextensions.CustomResourceDefinition, desired apiextensions.CustomResourceDefinition) bool {
	return metadataContains(existing.Labels, desired.Labels) && metadataContains(existing.Annotations, desired.Annotations)
}

func metadataContains(existing map[string]string, desired map[string]string) bool {
	for key, value := range desired {
		// Use a two-value lookup so that a desired empty value is distinguished from a missing key;
		// a single-value lookup would treat an absent key as matching an empty desired value.
		existingValue, found := existing[key]
		if !found || existingValue != value {
			return false
		}
	}

	return true
}

// mergeMetadata preserves existing metadata while allowing desired metadata to take precedence.
func mergeMetadata(existing map[string]string, desired map[string]string) map[string]string {
	return lo.Assign(existing, desired)
}

func mergeCRDMetadata(existing metav1.ObjectMeta, desired *apiextensions.CustomResourceDefinition) {
	desired.Labels = mergeMetadata(existing.Labels, desired.Labels)
	desired.Annotations = mergeMetadata(existing.Annotations, desired.Annotations)
}

// applyCRDLabels merges the user configured labels into the labels of the CRDs loaded from disk.
// Labels reserved for ASO's own use are never overwritten, as doing so would break CRD discovery and
// upgrade detection.
func applyCRDLabels(crds []apiextensions.CustomResourceDefinition, configuredLabels map[string]string) {
	safeLabels := lo.OmitBy(configuredLabels, func(key string, _ string) bool {
		return IsReservedLabel(key)
	})

	for i := range crds {
		crds[i].Labels = mergeMetadata(crds[i].Labels, safeLabels)
	}
}
