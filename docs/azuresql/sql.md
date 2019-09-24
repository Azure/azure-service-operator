# Azure SQL Operator

## Resources supported

1. Azure SQL server
2. SQL database
3. SQL firewall rule
4. Action (Rolling user credentials for the SQL server)

## Deploying SQL Resources

### Create a Resource Group

Create a resource group by running the following yaml file. Update the name variable to your preferred resource group name.

```kubectl create -f config/samples/azure_v1_resourcegroup.yaml```

### Create SQL resources

Create your SQL resources by running the following yaml files. Update the resource group name, and SQL resource names to your preferred names.

```kubectl create -f config/samples/azure_v1_sqldatabase.yaml```

```kubectl create -f config/samples/azure_v1_sqlfirewallrule.yaml```

```kubectl create -f config/samples/azure_v1_sqlserver.yaml```

### View SQL Resources

To view your created SQL resources, such as sqlserver, run the following command:

`k get sqlserver`

Your servers should be displayed with their name and age.

### Delete a SQL Resource

To delete an existing resource from Kubernetes and Azure, such as SQL server, run:

`k edit sqlserver $sqlservername`

Remove the lines under the finalizer, and then delete the sql server instance:

`k delete sqlserver $sqlservername`

The following message should appear:

`sqlserver.azure.microsoft.com "$sqlservername" deleted.`

## Troubleshooting resource provisioning

kubectl describe to see events

## How would you use the Azure SQL Operator from a real application

describe demo app
