# Quick start application deployment using ASO

In this guide we will:
 - Create a Service Principal
 - Create an Azure Container Registry
 - Create an AKS cluster with the ACR attached
 - Deploy ASO
 - Deploy an Azure Votes app that requires Azure SQL Server using ASO

## Set up the Kubernetes environment and deploy ASO

 ### Requirements
- az cli ([install instructinos](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest))
- kubectl cli ([install instructions](https://kubernetes.io/docs/tasks/tools/install-kubectl/))
- helm cli ([install instructions](https://helm.sh/docs/intro/install/))
- jq [install instructions](https://stedolan.github.io/jq/download/)

### Prepare credentials and variables

Set these variables now as they will help with copy/paste of commands later.

Note: only the first 3 need to be changed.

```
VOTES_APP_INSTANCE_NAME="<alias>" # only alpha numeric
TENANT_ID=""
SUBSCRIPTION_ID=""
LOCATION="westus2"
CREDENTIAL_DIR="$HOME/Desktop/votes-creds"

RESOURCE_GROUP="rg-"$VOTES_APP_INSTANCE_NAME
APP_RESOURCE_GROUP="app-rg-"$VOTES_APP_INSTANCE_NAME
ACR_NAME="acr"$VOTES_APP_INSTANCE_NAME
AKS_NAME="aks-"$VOTES_APP_INSTANCE_NAME

AZURE_SQL_SEVER_NAME="sql-"$VOTES_APP_INSTANCE_NAME
AZURE_SQL_DB_NAME="sqldb-"$VOTES_APP_INSTANCE_NAME
ACR_PUSH_SP_NAME="http://"$VOTES_APP_INSTANCE_NAME"-acr-push-creds"
OPERATOR_SP_NAME="http://"$VOTES_APP_INSTANCE_NAME"-operator-creds"

HELM_RELEASE_NAME=$VOTES_APP_INSTANCE_NAME

mkdir -p $CREDENTIAL_DIR
az login -t $TENANT_ID \
    && az account set -s $SUBSCRIPTION_ID
```

### Create a new resource group where your ACR and AKS cluster will live

Note: if you already have a Kubernetes cluster and Container registry you can skip some steps.

```
az group create -g $RESOURCE_GROUP -l $LOCATION
```

### Create the Azure Container Registry where the Votes app image will be pushed
```
az acr create \
    --name $ACR_NAME \
    --resource-group $RESOURCE_GROUP \
    --location $LOCATION \
    --sku Basic
```

### Create a Kubernetes cluster
- Create an AKS Cluster attached to the ACR with monitoring addon enabled

> Note: you may need to add the `--generate-ssh-keys` flag if you get an error stating the following: `An RSA key file or key value must be supplied to SSH Key Value`
```
az aks create \
    --location $LOCATION \
    --name $AKS_NAME \
    --resource-group $RESOURCE_GROUP \
    --node-count 2 \
    --node-vm-size "Standard_DS2_v2" \
    --attach-acr $ACR_NAME # remove this flag if using your own container registry

az aks get-credentials \
    --name $AKS_NAME \
    --resource-group $RESOURCE_GROUP
```

### Create Service Principal for Azure Service Operator

Note: The access granted for the purposes of this guide may not be appropriate for production deployments

```

az ad sp create-for-rbac \
    --sdk-auth \
    --scope "/subscriptions/"$SUBSCRIPTION_ID \
    --role Contributor \
    --name $OPERATOR_SP_NAME > $CREDENTIAL_DIR/operator_sp_creds.json
```

### Install Cert Manager

Cert Manager is used to provide certificates for Kube RBAC Proxy and the operator's web hooks.

```
kubectl create namespace cert-manager
kubectl label namespace cert-manager cert-manager.io/disable-validation=true
kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v0.12.0/cert-manager.yaml
```

Wait for Cert Manager to deploy before continuing...or run this handy command

```
kubectl rollout status -n cert-manager deploy/cert-manager-webhook
```

### Install Azure Service Operator
```
helm upgrade --install aso https://github.com/Azure/azure-service-operator/raw/master/charts/azure-service-operator-0.1.0.tgz \
    --namespace=azureoperator-system \
    --create-namespace \
    --set azureSubscriptionID=$SUBSCRIPTION_ID \
    --set azureTenantID=$(cat $CREDENTIAL_DIR/operator_sp_creds.json | jq -r .tenantId) \
    --set azureClientID=$(cat $CREDENTIAL_DIR/operator_sp_creds.json | jq -r .clientId) \
    --set azureClientSecret=$(cat $CREDENTIAL_DIR/operator_sp_creds.json | jq -r .clientSecret) \
    --set image.repository="mcr.microsoft.com/k8s/azureserviceoperator:0.1.11139"
```

## Build the Votes App and deploy

### Clone the Votes App and Helm Chart

Go to [Azure Service Operator Samples](git@github.com:Azure-Samples/azure-service-operator-samples.git)

Clone this repository. Move to the `azure-votes-sql` directory.

Get ACR host info for `docker build` commadn
```
echo $(az acr show -g $RESOURCE_GROUP -n $ACR_NAME --query loginServer -o tsv)
```

Log in to ACR

```
az acr login --name $ACR_NAME
```

#### Build/Push Docker image
```
docker build -t your_registry/sqldemo:1 .
docker push your_registry/sqldemo:1
```

#### Helm install Azure SQL Votes App

```
helm upgrade --install $HELM_RELEASE_NAME charts/azure-sql-demo \
    --set resourceGroupName=$APP_RESOURCE_GROUP \
    --set serverName=$AZURE_SQL_SEVER_NAME \
    --set databaseName=$AZURE_SQL_DB_NAME \
    --set region=$LOCATION \
    --set image=your_registry.com/sqldemo:1 \
    --set local=true # this opens up firewall access to the server
```

This command will deploy 5 kinds of Kube resoruces

- ResourceGroup
- AzureSqlServer
- AzureSqlDatabase
- AzureSqlFirewallRule
- Deployment

Use `kubectl get ${Kind}` to see the status of that object.

E.g.

```
kubectl get AzureSqlServer
NAME                  PROVISIONED   MESSAGE
sql-azurevotesdemo   true          successfully provisioned
```

Once all the Azure resources are successfully provisioned (ResourceGroup, AzureSqlServer, AzureSqlDatabase, AzureSqlFirewallRule) the Azure Votes App will be able to start.

Note: The Azure resources take 3-4 minutes to become ready. The application will attempt to start once the resources are provisioned but this doesn't mean all networking is in fact ready. For the reason the application is biult to handle failures in connection by exiting so Kubernetes can restart it.

Your deployment may restart 5 times before everything is actually available.

```
kubectl get pods -w
NAME                               READY   STATUS             RESTARTS   AGE
azure-votes-sql-7c5f5dbc86-r4lbv   0/1     CrashLoopBackOff   4          2m38s
azure-votes-sql-7c5f5dbc86-r4lbv   1/1     Running            5          2m58s
```

Now that the app is running, port forward to it using this command:

```
kubectl port-forward deployment/azure-votes-sql 8080:8080
```

Then, in your browser, go to "http://localhost:8080"

![image of app](https://github.com/Azure-Samples/azure-service-operator-samples/blob/master/azure-votes-sql/app.png?raw=true)

If you need to access the database via the portal or some SQL workbench tool simply access the secret with the server credentials.

```
kubectl get secret
NAME                            TYPE                                  DATA   AGE
your_sqlserver_name             Opaque                                5      33m
```
