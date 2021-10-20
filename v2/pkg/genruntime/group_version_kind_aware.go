/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupVersionKindAware is implemented by resources that are aware of which version of the resource was originally
// specified. This allows us to interace with ARM using an API version specified by an end user.
type GroupVersionKindAware interface {
	// OriginalGVK returns the GroupVersionKind originally used to create the resource (regardless of any conversions)
	OriginalGVK() *schema.GroupVersionKind
}
