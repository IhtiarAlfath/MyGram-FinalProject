package handler

import (
	"kominfo/h8-myGram-finalproject/database"
	"kominfo/h8-myGram-finalproject/docs"
	"kominfo/h8-myGram-finalproject/repository/comment_repository/comment_pg"
	"kominfo/h8-myGram-finalproject/repository/photo_repository/photo_pg"
	"kominfo/h8-myGram-finalproject/repository/socialmedia_repository/socialmedia_pg"
	"kominfo/h8-myGram-finalproject/repository/user_repository/user_pg"
	"kominfo/h8-myGram-finalproject/service"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	const PORT = "3030"

	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	userRepo := user_pg.NewUserPG(db)

	photoRepo := photo_pg.NewPhotoPG(db)

	socialMediaRepo := socialmedia_pg.NewSocialMediaPG(db)

	commentRepo := comment_pg.NewCommentPG(db)

	userService := service.NewUserService(userRepo)

	photoService := service.NewPhotoService(photoRepo)

	socialMediaService := service.NewSocialMediaService(socialMediaRepo)

	commentService := service.NewCommentService(commentRepo)

	userHandler := NewUserHandler(userService)

	photoHandler := NewPhotoHandler(photoService)

	socialMediaHandler := NewSocialMediaHandler(socialMediaService)

	commentHandler := NewCommentHandler(commentService)

	authService := service.NewAuthService(userRepo, photoRepo, socialMediaRepo, commentRepo)

	route := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "MyGram API"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3030"


	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRoutes := route.Group("/users")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
	}

	photoRoutes := route.Group("/photo")
	{
		photoRoutes.Use(authService.Authentication())

		photoRoutes.POST("/", photoHandler.CreatePhoto)
		photoRoutes.GET("/", photoHandler.GetPhotos)
		photoRoutes.GET("/:photoId", photoHandler.GetPhotoByID)
		photoRoutes.PUT("/:photoId", authService.PhotoAuthorization(), photoHandler.UpdatePhotoById)
		photoRoutes.DELETE("/:photoId", authService.PhotoAuthorization(), photoHandler.DeletePhotoById)
	}

	commentRoutes := route.Group("/photo/:photoId/comment")
	{
		commentRoutes.Use(authService.Authentication())

		commentRoutes.POST("/", commentHandler.CreateComment)
		commentRoutes.GET("/", commentHandler.GetComments)
		commentRoutes.GET("/:commentId", commentHandler.GetCommentById)
		commentRoutes.PUT("/:commentId", authService.CommentAuthorization(), commentHandler.UpdateCommentById)
		commentRoutes.DELETE("/:commentId", authService.CommentAuthorization(), commentHandler.DeleteCommentById)
	}

	socialMediaRoutes := route.Group("/socialmedia")
	{
		socialMediaRoutes.Use(authService.Authentication())

		socialMediaRoutes.POST("/", socialMediaHandler.CreateSocialMedia)
		socialMediaRoutes.GET("/", socialMediaHandler.GetSocialMedias)
		socialMediaRoutes.GET("/:socialMediaId", socialMediaHandler.GetSocialMediabyId)
		socialMediaRoutes.PUT("/:socialMediaId", authService.SocialMediaAuthorization(), socialMediaHandler.UpdateSocialMediabyId)
		socialMediaRoutes.DELETE("/:socialMediaId", authService.SocialMediaAuthorization(), socialMediaHandler.DeleteSocialMediabyId)
	}


	route.Run(":", PORT)
}
