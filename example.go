package main

import (
	"fmt"
	"net/http"

	"time"

	"github.com/arock95/go-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetMaxIdleConnections(5).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(50 * time.Millisecond).
		DisableTimeouts(true).
		Build()

	return client
}

func main() {
	for i:=0; i < 5; i++ {
		go func(){
			getUrls()
		}()
	}
	time.Sleep(time.Second * 2)
}

func getUrls() {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer 123ABC")

	response, err := githubHttpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status())
	fmt.Println(response.String())

	// var user User
	// response.UnmarshalJson(&user)

	// fmt.Println(user.FirstName)
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// func createUser(user User) {
// 	headers := make(http.Header)
// 	headers.Set("Authorization", "Bearer 123ABC")

// 	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()

// 	bytes, _ := ioutil.ReadAll(response.Body)
// 	fmt.Println(string(bytes))
// }
