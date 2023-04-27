---
title: Reducing Access
linktitle: Reducing Access
---

## Using a credential for ASO with reduced permissions

Most examples of installing ASO suggest using an identity that has Contributor access to the Subscription.
Such broadly scoped access is _**not**_ required to run the operator, it's just the easiest way to set things up and so
is often used in examples. 

Here we discuss a few approaches to restricting the access that the ASO identity has.

### Contributor access at a reduced scope

In some scenarios you may know that ASO will be operating on only certain resource groups. For example, if you have two
teams working in a cluster with ASO, "Marketing" and "Product". If you know ahead of time that the only resource groups these
teams will use are `marketing-dev`, `marketing-prod`, `product-dev` and `product-prod`, you can pre-create the following resources:

* The 4 resource groups.
* An identity for use by ASO, assigned Contributor permission to each of the 4 resource groups.

Now install ASO following the instructions above. Once ASO is installed, adopt the existing resource groups in ASO by applying a YAML in the cluster
for each of the 4 resource groups. Now users of ASO are free to create resources in those resource groups as normal. If they attempt
to use ASO to create a new resource group `foo` they will be rejected by Azure as the ASO identity doesn't have permission to create
arbitrary resource groups in the Subscription.

### Reduced access at Subscription scope

In other scenarios, you may want to give out reduced access at the Subscription scope. This can be done by determining what resources 
in ASO you will be using and creating a custom role scoped to just those permissions. 

For example if you want to use ASO to manage ResourceGroups, Redis and MySQL flexible servers, but _not_ anything else, you could 
[create a role](https://learn.microsoft.com/azure/role-based-access-control/custom-roles-cli) similar to `aso-operator.json`:

```json
{
  "Name": "ASO Operator",
  "IsCustom": true,
  "Description": "Role with access to perform only the operations which we allow ASO to perform",
  "Actions": [
    "Microsoft.Resources/subscriptions/resourceGroups/*",
    "Microsoft.Cache/*",
    "Microsoft.DBforMySQL/*"
  ],
  "NotActions": [
  ],
  "AssignableScopes": [
    "/subscriptions/{subscriptionId1}",
    "/subscriptions/{subscriptionId2}"
  ]
}
```
Then use the az cli to create that custom role definition: `az role definition create --role-definition ~/aso-operator.json`

See [list of resource provider operations](https://learn.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for a comprehensive
list of operations. 

**Note: We strongly recommend giving full permissions to the resource types ASO will be managing.** ASO needs `read`, `write`, and `delete` permissions to a resource
to fully manage it. In some cases, it also needs `action` permissions. It's recommended to give `*` permissions for a given resource type which ASO 
will be managing.

Once the role has been created, assign it to the ASO identity and install the operator as normal. It will be able to
manage any resource of the types you allowed. Any other resources will fail to be created/updated/deleted with a permissions error.
