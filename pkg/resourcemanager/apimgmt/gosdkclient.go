// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package apimgmt

import "context"

// GoSDKClient is used to pass information to an implemetation of the APIMgmtImplemetation interface that wraps calls to the Go SDK for Azure.
type GoSDKClient struct {
	Ctx               context.Context
	ResourceGroupName string
	ServiceName       string
	Email             string
	Name              string
}
