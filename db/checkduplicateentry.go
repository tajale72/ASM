package db

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"github.com/tajale72/asm/model"
)

type Document struct {
	Hash []byte `bson:"hash"`
}

func hash(data model.User) []byte {
	h := sha256.New()
	marshaldata, _ := json.Marshal(data)
	h.Write(marshaldata)
	return h.Sum(nil)
}

func hashExists(collection *mongo.Collection, hash []byte) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	var user model.User
	hashobjID, err := primitive.ObjectIDFromHex(string(hash))

	if err != nil {
		return false, err
	}
	err = collection.FindOne(ctx, bson.M{"hash": hashobjID}).Decode(&user)
	if err != nil {
		return false, err
	}
	return true, nil
}
