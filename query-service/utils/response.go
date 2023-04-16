package utils

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(message string, data interface{}) response {
	return response{
		Message: message,
		Data:    data,
	}
}
