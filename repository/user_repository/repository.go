package user_repository

import (
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
)

type UserRepository interface {
	GetuserByEmail(email string) (*entity.User, errs.MessageErr)
	GetuserByUsername(username string) (*entity.User, errs.MessageErr)
	Register(payload entity.User) errs.MessageErr
}
