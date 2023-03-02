package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"task-5-vix-btpns-SuburFirmansyah/helpers"
	"task-5-vix-btpns-SuburFirmansyah/models"
)

type RegisInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateInput struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (s *Server) CreateUser(c *gin.Context) {
	var input RegisInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input does not match the criteria"})
		return
	}

	user := models.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	if err := s.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (s *Server) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input does not match the criteria"})
		return
	}

	user := models.User{}
	dbRes := s.DB.Where("email = ?", input.Email).First(&user)
	if err := dbRes.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password is incorrect"})
		return
	}

	tokenString, err := helpers.GenerateJwt(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (s *Server) UpdateUser(c *gin.Context) {
	var input UpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input does not match the criteria"})
		return
	}
	userId, err := helpers.GetUserIdFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIdFromUrl := c.Param("userId")

	if userIdFromUrl != fmt.Sprint(userId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user := models.User{}
	_, err = user.GetUser(s.DB, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Name = input.Name
	user.Email = input.Email

	if _, err = user.UpdateUser(s.DB, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (s *Server) DeleteUser(c *gin.Context) {
	userId, err := helpers.GetUserIdFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIdFromUrl := c.Param("userId")

	if userIdFromUrl != fmt.Sprint(userId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user := models.User{}

	_, err = user.DeleteUser(s.DB, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
