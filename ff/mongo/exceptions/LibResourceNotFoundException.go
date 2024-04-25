package ff_mongo_exceptions

import "net/http"

type LibResourceNotFoundException struct {
	code     int
	messages []string
}

const STATUS_NOT_FOUND = http.StatusNotFound

func NewLibResourceNotFoundException(messages []string) *LibResourceNotFoundException {
	return &LibResourceNotFoundException{
		code:     STATUS_NOT_FOUND,
		messages: messages,
	}
}

func NewLibResourceNotFoundExceptionSglMsg(message string) *LibResourceNotFoundException {
	return &LibResourceNotFoundException{
		code:     STATUS_NOT_FOUND,
		messages: []string{message},
	}
}

func (this LibResourceNotFoundException) GetCode() int {
	return this.code
}

func (this LibResourceNotFoundException) GetMessages() []string {
	return this.messages
}

func (this LibResourceNotFoundException) Error() string {
	var message string
	for _, value := range this.messages {
		message = message + value
	}
	return message
}
