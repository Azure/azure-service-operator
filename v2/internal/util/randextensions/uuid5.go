// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package randextensions

import (
	"fmt"

	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Namespace is the ASOv2 UUIDv5 namespace UUID
var Namespace = uuid.Must(uuid.Parse("9a329043-7ad7-4b1c-812f-9c7a93d6392a"))

// MakeUniqueResourceScopedString generates a string that uniquely identifies a cluster resource. It includes the
// following distinguishing parts:
// * Group
// * Kind
// * Namespace
// * Name
func MakeUniqueResourceScopedString(gk schema.GroupKind, namespace string, name string) string {
	return fmt.Sprintf("%s/%s:%s/%s", gk.Group, gk.Kind, namespace, name)
}

// TODO: This method has a bug where it is called with an empty owner gk when the owner is an ARM ID.
// MakeUniqueOwnerScopedString generates a string that uniquely identifies a cluster resource.  It includes the
// following distinguishing parts:
// * Owner group
// * Owner kind
// * Owner name
// * Group
// * Kind
// * Namespace
// * Name
func MakeUniqueOwnerScopedString(ownerGK schema.GroupKind, ownerName string, gk schema.GroupKind, namespace string, name string) string {
	return fmt.Sprintf("%s/%s:%s/%s:%s/%s:%s/%s", ownerGK.Group, ownerGK.Kind, namespace, ownerName, gk.Kind, gk.Group, namespace, name)
}

// MakeUUIDName creates a stable UUID (v5) if the provided name is not already a UUID based on the specified
// uniqueString.
func MakeUUIDName(name string, uniqueString string) string {
	// If name is already a UUID we can just use that
	_, err := uuid.Parse(name)
	if err == nil {
		return name
	}

	return uuid.NewSHA1(Namespace, []byte(uniqueString)).String()
}
