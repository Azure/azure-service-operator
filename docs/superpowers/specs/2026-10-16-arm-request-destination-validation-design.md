# ARM request destination validation

## Context

ASO must not send credentials to a destination selected through an editable LRO
resume-token annotation. The Azure SDK reconstructs polling and final-result
requests from that token, but its ARM bearer-token policy does not restrict their
destination to the configured Resource Manager endpoint.

The protection belongs on outgoing requests, rather than on individual fields in
the SDK's internal token format. This also covers destinations supplied by later
responses and ARM SDK clients used by customisations.

Existing recordings support using the configured ARM origin as the boundary.
Legitimate polling paths can be outside the resource's resource group, provider,
or resource-ID prefix. Some KubernetesConfiguration responses contain a regional
`Location` alongside an ARM `Azure-AsyncOperation`; the recorded operation uses
the ARM polling URL and the original resource URL, not that regional location.
Checking unused token fields would reject those operations unnecessarily.

## Requirements

- Validate each ARM request against the effective client's Resource Manager
  origin before bearer authentication.
- Reject off-origin redirects before sending the redirected request.
- Use credential-specific cloud configuration when present, falling back to
  global configuration through the existing client-selection mechanism.
- Protect generic operations, resumed pollers, final-result requests,
  resource-provider registration, and ARM SDK clients inheriting ASO options.
- Preserve connection pooling, existing HTTP client settings, ordinary retry
  behaviour, and data-plane client behaviour.
- Introduce shared HTTP infrastructure in the ARM PR without implementing Graph,
  Key Vault, or database enforcement in that PR.
- Keep all work local until coordinated disclosure is cleared.

Token signing, SDK-token schema validation, new credential mechanisms, and
unrelated cloud-support improvements are outside this design.

## Options

| Approach | Benefits | Limitations |
| --- | --- | --- |
| Pre-authentication SDK policy plus client-local redirect guard | Rejects untrusted initial requests before authentication; covers redirects; fits inherited ARM client options | Requires both pipeline and HTTP redirect integration |
| Transport-only enforcement | Checks actual destinations independently of the SDK pipeline | Authentication may already have run before rejection |
| Resume-token validation or signing | Rejects altered persisted state before rehydration | Parsing depends on SDK-internal schemas; signing needs key management; neither alone checks later destinations or redirects |

## Decision

Use a shared origin validator, an ARM-specific pre-authentication policy, and a
client-local redirect guard. The credential's configured destination is the
authority; annotations and service responses never expand it.

### Shared HTTP components

Add a small package under `v2/internal/httpclient` containing origin validation, a
recognisable destination-validation error, and a helper for creating a
redirect-guarded copy of an HTTP client. Keep these components independent of ARM
and the Azure SDK so later HTTP clients can reuse them.

The validator accepts a trusted endpoint at construction and checks parsed request
URLs against that endpoint's origin:

- Both the configured endpoint and request destination must be absolute HTTPS
  URLs with a syntactically valid authority, a non-empty hostname, and a valid port.
- Canonicalise non-ASCII hostnames using the IDNA lookup conversion used by Go's
  HTTP transport, then compare the resulting ASCII hostnames case-insensitively.
  Preserve the transport's ASCII fast path, without imposing DNS-label naming
  rules. Reject failed non-ASCII conversions rather than using partial results.
  Compare effective ports; an omitted HTTPS port and explicit port `443` are
  equivalent.
- Reject userinfo, opaque URLs, malformed authorities, and invalid ports.
- Do not use string-prefix matching, hostname suffix wildcards, DNS resolution,
  or paths and query strings to determine origin equality.
- Validate URL authority syntax, not DNS-label naming conventions. The configured
  endpoint is trusted; adding DNS-name restrictions does not strengthen exact
  origin containment and could unnecessarily constrain custom endpoints.
- Do not add an additional allowlist setting or an enforcement-disable switch.

The helper returns a shallow copy of the supplied HTTP client. It retains the
underlying transport and connection pool, timeouts, cookie jar, and other client
settings. It never modifies the supplied or default client.

Its redirect callback checks the resolved destination against the trusted origin
before the redirect is sent. Same-origin relative redirects remain valid.
Existing redirect callbacks remain effective and can reject an otherwise allowed
redirect; their return values, including `http.ErrUseLastResponse`, are
preserved. A callback must not be able to mutate a permitted destination into an
unvalidated one. When the original callback is nil, retain Go's normal
ten-redirect limit.

### ARM integration

`GenericClient` obtains the trusted endpoint from
`cloud.Configuration.Services[cloud.ResourceManager].Endpoint`, as it does
today. The audience determines token scopes, not allowed destinations. Do not
derive the boundary from `AZURE_AUTHORITY_HOST` or from object state.

Construct the origin validator and guarded HTTP client before creating any ARM
pipeline. Install the destination policy in per-call policies, before bearer
authentication. Ensure it is present when the separate resource-provider
registration pipeline is constructed, not only on the main generic pipeline.

Retain the unmodified HTTP client separately from the ARM-guarded copy:

| Surface | Policies and HTTP client |
| --- | --- |
| Generic ARM pipeline | ARM destination policy and guarded HTTP client |
| `ClientOptions()` | ARM options containing both protections for SDK clients |
| Resource-provider registration | Same ARM boundary on its separate pipeline |
| `DataPlaneClientOptions()` | Existing data-plane policies and original, unmodified HTTP client |

This preserves the existing recording transport and connection pooling without
making ARM's origin a global transport restriction. Data-plane clients continue
to receive the existing cloud, retry, logging, telemetry, and user-agent settings.

Coverage includes BackupInstance sync pollers constructed directly by the ARM
SDK, and databasewatcher action pollers, as well as the generic create/update and
delete resume paths. Do not add special-case URL validation to each customisation.

### Request flow and failure behaviour

For an initial request, poll, or final-result request, the ARM policy validates
the destination before passing control to authentication and transport. The
policy does not inspect unused fields inside the resume token.

For an HTTP redirect, `http.Client` resolves the target and invokes the guarded
callback. A rejected target receives no request, regardless of whether the
standard library would otherwise forward authentication headers.

An invalid configured endpoint fails client construction. A request or redirect
rejected by our validator returns a recognisable validation error through the
existing error path. Its diagnostics identify the rejected and allowed origins,
or a safe malformed destination description, without including resume tokens,
bearer headers, URL userinfo, or query values. Preserve this redaction through
HTTP/SDK wrapping of our validation errors; a safe inner error alone is
insufficient if an outer error includes a full URL.

Sanitisation is limited to errors raised by our validator. Do not change other
HTTP or SDK errors, including malformed `Location` headers rejected by Go before
the redirect callback runs. Those existing errors may retain URL details; this
change is not a general URL-redaction feature.

Retain existing reconciliation recovery. Generic ARM create/update and delete
failure handlers clear the saved LRO state, report the failure through existing
condition handling, and allow normal retries. Do not add a terminal condition,
preserve rejected tokens, automatically repair tokens, or change custom-action
recovery behaviour. Reissued operations may conflict, restart, or return a new
polling URL; existing reconciliation handles those outcomes.

### Verification

Use fake credentials, recording transports, and local TLS servers; security
regression tests must not require live Azure credentials. Observe token requests
and transport/server calls rather than merely checking an error string.

| Area | Required evidence |
| --- | --- |
| Origin comparison | Transport-equivalent Unicode/punycode hostnames, ASCII casing, configured ASCII custom authorities and default ports pass; unrelated hosts, Unicode case-folding lookalikes with different IDNA destinations, deceptive suffixes, changed ports, HTTP, userinfo, opaque URLs, malformed authorities and failed non-ASCII conversions fail |
| Initial requests | An off-origin request is rejected before credential acquisition or transport |
| Resumed operations | Tampered polling and final-result destinations cannot be sent; include SDK async, Location, Operation-Location, body, and fake-poller forms supported by the pinned dependency |
| Redirects | A rejected target receives zero requests; cover unrelated hosts, subdomains, changed ports, and HTTPS downgrades |
| Existing HTTP semantics | Same-origin absolute and relative redirects, existing stricter callbacks, `ErrUseLastResponse`, and the default redirect limit retain their behaviour |
| Inherited clients | Generic polling, inherited ARM SDK options, custom-action paths, and resource-provider registration retain the guard |
| Cloud selection | Public, sovereign, and custom TLS endpoints use the selected credential's endpoint; one client's boundary does not affect another |
| Data-plane isolation | Data-plane options keep the original HTTP client and can reach a different origin; neither construction nor redirect checking mutates the original client |
| Recovery and diagnostics | Generic failure handling still clears LRO state; reported errors remain visible without token, userinfo, or query disclosure |

When a resumed operation has already polled successfully, test the final-result
request independently: earlier legitimate calls may have acquired a token, but
the rejected final destination must not reach authentication or transport.

Use a real `net/http.Transport` with a dial callback that records the address and
always returns a controlled error to verify hostname canonicalisation without
DNS or external network access. Include China and Government Unicode lookalike
resume URLs, and positive Unicode/punycode configuration equivalence. The
lookalikes must produce a typed destination error before authentication, sending,
or dialing; independent unguarded controls establish Go's actual IDNA addresses.

Use the recordings in `v2/samples/sampletesting/recordings` and
`v2/internal/controllers/recordings` to check legitimate polling origins and
paths. Include a focused KubernetesConfiguration regression with an unused
off-origin `Location` so implementation does not drift into token-wide rejection.
Existing playback behaviour must remain intact; no live re-recording is planned.
The recordings do not establish compatibility for every service or cloud.

Run targeted tests during implementation, followed by the repository's required
checks in order: `./dev.sh -c 'task format-code'`,
`./dev.sh -c 'task generator:quick-checks'`, and
`./dev.sh -c 'task controller:quick-checks'`. If a required check fails, resolve
the cause and restart the sequence. Keep the configured Go proxy.

Update directly related controller-setting documentation during implementation
to explain the HTTPS/origin requirement and rejection behaviour. Do not publish
security details before disclosure clearance.

## Status

The user approved the design, written spec, and subagent-driven execution on
16 October 2026. The ARM implementation and verification are committed locally.
Whole-branch review is complete, including a final re-review that closed the
Unicode/IDNA finding. Public disclosure is not yet cleared.

Prepare six separate changes, each with its own design and implementation plan:

1. ARM and shared HTTP infrastructure, in this session's existing feature branch.
2. Microsoft Graph / Entra.
3. Key Vault.
4. Database connections in `dbformysql.azure.com`.
5. Database connections in `dbforpostgresql.azure.com`.
6. Database connections in `sql.azure.com`.

Ordinary branch and PR stacking is permitted. Do not use GitHub's new PR-stacking
features. Follow-ups that require only the shared work can branch directly from
ARM; they need not form a linear stack or wait for ARM to merge. Do not push
branches or create public PRs until the user clears coordinated disclosure.

## Consequences

The boundary covers URLs actually requested without depending on SDK token
internals. Existing tokens need no migration, and legitimate ARM path variations
remain supported.

An operation that genuinely needs an authenticated cross-origin ARM request will
now fail. Investigate such a compatibility failure explicitly; do not silently
widen the boundary or fall back to an unguarded client. The existing configuration
is trusted: this change does not constrain an administrator who can change the
credential's endpoint.

This is destination control, not a complete network-egress or DNS-rebinding
defence. Existing TLS validation, trusted proxy configuration, and credential
security remain required.

Graph has its own SDK and redirect stack, Key Vault has resource-specific origins,
and databases do not use HTTP. Their later designs must establish their own
trusted destinations rather than assuming this HTTP helper solves every case.

## Experience Report

The selected storage-account and KubernetesConfiguration controller recordings,
and the selected Network and KubernetesConfiguration sample recordings, replayed
successfully. All 460 controller/sample recording hashes remained unchanged.

Focused non-short race tests exercised the new shared HTTP, ARM/SDK/LRO, recovery,
and data-plane isolation regressions. The required formatting, generator checks,
and controller checks passed after task-owned test lint corrections. The existing
short-mode test harness skips `genericarmclient`, so its focused tests must run
without `-short`.

Whole-branch review found a Unicode hostname comparison gap that the original
tests missed. Unicode lowercasing equated a malicious hostname with a trusted ARM
hostname, while Go's transport selected a different IDNA destination. The
validator now applies the transport's lookup conversion before ASCII case
folding. Network-blocked tests reproduce the old authenticated transport escape
and confirm rejection before authentication with the fix. All four selected
playback tests, the non-short race regressions, and the required checks passed
again after this correction; all 460 recording hashes remained unchanged.

The worktree initially lacked its pinned REST API specifications submodule.
Restoring that dependency allowed normal type/CRD generation and playback; no
generated files or recordings were edited manually. These are local playback and
test results, not evidence from a live deployment.

## References

- [Security model and disclosure policy](../../../SECURITY.md)
- [Generic ARM client and inherited options](../../../v2/internal/genericarmclient/generic_client.go)
- [Existing HTTP client](../../../v2/internal/genericarmclient/default_http_client.go)
- [Generic poller rehydration](../../../v2/internal/genericarmclient/poller.go)
- [Resource-provider registration pipeline](../../../v2/internal/genericarmclient/policy_register_rp.go)
- [Credential-specific ARM cloud selection](../../../v2/internal/reconcilers/arm/arm_client_cache.go)
- [Reconciliation and failure handling](../../../v2/internal/reconcilers/arm/azure_generic_arm_reconciler_instance.go)
- [Data-plane options regression tests](../../../v2/internal/genericarmclient/data_plane_client_options_test.go)
- [BackupInstance SDK poller](../../../v2/api/dataprotection/customizations/backup_vaults_backup_instance_extensions.go)
- [Network sample polling paths](../../../v2/samples/sampletesting/recordings/Test_Samples_CreationAndDeletion/Test_Network_v1api20241001_CreationAndDeletion.yaml)
- [Storage controller polling paths](../../../v2/internal/controllers/recordings/Test_CreateStorageAccountWithSkipReconcile_SecretsAreWritten.yaml)
- [KubernetesConfiguration sample](../../../v2/samples/sampletesting/recordings/Test_Samples_CreationAndDeletion/Test_Kubernetesconfiguration_v1api20230501_CreationAndDeletion.yaml)
