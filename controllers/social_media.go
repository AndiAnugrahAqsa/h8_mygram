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

type SocialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(socialMediaService services.SocialMediaService) SocialMediaController {
	return SocialMediaController{
		socialMediaService,
	}
}

func (cc *SocialMediaController) GetAll(c *gin.Context) {

	socialMedias, err := cc.socialMediaService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if len(*socialMedias) == 0 {
		c.JSON(http.StatusNoContent, &socialMedias)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"social_medias": &socialMedias,
	})
}

func (cc *SocialMediaController) Create(c *gin.Context) {
	var socialMediaRequest dto.SocialMediaRequest

	if err := c.Bind(&socialMediaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := socialMediaRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)

	userId := int(userData["id"].(float64))

	socialMediaRequest.UserID = userId

	socialMedia, err := cc.socialMediaService.Create(socialMediaRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, socialMedia)
}

func (cc *SocialMediaController) Update(c *gin.Context) {
	var socialMediaRequest dto.SocialMediaRequest

	if err := c.Bind(&socialMediaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := socialMediaRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	socialMediaIdString := c.Param("socialMediaId")

	socialMediaId, err := strconv.Atoi(socialMediaIdString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	socialMedia, err := cc.socialMediaService.Update(socialMediaId, socialMediaRequest)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "social media not found or has been deleted",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, socialMedia)
}

func (cc *SocialMediaController) Delete(c *gin.Context) {
	socialMediaIdString := c.Param("socialMediaId")

	socialMediaId, err := strconv.Atoi(socialMediaIdString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := cc.socialMediaService.Delete(socialMediaId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "your social media not found or has been deleted",
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
		"message": "your social media has been successfully deleted",
	})
}
