package socialmedia_repository

import (
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
)

type SocialMediaRepository interface {
	CreateSocialMedia(payload *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr)
	GetSocialMedias(userId int) ([]*entity.SocialMedia, errs.MessageErr)
	GetSocialMediabById(socialMediaId int) (*entity.SocialMedia, errs.MessageErr)
	UpdateSocialMediaById(payload *entity.SocialMedia) errs.MessageErr
	DeleteSocialMediaById(socialMediaId int) errs.MessageErr
}