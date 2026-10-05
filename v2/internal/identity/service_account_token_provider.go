/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package identity

import (
	"context"
	"sync"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/rotisserie/eris"
	authenticationv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
)

const serviceAccountTokenRefreshMargin = 5 * time.Minute

type ServiceAccountTokenProvider interface {
	GetAssertion(
		ctx context.Context,
		namespace string,
		serviceAccountName string,
		audience string,
		requestedLifetime time.Duration,
	) (string, error)
}

type serviceAccountTokenProvider struct {
	kubeClient kubeclient.Client
	clock      clock.Clock

	mu       sync.Mutex
	cache    map[serviceAccountTokenCacheKey]serviceAccountToken
	inFlight map[serviceAccountTokenCacheKey]*serviceAccountTokenRefresh
}

type serviceAccountTokenCacheKey struct {
	namespace string
	name      string
	audience  string
}

type serviceAccountToken struct {
	assertion string
	expiresAt time.Time
}

type serviceAccountTokenRefresh struct {
	done      chan struct{}
	assertion string
	err       error
}

var _ ServiceAccountTokenProvider = &serviceAccountTokenProvider{}

func NewServiceAccountTokenProvider(kubeClient kubeclient.Client, clk clock.Clock) ServiceAccountTokenProvider {
	return &serviceAccountTokenProvider{
		kubeClient: kubeClient,
		clock:      clk,
		cache:      make(map[serviceAccountTokenCacheKey]serviceAccountToken),
		inFlight:   make(map[serviceAccountTokenCacheKey]*serviceAccountTokenRefresh),
	}
}

func (p *serviceAccountTokenProvider) GetAssertion(
	ctx context.Context,
	namespace string,
	serviceAccountName string,
	audience string,
	duration time.Duration,
) (string, error) {
	key := serviceAccountTokenCacheKey{
		namespace: namespace,
		name:      serviceAccountName,
		audience:  audience,
	}

	p.mu.Lock()
	if cached, ok := p.cachedAssertion(key); ok {
		p.mu.Unlock()
		return cached, nil
	}

	if refresh, ok := p.inFlight[key]; ok {
		p.mu.Unlock()
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-refresh.done:
			return refresh.assertion, refresh.err
		}
	}

	refresh := &serviceAccountTokenRefresh{done: make(chan struct{})}
	p.inFlight[key] = refresh
	p.mu.Unlock()

	token, err := p.requestAssertion(ctx, key, duration)

	p.mu.Lock()
	if err == nil {
		p.cache[key] = token
	}
	refresh.assertion = token.assertion
	refresh.err = err
	delete(p.inFlight, key)
	close(refresh.done)
	p.mu.Unlock()

	return token.assertion, err
}

func (p *serviceAccountTokenProvider) cachedAssertion(key serviceAccountTokenCacheKey) (string, bool) {
	cached, ok := p.cache[key]
	if !ok || !cached.expiresAt.After(p.clock.Now().Add(serviceAccountTokenRefreshMargin)) {
		return "", false
	}

	return cached.assertion, true
}

func (p *serviceAccountTokenProvider) requestAssertion(
	ctx context.Context,
	key serviceAccountTokenCacheKey,
	requestedLifetime time.Duration,
) (serviceAccountToken, error) {
	expirationSeconds := int64(requestedLifetime / time.Second)
	request := &authenticationv1.TokenRequest{
		Spec: authenticationv1.TokenRequestSpec{
			Audiences:         []string{key.audience},
			ExpirationSeconds: &expirationSeconds,
		},
	}
	serviceAccount := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: key.namespace,
			Name:      key.name,
		},
	}

	err := p.kubeClient.SubResource("token").Create(ctx, serviceAccount, request)
	if err != nil {
		return serviceAccountToken{}, eris.Wrapf(
			err,
			"requesting a token for ServiceAccount %s/%s",
			key.namespace,
			key.name,
		)
	}

	if request.Status.Token == "" {
		return serviceAccountToken{}, eris.Errorf(
			"token request for ServiceAccount %s/%s returned an empty token",
			key.namespace,
			key.name,
		)
	}

	if request.Status.ExpirationTimestamp.IsZero() {
		return serviceAccountToken{}, eris.Errorf(
			"token request for ServiceAccount %s/%s returned an empty expiration timestamp",
			key.namespace,
			key.name,
		)
	}

	return serviceAccountToken{
		assertion: request.Status.Token,
		expiresAt: request.Status.ExpirationTimestamp.Time,
	}, nil
}
