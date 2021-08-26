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
)

func NewReadyConditionBuilder(builder PositiveConditionBuilderInterface) *ReadyConditionBuilder {
	return &ReadyConditionBuilder{
		builder: builder,
	}
}

type ReadyConditionBuilder struct {
	builder PositiveConditionBuilderInterface
}

func (b *ReadyConditionBuilder) Reconciling() Condition {
	return b.builder.MakeFalseCondition(
		ConditionTypeReady,
		ConditionSeverityInfo,
		ReasonReconciling,
		"The resource is in the process of being reconciled by the operator")
}

func (b *ReadyConditionBuilder) WaitingForOwner() Condition {
	return b.builder.MakeFalseCondition(
		ConditionTypeReady,
		ConditionSeverityWarning,
		ReasonWaitingForOwner,
		"The owner of this resource cannot be found in Kubernetes. Process is blocked until the owner is created.")
}

func (b *ReadyConditionBuilder) Deleting() Condition {
	return b.builder.MakeFalseCondition(
		ConditionTypeReady,
		ConditionSeverityInfo,
		ReasonDeleting,
		"The resource is being deleted")
}

func (b *ReadyConditionBuilder) Succeeded() Condition {
	return b.builder.MakeTrueCondition(ConditionTypeReady)
}
