package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"rs/auth/app/configs"
	"time"
)

type jwtT struct {
	secretKey string
}

var Jwt *jwtT

func init() {
	Jwt = &jwtT{secretKey: configs.Config.JWT_SECRET_KEY}
}

func (j *jwtT) GenerateToken(userID string, expiryDuration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expiryDuration).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtT) ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidId
}
