---
title: Extension Points
linktitle: Extension Points
menu:
  main:
    parent: Contributing
description: "How to extend Azure Service Operator v2 resources with custom behavior"
---

Azure Service Operator v2 provides several extension points that allow customization of resource behavior. These extension points enable contributors to compensate for variation in the behavior of resources by adding custom logic at various stages of the resource lifecycle.

## Overview

Extension points are Go interfaces defined in `v2/pkg/genruntime/extensions/` that resources can implement to customize their behavior. When a resource implements an extension interface, the controller will invoke the custom logic at the appropriate time during reconciliation.

### Extension Implementation Pattern

Extensions are typically implemented in resource-specific files under `v2/api/<service>/customizations/<resource>_extensions.go`. The general pattern is:

1. Declare that your extension type implements the interface:

   ```go
   var _ extensions.ARMResourceModifier = &MyResourceExtension{}
   ```

2. Implement the required method(s) of the interface

3. The controller automatically detects and uses the extension through a type-check at the appropriate time

### Available Extension Points

The following extension points are available for customizing resource behavior:

| Extension Point                 | Purpose                                        | When Invoked                              |
| ------------------------------- | ---------------------------------------------- | ----------------------------------------- |
| [ARMResourceModifier]           | Modify the ARM payload before sending to Azure | Just before PUT/PATCH to ARM              |
| [Deleter]                       | Customize resource deletion behavior           | When resource is being deleted            |
| [ErrorClassifier]               | Classify ARM errors as retryable or fatal      | When ARM returns an error                 |
| [Importer]                      | Customize resource import behavior             | During `asoctl import` operations         |
| [KubernetesSecretExporter]      | Export secrets to Kubernetes                   | After successful reconciliation           |
| [PostReconciliationChecker]     | Perform post-reconciliation validation         | After ARM reconciliation succeeds         |
| [PreReconciliationChecker]      | Validate before reconciling                    | Before sending requests to ARM            |
| [PreReconciliationOwnerChecker] | Validate owner state before reconciling        | Before any ARM operations (including GET) |
| [SuccessfulCreationHandler]     | Handle successful resource creation            | After initial resource creation           |

[ARMResourceModifier]: {{< relref "arm-resource-modifier" >}}
[Deleter]: {{< relref "deleter" >}}
[ErrorClassifier]: {{< relref "error-classifier" >}}
[Importer]: {{< relref "importer" >}}
[KubernetesSecretExporter]: {{< relref "kubernetes-secret-exporter" >}}
[PostReconciliationChecker]: {{< relref "post-reconciliation-checker" >}}
[PreReconciliationChecker]: {{< relref "pre-reconciliation-checker" >}}
[PreReconciliationOwnerChecker]: {{< relref "pre-reconciliation-owner-checker" >}}
[SuccessfulCreationHandler]: {{< relref "successful-creation-handler" >}}

## When to Use Extensions

Extensions should be used when:

- The generated code doesn't handle a specific Azure resource quirk
- Additional validation or logic is needed before/after ARM operations
- Custom error handling is required for specific scenarios
- Resources need special handling during creation, deletion, or import
- Secrets or configuration need custom export logic

Extensions should **not** be used for:

- Changes that could be made to the generator itself
- Logic that applies to all resources (consider modifying the controller instead)
- Working around bugs in the generator (fix the generator instead)

## Development Guidelines

1. **Keep extensions minimal**: Only add logic that cannot be handled by the generator
2. **Document thoroughly**: Explain why the extension is needed
3. **Type assert hub versions**: Include hub type assertions to catch breaking changes
4. **Handle errors gracefully**: Return appropriate error types and messages
5. **Test thoroughly**: Add unit tests for extension logic
6. **Call next**: Most extensions use a chain pattern - remember to call the `next` function

   ```go
   // Example: calling next in an ErrorClassifier
   func (ex *MyExtension) ClassifyError(
       cloudError *genericarmclient.CloudError,
       apiVersion string,
       log logr.Logger,
       next extensions.ErrorClassifierFunc,
   ) (core.CloudErrorDetails, error) {
       // First call the default classifier
       details, err := next(cloudError)
       if err != nil {
           return core.CloudErrorDetails{}, err
       }
       
       // Then apply custom logic
       if isMySpecialCase(cloudError) {
           details.Classification = core.ErrorRetryable
       }
       
       return details, nil
   }
   ```

## Related Resources

- [Adding a new code-generated resource]({{< relref "../add-a-new-code-generated-resource" >}})
- [Generator overview]({{< relref "../generator-overview" >}})
- [Testing]({{< relref "../testing" >}})
