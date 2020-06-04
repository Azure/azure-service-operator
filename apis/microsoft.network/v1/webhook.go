/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var backendaddresspoollog = logf.Log.WithName("backendaddresspool-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-backendaddresspool,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=backendaddresspools,verbs=create;update,versions=v1,name=default.backendaddresspool.infra.azure.com

var _ webhook.Defaulter = &BackendAddressPool{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *BackendAddressPool) Default() {
	backendaddresspoollog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-backendaddresspool,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=backendaddresspools,versions=v1,name=validation.backendaddresspool.infra.azure.com

var _ webhook.Validator = &BackendAddressPool{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *BackendAddressPool) ValidateCreate() error {
	backendaddresspoollog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *BackendAddressPool) ValidateUpdate(old runtime.Object) error {
	backendaddresspoollog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *BackendAddressPool) ValidateDelete() error {
	backendaddresspoollog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// log is for logging in this package.
var frontendipconfigurationlog = logf.Log.WithName("frontendipconfiguration-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-frontendipconfiguration,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=frontendipconfigurations,verbs=create;update,versions=v1,name=default.frontendipconfiguration.infra.azure.com

var _ webhook.Defaulter = &FrontendIPConfiguration{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *FrontendIPConfiguration) Default() {
	frontendipconfigurationlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-frontendipconfiguration,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=frontendipconfigurations,versions=v1,name=validation.frontendipconfiguration.infra.azure.com

var _ webhook.Validator = &FrontendIPConfiguration{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *FrontendIPConfiguration) ValidateCreate() error {
	frontendipconfigurationlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *FrontendIPConfiguration) ValidateUpdate(old runtime.Object) error {
	frontendipconfigurationlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *FrontendIPConfiguration) ValidateDelete() error {
	frontendipconfigurationlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// log is for logging in this package.
var inboundnatrulelog = logf.Log.WithName("inboundnatrule-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-inboundnatrule,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=inboundnatrules,verbs=create;update,versions=v1,name=default.inboundnatrule.infra.azure.com

var _ webhook.Defaulter = &InboundNatRule{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *InboundNatRule) Default() {
	inboundnatrulelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-inboundnatrule,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=inboundnatrules,versions=v1,name=validation.inboundnatrule.infra.azure.com

var _ webhook.Validator = &InboundNatRule{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *InboundNatRule) ValidateCreate() error {
	inboundnatrulelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *InboundNatRule) ValidateUpdate(old runtime.Object) error {
	inboundnatrulelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *InboundNatRule) ValidateDelete() error {
	inboundnatrulelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// log is for logging in this package.
var loadbalancerlog = logf.Log.WithName("loadbalancer-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-loadbalancer,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=loadbalancers,verbs=create;update,versions=v1,name=default.loadbalancer.infra.azure.com

var _ webhook.Defaulter = &LoadBalancer{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *LoadBalancer) Default() {
	loadbalancerlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-loadbalancer,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=loadbalancers,versions=v1,name=validation.loadbalancer.infra.azure.com

var _ webhook.Validator = &LoadBalancer{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *LoadBalancer) ValidateCreate() error {
	loadbalancerlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *LoadBalancer) ValidateUpdate(old runtime.Object) error {
	loadbalancerlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *LoadBalancer) ValidateDelete() error {
	loadbalancerlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// log is for logging in this package.
var loadbalancingrulelog = logf.Log.WithName("loadbalancingrule-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-loadbalancingrule,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=loadbalancingrules,verbs=create;update,versions=v1,name=default.loadbalancingrule.infra.azure.com

var _ webhook.Defaulter = &LoadBalancingRule{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *LoadBalancingRule) Default() {
	loadbalancingrulelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-loadbalancingrule,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=loadbalancingrules,versions=v1,name=validation.loadbalancingrule.infra.azure.com

var _ webhook.Validator = &LoadBalancingRule{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *LoadBalancingRule) ValidateCreate() error {
	loadbalancingrulelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *LoadBalancingRule) ValidateUpdate(old runtime.Object) error {
	loadbalancingrulelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *LoadBalancingRule) ValidateDelete() error {
	loadbalancingrulelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// log is for logging in this package.
var networksecuritygrouplog = logf.Log.WithName("networksecuritygroup-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-networksecuritygroup,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=networksecuritygroups,verbs=create;update,versions=v1,name=default.networksecuritygroup.infra.azure.com

var _ webhook.Defaulter = &NetworkSecurityGroup{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *NetworkSecurityGroup) Default() {
	networksecuritygrouplog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-networksecuritygroup,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=networksecuritygroups,versions=v1,name=validation.networksecuritygroup.infra.azure.com

var _ webhook.Validator = &NetworkSecurityGroup{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *NetworkSecurityGroup) ValidateCreate() error {
	networksecuritygrouplog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *NetworkSecurityGroup) ValidateUpdate(old runtime.Object) error {
	networksecuritygrouplog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *NetworkSecurityGroup) ValidateDelete() error {
	networksecuritygrouplog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// log is for logging in this package.
var outboundrulelog = logf.Log.WithName("outboundrule-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-outboundrule,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=outboundrules,verbs=create;update,versions=v1,name=default.outboundrule.infra.azure.com

var _ webhook.Defaulter = &OutboundRule{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *OutboundRule) Default() {
	outboundrulelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-outboundrule,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=outboundrules,versions=v1,name=validation.outboundrule.infra.azure.com

var _ webhook.Validator = &OutboundRule{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *OutboundRule) ValidateCreate() error {
	outboundrulelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *OutboundRule) ValidateUpdate(old runtime.Object) error {
	outboundrulelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *OutboundRule) ValidateDelete() error {
	outboundrulelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// log is for logging in this package.
var routelog = logf.Log.WithName("route-resource")

func (r *Route) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-route,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=routes,verbs=create;update,versions=v1,name=default.route.infra.azure.com

var _ webhook.Defaulter = &Route{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Route) Default() {
	routelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-route,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=routes,versions=v1,name=validation.route.infra.azure.com

var _ webhook.Validator = &Route{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Route) ValidateCreate() error {
	routelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Route) ValidateUpdate(old runtime.Object) error {
	routelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Route) ValidateDelete() error {
	routelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// log is for logging in this package.
var routetablelog = logf.Log.WithName("routetable-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-routetable,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=routetables,verbs=create;update,versions=v1,name=default.routetable.infra.azure.com

var _ webhook.Defaulter = &RouteTable{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *RouteTable) Default() {
	routetablelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-routetable,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=routetables,versions=v1,name=validation.routetable.infra.azure.com

var _ webhook.Validator = &RouteTable{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *RouteTable) ValidateCreate() error {
	routetablelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *RouteTable) ValidateUpdate(old runtime.Object) error {
	routetablelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *RouteTable) ValidateDelete() error {
	routetablelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// log is for logging in this package.
var securityrulelog = logf.Log.WithName("securityrule-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-securityrule,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=securityrules,verbs=create;update,versions=v1,name=default.securityrule.infra.azure.com

var _ webhook.Defaulter = &SecurityRule{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *SecurityRule) Default() {
	securityrulelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-securityrule,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=securityrules,versions=v1,name=validation.securityrule.infra.azure.com

var _ webhook.Validator = &SecurityRule{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *SecurityRule) ValidateCreate() error {
	securityrulelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *SecurityRule) ValidateUpdate(old runtime.Object) error {
	securityrulelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *SecurityRule) ValidateDelete() error {
	securityrulelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// log is for logging in this package.
var virtualnetworklog = logf.Log.WithName("virtualnetwork-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-virtualnetwork,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=virtualnetworks,verbs=create;update,versions=v1,name=default.virtualnetwork.infra.azure.com

var _ webhook.Defaulter = &VirtualNetwork{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *VirtualNetwork) Default() {
	virtualnetworklog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-virtualnetwork,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=virtualnetworks,versions=v1,name=validation.virtualnetwork.infra.azure.com

var _ webhook.Validator = &VirtualNetwork{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualNetwork) ValidateCreate() error {
	virtualnetworklog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualNetwork) ValidateUpdate(old runtime.Object) error {
	virtualnetworklog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualNetwork) ValidateDelete() error {
	virtualnetworklog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// log is for logging in this package.
var subnetlog = logf.Log.WithName("subnet-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-subnet,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=subnets,verbs=create;update,versions=v1,name=default.subnet.infra.azure.com

var _ webhook.Defaulter = &Subnet{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Subnet) Default() {
	subnetlog.Info("default", "name", r.Name)
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-subnet,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=subnets,versions=v1,name=validation.subnet.infra.azure.com

var _ webhook.Validator = &Subnet{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Subnet) ValidateCreate() error {
	subnetlog.Info("validate create", "name", r.Name)
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Subnet) ValidateUpdate(old runtime.Object) error {
	subnetlog.Info("validate update", "name", r.Name)
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Subnet) ValidateDelete() error {
	subnetlog.Info("validate delete", "name", r.Name)
	return nil
}
