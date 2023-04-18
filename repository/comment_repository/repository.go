package comment_repository

import (
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
)

type CommentRepository interface {
	CreateComment(payload entity.Comment) (*entity.Comment, errs.MessageErr)
	GetComments(photoId int) ([]*entity.Comment, errs.MessageErr)
	GetCommentById(photoId, commentId int) (*entity.Comment, errs.MessageErr)
	UpdateCommentById(payload entity.Comment) (errs.MessageErr)
	DeleteCommentById(photoId, commentId int) (errs.MessageErr)
}
