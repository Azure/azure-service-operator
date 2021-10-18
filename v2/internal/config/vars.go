// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config

import (
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	subscriptionIDVar   = "AZURE_SUBSCRIPTION_ID"
	targetNamespacesVar = "AZURE_TARGET_NAMESPACES"
	operatorModeVar     = "AZURE_OPERATOR_MODE"
	podNamespaceVar     = "POD_NAMESPACE"
)

// Values stores configuration values that are set for the operator.
type Values struct {
	// SubscriptionID is the Azure subscription the operator will use
	// for ARM communication.
	SubscriptionID string

	// PodNamespace is the namespace the operator pods are running in.
	PodNamespace string

	// OperatorMode determines whether the operator should run
	// watchers, webhooks or both.
	OperatorMode OperatorMode

	// TargetNamespaces lists the namespaces the operator will watch
	// for Azure resources (if the mode includes running watchers). If
	// it's empty the operator will watch all namespaces.
	TargetNamespaces []string

	// RequeueDelay overrides the default requeue delay
	// when the operator knows it needs to retry. This is only
	// used by testing code at the moment.
	RequeueDelay time.Duration
}

// ReadFromEnvironment loads configuration values from the AZURE_*
// environment variables.
func ReadFromEnvironment() (Values, error) {
	var result Values
	modeValue := os.Getenv(operatorModeVar)
	if modeValue == "" {
		result.OperatorMode = OperatorModeBoth
	} else {
		mode, err := ParseOperatorMode(modeValue)
		if err != nil {
			return Values{}, err
		}
		result.OperatorMode = mode
	}

	result.SubscriptionID = os.Getenv(subscriptionIDVar)
	result.PodNamespace = os.Getenv(podNamespaceVar)
	result.TargetNamespaces = parseTargetNamespaces(os.Getenv(targetNamespacesVar))
	// Not calling validate here to support using from tests where we
	// don't require consistent settings.
	return result, nil
}

// ReadAndValidate loads the configuration values and checks that
// they're consistent.
func ReadAndValidate() (Values, error) {
	result, err := ReadFromEnvironment()
	if err != nil {
		return Values{}, err
	}
	err = result.Validate()
	if err != nil {
		return Values{}, err
	}
	return result, nil
}

// Validate checks whether the configuration settings are consistent.
func (v Values) Validate() error {
	if v.SubscriptionID == "" {
		return errors.Errorf("missing value for %s", subscriptionIDVar)
	}
	if v.PodNamespace == "" {
		return errors.Errorf("missing value for %s", podNamespaceVar)
	}
	if !v.OperatorMode.IncludesWatchers() && len(v.TargetNamespaces) > 0 {
		return errors.Errorf("%s must include watchers to specify target namespaces", targetNamespacesVar)
	}
	return nil
}

// parseTargetNamespaces splits a comma-separated string into a slice
// of strings with spaces trimmed.
func parseTargetNamespaces(fromEnv string) []string {
	if len(strings.TrimSpace(fromEnv)) == 0 {
		return nil
	}
	items := strings.Split(fromEnv, ",")
	// Remove any whitespace used to separate items.
	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}
	return items
}
