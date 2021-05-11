package examples

import (
	"fmt"
)

type Endpoints struct {
	CurrentUser       string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("http://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Status Code: %d", response.StatusCode())
	fmt.Printf("Body: %s\n", response.String())

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Println(endpoints.CurrentUser)

	return &endpoints, nil
}
