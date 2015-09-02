package models

type User struct {
	UserId int    `json:"UserId"`
	Name   string `json:"Name"`
	Email  string `json:"Email"`
	Items  []Item `json:"Items"`
}
