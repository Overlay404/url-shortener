package models

import "encoding/json"

type Link struct {
	Url     string `json:"url"`
	Clicks  int    `json:"clicks"`
	Created int    `json:"created"`
}

func (l *Link) MarshalBinary() ([]byte, error) {
	return json.Marshal(l)
}

func (l *Link) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &l)
}
