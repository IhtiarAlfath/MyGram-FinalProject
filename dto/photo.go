package dto

import "time"

type PhotoResult struct {
	Id        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"myphoto"`
	Caption   string    `json:"caption" example:"myphoto description"`
	PhotoURL  string    `json:"photo_url" example:"photo.com"`
	UserId    int       `json:"user_id" example:"2"`
	CreatedAt time.Time `json:"created_at" example:"2023-04-14 09:44:00.462102+07"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-04-14 09:44:00.462102+07"`
}


type CreatePhotoRequest struct {
	Title    string `json:"title" valid:"required~title cannot be empty" example:"myphoto"`
	Caption  string `json:"caption" valid:"-" example:"myphoto description"`
	PhotoURL string `json:"photo_url" valid:"required~photoURL cannot be empty" example:"photo.com"`
}

type CreatePhotoResponse struct {
	Result     string 			`json:"result" example:"success"`
	StatusCode int    			`json:"statusCode" example:"201"`
	Message    string 			`json:"message" example:"successfully created photo"`
}

type GetPhotosResponse struct {
	Result     string 			`json:"result" example:"success"`
	StatusCode int    			`json:"statusCode" example:"200"`
	Message    string 			`json:"message" example:"photos data successfully sent"`
	Data	   []PhotoResult	`json:"data"`
}

type GetPhotoByIdResponse struct {
	Result     string 			`json:"result" example:"success"`
	StatusCode int    			`json:"statusCode" example:"200"`
	Message    string 			`json:"message" example:"photo data successfully sent"`
	Data	   PhotoResult		`json:"data"`
}

type UpdatePhotoResponse struct {
	Result     string 			`json:"result" example:"success"`
	StatusCode int    			`json:"statusCode" example:"200"`
	Message    string 			`json:"message" example:"successfully updated photo"`
}

type DeletePhotoResponse struct {
	Result     string 			`json:"result" example:"success"`
	StatusCode int    			`json:"statusCode" example:"200"`
	Message    string 			`json:"message" example:"successfully deleted photo"`
}