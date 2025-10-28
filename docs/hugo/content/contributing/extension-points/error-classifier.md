---
title: ErrorClassifier
linktitle: ErrorClassifier
weight: 40
---

## Description

`ErrorClassifier` allows resources to customize how the reconciler classifies and responds to errors returned by Azure Resource Manager (ARM). This extension is invoked whenever an ARM API call returns an error, giving the resource a chance to determine whether the error is retryable, fatal, or requires special handling.

The interface is called in the error handling path of all ARM operations (GET, PUT, PATCH, DELETE). Proper error classification is critical for determining reconciliation behavior and user feedback.

## Interface Definition

```go
type ErrorClassifier interface {
    ClassifyError(
        cloudError *genericarmclient.CloudError,
        apiVersion string,
        log logr.Logger,
        next ErrorClassifierFunc,
    ) (core.CloudErrorDetails, error)
}

type ErrorClassifierFunc func(
    cloudError *genericarmclient.CloudError,
) (core.CloudErrorDetails, error)
```

**Parameters:**
- `cloudError`: The error returned from ARM, including HTTP status, code, and message
- `apiVersion`: The ARM API version used for the request
- `log`: Logger for the current operation
- `next`: The default error classification function

**Returns:**
- `core.CloudErrorDetails`: Structured error information including classification
- `error`: Error if classification itself fails

The `CloudErrorDetails` structure includes:
- `Classification`: Fatal, Retryable, or other classifications
- `Code`: Error code from Azure
- `Message`: Human-readable error message
- `Retry`: Whether the operation should be retried

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

DNS Zone records can encounter specific transient errors during provisioning that should be retried:

```go
var _ extensions.ErrorClassifier = &DnsZonesARecordExtension{}

func (extension *DnsZonesARecordExtension) ClassifyError(
    cloudError *genericarmclient.CloudError,
    apiVersion string,
    log logr.Logger,
    next extensions.ErrorClassifierFunc,
) (core.CloudErrorDetails, error) {
    // First, call the default classifier
    details, err := next(cloudError)
    if err != nil {
        return core.CloudErrorDetails{}, err
    }

    // Check for DNS-specific retryable errors
    if isRetryableDNSZoneRecordError(cloudError) {
        log.V(Status).Info(
            "Classifying DNS zone record error as retryable",
            "Code", cloudError.Code(),
            "Message", cloudError.Message())
        
        // Override the classification to make it retryable
        details.Classification = core.ErrorRetryable
    }

    return details, nil
}

// Helper function to identify retryable DNS errors
func isRetryableDNSZoneRecordError(cloudError *genericarmclient.CloudError) bool {
    // DNS zones may temporarily fail during propagation
    code := cloudError.Code()
    return code == "DnsRecordInGracePeriod" ||
           code == "DnsZoneNotReady" ||
           code == "NameServerNotReady"
}
```

**Key aspects of this example:**

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

### Pattern 5: Multiple Error Conditions

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

    code := cloudError.Code()
    
    switch {
    case isTransientNetworkError(code):
        details.Classification = core.ErrorRetryable
        details.Retry = true
    case isQuotaError(code):
        details.Classification = core.ErrorFatal
        details.Message = "Quota exceeded. Please request a quota increase."
    case isAuthorizationError(code):
        details.Classification = core.ErrorFatal
        details.Message = "Insufficient permissions. Check service principal roles."
    default:
        // Keep default classification
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

Example test structure:

```go
func TestMyResourceExtension_ClassifyError(t *testing.T) {
    t.Run("default classification unchanged", func(t *testing.T) {
        // Test that unrecognized errors pass through
    })

    t.Run("transient error marked retryable", func(t *testing.T) {
        // Test custom retryable classification
    })

    t.Run("quota error marked fatal", func(t *testing.T) {
        // Test fatal classification
    })

    t.Run("enhanced error message", func(t *testing.T) {
        // Test message improvements
    })

    t.Run("api version specific handling", func(t *testing.T) {
        // Test version-specific logic
    })
}
```

Example test implementation:

```go
func TestDnsZonesARecordExtension_ClassifyError(t *testing.T) {
    extension := &DnsZonesARecordExtension{}
    
    // Create a mock default classifier
    defaultClassifier := func(err *genericarmclient.CloudError) (core.CloudErrorDetails, error) {
        return core.CloudErrorDetails{
            Classification: core.ErrorFatal,
            Code: err.Code(),
            Message: err.Message(),
        }, nil
    }

    t.Run("DnsRecordInGracePeriod is retryable", func(t *testing.T) {
        cloudError := &genericarmclient.CloudError{
            Code: "DnsRecordInGracePeriod",
            Message: "DNS record is in grace period",
        }

        details, err := extension.ClassifyError(
            cloudError,
            "2023-01-01",
            logr.Discard(),
            defaultClassifier)

        assert.NoError(t, err)
        assert.Equal(t, core.ErrorRetryable, details.Classification)
    })
}
```

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
