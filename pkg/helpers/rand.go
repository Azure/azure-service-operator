// copied from github.com/Azure/open-service-broker-azure/
package helpers

import (
	mathrand "math/rand"
	"sync"
	"time"
)

// Seeded contains a self-contained, seeded (math/rand).Rand. It allows you
// to generate random numbers in a concurrency-safe manner.
//
// Create one of these with NewSeeded()
type Seeded struct {
	seededRand *mathrand.Rand
	mut        *sync.Mutex
}

// NewSeeded creates a new Seeded
func NewSeeded() *Seeded {
	rnd := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	return &Seeded{
		seededRand: rnd,
		mut:        &sync.Mutex{},
	}
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n)
// from the internally held seeded random number generator.
//
// It panics if n <= 0.
//
// This function (and its documentation!) is similar to (math/rand).Intn(max)
func (s *Seeded) Intn(max int) int {
	s.mut.Lock()
	defer s.mut.Unlock()
	return s.seededRand.Intn(max)
}
