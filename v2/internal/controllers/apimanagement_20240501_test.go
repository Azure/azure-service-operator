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

	apim "github.com/Azure/azure-service-operator/v2/api/apimanagement/v20240501"
	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20200202"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

func Test_ApiManagement_20240501_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.NewTestResourceGroup()

	tc.CreateResourceAndWait(rg)

	// The v2 SKU has a much quicker start up time.
	// Use BasicV2 or StandardV2 SKU for testing
	sku := apim.ApiManagementServiceSkuProperties{
		Capacity: to.Ptr(1),
		Name:     to.Ptr(apim.ApiManagementServiceSkuProperties_Name_BasicV2),
	}

	// Create an APIM instance. APIM has a soft delete feature; if you find that you
	// hit this problem add the `restore` back in to resurrect it
	service := apim.Service{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("apimanagementv3")),
		Spec: apim.Service_Spec{
			Location:       to.Ptr("eastus"), // Not supported in West US 2
			Owner:          testcommon.AsOwner(rg),
			PublisherEmail: to.Ptr("ASO@testing.com"),
			PublisherName:  to.Ptr("ASOTesting"),
			Sku:            &sku,
		},
	}

	tc.CreateResourceAndWait(&service)

	tc.Expect(service.Status.Id).ToNot(BeNil())

	// Update the service to ensure that works
	tc.T.Log("updating tags on apim")
	old := service.DeepCopy()
	service.Spec.Tags = map[string]string{"scratchcard": "lanyard"}
	tc.PatchResourceAndWait(old, &service)
	tc.Expect(service.Status.Tags).To(HaveKey("scratchcard"))

	authorizationProviderSecrets := createAuthorizationProviderSecrets20240501(tc, "authorizationprovider20240501")

	// Run subtests
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "APIM Subscription CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Subscription20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Backend CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Backend20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Named Value CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_NamedValue20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Policy Fragment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_PolicyFragment20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Policy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Policy20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Product CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Product20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Api CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Api20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Product Api CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Product_Api20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Product Policy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Product_Policy20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Api Policy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_ApiPolicy20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Group CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Group20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Logger CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_Logger20240501_CRUD(tc, rg, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM User CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_User20240501_CRUD(tc, &service)
			},
		},
		testcommon.Subtest{
			Name: "APIM Authorization Provider CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_AuthorizationProvider20240501_CRUD(tc, &service, &authorizationProviderSecrets)
			},
		},
		testcommon.Subtest{
			Name: "APIM Authorization CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_AuthorizationProviders_Authorization20240501_CRUD(tc, &service, &authorizationProviderSecrets)
			},
		},
		testcommon.Subtest{
			Name: "APIM Authorization Access Policy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_AuthorizationProviders_Authorizations_AccessPolicy20240501_CRUD(tc, rg, &service, &authorizationProviderSecrets)
			},
		},
	)
}

func APIM_Subscription20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	// Create a subscription for all the apis
	subscription := apim.Subscription{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("sub1")),
		Spec: apim.Subscription_Spec{
			DisplayName: to.Ptr("Subscription for all APIs"),
			Scope:       to.Ptr("/apis"),
			Owner:       testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim subscriptions")
	tc.CreateResourceAndWait(&subscription)
	tc.Expect(subscription.Status).ToNot(BeNil())
	tc.Expect(subscription.Status.Id).ToNot(BeNil())

	// Run sub-tests on subscription in sequence
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				Subscription20240501_SecretsWrittenToSameKubeSecret(tc, &subscription)
			},
		},
		testcommon.Subtest{
			Name: "SecretsWrittenToDifferentKubeSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				Subscription20240501_SecretsWrittenToDifferentKubeSecrets(tc, &subscription)
			},
		},
	)

	defer tc.DeleteResourceAndWait(&subscription)

	tc.T.Log("cleaning up subscription")
}

func APIM_Backend20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	// Add a simple backend
	backend := apim.Backend{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("backend")),
		Spec: apim.Backend_Spec{
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

func APIM_NamedValue20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	// Add a Plain Text Named Value
	namedValue := apim.NamedValue{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("namedvalue")),
		Spec: apim.NamedValue_Spec{
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

func APIM_Policy20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	// Add a simple Policy
	policy := apim.Policy{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("policy")),
		Spec: apim.Policy_Spec{
			Value: to.Ptr("<policies><inbound /><backend><forward-request /></backend><outbound /></policies>"),
			Owner: testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim policy")
	tc.CreateResourceAndWait(&policy)
	defer tc.DeleteResourceAndWait(&policy)

	tc.Expect(policy.Status).ToNot(BeNil())

	tc.T.Log("cleaning up policy")
}

func APIM_PolicyFragment20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	// Add a simple Policy Fragment
	policyFragment := apim.PolicyFragment{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("policyfragment")),
		Spec: apim.PolicyFragment_Spec{
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

func APIM_Product20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	productName := tc.Namer.GenerateName("cust1")
	// Now add a product
	product := apim.Product{
		ObjectMeta: tc.MakeObjectMetaWithName(productName),
		Spec: apim.Product_Spec{
			Owner:                testcommon.AsOwner(service),
			DisplayName:          to.Ptr("Customer 1"),
			Description:          to.Ptr("A product for customer 1"),
			SubscriptionRequired: to.Ptr(false), // This creates a subscription which then makes the subscription test fail.
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

func APIM_Product_Policy20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	productName := tc.Namer.GenerateName("product1")
	// Now add a product
	product := apim.Product{
		ObjectMeta: tc.MakeObjectMetaWithName(productName),
		Spec: apim.Product_Spec{
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
		Spec: apim.ProductPolicy_Spec{
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

func APIM_Product_Api20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	productName := tc.Namer.GenerateName("product2")
	product := apim.Product{
		ObjectMeta: tc.MakeObjectMetaWithName(productName),
		Spec: apim.Product_Spec{
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
		Spec: apim.ApiVersionSet_Spec{
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
		Spec: apim.Api_Spec{
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
				apim.ApiCreateOrUpdateProperties_Protocols_Https,
			},

			TermsOfServiceUrl: to.Ptr("https://www.bing.com/tos"),
			Type:              to.Ptr(apim.ApiCreateOrUpdateProperties_Type_Http),
		},
	}

	tc.T.Log("creating apim api to attach to product")
	tc.CreateResourceAndWait(&api)

	// Now link the display name of the api to the product
	productAPI := apim.ProductApi{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("productapi")),
		Spec: apim.ProductApi_Spec{
			Owner:     testcommon.AsOwner(&product),
			AzureName: api.Spec.AzureName,
		},
	}

	tc.T.Log("creating apim product api")
	tc.CreateResourceAndWait(&productAPI)

	tc.Expect(productAPI.Status).ToNot(BeNil())

	defer tc.DeleteResourceAndWait(&product)
	defer tc.DeleteResourceAndWait(&productAPI)

	tc.T.Log("cleaning up product")
}

func APIM_Api20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	versionSet := apim.ApiVersionSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vs")),
		Spec: apim.ApiVersionSet_Spec{
			DisplayName:      to.Ptr("/apiVersionSets/account-api"),
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
		Spec: apim.Api_Spec{
			APIVersion:             to.Ptr("2.0.0"),
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
				apim.ApiCreateOrUpdateProperties_Protocols_Https,
			},

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

func APIM_AuthorizationProvider20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object, secretsMap *genruntime.SecretMapReference) {
	authorizationProvider := apim.AuthorizationProvider{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("authorizationprovider")),
		Spec: apim.AuthorizationProvider_Spec{
			DisplayName:      to.Ptr("sampleauthcode"),
			IdentityProvider: to.Ptr("aad"),
			Oauth2: &apim.AuthorizationProviderOAuth2Settings{
				GrantTypes: &apim.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: secretsMap,
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

func APIM_AuthorizationProviders_Authorization20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object, secretsMap *genruntime.SecretMapReference) {
	authorizationProvider := apim.AuthorizationProvider{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("authorizationprovider")),
		Spec: apim.AuthorizationProvider_Spec{
			DisplayName:      to.Ptr("sampleauthcode"),
			IdentityProvider: to.Ptr("aad"),
			Oauth2: &apim.AuthorizationProviderOAuth2Settings{
				GrantTypes: &apim.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: secretsMap,
				},
			},
			Owner: testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim authorization provider")
	tc.CreateResourceAndWait(&authorizationProvider)

	authorization := apim.AuthorizationProvidersAuthorization{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("authorization")),
		Spec: apim.AuthorizationProvidersAuthorization_Spec{
			AuthorizationType: to.Ptr(apim.AuthorizationContractProperties_AuthorizationType_OAuth2),
			Oauth2GrantType:   to.Ptr(apim.AuthorizationContractProperties_Oauth2GrantType_AuthorizationCode),
			Owner:             testcommon.AsOwner(&authorizationProvider),
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

func APIM_AuthorizationProviders_Authorizations_AccessPolicy20240501_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, service client.Object, secretsMap *genruntime.SecretMapReference) {
	authorizationProvider := apim.AuthorizationProvider{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("authorizationprovider")),
		Spec: apim.AuthorizationProvider_Spec{
			DisplayName:      to.Ptr("sampleauthcode"),
			IdentityProvider: to.Ptr("aad"),
			Oauth2: &apim.AuthorizationProviderOAuth2Settings{
				GrantTypes: &apim.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: secretsMap,
				},
			},
			Owner: testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim authorization provider")
	tc.CreateResourceAndWait(&authorizationProvider)

	authorization := apim.AuthorizationProvidersAuthorization{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("authorization")),
		Spec: apim.AuthorizationProvidersAuthorization_Spec{
			AuthorizationType: to.Ptr(apim.AuthorizationContractProperties_AuthorizationType_OAuth2),
			Oauth2GrantType:   to.Ptr(apim.AuthorizationContractProperties_Oauth2GrantType_AuthorizationCode),
			Owner:             testcommon.AsOwner(&authorizationProvider),
		},
	}

	tc.T.Log("creating apim authorization")
	tc.CreateResourceAndWait(&authorization)

	configMapName := "my-configmap"
	principalIDKey := "principalId"
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
						Key:  principalIDKey,
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
		Spec: apim.AuthorizationProvidersAuthorizationsAccessPolicy_Spec{
			Owner: testcommon.AsOwner(&authorization),
			TenantIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  tenantIDKey,
			},
			ObjectIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIDKey,
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

func Subscription20240501_SecretsWrittenToSameKubeSecret(tc *testcommon.KubePerTestContext, subscription *apim.Subscription) {
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

func Subscription20240501_SecretsWrittenToDifferentKubeSecrets(tc *testcommon.KubePerTestContext, subscription *apim.Subscription) {
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

func APIM_ApiPolicy20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	versionSet := apim.ApiVersionSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vs")),
		Spec: apim.ApiVersionSet_Spec{
			DisplayName:      to.Ptr("api-policy-test-vs"),
			Description:      to.Ptr("A version set for the api policy test"),
			Owner:            testcommon.AsOwner(service),
			VersioningScheme: to.Ptr(apim.ApiVersionSetContractProperties_VersioningScheme_Segment),
		},
	}

	tc.T.Log("creating apim version set for api policy test")
	tc.CreateResourceAndWait(&versionSet)

	versionSetReference := genruntime.ResourceReference{
		ARMID: *versionSet.Status.Id,
	}

	// Add a simple Api to attach the policy to
	api := apim.Api{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("api")),
		Spec: apim.Api_Spec{
			APIVersion:             to.Ptr("1.0.0"),
			ApiRevision:            to.Ptr("v1"),
			ApiVersionSetReference: &versionSetReference,
			Description:            to.Ptr("An API for policy testing"),
			DisplayName:            to.Ptr("api-policy-test-api"),
			Owner:                  testcommon.AsOwner(service),
			Path:                   to.Ptr("/api-policy-test"),
			SubscriptionRequired:   to.Ptr(false),
			IsCurrent:              to.Ptr(true),
			Protocols: []apim.ApiCreateOrUpdateProperties_Protocols{
				apim.ApiCreateOrUpdateProperties_Protocols_Https,
			},
			Type: to.Ptr(apim.ApiCreateOrUpdateProperties_Type_Http),
		},
	}

	tc.T.Log("creating apim api for api policy test")
	tc.CreateResourceAndWait(&api)

	// Add an Api Policy
	apiPolicy := apim.ApiPolicy{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("apipolicy")),
		Spec: apim.ApiPolicy_Spec{
			Owner: testcommon.AsOwner(&api),
			Value: to.Ptr("<policies><inbound><set-variable name=\"asoTest\" value=\"ApiPolicy Value\" /></inbound><backend><forward-request /></backend><outbound /></policies>"),
		},
	}

	tc.T.Log("creating apim api policy")
	tc.CreateResourceAndWait(&apiPolicy)

	tc.Expect(apiPolicy.Status).ToNot(BeNil())
	tc.Expect(apiPolicy.Status.Id).ToNot(BeNil())

	defer tc.DeleteResourceAndWait(&apiPolicy)
	defer tc.DeleteResourceAndWait(&api)

	tc.T.Log("cleaning up api policy")
}

func APIM_Group20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	group := apim.Group{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("group")),
		Spec: apim.Group_Spec{
			DisplayName: to.Ptr("Sample Group"),
			Description: to.Ptr("A sample APIM group"),
			Owner:       testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim group")
	tc.CreateResourceAndWait(&group)
	defer tc.DeleteResourceAndWait(&group)

	tc.Expect(group.Status).ToNot(BeNil())
	tc.Expect(group.Status.Id).ToNot(BeNil())

	tc.T.Log("cleaning up group")
}

func APIM_Logger20240501_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, service client.Object) {
	// Create an Application Insights component to use as the logger target
	//nolint:gosec // This is not a secret, it is the name of a secret
	instrumentationKeySecret := "appinsights-secret"
	instrumentationKeyKey := "instrumentationKey"
	applicationType := insights.ApplicationInsightsComponentProperties_Application_Type_Other
	component := &insights.Component{
		ObjectMeta: tc.MakeObjectMeta("component"),
		Spec: insights.Component_Spec{
			Location:         tc.AzureRegion,
			Owner:            testcommon.AsOwner(rg),
			Application_Type: &applicationType,
			Kind:             to.Ptr("web"),
			OperatorSpec: &insights.ComponentOperatorSpec{
				SecretExpressions: []*core.DestinationExpression{
					{
						Name:  instrumentationKeySecret,
						Key:   instrumentationKeyKey,
						Value: "self.status.InstrumentationKey",
					},
				},
			},
		},
	}

	tc.T.Log("creating application insights component")
	tc.CreateResourceAndWait(component)

	tc.Expect(component.Status.Id).ToNot(BeNil())
	tc.Expect(component.Status.InstrumentationKey).ToNot(BeNil())

	// Create the APIM logger using Application Insights
	logger := apim.Logger{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("logger")),
		Spec: apim.Logger_Spec{
			Description: to.Ptr("An Application Insights logger"),
			LoggerType:  to.Ptr(apim.LoggerContractProperties_LoggerType_ApplicationInsights),
			Owner:       testcommon.AsOwner(service),
			Credentials: &genruntime.SecretMapReference{
				Name: instrumentationKeySecret,
			},
			ResourceReference: &genruntime.ResourceReference{
				ARMID: *component.Status.Id,
			},
		},
	}

	tc.T.Log("creating apim logger")
	tc.CreateResourceAndWait(&logger)
	defer tc.DeleteResourceAndWait(&logger)

	tc.Expect(logger.Status).ToNot(BeNil())
	tc.Expect(logger.Status.Id).ToNot(BeNil())

	// Run ApiDiagnostic subtest now that the logger exists
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "APIM Api Diagnostic CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				APIM_ApiDiagnostic20240501_CRUD(tc, service, &logger)
			},
		},
	)

	tc.T.Log("cleaning up logger")
}

func APIM_ApiDiagnostic20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object, logger *apim.Logger) {
	versionSet := apim.ApiVersionSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vs")),
		Spec: apim.ApiVersionSet_Spec{
			DisplayName:      to.Ptr("diagnostic-test-vs"),
			Description:      to.Ptr("A version set for the diagnostic test"),
			Owner:            testcommon.AsOwner(service),
			VersioningScheme: to.Ptr(apim.ApiVersionSetContractProperties_VersioningScheme_Segment),
		},
	}

	tc.T.Log("creating apim version set for diagnostic test")
	tc.CreateResourceAndWait(&versionSet)

	versionSetReference := genruntime.ResourceReference{
		ARMID: *versionSet.Status.Id,
	}

	// Create an Api to own the diagnostic
	api := apim.Api{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("api")),
		Spec: apim.Api_Spec{
			APIVersion:             to.Ptr("1.0.0"),
			ApiRevision:            to.Ptr("v1"),
			ApiVersionSetReference: &versionSetReference,
			Description:            to.Ptr("An API for diagnostic testing"),
			DisplayName:            to.Ptr("diagnostic-test-api"),
			Owner:                  testcommon.AsOwner(service),
			Path:                   to.Ptr("/diagnostic-test"),
			SubscriptionRequired:   to.Ptr(false),
			IsCurrent:              to.Ptr(true),
			Protocols: []apim.ApiCreateOrUpdateProperties_Protocols{
				apim.ApiCreateOrUpdateProperties_Protocols_Https,
			},
			Type: to.Ptr(apim.ApiCreateOrUpdateProperties_Type_Http),
		},
	}

	tc.T.Log("creating apim api for diagnostic test")
	tc.CreateResourceAndWait(&api)

	// Create the ApiDiagnostic referencing the logger
	diagnostic := apim.ApiDiagnostic{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("diag")),
		Spec: apim.ApiDiagnostic_Spec{
			AzureName: "applicationinsights",
			Owner:     testcommon.AsOwner(&api),
			LoggerReference: &genruntime.ResourceReference{
				ARMID: *logger.Status.Id,
			},
			Sampling: &apim.SamplingSettings{
				Percentage:   to.Ptr(100),
				SamplingType: to.Ptr(apim.SamplingSettings_SamplingType_Fixed),
			},
			Verbosity: to.Ptr(apim.DiagnosticContractProperties_Verbosity_Information),
		},
	}

	tc.T.Log("creating apim api diagnostic")
	tc.CreateResourceAndWait(&diagnostic)

	tc.Expect(diagnostic.Status).ToNot(BeNil())
	tc.Expect(diagnostic.Status.Id).ToNot(BeNil())

	defer tc.DeleteResourceAndWait(&diagnostic)
	defer tc.DeleteResourceAndWait(&api)

	tc.T.Log("cleaning up api diagnostic")
}

func APIM_User20240501_CRUD(tc *testcommon.KubePerTestContext, service client.Object) {
	user := apim.User{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("user")),
		Spec: apim.User_Spec{
			Email:     to.Ptr("asotest@example.com"),
			FirstName: to.Ptr("ASO"),
			LastName:  to.Ptr("Test"),
			Owner:     testcommon.AsOwner(service),
		},
	}

	tc.T.Log("creating apim user")
	tc.CreateResourceAndWait(&user)
	defer tc.DeleteResourceAndWait(&user)

	tc.Expect(user.Status).ToNot(BeNil())
	tc.Expect(user.Status.Id).ToNot(BeNil())

	tc.T.Log("cleaning up user")
}

func createAuthorizationProviderSecrets20240501(tc *testcommon.KubePerTestContext, name string) genruntime.SecretMapReference {
	clientID := tc.Namer.GeneratePasswordOfLength(10)
	clientSecret := tc.Namer.GeneratePasswordOfLength(10)

	clientIDKey := "clientId"
	clientSecretKey := "clientSecret"

	stringData := map[string]string{
		clientIDKey:     clientID,
		clientSecretKey: clientSecret,
		"ResourceUri":   "https://www.contoso.com",
	}

	data := make(map[string][]byte, len(stringData))
	for k, v := range stringData {
		data[k] = []byte(v)
	}

	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMeta(name),
		Data:       data,
		Type:       "Opaque",
	}

	tc.CreateResource(secret)
	secretRef := genruntime.SecretMapReference{Name: secret.Name}

	return secretRef
}
