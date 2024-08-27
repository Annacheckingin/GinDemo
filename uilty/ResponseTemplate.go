package uilty

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
