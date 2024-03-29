package db

import (
	"database/sql"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/tajale72/asm/model"
)

type Service struct {
	Mongoclient *mongo.Client
	logger      *log.Logger
	Db          *sql.DB
	Name        string
}

type DatabaseInterface interface {
	Hello() (string, error)
	StoreFinancial(finance model.Finance, total float64)
	InserFinancialData(finance model.Finance, total float64) (*mongo.InsertOneResult, error)
	GetFinancial(id string) (*model.Finance, error)

	InsertAllUsers(users model.User) (*mongo.InsertOneResult, error)
	GetAllUsers() ([]model.GetUser, error)
	GetUserById(id string) (*model.User, error)
	UpdateUserById(users model.User, id string) (*mongo.UpdateResult, error)
	DeleteUserById(id string) (*mongo.DeleteResult, error)
}

func (s *Service) Hello() (string, error) {
	return "Main Website", nil
}
