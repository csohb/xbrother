package api_context

import (
	"github.com/labstack/echo"
	"net/http"
)

type CommonError struct {
	Message string `json:"message"`
}

type CommonResponse struct {
	Code      int          `json:"code"`
	IsSuccess bool         `json:"isSuccess"`
	Error     *CommonError `json:"error,omitempty"`
	Data      interface{}  `json:"data"`
}

type RedirectResponse struct {
	Message     string `json:"message"`
	RedirectUrl string `json:"redirect_url"`
}

func (resp CommonResponse) Send(c echo.Context) error {
	return c.JSON(http.StatusOK, &resp)
}

func BadJSON(msg string) *CommonResponse {
	return &CommonResponse{
		Code:      http.StatusBadRequest,
		IsSuccess: false,
		Error:     &CommonError{Message: msg},
		Data:      nil,
	}
}

func SuccessJSON(data interface{}) *CommonResponse {
	return &CommonResponse{
		Code:      http.StatusOK,
		IsSuccess: true,
		Data:      data,
	}
}

func EmptyJSON() *CommonResponse {
	return &CommonResponse{
		Code:      http.StatusOK,
		IsSuccess: true,
		Data:      nil,
	}
}

func InternalServerErrJSON(msg string) *CommonResponse {
	return &CommonResponse{
		Code:      http.StatusInternalServerError,
		IsSuccess: false,
		Error:     &CommonError{Message: msg},
		Data:      nil,
	}
}

func RedirectJSON(msg, redirectURL string) *CommonResponse {
	return &CommonResponse{
		Code:      0,
		IsSuccess: false,
		Error:     nil,
		Data:      nil,
	}
}
