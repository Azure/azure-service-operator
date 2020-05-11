FROM mcr.microsoft.com/azure-cli

LABEL maintainer="Microsoft Corporation"

WORKDIR /

#COPY . ./

COPY /bin/manager .

ENV REQUEUE_AFTER "30"

ENTRYPOINT ["/manager"]