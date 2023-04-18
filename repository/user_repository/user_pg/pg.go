package user_pg

import (
	"database/sql"
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/repository/user_repository"
)

const (
	createUserQuery = `
		INSERT INTO "users"
		(
			username,
			email,
			password,
			age
		)
		VALUES($1, $2, $3, $4);
	`

	getuserByEmailQuery = `
	SELECT 
		id, username, email, password, age 
	FROM 
		"users" 
	WHERE 
		email =$1
	`

	getuserByUsernameQuery = `
	SELECT 
		id, username, email, password, age 
	FROM 
		"users" 
	WHERE 
		email =$1
	`
)

type userPG struct {
	db *sql.DB
}

func NewUserPG(db *sql.DB) user_repository.UserRepository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) GetuserByEmail(email string) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(getuserByEmailQuery, email)

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) GetuserByUsername(username string) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(getuserByUsernameQuery, username)

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) Register(payload entity.User) errs.MessageErr {
	_, errCheck := u.GetuserByEmail(payload.Email)

	if errCheck == nil {
		return errs.NewBadRequest("email has been registered")
	}

	_, errCheck = u.GetuserByUsername(payload.Username)

	if errCheck == nil {
		return errs.NewBadRequest("username has been registered")
	}

	_, err := u.db.Exec(createUserQuery, payload.Username, payload.Email, payload.Password, payload.Age)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
