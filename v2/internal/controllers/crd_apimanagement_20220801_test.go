/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apim "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_ApiManagement_20220801_CRUD(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("skipping test in live mode as it takes a very long time to provision an APIM service (1+h)")
	}

	tc := globalTestContext.ForTest(t)
	rg := tc.NewTestResourceGroup()

	// APIM takes a long time to provision. When you are authoring tests in this file
	// I recommend you change this line to CreateResourceAndWaitWithoutCleanup. This way
	// apim will not be deleted until you have finished writing your tests. Also change
	// the code below when you create the &service
	tc.CreateResourceAndWait(rg)

	// There will be a New v2 SKU released 5/10/2023 which will have a much quicker start up
	// time. Move to that when it's available (BasicV2 or StandardV2 SKU)
	sku := apim.ApiManagementServiceSkuProperties{
		Capacity: to.Ptr(1),
		Name:     to.Ptr(apim.ApiManagementServiceSkuProperties_Name_Developer),
	}

	// Create an APIM instance. APIM has a soft delete feature; if you find that you
	// hit this problem add the `restore` back in to resurrect it
	service := apim.Service{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("apim")),
		Spec: apim.Service_Spec{
			Location:       tc.AzureRegion,
			Owner:          testcommon.AsOwner(rg),
			PublisherEmail: to.Ptr("ASO@testing.com"),
			PublisherName:  to.Ptr("ASOTesting"),
			Sku:            &sku,
		},
	}

	// APIM takes a long time to provision. When you are authoring tests in this file. I
	// recommend you change this line to CreateResourceAndWaitWithoutCleanup.
	tc.CreateResourceAndWait(&service)

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
		testcommon.Subtest{
			Name: "APIM Policy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Policy_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Product CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Product_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Api CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Api_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Product Api CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Product_Api_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Product Policy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Product_Policy_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Authorization Provider CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_AuthorizationProvider_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Authorization CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_AuthorizationProviders_Authorization_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Authorization Access Policy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_AuthorizationProviders_Authorizations_AccessPolicy_CRUD(tc, rg, &service)
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

	tc.T.Log("creating apim subscriptions")
	tc.CreateResourceAndWait(&subscription)
	tc.Expect(subscription.Status).ToNot(BeNil())
	tc.Expect(subscription.Status.Id).ToNot(BeNil())

	// The subscription will have been created and it would have generated it's own
	// PrimaryKey and SecondaryKey. Now let's see if we can set them from a Kubernetes secret.

	// There should be no secrets at this point
	secretList := &v1.SecretList{}
	tc.ListResources(secretList, client.InNamespace(tc.Namespace))
	tc.Expect(secretList.Items).To(HaveLen(0))

	// Run sub-tests on subscription in sequence
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				Subscription_SecretsWrittenToSameKubeSecret(tc, &subscription)
			},
		},
		testcommon.Subtest{
			Name: "SecretsWrittenToDifferentKubeSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				Subscription_SecretsWrittenToDifferentKubeSecrets(tc, &subscription)
			},
		},
	)

	defer tc.DeleteResourceAndWait(&subscription)

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

func APIM_Policy_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	// Add a simple Policy
	policy := apim.Policy{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("policy")),
		Spec: apim.Service_Policy_Spec{
			Value: to.Ptr("<policies><inbound><set-variable name=\"asoTest\" value=\"ProductPolicy Value\" /></inbound><backend><forward-request /></backend><outbound /></policies>"),
			Owner: testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim policy")
	tc.CreateResourceAndWait(&policy)
	defer tc.DeleteResourceAndWait(&policy)

	tc.Expect(policy.Status).ToNot(BeNil())

	tc.T.Log("cleaning up policy")
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

	productName := tc.Namer.GenerateName("cust1")
	// Now add a product
	product := apim.Product{
		ObjectMeta: tc.MakeObjectMetaWithName(productName),
		Spec: apim.Service_Product_Spec{
			Owner:                testcommon.AsOwner(service),
			DisplayName:          to.Ptr("Customer 1"),
			Description:          to.Ptr("A product for customer 1"),
			SubscriptionRequired: to.Ptr(false), // This creates a subscription anyway!
		},
	}

	tc.T.Log("creating apim product")
	tc.CreateResourceAndWait(&product)

	tc.Expect(product.Status).ToNot(BeNil())
	tc.Expect(product.Status.Id).ToNot(BeNil())

	// The product would have created a subscription automatically,
	// we have had to extend (products_extensions.go) to pass the following option
	// deleteSubscription = true

	defer tc.DeleteResourceAndWait(&product)

	tc.T.Log("cleaning up product")
}

func APIM_Product_Policy_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {

	productName := tc.Namer.GenerateName("product1")
	// Now add a product
	product := apim.Product{
		ObjectMeta: tc.MakeObjectMetaWithName(productName),
		Spec: apim.Service_Product_Spec{
			Owner:                testcommon.AsOwner(service),
			DisplayName:          to.Ptr("Product Policy Test"),
			Description:          to.Ptr("A product policy example"),
			SubscriptionRequired: to.Ptr(false), // This creates a subscription anyway!
		},
	}

	tc.T.Log("creating apim product to attach policy to")
	tc.CreateResourceAndWait(&product)

	tc.Expect(product.Status).ToNot(BeNil())
	tc.Expect(product.Status.Id).ToNot(BeNil())

	productPolicy := apim.ProductPolicy{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("productpolicy")),
		Spec: apim.Service_Products_Policy_Spec{
			Owner: testcommon.AsOwner(&product),
			Value: to.Ptr("<policies><inbound><set-variable name=\"asoTest\" value=\"ProductPolicy Value\" /></inbound><backend><forward-request /></backend><outbound /></policies>"),
		},
	}

	tc.T.Log("creating apim product policy")
	tc.CreateResourceAndWait(&productPolicy)

	tc.Expect(productPolicy.Status).ToNot(BeNil())
	tc.Expect(productPolicy.Status.Id).ToNot(BeNil())

	defer tc.DeleteResourceAndWait(&product)
	defer tc.DeleteResourceAndWait(&productPolicy)

	tc.T.Log("cleaning up product")
}

func APIM_Product_Api_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {

	productName := tc.Namer.GenerateName("product2")
	product := apim.Product{
		ObjectMeta: tc.MakeObjectMetaWithName(productName),
		Spec: apim.Service_Product_Spec{
			Owner:                testcommon.AsOwner(service),
			DisplayName:          to.Ptr("Product Api Test"),
			Description:          to.Ptr("A product Api example"),
			SubscriptionRequired: to.Ptr(false), // This creates a subscription anyway!
		},
	}

	tc.T.Log("creating apim product to attach api to")
	tc.CreateResourceAndWait(&product)

	tc.Expect(product.Status).ToNot(BeNil())
	tc.Expect(product.Status.Id).ToNot(BeNil())

	versionSet := apim.ApiVersionSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vs2")),
		Spec: apim.Service_ApiVersionSet_Spec{
			DisplayName:      to.Ptr("vs2"),
			Description:      to.Ptr("A version set for the account api"),
			Owner:            testcommon.AsOwner(service),
			VersioningScheme: to.Ptr(apim.ApiVersionSetContractProperties_VersioningScheme_Segment),
		},
	}

	tc.T.Log("creating apim version set")
	tc.CreateResourceAndWait(&versionSet)

	versionSetReference := genruntime.ResourceReference{
		ARMID: *versionSet.Status.Id,
	}

	// Add a simple Api
	api := apim.Api{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("api2")),
		Spec: apim.Service_Api_Spec{
			APIVersion:             to.Ptr("2.0.0"),
			ApiRevision:            to.Ptr("v1"),
			ApiRevisionDescription: to.Ptr("First Revision"),
			ApiVersionDescription:  to.Ptr("Second Version"),
			ApiVersionSetReference: &versionSetReference,
			Description:            to.Ptr("A Description about the api"),
			DisplayName:            to.Ptr("account-api2"),
			Owner:                  testcommon.AsOwner(service),
			Path:                   to.Ptr("/account-api2"),
			SubscriptionRequired:   to.Ptr(false),
			IsCurrent:              to.Ptr(true),
			Contact: &apim.ApiContactInformation{
				Email: to.Ptr("test@test.com"),
				Name:  to.Ptr("Test"),
				Url:   to.Ptr("https://www.bing.com"),
			},

			Protocols: []apim.ApiCreateOrUpdateProperties_Protocols{
				apim.ApiCreateOrUpdateProperties_Protocols_Https},

			TermsOfServiceUrl: to.Ptr("https://www.bing.com/tos"),
			Type:              to.Ptr(apim.ApiCreateOrUpdateProperties_Type_Http),
		},
	}

	tc.T.Log("creating apim api to attach to product")
	tc.CreateResourceAndWait(&api)

	// Now link the display name of the api to the product
	productApi := apim.ProductApi{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("productapi")),
		Spec: apim.Service_Products_Api_Spec{
			Owner:     testcommon.AsOwner(&product),
			AzureName: api.Spec.AzureName,
		},
	}

	tc.T.Log("creating apim product api")
	tc.CreateResourceAndWait(&productApi)

	tc.Expect(productApi.Status).ToNot(BeNil())

	defer tc.DeleteResourceAndWait(&product)
	defer tc.DeleteResourceAndWait(&productApi)

	tc.T.Log("cleaning up product")
}

func APIM_Api_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {

	versionSet := apim.ApiVersionSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vs")),
		Spec: apim.Service_ApiVersionSet_Spec{
			DisplayName:      to.Ptr("vs"),
			Description:      to.Ptr("A version set for the account api"),
			Owner:            testcommon.AsOwner(service),
			VersioningScheme: to.Ptr(apim.ApiVersionSetContractProperties_VersioningScheme_Segment),
		},
	}

	tc.T.Log("creating apim version set")
	tc.CreateResourceAndWait(&versionSet)

	versionSetReference := genruntime.ResourceReference{
		ARMID: *versionSet.Status.Id,
	}

	// Add a simple Api
	api := apim.Api{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("api")),
		Spec: apim.Service_Api_Spec{
			APIVersion:             to.Ptr("1.0.0"),
			ApiRevision:            to.Ptr("v1"),
			ApiRevisionDescription: to.Ptr("First Revision"),
			ApiVersionDescription:  to.Ptr("Second Version"),
			ApiVersionSetReference: &versionSetReference,
			Description:            to.Ptr("A Description about the api"),
			DisplayName:            to.Ptr("account-api"),
			Owner:                  testcommon.AsOwner(service),
			Path:                   to.Ptr("/account-api"),
			SubscriptionRequired:   to.Ptr(false),
			IsCurrent:              to.Ptr(true),
			Contact: &apim.ApiContactInformation{
				Email: to.Ptr("test@test.com"),
				Name:  to.Ptr("Test"),
				Url:   to.Ptr("https://www.bing.com"),
			},

			Protocols: []apim.ApiCreateOrUpdateProperties_Protocols{
				apim.ApiCreateOrUpdateProperties_Protocols_Https},

			TermsOfServiceUrl: to.Ptr("https://www.bing.com/tos"),
			Type:              to.Ptr(apim.ApiCreateOrUpdateProperties_Type_Http),
		},
	}

	tc.T.Log("creating apim api")
	tc.CreateResourceAndWait(&api)
	defer tc.DeleteResourceAndWait(&api)

	tc.Expect(api.Status).ToNot(BeNil())

	tc.T.Log("cleaning up api")
}

func APIM_AuthorizationProvider_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {

	authorizationProvider := apim.AuthorizationProvider{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("authorizationprovider")),
		Spec: apim.Service_AuthorizationProvider_Spec{
			DisplayName:      to.Ptr("sampleauthcode"),
			IdentityProvider: to.Ptr("aad"),
			Oauth2: &apim.AuthorizationProviderOAuth2Settings{
				GrantTypes: &apim.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: map[string]string{
						"ClientId":     "000",
						"ClientSecret": "*",
						"ResourceUri":  "https://www.contoso.com",
					},
				},
			},
			Owner: testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim authorization provider")
	tc.CreateResourceAndWait(&authorizationProvider)
	defer tc.DeleteResourceAndWait(&authorizationProvider)

	tc.Expect(authorizationProvider.Status).ToNot(BeNil())

	tc.T.Log("cleaning up authorizationProvider")
}

func APIM_AuthorizationProviders_Authorization_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {

	authorizationProvider := apim.AuthorizationProvider{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("authorizationprovider")),
		Spec: apim.Service_AuthorizationProvider_Spec{
			DisplayName:      to.Ptr("sampleauthcode"),
			IdentityProvider: to.Ptr("aad"),
			Oauth2: &apim.AuthorizationProviderOAuth2Settings{
				GrantTypes: &apim.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: map[string]string{
						"ClientId":     "000",
						"ClientSecret": "*",
						"ResourceUri":  "https://www.contoso.com",
					},
				},
			},
			Owner: testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim authorization provider")
	tc.CreateResourceAndWait(&authorizationProvider)

	authorization := apim.AuthorizationProvidersAuthorization{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("authorization")),
		Spec: apim.Service_AuthorizationProviders_Authorization_Spec{
			AuthorizationType: to.Ptr(apim.AuthorizationContractProperties_AuthorizationType_OAuth2),
			Error: &apim.AuthorizationError{
				Code:    to.Ptr("samplecode"),
				Message: to.Ptr("samplemessage"),
			},
			Oauth2GrantType: to.Ptr(apim.AuthorizationContractProperties_Oauth2GrantType_AuthorizationCode),
			Status:          to.Ptr("samplemessage"),
			Parameters:      map[string]string{},
			Owner:           testcommon.AsOwner(&authorizationProvider),
		},
	}

	tc.T.Log("creating apim authorization")
	tc.CreateResourceAndWait(&authorization)

	defer tc.DeleteResourceAndWait(&authorization)
	tc.Expect(authorization.Status).ToNot(BeNil())
	tc.T.Log("cleaning up authorization")

	defer tc.DeleteResourceAndWait(&authorizationProvider)
	tc.Expect(authorizationProvider.Status).ToNot(BeNil())
	tc.T.Log("cleaning up authorizationProvider")
}

func APIM_AuthorizationProviders_Authorizations_AccessPolicy_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, service client.Object) {

	authorizationProvider := apim.AuthorizationProvider{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("authorizationprovider")),
		Spec: apim.Service_AuthorizationProvider_Spec{
			DisplayName:      to.Ptr("sampleauthcode"),
			IdentityProvider: to.Ptr("aad"),
			Oauth2: &apim.AuthorizationProviderOAuth2Settings{
				GrantTypes: &apim.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: map[string]string{
						"ClientId":     "000",
						"ClientSecret": "*",
						"ResourceUri":  "https://www.contoso.com",
					},
				},
			},
			Owner: testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim authorization provider")
	tc.CreateResourceAndWait(&authorizationProvider)

	authorization := apim.AuthorizationProvidersAuthorization{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("authorization")),
		Spec: apim.Service_AuthorizationProviders_Authorization_Spec{
			AuthorizationType: to.Ptr(apim.AuthorizationContractProperties_AuthorizationType_OAuth2),
			Error: &apim.AuthorizationError{
				Code:    to.Ptr("samplecode"),
				Message: to.Ptr("samplemessage"),
			},
			Oauth2GrantType: to.Ptr(apim.AuthorizationContractProperties_Oauth2GrantType_AuthorizationCode),
			Status:          to.Ptr("samplemessage"),
			Parameters:      map[string]string{},
			Owner:           testcommon.AsOwner(&authorizationProvider),
		},
	}

	tc.T.Log("creating apim authorization")
	tc.CreateResourceAndWait(&authorization)

	configMapName := "my-configmap"
	principalIdKey := "principalId"
	tenantIDKey := "tenantId"

	// Create a managed identity to use as the AAD administrator
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  principalIdKey,
					},
					TenantId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  tenantIDKey,
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(mi)
	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())

	accessPolicy := apim.AuthorizationProvidersAuthorizationsAccessPolicy{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("authorization-access-policy")),
		Spec: apim.Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec{
			Owner: testcommon.AsOwner(&authorization),
			TenantIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  tenantIDKey,
			},
			ObjectIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
		},
	}

	tc.T.Log("creating apim authorization accessPolicy")
	tc.CreateResourceAndWait(&accessPolicy)

	defer tc.DeleteResourceAndWait(&accessPolicy)
	tc.Expect(accessPolicy.Status).ToNot(BeNil())
	tc.T.Log("cleaning up authorization accessPolicy")

	defer tc.DeleteResourceAndWait(&authorization)
	tc.Expect(authorization.Status).ToNot(BeNil())
	tc.T.Log("cleaning up authorization")

	defer tc.DeleteResourceAndWait(&authorizationProvider)
	tc.Expect(authorizationProvider.Status).ToNot(BeNil())
	tc.T.Log("cleaning up authorizationProvider")
}

func Subscription_SecretsWrittenToSameKubeSecret(tc *testcommon.KubePerTestContext, subscription *apim.Subscription) {
	old := subscription.DeepCopy()
	subscriptionSecret := "storagekeys"
	subscription.Spec.OperatorSpec = &apim.SubscriptionOperatorSpec{
		Secrets: &apim.SubscriptionOperatorSecrets{
			PrimaryKey:   &genruntime.SecretDestination{Name: subscriptionSecret, Key: "primary"},
			SecondaryKey: &genruntime.SecretDestination{Name: subscriptionSecret, Key: "secondary"},
		},
	}
	tc.PatchResourceAndWait(old, subscription)

	tc.ExpectSecretHasKeys(subscriptionSecret, "primary", "secondary")
}

func Subscription_SecretsWrittenToDifferentKubeSecrets(tc *testcommon.KubePerTestContext, subscription *apim.Subscription) {
	old := subscription.DeepCopy()
	primaryKeySecret := "secret1"
	secondaryKeySecret := "secret2"

	subscription.Spec.OperatorSpec = &apim.SubscriptionOperatorSpec{
		Secrets: &apim.SubscriptionOperatorSecrets{
			PrimaryKey: &genruntime.SecretDestination{
				Name: primaryKeySecret,
				Key:  "primarymasterkey",
			},
			SecondaryKey: &genruntime.SecretDestination{
				Name: secondaryKeySecret,
				Key:  "secondarymasterkey",
			},
		},
	}
	tc.PatchResourceAndWait(old, subscription)

	tc.ExpectSecretHasKeys(primaryKeySecret, "primarymasterkey")
	tc.ExpectSecretHasKeys(secondaryKeySecret, "secondarymasterkey")
}
