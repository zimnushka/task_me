package models

type AppError interface {
	call(AppError)
}
