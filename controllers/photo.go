package controllers

import (
	"errors"
	"mygram/dto"
	"mygram/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoController struct {
	photoService services.PhotoService
}

func NewPhotoController(photoService services.PhotoService) PhotoController {
	return PhotoController{
		photoService,
	}
}

func (pc *PhotoController) GetAll(c *gin.Context) {

	photos, err := pc.photoService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if len(*photos) == 0 {
		c.JSON(http.StatusNoContent, &photos)
		return
	}

	c.JSON(http.StatusOK, &photos)
}

func (pc *PhotoController) Create(c *gin.Context) {
	var photoRequest dto.PhotoRequest

	if err := c.Bind(&photoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := photoRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)

	userId := int(userData["id"].(float64))

	photoRequest.UserID = userId

	photo, err := pc.photoService.Create(photoRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, &photo)
}

func (pc *PhotoController) Update(c *gin.Context) {
	var photoRequest dto.PhotoRequest

	if err := c.Bind(&photoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := photoRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	photoIdString := c.Param("photoId")

	photoId, err := strconv.Atoi(photoIdString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	photo, err := pc.photoService.Update(photoId, photoRequest)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "photo not found or has been deleted",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, photo)
}

func (pc *PhotoController) Delete(c *gin.Context) {
	photoIdString := c.Param("photoId")

	photoId, err := strconv.Atoi(photoIdString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := pc.photoService.Delete(photoId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "your photo not found or has been deleted",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your photo has been successfully deleted",
	})
}
