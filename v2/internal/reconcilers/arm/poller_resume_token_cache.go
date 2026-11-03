/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package arm

import (
	"sync"

	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type cachedPollerResumeToken struct {
	id            string
	token         string
	generation    int64
	hasGeneration bool
}

type pollerResumeTokenCache struct {
	mu     sync.RWMutex
	tokens map[types.UID]cachedPollerResumeToken
}

func newPollerResumeTokenCache() *pollerResumeTokenCache {
	return &pollerResumeTokenCache{
		tokens: make(map[types.UID]cachedPollerResumeToken),
	}
}

func (c *pollerResumeTokenCache) add(obj genruntime.MetaObject) {
	id, token, ok := GetPollerResumeToken(obj)
	if !ok || obj.GetUID() == "" {
		return
	}

	generation, hasGeneration := GetLatestReconciledGeneration(obj)
	cached := cachedPollerResumeToken{
		id:            id,
		token:         token,
		generation:    generation,
		hasGeneration: hasGeneration,
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.tokens[obj.GetUID()] = cached
}

func (c *pollerResumeTokenCache) restore(obj genruntime.MetaObject) bool {
	if _, _, ok := GetPollerResumeToken(obj); ok {
		c.clear(obj)
		return false
	}

	c.mu.RLock()
	cached, ok := c.tokens[obj.GetUID()]
	c.mu.RUnlock()
	if !ok {
		return false
	}

	SetPollerResumeToken(obj, cached.id, cached.token)
	if cached.hasGeneration {
		setLatestReconciledGeneration(obj, cached.generation)
	}

	return true
}

func (c *pollerResumeTokenCache) clear(obj genruntime.MetaObject) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.tokens, obj.GetUID())
}
