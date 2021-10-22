---
name: Bug report
about: Create a report to help us improve
title: 'Bug: <Brief description of bug>'
labels: bug
assignees: ''

---

**Version of Azure Service Operator**
<!--- 
The version of the operator pod. 
Assuming your ASO is deployed in the default namespace, you can get this version from the container image which the controller is running.
Use one of the following commands:
ASO V1: `kubectl get deployment -n azureoperator-system azureoperator-controller-manager -o wide` and share the image being used by the manager container.
ASO V2: `kubectl get deployment -n azureserviceoperator-system azureserviceoperator-controller-manager -o wide` and share the image being used by the manager container.   
-->

**Describe the bug**
A clear and concise description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
<Fill in the steps>

**Expected behavior**
A clear and concise description of what you expected to happen.

**Screenshots**
If applicable, add screenshots to help explain your problem.

**Additional context**
Add any other context about the problem here.
