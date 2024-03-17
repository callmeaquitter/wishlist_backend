package db

import (
	"fmt"


	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func getDsn() string {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	Db_host, Db_username, Db_password, Db_name, Db_port := os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_PORT")
// 	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", Db_host, Db_username, Db_password, Db_name, Db_port)
// }

var Database *gorm.DB

func Setup(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in New | Establishing a connection", err)
	}

	err = db.AutoMigrate(
		&Gift{}, &Seller{}, &Service{}, &SellerToService{}, &ServiceReview{})
	if err != nil {
		fmt.Println("Couldn't Automigrate Database.", err)
	}
	Database = db
}
