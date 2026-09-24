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
	"github.com/onsi/gomega"
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

	for _, tc := range []struct {
		name    string
		results []string
		wantID  string
		wantErr bool
	}{
		{name: "found", results: []string{objectID}, wantID: objectID},
		{name: "not found"},
		{name: "missing object ID", results: []string{""}, wantErr: true},
		{name: "ambiguous", results: []string{objectID, "2251de93-281a-48c3-9842-e5e8619ad581"}, wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			g := gomega.NewWithT(t)
			adapter := &servicePrincipalTestAdapter{}
			adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
				uri, err := request.GetUri()
				g.Expect(err).NotTo(gomega.HaveOccurred())
				g.Expect(uri.Path).To(gomega.Equal("/beta/servicePrincipals"))
				g.Expect(uri.Query().Get("$filter")).To(gomega.Equal(fmt.Sprintf("appId eq '%s'", appID)))

				page := msgraphmodels.NewServicePrincipalCollectionResponse()
				principals := make([]msgraphmodels.ServicePrincipalable, 0, len(tc.results))
				for _, id := range tc.results {
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
				Spec: asoentra.ServicePrincipalSpec{AppId: stringPtr(appID)},
			}

			id, err := reconciler.tryAdopt(context.Background(), obj, logr.Discard())
			if tc.wantErr {
				g.Expect(err).To(gomega.HaveOccurred())
				return
			}
			g.Expect(err).NotTo(gomega.HaveOccurred())
			g.Expect(id).To(gomega.Equal(tc.wantID))
		})
	}
}

func TestServicePrincipalTryAdoptByDisplayName(t *testing.T) {
	t.Parallel()

	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const otherAppID = "00000003-0000-0000-c000-000000000000"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	const name = "Azure's Cassandra Service"

	for _, tc := range []struct {
		name        string
		appID       *string
		appMatches  []msgraphmodels.ServicePrincipalable
		nameMatches []msgraphmodels.ServicePrincipalable
		paginated   bool
		wantFilters []string
		wantID      string
		wantErr     string
	}{
		{
			name:        "name-only adoption",
			nameMatches: []msgraphmodels.ServicePrincipalable{testServicePrincipal(objectID, appID)},
			wantFilters: []string{"displayName eq 'Azure''s Cassandra Service'"},
			wantID:      objectID,
		},
		{
			name: "ambiguous name prevents creation",
			nameMatches: []msgraphmodels.ServicePrincipalable{
				testServicePrincipal(objectID, appID),
				testServicePrincipal("2251de93-281a-48c3-9842-e5e8619ad581", otherAppID),
			},
			wantFilters: []string{"displayName eq 'Azure''s Cassandra Service'"},
			wantErr:     "multiple existing Entra service principals",
		},
		{
			name: "ambiguous name across pages prevents creation",
			nameMatches: []msgraphmodels.ServicePrincipalable{
				testServicePrincipal(objectID, appID),
				testServicePrincipal("2251de93-281a-48c3-9842-e5e8619ad581", otherAppID),
			},
			paginated:   true,
			wantFilters: []string{"displayName eq 'Azure''s Cassandra Service'", ""},
			wantErr:     "multiple existing Entra service principals",
		},
		{
			name:        "GUID match takes priority regardless of name",
			appID:       stringPtr(appID),
			appMatches:  []msgraphmodels.ServicePrincipalable{testServicePrincipal(objectID, appID)},
			wantFilters: []string{"appId eq '" + appID + "'"},
			wantID:      objectID,
		},
		{
			name:        "name fallback matches GUID",
			appID:       stringPtr(appID),
			nameMatches: []msgraphmodels.ServicePrincipalable{testServicePrincipal(objectID, appID)},
			wantFilters: []string{"appId eq '" + appID + "'", "displayName eq 'Azure''s Cassandra Service'"},
			wantID:      objectID,
		},
		{
			name:        "name fallback conflicts with GUID",
			appID:       stringPtr(appID),
			nameMatches: []msgraphmodels.ServicePrincipalable{testServicePrincipal(objectID, otherAppID)},
			wantFilters: []string{"appId eq '" + appID + "'", "displayName eq 'Azure''s Cassandra Service'"},
			wantErr:     "expected \"" + appID + "\"",
		},
		{
			name:        "no name match",
			wantFilters: []string{"displayName eq 'Azure''s Cassandra Service'"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			g := gomega.NewWithT(t)
			filters := make([]string, 0, len(tc.wantFilters))
			adapter := &servicePrincipalTestAdapter{}
			adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
				g.Expect(request.Method).To(gomega.Equal(abstractions.GET))
				uri, err := request.GetUri()
				g.Expect(err).NotTo(gomega.HaveOccurred())
				g.Expect(uri.Path).To(gomega.Equal("/beta/servicePrincipals"))
				filter := uri.Query().Get("$filter")
				filters = append(filters, filter)
				page := msgraphmodels.NewServicePrincipalCollectionResponse()
				if tc.paginated && uri.Query().Get("$skiptoken") == "next" {
					page.SetValue(tc.nameMatches[1:])
				} else if tc.appID != nil && filter == "appId eq '"+*tc.appID+"'" {
					page.SetValue(tc.appMatches)
				} else if tc.paginated {
					page.SetValue(tc.nameMatches[:1])
					page.SetOdataNextLink(stringPtr("https://graph.microsoft.com/beta/servicePrincipals?$skiptoken=next"))
				} else {
					page.SetValue(tc.nameMatches)
				}
				return page, nil
			}
			reconciler := &EntraServicePrincipalReconciler{
				EntraClientFactory: servicePrincipalTestFactory(adapter),
			}
			obj := &asoentra.ServicePrincipal{
				Spec: asoentra.ServicePrincipalSpec{
					AppId:       tc.appID,
					DisplayName: stringPtr(name),
				},
			}

			id, err := reconciler.tryAdopt(context.Background(), obj, logr.Discard())
			g.Expect(filters).To(gomega.Equal(tc.wantFilters))
			if tc.wantErr != "" {
				g.Expect(err).To(gomega.MatchError(gomega.ContainSubstring(tc.wantErr)))
			} else {
				g.Expect(err).NotTo(gomega.HaveOccurred())
				g.Expect(id).To(gomega.Equal(tc.wantID))
			}
		})
	}
}

func TestServicePrincipalNameOnlyAdoptsExisting(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	const name = "existing principal"
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	calls := 0
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		calls++
		g.Expect(request.Method).To(gomega.Equal(abstractions.GET))
		uri, err := request.GetUri()
		g.Expect(err).NotTo(gomega.HaveOccurred())
		principal := testServicePrincipal(objectID, appID)
		principal.SetDisplayName(stringPtr(name))
		if uri.Path == "/beta/servicePrincipals" {
			g.Expect(uri.Query().Get("$filter")).To(gomega.Equal("displayName eq '" + name + "'"))
			page := msgraphmodels.NewServicePrincipalCollectionResponse()
			page.SetValue([]msgraphmodels.ServicePrincipalable{principal})
			return page, nil
		}
		g.Expect(uri.Path).To(gomega.Equal("/beta/servicePrincipals/" + objectID))
		return principal, nil
	}
	reconciler := &EntraServicePrincipalReconciler{EntraClientFactory: servicePrincipalTestFactory(adapter)}
	obj := &asoentra.ServicePrincipal{
		Spec: asoentra.ServicePrincipalSpec{DisplayName: stringPtr(name)},
	}

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(calls).To(gomega.Equal(2))
	g.Expect(obj.Status.EntraID).To(gomega.Equal(stringPtr(objectID)))
	g.Expect(obj.Status.AppId).To(gomega.Equal(stringPtr(appID)))
	g.Expect(obj.Annotations[servicePrincipalCreatedAnnotation]).To(gomega.Equal("false"))
}

func testServicePrincipal(objectID, appID string) msgraphmodels.ServicePrincipalable {
	principal := msgraphmodels.NewServicePrincipal()
	principal.SetId(&objectID)
	principal.SetAppId(&appID)
	return principal
}

func TestServicePrincipalNameOnlyMissingDoesNotCreate(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		g.Expect(request.Method).To(gomega.Equal(abstractions.GET))
		return msgraphmodels.NewServicePrincipalCollectionResponse(), nil
	}
	reconciler := &EntraServicePrincipalReconciler{EntraClientFactory: servicePrincipalTestFactory(adapter)}
	obj := &asoentra.ServicePrincipal{
		Spec: asoentra.ServicePrincipalSpec{DisplayName: stringPtr("new principal")},
	}

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).To(gomega.MatchError(gomega.ContainSubstring("cannot create service principal")))
}

func TestServicePrincipalCreationModes(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	reconciler := &EntraServicePrincipalReconciler{}
	obj := &asoentra.ServicePrincipal{}

	g.Expect(reconciler.canAdopt(obj)).To(gomega.BeTrue())
	g.Expect(reconciler.canCreate(obj)).To(gomega.BeTrue())
	obj.Spec.OperatorSpec = &asoentra.ServicePrincipalOperatorSpec{}
	g.Expect(reconciler.canAdopt(obj)).To(gomega.BeTrue())
	g.Expect(reconciler.canCreate(obj)).To(gomega.BeTrue())

	for _, tc := range []struct {
		mode      asoentra.CreationMode
		canAdopt  bool
		canCreate bool
	}{
		{mode: asoentra.AdoptOnly, canAdopt: true, canCreate: false},
		{mode: asoentra.AdoptOrCreate, canAdopt: true, canCreate: true},
		{mode: asoentra.AlwaysCreate, canAdopt: false, canCreate: true},
	} {
		t.Run(string(tc.mode), func(t *testing.T) {
			g := gomega.NewWithT(t)
			typedObj := &asoentra.ServicePrincipal{
				Spec: asoentra.ServicePrincipalSpec{
					OperatorSpec: &asoentra.ServicePrincipalOperatorSpec{CreationMode: &tc.mode},
				},
			}
			g.Expect(reconciler.canAdopt(typedObj)).To(gomega.Equal(tc.canAdopt))
			g.Expect(reconciler.canCreate(typedObj)).To(gomega.Equal(tc.canCreate))
		})
	}
}

func TestServicePrincipalCreateResolvesTenantObjectID(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	calls := 0
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		uri, err := request.GetUri()
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(uri.Path).To(gomega.Equal("/beta/servicePrincipals"))
		calls++

		switch request.Method {
		case abstractions.GET:
			g.Expect(uri.Query().Get("$filter")).To(gomega.Equal(fmt.Sprintf("appId eq '%s'", appID)))
			return msgraphmodels.NewServicePrincipalCollectionResponse(), nil
		case abstractions.POST:
			var body struct {
				AppID string `json:"appId"`
			}
			g.Expect(json.Unmarshal(request.Content, &body)).To(gomega.Succeed())
			g.Expect(body.AppID).To(gomega.Equal(appID))
			principal := msgraphmodels.NewServicePrincipal()
			principal.SetId(stringPtr(objectID))
			principal.SetAppId(stringPtr(appID))
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
		Spec: asoentra.ServicePrincipalSpec{AppId: stringPtr(appID)},
	}

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(calls).To(gomega.Equal(2))
	id, ok := getEntraID(obj)
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(id).To(gomega.Equal(objectID))
	g.Expect(obj.Status.EntraID).To(gomega.Equal(stringPtr(objectID)))
	g.Expect(obj.Status.AppId).To(gomega.Equal(stringPtr(appID)))
	g.Expect(obj.Annotations[servicePrincipalCreatedAnnotation]).To(gomega.Equal("true"))
}

func TestServicePrincipalAdoptResolvesTenantObjectID(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	calls := 0
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		uri, err := request.GetUri()
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(request.Method).To(gomega.Equal(abstractions.GET))
		calls++

		principal := msgraphmodels.NewServicePrincipal()
		principal.SetId(stringPtr(objectID))
		principal.SetAppId(stringPtr(appID))
		if uri.Path == "/beta/servicePrincipals" {
			g.Expect(uri.Query().Get("$filter")).To(gomega.Equal(fmt.Sprintf("appId eq '%s'", appID)))
			page := msgraphmodels.NewServicePrincipalCollectionResponse()
			page.SetValue([]msgraphmodels.ServicePrincipalable{principal})
			return page, nil
		}
		g.Expect(uri.Path).To(gomega.Equal("/beta/servicePrincipals/" + objectID))
		return principal, nil
	}
	reconciler := &EntraServicePrincipalReconciler{
		EntraClientFactory: servicePrincipalTestFactory(adapter),
	}
	obj := &asoentra.ServicePrincipal{
		Spec: asoentra.ServicePrincipalSpec{
			AppId:       stringPtr(appID),
			DisplayName: stringPtr("do not PATCH an adopted principal"),
		},
	}

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(calls).To(gomega.Equal(2))
	id, ok := getEntraID(obj)
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(id).To(gomega.Equal(objectID))
	g.Expect(obj.Status.EntraID).To(gomega.Equal(stringPtr(objectID)))
	g.Expect(obj.Status.AppId).To(gomega.Equal(stringPtr(appID)))
	g.Expect(obj.Annotations[servicePrincipalCreatedAnnotation]).To(gomega.Equal("false"))

	_, err = reconciler.Delete(context.Background(), logr.Discard(), nil, obj)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(calls).To(gomega.Equal(2))
}

func TestServicePrincipalMissingAdoptionTargetDoesNotCreate(t *testing.T) {
	t.Parallel()
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	mode := asoentra.AdoptOnly
	for _, tc := range []struct {
		name         string
		operatorSpec *asoentra.ServicePrincipalOperatorSpec
	}{
		{name: "explicit AdoptOnly", operatorSpec: &asoentra.ServicePrincipalOperatorSpec{CreationMode: &mode}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			adapter := &servicePrincipalTestAdapter{}
			adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
				g.Expect(request.Method).To(gomega.Equal(abstractions.GET))
				return msgraphmodels.NewServicePrincipalCollectionResponse(), nil
			}
			reconciler := &EntraServicePrincipalReconciler{
				EntraClientFactory: servicePrincipalTestFactory(adapter),
			}
			obj := &asoentra.ServicePrincipal{
				Spec: asoentra.ServicePrincipalSpec{AppId: stringPtr(appID), OperatorSpec: tc.operatorSpec},
			}

			_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
			g.Expect(err).To(gomega.HaveOccurred())
			g.Expect(err.Error()).To(gomega.ContainSubstring("not found for adoption"))
		})
	}
}

func TestServicePrincipalAdoptOnlyDoesNotMutateAnnotatedPrincipal(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	calls := 0
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		uri, err := request.GetUri()
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(uri.Path).To(gomega.Equal("/beta/servicePrincipals/" + objectID))
		g.Expect(request.Method).To(gomega.Equal(abstractions.GET))
		calls++
		principal := msgraphmodels.NewServicePrincipal()
		principal.SetId(stringPtr(objectID))
		principal.SetAppId(stringPtr(appID))
		return principal, nil
	}
	reconciler := &EntraServicePrincipalReconciler{
		EntraClientFactory: servicePrincipalTestFactory(adapter),
	}
	mode := asoentra.AdoptOnly
	obj := &asoentra.ServicePrincipal{
		Spec: asoentra.ServicePrincipalSpec{
			AppId:       stringPtr(appID),
			DisplayName: stringPtr("never update in AdoptOnly mode"),
			OperatorSpec: &asoentra.ServicePrincipalOperatorSpec{
				CreationMode: &mode,
			},
		},
	}
	setEntraID(obj, objectID)
	obj.Annotations[servicePrincipalCreatedAnnotation] = "true"

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	_, err = reconciler.Delete(context.Background(), logr.Discard(), nil, obj)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(calls).To(gomega.Equal(1))
}

func TestServicePrincipalDeleteAdoptedByDefault(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	obj := &asoentra.ServicePrincipal{}
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	setEntraID(obj, objectID)
	obj.Annotations[servicePrincipalCreatedAnnotation] = "false"

	persisted, err := json.Marshal(obj)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	reloaded := &asoentra.ServicePrincipal{}
	g.Expect(json.Unmarshal(persisted, reloaded)).To(gomega.Succeed())
	id, ok := getEntraID(reloaded)
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(id).To(gomega.Equal(objectID))
	g.Expect(reloaded.Annotations[servicePrincipalCreatedAnnotation]).To(gomega.Equal("false"))

	// A fresh reconciler has no in-memory adoption state; deletion must still be safe.
	reconciler := &EntraServicePrincipalReconciler{}
	_, err = reconciler.Delete(context.Background(), logr.Discard(), nil, reloaded)
	g.Expect(err).NotTo(gomega.HaveOccurred())
}

func TestServicePrincipalDisplayNameIsPatchedOnlyWhenSpecified(t *testing.T) {
	t.Parallel()
	const appID = "a232010e-820c-4083-83bb-3ace5fc29d0b"
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	for _, tc := range []struct {
		name          string
		displayName   *string
		created       bool
		expectPatches int
	}{
		{name: "created without display name", created: true},
		{name: "adopted with display name", displayName: stringPtr("requested name")},
		{name: "created with display name", displayName: stringPtr("requested name"), created: true, expectPatches: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			patches := 0
			adapter := &servicePrincipalTestAdapter{}
			adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
				uri, err := request.GetUri()
				g.Expect(err).NotTo(gomega.HaveOccurred())
				g.Expect(uri.Path).To(gomega.Equal("/beta/servicePrincipals/" + objectID))
				if request.Method == abstractions.PATCH {
					patches++
					var body map[string]any
					g.Expect(json.Unmarshal(request.Content, &body)).To(gomega.Succeed())
					g.Expect(body).To(gomega.HaveKeyWithValue("displayName", "requested name"))
					g.Expect(body).NotTo(gomega.HaveKey("appId"))
				} else {
					g.Expect(request.Method).To(gomega.Equal(abstractions.GET))
				}
				principal := msgraphmodels.NewServicePrincipal()
				principal.SetId(stringPtr(objectID))
				principal.SetAppId(stringPtr(appID))
				return principal, nil
			}
			reconciler := &EntraServicePrincipalReconciler{
				EntraClientFactory: servicePrincipalTestFactory(adapter),
			}
			obj := &asoentra.ServicePrincipal{
				Spec: asoentra.ServicePrincipalSpec{AppId: stringPtr(appID), DisplayName: tc.displayName},
			}
			setEntraID(obj, objectID)
			if tc.created {
				obj.Annotations[servicePrincipalCreatedAnnotation] = "true"
			}
			_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
			g.Expect(err).NotTo(gomega.HaveOccurred())
			g.Expect(patches).To(gomega.Equal(tc.expectPatches))
		})
	}
}

func TestServicePrincipalDeleteCreated(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	const objectID = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	calls := 0
	adapter := &servicePrincipalTestAdapter{}
	adapter.delete = func(request *abstractions.RequestInformation) error {
		uri, err := request.GetUri()
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(request.Method).To(gomega.Equal(abstractions.DELETE))
		g.Expect(uri.Path).To(gomega.Equal("/beta/servicePrincipals/" + objectID))
		calls++
		return nil
	}
	reconciler := &EntraServicePrincipalReconciler{
		EntraClientFactory: servicePrincipalTestFactory(adapter),
	}
	obj := &asoentra.ServicePrincipal{}
	setEntraID(obj, objectID)
	obj.Annotations[servicePrincipalCreatedAnnotation] = "true"
	mode := asoentra.AdoptOrCreate
	obj.Spec.OperatorSpec = &asoentra.ServicePrincipalOperatorSpec{CreationMode: &mode}

	_, err := reconciler.Delete(context.Background(), logr.Discard(), nil, obj)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(calls).To(gomega.Equal(1))
}
