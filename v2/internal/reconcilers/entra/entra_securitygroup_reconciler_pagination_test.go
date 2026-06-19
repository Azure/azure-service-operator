/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"context"
	"errors"
	"testing"

	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	. "github.com/onsi/gomega"

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

	pages := map[string]msgraphmodels.StringCollectionResponseable{
		"first": makeStringCollectionPage(
			[]string{
				"https://graph.microsoft.com/v1.0/directoryObjects/owner-a",
				"https://graph.microsoft.com/v1.0/directoryObjects/owner-b",
			},
			stringPtr("second"),
		),
		"second": makeStringCollectionPage(
			[]string{
				"owner-b",
				"https://graph.microsoft.com/v1.0/directoryObjects/owner-c",
				"   ",
			},
			nil,
		),
	}

	ids, err := collectDirectoryObjectIDs(
		context.Background(),
		func(context.Context) (msgraphmodels.StringCollectionResponseable, error) {
			return pages["first"], nil
		},
		func(nextLink string) (msgraphmodels.StringCollectionResponseable, error) {
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
		func(context.Context) (msgraphmodels.StringCollectionResponseable, error) {
			return nil, errors.New("first page failed")
		},
		func(string) (msgraphmodels.StringCollectionResponseable, error) {
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
		func(context.Context) (msgraphmodels.StringCollectionResponseable, error) {
			return makeStringCollectionPage(
				[]string{"https://graph.microsoft.com/v1.0/directoryObjects/owner-a"},
				stringPtr("next"),
			), nil
		},
		func(string) (msgraphmodels.StringCollectionResponseable, error) {
			return nil, errors.New("next page failed")
		},
	)

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("next page failed"))
}

func makeStringCollectionPage(values []string, nextLink *string) msgraphmodels.StringCollectionResponseable {
	response := msgraphmodels.NewStringCollectionResponse()
	response.SetValue(values)
	response.SetOdataNextLink(nextLink)
	return response
}

func stringPtr(value string) *string {
	return &value
}
