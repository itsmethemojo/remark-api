package main

import (
	bookmarkRepository "./repositories/bookmark"
	"./routes"
)

func main() {
	//TODO init connection and database migration might be done lazy colling the first route
	bookmarkRepository := bookmarkRepository.BookmarkRepository{}
	bookmarkRepository.InitializeDatabase()
	routes.Run()
}
