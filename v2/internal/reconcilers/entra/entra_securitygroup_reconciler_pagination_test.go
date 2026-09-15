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

	"github.com/google/uuid"
	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

func TestCollectDirectoryObjectIDs_PaginatesAndDedupes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	pages := map[string]msgraphmodels.DirectoryObjectCollectionResponseable{
		"first": makeDirectoryObjectPage(
			[]string{"AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA", "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"},
			stringPtr("second"),
		),
		"second": makeDirectoryObjectPage(
			[]string{"BBBBBBBB-BBBB-BBBB-BBBB-BBBBBBBBBBBB", "cccccccc-cccc-cccc-cccc-cccccccccccc", ""},
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
	g.Expect(ids).To(Equal([]uuid.UUID{
		uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
		uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
		uuid.MustParse("cccccccc-cccc-cccc-cccc-cccccccccccc"),
	}))
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
				[]string{"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"},
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
