package main

import (
	"amirhossein-shakeri/gin-boilerplate/auth"
	"amirhossein-shakeri/gin-boilerplate/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var PORT = os.Getenv("PORT")
var ADDRESS = os.Getenv("DOMAIN") + ":" + PORT
var APP_NAME = os.Getenv("APP_NAME")

func main() {
	log.Printf("🚀 Starting %v Server ...\n", APP_NAME)
	if PORT == "" || ADDRESS == ":" {
		log.Panicln("Please set environment variables ...")
	}
	db.InitMGM()
	gin.ForceConsoleColor()
	mainRouter := gin.Default()
	setupRoutes(mainRouter)
	mainRouter.Run(ADDRESS)
}

func setupRoutes(router *gin.Engine) {
	router.Any("/health", healthHandler)
	authRouter := router.Group("/auth")
	{
		authRouter.GET("/", auth.AuthorizeJWT(), auth.GetInfo) // get session info
		authRouter.POST("/", auth.Login)                       // login
		authRouter.POST("/signup", auth.Signup)                // signup
	}
}

func healthHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "I'm Alive!",
	})
}
