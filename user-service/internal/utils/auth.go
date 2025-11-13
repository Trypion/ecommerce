package utils

import (
	"time"

	"github.com/Trypion/ecommerce/user-service/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// return false if password matches
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err != nil
}

type tokenClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func SignJWT(user models.User) (string, error) {
	claims := tokenClaims{
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "place_holder",
			ID:        uuid.NewString(),
		},
	}

	// TODO: change sign method to RSA and use private key.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// TODO: place holder secret key. Change to env variable for production.
	return token.SignedString([]byte("mySuperSecretKey"))
}
