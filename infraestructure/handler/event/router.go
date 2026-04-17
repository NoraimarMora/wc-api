package event

import (
	"github.com/gin-gonic/gin"

	"wc-api/model"
)

func NewRouter(specification model.RouterSpecification) {
	publicRoutes(specification.Api, newHandler(specification.Dependencies.EventUseCase))
}

func publicRoutes(api *gin.RouterGroup, h handler) {
	routes := api.Group("/v1/events")
	routes.GET("", h.getAll)
}
