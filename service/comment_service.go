package service

import (
	"kominfo/h8-myGram-finalproject/dto"
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/pkg/helpers"
	"kominfo/h8-myGram-finalproject/repository/comment_repository"
	"net/http"
	"time"
)

type commentService struct {
	commentRepo comment_repository.CommentRepository
}

type CommentService interface {
	CreateComment(userId, photoId int, newCommentRequest dto.CreateCommentRequest) (*dto.CreateCommentResponse, errs.MessageErr)
	GetComments(photoId int) (*dto.GetCommentsResponse, errs.MessageErr)
	GetCommentById(photoId, commentId int) (*dto.GetCommentByIdResponse, errs.MessageErr)
	UpdateCommentById(photoId, commentId int, updateCommentRequest dto.CreateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr)
	DeleteCommentById(photoId, commentId int) (*dto.DeleteCommentResponse, errs.MessageErr)
}

func NewCommentService(commentRepo comment_repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (c *commentService) CreateComment(userId, photoId int, newCommentRequest dto.CreateCommentRequest) (*dto.CreateCommentResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newCommentRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Comment{
		UserId:  userId,
		PhotoId: photoId,
		Message: newCommentRequest.Message,
	}

	_, err = c.commentRepo.CreateComment(payload)

	if err != nil {
		return nil, err
	}

	response := dto.CreateCommentResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "successfully created comment",
	}

	return &response, nil
}

func (c *commentService) GetComments(photoId int) (*dto.GetCommentsResponse, errs.MessageErr) {
	result, err := c.commentRepo.GetComments(photoId)

	if err != nil {
		return nil, err
	}

	var comments []dto.CommentResult

	for _, eachComment := range result {
		comments = append(comments, eachComment.CommentEntityToDto())
	}

	response := dto.GetCommentsResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "comment data successfully sent",
		Data:       comments,
	}

	return &response, nil
}

func (c *commentService) GetCommentById(photoId, commentId int) (*dto.GetCommentByIdResponse, errs.MessageErr) {
	comment, err := c.commentRepo.GetCommentById(photoId, commentId)

	if err != nil {
		return nil, err
	}

	commentResult := comment.CommentEntityToDto()

	response := dto.GetCommentByIdResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "comment data successfully sent",
		Data:       commentResult,
	}

	return &response, nil
}

func (c *commentService) UpdateCommentById(photoId, commentId int, updateCommentRequest dto.CreateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(updateCommentRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Comment{
		Id:        commentId,
		PhotoId:   photoId,
		Message:   updateCommentRequest.Message,
		UpdatedAt: time.Now(),
	}

	err = c.commentRepo.UpdateCommentById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateCommentResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully updated comment",
	}

	return &response, nil
}

func (c *commentService) DeleteCommentById(photoId, commentId int) (*dto.DeleteCommentResponse, errs.MessageErr) {
	err := c.commentRepo.DeleteCommentById(photoId, commentId)

	if err != nil {
		return nil, err
	}

	response := dto.DeleteCommentResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully updated comment",
	}

	return &response, nil
}
