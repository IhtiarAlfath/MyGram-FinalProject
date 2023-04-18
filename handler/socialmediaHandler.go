package handler

import (
	"kominfo/h8-myGram-finalproject/dto"
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/pkg/helpers"
	"kominfo/h8-myGram-finalproject/service"

	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService service.SocialMediaService) socialMediaHandler {
	return socialMediaHandler{
		socialMediaService: socialMediaService,
	}
}

// CreateSocialMedia godoc
// @Tags Social Media
// @Description Create Social Media Data
// @ID create-social-media-data
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.CreateSocialMediaRequest true "request body json"
// @Success 201 {object} dto.CreatePhotoResponse
// @Router /socialmedia [POST]
func (sh *socialMediaHandler) CreateSocialMedia(c *gin.Context) {
	var newSocialMediaRequest dto.CreateSocialMediaRequest

	if err := c.ShouldBindJSON(&newSocialMediaRequest); err != nil {
		errBindJSON := errs.NewUnproccessibleEntityError("invalid request body")

		c.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	user := c.MustGet("userData").(entity.User)

	response, err := sh.socialMediaService.CreateSocialMedia(user.Id, newSocialMediaRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// GetAllSocialMedia godoc
// @Tags Social Media
// @Description Get All Social Media Data
// @ID get-all-social-media-data
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.GetSocialMediasResponse
// @Router /socialmedia [GET]
func (sh *socialMediaHandler) GetSocialMedias(c *gin.Context) {
	user := c.MustGet("userData").(entity.User)

	response, err := sh.socialMediaService.GetSocialMedias(user.Id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// GetSocialMediaById godoc
// @Tags Social Media
// @Description Get One Social Media Data By Id
// @ID get-one-social-media-data-by-id
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param socialMediaId path int true "Social media's id"
// @Success 200 {object} dto.GetSocialMediaByIdResponse
// @Router /socialmedia/{socialMediaId} [GET]
func (sh *socialMediaHandler) GetSocialMediabyId(c *gin.Context) {
	socialMediaId, err := helpers.GetParamID(c, "socialMediaId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := sh.socialMediaService.GetSocialMediaById(socialMediaId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// UpdateSocialMedia godoc
// @Tags Social Media
// @Description Update Social Media Data
// @ID update-social-media-data
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param socialMediaId path int true "Social media's id"
// @Param RequestBody body dto.CreateSocialMediaRequest true "request body json"
// @Success 200 {object} dto.UpdatePhotoResponse
// @Router /socialmedia/{socialMediaId} [PUT]
func (sh *socialMediaHandler) UpdateSocialMediabyId(c *gin.Context) {
	socialMediaId, err := helpers.GetParamID(c, "socialMediaId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	var SocialMediaRequest dto.CreateSocialMediaRequest

	if err := c.ShouldBindJSON(&SocialMediaRequest); err != nil {
		errBindJSON := errs.NewUnproccessibleEntityError("invalid request body")

		c.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	response, err := sh.socialMediaService.UpdateSocialMediaById(socialMediaId, SocialMediaRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// DelteSocialMediaById godoc
// @Tags Social Media
// @Description Delete Social Media Data By Id
// @ID delete-social-media-data-by-id
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param socialMediaId path int true "Social media's id"
// @Success 200 {object} dto.DeletePhotoResponse
// @Router /socialmedia/{socialMediaId} [DELETE]
func (sh *socialMediaHandler) DeleteSocialMediabyId(c *gin.Context) {
	socialMediaId, err := helpers.GetParamID(c, "socialMediaId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := sh.socialMediaService.DeleteSocialMediaById(socialMediaId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}
