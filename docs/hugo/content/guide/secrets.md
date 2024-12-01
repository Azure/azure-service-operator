---
title: Handling secrets
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

Resources supported by Azure Service Operator may require secrets as input (passwords, SSH keys, etc).
They may also produce secrets as "output" (storage keys, connection strings, etc). 

ASO has integration with [Kubernetes secrets](https://kubernetes.io/docs/concepts/configuration/secret/) to interact with 
these different types of secrets.

## How to provide secrets to Azure
Resources may have fields in their `spec` that expect a reference to a Kubernetes `Secret`. 
For example, in order to create a VM, you may want to specify an SSH password to enable SSH access to that VM.

The field in the `spec` will be a [SecretReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference),
which refers to a particular Kubernetes `Secret` key.

**Example (from [the MySQL FlexibleServer sample](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/dbformysql/v1api20230630/v1api20230630_flexibleserver.yaml)):**
```yaml
apiVersion: dbformysql.azure.com/v1api20230630
kind: FlexibleServer
metadata:
  name: samplemysql
  namespace: default
spec:
  location: eastus
  owner:
    name: aso-sample-rg
  version: "8.0.21"
  sku:
    name: Standard_D4ds_v4
    tier: GeneralPurpose
  administratorLogin: myAdmin
  administratorLoginPassword: # This is the name/key of a Kubernetes secret in the same namespace
    name: server-admin-pw
    key: password
  storage:
    storageSizeGB: 128
```

### Rotating credentials

Azure Service Operator is watching the referenced secret for changes, so rotating these credentials is as simple as 
editing the secret to contain a different credential. In the example above, if we update the content of 
the `server-admin-pw` secret `password` key to a new value, that will automatically change the password on the
`FlexibleServer` as well.

Note that if you have any applications that hardcoded the old password they will fail to authenticate until they
are also updated to the new password. This will also happen for any pods referring to the secret as an environment variable.
For this reason, we recommend if you expect to perform secret rotation to mount the secret as a pod volume instead.

## How to retrieve secrets created by Azure

Some Azure resources produce secrets themselves. ASO supports automatically querying these secrets 
and storing them in the [SecretDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination) 
you specify. This is done with the `.spec.operatorSpec.secrets` field, if you just need the secret itself,
or the [`.spec.operatorSpec.secretExpressions`]( {{< relref "expressions" >}} ) field, if you want the secret formatted in some way.

These secrets will be written to the destination(s) you specify once the resource has successfully been provisioned in Azure.
The resource will not move to [Condition]( {{< relref "conditions" >}} ) `Ready=True` 
until the secrets have been written.

**Example `.spec.operatorSpec.secretExpressions`:**
```yaml
apiVersion: documentdb.azure.com/v1alpha1api20210515
kind: DatabaseAccount
metadata:
  name: sample-db-account
  namespace: default
spec:
  location: westcentralus
  owner:
    name: aso-sample-rg
  kind: MongoDB
  databaseAccountOfferType: Standard
  locations:
    - locationName: westcentralus
  operatorSpec:
    secretExpressions:
      - name: mysecret
        key: primarymasterkey
        value: secret.primaryMasterKey
      - name: mysecret
        key: secondarymasterkey
        value: secret.secondaryMasterKey
      - name: myendpoint  # Can put different values into different Kubernetes secrets, if desired
        key: endpoint
        value: self.status.documentEndpoint
```

More complex expressions can be exported as well, see [Expressions]( {{< relref "expressions" >}} ) for more details.

**Example `.spec.operatorSpec.secrets`:**
```yaml
apiVersion: documentdb.azure.com/v1alpha1api20210515
kind: DatabaseAccount
metadata:
  name: sample-db-account
  namespace: default
spec:
  location: westcentralus
  owner:
    name: aso-sample-rg
  kind: MongoDB
  databaseAccountOfferType: Standard
  locations:
    - locationName: westcentralus
  operatorSpec:
    secrets:
      primaryMasterKey:
        name: mysecret
        key: primarymasterkey
      secondaryMasterKey:
        name: mysecret
        key: secondarymasterkey
      documentEndpoint: # Can put different secrets into different Kubernetes secrets, if desired
        name: myendpoint
        key: endpoint
```

### Rotating credentials

{{% alert title="Warning" color="warning" %}}
We recommend using Managed Identities wherever possible to avoid having to manually manage credentials.
{{% /alert %}}

Azure Service Operator does not currently support rotation Azure generated credentials through Kubernetes.
ASO will (after some time) pick up rotations that happen via `az cli` or other tools. In order to avoid downtime
during the sync time, applications must have access to both the primary and secondary key and fall back from the 
primary to the secondary in the case authentication fails.
The recommended pattern for rotating credentials that support a primary and secondary key is to rotate the primary 
key with the `az cli` or Azure portal, wait for 30m and then rotate the secondary key as well. 

As mentioned above, pods using the secrets ASO populates containing the
Azure generated secrets should mount the secrets as a volume so that they are automatically updated as soon as 
ASO picks up the new secret value.

Alternatively, instead of waiting 30m and having the application have logic to fallback to the other key during a 
rotation, a tool like [reloader](https://github.com/stakater/Reloader) can be used to automate the process.

The rotation process for an example resource such as CosmosDB then becomes: 

Starting with the following `operatorSpec`:
```yaml
operatorSpec:
  secrets:
    primaryMasterKey:
      name: cosmos-db
      key: key
```

1. Rotate the secondary key.
2. Update the `operatorSpec.secrets` `SecretDestination` to use the secondary key.
    ```yaml
    operatorSpec:
      secrets:
        primarySecondaryKey:
          name: cosmos-db
          key: key
    ```
3. Refresh workload (via `reloader`). Once workload pods are refreshed, proceed to next step.
4. Rotate the primary key.
5. Update the `operatorSpec.secrets` `SecretDestination` to use the primary key.
    ```yaml
    operatorSpec:
      secrets:
        primaryMasterKey:
          name: cosmos-db
          key: key
    ```
6. Refresh workload (via `reloader`).
