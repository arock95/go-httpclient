package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/arock95/go-httpclient/gohttp"
)

func TestGet(t *testing.T) {
	// Tell Http library to mock
	gohttp.StartMockServer()

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		mock := gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoint"),
		}

		gohttp.AddMock(mock)

		endpoints, err := GetEndpoints()
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		mock := gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			RequestBody:        `{"current_user_url": 123}`,
		}

		gohttp.AddMock(mock)

		endpoints, err := GetEndpoints()
	})

	t.Run("TestNoError", func(t *testing.T) {
		mock := gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			RequestBody:        `{"current_user_url": "https://api.github.com"}`,
		}

		gohttp.AddMock(mock)

		endpoints, err := GetEndpoints()
	})

	endpoints, err := GetEndpoints()
	if err != nil {
		t.Errorf("expected bleh received bleh2")
	}
	fmt.Println(err)
	fmt.Println(endpoints)
}
