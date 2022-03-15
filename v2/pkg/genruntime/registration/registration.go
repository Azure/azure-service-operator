/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package registration

import (
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// Index describes an index registration.
// See controller-runtime mgr.GetFieldIndexer().IndexField() for more details.
type Index struct {
	Key  string
	Func func(rawObj client.Object) []string
}

type EventHandlerFactory func(client client.Client, log logr.Logger) handler.EventHandler

// Watch describes a watch registration.
// See controller-runtime builder.Watches() for more details.
type Watch struct {
	Src              source.Source
	MakeEventHandler EventHandlerFactory
}

// StorageType describes a storage type to register, as well as any additional
// indexes or watches required.
// Obj is the object whose Kind should be registered as the storage type.
type StorageType struct {
	Obj     client.Object
	Indexes []Index
	Watches []Watch
}

// NewStorageType makes a new storage type for the specified object
func NewStorageType(obj client.Object) *StorageType {
	return &StorageType{
		Obj: obj,
	}
}
