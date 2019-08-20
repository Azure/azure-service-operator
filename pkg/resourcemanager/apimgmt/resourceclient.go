// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package apimgmt

// ResourceClient contains the helper functions for interacting with APIs / API Mgmt Svc.
type ResourceClient interface {
	CreateAPIMgmtSvcImpl() (result bool, err error)
	CreateOrUpdateAPIImpl(apiid string, properties APIProperties, ifMatch string) (result bool, err error)
	DeleteAPIImpl(apiid string) (result bool, err error)
	DeleteAPIMgmtSvcImpl() (result bool, err error)
	IsAPIMgmtSvcActivatedImpl() (result bool, err error)
}
