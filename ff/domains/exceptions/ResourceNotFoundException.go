/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_domains_exceptions

import "net/http"

type ResourceNotFoundException struct {
	code     int
	messages []string
}

const STATUS_NOT_FOUND = http.StatusNotFound

func NewResourceNotFoundException(messages []string) *ResourceNotFoundException {
	return &ResourceNotFoundException{
		code:     STATUS_NOT_FOUND,
		messages: messages,
	}
}

func NewResourceNotFoundExceptionSglMsg(message string) *ResourceNotFoundException {
	return &ResourceNotFoundException{
		code:     STATUS_NOT_FOUND,
		messages: []string{message},
	}
}

func (this ResourceNotFoundException) GetCode() int {
	return this.code
}

func (this ResourceNotFoundException) GetMessages() []string {
	return this.messages
}

func (this ResourceNotFoundException) Error() string {
	var message string
	for _, value := range this.messages {
		message = message + value
	}
	return message
}
