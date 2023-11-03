package main

import (
	"net/http"

	"github.com/Enhancifire/Petrolium/backend/models"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {

}

func GetPetrolList(c *gin.Context) {
	petrolList := []models.PetrolModel{
		*models.NewPetrolModel("1", 1, 1, 1),
		*models.NewPetrolModel("2", 1, 3, 1),
		*models.NewPetrolModel("3", 2, 1, 1),
		*models.NewPetrolModel("3", 1, 1, 5),
	}
	c.JSON(http.StatusOK, petrolList)
}
