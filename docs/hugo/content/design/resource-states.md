---
title: Proposal for reporting resource Status
linktitle: Resource status
---
## What Status are we talking about?
There are two types of Status that we're interested in understanding and reporting to the user when running an operator that creates resources in Azure:

1. The Status of the operator and its actions on the resource. Has the resource successfully been reconciled? Is the operator in the progress of driving the resource to its goal state (defined in the `spec`), or has it already done so and is now
   waiting for more changes before taking further action?
2. The State of the resource in Azure. What fields are set in Azure? There are often defaults and other values which, while you may not have specified them in the `spec` have been applied on the server side and we want to show to you.

This document is focused on the first type of Status defined above.

## Current state of ASO

### ASO v1 (handcrafted)

The handcrafted ASO resources have a 3-tuple on the `Status` field which describes the state of the resource:
```go
  Provisioning       bool                `json:"provisioning,omitempty"`
  Provisioned        bool                `json:"provisioned,omitempty"`
  FailedProvisioning bool                `json:"failedProvisioning,omitempty"`
```


There are also additional fields for conveying additional information, including the details of any error as well as information about the Resource ID of the resource
```go
  State              string              `json:"state,omitempty"`
  Message            string              `json:"message,omitempty"`
  ResourceId         string              `json:"resourceId,omitempty"`
```

#### Problems with this approach
* `Provisioning`, `Provisioned` and `FailedProvisioning` are mutually exclusive but the structure of the `status` doesn't convey that well. Additionally the fact that there are multiple
  fields means it's easier for bugs to creep in where we set one but didn't clear another.
* The `State` field is not widely used by all resources, and seems to conflict with `Provisioning`/`Provisioned`/`FailedProvisioning`. It's not clear when you should look at one or the other.

### ASO v2 (auto-generated)
The auto-generated resources don't currently have anything in their `Status` which conveys the Status _of the operator_. For the initial alpha releases, these fields are instead set as annotations on the object:
```go
  DeploymentIDAnnotation   = "deployment-id.infra.azure.com"
  DeploymentNameAnnotation = "deployment-name.infra.azure.com"
  ResourceStateAnnotation  = "resource-state.infra.azure.com"
  ResourceErrorAnnotation  = "resource-error.infra.azure.com"
```

This was done mostly to avoid the additional work to code-generate the correct fields in each resources `status` and set them. Since it's possible that the whole `status` can get lost, some of the things may need to stay 
as annotations (`deploymentId` possibly), while others are obviously status and can be derived again if lost. We took the temporary approach of putting it all in annotations while we learned more about what the best practices were.

#### Problems with this approach
* The `ResourceStateAnnotation` is reusing the same enum to describe state as the `armclient` uses for deployments. This is problematic because that only supports `Succeeded`, `Failed`, `Deleting`, and `Accepted`.
  `Succeeded`, `Failed`, and `Deleting` make sense in the context of the operator but `Accepted` doesn't really. There's no state for `InProgress` or `Reconciling`.
* There's an awkward tension between sharing the state of the ARM deployment and the state of the resource. They're often the same but sometimes not (such as in the case the deployment is being deleted but the resource is not).
* Using annotations is awkward because changing an annotation triggers another reconcile.
* Using annotations means that it's not clear to the user what the possible valid states are for the fields that are enums.
* Users expect to look in `Status` for status related fields. Having it exposed as an annotation was never intended as a long term solution, just a hack so that we could get things up and running.

## Examining other projects like ASO
A quick look across the field of projects similar to ASO suggests that there is a relatively standard approach to solving this problem:

* Crossplane reports status through a [Ready condition](https://crossplane.io/docs/v1.9/reference/troubleshoot.html#resource-status-and-conditions)
* ACK reports status through a [variety of conditions](https://github.com/aws-controllers-k8s/runtime/blob/8191f606c5975d5ba0a1433fab695533980733ba/apis/core/v1alpha1/conditions.go), including `ACK.Adopted`, `ACK.resourceSynced`, `ACK.Terminal`, etc.
* Cluster API originally used a `phase` and `failureReason`/`failureMessage` pattern, but has since [moved to use conditions](https://github.com/kubernetes-sigs/cluster-api/blob/master/docs/proposals/20200506-conditions.md) and is deprecating the old pattern.
* Pod uses a combination of [phase](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-phase) and [conditions](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-conditions), including `PodScheduled`, `ContainersReady`, `Initialized`, and `Ready`.

It seems that most of these (and many core Kubernetes resources) use a `Ready` condition to specify the status.

## More on conditions
These best practices were gathered from [conditions API conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties) as of July 2021.

> Some resources in the v1 API contain fields called phase, and associated message, reason, and other status fields. The pattern of using phase is deprecated. Newer API types should use conditions instead. Phase was essentially a state-machine enumeration field, that contradicted system-design principles and hampered evolution, since adding new enum values breaks backward compatibility. Rather than encouraging clients to infer implicit properties from phases, we prefer to explicitly expose the individual conditions that clients need to monitor. Conditions also have the benefit that it is possible to create some conditions with uniform meaning across all resource types, while still exposing others that are unique to specific resource types. See [#7856](https://github.com/kubernetes/kubernetes/issues/7856) for more details and discussion.

> In condition types, and everywhere else they appear in the API, Reason is intended to be a one-word, CamelCase representation of the category of cause of the current status, and Message is intended to be a human-readable phrase or sentence, which may contain specific details of the individual occurrence. Reason is intended to be used in concise output, such as one-line kubectl get output, and in summarizing occurrences of causes, whereas Message is intended to be presented to users in detailed status explanations, such as kubectl describe output.

> In general, condition values may change back and forth, but some condition transitions may be monotonic, depending on the resource and condition type. However, conditions are observations and not, themselves, state machines, nor do we define comprehensive state machines for objects, nor behaviors associated with state transitions. The system is level-based rather than edge-triggered, and should assume an Open World.

> Controllers should apply their conditions to a resource the first time they visit the resource, even if the status is Unknown. This allows other components in the system to know that the condition exists and the controller is making progress on reconciling that resource.

These from [standardizing conditions](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1623-standardize-conditions):

> reason is required and must not be empty. The actor setting the value should always describe why the condition is the way it is, even if that value is "unknown unknowns". No other actor has the information to make a better choice.

Some other best practices on [Kubernetes design principles](https://github.com/kubernetes/community/blob/4c9ef2d/contributors/design-proposals/architecture/principles.md):

> Object status must be 100% reconstructable by observation. Any history kept must be just an optimization and not required for correct operation.

> Do not define comprehensive state machines for objects with behaviors associated with state transitions and/or "assumed" states that cannot be ascertained by observation.


## Proposal

### Goals
* The status of the operator should be reported as a subsection in the `Status` of the resource itself. See [#1504](https://github.com/Azure/azure-service-operator/pull/1504).
* It should be clear to users where to look to determine what the current status of their resource is.
* The status should include information in both normal and failure cases.
* When a failure occurs, it should be clear to the user that a failure has happened and what the cause of the failure was.
* When a failure occurs, there should be an error or code that is programmatically consumable (for automation, etc).
* When a failure occurs, there should be an error or text that is human consumable and ideally more informative than the programmatically consumable error code.
* When a failure occurs, it should be clear to the user whether that failure is transient and the operator can self-recover.

### Details
All resource's `Status`'s will have a top level `Conditions` field which is a collection of type `Condition`. This collection supports extension (through conditions with different type names).
Initially we will expose a single `Ready` condition across all resources, representing the high level availability of the resource and its readiness for use.

The structure of the `Condition` is based on the [recommended shape of conditions KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1623-standardize-conditions) as well as the [Cluster API conditions proposal](https://github.com/kubernetes-sigs/cluster-api/blob/master/docs/proposals/20200506-conditions.md) and is as follows:

```go
// ConditionSeverity expresses the severity of a Condition.
type ConditionSeverity string
const (
  // ConditionSeverityError specifies that a failure of a condition type
  // should be viewed as an error. Errors are fatal to reconciliation and
  // mean that the user must take some action to resolve
  // the problem before reconciliation will be attempted again.
  ConditionSeverityError ConditionSeverity = "Error"

  // ConditionSeverityWarning specifies that a failure of a condition type
  // should be viewed as a warning. Warnings are informational. The operator
  // may be able to retry through the warning without any action from the user, but
  // in some cases user action to resolve the warning will be required.
  ConditionSeverityWarning ConditionSeverity = "Warning"

  // ConditionSeverityInfo specifies that a failure of a condition type
  // should be viewed as purely informational. Things are working.
  // This is the happy path.
  ConditionSeverityInfo ConditionSeverity = "Info"

  // ConditionSeverityNone specifies that there is no condition severity.
  // For conditions which have positive polarity (Status == True is their normal/healthy state), this will set when Status == True
  // For conditions which have negative polarity (Status == False is their normal/healthy state), this will be set when Status == False.
  // Conditions in Status == Unknown always have a severity of None as well.
  // This is the default state for conditions.
  ConditionSeverityNone ConditionSeverity = ""
)

type ConditionType string
const (
  ConditionTypeReady = "Ready"
)

// Condition defines an extension to status (an observation) of a resource
type Condition struct {
  // Type of condition.
  // +kubebuilder:validation:Required
  Type ConditionType `json:"type"`

  // Status of the condition, one of True, False, or Unknown.
  // +kubebuilder:validation:Required
  Status metav1.ConditionStatus `json:"status"`

  // Severity with which to treat failures of this type of condition.
  // For conditions which have positive polarity (Status == True is their normal/healthy state), this will be omitted when Status == True
  // For conditions which have negative polarity (Status == False is their normal/healthy state), this will be omitted when Status == False.
  // This is omitted in all cases when Status == Unknown
  // +kubebuilder:validation:Optional
  Severity ConditionSeverity `json:"severity,omitempty"`

  // LastTransitionTime is the last time the condition changed.
  // +kubebuilder:validation:Required
  LastTransitionTime metav1.Time `json:"lastTransitionTime"`

  // Reason for the condition's last transition.
  // Reasons are upper CamelCase (PascalCase) with no spaces. A reason is always provided, this field will not be empty.
  // +kubebuilder:validation:Required
  Reason string `json:"reason"`

  // Message is a human readable message indicating details about the transition. This field may be empty.
  // +kubebuilder:validation:Optional
  Message string `json:"message,omitempty"`
}
```

### The `Ready` condition
The `Condition` definition above discusses support for arbitrary conditions. In practice at least for now, ASO will only expose a single `Ready` condition.
The addition of a `Severity` field inspired by [Cluster API](https://github.com/kubernetes-sigs/cluster-api/blob/master/docs/proposals/20200506-conditions.md) allows different combinations of `Severity` and `Reason` to combine together and describe complex scenarios.

#### The `Reason` field
Inside of the `Ready` condition we always include a `Reason` which provides a programmatically consumable reason that the resource is in the given state. 
`Reason` is derived from multiple sources depending on where in the resource lifecycle we are. The value in `Reason` may be set by the operator itself, or be set to the result of an error or response code received from Azure. 

The operator defines the following reasons:

| Reason          | Severity        | Meaning                                                                                                       |
| --------------- | --------------- | ------------------------------------------------------------------------------------------------------------- |
| Reconciling     | Info            | A request has been submitted to Azure. The operator may be waiting for a response from Azure.                 |
| WaitingForOwner | Warning         | The `owner` of this resource cannot be found in Kubernetes. Progress is blocked until the `owner` is created. |
| Deleting        | Info            | The resource is being deleted.                                                                                |
| Succeeded       | None (""/empty) | The `spec` has successfully been applied. No additional changes are being attempted at this time.             |

There are no failure conditions (`Severity = Error`) specified here because currently all fatal errors come directly from Azure. When an error response is received from Azure, the `Code` from Azure is set as the `Reason`, and the `Message` from Azure is set as the `Message`.

#### `Severity` meaning in the context of the `Ready` condition

* `Error` means that we were unable to reconcile the resource successfully. A `Ready` condition with `Status=False, Reason=<something>, Severity=Error` is no longer attempting to be reconciled. The user must make an update to the resource to fix one or more configuration problems.
* `Warning` means that something is wrong, but we haven't given up. The resource will continue to be reconciled until we can progress past the cause of the `Warning`. Users should examine the `Warning` as some may be due to the operator waiting for action from them, as in the case of `Reason=WaitingForOwner, Severity=Warning`. Others may be due to transient unavailability in Azure, as in the case of `Reason=InternalServerError, Severity=Warning`.
* `Info` means that everything is proceeded as expected. This is the "happy path".

## Reference
1. [Article on conditions and status reporting in Kubernetes](https://dev.to/maelvls/what-the-heck-are-kubernetes-conditions-for-4je7). Note that this article is a bit old as Cluster API uses conditions now, but it gives a great overview of the topic.
2. [Recommended shape of Conditions](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1623-standardize-conditions)
3. [Update condition guidance](https://github.com/kubernetes/community/pull/4521) - and the actual guidance itself is [here](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties)
4. [More condition recommendations](https://github.com/kubernetes-sigs/cluster-api/issues/1658#issuecomment-584794287)

## Open questions
1. There are some small differences between this proposal and what CAPI is doing - are we ok with these differences?
   * CAPI says `Reason` is optional, but we're making it required.

## Open questions answered
Open questions which have since been answered are below.

**Question:** Will this work with [kstatus](https://github.com/kubernetes-sigs/cli-utils/blob/master/pkg/kstatus/README.md), which recommends negative polarity conditions?
**Answer:** Yes, kstatus supports the `Ready` condition as well, with a [few caveats](https://github.com/kubernetes-sigs/cli-utils/blob/master/pkg/kstatus/README.md#the-ready-condition). Anecdotally, we feel that positive polarity conditions (like `Ready`) are more clear. As mentioned [above](#examining-other-projects-like-aso), there are many operators that follow this `Ready` pattern including Crossplane and CAPI. If need-be, we can work around the major problem with kstatus and `Ready` by providing a webhook that automatically includes it on all resource creations.

**Question:** The KEP for `Condition` says that `LastTransitionTime` should be updated any time the condition changes. The CAPI proposal says it should change when `Status` changes, but the actual CAPI implementation changes it any time the `Condition` changes.Which behavior do we want?
**Answer:** We will follow the KEPs definition (and CAPI's actual implementation) and update `LastTransitionTime` any time a `Condition` changes, even if that change is from `Status=False` to `Status=False`.

## An aside on ARM templates
This isn't related to the core topic, but it might be useful to understand because it has a big impact on where failure is possible when communicating with ARM.

You can read some guidelines about how ARM templates work [here](https://armwiki.azurewebsites.net/api_contracts/guidelines/templatedeployment.html).

The process basically boils down to the following and is documented [here](https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/async-api-reference.md):
1. ARM: Validate basic template syntax. Return an error for invalid JSON, invalid resource ID structure, etc.
2. ARM: Send the resource to the individual resource provider (RP).
3. RP: Return a response for the resource. This can be:
   * `201` or `200` HTTP status code and a response body. At this point the resource _should_ exist, but with possibly with a non-terminal `provisioningState`.
   * `202` HTTP status code. At this point the resource _should not_ exist, and will not exist until the long-running operation completes successfully.
4. ARM: Poll long running operation until it completes (if needed).
5. ARM: Deployment status will be the status of the resource at the end of the long running operation, which should be in one of the 3 terminal states `Succeeded`, `Failed`, or `Canceled`

Possible outcomes:

| Status of deployment                         | State of resource                                                                                               |
| -------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| 4xx HTTP error code when creating deployment | Resource does not exist                                                                                         |
| Deployment accepted, but fails after polling | Resource may or may not exist (depends on the RP)                                                               |
| Deployment succeeds                          | Resource has `Failed` `provisioningState` - The spec doesn't say this is impossible but it's probably very rare |
| Deployment succeeds                          | Resource has `Succeeded` `provisioningState`                                                                    |

**Note:** When a deployment is accepted, the underlying resource may or may not have been created. 
This means we have to handle cases where the deployment failed but the resource was created anyway alongside cases where the deployment failed and no resource was created.
