# Azure Service Operator FAQ

This page provides help with the most common questions about Azure Service Operators.

### Do I have to order the creation of resources through the operator? For instance, do I need to first create the EventHub namespace before creating the EventHub?
No, you can issue the creation of all resources at the same time. The operator will take care of waiting and requeuing the requests until the parent resource is ready and all resources will eventually be ready.

### Is there a guide that talks about how to develop an operator for a new Azure service?
There is a [step-by-step guide](/docs/v1/howto/newoperatorguide.md) that walks you through this process.

### Are there any sample apps that demonstrate how to utilize the Azure Service Operators?
We have some samples that illustrate how to use Azure Service Operators along with your application [here](https://github.com/Azure-Samples/azure-service-operator-samples). The samples also show how to utilize post-deployment-secrets from the Azure Service Operator.
