package main

import "fmt"

/*
You lack discipline, control, and worst? YOU'RE SLOPPY.
Тебе не хватает дисциплины, контроля. И что хуже всего ты работаешь небрежно.
- Alastor, Hazbin Hotel
*/
func main() {
	//setups (must be sync)
	databaseSetup(fmt.Sprintf("host=%s user=%s password=%s dbname=%s client_encoding=utf8", "", "test", "KhQOjztlHb_V", "testdb"))
	serverSetup()

	//start (must be async)
	go serverStart()

	//wait
	select {}
}
