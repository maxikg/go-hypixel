package hypixel

// KeyInfoResponse is the bare response of the API containing maybe the record, maybe a cause and a bool determine if
// the request was successful.
type KeyInfoResponse struct {
	Record  *KeyInfo `json:"record"`
	Cause   string   `json:"cause"`
	Success bool     `json:"success"`
}

// KeyInfo contains the information of a /key request.
type KeyInfo struct {
	Key              string `json:"key"`
	OwnerUUID        string `json:"ownerUuid"`
	TotalQueries     int32  `json:"totalQueries"`
	QueriesInPastMin int32  `json:"queriesInPastMin"`
}

// GuildIdResponse is the bare response of the API containing maybe a guild id, maybe a cause and a bool determine if
// the request was successful.
type GuildIdResponse struct {
	Guild	string   `json:"guild"`
	Cause   string   `json:"cause"`
	Success bool     `json:"success"`
}

// FriendsResponse is the bare response of the API containing maybe an array of friends, maybe a cause and a bool
// determine if the request was successful.
type FriendsResponse struct {
	Records []*Friend `json:"records"`
	Cause   string    `json:"cause"`
	Success bool      `json:"success"`
}

// Friend contains the information of a /friends records entry.
type Friend struct {
	ID 		     string `json:"_id"`
	Receiver     string `json:"receiver"`
	Sender       string `json:"sender"`
	UUIDReceiver string `json:"uuidReceiver"`
	UUIDSender   string `json:"uuidSender"`
	Started      int64  `json:"started"`
}
