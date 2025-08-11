# Copilot instructions for adding a new resource to Azure Service Operator

You are an expert Go developer and an AI assistant helping to maintain the Azure Service Operator (ASO) project. Your goal is to add a new Azure resource to the operator based on the details in the assigned GitHub issue.

Instructions in this file are specific to adding a new resource (or a new version of an existing resource) to the project.

## Your Workflow

Follow these steps precisely. If you encounter an error you cannot resolve, stop and report your progress.

### 1. Identify the Resource to Add

From the GitHub issue, determine the resource's **Group**, **Version**, and **Kind** (GVK).

*   **Group**: The Azure service name (e.g., `storage`, `network`, `compute`).
*   **Version**: The API version, like `2021-06-01`, which becomes `v1api20210601`. Use the latest stable version unless the issue specifies a preview.
*   **Kind**: The singular name of the resource (e.g., `storageAccount`, `virtualNetwork`).

### 2. Configure the Code Generator

You will now edit `v2/azure-arm.yaml` to tell the code generator about the new resource.

1.  Locate the `objectModelConfiguration` section.
2.  Find the resource's `group` (e.g., `synapse:`). If it doesn't exist, add it in alphabetical order.
3.  Find the `version` (e.g., `2021-06-01:`). If it doesn't exist, add it in numerical order.
4.  Add the resource `Kind` (e.g., `Workspace:`), ensuring it is nested correctly.
5.  Add the `$exportAs` and `$supportedFrom` directives. The value for `$supportedFrom` should be the next unreleased version of ASO (check project milestones if unsure).

**Example:**
```yaml
# ...
  synapse:
    2021-06-01:
      Workspace:
        $exportAs: Workspace
        $supportedFrom: v2.10.0 # Check for the correct upcoming release
# ...
```

### 3. Run the Code Generator

Execute the following command from the root of the repository to generate the resource files.

```bash
task generator:quick-checks
```

This command will likely fail with errors about resource references. This is expected. Proceed to the next step.

### 4. Fix Generator Errors

The most common error is `"looks like a resource reference but was not labelled as one"`. You must inspect each property reported in the error and decide if it's a reference to another Azure resource.

*   If it **IS** a reference to another ARM resource, mark it with `$referenceType: arm`.
*   If it **IS NOT** a reference, mark it with `$referenceType: simple`.

Update `v2/azure-arm.yaml` with the necessary configuration.

**Example:**
```yaml
# ...
  network:
    2020-11-01:
      NetworkSecurityGroup:
        PrivateLinkResource:
          Id:
            $referenceType: arm # This property IS an ARM reference
# ...
```

After editing the file, re-run `task generator:quick-checks`. Repeat this process until the generator succeeds without errors.

### 5. Review the Generated Code

Inspect the newly generated files in `v2/api/<group>/<version>/`. Pay close attention to the `_types_gen.go` file for your resource.

*   **Check for missing secrets:** Identify any properties on the `Spec` that should be secrets (e.g., passwords, keys) but are currently strings. Mark them in `azure-arm.yaml` with `$isSecret: true`.
*   **Check for Azure-generated secrets:** Look at the `Status` object. If Azure generates keys, connection strings, or other sensitive data upon creation, configure them to be exported to a Kubernetes secret using `$azureGeneratedSecrets`.
*   **Check for read-only properties in the Spec:** Sometimes properties that are read-only in Azure are not marked as such in the OpenAPI spec. These must be removed from the `Spec` by adding a `remove: true` directive in `azure-arm.yaml`.

After making any changes, re-run `task generator:quick-checks`.

### 6. Write an Integration Test

1.  Navigate to `v2/internal/controllers/`.
2.  Find a test for a similar resource and copy it to a new file. The file should be named `<group>_<kind>_v<version>_crud_test.go`.
3.  Modify the test to create, verify, and delete your new resource. Ensure you are testing a representative set of properties.
4.  Delete any existing recording file for your new test from the `recordings` sub-directory.
5.  Run the test to record the interaction with Azure:
    ```bash
    TEST_FILTER=Test_<Group>_<Kind>_CRUD task controller:test-integration-envtest
    ```
    Replace `<Group>` and `<Kind>` with the appropriate values.
    If you're missing the AZURE_TENANT_ID and AZURE_SUBSCRIPTION_ID environment variables, you can't make the recordings yourself; instead, pause and ask for someone to create these recordings for you.


Tips

* Prefer creating all resources at the same time using `CreateResourcesAndWait()` instead of creating them in sequence. This is to flush out any dependency issues. 

### 7. Create a Sample

1.  Create a new sample YAML file in `v2/samples/<group>/<version>/`.
2.  The sample should demonstrate a simple, working configuration of the resource and all relevant dependencies. 
    * Be sure to include any other resources that are required.
    * Resources from other groups should be nested in a `refs` folder (see existing tests for how this pattern works).
3.  Run the samples test to record the new sample's creation and deletion:
    ```bash
    TEST_FILTER=Test_Samples_CreationAndDeletion task controller:test-integration-envtest
    ```
    Again, if you're missing the AZURE_TENANT_ID and AZURE_SUBSCRIPTION_ID environment variables, you can't make the recordings yourself; instead, pause and ask for someone to create these recordings for you.

### 8. Final Verification

Run local checks to ensure your changes haven't introduced any regressions. The full verification process takes too long to run it in a single go; instead, run each of these in turn.  If any step fails, fix the issues and **start again from the beginning of the list.**

1.  **Format Code:** 
    ```bash
    task format-code
    ```

2.  **Run Generator Checks:** This builds the code generator, runs its unit tests, generates any new code, and checks for issues. 
    ```bash
    task generator:quick-checks
    ```
    
3.  **Run Controller Checks:** This builds the main controller, and runs its unit tests.
    ```bash
    task controller:quick-checks
    ```

4. **Run asoctl Checks:** This builds `asoctl` and runs its unit tests
    ```bash
    task asoctl:quick-checks
    ```

## Reference Documentation

For more detailed instructions on adding a new resource, refer to the [Adding a new code generated resource to ASO v2 documentation](docs/hugo/content/contributing/add-a-new-code-generated-resource/_index.md). 

