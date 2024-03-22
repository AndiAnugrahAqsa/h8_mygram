package main

import (
	"mygram/config"
	"mygram/controllers"
	"mygram/database"
	"mygram/repositories"
	"mygram/routes"
	"mygram/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	database.InitDB()

	r := gin.Default()

	db := database.InitDB()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	photoRepository := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepository)
	photoController := controllers.NewPhotoController(photoService)

	commentRepository := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepository)
	commentController := controllers.NewCommentController(commentService)

	socialMediaRepository := repositories.NewSocialMediaRepository(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepository)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	controllerList := routes.ControllerList{
		UserController:        userController,
		PhotoController:       photoController,
		CommentController:     commentController,
		SocialMediaController: socialMediaController,
	}

	controllerList.InitRoute(r)

	r.Run(":" + config.Cfg.PORT) // listen and serve on 0.0.0.0:8080
}
