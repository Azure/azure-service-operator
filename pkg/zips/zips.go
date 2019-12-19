// zips package is where all the fun Azure client, cache, throttling, CRUD will go. Right now, it just provides an
// Apply and Delete interface
package zips

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
)

type (
	Resourcer interface {
		ToResource() (Resource, error)
		FromResource(Resource) (runtime.Object, error)
	}

	Applier interface {
		Apply(ctx context.Context, res Resource) error
		Delete(ctx context.Context, resourceID string) error
	}

	Resource struct {
		ID         string
		Name       string
		Location   string
		Type       string
		APIVersion string
		Properties map[string]interface{}
	}
)
