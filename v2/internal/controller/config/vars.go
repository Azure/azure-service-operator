// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config

import (
	"os"
	"strings"
)

const (
	targetNamespacesVar = "AZURE_TARGET_NAMESPACES"
	operatorModeVar     = "AZURE_OPERATOR_MODE"
	podNamespaceVar     = "POD_NAMESPACE"
)

// GetOperatorMode reads the selected operator mode from the
// environment.
func GetOperatorMode() (OperatorMode, error) {
	value := os.Getenv(operatorModeVar)
	if value == "" {
		return OperatorModeBoth, nil
	}
	return ParseOperatorMode(value)
}

// GetPodNamespace reads the operator pod namespace from the
// environment.
func GetPodNamespace() string {
	return os.Getenv(podNamespaceVar)
}

// GetTargetNamespaces reads the list of target namespaces from the
// environment.
func GetTargetNamespaces() []string {
	return ParseTargetNamespaces(os.Getenv(targetNamespacesVar))
}

// ParseTargetNamespaces splits a comma-separated string into a slice
// of strings with spaces trimmed.
func ParseTargetNamespaces(fromEnv string) []string {
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
