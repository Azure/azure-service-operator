package helpers

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Data wrapps the object that is needed for the services
type Data struct {
	Obj interface{}
}

type KubeParent struct {
	Key    types.NamespacedName
	Target runtime.Object
}
