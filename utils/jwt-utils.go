package utils

import (
	stdErrors "errors"
	"time"

	"github.com/Aboagye-Dacosta/shopBackend/internal/env"
	appErrors "github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId string) (string, error) {
	// Define the JWT claims
	claims := Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(env.GetStringEnv("JWT_SECRETE", "klwelwkewlek")))
}

func VerifyJWT(tokenString string) (string, error) {
	if tokenString == "" {
		return "", appErrors.NoTokenProvided(nil)
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, appErrors.InvalidToken(stdErrors.New("unexpected signing method"))
		}
		return []byte(env.GetStringEnv("JWT_SECRETE", "klwelwkewlek")), nil
	})

	if err != nil {
		if stdErrors.Is(err, jwt.ErrTokenExpired) {
			return "", appErrors.ExpiredToken(err)
		}
		return "", appErrors.InvalidToken(err)
	}

	if !token.Valid {
		return "", appErrors.InvalidToken(nil)
	}

	return claims.UserID, nil
}
