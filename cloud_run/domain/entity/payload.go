package entity

type Payload struct {
	Gamertag string `json:"gamertag"`
	Filter   string `json:"filter,omitempty"`
}
