/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package conditions

import "fmt"

const (
	ReasonReconciling                     = "Reconciling"
	ReasonWaitingForOwner                 = "WaitingForOwner"
	ReasonDeleting                        = "Deleting"
	ReasonReconciliationFailedPermanently = "ReconciliationFailedPermanently"
	ReasonAzureResourceNotFound           = "AzureResourceNotFound"
)

func NewReadyConditionBuilder(builder PositiveConditionBuilderInterface) *ReadyConditionBuilder {
	return &ReadyConditionBuilder{
		builder: builder,
	}
}

type ReadyConditionBuilder struct {
	builder PositiveConditionBuilderInterface
}

func (b *ReadyConditionBuilder) Reconciling(observedGeneration int64) Condition {
	return b.builder.MakeFalseCondition(
		ConditionTypeReady,
		ConditionSeverityInfo,
		observedGeneration,
		ReasonReconciling,
		"The resource is in the process of being reconciled by the operator")
}

func (b *ReadyConditionBuilder) WaitingForOwner(observedGeneration int64, ownerDetails string) Condition {
	return b.builder.MakeFalseCondition(
		ConditionTypeReady,
		ConditionSeverityWarning,
		observedGeneration,
		ReasonWaitingForOwner,
		fmt.Sprintf("Owner %q cannot be found. Progress is blocked until the owner is created.", ownerDetails))
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

func (b *ReadyConditionBuilder) AzureResourceNotFound(observedGeneration int64, message string) Condition {
	return b.builder.MakeFalseCondition(
		ConditionTypeReady,
		ConditionSeverityWarning,
		observedGeneration,
		ReasonAzureResourceNotFound,
		message)
}
