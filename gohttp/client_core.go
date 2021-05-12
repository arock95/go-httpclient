package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = time.Second * 5
	defaultConnectionTimeout  = time.Second * 1
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*Response, error) {
	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, errors.New("unable to create request")
	}

	
	if mock := mockupServer.getMock(method, url, string(requestBody)); mock != nil {
		resp, err := mock.GetResponse()
		return resp, err
	}


	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, errors.New("unable to create request")
	}

	request.Header = fullHeaders

	client := c.getHttpClient()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	finalResponse := Response{
		statusCode: response.StatusCode,
		headers:    response.Header,
		body:       responseBody,
		status:     response.Status,
	}
	return &finalResponse, nil
}

func (c *httpClient) getHttpClient() *http.Client {
	// sync.Once.Do --- a function that regardless of how many goroutines call it, will only get done once!
	c.clientOnce.Do(
		func() {
			c.client = &http.Client{
				Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
				Transport: &http.Transport{
					MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
					ResponseHeaderTimeout: c.getResponseTimeout(),
					DialContext: (&net.Dialer{
						Timeout: c.getConnectionTimeout(),
					}).DialContext,
				},
			}
		})

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultConnectionTimeout
}

func (c *httpClient) getRequestHeaders(headers http.Header) http.Header {
	result := make(http.Header)
	// add common headers
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// add custom headers
	for header, value := range headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}

}
