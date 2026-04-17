package match

import (
	"github.com/gin-gonic/gin"

	"wc-api/model"
)

func NewRouter(specification model.RouterSpecification) {
	publicRoutes(specification.Api, newHandler(specification.Dependencies.MatchUseCase))
}

func publicRoutes(api *gin.RouterGroup, h handler) {
	routes := api.Group("/v1/matches")
	routes.GET("", h.getMatches)
	routes.GET("/:id", h.getByID)
}
