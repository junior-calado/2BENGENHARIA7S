package controller

import (
	"2BENGENHARIA7S/model"
	"2BENGENHARIA7S/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCards(c *gin.Context) {
	cards, err := service.GetCards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cards)
}

func GetCardByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	card, err := service.GetCardByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}
	c.JSON(http.StatusOK, card)
}

func CreateCard(c *gin.Context) {
	var newCard model.Card
	if err := c.ShouldBindJSON(&newCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCard, err := service.AddCard(newCard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdCard)
}

func UpdateCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var card model.Card
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card.ID = id
	updatedCard, err := service.UpdateCard(card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCard)
}

func DeleteCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := service.DeleteCard(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Card deleted successfully"})
} 