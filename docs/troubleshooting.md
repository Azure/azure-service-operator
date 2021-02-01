# ASO troubleshooting

This document is written with a point of view of help newbies like me to understand gotchas and as the knowledge will grow across, this docuemnt will evolve and not needed one day. Learning with mistakes is sometimes to good to learn and this document will try to cover points of gotchas.

When you create azure service operator, you might occasionally come across problems. This article details some common problems and troubleshooting steps.

# What are the key learning so far


## Is Az-CLI version is important?
Yes, please make sure it is 2.0.53 or later, user `az --version` to find the right version in user machine. 

## Does the steps work out-of-order?
No, order is key for the first time user, setting\installing the Operator Lifecycle Manager (OLM) and operator service up the operator-namespace and then the steps 1 to 7. The step 1 to 7 mentioned in the operator hub page safely assumes that user have successfully installed the OLM.

## How to move forward for the out-of-order run \ or wrong input failure for the intermediate steps?

ASO tool is growing currently all these intermediate steps to set up ASO in users existing cluster will fail with either "already-existing" error if user have already tried to run the ASO steps half way. Intent of below help is to unblock user in case of out-of-order run of steps or if user faces some issues from the intermediate steps.

It is ok to learn by experiment and backward engineering. Kubernetes eco-system to provide that flexibility but only needs bit of tinkering around to get the ASO running. Let me take you through a scenario:

Lets assume that that as an end user Tats povided wrong service-principle which was incorrect for the ASO to be correctly set-up.

Example of: Service principle error:

* ToDo: I plan to add spoecific user-case scenario here: (possibly with screenshot)

* So at this point the issue happend in step 3, essentially all user need to do is re-new or refresh their service-principle credentials.

* Then after re-generating that user need to make sure that the only steps they need to manually run are step-5 onwards mentioned here: https://operatorhub.io/operator/azure-service-operator , But we can simplify a little more here as most of these command will be availabe here: use command : `az account show` 

* Decoding the env var for the users: https://docs.microsoft.com/en-us/azure/storage/common/storage-powershell-independent-clouds#get-endpoint-using-get-azenvironment 
* 

## What if I see operator-namespace already exit for the cluster?

If user will see already installed operator-namespace then it means user have tried installing the operator-namespace in-past, and it also means that user can now successfully move to the next step. 

But, let me here add a small note that in the edge case of anyone playing around with out-of-order running of command, and wondering why the existing operator running is not 

Minor detail of when steps are done out of sync and user see: 

* `namespace already exist` 

* `service-principle` issues, silent failures with service principle in use is expired or re-set in those case you can follow the as mentioned here https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli#create-a-service-principalor in step-4 here in official service operator page. https://operatorhub.io/operator/azure-service-operator

* `kubectl api-resource` failure with `error: unable to retirve the com0plete list of server APIs: metcs.k8s.io/vabeta=1: the server is currently unable to handle the request.` - This error is a big indicaiton of network error within cluster and this is jsut the causal occurrence and not something ASO needs handles.

 
## How to get operator crash loops for azureoperator-controller-manager?

`kubectl describe  pod operators`

To get Logs in case aso operator is crashlooping : `kubectl logs -n  operators pod/azureoperator-controller-manager-7cd684745f-dvqtq --conatiner manager`


## What if I see cert-manager-webhook timeout errors?

This is network related again, and possible check is there is any on-going network issue. Common error message like: `Post https://cert-manager-webhook.cert-manager.svc:443/mutate?timeout=30s: connection refused` 




Reference:

https://docs.microsoft.com/en-us/azure/aks/troubleshooting 






The Azure Service Operator comprises of:

The Custom Resource Definitions (CRDs) for each of the Azure services a Kubernetes user can provision. The Kubernetes controller that watches for requests to create Custom Resources for each of these CRDs and creates them.
The project was built using Kubebuilder



