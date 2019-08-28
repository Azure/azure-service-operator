// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package apimgmt

import (
	"context"
	"fmt"
	"testing"

	api "github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	"github.com/Azure/azure-service-operator/pkg/util"
)

// TestCreateOrUpdateAPI tests creating and deleting APIs
func TestCreateOrUpdateAPI(t *testing.T) {

	// setup config for pre-existing API Mgmt Svc
	ctx := context.Background()
	groupName := "heliumaksmsi"
	sdk := GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServiceName:       "heliumtestaksmsi",
		Email:             "test@microsoft.com",
		Name:              "test",
	}

	// only attempt to create the API if the api mgmt svc has been activated
	activated, err := sdk.IsAPIMgmtSvcActivated()
	if activated == true {
		apiName := generateName("apiname")
		//existingEndPoint := `{ "value": "https://heliumint.azurewebsites.net/swagger.json" }`
		existingEndPoint := "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/examples/v3.0/api-with-examples.yaml"
		protocols := []api.Protocol{api.ProtocolHTTP, api.ProtocolHTTPS}
		path := "/testpath"

		// initialize an apiproperties object
		apiProperties := APIProperties{
			// REQUIRED FOR ALL 3 - BLANK, OPENAPI, AND WADL
			// Format - Format of the Content in which the API is getting imported. Possible values include: 'WadlXML', 'WadlLinkJSON', 'SwaggerJSON', 'SwaggerLinkJSON', 'Wsdl', 'WsdlLink', 'Openapi', 'Openapijson', 'OpenapiLink'
			Format: api.OpenapiLink,
			// DisplayName - API name. Must be 1 to 300 characters long.
			DisplayName: &apiName,
			// REQUIRED FOR OPENAPI + WADL ONLY
			// Value - Content value when Importing an API (OpenAPI or WADL Specification?).
			Value: &existingEndPoint,
			// REQUIRED FOR OPENAPI ONLY
			// Protocols - Describes on which protocols the operations in this API can be invoked (HTTP, HTTPS, BOTH?)
			Protocols: &protocols,
			// OTHER FIELDS WE DISCOVERED WERE REQUIRED DURING TESTING
			// Path - Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
			Path: &path,

			APIID: generateName("apiid"),

			// TODO: set ifMatch
			IFMatch: "",
		}
		util.PrintAndLog("apiproperties object initialized")

		util.PrintAndLog("initiating CreateOrUpdateAPI")
		_, err = sdk.CreateOrUpdateAPI(apiProperties)
		if err != nil {
			util.PrintAndLog(fmt.Sprintf("cannot create api: %v", err))
			t.FailNow()
		} else {
			util.PrintAndLog("open api endpoint created")
		}

		// TODO: test delete of API
	}
}
