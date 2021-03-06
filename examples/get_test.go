package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/arock95/go-httpclient/gohttp"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start testing up in here...")
	gohttp.StartMockServer()

	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		mock := gohttp.Mock{
			Method: http.MethodGet,
			Url:    "http://api.github.com",
			Error:  errors.New("timeout getting github endpoint"),
		}

		gohttp.AddMock(mock)

		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("an error expected")
		}

		if err.Error() != "timeout getting github endpoint" {
			fmt.Println(err.Error())
			t.Error("invalid error message")
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		mock := gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "http://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		}

		gohttp.AddMock(mock)

		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("an error expected")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal") {
			t.Error("invalid error message")
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		mock := gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "http://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "http://api.github.com"}`,
		}

		gohttp.AddMock(mock)

		endpoints, err := GetEndpoints()
		if endpoints == nil {
			t.Error("endpoints expected at this point")
		}

		if err != nil {
			t.Error("no error expected")
		}

		if endpoints.CurrentUser != "http://api.github.com" {
			t.Error("invalid user url")
		}

	})

}
