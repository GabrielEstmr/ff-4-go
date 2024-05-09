/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

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
