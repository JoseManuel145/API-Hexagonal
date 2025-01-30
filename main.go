package main

import (
	infrastructureA "proyecto/src/accessories/infraestructure"
	infrastructureP "proyecto/src/pets/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	dbP := infrastructureP.NewMySQL()
	dbA := infrastructureA.NewMySQL()
	router := gin.Default()
	infrastructureP.InitPets(dbP, router)
	infrastructureA.InitAccessories(dbA, router)
	router.Run(":8080")
}
