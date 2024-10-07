---
name: Bug report
about: Create a report to help us improve
title: 'Bug: <Brief description of bug>'
labels: "bug \U0001FAB2"
assignees: ''
---

## Describe the bug

_Please give a clear description of what you see and why you think it is a bug._

Azure Service Operator Version: _What version of ASO are you using?_

<!--- 
The version of the operator pod. 
Assuming your ASO is deployed in the default namespace, you can get this version from the container image which the controller is running.
Use one of the following commands:
ASO V1: `kubectl get deployment -n azureoperator-system azureoperator-controller-manager -o wide` and share the image being used by the manager container.
ASO V2: `kubectl get deployment -n azureserviceoperator-system azureserviceoperator-controller-manager -o wide` and share the image being used by the manager container.   
-->

## Expected behavior

_A concise description of what you expected to happen._

## To Reproduce

_Describe the steps to recreate the issue._

## Screenshots

_If applicable, add screenshots to help explain your problem._

## Additional context

_Add any other context about the problem here._
