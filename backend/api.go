package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {

}

func GetPetrolList(c *gin.Context) {
	petrolList := []PetrolModel{
		*NewPetrolModel("1", 1, 1, 1),
		*NewPetrolModel("2", 1, 3, 1),
		*NewPetrolModel("3", 2, 1, 1),
		*NewPetrolModel("3", 1, 1, 5),
	}
	c.JSON(http.StatusOK, petrolList)
}
