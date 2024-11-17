/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type simpleObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

func (s simpleObject) DeepCopyObject() runtime.Object {
	panic("Not implemented")
}

var _ client.Object = &simpleObject{}

func TestOwnerLookup(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Build a structure that looks like:
	//         obj1
	//         /   \
	//       obj2   obj3
	//       / \      |
	//    obj4 obj5  obj6

	obj1 := &simpleObject{
		ObjectMeta: metav1.ObjectMeta{
			UID: types.UID("1"),
		},
	}
	obj2 := &simpleObject{
		ObjectMeta: metav1.ObjectMeta{
			UID: types.UID("2"),
			OwnerReferences: []metav1.OwnerReference{
				{
					UID: obj1.UID,
				},
			},
		},
	}
	obj3 := &simpleObject{
		ObjectMeta: metav1.ObjectMeta{
			UID: types.UID("3"),
			OwnerReferences: []metav1.OwnerReference{
				{
					UID: obj1.UID,
				},
			},
		},
	}
	obj4 := &simpleObject{
		ObjectMeta: metav1.ObjectMeta{
			UID: types.UID("4"),
			OwnerReferences: []metav1.OwnerReference{
				{
					UID: obj2.UID,
				},
			},
		},
	}
	obj5 := &simpleObject{
		ObjectMeta: metav1.ObjectMeta{
			UID: types.UID("5"),
			OwnerReferences: []metav1.OwnerReference{
				{
					UID: obj2.UID,
				},
			},
		},
	}
	obj6 := &simpleObject{
		ObjectMeta: metav1.ObjectMeta{
			UID: types.UID("6"),
			OwnerReferences: []metav1.OwnerReference{
				{
					UID: obj3.UID,
				},
			},
		},
	}

	objs := []client.Object{
		obj1, obj2, obj3, obj4, obj5, obj6,
	}
	ranks := objectRanksByOwner(objs...)
	g.Expect(ranks).To(HaveLen(3))
	g.Expect(ranks[0]).To(HaveLen(1))
	g.Expect(ranks[0][0]).To(Equal(obj1))
	g.Expect(ranks[1]).To(HaveLen(2))
	g.Expect(ranks[1][0]).To(Equal(obj2))
	g.Expect(ranks[1][1]).To(Equal(obj3))
	g.Expect(ranks[2]).To(HaveLen(3))
	g.Expect(ranks[2][0]).To(Equal(obj4))
	g.Expect(ranks[2][1]).To(Equal(obj5))
	g.Expect(ranks[2][2]).To(Equal(obj6))
}
