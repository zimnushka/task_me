package app_errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func New(statusCode int, message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func FromError(err error) *AppError {
	if err != nil {
		return nil
	}
	return &AppError{
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}

func (appError AppError) Call(context *gin.Context) {
	err := errors.New(appError.Message)
	context.Error(err)
	context.AbortWithStatusJSON(appError.StatusCode, err.Error())
}

const ERR_Forbiden = "forbiden"
const ERR_Wrong_auth = "wrong auth"
const ERR_Empty_field = "the field is empty"
const ERR_Not_found = "not found"
const ERR_Unexpected_repository_error = "repository unexpected error"
const ERR_User_already_register = "user already registered"
