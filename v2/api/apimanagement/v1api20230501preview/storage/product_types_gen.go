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
	"github.com/rotisserie/eris"
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
// Storage version of v1api20230501preview.Product
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimproducts.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}
type Product struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Product_Spec   `json:"spec,omitempty"`
	Status            Product_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Product{}

// GetConditions returns the conditions of the resource
func (product *Product) GetConditions() conditions.Conditions {
	return product.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (product *Product) SetConditions(conditions conditions.Conditions) {
	product.Status.Conditions = conditions
}

var _ conversion.Convertible = &Product{}

// ConvertFrom populates our Product from the provided hub Product
func (product *Product) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.Product)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/Product but received %T instead", hub)
	}

	return product.AssignProperties_From_Product(source)
}

// ConvertTo populates the provided hub Product from our Product
func (product *Product) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.Product)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/Product but received %T instead", hub)
	}

	return product.AssignProperties_To_Product(destination)
}

var _ configmaps.Exporter = &Product{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (product *Product) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if product.Spec.OperatorSpec == nil {
		return nil
	}
	return product.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Product{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (product *Product) SecretDestinationExpressions() []*core.DestinationExpression {
	if product.Spec.OperatorSpec == nil {
		return nil
	}
	return product.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &Product{}

// AzureName returns the Azure name of the resource
func (product *Product) AzureName() string {
	return product.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (product Product) GetAPIVersion() string {
	return "2023-05-01-preview"
}

// GetResourceScope returns the scope of the resource
func (product *Product) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (product *Product) GetSpec() genruntime.ConvertibleSpec {
	return &product.Spec
}

// GetStatus returns the status of this resource
func (product *Product) GetStatus() genruntime.ConvertibleStatus {
	return &product.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (product *Product) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/products"
func (product *Product) GetType() string {
	return "Microsoft.ApiManagement/service/products"
}

// NewEmptyStatus returns a new empty (blank) status
func (product *Product) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Product_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (product *Product) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(product.Spec)
	return product.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (product *Product) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Product_STATUS); ok {
		product.Status = *st
		return nil
	}

	// Convert status to required version
	var st Product_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	product.Status = st
	return nil
}

// AssignProperties_From_Product populates our Product from the provided source Product
func (product *Product) AssignProperties_From_Product(source *storage.Product) error {

	// ObjectMeta
	product.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Product_Spec
	err := spec.AssignProperties_From_Product_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_Product_Spec() to populate field Spec")
	}
	product.Spec = spec

	// Status
	var status Product_STATUS
	err = status.AssignProperties_From_Product_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_Product_STATUS() to populate field Status")
	}
	product.Status = status

	// Invoke the augmentConversionForProduct interface (if implemented) to customize the conversion
	var productAsAny any = product
	if augmentedProduct, ok := productAsAny.(augmentConversionForProduct); ok {
		err := augmentedProduct.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Product populates the provided destination Product from our Product
func (product *Product) AssignProperties_To_Product(destination *storage.Product) error {

	// ObjectMeta
	destination.ObjectMeta = *product.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Product_Spec
	err := product.Spec.AssignProperties_To_Product_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_Product_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Product_STATUS
	err = product.Status.AssignProperties_To_Product_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_Product_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForProduct interface (if implemented) to customize the conversion
	var productAsAny any = product
	if augmentedProduct, ok := productAsAny.(augmentConversionForProduct); ok {
		err := augmentedProduct.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (product *Product) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: product.Spec.OriginalVersion,
		Kind:    "Product",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501preview.Product
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimproducts.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}
type ProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Product `json:"items"`
}

type augmentConversionForProduct interface {
	AssignPropertiesFrom(src *storage.Product) error
	AssignPropertiesTo(dst *storage.Product) error
}

// Storage version of v1api20230501preview.Product_Spec
type Product_Spec struct {
	ApprovalRequired *bool `json:"approvalRequired,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Description     *string              `json:"description,omitempty"`
	DisplayName     *string              `json:"displayName,omitempty"`
	OperatorSpec    *ProductOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner                *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`
	PropertyBag          genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	State                *string                            `json:"state,omitempty"`
	SubscriptionRequired *bool                              `json:"subscriptionRequired,omitempty"`
	SubscriptionsLimit   *int                               `json:"subscriptionsLimit,omitempty"`
	Terms                *string                            `json:"terms,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Product_Spec{}

// ConvertSpecFrom populates our Product_Spec from the provided source
func (product *Product_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Product_Spec)
	if ok {
		// Populate our instance from source
		return product.AssignProperties_From_Product_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Product_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = product.AssignProperties_From_Product_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Product_Spec
func (product *Product_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Product_Spec)
	if ok {
		// Populate destination from our instance
		return product.AssignProperties_To_Product_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Product_Spec{}
	err := product.AssignProperties_To_Product_Spec(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_Product_Spec populates our Product_Spec from the provided source Product_Spec
func (product *Product_Spec) AssignProperties_From_Product_Spec(source *storage.Product_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ApprovalRequired
	if source.ApprovalRequired != nil {
		approvalRequired := *source.ApprovalRequired
		product.ApprovalRequired = &approvalRequired
	} else {
		product.ApprovalRequired = nil
	}

	// AzureName
	product.AzureName = source.AzureName

	// Description
	product.Description = genruntime.ClonePointerToString(source.Description)

	// DisplayName
	product.DisplayName = genruntime.ClonePointerToString(source.DisplayName)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec ProductOperatorSpec
		err := operatorSpec.AssignProperties_From_ProductOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_ProductOperatorSpec() to populate field OperatorSpec")
		}
		product.OperatorSpec = &operatorSpec
	} else {
		product.OperatorSpec = nil
	}

	// OriginalVersion
	product.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		product.Owner = &owner
	} else {
		product.Owner = nil
	}

	// State
	product.State = genruntime.ClonePointerToString(source.State)

	// SubscriptionRequired
	if source.SubscriptionRequired != nil {
		subscriptionRequired := *source.SubscriptionRequired
		product.SubscriptionRequired = &subscriptionRequired
	} else {
		product.SubscriptionRequired = nil
	}

	// SubscriptionsLimit
	product.SubscriptionsLimit = genruntime.ClonePointerToInt(source.SubscriptionsLimit)

	// Terms
	product.Terms = genruntime.ClonePointerToString(source.Terms)

	// Update the property bag
	if len(propertyBag) > 0 {
		product.PropertyBag = propertyBag
	} else {
		product.PropertyBag = nil
	}

	// Invoke the augmentConversionForProduct_Spec interface (if implemented) to customize the conversion
	var productAsAny any = product
	if augmentedProduct, ok := productAsAny.(augmentConversionForProduct_Spec); ok {
		err := augmentedProduct.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Product_Spec populates the provided destination Product_Spec from our Product_Spec
func (product *Product_Spec) AssignProperties_To_Product_Spec(destination *storage.Product_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(product.PropertyBag)

	// ApprovalRequired
	if product.ApprovalRequired != nil {
		approvalRequired := *product.ApprovalRequired
		destination.ApprovalRequired = &approvalRequired
	} else {
		destination.ApprovalRequired = nil
	}

	// AzureName
	destination.AzureName = product.AzureName

	// Description
	destination.Description = genruntime.ClonePointerToString(product.Description)

	// DisplayName
	destination.DisplayName = genruntime.ClonePointerToString(product.DisplayName)

	// OperatorSpec
	if product.OperatorSpec != nil {
		var operatorSpec storage.ProductOperatorSpec
		err := product.OperatorSpec.AssignProperties_To_ProductOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_ProductOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = product.OriginalVersion

	// Owner
	if product.Owner != nil {
		owner := product.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// State
	destination.State = genruntime.ClonePointerToString(product.State)

	// SubscriptionRequired
	if product.SubscriptionRequired != nil {
		subscriptionRequired := *product.SubscriptionRequired
		destination.SubscriptionRequired = &subscriptionRequired
	} else {
		destination.SubscriptionRequired = nil
	}

	// SubscriptionsLimit
	destination.SubscriptionsLimit = genruntime.ClonePointerToInt(product.SubscriptionsLimit)

	// Terms
	destination.Terms = genruntime.ClonePointerToString(product.Terms)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForProduct_Spec interface (if implemented) to customize the conversion
	var productAsAny any = product
	if augmentedProduct, ok := productAsAny.(augmentConversionForProduct_Spec); ok {
		err := augmentedProduct.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20230501preview.Product_STATUS
type Product_STATUS struct {
	ApprovalRequired     *bool                  `json:"approvalRequired,omitempty"`
	Conditions           []conditions.Condition `json:"conditions,omitempty"`
	Description          *string                `json:"description,omitempty"`
	DisplayName          *string                `json:"displayName,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State                *string                `json:"state,omitempty"`
	SubscriptionRequired *bool                  `json:"subscriptionRequired,omitempty"`
	SubscriptionsLimit   *int                   `json:"subscriptionsLimit,omitempty"`
	Terms                *string                `json:"terms,omitempty"`
	Type                 *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Product_STATUS{}

// ConvertStatusFrom populates our Product_STATUS from the provided source
func (product *Product_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Product_STATUS)
	if ok {
		// Populate our instance from source
		return product.AssignProperties_From_Product_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Product_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = product.AssignProperties_From_Product_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Product_STATUS
func (product *Product_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Product_STATUS)
	if ok {
		// Populate destination from our instance
		return product.AssignProperties_To_Product_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Product_STATUS{}
	err := product.AssignProperties_To_Product_STATUS(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_Product_STATUS populates our Product_STATUS from the provided source Product_STATUS
func (product *Product_STATUS) AssignProperties_From_Product_STATUS(source *storage.Product_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ApprovalRequired
	if source.ApprovalRequired != nil {
		approvalRequired := *source.ApprovalRequired
		product.ApprovalRequired = &approvalRequired
	} else {
		product.ApprovalRequired = nil
	}

	// Conditions
	product.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Description
	product.Description = genruntime.ClonePointerToString(source.Description)

	// DisplayName
	product.DisplayName = genruntime.ClonePointerToString(source.DisplayName)

	// Id
	product.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	product.Name = genruntime.ClonePointerToString(source.Name)

	// State
	product.State = genruntime.ClonePointerToString(source.State)

	// SubscriptionRequired
	if source.SubscriptionRequired != nil {
		subscriptionRequired := *source.SubscriptionRequired
		product.SubscriptionRequired = &subscriptionRequired
	} else {
		product.SubscriptionRequired = nil
	}

	// SubscriptionsLimit
	product.SubscriptionsLimit = genruntime.ClonePointerToInt(source.SubscriptionsLimit)

	// Terms
	product.Terms = genruntime.ClonePointerToString(source.Terms)

	// Type
	product.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		product.PropertyBag = propertyBag
	} else {
		product.PropertyBag = nil
	}

	// Invoke the augmentConversionForProduct_STATUS interface (if implemented) to customize the conversion
	var productAsAny any = product
	if augmentedProduct, ok := productAsAny.(augmentConversionForProduct_STATUS); ok {
		err := augmentedProduct.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Product_STATUS populates the provided destination Product_STATUS from our Product_STATUS
func (product *Product_STATUS) AssignProperties_To_Product_STATUS(destination *storage.Product_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(product.PropertyBag)

	// ApprovalRequired
	if product.ApprovalRequired != nil {
		approvalRequired := *product.ApprovalRequired
		destination.ApprovalRequired = &approvalRequired
	} else {
		destination.ApprovalRequired = nil
	}

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(product.Conditions)

	// Description
	destination.Description = genruntime.ClonePointerToString(product.Description)

	// DisplayName
	destination.DisplayName = genruntime.ClonePointerToString(product.DisplayName)

	// Id
	destination.Id = genruntime.ClonePointerToString(product.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(product.Name)

	// State
	destination.State = genruntime.ClonePointerToString(product.State)

	// SubscriptionRequired
	if product.SubscriptionRequired != nil {
		subscriptionRequired := *product.SubscriptionRequired
		destination.SubscriptionRequired = &subscriptionRequired
	} else {
		destination.SubscriptionRequired = nil
	}

	// SubscriptionsLimit
	destination.SubscriptionsLimit = genruntime.ClonePointerToInt(product.SubscriptionsLimit)

	// Terms
	destination.Terms = genruntime.ClonePointerToString(product.Terms)

	// Type
	destination.Type = genruntime.ClonePointerToString(product.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForProduct_STATUS interface (if implemented) to customize the conversion
	var productAsAny any = product
	if augmentedProduct, ok := productAsAny.(augmentConversionForProduct_STATUS); ok {
		err := augmentedProduct.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForProduct_Spec interface {
	AssignPropertiesFrom(src *storage.Product_Spec) error
	AssignPropertiesTo(dst *storage.Product_Spec) error
}

type augmentConversionForProduct_STATUS interface {
	AssignPropertiesFrom(src *storage.Product_STATUS) error
	AssignPropertiesTo(dst *storage.Product_STATUS) error
}

// Storage version of v1api20230501preview.ProductOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ProductOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_ProductOperatorSpec populates our ProductOperatorSpec from the provided source ProductOperatorSpec
func (operator *ProductOperatorSpec) AssignProperties_From_ProductOperatorSpec(source *storage.ProductOperatorSpec) error {
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

	// Invoke the augmentConversionForProductOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForProductOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_ProductOperatorSpec populates the provided destination ProductOperatorSpec from our ProductOperatorSpec
func (operator *ProductOperatorSpec) AssignProperties_To_ProductOperatorSpec(destination *storage.ProductOperatorSpec) error {
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

	// Invoke the augmentConversionForProductOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForProductOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForProductOperatorSpec interface {
	AssignPropertiesFrom(src *storage.ProductOperatorSpec) error
	AssignPropertiesTo(dst *storage.ProductOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&Product{}, &ProductList{})
}
