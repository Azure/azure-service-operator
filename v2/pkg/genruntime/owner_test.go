/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime_test

import (
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_ApplyObjAndEnsureOwner(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	owner := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "myconfigmap",
			Namespace: tc.Namespace,
		},
	}
	tc.CreateResource(owner)

	// Manually set TypeMeta stuff
	owner.Kind = "ConfigMap"
	owner.APIVersion = "v1"

	obj := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "mysecret",
			Namespace: tc.Namespace,
		},
		StringData: map[string]string{
			"a": "b",
		},
	}
	result, err := genruntime.ApplyObjAndEnsureOwner(tc.Ctx, tc.KubeClient(), owner, obj)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(result).To(Equal(controllerutil.OperationResultCreated))

	var storedObj v1.Secret
	tc.GetResource(types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetName()}, &storedObj)
	tc.Expect(storedObj.Data).ToNot(BeNil())
	tc.Expect(storedObj.Data).To(HaveLen(1)) // Only 1 key should be here
	tc.Expect(storedObj.Data).To(HaveKey("a"))
	tc.Expect(storedObj.Data["a"]).To(Equal([]byte("b")))

	// Update Obj
	obj.StringData = map[string]string{
		"c": "d",
	}
	result, err = genruntime.ApplyObjAndEnsureOwner(tc.Ctx, tc.KubeClient(), owner, obj)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(result).To(Equal(controllerutil.OperationResultUpdated))

	tc.GetResource(client.ObjectKey{Namespace: obj.Namespace, Name: obj.Name}, &storedObj)
	tc.Expect(storedObj.Data).ToNot(BeNil())
	tc.Expect(storedObj.Data).To(HaveLen(1)) // Only 1 key should be here
	tc.Expect(storedObj.Data).To(HaveKey("c"))
	tc.Expect(storedObj.Data["c"]).To(Equal([]byte("d")))
}
