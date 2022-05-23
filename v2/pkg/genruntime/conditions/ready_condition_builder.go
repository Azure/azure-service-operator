/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package conditions

const (
	ReasonReconciling                     = "Reconciling"
	ReasonWaitingForOwner                 = "WaitingForOwner"
	ReasonDeleting                        = "Deleting"
	ReasonReconciliationFailedPermanently = "ReconciliationFailedPermanently"
	ReasonAzureResourceNotFound           = "AzureResourceNotFound"
	ReasonSecretNotFound                  = "SecretNotFound"
	ReasonReferenceNotFound               = "ReferenceNotFound"
	ReasonSecretWriteFailure              = "FailedWritingSecret"

	// ReasonFailed is a catch-all error code for when we don't have a more specific error classification
	ReasonFailed = "Failed"
)

func NewReadyConditionBuilder(builder PositiveConditionBuilderInterface) *ReadyConditionBuilder {
	return &ReadyConditionBuilder{
		builder: builder,
	}
}

type ReadyConditionBuilder struct {
	builder PositiveConditionBuilderInterface
}

func (b *ReadyConditionBuilder) ReadyCondition(severity ConditionSeverity, observedGeneration int64, reason string, message string) Condition {
	return b.builder.MakeFalseCondition(
		ConditionTypeReady,
		severity,
		observedGeneration,
		reason,
		message)
}

func (b *ReadyConditionBuilder) Reconciling(observedGeneration int64) Condition {
	return b.builder.MakeFalseCondition(
		ConditionTypeReady,
		ConditionSeverityInfo,
		observedGeneration,
		ReasonReconciling,
		"The resource is in the process of being reconciled by the operator")
}

func (b *ReadyConditionBuilder) Deleting(observedGeneration int64) Condition {
	return b.builder.MakeFalseCondition(
		ConditionTypeReady,
		ConditionSeverityInfo,
		observedGeneration,
		ReasonDeleting,
		"The resource is being deleted")
}

func (b *ReadyConditionBuilder) Succeeded(observedGeneration int64) Condition {
	return b.builder.MakeTrueCondition(ConditionTypeReady, observedGeneration)
}
