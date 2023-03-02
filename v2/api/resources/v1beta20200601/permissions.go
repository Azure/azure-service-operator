/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1beta20200601

// TODO: it doesn't really matter where these are (as long as they're in 'apis', where is where we run controller-gen).
// These are the permissions required by the generic_controller. They're here because they can't go outside the 'apis'
// directory.

// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;patch;delete

// TODO: Somehow we need two different sets of permissions for the operator pod, one with create+update+read, and one
// TODO: with just read. The just read case is for when the user wants to manage the CRDs themselves to deny the operator
// TODO: pod the permissions. That mode probably can't be installed via Helm?

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch
