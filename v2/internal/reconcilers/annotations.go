/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

// Annotation labels, used to store metadata about the state of the resource.
const (
	PollerResumeTokenAnnotation = "serviceoperator.azure.com/poller-resume-token"
	PollerResumeIDAnnotation    = "serviceoperator.azure.com/poller-resume-id"
	LatestReconciledGeneration  = "serviceoperator.azure.com/latest-reconciled-generation"

	// OperatorNamespaceAnnotation names the namespace of the operator that has claimed the resource.
	// Two resources carrying different ones are managed by different operators.
	OperatorNamespaceAnnotation = "serviceoperator.azure.com/operator-namespace"
)
