package dto

type Result struct {
	Result interface{} `json:"result,omitempty"`
	Error  *string     `json:"error,omitempty"`
}
