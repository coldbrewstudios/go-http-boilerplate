package models

type Law struct {
	Name       string `json:"name,omitempty"`
	Definition string `json:"definition,omitempty"`
}
