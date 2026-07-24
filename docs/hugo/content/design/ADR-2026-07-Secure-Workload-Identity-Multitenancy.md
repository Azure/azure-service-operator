---
title: "2026-07: Secure Workload Identity for Single-Operator Multitenancy"
---

## Context

ASO supports [single-operator multitenancy]( {{< relref "credential-scope" >}} ) via a three-level credential hierarchy:

- **Global credentials** (via the `aso-controller-settings` secret in the operator namespace)
- **Per-namespace credentials** (via a fixed `aso-credential` secret in the resource's namespace)
- **Per-resource credentials** (via the `serviceoperator.azure.com/credential-from` annotation)

When using [Workload Identity](https://learn.microsoft.com/en-us/azure/aks/workload-identity-overview) for
authentication, each credential secret contains only `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, and `AZURE_CLIENT_ID`.
The operator then uses its own projected ServiceAccount token
(`system:serviceaccount:azureserviceoperator-system:azureserviceoperator-default`) to perform an OAuth2 token exchange
with Azure AD, presenting the `AZURE_CLIENT_ID` from the secret. This works because the Managed Identity has a
FederatedIdentityCredential registered for the ASO controller's ServiceAccount subject.

### The Security Problem

Unlike service principal and certificate authentication modes, none of the values in a Workload Identity credential 
secret (`AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`) are truly secret. They are GUIDs that can be discovered through:

- The Azure portal or CLI (e.g. `az identity show`)
- Azure Resource Graph queries
- Reading secrets in other namespaces if the user has RBAC access to do so
- Social engineering or documentation leaks

Because all namespaces share the same ServiceAccount token for the Azure token exchange, a user who can create ASO
resources in _any_ namespace can forge a credential secret referencing _any_ Managed Identity that has a
FederatedIdentityCredential configured for the ASO ServiceAccount. This is a **privilege escalation vulnerability**:
a tenant in namespace `team-a` can create a secret with `team-b`'s `AZURE_CLIENT_ID` and use `team-b`'s Azure identity
to manage Azure resources.

This has been raised as a security concern by multiple users
([#4810](https://github.com/Azure/azure-service-operator/issues/4810),
[#4807](https://github.com/Azure/azure-service-operator/issues/4807)):

The only current mitigation is deploying ASO in
[multi-operator multitenancy]( {{< relref "multitenant-deployment" >}} ) mode, where each namespace has its own
ASO pod and ServiceAccount. This provides true isolation because each pod's SA token can only be exchanged for the
Managed Identity it is federated with. However, multi-operator multitenancy is significantly harder to administer
(multiple CRD installs, webhook routing, independent upgrades) and does not scale well.

## Requirements

1. **Namespace-level credential isolation**: It must not be possible for a tenant in one namespace to use a Managed
   Identity that has been authorized for a different namespace, even if they know the Client ID.
2. **Privileged authorization**: The binding between a namespace and an Azure identity must be authorized by a user
   with appropriate privileges (cluster admin or similar), not merely by possession of credential values.
3. **Self-service bootstrapping**: Teams that have already created a Managed Identity with appropriate
   FederatedIdentityCredentials should be able to bootstrap their namespace without requiring write access to the
   `azureserviceoperator-system` namespace.
4. **Backward compatibility**: The existing credential hierarchy and secret format should continue to work for users
   who don't need the hardened mode. This should be opt-in.
5. **Workload Identity focus**: While the solution should be designed with extensibility in mind, the primary target
   is Workload Identity, as client secret and client certificate authentication modes have genuinely secret credential
   values that provide their own authorization factor.
6. **Single operator**: The solution must work within the single-operator multitenancy model (one ASO pod for the
   whole cluster).
7. **Manageable complexity**: The solution should be simpler to administer than multi-operator multitenancy — that is
   the entire motivation for single-operator mode.

## Options

### Option 1: ServiceAccount Impersonation (Per-Namespace ServiceAccounts for Azure Auth Only)

The ASO controller impersonates a namespace-scoped ServiceAccount to obtain a token for Azure AD token exchange,
instead of using its own ServiceAccount token directly.

**Setup (admin actions):**

1. Cluster admin creates a `ServiceAccount` in each tenant namespace (e.g., `aso-team-a` in namespace `team-a`).
2. Cluster admin creates a `FederatedIdentityCredential` in Azure for `team-a`'s Managed Identity with subject
   `system:serviceaccount:team-a:aso-team-a`.
3. Cluster admin creates an RBAC `Role`/`RoleBinding` or `ClusterRole`/`ClusterRoleBinding` granting ASO's controller
   ServiceAccount permission to create tokens for (`serviceaccounts/token`) the per-namespace SA.
4. The credential secret in `team-a` references the Client ID of `team-a`'s Managed Identity as before.

**Auth flow:**

1. User creates an ASO resource in namespace `team-a`.
2. ASO reads the credential secret from `team-a`, which includes the Client ID and the SA name (`aso-team-a`).
3. ASO uses the TokenRequest API to mint a short-lived token for `ServiceAccount team-a/aso-team-a` with audience
   `api://AzureADTokenExchange`. The token's subject is `system:serviceaccount:team-a:aso-team-a`.
4. ASO exchanges this token with Azure AD for an access token, presenting the Client ID from the secret.
5. Azure AD validates that the FederatedIdentityCredential on the Managed Identity matches the token's subject —
   since the FIC is scoped to `team-a`'s SA, other namespaces cannot use this identity even if they know the Client ID.

**Implementation details:**

- ASO would use the [Kubernetes TokenRequest API](https://kubernetes.io/docs/reference/kubernetes-api/authentication-resources/token-request-v1/)
  to create a short-lived token for the target namespace's SA with the `api://AzureADTokenExchange` audience.
- The controller needs `serviceaccounts/token: create` permission for the well-known SA name. With the fixed-name
  convention, this is a single `ClusterRole`/`ClusterRoleBinding` scoped via `resourceNames: ["aso-workload"]`,
  shipped as part of the ASO Helm chart.
- In the configurable variant (not chosen), a new field in the credential secret would indicate the SA name.
  With the fixed-name convention chosen in this design, no credential secret changes are needed.

**Variant: Fixed vs. configurable SA name**

There are two sub-variants for how the SA name is determined:

- **Configurable (per-secret):** The credential secret includes a field (`AZURE_WORKLOAD_IDENTITY_SERVICE_ACCOUNT`) specifying
  which SA to use. This is the most flexible — different credentials in the same namespace could use different SAs.
- **Fixed convention:** ASO uses a well-known SA name (e.g., `aso-workload`) in every namespace. No configuration
  is needed — if the SA exists and the RBAC grant is in place, ASO uses it automatically.

The fixed-name approach is simpler to administer and requires no changes to the credential secret format. The
configurable variant could be offered as a future extension if there's demand.

- When this mode is enabled, if ASO encounters a namespace without the designated SA or without the RBAC grant, it
  refuses to authenticate and sets a clear error condition on the resource.
- Note that the per-namespace SA only needs to exist and be tokenizable; it doesn't need any Kubernetes RBAC permissions itself.
- Both namespace-scoped credentials and per-resource credentials within that namespace would use the same SA for
  token exchange. Since Kubernetes RBAC is also namespace-scoped, there is no hard security boundary between these
  two levels within the same namespace — the isolation boundary is the namespace itself.
- The TokenRequest API call adds a small amount of latency to Azure authentication. However, the SA token can be
  requested with a configurable expiration (e.g., 1 hour) and cached, so the extra API call only occurs on initial
  auth and periodic refresh — not on every Azure API call.

**Pros:**

- Uses native Kubernetes primitives (ServiceAccount, TokenRequest API, RBAC) — no new CRDs.
- RBAC for `serviceAccounts/token` on the main ASO service account can be scoped to specific SA 
  names using `resourceNames`, if desired.
- The authorization gate is the FederatedIdentityCredential in Azure (which requires Azure-level permissions to
  create) and the RBAC grant allowing ASO to mint tokens for the per-namespace SA (which requires cluster-admin).
  Without both, ASO can't authenticate in that namespace. Since each serviceAccount must have a FederatedIdentityCredential
  bound to a unique subject (including that SA name + namespace), stealing a Client ID is useless since each Client 
  ID only works for the SA it is bound to.
- Credential secrets retain their existing format — no changes needed.
- Minimal change to ASO's Kubernetes RBAC model and controller architecture.
- Supports self-service bootstrapping: once the ClusterRole is installed with ASO, teams can create the SA in their
  own namespace and the FIC in Azure without needing write access to `azureserviceoperator-system`.

**Cons:**

- Requires cluster admin to create a ServiceAccount and RBAC binding per namespace — this is the "cost" of security,
  but it's the same cost ESO imposes and is significantly less than multi-operator multitenancy.
- Only addresses Workload Identity. Client secret/certificate auth modes don't benefit (though they don't need to,
  as their credentials are genuinely secret).
- Does not provide Kubernetes-level audit trail showing which SA performed which operations in which namespace.
- Each namespace requires its own FederatedIdentityCredential on the Managed Identity. Azure limits the number of
  FICs per identity (currently 20), so using the same identity across many namespaces will hit this limit. Users
  needing the same identity in more than 20 namespaces would need to create duplicate Managed Identities with the
  same role assignments, or use the [AKS Identity Bindings feature](https://learn.microsoft.com/en-us/azure/aks/identity-bindings-concepts).

### Option 2: ServiceAccount Impersonation for Both Azure Auth AND Kubernetes Operations

This option extends [Option 1](#option-1-serviceaccount-impersonation-per-namespace-serviceaccounts-for-azure-auth-only)
by also impersonating the per-namespace SA for _all_ Kubernetes operations in that namespace: reading secrets,
watching ASO resources, writing status, creating output secrets/configmaps — in addition to Azure token exchange.

**Pros:**

- All the pros of Option 1.
- Full isolation: each namespace's operations are performed under that namespace's SA identity.
- Kubernetes audit logs clearly attribute operations to the per-namespace SA.

**Cons:**

- Significantly more complex to implement. The ASO controller currently uses a single informer/cache for all
  namespaces. Impersonating different SAs for different namespaces would require per-namespace clients or
  impersonation headers on every API call.
- Each per-namespace SA would need its own RBAC grants for the ASO resource types, secrets, configmaps, and events
  in its namespace. This is a large amount of RBAC to administer.
- Performance implications: per-namespace informer caches or impersonation on every call adds overhead.
- Closer to multi-operator multitenancy in management complexity, undermining the simplicity benefit of
  single-operator mode.
- Does not meaningfully reduce the blast radius of an ASO pod compromise compared to Option 1. 
  In both cases, the ASO controller's own SA retains broad cluster-wide Kubernetes permissions — and in Option 2
  it also has RBAC to mint tokens for every managed namespace's SA, so compromising the ASO pod grants equivalent 
  access either way.
- Fundamentally changes ASO's controller architecture in ways that are difficult to make backward-compatible.

### Option 3: Admin-Controlled Credential CRD (ASO-Managed SecretStore Equivalent)

Introduce a new namespaced CRD (e.g., `ASOCredential` or `CredentialBinding`) that only cluster admins can create,
replacing the use of Secrets as the credential mechanism for Workload Identity. This is conceptually similar to
the approach used by the External Secrets Operator (see [Appendix: External Secrets Operator](#appendix-external-secrets-operator)
for a detailed comparison).

**Setup:**

1. Cluster admin creates an `ASOCredential` CR in each tenant namespace specifying the Client ID, Tenant ID,
   and Subscription ID.
2. ASO watches `ASOCredential` resources and uses them instead of (or in addition to) credential secrets.
3. RBAC for `ASOCredential` resources is restricted to cluster admins.
4. Optionally, ASO is configured with `AZURE_DISALLOW_WORKLOAD_IDENTITY_SECRETS=true` (in the global
   `aso-controller-settings` secret) that causes it to reject credential secrets containing Workload Identity
   configuration, ensuring `ASOCredential` CRDs are the only accepted path for Workload Identity.

**Pros:**

- Clear admin-controlled authorization boundary (CRD RBAC, not Secret RBAC).
- Auditable — CRD resources are visible via `kubectl get asocredentials`.
- Does not require per-namespace ServiceAccounts or RBAC for token creation.
- With secrets disabled for Workload Identity, the shared SA token is not exploitable — users have no way to
  get ASO to perform a token exchange with an unauthorized Client ID.
- Does not consume FederatedIdentityCredentials per namespace (avoids the 20 FIC-per-identity Azure limit).
- Does not require write access to `azureserviceoperator-system` for per-namespace setup, but does require a cluster
  admin to create the `ASOCredential` resource.

**Cons:**

- Adds a new CRD to ASO's already large CRD footprint.
- Requires a parallel credential resolution path alongside the existing Secret-based system. CRD cannot be used for 
  ServicePrincipal or Certificate based auth because those actually contain secret data and so must be stored in a 
  Kubernetes secret.
  - The CRD could theoretically hold ServicePrincipal or Certificate credentials too, but Kubernetes `Secret` resources benefit
    from encryption at rest (via `EncryptionConfiguration`), audit log redaction, and other special handling 
    that CRD specs do not receive by default.
- ASO already has a working Secret-based system and users using it for Workload Identity, making migration awkward.

### Option 4: Namespace Label/Annotation Allowlist

The ASO operator configuration includes an allowlist mapping Client IDs to allowed namespaces. The operator rejects
credential secrets whose Client ID is not explicitly allowed for the secret's namespace.

**Setup:**

1. Cluster admin configures ASO (via Helm values, global ConfigMap, or operator flags) with a mapping:
   `client-id-x → [namespace-a, namespace-b]`.
2. When ASO reads a credential secret, it checks the Client ID against the allowlist for that namespace.
3. If the Client ID is not allowed, reconciliation fails with a clear error.

**Pros:**

- No new CRDs or ServiceAccounts required.
- Centralized configuration — admin manages one allowlist.
- Simple conceptual model.

**Cons:**

- Centralized configuration does not scale: every namespace/identity change requires updating the operator config
  and potentially restarting the operator.
- Awkward to manage in GitOps workflows where namespace and identity provisioning are decentralized.
- If the allowlist is not enforced (i.e., made optional or left unconfigured), the security benefit is lost — same
  as Option 3 without `AZURE_DISALLOW_WORKLOAD_IDENTITY_SECRETS`.
- Does not support self-service bootstrapping: every namespace/identity change requires an admin to update the
  allowlist in `aso-controller-settings` (in the `azureserviceoperator-system` namespace).

### Option 5: Admission Webhook Validation

Deploy an admission webhook (built into ASO or external via OPA/Kyverno) that validates credential secrets on
creation/update, rejecting secrets with unauthorized Client IDs.

**Pros:**

- Can be implemented with existing tools (Kyverno, OPA Gatekeeper) without ASO changes.
- Preventive — blocks unauthorized secrets before they're created.

**Cons:**

- External policy engines are an additional dependency and management burden.
- Built-in webhook would need the same allowlist configuration as Option 4.
- Users could potentially work around the webhook (e.g., creating the secret before the webhook is deployed,
  or if the webhook is unavailable due to `failurePolicy: Ignore`).
- Not a defense-in-depth solution — it's a gate at one point, not a fundamental architectural fix.
- Policy management is typically centralized (cluster-scoped policies or operator namespace), limiting self-service
  bootstrapping by teams.

## Decision

We choose **Option 1: ServiceAccount Impersonation** (SA for Azure token exchange only), using the
**fixed convention** variant: ASO uses a well-known ServiceAccount name (`aso-workload`) in every namespace. This option will be
**disabled by default** and configured via the new `AZURE_WORKLOAD_IDENTITY_AUTH_MODE` setting in `aso-controller-settings`.
This setting accepts two values:

- **`relaxed`** (default): ASO uses its own SA token for all Workload Identity token exchanges, matching current behavior.
- **`strict`**: ASO requires a per-namespace `aso-workload` ServiceAccount for namespace-scoped and per-resource
  Workload Identity credentials. The global credential (`aso-controller-settings`) continues to use ASO's own SA
  token, since it is managed by cluster admins who are already trusted.

This option provides the strongest security guarantee: the FederatedIdentityCredential is bound to a per-namespace
ServiceAccount subject, so possessing a Client ID alone is insufficient to use an identity. The real authorization
boundary is twofold:

1. **Kubernetes side** (necessary but not the security boundary): The team must be able to create a ServiceAccount
   and a credential Secret in their namespace. These are common namespace-admin permissions. The ClusterRole shipped
   with ASO pre-authorizes token creation for the well-known SA name, so no per-namespace cluster-admin approval is
   needed.
2. **Azure side** (the real security boundary): The team must have Azure-level permissions to create a
   FederatedIdentityCredential on the Managed Identity, binding it to the namespace's SA subject. Without this,
   the token exchange will fail — even if the Kubernetes resources exist. This is the authorization gate that
   prevents unauthorized use of an identity.

Together these mean the Kubernetes setup is a prerequisite, but the actual trust decision is made in Azure: only
users who can configure a FIC for a given Managed Identity can authorize a namespace to use it.

Using a fixed SA name means no changes to the credential secret format are needed. The credential secret continues
to contain only `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, and `AZURE_CLIENT_ID` as before. It also means that 
a ClusterRole/ClusterRoleBinding can be automatically included at ASO install time enabling impersonation of specifically
those accounts, which makes management easy.

### Proposed Configuration

When `AZURE_WORKLOAD_IDENTITY_AUTH_MODE=strict` is set in the global `aso-controller-settings` secret,
ASO requires a ServiceAccount named `aso-workload` to exist in each namespace where namespace-scoped or per-resource
Workload Identity credentials are used. The global credential continues to use ASO's own SA token.
The credential secret format is unchanged from the current format.

When this mode is enabled:

1. ASO looks for a ServiceAccount named `aso-workload` in the resource's namespace.
2. ASO uses the TokenRequest API to mint a short-lived token for that SA with audience
   `api://AzureADTokenExchange`.
3. ASO uses this token (instead of its own projected SA token) for the OAuth2 token exchange with Azure AD.
4. The FederatedIdentityCredential on the Managed Identity must have subject
   `system:serviceaccount:<namespace>:aso-workload` (not the ASO controller's SA).

When `AZURE_WORKLOAD_IDENTITY_AUTH_MODE` is not set or is `relaxed`, ASO falls back to the current
behavior (using its own SA token for all namespaces), maintaining backward compatibility.

### Admin Setup Workflow

The RBAC grant for token creation can be done once at the cluster level using the fixed SA name. ASO ships this
`ClusterRole`/`ClusterRoleBinding` as part of its Helm chart when the feature is enabled:

```yaml
# Shipped with ASO (Helm chart) — grants ASO permission to create tokens
# for any SA named "aso-workload" in any namespace
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: aso-token-creator
rules:
  - apiGroups: [""]
    resources: ["serviceaccounts/token"]
    resourceNames: ["aso-workload"]
    verbs: ["create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: aso-token-creator
subjects:
  - kind: ServiceAccount
    name: azureserviceoperator-default
    namespace: azureserviceoperator-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: aso-token-creator
```

Per-namespace setup:

```bash
# 1. Create the well-known ServiceAccount in the tenant namespace
kubectl create serviceaccount aso-workload -n team-a

# 2. Create the FederatedIdentityCredential in Azure for team-a's Managed Identity
#    This can also be done using ASO itself, see:
#    https://azure.github.io/azure-service-operator/guide/authentication/#using-aso-to-create-its-own-credentials
az identity federated-credential create \
  --name team-a-aso \
  --identity-name team-a-identity \
  --resource-group rg-identities \
  --issuer "${AKS_OIDC_ISSUER}" \
  --subject "system:serviceaccount:team-a:aso-workload" \
  --audience "api://AzureADTokenExchange"

# 3. Create the credential secret (format unchanged)
kubectl apply -f - <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: aso-credential
  namespace: team-a
stringData:
  AZURE_SUBSCRIPTION_ID: "${SUBSCRIPTION_ID}"
  AZURE_TENANT_ID: "${TENANT_ID}"
  AZURE_CLIENT_ID: "${TEAM_A_CLIENT_ID}"
EOF
```

Because the `ClusterRole` uses `resourceNames: ["aso-workload"]`, ASO can only create tokens for SAs with that
exact name — it cannot mint tokens for arbitrary ServiceAccounts.

### FAQ

Q: Does this work with the global credential (`aso-controller-settings`)?

A: The global credential always uses ASO's own SA token, even in `strict` mode. The global credential is managed by
cluster admins who are already trusted, so per-namespace SA impersonation is unnecessary. `strict` mode only applies
to namespace-scoped and per-resource Workload Identity credentials.

Q: What about `ServicePrincipal` (client secret/certificate) credentials?

A: These authentication modes already use genuinely secret credential values (the client secret or certificate private
key) as the authorization factor. A user who doesn't know the client secret cannot forge a credential. The
ServiceAccount impersonation feature is specific to Workload Identity where the credential values are all
non-secret.

Q: What happens if the ServiceAccount doesn't exist or ASO doesn't have permission to create tokens for it?

A: ASO will fail the reconciliation with a clear error condition on the resource, indicating that the ServiceAccount
is missing or the RBAC is not configured. This is the same behavior as other credential resolution failures today.

Q: In `strict` mode, can a tenant's resources still fall back to the global credential?

A: Yes. If a namespace has no namespace-scoped or per-resource credential, ASO will still fall back to the global
credential (if configured), which uses the ASO controller's own SA token. This means tenant resources could be
managed using the global identity. Whether this is acceptable depends on the admin's intent — some admins
intentionally configure a global credential as a shared default, while security-conscious admins may want to prevent
this entirely. Disabling or restricting the global credential fallback is a separate concern that could be addressed
by a future configuration option (e.g., requiring every namespace to have an explicit credential), but it is
orthogonal to this design and does not affect the per-namespace isolation guarantees that `strict` mode provides. 
Today, if admins do not want to allow fallback to the global credential, just do not supply global credential details.

## Status

Proposed.

## Consequences

- When `AZURE_WORKLOAD_IDENTITY_AUTH_MODE=strict`, cluster admins (or teams with appropriate permissions) will need
  to create a ServiceAccount per tenant namespace. This is additional setup compared to today but significantly less
  than multi-operator multitenancy. In `relaxed` mode (the default), no changes are needed.
- Users who don't use Workload Identity or don't need cross-tenant isolation are unaffected.
- Documentation for the [credential scope guide]( {{< relref "credential-scope" >}} ) and
  [Workload Identity setup]( {{< relref "credential-format" >}} ) will need updates.
- The `AZURE_WORKLOAD_IDENTITY_AUTH_MODE=strict` setting provides a cluster-wide enforcement mechanism for security-conscious
  organizations.

## Experience Report

TBC

## References

- [#4810: Improve single-operator multitenancy credential management for Workload Identity](https://github.com/Azure/azure-service-operator/issues/4810)
- [#4807: What is the best practice for multi-tenancy when using workload identity?](https://github.com/Azure/azure-service-operator/issues/4807)
- [#3645: Support for Namespace scoped RoleAssignments that mitigate privilege escalation risks](https://github.com/Azure/azure-service-operator/issues/3645)
- [ADR-2022-09: Support For Multiple Credentials Under Global Operator]( {{< relref "ADR-2022-09-Multiple-Credential-Operator" >}} )
- [External Secrets Operator — Multi Tenancy Guide](https://external-secrets.io/latest/guides/multi-tenancy/)
- [External Secrets Operator — Security Best Practices](https://external-secrets.io/latest/guides/security-best-practices/)
- [Kubernetes TokenRequest API](https://kubernetes.io/docs/reference/kubernetes-api/authentication-resources/token-request-v1/)
- [Azure Workload Identity Overview](https://learn.microsoft.com/en-us/azure/aks/workload-identity-overview)

## Appendix: External Secrets Operator

The [External Secrets Operator](https://external-secrets.io/) (ESO) faces a similar challenge: a cluster-wide
controller that needs to authenticate to external APIs on behalf of different tenants. ESO's approach provides useful
contrast for this design.

### How ESO Works

ESO introduces dedicated CRDs to model the authentication boundary:

- **`SecretStore`** (namespaced): Defines _how_ to access an external secret provider, including auth credentials.
  A `SecretStore` is scoped to a single namespace and cannot reference resources across namespaces.
- **`ClusterSecretStore`** (cluster-scoped): A cluster-wide store that can be referenced from any namespace, with
  optional `namespaceSelector` conditions to restrict which namespaces may use it.
- **`ExternalSecret`** (namespaced): References a `SecretStore` or `ClusterSecretStore` and declares what to fetch.

For service-account-based auth (e.g. Vault Kubernetes auth), ESO requests a token for a `serviceAccountRef` defined
in the `SecretStore`. ESO's RBAC includes `serviceaccounts/token: create` permission, which can be scoped per-SA
using `resourceNames` constraints for hardening.

### ESO's Multitenancy Models

ESO documents three multitenancy models:

1. **Shared `ClusterSecretStore`**: A single CSS with `namespaceSelector` conditions. Simple but coarse-grained.
2. **Managed `SecretStore` per Namespace**: Cluster admins create namespaced `SecretStore` resources with
   per-namespace credentials/roles. Access is governed by the external API's own role system.
3. **ESO as a Service**: Tenants manage their own `SecretStore` and `ExternalSecret` resources autonomously.

### Comparison with ASO's Situation

| Aspect              | ESO                                                         | ASO (current)                                                     |
| ------------------- | ----------------------------------------------------------- | ----------------------------------------------------------------- |
| Auth configuration  | Dedicated CRD (`SecretStore`) created by admin              | Kubernetes `Secret` creatable by any user with secret-create RBAC |
| Namespace isolation | `SecretStore` enforces namespace boundary by design         | Secret is namespace-scoped but contents are not secret            |
| Admin control point | CRD creation requires explicit RBAC grant                   | Secret creation is a common permission many users already have    |
| SA token exchange   | ESO requests tokens for per-store SAs (`serviceAccountRef`) | ASO uses its own single SA for all namespaces                     |
| Audit trail         | `SecretStore` is a visible, auditable API resource          | Opaque secret contents are not easily auditable                   |

The key architectural difference is that ESO uses a **dedicated CRD** as the admin-controlled authorization boundary,
while ASO uses **Kubernetes Secrets**. Since Secrets are a general-purpose resource that many users have permission to
create, they do not serve as an effective authorization gate for Workload Identity where the credential values
themselves are not secret.
