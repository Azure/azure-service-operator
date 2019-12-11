package resourcemanager

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type KubeParent struct {
	Key    types.NamespacedName
	Target runtime.Object
}

type AsyncClient interface {
	Ensure(context.Context, runtime.Object) (bool, error)
	Delete(context.Context, runtime.Object) (bool, error)
	Parents(runtime.Object) ([]KubeParent, error)
}
