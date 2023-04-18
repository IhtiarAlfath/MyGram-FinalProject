package service

import (
	"kominfo/h8-myGram-finalproject/dto"
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/pkg/helpers"
	"kominfo/h8-myGram-finalproject/repository/photo_repository"
	"net/http"
	"time"
)

type photoService struct {
	photoRepo photo_repository.PhotoRepository
}

type PhotoService interface {
	CreatePhoto(userID int, payload dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr)
	GetPhotos(userId int) (*dto.GetPhotosResponse, errs.MessageErr)
	GetPhotoById(photoId int) (*dto.GetPhotoByIdResponse, errs.MessageErr)
	UpdatePhotoById(photoId int, payload dto.CreatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr)
	DeletePhotoById(photoId int) (*dto.DeletePhotoResponse, errs.MessageErr)
}

func NewPhotoService(photoRepo photo_repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepo: photoRepo,
	}
}

func (ps *photoService) CreatePhoto(userID int, payload dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	newPhoto := entity.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoURL: payload.PhotoURL,
		UserId:   userID,
	}

	_, err = ps.photoRepo.CreatePhoto(&newPhoto)

	if err != nil {
		return nil, err
	}

	response := dto.CreatePhotoResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "successfully created photo",
	}

	return &response, err
}

func (ps *photoService) GetPhotos(userId int) (*dto.GetPhotosResponse, errs.MessageErr) {
	photos, err := ps.photoRepo.GetPhotos(userId)

	if err != nil {
		return nil, err
	}

	var photoResponse []dto.PhotoResult

	for _, eachPhotos := range photos {
		photoResponse = append(photoResponse, eachPhotos.PhotoEntityToDto())
	}

	Response := dto.GetPhotosResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "photos data successfully sent",
		Data:       photoResponse,
	}

	return &Response, nil
}

func (ps *photoService) GetPhotoById(photoId int) (*dto.GetPhotoByIdResponse, errs.MessageErr) {
	photo, err := ps.photoRepo.GetPhotoByID(photoId)

	if err != nil {
		return nil, err
	}

	photoResponse := photo.PhotoEntityToDto()

	Response := dto.GetPhotoByIdResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "photo data successfully sent",
		Data:       photoResponse,
	}

	return &Response, nil
}

func (ps *photoService) UpdatePhotoById(photoId int, updateRequest dto.CreatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(updateRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Photo{
		Id: photoId,
		Title:     updateRequest.Title,
		Caption:   updateRequest.Caption,
		PhotoURL:  updateRequest.PhotoURL,
		UpdatedAt: time.Now(),
	}

	err = ps.photoRepo.UpdatePhotoById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.UpdatePhotoResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully updated photo",
	}

	return &response, err
}

func (ps *photoService) DeletePhotoById(photoId int) (*dto.DeletePhotoResponse, errs.MessageErr) {
	err := ps.photoRepo.DeletePhotoById(photoId)

	if err != nil {
		return nil, err
	}

	response := dto.DeletePhotoResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully deleted photo",
	}

	return &response, err
}