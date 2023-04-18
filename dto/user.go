package dto

type NewUserRequest struct {
	Username string `json:"username" valid:"required~username cannot be empty" example:"budi"`
	Email    string `json:"email" valid:"required~email cannot be empty, email~invalid email" example:"budi@gmail.com"`
	Password string `json:"password" valid:"required~password cannot be empty, minstringlength(6)~password minimal 6 characters" example:"budibudi"`
	Age		 int	`json:"age" valid:"required~age cannot be empty, int~age must be a number" example:"12"`
}

type NewUserResponse struct {
	Result     string `json:"result" example:"success"`
	StatusCode int    `json:"statusCode" example:"201"`
	Message    string `json:"message" example:"registered successfully"`
}

type TokenResponse struct{
	Token 	string	`json:"token" example:"fgsfsdfbawgy45563456hfgbe"`
}

type LoginRequest struct{
	Email    string `json:"email" valid:"required~email cannot be empty, email~invalid email" example:"budi@gmail.com"`
	Password string `json:"password" valid:"required~password cannot be empty, minstringlength(6)~password minimal 6 characters" example:"budibudi"`
}

type LoginResponse struct {
	Result     string 			`json:"result" example:"success"`
	StatusCode int    			`json:"statusCode" example:"200"`
	Message    string 			`json:"message" example:"login successfully"`
	Data	   TokenResponse	`json:"data"` 
}
