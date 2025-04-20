package controllers

import (
	"nbapp/config"
	"nbapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	// input.Username = strings.TrimSpace(input.Username) // Isso seria bom para evitar espaços no começo e no final dos inputs, [import "strings"]
	// input.Password = strings.TrimSpace(input.Password)

	var user models.User

	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha incorretos."})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha incorretos."})
		return
	}

	token, err := GenerateKey(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o access_token (JWT)"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login realiado com sucesso!", "access_token": token})
}
