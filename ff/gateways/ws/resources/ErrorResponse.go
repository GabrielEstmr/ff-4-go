/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_gateways_ws_resources

type ErrorResponse struct {
	Code     string   `json:"code"`
	Messages []string `json:"message"`
}

func NewErrorResponse(code string, messages []string) *ErrorResponse {
	return &ErrorResponse{
		Code:     code,
		Messages: messages,
	}
}

func NewErrorResponseSlgMsg(code string, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:     code,
		Messages: []string{message},
	}
}
