FROM mcr.microsoft.com/azure-cli

LABEL maintainer="Microsoft"

# ENV AZURE_CLIENT_ID "${AZURE_CLIENT_ID}"
# ENV AZURE_CLIENT_SECRET "{AZURE_CLIENT_SECRET}"
# ENV AZURE_SUBSCRIPTION_ID "${AZURE_SUBSCRIPTION_ID}"
# ENV AZURE_TENANT_ID "${AZURE_TENANT_ID}"
# ENV REQUEUE_AFTER "30"
# ENV AZURE_OPERATOR_KEYVAULT "${AZURE_OPERATOR_KEYVAULT}"
# ENV KUBECONFIG "./config"

RUN echo $TEST
# RUN az login --service-principal -u '${AZURE_CLIENT_ID}' -p '{AZURE_CLIENT_SECRET}' --tenant '{AZURE_TENANT_ID}'
# RUN az keyvault secret download --file ./config 

WORKDIR /

# COPY manager .

# ENTRYPOINT ["/manager"]