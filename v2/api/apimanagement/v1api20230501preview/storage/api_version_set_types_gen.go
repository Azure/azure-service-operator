// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230501preview.ApiVersionSet
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimapiversionsets.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}
type ApiVersionSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiVersionSet_Spec   `json:"spec,omitempty"`
	Status            ApiVersionSet_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ApiVersionSet{}

// GetConditions returns the conditions of the resource
func (versionSet *ApiVersionSet) GetConditions() conditions.Conditions {
	return versionSet.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (versionSet *ApiVersionSet) SetConditions(conditions conditions.Conditions) {
	versionSet.Status.Conditions = conditions
}

var _ conversion.Convertible = &ApiVersionSet{}

// ConvertFrom populates our ApiVersionSet from the provided hub ApiVersionSet
func (versionSet *ApiVersionSet) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.ApiVersionSet)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/ApiVersionSet but received %T instead", hub)
	}

	return versionSet.AssignProperties_From_ApiVersionSet(source)
}

// ConvertTo populates the provided hub ApiVersionSet from our ApiVersionSet
func (versionSet *ApiVersionSet) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.ApiVersionSet)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/ApiVersionSet but received %T instead", hub)
	}

	return versionSet.AssignProperties_To_ApiVersionSet(destination)
}

var _ configmaps.Exporter = &ApiVersionSet{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (versionSet *ApiVersionSet) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if versionSet.Spec.OperatorSpec == nil {
		return nil
	}
	return versionSet.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &ApiVersionSet{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (versionSet *ApiVersionSet) SecretDestinationExpressions() []*core.DestinationExpression {
	if versionSet.Spec.OperatorSpec == nil {
		return nil
	}
	return versionSet.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &ApiVersionSet{}

// AzureName returns the Azure name of the resource
func (versionSet *ApiVersionSet) AzureName() string {
	return versionSet.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (versionSet ApiVersionSet) GetAPIVersion() string {
	return "2023-05-01-preview"
}

// GetResourceScope returns the scope of the resource
func (versionSet *ApiVersionSet) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (versionSet *ApiVersionSet) GetSpec() genruntime.ConvertibleSpec {
	return &versionSet.Spec
}

// GetStatus returns the status of this resource
func (versionSet *ApiVersionSet) GetStatus() genruntime.ConvertibleStatus {
	return &versionSet.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (versionSet *ApiVersionSet) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/apiVersionSets"
func (versionSet *ApiVersionSet) GetType() string {
	return "Microsoft.ApiManagement/service/apiVersionSets"
}

// NewEmptyStatus returns a new empty (blank) status
func (versionSet *ApiVersionSet) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ApiVersionSet_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (versionSet *ApiVersionSet) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(versionSet.Spec)
	return versionSet.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (versionSet *ApiVersionSet) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ApiVersionSet_STATUS); ok {
		versionSet.Status = *st
		return nil
	}

	// Convert status to required version
	var st ApiVersionSet_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	versionSet.Status = st
	return nil
}

// AssignProperties_From_ApiVersionSet populates our ApiVersionSet from the provided source ApiVersionSet
func (versionSet *ApiVersionSet) AssignProperties_From_ApiVersionSet(source *storage.ApiVersionSet) error {

	// ObjectMeta
	versionSet.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec ApiVersionSet_Spec
	err := spec.AssignProperties_From_ApiVersionSet_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_ApiVersionSet_Spec() to populate field Spec")
	}
	versionSet.Spec = spec

	// Status
	var status ApiVersionSet_STATUS
	err = status.AssignProperties_From_ApiVersionSet_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_ApiVersionSet_STATUS() to populate field Status")
	}
	versionSet.Status = status

	// Invoke the augmentConversionForApiVersionSet interface (if implemented) to customize the conversion
	var versionSetAsAny any = versionSet
	if augmentedVersionSet, ok := versionSetAsAny.(augmentConversionForApiVersionSet); ok {
		err := augmentedVersionSet.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_ApiVersionSet populates the provided destination ApiVersionSet from our ApiVersionSet
func (versionSet *ApiVersionSet) AssignProperties_To_ApiVersionSet(destination *storage.ApiVersionSet) error {

	// ObjectMeta
	destination.ObjectMeta = *versionSet.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.ApiVersionSet_Spec
	err := versionSet.Spec.AssignProperties_To_ApiVersionSet_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_ApiVersionSet_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.ApiVersionSet_STATUS
	err = versionSet.Status.AssignProperties_To_ApiVersionSet_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_ApiVersionSet_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForApiVersionSet interface (if implemented) to customize the conversion
	var versionSetAsAny any = versionSet
	if augmentedVersionSet, ok := versionSetAsAny.(augmentConversionForApiVersionSet); ok {
		err := augmentedVersionSet.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (versionSet *ApiVersionSet) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: versionSet.Spec.OriginalVersion,
		Kind:    "ApiVersionSet",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501preview.ApiVersionSet
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimapiversionsets.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}
type ApiVersionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiVersionSet `json:"items"`
}

// Storage version of v1api20230501preview.ApiVersionSet_Spec
type ApiVersionSet_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                     `json:"azureName,omitempty"`
	Description     *string                    `json:"description,omitempty"`
	DisplayName     *string                    `json:"displayName,omitempty"`
	OperatorSpec    *ApiVersionSetOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                     `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner             *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`
	PropertyBag       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	VersionHeaderName *string                            `json:"versionHeaderName,omitempty"`
	VersionQueryName  *string                            `json:"versionQueryName,omitempty"`
	VersioningScheme  *string                            `json:"versioningScheme,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ApiVersionSet_Spec{}

// ConvertSpecFrom populates our ApiVersionSet_Spec from the provided source
func (versionSet *ApiVersionSet_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.ApiVersionSet_Spec)
	if ok {
		// Populate our instance from source
		return versionSet.AssignProperties_From_ApiVersionSet_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.ApiVersionSet_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = versionSet.AssignProperties_From_ApiVersionSet_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our ApiVersionSet_Spec
func (versionSet *ApiVersionSet_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.ApiVersionSet_Spec)
	if ok {
		// Populate destination from our instance
		return versionSet.AssignProperties_To_ApiVersionSet_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.ApiVersionSet_Spec{}
	err := versionSet.AssignProperties_To_ApiVersionSet_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_ApiVersionSet_Spec populates our ApiVersionSet_Spec from the provided source ApiVersionSet_Spec
func (versionSet *ApiVersionSet_Spec) AssignProperties_From_ApiVersionSet_Spec(source *storage.ApiVersionSet_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	versionSet.AzureName = source.AzureName

	// Description
	versionSet.Description = genruntime.ClonePointerToString(source.Description)

	// DisplayName
	versionSet.DisplayName = genruntime.ClonePointerToString(source.DisplayName)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec ApiVersionSetOperatorSpec
		err := operatorSpec.AssignProperties_From_ApiVersionSetOperatorSpec(source.OperatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_ApiVersionSetOperatorSpec() to populate field OperatorSpec")
		}
		versionSet.OperatorSpec = &operatorSpec
	} else {
		versionSet.OperatorSpec = nil
	}

	// OriginalVersion
	versionSet.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		versionSet.Owner = &owner
	} else {
		versionSet.Owner = nil
	}

	// VersionHeaderName
	versionSet.VersionHeaderName = genruntime.ClonePointerToString(source.VersionHeaderName)

	// VersionQueryName
	versionSet.VersionQueryName = genruntime.ClonePointerToString(source.VersionQueryName)

	// VersioningScheme
	versionSet.VersioningScheme = genruntime.ClonePointerToString(source.VersioningScheme)

	// Update the property bag
	if len(propertyBag) > 0 {
		versionSet.PropertyBag = propertyBag
	} else {
		versionSet.PropertyBag = nil
	}

	// Invoke the augmentConversionForApiVersionSet_Spec interface (if implemented) to customize the conversion
	var versionSetAsAny any = versionSet
	if augmentedVersionSet, ok := versionSetAsAny.(augmentConversionForApiVersionSet_Spec); ok {
		err := augmentedVersionSet.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_ApiVersionSet_Spec populates the provided destination ApiVersionSet_Spec from our ApiVersionSet_Spec
func (versionSet *ApiVersionSet_Spec) AssignProperties_To_ApiVersionSet_Spec(destination *storage.ApiVersionSet_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(versionSet.PropertyBag)

	// AzureName
	destination.AzureName = versionSet.AzureName

	// Description
	destination.Description = genruntime.ClonePointerToString(versionSet.Description)

	// DisplayName
	destination.DisplayName = genruntime.ClonePointerToString(versionSet.DisplayName)

	// OperatorSpec
	if versionSet.OperatorSpec != nil {
		var operatorSpec storage.ApiVersionSetOperatorSpec
		err := versionSet.OperatorSpec.AssignProperties_To_ApiVersionSetOperatorSpec(&operatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_ApiVersionSetOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = versionSet.OriginalVersion

	// Owner
	if versionSet.Owner != nil {
		owner := versionSet.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// VersionHeaderName
	destination.VersionHeaderName = genruntime.ClonePointerToString(versionSet.VersionHeaderName)

	// VersionQueryName
	destination.VersionQueryName = genruntime.ClonePointerToString(versionSet.VersionQueryName)

	// VersioningScheme
	destination.VersioningScheme = genruntime.ClonePointerToString(versionSet.VersioningScheme)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForApiVersionSet_Spec interface (if implemented) to customize the conversion
	var versionSetAsAny any = versionSet
	if augmentedVersionSet, ok := versionSetAsAny.(augmentConversionForApiVersionSet_Spec); ok {
		err := augmentedVersionSet.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20230501preview.ApiVersionSet_STATUS
type ApiVersionSet_STATUS struct {
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Description       *string                `json:"description,omitempty"`
	DisplayName       *string                `json:"displayName,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type              *string                `json:"type,omitempty"`
	VersionHeaderName *string                `json:"versionHeaderName,omitempty"`
	VersionQueryName  *string                `json:"versionQueryName,omitempty"`
	VersioningScheme  *string                `json:"versioningScheme,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ApiVersionSet_STATUS{}

// ConvertStatusFrom populates our ApiVersionSet_STATUS from the provided source
func (versionSet *ApiVersionSet_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.ApiVersionSet_STATUS)
	if ok {
		// Populate our instance from source
		return versionSet.AssignProperties_From_ApiVersionSet_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.ApiVersionSet_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = versionSet.AssignProperties_From_ApiVersionSet_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our ApiVersionSet_STATUS
func (versionSet *ApiVersionSet_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.ApiVersionSet_STATUS)
	if ok {
		// Populate destination from our instance
		return versionSet.AssignProperties_To_ApiVersionSet_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.ApiVersionSet_STATUS{}
	err := versionSet.AssignProperties_To_ApiVersionSet_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_ApiVersionSet_STATUS populates our ApiVersionSet_STATUS from the provided source ApiVersionSet_STATUS
func (versionSet *ApiVersionSet_STATUS) AssignProperties_From_ApiVersionSet_STATUS(source *storage.ApiVersionSet_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	versionSet.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Description
	versionSet.Description = genruntime.ClonePointerToString(source.Description)

	// DisplayName
	versionSet.DisplayName = genruntime.ClonePointerToString(source.DisplayName)

	// Id
	versionSet.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	versionSet.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	versionSet.Type = genruntime.ClonePointerToString(source.Type)

	// VersionHeaderName
	versionSet.VersionHeaderName = genruntime.ClonePointerToString(source.VersionHeaderName)

	// VersionQueryName
	versionSet.VersionQueryName = genruntime.ClonePointerToString(source.VersionQueryName)

	// VersioningScheme
	versionSet.VersioningScheme = genruntime.ClonePointerToString(source.VersioningScheme)

	// Update the property bag
	if len(propertyBag) > 0 {
		versionSet.PropertyBag = propertyBag
	} else {
		versionSet.PropertyBag = nil
	}

	// Invoke the augmentConversionForApiVersionSet_STATUS interface (if implemented) to customize the conversion
	var versionSetAsAny any = versionSet
	if augmentedVersionSet, ok := versionSetAsAny.(augmentConversionForApiVersionSet_STATUS); ok {
		err := augmentedVersionSet.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_ApiVersionSet_STATUS populates the provided destination ApiVersionSet_STATUS from our ApiVersionSet_STATUS
func (versionSet *ApiVersionSet_STATUS) AssignProperties_To_ApiVersionSet_STATUS(destination *storage.ApiVersionSet_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(versionSet.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(versionSet.Conditions)

	// Description
	destination.Description = genruntime.ClonePointerToString(versionSet.Description)

	// DisplayName
	destination.DisplayName = genruntime.ClonePointerToString(versionSet.DisplayName)

	// Id
	destination.Id = genruntime.ClonePointerToString(versionSet.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(versionSet.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(versionSet.Type)

	// VersionHeaderName
	destination.VersionHeaderName = genruntime.ClonePointerToString(versionSet.VersionHeaderName)

	// VersionQueryName
	destination.VersionQueryName = genruntime.ClonePointerToString(versionSet.VersionQueryName)

	// VersioningScheme
	destination.VersioningScheme = genruntime.ClonePointerToString(versionSet.VersioningScheme)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForApiVersionSet_STATUS interface (if implemented) to customize the conversion
	var versionSetAsAny any = versionSet
	if augmentedVersionSet, ok := versionSetAsAny.(augmentConversionForApiVersionSet_STATUS); ok {
		err := augmentedVersionSet.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForApiVersionSet interface {
	AssignPropertiesFrom(src *storage.ApiVersionSet) error
	AssignPropertiesTo(dst *storage.ApiVersionSet) error
}

// Storage version of v1api20230501preview.ApiVersionSetOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ApiVersionSetOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_ApiVersionSetOperatorSpec populates our ApiVersionSetOperatorSpec from the provided source ApiVersionSetOperatorSpec
func (operator *ApiVersionSetOperatorSpec) AssignProperties_From_ApiVersionSetOperatorSpec(source *storage.ApiVersionSetOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		operator.PropertyBag = propertyBag
	} else {
		operator.PropertyBag = nil
	}

	// Invoke the augmentConversionForApiVersionSetOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForApiVersionSetOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_ApiVersionSetOperatorSpec populates the provided destination ApiVersionSetOperatorSpec from our ApiVersionSetOperatorSpec
func (operator *ApiVersionSetOperatorSpec) AssignProperties_To_ApiVersionSetOperatorSpec(destination *storage.ApiVersionSetOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(operator.PropertyBag)

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForApiVersionSetOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForApiVersionSetOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForApiVersionSet_Spec interface {
	AssignPropertiesFrom(src *storage.ApiVersionSet_Spec) error
	AssignPropertiesTo(dst *storage.ApiVersionSet_Spec) error
}

type augmentConversionForApiVersionSet_STATUS interface {
	AssignPropertiesFrom(src *storage.ApiVersionSet_STATUS) error
	AssignPropertiesTo(dst *storage.ApiVersionSet_STATUS) error
}

type augmentConversionForApiVersionSetOperatorSpec interface {
	AssignPropertiesFrom(src *storage.ApiVersionSetOperatorSpec) error
	AssignPropertiesTo(dst *storage.ApiVersionSetOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&ApiVersionSet{}, &ApiVersionSetList{})
}
