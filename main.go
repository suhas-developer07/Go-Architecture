package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/suhas-developer07/Bank-service/db"
	"github.com/suhas-developer07/Bank-service/repository"
	"github.com/suhas-developer07/Bank-service/routes"
)

func main(){
	if err := godotenv.Load();err!=nil{
		log.Fatalln("env file is not loaded")
	}
	serverAddress := os.Getenv("APP_PORT")

	dbConn,err := db.Connect()

	if err!=nil {
		log.Fatalln("Failed to connect to database:", err)
	}
	log.Println("Datbase connected")
	
	repo := repository.NewPostgresRepository(dbConn)

	router := routes.MountRoutes(repo)

	http.ListenAndServe(serverAddress,router)
}