package main

import (
	"fmt"
	"log"
	"os"

	_ "wishlist/docs"

	"wishlist/db"
	"wishlist/server"

	"github.com/joho/godotenv"
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

// @title Wishlist
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 80.249.145.126:80
// @BasePath /

func main() {
	//setups (must be sync)
	db.Setup(getDsn())
	server.Setup()

	//start (must be async)
	go server.Start()

	//wait
	select {}
}
