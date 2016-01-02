package hypixel

import (
	"net/http"
	"html"
	"encoding/json"
	"io/ioutil"
	"errors"
)

const (
	defaultBaseURL = "https://api.hypixel.net/"
	userAgent      = "go-hypixel/0.1 (see https://github.com/maxikg/go-hypixel)"
)

// A Client manages communication with the Hypixel API.
type Client struct {
	// The http.Client which is used to communicate with the Hypixel API.
	client    *http.Client

	// The base url of the API requests. Default to the Hypixel API.
	BaseURL   string

	// The api key.
	APIKey    string

	// The user agent. Default to "go-hypixel/0.1 (see https://github.com/maxikg/go-hypixel)".
	UserAgent string
}

// NewClient returns a new Hypixel API client. If a nil httpClient is provided, http.DefaultClient will be used.
func NewClient(key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{client: httpClient, BaseURL: defaultBaseURL, UserAgent: userAgent, APIKey: key}
}

// CreateRequest creates an http.Request for querying the API. method specify which method should be called. Optional
// a map can be provided for additional parameters.
func(c *Client) CreateRequest(method string, params map[string]string) (*http.Request, error) {
	url := c.BaseURL + method + "?key=" + c.APIKey
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

// Find a guild id by an name of a guilds member. Returns an empty string as the id if no guild data is available.
func(c *Client) FindGuildByPlayer(player string) (string, error) {
	return c.findGuild("byPlayer", player)
}

// Find a guild id by an uuid of a guilds member. Returns an empty string as the id if no guild data is available.
func(c *Client) FindGuildByUuid(uuid string) (string, error) {
	return c.findGuild("byUuid", uuid)
}

// Find a guild id by the guilds name. Returns an empty string as the id if no guild data is available.
func(c *Client) FindGuildByName(name string) (string, error) {
	return c.findGuild("byName", name)
}

// Internal helper method which queries for guild information using a parameterName and a parameterValue. Returns an
// empty string as the id if no guild data is available.
func(c *Client) findGuild(parameterName string, parameterValue string) (string, error) {
	req, err := c.CreateRequest("findGuild", map[string]string{ parameterName: parameterValue })
	if err != nil {
		return "", err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	result := &GuildIdResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if result.Success == false {
		return "", errors.New(result.Cause)
	}
	return result.Guild, nil
}
