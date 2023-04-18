package entity

import (
	"kominfo/h8-myGram-finalproject/dto"
	"time"
)

type Photo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Caption		string    `json:"caption"`
	PhotoURL	string    `json:"photo_url"`
	UserId      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Photo) PhotoEntityToDto () dto.PhotoResult {
	return dto.PhotoResult{
		Id: p.Id,
		Title: p.Title,
		Caption: p.Caption,
		PhotoURL: p.PhotoURL,
		UserId: p.UserId,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}