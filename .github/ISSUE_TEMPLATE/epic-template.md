---
name: Epic template
about: Template to use for Epics
title: As a user, I can configure and deploy <Azure Service name> by specifying a
  CRD in the Kubernetes YAML file
labels: epic
assignees: ''

---

This epic covers

- Development of an Azure Service operator using Kubebuilder which creates a CRD for <Azure service name>

- An user should be able to deploy/delete/update the resource

- We'll use Azure Go SDK for the resource management/creation

- There will need to be tests and logging added

**Acceptance criteria**
Reference: [Done-Done checklist] [Link](https://github.com/Microsoft/code-with-engineering-playbook/blob/master/Engineering/BestPractices/DoneDone.md)

- [ ] Can you create, delete and update settings for the Azure API Management service?

- [ ] Are there unit tests for the resource helper functions?

- [ ] Are there integration tests that run as part of the release pipeline?

- [ ] Is there logging and telemetry added for diagnosability?

- [ ] Is there documentation added for how to deploy and use the operator?

- [ ] Code reviewed by product group?

- [ ] Code reviewed by customer?

**User Stories**
Add links to User Stories for this epic. User stories are intended to be completed in one sprint.
