package models

// @Success 200 {array} models.Patient

type Patient struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Gender  string `json:"Gender"`
	Address string `json:"Address"`
}
