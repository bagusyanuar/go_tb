package common

import (
	"os"

	"github.com/bagusyanuar/go_tb/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTClaims struct {
	jwt.StandardClaims
	Unique   uuid.UUID `json:"unique"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     string    `json:"roles"`
}

func CreateJWTAccessToken(model *model.APISignUpResponse) (accessToken string, err error) {
	issuer := os.Getenv("JWT_ISSUER")
	JWTSignatureKey := os.Getenv("JWT_SIGNATURE_KEY")
	JWTSigninMethod := jwt.SigningMethodHS256
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: issuer,
		},
		Unique:   model.ID,
		Username: model.Username,
		Email:    model.Email,
		Role:     model.Role,
	}
	token := jwt.NewWithClaims(JWTSigninMethod, claims)
	return token.SignedString([]byte(JWTSignatureKey))
}
