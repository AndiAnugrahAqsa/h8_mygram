package middlewares

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", verifyToken)

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		db := database.InitDB()

		photoId := c.Param("photoId")

		var photo models.Photo

		if err := db.Select("user_id").First(&photo, "id", photoId).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "data doesn't exist",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(float64))

		if photo.UserID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		db := database.InitDB()

		commentId := c.Param("commentId")

		var comment models.Comment

		if err := db.Select("user_id").First(&comment, "id", commentId).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "data doesn't exist",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(float64))

		if comment.UserID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		db := database.InitDB()

		socialMediaId := c.Param("socialMediaId")

		var socialMedia models.SocialMedia

		if err := db.Select("user_id").First(&socialMedia, "id", socialMediaId).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "data doesn't exist",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(float64))

		if socialMedia.UserID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
