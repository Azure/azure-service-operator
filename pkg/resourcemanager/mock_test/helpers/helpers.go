package helpers

import (
	"fmt"
	"github.com/Azure/go-autorest/autorest"
	"net/http"
)

func GetRestResponse(statusCode int) autorest.Response {
	return autorest.Response{
		Response: &http.Response{
			Status:     fmt.Sprintf("%d", statusCode),
			StatusCode: statusCode,
		},
	}
}
