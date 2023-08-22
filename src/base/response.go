package base

type BaseObjectResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateResponse(data interface{}, okMsg string, err error) *BaseObjectResponse {
	var code, message string

	if err != nil {
		code = "0"
		message = err.Error()
	} else {
		code = "1"
		if len(okMsg) == 0 {
			message = "OK"
		} else {
			message = okMsg
		}
	}

	return &BaseObjectResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func CreateResponseWithError(err error) *BaseObjectResponse {
	var empty interface{}
	return CreateResponse(empty, "", err)
}
