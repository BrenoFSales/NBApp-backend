package controllers

import (
	"encoding/base64"
	"nbapp/config"
	"nbapp/models"
	"net/http"
	"strings"

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

// Essa função apenas retorna a logo do time como o ID no parâmetro
func GetLogo(c *gin.Context) {
	var team models.Team
	if err := config.DB.First(&team, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time não encontrado"})
		return
	}

	if team.LogoBase64 == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Logo não disponível."})
		return
	}

	// Tratamento para pegar a imagem sem o prefixo base64 por conta do método padrão de Decode do GO
	logoBase64 := team.LogoBase64
	if strings.HasPrefix(logoBase64, "data:image") {
		parts := strings.Split(logoBase64, ",")
		if len(parts) == 2 {
			logoBase64 = parts[1] // Pega só a parte codificada
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Logo Base64 inválida."})
			return
		}
	}

	logoBytes, err := base64.StdEncoding.DecodeString(logoBase64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar a logo."})
		return
	}

	c.Data(http.StatusOK, "image/png", logoBytes)
}
