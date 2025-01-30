package controllers

import (
	"net/http"
	"proyecto/src/pets/application"

	"github.com/gin-gonic/gin"
)

type CreatePetController struct {
	petSaver *application.SavePet
}

func NewSavePetController(useCase *application.SavePet) *CreatePetController {
	return &CreatePetController{petSaver: useCase}
}

func (cp *CreatePetController) Run(c *gin.Context) {
	var json struct {
		Name string `json:"name"`
		Raza string `json:"raza"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cp.petSaver.Execute(json.Name, json.Raza)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pet saved successfully"})
}
