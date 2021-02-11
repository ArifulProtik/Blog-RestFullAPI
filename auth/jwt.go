package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const hello string = "key"

// CreateToken Generates the Acess Token
func CreateToken(UUID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["UUID"] = UUID
	claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(hello))

}

// TokenValid Checks The Vakidity of Acess Token
func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(hello), nil
	})
	if err != nil {
		return err
	}
	_ = token
	return nil
}

// ExtractToken Extracts Access Key From Header
func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// ExtractTokenID extract ID From Aceesskey
func ExtractTokenID(r *http.Request) (string, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(hello), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		authUuid, ok := claims["UUID"].(string) //convert the interface to string
		if !ok {
			return "", err
		}
		return authUuid, nil
	}
	return "", nil
}
