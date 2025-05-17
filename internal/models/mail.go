package models

type Mail struct {
	To     string `json:"to,omitempty"`
	Header string `json:"header,omitempty"`
	Body   string `json:"body,omitempty"`
}
