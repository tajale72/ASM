package controller

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type VerifyToken struct {
	Token string `json:"token"`
}

type JwtVerifyToken struct {
	Token     string    `json:"token"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	ExpiresAt time.Time `json:"expiresAt"`
	IssuesAt  time.Time `json:"issuesAt"`
}

func (s *Service) ProcessTokenBody(token_payload string) (*JwtVerifyToken, error) {
	token, err := jwt.ParseWithClaims(
		token_payload,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		log.Println("error parsing the claims")
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		log.Println("error with claims")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Println("token expired")
		return nil, err
	}
	Token := JwtVerifyToken{
		token_payload,
		claims.Username,
		claims.Email,
		claims.Password,
		time.Unix(0, claims.ExpiresAt),
		time.Unix(0, claims.IssuedAt)}
	return &Token, nil
}

func (s *Service) VerifyToken(body []byte) (*JwtVerifyToken, error) {
	log.Println("VerifyToken")
	var signedToken VerifyToken
	if len(body) == 0 {
		return nil, errors.New("token missing")
	}
	err := json.Unmarshal(body, &signedToken)
	if err != nil {
		return nil, errors.New("error unmarshalling the payload")
	}
	return s.ProcessTokenBody(signedToken.Token)

}
