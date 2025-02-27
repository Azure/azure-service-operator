/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package identity

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/msi-dataplane/pkg/dataplane"
)

var _ TokenCredentialProvider = &mockTokenCredentialProvider{}

type mockTokenCredentialProvider struct {
	TenantID          string
	ClientID          string
	ClientSecret      string
	ClientCertificate []byte
	Password          []byte
	TokenFilePath     string
	AdditionalTenants []string
	Cloud             cloud.Configuration
}

func (m *mockTokenCredentialProvider) NewClientSecretCredential(
	tenantID string,
	clientID string,
	clientSecret string,
	options *azidentity.ClientSecretCredentialOptions,
) (*azidentity.ClientSecretCredential, error) {
	m.TenantID = tenantID
	m.ClientID = clientID
	m.ClientSecret = clientSecret
	if options != nil {
		m.AdditionalTenants = options.AdditionallyAllowedTenants
		m.Cloud = options.Cloud
	}
	// We're not doing anything with the returned secrets so just return a dummy
	return &azidentity.ClientSecretCredential{}, nil
}

func (m *mockTokenCredentialProvider) NewClientCertificateCredential(
	tenantID string,
	clientID string,
	clientCertificate []byte,
	password []byte,
	options *azidentity.ClientCertificateCredentialOptions,
) (*azidentity.ClientCertificateCredential, error) {
	m.TenantID = tenantID
	m.ClientID = clientID
	m.ClientCertificate = clientCertificate
	m.Password = password

	if options != nil {
		m.AdditionalTenants = options.AdditionallyAllowedTenants
		m.Cloud = options.Cloud
	}

	return &azidentity.ClientCertificateCredential{}, nil
}

func (m *mockTokenCredentialProvider) NewManagedIdentityCredential(options *azidentity.ManagedIdentityCredentialOptions) (*azidentity.ManagedIdentityCredential, error) {
	return &azidentity.ManagedIdentityCredential{}, nil
}

func (m *mockTokenCredentialProvider) NewWorkloadIdentityCredential(options *azidentity.WorkloadIdentityCredentialOptions) (*azidentity.WorkloadIdentityCredential, error) {
	m.TenantID = options.TenantID
	m.ClientID = options.ClientID
	m.TokenFilePath = options.TokenFilePath
	m.AdditionalTenants = options.AdditionallyAllowedTenants
	m.Cloud = options.Cloud

	return &azidentity.WorkloadIdentityCredential{}, nil
}

func (m *mockTokenCredentialProvider) NewUserAssignedIdentityCredentials(ctx context.Context, credentialPath string, opts ...dataplane.Option) (azcore.TokenCredential, error) {
	return nil, nil
}
