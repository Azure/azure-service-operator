# Thoughts on performing provider-azure code generation for Crossplane

The patterns identified below were extrapolated from [the Crossplane VNET spec](https://github.com/crossplane/provider-azure/blob/master/config/crd/network.azure.crossplane.io_virtualnetworks.yaml).

## General notes

### Crossplane has the following static "special" properties
  1. ForProvider - this contains the actual "spec" as Azure would see it.
  2. `DeletionPolicy`
  3. `ProviderConfigRef`
  4. `WriteConnectionSecretToRef`

Since these are statically shaped they should be relatively easy to add with Crossplane specific generator pipeline stages. 

Items #2-4 are included automatically by embedding the  `runtimev1alpha1.ResourceSpec` from `runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"`.

### Cross resource references in Crossplane
1. `X` (e.g. `ResourceGroupName`) - An actual field that exists in the Azure API. Can be set to any string you'd post to the Azure API.
2. `XRef` (e.g. `ResourceGroupNameRef`) - A reference to a named Kubernetes object of a preordained kind. Used to resolve a value for `X`.
3. `XNameSelector` (e.g. `ResourceGroupNameSelector`) - A selector used to set `XRef`.

Note that these values cascade down from the selector, to the ref, to the underlying field. If the ref is already set the selector is ignored, 
and if the field is already set the ref is ignored.

Usually these are used to specify parentage for a resource (as in the case of the example above with `ResourceGroupName`) but may also be used for other
cross resource references.

Generating these might take a bit of work as we need to derive the name(s) for each of these for each of the parent resources. This should be doable 
in a pipeline stage by walking up the graph of owners that we already have and getting their names, then generating a set of these for each owner.

### Status 
Our status' include the "full" shape of the object, whereas Crossplane's include the things which are not in the spec (basically the server-only fields).
As pointed out by @negz, it's a bit weird to duplicate data across Spec and Status like we're currently doing so it may be worth us considering their approach.
[#269](https://github.com/Azure/k8s-infra/issues/269) is tracking this discussion.

## Missing features

These are features the k8s-infra code generator is currently lacking that would be required to do a good job of generating the Crossplane CRDs.

### [#266](https://github.com/Azure/k8s-infra/issues/266) Support for embedding structs/interfaces into generated types

We don't currently support embedding in the generated types, and we will need to in order to support 
embedding of `runtimev1alpha1.ResourceSpec` and `runtimev1alpha1.ResourceStatus`.

### Support for `additionalPrinterColumns` in the CRD YAML

Since these are likely custom per resource I am not sure if we want to generate them or just merge them in with hand-maintained Kustomize files.

### [#267](https://github.com/Azure/k8s-infra/issues/267) Detection and removal of embedded subresources.
Some of our CRDs aren't as clear as they should be due to the service teams embedding subresources in the parent resource as properties. For example on VNET
we have `subnets` and `virtualNetworkPeerings` properties that really probably shouldn't be there (since those are their own sub-resources).

### Improvements in JSON schema/Swagger upstream for marking read-only properties
This is minor as we can bypass properties using exclusions in the configuration file for now, but ideally we'd push this stuff up to the actual "source of truth" documents.

An example of this is: our VNET Spec has `provisioningState` in the `spec` which doesn't make sense.

### A plan for secrets
Right now we don't have a plan for fields which are classified as secret. This will matter for some things such as Microsoft SQL Server.
We do already have an issue tracking this at #154.

### Support for resource references other than ownership
We have a plan for supporting Cross resource references but don't currently have the data to easily detect them. Ideally we could detect these and then 
turn them into our resource reference shape, or into Crossplanes.

### [#268](https://github.com/Azure/k8s-infra/issues/268) Promote "Properties" Spec property
This shows up on quite a large number of resources and adds an extra level 
of nesting that doesn't look great.

## Specific resources

### VNET
**Note**: VNET in Crossplane is a `v1alpha1` API and as such doesn't follow all of their best practices (`v1beta1` APIs do)

- As mentioned above, our VNET today has `subnets` and `virtualNetworkPeerings`, which are really subresources. It also has `provisioningState` which it shouldn't.
  **If we excluded these properties in the configuration file, we could definitely hack towards generating a VNET that looks like Crossplane's now**.

### Redis

- Our version of Redis has a `Properties` field in the spec, which holds a bunch of properties. Some properties (`Location`, `Owner`, `Tags`) are peers of `Properties`.
**We should consider promoting the contents of `Properties` up a level into `Spec`, as mentioned above.** 
- There are a few trivial capitalization differences between our generated Redis and the one that is in Crossplane. For example: `EnableNonSSLPort` vs our `EnableNonSslPort`.
Crossplane is correct here, but the difference only shows up in the generated code and so isn't a big deal.
- We generate Enums for `Redis.Sku` fields `Name` and `Family`, whereas Crossplane just uses strings.
- Neither Crossplane nor we support `SubnetId` as a reference to a Kubernetes resource, but we probably should. (They have a TODO in the code).
- Crossplane makes use of a `+immutable` annotation, presumably to annotate that certain fields cannot be changed once they are set. We should look into doing that as well.
[#180](https://github.com/Azure/k8s-infra/issues/180) is tracking this. It's not clear to me that this annotation actually _does_ anything right now though.
- Crossplane is making use of a newer version of the Redis API than we see, because the 2018+ versions of the Redis API seem to have not made it into the
deployment template we're using as the source of truth. I filed [#1237](https://github.com/Azure/azure-resource-manager-schemas/issues/1237) in the schemas repro tracking this.

### Microsoft SQL

- The strange thing with SQL is somehow some of their APIs don't have the full API surface area... For example I think `2019-06-01-preview` exists but it's not in the JSON schema we're using. Other bits are split across API versions but then unioned together in their SDK:
[2015-05-01-preview](https://schema.management.azure.com/schemas/2015-05-01-preview/Microsoft.Sql.json#/resourceDefinitions) has servers, and [2017-03-01-preview](https://schema.management.azure.com/schemas/2017-03-01-preview/Microsoft.Sql.json#/resourceDefinitions) has databases.
[Their Swagger](https://github.com/Azure/azure-rest-api-specs/tree/master/specification/sql/resource-manager/Microsoft.Sql) is equally confusing.
- **As mentioned above, the big issue here is going to be secrets/credentials since we don't have an automated solution for that yet. Other than that, I think `servers` and `databases` (which are AFAIK the two interesting resources here) should be doable today.**

## Plugging into Crossplane framework
One other area we need to investigate is how to write a generic adapter that takes standard ARM deployments and turns them into something that plugs into the Crossplane framework (CRUD). We might need some discussion with the Crossplane folks here because one thing with deployments is that they are actually their own resource, so a Create VNET operation would have to
create a deployment that created the VNET. Once the deployment is done (resource in steady state), we'd want a way to conditionally delete the deployment (that doesn't delete the VNET). 

For those curious, we've got the scaffolding to do that in our generic controller in PR #250, you can see the workflow [here](https://github.com/Azure/k8s-infra/pull/250/files#diff-043a497af8a3c32ba21f705feaa1d3dcR222).