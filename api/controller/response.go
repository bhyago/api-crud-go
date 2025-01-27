package controller

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseMessageError struct {
	Error string `json:"error"`
}

func NewResponseMessage(message string) *ResponseMessage {
	return &ResponseMessage{Message: message}
}

func NewResponseMessageError(err string) *ResponseMessageError {
	return &ResponseMessageError{Error: err}
}
