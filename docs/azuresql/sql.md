# Azure SQL Operator

## Resources Supported

1. Azure SQL server
2. SQL database
3. SQL firewall rule
4. Action (Rolling user credentials for the SQL server)

## Deploying SQL Resources

We will use two terminal windows to deploy SQL resouces.

- Terminal Window #1 - used to run the operator
- Terminal Window #2 - used to create new Azure resources.

In the first terminal window, run the operator by running the following commands:

```bash
make install
make run
```

Now that the operator is running, open your second terminal window. Here we will use our yaml files to create different Azure resources.

### 1. Create a Resource Group

Create a resource group by running the following yaml file. Update the name variable to your preferred resource group name.

```bash
kubectl create -f config/samples/azure_v1_resourcegroup.yaml
```

### 2. Create SQL resources

Create your SQL resources by running the following yaml files. Update the resource group name, and SQL resource names to your preferred names.

```bash
kubectl create -f config/samples/azure_v1_sqldatabase.yaml
```

```bash
kubectl create -f config/samples/azure_v1_sqlfirewallrule.yaml
```

```bash
kubectl create -f config/samples/azure_v1_sqlserver.yaml
```

## Updating SQL Resources

### View SQL Resources

To view your created SQL resources, such as sqlserver, run the following command:

```bash
kubectl get sqlserver
```

Your servers should be displayed with their name and age.

### Delete a SQL Resource

To delete an existing resource from Kubernetes and Azure, such as SQL server, run:

```bash
kubectl edit sqlserver $sqlservername
```

Remove the lines under the finalizer, and then delete the sql server instance:

```bash
kubectl delete sqlserver $sqlservername
```

The following message should appear:

`sqlserver.azure.microsoft.com "$sqlservername" deleted.`

## Troubleshooting Provisioning Resources

kubectl describe to see events

## Demo App

Watch this demo <https://bit.ly/2lUIX6Y> to observe how you would you use the Azure SQL Operator from a real application.
