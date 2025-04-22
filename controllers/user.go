package controllers

import (
	"nbapp/config"
	"nbapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Struct para criar usuário
type UserCreateInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Struct para atualizar usuário
type UserUpdateInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Criar Usuário
func CreateUser(c *gin.Context) {
	var input UserCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados Inválidos."})
		return
	}

	// Faz o hash da senha no banco de dados
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{Username: input.Username, Password: string(hashedPassword)}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível criar usuário."})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso."})
}

// Listar todos os usuários
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}

// Buscar usuário por ID
func GetUserByID(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrato."})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Atualizar usuário
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado."})
		return
	}

	var input UserUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if input.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
	}

	if input.Username != "" {
		user.Username = input.Username
	}

	config.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso."})
}

// Deletar usuário
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado."})
		return
	}

	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso."})
}
