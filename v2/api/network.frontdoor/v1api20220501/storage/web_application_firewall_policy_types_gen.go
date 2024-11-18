// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.frontdoor.azure.com,resources=webapplicationfirewallpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.frontdoor.azure.com,resources={webapplicationfirewallpolicies/status,webapplicationfirewallpolicies/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220501.WebApplicationFirewallPolicy
// Generator information:
// - Generated from: /frontdoor/resource-manager/Microsoft.Network/stable/2022-05-01/webapplicationfirewall.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallPolicies/{policyName}
type WebApplicationFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebApplicationFirewallPolicy_Spec   `json:"spec,omitempty"`
	Status            WebApplicationFirewallPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &WebApplicationFirewallPolicy{}

// GetConditions returns the conditions of the resource
func (policy *WebApplicationFirewallPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *WebApplicationFirewallPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ configmaps.Exporter = &WebApplicationFirewallPolicy{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (policy *WebApplicationFirewallPolicy) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &WebApplicationFirewallPolicy{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (policy *WebApplicationFirewallPolicy) SecretDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &WebApplicationFirewallPolicy{}

// AzureName returns the Azure name of the resource
func (policy *WebApplicationFirewallPolicy) AzureName() string {
	return policy.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-05-01"
func (policy WebApplicationFirewallPolicy) GetAPIVersion() string {
	return "2022-05-01"
}

// GetResourceScope returns the scope of the resource
func (policy *WebApplicationFirewallPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *WebApplicationFirewallPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *WebApplicationFirewallPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *WebApplicationFirewallPolicy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/FrontDoorWebApplicationFirewallPolicies"
func (policy *WebApplicationFirewallPolicy) GetType() string {
	return "Microsoft.Network/FrontDoorWebApplicationFirewallPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *WebApplicationFirewallPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &WebApplicationFirewallPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *WebApplicationFirewallPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *WebApplicationFirewallPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*WebApplicationFirewallPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st WebApplicationFirewallPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// Hub marks that this WebApplicationFirewallPolicy is the hub type for conversion
func (policy *WebApplicationFirewallPolicy) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *WebApplicationFirewallPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion,
		Kind:    "WebApplicationFirewallPolicy",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220501.WebApplicationFirewallPolicy
// Generator information:
// - Generated from: /frontdoor/resource-manager/Microsoft.Network/stable/2022-05-01/webapplicationfirewall.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallPolicies/{policyName}
type WebApplicationFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebApplicationFirewallPolicy `json:"items"`
}

// Storage version of v1api20220501.APIVersion
// +kubebuilder:validation:Enum={"2022-05-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-05-01")

// Storage version of v1api20220501.WebApplicationFirewallPolicy_Spec
type WebApplicationFirewallPolicy_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                    `json:"azureName,omitempty"`
	CustomRules     *CustomRuleList                           `json:"customRules,omitempty"`
	Etag            *string                                   `json:"etag,omitempty"`
	Location        *string                                   `json:"location,omitempty"`
	ManagedRules    *ManagedRuleSetList                       `json:"managedRules,omitempty"`
	OperatorSpec    *WebApplicationFirewallPolicyOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                    `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner          *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PolicySettings *PolicySettings                    `json:"policySettings,omitempty"`
	PropertyBag    genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku            *Sku                               `json:"sku,omitempty"`
	Tags           map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &WebApplicationFirewallPolicy_Spec{}

// ConvertSpecFrom populates our WebApplicationFirewallPolicy_Spec from the provided source
func (policy *WebApplicationFirewallPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(policy)
}

// ConvertSpecTo populates the provided destination from our WebApplicationFirewallPolicy_Spec
func (policy *WebApplicationFirewallPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(policy)
}

// Storage version of v1api20220501.WebApplicationFirewallPolicy_STATUS
type WebApplicationFirewallPolicy_STATUS struct {
	Conditions            []conditions.Condition        `json:"conditions,omitempty"`
	CustomRules           *CustomRuleList_STATUS        `json:"customRules,omitempty"`
	Etag                  *string                       `json:"etag,omitempty"`
	FrontendEndpointLinks []FrontendEndpointLink_STATUS `json:"frontendEndpointLinks,omitempty"`
	Id                    *string                       `json:"id,omitempty"`
	Location              *string                       `json:"location,omitempty"`
	ManagedRules          *ManagedRuleSetList_STATUS    `json:"managedRules,omitempty"`
	Name                  *string                       `json:"name,omitempty"`
	PolicySettings        *PolicySettings_STATUS        `json:"policySettings,omitempty"`
	PropertyBag           genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	ProvisioningState     *string                       `json:"provisioningState,omitempty"`
	ResourceState         *string                       `json:"resourceState,omitempty"`
	RoutingRuleLinks      []RoutingRuleLink_STATUS      `json:"routingRuleLinks,omitempty"`
	SecurityPolicyLinks   []SecurityPolicyLink_STATUS   `json:"securityPolicyLinks,omitempty"`
	Sku                   *Sku_STATUS                   `json:"sku,omitempty"`
	Tags                  map[string]string             `json:"tags,omitempty"`
	Type                  *string                       `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &WebApplicationFirewallPolicy_STATUS{}

// ConvertStatusFrom populates our WebApplicationFirewallPolicy_STATUS from the provided source
func (policy *WebApplicationFirewallPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(policy)
}

// ConvertStatusTo populates the provided destination from our WebApplicationFirewallPolicy_STATUS
func (policy *WebApplicationFirewallPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(policy)
}

// Storage version of v1api20220501.CustomRuleList
// Defines contents of custom rules
type CustomRuleList struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rules       []CustomRule           `json:"rules,omitempty"`
}

// Storage version of v1api20220501.CustomRuleList_STATUS
// Defines contents of custom rules
type CustomRuleList_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rules       []CustomRule_STATUS    `json:"rules,omitempty"`
}

// Storage version of v1api20220501.FrontendEndpointLink_STATUS
// Defines the Resource ID for a Frontend Endpoint.
type FrontendEndpointLink_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220501.ManagedRuleSetList
// Defines the list of managed rule sets for the policy.
type ManagedRuleSetList struct {
	ManagedRuleSets []ManagedRuleSet       `json:"managedRuleSets,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220501.ManagedRuleSetList_STATUS
// Defines the list of managed rule sets for the policy.
type ManagedRuleSetList_STATUS struct {
	ManagedRuleSets []ManagedRuleSet_STATUS `json:"managedRuleSets,omitempty"`
	PropertyBag     genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220501.PolicySettings
// Defines top-level WebApplicationFirewallPolicy configuration settings.
type PolicySettings struct {
	CustomBlockResponseBody       *string                `json:"customBlockResponseBody,omitempty"`
	CustomBlockResponseStatusCode *int                   `json:"customBlockResponseStatusCode,omitempty"`
	EnabledState                  *string                `json:"enabledState,omitempty"`
	Mode                          *string                `json:"mode,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RedirectUrl                   *string                `json:"redirectUrl,omitempty"`
	RequestBodyCheck              *string                `json:"requestBodyCheck,omitempty"`
}

// Storage version of v1api20220501.PolicySettings_STATUS
// Defines top-level WebApplicationFirewallPolicy configuration settings.
type PolicySettings_STATUS struct {
	CustomBlockResponseBody       *string                `json:"customBlockResponseBody,omitempty"`
	CustomBlockResponseStatusCode *int                   `json:"customBlockResponseStatusCode,omitempty"`
	EnabledState                  *string                `json:"enabledState,omitempty"`
	Mode                          *string                `json:"mode,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RedirectUrl                   *string                `json:"redirectUrl,omitempty"`
	RequestBodyCheck              *string                `json:"requestBodyCheck,omitempty"`
}

// Storage version of v1api20220501.RoutingRuleLink_STATUS
// Defines the Resource ID for a Routing Rule.
type RoutingRuleLink_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220501.SecurityPolicyLink_STATUS
// Defines the Resource ID for a Security Policy.
type SecurityPolicyLink_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220501.Sku
// The pricing tier of the web application firewall policy.
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220501.Sku_STATUS
// The pricing tier of the web application firewall policy.
type Sku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220501.WebApplicationFirewallPolicyOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type WebApplicationFirewallPolicyOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20220501.CustomRule
// Defines contents of a web application rule
type CustomRule struct {
	Action                     *string                `json:"action,omitempty"`
	EnabledState               *string                `json:"enabledState,omitempty"`
	MatchConditions            []MatchCondition       `json:"matchConditions,omitempty"`
	Name                       *string                `json:"name,omitempty"`
	Priority                   *int                   `json:"priority,omitempty"`
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RateLimitDurationInMinutes *int                   `json:"rateLimitDurationInMinutes,omitempty"`
	RateLimitThreshold         *int                   `json:"rateLimitThreshold,omitempty"`
	RuleType                   *string                `json:"ruleType,omitempty"`
}

// Storage version of v1api20220501.CustomRule_STATUS
// Defines contents of a web application rule
type CustomRule_STATUS struct {
	Action                     *string                 `json:"action,omitempty"`
	EnabledState               *string                 `json:"enabledState,omitempty"`
	MatchConditions            []MatchCondition_STATUS `json:"matchConditions,omitempty"`
	Name                       *string                 `json:"name,omitempty"`
	Priority                   *int                    `json:"priority,omitempty"`
	PropertyBag                genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	RateLimitDurationInMinutes *int                    `json:"rateLimitDurationInMinutes,omitempty"`
	RateLimitThreshold         *int                    `json:"rateLimitThreshold,omitempty"`
	RuleType                   *string                 `json:"ruleType,omitempty"`
}

// Storage version of v1api20220501.ManagedRuleSet
// Defines a managed rule set.
type ManagedRuleSet struct {
	Exclusions         []ManagedRuleExclusion     `json:"exclusions,omitempty"`
	PropertyBag        genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	RuleGroupOverrides []ManagedRuleGroupOverride `json:"ruleGroupOverrides,omitempty"`
	RuleSetAction      *string                    `json:"ruleSetAction,omitempty"`
	RuleSetType        *string                    `json:"ruleSetType,omitempty"`
	RuleSetVersion     *string                    `json:"ruleSetVersion,omitempty"`
}

// Storage version of v1api20220501.ManagedRuleSet_STATUS
// Defines a managed rule set.
type ManagedRuleSet_STATUS struct {
	Exclusions         []ManagedRuleExclusion_STATUS     `json:"exclusions,omitempty"`
	PropertyBag        genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	RuleGroupOverrides []ManagedRuleGroupOverride_STATUS `json:"ruleGroupOverrides,omitempty"`
	RuleSetAction      *string                           `json:"ruleSetAction,omitempty"`
	RuleSetType        *string                           `json:"ruleSetType,omitempty"`
	RuleSetVersion     *string                           `json:"ruleSetVersion,omitempty"`
}

// Storage version of v1api20220501.ManagedRuleExclusion
// Exclude variables from managed rule evaluation.
type ManagedRuleExclusion struct {
	MatchVariable         *string                `json:"matchVariable,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Selector              *string                `json:"selector,omitempty"`
	SelectorMatchOperator *string                `json:"selectorMatchOperator,omitempty"`
}

// Storage version of v1api20220501.ManagedRuleExclusion_STATUS
// Exclude variables from managed rule evaluation.
type ManagedRuleExclusion_STATUS struct {
	MatchVariable         *string                `json:"matchVariable,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Selector              *string                `json:"selector,omitempty"`
	SelectorMatchOperator *string                `json:"selectorMatchOperator,omitempty"`
}

// Storage version of v1api20220501.ManagedRuleGroupOverride
// Defines a managed rule group override setting.
type ManagedRuleGroupOverride struct {
	Exclusions    []ManagedRuleExclusion `json:"exclusions,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RuleGroupName *string                `json:"ruleGroupName,omitempty"`
	Rules         []ManagedRuleOverride  `json:"rules,omitempty"`
}

// Storage version of v1api20220501.ManagedRuleGroupOverride_STATUS
// Defines a managed rule group override setting.
type ManagedRuleGroupOverride_STATUS struct {
	Exclusions    []ManagedRuleExclusion_STATUS `json:"exclusions,omitempty"`
	PropertyBag   genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RuleGroupName *string                       `json:"ruleGroupName,omitempty"`
	Rules         []ManagedRuleOverride_STATUS  `json:"rules,omitempty"`
}

// Storage version of v1api20220501.MatchCondition
// Define a match condition.
type MatchCondition struct {
	MatchValue      []string               `json:"matchValue,omitempty"`
	MatchVariable   *string                `json:"matchVariable,omitempty"`
	NegateCondition *bool                  `json:"negateCondition,omitempty"`
	Operator        *string                `json:"operator,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Selector        *string                `json:"selector,omitempty"`
	Transforms      []string               `json:"transforms,omitempty"`
}

// Storage version of v1api20220501.MatchCondition_STATUS
// Define a match condition.
type MatchCondition_STATUS struct {
	MatchValue      []string               `json:"matchValue,omitempty"`
	MatchVariable   *string                `json:"matchVariable,omitempty"`
	NegateCondition *bool                  `json:"negateCondition,omitempty"`
	Operator        *string                `json:"operator,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Selector        *string                `json:"selector,omitempty"`
	Transforms      []string               `json:"transforms,omitempty"`
}

// Storage version of v1api20220501.ManagedRuleOverride
// Defines a managed rule group override setting.
type ManagedRuleOverride struct {
	Action       *string                `json:"action,omitempty"`
	EnabledState *string                `json:"enabledState,omitempty"`
	Exclusions   []ManagedRuleExclusion `json:"exclusions,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RuleId       *string                `json:"ruleId,omitempty"`
}

// Storage version of v1api20220501.ManagedRuleOverride_STATUS
// Defines a managed rule group override setting.
type ManagedRuleOverride_STATUS struct {
	Action       *string                       `json:"action,omitempty"`
	EnabledState *string                       `json:"enabledState,omitempty"`
	Exclusions   []ManagedRuleExclusion_STATUS `json:"exclusions,omitempty"`
	PropertyBag  genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RuleId       *string                       `json:"ruleId,omitempty"`
}

func init() {
	SchemeBuilder.Register(&WebApplicationFirewallPolicy{}, &WebApplicationFirewallPolicyList{})
}
