package main

import (
	infrastructureA "proyecto/src/accessories/infraestructure"
	infrastructureP "proyecto/src/pets/infrastructure"

	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {

	//configurar cors
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router := gin.Default()

	router.Use(cors.New(config))

	dbP := infrastructureP.NewMySQL()
	dbA := infrastructureA.NewMySQL()
	infrastructureP.InitPets(dbP, router)
	infrastructureA.InitAccessories(dbA, router)
	router.Run(":8080")
}
