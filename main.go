package main

import (
	_ "github.com/JuanEsp14/go-spotify/docs"
	"curso-go/go-spotify/pkg"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)


// @title Swagger Spotify API
// @version 1.0
// @description Spotify API.

// @tag.name Spotify
// @tag.description Resolve calls to Spotify API

// @contact.name Juan Espinoza
// @contact.email juanmesp@hotmail.com

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	router := gin.Default()
	serviceSpotify := pkg.NewSpotifyAPIServiceImpl(new(http.Client))
	controller := pkg.NewSpotifyGinController(serviceSpotify)
	controller.SetupRouter(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8080")
}
