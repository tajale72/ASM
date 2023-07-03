package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"github.com/tajale72/asm/model"
)

func Mongo() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("error getting the mongo client", err.Error)
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to mongo dbß")
	//s.dbclient = client
	return client, nil
}

// InsertUser inserts data into the test database
func (s *Service) InserFinancialData(finance model.Finance, total float64) (*mongo.InsertOneResult, error) {
	todaysdate := time.Now().Format("01-02-2006")
	finance.Networth = total
	collection := s.Mongoclient.Database("test").Collection(todaysdate)
	res, err := collection.InsertOne(context.Background(), finance)
	if err != nil {
		return nil, err
	}
	return res, nil

}

// InsertUser inserts data into the test database
func (s *Service) GetFinancial(id string) (*model.Finance, error) {
	todaysdate := time.Now().Format("01-02-2006")
	collection := s.Mongoclient.Database("test").Collection(todaysdate)
	var user model.Finance
	err := collection.FindOne(context.Background(), bson.M{"discoveracardbalance": 2300}).Decode(&user)
	if err != nil {
		log.Println("error from getting an object ", err)
		return nil, err
	}
	return &user, nil

}
