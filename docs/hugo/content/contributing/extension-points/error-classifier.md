---
title: ErrorClassifier
linktitle: ErrorClassifier
weight: 40
---

## Description

`ErrorClassifier` allows resources to customize how the reconciler classifies and responds to errors returned by Azure Resource Manager (ARM). This extension is invoked whenever an ARM API call returns an error, giving the resource a chance to determine whether the error is retryable, fatal, or requires special handling.

The interface is called in the error handling path of all ARM operations (GET, PUT, PATCH, DELETE). Proper error classification is critical for determining reconciliation behavior and user feedback.

## Interface Definition

See the [ErrorClassifier interface definition](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/error_classifier.go) in the source code.


## Motivation

The `ErrorClassifier` extension exists to handle cases where:

1. **Resource-specific errors**: Some Azure resources return unique error codes that require special handling

2. **Retryable conditions**: Certain errors that appear permanent are actually transient for specific resources

3. **Better user feedback**: Providing more context or clearer messages for resource-specific errors

4. **API version differences**: Error behavior may vary across API versions for the same resource

5. **Conditional retry logic**: Some errors are retryable under certain conditions but fatal under others

The default error classifier handles common patterns, but some resources have unique error behaviors that need custom classification.

## When to Use

Implement `ErrorClassifier` when:

- ✅ A resource returns specific error codes that need special handling
- ✅ Transient errors are being classified as fatal (or vice versa)
- ✅ Error messages need resource-specific clarification
- ✅ Certain errors should trigger different retry behavior
- ✅ API version affects error classification
- ✅ User feedback for errors needs improvement

Do **not** use `ErrorClassifier` when:

- ❌ The default classification handles the error correctly
- ❌ The error handling should apply to all resources (modify default classifier)
- ❌ The issue is with ARM itself (report to Azure)
- ❌ You want to suppress valid errors (fix the root cause instead)

## Example: DNS Zone Record Error Classification

See the [full implementation in dns_zones_a_record_extension.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/network/customizations/dns_zones_a_record_extension.go).

**Key aspects of this implementation:**

1. **Delegation to default**: Calls `next()` first to get standard classification
2. **Selective override**: Only modifies classification for specific errors
3. **Logging**: Records the classification decision for debugging
4. **Helper function**: Encapsulates the error detection logic
5. **Error propagation**: Returns errors from `next()` without modification


## Common Patterns

### Pattern 1: Marking Specific Errors as Retryable

```go
func (ex *ResourceExtension) ClassifyError(
    cloudError *genericarmclient.CloudError,
    apiVersion string,
    log logr.Logger,
    next extensions.ErrorClassifierFunc,
) (core.CloudErrorDetails, error) {
    // Get default classification
    details, err := next(cloudError)
    if err != nil {
        return core.CloudErrorDetails{}, err
    }

    // Check for resource-specific transient errors
    if ex.isTransientError(cloudError) {
        details.Classification = core.ErrorRetryable
        details.Retry = true
        log.V(Status).Info("Marking error as retryable", "code", cloudError.Code())
    }

    return details, nil
}
```

### Pattern 2: API Version-Specific Classification

```go
func (ex *ResourceExtension) ClassifyError(
    cloudError *genericarmclient.CloudError,
    apiVersion string,
    log logr.Logger,
    next extensions.ErrorClassifierFunc,
) (core.CloudErrorDetails, error) {
    details, err := next(cloudError)
    if err != nil {
        return core.CloudErrorDetails{}, err
    }

    // Certain errors behave differently in older API versions
    if apiVersion == "2021-01-01" && cloudError.Code() == "ResourceQuotaExceeded" {
        // In this API version, quota errors are temporary during provisioning
        details.Classification = core.ErrorRetryable
        details.Message = "Resource quota temporarily exceeded during provisioning, will retry"
    }

    return details, nil
}
```

### Pattern 3: Enhanced Error Messages

```go
func (ex *ResourceExtension) ClassifyError(
    cloudError *genericarmclient.CloudError,
    apiVersion string,
    log logr.Logger,
    next extensions.ErrorClassifierFunc,
) (core.CloudErrorDetails, error) {
    details, err := next(cloudError)
    if err != nil {
        return core.CloudErrorDetails{}, err
    }

    // Provide more context for common user errors
    if cloudError.Code() == "InvalidParameterValue" {
        if strings.Contains(cloudError.Message(), "sku") {
            details.Message = fmt.Sprintf(
                "Invalid SKU specified: %s. "+
                "Available SKUs for this resource: %s. "+
                "See documentation: https://docs.microsoft.com/...",
                extractSKU(cloudError.Message()),
                strings.Join(ex.getValidSKUs(), ", "))
        }
    }

    return details, nil
}
```

### Pattern 4: Conditional Fatal Classification

```go
func (ex *ResourceExtension) ClassifyError(
    cloudError *genericarmclient.CloudError,
    apiVersion string,
    log logr.Logger,
    next extensions.ErrorClassifierFunc,
) (core.CloudErrorDetails, error) {
    details, err := next(cloudError)
    if err != nil {
        return core.CloudErrorDetails{}, err
    }

    // Some configuration errors are unrecoverable
    if cloudError.Code() == "InvalidConfiguration" {
        details.Classification = core.ErrorFatal
        details.Retry = false
        details.Message = fmt.Sprintf(
            "Configuration error: %s. This requires manual intervention.",
            cloudError.Message())
    }

    return details, nil
}
```

## Error Classifications

The `core` package defines several error classifications:

- **`ErrorRetryable`**: The operation should be retried after a delay
- **`ErrorFatal`**: The operation cannot succeed without user intervention
- **`ErrorUnknown`**: Classification uncertain, treated conservatively

The controller uses these classifications to determine:
- Whether to requeue the reconciliation
- What condition to set on the resource
- How long to wait before retrying
- Whether to emit events

## Testing

When testing `ErrorClassifier` extensions:

1. **Test default pass-through**: Verify normal errors aren't affected
2. **Test specific error codes**: Cover all custom classification logic
3. **Test API version variations**: If version-specific, test all versions
4. **Test error message enhancement**: Verify improved user messages
5. **Test classification changes**: Assert correct classification results


## Important Notes

- **Always call `next()` first**: This ensures default classification logic runs
- **Be conservative**: When in doubt, prefer retryable over fatal
- **Document error codes**: Comment which Azure error codes you're handling
- **Log decisions**: Help debugging by logging classification changes
- **Consider idempotency**: Retryable errors will cause repeated operations
- **Test thoroughly**: Incorrect classification can cause user frustration

## Common Azure Error Codes

Some error codes you might encounter:

- `ResourceNotFound`: Resource doesn't exist (usually retryable for dependencies)
- `ResourceGroupNotFound`: Parent resource group missing
- `InvalidParameterValue`: Configuration error (usually fatal)
- `QuotaExceeded`: Subscription limits reached (fatal)
- `InternalServerError`: Azure service issue (retryable)
- `TooManyRequests`: Rate limiting (retryable with backoff)
- `Conflict`: Concurrent operation conflict (retryable)

## Related Extension Points

- [ARMResourceModifier]({{< relref "arm-resource-modifier" >}}): Modify requests before they can error
- [PreReconciliationChecker]({{< relref "pre-reconciliation-checker" >}}): Prevent operations that would error
- [PostReconciliationChecker]({{< relref "post-reconciliation-checker" >}}): Validate after operations
- [Deleter]({{< relref "deleter" >}}): Custom deletion with error handling
