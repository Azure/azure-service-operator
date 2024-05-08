// ======
// Params
// ======
@description('The principalId of the identity')
param principalId string

@description('The principal type of the identity')
param principalType string

@description('The roleDefinitionId of the role to grant')
param roleDefinitionId string


// ==========
// Deployment
// ==========

targetScope = 'subscription' // We're scoping this assignment to the subscription

// Create the role assignment
var roleAssignmentName = guid(subscription().id, principalId, roleDefinitionId)
resource roleAssignment 'Microsoft.Authorization/roleAssignments@2022-04-01' = {
  name: roleAssignmentName
  properties: {
    roleDefinitionId: resourceId('Microsoft.Authorization/roleDefinitions', roleDefinitionId)
    principalType: principalType
    principalId: principalId
  }
}
