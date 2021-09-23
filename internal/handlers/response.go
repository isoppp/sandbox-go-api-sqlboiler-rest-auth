package handlers

type successResponse struct {
	Data interface{} `json:"data"`
}

func JsonSuccessResponse(data interface{}) interface{} {
	return successResponse{
		Data: data,
	}
}
