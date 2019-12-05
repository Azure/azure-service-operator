package helpers

import (
	"context"

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

type AsyncClient interface {
	Ensure(context.Context, runtime.Object) (bool, error)
	Delete(context.Context, runtime.Object) (bool, error)
	Parents(runtime.Object) ([]KubeParent, error)
}
