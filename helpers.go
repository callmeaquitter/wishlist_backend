package main 


var sessions = map[string]string{
	"loveyou":         "Axtem",
	"callmeback":      "Asya",
	"cheatcode":       "Misha",
	"totellthetruth":  "Tolya",
	"prostopelmeshki": "Zlata",
}

func getUserID(session string) (string, bool) {
	user, ok := sessions[session] //TODO: change to db operation
	return user, ok
}

func getUser(login, password string) (string, bool) {
	for session, user := range sessions { //TODO: change to db operation
		if user == login {
			return session, true
		}
	}
	return "", false
}

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
