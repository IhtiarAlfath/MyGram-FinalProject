package service

import (
	"kominfo/h8-myGram-finalproject/dto"
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/pkg/helpers"
	"kominfo/h8-myGram-finalproject/repository/socialmedia_repository"
	"net/http"
	"time"
)

type socialMediaService struct {
	socialMediaRepo socialmedia_repository.SocialMediaRepository
}

type SocialMediaService interface {
	CreateSocialMedia(userId int, newSocialMediaRequest dto.CreateSocialMediaRequest) (*dto.CreateSocialMediaResponse, errs.MessageErr)
	GetSocialMedias(userId int) (*dto.GetSocialMediasResponse, errs.MessageErr)
	GetSocialMediaById(socialMediaId int) (*dto.GetSocialMediaByIdResponse, errs.MessageErr)
	UpdateSocialMediaById(socialMediaId int, updateRequest dto.CreateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr)
	DeleteSocialMediaById(socialMediaId int) (*dto.DeleteSocialMediaResponse, errs.MessageErr)
}

func NewSocialMediaService(socialMediaRepo socialmedia_repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{
		socialMediaRepo: socialMediaRepo,
	}
}

func (s *socialMediaService) CreateSocialMedia(userId int, newSocialMediaRequest dto.CreateSocialMediaRequest) (*dto.CreateSocialMediaResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newSocialMediaRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.SocialMedia{
		Name:           newSocialMediaRequest.Name,
		SocialMediaURL: newSocialMediaRequest.SocialMediaURL,
		UserId:         userId,
	}

	_, err = s.socialMediaRepo.CreateSocialMedia(&payload)

	if err != nil {
		return nil, err
	}

	response := dto.CreateSocialMediaResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "successfully created social media",
	}

	return &response, nil
}

func (s *socialMediaService) GetSocialMedias(userId int) (*dto.GetSocialMediasResponse, errs.MessageErr) {
	var socialMedias []dto.SocialMediaResult

	result, err := s.socialMediaRepo.GetSocialMedias(userId)

	if err != nil {
		return nil, err
	}

	for _, eachSocialMedia := range result {
		socialMedias = append(socialMedias, eachSocialMedia.SocialMediaEntityToDto())
	}

	response := dto.GetSocialMediasResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "social media data successfully sent",
		Data:       socialMedias,
	}

	return &response, nil
}

func (s *socialMediaService) GetSocialMediaById(socialMediaId int) (*dto.GetSocialMediaByIdResponse, errs.MessageErr) {
	result, err := s.socialMediaRepo.GetSocialMediabById(socialMediaId)

	if err != nil {
		return nil, err
	}

	socialMedia := result.SocialMediaEntityToDto()

	response := dto.GetSocialMediaByIdResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "social media data successfully sent",
		Data:       socialMedia,
	}

	return &response, nil
}

func (s *socialMediaService) UpdateSocialMediaById(socialMediaId int, updateRequest dto.CreateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(updateRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.SocialMedia{
		Id:             socialMediaId,
		Name:           updateRequest.Name,
		SocialMediaURL: updateRequest.SocialMediaURL,
		UpdatedAt:      time.Now(),
	}

	err = s.socialMediaRepo.UpdateSocialMediaById(&payload)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateSocialMediaResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully updated social media",
	}

	return &response, nil
}

func (s *socialMediaService) DeleteSocialMediaById(socialMediaId int) (*dto.DeleteSocialMediaResponse, errs.MessageErr) {
	err := s.socialMediaRepo.DeleteSocialMediaById(socialMediaId)

	if err != nil {
		return nil, err
	}

	response := dto.DeleteSocialMediaResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully deleted social media",
	}

	return &response, nil
}
