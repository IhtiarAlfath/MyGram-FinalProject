package service

import (
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/pkg/helpers"
	"kominfo/h8-myGram-finalproject/repository/comment_repository"
	"kominfo/h8-myGram-finalproject/repository/photo_repository"
	"kominfo/h8-myGram-finalproject/repository/socialmedia_repository"
	"kominfo/h8-myGram-finalproject/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	PhotoAuthorization() gin.HandlerFunc
	SocialMediaAuthorization() gin.HandlerFunc
	CommentAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo        user_repository.UserRepository
	photoRepo       photo_repository.PhotoRepository
	socialMediaRepo socialmedia_repository.SocialMediaRepository
	commentRepo     comment_repository.CommentRepository
}

func NewAuthService(userRepo user_repository.UserRepository, photoRepo photo_repository.PhotoRepository, socialMediaRepo socialmedia_repository.SocialMediaRepository, commentRepo comment_repository.CommentRepository) AuthService {
	return &authService{
		userRepo:        userRepo,
		photoRepo:       photoRepo,
		socialMediaRepo: socialMediaRepo,
		commentRepo:     commentRepo,
	}
}

func (au *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invalidTokenErr := errs.NewUnauthorizedError("invalid token")

		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := au.userRepo.GetuserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), errs.NewNotFoundError("Usernotfound"))
			return
		}

		_ = result

		ctx.Set("userData", user)

		ctx.Next()
	}
}

func (au *authService) PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		photoId, err := helpers.GetParamID(ctx, "photoId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		photo, err := au.photoRepo.GetPhotoByID(photoId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if user.Id != photo.UserId {
			Unauthorized := errs.NewUnauthorizedError("you are not authorized")
			ctx.AbortWithStatusJSON(Unauthorized.Status(), Unauthorized)
			return
		}
		ctx.Next()
	}
}

func (au *authService) SocialMediaAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		socialMediaId, err := helpers.GetParamID(ctx, "socialMediaId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		socialMedia, err := au.socialMediaRepo.GetSocialMediabById(socialMediaId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if socialMedia.UserId != user.Id {
			Unauthorized := errs.NewUnauthorizedError("you are not authorized")
			ctx.AbortWithStatusJSON(Unauthorized.Status(), Unauthorized)
			return
		}
		ctx.Next()
	}
}

func (au *authService) CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		photoId, err := helpers.GetParamID(ctx, "photoId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		commentId, err := helpers.GetParamID(ctx, "commentId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
		}

		comment, err := au.commentRepo.GetCommentById(photoId, commentId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
		}

		if comment.UserId != user.Id {
			Unauthorized := errs.NewUnauthorizedError("you are not authorized")
			ctx.AbortWithStatusJSON(Unauthorized.Status(), Unauthorized)
			return
		}
		ctx.Next()
	}
}
