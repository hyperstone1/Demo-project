package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/hyperstone1/TestCRUD/internal/model"
	"github.com/pkg/errors"
)

var SignKey = []byte("hyperstone")

const (
	salt       = "hjqrhjqw124617ajfhajs"
	tokenttl   = 12 * time.Hour
	//usersTable = "users"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *CatService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.rep.Cat.CreateUser(user)
}

func (s *CatService) GenerateJWT(username, password string) (string, error) {
	log, err := s.rep.Cat.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(tokenttl)),
			IssuedAt:  jwt.At(time.Now()),
		},
		log.Id,
	})

	return token.SignedString([]byte(SignKey))

}

func (s CatService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(SignKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
