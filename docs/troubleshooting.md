# Azure Service Operator (ASO) troubleshooting

This document is written from a beginner perspective to help those new to project understand common issues they might face. We hope to iterate and grow this document, so feel free to contribute!

# What are the key learning so far

## What are Operators and OLM?

To learn more about operators, Red Hat OpenShift has a good article on what they are conceptually and how valuable they can be. You can also find details on the Operator Lifecycle Manager (OLM) [here](https://docs.openshift.com/container-platform/4.6/operators/understanding/olm-what-operators-are.html)

## Is Az-CLI version is important?

Yes, please make sure it is 2.0.53 or later. To check your current version, run `az --version`. If you need to update to a later version, check out [Update Azure CLI](https://docs.microsoft.com/en-us/cli/azure/update-azure-cli).

## Does the steps work out-of-order?

Nope, for your first run through, make sure you follow the steps in order. This means setting up and installing the Operator Lifecycle Manager (OLM) and proceeding with steps 1 to 7. This is also true for the steps documented in OperatorHub catalog. The steps there assume that you have successfully installed OLM.

## So what if I messed up while installing Azure Service Operator?

If you've accidentally run through the steps out of order, or are having some trouble with the install process, check out the tips below. Thankfully, the Kubernetes eco-system provides enough flexibility that you should only need a bit of tinkering around to get the ASO running. Let's go through a scenario:

Let's assume that as an end user Tats provided the wrong service principal when trying to set up ASO.

**Sample scenario**: `user-A` has mistakenly used an outdated or incorrect service principal secret for `azure-service-operator` which resulted in an unsuccessful install. What can we do now?

We can recognize that we reached this issue at step 4 of our install guide[here](https://operatorhub.io/operator/azure-service-operator). To remediate this, all the user needs to do is renew or refresh their service-principal credentials.
    
* To renew the service principal, see the command mentioned in step 4 [here](https://operatorhub.io/operator/azure-service-operator). 
* Next, after regenerating/renewing the credentials, the user needs to re-run step 5 onwards as mentioned [here] (https://operatorhub.io/operator/azure-service-operator). However, only a few of these steps are impacted by the service principal change. For example, the step to retrieve and set the `AZURE_TENANT_ID` can be skipped.
    
Next, find the cloud env var for your cluster. You can find instructions [here](https://docs.microsoft.com/en-us/azure/storage/common/storage-powershell-independent-clouds#get-endpoint-using-get-azenvironment). 

Once you have that, create the `azureoperatorsettings.yaml` file from step 6 mentioned [here](https://operatorhub.io/operator/azure-service-operator), and replace the placeholder values into the respective fields. Finally, run `kubectl apply -f <location of your azureoperatorsettings.yaml file>`.

You should be able to see `azureoperator-controller-manager-***` pods running in your cluster. To view the pods, run `kubectl get pods -n operators` 

## I have already installed `azureserviceoperator` but I'm not sure which service principal secret is in use?
This command might be able to help you see your secret locally. 
```kubectl get secret azureoperatorsettings -n operators -o go-template='{{range $k,$v := .data}}{{printf \"%s: \" $k}}{{if not $v}}{{$v}}{{else}}{{$v | base64decode}}{{end}}{{\"\n\"}}{{end}}'```
[SO thread for reference](https://stackoverflow.com/questions/56909180/decoding-kubernetes-secret/58117444#58117444)

## What if I see the operator namespace already exists for the cluster?

If you see that the operator namespace already exists for the cluster, it's likely the result of trying to install the operator previously. It also means that you can now successfully move on to the next step. 
However, this can also be a symptom of incorrectly installing the operator, for example, trying to run the install guide out of order. Other common issues related to an "out-of-order install":

* `namespace already exist` - This means that azure-service-operator (ASO) has already been run for the cluster before.
* `service-principal` issues: If you're seeing silent failures with service principal in use, meaning it is expired or needs to be reset, you can follow the steps as mentioned [here](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli#create-a-service-principal) or in step 4 here in the [OperatorHub guidance](https://operatorhub.io/operator/azure-service-operator).
* `kubectl api-resource` failure with `error: unable to retrieve the complete list of server APIs: metcs.k8s.io/vabeta=1: the server is currently unable to handle the request.` : This error indicates a networking error within your cluster and is not related to your Azure Service Operator installation or configuration.

 ## When trying to install cert-manager, I'm seeing that the connection is refused for `cert-manager-package`?
This could be related to following open error: https://github.com/jetstack/cert-manager/issues/2752
 
 ## The azureoperator-controller-manager seems to be in a crash loop. How do I investigate?

`kubectl describe  pod operators`

To get logs in case ASO operator is crashlooping : `kubectl logs -n  operators pod/azureoperator-controller-manager-****** --container manager`

## What if I see cert-manager-webhook timeout errors?

This is network related, and you should check to see if there are any ongoing network issues. This is evidenced by common error messages like: `Post https://cert-manager-webhook.cert-manager.svc:443/mutate?timeout=30s: connection refused` 
## What if I am seeing APIserver errors like `metrics.k8s.io/v1beta1: the server is currently unable to handle request`?

This could be a result of a Kubernetes internal error like control plane & data plane not being able to communicate. Follow normal Kubernetes troubleshooting steps.
## How do I create a service principal?

This link might come handy: https://docs.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal 