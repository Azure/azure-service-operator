package apis

import (
	corev1 "k8s.io/api/core/v1"
)

type (
	Grouped interface {
		GetResourceGroupObjectRef() *corev1.ObjectReference
	}
)
