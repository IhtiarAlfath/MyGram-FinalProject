package handler

import (
	"kominfo/h8-myGram-finalproject/dto"
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/pkg/helpers"
	"kominfo/h8-myGram-finalproject/service"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler (commentService service.CommentService) commentHandler {
	return commentHandler{
		commentService: commentService,
	}
}

// CreateComment godoc
// @Tags Comment
// @Description Create Comment Data
// @ID create-comment-data
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "Photo's id"
// @Param RequestBody body dto.CreateCommentRequest true "request body json"
// @Success 201 {object} dto.CreateCommentResponse
// @Router /photo/{photoId}/comment [POST]
func (ch commentHandler) CreateComment (c *gin.Context) {
	photoId, err := helpers.GetParamID(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	var commentRequest dto.CreateCommentRequest

	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		errBindJson := errs.NewUnproccessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	response, err := ch.commentService.CreateComment(user.Id, photoId, commentRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// GetComment godoc
// @Tags Comment
// @Description Get All Comments Data
// @ID get-all-comment
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "Photo's id"
// @Success 200 {object} dto.GetCommentsResponse
// @Router /photo/{photoId}/comment [GET]
func (ch *commentHandler) GetComments (c *gin.Context) {
	photoId, err := helpers.GetParamID(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ch.commentService.GetComments(photoId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// GetCommentById godoc
// @Tags Comment
// @Description Get One Comment Data By Id
// @ID get-one-comment-data-by-id
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "Photo's id"
// @Param commentId path int true "Comment's id"
// @Success 200 {object} dto.GetCommentByIdResponse
// @Router /photo/{photoId}/comment/{commentId} [GET]
func (ch *commentHandler) GetCommentById (c *gin.Context) {
	photoId, err := helpers.GetParamID(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	commentId, err := helpers.GetParamID(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ch.commentService.GetCommentById(photoId, commentId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// UpdateComment godoc
// @Tags Comment
// @Description Update Comment Data
// @ID update-comment-data
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "Photo's id"
// @Param commentId path int true "Comment's id"
// @Param RequestBody body dto.CreateCommentRequest true "request body json"
// @Success 200 {object} dto.UpdatePhotoResponse
// @Router /photo/{photoId}/comment/{commentId} [PUT]
func (ch *commentHandler) UpdateCommentById (c *gin.Context){
	photoId, err := helpers.GetParamID(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	commentId, err := helpers.GetParamID(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	var updateRequest dto.CreateCommentRequest

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		errBindJson := errs.NewUnproccessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := ch.commentService.UpdateCommentById(photoId, commentId, updateRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// DeleteComment godoc
// @Tags Comment
// @Description Delete Comment Data By Id
// @ID delete-comment-data-by-id
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "Photo's id"
// @Param commentId path int true "Comment's id"
// @Success 200 {object} dto.DeleteCommentResponse
// @Router /photo/{photoId}/comment/{commentId} [DELETE]
func (ch *commentHandler) DeleteCommentById (c *gin.Context){
	photoId, err := helpers.GetParamID(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	commentId, err := helpers.GetParamID(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ch.commentService.DeleteCommentById(photoId, commentId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}