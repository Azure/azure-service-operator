FROM mcr.microsoft.com/azure-cli

LABEL maintainer="Microsoft"

ENV KUBECONFIG "./config"

RUN az login --service-principal -u AZURE_CLIENT_ID -p=AZURE_CLIENT_SECRET --tenant AZURE_TENANT_ID
# RUN az keyvault secret download --file ./config 

WORKDIR /

# COPY manager .

# ENTRYPOINT ["/manager"]