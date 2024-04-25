package ff_mongo_exceptions

import "net/http"

type LibInternalServerErrorException struct {
	code     int
	messages []string
}

const STATUS_INTERNAL_SERVER_ERROR = http.StatusInternalServerError

func NewLibInternalServerErrorException(messages []string) *LibInternalServerErrorException {
	return &LibInternalServerErrorException{
		code:     STATUS_INTERNAL_SERVER_ERROR,
		messages: messages,
	}
}

func NewLibInternalServerErrorExceptionSglMsg(message string) *LibInternalServerErrorException {
	return &LibInternalServerErrorException{
		code:     STATUS_INTERNAL_SERVER_ERROR,
		messages: []string{message},
	}
}

func (this LibInternalServerErrorException) GetCode() int {
	return this.code
}

func (this LibInternalServerErrorException) GetMessages() []string {
	return this.messages
}

func (this LibInternalServerErrorException) Error() string {
	var message string
	for _, value := range this.messages {
		message = message + value
	}
	return message
}
