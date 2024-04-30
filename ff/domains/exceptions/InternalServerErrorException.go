package ff_domains_exceptions

import "net/http"

type InternalServerErrorException struct {
	code     int
	messages []string
}

const STATUS_INTERNAL_SERVER_ERROR = http.StatusInternalServerError

func NewInternalServerErrorException(messages []string) *InternalServerErrorException {
	return &InternalServerErrorException{
		code:     STATUS_INTERNAL_SERVER_ERROR,
		messages: messages,
	}
}

func NewInternalServerErrorExceptionSglMsg(message string) *InternalServerErrorException {
	return &InternalServerErrorException{
		code:     STATUS_INTERNAL_SERVER_ERROR,
		messages: []string{message},
	}
}

func (this InternalServerErrorException) GetCode() int {
	return this.code
}

func (this InternalServerErrorException) GetMessages() []string {
	return this.messages
}

func (this InternalServerErrorException) Error() string {
	var message string
	for _, value := range this.messages {
		message = message + value
	}
	return message
}
