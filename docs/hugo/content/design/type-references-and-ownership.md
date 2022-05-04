---
title: Type References & ownership
---
## Related reading
- [Kubernetes garbage collection](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/).
- [Kubernetes RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)

## Related Projects
- ASO: [Azure Service Operator](https://github.com/Azure/azure-service-operator)
- k8s-infra: The handcrafted precurser to the code generation tool being designed.

## Goals
- Provide a way for customers to express relationships between Azure resources in an idiomatic Kubernetes way.
- Provide automatic ownership and garbage collection (using [Kubernetes garbage collection](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/)) 
  where appropriate (e.g. ResourceGroup as an owner of all the resources inside of it)
  - Ideally ResourceGroup is handled the same as other owners and isn't special cased. 
- Define how Kubernetes interacts with Azure resources _not_ created/managed by Kubernetes, for example resources which were created prior to the 
  customer onboarding to the Azure Service Operator.
- References should be extensible to work across multiple Azure subscriptions, although
  initially we may not support that.
 
## Non-goals
- Managing ownership for resources/resource hierarchies that were not created
  by the service operator.
  While this proposal allows references to point to external resources not managed by the service operator,
  the operator is not watching/monitoring the resource in question
  and as such cannot propagate deletes. Put another way: for the operator to manage ownership/object lifecycles, the entire resource hierarchy must exist _within Kubernetes_. If only part of the resource hierarchy
  is managed by the service operator, only those parts can have their lifecycles managed.  

## Different kinds of resource relationships in Azure
#### Related/linked resources 
Two resources are related to one another ("has-a" relationship), but there is no ownership.
Example: [VMSS → Subnet](https://github.com/Azure/k8s-infra/blob/master/apis/microsoft.compute/v1/virtualmachinescaleset_types.go#L169) ([json schema](https://schema.management.azure.com/schemas/2019-07-01/Microsoft.Compute.json)).

This relationship is always one-way (a VMSS refers to a Subnet, but a Subnet _does not_ refer to a VMSS).

#### Owner and dependent
Two resources have a relationship where one is owned by the other.

Examples:
- a [RouteTable owns many Routes](https://github.com/Azure/k8s-infra/blob/master/apis/microsoft.network/v1/routetable_types.go#L18) ([json schema](https://schema.management.azure.com/schemas/2020-03-01/Microsoft.Network.json))
- a BatchAccount owns many Pools ([json schema](https://schema.management.azure.com/schemas/2017-09-01/Microsoft.Batch.json))
- a ResourceGroup owns any resource

A relationship like those shown here tells us two things:
- Where to create/manage the dependent resource (this `Route` goes in that particular `RouteTable`, this `RouteTable` has that `Route`)
- That the dependent resource should be deleted when the parent resource is deleted. There are theoretically two cases here:
    - The dependent resource must be deleted before the parent can be deleted.
    - Deletion of the parent automatically cascades to all dependent resources. Due to how Azure ARM
      resources express ownership (via `id` which is part of the URL, with dependent resources being a subdirectory under the
      owning resources URL) all ARM resources _should_ fall into this case.

Note that sometimes an owning resource has its dependent resources embedded directly
(for example: `RouteTable` has the property [RouteTablePropertiesFormat](https://schema.management.azure.com/schemas/2020-03-01/Microsoft.Network.json)).
Most types do not embed the dependent resource directly in the owning resource. We will need to cater for both the embedded and non-embedded cases.

## What do these relationships look like in existing solutions?

This section examines how other operator solutions have tackled these problems. We look at:
- ARM templates
- Azure Service Operator (ASO)
- k8s-infra

### Related/Linked resources

#### What does ARM do?
These are just properties (often but not always called `id`) which refer to the fully qualified ARM ID of another resource. For example see a
[sample deployment template for a VMSS refering to an existing vnet](https://github.com/Azure/azure-quickstart-templates/blob/master/201-vmss-existing-vnet/azuredeploy.json#L136).

```json
"properties": {
  "subnet": {
    "id": "[resourceId(parameters('existingVnetResourceGroupName'), 'Microsoft.Network/virtualNetworks/subnets', parameters('existingVnetName'), parameters('existingSubNetName'))]"
  },
}
```

#### What does ASO do?
Similar to how ARM templates behave, ASO uses the decomposition of fully qualified resource id to reference another resource, as seen [here for VMSS → VNet](https://github.com/Azure/azure-service-operator/blob/92240406aff3863f3a267d8a1dc1e28aa3e841ae/api/v1alpha1/azurevmscaleset_types.go#L25)

```go
type AzureVMScaleSetSpec struct {
    ...
	Location               string `json:"location"`
	ResourceGroup          string `json:"resourceGroup"`
	VirtualNetworkName     string `json:"virtualNetworkName"`
	SubnetName             string `json:"subnetName"`
    ...
}
```

These properties are combined into a fully qualified ARM ID like so:

```go
subnetIDInput := helpers.MakeResourceID(
    client.SubscriptionID,
    resourceGroupName,
    "Microsoft.Network",
    "virtualNetworks",
    vnetName,
    "subnets",
    subnetName,
)
```

This produces a resource ID: `/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}`.

Currently ASO does not support cross-subscription references (and some of the resources such as VMSS don't allow cross-resource group references), but it in theory could by adding parameters.

#### What does k8s-infra do?
k8s-infra is a bit different in that resource references are in Kubernetes style (namespace + name) and not Azure style (resource-group + resource-name).
All resource references are done using the special type `KnownTypeReference` which contains the fully qualified Kubernetes name for the resource.

### Dependent Resources
#### What does ARM do?
ARM template deployments support [two different ways of deploying dependent resources](https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/child-resource-name-type):

- Deploy the resources in the same ARM template with dependent resources embedded inside owning resources using the `resources` property.
    - The dependent resource need not specify which resource is its owner, because it is implied by embedded structure.
    - The owning resource and the dependent resource must be created at the same time.
- Deploy the resources separately.
    - The dependent resource must specify which resource is its owner by including in its `name` field both the owning resource name and the dependent
      resource name separated by a `/`. Each segment of the name corresponds to an owning resource. For example creating a Batch Pool `foo` in Batch account `account`
      would have `name` = `account/foo`.
    - The dependent resource can be created after the owning resource has already been created, or can be created at the same time as the owning resource. If
      created at the same time, the `dependsOn` field must be used to inform ARM of the order to perform resource creation.

#### What does ASO do?
Dependent resources in ASO have properties which map to the name/path to their owner. For example 
[MySQLFirewallRuleSpec](https://github.com/Azure/azure-service-operator/blob/92240406aff3863f3a267d8a1dc1e28aa3e841ae/api/v1alpha1/mysqlfirewallrule_types.go#L15) looks like this:

```go
type MySQLFirewallRuleSpec struct {
	ResourceGroup  string `json:"resourceGroup"`
	Server         string `json:"server"`
	StartIPAddress string `json:"startIpAddress"`
	EndIPAddress   string `json:"endIpAddress"`
}
```

The `ResourceGroup` and `Server` are references to the owners of this type.

#### What does k8s-infra do?
k8s-infra uses the same `KnownTypeReference` type mentioned above for ownership references too.
There are two patterns for ownership in k8s-infra today.

One pattern is used for ResourceGroup, where 
[top level resources have a link to the resource group they are in](https://github.com/Azure/k8s-infra/blob/14105b1cb3f6967cd086c9f8f75fb16bb85d6318/apis/microsoft.compute/v1/virtualmachinescaleset_types.go#L323).

```go
type VirtualMachineScaleSetSpec struct {
    // +kubebuilder:validation:Required
    ResourceGroupRef *azcorev1.KnownTypeReference `json:"resourceGroupRef" group:"microsoft.resources.infra.azure.com" kind:"ResourceGroup"`
    ...
}
```

The other pattern is where the owning resource has links to the dependent resources it is expecting to have:

```go
type RouteTableSpecProperties struct {
    DisableBGPRoutePropagation bool                          `json:"disableBgpRoutePropagation,omitempty"`
    RouteRefs                  []azcorev1.KnownTypeReference `json:"routeRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"Route" owned:"true"`
}
```

If the dependent resources aren't there, the status of the owning resource reflects an error explaining that.

## Proposal

### A note on names before we get started
In Kubernetes each resource must have a unique name for its group-kind. For example, if we had a `RouteTable` CRD, each `RouteTable` object would need to have a unique name.
In ARM, resources do not need to be uniquely named. There can be two `RouteTable` resources with the same name provided they are in different resource groups. The owner-dependent resource
relationship impacts uniqueness in Azure in a way that it doesn't in Kubernetes.

#### Proposed solution
All Kubernetes resources will have two fields which are used in combination to build the Azure name: `Metadata.Name` and `Spec.AzureName`.
When `Spec.AzureName` is empty, `Metadata.Name` is used as the resource name. When `Spec.AzureName` is provided, it takes precedence and is used when interacting with ARM, but
the resource in Kubernetes is still called by its  `Metadata.Name`.

### Overview

We propose the following high-level solution:
- All references will be via Kubernetes group, kind, and name.
- If a resource not managed by Kubernetes must be referenced, that resource must be imported into Kubernetes as an `Unmanaged` resource.
- Dependent resources will refer to their parent via an `Owner` property.
- The `Owner` property will automatically detect `group` and `kind`, making specifying an owner as simple as providing the Kubernetes resource name.
- Dependent resources with the `Owner` property set will automatically have their `ownerReferences` configured so that Kubernetes garbage collection will 
  delete the dependent resources when the owner is deleted.
- References to related resources will be automatically detected by the code generator and transformed into the correct reference type.
- At serialization time, the controller will transform the Kubernetes types (including related resource references and owner references) into the correct Azure
  resource definitions (including fully qualified ARM IDs).

More specific details about how this will be achieved are in the following sections.

### How to represent references
There are two kinds of references we need to represent: References to a resource whose type we know statically at compile time,
and references to a resource whose type we do not know at compile time.

We could use the same type for both kinds of references, but that has the downside of allowing a situation where
we know the group and version statically at compile time, but the customer has also provided it and it doesn't match.
Two types allows us to clearly express what we're expecting for each reference. The resulting YAMLs look basically the same
to the customer, and the required-ness of the fields will give push-back when customers need to specify a group or kind and
have not.

```go
// KnownResourceReference is a resource reference to a known type.
type KnownResourceReference struct {
	// This is the name of the Kubernetes resource to reference.
	Name string `json:"name"`

	// References across namespaces are not supported.

	// Note that ownership across namespaces in Kubernetes is not allowed, but technically resource
	// references are. There are RBAC considerations here though so probably easier to just start by
	// disallowing cross-namespace references for now
}

type ResourceReference struct {
	// The group of the referenced resource.
    Group string `json:"group"`
    // The kind of the referenced resource.
    Kind string `json:"kind"`
    // The name of the referenced resource.
    Name string `json:"name"`

    // Note: Version is not required here because references are all about linking one Kubernetes
    // resource to another, and Kubernetes resources are uniquely identified by group, kind, (optionally namespace) and
    // name - the versions are just giving a different view on the same resource
}
```

### How to represent ownership and dependent resources
We will use the same `KnownResourceReference` type as an additional `Owner` field on dependent resource specifications.

When we determine that a resource is a dependent resource of another resource kind, we will code-generate
an `Owner` property in the dependent resource `Spec`. This will also include an annotation 
about the expected type of the resource (group and kind) so that the customer doesn't have to specify that in the YAML. 

```go
type SubnetSpec struct {
    Owner KnownResourceReference `json:"owner" group:"microsoft.network.infra.azure.com" kind:"VirtualNetwork"`
    ...
}
```

When users submit a dependent object we will validate that the provided owner reference is present.
This can be accomplished by making the property required in the CRD.

A YAML snippet showing how this will look from the customer's perspective:
```yaml
...
  spec:
    owner:
      name: my-vnet
...
```

One major advantage of this approach is that the customer cannot really get the owning type wrong, because we've autogenerated the expected
group/kind information all names they supply must point to the right kind of resource.

### How to represent a resource generically
In addition to representing references generically, we will need the ability to reference ARM resources generically,
so that the generic controller can act on them without needing to cast to their specific type.

```go
// TODO: There may be more in this interface, or it may get rolled into MetaObject depending on yet to be determined implementation details
type ArmResource interface {
	// Name returns the ARM resource name
	Name() string

    // Owner returns the ResourceReference so that we can extract the Group/Kind for easier lookups
    Owner() *ResourceReference
}
```

### How to identify resource relationships

**For related (not owned) resources** we must find each field that represents a resource reference and transform its type to `ResourceReference`.
There is no specific marker which means: "This field is a reference" - most are called `id` but that's not a guarantee.
For example on the [VirtualMachineScaleSetIPConfigurationProperties](https://schema.management.azure.com/schemas/2019-07-01/Microsoft.Compute.json) 
the `subnet` field is of custom type `ApiEntityReference`, which has an `id` field where you put the ARM ID for a subnet.
This may require some manual work. One thing we can investigate doing long term
is see if there's a way to get teams to annotate "links" in their Swagger somehow.

**For dependent resources** we must identify all of the owner to dependent relationships between resources.
As discussed in what ARM does, this can be done using the `resources` property in the ARM deployment templates.
These are much easier to automatically detect than related resources as the dependent types are called out in the `resources` property explicitly.

### How to choose the right reference type (ResourceReference vs KnownResourceReference) at generation time
Because we are code-generating all of the `Owner` fields based on the `resources` property in the JSON schema, and each ARM resource can be owned by 
at most 1 other resource, we can always supply the annotations for group and kind automatically for the `Owner` field.
**This is not the case for abitrary references (`id`'s) to other resources**. We do not actually know
programmatically what type those references are. In some cases it may actually be allowed to point to multiple different types 
(for example: custom image vs shared image gallery).

In the `KnownResourceReference` case, we know the type we're looking for and can fail fast if the customer specifies the wrong type.
In the `ResourceReference` case, we cannot know the type we're looking for, so we must accept what the customer has provided and ensure that we have good error messages
if they have provided a link to an invalid resource (usually the error from Azure should suffice).

### ResourceLifeCycle and unmanaged resources
In order to keep references Kubernetes-native, allow a "single pane of glass" for customers looking at their Azure resources through Kubernetes, and 
allow references to resources that were created before the customer onboarded to the operator, we introduce a new mode to each resource: `ResourceLifeCycle`.

`ResourceLifeCycle` can be either `Managed` or `Unmanaged`.

`ResourceLifeCycle` is not specified by the customer explicitly in the `Spec`, instead it is inferred based on how
the resource was created in Kubernetes. If a resource is created as just a reference (id, name, no spec details) then
it is `Unmanaged`. If a resource is created with a populated spec, then the resource is `Managed`.

### RouteTable + Routes issue (multiple routes of the same name are allowed)

Options for this:
Note that all of these options share this restriction: Each resource must be imported, e.g. to import a VNET you may need to import
the resource group the VNET is in, and then the VNET (with an `Owner` reference pointing to the imported `ResourceGroup`). 

**Option 1:** Users must create a _valid_ resource with the same name as the resource they want to track. If this resource's spec differs from what is in Azure, an error is logged 
but we never actually apply any state to Azure (i.e. we don't try to sync to the spec). A tag in the metadata must be added to inform the operator not to sync.   

**Advantages**
- Swapping from unmanaged to managed is super easy, just remove the tag blocking the reconciliation loop.

**Disadvantages**
- There is possibly a significant amount of extra effort required to re-specify a resource whose shape we really just want to "import" from ARM. Worse for large trees of objects
  or deeply nested objects.
- If the tag is forgotten (or has a typo) we will try to manage a resource which we shouldn't be managing. This could be _very_ problematic depending on how different
  the specification is from what exists in Azure.
- The existence of a spec may suggest we are actually seeking towards it -- which we are not. ASO does have a similar feature though so maybe not that big of a problem.

**Option 2:** Users create an entity with just the "identifying fields" set: `Metadata.Name`, `Owner`, and optionally `Spec.AzureName`.
When an entity is created like this, the controller knows to treat it specially (optionally may also add a tag automatically?).
These entities will only be watched by the controller, no mutating update will be sent to ARM.

**Advantages**
- Relatively easy to import even complex object hierarchies.

**Disadvantages**
- This screws up the "required-ness" of non-identifying fields in a spec. For example: a Virtual Network requires a `Properties VirtualNetworkProperties` field to be set,
  but since we have to allow that field to be `nil` when importing a Virtual Network we can't set the `Properties` field with a required annotation for Kubebuilder.

Option 3: Same as option 2, but use `anyOf` to specify two valid structures:
```yaml
  spec:
    type: object
    properties:
      owner:
        properties:
          name:
            type: string
      azureName:
        type: string
      foo:
        type: integer
      anyOf:
        - required: ["owner"]
        - required: ["owner", "foo"]
```

Note that it has to be `anyOf` because `oneOf` disallows multiple matches, and `owner` + `foo` matches both sets in the example above.

**Advantages**
- Represents what we want and maintains better automatic validation.

**Disadvantages**
- Kubebuilder doesn't support generating this, so we would have to come up with another way to do it, or possibly upstream changes to Kubebuilder to support it.

**Option 4:** Other ideas...
Do away with Kubebuilder validation entirely and use our own (including our own validating webhooks).
Use Kustomize and our own code-generator/parser to generate amendments to Kubebuilder's generated CRDs to get the `anyOf` shape we want above.

### How to transform Kubernetes objects to ARM objects (and back)
In the case of resource ownership, the proposed `Owner` property exists on dependent resources in the CRD but must not go to Azure as Azure doesn't understand it.
In the case of a generic resource reference, the `ResourceReference` in the CRD must become an `id` (with fully-qualified ARM ID) when serialized to ARM.
In both cases, we need two representations of the entity: one to Kubernetes as the CRD, and one to Azure. These two types are structurally similar but not identical.
We cannot just override JSON serialization to solve this problem due to the fact that there are actually two distinct JSON representations we need.

The proposed solution is that the code generator intelligently generates 2 types for cases where we know the CRD shape differs from ARM.
We will add an interface which types can optionally implement which allows them to transform themselves to
another type prior to serialization to/from ARM. This is also a useful hook for any manual customization for serialization we may need.

The interface will look something like this:
```go
type ARMTransformer interface {
	ToArm(owningName string) (interface{}, error)
	FromArm(owner KnownResourceReference, input interface{}) error
}
```

Here's an example of how it will be implemented:

```go
func CreateArmResourceNameForDeployment(owningName string, name string) string {
	result := owningName + "/" + name
	return result
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
type VirtualNetworksSubnets struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworksSubnetsSpec `json:"spec,omitempty"`
}

var _ ArmResource = &VirtualNetworksSubnets{}

func (resource *VirtualNetworksSubnets) Owner() *ResourceReference {
    r := reflect.TypeOf(resource.Spec)
	ownerField, found := r.FieldByName("Owner")
    if !found {
        return nil
    }
    
    group := ownerField.Tag.Get("group")
    kind := ownerField.Tag.Get("kind")

    return &ResourceReference {
        group: group,
        kind: kind,
        name: resource.Spec.Owner.Name
    }
}
func (resource *VirtualNetworksSubnets) Name() string {
	return resource.Spec.Name
}

type VirtualNetworksSubnetsSpec struct {

	// +kubebuilder:validation:Required
	ApiVersion VirtualNetworksSubnetsSpecApiVersion `json:"apiVersion"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `json:"owner" group:"microsoft.network" kind:"VirtualNetworks"`

	// +kubebuilder:validation:Required
	//Properties: Properties of the subnet.
	Properties SubnetPropertiesFormat `json:"properties"`

	// +kubebuilder:validation:Required
	Type VirtualNetworksSubnetsSpecType `json:"type"`
}

// No KubeBuilder comments required here because not ever used to generate CRD
type VirtualNetworksSubnetsSpecArm struct {
	ApiVersion VirtualNetworksSubnetsSpecApiVersion `json:"apiVersion"`
	Name       string                               `json:"name"`

	//Properties: Properties of the subnet.
	Properties SubnetPropertiesFormat         `json:"properties"`
	Type       VirtualNetworksSubnetsSpecType `json:"type"`
}

// This interface implementation would be autogenerated for ARM resources with references
var _ genruntime.ArmTransformer = &VirtualNetworksSubnetsSpec{}

func (transformer *VirtualNetworksSubnetsSpec) ToArm(owningName string) (interface{}, error) {
	result = VirtualNetworksSubnetsSpecArm{}
	result.ApiVersion = transformer.ApiVersion
	result.Name = CreateArmResourceNameForDeployment(owningName, transformer.Name)
	result.Properties = transformer.Properties
	result.Type = transformer.Type
	return result, nil
}

func (transformer *VirtualNetworksSubnetsSpec) FromArm(owner genruntime.KnownResourceReference, input interface{}) error {
	typedInput, ok := input.(VirtualNetworksSubnetsSpecArm)
	if !ok {
		return fmt.Errorf("unexepected type supplied for FromArm function. Expected VirtualNetworksSubnetsSpecArm, got %T", input)
	}
	transformer.ApiVersion = typedInput.ApiVersion
	transformer.Name = ExtractKubernetesResourceNameFromArmName(typedInput.Name)
	transformer.Owner = owner
	transformer.Properties = typedInput.Properties
	transformer.Type = typedInput.Type
	return nil
}
```

### Controller example
Putting it all together, here's what a generic controller reconciliation loop would look like using the interfaces discussed previously.

```go
// Example usage -- error handling elided for brevity
func (gr *GenericReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
    scheme := ...
    gvk := ...
    client := ...
    
    // Load the object from etcd
    obj := scheme.New(gvk)
    resource := client.Get(req.NamespacedName, obj)

    // Get the owner details
    armResource := obj.(ARMResource)
    ownerRef := armResource.Owner()

    // Perform a get from Azure to see current resource state    
    armId := helpers.GetArmId(resource)
    objFromAzure := scheme.new(gvk) // We need to provide the empty type to deserialize into

    // Somehow construct a new object of type etcdObject
    if armTransformer, ok := objFromAzure.(ARMTransformer); ok {
        result := armTransformer.ToArm("") // This just converts from an empty kube shape to an empty arm shape
        armClient.GetIt(armId, result)
    
        armTransformer.FromArm(ownerRef, result)
    }

    // Perform a put to update resource state
    // Walk the owner hierarchy (assuming owner has no owner here for simplicity) to build owner name
    ownerGvk := ownerRef.ToGvk()
    owner := scheme.New(ownerGvk)
    ownerArmResource := owner.(ARMResource)
    ownerId := owner.Name()

    var toSerialize interface{}
    toSerialize = resource
    if armTransformer, ok := toSerialize.(ARMTransformer); ok {
        toSerialize = armTransformer.ToArm(ownerArmId)
    }
    json := json.Marshal(toSerialize)
    armClient.SendIt(json)
}
```

## FAQ

### What happens when a dependent resource specifies an `Owner` that doesn't exist?
The dependent resource will be stuck in an unprovisioned state with an error stating that the owner doesn't exist.
If the owner is created, the dependent resource will then be created by the reconciliation loop automatically.

### What happens when a resource contains a link to another resource which doesn't exist?
The resource with the link will be stuck in an unprovisioned state with an error stating that the linked resource doesn't exist.
This behavior is the same as for a dependent resource with a non-existent owner.

### How are the CRD entities going to be rendered as ARM deployments?
There are a few different ways to perform ARM deployments as [discussed in Dependent Resources](#dependent-resources).
Due to the nature of Kubernetes CRDs, each resource is managed separately and has its own reconcilation loop. It doesn't make sense to try to 
deploy a single ARM template with the entire resource graph. Each resource will be done in its own deployment (with
a `dependsOn` specified if required).

### Aren't there going to be races in resource creation?
Yes. If you have a complex hierarchy of resources (where resources have involved relationships between one another) and submit 
all of their YAMLs to the operator at the same time it is likely that some requests when sent to ARM will fail because of missing dependencies.
Those resources that failed to deploy initially will be in an unprovisioned state in Kubernetes, and eventually
all the resources will be created through multiple iterations of the reconciliation loop. 


### Aren't there going to be races in resource deletion?
Yes. `Owner` as discussed in this specification is informing Kubernetes _how Azure behaves_. The fact that 
a `ResourceGroup` is the owner of a `VirtualMachineScaleSet` means that when the `ResourceGroup` is deleted in 
Azure, the `VirtualMachineScaleSet` will be too. 

This means that practically speaking, we don't need Kubernetes garbage collection to 
perform deletion of resources in Azure. Azure is already going to do that automatically. We need Kubernetes garbage collection
to easily maintain sync with Azure.

As far as implementation goes this just means that when we are performing deletes in the generic controller and the 
resource is already deleted in Azure we just swallow that error and allow the Kubernetes object to be deleted.

### What exactly happens when a resource with an `Owner` is created?
Once the resource has been accepted by the various admissions controllers and has been cofirmed to match 
the structural schema defined in the CRD, the generic controller will attempt to look up
the owning resource in etcd (or in ARM if it's an `AzureReference`).

If the generic controller finds the owning resource, it updates the `ownerReference` in the object metadata
to include the `uid` of the owning resource and then submits an ARM template to ARM using the 
name of the owner and the name of the resource to build the name specified in the ARM template. It will
include the name of the owner in the `dependsOn` field. 

### What happens if an owning resource is deleted and immediately recreated?
Kubernetes garbage collection is based on object `uid`'s. As discussed above, we bind to that `uid` on 
dependent resource creation. If a resource is deleted and then recreated Kubernetes will still understand
that the new resource is fundamentally different than the old resource and garbage collection will happen
as expected. The result will be that there is a new owning resource but all of its dependent resources were 
deleted (in Azure and in k8s).

## TODOs
- How can we allow customers to easily find all dependents for a particular owner (i.e. all subnets of a vnet) using `kubectl`?
- Cross subscription refs? Note that these are supported by a few Azure resources (VNET for example), but aren't supported in most places.

## Questions
These are questions I am posing to the group - I don't expect to have an answer without input from the group.

- What to do with awkward resources where the owner requires at least 1 dependent to also be created with it?
  David Justice pointed out [this one](https://github.com/Azure/k8s-infra/blob/14105b1cb3f6967cd086c9f8f75fb16bb85d6318/apis/microsoft.network/v1/networkinterface_types.go#L43)
- Do we want to use the same type for ownership relationships and "related" relationships? Ownership has other angles such
  as how deletes propagate which in theory don't apply for other kinds of relationships.
- Do we need to worry about letting customers choose between [foreground cascading deletion](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/#foreground-cascading-deletion)
  and [background cascading deletion](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/#background-cascading-deletion) or do we just pick 
  one behavior which is best for our case?

# The road not travelled

## Shape of Azure References
We considered avoiding the complexity of `ResourceLifecycle` (`Managed` vs `Unmanaged`), instead allowing references to Azure resources directly by ARM ID.

References would look like this:
```go
type KnownResourceReference struct {
	Kubernetes KnownKubernetesReference `json:"kubernetes"`
	Azure      string                   `json:"azure"`
}

type ResourceReference struct {
	Kubernetes KubernetesReference `json:"kubernetes"`
	Azure      string              `json:"azure"`
}

type KnownKubernetesReference struct {
	// This is the name of the Kubernetes resource to reference.
	Name string `json:"name"`
	// References across namespaces are not supported.
	// Note that ownership across namespaces in Kubernetes is not allowed, but technically resource
	// references are. There are RBAC considerations here though so probably easier to just start by
	// disallowing cross-namespace references for now
}

type KubernetesReference struct {
	// The group of the referenced resource.
	Group string `json:"group"`
	// The kind of the referenced resource.
	Kind string `json:"kind"`
	// The name of the referenced resource.
	Name string `json:"name"`
	// Note: Version is not required here because references are all about linking one Kubernetes
	// resource to another, and Kubernetes resources are uniquely identified by group, kind, (optionally namespace) and
	// name - the versions are just giving a different view on the same resource
}
```

### Advantages compared to what we chose
- Can track resources which are not tracked by Kubernetes.
- Doesn't need to introduce `ResourceLifeCycle`. `ResourceLifeCycle` complicates the mental model of individual resources
  as now apply on a resource can fail due to `ResourceLifeCycle` being `Unmanaged`.
- Can support references to resource types which the operator doesn't yet support. It's likely that we can work around
  this in the chosen architecture if it becomes a big problem though.

### Disadvantages compared to what we chose
- References are not always Kubernetes-native looking.
- The reference structure is a more complex nested type, which makes references (which are common) more complicated.
- Moving from a resource link being to Azure directly to that same resource being managed/tracked by Kubernetes requires
  sweeping updates across all types referencing the migrated resource.
- Doesn't allow for a "single pane of glass" experience where customers can easily view all of their resources in a 
  Kubernetes native way.

## How to represent references
### Use fully qualified ARM ID (a single string) for all references
#### Pros
- Super simple to implement, because it's what ARM expects at the end of the day anyway.

#### Cons
- You can't easily transplant your YAML between subscriptions/resource groups because those IDs are in the YAML - you need
  templates and variables so that you can easily move between different resource groups or Subscriptions.
- Customers can't stay in Kubernetes-land, they have to move their mental model to an "Azure" model.

### Use built-in OwnerReference for owner references (customer setting these directly)
#### Pros
- Basically none - customers are not supposed to set this directly.

#### Cons
- OwnerReference requires the object UID, which cannot be known at template authoring time.
- OwnerReference only works for ownership relationships, not for references.

## Where ownership references are specified
### Ownership is from owner to dependent
#### Pros
- It makes getting a list of all resources under a particular owner very easy.

#### Cons
- Adding/deleting a new dependent resource requires an update to the owner.
- The owner can be in a failed state because dependent resources are missing. It feels like we're repeating our intent here:
  On the one hand, we told the owner that it should have 3 dependents, while on the other hand we only created 2 of those 3.
  It feels like the state of the resources in kubernetes (i.e. how many dependents there actually _are_) is already expressing 
  the intent for how many we want, so having that also on the owner seems duplicate.
