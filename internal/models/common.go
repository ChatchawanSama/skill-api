package models

type AppResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type AppErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)
