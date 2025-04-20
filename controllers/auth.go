package controllers

import (
	"nbapp/config"
	"nbapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos."})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos."})
		return
	}

	token, err := GenerateKey(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o acess_token (JWT)"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login realiado com sucesso!", "acess_token": token})
}
