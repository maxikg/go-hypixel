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

// Query queries the remote API for a specified method, an optional map of parameters and an instance of a struct in
// which the response will be deserialized.
func(c *Client) Query(method string, parameters map[string]string, result interface{}) error {
	req, err := c.CreateRequest(method, parameters)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, result)
}

// KeyInfo calls the API at /key. It will returns an error or a KeyInfo which contains statistics about the usage of
// the currently used key.
func(c *Client) KeyInfo() (map[string]interface{}, error) {
	result := &KeyInfoResponse{}
	err := c.Query("key", nil, result)
	if err != nil {
		return nil, err
	} else if result.Success == false {
		return nil, errors.New(result.Cause)
	}
	return result.Record, nil
}

// FindGuildByPlayer finds a guild id by an name of a guilds member. Returns an empty string as the id if no guild data
// is available.
func(c *Client) FindGuildByPlayer(player string) (string, error) {
	return c.findGuild("byPlayer", player)
}

// FindGuildByUUID finds a guild id by an uuid of a guilds member. Returns an empty string as the id if no guild data is
// available.
func(c *Client) FindGuildByUUID(uuid string) (string, error) {
	return c.findGuild("byUuid", uuid)
}

// FindGuildByName finds a guild id by the guilds name. Returns an empty string as the id if no guild data is available.
func(c *Client) FindGuildByName(name string) (string, error) {
	return c.findGuild("byName", name)
}

// Internal helper method which queries for guild information using a parameterName and a parameterValue. Returns an
// empty string as the id if no guild data is available.
func(c *Client) findGuild(parameterName string, parameterValue string) (string, error) {
	result := &GuildIDResponse{}
	err := c.Query("findGuild", map[string]string{ parameterName: parameterValue }, result)
	if err != nil {
		return "", err
	} else if result.Success == false {
		return "", errors.New(result.Cause)
	}
	return result.Guild, nil
}

// Guild queries a guilds information by its id. Returns nil if no guild with this id is known.
func(c *Client) Guild(id string) (map[string]interface{}, error) {
	result := &GuildResponse{}
	err := c.Query("guild", map[string]string{ "id": id }, result)
	if err != nil {
		return nil, err
	} else if result.Success == false {
		return nil, errors.New(result.Cause)
	}
	return result.Guild, nil
}

// FriendsByName queries friends of the player name.
func(c *Client) FriendsByName(name string) ([]map[string]interface{}, error) {
	return c.friends("player", name)
}

// FriendsByUUID queries friends of the player uuid.
func(c *Client) FriendsByUUID(uuid string) ([]map[string]interface{}, error) {
	return c.friends("uuid", uuid)
}

// Internal helper method which queries for friends using a parameterName and a parameterValue. Returns an array of
// map[string]interface{}.
func(c *Client) friends(parameterName string, parameterValue string) ([]map[string]interface{}, error) {
	result := &FriendsResponse{}
	err := c.Query("friends", map[string]string{ parameterName: parameterValue }, result)

	if err != nil {
		return nil, err
	} else if result.Success == false {
		return nil, errors.New(result.Cause)
	}
	return result.Records, nil
}

// PlayerByName queries a player by its name.
func(c *Client) PlayerByName(name string) (map[string]interface{}, error) {
	return c.player("name", name)
}

// PlayerByUUID queries a player by its uuid.
func(c *Client) PlayerByUUID(uuid string) (map[string]interface{}, error) {
	return c.player("uuid", uuid)
}
// Internal helper method which queries for a player using a parameterName and a parameterValue. Returns a
// map[string]interface{}.
func(c *Client) player(parameterName string, parameterValue string) (map[string]interface{}, error) {
	result := &PlayerResponse{}
	err := c.Query("player", map[string]string{ parameterName: parameterValue }, result)

	if err != nil {
		return nil, err
	} else if result.Success == false {
		return nil, errors.New(result.Cause)
	}
	return result.Player, nil
}

// SessionByName queries a players session by its name.
func(c *Client) SessionByName(name string) (map[string]interface{}, error) {
	return c.session("player", name)
}

// SessionByUUID queries a players session by its uuid.
func(c *Client) SessionByUUID(uuid string) (map[string]interface{}, error) {
	return c.session("uuid", uuid)
}

// Internal helper method which queries for a players session using a parameterName and a parameterValue. Returns a
// map[string]interface{}.
func(c *Client) session(parameterName string, parameterValue string) (map[string]interface{}, error) {
	result := &SessionResponse{}
	err := c.Query("session", map[string]string{ parameterName: parameterValue }, result)

	if err != nil {
		return nil, err
	} else if result.Success == false {
		return nil, errors.New(result.Cause)
	}
	return result.Session, nil
}
