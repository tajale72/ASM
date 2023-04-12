package controller

import (
	"encoding/json"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"myapi/model"
)

func (s *Service) InsertAllUsers(body []byte) (*mongo.InsertOneResult, error) {
	var user model.User
	err := json.Unmarshal(body, &user)
	user.FirstName = EncrytionService(user.FirstName)
	if err != nil {
		return nil, err
	}
	log.Println(user.FirstName)
	return s.DB.InsertAllUsers(user)

}

func (s *Service) UpdateUserById(body []byte, id string) (*mongo.UpdateResult, error) {
	var user model.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return s.DB.UpdateUserById(user, id)

}

func (s *Service) GetAllUsers() ([]model.GetUser, error) {
	return s.DB.GetAllUsers()
}

func (s *Service) GetUserById(id string) (*model.User, error) {
	if id == "" {
		return nil, errors.New("missing id")
	}
	return s.DB.GetUserById(id)
}

func (s *Service) DeleteUserById(id string) (*mongo.DeleteResult, error) {
	if id == "" {
		return nil, errors.New("missing id")
	}
	return s.DB.DeleteUserById(id)
}
