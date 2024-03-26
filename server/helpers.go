package server

import "wishlist/db"

// var sessions = map[string]string{
// 	"loveyou":         "Axtem",
// 	"callmeback":      "Asya",
// 	"cheatcode":       "Misha",
// 	"totellthetruth":  "Tolya",
// 	"prostopelmeshki": "Zlata",
// }



//Step by step guide to add a new feature:
//1. Create a model (if doesn't exist)
// - Add to AutoMigrate in database.go
//2. Create a route & handler
// - Define route in serverSetup
// - Define handler in handlers.go
//3. Create a db operation
// - Define operation in operations.go

//Step by step guide to authenticate a user:
//1. Create a middleware (attach to group of routes)
// - Take session/jwt token from Authorization header
// - Get user from session/jwt token (error if not found)
// - Add user to context (c.Locals)
//2. Use middleware in routes
// - Take user from context (c.Locals)
// - Use user in handler
//3. Create a register & login handlers
// - Register: add user to db, return session/jwt token
// - Login: check user in db, return session/jwt token

func test(wishlists []db.UserWishlist) []string {
	var WishlistID []string
	for _, wishlist := range wishlists {
		WishlistID = append(WishlistID, wishlist.ID)
	}
	return WishlistID
}

// func findGiftToSelectionInGifts(wishes in wishlist) []db.Gift {
// 	var GiftID []string

// }

