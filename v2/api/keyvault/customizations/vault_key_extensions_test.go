/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20230701/storage"
	keys "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

const (
	testKeyName = "my-key"
	testVaultID = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg" +
		"/providers/Microsoft.KeyVault/vaults/myvault"
	testKeyID = testVaultID + "/keys/" + testKeyName
)

// ---------------------------------------------------------------------------------------------
// Fake Key Vault
//
// The extension talks to Key Vault through the real azkeys client, pointed at an httptest server
// via the key's status.keyUri (or, for the ARM fallback, via the ARM client's endpoint). The fake
// answers unauthenticated requests with the 401 challenge Key Vault issues, so the SDK's challenge
// handshake runs for real and every request it records arrives authenticated and with its body.
// ---------------------------------------------------------------------------------------------

type fakeResponse struct {
	status int
	body   string
}

type fakeKeyVault struct {
	server *httptest.Server
	mu     sync.Mutex
	routes map[string]fakeResponse
	calls  []string
	bodies map[string][]string
}

// newFakeKeyVault serves canned responses keyed by "METHOD /path" (query string and any trailing
// slash ignored). Unrouted requests get a 500 so an unexpected call fails its test loudly.
func newFakeKeyVault() *fakeKeyVault {
	f := &fakeKeyVault{routes: map[string]fakeResponse{}, bodies: map[string][]string{}}
	f.server = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Method + " " + strings.TrimSuffix(r.URL.Path, "/")

		if r.Header.Get("Authorization") == "" {
			// Elicit the SDK's challenge handshake. The SDK insists that the challenge's resource host be a
			// parent domain of the vault host; for the loopback address the closest thing is its dotted
			// suffix, which satisfies that check without disabling verification in production code.
			// (Should httptest ever bind an undotted host such as [::1], the check fails and so do these
			// tests, loudly.)
			resource := "https://" + r.Host
			if i := strings.Index(r.Host, "."); i >= 0 {
				resource = "https://" + r.Host[i+1:]
			}
			w.Header().Set("WWW-Authenticate", fmt.Sprintf(
				`Bearer authorization="https://login.microsoftonline.com/00000000-0000-0000-0000-000000000000", resource=%q`,
				resource,
			))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		body, _ := io.ReadAll(r.Body)

		f.mu.Lock()
		f.calls = append(f.calls, key)
		f.bodies[key] = append(f.bodies[key], string(body))
		resp, ok := f.routes[key]
		f.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		if !ok {
			// The offending route is visible through Calls(), so the body needn't echo it
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"error":{"code":"Unrouted","message":"unexpected request to the fake Key Vault"}}`))
			return
		}

		w.WriteHeader(resp.status)
		_, _ = w.Write([]byte(resp.body))
	}))

	return f
}

func (f *fakeKeyVault) route(method string, path string, status int, body string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.routes[method+" "+path] = fakeResponse{status: status, body: body}
}

func (f *fakeKeyVault) Calls() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string(nil), f.calls...)
}

// LastBody returns the body of the most recent request to the given "METHOD /path".
func (f *fakeKeyVault) LastBody(key string) string {
	f.mu.Lock()
	defer f.mu.Unlock()

	bodies := f.bodies[key]
	if len(bodies) == 0 {
		return ""
	}

	return bodies[len(bodies)-1]
}

func (f *fakeKeyVault) Close() {
	f.server.Close()
}

func newARMClient(g *WithT, server *httptest.Server) *genericarmclient.GenericClient {
	cfg := cloud.Configuration{
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Endpoint: server.URL,
				Audience: cloud.AzurePublic.Services[cloud.ResourceManager].Audience,
			},
		},
	}

	armClient, err := genericarmclient.NewGenericClient(cfg, creds.MockTokenCredential{}, &genericarmclient.GenericClientOptions{
		HTTPClient: server.Client(),
		Metrics:    asometrics.NewARMClientMetrics(),
	})
	g.Expect(err).ToNot(HaveOccurred())

	return armClient
}

const (
	keyPath            = "/keys/" + testKeyName
	deletedKeyPath     = "/deletedkeys/" + testKeyName
	rotationPolicyPath = keyPath + "/rotationpolicy"

	notFoundBody  = `{"error":{"code":"KeyNotFound","message":"A key with (name/id) my-key was not found in this key vault."}}`
	forbiddenBody = `{"error":{"code":"Forbidden","message":"The user, group or application does not have keys get permission on key vault"}}`
	conflictBody  = `{"error":{"code":"Conflict","message":"Key is currently being recovered."}}`
)

func rsaModulus() string {
	return base64.RawURLEncoding.EncodeToString(bytes.Repeat([]byte{0xff}, 256)) // 2048 bits
}

// rsaKeyJSON is a key bundle matching testVaultKey's spec unless the arguments say otherwise.
func rsaKeyJSON(enabled bool, managed bool) string {
	managedField := ""
	if managed {
		managedField = `,"managed":true`
	}

	return fmt.Sprintf(
		`{"key":{"kid":"https://myvault.vault.azure.net/keys/%s/1","kty":"RSA","n":%q,"e":"AQAB",`+
			`"key_ops":["encrypt","decrypt"]},"attributes":{"enabled":%t},"tags":{"env":"test"}%s}`,
		testKeyName, rsaModulus(), enabled, managedField,
	)
}

const ecKeyJSON = `{"key":{"kid":"https://myvault.vault.azure.net/keys/my-key/1","kty":"EC","crv":"P-256",` +
	`"key_ops":["sign","verify"]},"attributes":{"enabled":true}}`

const rotationPolicyJSON = `{"id":"https://myvault.vault.azure.net/keys/my-key/rotationpolicy",` +
	`"attributes":{"expiryTime":"P2Y"},"lifetimeActions":[` +
	`{"action":{"type":"Rotate"},"trigger":{"timeAfterCreate":"P90D"}},` +
	`{"action":{"type":"Notify"},"trigger":{"timeBeforeExpiry":"P30D"}}]}`

// testVaultKey returns a claimed VaultKey whose status.keyUri points at the fake.
func testVaultKey(fake *fakeKeyVault, mutate func(*keys.VaultKey)) *keys.VaultKey {
	key := &keys.VaultKey{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testKeyName,
			Namespace: "default",
		},
		Spec: keys.VaultKey_Spec{
			AzureName: testKeyName,
			Properties: &keys.KeyProperties{
				Kty:     to.Ptr("RSA"),
				KeySize: to.Ptr(2048),
				KeyOps:  []string{"encrypt", "decrypt"},
				Attributes: &keys.KeyAttributes{
					Enabled: to.Ptr(true),
				},
			},
			Tags: map[string]string{"env": "test"},
		},
		Status: keys.VaultKey_STATUS{
			KeyUri: to.Ptr(fake.server.URL + keyPath),
		},
	}
	genruntime.SetResourceID(key, testKeyID)

	if mutate != nil {
		mutate(key)
	}

	return key
}

func withCreateMode(mode string) func(*keys.VaultKey) {
	return func(key *keys.VaultKey) {
		key.Spec.OperatorSpec = &keys.VaultKeyOperatorSpec{CreateMode: to.Ptr(mode)}
	}
}

func withDeleteMode(mode string) func(*keys.VaultKey) {
	return func(key *keys.VaultKey) {
		key.Spec.OperatorSpec = &keys.VaultKeyOperatorSpec{DeleteMode: to.Ptr(mode)}
	}
}

func preNext(called *bool) extensions.PreReconcileCheckFunc {
	return func(
		context.Context, genruntime.MetaObject, *resolver.Resolver, *genericarmclient.GenericClient, logr.Logger,
	) (extensions.PreReconcileCheckResult, error) {
		*called = true
		return extensions.ProceedWithReconcile(), nil
	}
}

func postNext(called *bool) extensions.PostReconcileCheckFunc {
	return func(
		context.Context, genruntime.MetaObject, genruntime.MetaObject, *resolver.Resolver,
		*genericarmclient.GenericClient, logr.Logger, annotations.ResolvedReconcilePolicies,
	) (extensions.PostReconcileCheckResult, error) {
		*called = true
		return extensions.PostReconcileCheckResultSuccess(), nil
	}
}

func managePolicy() annotations.ResolvedReconcilePolicies {
	return annotations.ResolvedReconcilePolicies{Effective: annotations.ReconcilePolicyManage}
}

// expectFatal asserts that err is a Ready-condition error of severity Error, i.e. one the reconciler
// reports as fatal until the spec changes.
func expectFatal(g *WithT, err error, substr string) {
	g.Expect(err).To(HaveOccurred())

	var readyErr *conditions.ReadyConditionImpactingError
	g.Expect(eris.As(err, &readyErr)).To(BeTrue(), "expected a ReadyConditionImpactingError, got: %s", err.Error())
	g.Expect(readyErr.Severity).To(Equal(conditions.ConditionSeverityError))
	g.Expect(err.Error()).To(ContainSubstring(substr))
}

// ---------------------------------------------------------------------------------------------
// Delete
// ---------------------------------------------------------------------------------------------

func Test_VaultKeyDeleteMode_DefaultsToDetach(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(vaultKeyDeleteMode(&keys.VaultKey{})).To(Equal(DeleteMode_Detach))

	g.Expect(vaultKeyDeleteMode(&keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			OperatorSpec: &keys.VaultKeyOperatorSpec{},
		},
	})).To(Equal(DeleteMode_Detach))

	g.Expect(vaultKeyDeleteMode(&keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			OperatorSpec: &keys.VaultKeyOperatorSpec{
				DeleteMode: to.Ptr(DeleteMode_Delete),
			},
		},
	})).To(Equal(DeleteMode_Delete))
}

// Detach must not require any Azure connectivity at all: it succeeds with no resolver and no ARM
// client, proving no data-plane or ARM call can be involved in the default deletion path.
func Test_VaultKeyExtension_Delete_DetachTouchesNothing(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ext := &VaultKeyExtension{}
	key := &keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			AzureName: "my-key",
		},
	}

	result, err := ext.Delete(context.Background(), logr.Discard(), nil, nil, key, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.Completed()).To(BeTrue())
}

func Test_VaultKeyExtension_Delete_RejectsUnexpectedResourceType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ext := &VaultKeyExtension{}

	_, err := ext.Delete(context.Background(), logr.Discard(), nil, nil, &keyvault.Vault{}, nil)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("unexpected resource type"))
}

func Test_VaultKeyExtension_Delete_RejectsUnknownDeleteMode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ext := &VaultKeyExtension{}
	key := &keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			AzureName: "my-key",
			OperatorSpec: &keys.VaultKeyOperatorSpec{
				DeleteMode: to.Ptr("obliterate"),
			},
		},
	}

	_, err := ext.Delete(context.Background(), logr.Discard(), nil, nil, key, nil)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("deleteMode"))
}

func Test_VaultKeyExtension_Delete_DataPlane(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		deleteMode    string
		liveStatus    int // response to the GET that precedes the operation; 200 with a matching key if zero
		liveBody      string
		method        string // the operation itself, if one is expected
		status        int
		body          string
		expectCalls   []string
		expectBody    string
		expectDone    bool
		expectBlocked string
		expectError   string
	}{
		"disable patches the key": {
			deleteMode:  DeleteMode_Disable,
			method:      http.MethodPatch,
			status:      http.StatusOK,
			body:        rsaKeyJSON(false, false),
			expectCalls: []string{"GET " + keyPath, "PATCH " + keyPath},
			expectBody:  `{"attributes":{"enabled":false}}`,
			expectDone:  true,
		},
		"delete soft-deletes the key": {
			deleteMode:  DeleteMode_Delete,
			method:      http.MethodDelete,
			status:      http.StatusOK,
			body:        rsaKeyJSON(true, false),
			expectCalls: []string{"GET " + keyPath, "DELETE " + keyPath},
			expectDone:  true,
		},
		"a key that is already gone counts as done": {
			deleteMode:  DeleteMode_Delete,
			liveStatus:  http.StatusNotFound,
			liveBody:    notFoundBody,
			expectCalls: []string{"GET " + keyPath},
			expectDone:  true,
		},
		"a key this resource never adopted is left untouched": {
			deleteMode:  DeleteMode_Delete,
			liveStatus:  http.StatusOK,
			liveBody:    ecKeyJSON, // the spec describes an RSA key
			expectCalls: []string{"GET " + keyPath},
			expectDone:  true,
		},
		"a certificate-backed key is left untouched": {
			deleteMode:  DeleteMode_Disable,
			liveStatus:  http.StatusOK,
			liveBody:    rsaKeyJSON(true, true),
			expectCalls: []string{"GET " + keyPath},
			expectDone:  true,
		},
		"a key that drifted in mutable properties is still ours and is deleted": {
			deleteMode: DeleteMode_Delete,
			liveStatus: http.StatusOK,
			liveBody: `{"key":{"kty":"RSA","n":"` + rsaModulus() + `","key_ops":["sign"]},` +
				`"attributes":{"enabled":false},"tags":{"other":"tag"}}`,
			method:      http.MethodDelete,
			status:      http.StatusOK,
			body:        rsaKeyJSON(true, false),
			expectCalls: []string{"GET " + keyPath, "DELETE " + keyPath},
			expectDone:  true,
		},
		"a missing read permission blocks deletion with guidance": {
			deleteMode:    DeleteMode_Delete,
			liveStatus:    http.StatusForbidden,
			liveBody:      forbiddenBody,
			expectCalls:   []string{"GET " + keyPath},
			expectBlocked: "Microsoft.KeyVault/vaults/keys/read",
		},
		"a missing update permission blocks deletion with guidance": {
			deleteMode:    DeleteMode_Disable,
			method:        http.MethodPatch,
			status:        http.StatusForbidden,
			body:          forbiddenBody,
			expectCalls:   []string{"GET " + keyPath, "PATCH " + keyPath},
			expectBlocked: "Microsoft.KeyVault/vaults/keys/update/action",
		},
		"other failures are errors": {
			deleteMode:  DeleteMode_Delete,
			method:      http.MethodDelete,
			status:      http.StatusInternalServerError,
			body:        `{"error":{"code":"InternalError","message":"boom"}}`,
			expectCalls: []string{"GET " + keyPath, "DELETE " + keyPath},
			expectError: "failed to soft-delete key",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			fake := newFakeKeyVault()
			defer fake.Close()
			liveStatus, liveBody := http.StatusOK, rsaKeyJSON(true, false)
			if c.liveStatus != 0 {
				liveStatus, liveBody = c.liveStatus, c.liveBody
			}
			fake.route(http.MethodGet, keyPath, liveStatus, liveBody)
			if c.method != "" {
				fake.route(c.method, keyPath, c.status, c.body)
			}

			key := testVaultKey(fake, withDeleteMode(c.deleteMode))
			ext := &VaultKeyExtension{}

			result, err := ext.Delete(context.Background(), logr.Discard(), nil, newARMClient(g, fake.server), key, nil)

			g.Expect(fake.Calls()).To(Equal(c.expectCalls))
			if c.expectBody != "" {
				g.Expect(fake.LastBody(c.expectCalls[len(c.expectCalls)-1])).To(MatchJSON(c.expectBody))
			}
			switch {
			case c.expectError != "":
				g.Expect(err).To(MatchError(ContainSubstring(c.expectError)))
			case c.expectBlocked != "":
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(result.BlockDeletion()).To(BeTrue())
				g.Expect(result.Message()).To(ContainSubstring(c.expectBlocked))
			default:
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(result.Completed()).To(Equal(c.expectDone))
			}
		})
	}
}

// Before the key exists, or when status was never populated, Delete finds the vault the same way
// PreReconcileCheck does: through the parent of the key's own ARM ID.
func Test_VaultKeyExtension_Delete_FallsBackToARMForVaultURL(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fake := newFakeKeyVault()
	defer fake.Close()
	fake.route(http.MethodGet, testVaultID, http.StatusOK,
		fmt.Sprintf(`{"id":%q,"name":"myvault","properties":{"vaultUri":%q}}`, testVaultID, fake.server.URL+"/"))
	fake.route(http.MethodGet, keyPath, http.StatusOK, rsaKeyJSON(true, false))
	fake.route(http.MethodDelete, keyPath, http.StatusOK, rsaKeyJSON(true, false))

	key := testVaultKey(fake, func(key *keys.VaultKey) {
		key.Status.KeyUri = nil
		key.Spec.OperatorSpec = &keys.VaultKeyOperatorSpec{DeleteMode: to.Ptr(DeleteMode_Delete)}
	})

	result, err := (&VaultKeyExtension{}).Delete(context.Background(), logr.Discard(), nil, newARMClient(g, fake.server), key, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.Completed()).To(BeTrue())
	g.Expect(fake.Calls()).To(Equal([]string{"GET " + testVaultID, "GET " + keyPath, "DELETE " + keyPath}))
}

// ---------------------------------------------------------------------------------------------
// PreReconcileCheck
// ---------------------------------------------------------------------------------------------

func Test_VaultKeyCreateMode_DefaultsToDefault(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(vaultKeyCreateMode(&keys.VaultKey{})).To(Equal(CreateMode_Default))

	g.Expect(vaultKeyCreateMode(&keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			OperatorSpec: &keys.VaultKeyOperatorSpec{
				CreateMode: to.Ptr(CreateMode_CreateOrRecover),
			},
		},
	})).To(Equal(CreateMode_CreateOrRecover))
}

func Test_VaultKeyExtension_PreReconcileCheck_RejectsUnexpectedResourceType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ext := &VaultKeyExtension{}

	_, err := ext.PreReconcileCheck(context.Background(), &keyvault.Vault{}, nil, nil, logr.Discard(), nil)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("unexpected resource type"))
}

func Test_VaultKeyExtension_PreReconcileCheck_DataPlane(t *testing.T) {
	t.Parallel()

	type route struct {
		method string
		path   string
		status int
		body   string
	}

	cases := map[string]struct {
		mutate      func(*keys.VaultKey)
		routes      []route
		expectCalls []string
		expectNext  bool
		expectBlock string
		expectFatal string
		expectError string
	}{
		"a matching live key is adopted": {
			routes:      []route{{http.MethodGet, keyPath, http.StatusOK, rsaKeyJSON(true, false)}},
			expectCalls: []string{"GET " + keyPath},
			expectNext:  true,
		},
		"a live key with different generation-time properties blocks with guidance": {
			mutate: func(key *keys.VaultKey) {
				key.Spec.Properties.Kty = to.Ptr("EC")
				key.Spec.Properties.KeySize = nil
			},
			routes:      []route{{http.MethodGet, keyPath, http.StatusOK, rsaKeyJSON(true, false)}},
			expectCalls: []string{"GET " + keyPath},
			expectBlock: "delete this resource and recreate it with matching properties",
		},
		"a key managed by Key Vault blocks with guidance": {
			routes:      []route{{http.MethodGet, keyPath, http.StatusOK, rsaKeyJSON(true, true)}},
			expectCalls: []string{"GET " + keyPath},
			expectBlock: "managed by Key Vault",
		},
		"a missing read permission is a retryable error with guidance": {
			routes:      []route{{http.MethodGet, keyPath, http.StatusForbidden, forbiddenBody}},
			expectCalls: []string{"GET " + keyPath},
			expectError: "refused the request (Forbidden)",
		},
		"no key and createMode default proceeds straight to ARM": {
			routes:      []route{{http.MethodGet, keyPath, http.StatusNotFound, notFoundBody}},
			expectCalls: []string{"GET " + keyPath},
			expectNext:  true,
		},
		"createMode recover with nothing to recover is blocked": {
			mutate: withCreateMode(CreateMode_Recover),
			routes: []route{
				{http.MethodGet, keyPath, http.StatusNotFound, notFoundBody},
				{http.MethodGet, deletedKeyPath, http.StatusNotFound, notFoundBody},
			},
			expectCalls: []string{"GET " + keyPath, "GET " + deletedKeyPath},
			expectBlock: "no soft-deleted key",
		},
		"a missing permission to look for soft-deleted keys is a retryable error with guidance": {
			mutate: withCreateMode(CreateMode_CreateOrRecover),
			routes: []route{
				{http.MethodGet, keyPath, http.StatusNotFound, notFoundBody},
				{http.MethodGet, deletedKeyPath, http.StatusForbidden, forbiddenBody},
			},
			expectCalls: []string{"GET " + keyPath, "GET " + deletedKeyPath},
			expectError: "failed to check for soft-deleted key",
		},
		"createMode createOrRecover with nothing to recover proceeds": {
			mutate: withCreateMode(CreateMode_CreateOrRecover),
			routes: []route{
				{http.MethodGet, keyPath, http.StatusNotFound, notFoundBody},
				{http.MethodGet, deletedKeyPath, http.StatusNotFound, notFoundBody},
			},
			expectCalls: []string{"GET " + keyPath, "GET " + deletedKeyPath},
			expectNext:  true,
		},
		"createMode createOrRecover recovers a matching soft-deleted key and waits": {
			mutate: withCreateMode(CreateMode_CreateOrRecover),
			routes: []route{
				{http.MethodGet, keyPath, http.StatusNotFound, notFoundBody},
				{http.MethodGet, deletedKeyPath, http.StatusOK, rsaKeyJSON(true, false)},
				{http.MethodPost, deletedKeyPath + "/recover", http.StatusOK, rsaKeyJSON(true, false)},
			},
			expectCalls: []string{"GET " + keyPath, "GET " + deletedKeyPath, "POST " + deletedKeyPath + "/recover"},
			expectBlock: "recovery of soft-deleted key",
		},
		"createMode recover refuses a soft-deleted key the spec does not describe": {
			mutate: withCreateMode(CreateMode_Recover),
			routes: []route{
				{http.MethodGet, keyPath, http.StatusNotFound, notFoundBody},
				{http.MethodGet, deletedKeyPath, http.StatusOK, ecKeyJSON},
			},
			expectCalls: []string{"GET " + keyPath, "GET " + deletedKeyPath},
			expectBlock: "cannot recover soft-deleted key",
		},
		"a recovery already in progress is blocked, not an error": {
			mutate: withCreateMode(CreateMode_CreateOrRecover),
			routes: []route{
				{http.MethodGet, keyPath, http.StatusNotFound, notFoundBody},
				{http.MethodGet, deletedKeyPath, http.StatusOK, rsaKeyJSON(true, false)},
				{http.MethodPost, deletedKeyPath + "/recover", http.StatusConflict, conflictBody},
			},
			expectCalls: []string{"GET " + keyPath, "GET " + deletedKeyPath, "POST " + deletedKeyPath + "/recover"},
			expectBlock: "in progress",
		},
		"createMode purgeThenCreate purges and proceeds": {
			mutate: withCreateMode(CreateMode_PurgeThenCreate),
			routes: []route{
				{http.MethodGet, keyPath, http.StatusNotFound, notFoundBody},
				{http.MethodGet, deletedKeyPath, http.StatusOK, rsaKeyJSON(true, false)},
				{http.MethodDelete, deletedKeyPath, http.StatusNoContent, ""},
			},
			expectCalls: []string{"GET " + keyPath, "GET " + deletedKeyPath, "DELETE " + deletedKeyPath},
			expectNext:  true,
		},
		"a purge refused by the service surfaces the failure": {
			mutate: withCreateMode(CreateMode_PurgeThenCreate),
			routes: []route{
				{http.MethodGet, keyPath, http.StatusNotFound, notFoundBody},
				{http.MethodGet, deletedKeyPath, http.StatusOK, rsaKeyJSON(true, false)},
				{http.MethodDelete, deletedKeyPath, http.StatusForbidden, forbiddenBody},
			},
			expectCalls: []string{"GET " + keyPath, "GET " + deletedKeyPath, "DELETE " + deletedKeyPath},
			expectError: "failed to purge soft-deleted key",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			fake := newFakeKeyVault()
			defer fake.Close()
			for _, r := range c.routes {
				fake.route(r.method, r.path, r.status, r.body)
			}

			key := testVaultKey(fake, c.mutate)
			ext := &VaultKeyExtension{}
			nextCalled := false

			result, err := ext.PreReconcileCheck(
				context.Background(), key, nil, newARMClient(g, fake.server), logr.Discard(), preNext(&nextCalled),
			)

			g.Expect(fake.Calls()).To(Equal(c.expectCalls))
			g.Expect(nextCalled).To(Equal(c.expectNext))
			switch {
			case c.expectFatal != "":
				expectFatal(g, err, c.expectFatal)
			case c.expectError != "":
				g.Expect(err).To(MatchError(ContainSubstring(c.expectError)))
				var readyErr *conditions.ReadyConditionImpactingError
				g.Expect(eris.As(err, &readyErr)).To(BeFalse(), "a retryable failure must not be reported as fatal")
			case c.expectBlock != "":
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(result.BlockReconciliation()).To(BeTrue())
				g.Expect(result.Message()).To(ContainSubstring(c.expectBlock))
			default:
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(result.BlockReconciliation()).To(BeFalse())
			}
		})
	}
}

// Before the key exists its status carries no URI, so the vault's data-plane URL comes from an ARM
// read of the vault, located through the parent of the key's own ARM ID (no Kubernetes owner lookup).
func Test_VaultKeyExtension_PreReconcileCheck_FallsBackToARMForVaultURL(t *testing.T) {
	t.Parallel()

	t.Run("vault found", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)

		fake := newFakeKeyVault()
		defer fake.Close()
		fake.route(http.MethodGet, testVaultID, http.StatusOK,
			fmt.Sprintf(`{"id":%q,"name":"myvault","properties":{"vaultUri":%q}}`, testVaultID, fake.server.URL+"/"))
		fake.route(http.MethodGet, keyPath, http.StatusNotFound, notFoundBody)

		key := testVaultKey(fake, func(key *keys.VaultKey) { key.Status.KeyUri = nil })
		nextCalled := false

		result, err := (&VaultKeyExtension{}).PreReconcileCheck(
			context.Background(), key, nil, newARMClient(g, fake.server), logr.Discard(), preNext(&nextCalled),
		)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(result.BlockReconciliation()).To(BeFalse())
		g.Expect(nextCalled).To(BeTrue())
		g.Expect(fake.Calls()).To(Equal([]string{"GET " + testVaultID, "GET " + keyPath}))
	})

	t.Run("vault not found blocks until it exists", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)

		fake := newFakeKeyVault()
		defer fake.Close()
		fake.route(http.MethodGet, testVaultID, http.StatusNotFound,
			`{"error":{"code":"ResourceNotFound","message":"The Resource 'Microsoft.KeyVault/vaults/myvault' was not found."}}`)

		key := testVaultKey(fake, func(key *keys.VaultKey) { key.Status.KeyUri = nil })
		nextCalled := false

		result, err := (&VaultKeyExtension{}).PreReconcileCheck(
			context.Background(), key, nil, newARMClient(g, fake.server), logr.Discard(), preNext(&nextCalled),
		)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(result.BlockReconciliation()).To(BeTrue())
		g.Expect(result.Message()).To(ContainSubstring("not found"))
		g.Expect(nextCalled).To(BeFalse())
	})

	t.Run("no ARM ID is an error", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)

		fake := newFakeKeyVault()
		defer fake.Close()

		key := testVaultKey(fake, func(key *keys.VaultKey) {
			key.Status.KeyUri = nil
			key.SetAnnotations(nil)
		})

		_, err := (&VaultKeyExtension{}).PreReconcileCheck(
			context.Background(), key, nil, newARMClient(g, fake.server), logr.Discard(), preNext(new(bool)),
		)
		g.Expect(err).To(MatchError(ContainSubstring("ARM resource ID")))
		g.Expect(fake.Calls()).To(BeEmpty())
	})
}

// ---------------------------------------------------------------------------------------------
// PostReconcileCheck
// ---------------------------------------------------------------------------------------------

func Test_VaultKeyExtension_PostReconcileCheck_RejectsUnexpectedResourceType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ext := &VaultKeyExtension{}

	_, err := ext.PostReconcileCheck(
		context.Background(), &keyvault.Vault{}, nil, nil, nil, logr.Discard(), annotations.ResolvedReconcilePolicies{}, nil,
	)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("unexpected resource type"))
}

func Test_VaultKeyExtension_PostReconcileCheck_HonoursReconcilePolicy(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fake := newFakeKeyVault()
	defer fake.Close()

	// The live key is disabled while the spec wants it enabled; under skip nothing may be written
	fake.route(http.MethodGet, keyPath, http.StatusOK, rsaKeyJSON(false, false))

	key := testVaultKey(fake, nil)
	nextCalled := false

	result, err := (&VaultKeyExtension{}).PostReconcileCheck(
		context.Background(), key, nil, nil, newARMClient(g, fake.server), logr.Discard(),
		annotations.ResolvedReconcilePolicies{Effective: annotations.ReconcilePolicySkip}, postNext(&nextCalled),
	)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeTrue())
	g.Expect(nextCalled).To(BeTrue())
	g.Expect(fake.Calls()).To(BeEmpty())
}

func Test_VaultKeyExtension_PostReconcileCheck_DataPlane(t *testing.T) {
	t.Parallel()

	type route struct {
		method string
		path   string
		status int
		body   string
	}

	withRotationPolicy := func(key *keys.VaultKey) {
		key.Spec.Properties.RotationPolicy = &keys.RotationPolicy{
			Attributes: &keys.KeyRotationPolicyAttributes{ExpiryTime: to.Ptr("P2Y")},
			LifetimeActions: []keys.LifetimeAction{{
				Action:  &keys.Action{Type: to.Ptr("rotate")},
				Trigger: &keys.Trigger{TimeAfterCreate: to.Ptr("P90D")},
			}},
		}
	}

	releasePolicy := []byte(`{"anyOf":[{"allOf":[{"claim":"x-ms-sgx-is-debuggable","equals":"false"}]}]}`)
	withReleasePolicy := func(key *keys.VaultKey) {
		key.Spec.Properties.Release_Policy = &keys.KeyReleasePolicy{
			Data: to.Ptr(base64.RawURLEncoding.EncodeToString(releasePolicy)),
		}
	}
	keyWithEmptyPolicy := `{"key":{"kty":"RSA","n":"` + rsaModulus() + `","key_ops":["encrypt","decrypt"]},"attributes":{"enabled":true},` +
		`"tags":{"env":"test"},"release_policy":{"data":"e30","immutable":false}}`

	cases := map[string]struct {
		mutate      func(*keys.VaultKey)
		routes      []route
		expectCalls []string
		expectBody  string
		expectFatal string
		expectError string
	}{
		"a key matching the spec is left alone": {
			routes:      []route{{http.MethodGet, keyPath, http.StatusOK, rsaKeyJSON(true, false)}},
			expectCalls: []string{"GET " + keyPath},
		},
		"drifted mutable properties are patched, and only those": {
			routes: []route{
				{http.MethodGet, keyPath, http.StatusOK, rsaKeyJSON(false, false)},
				{http.MethodPatch, keyPath, http.StatusOK, rsaKeyJSON(true, false)},
			},
			expectCalls: []string{"GET " + keyPath, "PATCH " + keyPath},
			expectBody:  `{"attributes":{"enabled":true}}`,
		},
		"a release policy differing from a mutable one is patched": {
			mutate: withReleasePolicy,
			routes: []route{
				{http.MethodGet, keyPath, http.StatusOK, keyWithEmptyPolicy},
				{http.MethodPatch, keyPath, http.StatusOK, rsaKeyJSON(true, false)},
			},
			expectCalls: []string{"GET " + keyPath, "PATCH " + keyPath},
			expectBody:  fmt.Sprintf(`{"release_policy":{"data":%q}}`, base64.RawURLEncoding.EncodeToString(releasePolicy)),
		},
		"a content type change carries the key's current policy along": {
			mutate: func(key *keys.VaultKey) {
				key.Spec.Properties.Release_Policy = &keys.KeyReleasePolicy{
					ContentType: to.Ptr("application/json; charset=utf-8"),
				}
			},
			routes: []route{
				{http.MethodGet, keyPath, http.StatusOK, keyWithEmptyPolicy},
				{http.MethodPatch, keyPath, http.StatusOK, rsaKeyJSON(true, false)},
			},
			expectCalls: []string{"GET " + keyPath, "PATCH " + keyPath},
			expectBody:  `{"release_policy":{"contentType":"application/json; charset=utf-8","data":"e30"}}`,
		},
		"a rotation policy matching the spec is left alone, service-added notify included": {
			mutate: withRotationPolicy,
			routes: []route{
				{http.MethodGet, keyPath, http.StatusOK, rsaKeyJSON(true, false)},
				{http.MethodGet, rotationPolicyPath, http.StatusOK, rotationPolicyJSON},
			},
			expectCalls: []string{"GET " + keyPath, "GET " + rotationPolicyPath},
		},
		"a drifted rotation policy is replaced with exactly the spec's policy": {
			mutate: withRotationPolicy,
			routes: []route{
				{http.MethodGet, keyPath, http.StatusOK, rsaKeyJSON(true, false)},
				{http.MethodGet, rotationPolicyPath, http.StatusOK, `{"attributes":{"expiryTime":"P1Y"},"lifetimeActions":[]}`},
				{http.MethodPut, rotationPolicyPath, http.StatusOK, rotationPolicyJSON},
			},
			expectCalls: []string{"GET " + keyPath, "GET " + rotationPolicyPath, "PUT " + rotationPolicyPath},
			expectBody:  `{"attributes":{"expiryTime":"P2Y"},"lifetimeActions":[{"action":{"type":"Rotate"},"trigger":{"timeAfterCreate":"P90D"}}]}`,
		},
		"a rotation policy the identity may not write is a retryable error with guidance": {
			mutate: withRotationPolicy,
			routes: []route{
				{http.MethodGet, keyPath, http.StatusOK, rsaKeyJSON(true, false)},
				{http.MethodGet, rotationPolicyPath, http.StatusOK, `{"attributes":{"expiryTime":"P1Y"},"lifetimeActions":[]}`},
				{http.MethodPut, rotationPolicyPath, http.StatusForbidden, forbiddenBody},
			},
			expectCalls: []string{"GET " + keyPath, "GET " + rotationPolicyPath, "PUT " + rotationPolicyPath},
			expectError: "failed to update rotation policy",
		},
		"an immutable release policy that differs from the spec is fatal": {
			mutate: func(key *keys.VaultKey) {
				key.Spec.Properties.Release_Policy = &keys.KeyReleasePolicy{
					Data: to.Ptr(base64.RawURLEncoding.EncodeToString([]byte(`{"anyOf":[{"allOf":[{"claim":"x","equals":"1"}]}]}`))),
				}
			},
			routes: []route{{
				http.MethodGet, keyPath, http.StatusOK,
				`{"key":{"kty":"RSA","n":"` + rsaModulus() + `","key_ops":["encrypt","decrypt"]},"attributes":{"enabled":true},` +
					`"tags":{"env":"test"},"release_policy":{"data":"e30","immutable":true}}`,
			}},
			expectCalls: []string{"GET " + keyPath},
			expectFatal: "immutable",
		},
		"a missing update permission is a retryable error with guidance": {
			routes: []route{
				{http.MethodGet, keyPath, http.StatusOK, rsaKeyJSON(false, false)},
				{http.MethodPatch, keyPath, http.StatusForbidden, forbiddenBody},
			},
			expectCalls: []string{"GET " + keyPath, "PATCH " + keyPath},
			expectError: "permission",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			fake := newFakeKeyVault()
			defer fake.Close()
			for _, r := range c.routes {
				fake.route(r.method, r.path, r.status, r.body)
			}

			key := testVaultKey(fake, c.mutate)
			nextCalled := false

			result, err := (&VaultKeyExtension{}).PostReconcileCheck(
				context.Background(), key, nil, nil, newARMClient(g, fake.server), logr.Discard(), managePolicy(), postNext(&nextCalled),
			)

			g.Expect(fake.Calls()).To(Equal(c.expectCalls))
			if c.expectBody != "" {
				g.Expect(fake.LastBody(c.expectCalls[len(c.expectCalls)-1])).To(MatchJSON(c.expectBody))
			}
			switch {
			case c.expectFatal != "":
				expectFatal(g, err, c.expectFatal)
				g.Expect(nextCalled).To(BeFalse())
			case c.expectError != "":
				g.Expect(err).To(MatchError(ContainSubstring(c.expectError)))
				g.Expect(nextCalled).To(BeFalse())
			default:
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(result.ReconciliationSucceeded()).To(BeTrue())
				g.Expect(nextCalled).To(BeTrue())
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------
// Pure helpers
// ---------------------------------------------------------------------------------------------

func Test_IntrinsicMismatch(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	keyWith := func(props *keys.KeyProperties) *keys.VaultKey {
		return &keys.VaultKey{
			Spec: keys.VaultKey_Spec{
				AzureName:  "my-key",
				Properties: props,
			},
		}
	}

	rsa2048 := &azkeys.JSONWebKey{
		Kty: to.Ptr(azkeys.KeyTypeRSA),
		N:   make([]byte, 256), // 2048-bit modulus
	}
	ecP256 := &azkeys.JSONWebKey{
		Kty: to.Ptr(azkeys.KeyTypeEC),
		Crv: to.Ptr(azkeys.CurveNameP256),
	}

	// A spec with no properties can't conflict
	g.Expect(intrinsicMismatch(keyWith(nil), rsa2048)).To(BeEmpty())

	// Matching properties pass
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty:     to.Ptr("RSA"),
		KeySize: to.Ptr(2048),
	}), rsa2048)).To(BeEmpty())
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty:       to.Ptr("EC"),
		CurveName: to.Ptr("P-256"),
	}), ecP256)).To(BeEmpty())

	// Unset spec properties match anything
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{}), rsa2048)).To(BeEmpty())

	// Mismatches are reported per property
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty: to.Ptr("EC"),
	}), rsa2048)).To(ContainSubstring("kty"))
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty:     to.Ptr("RSA"),
		KeySize: to.Ptr(4096),
	}), rsa2048)).To(ContainSubstring("keySize"))
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty:       to.Ptr("EC"),
		CurveName: to.Ptr("P-384"),
	}), ecP256)).To(ContainSubstring("curveName"))

	// keySize can't be verified for keys without an RSA modulus, so it's left to the service
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty:     to.Ptr("EC"),
		KeySize: to.Ptr(2048),
	}), ecP256)).To(BeEmpty())
}

func Test_KeyUpdatesNeeded_NothingManagedMeansNoUpdate(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// A spec that sets nothing mutable requires no update, whatever the key looks like; an empty
	// keyOps list or tags map counts as unset
	key := &keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			AzureName: "my-key",
			Tags:      map[string]string{},
			Properties: &keys.KeyProperties{
				Kty:    to.Ptr("RSA"),
				KeyOps: []string{},
			},
		},
	}
	actual := azkeys.KeyBundle{
		Attributes: &azkeys.KeyAttributes{Enabled: to.Ptr(false)},
		Tags:       map[string]*string{"unmanaged": to.Ptr("tag")},
		Key: &azkeys.JSONWebKey{
			KeyOps: []*azkeys.KeyOperation{to.Ptr(azkeys.KeyOperationEncrypt)},
		},
	}

	_, changed, err := keyUpdatesNeeded(key, actual)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeFalse())
}

func Test_KeyUpdatesNeeded_MatchingSpecMeansNoUpdate(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	exp := 1893456000 // 2030-01-01T00:00:00Z
	key := &keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			AzureName: "my-key",
			Tags:      map[string]string{"env": "test"},
			Properties: &keys.KeyProperties{
				KeyOps: []string{"encrypt", "decrypt"},
				Attributes: &keys.KeyAttributes{
					Enabled: to.Ptr(true),
					Exp:     to.Ptr(exp),
				},
			},
		},
	}
	actual := azkeys.KeyBundle{
		Attributes: &azkeys.KeyAttributes{
			Enabled: to.Ptr(true),
			Expires: to.Ptr(time.Unix(int64(exp), 0)),
		},
		Tags: map[string]*string{"env": to.Ptr("test")},
		Key: &azkeys.JSONWebKey{
			// Order deliberately differs from the spec; keyOps compare as a set
			KeyOps: []*azkeys.KeyOperation{
				to.Ptr(azkeys.KeyOperationDecrypt),
				to.Ptr(azkeys.KeyOperationEncrypt),
			},
		},
	}

	_, changed, err := keyUpdatesNeeded(key, actual)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeFalse())
}

func Test_KeyUpdatesNeeded_DivergedPropertiesAreUpdated(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	nbf := 1735689600 // 2025-01-01T00:00:00Z
	key := &keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			AzureName: "my-key",
			Tags:      map[string]string{"env": "prod"},
			Properties: &keys.KeyProperties{
				KeyOps: []string{"sign", "verify"},
				Attributes: &keys.KeyAttributes{
					Enabled: to.Ptr(true),
					Nbf:     to.Ptr(nbf),
				},
			},
		},
	}
	actual := azkeys.KeyBundle{
		// No attributes reported at all: everything the spec sets must be applied
		Tags: map[string]*string{"env": to.Ptr("test")},
		Key: &azkeys.JSONWebKey{
			KeyOps: []*azkeys.KeyOperation{to.Ptr(azkeys.KeyOperationEncrypt)},
		},
	}

	params, changed, err := keyUpdatesNeeded(key, actual)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeTrue())
	g.Expect(params.KeyOps).To(HaveLen(2))
	g.Expect(params.KeyAttributes).ToNot(BeNil())
	g.Expect(params.KeyAttributes.Enabled).To(HaveValue(BeTrue()))
	g.Expect(params.KeyAttributes.NotBefore).To(HaveValue(Equal(time.Unix(int64(nbf), 0).UTC())))
	// Expiry was not set in the spec, so the update must not touch it
	g.Expect(params.KeyAttributes.Expires).To(BeNil())
	g.Expect(params.Tags).To(HaveKeyWithValue("env", HaveValue(Equal("prod"))))

	// A diverged expiry is applied on its own
	exp := 1893456000 // 2030-01-01T00:00:00Z
	key.Spec.Properties.Attributes = &keys.KeyAttributes{Exp: to.Ptr(exp)}
	key.Spec.Properties.KeyOps = nil
	key.Spec.Tags = nil
	actual.Attributes = &azkeys.KeyAttributes{Expires: to.Ptr(time.Unix(int64(exp)-86400, 0))}

	params, changed, err = keyUpdatesNeeded(key, actual)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeTrue())
	g.Expect(params.KeyOps).To(BeNil())
	g.Expect(params.Tags).To(BeNil())
	g.Expect(params.KeyAttributes.Enabled).To(BeNil())
	g.Expect(params.KeyAttributes.Expires).To(HaveValue(Equal(time.Unix(int64(exp), 0).UTC())))
}

func Test_KeyUpdatesNeeded_ImmutableReleasePolicyIsFatal(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	key := &keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			AzureName: "my-key",
			Properties: &keys.KeyProperties{
				Release_Policy: &keys.KeyReleasePolicy{
					Data: to.Ptr(base64.RawURLEncoding.EncodeToString([]byte(`{"anyOf":[]}`))),
				},
			},
		},
	}
	actual := azkeys.KeyBundle{
		ReleasePolicy: &azkeys.KeyReleasePolicy{
			EncodedPolicy: []byte(`{}`),
			Immutable:     to.Ptr(true),
		},
	}

	_, _, err := keyUpdatesNeeded(key, actual)
	expectFatal(g, err, "immutable")

	// An immutable policy that already matches the spec is no problem at all
	actual.ReleasePolicy.EncodedPolicy = []byte(`{ "anyOf" : [ ] }`)
	_, changed, err := keyUpdatesNeeded(key, actual)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeFalse())
}

func Test_ReleasePolicyUpdateNeeded(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	policyJSON := []byte(`{"anyOf":[{"allOf":[{"claim":"x-ms-sgx-is-debuggable","equals":"false"}]}]}`)
	encoded := base64.RawURLEncoding.EncodeToString(policyJSON)

	spec := &keys.KeyReleasePolicy{Data: to.Ptr(encoded)}

	// Matching policy: no update
	_, changed, err := releasePolicyUpdateNeeded(spec, &azkeys.KeyReleasePolicy{EncodedPolicy: policyJSON})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeFalse())

	// The same policy formatted differently by the service is still a match
	reformatted := []byte("{\n  \"anyOf\": [ { \"allOf\": [ { \"equals\": \"false\", \"claim\": \"x-ms-sgx-is-debuggable\" } ] } ]\n}")
	_, changed, err = releasePolicyUpdateNeeded(spec, &azkeys.KeyReleasePolicy{EncodedPolicy: reformatted})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeFalse())

	// Padded base64url in the spec decodes the same way
	padded := &keys.KeyReleasePolicy{Data: to.Ptr(base64.URLEncoding.EncodeToString(policyJSON))}
	_, changed, err = releasePolicyUpdateNeeded(padded, &azkeys.KeyReleasePolicy{EncodedPolicy: policyJSON})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeFalse())

	// Different policy: update with the decoded bytes
	desired, changed, err := releasePolicyUpdateNeeded(spec, &azkeys.KeyReleasePolicy{EncodedPolicy: []byte(`{}`)})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeTrue())
	g.Expect(desired.EncodedPolicy).To(Equal(policyJSON))

	// No policy on the key at all: update
	_, changed, err = releasePolicyUpdateNeeded(spec, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeTrue())

	// Content type is compared when the spec sets it
	typed := &keys.KeyReleasePolicy{Data: to.Ptr(encoded), ContentType: to.Ptr("application/json; charset=utf-8")}
	_, changed, err = releasePolicyUpdateNeeded(typed, &azkeys.KeyReleasePolicy{EncodedPolicy: policyJSON})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeTrue())

	// A content type on its own carries the key's current policy, since the service won't take one
	// without the other; with no policy on the key there is nothing to change
	typeOnly := &keys.KeyReleasePolicy{ContentType: to.Ptr("application/json; charset=utf-8")}
	desired, changed, err = releasePolicyUpdateNeeded(typeOnly, &azkeys.KeyReleasePolicy{EncodedPolicy: policyJSON})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeTrue())
	g.Expect(desired.EncodedPolicy).To(Equal(policyJSON))
	_, changed, err = releasePolicyUpdateNeeded(typeOnly, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeFalse())

	// Invalid base64 in the spec is fatal, not a silent no-op and not a retry loop
	_, _, err = releasePolicyUpdateNeeded(&keys.KeyReleasePolicy{Data: to.Ptr("!!not-base64!!")}, nil)
	expectFatal(g, err, "base64url")
}

func Test_RotationPolicyUpdateNeeded(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spec := &keys.RotationPolicy{
		Attributes: &keys.KeyRotationPolicyAttributes{ExpiryTime: to.Ptr("P2Y")},
		LifetimeActions: []keys.LifetimeAction{
			{
				// ARM spells the action lowercase; the data plane answers with "Rotate".
				// The comparison must be case-insensitive per the service contract.
				Action:  &keys.Action{Type: to.Ptr("rotate")},
				Trigger: &keys.Trigger{TimeAfterCreate: to.Ptr("P90D")},
			},
		},
	}

	rotate := func(afterCreate string) *azkeys.LifetimeAction {
		return &azkeys.LifetimeAction{
			Action:  &azkeys.LifetimeActionType{Type: to.Ptr(azkeys.KeyRotationPolicyActionRotate)},
			Trigger: &azkeys.LifetimeActionTrigger{TimeAfterCreate: to.Ptr(afterCreate)},
		}
	}
	notify := func(beforeExpiry string) *azkeys.LifetimeAction {
		return &azkeys.LifetimeAction{
			Action:  &azkeys.LifetimeActionType{Type: to.Ptr(azkeys.KeyRotationPolicyActionNotify)},
			Trigger: &azkeys.LifetimeActionTrigger{TimeBeforeExpiry: to.Ptr(beforeExpiry)},
		}
	}
	expiry := &azkeys.KeyRotationPolicyAttributes{ExpiryTime: to.Ptr("P2Y")}

	// The service-added default notify action must be tolerated, or every reconcile would apply
	// the policy again forever
	_, changed := rotationPolicyUpdateNeeded(spec, azkeys.KeyRotationPolicy{
		Attributes:      expiry,
		LifetimeActions: []*azkeys.LifetimeAction{rotate("P90D"), notify("P30D")},
	})
	g.Expect(changed).To(BeFalse())

	// Different trigger: update, carrying the spec's policy with canonical action casing
	desired, changed := rotationPolicyUpdateNeeded(spec, azkeys.KeyRotationPolicy{
		Attributes:      expiry,
		LifetimeActions: []*azkeys.LifetimeAction{rotate("P180D")},
	})
	g.Expect(changed).To(BeTrue())
	g.Expect(desired.Attributes.ExpiryTime).To(HaveValue(Equal("P2Y")))
	g.Expect(desired.LifetimeActions).To(HaveLen(1))
	g.Expect(desired.LifetimeActions[0].Action.Type).To(HaveValue(Equal(azkeys.KeyRotationPolicyActionRotate)))

	// Missing expiryTime on the actual policy: update
	_, changed = rotationPolicyUpdateNeeded(spec, azkeys.KeyRotationPolicy{
		LifetimeActions: []*azkeys.LifetimeAction{rotate("P90D"), notify("P30D")},
	})
	g.Expect(changed).To(BeTrue())

	// The policy is managed as a whole: an action the spec doesn't declare (here a notify that is
	// not the service default) must be removed, so this is an update
	_, changed = rotationPolicyUpdateNeeded(spec, azkeys.KeyRotationPolicy{
		Attributes:      expiry,
		LifetimeActions: []*azkeys.LifetimeAction{rotate("P90D"), notify("P60D")},
	})
	g.Expect(changed).To(BeTrue())

	// Likewise a second rotate action the spec doesn't declare
	_, changed = rotationPolicyUpdateNeeded(spec, azkeys.KeyRotationPolicy{
		Attributes:      expiry,
		LifetimeActions: []*azkeys.LifetimeAction{rotate("P90D"), rotate("P30D"), notify("P30D")},
	})
	g.Expect(changed).To(BeTrue())

	// An expiry time the spec doesn't set is not managed: it neither causes a diff nor is it
	// erased when the policy is replaced for another reason
	specNoExpiry := &keys.RotationPolicy{LifetimeActions: spec.LifetimeActions}
	_, changed = rotationPolicyUpdateNeeded(specNoExpiry, azkeys.KeyRotationPolicy{
		Attributes:      &azkeys.KeyRotationPolicyAttributes{ExpiryTime: to.Ptr("P1Y")},
		LifetimeActions: []*azkeys.LifetimeAction{rotate("P90D"), notify("P30D")},
	})
	g.Expect(changed).To(BeFalse())
	desired, changed = rotationPolicyUpdateNeeded(specNoExpiry, azkeys.KeyRotationPolicy{
		Attributes:      &azkeys.KeyRotationPolicyAttributes{ExpiryTime: to.Ptr("P1Y")},
		LifetimeActions: []*azkeys.LifetimeAction{rotate("P180D")},
	})
	g.Expect(changed).To(BeTrue())
	g.Expect(desired.Attributes).ToNot(BeNil())
	g.Expect(desired.Attributes.ExpiryTime).To(HaveValue(Equal("P1Y")))

	// When the spec declares its own notify action, the service default is no longer tolerated
	specWithNotify := &keys.RotationPolicy{
		Attributes: spec.Attributes,
		LifetimeActions: append(spec.LifetimeActions, keys.LifetimeAction{
			Action:  &keys.Action{Type: to.Ptr("notify")},
			Trigger: &keys.Trigger{TimeBeforeExpiry: to.Ptr("P60D")},
		}),
	}
	_, changed = rotationPolicyUpdateNeeded(specWithNotify, azkeys.KeyRotationPolicy{
		Attributes:      expiry,
		LifetimeActions: []*azkeys.LifetimeAction{rotate("P90D"), notify("P30D")},
	})
	g.Expect(changed).To(BeTrue())
	_, changed = rotationPolicyUpdateNeeded(specWithNotify, azkeys.KeyRotationPolicy{
		Attributes:      expiry,
		LifetimeActions: []*azkeys.LifetimeAction{rotate("P90D"), notify("P60D")},
	})
	g.Expect(changed).To(BeFalse())
}

func Test_CanonicalRotationAction(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(canonicalRotationAction("rotate")).To(Equal(azkeys.KeyRotationPolicyActionRotate))
	g.Expect(canonicalRotationAction("Notify")).To(Equal(azkeys.KeyRotationPolicyActionNotify))
	g.Expect(canonicalRotationAction("archive")).To(Equal(azkeys.KeyRotationPolicyAction("archive")))
}

func Test_VaultURLFromKeyURI(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	vaultURL, err := vaultURLFromKeyURI("https://myvault.vault.azure.net/keys/my-key")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(vaultURL).To(Equal("https://myvault.vault.azure.net"))

	// Sovereign clouds use a different DNS suffix; nothing may be hardcoded
	vaultURL, err = vaultURLFromKeyURI("https://myvault.vault.azure.cn/keys/my-key/abc123")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(vaultURL).To(Equal("https://myvault.vault.azure.cn"))

	_, err = vaultURLFromKeyURI("not-a-uri")
	g.Expect(err).To(HaveOccurred())
}
