package routes

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

type ControllerList struct {
	UserController        controllers.UserController
	PhotoController       controllers.PhotoController
	CommentController     controllers.CommentController
	SocialMediaController controllers.SocialMediaController
}

func (c ControllerList) InitRoute(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", c.UserController.Register)
		userRoutes.POST("/login", c.UserController.Login)
		userRoutes.Use(middlewares.Authentication())
		userRoutes.PUT("", c.UserController.Update)
		userRoutes.DELETE("", c.UserController.Delete)
	}

	photoRoutes := r.Group("/photos")
	photoRoutes.Use(middlewares.Authentication())
	{
		photoRoutes.GET("", c.PhotoController.GetAll)
		photoRoutes.POST("", c.PhotoController.Create)
		photoRoutes.Use(middlewares.PhotoAuthorization())
		photoRoutes.PUT("/:photoId", c.PhotoController.Update)
		photoRoutes.DELETE("/:photoId", c.PhotoController.Delete)
	}

	commentRoute := r.Group("/comments")
	commentRoute.Use(middlewares.Authentication())
	{
		commentRoute.GET("", c.CommentController.GetAll)
		commentRoute.POST("", c.CommentController.Create)
		commentRoute.Use(middlewares.CommentAuthorization())
		commentRoute.PUT("/:commentId", c.CommentController.Update)
		commentRoute.DELETE("/:commentId", c.CommentController.Delete)
	}

	socialMediaRoute := r.Group("/socialmedias")
	socialMediaRoute.Use(middlewares.Authentication())
	{
		socialMediaRoute.GET("", c.SocialMediaController.GetAll)
		socialMediaRoute.POST("", c.SocialMediaController.Create)
		socialMediaRoute.Use(middlewares.SocialMediaAuthorization())
		socialMediaRoute.PUT("/:socialMediaId", c.SocialMediaController.Update)
		socialMediaRoute.DELETE("/:socialMediaId", c.SocialMediaController.Delete)
	}
}
