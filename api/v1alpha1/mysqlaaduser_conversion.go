// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	"encoding/json"
	"sort"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
)

var _ conversion.Convertible = &MySQLAADUser{}

// To avoid losing information converting a v1alpha2 instance into
// v1alpha1, we stash the affected fields (roles and database roles)
// into a json-serialised struct in the annotations. When converting
// back to a v1alpha2 instance we use any stashed values, layering any
// changes to the database roles back over the top. The conversions
// will only return errors if JSON marshalling/unmarshalling fails.

func (src *MySQLAADUser) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.MySQLAADUser)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// If there are stashed values in the annotations then we need to
	// retrieve them first.
	if encoded, found := src.getStashedAnnotation(src.ObjectMeta); found {
		var stashedValues stashedMySQLAADUserFields
		err := json.Unmarshal([]byte(encoded), &stashedValues)
		if err != nil {
			return errors.Wrap(err, "decoding stashed fields")
		}
		dst.Spec.Roles = stashedValues.Roles
		dst.Spec.DatabaseRoles = stashedValues.DatabaseRoles

		// Clear out the annotation to avoid confusion.
		delete(dst.ObjectMeta.Annotations, conversionStashAnnotation)
		if len(dst.ObjectMeta.Annotations) == 0 {
			dst.ObjectMeta.Annotations = nil
		}
	}

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Server = src.Spec.Server
	dst.Spec.AADID = src.Spec.AADID
	dst.Spec.Username = src.Spec.Username

	if dst.Spec.Roles == nil {
		dst.Spec.Roles = []string{}
	}
	if dst.Spec.DatabaseRoles == nil {
		dst.Spec.DatabaseRoles = make(map[string][]string)
	}
	dst.Spec.DatabaseRoles[src.Spec.DBName] = append([]string(nil), src.Spec.Roles...)

	// Status
	dst.Status = v1alpha2.ASOStatus(src.Status)

	return nil
}

func (dst *MySQLAADUser) getStashedAnnotation(meta metav1.ObjectMeta) (string, bool) {
	if meta.Annotations == nil {
		return "", false
	}
	value, found := meta.Annotations[conversionStashAnnotation]
	return value, found
}

func (dst *MySQLAADUser) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.MySQLAADUser)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	if len(src.Spec.DatabaseRoles) != 1 || len(src.Spec.Roles) != 0 {
		// If this can't be represented exactly as a v1alpha1, store
		// the original server-level and database roles in an
		// annotation.
		if dst.ObjectMeta.Annotations == nil {
			dst.ObjectMeta.Annotations = make(map[string]string)
		}
		stashValues := stashedMySQLAADUserFields{
			DatabaseRoles: src.Spec.DatabaseRoles,
			Roles:         src.Spec.Roles,
		}
		encoded, err := json.Marshal(stashValues)
		if err != nil {
			return errors.Wrap(err, "encoding stashed fields")
		}
		dst.ObjectMeta.Annotations[conversionStashAnnotation] = string(encoded)
	}

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Server = src.Spec.Server
	dst.Spec.AADID = src.Spec.AADID
	dst.Spec.Username = src.Spec.Username

	// Pick the first database name to include as the DBName.
	var dbNames []string
	for dbName := range src.Spec.DatabaseRoles {
		dbNames = append(dbNames, dbName)
	}
	sort.Strings(dbNames)
	var (
		dbName string
		roles  []string
	)
	if len(dbNames) != 0 {
		dbName = dbNames[0]
		roles = src.Spec.DatabaseRoles[dbName]
	}

	dst.Spec.DBName = dbName
	dst.Spec.Roles = append(dst.Spec.Roles, roles...)

	// Status
	dst.Status = ASOStatus(src.Status)

	return nil

}

const conversionStashAnnotation = "azure.microsoft.com/convert-stash"

// stashedMySQLAADUserFields stores values that can't be represented
// directly on a v1alpha1 spec struct, so that they can be stored in
// an annotation and used when converting to v1alpha2.
type stashedMySQLAADUserFields struct {
	DatabaseRoles map[string][]string `json:"databaseRoles,omitempty"`
	Roles         []string            `json:"roles,omitempty"`
}
