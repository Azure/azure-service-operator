---
title: Adding a new code generated resource to ASO v2
linktitle: Add a resource
---

This document discusses how to add a new resource to the ASO v2 code generation configuration. Check out 
[this PR](https://github.com/Azure/azure-service-operator/pull/2860) if you'd like to see what the end product looks like.

## What resources can be code generated?
Any ARM resource can be generated. If you're not sure if the resource you are interested in is an ARM resource, check if it
is defined in a `resource-manager` folder in the [Azure REST API specs](https://github.com/Azure/azure-rest-api-specs/tree/main/specification) repo.
If it is, it's an ARM resource.

## Determine a resource to add
There are three key pieces of information required before adding a resource to the code generation configuration file, 
and each of them can be found in the [Azure REST API specs](https://github.com/Azure/azure-rest-api-specs/tree/main/specification)
repo. We will walk through an example of adding Azure Synapse Workspace.

**Note**: In many cases there will be multiple API versions for a given resource, you can see this in the 
[Synapse](https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable)
folder for example, there are 4 API versions as of April 2023. It is _strongly_ recommended that you use the latest 
available non-preview `api-version` when choosing the version of the resource to add.

The three key pieces of information needed to code generate a resource are:

1. The `name` of the resource.

   You usually know this going in. In our example above, the name of the resource is `workspaces`. If you're not sure,
   look in the Swagger/OpenAPI specification file for the service and find the documented PUT for the resource you're 
   interested in. The resource name will be the 
   [second to last section of the URL](https://github.com/Azure/azure-rest-api-specs/blob/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/workspace.json#L71).
2. The `group` the resource is in. 

   This is named after the Azure service, for example `resources` or `documentdb`. 
   In our example entry from above, this is `synapse` (from `Microsoft.Synapse`, the provider documented in the resource URL).
3. The `api-version` of the resource.

   This is usually a date, sometimes with a `-preview` suffix. In our example entry from above, this is `2021-06-01`.

## Adding the resource to the code generation configuration file
The code generation configuration file is located [here](https://github.com/Azure/azure-service-operator/blob/main/v2/azure-arm.yaml). 
To add a new resource to this file, find the `objectModelConfiguration` section of the file.

Find the configuration for the `group` you want; if it's not there, create a new one, inserting it into the existing list 
in alphabetical order. Within the group, find the `version` you want; again, create a new one if it's not already there.

Add your new resource to the list for that version, including the directive `$export: true` nested beneath.
You must also include the `$supportedFrom:` annotation. This should be the _next release_ which will contain support
for the resource in question. You can determine the name of the next ASO release by looking at our 
[milestones](https://github.com/Azure/azure-service-operator/milestones).

The final result should look like this:

``` yaml
<group>:
  <version>:
    <resource name>: # singular, typically just remove the trailing "s"
      $export: true
      $supportedFrom: <the upcoming release>
```

For example, taking the *Azure Synapse Workspace* sample from above:

``` yaml
synapse:
 2021-06-01:
   Workspace:
     $export: true
     $supportedFrom: v2.0.0
```

If ASO was already configured to generate resources from this group (or version), you will need to add your new 
configuration around the existing values.

## Run the code generator

Follow the steps in the [contributing guide]( {{< relref "." >}} ) to set up your development environment.
Once you have a working development environment, run the `task` command to run the code generator.

## Fix any errors raised by the code generator

### \<Resource\> looks like a resource reference but was not labelled as one
Example:
>  Replace cross-resource references with genruntime.ResourceReference: 
> ["github.com/Azure/azure-service-operator/hack/generated/_apis/containerservice/v1alpha1api20210501/PrivateLinkResource.Id" looks like a resource reference but was not labelled as one. 

To fix this error, determine whether the property in question is an ARM ID or not, and then update the `objectModelConfiguration` section in the configuration file.

Find the section you added earlier, adding your property with an `$armReference:` declaration nested below.

If the property is an ARM ID, use `$armReference: true` to flag that property as a reference:

```yaml
network:
  2020-11-01:
    NetworkSecurityGroup:
      $export: true
      PrivateLinkResource: 
        $armReference: true # the property IS an ARM reference
```

If the property is ***not*** an ARM ID, use `$armReference: false` instead:

```yaml
network:
  2020-11-01:
    NetworkSecurityGroup:
      $export: true
      PrivateLinkResource: 
        $armReference: false # the property IS NOT an ARM reference
```

TODO: expand on other common errors

## Examine the generated resource
After running the generator, the new resource you added should be in the 
[apis](https://github.com/Azure/azure-service-operator/blob/main/v2/api/) directory. 

Have a look through the files in the directory named after the `group` and `version` of the resource that was added.
In our `Workspaces` example, the best place to start is `v2/api/synapse/v1api20210601/workspace_types_gen.go`
There may be other resources that already exist in that same directory - that's expected if ASO already supported some 
resources from that provider and API version.

Starting with the `workspace_types_gen.go` file, find the struct representing the resource you just added. 
It should be near the top and look something like this:
```go
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/workspace.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}
type Workspace struct {
metav1.TypeMeta   `json:",inline"`
metav1.ObjectMeta `json:"metadata,omitempty"`
Spec              Workspace_Spec   `json:"spec,omitempty"`
Status            Workspace_STATUS `json:"status,omitempty"`
}

```

Look over the `Spec` and `Status` types and their properties (and the properties of their properties and so-on).
The Azure REST API specs which these types were derived from are not perfect. Sometimes they mark a `readonly` property 
as mutable or have another error or mistake. 

### Common issues
This is a non-exhaustive list of common issues which may need to be fixed in our configuration file.

#### Properties that should have been marked read-only but weren't

There is a section in the config detailing a number of these errors which have been fixed. Look for "Deal with properties that should have been marked readOnly but weren't".

Here's an example, where DocumentDB missed setting the `provisioningState` property to read-only. We override it in the config:

```yaml
  - group: documentdb
    name: Location  # This type is subsequently flattened into NamespacesTopics_Spec
    property: ProvisioningState
    remove: true
    because: This property should have been marked readonly but wasn't.
```

#### Types that can't be found

If you get an error indicating the generator can't find a type, but you're sure it exists:

```
E1214 10:34:15.476761   95884 gen_kustomize.go:111] 
Error during code generation:
failed during pipeline stage 23/67 [filterTypes]: 
Apply export filters to reduce the number of generated types: 
group cdn: version 2021-06-01: 
type DodgyResource not seen (did you mean ResourceReference?): 
type DodgyResource: $exportAs: ReputableResource not consumed
```

It's possible the submodule `v2/specs/azure-rest-api-specs` is out of date. Try running `git submodule update --init --recursive` to update the submodule.

### Debugging

Sometimes it is useful to see what each stage of the generator pipeline has changed. To write detailed debug logs detailing internal stage after each stage of the pipeline has run, use the `--debug` flag to specify the group (or groups) to include.

``` bash
PS> aso-gen gen-types azure-arm.yaml --debug network
I0622 12:28:01.913420    5572 gen_types.go:49] Debug output will be written to the folder C:\...\Temp\aso-gen-debug-1580100077
... elided ...
I0622 12:29:15.836643    5572 gen_types.go:53] Debug output is available in folder C:\...\Temp\aso-gen-debug-1580100077
```

The volume of output is high, so the logs are written into a new temp directory. As shown, the name of this directory is included in the output of the generator twice, once at the start and again at the end.

A separate output file is generated after each pipeline stage runs, allowing for easy comparison to identify what each stage achieves.

![Debug Output](../images/debug-output.png)

In this screenshot, I'm comparing the output after stage 44 with the output after stage 52, reviewing the results of all the intervening stages.

![Stage diffs](../images/stage-diff.png)

Normal use of the `--debug` flag is to specify a single output group (e.g. `--debug network`) but you can also specify multiple groups using semicolons (e.g. `--debug network;compute`) or wildcards (e.g. `--debug db*`).

## Determine if the resource has secrets generated by Azure

Some resources, such as `microsoft.storage/accounts` and `microsoft.documentdb/databaseAccounts`, have keys and endpoints generated by
Azure. Unfortunately, there is no good way to automatically detect these in the specs which the code generator runs on.

**You must manually identify if the resource you are adding has keys (or endpoints) which users will want exported to a secret store**.
If the resource in question _does_ have Azure generated secrets or endpoints, identify those endpoints in the configuration file
by specifying `$azureGeneratedSecrets`.

Our example resource above does not have any Azure generated secrets. As mentioned above, `microsoft.documentdb/databaseAccounts` has 
Azure generated secrets. Here is the snippet from the configuration file showing how they were configured.
```yaml
  documentdb:
    2021-05-15:
      DatabaseAccount:
        $export: true
        $azureGeneratedSecrets:
          - PrimaryMasterKey
          - SecondaryMasterKey
          - PrimaryReadonlyMasterKey
          - SecondaryReadonlyMasterKey
          - DocumentEndpoint
```

Since these properties are manually configured, they must also be retrieved from Azure manually. This can be done using
the extension framework supported by the operator. Use the [documentdb secrets extension](https://github.com/Azure/azure-service-operator/blob/main/v2/api/documentdb/customizations/database_account_extensions.go)
as a template for authoring an extension for your resource.

<!-- TODO: Do we have a link to give more details about extensions and how they work? -->

## Determine if the resource has secrets for input

Some resources take secrets as input from the user, for example when creating a virtual machine the user may supply a `pasword`.
This is a `string` field but because it represents a secret value we don't want to just stick it on the `Spec` as a string.
It needs to be a `genruntime.SecretReference` (a reference to a secret) instead, so that users can safely supply their
secret to the operator without it being accessible via plain-text.

In most cases, the Swagger/OpenAPI document will annotate these properties with an extension `x-ms-secret`, which indicates
that value is a secret. The ASO code-generator uses that annotation to automatically transform the annotated field to
a `genruntime.SecretReference`.

In some cases, the resources Swagger specification is missing this annotation. If that happens, you can use an override in
our `azure-arm.yaml` to force a particular property to be treated as a secret. For example, under `synapse`:

```yaml
WorkspaceProperties:
  SqlAdministratorLoginPassword:
    $isSecret: true
```

## Write a CRUD test for the resource
The best way to do this is to start from an [existing test](https://github.com/Azure/azure-service-operator/blob/main/v2/internal/controllers/crd_cosmosdb_mongodb_test.go) and modify it to work for your resource. It can also be helpful to refer to examples in the [ARM templates GitHub repo](https://github.com/Azure/azure-quickstart-templates).

## Run the CRUD test for the resource and commit the recording
See [the code generator README](../#running-integration-tests) for how to run recording tests.

## Add a new sample
The samples are located in the [samples directory](https://github.com/Azure/azure-service-operator/blob/main/v2/samples). There should be at least one sample for each kind of supported resource. These currently need to be added manually. It's possible in the future we will automatically generate samples similar to how we automatically generate CRDs and types, but that doesn't happen today.

## Run test for added sample and commit the recording
The added new sample needs to be tested and recorded. 

If a recording for the test already exists, delete it.  
Look in the [recordings directory](https://github.com/Azure/azure-service-operator/blob/main/v2/internal/controllers/recordings/Test_Samples_CreationAndDeletion) for a file with the same name as your new test.  
Typically these are named `Test_<GROUP>_<VERSION_PREFIX>_CreationAndDeletion.yaml`.
For example, if we're adding sample for NetworkSecurityGroup resource, check for `Test_Network_v1beta_CreationAndDeletion.yaml`

Run the test and record it:

``` bash
$ TEST_FILTER=Test_Samples_CreationAndDeletion task controller:test-integration-envtest
```

Some Azure resources take longer to provision or delete than the default test timeout of 15m, so you may need to add the `TIMEOUT` environment variable to the command above. For example, to give your test a 60m timeout, use:

``` bash
$ TIMEOUT=60m TEST_FILTER=Test_Samples_CreationAndDeletion task controller:test-integration-envtest
```


## Send a PR

You're all done!
