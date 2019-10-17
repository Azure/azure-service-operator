package autorest

// import (
// 	"net/http"
// 	"strings"
// 	. "github.com/Azure/go-autorest/autorest"
// )

// type SharedKeyAuthorizer struct {
// 	storageAccountName 	string
// 	storageAccountKey 	string
// }

// func NewSharedKeyAuthorizer(accountName string, accountKey string) *SharedKeyAuthorizer {
// 	return &SharedKeyAuthorizer{
// 		storageAccountName: accountName,
// 		storageAccountKey: accountKey,
// 	}
// }

// func (ska *SharedKeyAuthorizer) WithAuthorization() PrepareDecorator {
// 	return func(p Preparer) Preparer {
// 		return PreparerFunc(func(r *http.Request) (*http.Request, error) {
// 			r, err := p.Prepare(r)
// 			if err != nil {
// 				return r, err
// 			}

// 			key, err := buildSharedKey(ska.storageAccountName, ska.storageAccountKey, r)
// 			if err != nil {
// 				return r, err
// 			}

// 			sharedKeyHeader := formatSharedKeyAuthorizationHeader(ska.storageAccountName, *key)
// 			return Prepare(r, WithHeader(HeaderAuthorization, sharedKeyHeader))
// 		})
// 	}
// }