/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package identity

import "github.com/Azure/azure-sdk-for-go/sdk/azidentity"

var _ TokenCredentialProvider = &mockTokenCredentialProvider{}

type mockTokenCredentialProvider struct {
	TenantID          string
	ClientID          string
	ClientSecret      string
	ClientCertificate []byte
	Password          []byte
	TokenFilePath     string
}

func (m *mockTokenCredentialProvider) NewClientSecretCredential(tenantID string, clientID string, clientSecret string, options *azidentity.ClientSecretCredentialOptions) (*azidentity.ClientSecretCredential, error) {
	m.TenantID = tenantID
	m.ClientID = clientID
	m.ClientSecret = clientSecret
	// We're not doing anything with the returned secrets so just return a dummy
	return &azidentity.ClientSecretCredential{}, nil
}

func (m *mockTokenCredentialProvider) NewClientCertificateCredential(tenantID, clientID string, clientCertificate, password []byte) (*azidentity.ClientCertificateCredential, error) {
	m.TenantID = tenantID
	m.ClientID = clientID
	m.ClientCertificate = clientCertificate
	m.Password = password

	return &azidentity.ClientCertificateCredential{}, nil
}

func (m *mockTokenCredentialProvider) NewManagedIdentityCredential(options *azidentity.ManagedIdentityCredentialOptions) (*azidentity.ManagedIdentityCredential, error) {
	return &azidentity.ManagedIdentityCredential{}, nil
}

func (m *mockTokenCredentialProvider) NewWorkloadIdentityCredential(options *azidentity.WorkloadIdentityCredentialOptions) (*azidentity.WorkloadIdentityCredential, error) {
	m.TenantID = options.TenantID
	m.ClientID = options.ClientID
	m.TokenFilePath = options.TokenFilePath

	return &azidentity.WorkloadIdentityCredential{}, nil
}
