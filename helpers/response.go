package helpers

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseApi(msg string, data interface{}) ApiResponse {
	result := ApiResponse{
		Message: msg,
		Data:    data,
	}

	return result
}
