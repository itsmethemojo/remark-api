package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.default")
	godotenv.Overload()
	//TODO init connection and database migration might be done lazy calling the first route
	bookmarkRepository := BookmarkRepository{}
	bookmarkRepository.InitializeDatabase()
	RoutesRun()
}
