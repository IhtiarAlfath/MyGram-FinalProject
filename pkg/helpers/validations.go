package helpers

import (
	"kominfo/h8-myGram-finalproject/pkg/errs"

	"github.com/asaskevich/govalidator"
)

func ValidateStruct(payload interface{}) errs.MessageErr{
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func ValidateAgeUser(num int) errs.MessageErr{
	if num < 8{
		return errs.NewBadRequest("age must be over 8")
	}
	return nil
}

// func CheckEmailUser(email string) errs.MessageErr{
// 	db := database.GetDatabaseInstance()
	
// 	userRepo := user_pg.NewuserPG()
// }