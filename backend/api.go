package main

import (
	"log"
	"net/http"

	"github.com/Enhancifire/Petrolium/backend/modules/petrol"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Hello Petrolium"})

}

func GetPetrolList(c *gin.Context) {
	petrolList, err := petrol.GetPetrolList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, petrolList)
}

func GetPetrolInstance(c *gin.Context) {
	id := c.Param("id")
	petrol, err := petrol.GetPetrolIndividual(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, petrol)
}

func AddPetrolData(c *gin.Context) {
	var input petrol.CreatePetrolInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// var model petrol.PetrolModel
	model, err := petrol.CreatePetrolModel(&input)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &model})
}

func UpdatePetrolData(c *gin.Context) {
	id := c.Param("id")
	var input petrol.UpdatePetrolInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := petrol.UpdatePetrolModel(id, input)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &item})
}

func DeletePetrolInstance(c *gin.Context) {
	id := c.Param("id")

	if err := petrol.DeletePetrolModel(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
