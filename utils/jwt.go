package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = "SECRET_TOKEN"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return webToken, nil
}

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenStr)
	if err != nil {
		return nil, err
	}
	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
