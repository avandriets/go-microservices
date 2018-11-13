package main

import (
	"./blob-reader"
	"./model"
	"fmt"
	"log"
)
import _ "github.com/lib/pq"
import "database/sql"

func main() {
	config := model.Config{}
	config.Init()
	model.SetConfig(&config)

	db := connectToDatabase()
	defer db.Close()

	blob_reader.ReadCSVFile("./data/data.csv", &model.ContactInfo{})
}

func connectToDatabase() *sql.DB {
	config := model.GetConfig()
	//connStr := "postgresql://nprl:nprl@localhost:5432/nprl?sslmode=disable"
	connStr := config.GetDBConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(fmt.Errorf("unable connect to DB: %v", err))
	}

	model.SetDatabase(db)

	return db
}
