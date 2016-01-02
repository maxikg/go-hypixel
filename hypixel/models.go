package hypixel

// The bare response of the API containing maybe the record, maybe a cause and a bool determine if the request was
// success.
type KeyInfoResponse struct {
	Record  *KeyInfo `json:"record"`
	Cause   string   `json:"cause"`
	Success bool     `json:"success"`
}

// The information of a KeyInfo request.
type KeyInfo struct {
	Key              string `json:"key"`
	OwnerUuid        string `json:"ownerUuid"`
	TotalQueries     int32  `json:"totalQueries"`
	QueriesInPastMin int32  `json:"queriesInPastMin"`
}
