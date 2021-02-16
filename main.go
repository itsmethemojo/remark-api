package main

import (
	. "./repositories/bookmark"
	"./routes"
)

func main() {
	//TODO init connection and database migration might be done lazy colling the first route
	bookmarkRepository := BookmarkRepository{}
	bookmarkRepository.InitializeDatabase()
	routes.Run()
}
