package entity

type AppResponse struct {
	Status string
	Data   interface{}
}

type AppErrorResponse struct {
	Status  string
	Message string
	Reason  string
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)
