package handler

import (
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"Pizzas": data.Pizzas,
	})
}

func PostPizzas(c *gin.Context) {
	var newPizza models.Pizza

	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var maxID int

	for _, pizza := range data.Pizzas {
		if pizza.ID > maxID {
			maxID = pizza.ID
		}
	}

	newPizza.ID = maxID + 1

	data.Pizzas = append(data.Pizzas, newPizza)

	c.JSON(201, newPizza)
	data.SavePizza()
}

func GetPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for _, p := range data.Pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Pizza not found",
	})
}

func DeletePizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizza()
			c.JSON(200, gin.H{
				"message": "Pizza deleted successfully.",
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Pizza not found.",
	})
}

func UpdatePizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var updatedPizza models.Pizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas[i] = updatedPizza
			data.Pizzas[i].ID = id
			data.SavePizza()
			c.JSON(200, data.Pizzas[i])
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Pizza not found.",
	})
}
