---
title: Handling secrets
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

**Example (from [the MySQL FlexibleServer sample](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/dbformysql/v1api/v1api20210501_flexibleserver.yaml)):**
```yaml
apiVersion: dbformysql.azure.com/v1alpha1api20210501
kind: FlexibleServer
metadata:
  name: samplemysql
  namespace: default
spec:
  location: westus2
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
and storing them in the [SecretDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination) you specify.

These secrets will be written to the destination(s) you specify once the resource has successfully been provisioned in Azure.
The resource will not move to [Condition](https://azure.github.io/azure-service-operator/guide/conditions/) `Ready=True` 
until the secrets have been written.

**Example:**
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

Azure Service Operator does not currently support rotation Azure generated credentials through Kubernetes.
ASO will (after some time) pick up rotations that happen via `az cli` or other tools. In order to avoid downtime
during the sync time, applications must have access to both the primary and secondary key and fall back from the 
primary to the secondary in the case authentication fails.
The recommended pattern for rotating credentials that support a primary and secondary key is to rotate the primary key with the `az cli` or Azure portal,
wait for 30m and then rotate the secondary key as well. As mentioned above, pods using the secrets ASO populates containing the
Azure generated secrets should mount the secrets as a volume so that they are automatically updated as soon as ASO picks up the new secret
value.  
