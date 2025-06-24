package main

import (
	"backend-login/config"
	"backend-login/db"
	"backend-login/handlers"
	"backend-login/services"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//Se obtienen las variables de entorno del .env y se inicializa la base de datos y Firebase
	config.LoadEnv()
	db.Connect()
	services.InitFirebase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://login.tssw.cl"}, // Cambiar en prod
		AllowMethods:     []string{"POST", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Endpoint para verificar el token de autenticación, redirige a la función VerifyToken del handler
	r.POST("/auth/verify", handlers.VerifyToken)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
