// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type WebApplicationFirewallPolicy_Spec struct {
	// Etag: Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the web application firewall policy.
	Properties *WebApplicationFirewallPolicyProperties `json:"properties,omitempty"`

	// Sku: The pricing tier of web application firewall policy. Defaults to Classic_AzureFrontDoor if not specified.
	Sku *Sku `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &WebApplicationFirewallPolicy_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-05-01"
func (policy WebApplicationFirewallPolicy_Spec) GetAPIVersion() string {
	return "2022-05-01"
}

// GetName returns the Name of the resource
func (policy *WebApplicationFirewallPolicy_Spec) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/FrontDoorWebApplicationFirewallPolicies"
func (policy *WebApplicationFirewallPolicy_Spec) GetType() string {
	return "Microsoft.Network/FrontDoorWebApplicationFirewallPolicies"
}

// The pricing tier of the web application firewall policy.
type Sku struct {
	// Name: Name of the pricing tier.
	Name *Sku_Name `json:"name,omitempty"`
}

// Defines web application firewall policy properties.
type WebApplicationFirewallPolicyProperties struct {
	// CustomRules: Describes custom rules inside the policy.
	CustomRules *CustomRuleList `json:"customRules,omitempty"`

	// ManagedRules: Describes managed rules inside the policy.
	ManagedRules *ManagedRuleSetList `json:"managedRules,omitempty"`

	// PolicySettings: Describes settings for the policy.
	PolicySettings *PolicySettings `json:"policySettings,omitempty"`
}

// Defines contents of custom rules
type CustomRuleList struct {
	// Rules: List of rules
	Rules []CustomRule `json:"rules,omitempty"`
}

// Defines the list of managed rule sets for the policy.
type ManagedRuleSetList struct {
	// ManagedRuleSets: List of rule sets.
	ManagedRuleSets []ManagedRuleSet `json:"managedRuleSets,omitempty"`
}

// Defines top-level WebApplicationFirewallPolicy configuration settings.
type PolicySettings struct {
	// CustomBlockResponseBody: If the action type is block, customer can override the response body. The body must be
	// specified in base64 encoding.
	CustomBlockResponseBody *string `json:"customBlockResponseBody,omitempty"`

	// CustomBlockResponseStatusCode: If the action type is block, customer can override the response status code.
	CustomBlockResponseStatusCode *int `json:"customBlockResponseStatusCode,omitempty"`

	// EnabledState: Describes if the policy is in enabled or disabled state. Defaults to Enabled if not specified.
	EnabledState *PolicySettings_EnabledState `json:"enabledState,omitempty"`

	// Mode: Describes if it is in detection mode or prevention mode at policy level.
	Mode *PolicySettings_Mode `json:"mode,omitempty"`

	// RedirectUrl: If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl *string `json:"redirectUrl,omitempty"`

	// RequestBodyCheck: Describes if policy managed rules will inspect the request body content.
	RequestBodyCheck *PolicySettings_RequestBodyCheck `json:"requestBodyCheck,omitempty"`
}

// +kubebuilder:validation:Enum={"Classic_AzureFrontDoor","Premium_AzureFrontDoor","Standard_AzureFrontDoor"}
type Sku_Name string

const (
	Sku_Name_Classic_AzureFrontDoor  = Sku_Name("Classic_AzureFrontDoor")
	Sku_Name_Premium_AzureFrontDoor  = Sku_Name("Premium_AzureFrontDoor")
	Sku_Name_Standard_AzureFrontDoor = Sku_Name("Standard_AzureFrontDoor")
)

// Mapping from string to Sku_Name
var sku_Name_Values = map[string]Sku_Name{
	"classic_azurefrontdoor":  Sku_Name_Classic_AzureFrontDoor,
	"premium_azurefrontdoor":  Sku_Name_Premium_AzureFrontDoor,
	"standard_azurefrontdoor": Sku_Name_Standard_AzureFrontDoor,
}

// Defines contents of a web application rule
type CustomRule struct {
	// Action: Describes what action to be applied when rule matches.
	Action *ActionType `json:"action,omitempty"`

	// EnabledState: Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
	EnabledState *CustomRule_EnabledState `json:"enabledState,omitempty"`

	// MatchConditions: List of match conditions.
	MatchConditions []MatchCondition `json:"matchConditions,omitempty"`

	// Name: Describes the name of the rule.
	Name *string `json:"name,omitempty"`

	// Priority: Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
	Priority *int `json:"priority,omitempty"`

	// RateLimitDurationInMinutes: Time window for resetting the rate limit count. Default is 1 minute.
	RateLimitDurationInMinutes *int `json:"rateLimitDurationInMinutes,omitempty"`

	// RateLimitThreshold: Number of allowed requests per client within the time window.
	RateLimitThreshold *int `json:"rateLimitThreshold,omitempty"`

	// RuleType: Describes type of rule.
	RuleType *CustomRule_RuleType `json:"ruleType,omitempty"`
}

// Defines a managed rule set.
type ManagedRuleSet struct {
	// Exclusions: Describes the exclusions that are applied to all rules in the set.
	Exclusions []ManagedRuleExclusion `json:"exclusions,omitempty"`

	// RuleGroupOverrides: Defines the rule group overrides to apply to the rule set.
	RuleGroupOverrides []ManagedRuleGroupOverride `json:"ruleGroupOverrides,omitempty"`

	// RuleSetAction: Defines the rule set action.
	RuleSetAction *ManagedRuleSetActionType `json:"ruleSetAction,omitempty"`

	// RuleSetType: Defines the rule set type to use.
	RuleSetType *string `json:"ruleSetType,omitempty"`

	// RuleSetVersion: Defines the version of the rule set to use.
	RuleSetVersion *string `json:"ruleSetVersion,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type PolicySettings_EnabledState string

const (
	PolicySettings_EnabledState_Disabled = PolicySettings_EnabledState("Disabled")
	PolicySettings_EnabledState_Enabled  = PolicySettings_EnabledState("Enabled")
)

// Mapping from string to PolicySettings_EnabledState
var policySettings_EnabledState_Values = map[string]PolicySettings_EnabledState{
	"disabled": PolicySettings_EnabledState_Disabled,
	"enabled":  PolicySettings_EnabledState_Enabled,
}

// +kubebuilder:validation:Enum={"Detection","Prevention"}
type PolicySettings_Mode string

const (
	PolicySettings_Mode_Detection  = PolicySettings_Mode("Detection")
	PolicySettings_Mode_Prevention = PolicySettings_Mode("Prevention")
)

// Mapping from string to PolicySettings_Mode
var policySettings_Mode_Values = map[string]PolicySettings_Mode{
	"detection":  PolicySettings_Mode_Detection,
	"prevention": PolicySettings_Mode_Prevention,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type PolicySettings_RequestBodyCheck string

const (
	PolicySettings_RequestBodyCheck_Disabled = PolicySettings_RequestBodyCheck("Disabled")
	PolicySettings_RequestBodyCheck_Enabled  = PolicySettings_RequestBodyCheck("Enabled")
)

// Mapping from string to PolicySettings_RequestBodyCheck
var policySettings_RequestBodyCheck_Values = map[string]PolicySettings_RequestBodyCheck{
	"disabled": PolicySettings_RequestBodyCheck_Disabled,
	"enabled":  PolicySettings_RequestBodyCheck_Enabled,
}

// Defines the action to take on rule match.
// +kubebuilder:validation:Enum={"Allow","AnomalyScoring","Block","Log","Redirect"}
type ActionType string

const (
	ActionType_Allow          = ActionType("Allow")
	ActionType_AnomalyScoring = ActionType("AnomalyScoring")
	ActionType_Block          = ActionType("Block")
	ActionType_Log            = ActionType("Log")
	ActionType_Redirect       = ActionType("Redirect")
)

// Mapping from string to ActionType
var actionType_Values = map[string]ActionType{
	"allow":          ActionType_Allow,
	"anomalyscoring": ActionType_AnomalyScoring,
	"block":          ActionType_Block,
	"log":            ActionType_Log,
	"redirect":       ActionType_Redirect,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type CustomRule_EnabledState string

const (
	CustomRule_EnabledState_Disabled = CustomRule_EnabledState("Disabled")
	CustomRule_EnabledState_Enabled  = CustomRule_EnabledState("Enabled")
)

// Mapping from string to CustomRule_EnabledState
var customRule_EnabledState_Values = map[string]CustomRule_EnabledState{
	"disabled": CustomRule_EnabledState_Disabled,
	"enabled":  CustomRule_EnabledState_Enabled,
}

// +kubebuilder:validation:Enum={"MatchRule","RateLimitRule"}
type CustomRule_RuleType string

const (
	CustomRule_RuleType_MatchRule     = CustomRule_RuleType("MatchRule")
	CustomRule_RuleType_RateLimitRule = CustomRule_RuleType("RateLimitRule")
)

// Mapping from string to CustomRule_RuleType
var customRule_RuleType_Values = map[string]CustomRule_RuleType{
	"matchrule":     CustomRule_RuleType_MatchRule,
	"ratelimitrule": CustomRule_RuleType_RateLimitRule,
}

// Exclude variables from managed rule evaluation.
type ManagedRuleExclusion struct {
	// MatchVariable: The variable type to be excluded.
	MatchVariable *ManagedRuleExclusion_MatchVariable `json:"matchVariable,omitempty"`

	// Selector: Selector value for which elements in the collection this exclusion applies to.
	Selector *string `json:"selector,omitempty"`

	// SelectorMatchOperator: Comparison operator to apply to the selector when specifying which elements in the collection
	// this exclusion applies to.
	SelectorMatchOperator *ManagedRuleExclusion_SelectorMatchOperator `json:"selectorMatchOperator,omitempty"`
}

// Defines a managed rule group override setting.
type ManagedRuleGroupOverride struct {
	// Exclusions: Describes the exclusions that are applied to all rules in the group.
	Exclusions []ManagedRuleExclusion `json:"exclusions,omitempty"`

	// RuleGroupName: Describes the managed rule group to override.
	RuleGroupName *string `json:"ruleGroupName,omitempty"`

	// Rules: List of rules that will be disabled. If none specified, all rules in the group will be disabled.
	Rules []ManagedRuleOverride `json:"rules,omitempty"`
}

// Defines the action to take when a managed rule set score threshold is met.
// +kubebuilder:validation:Enum={"Block","Log","Redirect"}
type ManagedRuleSetActionType string

const (
	ManagedRuleSetActionType_Block    = ManagedRuleSetActionType("Block")
	ManagedRuleSetActionType_Log      = ManagedRuleSetActionType("Log")
	ManagedRuleSetActionType_Redirect = ManagedRuleSetActionType("Redirect")
)

// Mapping from string to ManagedRuleSetActionType
var managedRuleSetActionType_Values = map[string]ManagedRuleSetActionType{
	"block":    ManagedRuleSetActionType_Block,
	"log":      ManagedRuleSetActionType_Log,
	"redirect": ManagedRuleSetActionType_Redirect,
}

// Define a match condition.
type MatchCondition struct {
	// MatchValue: List of possible match values.
	MatchValue []string `json:"matchValue,omitempty"`

	// MatchVariable: Request variable to compare with.
	MatchVariable *MatchCondition_MatchVariable `json:"matchVariable,omitempty"`

	// NegateCondition: Describes if the result of this condition should be negated.
	NegateCondition *bool `json:"negateCondition,omitempty"`

	// Operator: Comparison type to use for matching with the variable value.
	Operator *MatchCondition_Operator `json:"operator,omitempty"`

	// Selector: Match against a specific key from the QueryString, PostArgs, RequestHeader or Cookies variables. Default is
	// null.
	Selector *string `json:"selector,omitempty"`

	// Transforms: List of transforms.
	Transforms []TransformType `json:"transforms,omitempty"`
}

// +kubebuilder:validation:Enum={"QueryStringArgNames","RequestBodyJsonArgNames","RequestBodyPostArgNames","RequestCookieNames","RequestHeaderNames"}
type ManagedRuleExclusion_MatchVariable string

const (
	ManagedRuleExclusion_MatchVariable_QueryStringArgNames     = ManagedRuleExclusion_MatchVariable("QueryStringArgNames")
	ManagedRuleExclusion_MatchVariable_RequestBodyJsonArgNames = ManagedRuleExclusion_MatchVariable("RequestBodyJsonArgNames")
	ManagedRuleExclusion_MatchVariable_RequestBodyPostArgNames = ManagedRuleExclusion_MatchVariable("RequestBodyPostArgNames")
	ManagedRuleExclusion_MatchVariable_RequestCookieNames      = ManagedRuleExclusion_MatchVariable("RequestCookieNames")
	ManagedRuleExclusion_MatchVariable_RequestHeaderNames      = ManagedRuleExclusion_MatchVariable("RequestHeaderNames")
)

// Mapping from string to ManagedRuleExclusion_MatchVariable
var managedRuleExclusion_MatchVariable_Values = map[string]ManagedRuleExclusion_MatchVariable{
	"querystringargnames":     ManagedRuleExclusion_MatchVariable_QueryStringArgNames,
	"requestbodyjsonargnames": ManagedRuleExclusion_MatchVariable_RequestBodyJsonArgNames,
	"requestbodypostargnames": ManagedRuleExclusion_MatchVariable_RequestBodyPostArgNames,
	"requestcookienames":      ManagedRuleExclusion_MatchVariable_RequestCookieNames,
	"requestheadernames":      ManagedRuleExclusion_MatchVariable_RequestHeaderNames,
}

// +kubebuilder:validation:Enum={"Contains","EndsWith","Equals","EqualsAny","StartsWith"}
type ManagedRuleExclusion_SelectorMatchOperator string

const (
	ManagedRuleExclusion_SelectorMatchOperator_Contains   = ManagedRuleExclusion_SelectorMatchOperator("Contains")
	ManagedRuleExclusion_SelectorMatchOperator_EndsWith   = ManagedRuleExclusion_SelectorMatchOperator("EndsWith")
	ManagedRuleExclusion_SelectorMatchOperator_Equals     = ManagedRuleExclusion_SelectorMatchOperator("Equals")
	ManagedRuleExclusion_SelectorMatchOperator_EqualsAny  = ManagedRuleExclusion_SelectorMatchOperator("EqualsAny")
	ManagedRuleExclusion_SelectorMatchOperator_StartsWith = ManagedRuleExclusion_SelectorMatchOperator("StartsWith")
)

// Mapping from string to ManagedRuleExclusion_SelectorMatchOperator
var managedRuleExclusion_SelectorMatchOperator_Values = map[string]ManagedRuleExclusion_SelectorMatchOperator{
	"contains":   ManagedRuleExclusion_SelectorMatchOperator_Contains,
	"endswith":   ManagedRuleExclusion_SelectorMatchOperator_EndsWith,
	"equals":     ManagedRuleExclusion_SelectorMatchOperator_Equals,
	"equalsany":  ManagedRuleExclusion_SelectorMatchOperator_EqualsAny,
	"startswith": ManagedRuleExclusion_SelectorMatchOperator_StartsWith,
}

// Defines a managed rule group override setting.
type ManagedRuleOverride struct {
	// Action: Describes the override action to be applied when rule matches.
	Action *ActionType `json:"action,omitempty"`

	// EnabledState: Describes if the managed rule is in enabled or disabled state. Defaults to Disabled if not specified.
	EnabledState *ManagedRuleEnabledState `json:"enabledState,omitempty"`

	// Exclusions: Describes the exclusions that are applied to this specific rule.
	Exclusions []ManagedRuleExclusion `json:"exclusions,omitempty"`

	// RuleId: Identifier for the managed rule.
	RuleId *string `json:"ruleId,omitempty"`
}

// +kubebuilder:validation:Enum={"Cookies","PostArgs","QueryString","RemoteAddr","RequestBody","RequestHeader","RequestMethod","RequestUri","SocketAddr"}
type MatchCondition_MatchVariable string

const (
	MatchCondition_MatchVariable_Cookies       = MatchCondition_MatchVariable("Cookies")
	MatchCondition_MatchVariable_PostArgs      = MatchCondition_MatchVariable("PostArgs")
	MatchCondition_MatchVariable_QueryString   = MatchCondition_MatchVariable("QueryString")
	MatchCondition_MatchVariable_RemoteAddr    = MatchCondition_MatchVariable("RemoteAddr")
	MatchCondition_MatchVariable_RequestBody   = MatchCondition_MatchVariable("RequestBody")
	MatchCondition_MatchVariable_RequestHeader = MatchCondition_MatchVariable("RequestHeader")
	MatchCondition_MatchVariable_RequestMethod = MatchCondition_MatchVariable("RequestMethod")
	MatchCondition_MatchVariable_RequestUri    = MatchCondition_MatchVariable("RequestUri")
	MatchCondition_MatchVariable_SocketAddr    = MatchCondition_MatchVariable("SocketAddr")
)

// Mapping from string to MatchCondition_MatchVariable
var matchCondition_MatchVariable_Values = map[string]MatchCondition_MatchVariable{
	"cookies":       MatchCondition_MatchVariable_Cookies,
	"postargs":      MatchCondition_MatchVariable_PostArgs,
	"querystring":   MatchCondition_MatchVariable_QueryString,
	"remoteaddr":    MatchCondition_MatchVariable_RemoteAddr,
	"requestbody":   MatchCondition_MatchVariable_RequestBody,
	"requestheader": MatchCondition_MatchVariable_RequestHeader,
	"requestmethod": MatchCondition_MatchVariable_RequestMethod,
	"requesturi":    MatchCondition_MatchVariable_RequestUri,
	"socketaddr":    MatchCondition_MatchVariable_SocketAddr,
}

// +kubebuilder:validation:Enum={"Any","BeginsWith","Contains","EndsWith","Equal","GeoMatch","GreaterThan","GreaterThanOrEqual","IPMatch","LessThan","LessThanOrEqual","RegEx"}
type MatchCondition_Operator string

const (
	MatchCondition_Operator_Any                = MatchCondition_Operator("Any")
	MatchCondition_Operator_BeginsWith         = MatchCondition_Operator("BeginsWith")
	MatchCondition_Operator_Contains           = MatchCondition_Operator("Contains")
	MatchCondition_Operator_EndsWith           = MatchCondition_Operator("EndsWith")
	MatchCondition_Operator_Equal              = MatchCondition_Operator("Equal")
	MatchCondition_Operator_GeoMatch           = MatchCondition_Operator("GeoMatch")
	MatchCondition_Operator_GreaterThan        = MatchCondition_Operator("GreaterThan")
	MatchCondition_Operator_GreaterThanOrEqual = MatchCondition_Operator("GreaterThanOrEqual")
	MatchCondition_Operator_IPMatch            = MatchCondition_Operator("IPMatch")
	MatchCondition_Operator_LessThan           = MatchCondition_Operator("LessThan")
	MatchCondition_Operator_LessThanOrEqual    = MatchCondition_Operator("LessThanOrEqual")
	MatchCondition_Operator_RegEx              = MatchCondition_Operator("RegEx")
)

// Mapping from string to MatchCondition_Operator
var matchCondition_Operator_Values = map[string]MatchCondition_Operator{
	"any":                MatchCondition_Operator_Any,
	"beginswith":         MatchCondition_Operator_BeginsWith,
	"contains":           MatchCondition_Operator_Contains,
	"endswith":           MatchCondition_Operator_EndsWith,
	"equal":              MatchCondition_Operator_Equal,
	"geomatch":           MatchCondition_Operator_GeoMatch,
	"greaterthan":        MatchCondition_Operator_GreaterThan,
	"greaterthanorequal": MatchCondition_Operator_GreaterThanOrEqual,
	"ipmatch":            MatchCondition_Operator_IPMatch,
	"lessthan":           MatchCondition_Operator_LessThan,
	"lessthanorequal":    MatchCondition_Operator_LessThanOrEqual,
	"regex":              MatchCondition_Operator_RegEx,
}

// Describes what transforms applied before matching.
// +kubebuilder:validation:Enum={"Lowercase","RemoveNulls","Trim","Uppercase","UrlDecode","UrlEncode"}
type TransformType string

const (
	TransformType_Lowercase   = TransformType("Lowercase")
	TransformType_RemoveNulls = TransformType("RemoveNulls")
	TransformType_Trim        = TransformType("Trim")
	TransformType_Uppercase   = TransformType("Uppercase")
	TransformType_UrlDecode   = TransformType("UrlDecode")
	TransformType_UrlEncode   = TransformType("UrlEncode")
)

// Mapping from string to TransformType
var transformType_Values = map[string]TransformType{
	"lowercase":   TransformType_Lowercase,
	"removenulls": TransformType_RemoveNulls,
	"trim":        TransformType_Trim,
	"uppercase":   TransformType_Uppercase,
	"urldecode":   TransformType_UrlDecode,
	"urlencode":   TransformType_UrlEncode,
}

// Describes if the managed rule is in enabled or disabled state.
// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ManagedRuleEnabledState string

const (
	ManagedRuleEnabledState_Disabled = ManagedRuleEnabledState("Disabled")
	ManagedRuleEnabledState_Enabled  = ManagedRuleEnabledState("Enabled")
)

// Mapping from string to ManagedRuleEnabledState
var managedRuleEnabledState_Values = map[string]ManagedRuleEnabledState{
	"disabled": ManagedRuleEnabledState_Disabled,
	"enabled":  ManagedRuleEnabledState_Enabled,
}
