/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"context"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

// TODO: ConfigMap stuff
// TODO: how is the handling going to work... I guess we need an optional interface in genruntime?
// TODO: can we use that same interface for secrets?
// TODO: ExportsKubernetesResources or something like that... yea I like this
// TODO: Problem is that secrets calls need to actually go to Azure, whereas other exports may not

type KubernetesExporter interface {
	ExportKubernetesResources(
		ctx context.Context,
		obj MetaObject, // TODO: The metaObject parameter is pointless unless we're reusing this as an extension too, since the type implementing this should be the storage type...
		armClient *genericarmclient.GenericClient,
		log logr.Logger) ([]client.Object, error)
}
