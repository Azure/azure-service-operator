param location string = resourceGroup().location

var poolName = 'asov1-1es-pool'
var sku = 'Standard_D4ds_v4' // Smaller VM size for ASOv1

// Base 1ES Image Resource IDs
var ubuntu2204GalleryVersionResourceId = '/subscriptions/723b64f0-884d-4994-b6de-8960d049cb7e/resourceGroups/CloudTestImages/providers/Microsoft.Compute/galleries/CloudTestGallery/images/MMSUbuntu22.04-Secure/versions/latest'

var poolSettings = {
  maxPoolSize: 2 // Keep the pool small for ASOv1
}

var msiName = 'asov1-ci-identity'

resource msi 'Microsoft.ManagedIdentity/userAssignedIdentities@2023-01-31' = {
  name: msiName
  location: location
}

// Followed https://devblogs.microsoft.com/devops/introduction-to-azure-devops-workload-identity-federation-oidc-with-terraform/
resource federatedIdentityCredential 'Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials@2023-01-31' = {
  name: '${msi.name}/fic'
  properties: {
    issuer: 'https://vstoken.dev.azure.com/4e7c52b2-a1d5-4777-9ab9-52d734b204cd'
    subject: 'sc://azure/azure-service-operator/ASO Subscription via MI' // From the ADO service connection UI
    audiences: [
      'api://AzureADTokenExchange'
    ]
  }
}

resource agentImage 'Microsoft.CloudTest/images@2020-05-07' = {
  name: 'asov1-1es-ubuntu-22.04'
  location: location
  properties: {
    imageType: 'SharedImageGallery'
    resourceId: ubuntu2204GalleryVersionResourceId
  }
}

resource hostedPool 'Microsoft.CloudTest/hostedpools@2020-05-07' = {
  name: poolName
  location: location
  properties: {
    organizationProfile: {
      type: 'AzureDevOps'
      url: 'https://dev.azure.com/azure'
      organizations: [
        {
          url: 'https://dev.azure.com/azure'
        }
      ]
    }
    sku: {
      name: sku
      tier: 'Standard' // Supports premium but we don't need it as work is done on temp disk anyway. See https://eng.ms/docs/cloud-ai-platform/devdiv/one-engineering-system-1es/1es-docs/1es-hosted-azure-devops-pools/demands.
    }
    images: [
      {
        subscriptionId: subscription().subscriptionId
        imageName: agentImage.name
        poolBufferPercentage: '100'
      }
    ]
    maxPoolSize: poolSettings.maxPoolSize
    agentProfile: {
      type: 'Stateless'
    }
    networkProfile: {
      firewallProfile: {
          policyName:  'Default' // allows access to everything apart from a list of known malicious endpoints
      }
    }
  }
  identity: {
    type: 'UserAssigned'
    userAssignedIdentities: {
      '${msi.id}': {}
    }
  }
}

var roleDefinitionId = '8e3af657-a8ff-443c-a75c-2fe8c4bcb635' // Owner

// This has to be in a module due to the following error:
// Error BCP120: This expression is being used in an assignment to the "name" property of the "Microsoft.Authorization/roleAssignments"
// type, which requires a value that can be calculated at the start of the deployment. Properties of msi which can be
// calculated at the start include "apiVersion", "id", "name", "type".
module roleAssignment 'identity.bicep' = {
  name: '${deployment().name}-identity'
  scope: subscription()
  params: {
    principalId: msi.properties.principalId
    principalType: 'ServicePrincipal'
    roleDefinitionId: roleDefinitionId
  }
}
