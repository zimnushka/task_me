package models

import "github.com/gin-gonic/gin"

const HeaderAuth = "Authorization"

type Controller interface {
	Init(router *gin.Engine)
}
