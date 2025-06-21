package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateJWTFromPayload(payload map[string]interface{}) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)

    claims := jwt.MapClaims{}
    for k, v := range payload {
        claims[k] = v
    }

    // Set standard claims
    claims["exp"] = expirationTime.Unix()
    claims["iat"] = time.Now().Unix()

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte("your_secret_key")) 
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
