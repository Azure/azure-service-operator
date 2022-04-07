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

>Only the first 3 environment variables need to be changed before pasting.

`VOTES_APP_INSTANCE_NAME` should be an alphanumeric string that is likely to be globally unique. The tutorial will use `azvotessqldemo1` but if you are following along you should pick something else.

Find the `TENANT_ID` on the Azure Portal
 1. go to https://portal.azure.com/
 2. select "Azure Active Directory"
 3. from the left, select "Properties"
 4. find the section labeled "Tenant ID" and copy the ID

Find the `SUBSCRIPTION_ID` on the Azure Portal
 1. go to https://portal.azure.com/
 2. select "Subscriptions"
 3. copy the ID from one of the listed subscriptions

```
VOTES_APP_INSTANCE_NAME="<alias>"
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
OPERATOR_SP_NAME="http://"$VOTES_APP_INSTANCE_NAME"-operator-creds"
HELM_RELEASE_NAME=$VOTES_APP_INSTANCE_NAME

mkdir -p $CREDENTIAL_DIR
az login -t $TENANT_ID \
    && az account set -s $SUBSCRIPTION_ID
```

### Create a new resource group where your ACR and AKS cluster will live

>If you already have a Kubernetes cluster and Container registry you can skip this step.

```
az group create -g $RESOURCE_GROUP -l $LOCATION
```

### Create the Azure Container Registry where the Votes app image will be pushed
>If you already have a Container registry you can skip this step.

>This guide attaches ACR directly to the AKS cluster. If you use your own Container Registry you may need to create an image pull secret.

```
az acr create \
    --name $ACR_NAME \
    --resource-group $RESOURCE_GROUP \
    --location $LOCATION \
    --sku Basic
```

### Create a Kubernetes cluster
>If you already have a Kubernetes cluster you can skip this step. Any Kubernetes cluster v1.13+ should work here.

#### Create an AKS Cluster attached to the ACR

>You may need to add the `--generate-ssh-keys` flag if you get an error stating the following: `An RSA key file or key value must be supplied to SSH Key Value`
```
az aks create \
    --location $LOCATION \
    --name $AKS_NAME \
    --resource-group $RESOURCE_GROUP \
    --node-count 2 \
    --node-vm-size "Standard_DS2_v2" \
    --attach-acr $ACR_NAME # remove this flag if using your own container registry
```

#### Configure local kubectl to point to new cluster
```
az aks get-credentials \
    --name $AKS_NAME \
    --resource-group $RESOURCE_GROUP
```

### Create Service Principal for Azure Service Operator

>The access granted for the purposes of this guide may not be appropriate for production deployments. We'll store the credentials in a JSON file to access later when deploying ASO.

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
kubectl apply-f https://github.com/jetstack/cert-manager/releases/download/v0.12.0/cert-manager.yaml
```

Wait for Cert Manager to deploy before continuing...or run this handy command

```
kubectl rollout status -n cert-manager deploy/cert-manager-webhook
```

### Install Azure Service Operator
```
helm upgrade --install aso https://github.com/Azure/azure-service-operator/raw/main/charts/azure-service-operator-0.1.0.tgz \
    --namespace=azureoperator-system \
    --create-namespace \
    --set azureSubscriptionID=$SUBSCRIPTION_ID \
    --set azureTenantID=$(cat $CREDENTIAL_DIR/operator_sp_creds.json | jq -r .tenantId) \
    --set azureClientID=$(cat $CREDENTIAL_DIR/operator_sp_creds.json | jq -r .clientId) \
    --set azureClientSecret=$(cat $CREDENTIAL_DIR/operator_sp_creds.json | jq -r .clientSecret) \
    --set image.repository="mcr.microsoft.com/k8s/azureserviceoperator:0.1.11139"
```

Successful output should look like this:

```
Release "aso" does not exist. Installing it now.
NAME: aso
LAST DEPLOYED: Mon Jun 29 15:35:56 2020
NAMESPACE: azureoperator-system
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

To check on the operator run `kubectl get po`

```
kubectl get pods -n azureoperator-system
NAME                                                READY   STATUS    RESTARTS   AGE
azureoperator-controller-manager-7dd75bbd97-djx8g   2/2     Running   1          2m50s
```

## Build the Votes App and deploy

### Clone the Votes App and Helm Chart

[Azure Service Operator Samples](git@github.com:Azure-Samples/azure-service-operator-samples.git)

Clone this repository. Move to the `azure-votes-sql` directory.

```
git clone git@github.com:Azure-Samples/azure-service-operator-samples.git
cd azure-service-operator-samples/azure-votes-sql/
```

Get ACR host info for `docker build` command
```
az acr show -g $RESOURCE_GROUP -n $ACR_NAME --query loginServer -o tsv
```

e.g.

```
# In this example VOTES_APP_INSTANCE_NAME="azvotessqldemo1"
>echo $(az acr show -g $RESOURCE_GROUP -n $ACR_NAME --query loginServer -o tsv)
acrazvotessqldemo1.azurecr.io
```

Use the value from the `echo` command to create this variable

```
DOCKER_HOST=$(az acr show -g $RESOURCE_GROUP -n $ACR_NAME --query loginServer -o tsv)
```

Log in to ACR

```
az acr login --name $ACR_NAME
```

#### Build/Push Docker image
```
docker build -t $DOCKER_HOST/sqldemo:1 .
docker push $DOCKER_HOST/sqldemo:1
```

#### Helm install Azure SQL Votes App

```
helm upgrade --install $HELM_RELEASE_NAME charts/azure-sql-demo \
    --set resourceGroupName=$APP_RESOURCE_GROUP \
    --set serverName=$AZURE_SQL_SEVER_NAME \
    --set databaseName=$AZURE_SQL_DB_NAME \
    --set region=$LOCATION \
    --set image=$DOCKER_HOST/sqldemo:1 \
    --set local=true # this opens up firewall access to the server
```

Successful output for this command looks like:
```
Release "azvotessqldemo1" does not exist. Installing it now.
NAME: azvotessqldemo1
LAST DEPLOYED: Mon Jun 29 15:51:28 2020
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
Demo App Deployed!

Run `kubectl port-forward deployment/azure-votes-sql 8080:8080`
Then visit `http://localhost:8080` in your browser to test.
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
kubectl get ResourceGroup
NAME                     PROVISIONED   MESSAGE
app-rg-azvotessqldemo1   true          successfully provisioned

kubectl get AzureSqlServer
NAME                  PROVISIONED   MESSAGE
sql-azurevotesdemo   true          successfully provisioned

kubectl get AzureSqlDatabase
NAME                    PROVISIONED   MESSAGE
sqldb-azvotessqldemo1   true          successfully provisioned

kubectl get AzureSqlFirewallRule
NAME                    PROVISIONED   MESSAGE
saso-sql-fwrule         true          successfully provisioned
saso-sql-fwrule-local   true          successfully provisioned
```

Once all the Azure resources are successfully provisioned (ResourceGroup, AzureSqlServer, AzureSqlDatabase, AzureSqlFirewallRule) the Azure Votes App will be able to start.

>[!NOTE]
>The Azure resources typically take 3-4 minutes to become ready. The application will attempt to start once the resources are provisioned but this doesn't mean all networking is in fact ready. For this reason, the application is built to handle connection failures by exiting, so Kubernetes can restart it.

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
sql-azvotessqldemo1             Opaque                                5      5m4s
```
