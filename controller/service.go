package controller

import (
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"

	"myapi/db"
	"myapi/model"
)

type Service struct {
	DB db.DatabaseInterface
}

type ControllerInterface interface {
	Hello() (string, error)
	VerifyToken(body []byte) (*JwtVerifyToken, error)
	ProcessTokenBody(token_payload string) (*JwtVerifyToken, error)
	GenerateToken(body []byte) (string, error)
	StoreFinancial(body []byte) (*mongo.InsertOneResult, error)
	GetFinancial(body []byte) (*model.Finance, error)

	InsertAllUsers(body []byte) (*mongo.InsertOneResult, error)
	GetAllUsers() ([]model.GetUser, error)
	GetUserById(id string) (*model.User, error)
	UpdateUserById(body []byte, id string) (*mongo.UpdateResult, error)
	DeleteUserById(id string) (*mongo.DeleteResult, error)
}

func (s *Service) Hello() (string, error) {
	return s.DB.Hello()
}
