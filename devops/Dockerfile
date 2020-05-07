FROM mcr.microsoft.com/azure-cli

LABEL maintainer="Microsoft Corporation"

RUN az login --service-principal -u "${AZURE_CLIENT_ID}" -p="${AZURE_CLIENT_SECRET}" --tenant "${AZURE_TENANT_ID}" && \
az account set -s "${AZURE_SUBSCRIPTION_ID}" && \
az aks get-credentials -n "${CLUSTERNAME}" -g "${RESOURCEGROUP}" --admin

WORKDIR /

COPY . ./

ENV REQUEUE_AFTER "30"

ENTRYPOINT ["/manager"]