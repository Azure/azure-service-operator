/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package identity

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/benbjohnson/clock"
	authenticationv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
)

const (
	testServiceAccountNamespace = "tenant-a"
	testServiceAccountName      = "aso-workload"
	testAudience                = "api://AzureADTokenExchange"
	testRequestedLifetime       = time.Hour
)

type tokenRequestHandler func(
	ctx context.Context,
	serviceAccount *corev1.ServiceAccount,
	tokenRequest *authenticationv1.TokenRequest,
) error

type waiterTrackingContext struct {
	context.Context
	entered chan struct{}
	once    sync.Once
}

func (c *waiterTrackingContext) Done() <-chan struct{} {
	c.once.Do(func() {
		close(c.entered)
	})

	return c.Context.Done()
}

func TestServiceAccountTokenProvider_CreatesExpectedTokenRequestAndCachesAssertion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	testClock := clock.NewMock()
	testClock.Set(time.Date(2026, time.September, 25, 12, 0, 0, 0, time.UTC))

	requestCount := 0
	provider := newTestServiceAccountTokenProvider(t, testClock, func(
		ctx context.Context,
		serviceAccount *corev1.ServiceAccount,
		tokenRequest *authenticationv1.TokenRequest,
	) error {
		requestCount++
		g.Expect(ctx).NotTo(BeNil())
		g.Expect(serviceAccount.Namespace).To(Equal(testServiceAccountNamespace))
		g.Expect(serviceAccount.Name).To(Equal(testServiceAccountName))
		g.Expect(tokenRequest.Spec.Audiences).To(Equal([]string{testAudience}))
		g.Expect(tokenRequest.Spec.ExpirationSeconds).NotTo(BeNil())
		g.Expect(*tokenRequest.Spec.ExpirationSeconds).To(Equal(int64(testRequestedLifetime / time.Second)))
		setTokenRequestStatus(tokenRequest, "assertion-1", testClock.Now().Add(testRequestedLifetime))
		return nil
	})

	assertion, err := provider.GetAssertion(
		context.Background(),
		testServiceAccountNamespace,
		testServiceAccountName,
		testAudience,
		testRequestedLifetime,
	)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(assertion).To(Equal("assertion-1"))

	assertion, err = provider.GetAssertion(
		context.Background(),
		testServiceAccountNamespace,
		testServiceAccountName,
		testAudience,
		testRequestedLifetime,
	)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(assertion).To(Equal("assertion-1"))
	g.Expect(requestCount).To(Equal(1))
}

func TestServiceAccountTokenProvider_RefreshesJustInTimeUsingServerExpiration(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	testClock := clock.NewMock()
	testClock.Set(time.Date(2026, time.September, 25, 12, 0, 0, 0, time.UTC))

	requestCount := 0
	provider := newTestServiceAccountTokenProvider(t, testClock, func(
		_ context.Context,
		_ *corev1.ServiceAccount,
		tokenRequest *authenticationv1.TokenRequest,
	) error {
		requestCount++
		setTokenRequestStatus(
			tokenRequest,
			fmt.Sprintf("assertion-%d", requestCount),
			testClock.Now().Add(10*time.Minute),
		)
		return nil
	})

	assertion, err := provider.GetAssertion(
		context.Background(),
		testServiceAccountNamespace,
		testServiceAccountName,
		testAudience,
		testRequestedLifetime,
	)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(assertion).To(Equal("assertion-1"))

	testClock.Add(4*time.Minute + 59*time.Second)
	assertion, err = provider.GetAssertion(
		context.Background(),
		testServiceAccountNamespace,
		testServiceAccountName,
		testAudience,
		testRequestedLifetime,
	)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(assertion).To(Equal("assertion-1"))
	g.Expect(requestCount).To(Equal(1))

	testClock.Add(time.Second)
	g.Expect(requestCount).To(Equal(1), "an idle provider must not refresh in the background")

	assertion, err = provider.GetAssertion(
		context.Background(),
		testServiceAccountNamespace,
		testServiceAccountName,
		testAudience,
		testRequestedLifetime,
	)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(assertion).To(Equal("assertion-2"))
	g.Expect(requestCount).To(Equal(2))
}

func TestServiceAccountTokenProvider_CoalescesConcurrentRequestsForSameKey(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	testClock := clock.NewMock()
	requestStarted := make(chan struct{})
	releaseRequest := make(chan struct{})
	requestCount := 0

	provider := newTestServiceAccountTokenProvider(t, testClock, func(
		_ context.Context,
		_ *corev1.ServiceAccount,
		tokenRequest *authenticationv1.TokenRequest,
	) error {
		requestCount++
		close(requestStarted)
		<-releaseRequest
		setTokenRequestStatus(tokenRequest, "shared-assertion", testClock.Now().Add(time.Hour))
		return nil
	})

	const callerCount = 20
	results := make(chan string, callerCount)
	errs := make(chan error, callerCount)
	go func() {
		assertion, err := provider.GetAssertion(
			context.Background(),
			testServiceAccountNamespace,
			testServiceAccountName,
			testAudience,
			testRequestedLifetime,
		)
		results <- assertion
		errs <- err
	}()
	<-requestStarted

	waitersEntered := make([]chan struct{}, 0, callerCount-1)
	for range callerCount - 1 {
		entered := make(chan struct{})
		waitersEntered = append(waitersEntered, entered)
		go func() {
			ctx := &waiterTrackingContext{Context: context.Background(), entered: entered}
			assertion, err := provider.GetAssertion(
				ctx,
				testServiceAccountNamespace,
				testServiceAccountName,
				testAudience,
				testRequestedLifetime,
			)
			results <- assertion
			errs <- err
		}()
	}

	for _, entered := range waitersEntered {
		<-entered
	}
	close(releaseRequest)

	for range callerCount {
		g.Expect(<-errs).NotTo(HaveOccurred())
		g.Expect(<-results).To(Equal("shared-assertion"))
	}
	g.Expect(requestCount).To(Equal(1))
}

func TestServiceAccountTokenProvider_RefreshesDifferentKeysIndependently(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	testClock := clock.NewMock()
	requestsStarted := make(chan string, 2)
	releaseRequests := make(chan struct{})

	provider := newTestServiceAccountTokenProvider(t, testClock, func(
		_ context.Context,
		serviceAccount *corev1.ServiceAccount,
		tokenRequest *authenticationv1.TokenRequest,
	) error {
		requestsStarted <- serviceAccount.Namespace
		<-releaseRequests
		setTokenRequestStatus(tokenRequest, serviceAccount.Namespace, testClock.Now().Add(time.Hour))
		return nil
	})

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	for _, namespace := range []string{"tenant-a", "tenant-b"} {
		go func() {
			defer waitGroup.Done()
			_, _ = provider.GetAssertion(
				context.Background(),
				namespace,
				testServiceAccountName,
				testAudience,
				testRequestedLifetime,
			)
		}()
	}

	started := map[string]bool{
		<-requestsStarted: true,
		<-requestsStarted: true,
	}
	g.Expect(started).To(Equal(map[string]bool{"tenant-a": true, "tenant-b": true}))
	close(releaseRequests)
	waitGroup.Wait()
}

func TestServiceAccountTokenProvider_DoesNotCacheRefreshErrors(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	testClock := clock.NewMock()
	requestStarted := make(chan struct{})
	releaseRequest := make(chan struct{})

	var mu sync.Mutex
	requestCount := 0
	provider := newTestServiceAccountTokenProvider(t, testClock, func(
		_ context.Context,
		_ *corev1.ServiceAccount,
		tokenRequest *authenticationv1.TokenRequest,
	) error {
		mu.Lock()
		requestCount++
		currentRequest := requestCount
		mu.Unlock()
		if currentRequest == 1 {
			close(requestStarted)
			<-releaseRequest
			return apierrors.NewServiceUnavailable("temporarily unavailable")
		}

		setTokenRequestStatus(tokenRequest, "recovered-assertion", testClock.Now().Add(time.Hour))
		return nil
	})

	const callerCount = 20
	errs := make(chan error, callerCount)
	go func() {
		_, err := provider.GetAssertion(
			context.Background(),
			testServiceAccountNamespace,
			testServiceAccountName,
			testAudience,
			testRequestedLifetime,
		)
		errs <- err
	}()
	<-requestStarted

	waitersEntered := make([]chan struct{}, 0, callerCount-1)
	for range callerCount - 1 {
		entered := make(chan struct{})
		waitersEntered = append(waitersEntered, entered)
		go func() {
			ctx := &waiterTrackingContext{Context: context.Background(), entered: entered}
			_, err := provider.GetAssertion(
				ctx,
				testServiceAccountNamespace,
				testServiceAccountName,
				testAudience,
				testRequestedLifetime,
			)
			errs <- err
		}()
	}

	for _, entered := range waitersEntered {
		<-entered
	}
	close(releaseRequest)
	for range callerCount {
		g.Expect(<-errs).To(MatchError(ContainSubstring("temporarily unavailable")))
	}

	assertion, err := provider.GetAssertion(
		context.Background(),
		testServiceAccountNamespace,
		testServiceAccountName,
		testAudience,
		testRequestedLifetime,
	)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(assertion).To(Equal("recovered-assertion"))
	g.Expect(requestCount).To(Equal(2))
}

func TestServiceAccountTokenProvider_WrapsKubernetesErrorsAndPreservesClassification(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		apiError error
		classify func(error) bool
	}{
		{
			name: "missing ServiceAccount",
			apiError: apierrors.NewNotFound(
				schema.GroupResource{Group: "", Resource: "serviceaccounts"},
				testServiceAccountName,
			),
			classify: apierrors.IsNotFound,
		},
		{
			name: "token creation forbidden",
			apiError: apierrors.NewForbidden(
				schema.GroupResource{Group: "", Resource: "serviceaccounts/token"},
				testServiceAccountName,
				errors.New("RBAC denied"),
			),
			classify: apierrors.IsForbidden,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			provider := newTestServiceAccountTokenProvider(t, clock.NewMock(), func(
				_ context.Context,
				_ *corev1.ServiceAccount,
				_ *authenticationv1.TokenRequest,
			) error {
				return test.apiError
			})

			_, err := provider.GetAssertion(
				context.Background(),
				testServiceAccountNamespace,
				testServiceAccountName,
				testAudience,
				testRequestedLifetime,
			)
			g.Expect(err).To(MatchError(ContainSubstring(
				"requesting a token for ServiceAccount tenant-a/aso-workload",
			)))
			g.Expect(test.classify(err)).To(BeTrue())
		})
	}
}

func TestServiceAccountTokenProvider_PropagatesCancellation(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	provider := newTestServiceAccountTokenProvider(t, clock.NewMock(), func(
		ctx context.Context,
		_ *corev1.ServiceAccount,
		_ *authenticationv1.TokenRequest,
	) error {
		<-ctx.Done()
		return ctx.Err()
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := provider.GetAssertion(
		ctx,
		testServiceAccountNamespace,
		testServiceAccountName,
		testAudience,
		testRequestedLifetime,
	)
	g.Expect(errors.Is(err, context.Canceled)).To(BeTrue())
}

func newTestServiceAccountTokenProvider(
	t *testing.T,
	testClock clock.Clock,
	handler tokenRequestHandler,
) ServiceAccountTokenProvider {
	t.Helper()
	g := NewGomegaWithT(t)
	scheme := runtime.NewScheme()
	g.Expect(corev1.AddToScheme(scheme)).To(Succeed())
	g.Expect(authenticationv1.AddToScheme(scheme)).To(Succeed())

	fakeClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithInterceptorFuncs(interceptor.Funcs{
			SubResourceCreate: func(
				ctx context.Context,
				_ client.Client,
				subResourceName string,
				obj client.Object,
				subResource client.Object,
				_ ...client.SubResourceCreateOption,
			) error {
				g.Expect(subResourceName).To(Equal("token"))
				serviceAccount, ok := obj.(*corev1.ServiceAccount)
				g.Expect(ok).To(BeTrue())
				tokenRequest, ok := subResource.(*authenticationv1.TokenRequest)
				g.Expect(ok).To(BeTrue())
				return handler(ctx, serviceAccount, tokenRequest)
			},
		}).
		Build()

	return NewServiceAccountTokenProvider(kubeclient.NewClient(fakeClient), testClock)
}

func setTokenRequestStatus(
	tokenRequest *authenticationv1.TokenRequest,
	assertion string,
	expiresAt time.Time,
) {
	tokenRequest.Status = authenticationv1.TokenRequestStatus{
		Token:               assertion,
		ExpirationTimestamp: metav1.NewTime(expiresAt),
	}
}
