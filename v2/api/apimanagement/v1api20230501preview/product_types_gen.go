// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimproducts.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}
type Product struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Service_Product_Spec   `json:"spec,omitempty"`
	Status            Service_Product_STATUS `json:"status,omitempty"`
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
	// intermediate variable for conversion
	var source storage.Product

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = product.AssignProperties_From_Product(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to product")
	}

	return nil
}

// ConvertTo populates the provided hub Product from our Product
func (product *Product) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.Product
	err := product.AssignProperties_To_Product(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from product")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-apimanagement-azure-com-v1api20230501preview-product,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=products,verbs=create;update,versions=v1api20230501preview,name=default.v1api20230501preview.products.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &Product{}

// Default applies defaults to the Product resource
func (product *Product) Default() {
	product.defaultImpl()
	var temp any = product
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (product *Product) defaultAzureName() {
	if product.Spec.AzureName == "" {
		product.Spec.AzureName = product.Name
	}
}

// defaultImpl applies the code generated defaults to the Product resource
func (product *Product) defaultImpl() { product.defaultAzureName() }

var _ genruntime.KubernetesResource = &Product{}

// AzureName returns the Azure name of the resource
func (product *Product) AzureName() string {
	return product.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (product Product) GetAPIVersion() string {
	return string(APIVersion_Value)
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
	return &Service_Product_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (product *Product) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(product.Spec)
	return product.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (product *Product) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Service_Product_STATUS); ok {
		product.Status = *st
		return nil
	}

	// Convert status to required version
	var st Service_Product_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	product.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-apimanagement-azure-com-v1api20230501preview-product,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=products,verbs=create;update,versions=v1api20230501preview,name=validate.v1api20230501preview.products.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &Product{}

// ValidateCreate validates the creation of the resource
func (product *Product) ValidateCreate() (admission.Warnings, error) {
	validations := product.createValidations()
	var temp any = product
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (product *Product) ValidateDelete() (admission.Warnings, error) {
	validations := product.deleteValidations()
	var temp any = product
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (product *Product) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := product.updateValidations()
	var temp any = product
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (product *Product) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){product.validateResourceReferences, product.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (product *Product) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (product *Product) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return product.validateResourceReferences()
		},
		product.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return product.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (product *Product) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(product)
}

// validateResourceReferences validates all resource references
func (product *Product) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&product.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (product *Product) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*Product)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, product)
}

// AssignProperties_From_Product populates our Product from the provided source Product
func (product *Product) AssignProperties_From_Product(source *storage.Product) error {

	// ObjectMeta
	product.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Service_Product_Spec
	err := spec.AssignProperties_From_Service_Product_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Service_Product_Spec() to populate field Spec")
	}
	product.Spec = spec

	// Status
	var status Service_Product_STATUS
	err = status.AssignProperties_From_Service_Product_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Service_Product_STATUS() to populate field Status")
	}
	product.Status = status

	// No error
	return nil
}

// AssignProperties_To_Product populates the provided destination Product from our Product
func (product *Product) AssignProperties_To_Product(destination *storage.Product) error {

	// ObjectMeta
	destination.ObjectMeta = *product.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Service_Product_Spec
	err := product.Spec.AssignProperties_To_Service_Product_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Service_Product_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Service_Product_STATUS
	err = product.Status.AssignProperties_To_Service_Product_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Service_Product_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (product *Product) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: product.Spec.OriginalVersion(),
		Kind:    "Product",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimproducts.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}
type ProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Product `json:"items"`
}

type Service_Product_Spec struct {
	// ApprovalRequired: whether subscription approval is required. If false, new subscriptions will be approved automatically
	// enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually
	// approve the subscription before the developer can any of the product’s APIs. Can be present only if
	// subscriptionRequired property is present and has a value of false.
	ApprovalRequired *bool `json:"approvalRequired,omitempty"`

	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:MaxLength=1000
	// +kubebuilder:validation:MinLength=0
	// Description: Product description. May include HTML formatting tags.
	Description *string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=300
	// +kubebuilder:validation:MinLength=1
	// DisplayName: Product name.
	DisplayName *string `json:"displayName,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`

	// State: whether product is published or not. Published products are discoverable by users of developer portal. Non
	// published products are visible only to administrators. Default state of Product is notPublished.
	State *ProductContractProperties_State `json:"state,omitempty"`

	// SubscriptionRequired: Whether a product subscription is required for accessing APIs included in this product. If true,
	// the product is referred to as "protected" and a valid subscription key is required for a request to an API included in
	// the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product
	// can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be
	// true.
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty"`

	// SubscriptionsLimit: Whether the number of subscriptions a user can have to this product at the same time. Set to null or
	// omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has
	// a value of false.
	SubscriptionsLimit *int `json:"subscriptionsLimit,omitempty"`

	// Terms: Product terms of use. Developers trying to subscribe to the product will be presented and required to accept
	// these terms before they can complete the subscription process.
	Terms *string `json:"terms,omitempty"`
}

var _ genruntime.ARMTransformer = &Service_Product_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (product *Service_Product_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if product == nil {
		return nil, nil
	}
	result := &Service_Product_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if product.ApprovalRequired != nil ||
		product.Description != nil ||
		product.DisplayName != nil ||
		product.State != nil ||
		product.SubscriptionRequired != nil ||
		product.SubscriptionsLimit != nil ||
		product.Terms != nil {
		result.Properties = &ProductContractProperties_ARM{}
	}
	if product.ApprovalRequired != nil {
		approvalRequired := *product.ApprovalRequired
		result.Properties.ApprovalRequired = &approvalRequired
	}
	if product.Description != nil {
		description := *product.Description
		result.Properties.Description = &description
	}
	if product.DisplayName != nil {
		displayName := *product.DisplayName
		result.Properties.DisplayName = &displayName
	}
	if product.State != nil {
		state := *product.State
		result.Properties.State = &state
	}
	if product.SubscriptionRequired != nil {
		subscriptionRequired := *product.SubscriptionRequired
		result.Properties.SubscriptionRequired = &subscriptionRequired
	}
	if product.SubscriptionsLimit != nil {
		subscriptionsLimit := *product.SubscriptionsLimit
		result.Properties.SubscriptionsLimit = &subscriptionsLimit
	}
	if product.Terms != nil {
		terms := *product.Terms
		result.Properties.Terms = &terms
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (product *Service_Product_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Service_Product_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (product *Service_Product_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Service_Product_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Service_Product_Spec_ARM, got %T", armInput)
	}

	// Set property "ApprovalRequired":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ApprovalRequired != nil {
			approvalRequired := *typedInput.Properties.ApprovalRequired
			product.ApprovalRequired = &approvalRequired
		}
	}

	// Set property "AzureName":
	product.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Description":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			product.Description = &description
		}
	}

	// Set property "DisplayName":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DisplayName != nil {
			displayName := *typedInput.Properties.DisplayName
			product.DisplayName = &displayName
		}
	}

	// Set property "Owner":
	product.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "State":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			product.State = &state
		}
	}

	// Set property "SubscriptionRequired":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.SubscriptionRequired != nil {
			subscriptionRequired := *typedInput.Properties.SubscriptionRequired
			product.SubscriptionRequired = &subscriptionRequired
		}
	}

	// Set property "SubscriptionsLimit":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.SubscriptionsLimit != nil {
			subscriptionsLimit := *typedInput.Properties.SubscriptionsLimit
			product.SubscriptionsLimit = &subscriptionsLimit
		}
	}

	// Set property "Terms":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Terms != nil {
			terms := *typedInput.Properties.Terms
			product.Terms = &terms
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Service_Product_Spec{}

// ConvertSpecFrom populates our Service_Product_Spec from the provided source
func (product *Service_Product_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Service_Product_Spec)
	if ok {
		// Populate our instance from source
		return product.AssignProperties_From_Service_Product_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Service_Product_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = product.AssignProperties_From_Service_Product_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Service_Product_Spec
func (product *Service_Product_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Service_Product_Spec)
	if ok {
		// Populate destination from our instance
		return product.AssignProperties_To_Service_Product_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Service_Product_Spec{}
	err := product.AssignProperties_To_Service_Product_Spec(dst)
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

// AssignProperties_From_Service_Product_Spec populates our Service_Product_Spec from the provided source Service_Product_Spec
func (product *Service_Product_Spec) AssignProperties_From_Service_Product_Spec(source *storage.Service_Product_Spec) error {

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
	if source.Description != nil {
		description := *source.Description
		product.Description = &description
	} else {
		product.Description = nil
	}

	// DisplayName
	if source.DisplayName != nil {
		displayName := *source.DisplayName
		product.DisplayName = &displayName
	} else {
		product.DisplayName = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		product.Owner = &owner
	} else {
		product.Owner = nil
	}

	// State
	if source.State != nil {
		state := *source.State
		stateTemp := genruntime.ToEnum(state, productContractProperties_State_Values)
		product.State = &stateTemp
	} else {
		product.State = nil
	}

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

	// No error
	return nil
}

// AssignProperties_To_Service_Product_Spec populates the provided destination Service_Product_Spec from our Service_Product_Spec
func (product *Service_Product_Spec) AssignProperties_To_Service_Product_Spec(destination *storage.Service_Product_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

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
	if product.Description != nil {
		description := *product.Description
		destination.Description = &description
	} else {
		destination.Description = nil
	}

	// DisplayName
	if product.DisplayName != nil {
		displayName := *product.DisplayName
		destination.DisplayName = &displayName
	} else {
		destination.DisplayName = nil
	}

	// OriginalVersion
	destination.OriginalVersion = product.OriginalVersion()

	// Owner
	if product.Owner != nil {
		owner := product.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// State
	if product.State != nil {
		state := string(*product.State)
		destination.State = &state
	} else {
		destination.State = nil
	}

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

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (product *Service_Product_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (product *Service_Product_Spec) SetAzureName(azureName string) { product.AzureName = azureName }

type Service_Product_STATUS struct {
	// ApprovalRequired: whether subscription approval is required. If false, new subscriptions will be approved automatically
	// enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually
	// approve the subscription before the developer can any of the product’s APIs. Can be present only if
	// subscriptionRequired property is present and has a value of false.
	ApprovalRequired *bool `json:"approvalRequired,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Description: Product description. May include HTML formatting tags.
	Description *string `json:"description,omitempty"`

	// DisplayName: Product name.
	DisplayName *string `json:"displayName,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// State: whether product is published or not. Published products are discoverable by users of developer portal. Non
	// published products are visible only to administrators. Default state of Product is notPublished.
	State *ProductContractProperties_State_STATUS `json:"state,omitempty"`

	// SubscriptionRequired: Whether a product subscription is required for accessing APIs included in this product. If true,
	// the product is referred to as "protected" and a valid subscription key is required for a request to an API included in
	// the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product
	// can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be
	// true.
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty"`

	// SubscriptionsLimit: Whether the number of subscriptions a user can have to this product at the same time. Set to null or
	// omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has
	// a value of false.
	SubscriptionsLimit *int `json:"subscriptionsLimit,omitempty"`

	// Terms: Product terms of use. Developers trying to subscribe to the product will be presented and required to accept
	// these terms before they can complete the subscription process.
	Terms *string `json:"terms,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Service_Product_STATUS{}

// ConvertStatusFrom populates our Service_Product_STATUS from the provided source
func (product *Service_Product_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Service_Product_STATUS)
	if ok {
		// Populate our instance from source
		return product.AssignProperties_From_Service_Product_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Service_Product_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = product.AssignProperties_From_Service_Product_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Service_Product_STATUS
func (product *Service_Product_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Service_Product_STATUS)
	if ok {
		// Populate destination from our instance
		return product.AssignProperties_To_Service_Product_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Service_Product_STATUS{}
	err := product.AssignProperties_To_Service_Product_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Service_Product_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (product *Service_Product_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Service_Product_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (product *Service_Product_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Service_Product_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Service_Product_STATUS_ARM, got %T", armInput)
	}

	// Set property "ApprovalRequired":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ApprovalRequired != nil {
			approvalRequired := *typedInput.Properties.ApprovalRequired
			product.ApprovalRequired = &approvalRequired
		}
	}

	// no assignment for property "Conditions"

	// Set property "Description":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			product.Description = &description
		}
	}

	// Set property "DisplayName":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DisplayName != nil {
			displayName := *typedInput.Properties.DisplayName
			product.DisplayName = &displayName
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		product.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		product.Name = &name
	}

	// Set property "State":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			product.State = &state
		}
	}

	// Set property "SubscriptionRequired":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.SubscriptionRequired != nil {
			subscriptionRequired := *typedInput.Properties.SubscriptionRequired
			product.SubscriptionRequired = &subscriptionRequired
		}
	}

	// Set property "SubscriptionsLimit":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.SubscriptionsLimit != nil {
			subscriptionsLimit := *typedInput.Properties.SubscriptionsLimit
			product.SubscriptionsLimit = &subscriptionsLimit
		}
	}

	// Set property "Terms":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Terms != nil {
			terms := *typedInput.Properties.Terms
			product.Terms = &terms
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		product.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Service_Product_STATUS populates our Service_Product_STATUS from the provided source Service_Product_STATUS
func (product *Service_Product_STATUS) AssignProperties_From_Service_Product_STATUS(source *storage.Service_Product_STATUS) error {

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
	if source.State != nil {
		state := *source.State
		stateTemp := genruntime.ToEnum(state, productContractProperties_State_STATUS_Values)
		product.State = &stateTemp
	} else {
		product.State = nil
	}

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

	// No error
	return nil
}

// AssignProperties_To_Service_Product_STATUS populates the provided destination Service_Product_STATUS from our Service_Product_STATUS
func (product *Service_Product_STATUS) AssignProperties_To_Service_Product_STATUS(destination *storage.Service_Product_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

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
	if product.State != nil {
		state := string(*product.State)
		destination.State = &state
	} else {
		destination.State = nil
	}

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

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"notPublished","published"}
type ProductContractProperties_State string

const (
	ProductContractProperties_State_NotPublished = ProductContractProperties_State("notPublished")
	ProductContractProperties_State_Published    = ProductContractProperties_State("published")
)

// Mapping from string to ProductContractProperties_State
var productContractProperties_State_Values = map[string]ProductContractProperties_State{
	"notpublished": ProductContractProperties_State_NotPublished,
	"published":    ProductContractProperties_State_Published,
}

type ProductContractProperties_State_STATUS string

const (
	ProductContractProperties_State_STATUS_NotPublished = ProductContractProperties_State_STATUS("notPublished")
	ProductContractProperties_State_STATUS_Published    = ProductContractProperties_State_STATUS("published")
)

// Mapping from string to ProductContractProperties_State_STATUS
var productContractProperties_State_STATUS_Values = map[string]ProductContractProperties_State_STATUS{
	"notpublished": ProductContractProperties_State_STATUS_NotPublished,
	"published":    ProductContractProperties_State_STATUS_Published,
}

func init() {
	SchemeBuilder.Register(&Product{}, &ProductList{})
}
