// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	"encoding/json"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
)

const conversionStashAnnotation = "azure.microsoft.com/convert-stash"

// getStashedAnnotation unmarshals the convert-stash annotation value
// into the target pointer, if the annotation is present. It returns
// whether the annotation was present, and any error from unmarshalling.
func getStashedAnnotation(meta metav1.ObjectMeta, target interface{}) (bool, error) {
	if _, err := conversion.EnforcePtr(target); err != nil {
		return false, errors.Errorf("getStashedAnnotation target must be a pointer, not %T", target)
	}
	if meta.Annotations == nil {
		return false, nil
	}
	value, found := meta.Annotations[conversionStashAnnotation]
	if !found {
		return false, nil
	}
	if err := json.Unmarshal([]byte(value), target); err != nil {
		return false, errors.Wrap(err, "decoding stashed fields")
	}
	return true, nil
}

func setStashedAnnotation(meta *metav1.ObjectMeta, stashValues interface{}) error {
	if meta.Annotations == nil {
		meta.Annotations = make(map[string]string)
	}
	encoded, err := json.Marshal(stashValues)
	if err != nil {
		return errors.Wrap(err, "encoding stashed fields")
	}
	meta.Annotations[conversionStashAnnotation] = string(encoded)
	return nil
}

func clearStashedAnnotation(meta *metav1.ObjectMeta) {
	delete(meta.Annotations, conversionStashAnnotation)
	if len(meta.Annotations) == 0 {
		meta.Annotations = nil
	}
}
