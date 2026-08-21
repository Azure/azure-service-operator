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
)

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
