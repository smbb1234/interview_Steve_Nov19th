package utils

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func ResponseSuccess(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func ResponseError(err string) Response {
	return Response{
		Success: false,
		Error:   err,
	}
}
