// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package customizations

import (
	"context"
	"fmt"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	databasewatcher "github.com/Azure/azure-service-operator/v2/api/databasewatcher/v20241001preview/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

const (
	watcherStatusRunning = "Running"
	watcherStatusStopped = "Stopped"
)

// StartPollerResumeTokenAnnotation holds a watcher's start on the target, the resource written back.
const StartPollerResumeTokenAnnotation = "serviceoperator.azure.com/watcher-start-resume-token"

var _ extensions.PostReconciliationChecker = &TargetExtension{}

// PostReconcileCheck starts the watcher that owns this target. ARM creates every watcher stopped, with no
// property in the spec to change that, and refuses to start one until it has a target - so the target
// invokes the action rather than the watcher.
func (extension *TargetExtension) PostReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	resourceResolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	reconcilePolicies annotations.ResolvedReconcilePolicies,
	next extensions.PostReconcileCheckFunc,
) (extensions.PostReconcileCheckResult, error) {
	target, ok := obj.(*databasewatcher.Target)
	if !ok {
		return extensions.PostReconcileCheckResult{},
			eris.Errorf("cannot run on unknown resource type %T, expected *databasewatcher.Target", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = target

	// A target owned by an ARM ID has no ASO-managed watcher to start
	watcher, ok := owner.(*databasewatcher.Watcher)
	if !ok {
		return next(ctx, obj, owner, resourceResolver, armClient, log, reconcilePolicies)
	}

	// If we have an existing start operation to monitor, pick it up and see if it's done.
	// A watcher has no failure state of its own, so a start already submitted is followed to its end
	if token, submitted := startResumeToken(target); submitted {
		done, err := resumeStart(ctx, armClient, token)
		if err != nil {
			clearStartResumeToken(target)

			return extensions.PostReconcileCheckResult{},
				eris.Wrapf(err, "cannot start watcher %s, which owns target %s", watcher.Name, target.Name)
		}

		if !done {
			// Return a failure result so that the controller will requeue and check again later,
			//  allowing us to continue monitoring the operation.
			return extensions.PostReconcileCheckResultFailure(
				fmt.Sprintf("waiting for the watcher %q to start", watcher.Name),
			), nil
		}

		// Successful completion of the start operation
		clearStartResumeToken(target)

		return next(ctx, obj, owner, resourceResolver, armClient, log, reconcilePolicies)
	}

	// The status on the watcher is only as fresh as its own last reconcile, so ask Azure instead
	status, err := readWatcherStatus(ctx, armClient, watcher)
	if err != nil {
		return extensions.PostReconcileCheckResult{},
			eris.Wrapf(err, "cannot read watcher %q, which owns target %s", watcher.Name, target.Name)
	}

	// If the watcher is running we have nothing to do, continue the usual reconciliation flow.
	if status == watcherStatusRunning {
		return next(ctx, obj, owner, resourceResolver, armClient, log, reconcilePolicies)
	}

	// Any other status means an action is already under way, and a start now would conflict with it
	if status != watcherStatusStopped {
		// Stay short of ready so we're asked again, which is how the start is seen to have worked. Nothing is
		// owned by a target, so this can't withhold anything the start itself needs.
		return extensions.PostReconcileCheckResultFailure(
			fmt.Sprintf("waiting for the watcher %q to run", watcher.Name),
		), nil
	}

	// The watcher is at status "Stopped", so we may be able to start it.

	// This check still runs when the policy forbids modification, and a target ARM was never given is no
	// reason to start anything
	if !reconcilePolicies.Effective.AllowsModify() {
		return next(ctx, obj, owner, resourceResolver, armClient, log, reconcilePolicies)
	}

	// Nothing below holds for a watcher another operator has claimed
	if reason, ok := foreignWatcher(target, watcher); ok {
		return extensions.PostReconcileCheckResultFailure(reason), nil
	}

	// A policy that forbids modifying the watcher forbids starting it
	allowed, err := startAllowed(reconcilePolicies, watcher)
	if err != nil {
		// We couldn't work out whether starting the watcher is allowed, returning the error for visibility
		return extensions.PostReconcileCheckResult{}, err
	}
	if !allowed {
		// Not allowed to start the watcher, but we can continue with the rest of the reconciliation flow
		return next(ctx, obj, owner, resourceResolver, armClient, log, reconcilePolicies)
	}

	// Check the watcher is set up to be started.
	if reason, ok := watcherConfigured(target, watcher); ok {
		return extensions.PostReconcileCheckResultFailure(reason), nil
	}

	// Everything is aligned, let's ask Azure to start the watcher.
	// This is usually a long-running action, so we don't wait for it to finish.
	token, err := submitStart(ctx, watcher, armClient, log)
	if err != nil {
		return extensions.PostReconcileCheckResult{},
			eris.Wrapf(err, "cannot start watcher %q, which owns this target", watcher.Name)
	}

	if token != "" {
		// Store the token so we can resume the operation on a later reconcile
		setStartResumeToken(target, token)
	}

	// Stay short of ready so we're asked again, which is how the start is seen to have worked. Nothing is
	// owned by a target, so this can't withhold anything the start itself needs.
	return extensions.PostReconcileCheckResultFailure("waiting for the watcher to run"), nil
}

func startResumeToken(target *databasewatcher.Target) (string, bool) {
	token, ok := target.GetAnnotations()[StartPollerResumeTokenAnnotation]
	return token, ok
}

func setStartResumeToken(target *databasewatcher.Target, token string) {
	genruntime.AddAnnotation(target, StartPollerResumeTokenAnnotation, token)
}

func clearStartResumeToken(target *databasewatcher.Target) {
	genruntime.RemoveAnnotation(target, StartPollerResumeTokenAnnotation)
}

// resumeStart picks up an earlier start, reporting whether it finished. A failed one is an error here.
func resumeStart(ctx context.Context, armClient *genericarmclient.GenericClient, token string) (bool, error) {
	poller := armClient.ResumeActionPoller(genericarmclient.ActionPollerID)

	err := poller.Resume(ctx, armClient, token)
	if err != nil {
		return false, err
	}

	return poller.Poller.Done(), nil
}

// foreignWatcher reports why this target cannot act on the watcher at all, and is empty when it can. A
// resource is claimed before any extension runs, so one carrying no operator is unknown, not ours.
func foreignWatcher(
	target *databasewatcher.Target,
	watcher *databasewatcher.Watcher,
) (string, bool) {
	ours := operatorNamespace(target)
	if theirs := operatorNamespace(watcher); theirs == "" || theirs != ours || ours == "" {
		return fmt.Sprintf(
			"cannot start watcher %q, which is managed by the operator in %s while this target is managed by the operator in %s",
			watcher.Name,
			describeOperator(watcher),
			describeOperator(target),
		), true
	}

	return "", false
}

func watcherConfigured(
	target *databasewatcher.Target,
	watcher *databasewatcher.Watcher,
) (string, bool) {
	// ARM rejects the start (WatcherStartFailedDueToNoDataStore), so say what's missing
	// This check is mostly useful to stop us from consuming request quota on an action that can't possibly succeed
	if watcher.Spec.Datastore == nil {
		return fmt.Sprintf(
				"watcher %q has no datastore, so it cannot be started",
				watcher.Name,
			),
			true
	}

	if differingCredential(target, watcher) {
		return fmt.Sprintf(
			"cannot start watcher %q, which asks for %s while this target asks for %s",
			watcher.Name,
			describeCredential(watcher),
			describeCredential(target),
		), true
	}

	return "", false
}

func operatorNamespace(obj genruntime.MetaObject) string {
	return obj.GetAnnotations()[reconcilers.OperatorNamespaceAnnotation]
}

func describeOperator(obj genruntime.MetaObject) string {
	namespace := operatorNamespace(obj)
	if namespace == "" {
		return "an unknown namespace"
	}

	return fmt.Sprintf("namespace %q", namespace)
}

// differingCredential reports whether the watcher is managed with a credential this target cannot prove is
// its own. Only annotations can be compared, so anything short of equal is refused rather than assumed.
func differingCredential(target *databasewatcher.Target, watcher *databasewatcher.Watcher) bool {
	watcherCredential, watcherAsks := credentialAnnotation(watcher)
	targetCredential, targetAsks := credentialAnnotation(target)

	return watcherAsks != targetAsks || watcherCredential != targetCredential
}

// credentialAnnotation reports the secret a resource asks for, and whether it asks at all - naming an
// empty secret is not the same as naming none.
func credentialAnnotation(obj genruntime.MetaObject) (string, bool) {
	credential, ok := obj.GetAnnotations()[annotations.PerResourceSecret]
	return credential, ok
}

func describeCredential(obj genruntime.MetaObject) string {
	credential, asks := credentialAnnotation(obj)
	if !asks {
		return "the credential its namespace or the operator provides"
	}

	return fmt.Sprintf("credential %q", credential)
}

// startAllowed reports whether the watcher's own policy permits modifying it. An owner always shares the
// target's namespace, so a mismatch here is a resolution the policies can't answer rather than a refusal.
func startAllowed(
	policies annotations.ResolvedReconcilePolicies,
	watcher *databasewatcher.Watcher,
) (bool, error) {
	policy, err := policies.ForResource(watcher)
	if err != nil {
		return false, eris.Wrapf(err, "resolving the reconcile policy of watcher %q", watcher.Name)
	}

	return policy.AllowsModify(), nil
}

// submitStart asks Azure to start the watcher, returning a token for the operation. Nothing waits here.
func submitStart(
	ctx context.Context,
	watcher *databasewatcher.Watcher,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
) (string, error) {
	id, hasID := genruntime.GetResourceID(watcher)
	if !hasID {
		return "", eris.Errorf("cannot start watcher %s, it has no resource ID", watcher.Name)
	}

	log.V(Status).Info("Starting watcher", "name", watcher.AzureName())

	poller, err := armClient.BeginPostActionByID(ctx, id, "start", watcher.GetAPIVersion())
	if err != nil {
		return "", err
	}

	// A start Azure finished while answering has no operation to follow, and no token to ask for
	if poller.Poller.Done() {
		return "", nil
	}

	return poller.Poller.ResumeToken()
}

// watcherState is the part of a watcher read back from ARM directly.
type watcherState struct {
	Properties struct {
		Status string `json:"status"`
	} `json:"properties"`
}

func readWatcherStatus(
	ctx context.Context,
	armClient *genericarmclient.GenericClient,
	watcher *databasewatcher.Watcher,
) (string, error) {
	id, hasID := genruntime.GetResourceID(watcher)
	if !hasID {
		return "", eris.Errorf("cannot read watcher %s, it has no resource ID", watcher.Name)
	}

	var state watcherState
	_, err := armClient.GetByID(ctx, id, watcher.GetAPIVersion(), &state)
	if err != nil {
		return "", eris.Wrap(err, "reading watcher")
	}

	return state.Properties.Status, nil
}
