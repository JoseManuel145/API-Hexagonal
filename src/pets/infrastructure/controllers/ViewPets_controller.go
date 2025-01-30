package controllers

import (
	"net/http"
	"proyecto/src/pets/application"

	"github.com/gin-gonic/gin"
)

type ViewPetsController struct {
	vps *application.ViewPets
}

func NewViewPetsController(useCase *application.ViewPets) *ViewPetsController {
	return &ViewPetsController{vps: useCase}
}

func (vpc *ViewPetsController) Run(c *gin.Context) {
	pets, err := vpc.vps.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"pets": pets})
}
