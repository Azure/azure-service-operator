## Deploying the operator - how it works

![](/docs/v1/images/Deploy%20Flow.png)

1. The user provisions the Managed Identity (or service principal) in their environment for the Azure Service Operator to use.
2. The user then deploys cert-manager for the Azure Service Operator to use when deployed.

3. The user then uses the Helm chart to deploy the operator. This includes the controller pod and the manager pod, created from images stored in the public Microsoft Container Registry. The Helm chart also deploys dependencies like [aad-pod-identity](https://github.com/Azure/aad-pod-identity) in addition to the Azure Service Operator. This deployment also configures the needed Kubernetes RBAC rules.


## Resource provisioning - how it works

![](/docs/v1/images/ASO%20flow.png)

1. The user deploys an application that includes the custom resource manifest for installing a Azure service that the app depends on.

2. The application is deployed using its manifest. However, the deployment does not yet succeed as it waits on the Azure service to be successfully created. The application references a secret that provides the information required to consume the Azure service, and the secret does not exist yet.

3. The Azure Service Operator continously watches the custom resource definitions (CRDs) corresponding to the Azure services and recognizes the request for a custom resource.

4. The Azure Service Operator then updates the Kubernetes instance for the requested resource with the correct status and events.

5. The Azure Service Operator requests an authorizer from Azure Active Directory for the Azure resource management endpoint, as the identity it is running as and receives an authorizer token.

6. The Azure Service Operator then sends the provisioning request to Azure API, along with the authorizer token in the request.

7. Azure API provisions/deprovisions the resource and returns the Resource object to the Service Operator.

8. The Azure Service Operator retrieves the information required to access/consume the Azure resource from the Resource object and stores it in a Kubernetes secret, or as a secret in a pre-specified Azure KeyVault.

9. The app is deployed successfully now that the Azure service it depends on is provisioned, and the secret it references exists.

## Security considerations

1. The Azure Service Operator requests the authorizer from AAD for every provision/deprovision request. There is no caching of security tokens.

2. Running the Azure Service Operator under a Managed Identity is recommended for security reasons. Azure Service Operator can also be run using service principals, but this is not recommended.

3. It is recommended to use Azure KeyVault to store connection information, AKA keys that are an output of the provisioning process. These connection strings can be stored as Kubernetes secrets, but this is not recommended.

4. There is no implicit deletion of resources. Resources will be deprovisioned/deleted only on an explicit delete operation.
