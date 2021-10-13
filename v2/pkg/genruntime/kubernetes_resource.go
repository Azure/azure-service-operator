package genruntime

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// KubernetesResource is an Azure resource. This interface contains the common set of
// methods that apply to all ASO resources.
type KubernetesResource interface {
	conditions.Conditioner

	// Owner returns the ResourceReference of the owner, or nil if there is no owner
	Owner() *ResourceReference

	// TODO: I think we need this?
	// KnownOwner() *KnownResourceReference

	// AzureName returns the Azure name of the resource
	AzureName() string

	// GetType returns the type of the resource according to Azure. For example Microsoft.Resources/resourceGroups or
	// Microsoft.Network/networkSecurityGroups/securityRules
	GetType() string

	// GetResourceKind returns the ResourceKind of the resource.
	GetResourceKind() ResourceKind

	// Some types, but not all, have a corresponding:
	// 	SetAzureName(name string)
	// They do not if the name must be a fixed value (like 'default').

	// TODO: GetAPIVersion here?

	// GetSpec returns the specification of the resource
	GetSpec() ConvertibleSpec

	// GetStatus returns the current status of the resource
	GetStatus() ConvertibleStatus

	// NewEmptyStatus returns a blank status ready for population
	NewEmptyStatus() ConvertibleStatus

	// SetStatus updates the status of the resource
	SetStatus(status ConvertibleStatus) error
}

// NewEmptyVersionedResource returns a new blank resource based on the passed metaObject; the original API version used
// (if available) from when the resource was first created is used to identify the version to return.
// returns the resource, if converted; nil if already the correct version
func NewEmptyVersionedResource(metaObject MetaObject, scheme *runtime.Scheme) (KubernetesResource, error) {
	aware, ok := metaObject.(GroupVersionKindAware)
	if !ok {
		// Resource is not aware of the original GVK used at creation, so no conversion possible
		return nil, nil
	}

	currentGVK := metaObject.GetObjectKind().GroupVersionKind()
	originalGVK := aware.OriginalGVK()

	if currentGVK == *originalGVK {
		// Already have the desired version of the resource, so no conversion is needed
		return nil, nil
	}

	// Create an empty resource at the desired version
	rsrc, err := scheme.New(*originalGVK)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create new %s", originalGVK)
	}

	// Convert it to our interface
	kr, ok := rsrc.(KubernetesResource)
	if !ok {
		return nil, errors.Errorf("expected resource %s to implement genruntime.KubernetesResource", originalGVK)
	}

	// Return the empty resource
	return kr, nil
}
