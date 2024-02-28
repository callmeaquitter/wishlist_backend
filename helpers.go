package main

func getUserID() string {
	return "user_1"
}

//Step by step guide
//1. Create a model (if doesn't exist)
// - Add to AutoMigrate in database.go
//2. Create a route & handler
// - Define route in serverSetup
// - Define handler in handlers.go
//3. Create a db operation
// - Define operation in operations.go
