package controllers

import (
	"errors"
	"mygram/dto"
	"mygram/helpers"
	"mygram/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var userRequest dto.UserRegister

	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := userRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	hashingPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	userRequest.Password = string(hashingPassword)

	user, err := uc.userService.Create(userRequest)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "the email or username has been registered",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, &user)
}

func (uc *UserController) Login(c *gin.Context) {
	var userRequest dto.UserLogin

	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := userRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	user, err := uc.userService.GetOneByFilter("email", userRequest.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "the email has not been registered",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (uc *UserController) Update(c *gin.Context) {
	var userRequest dto.UserEdit

	userData := c.MustGet("userData").(jwt.MapClaims)

	userId := int(userData["id"].(float64))

	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	if err := userRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "your request invalid, please check your request",
		})
		return
	}

	user, err := uc.userService.Update(userId, userRequest)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "your account not found or has been deleted",
			})
			return
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "email or username has been used",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Delete(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	userId := int(userData["id"].(float64))

	err := uc.userService.Delete(userId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "your account not found or has been deleted",
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
		"message": "your account has been successfully deleted",
	})
}
