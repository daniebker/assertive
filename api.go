package assertive

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

// API Struct
type API struct {
	client  *http.Client
	url     *url.URL
	request *http.Request
}

// NewAPI sets up an api url
func NewAPI(host string) *API {
	api := API{}
	api.client = &http.Client{}
	parsedURL, _ := url.Parse(host)
	api.url = parsedURL
	return &api
}

// Get creates a GET request
func (api *API) Get(path string) *API {
	request := makeRequest("GET")

	api.url.Path = path
	request.URL = api.url
	api.request = request
	return api
}

// PostJSON creates a JSON POST request
func (api *API) PostJSON(path string, json []byte) *API {
	request, _ := http.NewRequest("POST", fmt.Sprintf("%s://%s/%s", api.url.Scheme, api.url.Host, path), bytes.NewBuffer(json))
	request.Header = make(map[string][]string)
	api.url.Path = path
	api.request = request
	return api.WithHeader("Content-Type", "application/json")
}

// DoRequest makes the request to the API
func (api *API) DoRequest() (*http.Response, error) {
	response, err := api.client.Do(api.request)

	if err != nil {
		err = fmt.Errorf("Request %s failed with %s", api.request.URL, err)
		return response, err
	}

	return response, nil
}

// WithHeader adds a header to the request
func (api *API) WithHeader(key, value string) *API {
	api.request.Header.Add(key, value)
	return api
}

func makeRequest(method string) *http.Request {
	request := &http.Request{}
	request.Method = method
	request.Header = make(map[string][]string)
	request.Header.Add("User-Agent", "Assertive test Suite")
	return request
}
