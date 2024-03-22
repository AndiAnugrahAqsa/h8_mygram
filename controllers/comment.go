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

type CommentController struct {
	commentService services.CommentService
}

func NewCommentController(commentService services.CommentService) CommentController {
	return CommentController{
		commentService,
	}
}

func (cc *CommentController) GetAll(c *gin.Context) {

	comments, err := cc.commentService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if len(*comments) == 0 {
		c.JSON(http.StatusNoContent, &comments)
		return
	}

	c.JSON(http.StatusOK, &comments)
}

func (cc *CommentController) Create(c *gin.Context) {
	var commentRequest dto.CommentRequest

	if err := c.Bind(&commentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := commentRequest.Validate(); err != nil || commentRequest.PhotoID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)

	userId := int(userData["id"].(float64))

	commentRequest.UserID = userId

	comment, err := cc.commentService.Create(commentRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, &comment)
}

func (cc *CommentController) Update(c *gin.Context) {
	var commentRequest dto.CommentRequest

	if err := c.Bind(&commentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := commentRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	commentIdString := c.Param("commentId")

	commentId, err := strconv.Atoi(commentIdString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	comment, err := cc.commentService.Update(commentId, commentRequest)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "comment not found or has been deleted",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, comment)
}

func (cc *CommentController) Delete(c *gin.Context) {
	commentIdString := c.Param("commentId")

	commentId, err := strconv.Atoi(commentIdString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := cc.commentService.Delete(commentId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "your comment not found or has been deleted",
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
		"message": "your comment has been successfully deleted",
	})
}
