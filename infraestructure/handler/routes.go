package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"wc-api/infraestructure/handler/ball"
	"wc-api/infraestructure/handler/city"
	"wc-api/infraestructure/handler/event"
	"wc-api/infraestructure/handler/mascot"
	"wc-api/infraestructure/handler/match"
	"wc-api/infraestructure/handler/news"
	"wc-api/infraestructure/handler/ranking"
	"wc-api/infraestructure/handler/record"
	"wc-api/infraestructure/handler/sound"
	"wc-api/infraestructure/handler/standings"
	"wc-api/infraestructure/handler/team"
	"wc-api/model"
)

func InitRoutes(specification model.RouterSpecification) {
	frontRoutes(specification.Api)
	healthRoutes(specification.Api)

	ball.NewRouter(specification)
	city.NewRouter(specification)
	event.NewRouter(specification)
	mascot.NewRouter(specification)
	match.NewRouter(specification)
	news.NewRouter(specification)
	ranking.NewRouter(specification)
	record.NewRouter(specification)
	sound.NewRouter(specification)
	standings.NewRouter(specification)
	team.NewRouter(specification)
}

func frontRoutes(api *gin.RouterGroup) {
	routes := api.Group("")

	routes.StaticFile("/styles.css", "./front/assets/styles.css")
	routes.StaticFile("/functions.js", "./front/assets/functions.js")

	routes.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "World Cup Qualifiers API",
		})
	})
}

func healthRoutes(api *gin.RouterGroup) {
	routes := api.Group("")

	routes.GET("health", health)
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
