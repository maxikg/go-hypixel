package hypixel

import (
	"net/http"
	"html"
	"encoding/json"
	"io/ioutil"
	"errors"
)

const (
	defaultBaseUrl = "https://api.hypixel.net/"
	userAgent      = "go-hypixel/0.1 (see https://github.com/maxikg/go-hypixel)"
)

// A Client manages communication with the Hypixel API.
type Client struct {
	// The http.Client which is used to communicate with the Hypixel API.
	client    *http.Client

	// The base url of the API requests. Default to the Hypixel API.
	BaseUrl   string

	// The api key.
	ApiKey    string

	// The user agent. Default to "go-hypixel/0.1 (see https://github.com/maxikg/go-hypixel)".
	UserAgent string
}

// NewClient returns a new Hypixel API client. If a nil httpClient is provided, http.DefaultClient will be used.
func NewClient(key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{client: httpClient, BaseUrl: defaultBaseUrl, UserAgent: userAgent, ApiKey: key}
}

// CreateRequest creates an http.Request for querying the API. method specify which method should be called. Optional
// a map can be provided for additional parameters.
func(c *Client) CreateRequest(method string, params map[string]string) (*http.Request, error) {
	url := c.BaseUrl + method + "?key=" + c.ApiKey
	if params != nil {
		for key, value := range params {
			url += "&" + html.EscapeString(key) + "=" + html.EscapeString(value)
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// KeyInfo calls the API at /key. It will returns an error or a KeyInfo which contains statistics about the usage of
// the currently used key.
func(c *Client) KeyInfo() (*KeyInfo, error) {
	req, err := c.CreateRequest("key", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &KeyInfoResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.Success == false {
		return nil, errors.New(result.Cause)
	}
	return result.Record, nil
}
