# API Versions

Specification for how we will ensure the ARM API version we use for interaction with ARM matches the version originally requested by a user when they created the resource in their Kubernetes cluster.

## Why do we need this?

Sometimes, in addition to structural changes, there are behaviour changes between ARM API versions. It's therefore important that we use the requested API version when interacting with ARM to ensure that we get the behaviour requested.

### Example

Revisting the CRM example from the [Versioning](versioning.md) specification, consider what happens if we have two available versions of the resource `Person`, lets call them **v1** and **v2**. In **v2** the new property `PostalAddress` is mandatory, requiring that everyone have a mailing address. 

If we have a valid **v1** `Person`, trying to submit that through the **v2** ARM API will fail because it's missing a postal address. 

## Proposed Solution

We need to preserve the original API Version of each resource, and use that to create an appropriate resource for ARM.

### API Preservation

When generating the storage variants of each resource, we'll inject a new `ApiVersion` property of type **string**.

The generated `AssignPropertiesTo*()` methods for each API resource will set the `ApiVersion` property to the *kind* (in the Kubernetes *group-version-kind* meaning) of the API version.

The generated `AssignPropertiesFrom*()` methods for each API resource will ignore the `ApiVersion` property.

Generated property assignment methods for conversion between storage variants will copy the `ApiVersion` property across, preserving the value.

### ARM Resource Generation

When an ARM resource is required, the running hub version of the resource will use the `ApiVersion` property to convert to the named API version, which will in turn be used to generate the required ARM resource.

