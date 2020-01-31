package ctrlutil

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// IsSucceeded returns true if the Status.ProvisioningState == Succeeded
func IsSucceeded(obj *unstructured.Unstructured) (bool, error) {
	ready, found, err := unstructured.NestedString(obj.Object, "status", "provisioningState")
	if err != nil {
		return false, errors.Wrapf(err, "failed to determine %v %q readiness",
			obj.GroupVersionKind(), obj.GetName())
	}
	return ready == "Succeeded" && found, nil
}
