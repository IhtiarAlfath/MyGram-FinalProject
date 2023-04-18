package photo_repository

import (
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
)

type PhotoRepository interface {
	CreatePhoto(payload *entity.Photo) (*entity.Photo, errs.MessageErr)
	GetPhotos(userId int) ([]*entity.Photo, errs.MessageErr)
	GetPhotoByID(photoID int) (*entity.Photo, errs.MessageErr)
	UpdatePhotoById(payload entity.Photo) (errs.MessageErr)
	DeletePhotoById(photoId int) (errs.MessageErr)
}
