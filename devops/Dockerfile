FROM mcr.microsoft.com/azure-cli

LABEL maintainer="Microsoft"

ENV KUBECONFIG "./config"

RUN az login --service-principal -u AZURE_CLIENT_ID -p=AZURE_CLIENT_SECRET --tenant AZURE_TENANT_ID
RUN az account set -s AZURE_SUBSCRIPTION_ID
RUN az keyvault secret download --vault-name skullykv1 --id https://skullykv1.vault.azure.net/secrets/default-sqlserverskulz0 -f ./secret
RUN ls -als

WORKDIR /

# COPY ./secret .
# COPY manager .

ENTRYPOINT ["/manager"]