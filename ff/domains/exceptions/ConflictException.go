package ff_domains_exceptions

import (
	"net/http"
)

type ConflictException struct {
	code     int
	messages []string
}

const STATUS_CONFLICT = http.StatusConflict

func NewConflictException(messages []string) *ConflictException {
	return &ConflictException{
		code:     STATUS_CONFLICT,
		messages: messages,
	}
}

func NewConflictExceptionSglMsg(message string) *ConflictException {
	return &ConflictException{
		code:     STATUS_CONFLICT,
		messages: []string{message},
	}
}

func (this ConflictException) GetCode() int {
	return this.code
}

func (this ConflictException) GetMessages() []string {
	return this.messages
}

func (this ConflictException) Error() string {
	var message string
	for _, value := range this.messages {
		message = message + value
	}
	return message
}
