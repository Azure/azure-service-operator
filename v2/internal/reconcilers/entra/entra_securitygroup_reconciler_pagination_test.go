/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"context"
	"errors"
	"testing"

	. "github.com/onsi/gomega"

	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
)

func TestRelationshipSidesToManage_OmittedVsEmpty(t *testing.T) {
	t.Parallel()

	ownerRef := asoentra.SecurityGroupMemberReference{ObjectID: stringPtr("11111111-1111-1111-1111-111111111111")}
	memberRef := asoentra.SecurityGroupMemberReference{ObjectID: stringPtr("22222222-2222-2222-2222-222222222222")}

	cases := map[string]struct {
		spec        asoentra.SecurityGroupSpec
		wantOwners  bool
		wantMembers bool
	}{
		"both omitted are unmanaged": {
			spec:        asoentra.SecurityGroupSpec{},
			wantOwners:  false,
			wantMembers: false,
		},
		"owners explicit empty is managed": {
			spec:        asoentra.SecurityGroupSpec{Owners: []asoentra.SecurityGroupMemberReference{}},
			wantOwners:  true,
			wantMembers: false,
		},
		"members explicit empty is managed": {
			spec:        asoentra.SecurityGroupSpec{Members: []asoentra.SecurityGroupMemberReference{}},
			wantOwners:  false,
			wantMembers: true,
		},
		"both present are managed": {
			spec: asoentra.SecurityGroupSpec{
				Owners:  []asoentra.SecurityGroupMemberReference{ownerRef},
				Members: []asoentra.SecurityGroupMemberReference{memberRef},
			},
			wantOwners:  true,
			wantMembers: true,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			owners, members := relationshipSidesToManage(tc.spec)
			g.Expect(owners).To(Equal(tc.wantOwners))
			g.Expect(members).To(Equal(tc.wantMembers))
		})
	}
}

func TestCollectDirectoryObjectIDs_PaginatesAndDedupes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	pages := map[string]msgraphmodels.DirectoryObjectCollectionResponseable{
		"first": makeDirectoryObjectPage(
			[]string{"owner-a", "owner-b"},
			stringPtr("second"),
		),
		"second": makeDirectoryObjectPage(
			[]string{"owner-b", "owner-c", ""},
			nil,
		),
	}

	ids, err := collectDirectoryObjectIDs(
		context.Background(),
		func(context.Context) (msgraphmodels.DirectoryObjectCollectionResponseable, error) {
			return pages["first"], nil
		},
		func(nextLink string) (msgraphmodels.DirectoryObjectCollectionResponseable, error) {
			return pages[nextLink], nil
		},
	)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(ids).To(Equal([]string{"owner-a", "owner-b", "owner-c"}))
}

func TestCollectDirectoryObjectIDs_FirstPageError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	_, err := collectDirectoryObjectIDs(
		context.Background(),
		func(context.Context) (msgraphmodels.DirectoryObjectCollectionResponseable, error) {
			return nil, errors.New("first page failed")
		},
		func(string) (msgraphmodels.DirectoryObjectCollectionResponseable, error) {
			return nil, nil
		},
	)

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("first page failed"))
}

func TestCollectDirectoryObjectIDs_NextPageError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	_, err := collectDirectoryObjectIDs(
		context.Background(),
		func(context.Context) (msgraphmodels.DirectoryObjectCollectionResponseable, error) {
			return makeDirectoryObjectPage(
				[]string{"owner-a"},
				stringPtr("next"),
			), nil
		},
		func(string) (msgraphmodels.DirectoryObjectCollectionResponseable, error) {
			return nil, errors.New("next page failed")
		},
	)

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("next page failed"))
}

func makeDirectoryObjectPage(ids []string, nextLink *string) msgraphmodels.DirectoryObjectCollectionResponseable {
	response := msgraphmodels.NewDirectoryObjectCollectionResponse()
	values := make([]msgraphmodels.DirectoryObjectable, 0, len(ids))
	for _, id := range ids {
		obj := msgraphmodels.NewDirectoryObject()
		if id != "" {
			idCopy := id
			obj.SetId(&idCopy)
		}
		values = append(values, obj)
	}
	response.SetValue(values)
	response.SetOdataNextLink(nextLink)
	return response
}

func stringPtr(value string) *string {
	return &value
}
