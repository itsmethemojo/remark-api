package main

import (
	"./routes"
	bookmarkRepository "./repositories/bookmark"
)

func main() {
	bookmarkRepository.InitializeDatabase()
	routes.Run()
}
