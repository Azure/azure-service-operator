FROM mcr.microsoft.com/azure-cli

LABEL maintainer="Microsoft"

RUN az login --service-principal -u AZURE_CLIENT_ID -p=AZURE_CLIENT_SECRET --tenant AZURE_TENANT_ID
RUN az account set -s AZURE_SUBSCRIPTION_ID
RUN az aks get-credentials -n aso-azdo-aks -g resourcegroup-azure-operators --admin

WORKDIR /

COPY . ./

RUN ls -als

ENTRYPOINT ["/manager"]