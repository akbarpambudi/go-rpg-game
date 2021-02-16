package gokithelper

type defaultResponse struct {
	successEvent interface{}
	failedEvent  error
}

func NewSuccessResponse(successEvent interface{}) *defaultResponse {
	return &defaultResponse{
		successEvent: successEvent,
		failedEvent:  nil,
	}
}

func NewFailedResponse(err error) *defaultResponse {
	return &defaultResponse{
		successEvent: nil,
		failedEvent:  err,
	}
}

func (s defaultResponse) SuccessEvent() interface{} {
	return s.successEvent
}

func (s defaultResponse) FailedEvent() *ErrResponse {
	if s.failedEvent == nil {
		return nil
	}
	return &ErrResponse{Message: s.failedEvent.Error()}
}
