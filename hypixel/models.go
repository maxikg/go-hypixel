package hypixel

// KeyInfoResponse is the bare response of the API containing maybe the record, maybe a cause and a bool determine if
// the request was successful.
type KeyInfoResponse struct {
	Record  map[string]interface{} `json:"record"`
	Cause   string                 `json:"cause"`
	Success bool                   `json:"success"`
}

// GuildIdResponse is the bare response of the API containing maybe a guild id, maybe a cause and a bool determine if
// the request was successful.
type GuildIdResponse struct {
	Guild	string   `json:"guild"`
	Cause   string   `json:"cause"`
	Success bool     `json:"success"`
}

// GuildResponse is the bare response of the API containing maybe a guild, maybe a cause and a bool determine if the
// request was successful.
type GuildResponse struct {
	Guild   map[string]interface{} `json:"guild"`
	Cause   string                 `json:"cause"`
	Success bool                   `json:"success"`
}

// FriendsResponse is the bare response of the API containing maybe an array of friends, maybe a cause and a bool
// determine if the request was successful.
type FriendsResponse struct {
	Records []map[string]interface{} `json:"records"`
	Cause   string                   `json:"cause"`
	Success bool                     `json:"success"`
}
