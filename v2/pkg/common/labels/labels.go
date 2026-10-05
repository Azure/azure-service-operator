// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package labels

import (
	"strings"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/validate/content"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const (
	OwnerNameLabel             = "serviceoperator.azure.com/owner-name"
	OwnerGroupKindLabel        = "serviceoperator.azure.com/owner-group-kind"
	OwnerUIDLabel              = "serviceoperator.azure.com/owner-uid"
	LastReconciledVersionLabel = "serviceoperator.azure.com/last-reconciled-version"
)

// ServiceOperatorLabelPrefix is the label domain reserved for ASO's own use.
const ServiceOperatorLabelPrefix = "serviceoperator.azure.com/"

// Labels applied to the CRDs ASO manages. These values must match the values injected by config/crd/labels.yaml.
const (
	// ServiceOperatorVersionLabelOld is the legacy label the CRDs have on them containing the ASO version.
	ServiceOperatorVersionLabelOld = ServiceOperatorLabelPrefix + "version"
	// ServiceOperatorVersionLabel is the label the CRDs have on them containing the ASO version.
	ServiceOperatorVersionLabel = "app.kubernetes.io/version"
	// ServiceOperatorAppLabel is the label used to identify the CRDs managed by ASO.
	ServiceOperatorAppLabel = "app.kubernetes.io/name"
	// ServiceOperatorAppValue is the value of ServiceOperatorAppLabel on CRDs managed by ASO.
	ServiceOperatorAppValue = "azure-service-operator"
)

// SetOwnerNameLabel sets the owner name label on the given object, or truncates it if it exceeds the character limit.
func SetOwnerNameLabel(logger logr.Logger, obj genruntime.ARMMetaObject) {
	if obj.Owner() != nil && obj.Owner().Name != "" {
		ownerName := obj.Owner().Name
		ownerName, truncated := truncateLabelValue(ownerName)
		if truncated {
			logger.V(Status).Info("WARNING: Owner name label truncated to fit Kubernetes label limits", "ownerName", ownerName)
		}
		genruntime.AddLabel(obj, OwnerNameLabel, ownerName)
	}
}

// SetOwnerGroupKindLabel sets the owner group kind label on the given object, or truncates it if it exceeds the character limit.
func SetOwnerGroupKindLabel(logger logr.Logger, obj genruntime.ARMMetaObject) {
	if obj.Owner() != nil && obj.Owner().IsKubernetesReference() {
		groupKind := obj.Owner().GroupKind().String()
		groupKind, truncated := truncateLabelValue(groupKind)
		if truncated {
			logger.V(Status).Info("WARNING: GroupKind name truncated to fit Kubernetes label limits", "groupKind", groupKind)
		}

		genruntime.AddLabel(obj, OwnerGroupKindLabel, groupKind)
	}
}

func truncateLabelValue(value string) (string, bool) {
	if len(value) <= content.LabelValueMaxLength {
		return value, false
	}

	// A simple truncation can leave trailing non-alphanumeric characters, which are invalid for Kubernetes label values. See #5734.
	return strings.TrimRightFunc(
			value[:content.LabelValueMaxLength],
			func(r rune) bool {
				return !isASCIIAlphaNumeric(r)
			}),
		true
}

func isASCIIAlphaNumeric(r rune) bool {
	return r >= 'a' && r <= 'z' ||
		r >= 'A' && r <= 'Z' ||
		r >= '0' && r <= '9'
}

// SetOwnerUIDLabel sets the owner UID label on the given object if the owner reference is found in the object's owner references.
func SetOwnerUIDLabel(obj genruntime.ARMMetaObject) {
	ownerRefs := obj.GetOwnerReferences()
	if len(ownerRefs) == 0 || obj.Owner() == nil || !obj.Owner().IsKubernetesReference() {
		return
	}

	groupKind := obj.Owner().GroupKind()

	for _, ref := range ownerRefs {
		ownerGroup := strings.Split(ref.APIVersion, "/")[0]
		if ref.Kind == groupKind.Kind && ownerGroup == groupKind.Group {
			// Set the label with the UID of the owner
			genruntime.AddLabel(obj, OwnerUIDLabel, string(ref.UID))
			return
		}
	}
}
