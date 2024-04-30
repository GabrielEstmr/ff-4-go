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
