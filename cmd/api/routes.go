package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "swagger-intro/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) routes() http.Handler {
	// router := httprouter.New()

	// router.NotFound = http.HandlerFunc(app.notFoundResponse)
	// router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	router := gin.Default()

	api := router.Group("/v1")

	{
		api.GET("/healthcheck", app.healthcheckHandler)
		api.POST("/movies", app.createMovieHandler)
		api.GET("/movies/:id", app.showMovieHandler)

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
