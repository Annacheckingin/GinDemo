package uilty

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response[T any] struct {
	Data    *T     `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func SuccessResponseArray[T any](t *[]T) Response[[]T] {
	return Response[[]T]{Code: 0, Message: "success", Data: t}
}

func SuccessResponse[T any](t *T) Response[T] {
	return Response[T]{Code: 0, Message: "success", Data: t}
}

func ErrorResponse[T any](t *T) Response[T] {
	return Response[T]{Code: -1, Message: "error", Data: t}
}

type ErrorResponseDefault struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func Error(c *gin.Context, er error) {
	ErrorMessage(c, er.Error())
}
func ErrorMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, ErrorResponseDefault{Code: -1, Message: message})
}

func Done(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse[string](nil))
}

func DoneWithReturn[T any](c *gin.Context, object T) {
	c.JSON(http.StatusOK, SuccessResponse(&object))
}
func DoneWithReturnArray[T any](c *gin.Context, object []T) {
	c.JSON(http.StatusOK, SuccessResponseArray(&object))
}
