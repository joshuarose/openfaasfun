package function

import (
	"fmt"
	"io"
	"net/http"

	handler "github.com/openfaas/templates-sdk/go-http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	resp, err := http.Get("https://api.kanye.rest/")

	if err != nil {
		// handle error
		return handler.Response{
			Body:       []byte(fmt.Sprint("Kanye isn't home")),
			StatusCode: http.StatusBadRequest,
		}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return handler.Response{
		Body:       body,
		StatusCode: http.StatusOK,
	}, err
}
