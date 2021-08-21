package entity

// Authentication data model
type Authentication struct {
	Token      string `json:"token"`
	ValidTime  int64  `json:"valid_time"`
	ServerTime int64  `json:"server_time"`
}
