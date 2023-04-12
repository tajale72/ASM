package db

import (
	"database/sql"
	"log"

	"myapi/model"
)

func PostgressClient() (*sql.DB, error) {
	conn := "host=localhost port=5432 user=romittajale dbname=finance sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected!")
	return db, nil
}

func (s *Service) StoreFinancial(finance model.Finance, total float64) {
	// sqlStatement := `INSERT INTO users (age, email, first_name, last_name)
	// VALUES ($1, $2, $3, $4)
	// RETURNING id`
	sql := `INSERT INTO finance (date,discovercardbalance,debitcardbalance,stockmarket_buyingpower,rent,utilities,phonebill,carinsurance,walmartgrocery,total_net_worth)VALUES ($1, $2, $3, $4 , $5, $6, $7, $8, $9, $10);`
	id := 0
	//err := s.Db.QueryRow(sqlStatement, 30, "hello@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
	err := s.Db.QueryRow(sql, finance.Date, finance.Discoveracardbalance, finance.DebitCardbalance, finance.Stockmarket_buyingpower, finance.Rent, finance.Utilities, finance.Phonebill, finance.Carinsurance, finance.WalmartGrocery, total).Scan(&id)
	if err != nil {
		log.Println("error from queryrow ", err)
	}
}
