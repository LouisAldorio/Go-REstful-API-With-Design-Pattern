package main

import (
	"log"
	"myapp/config"
	"myapp/middlewares"
	"myapp/migrations"
	"myapp/routers"
	"os"

	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectGorm()
	migrations.RunMigrate()
}


// @title           Swagger Example API
// @version         1.0
func main() {
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware(), sentrygin.New(sentrygin.Options{
		Repanic: true,
	}), middlewares.Panic)

	routers.ApiRoute(router)
	routers.DocsRoute(router)

	log.Println("Listen and serve at http://localhost:" + port)
	router.Run(":" + port)
}
