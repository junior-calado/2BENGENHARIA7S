package service

import (
	"2BENGENHARIA7S/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("segredoSegredozo") // Em produção, use uma variável de ambiente

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func RegisterUser(username, password, email string) (model.User, error) {
	if DB == nil {
		return model.User{}, errors.New("database not initialized")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
	}

	if err := DB.Create(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func LoginUser(username, password string) (string, error) {
	if DB == nil {
		return "", errors.New("database not initialized")
	}

	var user model.User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("user not found")
		}
		return "", err
	}

	if !CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, err := GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
} 