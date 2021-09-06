package genruntime

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// HasOriginalGVK is implemented by resources that are aware of which version of the resource was originally specified.
// This allows us to interace with ARM using an API version specified by an end user.
type HasOriginalGVK interface {
	// OriginalGVK returns the GroupVersionKind originally used to create the resource (regardless of any conversions)
	OriginalGVK() *schema.GroupVersionKind
}
