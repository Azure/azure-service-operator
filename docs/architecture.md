# Azure Service Operator - Architecture and Design

Below is a high level architecture of the Azure Service Operator.

<In progress - arch diagram>

## Deploying the operator - How it works

![](/docs/images/DeployFlow.png)

1. The user provisions the Managed Identity (or service principal) in their environment for the Azure Service Operator to use.
2. The user then uses the Helm chart to deploy the operator.
3. The Helm chart uses the image from the public Container Registry.
4. It deploys dependencies like [aad-pod-identity](https://github.com/Azure/aad-pod-identity) in addition to the Azure Service Operator.
5. This deployment takes care of deploying the operator while also configuring the needed Kubernetes RBAC rules.
6. Due to ordering of deployment, currently, cert-manager needs to be deployed manually before using the Helm chart to deploy the Operator.

## Resource provisioning - How it works

<resource provisioning diagram>

1. The developer deploys an application that includes the manifest for installing a Azure service that the app depends on.

2. The Azure Service Operator watches the CRDs and recognizes the request for a custom resource.

3. The Azure Seervice Operator creates the Kubernetes instance for the requested resource. It keeps the Kubernetes instance updated with the correct Status and with events.

4. The Azure Service Operator requests an authorizer from Azure Active Directory for the Azure resource management endpoint, as the identity it is running as.

5. The Azure Service Operator then sends the provisioning request to Azure API, along with the authorizer in the request.

6. Azure API provisions/deprovisions the resource and returns the Resource object to the Service Operator.

7. The Azure Service Operator retrieves the information required to access/consume the Azure resource from the Resource objecrt and stores it in a Kubernetes secret or as a secret in a specified Azure KeyVault.

8. The Azure Service Operator marks the Kubernetes instance for the resource as successfully provisioned.

## Security considerations

1. The Azure Service Operator requests the authorizer from AAD for every provision/deprovision request. There is no caching of security tokens.

2. Running the Azure Service Operator under a Managed Identity is recommended for security reasons. There is support to use Service Principals if needed, but not recommended.

3. It is recommended to use Azure KeyVault to store connection information, keys that are an output of the provisioning process. There is support to store as Kubernetes secrets but not recommended.

3. There is no implicit deletion of resources. Resources will be deprovisioned/deleted only on an explicit delete operation.