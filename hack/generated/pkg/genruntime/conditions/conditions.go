/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conditions

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Conditions []Condition

// FindIndexByType returns the index of the condition with the given ConditionType if it exists.
// If the Condition with the specified ConditionType is doesn't exist, the second boolean parameter is false.
func (c Conditions) FindIndexByType(conditionType ConditionType) (int, bool) {
	for i := range c {
		condition := c[i]
		if condition.Type == conditionType {
			return i, true
		}
	}

	return -1, false
}

// TODO: Hah, name...
type Conditioner interface {
	GetConditions() Conditions
	SetConditions(conditions Conditions)
}

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
	// may be able to retry and resolve the warning without any action from the user, but
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
	// ConditionTypeReady is a condition indicating if the resource is ready or not.
	// A ready resource is one that has been successfully provisioned to Azure according to the
	// resource spec. It has reached the goal state. This usually means that the resource is ready
	// to use, but the exact meaning of Ready may vary slightly from resource to resource. Resources with
	// caveats to Ready's meaning will call that out in the resource specific documentation.
	ConditionTypeReady = "Ready"
)

var _ fmt.Stringer = Condition{}

// Condition defines an extension to status (an observation) of a resource
// +kubebuilder:object:generate=true
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

	// LastTransitionTime is the last time the condition transitioned from one status to another.
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

// IsEquivalent returns true if this condition is equivalent to the passed in condition.
// Two conditions are equivalent if all of their fields EXCEPT LastTransitionTime are the same.
func (c Condition) IsEquivalent(other Condition) bool {
	return c.Type == other.Type &&
		c.Status == other.Status &&
		c.Severity == other.Severity &&
		c.Reason == other.Reason &&
		c.Message == other.Message
}

// Copy returns an independent copy of the Condition
func (c Condition) Copy() Condition {
	return c
}

// String returns a string representation of this condition
func (c Condition) String() string {
	return fmt.Sprintf(
		"Condition [%s], Status = %q, Severity = %q, Reason = %q, Message = %q, LastTransitionTime = %q",
		c.Type,
		c.Status,
		c.Severity,
		c.Reason,
		c.Message,
		c.LastTransitionTime)
}

// SetCondition sets the provided Condition on the Conditioner. The condition is only
// set if the new condition is in a different state than the existing condition of
// the same type.
func SetCondition(o Conditioner, new Condition) {
	if o == nil {
		return
	}

	conditions := o.GetConditions()
	i, exists := conditions.FindIndexByType(new.Type)
	if exists {
		if conditions[i].IsEquivalent(new) {
			// Nothing to do, the conditions are the same
			return
		}
		conditions[i] = new
	} else {
		conditions = append(conditions, new)
	}

	// TODO: do we sort conditions here? CAPI does.

	o.SetConditions(conditions)
}

// GetCondition gets the Condition with the specified type from the provided Conditioner.
// Returns the Condition and true if a Condition with the specified type is found, or an empty Condition
// and false if not.
func GetCondition(o Conditioner, conditionType ConditionType) (Condition, bool) {
	if o == nil {
		return Condition{}, false
	}

	conditions := o.GetConditions()
	i, exists := conditions.FindIndexByType(conditionType)
	if exists {
		return conditions[i], true
	}

	return Condition{}, false
}
