/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/go-logr/logr"
)

type deleteAwareRoundTripper struct {
	inner        http.RoundTripper
	deletedPaths map[string]struct{}
	lroPollURLs  map[string]struct{}
	log          logr.Logger
	lock         sync.Mutex
}

var _ http.RoundTripper = &deleteAwareRoundTripper{}

func newDeleteAwareRoundTripper(inner http.RoundTripper, log logr.Logger) *deleteAwareRoundTripper {
	return &deleteAwareRoundTripper{
		inner:        inner,
		deletedPaths: make(map[string]struct{}),
		lroPollURLs:  make(map[string]struct{}),
		log:          log,
	}
}

func (r *deleteAwareRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	if request.Method == http.MethodGet {
		return r.roundTripGet(request)
	}

	response, err := r.inner.RoundTrip(request)
	if err != nil {
		return response, err
	}

	if response.StatusCode >= http.StatusOK && response.StatusCode < http.StatusMultipleChoices {
		requestPath := urlPath(request.URL.String())
		switch request.Method {
		case http.MethodDelete:
			r.recordDeletion(requestPath)
			r.recordLROPollURLs(request, response)
		case http.MethodPut, http.MethodPost, http.MethodPatch:
			r.clearDeletedAncestors(requestPath)
		}
	}

	return response, nil
}

func (r *deleteAwareRoundTripper) roundTripGet(request *http.Request) (*http.Response, error) {
	if r.isLROPollURL(request.URL.String()) {
		return r.inner.RoundTrip(request)
	}

	for {
		response, err := r.inner.RoundTrip(request)
		if err != nil {
			return response, err
		}

		requestPath := urlPath(request.URL.String())
		if !r.wasDeleted(requestPath) ||
			(response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated) {
			return response, nil
		}

		r.log.V(1).Info("Discarding stale GET response after deletion", "url", request.URL.String())
		if response.Body != nil {
			if err := response.Body.Close(); err != nil {
				return nil, fmt.Errorf("closing stale GET response body: %w", err)
			}
		}
	}
}

func (r *deleteAwareRoundTripper) recordLROPollURLs(request *http.Request, response *http.Response) {
	for _, header := range []string{"Operation-Location", "Azure-AsyncOperation", "Location"} {
		value := response.Header.Get(header)
		if value == "" {
			continue
		}

		pollURL, err := request.URL.Parse(value)
		if err != nil {
			continue
		}

		r.lock.Lock()
		r.lroPollURLs[pollURL.String()] = struct{}{}
		r.lock.Unlock()
		return
	}
}

func (r *deleteAwareRoundTripper) isLROPollURL(rawURL string) bool {
	r.lock.Lock()
	defer r.lock.Unlock()
	_, ok := r.lroPollURLs[rawURL]
	return ok
}

func (r *deleteAwareRoundTripper) recordDeletion(path string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.deletedPaths[path] = struct{}{}
}

func (r *deleteAwareRoundTripper) clearDeletedAncestors(path string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	for deletedPath := range r.deletedPaths {
		if r.isSameOrDescendantPath(path, deletedPath) {
			delete(r.deletedPaths, deletedPath)
		}
	}
}

func (r *deleteAwareRoundTripper) wasDeleted(path string) bool {
	r.lock.Lock()
	defer r.lock.Unlock()
	for deletedPath := range r.deletedPaths {
		if r.isSameOrDescendantPath(path, deletedPath) {
			return true
		}
	}

	return false
}

func (*deleteAwareRoundTripper) isSameOrDescendantPath(path string, ancestor string) bool {
	path = strings.TrimSuffix(path, "/")
	ancestor = strings.TrimSuffix(ancestor, "/")
	if strings.EqualFold(path, ancestor) {
		return true
	}
	if ancestor == "" {
		return false
	}
	if ancestor == "/" {
		return strings.HasPrefix(path, "/")
	}
	if len(path) <= len(ancestor) {
		return false
	}

	return strings.EqualFold(path[:len(ancestor)], ancestor) && path[len(ancestor)] == '/'
}
