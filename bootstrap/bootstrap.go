package bootstrap

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"wc-api/infraestructure/handler"
	"wc-api/infraestructure/repository/local"
	"wc-api/model"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	appName := os.Getenv("APP_NAME")
	dependencies := NewDependencies(ctx, appName)
	localProvider := local.New()

	appDependencies := dependencies.Build(ctx, localProvider)

	api := gin.New()
	api.Use(gin.Recovery())
	api.SetHTMLTemplate(template.Must(template.ParseFiles("./front/html/index.html")))
	apiGroup := api.Group(appName + "/api")
	handler.InitRoutes(model.RouterSpecification{
		Api:          apiGroup,
		Dependencies: appDependencies,
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        api,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}
