package apps

import "github.com/gin-gonic/gin"

type AppController interface {
	GetHandler() gin.HandlerFunc
}

type Response struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"error"`
	Data     interface{} `json:"data"`
}

func ErrorResponse(err error) *Response {
	return &Response{Code: -1, ErrorMsg: err.Error()}
}

func SuccessResponse(data interface{}) *Response {
	return &Response{Code: 0, ErrorMsg: "ok", Data: data}
}
