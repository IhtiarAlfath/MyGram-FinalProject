package dto

import "time"

type CommentResult struct {
	Id        int       `json:"id" example:"1"`
	UserId    int       `json:"user_id" example:"1"`
	PhotoId   int       `json:"photo_id" example:"1"`
	Message   string    `json:"message" example:"wow..."`
	CreatedAt time.Time `json:"created_at" example:"2023-04-14 09:44:00.462102+07"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-04-14 09:44:00.462102+07"`
}

type CreateCommentRequest struct {
	Message string `json:"message" valid:"required~message cannot be empty" example:"wow..."`
}

type CreateCommentResponse struct {
	Result     string `json:"result" example:"success"`
	StatusCode int    `json:"statusCode" example:"201"`
	Message    string `json:"message" example:"successfully created comment"`
}

type GetCommentsResponse struct {
	Result     string          `json:"result" example:"success"`
	StatusCode int             `json:"statusCode" example:"200"`
	Message    string          `json:"message" example:"comment data successfully sent"`
	Data       []CommentResult `json:"data"`
}

type GetCommentByIdResponse struct {
	Result     string        `json:"result" example:"success"`
	StatusCode int           `json:"statusCode" example:"200"`
	Message    string        `json:"message" example:"comment data successfully sent"`
	Data       CommentResult `json:"data"`
}

type UpdateCommentResponse struct {
	Result     string `json:"result" example:"success"`
	StatusCode int    `json:"statusCode" example:"200"`
	Message    string `json:"message" example:"successfully updated comment"`
}

type DeleteCommentResponse struct {
	Result     string `json:"result" example:"success"`
	StatusCode int    `json:"statusCode" example:"200"`
	Message    string `json:"message" example:"successfully deleted comment"`
}
