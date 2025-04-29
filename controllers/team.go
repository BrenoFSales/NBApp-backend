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

// Struct para upload da logo do time
type UploadLogoBase64 struct {
	LogoBase64 string `json:"logoBase64" binding:"required"`
}

// Atualizar a logo do time
func UploadLogoTeam(c *gin.Context) {
	var team models.Team
	if err := config.DB.First(&team, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time não encontrado."})
		return
	}

	var input UploadLogoBase64
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if input.LogoBase64 != "" {
		team.LogoBase64 = input.LogoBase64
	}

	config.DB.Save(&team)
	c.JSON(http.StatusOK, gin.H{"message": "Logo do time foi importada com sucesso."})
}
