/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
*/

package apimgmt

import (
	"context"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

// APIManager is an interface for API Management
type APIManager interface {
	CreateAPI(ctx context.Context, resourcegroup string, apiname string, properties interface{}) (string, error)
	DeleteAPI(ctx context.Context, resourcegroup string, apiname string) (string, error)
	GetAPI(ctx context.Context, resourcegroup string, servername string, database string) (string, error)
	resourcemanager.ARMClient
}
