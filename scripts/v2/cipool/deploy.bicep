param location string = resourceGroup().location

var poolName = 'aso-1es-pool'
var sku = 'Standard_D8ds_v4'

// Base 1ES Image Resource IDs
var ubuntu2204GalleryVersionResourceId = '/subscriptions/723b64f0-884d-4994-b6de-8960d049cb7e/resourceGroups/CloudTestImages/providers/Microsoft.Compute/galleries/CloudTestGallery/images/MMSUbuntu22.04-Secure/versions/latest'


var poolSettings = {
  maxPoolSize: 4 // We run 2 jobs per CI run on this pool, so a max of 4 means we can run 2 CI builds in parallel
  resourcePredictions: [
    {
      '21:00': 2  // 9 AM Monday NZT
    }
    {
      '05:00': 0  // 5 PM Monday NZT
      '16:00': 2  // 9 AM Monday PST
    }
    {
      '05:00': 0  // 5 PM Tuesday NZT
      '16:00': 2  // 9 AM Tuesday PST
    }
    {
      '05:00': 0  // 5 PM Wednesday NZT
      '16:00': 2  // 9 AM Wednesday PST
    }
    {
      '05:00': 0  // 5 PM Thursday NZT
      '16:00': 2  // 9 AM Thursday PST
    }
    {
      '05:00': 0  // 5 PM Friday NZT
      '16:00': 2  // 9 AM Friday PST
    }
    {
      '00:00': 0 // 5 PM Friday PST
    }
  ]
}

resource agentImage 'Microsoft.CloudTest/images@2020-05-07' = {
  name: '1es-ubuntu-22.04'
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
      type: 'GitHub'
      organizationName: 'Azure'
      url: 'https://github.com/Azure/azure-service-operator'
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
      resourcePredictions: poolSettings.resourcePredictions
    }
  }
}
