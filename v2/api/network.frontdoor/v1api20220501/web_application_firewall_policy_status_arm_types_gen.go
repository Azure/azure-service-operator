// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220501

type WebApplicationFirewallPolicy_STATUS_ARM struct {
	// Etag: Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the web application firewall policy.
	Properties *WebApplicationFirewallPolicyProperties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: The pricing tier of web application firewall policy. Defaults to Classic_AzureFrontDoor if not specified.
	Sku *Sku_STATUS_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The pricing tier of the web application firewall policy.
type Sku_STATUS_ARM struct {
	// Name: Name of the pricing tier.
	Name *Sku_Name_STATUS_ARM `json:"name,omitempty"`
}

// Defines web application firewall policy properties.
type WebApplicationFirewallPolicyProperties_STATUS_ARM struct {
	// CustomRules: Describes custom rules inside the policy.
	CustomRules *CustomRuleList_STATUS_ARM `json:"customRules,omitempty"`

	// FrontendEndpointLinks: Describes Frontend Endpoints associated with this Web Application Firewall policy.
	FrontendEndpointLinks []FrontendEndpointLink_STATUS_ARM `json:"frontendEndpointLinks,omitempty"`

	// ManagedRules: Describes managed rules inside the policy.
	ManagedRules *ManagedRuleSetList_STATUS_ARM `json:"managedRules,omitempty"`

	// PolicySettings: Describes settings for the policy.
	PolicySettings *PolicySettings_STATUS_ARM `json:"policySettings,omitempty"`

	// ProvisioningState: Provisioning state of the policy.
	ProvisioningState *string                                                          `json:"provisioningState,omitempty"`
	ResourceState     *WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM `json:"resourceState,omitempty"`

	// RoutingRuleLinks: Describes Routing Rules associated with this Web Application Firewall policy.
	RoutingRuleLinks []RoutingRuleLink_STATUS_ARM `json:"routingRuleLinks,omitempty"`

	// SecurityPolicyLinks: Describes Security Policy associated with this Web Application Firewall policy.
	SecurityPolicyLinks []SecurityPolicyLink_STATUS_ARM `json:"securityPolicyLinks,omitempty"`
}

// Defines contents of custom rules
type CustomRuleList_STATUS_ARM struct {
	// Rules: List of rules
	Rules []CustomRule_STATUS_ARM `json:"rules,omitempty"`
}

// Defines the Resource ID for a Frontend Endpoint.
type FrontendEndpointLink_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Defines the list of managed rule sets for the policy.
type ManagedRuleSetList_STATUS_ARM struct {
	// ManagedRuleSets: List of rule sets.
	ManagedRuleSets []ManagedRuleSet_STATUS_ARM `json:"managedRuleSets,omitempty"`
}

// Defines top-level WebApplicationFirewallPolicy configuration settings.
type PolicySettings_STATUS_ARM struct {
	// CustomBlockResponseBody: If the action type is block, customer can override the response body. The body must be
	// specified in base64 encoding.
	CustomBlockResponseBody *string `json:"customBlockResponseBody,omitempty"`

	// CustomBlockResponseStatusCode: If the action type is block, customer can override the response status code.
	CustomBlockResponseStatusCode *int `json:"customBlockResponseStatusCode,omitempty"`

	// EnabledState: Describes if the policy is in enabled or disabled state. Defaults to Enabled if not specified.
	EnabledState *PolicySettings_EnabledState_STATUS_ARM `json:"enabledState,omitempty"`

	// Mode: Describes if it is in detection mode or prevention mode at policy level.
	Mode *PolicySettings_Mode_STATUS_ARM `json:"mode,omitempty"`

	// RedirectUrl: If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl *string `json:"redirectUrl,omitempty"`

	// RequestBodyCheck: Describes if policy managed rules will inspect the request body content.
	RequestBodyCheck *PolicySettings_RequestBodyCheck_STATUS_ARM `json:"requestBodyCheck,omitempty"`
}

// Defines the Resource ID for a Routing Rule.
type RoutingRuleLink_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Defines the Resource ID for a Security Policy.
type SecurityPolicyLink_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type Sku_Name_STATUS_ARM string

const (
	Sku_Name_STATUS_ARM_Classic_AzureFrontDoor  = Sku_Name_STATUS_ARM("Classic_AzureFrontDoor")
	Sku_Name_STATUS_ARM_Premium_AzureFrontDoor  = Sku_Name_STATUS_ARM("Premium_AzureFrontDoor")
	Sku_Name_STATUS_ARM_Standard_AzureFrontDoor = Sku_Name_STATUS_ARM("Standard_AzureFrontDoor")
)

// Mapping from string to Sku_Name_STATUS_ARM
var sku_Name_STATUS_ARM_Values = map[string]Sku_Name_STATUS_ARM{
	"classic_azurefrontdoor":  Sku_Name_STATUS_ARM_Classic_AzureFrontDoor,
	"premium_azurefrontdoor":  Sku_Name_STATUS_ARM_Premium_AzureFrontDoor,
	"standard_azurefrontdoor": Sku_Name_STATUS_ARM_Standard_AzureFrontDoor,
}

type WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM string

const (
	WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Creating  = WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM("Creating")
	WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Deleting  = WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM("Deleting")
	WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Disabled  = WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM("Disabled")
	WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Disabling = WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM("Disabling")
	WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Enabled   = WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM("Enabled")
	WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Enabling  = WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM("Enabling")
)

// Mapping from string to WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM
var webApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Values = map[string]WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM{
	"creating":  WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Creating,
	"deleting":  WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Deleting,
	"disabled":  WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Disabled,
	"disabling": WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Disabling,
	"enabled":   WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Enabled,
	"enabling":  WebApplicationFirewallPolicyProperties_ResourceState_STATUS_ARM_Enabling,
}

// Defines contents of a web application rule
type CustomRule_STATUS_ARM struct {
	// Action: Describes what action to be applied when rule matches.
	Action *ActionType_STATUS_ARM `json:"action,omitempty"`

	// EnabledState: Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
	EnabledState *CustomRule_EnabledState_STATUS_ARM `json:"enabledState,omitempty"`

	// MatchConditions: List of match conditions.
	MatchConditions []MatchCondition_STATUS_ARM `json:"matchConditions,omitempty"`

	// Name: Describes the name of the rule.
	Name *string `json:"name,omitempty"`

	// Priority: Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
	Priority *int `json:"priority,omitempty"`

	// RateLimitDurationInMinutes: Time window for resetting the rate limit count. Default is 1 minute.
	RateLimitDurationInMinutes *int `json:"rateLimitDurationInMinutes,omitempty"`

	// RateLimitThreshold: Number of allowed requests per client within the time window.
	RateLimitThreshold *int `json:"rateLimitThreshold,omitempty"`

	// RuleType: Describes type of rule.
	RuleType *CustomRule_RuleType_STATUS_ARM `json:"ruleType,omitempty"`
}

// Defines a managed rule set.
type ManagedRuleSet_STATUS_ARM struct {
	// Exclusions: Describes the exclusions that are applied to all rules in the set.
	Exclusions []ManagedRuleExclusion_STATUS_ARM `json:"exclusions,omitempty"`

	// RuleGroupOverrides: Defines the rule group overrides to apply to the rule set.
	RuleGroupOverrides []ManagedRuleGroupOverride_STATUS_ARM `json:"ruleGroupOverrides,omitempty"`

	// RuleSetAction: Defines the rule set action.
	RuleSetAction *ManagedRuleSetActionType_STATUS_ARM `json:"ruleSetAction,omitempty"`

	// RuleSetType: Defines the rule set type to use.
	RuleSetType *string `json:"ruleSetType,omitempty"`

	// RuleSetVersion: Defines the version of the rule set to use.
	RuleSetVersion *string `json:"ruleSetVersion,omitempty"`
}

type PolicySettings_EnabledState_STATUS_ARM string

const (
	PolicySettings_EnabledState_STATUS_ARM_Disabled = PolicySettings_EnabledState_STATUS_ARM("Disabled")
	PolicySettings_EnabledState_STATUS_ARM_Enabled  = PolicySettings_EnabledState_STATUS_ARM("Enabled")
)

// Mapping from string to PolicySettings_EnabledState_STATUS_ARM
var policySettings_EnabledState_STATUS_ARM_Values = map[string]PolicySettings_EnabledState_STATUS_ARM{
	"disabled": PolicySettings_EnabledState_STATUS_ARM_Disabled,
	"enabled":  PolicySettings_EnabledState_STATUS_ARM_Enabled,
}

type PolicySettings_Mode_STATUS_ARM string

const (
	PolicySettings_Mode_STATUS_ARM_Detection  = PolicySettings_Mode_STATUS_ARM("Detection")
	PolicySettings_Mode_STATUS_ARM_Prevention = PolicySettings_Mode_STATUS_ARM("Prevention")
)

// Mapping from string to PolicySettings_Mode_STATUS_ARM
var policySettings_Mode_STATUS_ARM_Values = map[string]PolicySettings_Mode_STATUS_ARM{
	"detection":  PolicySettings_Mode_STATUS_ARM_Detection,
	"prevention": PolicySettings_Mode_STATUS_ARM_Prevention,
}

type PolicySettings_RequestBodyCheck_STATUS_ARM string

const (
	PolicySettings_RequestBodyCheck_STATUS_ARM_Disabled = PolicySettings_RequestBodyCheck_STATUS_ARM("Disabled")
	PolicySettings_RequestBodyCheck_STATUS_ARM_Enabled  = PolicySettings_RequestBodyCheck_STATUS_ARM("Enabled")
)

// Mapping from string to PolicySettings_RequestBodyCheck_STATUS_ARM
var policySettings_RequestBodyCheck_STATUS_ARM_Values = map[string]PolicySettings_RequestBodyCheck_STATUS_ARM{
	"disabled": PolicySettings_RequestBodyCheck_STATUS_ARM_Disabled,
	"enabled":  PolicySettings_RequestBodyCheck_STATUS_ARM_Enabled,
}

// Defines the action to take on rule match.
type ActionType_STATUS_ARM string

const (
	ActionType_STATUS_ARM_Allow          = ActionType_STATUS_ARM("Allow")
	ActionType_STATUS_ARM_AnomalyScoring = ActionType_STATUS_ARM("AnomalyScoring")
	ActionType_STATUS_ARM_Block          = ActionType_STATUS_ARM("Block")
	ActionType_STATUS_ARM_Log            = ActionType_STATUS_ARM("Log")
	ActionType_STATUS_ARM_Redirect       = ActionType_STATUS_ARM("Redirect")
)

// Mapping from string to ActionType_STATUS_ARM
var actionType_STATUS_ARM_Values = map[string]ActionType_STATUS_ARM{
	"allow":          ActionType_STATUS_ARM_Allow,
	"anomalyscoring": ActionType_STATUS_ARM_AnomalyScoring,
	"block":          ActionType_STATUS_ARM_Block,
	"log":            ActionType_STATUS_ARM_Log,
	"redirect":       ActionType_STATUS_ARM_Redirect,
}

type CustomRule_EnabledState_STATUS_ARM string

const (
	CustomRule_EnabledState_STATUS_ARM_Disabled = CustomRule_EnabledState_STATUS_ARM("Disabled")
	CustomRule_EnabledState_STATUS_ARM_Enabled  = CustomRule_EnabledState_STATUS_ARM("Enabled")
)

// Mapping from string to CustomRule_EnabledState_STATUS_ARM
var customRule_EnabledState_STATUS_ARM_Values = map[string]CustomRule_EnabledState_STATUS_ARM{
	"disabled": CustomRule_EnabledState_STATUS_ARM_Disabled,
	"enabled":  CustomRule_EnabledState_STATUS_ARM_Enabled,
}

type CustomRule_RuleType_STATUS_ARM string

const (
	CustomRule_RuleType_STATUS_ARM_MatchRule     = CustomRule_RuleType_STATUS_ARM("MatchRule")
	CustomRule_RuleType_STATUS_ARM_RateLimitRule = CustomRule_RuleType_STATUS_ARM("RateLimitRule")
)

// Mapping from string to CustomRule_RuleType_STATUS_ARM
var customRule_RuleType_STATUS_ARM_Values = map[string]CustomRule_RuleType_STATUS_ARM{
	"matchrule":     CustomRule_RuleType_STATUS_ARM_MatchRule,
	"ratelimitrule": CustomRule_RuleType_STATUS_ARM_RateLimitRule,
}

// Exclude variables from managed rule evaluation.
type ManagedRuleExclusion_STATUS_ARM struct {
	// MatchVariable: The variable type to be excluded.
	MatchVariable *ManagedRuleExclusion_MatchVariable_STATUS_ARM `json:"matchVariable,omitempty"`

	// Selector: Selector value for which elements in the collection this exclusion applies to.
	Selector *string `json:"selector,omitempty"`

	// SelectorMatchOperator: Comparison operator to apply to the selector when specifying which elements in the collection
	// this exclusion applies to.
	SelectorMatchOperator *ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM `json:"selectorMatchOperator,omitempty"`
}

// Defines a managed rule group override setting.
type ManagedRuleGroupOverride_STATUS_ARM struct {
	// Exclusions: Describes the exclusions that are applied to all rules in the group.
	Exclusions []ManagedRuleExclusion_STATUS_ARM `json:"exclusions,omitempty"`

	// RuleGroupName: Describes the managed rule group to override.
	RuleGroupName *string `json:"ruleGroupName,omitempty"`

	// Rules: List of rules that will be disabled. If none specified, all rules in the group will be disabled.
	Rules []ManagedRuleOverride_STATUS_ARM `json:"rules,omitempty"`
}

// Defines the action to take when a managed rule set score threshold is met.
type ManagedRuleSetActionType_STATUS_ARM string

const (
	ManagedRuleSetActionType_STATUS_ARM_Block    = ManagedRuleSetActionType_STATUS_ARM("Block")
	ManagedRuleSetActionType_STATUS_ARM_Log      = ManagedRuleSetActionType_STATUS_ARM("Log")
	ManagedRuleSetActionType_STATUS_ARM_Redirect = ManagedRuleSetActionType_STATUS_ARM("Redirect")
)

// Mapping from string to ManagedRuleSetActionType_STATUS_ARM
var managedRuleSetActionType_STATUS_ARM_Values = map[string]ManagedRuleSetActionType_STATUS_ARM{
	"block":    ManagedRuleSetActionType_STATUS_ARM_Block,
	"log":      ManagedRuleSetActionType_STATUS_ARM_Log,
	"redirect": ManagedRuleSetActionType_STATUS_ARM_Redirect,
}

// Define a match condition.
type MatchCondition_STATUS_ARM struct {
	// MatchValue: List of possible match values.
	MatchValue []string `json:"matchValue,omitempty"`

	// MatchVariable: Request variable to compare with.
	MatchVariable *MatchCondition_MatchVariable_STATUS_ARM `json:"matchVariable,omitempty"`

	// NegateCondition: Describes if the result of this condition should be negated.
	NegateCondition *bool `json:"negateCondition,omitempty"`

	// Operator: Comparison type to use for matching with the variable value.
	Operator *MatchCondition_Operator_STATUS_ARM `json:"operator,omitempty"`

	// Selector: Match against a specific key from the QueryString, PostArgs, RequestHeader or Cookies variables. Default is
	// null.
	Selector *string `json:"selector,omitempty"`

	// Transforms: List of transforms.
	Transforms []TransformType_STATUS_ARM `json:"transforms,omitempty"`
}

type ManagedRuleExclusion_MatchVariable_STATUS_ARM string

const (
	ManagedRuleExclusion_MatchVariable_STATUS_ARM_QueryStringArgNames     = ManagedRuleExclusion_MatchVariable_STATUS_ARM("QueryStringArgNames")
	ManagedRuleExclusion_MatchVariable_STATUS_ARM_RequestBodyJsonArgNames = ManagedRuleExclusion_MatchVariable_STATUS_ARM("RequestBodyJsonArgNames")
	ManagedRuleExclusion_MatchVariable_STATUS_ARM_RequestBodyPostArgNames = ManagedRuleExclusion_MatchVariable_STATUS_ARM("RequestBodyPostArgNames")
	ManagedRuleExclusion_MatchVariable_STATUS_ARM_RequestCookieNames      = ManagedRuleExclusion_MatchVariable_STATUS_ARM("RequestCookieNames")
	ManagedRuleExclusion_MatchVariable_STATUS_ARM_RequestHeaderNames      = ManagedRuleExclusion_MatchVariable_STATUS_ARM("RequestHeaderNames")
)

// Mapping from string to ManagedRuleExclusion_MatchVariable_STATUS_ARM
var managedRuleExclusion_MatchVariable_STATUS_ARM_Values = map[string]ManagedRuleExclusion_MatchVariable_STATUS_ARM{
	"querystringargnames":     ManagedRuleExclusion_MatchVariable_STATUS_ARM_QueryStringArgNames,
	"requestbodyjsonargnames": ManagedRuleExclusion_MatchVariable_STATUS_ARM_RequestBodyJsonArgNames,
	"requestbodypostargnames": ManagedRuleExclusion_MatchVariable_STATUS_ARM_RequestBodyPostArgNames,
	"requestcookienames":      ManagedRuleExclusion_MatchVariable_STATUS_ARM_RequestCookieNames,
	"requestheadernames":      ManagedRuleExclusion_MatchVariable_STATUS_ARM_RequestHeaderNames,
}

type ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM string

const (
	ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM_Contains   = ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM("Contains")
	ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM_EndsWith   = ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM("EndsWith")
	ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM_Equals     = ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM("Equals")
	ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM_EqualsAny  = ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM("EqualsAny")
	ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM_StartsWith = ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM("StartsWith")
)

// Mapping from string to ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM
var managedRuleExclusion_SelectorMatchOperator_STATUS_ARM_Values = map[string]ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM{
	"contains":   ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM_Contains,
	"endswith":   ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM_EndsWith,
	"equals":     ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM_Equals,
	"equalsany":  ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM_EqualsAny,
	"startswith": ManagedRuleExclusion_SelectorMatchOperator_STATUS_ARM_StartsWith,
}

// Defines a managed rule group override setting.
type ManagedRuleOverride_STATUS_ARM struct {
	// Action: Describes the override action to be applied when rule matches.
	Action *ActionType_STATUS_ARM `json:"action,omitempty"`

	// EnabledState: Describes if the managed rule is in enabled or disabled state. Defaults to Disabled if not specified.
	EnabledState *ManagedRuleEnabledState_STATUS_ARM `json:"enabledState,omitempty"`

	// Exclusions: Describes the exclusions that are applied to this specific rule.
	Exclusions []ManagedRuleExclusion_STATUS_ARM `json:"exclusions,omitempty"`

	// RuleId: Identifier for the managed rule.
	RuleId *string `json:"ruleId,omitempty"`
}

type MatchCondition_MatchVariable_STATUS_ARM string

const (
	MatchCondition_MatchVariable_STATUS_ARM_Cookies       = MatchCondition_MatchVariable_STATUS_ARM("Cookies")
	MatchCondition_MatchVariable_STATUS_ARM_PostArgs      = MatchCondition_MatchVariable_STATUS_ARM("PostArgs")
	MatchCondition_MatchVariable_STATUS_ARM_QueryString   = MatchCondition_MatchVariable_STATUS_ARM("QueryString")
	MatchCondition_MatchVariable_STATUS_ARM_RemoteAddr    = MatchCondition_MatchVariable_STATUS_ARM("RemoteAddr")
	MatchCondition_MatchVariable_STATUS_ARM_RequestBody   = MatchCondition_MatchVariable_STATUS_ARM("RequestBody")
	MatchCondition_MatchVariable_STATUS_ARM_RequestHeader = MatchCondition_MatchVariable_STATUS_ARM("RequestHeader")
	MatchCondition_MatchVariable_STATUS_ARM_RequestMethod = MatchCondition_MatchVariable_STATUS_ARM("RequestMethod")
	MatchCondition_MatchVariable_STATUS_ARM_RequestUri    = MatchCondition_MatchVariable_STATUS_ARM("RequestUri")
	MatchCondition_MatchVariable_STATUS_ARM_SocketAddr    = MatchCondition_MatchVariable_STATUS_ARM("SocketAddr")
)

// Mapping from string to MatchCondition_MatchVariable_STATUS_ARM
var matchCondition_MatchVariable_STATUS_ARM_Values = map[string]MatchCondition_MatchVariable_STATUS_ARM{
	"cookies":       MatchCondition_MatchVariable_STATUS_ARM_Cookies,
	"postargs":      MatchCondition_MatchVariable_STATUS_ARM_PostArgs,
	"querystring":   MatchCondition_MatchVariable_STATUS_ARM_QueryString,
	"remoteaddr":    MatchCondition_MatchVariable_STATUS_ARM_RemoteAddr,
	"requestbody":   MatchCondition_MatchVariable_STATUS_ARM_RequestBody,
	"requestheader": MatchCondition_MatchVariable_STATUS_ARM_RequestHeader,
	"requestmethod": MatchCondition_MatchVariable_STATUS_ARM_RequestMethod,
	"requesturi":    MatchCondition_MatchVariable_STATUS_ARM_RequestUri,
	"socketaddr":    MatchCondition_MatchVariable_STATUS_ARM_SocketAddr,
}

type MatchCondition_Operator_STATUS_ARM string

const (
	MatchCondition_Operator_STATUS_ARM_Any                = MatchCondition_Operator_STATUS_ARM("Any")
	MatchCondition_Operator_STATUS_ARM_BeginsWith         = MatchCondition_Operator_STATUS_ARM("BeginsWith")
	MatchCondition_Operator_STATUS_ARM_Contains           = MatchCondition_Operator_STATUS_ARM("Contains")
	MatchCondition_Operator_STATUS_ARM_EndsWith           = MatchCondition_Operator_STATUS_ARM("EndsWith")
	MatchCondition_Operator_STATUS_ARM_Equal              = MatchCondition_Operator_STATUS_ARM("Equal")
	MatchCondition_Operator_STATUS_ARM_GeoMatch           = MatchCondition_Operator_STATUS_ARM("GeoMatch")
	MatchCondition_Operator_STATUS_ARM_GreaterThan        = MatchCondition_Operator_STATUS_ARM("GreaterThan")
	MatchCondition_Operator_STATUS_ARM_GreaterThanOrEqual = MatchCondition_Operator_STATUS_ARM("GreaterThanOrEqual")
	MatchCondition_Operator_STATUS_ARM_IPMatch            = MatchCondition_Operator_STATUS_ARM("IPMatch")
	MatchCondition_Operator_STATUS_ARM_LessThan           = MatchCondition_Operator_STATUS_ARM("LessThan")
	MatchCondition_Operator_STATUS_ARM_LessThanOrEqual    = MatchCondition_Operator_STATUS_ARM("LessThanOrEqual")
	MatchCondition_Operator_STATUS_ARM_RegEx              = MatchCondition_Operator_STATUS_ARM("RegEx")
)

// Mapping from string to MatchCondition_Operator_STATUS_ARM
var matchCondition_Operator_STATUS_ARM_Values = map[string]MatchCondition_Operator_STATUS_ARM{
	"any":                MatchCondition_Operator_STATUS_ARM_Any,
	"beginswith":         MatchCondition_Operator_STATUS_ARM_BeginsWith,
	"contains":           MatchCondition_Operator_STATUS_ARM_Contains,
	"endswith":           MatchCondition_Operator_STATUS_ARM_EndsWith,
	"equal":              MatchCondition_Operator_STATUS_ARM_Equal,
	"geomatch":           MatchCondition_Operator_STATUS_ARM_GeoMatch,
	"greaterthan":        MatchCondition_Operator_STATUS_ARM_GreaterThan,
	"greaterthanorequal": MatchCondition_Operator_STATUS_ARM_GreaterThanOrEqual,
	"ipmatch":            MatchCondition_Operator_STATUS_ARM_IPMatch,
	"lessthan":           MatchCondition_Operator_STATUS_ARM_LessThan,
	"lessthanorequal":    MatchCondition_Operator_STATUS_ARM_LessThanOrEqual,
	"regex":              MatchCondition_Operator_STATUS_ARM_RegEx,
}

// Describes what transforms applied before matching.
type TransformType_STATUS_ARM string

const (
	TransformType_STATUS_ARM_Lowercase   = TransformType_STATUS_ARM("Lowercase")
	TransformType_STATUS_ARM_RemoveNulls = TransformType_STATUS_ARM("RemoveNulls")
	TransformType_STATUS_ARM_Trim        = TransformType_STATUS_ARM("Trim")
	TransformType_STATUS_ARM_Uppercase   = TransformType_STATUS_ARM("Uppercase")
	TransformType_STATUS_ARM_UrlDecode   = TransformType_STATUS_ARM("UrlDecode")
	TransformType_STATUS_ARM_UrlEncode   = TransformType_STATUS_ARM("UrlEncode")
)

// Mapping from string to TransformType_STATUS_ARM
var transformType_STATUS_ARM_Values = map[string]TransformType_STATUS_ARM{
	"lowercase":   TransformType_STATUS_ARM_Lowercase,
	"removenulls": TransformType_STATUS_ARM_RemoveNulls,
	"trim":        TransformType_STATUS_ARM_Trim,
	"uppercase":   TransformType_STATUS_ARM_Uppercase,
	"urldecode":   TransformType_STATUS_ARM_UrlDecode,
	"urlencode":   TransformType_STATUS_ARM_UrlEncode,
}

// Describes if the managed rule is in enabled or disabled state.
type ManagedRuleEnabledState_STATUS_ARM string

const (
	ManagedRuleEnabledState_STATUS_ARM_Disabled = ManagedRuleEnabledState_STATUS_ARM("Disabled")
	ManagedRuleEnabledState_STATUS_ARM_Enabled  = ManagedRuleEnabledState_STATUS_ARM("Enabled")
)

// Mapping from string to ManagedRuleEnabledState_STATUS_ARM
var managedRuleEnabledState_STATUS_ARM_Values = map[string]ManagedRuleEnabledState_STATUS_ARM{
	"disabled": ManagedRuleEnabledState_STATUS_ARM_Disabled,
	"enabled":  ManagedRuleEnabledState_STATUS_ARM_Enabled,
}
