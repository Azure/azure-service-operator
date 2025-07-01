# Roadmap

We track upcoming work on Azure Service Operator (ASO) using [GitHub Issues](https://github.com/Azure/azure-service-operator/issues?q=is%3Aissue+is%3Aopen).

## Release Cadence & Planning

We target a new release of ASO approximately every two months, though this may vary depending on the amount of work we have planned as well as external factors.

Our current release plan:

| Version                                                                |    Estimated Release |
| ---------------------------------------------------------------------- | -------------------: |
| [2.15.0](https://github.com/Azure/azure-service-operator/milestone/34) |      Mid August 2025 |
| [2.16.0](https://github.com/Azure/azure-service-operator/milestone/35) |     Mid October 2025 |
| [2.17.0](https://github.com/Azure/azure-service-operator/milestone/36) |  Early December 2025 |

Where linked, versions go to a list of feature and bugs that are planned to be included in that release.

Any items note completed in time for one release will be carried over to the next, and may result in us pushing other items to a later release.

Partway through each release cycle, we'll review the list of issues assigned to upcoming releases and redistribute issues as needed. This usually involves some issues being moved to later releases.

If you're waiting on a particular resource or feature to be released, please comment on the relevant issue (or create a new issue if there isn't already one tracking the request) to let us know. We'll do our best to keep you updated on progress.

## Experimental Releases

We publish an [experimental release](https://github.com/Azure/azure-service-operator/releases/tag/experimental) on a regular basis, suitable for testing and other non-production use cases. Such releases may not be stable and should not be used in production.

## Prior Releases

Prior GA releases of ASO v2:

| Version                                                                        |     Release Date |                                                                                        |
| ------------------------------------------------------------------------------ | ---------------: | -------------------------------------------------------------------------------------- |
| [2.14.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.14.0) |     24 June 2025 |                                                                                        |
| [2.13.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.13.0) |    23 April 2025 |                                                                                        |
| [2.12.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.12.0) | 11 February 2025 |                                                                                        |
| [2.11.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.11.0) | 12 November 2024 |                                                                                        |
| [2.10.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.10.0) |  22 October 2024 |                                                                                        |
| [2.9.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.9.0)   |   22 August 2024 |                                                                                        |
| [2.8.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.8.0)   |     25 June 2024 |                                                                                        |
| [2.7.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.7.0)   |    25 April 2024 |                                                                                        |
| [2.6.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.6.0)   | 23 February 2024 |                                                                                        |
| [2.5.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.5.0)   |  7 December 2023 | Release cycle was abbreviated to get key items out before the 2023/2024 holiday season |
| [2.4.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.4.0)   | 14 November 2023 |                                                                                        |
| [2.3.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.3.0)   | 5 September 2023 |                                                                                        |
| [2.2.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.2.0)   |     21 July 2023 |                                                                                        |
| [2.1.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.1.0)   |      2 June 2023 |                                                                                        |
| [2.0.0](https://github.com/Azure/azure-service-operator/releases/tag/v2.0.0)   |    15 April 2023 |                                                                                        |

## Issue Triage

We triage new issues weekly and assign most of them to an upcoming release milestone, based on our understanding of priority, complexity, and available resourcing.

## What about ASO v1?

Azure Service Operator v1 is end of life, with the code preserved in the [`asov1` branch](https://github.com/Azure/azure-service-operator/blob/asov1). It is receiving neither security updates nor bug fixes. Existing users are advised to [migrate to ASO v2](https://azure.github.io/azure-service-operator/guide/asov1-asov2-migration/).
