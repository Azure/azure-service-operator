# Enable Microsoft Entra-only authentication with separate resources

This sample creates a SQL server, its Microsoft Entra administrator, and its
Entra-only authentication setting as three separate ASO resources. You can apply
them together without waiting for each resource to become ready.

Both child resources depend on the server, but neither depends on the other.
Once the server is ready, they can reconcile concurrently. If authentication is
enabled before Azure has configured the administrator, Azure returns
`InvalidServerAADOnlyAuthNoAADAdminPropertyName`. ASO treats this error
case-insensitively as retryable, using its slow retry schedule (a few minutes,
with backoff and jitter), so the authentication setting can recover automatically.
There's no need to delete and recreate it.

The administrator is a user-assigned managed identity created by this sample.
ASO exports its principal ID and tenant ID to a ConfigMap consumed by
`ServersAdministrator`. This makes the demo independent of an existing Entra user.
The server deliberately doesn't configure `spec.administrators`: doing so would
avoid the separate-resource scenario demonstrated here.

## Run the sample

Create a resource group named `aso-sample-rg` using ASO, and create a Secret named
`sql-aadonly-password` with a `password` key containing a strong, unique password.
The server initially allows SQL authentication; Entra-only authentication disables
that access once the administrator has been configured. Don't use this initial
SQL-authentication window for a production deployment that must be Entra-only
from creation.

You also need a subscription whose policies allow that initial SQL-authentication
window. A policy requiring Entra-only authentication at server creation rejects
this scenario before either child resource can reconcile.

Apply the identity and all three SQL resources:

```bash
kubectl apply -f refs/
kubectl apply -f .
kubectl get servers,serversadministrators,serversazureadonlyauthentications
```

The sample test creates the resource group and password Secret automatically,
applies the resources together, waits for readiness, and deletes the resource
group. The race is timing-dependent: a successful run need not encounter the
missing-administrator error.

See [issue #4956](https://github.com/Azure/azure-service-operator/issues/4956) for
the original scenario.
