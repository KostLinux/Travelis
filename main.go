package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"travelis/m/v2/middleware"
	"travelis/m/v2/repository"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(fmt.Errorf("error loading .env file: %v", err))
	}

	db := repository.InitDB()
	defer db.Close()

	router := middleware.Router()

	fmt.Println("Starting server on the port " + os.Getenv("SERVER_PORT") + "...")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router))
}
