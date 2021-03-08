package controllers

import "github.com/gin-gonic/gin"

type Response struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
}

func HTTPRes(c *gin.Context, httpCode int, msg string) {
	c.JSON(httpCode, Response{
		Code: httpCode,
		Error:  msg,
	})
	return
}
