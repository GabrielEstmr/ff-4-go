/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

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
