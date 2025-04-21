package utils

import (
	
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

// JWTClaims is a struct that represents the claims in a JWT token.
type JWTClaims struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

//GenerateToken creates a JWT token for a gifen user ID.
func GenerateToken(userID uint) (string, error) {
	//Create the claims payload with custom and registered claims
	claims := &JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	//Create the token using HS256 signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	//Sign the token with the secret key and return it
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//validateToken checks if the given token is valid and returns the claims if it is.

func ValidateToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error){
		return jwtKey, nil
	})
	//step 2 Check if the token is valid and not expired and any other errors
	if err != nil {
		return nil, err
	}
	return claims, nil

	
}