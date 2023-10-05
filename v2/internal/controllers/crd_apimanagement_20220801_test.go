/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apim "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_ApiManagement_20220801_CRUD(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("skipping test in live mode as it takes a very long time to provision an APIM service (1+h)")
	}

	tc := globalTestContext.ForTest(t)

	// We don't want to delete the resource group at the end of the test as APIM
	// takes a long time to provision. We'll clean it up manually
	rg := tc.NewTestResourceGroup()
	tc.CreateResourceAndWaitWithoutCleanup(rg)

	// There will be a New v2 SKU released 5/10/2023 which will have a much quicker start up
	// time. Move to that when it's available (BasicV2 or StandardV2 SKU)
	sku := apim.ApiManagementServiceSkuProperties{
		Capacity: to.Ptr(1),
		Name:     to.Ptr(apim.ApiManagementServiceSkuProperties_Name_Developer),
	}

	// Create an APIM instance. APIM has a soft delete feature; if you find that you
	// hit this problem add the `restore`` back in to resurrect it
	service := apim.Service{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("apim")),
		Spec: apim.Service_Spec{
			Location:       tc.AzureRegion,
			Owner:          testcommon.AsOwner(rg),
			PublisherEmail: to.Ptr("ASO@testing.com"),
			PublisherName:  to.Ptr("ASOTesting"),
			Sku:            &sku,
			// Restore:        to.Ptr(true),
		},
	}

	// When you are debugging, you can use this to create the APIM service once and not delete it
	tc.CreateResourceAndWaitWithoutCleanup(&service)
	// tc.CreateResourceAndWait(&service)

	tc.Expect(service.Status.Id).ToNot(BeNil())

	// Update the service to ensure that works
	tc.T.Log("updating tags on apim")
	old := service.DeepCopy()
	service.Spec.Tags = map[string]string{"scratchcard": "lanyard"}
	tc.PatchResourceAndWait(old, &service)
	tc.Expect(service.Status.Tags).To(HaveKey("scratchcard"))

	// Run sub-tests
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "APIM Subscription CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Subscription_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Backend CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Backend_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Named Value CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_NamedValue_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Policy Fragment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_PolicyFragment_CRUD(tc, &service)
			},
		},
		// testcommon.Subtest{
		// 	Name: "APIM Product CRUD",
		// 	Test: func(tc *testcommon.KubePerTestContext) {
		// 		APIM_Product_CRUD(tc, &service)
		// 	},
		// },
		testcommon.Subtest{
			Name: "APIM ProductAPis CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_ProductAPI_CRUD(tc, &service)
			},
		},
	)
}

func APIM_Subscription_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {

	// Create a subscription for all the apis
	subscription := apim.Subscription{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("sub1")),
		Spec: apim.Service_Subscription_Spec{
			DisplayName: to.Ptr("Subscription for all APIs"),
			Scope:       to.Ptr("/apis"),
			Owner:       testcommon.AsOwner(service),
		},
	}

	// Create a Subscription for a single API
	// subscription2 := apim.Subscription{
	// 	ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("sub2")),
	// 	Spec: apim.Service_Subscription_Spec{
	// 		DisplayName: to.Ptr("Subscription for single API"),
	// 		Scope:       to.Ptr("/apis/echo-api"),
	// 		Owner:       testcommon.AsOwner(service),
	// 	},
	// }

	tc.T.Log("creating apim subscriptions")
	tc.CreateResourceAndWait(&subscription)
	defer tc.DeleteResourceAndWait(&subscription)
	// tc.CreateResourceAndWait(&subscription2)
	// defer tc.DeleteResourceAndWait(&subscription2)

	tc.Expect(subscription.Status).ToNot(BeNil())

	tc.T.Log("cleaning up subscription")
}

func APIM_Backend_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {

	// Add a simple backend
	backend := apim.Backend{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("backend")),
		Spec: apim.Service_Backend_Spec{
			AzureName:   "test_backend",
			Description: to.Ptr("A Description about the backend"),
			Protocol:    to.Ptr(apim.BackendContractProperties_Protocol_Http),
			Url:         to.Ptr("https://www.bing.com"),
			Owner:       testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim backend")
	tc.CreateResourceAndWait(&backend)
	defer tc.DeleteResourceAndWait(&backend)

	tc.Expect(backend.Status).ToNot(BeNil())

	tc.T.Log("cleaning up backend")
}

func APIM_NamedValue_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {

	// Add a Plain Text Named Value
	namedValue := apim.NamedValue{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("namedvalue")),
		Spec: apim.Service_NamedValue_Spec{
			AzureName:   "test_namedvalue",
			DisplayName: to.Ptr("My_Key"),
			Value:       to.Ptr("It's value"),
			Secret:      to.Ptr(false),
			Owner:       testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim namedValue")
	tc.CreateResourceAndWait(&namedValue)
	defer tc.DeleteResourceAndWait(&namedValue)

	tc.Expect(namedValue.Status).ToNot(BeNil())

	tc.T.Log("cleaning up namedValue")
}

func APIM_PolicyFragment_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {

	// Add a simple Policy Fragment
	policyFragment := apim.PolicyFragment{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("policyfragment")),
		Spec: apim.Service_PolicyFragment_Spec{
			Description: to.Ptr("A Description about the policy fragment"),
			Value:       to.Ptr("<fragment></fragment>"),
			Owner:       testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim policy fragment")
	tc.CreateResourceAndWait(&policyFragment)
	defer tc.DeleteResourceAndWait(&policyFragment)

	tc.Expect(policyFragment.Status).ToNot(BeNil())

	tc.T.Log("cleaning up policyFragment")
}

func APIM_Product_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {

	// Now add a product
	product := apim.Product{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("cust1")),
		Spec: apim.Service_Product_Spec{
			Owner:                testcommon.AsOwner(service),
			DisplayName:          to.Ptr("Customer 1"),
			Description:          to.Ptr("A product for customer 1"),
			SubscriptionRequired: to.Ptr(false), // This creates a subscription which then makes the subscription test fail.
		},
	}

	tc.T.Log("creating apim product")
	tc.CreateResourceAndWait(&product)
	defer tc.DeleteResourceAndWait(&product)

	tc.Expect(product.Status).ToNot(BeNil())

	tc.T.Log("cleaning up product")
}

func APIM_ProductAPI_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {

	versionSet := apim.VersionSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vs")),
		Spec: apim.Service_ApiVersionSet_Spec{
			DisplayName:      to.Ptr("/apiVersionSets/account-api"),
			Description:      to.Ptr("A version set for the account api"),
			Owner:            testcommon.AsOwner(service),
			VersioningScheme: to.Ptr(apim.ApiVersionSetContractProperties_VersioningScheme_Segment),
		},
	}

	tc.T.Log("creating apim version set")
	tc.CreateResourceAndWait(&versionSet)

	// Add a simple Api
	api := apim.Api{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("api")),
		Spec: apim.Service_Api_Spec{
			APIVersion:             to.Ptr("2.0.0"),
			ApiRevision:            to.Ptr("v1"),
			ApiRevisionDescription: to.Ptr("First Revision"),
			ApiVersionDescription:  to.Ptr("Second Version"),
			ApiVersionSetId:        versionSet.Status.Id,
			Description:            to.Ptr("A Description about the api"),
			DisplayName:            to.Ptr("account-api"),
			Owner:                  testcommon.AsOwner(service),
			Path:                   to.Ptr("/account-api"),
			SubscriptionRequired:   to.Ptr(false),
			IsCurrent:              to.Ptr(true),
			Contact:                &apim.ApiContactInformation{
										Email: to.Ptr("test@test.com"),
										Name:  to.Ptr("Test"),
										Url:   to.Ptr("https://www.bing.com"),
			},
			//Format: 				to.Ptr(apim.ApiCreateOrUpdateProperties_Format_Swagger),
		
			Protocols: 				[]apim.ApiCreateOrUpdateProperties_Protocols{
										apim.ApiCreateOrUpdateProperties_Protocols_Https},
			//ServiceUrl: "https://www.bing.com",

			TermsOfServiceUrl: 		to.Ptr("https://www.bing.com/tos"),
			Type: 					to.Ptr(apim.ApiCreateOrUpdateProperties_Type_Http),
		},
	}

	tc.T.Log("creating apim api")
	tc.CreateResourceAndWait(&api)
	defer tc.DeleteResourceAndWait(&api)

	tc.Expect(api.Status).ToNot(BeNil())

	// Now add a product
	// product := apim.Product{
	// 	ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("custA")),
	// 	Spec: apim.Service_Product_Spec{
	// 		Owner: 	 testcommon.AsOwner(service),
	// 		DisplayName: to.Ptr("Customer A"),
	// 		Description: to.Ptr("A product for customer A"),
	// 		SubscriptionRequired: to.Ptr(true),
	// 	},
	// }

	// tc.T.Log("creating apim product")
	// tc.CreateResourceAndWait(&product)

	tc.T.Log("cleaning up api")
}