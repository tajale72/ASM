package controller

import (
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/tajale72/asm/model"
)

func (s *Service) StoreFinancial(body []byte) (*mongo.InsertOneResult, error) {
	var finance model.Finance
	json.Unmarshal(body, &finance)
	finance.Date = time.Now()
	total := -finance.Discoveracardbalance + finance.DebitCardbalance + finance.Stockmarket_buyingpower - finance.Rent - finance.Utilities - finance.Phonebill - finance.Carinsurance - finance.WalmartGrocery
	s.DB.StoreFinancial(finance, total)
	res, err := s.DB.InserFinancialData(finance, total)
	if err != nil {
		log.Println("error from db error", err)
		return nil, err
	}
	return res, nil
}

func (s *Service) GetFinancial(body []byte) (*model.Finance, error) {
	id := string(body)
	return s.DB.GetFinancial(id)
}
