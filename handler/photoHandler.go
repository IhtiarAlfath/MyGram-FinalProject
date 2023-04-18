package handler

import (
	"kominfo/h8-myGram-finalproject/dto"
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/pkg/helpers"
	"kominfo/h8-myGram-finalproject/service"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) photoHandler {
	return photoHandler{
		photoService: photoService,
	}
}

// CreatePhoto godoc
// @Tags Photo
// @Description Create Photo Data
// @ID create-photo-data
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.CreatePhotoRequest true "request body json"
// @Success 201 {object} dto.CreatePhotoResponse
// @Router /photo [POST]
func (ph *photoHandler) CreatePhoto(c *gin.Context) {
	var payload dto.CreatePhotoRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		errBindJSON := errs.NewUnproccessibleEntityError("invalid request body")

		c.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	user := c.MustGet("userData").(entity.User)

	response, err := ph.photoService.CreatePhoto(user.Id, payload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// GetPhoto godoc
// @Tags Photo
// @Description Get All Photo Data
// @ID get-all-photo-data
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.GetPhotosResponse
// @Router /photo [GET]
func (ph *photoHandler) GetPhotos(c *gin.Context) {
	user := c.MustGet("userData").(entity.User)

	response, err := ph.photoService.GetPhotos(user.Id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// GetPhotoById godoc
// @Tags Photo
// @Description Get One Photo Data By Id
// @ID get-one-photo-data-by-id
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "Photo's id"
// @Success 200 {object} dto.GetPhotoByIdResponse
// @Router /photo/{photoId} [GET]
func (ph *photoHandler) GetPhotoByID(c *gin.Context) {
	photoId, err := helpers.GetParamID(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ph.photoService.GetPhotoById(photoId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// UpdatePhoto godoc
// @Tags Photo
// @Description Update Photo Data
// @ID update-photo-data
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "Photo's id"
// @Param RequestBody body dto.CreatePhotoRequest true "request body json"
// @Success 200 {object} dto.UpdatePhotoResponse
// @Router /photo/{photoId} [PUT]
func (ph *photoHandler) UpdatePhotoById(c *gin.Context) {
	photoId, err := helpers.GetParamID(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	var photoRequest dto.CreatePhotoRequest

	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		errBindJSON := errs.NewUnproccessibleEntityError("invalid request body")

		c.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	response, err := ph.photoService.UpdatePhotoById(photoId, photoRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// DeletePhoto godoc
// @Tags Photo
// @Description Delete Photo Data By Id
// @ID delete-photo-data-by-id
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "Photo's id"
// @Success 200 {object} dto.DeletePhotoResponse
// @Router /photo/{photoId} [DELETE]
func (ph *photoHandler) DeletePhotoById(c *gin.Context) {
	photoId, err := helpers.GetParamID(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ph.photoService.DeletePhotoById(photoId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}
