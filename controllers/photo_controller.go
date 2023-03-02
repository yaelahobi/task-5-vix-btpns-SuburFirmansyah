package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"task-5-vix-btpns-SuburFirmansyah/helpers"
	"task-5-vix-btpns-SuburFirmansyah/models"
)

type InputPhoto struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoUrl string `json:"photo_url" validate:"required,url"`
}

type GetPhotoResp struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	Name     string `json:"name"`
}

func (s *Server) CreatePhoto(c *gin.Context) {
	var input InputPhoto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := helpers.GetUserIdFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	photo := models.Photo{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoUrl: input.PhotoUrl,
		UserID:   userId,
	}

	if err := s.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (s *Server) GetPhotos(c *gin.Context) {
	var photos []models.Photo
	if err := s.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// This method is unoptimized, gonna find another way

	var respStruct []GetPhotoResp
	for _, photo := range photos {
		var user models.User
		if err := s.DB.First(&user, photo.UserID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		photoRespStruct := GetPhotoResp{}

		photoRespStruct.ID = photo.ID
		photoRespStruct.Title = photo.Title
		photoRespStruct.Caption = photo.Caption
		photoRespStruct.PhotoUrl = photo.PhotoUrl
		photoRespStruct.Name = user.Name

		respStruct = append(respStruct, photoRespStruct)
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": respStruct})
}

func (s *Server) UpdatePhoto(c *gin.Context) {
	var input InputPhoto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photoId := c.Param("photoId")

	photo := models.Photo{}

	if err := s.DB.First(&photo, photoId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId, err := helpers.GetUserIdFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userId != photo.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you don't have permission to update this photo"})
		return
	}

	photo.PhotoUrl = input.PhotoUrl
	photo.Title = input.Title
	photo.Caption = input.Caption

	if err := s.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (s *Server) DeletePhoto(c *gin.Context) {
	photoId := c.Param("photoId")

	var photo models.Photo
	if err := s.DB.First(&photo, photoId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId, err := helpers.GetUserIdFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userId != photo.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you don't have permission to delete this photo"})
		return
	}

	if err := s.DB.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
