package gokithelper

type ErrResponse struct {
	Message string `json:"message"`
}

type Response interface {
	SuccessEvent() interface{}
	FailedEvent() *ErrResponse
}
