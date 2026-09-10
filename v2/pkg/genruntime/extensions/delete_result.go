/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package extensions

import (
	"fmt"
	"time"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

type DeleteResult struct {
	action         deleteResultType
	severity       conditions.ConditionSeverity
	reason         conditions.Reason
	message        string
	pollerResponse *genericarmclient.PollerResponse[genericarmclient.GenericDeleteResponse]
}

// DeleteCompleted is returned if deletion of the resource has completed successfully.
// No further action is needed.
func DeleteCompleted() DeleteResult {
	return DeleteResult{
		action: deleteResultTypeComplete,
	}
}

// BlockDelete is returned if deletion of the resource should be blocked for now, but can be retried later
// The deletion will automatically be retried after a short delay.
// message is an explanatory reason to show to the user via a warning condition on the resource.
func BlockDelete(message string, reason conditions.Reason) DeleteResult {
	return DeleteResult{
		action:   deleteResultTypeBlock,
		message:  message,
		severity: conditions.ConditionSeverityWarning,
		reason:   reason,
	}
}

// MonitorDelete is returned if deletion of a resource in Azure is in progress.
// pollerResponse is the response from the initial delete request,
// which will be used to monitor the status of the deletion operation.
func MonitorDelete(
	pollerResponse *genericarmclient.PollerResponse[genericarmclient.GenericDeleteResponse],
) DeleteResult {
	return DeleteResult{
		action:         deleteResultTypeMonitor,
		pollerResponse: pollerResponse,
	}
}

// Completed returns true if the deletion of the resource has completed successfully, false otherwise.
func (r DeleteResult) Completed() bool {
	return r.action == deleteResultTypeComplete
}

// BlockDeletion returns true if the deletion of the resource is currently blocked, false otherwise.
func (r DeleteResult) BlockDeletion() bool {
	return r.action == deleteResultTypeBlock
}

// MonitorDeletion returns true if deletion of the resource has started and needs to be monitored, false otherwise.
func (r DeleteResult) MonitorDeletion() bool {
	return r.action == deleteResultTypeMonitor
}

// Message returns the message associated with the DeleteResult, if any.
// This is typically used to provide an explanatory reason to show to the user via a warning condition on the resource.
func (r DeleteResult) Message() string {
	return r.message
}

// OperationID returns the operation ID associated monitoring the deletion operation, if any.
func (r DeleteResult) OperationID() (string, bool) {
	if r.pollerResponse == nil || r.pollerResponse.ID == "" {
		return "", false
	}

	return r.pollerResponse.ID, true
}

// OperationToken returns the operation token associated with monitoring the deletion operation, if any.
func (r DeleteResult) OperationToken() (string, error) {
	if r.pollerResponse == nil {
		return "", eris.New("no poller response available")
	}

	token, err := r.pollerResponse.Poller.ResumeToken()
	return token, eris.Wrap(err, "unable to create resume token")
}

func (r DeleteResult) RetryAfter() time.Duration {
	if r.pollerResponse == nil {
		return 0
	}

	return genericarmclient.GetRetryAfter(r.pollerResponse.RawResponse)
}

func (r DeleteResult) Reason() conditions.Reason {
	return r.reason
}

func (r DeleteResult) CreateConditionError() error {
	return conditions.NewReadyConditionImpactingError(
		eris.New(r.message),
		r.severity,
		r.reason,
	)
}

func (r DeleteResult) String() string {
	return fmt.Sprintf(
		"DeleteResult{action=%s, severity=%s, reason=%s, message=%s}",
		r.action,
		r.severity,
		r.reason,
		r.message,
	)
}

type deleteResultType string

const (
	deleteResultTypeBlock    deleteResultType = "Block"
	deleteResultTypeProceed  deleteResultType = "Proceed"
	deleteResultTypeComplete deleteResultType = "Complete"
	deleteResultTypeMonitor  deleteResultType = "Monitor"
)
