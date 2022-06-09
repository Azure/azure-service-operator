---
title: Improving the Reconciler interface
---

# Reconciler Interface

Today the interface is:

```go
type Reconciler interface {
    Reconcile(
        ctx context.Context,
        log logr.Logger,
        eventRecorder record.EventRecorder,
        obj MetaObject) (ctrl.Result, error)
}
```

This is pretty similar to the interface that controller-runtime gives us:

```go
type Reconciler interface {
    // Reconcile performs a full reconciliation for the object referred to by the Request.
    // The Controller will requeue the Request to be processed again if an error is non-nil or
    // Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
    Reconcile(context.Context, Request) (Result, error)
}
```


The problem with the interface at this level is that each controller must duplicate certain checks. This not only involves needless code duplication, but also opens the door for subtle bugs when differences occur.

## What's duplicated today

### Create vs Delete check

```go
    var result ctrl.Result
    var err error
    if !obj.GetDeletionTimestamp().IsZero() {
        result, err = r.Delete(ctx, log, typedObj)
    } else {
        result, err = r.CreateOrUpdate(ctx, log, typedObj)
    }

    if err != nil {
        return ctrl.Result{}, err
    }
```

### Handling reconcile policy

In delete: 
```go
    reconcilePolicy := reconcilers.GetReconcilePolicy(user, log)
    if !reconcilePolicy.AllowsDelete() {
        log.V(Info).Info("Bypassing delete of resource due to policy", "policy", reconcilePolicy)
        return ctrl.Result{}, r.RemoveResourceFinalizer(ctx, log, user)
    }
```

In create/update:
```go
    reconcilePolicy := reconcilers.GetReconcilePolicy(user, log)
    if !reconcilePolicy.AllowsModify() {
        return r.handleSkipReconcile(ctx, log, user)
    }
```

Most of the implementation of `handleSkipReconcile` is also duplicated. For ARM resources we _also_ want to populate status as part of that call, which requires a GET call to ARM; not every kind of resource will need this extra step.

### Claiming a resource

```go
    if r.NeedToClaimResource(user) {
        // TODO: This assumes we're requeing and doing the "stages" thing, which I am not sure if it makes sense to do for everybody anymore
        _, err := r.ClaimResource(ctx, log, user)
        if err != nil {
            return ctrl.Result{}, err
        }
    }
```

Each resource must add the finalizer, so that check could be lifted out into the generic controller. Other claiming activities are resource specific. 
An example of a resource specific claiming activity is checking that the owner exists. Only resource types that support owner will do this. There may be additional parts to the `Claim` workflow as well, such as 
determining what the ARM ID of the resource is going to be and saving that as an annotation.

### Checking if a resource has been claimed before deletion

```go
    // If we have no resourceID to begin with, or no finalizer, the Azure resource was never created
    hasFinalizer := controllerutil.ContainsFinalizer(r.Obj, reconcilers.GenericControllerFinalizer)
    resourceID := genruntime.GetResourceIDOrDefault(r.Obj)
    if resourceID == "" || !hasFinalizer {
        return ctrl.Result{}, r.RemoveResourceFinalizer(ctx, r.Log, r.Obj)
    }

```

Note that this check may need to be expanded by individual reconcilers. For example the one above checks both resourceID (ARM specific) and the standard finalizer (generic)

### Some logging/eventing

```go
    msg := "Starting delete of resource"
    r.Log.V(Status).Info(msg)
    r.Recorder.Event(r.Obj, v1.EventTypeNormal, string(DeleteActionBeginDelete), msg)
```

## Proposal

We create a more detailed interface and do as much work as we can in the `GenericController`. Each of these proposed methods would be idempotent:

```go
type Reconciler interface {
    CreateOrUpdate(
        ctx context.Context,
        log logr.Logger,
        eventRecorder record.EventRecorder,
        obj MetaObject) (ctrl.Result, error)

    Delete(
        ctx context.Context,
        log logr.Logger,
        eventRecorder record.EventRecorder,
        obj MetaObject) (ctrl.Result, error)

    // No ctrl.Result return here. An error means we can't proceed, no error means proceed directly to CreateOrUpdate.
    Claim(
        ctx context.Context,
        log logr.Logger,
        eventRecorder record.EventRecorder,
        obj MetaObject) error 

    // This is only called by the GenericController in cases where CreateOrUpdate couldn't be called due to reconcile policy
    // No ctrl.Result return here. An error means we couldn't get the status. Generic Controller will set the Ready condition
    // with this error and automatically retry.
    UpdateStatus(
        ctx context.Context,
        log logr.Logger,
        eventRecorder record.EventRecorder,
        obj MetaObject) error
}
```

Here's a sketch of what this would look like in the generic controller. This code is just illustrative, it'd be refactored to be a bit less verbose in the actual implementation:
```go
    // Check namespace annotation to ensure we aren't reconciling resource that this controller doesn't "own". This exists in generic controller today
    if !CheckNamespaceAnnotation() {
        return ctrl.Result{}, nil
    }

    // Claim the resource
    err := gr.Reconciler.Claim(ctx, log, gr.Recorder, metaObj)
    if readyErr, ok := conditions.AsReadyConditionImpactingError(err); ok {
        log.Error(readyErr, "Encountered error impacting Ready condition")
        err = gr.WriteReadyConditionError(ctx, metaObj, readyErr)
    }
    if err != nil {
        log.Error(err, "Error claiming resource")
        return ctrl.Result{}, reconcilers.IgnoreNotFoundAndConflict(err)
    }
    log.V(Info).Info("adding finalizer")
    controllerutil.AddFinalizer(obj, GenericControllerFinalizer)
    err = gr.KubeClient.CommitObject(ctx, obj)
    if err != nil {
        log.Error(err, "Error adding finalizer")
        return ctrl.Result{}, reconcilers.IgnoreNotFoundAndConflict(err)
    }

    // Take the action needed
    var result ctrl.Result
    var err error
    if !metaObj.GetDeletionTimestamp().IsZero() {
        // Check the reconcile policy to ensure we're allowed to issue a delete
        reconcilePolicy := GetReconcilePolicy(metaObj, log)
        if !reconcilePolicy.AllowsDelete() {
            log.V(Info).Info("Bypassing delete of resource due to policy", "policy", reconcilePolicy)
            return ctrl.Result{}, RemoveResourceFinalizer(ctx, log, metaObj)
        }

        // Check if we actually need to issue a delete
        hasFinalizer := controllerutil.ContainsFinalizer(metaObj, reconcilers.GenericControllerFinalizer)
        if !hasFinalizer {
            log.Info("Deleted resource")
            return ctrl.Result{}, nil
        }
        result, err = r.Delete(ctx, log, gr.Recorder, metaObj) // It would be the responsibility of the reconcilers delete to check anything specific that it did during Claim and short-circuit if needed
    } else {
        // Check the reconcile policy to ensure we're allowed to issue a create or update
        reconcilePolicy := reconcilers.GetReconcilePolicy(metaObj, log)
        if !reconcilePolicy.AllowsModify() {
            return handleSkipReconcile(ctx, log, metaObj) // This would call GetStatus to optionally fill out the status and Commit that
        }

        result, err = r.CreateOrUpdate(ctx, log, gr.Recorder, metaObj)
    }

    if readyErr, ok := conditions.AsReadyConditionImpactingError(err); ok {
        log.Error(readyErr, "Encountered error impacting Ready condition")
        err = gr.WriteReadyConditionError(ctx, metaObj, readyErr)
    }
    if err != nil {
        log.Error(err, "Error claiming resource")
        return ctrl.Result{}, reconcilers.IgnoreNotFoundAndConflict(err)
    }
```


# FAQ

**Q:** What component would own commiting the object to etcd?

**A:** `GenericController` would always do a commit to etcd after a reconcile. It may need to check that the actual individual reconciler changed `Obj`. If no change it won't commit. Commit ability will not be taken from individual reconcilers though, if 
       they need it they can do extra commits using the `KubeClient` they have access to.

**Q:** What component would own setting the ready condition?

**A:** The `GenericController`. Custom errors would be propagated via the error mechanism we have now from the individual reconciler to the `GenericController`. 
       There are no custom success states currently, but if there were we wouldn't block the reconciler from setting them manually if we wanted some in the future.

**Q:** Why split out `Claim`?

**A:** Part of `Claim` is general (adding the standard finalizer). Part of it is reconciler specific (for example checking if the owner exists). `Claim` is part of a fault-tolerance model. The resource must actually succeed in writing its finalizer to
       etcd before it interacts with ARM or SQL or whatever other data store the CR is telling it to talk to. It's important (but subtle) to get that right, so it makes sense to lift it out of the individual reconcilers and ensure that the `Commit` happens.
       Lifting `Claim` to be part of the interface also makes it easy to ensure that all reconcilers are doing things in the right order (and thus maintaining the fault tolerance guarantees we need).

**Q:** Why split out `UpdateStatus`?

**A:** In the standard case, it's expected that `CreateOrUpdate` will do this itself. When `reconcilePolicy` is set to skip updates, we just want to update the `Status` but not actually perform an `CreateOrUpdate` actions.
   This is split out so that the `GenericController` can make the check and perform the minimal required operation (just refreshing `Status`).

**Q:** What are the downsides of this?

**A:** Pretty much the standard framework vs library argument. We're restricting ourselves by pulling some of these calls up into `GenericController`. Now all of our reconcilers must look a particular way. We'll still have a library for doing most things
       (`Commit`, etc), but we're forcing a particular lifecycle for resources and it's possible in the future there is friction there.

**Q:** What are the upsides of this?

**A:** Consistency and testing are the big ones. By pulling this code up into `GenericController` and forcing each reconciler to implement it we remove the possibility that an implementer forgets to handle `reconcilePolicy`, or does it in a way that's incorrect.
       We reduce the amount of boilerplate required to write a custom reconciler as well. 
       Since all code takes the same codepath for common things like finalizer addition and `reconcilePolicy` checks, we can be confident that those features work for all resources without needing to test them for each handcrafted resource manually. That's 
       especially important because the handcrafted resources may not be HTTP based (`MySQLUser` isn't, for example), which means tests for those resources are quite expensive (in time cost) and fragile (susceptible to flakiness due to Azure).