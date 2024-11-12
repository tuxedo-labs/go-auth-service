package utils

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey string

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return webtoken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOke := token.Claims.(jwt.MapClaims)
	if isOke && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid Token")
}

func init() {
	SecretKey = os.Getenv("SECRET_KEY")
	if SecretKey == "" {
		SecretKey = "n3^|e{jJ,|UmsT(ch42^yl8x^=7#zp}q"
	}
}
