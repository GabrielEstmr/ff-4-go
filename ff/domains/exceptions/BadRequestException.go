package ff_domains_exceptions

import "net/http"

type BadRequestException struct {
	code     int
	messages []string
}

const STATUS_BAD_REQUEST = http.StatusBadRequest

func NewBadRequestException(messages []string) *BadRequestException {
	return &BadRequestException{
		code:     STATUS_BAD_REQUEST,
		messages: messages,
	}
}

func NewBadRequestExceptionSglMsg(message string) *BadRequestException {
	return &BadRequestException{
		code:     STATUS_BAD_REQUEST,
		messages: []string{message},
	}
}

func (this BadRequestException) GetCode() int {
	return this.code
}

func (this BadRequestException) GetMessages() []string {
	return this.messages
}

func (this BadRequestException) Error() string {
	var message string
	for _, value := range this.messages {
		message = message + value
	}
	return message
}
