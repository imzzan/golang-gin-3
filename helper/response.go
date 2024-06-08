package helper

import "golang-gin3/dto"

type ResponseWithData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(params dto.ResponsePrams) any {
	var response any
	var status string

	if params.StatusCode >= 200 && params.StatusCode < 300 {
		status = "OK"
	} else {
		status = "Faild"
	}

	if params.Data != nil {
		response = &ResponseWithData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
			Data:    params.Data,
		}
	} else {
		response = &ResponseWithoutData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
		}
	}

	return response
}
