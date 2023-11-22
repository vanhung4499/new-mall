package ctl

import (
	"github.com/gin-gonic/gin"
	"new-mall/pkg/e"
)

// Response is the base serializer
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// ResSuccess returns a success response with data
func ResSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	if data == nil {
		data = "Operation successful"
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}

	return r
}

// ResError returns an error response
func ResError(ctx *gin.Context, err error, data string, code ...int) *Response {
	status := e.ERROR
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Msg:    e.GetMsg(status),
		Data:   data,
		Error:  err.Error(),
	}

	return r
}
