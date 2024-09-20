package models

type Email struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Sender  string `json:"sender"`
	Body    string `json:"body"`
}
