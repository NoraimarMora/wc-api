package ranking

import (
	"github.com/gin-gonic/gin"

	"wc-api/model"
)

func NewRouter(specification model.RouterSpecification) {
	publicRoutes(specification.Api, newHandler(specification.Dependencies.RankingUseCase))
}

func publicRoutes(api *gin.RouterGroup, h handler) {
	routes := api.Group("/v1/ranking")
	routes.GET("", h.getWorldRanking)
}
