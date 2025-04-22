package controllers

import (
	"nbapp/config"
	"nbapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeams(c *gin.Context) {
	var teams []models.Team
	if err := config.DB.Find(&teams).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Times não encontrados."})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func GetTeamsById(c *gin.Context) {
	var team models.Team
	if err := config.DB.First(&team, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time com o `ID` não encontrado."})
		return
	}

	c.JSON(http.StatusOK, team)
}
