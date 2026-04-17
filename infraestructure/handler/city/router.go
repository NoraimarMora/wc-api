package city

import (
	"github.com/gin-gonic/gin"

	"wc-api/model"
)

func NewRouter(specification model.RouterSpecification) {
	publicRoutes(specification.Api, newHandler(specification.Dependencies.CityUseCase))
}

func publicRoutes(api *gin.RouterGroup, h handler) {
	routes := api.Group("/v1/cities")
	routes.GET("", h.getCities)
	routes.GET("/:id", h.getCity)
}
