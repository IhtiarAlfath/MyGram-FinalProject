package dto

import "time"

type SocialMediaResult struct {
	Id             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"budi_"`
	SocialMediaURL string    `json:"social_media_url" example:"budi_.com"`
	UserId         int       `json:"user_id" example:"1"`
	CreatedAt      time.Time `json:"created_at" example:"2023-04-14 09:44:00.462102+07"` 
	UpdatedAt      time.Time `json:"updated_at" example:"2023-04-14 09:44:00.462102+07"`
}

type CreateSocialMediaRequest struct {
	Name           string `json:"name" valid:"required~name cannot be empty"`
	SocialMediaURL string `json:"photo_url" valid:"required~social media URL cannot be empty"`
}

type CreateSocialMediaResponse struct {
	Result     string `json:"result" example:"success"`
	StatusCode int    `json:"statusCode" example:"201"`
	Message    string `json:"message" example:"successfully created social media"`
}

type GetSocialMediasResponse struct {
	Result     string              `json:"result" example:"success"`
	StatusCode int                 `json:"statusCode" example:"200"`
	Message    string              `json:"message" example:"social media data successfully sent"`
	Data       []SocialMediaResult `json:"data"`
}

type GetSocialMediaByIdResponse struct {
	Result     string               `json:"result" example:"success"`
	StatusCode int                  `json:"statusCode" example:"200"`
	Message    string               `json:"message" example:"social media data successfully sent"`
	Data       SocialMediaResult 	`json:"data"`
}

type UpdateSocialMediaResponse struct {
	Result     string `json:"result" example:"success"`
	StatusCode int    `json:"statusCode" example:"200"`
	Message    string `json:"message" example:"successfully updated social media"`
}

type DeleteSocialMediaResponse struct {
	Result     string `json:"result" example:"success"`
	StatusCode int    `json:"statusCode" example:"200"`
	Message    string `json:"message" example:"successfully deleted social media"`
}