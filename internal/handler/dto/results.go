package dto

type Result struct {
	Result any     `json:"result,omitempty"`
	Error  *string `json:"error,omitempty"`
}
