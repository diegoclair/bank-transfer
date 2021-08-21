package viewmodel

type Login struct {
	DocumentNumber string `json:"document_number,omitempty"`
	Secret         string `json:"secret,omitempty"`
}
