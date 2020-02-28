package resourcemanager

import (
	"context"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

const (
	SuccessMsg string = "successfully provisioned"
)

type KubeParent struct {
	Key    types.NamespacedName
	Target runtime.Object
}

type ARMClient interface {
	Ensure(context.Context, runtime.Object) (bool, error)
	Delete(context.Context, runtime.Object) (bool, error)
	GetParents(runtime.Object) ([]KubeParent, error)
	GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error)
}
