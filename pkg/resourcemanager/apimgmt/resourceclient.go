// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package apimgmt

import (
	api "github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	"github.com/Azure/go-autorest/autorest"
)

// ResourceClient contains the helper functions for interacting with APIs / API Mgmt Svc.
type ResourceClient interface {
	CreateAPIMgmtSvc() (result api.ServiceResource, err error)
	CreateOrUpdateAPI(properties APIProperties) (result api.APIContract, err error)
	DeleteAPI(apiid string, ifMatch string) (result autorest.Response, err error)
	DeleteAPIMgmtSvc() (result api.ServiceResource, err error)
	IsAPIMgmtSvcActivated() (result bool, err error)
}
