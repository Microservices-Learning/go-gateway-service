package model

import (
	"go-gateway-service/client/gen-proto/common"
)

type ErrorResponse struct {
	ErrorCode        string            `json:"error_code"`
	ErrorDescription string            `json:"error_description"`
	Errors           map[string]string `json:"errors"`
	Exception        string            `json:"exception"`
}

func AsErrorResponse(response *common.Error) ErrorResponse {
	return ErrorResponse{
		ErrorCode:        response.Code.String(),
		ErrorDescription: response.Message,
		Exception:        response.Exception,
		Errors:           response.Details,
	}
}

/* Swagger Docs */

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PageResponse struct {
	Page          int32       `json:"page"`
	Size          int32       `json:"size"`
	TotalElements int64       `json:"total_elements"`
	TotalPages    int32       `json:"total_pages"`
	Items         interface{} `json:"items"`
}
