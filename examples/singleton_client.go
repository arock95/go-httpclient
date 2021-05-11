package examples

import (
	"time"

	"github.com/arock95/go-httpclient/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(50 * time.Millisecond).
		Build()

	return client
}
