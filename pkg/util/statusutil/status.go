/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package statusutil

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func GetProvisioningState(obj runtime.Object) (string, error) {
	unObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return "", fmt.Errorf("unable to convert to unstructured with: %w", err)
	}

	status, _, err := unstructured.NestedString(unObj, "status", "provisioningState")
	if err != nil {
		return "", fmt.Errorf("unable to fetch nested string status with: %w", err)
	}

	return status, nil
}
