package services

import "github.com/gin-gonic/gin"

type responseStruct struct {
	Code resCode     `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func response(c *gin.Context, statusCode int, code resCode, msg string, data interface{}) {
	c.JSON(statusCode, responseStruct{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func responseSuccess(c *gin.Context, data interface{}) {
	response(c, codeSuccess.StatusCode(), codeSuccess, "", data)
}

func responseError(c *gin.Context, code resCode) {
	response(c, code.StatusCode(), code, code.Msg(), nil)
}

// func responseErrorWithMsg(c *gin.Context, code resCode, msg string) {
// 	response(c, code.StatusCode(), code, msg, nil)
// }

// func responseErrorWithData(c *gin.Context, code resCode, data interface{}) {
// 	response(c, code.StatusCode(), code, code.Msg(), data)
// }
