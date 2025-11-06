package utils

import (
	stdErrors "errors"
	"time"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"github.com/Aboagye-Dacosta/shopBackend/internal/env"
	appErrors "github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID      string   `json:"user_id"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId string, roles []models.Role) (string, error) {

	// Define the JWT claims
	claims := Claims{
		UserID:      userId,
		Roles:       getRoles(roles),
		Permissions: getPermissions(roles),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(env.GetStringEnv("JWT_SECRETE", "klwelwkewlek")))
}

func VerifyJWT(tokenString string) (string, []string, error) {
	if tokenString == "" {
		return "", nil, appErrors.NoTokenProvided(nil)
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
			return "", nil, appErrors.ExpiredToken(err)
		}
		return "", nil, appErrors.InvalidToken(err)
	}

	if !token.Valid {
		return "", nil, appErrors.InvalidToken(nil)
	}

	return claims.UserID, claims.Permissions, nil
}

func getPermissions(roles []models.Role) []string {
	var permissions = make([]string, 0)

	for _, role := range roles {
		for _, perm := range role.Permissions {
			permissions = append(permissions, perm.Name)
		}
	}

	return permissions
}
func getRoles(roles []models.Role) []string {
	var values = make([]string, 0)

	for _, role := range roles {
		values = append(values, role.Name)

	}

	return values
}
