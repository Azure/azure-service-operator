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


1. The developer deploys an application that includes the manifest for installing a Azure service that the app depends on.

2. The Azure Service Operator watches the CRDs and recognizes the request as that for one of those custom resources.

3. The Azure Service Operator requests an authorizer from Azure Active Directory for the Azure resource management endpoint, as the identity it is running as.

4. 


## Security considerations


