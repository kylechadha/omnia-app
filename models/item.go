package models

type Item struct {
	ItemId      int    `json:"ItemId"`
	Description string `json:"Description"`
	Category    string `json:"Category"`
}
