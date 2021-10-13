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

