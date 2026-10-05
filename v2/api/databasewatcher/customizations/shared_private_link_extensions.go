// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package customizations

import (
	"context"
	"fmt"
	"path"
	"strings"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	databasewatcher "github.com/Azure/azure-service-operator/v2/api/databasewatcher/v20241001preview/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

// authorizationFailedErrorCode is what ARM answers a request the credential has no rights to make.
const authorizationFailedErrorCode = "AuthorizationFailed"

// connectionNotPendingErrorCode is what ARM answers an approval of a connection that has already left Pending.
const connectionNotPendingErrorCode = "PrivateEndpointConnectionStatusNotPending"

const (
	connectionStateApproved     = "Approved"
	connectionStatePending      = "Pending"
	connectionStateRejected     = "Rejected"
	connectionStateDisconnected = "Disconnected"
)

// connectionStates are the states ARM documents, under the casing it documents them in.
var connectionStates = []string{
	connectionStateApproved,
	connectionStatePending,
	connectionStateRejected,
	connectionStateDisconnected,
}

// ApprovalPollerResumeTokenAnnotation holds an approval in flight on the resource a link points at.
const ApprovalPollerResumeTokenAnnotation = "serviceoperator.azure.com/shared-private-link-approval-resume-token"

var _ extensions.PostReconciliationChecker = &SharedPrivateLinkExtension{}

// PostReconcileCheck completes the private endpoint connection a shared private link opens on the resource it
// points at. Azure opens that connection pending the resource owner's approval, and reports nothing about it
// back on the link, so the connection is read - and approved - on the resource itself.
func (extension *SharedPrivateLinkExtension) PostReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	resourceResolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	reconcilePolicies annotations.ResolvedReconcilePolicies,
	next extensions.PostReconcileCheckFunc,
) (extensions.PostReconcileCheckResult, error) {
	link, ok := obj.(*databasewatcher.SharedPrivateLink)
	if !ok {
		return extensions.PostReconcileCheckResult{},
			eris.Errorf("cannot run on unknown resource type %T, expected *databasewatcher.SharedPrivateLink", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = link

	// A resource named by ARM ID alone tells us no API version to read its connections with, so a link to one
	// keeps the readiness it has always had
	ref := link.Spec.PrivateLinkResourceReference
	if ref == nil || !ref.IsKubernetesReference() {
		return next(ctx, obj, owner, resourceResolver, armClient, log, reconcilePolicies)
	}

	resource, err := resourceResolver.ResolveReference(ctx, ref.AsNamespacedRef(link.Namespace))
	if err != nil {
		return extensions.PostReconcileCheckResult{},
			eris.Wrapf(err, "cannot resolve the resource shared private link %s points at", link.Name)
	}

	resourceID, hasID := genruntime.GetResourceID(resource)
	if !hasID {
		return extensions.PostReconcileCheckResultFailure(
			fmt.Sprintf("waiting for %s to be created in Azure", resource.GetName()),
		), nil
	}

	// If we have an approval in flight, pick it up and see if it's done
	if token, submitted := approvalResumeToken(link); submitted {
		done, err := resumeApproval(ctx, armClient, token)
		if err != nil {
			// A token that cannot be followed is worth less than a second approval, which Azure either
			// refuses as already decided or applies again harmlessly
			clearApprovalResumeToken(link)

			return extensions.PostReconcileCheckResult{},
				eris.Wrapf(err, "cannot approve the connection shared private link %s opened on %s", link.Name, resource.GetName())
		}

		if !done {
			// Stay short of ready so we're asked again, which is how the approval is seen through
			return extensions.PostReconcileCheckResultFailure(
				fmt.Sprintf("waiting for the private endpoint connection on %s to be approved", resource.GetName()),
			), nil
		}

		clearApprovalResumeToken(link)
	}

	connections, err := linkConnections(ctx, armClient, link, resourceID, resource.GetAPIVersion())
	if err != nil {
		// Reading the connections needs rights on the resource, and a credential without them leaves the link
		// the readiness it had before this extension existed rather than a failure it cannot act on
		if unauthorized(err) {
			log.V(Status).Info(
				"Not authorized to read the private endpoint connections a shared private link opened",
				"link", link.Name,
				"resource", resource.GetName(),
			)

			return next(ctx, obj, owner, resourceResolver, armClient, log, reconcilePolicies)
		}

		return extensions.PostReconcileCheckResult{}, err
	}

	// Azure opens the connection after the link itself is created, so it may not be there yet
	if len(connections) == 0 {
		return extensions.PostReconcileCheckResultFailure(
			fmt.Sprintf("waiting for Azure to open the private endpoint connection on %s", resource.GetName()),
		), nil
	}

	// Readiness is a fact about the connection, so it holds however the resource is managed and whichever of
	// these is this link's
	if allApproved(connections) {
		return next(ctx, obj, owner, resourceResolver, armClient, log, reconcilePolicies)
	}

	// A link name is unique only under its own watcher, so two watchers can open connections that look alike
	// here. Approving either would be a guess, so an ambiguous resource is left to its owner
	if len(connections) > 1 {
		return extensions.PostReconcileCheckResultFailure(
			fmt.Sprintf(
				"cannot tell which private endpoint connection on %s this link opened, so it has to be approved there",
				resource.GetName(),
			),
		), nil
	}

	connection := &connections[0]

	switch state := connection.state(); state {
	case connectionStateRejected, connectionStateDisconnected:
		// Someone acted on this connection deliberately, and approving it now would undo that
		return extensions.PostReconcileCheckResultFailure(
			fmt.Sprintf("the private endpoint connection on %s was %s", resource.GetName(), strings.ToLower(state)),
		), nil

	case connectionStatePending:
		// Handled below, this is the only state we approve from

	default:
		return extensions.PostReconcileCheckResultFailure(
			fmt.Sprintf("waiting for the private endpoint connection on %s to leave state %q", resource.GetName(), state),
		), nil
	}

	// Everything from here writes to the resource the link points at, so the checks that answer whether this
	// operator may do that belong to this path alone

	// This check still runs when the policy forbids modification, so the connection is reported as it stands
	if !reconcilePolicies.Effective.AllowsModify() {
		return extensions.PostReconcileCheckResultFailure(
			fmt.Sprintf(
				"the private endpoint connection on %s requires approval, which the reconcile policy on this link forbids",
				resource.GetName(),
			),
		), nil
	}

	// Nothing below holds for a resource another operator has claimed, and this comes before its policy is
	// resolved: none of that operator's policies or credentials is visible from here
	if reason, ok := foreignPrivateLinkResource(link, resource); ok {
		return extensions.PostReconcileCheckResultFailure(reason), nil
	}

	// Approving writes to the resource the link points at, so its own policy governs it
	allowed, err := modifyAllowed(reconcilePolicies, resource)
	if err != nil {
		return extensions.PostReconcileCheckResult{}, err
	}

	if !allowed {
		return extensions.PostReconcileCheckResultFailure(
			fmt.Sprintf("the private endpoint connection on %s requires approval", resource.GetName()),
		), nil
	}

	token, err := submitApproval(ctx, armClient, link, connection, resource, log)
	if err != nil {
		// Somebody decided the connection between it being read and being approved, so it is read again
		if alreadyDecided(err) {
			return extensions.PostReconcileCheckResultFailure(
				fmt.Sprintf("waiting to read the private endpoint connection on %s again", resource.GetName()),
			), nil
		}

		// Opening the connection needs no rights on the resource it is opened against, so a credential that
		// got this far may still have no say in approving it, leaving that to whoever owns the resource
		if unauthorized(err) {
			return extensions.PostReconcileCheckResultFailure(
				fmt.Sprintf(
					"the private endpoint connection on %s requires approval, which this operator's credential is not authorized to give",
					resource.GetName(),
				),
			), nil
		}

		return extensions.PostReconcileCheckResult{},
			eris.Wrapf(err, "cannot approve the connection shared private link %s opened on %s", link.Name, resource.GetName())
	}

	if token != "" {
		setApprovalResumeToken(link, token)
	}

	// Stay short of ready so we're asked again, which is how the approval is seen through
	return extensions.PostReconcileCheckResultFailure(
		fmt.Sprintf("approving the private endpoint connection on %s", resource.GetName()),
	), nil
}

func approvalResumeToken(link *databasewatcher.SharedPrivateLink) (string, bool) {
	token, ok := link.GetAnnotations()[ApprovalPollerResumeTokenAnnotation]
	return token, ok
}

func setApprovalResumeToken(link *databasewatcher.SharedPrivateLink, token string) {
	genruntime.AddAnnotation(link, ApprovalPollerResumeTokenAnnotation, token)
}

func clearApprovalResumeToken(link *databasewatcher.SharedPrivateLink) {
	genruntime.RemoveAnnotation(link, ApprovalPollerResumeTokenAnnotation)
}

// resumeApproval picks up an earlier approval, reporting whether it finished. A failed one is an error here.
func resumeApproval(ctx context.Context, armClient *genericarmclient.GenericClient, token string) (bool, error) {
	poller := armClient.ResumeCreatePoller(genericarmclient.CreatePollerID)

	err := poller.Resume(ctx, armClient, token)
	if err != nil {
		return false, err
	}

	return poller.Poller.Done(), nil
}

// foreignPrivateLinkResource reports why the link cannot approve a connection on the resource it points at,
// and is empty when it can. Approving writes to that resource with the link's credential, so the operator
// and the credential behind both have to be the same.
func foreignPrivateLinkResource(
	link *databasewatcher.SharedPrivateLink,
	resource genruntime.ARMMetaObject,
) (string, bool) {
	if differingOperator(link, resource) {
		return fmt.Sprintf(
			"cannot approve the private endpoint connection on %s; it is managed by the operator in %s while this link is managed by the operator in %s",
			resource.GetName(),
			describeOperator(resource),
			describeOperator(link),
		), true
	}

	if differingCredential(link, resource) {
		return fmt.Sprintf(
			"cannot approve the private endpoint connection on %s; it asks for %s while this link asks for %s",
			resource.GetName(),
			describeCredential(resource),
			describeCredential(link),
		), true
	}

	return "", false
}

// unauthorized reports whether ARM refused the request for want of permission.
func unauthorized(err error) bool {
	return armErrorCode(err) == authorizationFailedErrorCode
}

// alreadyDecided reports whether ARM refused the approval because the connection had left Pending.
func alreadyDecided(err error) bool {
	return armErrorCode(err) == connectionNotPendingErrorCode
}

func armErrorCode(err error) string {
	var cloudError *genericarmclient.CloudError
	if !eris.As(err, &cloudError) {
		return ""
	}

	return cloudError.Code()
}

// privateEndpointConnection is the part of a connection on the linked resource that its state is read from.
type privateEndpointConnection struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Properties struct {
		PrivateEndpoint struct {
			ID string `json:"id"`
		} `json:"privateEndpoint"`
		PrivateLinkServiceConnectionState struct {
			Status      string `json:"status"`
			Description string `json:"description"`
		} `json:"privateLinkServiceConnectionState"`
	} `json:"properties"`
}

func (connection *privateEndpointConnection) state() string {
	return canonicalState(connection.Properties.PrivateLinkServiceConnectionState.Status, connectionStates)
}

// privateEndpointConnectionApproval is the approval written back to a connection. Only the state is sent, so
// that nothing else Azure holds on the connection is overwritten.
type privateEndpointConnectionApproval struct {
	Properties struct {
		PrivateLinkServiceConnectionState struct {
			Status      string `json:"status"`
			Description string `json:"description"`
		} `json:"privateLinkServiceConnectionState"`
	} `json:"properties"`
}

// linkConnections returns the connections on the resource that carry the link's name, of which Azure opens one.
// The name is all that ties the two together: the link reports no connection of its own, and the private
// endpoint behind it lives in a subscription Microsoft owns.
func linkConnections(
	ctx context.Context,
	armClient *genericarmclient.GenericClient,
	link *databasewatcher.SharedPrivateLink,
	resourceID string,
	apiVersion string,
) ([]privateEndpointConnection, error) {
	container := resourceID + "/privateEndpointConnections"

	connections, err := genericarmclient.ListByContainerID[privateEndpointConnection](ctx, armClient, container, apiVersion)
	if err != nil {
		return nil, eris.Wrapf(err, "cannot list the private endpoint connections on %s", resourceID)
	}

	opened := make([]privateEndpointConnection, 0, 1)

	for i := range connections {
		if connectionOpenedBy(&connections[i], link.AzureName()) {
			opened = append(opened, connections[i])
		}
	}

	return opened, nil
}

// connectionOpenedBy reports whether a connection was opened by the named link. Azure names the managed
// private endpoint behind it after the link exactly, where the connection itself carries a GUID as well.
func connectionOpenedBy(connection *privateEndpointConnection, linkName string) bool {
	endpoint := connection.Properties.PrivateEndpoint.ID
	if endpoint == "" {
		return false
	}

	// ARM does not promise the casing it returns an ID in
	return strings.EqualFold(path.Base(endpoint), linkName)
}

// allApproved reports whether every connection is approved, which makes the link usable whichever is its own.
func allApproved(connections []privateEndpointConnection) bool {
	for i := range connections {
		if connections[i].state() != connectionStateApproved {
			return false
		}
	}

	return true
}

// submitApproval asks Azure to approve the connection, returning a token for the operation. Nothing waits here.
func submitApproval(
	ctx context.Context,
	armClient *genericarmclient.GenericClient,
	link *databasewatcher.SharedPrivateLink,
	connection *privateEndpointConnection,
	resource genruntime.ARMMetaObject,
	log logr.Logger,
) (string, error) {
	log.V(Status).Info(
		"Approving the private endpoint connection a shared private link opened",
		"connection", connection.Name,
		"resource", resource.GetName(),
	)

	var approval privateEndpointConnectionApproval
	approval.Properties.PrivateLinkServiceConnectionState.Status = connectionStateApproved
	approval.Properties.PrivateLinkServiceConnectionState.Description = approvalDescription(link)

	poller, err := armClient.BeginCreateOrUpdateByID(ctx, connection.ID, resource.GetAPIVersion(), approval)
	if err != nil {
		return "", err
	}

	// An approval Azure finished while answering has no operation to follow, and no token to ask for
	if poller.Poller.Done() {
		return "", nil
	}

	return poller.Poller.ResumeToken()
}

// approvalDescription is recorded on the connection, where it is the only account of what approved it.
func approvalDescription(link *databasewatcher.SharedPrivateLink) string {
	return fmt.Sprintf("Approved by Azure Service Operator for shared private link %s", link.AzureName())
}
