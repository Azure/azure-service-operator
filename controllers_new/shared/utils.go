package shared

import (
	"strings"

	"github.com/Azure/azure-service-operator/pkg/reconciler"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

func IsNotFalse(s string) bool {
	return !IsFalse(s)
}

func IsFalse(s string) bool {
	return strings.ToLower(s) == "false" || strings.HasPrefix(strings.ToLower(s), "n") || strings.ToLower(s) == "0"
}

func GetOwnerIfManaged(meta controllerruntime.ObjectMeta, ownerDependency func() *reconciler.Dependency) *reconciler.Dependency {
	// get the metadata annotation to check if the eventhubnamespace belongs to a resourcegroup that is
	// managed by Kubernetes, and if so, make this resourcegroup a dependency
	managedResourceGroup := meta.Annotations[ManagedParentAnnotation]
	// defaults to true
	isManaged := IsNotFalse(managedResourceGroup)
	var owner *reconciler.Dependency = nil
	if isManaged {
		owner = ownerDependency()
	}
	return owner
}
