package main

import (
	"fmt"
	"log"

	"ciam_exam/src/config"
	"ciam_exam/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"

	docs "ciam_exam/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger

// @contact.name nathan.lu
// @contact.url https://tedmax100.github.io/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8008
// schemes http
func main() {
	config.InitOs()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(favicon.New("./favicon.ico"))

	if mode := gin.Mode(); mode == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.LoadHTMLGlob("views/*")

	routes.UseAuth(r)

	if err := r.Run(fmt.Sprintf(":%s", config.GetEnv().SERVER_PORT)); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
