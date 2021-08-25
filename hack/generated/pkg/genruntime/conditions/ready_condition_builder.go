/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package conditions

import "fmt"

const (
	ReasonReconciling     = "Reconciling"
	ReasonWaitingForOwner = "WaitingForOwner"
	ReasonDeleting        = "Deleting"
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

func (b *ReadyConditionBuilder) WaitingForOwner(ownerDetails string) Condition {
	return b.builder.MakeFalseCondition(
		ConditionTypeReady,
		ConditionSeverityWarning,
		ReasonWaitingForOwner,
		fmt.Sprintf("Owner %q cannot be found. Progress is blocked until the owner is created.", ownerDetails))
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
