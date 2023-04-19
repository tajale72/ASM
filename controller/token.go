package controller

import (
	// import the jwt-go library
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Create a struct to read the username and password from the request body
type GenerateToken struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type JWTClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// Create the Signin handler
func (s *Service) GenerateToken(body []byte) (string, error) {
	var generatetoken GenerateToken
	err := json.Unmarshal(body, &generatetoken)
	if err != nil {
		s.logger.Println("error unmarshalling the token: " + err.Error())
		return "", errors.New("error unmarshalling the token: " + err.Error())
	}
	issuedat := time.Now().Unix()
	expirationTime := time.Now().Add(5 * time.Minute).Unix()
	claims := &JWTClaims{
		Email:    generatetoken.Email,
		Username: generatetoken.Username,
		Password: generatetoken.Password,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issuedat,
			ExpiresAt: expirationTime,
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		s.logger.Println("error unmarshalling the token: " + err.Error())
		return "", errors.New("error generating the token using the signekey: " + err.Error())
	}
	return tokenString, nil
}
