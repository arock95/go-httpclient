package gohttp

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// init
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	client.builder.headers = commonHeaders

	headers := make(http.Header)
	headers.Set("Authorization", "Token asdf")

	// exec
	finalHeaders := client.getRequestHeaders(headers)

	// validation
	if len(finalHeaders) != 2 {
		t.Error("expected 2 headers!")
	}

	if finalHeaders.Get("Authorization") != "Token asdf" {
		t.Error("authorization header is incorrect")
	}
}


func TestGetRequestBody(t *testing.T) {
	// init
	client := httpClient{}

	t.Run("BodyWithJson", func (t *testing.T) {
		contentType := "application/json"

		type User struct {
			FirstName string `json:"first_name"`
			LastName string `json:"last_name"`
		}

		user := User {
			FirstName: "Anthony",
			LastName: "Rocchio",
		}
		valUser := &User{}

		// exe
		bytes, _ := client.getRequestBody(contentType, user)

		// val
		json.Unmarshal(bytes, valUser)
		if user != *valUser {
			t.Errorf("error creating request body!")
		}
	})

	t.Run("nilBody", func(t *testing.T) {
		contentType := "application/json"

		// exe
		bytes, err := client.getRequestBody(contentType, nil)

		if err != nil {
			t.Errorf("nil body error!")
		}

		// val
		if bytes != nil {
			t.Errorf("nil body error!")
		}
	})

	t.Run("BodyWithDefaultContentType", func (t *testing.T) {
		contentType := "fake"

		type User struct {
			FirstName string `json:"first_name"`
			LastName string `json:"last_name"`
		}

		user := User {
			FirstName: "Anthony",
			LastName: "Rocchio",
		}
		valUser := &User{}

		// exe
		bytes, _ := client.getRequestBody(contentType, user)

		// val
		json.Unmarshal(bytes, valUser)
		if user != *valUser {
			t.Errorf("error creating request body!")
		}
	})
}
