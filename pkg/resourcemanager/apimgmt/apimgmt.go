// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package apimgmt

// CreateAPIMgmtSvc creates an instance of an API Management Service. Returns "true" if successful.
func CreateAPIMgmtSvc(provider ResourceClient) (result bool, err error) {
	return provider.CreateAPIMgmtSvcImpl()
}

// CreateOrUpdateAPI creates an API endpoint on an API Management Service. Returns "true" if successful.
func CreateOrUpdateAPI(provider ResourceClient, apiid string, properties APIProperties, ifMatch string) (result bool, err error) {
	return provider.CreateOrUpdateAPIImpl(apiid, properties, ifMatch)
}

// DeleteAPI deletes an API endpoint on an API Management Service. Returns "true" if successful.
func DeleteAPI(provider ResourceClient, apiid string) (result bool, err error) {
	return provider.DeleteAPIImpl(apiid)
}

// DeleteAPIMgmtSvc deletes an instance of an API Management Service. Returns "true" if successful.
func DeleteAPIMgmtSvc(provider ResourceClient) (result bool, err error) {
	return provider.DeleteAPIMgmtSvcImpl()
}

// IsAPIMgmtSvcActivated checks to make sure the service is availble, returns true if it is
func IsAPIMgmtSvcActivated(provider ResourceClient) (result bool, err error) {
	return provider.IsAPIMgmtSvcActivatedImpl()
}
