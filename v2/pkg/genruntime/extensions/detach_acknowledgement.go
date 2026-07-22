/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package extensions

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// DetachAcknowledgementRequirer is implemented by resources for which detaching
// (reconcile-policy=detach-on-delete) has special safety implications requiring
// explicit per-object acknowledgment before the operator will honor a detach.
//
// This is a safety/consent mechanism, not an authorization control - any principal
// with write access to the object can satisfy the acknowledgment. It exists to
// prevent *accidental* silent detachment (e.g. via a namespace-wide reconcile-policy
// applied for unrelated reasons) from bypassing a resource's delete-time safety logic,
// not to guard against a malicious actor.
type DetachAcknowledgementRequirer interface {
	// RequireDetachAcknowledgement returns an error if the object may NOT be
	// detached in its current state (e.g. missing acknowledgment annotation).
	// A non-nil error means: do not honor the detach; fall back to normal delete handling.
	RequireDetachAcknowledgement(obj genruntime.MetaObject) error
}
