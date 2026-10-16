# ARM Request Destination Validation Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Reject ARM requests and redirects outside the selected client's trusted origin without changing normal reconciliation recovery or constraining data-plane clients.

**Architecture:** A shared HTTP package compares parsed HTTPS origins and creates redirect-guarded client copies. An ARM per-call policy validates destinations before authentication, including pipelines inherited by SDK clients and resource-provider registration. Data-plane options retain the original HTTP client.

**Tech Stack:** Go 1.26.0 as required by `v2/go.mod`, azcore v1.23.3, `net/http`, Gomega, existing ASO task and recording infrastructure. No new dependencies.

## Global Constraints

- "Validate each ARM request against the effective client's Resource Manager origin before bearer authentication."
- "Reject off-origin redirects before sending the redirected request."
- "Use credential-specific cloud configuration when present, falling back to global configuration through the existing client-selection mechanism."
- "Preserve connection pooling, existing HTTP client settings, ordinary retry behaviour, and data-plane client behaviour."
- "Keep all work local until coordinated disclosure is cleared."
- "Do not add an additional allowlist setting or an enforcement-disable switch."
- "Do not use GitHub's new PR-stacking features."
- Keep the configured Go proxy. Never switch to direct module downloads.
- Work in this existing feature worktree and branch, not the main checkout.
- Do not edit generated files or introduce a resume-token parser in production.
- Keep Graph, Key Vault and each database group's enforcement out of this change.
- Include `Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>` in commits.

---

## Approved spec and execution preflight

Read [the approved design](../specs/2026-10-16-arm-request-destination-validation-design.md).
Ordinary stacked branches/PRs are allowed; no pushes or public PRs are currently
authorised. Subsequent changes are Graph, Key Vault, `dbformysql.azure.com`,
`dbforpostgresql.azure.com`, and `sql.azure.com`, each with its own design, plan,
branch and PR. Do not implement those changes while executing this plan.

Before changing Go files, run:

```bash
git status --short
./dev.sh -c 'go version'
```

Record any pre-existing changes and leave unrelated work intact. Stop only if
unexpected changes conflict with this task. `dev.sh` performs the repository's
required environment setup; allow it to finish. If setup fails, stop and report
the failure. The Go version must satisfy `v2/go.mod`. The clone was already
confirmed non-shallow during design.

All new Go files use the repository's Microsoft copyright/MIT header.
Production snippets below are complete new files or exact replacement blocks;
existing files outside those blocks stay unchanged.

## File responsibilities

| File | Responsibility |
| --- | --- |
| `v2/internal/httpclient/origin.go` | Origin parsing/comparison and safe typed errors, independent of Azure SDK |
| `v2/internal/httpclient/origin_test.go` | URL and port equivalence/rejection tests |
| `v2/internal/httpclient/redirects.go` | Copy a client, enforce redirect origins, preserve callbacks and sanitise destination errors |
| `v2/internal/httpclient/redirects_test.go` | Redirect non-transmission, callback semantics, client isolation and safe errors |
| `v2/internal/genericarmclient/destination_policy.go` | Pre-authentication ARM destination policy |
| `v2/internal/genericarmclient/generic_client.go` | Bind effective ARM origin, keep base and guarded clients separate, propagate policies |
| `v2/internal/genericarmclient/destination_policy_test.go` | Inherited-pipeline integration and reusable offline test harness |
| `v2/internal/genericarmclient/destination_poller_test.go` | Every supported SDK poller form, final GETs and response-derived destinations |
| `v2/internal/genericarmclient/destination_sdk_test.go` | BackupInstance SDK resume path and per-client cloud isolation |
| `v2/internal/genericarmclient/data_plane_client_options_test.go` | Preserve data-plane access to another origin |
| `v2/internal/reconcilers/arm/destination_failure_test.go` | Existing clear-and-retry recovery remains intact |
| `docs/hugo/content/guide/aso-controller-settings-options.md` | Document the endpoint boundary |

Tasks 1 and 2 are shared infrastructure. Task 3 integrates it with ARM. Task 4
proves complete LRO/recovery coverage. Task 5 documents and verifies the change.
Each task depends on the preceding tasks.

### Task 1: Implement shared HTTPS-origin validation

**Files:** Create `v2/internal/httpclient/origin.go` and
`v2/internal/httpclient/origin_test.go`.

**Interfaces:**
- Consumes: standard-library `net/url.URL`.
- Produces: `NewOrigin(endpoint string) (Origin, error)`,
  `(Origin).Validate(destination *url.URL) error`, `(Origin).String() string`,
  and `*DestinationError` with safe `Allowed` and `Destination` strings.

- [ ] **Step 1: Create the failing tests in `origin_test.go`.**

```go
package httpclient

import (
	"errors"
	"net/url"
	"strings"
	"testing"
)

func TestOrigin(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name string
		raw  string
		ok   bool
	}{
		{"same", "https://management.example.invalid/subscriptions/s/operations/o?q=x", true},
		{"case", "https://MANAGEMENT.example.invalid/path", true},
		{"default port", "https://management.example.invalid:443/path", true},
		{"numeric port", "https://management.example.invalid:0443/path", true},
		{"other host", "https://other.example.invalid/path", false},
		{"subdomain", "https://child.management.example.invalid/path", false},
		{"suffix", "https://management.example.invalid.evil.invalid/path", false},
		{"trailing dot", "https://management.example.invalid./path", false},
		{"port", "https://management.example.invalid:444/path", false},
		{"downgrade", "http://management.example.invalid/path", false},
		{"userinfo", "https://user:private@management.example.invalid/path", false},
		{"relative", "/operations/1", false},
		{"opaque", "https:management.example.invalid/path", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			u, parseErr := url.Parse(c.raw)
			if parseErr != nil {
				t.Fatal(parseErr)
			}
			err := allowed.Validate(u)
			if (err == nil) != c.ok {
				t.Fatalf("allowed=%v, error=%v", c.ok, err)
			}
			if err != nil {
				var rejected *DestinationError
				if !errors.As(err, &rejected) {
					t.Fatalf("unexpected error type: %T", err)
				}
				if strings.Contains(err.Error(), "private") || strings.Contains(err.Error(), "?") {
					t.Fatalf("unsafe diagnostic: %v", err)
				}
			}
		})
	}
	if err := allowed.Validate(nil); err == nil {
		t.Fatal("nil URL accepted")
	}
}

func TestOrigin_InvalidEndpoint(t *testing.T) {
	t.Parallel()
	for _, raw := range []string{
		"", "/relative", "http://example.invalid", "https://",
		"https://u:secret@example.invalid", "https://example.invalid:",
		"https://example.invalid:0", "https://example.invalid:65536",
		"https://example.invalid:bad", "https://[broken",
	} {
		if _, err := NewOrigin(raw); err == nil {
			t.Fatalf("accepted invalid endpoint %q", raw)
		}
	}
}

func TestOrigin_CustomPortsAndIPv6(t *testing.T) {
	t.Parallel()
	for _, endpoint := range []string{"https://localhost:8443/base", "https://[::1]:8443/base"} {
		allowed, err := NewOrigin(endpoint)
		if err != nil {
			t.Fatal(err)
		}
		u, err := url.Parse(endpoint + "/operations/1")
		if err != nil {
			t.Fatal(err)
		}
		if err := allowed.Validate(u); err != nil {
			t.Fatal(err)
		}
	}
}
```

- [ ] **Step 2: Run the tests and record the red result.**

```bash
./dev.sh -c 'cd v2 && go test -count=1 ./internal/httpclient -run "^TestOrigin"'
```

Expected: compilation fails because the new interfaces do not exist.

- [ ] **Step 3: Create `origin.go`.**

```go
package httpclient

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

// DestinationError identifies a rejected destination without retaining a full URL.
type DestinationError struct {
	Allowed     string
	Destination string
}

func (e *DestinationError) Error() string {
	return fmt.Sprintf("request destination %s is not allowed; expected %s", e.Destination, e.Allowed)
}

// Origin is an immutable HTTPS hostname and effective port.
type Origin struct {
	host string
	port string
}

// NewOrigin validates a configured HTTPS endpoint.
func NewOrigin(endpoint string) (Origin, error) {
	u, err := url.Parse(endpoint)
	if err == nil {
		if origin, ok := originOf(u); ok {
			return origin, nil
		}
	}
	return Origin{}, &DestinationError{
		Allowed: "a valid HTTPS origin", Destination: "<invalid HTTPS origin>",
	}
}

func (o Origin) String() string {
	return "https://" + net.JoinHostPort(o.host, o.port)
}

// Validate checks a destination without comparing its path or query.
func (o Origin) Validate(destination *url.URL) error {
	candidate, ok := originOf(destination)
	if ok && o.host != "" && candidate == o {
		return nil
	}
	safe := "<invalid HTTPS origin>"
	if ok {
		safe = candidate.String()
	}
	return &DestinationError{Allowed: o.String(), Destination: safe}
}

func originOf(u *url.URL) (Origin, bool) {
	if u == nil || !strings.EqualFold(u.Scheme, "https") || u.Opaque != "" || u.User != nil || u.Host == "" {
		return Origin{}, false
	}
	authority, err := url.Parse("https://" + u.Host)
	if err != nil || authority.Host != u.Host || authority.User != nil {
		return Origin{}, false
	}
	host := authority.Hostname()
	if host == "" || strings.ContainsAny(host, "%\\/?#@ \t\r\n") {
		return Origin{}, false
	}
	if strings.Contains(host, ":") && net.ParseIP(host) == nil {
		return Origin{}, false
	}
	port := authority.Port()
	if port == "" {
		if strings.HasSuffix(authority.Host, ":") {
			return Origin{}, false
		}
		port = "443"
	}
	n, err := strconv.Atoi(port)
	if err != nil || n < 1 || n > 65535 {
		return Origin{}, false
	}
	return Origin{host: strings.ToLower(host), port: strconv.Itoa(n)}, true
}
```

- [ ] **Step 4: Rerun the exact test command; require PASS.**
- [ ] **Step 5: Commit only these files.**

```bash
git add v2/internal/httpclient/origin.go v2/internal/httpclient/origin_test.go
git commit -m "feat: add shared HTTPS origin validation" -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

### Task 2: Guard redirects without modifying shared clients

**Files:** Create `v2/internal/httpclient/redirects.go` and
`v2/internal/httpclient/redirects_test.go`.

**Interfaces:**
- Consumes: `Origin.Validate(*url.URL) error` and `*DestinationError` from Task 1.
- Produces: `NewRedirectGuardedClient(base *http.Client, allowed Origin) *Client`
  and `(*Client).Do(req *http.Request) (*http.Response, error)`.
- `Client` satisfies the Azure transport interface structurally, without importing
  azcore. It guards redirects; initial-request validation remains the ARM policy's
  responsibility.

- [ ] **Step 1: Create `redirects_test.go` with these offline tests.**

```go
package httpclient

import (
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestRedirects_OriginBoundary(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		target string
		ok     bool
	}{
		{"/done", true},
		{"https://management.example.invalid:443/done", true},
		{"https://other.example.invalid/done?secret=private", false},
		{"https://child.management.example.invalid/done", false},
		{"https://management.example.invalid:444/done", false},
		{"http://management.example.invalid/done", false},
		{"https://user:private@management.example.invalid/done", false},
	} {
		t.Run(c.target, func(t *testing.T) {
			calls := 0
			base := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				calls++
				status := http.StatusOK
				h := make(http.Header)
				if calls == 1 {
					status = http.StatusFound
					h.Set("Location", c.target)
				}
				return &http.Response{
					StatusCode: status, Header: h, Request: r,
					Body: io.NopCloser(strings.NewReader("{}")),
				}, nil
			})}
			client := NewRedirectGuardedClient(base, allowed)
			req, err := http.NewRequest(http.MethodGet, "https://management.example.invalid/start?secret=private", nil)
			if err != nil {
				t.Fatal(err)
			}
			resp, err := client.Do(req)
			if resp != nil {
				_ = resp.Body.Close()
			}
			if c.ok {
				if err != nil || calls != 2 {
					t.Fatalf("valid redirect failed: calls=%d error=%v", calls, err)
				}
			} else {
				var rejected *DestinationError
				if !errors.As(err, &rejected) || calls != 1 {
					t.Fatalf("redirect not blocked: calls=%d error=%v", calls, err)
				}
				if strings.Contains(err.Error(), "private") || strings.Contains(err.Error(), "?") {
					t.Fatalf("unsafe outer error: %v", err)
				}
			}
		})
	}
}

func TestRedirects_ClientSettingsAndCallbacks(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
	callbackErr := errors.New("caller rejected redirect")
	calls := 0
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	base := &http.Client{
		Timeout: time.Minute, Jar: jar, Transport: &http.Transport{},
		CheckRedirect: func(*http.Request, []*http.Request) error {
			calls++
			return callbackErr
		},
	}
	client := NewRedirectGuardedClient(base, allowed)
	if client.client == base || client.client.Timeout != base.Timeout || client.client.Transport != base.Transport || client.client.Jar != base.Jar {
		t.Fatal("client was not copied with settings intact")
	}
	req, err := http.NewRequest(http.MethodGet, "https://management.example.invalid/done", nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := client.client.CheckRedirect(req, nil); err != callbackErr || calls != 1 {
		t.Fatalf("caller callback lost: %v", err)
	}
	if err := base.CheckRedirect(req, nil); err != callbackErr || calls != 2 {
		t.Fatal("base client callback was modified")
	}
	base.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	client = NewRedirectGuardedClient(base, allowed)
	if err := client.client.CheckRedirect(req, nil); err != http.ErrUseLastResponse {
		t.Fatalf("ErrUseLastResponse not preserved: %v", err)
	}
	base.CheckRedirect = func(r *http.Request, _ []*http.Request) error {
		r.URL.Host = "other.example.invalid"
		return nil
	}
	client = NewRedirectGuardedClient(base, allowed)
	if err := client.client.CheckRedirect(req, nil); err == nil {
		t.Fatal("callback mutation bypassed origin check")
	}
}

func TestRedirects_DefaultLimit(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
	calls := 0
	base := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		calls++
		return &http.Response{
			StatusCode: http.StatusFound, Request: r,
			Header: http.Header{"Location": []string{"/loop"}},
			Body: io.NopCloser(strings.NewReader("")),
		}, nil
	})}
	req, err := http.NewRequest(http.MethodGet, "https://management.example.invalid/loop", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := NewRedirectGuardedClient(base, allowed).Do(req)
	if resp != nil {
		_ = resp.Body.Close()
	}
	if err == nil || calls != 10 {
		t.Fatalf("redirect limit changed: calls=%d error=%v", calls, err)
	}
}
```

- [ ] **Step 2: Run the red test command.**

```bash
./dev.sh -c 'cd v2 && go test -count=1 ./internal/httpclient -run "^TestRedirects"'
```

Expected: compilation fails for missing `NewRedirectGuardedClient`.

- [ ] **Step 3: Create `redirects.go`.**

```go
package httpclient

import (
	"errors"
	"net/http"
)

// Client preserves HTTP behaviour while checking redirected destinations.
type Client struct {
	client *http.Client
}

// NewRedirectGuardedClient copies base without changing its shared transport.
func NewRedirectGuardedClient(base *http.Client, allowed Origin) *Client {
	if base == nil {
		base = http.DefaultClient
	}
	client := *base
	previous := base.CheckRedirect
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if err := allowed.Validate(req.URL); err != nil {
			return err
		}
		if previous != nil {
			if err := previous(req, via); err != nil {
				return err
			}
		} else if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		return allowed.Validate(req.URL)
	}
	return &Client{client: &client}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	var rejected *DestinationError
	if errors.As(err, &rejected) {
		// http.Client wraps redirect errors in a url.Error containing the original URL.
		return resp, rejected
	}
	return resp, err
}
```

- [ ] **Step 4: Run all shared-package tests and require PASS.**

```bash
./dev.sh -c 'cd v2 && go test -count=1 -race ./internal/httpclient'
```

- [ ] **Step 5: Commit the redirect helper and tests.**

```bash
git add v2/internal/httpclient/redirects.go v2/internal/httpclient/redirects_test.go
git commit -m "feat: guard HTTP redirects with client-local origin boundaries" -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

### Task 3: Install the ARM policy and isolate the base transport

**Files:** Create `v2/internal/genericarmclient/destination_policy.go` and
`v2/internal/genericarmclient/destination_policy_test.go`. Modify
`v2/internal/genericarmclient/generic_client.go` and
`v2/internal/genericarmclient/data_plane_client_options_test.go`. Preserve
existing RP registration code; the new policy must be installed before its
pipeline is constructed.

**Interfaces:**
- Consumes: `httpclient.Origin`, `NewOrigin`, `NewRedirectGuardedClient`.
- Produces: `newDestinationPolicy(allowed httpclient.Origin) policy.Policy`.
- Preserves public `GenericClient` APIs; adds a private `httpClient *http.Client`
  field holding the original client.
- Produces offline test helpers in package `genericarmclient_test`:
  `boundaryClient(t *testing.T, endpoint string, send func(*http.Request) (*http.Response, error)) (*genericarmclient.GenericClient, *boundaryCredential)`,
  `boundaryResponse(r *http.Request, status int, headers http.Header, body string) *http.Response`,
  and `boundaryToken(t *testing.T, resultType string, state map[string]any) string`.

- [ ] **Step 1: Create `destination_policy_test.go`.**

```go
package genericarmclient_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/httpclient"
)

type boundaryCredential struct{ calls atomic.Int32 }

func (c *boundaryCredential) GetToken(context.Context, policy.TokenRequestOptions) (azcore.AccessToken, error) {
	c.calls.Add(1)
	return azcore.AccessToken{Token: "boundary-test-token", ExpiresOn: time.Now().Add(time.Hour)}, nil
}

type boundaryRoundTripper func(*http.Request) (*http.Response, error)

func (f boundaryRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func boundaryResponse(r *http.Request, status int, headers http.Header, body string) *http.Response {
	if headers == nil {
		headers = make(http.Header)
	}
	headers.Set("Content-Type", "application/json")
	return &http.Response{
		StatusCode: status, Header: headers, Request: r,
		Body: io.NopCloser(strings.NewReader(body)),
	}
}

func boundaryClient(
	t *testing.T,
	endpoint string,
	send func(*http.Request) (*http.Response, error),
) (*genericarmclient.GenericClient, *boundaryCredential) {
	t.Helper()
	g := NewGomegaWithT(t)
	cred := &boundaryCredential{}
	cfg := cloud.Configuration{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
		cloud.ResourceManager: {Endpoint: endpoint, Audience: "https://management.core.windows.net/"},
	}}
	client, err := genericarmclient.NewGenericClient(cfg, cred, &genericarmclient.GenericClientOptions{
		HTTPClient: &http.Client{Transport: boundaryRoundTripper(send)},
	})
	g.Expect(err).ToNot(HaveOccurred())
	return client, cred
}

func boundaryToken(t *testing.T, resultType string, state map[string]any) string {
	t.Helper()
	b, err := json.Marshal(map[string]any{"type": resultType, "token": state})
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func Test_DestinationPolicy_RejectsBeforeAuthentication(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	calls := 0
	client, cred := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
		calls++
		return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
	})
	pl, err := armruntime.NewPipeline("boundary-test", "v0.0.0", cred, runtime.PipelineOptions{}, client.ClientOptions())
	g.Expect(err).ToNot(HaveOccurred())
	req, err := runtime.NewRequest(context.Background(), http.MethodGet, "https://attacker.example.invalid/path?secret=private")
	g.Expect(err).ToNot(HaveOccurred())
	_, err = pl.Do(req)
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(cred.calls.Load()).To(BeZero())
	g.Expect(calls).To(BeZero())
	g.Expect(err.Error()).ToNot(ContainSubstring("private"))
}

func Test_DestinationPolicy_InvalidEndpoint(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	cfg := cloud.Configuration{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
		cloud.ResourceManager: {Endpoint: "http://management.example.invalid", Audience: "https://management.core.windows.net/"},
	}}
	_, err := genericarmclient.NewGenericClient(cfg, &boundaryCredential{}, nil)
	g.Expect(err).To(HaveOccurred())
}

func Test_DestinationPolicy_RedirectDiagnosticIsSafe(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	calls := 0
	client, _ := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
		calls++
		return boundaryResponse(r, http.StatusFound,
			http.Header{"Location": []string{"https://attacker.example.invalid/path?secret=private"}}, ""), nil
	})
	var result genericarmclient.GenericResource
	_, err := client.GetByID(context.Background(),
		"/subscriptions/s/resourceGroups/rg/providers/Microsoft.Fake/resources/r", "2020-01-01", &result)
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(calls).To(Equal(1))
	g.Expect(err.Error()).ToNot(ContainSubstring("private"))
	g.Expect(err.Error()).ToNot(ContainSubstring("boundary-test-token"))
}
```

Change the `Endpoint` in the first `Test_DataPlaneClientOptions` cloud
configuration from `server.URL` to:

```go
Endpoint: "https://management.example.invalid",
```

The test must still send its data-plane request to `server.URL`. This proves the
original transport reaches a different origin while retaining recording/TLS
settings and the existing user-agent policy.

- [ ] **Step 2: Run the red integration tests.**

```bash
./dev.sh -c 'cd v2 && go test -short -tags=noexit -count=1 ./internal/genericarmclient -run "^Test_(DestinationPolicy|DataPlaneClientOptions)"'
```

Expected: the new boundary assertions fail against the unguarded ARM client.
The data-plane test may already pass; retain it as an isolation regression.

- [ ] **Step 3: Create `destination_policy.go`.**

```go
package genericarmclient

import (
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"

	"github.com/Azure/azure-service-operator/v2/internal/httpclient"
)

type destinationPolicy struct{ allowed httpclient.Origin }

func newDestinationPolicy(allowed httpclient.Origin) policy.Policy {
	return &destinationPolicy{allowed: allowed}
}

func (p *destinationPolicy) Do(req *policy.Request) (*http.Response, error) {
	if err := p.allowed.Validate(req.Raw().URL); err != nil {
		return nil, err
	}
	return req.Next()
}
```

In `generic_client.go`, add the import:

```go
"github.com/Azure/azure-service-operator/v2/internal/httpclient"
```

Replace the `GenericClient` declaration with:

```go
type GenericClient struct {
	endpoint   string
	pl         runtime.Pipeline
	creds      azcore.TokenCredential
	opts       *arm.ClientOptions
	userAgent  string
	httpClient *http.Client
}
```

Immediately after the existing empty-endpoint check, add:

```go
allowed, err := httpclient.NewOrigin(rmConfig.Endpoint)
if err != nil {
	return nil, eris.Wrap(err, "invalid Resource Manager endpoint")
}
```

In the initial `opts.ClientOptions.PerCallPolicies` literal, use:

```go
PerCallPolicies: []policy.Policy{
	newDestinationPolicy(allowed),
	NewUserAgentPolicy(ua),
},
```

Replace the existing HTTP-client assignment block, before
`NewRPRegistrationPolicy`, with:

```go
httpClient := options.HTTPClient
if httpClient == nil {
	httpClient = defaultHTTPClient
}
opts.Transport = httpclient.NewRedirectGuardedClient(httpClient, allowed)
```

Add this entry to the returned `GenericClient` literal:

```go
httpClient: httpClient,
```

In `DataPlaneClientOptions`, replace its transport entry with:

```go
Transport: client.httpClient,
```

Replace the `DataPlaneClientOptions` doc comment with:

```go
// DataPlaneClientOptions returns options for an Azure data-plane SDK client.
// Cloud, retry, logging, telemetry, and user-agent settings match the ARM client.
// The original HTTP client and its connection pool are shared without ARM-specific
// policies or redirect restrictions. The transport must not be modified by callers.
// Logging settings remain shared, including request and response body logging.
```

Do not move guard installation below `NewRPRegistrationPolicy`: that function
creates its own pipeline from the supplied options. Existing prepend/append
operations for registration and metrics remain unchanged.

- [ ] **Step 4: Run the targeted tests, including existing RP registration and actions.**

```bash
./dev.sh -c 'cd v2 && go test -short -tags=noexit -count=1 ./internal/httpclient ./internal/genericarmclient -run "TestOrigin|TestRedirects|Test_(DestinationPolicy|DataPlaneClientOptions|NewResourceGroup_SubscriptionNotRegisteredError|BeginPostActionByID|ResumeActionPoller)"'
```

Expected: PASS, including the existing local TLS server tests. The source
registration test is `Test_NewResourceGroup_SubscriptionNotRegisteredError`.

- [ ] **Step 5: Commit only the ARM integration and its tests.**

```bash
git add v2/internal/genericarmclient/destination_policy.go v2/internal/genericarmclient/destination_policy_test.go v2/internal/genericarmclient/generic_client.go v2/internal/genericarmclient/data_plane_client_options_test.go
git commit -m "fix: restrict authenticated ARM requests to their configured origin" -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

### Task 4: Prove poller, SDK, cloud and recovery coverage

**Files:** Create
`v2/internal/genericarmclient/destination_poller_test.go`,
`v2/internal/genericarmclient/destination_sdk_test.go`, and
`v2/internal/reconcilers/arm/destination_failure_test.go`.

**Interfaces:**
- Consumes the three `boundary*` helpers from Task 3 and
  `httpclient.NewOrigin(string) (Origin, error)`.
- Uses existing `ResumeCreatePoller`, `ResumeDeletePoller`, `ResumeActionPoller`,
  `PollerResponse.Resume`, `ClientOptions`, and reconciler failure handlers.
- Produces regression coverage only; no additional production control flow.

SDK token-shaped fixtures are deliberately confined to tests. They exercise
azcore v1.23.3's actual resume factory, including forms that ASO would not
ordinarily generate. Do not make production enforcement depend on these fields.

- [ ] **Step 1: Create `destination_poller_test.go`.**

```go
package genericarmclient_test

import (
	"context"
	"errors"
	"maps"
	"net/http"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/httpclient"
)

func Test_DestinationPoller_RejectsEveryPollingForm(t *testing.T) {
	t.Parallel()
	states := map[string]map[string]any{
		"async": {"asyncURL": "https://attacker.example.invalid/poll?secret=private", "state": "InProgress"},
		"operation": {"oplocURL": "https://attacker.example.invalid/poll", "state": "InProgress"},
		"location": {"type": "loc", "pollURL": "https://attacker.example.invalid/poll", "state": "InProgress"},
		"body": {"type": "body", "pollURL": "https://attacker.example.invalid/poll", "state": "InProgress"},
		"fake": {"fakeURL": "https://attacker.example.invalid/poll", "apiName": "test", "status": "InProgress"},
	}
	for name, state := range states {
		t.Run(name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			calls := 0
			client, cred := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
				calls++
				return boundaryResponse(r, http.StatusOK, nil, `{"status":"Succeeded"}`), nil
			})
			token := boundaryToken(t, "GenericResource", state)
			resumed := client.ResumeCreatePoller(genericarmclient.CreatePollerID)
			err := resumed.Resume(context.Background(), client, token)
			var rejected *httpclient.DestinationError
			g.Expect(errors.As(err, &rejected)).To(BeTrue())
			g.Expect(cred.calls.Load()).To(BeZero())
			g.Expect(calls).To(BeZero())
			g.Expect(err.Error()).ToNot(ContainSubstring("private"))
		})
	}
}

func Test_DestinationPoller_FinalAndUnusedURLs(t *testing.T) {
	t.Parallel()
	const endpoint = "https://management.example.invalid"
	const offOrigin = "https://regional.example.invalid/result"
	cases := []struct {
		name       string
		state      map[string]any
		response   string
		rejected   bool
		wantCalls  int
	}{
		{"async original", map[string]any{"asyncURL": endpoint + "/poll", "origURL": offOrigin, "method": "PUT"}, `{"status":"Succeeded"}`, true, 1},
		{"async location", map[string]any{"asyncURL": endpoint + "/poll", "locURL": offOrigin, "method": "POST"}, `{"status":"Succeeded"}`, true, 1},
		{"operation original", map[string]any{"oplocURL": endpoint + "/poll", "origURL": offOrigin, "method": "PUT"}, `{"status":"Succeeded"}`, true, 1},
		{"operation location", map[string]any{"oplocURL": endpoint + "/poll", "locURL": offOrigin, "method": "POST", "finalState": "location"}, `{"status":"Succeeded"}`, true, 1},
		{"operation response location", map[string]any{"oplocURL": endpoint + "/poll", "method": "POST"}, `{"status":"Succeeded","resourceLocation":"https://regional.example.invalid/result"}`, true, 1},
		{"unused regional location", map[string]any{"asyncURL": endpoint + "/poll", "origURL": endpoint + "/resource", "locURL": offOrigin, "method": "PUT"}, `{"status":"Succeeded"}`, false, 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			calls := 0
			client, cred := boundaryClient(t, endpoint, func(r *http.Request) (*http.Response, error) {
				calls++
				g.Expect(r.URL.Host).To(Equal("management.example.invalid"))
				body := c.response
				if r.URL.Path == "/resource" {
					body = "{}"
				}
				return boundaryResponse(r, http.StatusOK, nil, body), nil
			})
			state := maps.Clone(c.state)
			state["state"] = "InProgress"
			resumed := client.ResumeCreatePoller(genericarmclient.CreatePollerID)
			err := resumed.Resume(context.Background(), client, boundaryToken(t, "GenericResource", state))
			if c.rejected {
				var rejected *httpclient.DestinationError
				g.Expect(errors.As(err, &rejected)).To(BeTrue())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(resumed.Poller.Done()).To(BeTrue())
			}
			g.Expect(calls).To(Equal(c.wantCalls))
			g.Expect(cred.calls.Load()).To(Equal(int32(1)))
		})
	}
}

func Test_DestinationPoller_ResponseCannotExpandLocationOrigin(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	calls := 0
	client, _ := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
		calls++
		return boundaryResponse(r, http.StatusAccepted,
			http.Header{"Location": []string{"https://attacker.example.invalid/next"}}, "{}"), nil
	})
	state := map[string]any{"type": "loc", "pollURL": "https://management.example.invalid/poll", "state": "InProgress"}
	resumed := client.ResumeDeletePoller(genericarmclient.DeletePollerID)
	g.Expect(resumed.Resume(context.Background(), client, boundaryToken(t, "GenericDeleteResponse", state))).To(Succeed())
	token, err := resumed.Poller.ResumeToken()
	g.Expect(err).ToNot(HaveOccurred())
	next := client.ResumeDeletePoller(genericarmclient.DeletePollerID)
	err = next.Resume(context.Background(), client, token)
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(calls).To(Equal(1))
}

func Test_DestinationPoller_ActionUsesSameBoundary(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	calls := 0
	client, cred := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
		calls++
		return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
	})
	state := map[string]any{"asyncURL": "https://attacker.example.invalid/poll", "state": "InProgress"}
	resumed := client.ResumeActionPoller(genericarmclient.ActionPollerID)
	err := resumed.Resume(context.Background(), client, boundaryToken(t, "GenericActionResponse", state))
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(calls).To(BeZero())
	g.Expect(cred.calls.Load()).To(BeZero())
}
```

- [ ] **Step 2: Create `destination_sdk_test.go`.**

```go
package genericarmclient_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	. "github.com/onsi/gomega"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"

	"github.com/Azure/azure-service-operator/v2/internal/httpclient"
)

func Test_DestinationSDK_BackupInstanceResume(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	calls := 0
	client, cred := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
		calls++
		return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
	})
	factory, err := armdataprotection.NewClientFactory("subscription", cred, client.ClientOptions())
	g.Expect(err).ToNot(HaveOccurred())
	state := map[string]any{"asyncURL": "https://attacker.example.invalid/poll", "state": "InProgress"}
	poller, err := factory.NewBackupInstancesClient().BeginSyncBackupInstance(
		context.Background(), "rg", "vault", "instance", armdataprotection.SyncBackupInstanceRequest{},
		&armdataprotection.BackupInstancesClientBeginSyncBackupInstanceOptions{
			ResumeToken: boundaryToken(t, "BackupInstancesClientSyncBackupInstanceResponse", state),
		})
	g.Expect(err).ToNot(HaveOccurred())
	_, err = poller.Poll(context.Background())
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(calls).To(BeZero())
	g.Expect(cred.calls.Load()).To(BeZero())
}

func Test_DestinationSDK_CloudBoundariesAreIndependent(t *testing.T) {
	t.Parallel()
	endpoints := []string{
		"https://management.azure.com", "https://management.chinacloudapi.cn",
		"https://management.usgovcloudapi.net", "https://custom.example.invalid:8443/base",
	}
	for _, endpoint := range endpoints {
		t.Run(endpoint, func(t *testing.T) {
			g := NewGomegaWithT(t)
			calls := 0
			client, cred := boundaryClient(t, endpoint, func(r *http.Request) (*http.Response, error) {
				calls++
				return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
			})
			pl, err := armruntime.NewPipeline("boundary-test", "v0.0.0", cred, runtime.PipelineOptions{}, client.ClientOptions())
			g.Expect(err).ToNot(HaveOccurred())
			req, err := runtime.NewRequest(context.Background(), http.MethodGet, endpoint+"/operations/1")
			g.Expect(err).ToNot(HaveOccurred())
			resp, err := pl.Do(req)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(resp.Body.Close()).To(Succeed())
			other, err := runtime.NewRequest(context.Background(), http.MethodGet, "https://different.example.invalid/operations/1")
			g.Expect(err).ToNot(HaveOccurred())
			_, err = pl.Do(other)
			var rejected *httpclient.DestinationError
			g.Expect(errors.As(err, &rejected)).To(BeTrue())
			g.Expect(calls).To(Equal(1))
		})
	}
}
```

- [ ] **Step 3: Create `destination_failure_test.go`.**

```go
package arm

import (
	"net/url"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/httpclient"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func Test_DestinationFailure_ClearsLROForNormalRecovery(t *testing.T) {
	t.Parallel()
	for _, deleting := range []bool{false, true} {
		name := "create"
		if deleting {
			name = "delete"
		}
		t.Run(name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			scheme := runtime.NewScheme()
			g.Expect(resources.AddToScheme(scheme)).To(Succeed())
			obj := &resources.ResourceGroup{}
			obj.SetGroupVersionKind(resources.GroupVersion.WithKind("ResourceGroup"))
			SetPollerResumeToken(obj, genericarmclient.CreatePollerID, "saved-operation")
			r := &azureDeploymentReconcilerInstance{
				ARMOwnedResourceReconcilerCommon: reconcilers.ARMOwnedResourceReconcilerCommon{
					ResourceResolver: resolver.NewResolver(NewFakeKubeClient(scheme)),
				},
				Obj: obj, Log: logr.Discard(),
			}
			allowed, err := httpclient.NewOrigin("https://management.example.invalid")
			g.Expect(err).ToNot(HaveOccurred())
			target, err := url.Parse("https://attacker.example.invalid/path?secret=private")
			g.Expect(err).ToNot(HaveOccurred())
			err = allowed.Validate(target)
			g.Expect(err).To(HaveOccurred())
			if deleting {
				err = r.handleDeleteFailed(err)
			} else {
				err = r.handleCreateOrUpdateFailed(err)
			}
			conditionError, ok := conditions.AsReadyConditionImpactingError(err)
			g.Expect(ok).To(BeTrue())
			g.Expect(conditionError.Severity).To(Equal(conditions.ConditionSeverityWarning))
			g.Expect(err.Error()).ToNot(ContainSubstring("private"))
			_, _, hasToken := GetPollerResumeToken(obj)
			g.Expect(hasToken).To(BeFalse())
			g.Expect(obj.GetAnnotations()).ToNot(HaveKey(reconcilers.PollerResumeTokenAnnotation))
			g.Expect(obj.GetAnnotations()).ToNot(HaveKey(reconcilers.PollerResumeIDAnnotation))
		})
	}
}
```

- [ ] **Step 4: Run the regression suite and require PASS.**

```bash
./dev.sh -c 'cd v2 && go test -short -tags=noexit -count=1 ./internal/httpclient ./internal/genericarmclient ./internal/reconcilers/arm -run "TestOrigin|TestRedirects|Test_(Destination|DataPlaneClientOptions|NewResourceGroup_SubscriptionNotRegisteredError|BeginPostActionByID|ResumeActionPoller)"'
```

These are coverage tests for Task 3's integration, so they may pass immediately.
Prove their red side before claiming a regression is covered: temporarily remove
only `newDestinationPolicy(allowed)` from the initial ARM per-call list and rerun
`Test_DestinationPoller` and `Test_DestinationSDK`; require failures because
credentials/transport were reached or the typed rejection was absent. Restore
that exact line and rerun the suite. Do not use a destructive git command or
change another contributor's work for this check.

Do not treat the final-GET credential count alone as proof that authentication
was skipped: the SDK caches tokens. The inherited-pipeline test proves policy
ordering; final-GET tests prove no off-origin transmission after a legitimate
poll acquired a token.

- [ ] **Step 5: Commit the three regression files.**

```bash
git add v2/internal/genericarmclient/destination_poller_test.go v2/internal/genericarmclient/destination_sdk_test.go v2/internal/reconcilers/arm/destination_failure_test.go
git commit -m "test: cover ARM LRO destinations and existing recovery" -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

### Task 5: Document the boundary and verify playback and required checks

**Files:** Modify
`docs/hugo/content/guide/aso-controller-settings-options.md:210-222`.
Do not modify recordings or generated files manually.

**Interfaces:**
- Consumes the completed ARM implementation and regression suite.
- Produces endpoint documentation and recorded results for exact validation
  commands. No new runtime APIs.

- [ ] **Step 1: Add this paragraph after the Resource Manager endpoint description,
  before its Format field.**

```markdown
The endpoint must use HTTPS. ASO restricts ARM requests and redirects to this
endpoint's origin: its hostname and effective port. An omitted port is equivalent
to port 443. Polling and result URLs can use different paths on that origin.
Requests to another origin are rejected and reported through normal
reconciliation error handling. When multi-environment management supplies a
credential-specific endpoint, that endpoint defines the client's boundary.
```

- [ ] **Step 2: Invoke `testing-aso-recordings` before running playback. Confirm
  these exact recordings exist; do not allow missing fixtures to trigger live
  recording.**

```bash
test -f v2/internal/controllers/recordings/Test_CreateStorageAccountWithSkipReconcile_SecretsAreWritten.yaml
test -f v2/internal/controllers/recordings/Test_KubernetesConfiguration_Extension_20241101_CRUD.yaml
test -f v2/samples/sampletesting/recordings/Test_Samples_CreationAndDeletion/Test_Network_v1api20241001_CreationAndDeletion.yaml
test -f v2/samples/sampletesting/recordings/Test_Samples_CreationAndDeletion/Test_Kubernetesconfiguration_v1api20230501_CreationAndDeletion.yaml
```

The skill's old `internal/testsamples` references do not match this worktree;
the current Taskfile uses `v2/samples/sampletesting`. Use the commands below.
Do not source live credential files or request Azure credentials for playback.

- [ ] **Step 3: Run the two focused playback commands, separately.**

```bash
./dev.sh -c 'TEST_FILTER="^(Test_CreateStorageAccountWithSkipReconcile_SecretsAreWritten|Test_KubernetesConfiguration_Extension_20241101_CRUD)$" task controller:test-controllers'
./dev.sh -c 'TEST_FILTER="^Test_Samples_CreationAndDeletion$/(Test_Network_v1api20241001_CreationAndDeletion|Test_Kubernetesconfiguration_v1api20230501_CreationAndDeletion)$" task controller:test-samples'
```

Expected: both selected controller tests and both selected sample subtests PASS.
The task commands already write logs; do not add another `tee`. Inspect complete
command results and the selected PASS/FAIL lines in
`reports/test-controllers.log` and `reports/test-samples.log`. The task pipelines
must not turn a failing `go test` into a success claim. No selected recording may
be added or changed. If playback fails, invoke `diagnosing-vcr-failures` and fix
the root cause; do not re-record to conceal a regression.

- [ ] **Step 4: Run the required repository checks in order, one at a time.**

```bash
./dev.sh -c 'task format-code'
./dev.sh -c 'task generator:quick-checks'
./dev.sh -c 'task controller:quick-checks'
```

Expected: exit code zero for each. Allow builds/tests to finish; do not cancel
them. If any step fails, fix the issue and restart this three-command sequence
from formatting. Inspect formatter and generator changes; retain only changes
caused by this work. Do not revert another contributor's changes.

- [ ] **Step 5: Inspect the final diff and commit the documentation and any
  task-owned formatting corrections.**

```bash
git diff --check
git status --short
git diff --stat
git add docs/hugo/content/guide/aso-controller-settings-options.md
git commit -m "docs: explain ARM request origin enforcement" -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

If formatting changed task-owned Go files, stage those exact paths separately
before this commit rather than using `git add .`. All listed commits remain
local. Do not open a PR, push, create a GitHub stack, or start another client
implementation as part of this task.

## Completion evidence

The implementation handoff must identify the local branch/commits, exact test
commands and outcomes, and any unresolved compatibility or environment blocker.
Confirm that the rejected destinations received no requests, ordinary recovery
is unchanged, data-plane isolation passed, and selected recordings did not
change. This plan itself contains no execution results.
