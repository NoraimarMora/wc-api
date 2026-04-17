package model

import "github.com/gin-gonic/gin"

type RouterSpecification struct {
	Api          *gin.RouterGroup
	Dependencies Dependencies
}
