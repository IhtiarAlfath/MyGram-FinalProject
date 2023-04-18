package entity

import (
	"kominfo/h8-myGram-finalproject/dto"
	"time"
)

type Comment struct {
	Id          	int       `json:"id"`
	UserId      	int       `json:"user_id"`
	PhotoId      	int       `json:"photo_id"`
	Message        	string    `json:"message"`
	CreatedAt   	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

func (c *Comment) CommentEntityToDto () dto.CommentResult{
	return dto.CommentResult{
		Id: c.Id,
		UserId: c.UserId,
		PhotoId: c.PhotoId,
		Message: c.Message,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}