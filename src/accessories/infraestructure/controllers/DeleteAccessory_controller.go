package controllers

import (
	"net/http"
	"proyecto/src/accessories/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteAccessoryController struct {
	remover *application.DeleteAccessory
}

func NewDeleteAccessoryController(useCase *application.DeleteAccessory) *DeleteAccessoryController {
	return &DeleteAccessoryController{remover: useCase}
}

func (dp *DeleteAccessoryController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Accessory ID"})
		return
	}
	err = dp.remover.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Accesorio eliminado correctamente"})
}
