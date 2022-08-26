// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources=workspacesconnections,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources={workspacesconnections/status,workspacesconnections/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210701.WorkspacesConnection
// Generator information:
// - Generated from: /machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/machineLearningServices.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}
type WorkspacesConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
<<<<<<< HEAD
	Spec              WorkspacesConnection_Spec   `json:"spec,omitempty"`
	Status            WorkspacesConnection_STATUS `json:"status,omitempty"`
=======
	Spec              Workspaces_Connections_Spec `json:"spec,omitempty"`
	Status            WorkspaceConnection_STATUS  `json:"status,omitempty"`
>>>>>>> main
}

var _ conditions.Conditioner = &WorkspacesConnection{}

// GetConditions returns the conditions of the resource
func (connection *WorkspacesConnection) GetConditions() conditions.Conditions {
	return connection.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (connection *WorkspacesConnection) SetConditions(conditions conditions.Conditions) {
	connection.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &WorkspacesConnection{}

// AzureName returns the Azure name of the resource
func (connection *WorkspacesConnection) AzureName() string {
	return connection.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (connection WorkspacesConnection) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (connection *WorkspacesConnection) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (connection *WorkspacesConnection) GetSpec() genruntime.ConvertibleSpec {
	return &connection.Spec
}

// GetStatus returns the status of this resource
func (connection *WorkspacesConnection) GetStatus() genruntime.ConvertibleStatus {
	return &connection.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces/connections"
func (connection *WorkspacesConnection) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces/connections"
}

// NewEmptyStatus returns a new empty (blank) status
func (connection *WorkspacesConnection) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &WorkspacesConnection_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (connection *WorkspacesConnection) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(connection.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  connection.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (connection *WorkspacesConnection) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*WorkspacesConnection_STATUS); ok {
		connection.Status = *st
		return nil
	}

	// Convert status to required version
	var st WorkspacesConnection_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	connection.Status = st
	return nil
}

// Hub marks that this WorkspacesConnection is the hub type for conversion
func (connection *WorkspacesConnection) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (connection *WorkspacesConnection) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: connection.Spec.OriginalVersion,
		Kind:    "WorkspacesConnection",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210701.WorkspacesConnection
// Generator information:
// - Generated from: /machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/machineLearningServices.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}
type WorkspacesConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspacesConnection `json:"items"`
}

// Storage version of v1beta20210701.WorkspacesConnection_Spec
type WorkspacesConnection_Spec struct {
	AuthType *string `json:"authType,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Category        *string `json:"category,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a machinelearningservices.azure.com/Workspace resource
	Owner       *genruntime.KnownResourceReference `group:"machinelearningservices.azure.com" json:"owner,omitempty" kind:"Workspace"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Target      *string                            `json:"target,omitempty"`
	Value       *string                            `json:"value,omitempty"`
	ValueFormat *string                            `json:"valueFormat,omitempty"`
}

var _ genruntime.ConvertibleSpec = &WorkspacesConnection_Spec{}

// ConvertSpecFrom populates our WorkspacesConnection_Spec from the provided source
func (connection *WorkspacesConnection_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == connection {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(connection)
}

// ConvertSpecTo populates the provided destination from our WorkspacesConnection_Spec
func (connection *WorkspacesConnection_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == connection {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(connection)
}

// Storage version of v1beta20210701.WorkspacesConnection_STATUS
type WorkspacesConnection_STATUS struct {
	AuthType    *string                `json:"authType,omitempty"`
	Category    *string                `json:"category,omitempty"`
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Target      *string                `json:"target,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Value       *string                `json:"value,omitempty"`
	ValueFormat *string                `json:"valueFormat,omitempty"`
}

var _ genruntime.ConvertibleStatus = &WorkspacesConnection_STATUS{}

// ConvertStatusFrom populates our WorkspacesConnection_STATUS from the provided source
func (connection *WorkspacesConnection_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == connection {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(connection)
}

// ConvertStatusTo populates the provided destination from our WorkspacesConnection_STATUS
func (connection *WorkspacesConnection_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == connection {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(connection)
}

<<<<<<< HEAD
=======
// Storage version of v1beta20210701.Workspaces_Connections_Spec
type Workspaces_Connections_Spec struct {
	AuthType *string `json:"authType,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Category        *string `json:"category,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a machinelearningservices.azure.com/Workspace resource
	Owner       *genruntime.KnownResourceReference `group:"machinelearningservices.azure.com" json:"owner,omitempty" kind:"Workspace"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
	Target      *string                            `json:"target,omitempty"`
	Value       *string                            `json:"value,omitempty"`
	ValueFormat *string                            `json:"valueFormat,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Workspaces_Connections_Spec{}

// ConvertSpecFrom populates our Workspaces_Connections_Spec from the provided source
func (connections *Workspaces_Connections_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == connections {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(connections)
}

// ConvertSpecTo populates the provided destination from our Workspaces_Connections_Spec
func (connections *Workspaces_Connections_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == connections {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(connections)
}

>>>>>>> main
func init() {
	SchemeBuilder.Register(&WorkspacesConnection{}, &WorkspacesConnectionList{})
}
