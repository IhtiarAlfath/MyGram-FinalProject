package entity

import (
	"kominfo/h8-myGram-finalproject/dto"
	"time"
)

type SocialMedia struct {
	Id          	int       `json:"id"`
	Name        	string    `json:"name"`
	SocialMediaURL	string    `json:"social_media_url"`
	UserId      	int       `json:"user_id"`
	CreatedAt   	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func (s *SocialMedia) SocialMediaEntityToDto() dto.SocialMediaResult{
	return dto.SocialMediaResult{
		Id: s.Id,
		Name: s.Name,
		SocialMediaURL: s.SocialMediaURL,
		UserId: s.UserId,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}