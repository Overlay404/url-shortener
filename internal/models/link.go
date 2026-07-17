package models

import "encoding/json"

type Link struct {
	Url     string
	Clicks  int
	Created int
}

func (l *Link) MarshalBinary() ([]byte, error) {
	return json.Marshal(l)
}

func (l *Link) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &l)
}
