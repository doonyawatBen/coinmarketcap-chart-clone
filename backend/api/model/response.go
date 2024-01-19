package model

type Response struct {
	Status    string          `json:"status,omitempty"`
	Timestamp string          `json:"timestamp,omitempty"`
	Error     []ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	FailedField  string `json:"failedField,omitempty"`
	Tag          string `json:"tag,omitempty"`
	Value        string `json:"value,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
