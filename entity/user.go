package entity

import (
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int       `json:"id"`
	Username  string	`json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age		  int		`json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var secret_key = "RAHASIA"

var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")

func (u *User) HashPassword() errs.MessageErr {
	salt := 8

	userPassword := []byte(u.Password)

	bs, err := bcrypt.GenerateFromPassword(userPassword, salt)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	u.Password = string(bs)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) SignToken(claims jwt.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstring , _ := token.SignedString([]byte(secret_key))

	return tokenstring
}	

func (u *User) TokenClaim() jwt.MapClaims{
	return jwt.MapClaims{
		"id" : u.Id,
		"email" : u.Email,
		"exp" : time.Now().Add(time.Hour * 10).Unix(),
	}
}

func (u *User) GenerateToke() string{
	claims := u.TokenClaim()

	return u.SignToken(claims)
}

func (u *User) ParseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenErr
		}

		return []byte(secret_key), nil
	})

	if err != nil {
		return nil, invalidTokenErr
	}

	return token, nil
}

func (u *User) bindTokenUserEntity(claim jwt.MapClaims) errs.MessageErr{
	if id, ok := claim["id"].(float64); !ok {
		return invalidTokenErr
	} else {
		u.Id = int(id)
	}

	if email, ok := claim["email"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Email = string(email)
	}

	return nil
}

func (u *User) ValidateToken (bearerToken string) errs.MessageErr {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return invalidTokenErr
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return invalidTokenErr
	}

	tokenString := splitToken[1]

	token, err := u.ParseToken(tokenString)

	if err != nil {
		return invalidTokenErr
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid{
		return invalidTokenErr
	} else {
		mapClaims = claims
	}

	err = u.bindTokenUserEntity(mapClaims)

	return err
}