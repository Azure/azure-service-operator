/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package client

import (
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func NewClient() (*v1.ApiextensionsV1Client, error) {
	cfg := config.GetConfigOrDie()
	client, err := v1.NewForConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create kubernetes client")
	}

	return client, nil
}
