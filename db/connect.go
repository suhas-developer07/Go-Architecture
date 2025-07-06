package db

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect()(*sql.DB,error){
	url := os.Getenv("POSTGRES_URL")

	if url == ""{
		log.Fatalln("Postgress url is not loaded")
	}
	db,err := sql.Open("pgx",url)

	if err != nil {
		log.Fatalln("Database connection failed",err)
		return nil,err
	}
	err = db.Ping()
	
	if err != nil{
		return nil, err
	}
	return db,nil
}