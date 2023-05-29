package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	config "user-service/pkg/common"
)

// JWTManager represents a JWT manager.
type JWTManager struct {
	secretKey       string
	tokenExpiration time.Duration
}

var jwtInstanceByUser *JWTManager
var jwtInstanceByPhone *JWTManager

// GetJWTManager returns the singleton instance of JWTManager.
func GetJWTManager() *JWTManager {
	if jwtInstanceByUser == nil {
		jwtInstanceByUser = &JWTManager{
			secretKey:       config.Config.JwtAuthKeyUserID,
			tokenExpiration: time.Hour * 3600,
		}
	}
	return jwtInstanceByUser
}

func GetPhoneJWTManager() *JWTManager {
	if jwtInstanceByPhone == nil {
		jwtInstanceByPhone = &JWTManager{
			secretKey:       config.Config.JwtAuthKeyPhone,
			tokenExpiration: time.Hour * 2,
		}
	}
	return jwtInstanceByPhone
}

// GenerateToken generates a new JWT token for the given claims.
func (jm *JWTManager) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jm.secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// VerifyToken verifies the validity of the given token and returns the claims if valid.
func (jm *JWTManager) VerifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jm.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token.Claims, nil
}
