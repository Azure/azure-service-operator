// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"context"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:rbac:groups=kubernetesconfiguration.azure.com,resources=extensions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kubernetesconfiguration.azure.com,resources={extensions/status,extensions/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230501.Extension
// Generator information:
// - Generated from: /kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/extensions.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions/{extensionName}
type Extension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Extension_Spec   `json:"spec,omitempty"`
	Status            Extension_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Extension{}

// GetConditions returns the conditions of the resource
func (extension *Extension) GetConditions() conditions.Conditions {
	return extension.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (extension *Extension) SetConditions(conditions conditions.Conditions) {
	extension.Status.Conditions = conditions
}

var _ configmaps.Exporter = &Extension{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (extension *Extension) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if extension.Spec.OperatorSpec == nil {
		return nil
	}
	return extension.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Extension{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (extension *Extension) SecretDestinationExpressions() []*core.DestinationExpression {
	if extension.Spec.OperatorSpec == nil {
		return nil
	}
	return extension.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesConfigExporter = &Extension{}

// ExportKubernetesConfigMaps defines a resource which can create ConfigMaps in Kubernetes.
func (extension *Extension) ExportKubernetesConfigMaps(_ context.Context, _ genruntime.MetaObject, _ *genericarmclient.GenericClient, _ logr.Logger) ([]client.Object, error) {
	collector := configmaps.NewCollector(extension.Namespace)
	if extension.Spec.OperatorSpec != nil && extension.Spec.OperatorSpec.ConfigMaps != nil {
		if extension.Status.AksAssignedIdentity != nil {
			if extension.Status.AksAssignedIdentity.PrincipalId != nil {
				collector.AddValue(extension.Spec.OperatorSpec.ConfigMaps.PrincipalId, *extension.Status.AksAssignedIdentity.PrincipalId)
			}
		}
	}
	result, err := collector.Values()
	if err != nil {
		return nil, err
	}
	return configmaps.SliceToClientObjectSlice(result), nil
}

var _ genruntime.KubernetesResource = &Extension{}

// AzureName returns the Azure name of the resource
func (extension *Extension) AzureName() string {
	return extension.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (extension Extension) GetAPIVersion() string {
	return "2023-05-01"
}

// GetResourceScope returns the scope of the resource
func (extension *Extension) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeExtension
}

// GetSpec returns the specification of this resource
func (extension *Extension) GetSpec() genruntime.ConvertibleSpec {
	return &extension.Spec
}

// GetStatus returns the status of this resource
func (extension *Extension) GetStatus() genruntime.ConvertibleStatus {
	return &extension.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (extension *Extension) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.KubernetesConfiguration/extensions"
func (extension *Extension) GetType() string {
	return "Microsoft.KubernetesConfiguration/extensions"
}

// NewEmptyStatus returns a new empty (blank) status
func (extension *Extension) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Extension_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (extension *Extension) Owner() *genruntime.ResourceReference {
	return extension.Spec.Owner.AsResourceReference()
}

// SetStatus sets the status of this resource
func (extension *Extension) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Extension_STATUS); ok {
		extension.Status = *st
		return nil
	}

	// Convert status to required version
	var st Extension_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	extension.Status = st
	return nil
}

// Hub marks that this Extension is the hub type for conversion
func (extension *Extension) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (extension *Extension) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: extension.Spec.OriginalVersion,
		Kind:    "Extension",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501.Extension
// Generator information:
// - Generated from: /kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/extensions.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions/{extensionName}
type ExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Extension `json:"items"`
}

// Storage version of v1api20230501.APIVersion
// +kubebuilder:validation:Enum={"2023-05-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2023-05-01")

// Storage version of v1api20230501.Extension_Spec
type Extension_Spec struct {
	AksAssignedIdentity     *Extension_Properties_AksAssignedIdentity_Spec `json:"aksAssignedIdentity,omitempty"`
	AutoUpgradeMinorVersion *bool                                          `json:"autoUpgradeMinorVersion,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                      string                         `json:"azureName,omitempty"`
	ConfigurationProtectedSettings *genruntime.SecretMapReference `json:"configurationProtectedSettings,omitempty"`
	ConfigurationSettings          map[string]string              `json:"configurationSettings,omitempty"`
	ExtensionType                  *string                        `json:"extensionType,omitempty"`
	Identity                       *Identity                      `json:"identity,omitempty"`
	OperatorSpec                   *ExtensionOperatorSpec         `json:"operatorSpec,omitempty"`
	OriginalVersion                string                         `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
	// extension resource, which means that any other Azure resource can be its owner.
	Owner        *genruntime.ArbitraryOwnerReference `json:"owner,omitempty"`
	Plan         *Plan                               `json:"plan,omitempty"`
	PropertyBag  genruntime.PropertyBag              `json:"$propertyBag,omitempty"`
	ReleaseTrain *string                             `json:"releaseTrain,omitempty"`
	Scope        *Scope                              `json:"scope,omitempty"`
	SystemData   *SystemData                         `json:"systemData,omitempty"`
	Version      *string                             `json:"version,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Extension_Spec{}

// ConvertSpecFrom populates our Extension_Spec from the provided source
func (extension *Extension_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == extension {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(extension)
}

// ConvertSpecTo populates the provided destination from our Extension_Spec
func (extension *Extension_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == extension {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(extension)
}

// Storage version of v1api20230501.Extension_STATUS
// The Extension object.
type Extension_STATUS struct {
	AksAssignedIdentity            *Extension_Properties_AksAssignedIdentity_STATUS `json:"aksAssignedIdentity,omitempty"`
	AutoUpgradeMinorVersion        *bool                                            `json:"autoUpgradeMinorVersion,omitempty"`
	Conditions                     []conditions.Condition                           `json:"conditions,omitempty"`
	ConfigurationProtectedSettings map[string]string                                `json:"configurationProtectedSettings,omitempty"`
	ConfigurationSettings          map[string]string                                `json:"configurationSettings,omitempty"`
	CurrentVersion                 *string                                          `json:"currentVersion,omitempty"`
	CustomLocationSettings         map[string]string                                `json:"customLocationSettings,omitempty"`
	ErrorInfo                      *ErrorDetail_STATUS                              `json:"errorInfo,omitempty"`
	ExtensionType                  *string                                          `json:"extensionType,omitempty"`
	Id                             *string                                          `json:"id,omitempty"`
	Identity                       *Identity_STATUS                                 `json:"identity,omitempty"`
	IsSystemExtension              *bool                                            `json:"isSystemExtension,omitempty"`
	Name                           *string                                          `json:"name,omitempty"`
	PackageUri                     *string                                          `json:"packageUri,omitempty"`
	Plan                           *Plan_STATUS                                     `json:"plan,omitempty"`
	PropertyBag                    genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	ProvisioningState              *string                                          `json:"provisioningState,omitempty"`
	ReleaseTrain                   *string                                          `json:"releaseTrain,omitempty"`
	Scope                          *Scope_STATUS                                    `json:"scope,omitempty"`
	Statuses                       []ExtensionStatus_STATUS                         `json:"statuses,omitempty"`
	SystemData                     *SystemData_STATUS                               `json:"systemData,omitempty"`
	Type                           *string                                          `json:"type,omitempty"`
	Version                        *string                                          `json:"version,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Extension_STATUS{}

// ConvertStatusFrom populates our Extension_STATUS from the provided source
func (extension *Extension_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == extension {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(extension)
}

// ConvertStatusTo populates the provided destination from our Extension_STATUS
func (extension *Extension_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == extension {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(extension)
}

// Storage version of v1api20230501.ErrorDetail_STATUS
// The error detail.
type ErrorDetail_STATUS struct {
	AdditionalInfo []ErrorAdditionalInfo_STATUS  `json:"additionalInfo,omitempty"`
	Code           *string                       `json:"code,omitempty"`
	Details        []ErrorDetail_STATUS_Unrolled `json:"details,omitempty"`
	Message        *string                       `json:"message,omitempty"`
	PropertyBag    genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Target         *string                       `json:"target,omitempty"`
}

// Storage version of v1api20230501.Extension_Properties_AksAssignedIdentity_Spec
type Extension_Properties_AksAssignedIdentity_Spec struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20230501.Extension_Properties_AksAssignedIdentity_STATUS
type Extension_Properties_AksAssignedIdentity_STATUS struct {
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId    *string                `json:"tenantId,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20230501.ExtensionOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ExtensionOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	ConfigMaps           *ExtensionOperatorConfigMaps  `json:"configMaps,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20230501.ExtensionStatus_STATUS
// Status from the extension.
type ExtensionStatus_STATUS struct {
	Code          *string                `json:"code,omitempty"`
	DisplayStatus *string                `json:"displayStatus,omitempty"`
	Level         *string                `json:"level,omitempty"`
	Message       *string                `json:"message,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Time          *string                `json:"time,omitempty"`
}

// Storage version of v1api20230501.Identity
// Identity for the resource.
type Identity struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20230501.Identity_STATUS
// Identity for the resource.
type Identity_STATUS struct {
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId    *string                `json:"tenantId,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20230501.Plan
// Plan for the resource.
type Plan struct {
	Name          *string                `json:"name,omitempty"`
	Product       *string                `json:"product,omitempty"`
	PromotionCode *string                `json:"promotionCode,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Publisher     *string                `json:"publisher,omitempty"`
	Version       *string                `json:"version,omitempty"`
}

// Storage version of v1api20230501.Plan_STATUS
// Plan for the resource.
type Plan_STATUS struct {
	Name          *string                `json:"name,omitempty"`
	Product       *string                `json:"product,omitempty"`
	PromotionCode *string                `json:"promotionCode,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Publisher     *string                `json:"publisher,omitempty"`
	Version       *string                `json:"version,omitempty"`
}

// Storage version of v1api20230501.Scope
// Scope of the extension. It can be either Cluster or Namespace; but not both.
type Scope struct {
	Cluster     *ScopeCluster          `json:"cluster,omitempty"`
	Namespace   *ScopeNamespace        `json:"namespace,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.Scope_STATUS
// Scope of the extension. It can be either Cluster or Namespace; but not both.
type Scope_STATUS struct {
	Cluster     *ScopeCluster_STATUS   `json:"cluster,omitempty"`
	Namespace   *ScopeNamespace_STATUS `json:"namespace,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.SystemData
// Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.SystemData_STATUS
// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.ErrorAdditionalInfo_STATUS
// The resource management error additional info.
type ErrorAdditionalInfo_STATUS struct {
	Info        map[string]v1.JSON     `json:"info,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20230501.ErrorDetail_STATUS_Unrolled
type ErrorDetail_STATUS_Unrolled struct {
	AdditionalInfo []ErrorAdditionalInfo_STATUS `json:"additionalInfo,omitempty"`
	Code           *string                      `json:"code,omitempty"`
	Message        *string                      `json:"message,omitempty"`
	PropertyBag    genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Target         *string                      `json:"target,omitempty"`
}

// Storage version of v1api20230501.ExtensionOperatorConfigMaps
type ExtensionOperatorConfigMaps struct {
	PrincipalId *genruntime.ConfigMapDestination `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.ScopeCluster
// Specifies that the scope of the extension is Cluster
type ScopeCluster struct {
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReleaseNamespace *string                `json:"releaseNamespace,omitempty"`
}

// Storage version of v1api20230501.ScopeCluster_STATUS
// Specifies that the scope of the extension is Cluster
type ScopeCluster_STATUS struct {
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReleaseNamespace *string                `json:"releaseNamespace,omitempty"`
}

// Storage version of v1api20230501.ScopeNamespace
// Specifies that the scope of the extension is Namespace
type ScopeNamespace struct {
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TargetNamespace *string                `json:"targetNamespace,omitempty"`
}

// Storage version of v1api20230501.ScopeNamespace_STATUS
// Specifies that the scope of the extension is Namespace
type ScopeNamespace_STATUS struct {
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TargetNamespace *string                `json:"targetNamespace,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Extension{}, &ExtensionList{})
}
