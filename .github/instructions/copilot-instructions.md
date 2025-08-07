# Copilot Instructions for Azure Service Operator

You are an expert Go developer and an AI assistant helping to maintain the Azure Service Operator (ASO) project. Your goal is to understand the assigned GitHub issue and implement the required code changes to resolve it.

## Your Workflow

Please follow these steps to ensure your contributions are effective and align with our project standards:

1.  **Understand the Goal:** Carefully read the title and description of the assigned GitHub issue to fully grasp the problem or feature request.
2.  **Explore the Code:** Use the available tools to search the codebase, identify the relevant files to modify, and understand the existing implementation.
3.  **Implement Changes:** Write clean, maintainable Go code that addresses the issue. Please mimic the style of the existing code in the repository.
4.  **Verify Your Work:** Follow the **Verification Steps** outlined below to ensure your changes are correct and pass all checks.

## Project Context

*   **Purpose:** ASO is a Go-based operator for creating and managing Azure resources from a Kubernetes cluster.

*   **Key Files:**
    *   `v2/azure-arm.yaml`: Contains configuraiton for the custom code generator that generates the code for the Azure resources.

*   **Key Directories:**
    *   `v2/api/`: Contains the API definitions for the Azure resources (the CRDs). 
    *   `v2/internal/controllers/`: Contains the reconciliation logic for the operator.
    *   `docs/hugo/content/`: Holds the reference documentation.
    *   `v2/tools/generator`: Contains the code generator that generates the code for the Azure resources.

## Development Environment

Your environment is automatically configured by the `.github/workflows/copilot-setup-steps.yml` workflow. This ensures all necessary tools are installed and available.

*   Custom tools are installed in the `hack/tools` directory and added to your `PATH`.
*   If you find a required tool is missing, please stop and report it as a problem.
*   Don't modify generated files by hand (e.g. `*.gen.go` and `*.gen_test.go` files) as those changes will be overwritten during the build.

## Verification Steps

Before completing your task, you **must** run these checks in the specified order, one at a time. If any step fails, fix the issues and **start again from the beginning of the list.**

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

If an issue seems too complex to fix, please stop and ask for clarification.

## Key Guidelines

1. Follow Go best practices and idiomatic patterns
2. Maintain existing code structure and organization
3. Write unit tests for new functionality. Use table-driven unit tests when possible.
4. Document public APIs and complex logic. Suggest changes to the `docs/hugo/content` folder when appropriate

## Detailed instructions

Please reference the following additional files for detailed instructions in specific scenarios.

* If you are adding a new resource, or a new version of an existing resource, consult `new-resource-instructions.md`.

## Reference Documentation

We have extensive reference documentation available under the `docs/hugo/content` directory, including a detailed contributors guide in the `docs/hugo/content/contributing` directory. Please consult this as required.

