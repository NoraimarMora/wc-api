package team

import (
	"github.com/gin-gonic/gin"

	"wc-api/model"
)

func NewRouter(specification model.RouterSpecification) {
	publicRoutes(specification.Api, newHandler(specification.Dependencies.TeamUseCase))
}

func publicRoutes(api *gin.RouterGroup, h handler) {
	routes := api.Group("/v1/teams")
	routes.GET("", h.getTeams)
	routes.GET("/:id", h.getTeam)
}
