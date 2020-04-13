# App Insights Operator

This operator deploys an Application Insights instance into a specified resource group at the specified location.

Learn more about Application Insights [here](https://docs.microsoft.com/en-us/azure/azure-monitor/app/app-insights-overview).

Here is a [sample YAML](/config/samples/azure_v1alpha1_appinsights.yaml) to provision an Application Insights instance.

### Required Fields

An Application Insights instance needs the following fields to deploy, along with a location and resource group.

* `Kind` specify the kind of application this component refers to, possible values are: 'web', 'ios', 'other', 'store', 'java', 'phone'
* `ApplicationType` specify the types of application being monitored, possible values are: 'web', 'other'

### Secrets

After creating an Application Insights instance, the operator stores a JSON formatted secret with the following fields. This secret is stored as a Kubernetes Secret or in an Azure Keyvault, with the same name as your Application Insights instance.
For more details on where the secrets are stored, look [here](/docs/secrets.md)

* `instrumentationKey`

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
