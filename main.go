package main

import (
	"fmt"
	"log"
	"travelis/m/v2/modules/db"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(fmt.Errorf("error loading .env file: %v", err))
	}

	db := db.InitDB()
	defer db.Close()
}
