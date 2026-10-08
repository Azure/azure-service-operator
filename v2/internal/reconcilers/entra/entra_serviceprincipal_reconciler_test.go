// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package entra

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-logr/logr"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-beta-sdk-go"
	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type servicePrincipalTestAdapter struct {
	abstractions.RequestAdapter
	baseURL string
	get     func(*abstractions.RequestInformation) (serialization.Parsable, error)
	delete  func(*abstractions.RequestInformation) error
}

func (a *servicePrincipalTestAdapter) GetBaseUrl() string {
	return a.baseURL
}

func (a *servicePrincipalTestAdapter) SetBaseUrl(url string) {
	a.baseURL = url
}

func (a *servicePrincipalTestAdapter) EnableBackingStore(store.BackingStoreFactory) {}

func (a *servicePrincipalTestAdapter) Send(
	_ context.Context,
	request *abstractions.RequestInformation,
	_ serialization.ParsableFactory,
	_ abstractions.ErrorMappings,
) (serialization.Parsable, error) {
	return a.get(request)
}

func (a *servicePrincipalTestAdapter) SendNoContent(
	_ context.Context,
	request *abstractions.RequestInformation,
	_ abstractions.ErrorMappings,
) error {
	return a.delete(request)
}

func (a *servicePrincipalTestAdapter) GetSerializationWriterFactory() serialization.SerializationWriterFactory {
	return jsonserialization.NewJsonSerializationWriterFactory()
}

type servicePrincipalTestConnection struct {
	client *msgraphsdk.GraphServiceClient
}

func (c *servicePrincipalTestConnection) Client() *msgraphsdk.GraphServiceClient {
	return c.client
}

func (c *servicePrincipalTestConnection) CredentialFrom() types.NamespacedName {
	return types.NamespacedName{}
}

func servicePrincipalTestFactory(adapter *servicePrincipalTestAdapter) EntraConnectionFactory {
	graph := msgraphsdk.NewGraphServiceClient(adapter)
	return func(context.Context, genruntime.EntraMetaObject) (Connection, error) {
		return &servicePrincipalTestConnection{client: graph}, nil
	}
}

func TestServicePrincipalTryAdoptByAppId(t *testing.T) {
	t.Parallel()

	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"

	cases := map[string]struct {
		results      []string
		wantID       string
		expectedErrs []string
	}{
		"found":     {results: []string{objectID}, wantID: objectID},
		"not found": {},
		"missing object ID": {
			results: []string{""},
			expectedErrs: []string{
				"service principal with appId",
				appID,
				"has no object ID",
			},
		},
		"ambiguous": {
			results: []string{objectID, "2251de93-281a-48c3-9842-e5e8619ad581"},
			expectedErrs: []string{
				"multiple service principals found with appId",
				appID,
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			adapter := &servicePrincipalTestAdapter{}
			adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
				uri, err := request.GetUri()
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(uri.Path).To(Equal("/beta/servicePrincipals"))
				g.Expect(uri.Query().Get("$filter")).To(Equal(fmt.Sprintf("appId eq '%s'", appID)))

				page := msgraphmodels.NewServicePrincipalCollectionResponse()
				principals := make([]msgraphmodels.ServicePrincipalable, 0, len(c.results))
				for _, id := range c.results {
					principal := msgraphmodels.NewServicePrincipal()
					principal.SetId(&id)
					principals = append(principals, principal)
				}
				page.SetValue(principals)
				return page, nil
			}
			reconciler := &EntraServicePrincipalReconciler{
				EntraClientFactory: servicePrincipalTestFactory(adapter),
			}
			obj := &asoentra.ServicePrincipal{
				Spec: asoentra.ServicePrincipalSpec{AppId: new(appID)},
			}

			id, err := reconciler.tryAdopt(context.Background(), obj, logr.Discard())
			if len(c.expectedErrs) > 0 {
				for _, expected := range c.expectedErrs {
					g.Expect(err).To(MatchError(ContainSubstring(expected)))
				}
				return
			}
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(id).To(Equal(c.wantID))
		})
	}
}

func TestServicePrincipalTryAdoptByDisplayName(t *testing.T) {
	t.Parallel()

	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const otherAppID = "00000003-0000-0000-c000-000000000000"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	const servicePrincipalName = "Azure's Cassandra Service"
	const displayNameFilter = "displayName eq 'Azure''s Cassandra Service'"

	cases := map[string]struct {
		appID       *string
		appMatches  []msgraphmodels.ServicePrincipalable
		nameMatches []msgraphmodels.ServicePrincipalable
		paginated   bool
		wantFilters []string
		wantID      string
		wantErr     string
	}{
		"name-only adoption": {
			nameMatches: []msgraphmodels.ServicePrincipalable{testServicePrincipal(objectID, appID)},
			wantFilters: []string{displayNameFilter},
			wantID:      objectID,
		},
		"ambiguous name prevents creation": {
			nameMatches: []msgraphmodels.ServicePrincipalable{
				testServicePrincipal(objectID, appID),
				testServicePrincipal("2251de93-281a-48c3-9842-e5e8619ad581", otherAppID),
			},
			wantFilters: []string{displayNameFilter},
			wantErr:     "multiple existing Entra service principals",
		},
		"ambiguous name across pages prevents creation": {
			nameMatches: []msgraphmodels.ServicePrincipalable{
				testServicePrincipal(objectID, appID),
				testServicePrincipal("2251de93-281a-48c3-9842-e5e8619ad581", otherAppID),
			},
			paginated:   true,
			wantFilters: []string{displayNameFilter, ""},
			wantErr:     "multiple existing Entra service principals",
		},
		"GUID match takes priority regardless of name": {
			appID:       new(appID),
			appMatches:  []msgraphmodels.ServicePrincipalable{testServicePrincipal(objectID, appID)},
			wantFilters: []string{"appId eq '" + appID + "'"},
			wantID:      objectID,
		},
		"name fallback matches GUID": {
			appID:       new(appID),
			nameMatches: []msgraphmodels.ServicePrincipalable{testServicePrincipal(objectID, appID)},
			wantFilters: []string{"appId eq '" + appID + "'", displayNameFilter},
			wantID:      objectID,
		},
		"name fallback conflicts with GUID": {
			appID:       new(appID),
			nameMatches: []msgraphmodels.ServicePrincipalable{testServicePrincipal(objectID, otherAppID)},
			wantFilters: []string{"appId eq '" + appID + "'", displayNameFilter},
			wantErr:     "expected \"" + appID + "\"",
		},
		"no name match": {
			wantFilters: []string{displayNameFilter},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			filters := make([]string, 0, len(c.wantFilters))
			adapter := &servicePrincipalTestAdapter{}
			adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
				g.Expect(request.Method).To(Equal(abstractions.GET))
				uri, err := request.GetUri()
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(uri.Path).To(Equal("/beta/servicePrincipals"))
				filter := uri.Query().Get("$filter")
				filters = append(filters, filter)
				page := msgraphmodels.NewServicePrincipalCollectionResponse()
				if c.paginated && uri.Query().Get("$skiptoken") == "next" {
					page.SetValue(c.nameMatches[1:])
				} else if c.appID != nil && filter == "appId eq '"+*c.appID+"'" {
					page.SetValue(c.appMatches)
				} else if c.paginated {
					page.SetValue(c.nameMatches[:1])
					page.SetOdataNextLink(new("https://graph.microsoft.com/beta/servicePrincipals?$skiptoken=next"))
				} else {
					page.SetValue(c.nameMatches)
				}
				return page, nil
			}
			reconciler := &EntraServicePrincipalReconciler{
				EntraClientFactory: servicePrincipalTestFactory(adapter),
			}
			obj := &asoentra.ServicePrincipal{
				Spec: asoentra.ServicePrincipalSpec{
					AppId:       c.appID,
					DisplayName: new(servicePrincipalName),
				},
			}

			id, err := reconciler.tryAdopt(context.Background(), obj, logr.Discard())
			g.Expect(filters).To(Equal(c.wantFilters))
			if c.wantErr != "" {
				g.Expect(err).To(MatchError(ContainSubstring(c.wantErr)))
			} else {
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(id).To(Equal(c.wantID))
			}
		})
	}
}

func TestServicePrincipalNameOnlyAdoptsAndUpdatesExisting(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	const name = "existing principal"
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	calls := 0
	patches := 0
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		calls++
		uri, err := request.GetUri()
		g.Expect(err).NotTo(HaveOccurred())
		principal := testServicePrincipal(objectID, appID)
		principal.SetDisplayName(new(name))
		if uri.Path == "/beta/servicePrincipals" {
			g.Expect(request.Method).To(Equal(abstractions.GET))
			g.Expect(uri.Query().Get("$filter")).To(Equal("displayName eq '" + name + "'"))
			page := msgraphmodels.NewServicePrincipalCollectionResponse()
			page.SetValue([]msgraphmodels.ServicePrincipalable{principal})
			return page, nil
		}
		g.Expect(uri.Path).To(Equal("/beta/servicePrincipals/" + objectID))
		if request.Method == abstractions.PATCH {
			patches++
		} else {
			g.Expect(request.Method).To(Equal(abstractions.GET))
		}
		return principal, nil
	}
	reconciler := &EntraServicePrincipalReconciler{EntraClientFactory: servicePrincipalTestFactory(adapter)}
	obj := &asoentra.ServicePrincipal{
		Spec: asoentra.ServicePrincipalSpec{DisplayName: new(name)},
	}

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(calls).To(Equal(3))
	g.Expect(patches).To(Equal(1))
	g.Expect(obj.Status.EntraID).To(Equal(new(objectID)))
	g.Expect(obj.Status.AppId).To(Equal(new(appID)))
}

func testServicePrincipal(objectID, appID string) msgraphmodels.ServicePrincipalable {
	principal := msgraphmodels.NewServicePrincipal()
	principal.SetId(&objectID)
	principal.SetAppId(&appID)
	return principal
}

func TestServicePrincipalNameOnlyMissingDoesNotCreate(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		g.Expect(request.Method).To(Equal(abstractions.GET))
		return msgraphmodels.NewServicePrincipalCollectionResponse(), nil
	}
	reconciler := &EntraServicePrincipalReconciler{EntraClientFactory: servicePrincipalTestFactory(adapter)}
	obj := &asoentra.ServicePrincipal{
		Spec: asoentra.ServicePrincipalSpec{DisplayName: new("new principal")},
	}

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).To(MatchError(ContainSubstring("cannot create service principal")))
}

func TestServicePrincipalCreationModes(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		mode      asoentra.CreationMode
		canAdopt  bool
		canCreate bool
	}{
		"AdoptOnly": {
			mode:      asoentra.AdoptOnly,
			canAdopt:  true,
			canCreate: false,
		},
		"AdoptOrCreate": {
			mode:      asoentra.AdoptOrCreate,
			canAdopt:  true,
			canCreate: true,
		},
		"AlwaysCreate": {
			mode:      asoentra.AlwaysCreate,
			canAdopt:  false,
			canCreate: true,
		},
	}

	reconciler := &EntraServicePrincipalReconciler{}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)

			typedObj := &asoentra.ServicePrincipal{
				Spec: asoentra.ServicePrincipalSpec{
					OperatorSpec: &asoentra.ServicePrincipalOperatorSpec{CreationMode: &c.mode},
				},
			}
			g.Expect(reconciler.canAdopt(typedObj)).To(Equal(c.canAdopt))
			g.Expect(reconciler.canCreate(typedObj)).To(Equal(c.canCreate))
		})
	}
}

func TestServicePrincipalCreateResolvesTenantObjectID(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	calls := 0
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		uri, err := request.GetUri()
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(uri.Path).To(Equal("/beta/servicePrincipals"))
		calls++

		switch request.Method {
		case abstractions.GET:
			g.Expect(uri.Query().Get("$filter")).To(Equal(fmt.Sprintf("appId eq '%s'", appID)))
			return msgraphmodels.NewServicePrincipalCollectionResponse(), nil
		case abstractions.POST:
			var body struct {
				AppID string `json:"appId"`
			}
			g.Expect(json.Unmarshal(request.Content, &body)).To(Succeed())
			g.Expect(body.AppID).To(Equal(appID))
			principal := msgraphmodels.NewServicePrincipal()
			principal.SetId(new(objectID))
			principal.SetAppId(new(appID))
			return principal, nil
		default:
			t.Fatalf("unexpected Graph request method %s", request.Method)
			return nil, nil
		}
	}
	reconciler := &EntraServicePrincipalReconciler{
		EntraClientFactory: servicePrincipalTestFactory(adapter),
	}
	obj := &asoentra.ServicePrincipal{
		Spec: asoentra.ServicePrincipalSpec{AppId: new(appID)},
	}

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(calls).To(Equal(2))
	id, ok := getEntraID(obj)
	g.Expect(ok).To(BeTrue())
	g.Expect(id).To(Equal(objectID))
	g.Expect(obj.Status.EntraID).To(Equal(new(objectID)))
	g.Expect(obj.Status.AppId).To(Equal(new(appID)))
}

func TestServicePrincipalAdoptOrCreateMutatesAndDeletesAdoptedPrincipal(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	patches := 0
	deletes := 0
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		uri, err := request.GetUri()
		g.Expect(err).NotTo(HaveOccurred())

		principal := msgraphmodels.NewServicePrincipal()
		principal.SetId(new(objectID))
		principal.SetAppId(new(appID))
		if uri.Path == "/beta/servicePrincipals" {
			g.Expect(request.Method).To(Equal(abstractions.GET))
			g.Expect(uri.Query().Get("$filter")).To(Equal(fmt.Sprintf("appId eq '%s'", appID)))
			page := msgraphmodels.NewServicePrincipalCollectionResponse()
			page.SetValue([]msgraphmodels.ServicePrincipalable{principal})
			return page, nil
		}
		g.Expect(uri.Path).To(Equal("/beta/servicePrincipals/" + objectID))
		if request.Method == abstractions.PATCH {
			patches++
		} else {
			g.Expect(request.Method).To(Equal(abstractions.GET))
		}
		return principal, nil
	}
	adapter.delete = func(request *abstractions.RequestInformation) error {
		uri, err := request.GetUri()
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(request.Method).To(Equal(abstractions.DELETE))
		g.Expect(uri.Path).To(Equal("/beta/servicePrincipals/" + objectID))
		deletes++
		return nil
	}
	reconciler := &EntraServicePrincipalReconciler{
		EntraClientFactory: servicePrincipalTestFactory(adapter),
	}
	obj := &asoentra.ServicePrincipal{
		Spec: asoentra.ServicePrincipalSpec{
			AppId:       new(appID),
			DisplayName: new("update an adopted principal"),
		},
	}

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(patches).To(Equal(1))
	id, ok := getEntraID(obj)
	g.Expect(ok).To(BeTrue())
	g.Expect(id).To(Equal(objectID))
	g.Expect(obj.Status.EntraID).To(Equal(new(objectID)))
	g.Expect(obj.Status.AppId).To(Equal(new(appID)))

	_, err = reconciler.Delete(context.Background(), logr.Discard(), nil, obj)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(deletes).To(Equal(1))
}

func TestServicePrincipalMissingAdoptionTargetDoesNotCreate(t *testing.T) {
	t.Parallel()
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"

	cases := map[string]struct {
		operatorSpec *asoentra.ServicePrincipalOperatorSpec
	}{
		"explicit AdoptOnly": {
			operatorSpec: &asoentra.ServicePrincipalOperatorSpec{
				CreationMode: new(asoentra.AdoptOnly),
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)

			adapter := &servicePrincipalTestAdapter{}
			adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
				g.Expect(request.Method).To(Equal(abstractions.GET))
				return msgraphmodels.NewServicePrincipalCollectionResponse(), nil
			}
			reconciler := &EntraServicePrincipalReconciler{
				EntraClientFactory: servicePrincipalTestFactory(adapter),
			}
			obj := &asoentra.ServicePrincipal{
				Spec: asoentra.ServicePrincipalSpec{AppId: new(appID), OperatorSpec: c.operatorSpec},
			}

			_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(ContainSubstring("not found for adoption"))
		})
	}
}

func TestServicePrincipalAdoptOnlyDoesNotMutateAnnotatedPrincipal(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	calls := 0
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		uri, err := request.GetUri()
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(uri.Path).To(Equal("/beta/servicePrincipals/" + objectID))
		g.Expect(request.Method).To(Equal(abstractions.GET))
		calls++
		principal := msgraphmodels.NewServicePrincipal()
		principal.SetId(new(objectID))
		principal.SetAppId(new(appID))
		return principal, nil
	}
	reconciler := &EntraServicePrincipalReconciler{
		EntraClientFactory: servicePrincipalTestFactory(adapter),
	}
	mode := asoentra.AdoptOnly
	obj := &asoentra.ServicePrincipal{
		Spec: asoentra.ServicePrincipalSpec{
			AppId:       new(appID),
			DisplayName: new("never update in AdoptOnly mode"),
			OperatorSpec: &asoentra.ServicePrincipalOperatorSpec{
				CreationMode: &mode,
			},
		},
	}
	setEntraID(obj, objectID)

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).NotTo(HaveOccurred())
	_, err = reconciler.Delete(context.Background(), logr.Discard(), nil, obj)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(calls).To(Equal(1))
}

func TestServicePrincipalDeleteAdoptedByDefault(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	deletes := 0
	adapter := &servicePrincipalTestAdapter{}
	adapter.delete = func(request *abstractions.RequestInformation) error {
		uri, err := request.GetUri()
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(request.Method).To(Equal(abstractions.DELETE))
		g.Expect(uri.Path).To(Equal("/beta/servicePrincipals/" + objectID))
		deletes++
		return nil
	}
	reconciler := &EntraServicePrincipalReconciler{
		EntraClientFactory: servicePrincipalTestFactory(adapter),
	}
	obj := &asoentra.ServicePrincipal{}
	setEntraID(obj, objectID)

	_, err := reconciler.Delete(context.Background(), logr.Discard(), nil, obj)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(deletes).To(Equal(1))
}

func TestServicePrincipalDisplayNameIsPatchedWhenSpecified(t *testing.T) {
	t.Parallel()
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"

	cases := map[string]struct {
		displayName   *string
		expectPatches int
	}{
		"without display name": {
			expectPatches: 0,
		},
		"with display name": {
			displayName:   new("requested name"),
			expectPatches: 1,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)

			patches := 0
			adapter := &servicePrincipalTestAdapter{}
			adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
				uri, err := request.GetUri()
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(uri.Path).To(Equal("/beta/servicePrincipals/" + objectID))
				if request.Method == abstractions.PATCH {
					patches++
					var body map[string]any
					g.Expect(json.Unmarshal(request.Content, &body)).To(Succeed())
					g.Expect(body).To(HaveKeyWithValue("displayName", "requested name"))
					g.Expect(body).NotTo(HaveKey("appId"))
				} else {
					g.Expect(request.Method).To(Equal(abstractions.GET))
				}
				principal := msgraphmodels.NewServicePrincipal()
				principal.SetId(new(objectID))
				principal.SetAppId(new(appID))
				return principal, nil
			}
			reconciler := &EntraServicePrincipalReconciler{
				EntraClientFactory: servicePrincipalTestFactory(adapter),
			}
			obj := &asoentra.ServicePrincipal{
				Spec: asoentra.ServicePrincipalSpec{AppId: new(appID), DisplayName: tc.displayName},
			}
			setEntraID(obj, objectID)
			_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(patches).To(Equal(tc.expectPatches))
		})
	}
}

func TestServicePrincipalDeleteCreated(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	calls := 0
	adapter := &servicePrincipalTestAdapter{}
	adapter.delete = func(request *abstractions.RequestInformation) error {
		uri, err := request.GetUri()
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(request.Method).To(Equal(abstractions.DELETE))
		g.Expect(uri.Path).To(Equal("/beta/servicePrincipals/" + objectID))
		calls++
		return nil
	}
	reconciler := &EntraServicePrincipalReconciler{
		EntraClientFactory: servicePrincipalTestFactory(adapter),
	}
	obj := &asoentra.ServicePrincipal{}
	setEntraID(obj, objectID)
	mode := asoentra.AdoptOrCreate
	obj.Spec.OperatorSpec = &asoentra.ServicePrincipalOperatorSpec{CreationMode: &mode}

	_, err := reconciler.Delete(context.Background(), logr.Discard(), nil, obj)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(calls).To(Equal(1))
}
