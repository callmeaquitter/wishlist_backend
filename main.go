package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "docs/docs.go"
	 
)

/*
You lack discipline, control, and worst? YOU'RE SLOPPY.
Тебе не хватает дисциплины, контроля. И что хуже всего ты работаешь небрежно.
- Alastor, Hazbin Hotel
*/

func getDsn() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Db_host, Db_username, Db_password, Db_name, Db_port := os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_PORT")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", Db_host, Db_username, Db_password, Db_name, Db_port)
}

func main() {
	//setups (must be sync)
	databaseSetup(getDsn())
	serverSetup()

	//start (must be async)
	go serverStart()

	//wait
	select {}
}
