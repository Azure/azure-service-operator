# API Management v1api20240501 Samples

This directory contains sample YAML configurations for Azure API Management using the `v1api20240501` API version.

## Available Samples

### Core Resources
- `v1api20240501_service.yaml` - API Management service instance
- `v1api20240501_subscription.yaml` - API subscription for accessing APIs
- `v1api20240501_backend.yaml` - Backend service configuration
- `v1api20240501_namedvalue.yaml` - Named value for configuration parameters

### API and Product Management
- `v1api20240501_apiversionset.yaml` - API version set for managing multiple API versions
- `v1api20240501_api.yaml` - API definition and configuration
- `v1api20240501_product.yaml` - Product definition for grouping APIs
- `v1api20240501_productapi.yaml` - Link between product and API

### Policies
- `v1api20240501_policy.yaml` - Global policy configuration
- `v1api20240501_policyfragment.yaml` - Reusable policy fragment
- `v1api20240501_productpolicy.yaml` - Product-specific policy

### Authorization
- `v1api20240501_authorizationprovider.yaml` - OAuth2 authorization provider
- `v1api20240501_authorizationprovidersauthorization.yaml` - Authorization configuration
- `v1api20240501_authorizationprovidersauthorizationsaccesspolicy.yaml` - Access policy for authorization

### Reference Resources
The `refs/` directory contains supporting resources:
- `v1api20230131_userassignedidentity.yaml` - Managed identity for authorization
- `v1api20240501_secrets.yaml` - Secret containing OAuth2 client credentials

## Dependencies

Most samples depend on the existence of a resource group named `aso-sample-rg`. The authorization samples also require the reference resources in the `refs/` directory.

## API Version Features

The `v1api20240501` API version is the latest stable release for Azure API Management, providing:
- All features from previous versions
- Latest stable API capabilities
- Production-ready functionality
- Enhanced security and performance features

This version replaces `v1api20230501preview` as the recommended version for new deployments.