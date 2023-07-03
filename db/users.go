package db

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/tajale72/asm/model"
)

// InsertUser inserts data into the test database
func (s *Service) InsertAllUsers(users model.User) (*mongo.InsertOneResult, error) {
	collection := s.Mongoclient.Database("test").Collection("consultants")

	hashValue := hash(users)

	// // Check if the hash value exists in the collection
	// exists, err := hashExists(collection, hashValue)
	// if err != nil {
	// 	s.logger.Println(err)
	// }
	// if exists {
	// 	fmt.Printf("Hash value %x exists in the collection\n", hashValue)
	// 	return nil, err
	// }

	users.HashData = string(hashValue)
	res, err := collection.InsertOne(context.Background(), users)
	if err != nil {
		return nil, err
	}
	return res, nil

}

// InsertUser inserts data into the test database
func (s *Service) UpdateUserById(users model.User, id string) (*mongo.UpdateResult, error) {

	collection := s.Mongoclient.Database("test").Collection("consultants")
	objID, err := primitive.ObjectIDFromHex(id)

	res, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.D{
			{"$set", users},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil

}

// InsertUser inserts data into the test database
func (s *Service) GetAllUsers() ([]model.GetUser, error) {
	var users []model.GetUser
	collection := s.Mongoclient.Database("test").Collection("consultants")

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("error from getting an object ", err)
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var user model.GetUser
		err := cur.Decode(&user)
		if err != nil {
			log.Println("error from getting single data from cur", err)
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		err := errors.New("empty collection")
		log.Println("error from getting an object ", err)
		return nil, err
	}
	return users, nil

}

// InsertUser inserts data into the test database
func (s *Service) GetUserById(id string) (*model.User, error) {
	var user model.User
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := s.Mongoclient.Database("test").Collection("consultants")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

// InsertUser inserts data into the test database
func (s *Service) DeleteUserById(id string) (*mongo.DeleteResult, error) {
	collection := s.Mongoclient.Database("test").Collection("consultants")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	log.Println(objID)
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		log.Println("error from getting an object ", err)
		return nil, err
	}

	return res, nil

}
