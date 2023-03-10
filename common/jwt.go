package common

import (
	"errors"
	"os"

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

type JWTSignReturn struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
}

func CreateJWTAccessToken(jwtSign *JWTSignReturn) (accessToken string, err error) {
	issuer := os.Getenv("JWT_ISSUER")
	JWTSignatureKey := os.Getenv("JWT_SIGNATURE_KEY")
	JWTSigninMethod := jwt.SigningMethodHS256
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: issuer,
		},
		Unique:   jwtSign.ID,
		Username: jwtSign.Username,
		Email:    jwtSign.Email,
		Role:     jwtSign.Role,
	}
	token := jwt.NewWithClaims(JWTSigninMethod, claims)
	return token.SignedString([]byte(JWTSignatureKey))
}

func ClaimToken(authoriztion string) (interface{}, error) {
	if authoriztion == "" {
		return nil, errors.New("unauthorized")
	}
	bearer := string(authoriztion[0:7])
	tokenString := string(authoriztion[7:])
	JWTSignatureKey := os.Getenv("JWT_SIGNATURE_KEY")

	if bearer != "Bearer " {
		return nil, errors.New("invalid bearer")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid jwt sign in method")
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid jwt sign in method")
		}
		return []byte(JWTSignatureKey), nil
	})

	if err != nil {
		return nil, errors.New("error parse jwt")
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("error get jwt claim")
	}
	return claim, nil
}
