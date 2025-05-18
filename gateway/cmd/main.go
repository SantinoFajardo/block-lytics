package main

import (
	"log"
	"os"

	"github.com/block-lytics/gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func main() {
	app := &App{
		router: gin.Default(),
	}

	app.router.POST("/accounts", handlers.CreateAccountHandler)
	app.router.GET("/accounts/:address", handlers.GetAccountByAdressHandler)

	app.router.POST("/users", handlers.CreateUserHandler)

	app.router.PUT("/balances", handlers.UpdateBalancesHandler)

	port := os.Getenv("PORT")
	if err := app.router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
