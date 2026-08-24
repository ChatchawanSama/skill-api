package dto

type AppResponse struct {
	Status string
	Data   interface{}
}

type AppErrorResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)
